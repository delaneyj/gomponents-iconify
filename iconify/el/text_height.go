package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextHeight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M991.699 0L783.398 209.399H922.85V990.6H783.398L991.699 1200L1200 990.601h-139.38V209.399H1200L991.699 0zM5.933 207.642L0 420.703h74.048c0-146.136.1-145.972 186.841-145.972V862.5c0 73.264-.037 73.141-102.539 71.338v58.521h379.468v-58.521c-102.335 0-102.539-.179-102.539-71.338V274.731c186.839 0 187.354.041 187.354 145.972h73.535l-5.347-213.062l-684.888.001z"/>`),
		g.Group(children),
	)
}