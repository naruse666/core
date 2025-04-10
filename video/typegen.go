// Code generated by "core generate"; DO NOT EDIT.

package video

import (
	"github.com/naruse666/core/tree"
	"github.com/naruse666/core/types"
	"github.com/cogentcore/reisen"
)

var _ = types.AddType(&types.Type{Name: "github.com/naruse666/core/video.Video", IDName: "video", Doc: "Video represents a video playback widget without any controls.\nSee [Player] for a version with controls.", Directives: []types.Directive{{Tool: "types", Directive: "add"}}, Embeds: []types.Field{{Name: "WidgetBase"}}, Fields: []types.Field{{Name: "Media", Doc: "Media is the video media."}, {Name: "Rotation", Doc: "degrees of rotation to apply to the video images\n90 = left 90, -90 = right 90"}, {Name: "Stop", Doc: "setting this to true will stop the playing"}, {Name: "frameBuffer"}, {Name: "lastFrame", Doc: "last frame we have rendered"}, {Name: "frameTarg", Doc: "target frame number to be played"}, {Name: "framePlayed", Doc: "actual frame number displayed"}, {Name: "frameStop", Doc: "frame number to stop playing at, if > 0"}}})

// NewVideo returns a new [Video] with the given optional parent:
// Video represents a video playback widget without any controls.
// See [Player] for a version with controls.
func NewVideo(parent ...tree.Node) *Video { return tree.New[Video](parent...) }

// SetMedia sets the [Video.Media]:
// Media is the video media.
func (t *Video) SetMedia(v *reisen.Media) *Video { t.Media = v; return t }

// SetRotation sets the [Video.Rotation]:
// degrees of rotation to apply to the video images
// 90 = left 90, -90 = right 90
func (t *Video) SetRotation(v float32) *Video { t.Rotation = v; return t }

// SetStop sets the [Video.Stop]:
// setting this to true will stop the playing
func (t *Video) SetStop(v bool) *Video { t.Stop = v; return t }
