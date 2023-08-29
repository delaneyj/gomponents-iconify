package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TreatedWaterNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsTreatedWaterNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM26 28.581l-.402.419h4.688c.947 0 1.714.767 1.714 1.714v1.716c-2.193 1.07-4.386.665-6.58-1.216c-2.796 2.397-5.592 2.397-8.388 0c-2.344 2.01-4.688 2.334-7.032.975v-1.475c0-.947.768-1.714 1.714-1.714h4.688L16 28.581V25h2v4h2v-4h2v4h2v-4h2v3.581ZM29 42v-1.798H13V42h-2v-1.924a4.002 4.002 0 0 1-3-3.874v-5.488A3.71 3.71 0 0 1 11.714 27s-.683-.713-1.463-1.81C9.21 23.726 8 21.582 8 19.53V5.643h26V19.53c0 2.053-1.211 4.197-2.25 5.66c-.781 1.097-1.464 1.81-1.464 1.81a3.698 3.698 0 0 1 2.995 1.518h1.427v-1.9h4.494v2h-3.1a6.002 6.002 0 0 1 4.898 5.9v1.19h-4.494v-1.696a1 1 0 0 0-1-1H34v3.19a4.002 4.002 0 0 1-3 3.874V42h-2Zm3-34.357v7.505c-8.545 1.237-13.536 1.25-22-.148V7.643h22Zm9 32.539C41 41.186 40.328 42 39.5 42s-1.5-.814-1.5-1.818C38 38.59 39.5 37 39.5 37s1.5 1.59 1.5 3.182Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsTreatedWaterNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}