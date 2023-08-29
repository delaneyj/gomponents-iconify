package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hngry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 27.277l-4.729-3.608a1.157 1.157 0 0 0-1.419.012l-3.688 2.914a1.157 1.157 0 0 1-1.588-.146l-1.61-1.839a.579.579 0 0 0-1.007.467l2.61 17.437a1.157 1.157 0 0 0 1.145.986H24m0-16.223l4.729-3.608a1.157 1.157 0 0 1 1.419.012l3.688 2.914a1.157 1.157 0 0 0 1.588-.146l1.61-1.839a.579.579 0 0 1 1.007.467l-2.61 17.437a1.157 1.157 0 0 1-1.145.986H24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.392 26.239V9.432A4.932 4.932 0 0 1 17.324 4.5h0a4.932 4.932 0 0 1 4.932 4.932v16.514m-4.932-13.659l4.932-2.531m-4.932 8.079l4.932-2.53M35.8 26.02c.798-3.22-.863-4.713 0-9.45l-1.279-3.115l-2.077 1.363l-.844-2.596l-2.79.973c-.52 3.44-4.721 4.77-4.81 9.215a7.786 7.786 0 0 0 .955 4.138"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.81 13.195c.52-2.595 2.66-6.489 6.619-4.282s.37 7.657.37 7.657"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.15 6.577a2.562 2.562 0 0 0-1.393 2.011"/>`),
		g.Group(children),
	)
}