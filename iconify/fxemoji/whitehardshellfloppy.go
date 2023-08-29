package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whitehardshellfloppy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#CFD3D3" d="m65.46 487.524l-46.782-70.892a18.65 18.65 0 0 1-3.084-10.271V49.188c0-19.882 16.118-36 36-36h410.812c19.882 0 36 16.118 36 36V460c0 19.882-16.118 36-36 36H81.208a18.867 18.867 0 0 1-15.748-8.476z"/><path fill="#EDEDED" d="M89.909 13.188h334.182v214.201c0 10.837-8.785 19.623-19.623 19.623H109.532c-10.837 0-19.623-8.785-19.623-19.623V13.188z"/><path fill="#597B91" d="M380.013 496H141.981V358.73c0-7.92 6.48-14.4 14.4-14.4h209.232c7.92 0 14.4 6.48 14.4 14.4V496z"/><path fill="#EDEDED" d="M243.99 488.556h-82c-6.6 0-12-5.4-12-12V365.012c0-6.6 5.4-12 12-12h82c6.6 0 12 5.4 12 12v111.544c0 6.6-5.4 12-12 12z"/>`),
		g.Group(children),
	)
}