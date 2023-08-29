package emojione_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagForWesternSahara(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="#e6e7e8" d="M1 26h63v14H1z"/><path fill="#137a08" d="M18.68 40L4.121 53.238C5.656 54.345 7.609 55 10 55h44c6.627 0 10-4.925 10-11v-4H18.68z"/><path fill="#25333a" d="M54 11H10c-2.201 0-4.03.553-5.514 1.5L18.879 26h45.12v-4c0-6.075-3.373-11-10-11"/><path fill="#ec1c24" d="M26.416 33L4.486 12.5C1.513 14.407 0 17.942 0 22v22c0 3.883 1.385 7.289 4.11 9.246l22.344-20.219l-.038-.027zm11.614-1.247l-2.627.004l-.816-2.657l-.809 2.657l-2.628-.004l2.13 1.619l-.827 2.638l2.138-1.64l2.139 1.64l-.827-2.638z"/><path fill="#ec1c24" fill-rule="evenodd" d="M34.37 27.906c-2.711.661-4.02 2.967-3.993 4.998c.025 2.305 1.573 4.356 3.981 4.923c-1.159.296-3.177-.207-4.438-1.686c-1.453-1.703-1.601-4.255-.331-6.104c1.247-1.818 3.387-2.419 4.781-2.131"/>`),
		g.Group(children),
	)
}