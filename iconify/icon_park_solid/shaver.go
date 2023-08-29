package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shaver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSShaver0"><g fill="none" stroke-width="4"><path fill="#fff" stroke="#fff" d="M20.785 18.73a.719.719 0 0 1 1.016 0l7.469 7.469a.72.72 0 0 1 0 1.016L17.757 38.728a6 6 0 0 1-8.485-8.486L20.785 18.73Z"/><path stroke="#fff" d="M42.581 22.389a.894.894 0 0 0 0-1.264L26.874 5.418a.894.894 0 0 0-1.263 0l-.783.783c-4.686 4.686-4.686 12.284 0 16.97c4.687 4.687 12.285 4.687 16.97 0l.783-.782Z"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" d="m19.879 28.121l-1.415 1.414"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" d="m27.657 17.515l2.828 2.828"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSShaver0)"/>`),
		g.Group(children),
	)
}