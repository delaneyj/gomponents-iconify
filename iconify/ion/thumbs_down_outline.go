package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ThumbsDownOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M192 53.84S208 48 256 48s74 16 96 32h64a64 64 0 0 1 64 64v48a64 64 0 0 1-64 64h-30a32.34 32.34 0 0 0-27.37 15.4S350 290.19 324 335.22S248 448 240 464c-29 0-43-22-34-47.71c10.28-29.39 23.71-54.38 27.46-87.09c.54-4.78-3.14-12-8-12L96 307"/><path fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32" d="m96 241l80 2c20 1.84 32 12.4 32 30s-14 28.84-32 30l-80 4c-17.6 0-32-16.4-32-34v-.17A32 32 0 0 1 96 241Zm-32-65l112 2c18 .84 32 12.41 32 30c0 17.61-14 28.86-32 30l-112 2a32.1 32.1 0 0 1-32-32a32.1 32.1 0 0 1 32-32Zm48-128l64 3c21 1.84 32 11.4 32 29s-14.4 30-32 30l-64 2a32.09 32.09 0 0 1-32-32a32.09 32.09 0 0 1 32-32Zm-32 64l96 2c19 .84 32 12.4 32 30s-13 28.84-32 30l-96 2a32.09 32.09 0 0 1-32-32a32.09 32.09 0 0 1 32-32Z"/>`),
		g.Group(children),
	)
}