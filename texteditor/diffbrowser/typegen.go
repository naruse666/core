// Code generated by "core generate"; DO NOT EDIT.

package diffbrowser

import (
	"github.com/naruse666/core/base/fileinfo"
	"github.com/naruse666/core/tree"
	"github.com/naruse666/core/types"
)

var _ = types.AddType(&types.Type{Name: "github.com/naruse666/core/texteditor/diffbrowser.Browser", IDName: "browser", Doc: "Browser is a diff browser, for browsing a set of paired files\nfor viewing differences between them, organized into a tree\nstructure, e.g., reflecting their source in a filesystem.", Methods: []types.Method{{Name: "OpenFiles", Doc: "OpenFiles Updates the tree based on files", Directives: []types.Directive{{Tool: "types", Directive: "add"}}}}, Embeds: []types.Field{{Name: "Frame"}}, Fields: []types.Field{{Name: "PathA", Doc: "starting paths for the files being compared"}, {Name: "PathB", Doc: "starting paths for the files being compared"}}})

// NewBrowser returns a new [Browser] with the given optional parent:
// Browser is a diff browser, for browsing a set of paired files
// for viewing differences between them, organized into a tree
// structure, e.g., reflecting their source in a filesystem.
func NewBrowser(parent ...tree.Node) *Browser { return tree.New[Browser](parent...) }

// SetPathA sets the [Browser.PathA]:
// starting paths for the files being compared
func (t *Browser) SetPathA(v string) *Browser { t.PathA = v; return t }

// SetPathB sets the [Browser.PathB]:
// starting paths for the files being compared
func (t *Browser) SetPathB(v string) *Browser { t.PathB = v; return t }

var _ = types.AddType(&types.Type{Name: "github.com/naruse666/core/texteditor/diffbrowser.Node", IDName: "node", Doc: "Node is an element in the diff tree", Embeds: []types.Field{{Name: "Tree"}}, Fields: []types.Field{{Name: "FileA", Doc: "file names (full path) being compared. Name of node is just the filename.\nTypically A is the older, base version and B is the newer one being compared."}, {Name: "FileB", Doc: "file names (full path) being compared. Name of node is just the filename.\nTypically A is the older, base version and B is the newer one being compared."}, {Name: "RevA", Doc: "VCS revisions for files if applicable"}, {Name: "RevB", Doc: "VCS revisions for files if applicable"}, {Name: "Status", Doc: "Status of the change from A to B: A=Added, D=Deleted, M=Modified, R=Renamed"}, {Name: "TextA", Doc: "Text content of the files"}, {Name: "TextB", Doc: "Text content of the files"}, {Name: "Info", Doc: "Info about the B file, for getting icons etc"}}})

// NewNode returns a new [Node] with the given optional parent:
// Node is an element in the diff tree
func NewNode(parent ...tree.Node) *Node { return tree.New[Node](parent...) }

// SetFileA sets the [Node.FileA]:
// file names (full path) being compared. Name of node is just the filename.
// Typically A is the older, base version and B is the newer one being compared.
func (t *Node) SetFileA(v string) *Node { t.FileA = v; return t }

// SetFileB sets the [Node.FileB]:
// file names (full path) being compared. Name of node is just the filename.
// Typically A is the older, base version and B is the newer one being compared.
func (t *Node) SetFileB(v string) *Node { t.FileB = v; return t }

// SetRevA sets the [Node.RevA]:
// VCS revisions for files if applicable
func (t *Node) SetRevA(v string) *Node { t.RevA = v; return t }

// SetRevB sets the [Node.RevB]:
// VCS revisions for files if applicable
func (t *Node) SetRevB(v string) *Node { t.RevB = v; return t }

// SetStatus sets the [Node.Status]:
// Status of the change from A to B: A=Added, D=Deleted, M=Modified, R=Renamed
func (t *Node) SetStatus(v string) *Node { t.Status = v; return t }

// SetTextA sets the [Node.TextA]:
// Text content of the files
func (t *Node) SetTextA(v string) *Node { t.TextA = v; return t }

// SetTextB sets the [Node.TextB]:
// Text content of the files
func (t *Node) SetTextB(v string) *Node { t.TextB = v; return t }

// SetInfo sets the [Node.Info]:
// Info about the B file, for getting icons etc
func (t *Node) SetInfo(v fileinfo.FileInfo) *Node { t.Info = v; return t }
