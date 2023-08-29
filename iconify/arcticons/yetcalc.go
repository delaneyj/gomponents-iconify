package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Yetcalc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 5.5v37M5.5 24h37m-13.773 7.096h9m-9 3.982h9"/><circle cx="14.936" cy="12.684" r=".75" fill="currentColor"/><circle cx="14.936" cy="16.668" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.436 14.676h9m-9 0h9m-2.357 19.586c-.253.505-.841.841-1.43.841a1.687 1.687 0 0 1-1.682-1.682v-1.093c0-.925.757-1.682 1.682-1.682s1.682.757 1.682 1.682v.589h-3.364m-1.737.504v2.27c0 .926-.756 1.683-1.68 1.683c-.505 0-.926-.169-1.178-.505"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.23 30.646v2.775c0 .925-.756 1.682-1.68 1.682s-1.683-.757-1.683-1.682v-2.775m10.57-1.431v5.046c0 .504.336.84.84.84h.253m-1.934-4.455h1.766m13.37-14.673c0 .925-.757 1.682-1.682 1.682s-1.682-.757-1.682-1.682V14.88c0-.925.757-1.681 1.682-1.681s1.682.756 1.682 1.681m0 2.774v-4.456m-4.625 3.616c-.253.504-.841.84-1.43.84a1.687 1.687 0 0 1-1.682-1.681V14.88c0-.925.757-1.682 1.682-1.682c.589 0 1.177.336 1.43.84m6.066-3.111v5.886c0 .505.336.841.84.841h.253m4.194-.84c-.253.504-.841.84-1.43.84a1.687 1.687 0 0 1-1.682-1.681V14.88c0-.925.757-1.682 1.682-1.682c.589 0 1.177.336 1.43.84"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}