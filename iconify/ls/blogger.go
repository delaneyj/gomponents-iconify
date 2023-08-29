package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Blogger(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M326 306h-73c-12 0-23-11-23-24c0-12 11-23 23-23h73c13 0 23 11 23 23c0 13-10 24-23 24zm-73 98h145c13 0 23 10 23 23s-10 23-23 23H253c-12 0-23-10-23-23s11-23 23-23zM98 29h454c54 0 98 44 98 98v454c0 54-44 98-98 98H98c-54 0-98-44-98-98V127c0-54 44-98 98-98zm428 403v-96c0-18-5-30-27-30h-17c-18 0-28-11-28-24v-11c0-57-62-117-116-117h-85c-71 0-127 56-127 127v152c0 67 58 121 117 121h161c65 0 122-55 122-122z"/>`),
		g.Group(children),
	)
}