// Code generated by "core generate"; DO NOT EDIT.

package tree

import (
	"cogentcore.org/core/gti"
)

// NodeBaseType is the [gti.Type] for [NodeBase]
var NodeBaseType = gti.AddType(&gti.Type{Name: "cogentcore.org/core/tree.NodeBaseBase", IDName: "node-base", Doc: "The NodeBase struct implements the [Node] interface and provides the core functionality\nfor the Cogent Core tree system. You can use the NodeBase as an embedded struct or as a struct\nfield; the embedded version supports full JSON saving and loading. All types that\nimplement the [Node] interface will automatically be added to gti in `core generate`, which\nis required for various pieces of core functionality.", Fields: []gti.Field{{Name: "Nm", Doc: "Nm is the user-supplied name of this node, which can be empty and/or non-unique."}, {Name: "Flags", Doc: "Flags are bit flags for internal node state, which can be extended using the enums package."}, {Name: "Props", Doc: "Props is a property map for arbitrary extensible properties."}, {Name: "Par", Doc: "Par is the parent of this node, which is set automatically when this node is added as a child of a parent."}, {Name: "Kids", Doc: "Kids is the list of children of this node. All of them are set to have this node\nas their parent. They can be reordered, but you should generally use Ki Node methods\nto Add / Delete to ensure proper usage."}, {Name: "Ths", Doc: "Ths is a pointer to ourselves as a Ki. It can always be used to extract the true underlying type\nof an object when [Node] is embedded in other structs; function receivers do not have this ability\nso this is necessary. This is set to nil when deleted. Typically use [Ki.This] convenience accessor\nwhich protects against concurrent access."}, {Name: "NumLifetimeKids", Doc: "NumLifetimeKids is the number of children that have ever been added to this node, which is used for automatic unique naming."}, {Name: "index", Doc: "index is the last value of our index, which is used as a starting point for finding us in our parent next time.\nIt is not guaranteed to be accurate; use the [Ki.IndexInParent] method."}, {Name: "depth", Doc: "depth is an optional depth parameter of this node, which is only valid during specific contexts, not generally.\nFor example, it is used in the WalkBreadth function"}}, Instance: &NodeBase{}})

// NewNodeBase adds a new [NodeBase] with the given name to the given parent:
// The NodeBase struct implements the [Node] interface and provides the core functionality
// for the Cogent Core tree system. You can use the NodeBase as an embedded struct or as a struct
// field; the embedded version supports full JSON saving and loading. All types that
// implement the [Node] interface will automatically be added to gti in `core generate`, which
// is required for various pieces of core functionality.
func NewNodeBase(parent Node, name ...string) *NodeBase {
	return parent.NewChild(NodeBaseType, name...).(*NodeBase)
}

// KiType returns the [*gti.Type] of [NodeBase]
func (t *NodeBase) KiType() *gti.Type { return NodeBaseType }

// New returns a new [*NodeBase] value
func (t *NodeBase) New() Node { return &NodeBase{} }
