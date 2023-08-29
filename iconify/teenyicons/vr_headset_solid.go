package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VrHeadsetSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12.467 3.572l.394.985c.478.106.953.227 1.425.362c.423.12.714.507.714.946V9.21A3.79 3.79 0 0 1 11.21 13h-.542a.783.783 0 0 1-.687-.408c-1.071-1.963-3.89-1.963-4.962 0a.783.783 0 0 1-.687.408H3.79A3.79 3.79 0 0 1 0 9.21V5.865c0-.44.291-.825.714-.946c.472-.135.947-.256 1.425-.362l.394-.985A2.5 2.5 0 0 1 4.854 2h5.292a2.5 2.5 0 0 1 2.321 1.572Zm-9.006.37A1.5 1.5 0 0 1 4.854 3h5.292a1.5 1.5 0 0 1 1.393.943l.153.384a24.703 24.703 0 0 0-8.384 0l.153-.384Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}