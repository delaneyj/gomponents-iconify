package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Serverless(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#F26D61" d="M0 161.202h45.312l-14.039 42.396H0v-42.396ZM0 80.6h72l-14.036 42.396H0V80.601ZM0 0h98.692l-14.04 42.395H0V0Zm143.349 0H256v42.395H129.312L143.349 0ZM116.66 80.6H256v42.397H102.622l14.038-42.396Zm-26.69 80.602H256v42.396H75.933l14.037-42.396Z"/>`),
		g.Group(children),
	)
}