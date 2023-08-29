package foundation

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DieThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="currentColor" d="M83.921 23.292a7.44 7.44 0 0 0-7.213-7.264v-.011H23.292v.008a7.443 7.443 0 0 0-7.267 7.267h-.009v53.416h.009a7.44 7.44 0 0 0 7.267 7.267v.008h53.416v-.009a7.44 7.44 0 0 0 7.267-7.267h.009V23.292h-.063zM32.99 40.07a7.035 7.035 0 1 1 .002-14.068a7.035 7.035 0 0 1-.002 14.068zm17.009 16.965a7.033 7.033 0 0 1-7.034-7.033a7.034 7.034 0 1 1 14.069 0a7.034 7.034 0 0 1-7.035 7.033zm17.01 16.944a7.033 7.033 0 0 1-7.034-7.033a7.034 7.034 0 1 1 14.069 0a7.034 7.034 0 0 1-7.035 7.033z"/>`),
		g.Group(children),
	)
}