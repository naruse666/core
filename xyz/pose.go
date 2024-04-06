// Copyright (c) 2019, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package xyz

import (
	"cogentcore.org/core/mat32"
)

// Pose contains the full specification of position and orientation,
// always relevant to the parent element.
type Pose struct {

	// position of center of element (relative to parent)
	Pos mat32.Vec3

	// scale (relative to parent)
	Scale mat32.Vec3

	// Node rotation specified as a Quat (relative to parent)
	Quat mat32.Quat

	// Local matrix. Contains all position/rotation/scale information (relative to parent)
	Matrix mat32.Mat4 `view:"-"`

	// Parent's world matrix -- we cache this so that we can independently update our own matrix
	ParMatrix mat32.Mat4 `view:"-"`

	// World matrix. Contains all absolute position/rotation/scale information (i.e. relative to very top parent, generally the scene)
	WorldMatrix mat32.Mat4 `view:"-"`

	// model * view matrix -- tranforms into camera-centered coords
	MVMatrix mat32.Mat4 `view:"-"`

	// model * view * projection matrix -- full final render matrix
	MVPMatrix mat32.Mat4 `view:"-"`

	// normal matrix has no offsets, for normal vector rotation only, based on MVMatrix
	NormMatrix mat32.Mat3 `view:"-"`
}

// Defaults sets defaults only if current values are nil
func (ps *Pose) Defaults() {
	if ps.Scale == (mat32.Vec3{}) {
		ps.Scale.Set(1, 1, 1)
	}
	if ps.Quat.IsNil() {
		ps.Quat.SetIdentity()
	}
}

// CopyFrom copies just the pose information from the other pose, critically
// not copying the ParMatrix so that is preserved in the receiver.
func (ps *Pose) CopyFrom(op *Pose) {
	ps.Pos = op.Pos
	ps.Scale = op.Scale
	ps.Quat = op.Quat
	ps.UpdateMatrix()
}

func (ps *Pose) String() string {
	return "Pos: " + ps.Pos.String() + "; Scale: " + ps.Scale.String() + "; Quat: " + ps.Quat.String()
}

// GenGoSet returns code to set values at given path (var.member etc)
func (ps *Pose) GenGoSet(path string) string {
	return ps.Pos.GenGoSet(path+".Pos") + "; " + ps.Scale.GenGoSet(path+".Scale") + "; " + ps.Quat.GenGoSet(path+".Quat")
}

// UpdateMatrix updates the local transform matrix based on its position, quaternion, and scale.
// Also checks for degenerate nil values
func (ps *Pose) UpdateMatrix() {
	ps.Defaults()
	ps.Matrix.SetTransform(ps.Pos, ps.Quat, ps.Scale)
}

// MulMatrix multiplies current pose Matrix by given Matrix, and re-extracts the
// Pos, Scale, Quat from resulting matrix.
func (ps *Pose) MulMatrix(mat *mat32.Mat4) {
	ps.Matrix.SetMul(mat)
	pos, quat, sc := ps.Matrix.Decompose()
	ps.Pos = pos
	ps.Quat = quat
	ps.Scale = sc
}

// UpdateWorldMatrix updates the world transform matrix based on Matrix and parent's WorldMatrix.
// Does NOT call UpdateMatrix so that can include other factors as needed.
func (ps *Pose) UpdateWorldMatrix(parWorld *mat32.Mat4) {
	if parWorld != nil {
		ps.ParMatrix.CopyFrom(parWorld)
	}
	ps.WorldMatrix.MulMatrices(&ps.ParMatrix, &ps.Matrix)
}

// UpdateMVPMatrix updates the model * view, * projection matricies based on camera view, prjn matricies
// Assumes that WorldMatrix has been updated
func (ps *Pose) UpdateMVPMatrix(viewMat, prjnMat *mat32.Mat4) {
	ps.MVMatrix.MulMatrices(viewMat, &ps.WorldMatrix)
	ps.NormMatrix.SetNormalMatrix(&ps.MVMatrix)
	ps.MVPMatrix.MulMatrices(prjnMat, &ps.MVMatrix)
}

///////////////////////////////////////////////////////
// 		Moving

// Note: you can just directly add to .Pos too..

// MoveOnAxis moves (translates) the specified distance on the specified local axis,
// relative to the current rotation orientation.
func (ps *Pose) MoveOnAxis(x, y, z, dist float32) {
	ps.Pos.SetAdd(mat32.V3(x, y, z).Normal().MulQuat(ps.Quat).MulScalar(dist))
}

// MoveOnAxisAbs moves (translates) the specified distance on the specified local axis,
// in absolute X,Y,Z coordinates.
func (ps *Pose) MoveOnAxisAbs(x, y, z, dist float32) {
	ps.Pos.SetAdd(mat32.V3(x, y, z).Normal().MulScalar(dist))
}

///////////////////////////////////////////////////////
// 		Rotating

// SetEulerRotation sets the rotation in Euler angles (degrees).
func (ps *Pose) SetEulerRotation(x, y, z float32) {
	ps.Quat.SetFromEuler(mat32.V3(x, y, z).MulScalar(mat32.DegToRadFactor))
}

// SetEulerRotationRad sets the rotation in Euler angles (radians).
func (ps *Pose) SetEulerRotationRad(x, y, z float32) {
	ps.Quat.SetFromEuler(mat32.V3(x, y, z))
}

// EulerRotation returns the current rotation in Euler angles (degrees).
func (ps *Pose) EulerRotation() mat32.Vec3 {
	return ps.Quat.ToEuler().MulScalar(mat32.RadToDegFactor)
}

// EulerRotationRad returns the current rotation in Euler angles (radians).
func (ps *Pose) EulerRotationRad() mat32.Vec3 {
	return ps.Quat.ToEuler()
}

// SetAxisRotation sets rotation from local axis and angle in degrees.
func (ps *Pose) SetAxisRotation(x, y, z, angle float32) {
	ps.Quat.SetFromAxisAngle(mat32.V3(x, y, z), mat32.DegToRad(angle))
}

// SetAxisRotationRad sets rotation from local axis and angle in radians.
func (ps *Pose) SetAxisRotationRad(x, y, z, angle float32) {
	ps.Quat.SetFromAxisAngle(mat32.V3(x, y, z), angle)
}

// RotateOnAxis rotates around the specified local axis the specified angle in degrees.
func (ps *Pose) RotateOnAxis(x, y, z, angle float32) {
	ps.Quat.SetMul(mat32.NewQuatAxisAngle(mat32.V3(x, y, z), mat32.DegToRad(angle)))
}

// RotateOnAxisRad rotates around the specified local axis the specified angle in radians.
func (ps *Pose) RotateOnAxisRad(x, y, z, angle float32) {
	ps.Quat.SetMul(mat32.NewQuatAxisAngle(mat32.V3(x, y, z), angle))
}

// RotateEuler rotates by given Euler angles (in degrees) relative to existing rotation.
func (ps *Pose) RotateEuler(x, y, z float32) {
	ps.Quat.SetMul(mat32.NewQuatEuler(mat32.V3(x, y, z).MulScalar(mat32.DegToRadFactor)))
}

// RotateEulerRad rotates by given Euler angles (in radians) relative to existing rotation.
func (ps *Pose) RotateEulerRad(x, y, z, angle float32) {
	ps.Quat.SetMul(mat32.NewQuatEuler(mat32.V3(x, y, z)))
}

// SetMatrix sets the local transformation matrix and updates Pos, Scale, Quat.
func (ps *Pose) SetMatrix(m *mat32.Mat4) {
	ps.Matrix = *m
	ps.Pos, ps.Quat, ps.Scale = ps.Matrix.Decompose()
}

// LookAt points the element at given target location using given up direction.
func (ps *Pose) LookAt(target, upDir mat32.Vec3) {
	ps.Quat.SetFromRotationMatrix(mat32.NewLookAt(ps.Pos, target, upDir))
}

///////////////////////////////////////////////////////
// 		World values

// WorldPos returns the current world position.
func (ps *Pose) WorldPos() mat32.Vec3 {
	pos := mat32.Vec3{}
	pos.SetFromMatrixPos(&ps.WorldMatrix)
	return pos
}

// WorldQuat returns the current world quaternion.
func (ps *Pose) WorldQuat() mat32.Quat {
	_, quat, _ := ps.WorldMatrix.Decompose()
	return quat
}

// WorldEulerRotation returns the current world rotation in Euler angles.
func (ps *Pose) WorldEulerRotation() mat32.Vec3 {
	return ps.Quat.ToEuler()
}

// WorldScale returns he current world scale.
func (ps *Pose) WorldScale() mat32.Vec3 {
	_, _, scale := ps.WorldMatrix.Decompose()
	return scale
}

/*

// PoseProps define the Toolbar and MenuBar for StructView
var PoseProps = tree.Props{
	"Toolbar": tree.PropSlice{
		{"GenGoSet", tree.Props{
			"label":       "Go Code",
			"desc":        "returns Go Code that sets the current Pose, based on given path to Pose.",
			"icon":        icons.Code,
			"show-return": true,
			"Args": tree.PropSlice{
				{"Path", tree.BlankProp{}},
			},
		}},
		{"SetEulerRotation", tree.Props{
			"desc": "Set the local rotation (relative to parent) using Euler angles, in degrees.",
			"icon": icons.X3DRotation,
			"Args": tree.PropSlice{
				{"Pitch", tree.Props{
					"desc": "rotation up / down along the X axis (in the Y-Z plane), e.g., the altitude (climbing, descending) for motion along the Z depth axis",
				}},
				{"Yaw", tree.Props{
					"desc": "rotation along the Y axis (in the horizontal X-Z plane), e.g., the bearing or direction for motion along the Z depth axis",
				}},
				{"Roll", tree.Props{
					"desc": "rotation along the Z axis (in the X-Y plane), e.g., the bank angle for motion along the Z depth axis",
				}},
			},
		}},
		{"SetAxisRotation", tree.Props{
			"desc": "Set the local rotation (relative to parent) using Axis about which to rotate, and the angle.",
			"icon": icons.X3DRotation,
			"Args": tree.PropSlice{
				{"X", tree.BlankProp{}},
				{"Y", tree.BlankProp{}},
				{"Z", tree.BlankProp{}},
				{"Angle", tree.BlankProp{}},
			},
		}},
		{"RotateEuler", tree.Props{
			"desc": "rotate (relative to current rotation) using Euler angles, in degrees.",
			"icon": icons.X3DRotation,
			"Args": tree.PropSlice{
				{"Pitch", tree.Props{
					"desc": "rotation up / down along the X axis (in the Y-Z plane), e.g., the altitude (climbing, descending) for motion along the Z depth axis",
				}},
				{"Yaw", tree.Props{
					"desc": "rotation along the Y axis (in the horizontal X-Z plane), e.g., the bearing or direction for motion along the Z depth axis",
				}},
				{"Roll", tree.Props{
					"desc": "rotation along the Z axis (in the X-Y plane), e.g., the bank angle for motion along the Z depth axis",
				}},
			},
		}},
		{"RotateOnAxis", tree.Props{
			"desc": "Rotate (relative to current rotation) using Axis about which to rotate, and the angle.",
			"icon": icons.X3DRotation,
			"Args": tree.PropSlice{
				{"X", tree.BlankProp{}},
				{"Y", tree.BlankProp{}},
				{"Z", tree.BlankProp{}},
				{"Angle", tree.BlankProp{}},
			},
		}},
		{"LookAt", tree.Props{
			"icon": icons.X3DRotation,
			"Args": tree.PropSlice{
				{"Target", tree.BlankProp{}},
				{"UpDir", tree.BlankProp{}},
			},
		}},
		{"EulerRotation", tree.Props{
			"desc":        "The local rotation (relative to parent) in Euler angles in degrees (X = Pitch, Y = Yaw, Z = Roll)",
			"icon":        icons.X3DRotation,
			"show-return": "true",
		}},
		{"sep-rot", tree.BlankProp{}},
		{"MoveOnAxis", tree.Props{
			"desc": "Move given distance on given X,Y,Z axis relative to current rotation orientation.",
			"icon": icons.PanTool,
			"Args": tree.PropSlice{
				{"X", tree.BlankProp{}},
				{"Y", tree.BlankProp{}},
				{"Z", tree.BlankProp{}},
				{"Dist", tree.BlankProp{}},
			},
		}},
		{"MoveOnAxisAbs", tree.Props{
			"desc": "Move given distance on given X,Y,Z axis in absolute coords, not relative to current rotation orientation.",
			"icon": icons.PanTool,
			"Args": tree.PropSlice{
				{"X", tree.BlankProp{}},
				{"Y", tree.BlankProp{}},
				{"Z", tree.BlankProp{}},
				{"Dist", tree.BlankProp{}},
			},
		}},
	},
}

*/
