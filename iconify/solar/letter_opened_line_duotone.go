package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterOpenedLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M18 6.102c1.3.127 2.175.416 2.828 1.07C22 8.343 22 10.229 22 14c0 3.771 0 5.657-1.172 6.828C19.657 22 17.771 22 14 22h-4c-3.771 0-5.657 0-6.828-1.172C2 19.657 2 17.771 2 14c0-3.771 0-5.657 1.172-6.828C3.825 6.518 4.7 6.229 6 6.102" opacity=".5"/><path stroke-linecap="round" d="M10 6h4m-3 3h2m-4.841 2.8l-.72-.6c-.707-.59-1.061-.885-1.25-1.288C6 9.508 6 9.048 6 8.126V7c0-2.357 0-3.536.732-4.268C7.464 2 8.643 2 11 2h2c2.357 0 3.535 0 4.268.732C18 3.464 18 4.643 18 7v1.127c0 .92 0 1.381-.189 1.785c-.189.403-.543.698-1.25 1.287l-.72.6c-1.836 1.53-2.755 2.296-3.841 2.296c-1.086 0-2.004-.765-3.841-2.296Z"/><path stroke-linecap="round" d="m6 10l2.159 1.8c1.837 1.53 2.755 2.295 3.841 2.295c1.086 0 2.005-.765 3.841-2.296L18 10"/></g>`),
		g.Group(children),
	)
}