package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Stapler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M640 299.3V432c0 26.5-21.5 48-48 48H64c-17.7 0-32-14.3-32-32s14.3-32 32-32h384v-48H96c-17.7 0-32-14.3-32-32V219.4L33.8 214C14.2 210.5 0 193.5 0 173.7c0-8.9 2.9-17.5 8.2-24.6l35.6-47.5C76.7 57.8 128.2 32 182.9 32c27 0 53.6 6.3 77.8 18.4l326.2 163.1c32.6 16.2 53.1 49.5 53.1 85.8zM448 304v-16l-320-57.1V304h320z"/>`),
		g.Group(children),
	)
}