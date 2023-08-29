package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlockedSiteSolidTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="M1877 250v501q0 144-37 274t-103 248t-155 221t-193 194t-219 168t-231 143q-116-61-230-138t-219-169t-194-197t-155-224t-103-248T0 751V250q84 0 159-5t147-22t142-44t143-75q42-26 83-45t83-33t87-19t95-7q51 0 96 6t87 20t83 33t83 47q71 45 141 73t142 44t148 21t158 6zm-438 1038l-349-349l351-352l-151-150l-351 351l-352-351l-150 150l351 352l-351 351l150 151l352-351l349 349l151-151z"/>`),
		g.Group(children),
	)
}