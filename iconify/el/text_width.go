package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextWidth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="m257.813 0l-5.859 213.062h74.048c0-146.136.027-145.972 186.768-145.972v587.77c0 73.264-.036 73.142-102.538 71.338v58.521h379.54v-58.521c-102.335 0-102.538-.18-102.538-71.338V67.09c186.839 0 187.279.041 187.279 145.972h73.535L942.7 0H257.813zm-48.414 783.398L0 991.699L209.399 1200v-139.38H990.6V1200L1200 991.699L990.601 783.398V922.85H209.399V783.398z"/>`),
		g.Group(children),
	)
}