package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Taxifix(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.863 10.985h7.599M13.64 22.457l2.023-11.472m16.956 3.872l-7.076 7.6m5.736 0l-4.396-7.6m-3.531 15.201l-7.075 7.6m5.735 0l-4.395-7.6"/><ellipse cx="36.36" cy="11.344" fill="currentColor" rx=".819" ry=".687" transform="rotate(-42.481 36.36 11.344)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m35.741 14.857l-1.34 7.6M23.13 19.589c-.28 1.584-1.78 2.868-3.352 2.868h-1c-1.04 0-1.735-.85-1.55-1.9s1.18-1.9 2.221-1.9h3.85"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.818 15.61c.813-.692 1.308-.753 2.642-.753c1.514 0 2.392.666 2.074 2.466l-.905 5.134M8.5 30.056l6.335.008l-1.34 7.596"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m8.665 37.66l1.67-9.468c.195-1.109 1.251-2.003 2.36-2.003h.781c.988 0 1.568.288 1.895.835m10.64 10.636l2.018-11.471H39.5L26.011 37.66zm9.194-11.471l-8.463 7.312m5.005-7.312l-4.362 3.656"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}