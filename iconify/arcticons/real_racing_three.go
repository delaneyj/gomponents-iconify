package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RealRacingThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.838 26.043a8.484 8.484 0 0 1 9.542-3.752a13.32 13.32 0 0 0 4.47-.126c1.121-1.283 1.09-1.968 3.688-2.71s9.954-.067 10.233 0a5.996 5.996 0 0 1 4.359 2.789s4.43-.929 4.911 1.655"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.838 26.043s-1.388.35-.86 2.605s-.76 2.174-.42 3.273s.785 1.966 3.101 2.739a43.466 43.466 0 0 0 11.357 2.073a26.51 26.51 0 0 0 9.184-.606s-.114-7.689 2.008-9.793c2.123-2.105 3.515-.292 3.623 1.12a13.914 13.914 0 0 1-1.104 6.399l8.251-3.971c-.238-.176-.214-6.672 1.802-5.908c.96.363.683 3.288.683 3.288"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.008 28.164S16.6 29.73 19.1 29.58s3.675.04 5.587-1.762s2.666-2.474.198-2.377s-5.877 2.723-5.877 2.723Zm-13.17-2.121c.82.675 3.12 1.585 3.515-2.903m19.172 12.315s2.528 1.976 4.202-1.601m8.251-3.972s2.241 1.878 2.485-2.62M5.197 18.436c.552.463 1.148.675 2.487.675h.312a2 2 0 0 0 2-2h0a2 2 0 0 0-2-1.999m-2.8-3.332c.553-.462 1.15-.672 2.49-.669l.31.002a2 2 0 0 1 2 2h0a2 2 0 0 1-2 1.999m-1.537 0h1.537"/>`),
		g.Group(children),
	)
}