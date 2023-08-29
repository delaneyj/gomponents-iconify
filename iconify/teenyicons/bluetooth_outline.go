package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" d="M7.5 14.5H7a.5.5 0 0 0 .787.41L7.5 14.5Zm0-14l.287-.41A.5.5 0 0 0 7 .5h.5Zm5 3.5l.282.413a.5.5 0 0 0 .005-.823L12.5 4Zm0 7l.287.41a.5.5 0 0 0-.005-.823L12.5 11ZM8 14.5V7.41H7v7.09h1Zm0-7.09V.5H7v6.91h1ZM7.213.91l5 3.5l.574-.82l-5-3.5l-.574.82Zm5.005 2.677l-5 3.409l.564.826l5-3.409l-.564-.826Zm-5 3.409l-6 4.09l.564.827l6-4.09l-.564-.827Zm.569 7.914l5-3.5l-.574-.82l-5 3.5l.574.82Zm4.995-4.323l-11-7.5l-.564.826l11 7.5l.564-.826Z"/>`),
		g.Group(children),
	)
}