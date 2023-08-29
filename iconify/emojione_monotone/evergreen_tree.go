package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EvergreenTree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="m61 53.684l-12.88-8.936c4.842-.894 8.047-1.832 8.047-1.832l-10.834-9.02c3.628-.865 6.001-1.735 6.001-1.735l-9.123-9.523c2.6-.894 4.289-1.771 4.289-1.771L31.999 2L17.5 20.867s1.689.877 4.288 1.771l-9.121 9.523s2.372.87 5.999 1.734L7.833 42.916s3.205.938 8.046 1.832L3 53.684s10.053 2.456 22.233 3.209V62h13.534v-5.107C50.946 56.14 61 53.684 61 53.684z"/>`),
		g.Group(children),
	)
}