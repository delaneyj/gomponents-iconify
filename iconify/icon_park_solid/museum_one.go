package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MuseumOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSMuseumOne0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" stroke-linejoin="round" d="M4 44h40"/><path fill="#fff" stroke-linejoin="round" d="M8 8.364L24 4l16 4.364V14H8V8.364Z"/><path stroke-linecap="round" d="M10 14v24m7-24v24m7-24v24m7-24v24m7-24v24"/><path stroke-linejoin="round" d="M7 38h34v6H7z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSMuseumOne0)"/>`),
		g.Group(children),
	)
}