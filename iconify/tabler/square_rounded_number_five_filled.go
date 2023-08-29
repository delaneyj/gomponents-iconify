package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SquareRoundedNumberFiveFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M0 0h24v24H0z"/><path fill="currentColor" d="m12 2l.642.005l.616.017l.299.013l.579.034l.553.046c4.687.455 6.65 2.333 7.166 6.906l.03.29l.046.553l.041.727l.006.15l.017.617L22 12l-.005.642l-.017.616l-.013.299l-.034.579l-.046.553c-.455 4.687-2.333 6.65-6.906 7.166l-.29.03l-.553.046l-.727.041l-.15.006l-.617.017L12 22l-.642-.005l-.616-.017l-.299-.013l-.579-.034l-.553-.046c-4.687-.455-6.65-2.333-7.166-6.906l-.03-.29l-.046-.553l-.041-.727l-.006-.15l-.017-.617l-.004-.318v-.648l.004-.318l.017-.616l.013-.299l.034-.579l.046-.553c.455-4.687 2.333-6.65 6.906-7.166l.29-.03l.553-.046l.727-.041l.15-.006l.617-.017c.21-.003.424-.005.642-.005zm2 5h-4a1 1 0 0 0-.993.883L9 8v4a1 1 0 0 0 .883.993L10 13h3v2h-2l-.007-.117A1 1 0 0 0 9 15a2 2 0 0 0 1.85 1.995L11 17h2a2 2 0 0 0 1.995-1.85L15 15v-2a2 2 0 0 0-1.85-1.995L13 11h-2V9h3a1 1 0 0 0 .993-.883L15 8a1 1 0 0 0-.883-.993L14 7z"/></g>`),
		g.Group(children),
	)
}