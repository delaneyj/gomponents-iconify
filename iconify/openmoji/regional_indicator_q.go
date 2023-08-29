package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RegionalIndicatorQ(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<defs><path id="openmojiRegionalIndicatorQ0" d="M36 21.012c-6.065 0-11 4.935-11 11v8c0 6.066 4.935 11 11 11c2.657 0 5.097-.947 7-2.522l2.248 2.232a1 1 0 1 0 1.408-1.419l-2.239-2.223A10.946 10.946 0 0 0 47 40.011v-8c0-6.065-4.935-11-11-11Zm9 19c0 2.14-.754 4.107-2.007 5.654l-3.089-3.067a1 1 0 0 0-1.408 1.419l3.073 3.05A8.948 8.948 0 0 1 36 49.013c-4.963 0-9-4.037-9-9v-8c0-4.962 4.037-9 9-9s9 4.038 9 9v8Z"/></defs><use href="#openmojiRegionalIndicatorQ0"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><circle cx="36" cy="36" r="28"/><use href="#openmojiRegionalIndicatorQ0"/></g>`),
		g.Group(children),
	)
}