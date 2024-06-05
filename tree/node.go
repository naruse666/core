// Copyright (c) 2018, Cogent Core. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package tree provides a powerful and extensible tree system,
// centered on the core Node interface.
package tree

//go:generate core generate
//go:generate core generate ./testdata

import (
	"cogentcore.org/core/enums"
	"cogentcore.org/core/types"
)

// Node is an interface that describes the core functionality of a tree node.
// Each Node is a node in a tree and can have child nodes, and no cycles
// are allowed (i.e., each node can only appear once in the tree).
// All the usual methods are included for accessing and managing Children,
// and efficiently traversing the tree and calling functions on the nodes.
//
// When adding a new node, if you do not specify its name, it will automatically
// be assigned a unique name of the ID (kebab-case) name of the type, plus the
// [Node.NumLifetimeChildren] of the parent. In general, the names of the children
// of a given node should all be unique.
//
// Use the [MoveToParent] function to move a node between trees or within a tree;
// otherwise, nodes are typically created and deleted but not moved.
//
// Most Node functions are only implemented once, by the [tree.NodeBase] type.
// Other Node types extend [tree.NodeBase] and provide their own functionality,
// which can override methods defined by embedded types through a system of virtual
// method calling, as described below.
//
// Each Node stores the Node interface version of itself, as [Node.This],
// which enables full virtual function calling by calling the method
// on that interface instead of directly on the receiver Node itself.
// This allows, for example, a WidgetBase type to call methods defined
// by higher-level Widgets. This requires proper initialization of nodes
// via [Node.InitName], which is called automatically when adding children
// and using [NewRoot].
//
// Nodes support full JSON I/O.
//
// All types that implement the Node interface will automatically
// be added to the Cogent Core type registry (types)
// in `core generate`, which is required for various
// pieces of core functionality.
type Node interface {

	// This returns the Node as its true underlying type.
	// It returns nil if the node is nil, has been destroyed,
	// or is improperly constructed.
	This() Node

	// AsTree returns the [NodeBase] for this Node.
	AsTree() *NodeBase

	// Name returns the user-defined name of the Node, which can be
	// used for finding elements, generating paths, I/O, etc.
	Name() string

	// SetName sets the name of this node. Names should generally be unique
	// across children of each node. If the node requires some non-unique name,
	// add a separate Label field.
	SetName(name string)

	// NodeType returns the [types.Type] record for this Node.
	// This is auto-generated by the typegen generator Node types.
	NodeType() *types.Type

	// New returns a new token of the type of this Node.
	// This new Node must still be initialized.
	// This is auto-generated by the typegen generator for Node types.
	New() Node

	// BaseType returns the base node type for all elements within this tree.
	// This is used in the GUI for determining what types of children can be created.
	BaseType() *types.Type

	// Parents:

	// Parent returns the parent of this Node.
	// Each Node can only have one parent.
	Parent() Node

	// ParentByName finds first parent recursively up hierarchy that matches
	// given name. Returns nil if not found.
	ParentByName(name string) Node

	// ParentByType finds parent recursively up hierarchy, by type, and
	// returns nil if not found. If embeds is true, then it looks for any
	// type that embeds the given type at any level of anonymous embedding.
	ParentByType(t *types.Type, embeds bool) Node

	// Children:

	// HasChildren returns whether this node has any children.
	HasChildren() bool

	// NumChildren returns the number of children this node has.
	NumChildren() int

	// NumLifetimeChildren returns the number of children that this node
	// has ever had added to it (it is not decremented when a child is removed).
	// It is used for unique naming of children.
	NumLifetimeChildren() uint64

	// Children returns a pointer to the slice of children of this node.
	// The resultant slice can be modified directly (e.g., sort, reorder),
	// but new children should be added via New/Add/Insert Child methods on
	// Node to ensure proper initialization.
	Children() *Slice

	// Child returns the child of this node at the given index and returns nil if
	// the index is out of range.
	Child(i int) Node

	// ChildByName returns the first child that has the given name, and nil
	// if no such element is found. startIndex arg allows for optimized
	// bidirectional find if you have an idea where it might be, which
	// can be a key speedup for large lists. If no value is specified for
	// startIndex, it starts in the middle, which is a good default.
	ChildByName(name string, startIndex ...int) Node

	// ChildByType returns the first child that has the given type, and nil
	// if not found. If embeds is true, then it also looks for any type that
	// embeds the given type at any level of anonymous embedding.
	// startIndex arg allows for optimized bidirectional find if you have an
	// idea where it might be, which can be a key speedup for large lists. If
	// no value is specified for startIndex, it starts in the middle, which is a
	// good default.
	ChildByType(t *types.Type, embeds bool, startIndex ...int) Node

	// Paths:

	// Path returns the path to this node from the tree root,
	// using [Node.Name]s separated by / and fields by .
	// Path is only valid for finding items when child names
	// are unique. Any existing / and . characters in names
	// are escaped to \\ and \,
	Path() string

	// PathFrom returns path to this node from the given parent node, using
	// [Node.Name]s separated by / and fields by .
	// Path is only valid for finding items when child names
	// are unique. Any existing / and . characters in names
	// are escaped to \\ and \,
	//
	// The paths that it returns exclude the
	// name of the parent and the leading slash; for example, in the tree
	// a/b/c/d/e, the result of d.PathFrom(b) would be c/d. PathFrom
	// automatically gets the [Node.This] version of the given parent,
	// so a base type can be passed in without manually calling [Node.This].
	PathFrom(parent Node) string

	// FindPath returns the node at the given path from this node.
	// FindPath only works correctly when names are unique.
	// Path has [Node.Name]s separated by / and fields by .
	// Node names escape any existing / and . characters to \\ and \,
	// There is also support for [idx] index-based access for any given path
	// element, for cases when indexes are more useful than names.
	// Returns nil if not found.
	FindPath(path string) Node

	// FieldByName returns the node that is a direct field with the given name.
	// This must be implemented for any types that have Node fields that
	// are processed as part of the overall Node tree. This is only used
	// by [Node.FindPath]. Returns error if not found.
	FieldByName(field string) (Node, error)

	// Adding and Inserting Children:

	// AddChild adds given child at end of children list.
	// The kid node is assumed to not be on another tree (see [MoveToParent])
	// and the existing name should be unique among children.
	// Any error is automatically logged in addition to being returned.
	AddChild(kid Node) error

	// NewChild creates a new child of the given type and adds it at the end
	// of the list of children. The name defaults to the ID (kebab-case) name
	// of the type, plus the [Node.NumLifetimeChildren] of the parent.
	NewChild(typ *types.Type) Node

	// SetChild sets the child at the given index to be the given item.
	// It just calls Init and SetParent on the child. The name defaults
	// to the ID (kebab-case) name of the type, plus the
	// [Node.NumLifetimeChildren] of the parent.
	// Any error is automatically logged in addition to being returned.
	SetChild(kid Node, idx int) error

	// InsertChild adds given child at position in children list.
	// The kid node is assumed to not be on another tree (see [MoveToParent])
	// and the existing name should be unique among children.
	// Any error is automatically logged in addition to being returned.
	InsertChild(kid Node, at int) error

	// InsertNewChild creates a new child of given type and add at position
	// in children list. The name defaults to the ID (kebab-case) name
	// of the type, plus the [Node.NumLifetimeChildren] of the parent.
	InsertNewChild(typ *types.Type, at int) Node

	// Deleting Children:

	// DeleteChildAtIndex deletes child at given index. It returns false
	// if there is no child at the given index.
	DeleteChildAtIndex(idx int) bool

	// DeleteChild deletes the given child node, returning false if
	// it can not find it.
	DeleteChild(child Node) bool

	// DeleteChildByName deletes child node by name, returning false
	// if it can not find it.
	DeleteChildByName(name string) bool

	// DeleteChildren deletes all children nodes.
	DeleteChildren()

	// Delete deletes this node from its parent's children list.
	Delete()

	// Destroy recursively deletes and destroys all children and
	// their children's children, etc.
	Destroy()

	// Flags:

	// Is checks if the given flag is set, using atomic,
	// which is safe for concurrent access.
	Is(f enums.BitFlag) bool

	// SetFlag sets the given flag(s) to the given state
	// using atomic, which is safe for concurrent access.
	SetFlag(on bool, f ...enums.BitFlag)

	// FlagType returns the flags of the node as the true flag type of the node,
	// which may be a type that extends the standard [Flags]. Each node type
	// that extends the flag type should define this method; for example:
	//	func (wb *WidgetBase) FlagType() enums.BitFlagSetter {
	//		return (*WidgetFlags)(&wb.Flags)
	//	}
	FlagType() enums.BitFlagSetter

	// Tree Walking:

	// WalkUp calls the given function on the node and all of its parents,
	// sequentially in the current goroutine (generally necessary for going up,
	// which is typically quite fast anyway). It stops walking if the function
	// returns [Break] and keeps walking if it returns [Continue]. It returns
	// whether walking was finished (false if it was aborted with [Break]).
	WalkUp(fun func(n Node) bool) bool

	// WalkUpParent calls the given function on all of the node's parents (but not
	// the node itself), sequentially in the current goroutine (generally necessary
	// for going up, which is typically quite fast anyway). It stops walking if the
	// function returns [Break] and keeps walking if it returns [Continue]. It returns
	// whether walking was finished (false if it was aborted with [Break]).
	WalkUpParent(fun func(n Node) bool) bool

	// WalkDown calls the given function on the node and all of its children
	// in a depth-first manner over all of the children, sequentially in the
	// current goroutine. It stops walking the current branch of the tree if
	// the function returns [Break] and keeps walking if it returns [Continue].
	// It is non-recursive and safe for concurrent calling. The [Node.NodeWalkDown]
	// method is called for every node after the given function, which enables nodes
	// to also traverse additional nodes, like widget parts.
	WalkDown(fun func(n Node) bool)

	// NodeWalkDown is a method that nodes can implement to traverse additional nodes
	// like widget parts during [Node.WalkDown]. It is called with the function passed
	// to [Node.WalkDown] after the function is called with the node itself.
	NodeWalkDown(fun func(n Node) bool)

	// WalkDownPost iterates in a depth-first manner over the children, calling
	// doChildTest on each node to test if processing should proceed (if it returns
	// [Break] then that branch of the tree is not further processed),
	// and then calls the given function after all of a node's children
	// have been iterated over. In effect, this means that the given function
	// is called for deeper nodes first. This uses node state information to manage
	// the traversal and is very fast, but can only be called by one goroutine at a
	// time, so you should use a Mutex if there is a chance of multiple threads
	// running at the same time. The nodes are processed in the current goroutine.
	WalkDownPost(doChildTest func(n Node) bool, fun func(n Node) bool)

	// WalkDownBreadth calls the given function on the node and all of its children
	// in breadth-first order. It stops walking the current branch of the tree if the
	// function returns [Break] and keeps walking if it returns [Continue]. It is
	// non-recursive, but not safe for concurrent calling.
	WalkDownBreadth(fun func(n Node) bool)

	// Deep Copy:

	// CopyFrom copies the data and children of the given node to this node.
	// It is essential that the source node has unique names. It is very efficient
	// by using the [Node.ConfigChildren] method which attempts to preserve any
	// existing nodes in the destination if they have the same name and type, so a
	// copy from a source to a target that only differ minimally will be
	// minimally destructive. Only copying to the same type is supported.
	// The struct field tag copier:"-" can be added for any fields that
	// should not be copied. Also, unexported fields are not copied.
	// See [Node.CopyFieldsFrom] for more information on field copying.
	CopyFrom(src Node)

	// Clone creates and returns a deep copy of the tree from this node down.
	// Any pointers within the cloned tree will correctly point within the new
	// cloned tree (see [Node.CopyFrom] for more information).
	Clone() Node

	// CopyFieldsFrom copies the fields of the node from the given node.
	// By default, it is [NodeBase.CopyFieldsFrom], which automatically does
	// a deep copy of all of the fields of the node that do not a have a
	// `copier:"-"` struct tag. Node types should only implement a custom
	// CopyFieldsFrom method when they have fields that need special copying
	// logic that can not be automatically handled. All custom CopyFieldsFrom
	// methods should call [NodeBase.CopyFieldsFrom] first and then only do manual
	// handling of specific fields that can not be automatically copied. See
	// [cogentcore.org/core/core.WidgetBase.CopyFieldsFrom] for an example of a
	// custom CopyFieldsFrom method.
	CopyFieldsFrom(from Node)

	// Event methods:

	// Init is called when the node is
	// initialized (ie: through [Node.InitName]).
	// It is called before the node is added to the tree,
	// so it will not have any parents or siblings.
	// It will be called only once in the lifetime of the node.
	// It does nothing by default, but it can be implemented
	// by higher-level types that want to do something.
	Init()

	// OnAdd is called when the node is added to a parent.
	// It will be called only once in the lifetime of the node,
	// unless the node is moved. It will not be called on root
	// nodes, as they are never added to a parent.
	// It does nothing by default, but it can be implemented
	// by higher-level types that want to do something.
	OnAdd()

	// OnChildAdded is called when a node is added to
	// this node or any of its children. When a node is added to
	// a tree, it calls [OnAdd] and then this function on each of its parents,
	// going in order from the closest parent to the furthest parent.
	// This function does nothing by default, but it can be
	// implemented by higher-level types that want to do something.
	OnChildAdded(child Node)
}
