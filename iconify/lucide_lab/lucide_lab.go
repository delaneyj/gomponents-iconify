package lucide_lab

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const IconifyVersion = ""

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 24 24")
)

func AmpersandSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M16 17c-4-2-7-6-7-8a2 2 0 0 1 4 0c0 3-5 1.5-5 5c0 1.7 1.3 3 3 3c3 0 5-2 5-4"/></g>`), g.Group(children),
	)
}

func AppleCore(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 2a2 2 0 0 0-2 2v2.53"/><path d="M12 6.53a5.98 5.98 0 0 0-8.5.5a4 4 0 0 1 4.02 5.86a4 4 0 0 1-1.76 7.04C6.82 21.17 7.97 22 9 22c1.5 0 1.5-1 3-1s1.5 1 3 1c1.03 0 2.18-.83 3.24-2.07a4 4 0 0 1-1.76-7.03a4 4 0 0 1 4.02-5.87a5.99 5.99 0 0 0-8.5-.5"/></g>`), g.Group(children),
	)
}

func ArrowsUpDownSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m6 9l3-3l3 3M9 6v6"/><rect width="20" height="20" x="2" y="2" rx="2"/><path d="M15 18v-6m3 3l-3 3l-3-3"/></g>`), g.Group(children),
	)
}

func AstronautHelmet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10.6 13.4l6.3 6.1c.3.5.1 1.1-.4 1.4c-1.4.7-2.9 1.1-4.5 1.1a2 2 0 0 1-1.4-.6l-8-8A2 2 0 0 1 2 12a10 10 0 0 1 19.44-3.3c.3.7-.3 1.3-1 1.3H12"/><circle cx="12" cy="12" r="2"/><path d="M16.2 18.8c3-1.9 4.4-5.5 3.5-8.8"/></g>`), g.Group(children),
	)
}

func AtSignCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M14 17.7a6 6 0 1 1 4-5.7a2 2 0 0 1-4 0"/><circle cx="12" cy="12" r="2"/></g>`), g.Group(children),
	)
}

func AtSignSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="2"/><path d="M14 17.7a6 6 0 1 1 4-5.7a2 2 0 0 1-4 0"/><circle cx="12" cy="12" r="2"/></g>`), g.Group(children),
	)
}

func Avocado(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 7a4.95 4.95 0 0 0-8.6-3.4c-1.5 1.6-1.6 1.8-5 2.6a8 8 0 1 0 9.4 9.5c.7-3.4 1-3.6 2.6-5c1-1 1.6-2.3 1.6-3.7"/><circle cx="10" cy="14" r="3.5"/></g>`), g.Group(children),
	)
}

func BabyPacifier(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.1 7.4a1.95 1.95 0 0 0 3.7-1.5c-.8-2-3.2-3-5.2-2.2c-2.9 1.2-4.8 3.7-5.4 6.5a1.95 1.95 0 0 0 0 3.6A9.05 9.05 0 0 0 7 19.42m10.1-.02c2-1.3 3.3-3.4 3.8-5.6a2 2 0 0 0 0-3.6a9.83 9.83 0 0 0-3.2-5M8 12h.01M16 12h.01"/><circle cx="12" cy="16" r="2"/><path d="M10 16h-.5A2.5 2.5 0 0 0 7 18.5v1A2.5 2.5 0 0 0 9.5 22h5a2.5 2.5 0 0 0 2.5-2.5v-1a2.5 2.5 0 0 0-2.5-2.5H14"/></g>`), g.Group(children),
	)
}

func Bacon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 21a1 1 0 0 0 1-1v-5.35c0-.457.316-.844.727-1.041a4 4 0 0 0-2.134-7.589a5 5 0 0 0-9.186 0a4 4 0 0 0-2.134 7.588c.411.198.727.585.727 1.041V20a1 1 0 0 0 1 1ZM6 17h12"/>`), g.Group(children),
	)
}

func BagHand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 8c0-2.8 1.8-5 4-5s4 2.2 4 5m5 10.6l-2-9.8c-.1-.5-.5-.8-1-.8H6c-.5 0-.9.3-1 .8l-2 9.8v.4a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2ZM12 12v4"/><path d="M18 8A6 6 0 0 1 6 8"/></g>`), g.Group(children),
	)
}

func Barbecue(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 4c0-1 2-1 2-2m4 2c0-1 2-1 2-2m4 2c0-1 2-1 2-2M3 8a9.06 9 0 0 0 18 0Zm6.2 7.6l-1.3 2.6"/><circle cx="7" cy="20" r="2"/><path d="M9 20h8m-2.2-4.4L18 22"/></g>`), g.Group(children),
	)
}

func BarberPole(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6h10M7 22h10m-9 0V6a4 4 0 0 1 8 0v16M8 11.5l8-4M8 16l8-4m-8 8.5l8-4"/>`), g.Group(children),
	)
}

func Barn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12H2l2-6l8-4l8 4Z"/><path d="M10 8h4v4h-4zM7 22l10-10v10L7 12Z"/><path d="M21 12v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-8"/></g>`), g.Group(children),
	)
}

func Baseball(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12c5.5 0 10-4.5 10-10"/><circle cx="12" cy="12" r="10"/><path d="M22 12c-5.5 0-10 4.5-10 10M8 11.5l-1.5-2m5-1.5l-2-1.5m5 11l-2-1.5m5-1.5l-1.5-2"/></g>`), g.Group(children),
	)
}

func BaselineSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9.5 12h5M9 13l3-6l3 6m-8 4h10"/></g>`), g.Group(children),
	)
}

func Basketball(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M2.1 13.4A10.1 10.1 0 0 0 13.4 2.1M5 4.9l14 14.2m2.9-8.5a10.1 10.1 0 0 0-11.3 11.3"/></g>`), g.Group(children),
	)
}

func BatBall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="4"/><path d="m4 8l10 10m6.8-2.8c1.9-3.4 1.4-7.7-1.4-10.6c-3.5-3.5-9.1-3.5-12.5 0c-4.7 4.7-5.1 6.9-1.4 11.1l-2.9 2.9c-.8.8-.8 2 0 2.8s2 .8 2.8 0l2.9-2.9c2.6 2.3 4.5 3 6.6 2.1"/></g>`), g.Group(children),
	)
}

func BathBubble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 3h.01"/><circle cx="11.5" cy="6.5" r=".5"/><circle cx="16.5" cy="7.5" r=".5"/><path d="M2 12h6m5 3H8v-3c0-.6.4-1 1-1h3c.6 0 1 .4 1 1Zm0-3h9"/><path d="M4 12v5a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5M7 19v2m10-2v2"/></g>`), g.Group(children),
	)
}

func BeachBall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M19.1 4.9c-1.6-1.6-6 .3-9.9 4.2S3.4 17.4 5 19s6-.3 9.9-4.2c3.8-3.9 5.7-8.3 4.2-9.9"/></g>`), g.Group(children),
	)
}

func BearFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 7l.5.5M18 7l-.5.5m3.3-3.3c-1.6-1.6-4.1-1.6-5.7 0l-1 1a13.6 13.6 0 0 0-4.2 0l-1-1a4 4 0 0 0-5.8 5.55A7 7 0 0 0 2 13.5C2 18.2 6.5 22 12 22s10-3.8 10-8.5a7 7 0 0 0-1.1-3.8c1.5-1.6 1.5-4-.1-5.5M10 12v-.5m4 .5v-.5m0 4.5h-4m2 0v2"/>`), g.Group(children),
	)
}

func BedBunk(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 22V2m0 3h18a2 2 0 0 1 2 2v15M6 5v5m-4 0h20M2 14h20m0 5H2m4-5v5"/>`), g.Group(children),
	)
}

func Bee(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 2l1.88 1.88m4.24 0L16 2M9 7V6a3 3 0 1 1 6 0v1M5 7a3 3 0 1 0 2.2 5.1C9.1 10 12 7 12 7s2.9 3 4.8 5.1A3 3 0 1 0 19 7Zm2.56 5h8.87M7.5 17h9"/><path d="M15.5 10.7c.9.9 1.4 2.1 1.5 3.3c0 5.8-5 8-5 8s-5-2.2-5-8c.1-1.2.6-2.4 1.5-3.3"/></g>`), g.Group(children),
	)
}

func BeeHive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="10" height="4" x="7" y="2" rx="2"/><rect width="16" height="4" x="4" y="6" rx="2"/><path d="M14 14h6a2 2 0 1 0 0-4H4a2 2 0 1 0 0 4h6"/><rect width="4" height="8" x="10" y="10" rx="2"/><path d="M19 14a2 2 0 1 1 0 4H5a2 2 0 1 1 0-4"/><rect width="14" height="4" x="5" y="18" rx="2"/></g>`), g.Group(children),
	)
}

func BeetleScarab(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 2l1.88 1.88m4.24 0L16 2M9 7.13V6a3 3 0 1 1 6 0v1.13"/><path d="M12 20c-3.3 0-6-2.7-6-6v-3a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v3c0 3.3-2.7 6-6 6m0 0v-9"/><path d="M5 4.8C3.2 6.2 2 8.5 2 11h20c0-2.5-1.2-4.8-3-6.2M6.08 15h-4c.2 2.4 1.25 4.4 2.8 6m14.22 0a9 9.4 0 0 0 2.82-6h-4"/></g>`), g.Group(children),
	)
}

func BellConcierge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 20a1 1 0 0 1-1-1v-1a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1Zm17-4a8 8 0 1 0-16 0m8-12v4m-2-4h4"/>`), g.Group(children),
	)
}

func BellConciergeDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="19" cy="9" r="3"/><path d="M2 18a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v2H2ZM12 4v4c-4.4 0-8 3.6-8 8m6-12h4"/></g>`), g.Group(children),
	)
}

func BellConciergeOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 2l20 20M12 4v2.3M10 4h4m5.8 10.1a8 8 0 0 0-5.9-5.9m-5.2.5C5.9 10 4 12.8 4 16"/><path d="M16 16H4a2 2 0 0 0-2 2v2h18"/></g>`), g.Group(children),
	)
}

func Belt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7.3 9H3c-.6 0-1-.4-1-1V4c0-.6.4-1 1-1h4.3M6 6h3m4 0h.01"/><rect width="10" height="8" x="7" y="2" rx="2"/><path d="M16.7 3H21c.6 0 1 .4 1 1v4c0 .6-.4 1-1 1h-4.3m-6.2 1l-8.1 6.2m19.2-7.4L12.2 16M3 22c-.6 0-1-.4-1-1v-4c0-.6.4-1 1-1h16l3 3l-3 3Z"/></g>`), g.Group(children),
	)
}

func BoldSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 12h5.5a2.5 2.5 0 0 1 0 5H8V7h5a2.5 2.5 0 0 1 0 5"/></g>`), g.Group(children),
	)
}

func BottleBaby(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 11c1.1-1.4 1.3-3.3.7-4.9l.8-.8a1.5 1.5 0 0 0-2.8-2.8l-.8.8A5.33 5.33 0 0 0 13 4"/><path d="M11.3 3.7a1 1 0 0 1 1.4 0l7.6 7.6a1 1 0 0 1 0 1.4l-1.6 1.6a1 1 0 0 1-1.4 0L9.7 6.7a1 1 0 0 1 0-1.4Z"/><path d="m10 7l-7.3 7.3c-.9.9-.9 2.5 0 3.4l3.6 3.6c.9.9 2.5.9 3.4 0L17 14M4 13l2 2m1-5l2 2"/></g>`), g.Group(children),
	)
}

func BottleChampagne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 2h.01M12 3h.01M19 8l-3-3M9.7 21.3a2.4 2.4 0 0 1-3.4 0l-3.6-3.6a2.41 2.41 0 0 1 0-3.4l6.27-6.27A3.5 3.5 0 0 1 11.45 7h1.1a3.5 3.5 0 0 0 2.47-1.03l3.62-3.61a1.21 1.21 0 0 1 1.72 0l1.28 1.28a1.2 1.2 0 0 1 0 1.72l-3.62 3.61A3.5 3.5 0 0 0 17 11.45v1.1a3.5 3.5 0 0 1-1.03 2.48Z"/><path d="m9.06 8l3.23 3.24a1 1 0 0 1 0 1.41L8.65 16.3a1 1 0 0 1-1.41 0L4 13.06M21 12h.01m.99 4h.01"/></g>`), g.Group(children),
	)
}

func BottleDispenser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18.5" cy="5.5" r=".5"/><path d="M20 10h.01M9 2h7m-5 0v4"/><rect width="4" height="4" x="9" y="6" rx="1"/><path d="M9 10c-1.7 0-3 1.3-3 3v7a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-7c0-1.7-1.3-3-3-3Z"/><path d="M6 14.5a6 6 0 0 1 5 0s2 1.25 5 0"/></g>`), g.Group(children),
	)
}

func BottlePerfume(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3h12v4H6zm3 4h6v4H9z"/><rect width="18" height="10" x="3" y="11" rx="2"/></g>`), g.Group(children),
	)
}

func BottlePlastic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 6.1V3c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v3.1"/><path d="M17 14.5c0-1.2-.9-2.2-2-2.4V12a2 2 0 0 0 2-2a4 4 0 0 0-4-4h-2a4 4 0 0 0-4 4a2 2 0 0 0 2 2v.1a2.5 2.5 0 0 0 0 4.8v.1a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-1a2 2 0 0 0-2-2v-.1c1.1-.2 2-1.2 2-2.4"/></g>`), g.Group(children),
	)
}

func BottleSpray(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 2v4m-6 4h4m0-2a2 2 0 0 1 2-2h3c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1H5C3.3 2 2 3.3 2 5c0 .6.4 1 1 1h1a2 2 0 0 1 2 2v2l-2.3 2.3c-.4.4-.7 1.1-.7 1.7v6a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-6c0-.6-.3-1.3-.7-1.7L10 10Z"/><path d="M14 6c0 2 0 3 2 4M3 16.5a6 6 0 0 1 5 0s2 1.25 5 0M22 2h.01M20 5.5h.01M22 9h.01"/></g>`), g.Group(children),
	)
}

func BottleToothbrushComb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 3h4v4H2zm4 4v2c0 .6.1 1.4.2 2L8 20.8v.2c0 .6-.4 1-1 1H3c-.6 0-1-.4-1-1V7m12-5v7l-2 5v8"/><path d="M10 3h4v6h-4zm8 3h4m-4 4h4m-4 4h4m-4 4h4M18 2h2a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2h-2"/></g>`), g.Group(children),
	)
}

func BowlChopsticks(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13 2l-3 11M22 2l-8 11"/><ellipse cx="12" cy="12" rx="10" ry="5"/><path d="M22 12a10 10 0 0 1-20 0"/></g>`), g.Group(children),
	)
}

func BowlOverflow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4C6.5 4 2 6.2 2 9c0 2.4 3.4 4.4 8 4.9c1.1.1 2 1 2 2.1v3a2 2 0 0 0 4 0v-3.4c0-1.1.9-2.2 1.9-2.6c2.5-.9 4.1-2.4 4.1-4c0-2.8-4.5-5-10-5"/><path d="M2 9c0 5.5 4.5 10 10 10m4-.8c3.5-1.5 6-5.1 6-9.2m-6 6.6c0-2.6 3-2.6 3-4.6c0-1.7-3.1-3-7-3s-7 1.3-7 3c0 1.4 2.1 2.5 5 2.9"/></g>`), g.Group(children),
	)
}

func Bowling(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 10h.01M6 13h.01M10 14h.01m1.08-7.93a8 8 0 1 0 .32 15.81M16 9h4m-5-4c0 1.5 1 2 1 4c0 2.5-2 4.5-2 7c0 2.6 1.9 6 1.9 6H20s2-3.4 2-6c0-2.5-2-4.5-2-7c0-2 1-2.5 1-4a3 3 0 1 0-6 0"/>`), g.Group(children),
	)
}

func BraSports(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 21l-9-9V3H4v10.6A4 4 0 0 0 6 21h12a4 4 0 0 0 2-7.4V3h-4v9l-4 4m-4-5h8"/>`), g.Group(children),
	)
}

func BriefcasePlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 7V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/><rect width="20" height="14" x="2" y="7" rx="2"/><path d="M15 14H9m3-3v6"/></g>`), g.Group(children),
	)
}

func Broom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 11l9-9m-7.4 10.6c.8.8.9 2.1.2 3L10 22l-8-8l6.4-4.8c.9-.7 2.2-.6 3 .2Zm-7.8-2.2l6.8 6.8M5 17l1.4-1.4"/>`), g.Group(children),
	)
}

func Bucket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 7c0-2.8 2.2-5 5-5h2c2.8 0 5 2.2 5 5M5 11h14m-1 0l-.8 9c-.1 1.1-1.1 2-2.2 2H9c-1.1 0-2.1-.9-2.2-2L6 11"/>`), g.Group(children),
	)
}

func BullHead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 10a5 5 0 0 1-4-8a4 4 0 0 0 4 4h10a4 4 0 0 0 4-4a5 5 0 0 1-4 8"/><path d="M6.4 15c-.3-.6-.4-1.3-.4-2c0-4 3-3 3-7m1 6.5v1.6m7.6.9c.3-.6.4-1.3.4-2c0-4-3-3-3-7m-1 6.5v1.6"/><path d="M15 22a4 4 0 1 0-3-6.7A4 4 0 1 0 9 22Zm-6-4h.01M15 18h.01"/></g>`), g.Group(children),
	)
}

func Burger(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 12a2 2 0 0 1-2-2a9 7 0 0 1 18 0a2 2 0 0 1-2 2l-3.5 4.1c-.8 1-2.4 1.1-3.4.3L7 12"/><path d="M11.7 16H4a2 2 0 0 1 0-4h16a2 2 0 0 1 0 4h-4.3M5 16a2 2 0 0 0-2 2c0 1.7 1.3 3 3 3h12c1.7 0 3-1.3 3-3a2 2 0 0 0-2-2"/></g>`), g.Group(children),
	)
}

func Butterfly(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.8 2C12 3.8 12 9 12 9s0-5.2-3.8-7M12 9v11"/><path d="M20 5c-3.5 0-6.5 3.9-8 6.3C10.5 8.9 7.5 5 4 5a2 2 0 0 0-2 2c0 2.3.6 4.4 1.5 5.6C4 13.5 4.9 14 6 14h2c-.9.4-2.1.9-2.6 1.5c-1.6 1.6-.9 3.4.7 4.9c1.6 1.6 3.4 2.3 4.9.7c.3-.3 1-1.1 1-1.1s.6.8 1 1.1c1.6 1.6 3.4.9 4.9-.7c1.6-1.6 2.3-3.4.7-4.9c-.5-.5-1.7-1.1-2.6-1.5h2c1.1 0 2-.5 2.5-1.4c.9-1.2 1.5-3.3 1.5-5.6a2 2 0 0 0-2-2"/></g>`), g.Group(children),
	)
}

func Cabin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2.6 10.4a2.12 2.12 0 1 0 3.02 2.98L12 7l6.4 6.4a2.12 2.12 0 1 0 2.979-3.021L13.7 2.7a2.4 2.4 0 0 0-3.404.004Z"/><path d="M14 22v-7a2 2 0 0 0-4 0v7"/><path d="M14 14h6v6a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-6h6m-6 4h6m4 0h6"/></g>`), g.Group(children),
	)
}

func CabinetFiling(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 12h16"/><rect width="16" height="20" x="4" y="2" rx="2"/><path d="M10 6h4m-4 10h4"/></g>`), g.Group(children),
	)
}

func Cactus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8v6a2 2 0 0 0 2 2h2m6-2h2a2 2 0 0 0 2-2V6M9 22V5a3 3 0 1 1 6 0v17m-8 0h10"/>`), g.Group(children),
	)
}

func CandleHolder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 11h4v7H8z"/><path d="M18 18a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4h18a2 2 0 1 0-2-2Zm-8-9v2"/></g>`), g.Group(children),
	)
}

func CandleHolderLit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 2S8 3.9 8 5s.9 2 2 2s2-.9 2-2s-2-3-2-3m-2 9h4v7H8zm5 2l-1-2"/><path d="M18 18a4 4 0 0 1-4 4H6a4 4 0 0 1-4-4h18a2 2 0 1 0-2-2Z"/></g>`), g.Group(children),
	)
}

func CandleTealight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 12V7"/><ellipse cx="12" cy="13" rx="10" ry="3"/><path d="M2 13v6c0 1.7 4.5 3 10 3s10-1.3 10-3v-6"/></g>`), g.Group(children),
	)
}

func CandleTealightLit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 2s-2 1.9-2 3s.9 2 2 2s2-.9 2-2s-2-3-2-3m0 10V7"/><ellipse cx="12" cy="13" rx="10" ry="3"/><path d="M2 13v6c0 1.7 4.5 3 10 3s10-1.3 10-3v-6M8 16v1m4-1v2"/></g>`), g.Group(children),
	)
}

func Candlestick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 11h4v11h-4zm2-3v3"/>`), g.Group(children),
	)
}

func CandlestickBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 22H8v-7c0-.6.4-1 1-1h6c.6 0 1 .4 1 1Zm-4-11v3"/>`), g.Group(children),
	)
}

func CandlestickBigLit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 2S9 5.3 9 7s1.3 3 3 3s3-1.3 3-3s-3-5-3-5m4 20H8v-7c0-.6.4-1 1-1h6c.6 0 1 .4 1 1Zm-4-8v3"/><path d="M17 17s-.7-1.4-1.1-2.4"/></g>`), g.Group(children),
	)
}

func CandlestickLit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2s-2 1.9-2 3s.9 2 2 2s2-.9 2-2s-2-3-2-3m-2 9h4v11h-4zm5 2l-1-2"/>`), g.Group(children),
	)
}

func CardCredit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="14" x="2" y="5" rx="2"/><path d="M2 10h20"/></g>`), g.Group(children),
	)
}

func CardSd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 22a2 2 0 0 1-2-2V6l4-4h10a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2Zm2-12V7.5M12 6v4m4-4v4"/>`), g.Group(children),
	)
}

func CardSim(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h9l5 5v13a2 2 0 0 1-2 2Z"/><rect width="8" height="8" x="8" y="10" rx="2"/><path d="M8 14h8m-4 0v4"/></g>`), g.Group(children),
	)
}

func Carton(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 6V3c0-.6-.4-1-1-1H9c-.6 0-1 .4-1 1v3l-3 4v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V10Z"/><path d="M13 22V10l3-4H8m-3 4h8"/></g>`), g.Group(children),
	)
}

func CartonOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.1 8.5L16 6V3c0-.6-.4-1-1-1H9"/><path d="M11.7 6H16l3 4v3.3M2 2l20 20m-3-3v1a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V10l2.1-2.9M13 13v9"/></g>`), g.Group(children),
	)
}

func CaseCamel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7" cy="13" r="3"/><path d="M10 10v6m4-4h4.5a2 2 0 0 1 0 4H14V8h4a2 2 0 0 1 0 4"/></g>`), g.Group(children),
	)
}

func CaseKebab(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="4.5" cy="13.5" r="2.5"/><path d="M7 11v5m4-3h2"/><circle cx="19.5" cy="13.5" r="2.5"/><path d="M17 9v7"/></g>`), g.Group(children),
	)
}

func CaseSnake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 11v5"/><circle cx="4.5" cy="13.5" r="2.5"/><path d="M11 16h2"/><circle cx="19.5" cy="13.5" r="2.5"/><path d="M17 9v7"/></g>`), g.Group(children),
	)
}

func CaseSnakeUpper(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 16v-5.5a2.5 2.5 0 0 1 5 0V16m0-3H2m9 3h2m4-4h3a2 2 0 1 1 0 4h-3V8h2.5a2 2 0 0 1 .1 4"/>`), g.Group(children),
	)
}

func CatBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 7l.5.5M18 7l-.5.5M5 13a5 5 0 1 0 6.8 7.2l3-3.6A1 1 0 0 0 14 15h-4a1 1 0 0 0-.8 1.6l3 3.6A5 5 0 1 0 19 13h3c0-1.2-.4-2.4-1-3.4a3 3 0 0 0-5.8-5.3l-1 1a7 4 0 0 0-4.4 0l-1-1A3 3 0 0 0 3 9.6c-.6 1-1 2.2-1 3.4Zm5-2v-.5m4 .5v-.5M5 18H2m17 0h3"/>`), g.Group(children),
	)
}

func Cauldron(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="8" cy="3.5" r=".5"/><circle cx="14" cy="6" r="2"/><path d="M19 2h.01M22 8H2m5 4V8m4 2V8M4.4 8C2.9 9.5 2 11.4 2 13.5C2 18.2 6.5 22 12 22s10-3.8 10-8.5c0-2.1-.9-4-2.4-5.5M5 20l-1 2m15-2l1 2"/></g>`), g.Group(children),
	)
}

func Cent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 8a8 8 0 1 0 0 8M12 2v20"/>`), g.Group(children),
	)
}

func CentCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v12m4-9a5 5 0 1 0 0 6"/></g>`), g.Group(children),
	)
}

func CentSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 7v10m3.4-7a4 4 0 1 0 0 4"/></g>`), g.Group(children),
	)
}

func ChairsTableParasol(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v13m9-13H3l9-6ZM7 13h10m-9 8v-3a1 1 0 0 0-1-1H3m0-5v9m13 0v-3c0-.5.5-1 1-1h4m0-5v9"/>`), g.Group(children),
	)
}

func ChairsTablePlatter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6V5m-4 5a4 4 0 0 1 8 0M6 10h12m-6 0v9m-4 0v-4c0-.6-.4-1-1-1H2m0-6v11m14 0v-4a1 1 0 0 1 1-1h5m0-6v11"/>`), g.Group(children),
	)
}

func Chameleon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 22c-5 0-9-4.5-9-10S6 2 11 2c2.2 0 4.2.9 5.7 2.3L19.3 2c3.1 3.1 3.5 7.9 1.3 11.4c-.6.9-1.9.9-2.7.1l-1.2-1.2C15.2 10.9 13.2 10 11 10a6 6 0 0 0 0 12a4 4 0 0 0 0-8a2 2 0 0 0 0 4m3-11h.01"/><circle cx="14.5" cy="7" r="3.5"/><path d="M8 10.8L6 10l1-2m15 14a2 2 0 0 1-2-2v-6.1"/></g>`), g.Group(children),
	)
}

func Cheese(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 19v-7c-1-6-7-9-7-9l-2.1 1.5a2 2 0 0 1-3 2.2L3 11v9c0 .6.4 1 1 1h3a2 2 0 0 1 4 0h8M9 12H3"/><path d="M9 12c0-.8 1.3-1.5 3-1.5s3 .7 3 1.5a3 3 0 1 1-6 0m12 0h-6"/><circle cx="19" cy="19" r="2"/></g>`), g.Group(children),
	)
}

func Chest(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 19a2 2 0 0 0 2-2V9a4 4 0 0 0-8 0v8a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a4 4 0 0 0-4-4H6m-4 6h20m-6 0v3"/>`), g.Group(children),
	)
}

func ChevronsLeftRightSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m10 15l-3-3l3-3m4 0l3 3l-3 3"/></g>`), g.Group(children),
	)
}

func ChevronsUpDownSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m9 10l3-3l3 3m0 4l-3 3l-3-3"/></g>`), g.Group(children),
	)
}

func Cloth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5a1.41 1.41 0 0 0 0 2l1 1a1.41 1.41 0 0 1 0 2l-1 1a1.41 1.41 0 0 0 0 2l1 1a1.41 1.41 0 0 1 0 2l-1 1a1.41 1.41 0 0 0 0 2l2 2a1.41 1.41 0 0 0 2 0l1-1a1.41 1.41 0 0 1 2 0l1 1a1.41 1.41 0 0 0 2 0l1-1a1.41 1.41 0 0 1 2 0l1 1a1.41 1.41 0 0 0 2 0l2-2a1.41 1.41 0 0 0 0-2l-1-1a1.41 1.41 0 0 1 0-2l1-1a1.41 1.41 0 0 0 0-2l-1-1a1.41 1.41 0 0 1 0-2l1-1a1.41 1.41 0 0 0 0-2l-2-2a1.41 1.41 0 0 0-2 0l-1 1a1.41 1.41 0 0 1-2 0l-1-1a1.41 1.41 0 0 0-2 0l-1 1a1.41 1.41 0 0 1-2 0L7 3a1.41 1.41 0 0 0-2 0Z"/>`), g.Group(children),
	)
}

func CoatHanger(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 5a3 3 0 1 1 5.1 2.1l-1.5 1.5A2 2 0 0 0 12 10v1M4 21a2 2 0 0 1-1.1-3.7L12 11l9.2 6.4A2 2 0 0 1 20 21Z"/>`), g.Group(children),
	)
}

func Cocktail(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 6L6.6 2.8C6.3 2.4 5.6 2 5 2H2m16 4l-7 8l-7-8Z"/><path d="M15.4 9.1A4 4 0 1 0 14 6m-3 8v8m-4 0h8"/></g>`), g.Group(children),
	)
}

func Coconut(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="9" rx="10" ry="7"/><path d="M2 9v3a10 10 0 0 0 20 0V9"/><ellipse cx="12" cy="9" rx="6" ry="3"/><path d="m14 8l6-6h2"/></g>`), g.Group(children),
	)
}

func CoffeeBean(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.05 19.95a11.24 8.585 135 0 0 15.9-15.9a11.24 8.585 135 0 0-15.9 15.9"/><path d="M19.8 4.2C20 14 4 10 4.2 19.8"/></g>`), g.Group(children),
	)
}

func Coffeemaker(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 22V12a2 2 0 0 0-2-2H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v18H2"/><path d="M10 2v2a2 2 0 1 1-4 0V2m16 4h-4m4 4h-4m0 12v-6a2 2 0 0 1 2-2h2M7 10v2m0 10c-1.7 0-3-1.3-3-3v-3h6v3c0 1.7-1.3 3-3 3m-5-4a2 2 0 0 1 2-2"/></g>`), g.Group(children),
	)
}

func CoinsExchange(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 10V5c0-1.7 1.3-3 3-3h1"/><path d="m3 7l3 3l3-3"/><circle cx="18" cy="6" r="4"/><path d="M18 14v5c0 1.7-1.3 3-3 3h-1"/><path d="m21 17l-3-3l-3 3"/><circle cx="6" cy="18" r="4"/></g>`), g.Group(children),
	)
}

func CoinsStack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="6" rx="9" ry="3"/><path d="M3 10c0 1.7 4 3 9 3s9-1.3 9-3M3 14c0 1.7 4 3 9 3s9-1.3 9-3"/><path d="M3 6v12c0 1.7 4 3 9 3s9-1.3 9-3V6"/></g>`), g.Group(children),
	)
}

func CopyCode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 16a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2"/><rect width="14" height="14" x="8" y="8" rx="2"/><path d="m13 13l-1 2l1 2m4-4l1 2l-1 2"/></g>`), g.Group(children),
	)
}

func CopyDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 16a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2"/><rect width="14" height="14" x="8" y="8" rx="2"/><path d="M15 12v6m-3-3l3 3l3-3"/></g>`), g.Group(children),
	)
}

func CopyFilePath(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/><rect width="14" height="14" x="8" y="8" rx="2"/><path d="M12 18h.01M18 12l-2 6"/></g>`), g.Group(children),
	)
}

func CopyImage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 16a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2"/><rect width="14" height="14" x="8" y="8" rx="2"/><circle cx="14" cy="14" r="2"/><path d="m13.4 22l4.7-3.9c.8-.8 2-.8 2.8 0l1.1 1.1"/></g>`), g.Group(children),
	)
}

func CopyText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 16a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2"/><rect width="14" height="14" x="8" y="8" rx="2"/><path d="M12 13h6m-6 4h6"/></g>`), g.Group(children),
	)
}

func CopyType(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="14" x="8" y="8" rx="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2m-4 9v-1h6v1m-3-1v6m-1 0h2"/></g>`), g.Group(children),
	)
}

func CowHead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.8 15.1a10 10 0 0 0 .9-7.1h.3c1.7 0 3-1.3 3-3V3h-3c-1.3 0-2.4.8-2.8 1.9a10 10 0 0 0-8.4 0C7.4 3.8 6.3 3 5 3H2v2c0 1.7 1.3 3 3 3h.3a10 10 0 0 0 .9 7.1M9 9.5v.5m6-.5v.5"/><path d="M15 22a4 4 0 1 0-3-6.6A4 4 0 1 0 9 22Zm-6-4h.01M15 18h.01"/></g>`), g.Group(children),
	)
}

func CowUdderDroplets(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 2c0 2 .6 4 1.7 5.5L2.3 10a2 2 0 0 0 3.4 2l.9-1.6c1 .6 2.1 1.1 3.4 1.4V14a2 2 0 0 0 4 0v-2.2a8.5 8.5 0 0 0 3.4-1.4l.9 1.6a1.94 1.94 0 1 0 3.4-2l-1.4-2.5C21.4 6 22 4 22 2Zm5.9 16.6c-.6-.6-1.1-1.3-1.4-2.1c-.3.8-.8 1.5-1.4 2.1a1.93 1.93 0 1 0 2.8 0m11 0c-.6-.6-1.1-1.3-1.4-2.1c-.3.8-.8 1.5-1.4 2.1a1.93 1.93 0 1 0 2.8 0"/>`), g.Group(children),
	)
}

func Crab(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7.5 14A6 6 0 1 1 10 2.36L8 5l2 2S7 8 2 8m14.5 6A6 6 0 1 0 14 2.36L16 5l-2 2s3 1 8 1m-12 5v-2m4 2v-2"/><ellipse cx="12" cy="17.5" rx="7" ry="4.5"/><path d="M2 16c2 0 3 1 3 1m-3 5c0-1.7 1.3-3 3-3m14-2s1-1 3-1m-3 3c1.7 0 3 1.3 3 3"/></g>`), g.Group(children),
	)
}

func CricketBall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 8L3.6 6.6M8 11l1 1m3 3l1 1m3 3l1.4 1.4"/><circle cx="12" cy="12" r="10"/><path d="M8 5L6.6 3.6M11 8l1 1m3 3l1 1m4.4 4.4L19 16"/></g>`), g.Group(children),
	)
}

func CricketWicket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m6 2l4 2m4-1l4-1"/><circle cx="12" cy="13" r="2"/><path d="M6 7v15m7-15l-.3 4.1m-.2 3.8L12 22m6-15v15"/></g>`), g.Group(children),
	)
}

func CrossSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="3"/><path d="M14 10V7c0-.6-.4-1-1-1h-2c-.6 0-1 .4-1 1v3H7c-.6 0-1 .4-1 1v2c0 .6.4 1 1 1h3v3c0 .6.4 1 1 1h2c.6 0 1-.4 1-1v-3h3c.6 0 1-.4 1-1v-2c0-.6-.4-1-1-1Z"/></g>`), g.Group(children),
	)
}

func CrosshairPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16v5m0-18v5m4 4h5M3 12h5"/>`), g.Group(children),
	)
}

func CrosshairPlusDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16v5m0-18v5m4 4h5M3 12h5m4 0h.01"/>`), g.Group(children),
	)
}

func CrosshairSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3H5a2 2 0 0 0-2 2v3m9-5v5m9 0V5a2 2 0 0 0-2-2h-3m0 9h5m-5 9h3a2 2 0 0 0 2-2v-3m-9 0v5m-9-5v3a2 2 0 0 0 2 2h3m-5-9h5m4 0h.01"/>`), g.Group(children),
	)
}

func CrosshairTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="8"/><path d="M12 6V2m10 10h-4M6 12H2m10 10v-4"/></g>`), g.Group(children),
	)
}

func CrosshairTwoDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="8"/><path d="M12 6V2m10 10h-4M6 12H2m10 10v-4m0-6h.01"/></g>`), g.Group(children),
	)
}

func CupSaucer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 18a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4ZM6 8h12v6a4 4 0 0 1-4 4h-4a4 4 0 0 1-4-4Zm12 0h1a3 3 0 0 1 3 3v0a3 3 0 0 1-3 3h-1M6 4a1 1 0 0 1 1-1a1 1 0 0 0 1-1m4 2a1 1 0 0 1 1-1a1 1 0 0 0 1-1m4 2a1 1 0 0 1 1-1a1 1 0 0 0 1-1"/>`), g.Group(children),
	)
}

func CupToGo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7h16m-1.8 4l.8-4l-.8-4c-.1-.5-.6-1-1.2-1H7c-.6 0-1.1.4-1.2 1C5.5 4.4 5 7 5 7l.8 4M18 18H6l-1-7h14ZM7.2 18l.6 3c.1.5.6 1 1.2 1h6c.6 0 1.1-.4 1.2-1l.6-3"/>`), g.Group(children),
	)
}

func CurrencySquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m17 7l-2.17 2.17M17 17l-2.17-2.17M7 17l2.17-2.17M7 7l2.17 2.17"/><circle cx="12" cy="12" r="4"/></g>`), g.Group(children),
	)
}

func DeskLamp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 7l1-2m-3 2h4v4h-4zM6 7v4m3-4H3l1-5h4Zm13 15V12c0-.6-.4-1-1-1H3c-.6 0-1 .4-1 1v10m8-7H2m8-4v8m12 0H2"/>`), g.Group(children),
	)
}

func Diaper(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 9h4m16 0h-4M9 20a7 7 0 0 1-7-7V7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v6a7 7 0 0 1-7 7Z"/><path d="M2 13a7 7 0 0 1 7 7m6 0a7 7 0 0 1 7-7"/></g>`), g.Group(children),
	)
}

func Dishwasher(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 6v2M3 7h18"/><rect width="18" height="20" x="3" y="2" rx="2"/><path d="m9 11l-2 7"/><circle cx="14.5" cy="15.5" r="2.5"/><path d="m13 11l-2 7"/></g>`), g.Group(children),
	)
}

func DollarSignCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8m4 2V6"/></g>`), g.Group(children),
	)
}

func DollarSignSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 17V7m4 1h-6a2 2 0 0 0 0 4h4a2 2 0 0 1 0 4H8"/></g>`), g.Group(children),
	)
}

func DoorbellIntercom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.8 4a2.9 2.9 0 0 0-5.6 0H5c-.6 0-1 .4-1 1v14c0 .6.4 1 1 1h4.2a2.9 2.9 0 0 0 5.6 0H19c.6 0 1-.4 1-1V5c0-.6-.4-1-1-1ZM8 8h.01M12 8h.01M16 8h.01"/><circle cx="12" cy="14" r="2"/></g>`), g.Group(children),
	)
}

func Dress(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 2v3a5.14 5.14 0 0 1 .7 4.8l-.2.5a7.64 7.64 0 0 0 .4 6.3C17.7 17.9 19 20 19 20s-3.1 2-7 2s-7-2-7-2s1.3-2.1 2.1-3.5a7.64 7.64 0 0 0 .4-6.2l-.2-.5A5.66 5.66 0 0 1 8 5V2"/><path d="M16 5c-1.8 0-3.3 1-4 2.5C11.3 6 9.8 5 8 5"/></g>`), g.Group(children),
	)
}

func EggCup(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 11c0-3.3-2.7-9-6-9s-6 5.7-6 9m13 0a7 7 0 1 1-14 0Zm-7 7v4m-3 0h6"/>`), g.Group(children),
	)
}

func Elephant(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 12H14c-2.8 0-5-2.2-5-5V5a2 2 0 0 1 2-2h2c1.5 0 2.8.8 3.4 2H19c1.7 0 3 1.3 3 3v10m-4-8h.01"/><path d="M14 10a4 4 0 0 0 4 4a4 4 0 0 1 4 4a2 2 0 0 1-4 0m-8-2v5"/><path d="M18 14a4 4 0 0 0-4 4v3H6v-2.6c0-1.1-.8-2.3-1.7-3C2.9 14.3 2 12.8 2 11c0-3.3 3.1-6 7-6m-7 6v7"/></g>`), g.Group(children),
	)
}

func ElephantFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 10a4 4 0 0 0 4 4a2 2 0 0 1 0 4a7 7 0 0 1-2.8-.6c-.5-.2-.9 0-1 .6l-.1 1l-.9.9c-.4.4-.3.9.2 1.2c1.4.6 3 .9 4.6.9c3.3 0 6-2.7 6-6V8a4 4 0 0 0-4-4h-4.6c-.7-1.2-2-2-3.4-2H6C4.3 2 3 3.3 3 5v1a7 7 0 0 0 7 7h2.4m3.1-3H15"/>`), g.Group(children),
	)
}

func EscalatorArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="3" r="1"/><path d="M9 7v3m8-7c-.6 0-1.3.3-1.7.7L6 13H4a2 2 0 0 0 0 4h3c.6 0 1.3-.3 1.7-.7L18 7h2a2 2 0 0 0 0-4Zm5 10l-9 9m0-4v4h4"/></g>`), g.Group(children),
	)
}

func EscalatorArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="3" r="1"/><path d="M9 7v3m8-7c-.6 0-1.3.3-1.7.7L6 13H4a2 2 0 0 0 0 4h3c.6 0 1.3-.3 1.7-.7L18 7h2a2 2 0 0 0 0-4Zm5 10l-9 9m5-9h4v4"/></g>`), g.Group(children),
	)
}

func EuroCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M6 12h7m3-3a5 5 0 1 0 0 6"/></g>`), g.Group(children),
	)
}

func EuroSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 12h7m2-4a5.14 5.14 0 0 0-8 4a4.95 4.95 0 0 0 8 4"/></g>`), g.Group(children),
	)
}

func FaceAlien(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12a7.85 7.85 0 0 1-3.7 6.6l-4 2.7a3.9 3.9 0 0 1-4.5 0l-4-2.7A7.57 7.57 0 0 1 2 12a10 10 0 0 1 20 0"/><path d="M10.7 11.3c-1.4-1.3-3.3-1.7-4.2-.8s-.5 2.8.8 4.2c1.4 1.4 3.2 1.8 4.2.8c.9-.9.5-2.8-.8-4.2"/><path d="M17.5 10.5c-.9-.9-2.8-.5-4.2.8c-1.4 1.4-1.8 3.2-.8 4.2c.9.9 2.8.5 4.2-.8c1.3-1.4 1.7-3.3.8-4.2"/></g>`), g.Group(children),
	)
}

func FanHandheld(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4c0-.6.4-1 1-1a17.8 17.8 0 0 1 16.9 16.9c0 .6-.4 1-1 1.1H5c-1.1.1-2-.8-2-1.9Zm6.9.4L3 19M15.7 8.3L3.6 20.4m16-6.3L5 21"/>`), g.Group(children),
	)
}

func Farm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 14V4.5a2.5 2.5 0 0 0-5 0V14m5-6l6-5l8 6m-2-5v10m-8-4h4v4h-4zM2 14h20M2 22l5-8m0 8l5-8m10 8H12l5-8m-2 4h7"/>`), g.Group(children),
	)
}

func Faucet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.22 4.9L5.4 6H5a2 2 0 0 1 0-4h.4l4.86 1"/><circle cx="12" cy="4" r="2"/><path d="m13.78 4.9l4.8 1h.4a2 2 0 0 0 0-4h-.4l-4.92 1M12 6v3m6 1h4v6h-4zm4-1v8"/><path d="M18 11h-2.6a3.87 3.87 0 0 0-6.8 0H7c-2.8 0-5 2.2-5 5v1h4v-1c0-.6.4-1 1-1h1.6a3.87 3.87 0 0 0 6.8 0H18"/><path d="M3.5 17S2 19 2 20a2 2 0 0 0 4 0c0-1-1.5-3-1.5-3"/></g>`), g.Group(children),
	)
}

func FeatherPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 7h6M5 4v6m.1 7H14l8-8.2c-2.3-2.3-6.1-2.3-8.5 0L2.1 20M18 13H9.2"/>`), g.Group(children),
	)
}

func FeatherSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.3 3H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-8.8"/><path d="m7 12l8.5-8.5c2-2 4.5-2 6.5 0L16.5 9H10"/></g>`), g.Group(children),
	)
}

func FeatherText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.9 8H9.2m-4.1 4H14l8-8.2c-2.3-2.3-6.1-2.3-8.5 0L2 15m0 4h8m11-2v1c0 1 1 1.5 1 2.5c0 .8-.7 1.5-1.5 1.5h-5c-.8 0-1.5-.7-1.5-1.5c0-1 1-1.5 1-2.5v-1m-1 0h8"/>`), g.Group(children),
	)
}

func Flippers(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 17c0-4 2-7 2-13.5c0-.3-.2-.5-.5-.5C19 3 17 4 17 4s-2-1-4.5-1h-1C9 3 7 4 7 4S5 3 2.5 3c-.3 0-.5.2-.5.5C2 10 4 13 4 17"/><path d="M12 3v.5C12 10 10 13 10 17"/><rect width="6" height="7" x="4" y="14" rx="3"/><path d="M12 3.5C12 10 14 13 14 17"/><rect width="6" height="7" x="14" y="14" rx="3"/><path d="M7 4v6m10-6v6"/></g>`), g.Group(children),
	)
}

func FloorPlan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-5M9 3v7m12 0h-7M3 15h9"/>`), g.Group(children),
	)
}

func FloppyDisk(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 3v5h8"/><path d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2Z"/><path d="M17 21v-8H7v8"/></g>`), g.Group(children),
	)
}

func FloppyDiskRear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2Z"/><circle cx="12" cy="12" r="2"/><path d="M12 21v-3"/></g>`), g.Group(children),
	)
}

func FloppyDiskTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2Z"/><path d="M7 3h7v5H7z"/><circle cx="12" cy="14" r="2"/></g>`), g.Group(children),
	)
}

func FloppyDisks(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 4a2 2 0 0 1 2-2h10l4 4v10.2a2 2 0 0 1-2 1.8H8a2 2 0 0 1-2-2Z"/><path d="M10 2v4h6m2 12v-7h-8v7"/><path d="M18 22H4a2 2 0 0 1-2-2V6"/></g>`), g.Group(children),
	)
}

func FloppyDisksRear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 18a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10l4 4v10a2 2 0 0 1-2 2Z"/><circle cx="14" cy="10" r="2"/><path d="M14 18v-2m4 6H4a2 2 0 0 1-2-2V6"/></g>`), g.Group(children),
	)
}

func FloppyDisksTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 18a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h10l4 4v10a2 2 0 0 1-2 2Z"/><path d="M10 2h7v4h-7z"/><circle cx="14" cy="12" r="2"/><path d="M18 22H4a2 2 0 0 1-2-2V6"/></g>`), g.Group(children),
	)
}

func FlowerLotus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 20c0-5.5-4.5-10-10-10c0 5.5 4.5 10 10 10"/><path d="M9.7 8.3c-1.8-2-3.8-3.1-3.8-3.1s-.8 2.5-.5 5.4"/><path d="M15 12.9V12c0-4.4-3-8-3-8s-3 3.6-3 8v.9"/><path d="M18.6 10.6c.3-2.9-.5-5.4-.5-5.4s-2 1-3.8 3.1"/><path d="M12 20c5.5 0 10-4.5 10-10c-5.5 0-10 4.5-10 10"/></g>`), g.Group(children),
	)
}

func FlowerPot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 8h1m2-3v1m3 2h-1"/><circle cx="12" cy="8" r="2"/><path d="M12 11a3 3 0 1 1-3-3a3 3 0 1 1 3-3a3 3 0 1 1 3 3a3 3 0 1 1-3 3m0-1v8m3 0l-1 4h-4l-1-4m-1 0h8"/></g>`), g.Group(children),
	)
}

func FlowerRose(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 6a4 4 0 1 1-2-3.46"/><circle cx="12" cy="6" r="2"/><path d="M10 6a4 4 0 0 1 8 0v2A6 6 0 0 1 6 8V6m6 8v8m0 0c-4.2 0-7-1.667-7-5c4.2 0 7 1.667 7 5m0 0c4.2 0 7-1.667 7-5c-4.2 0-7 1.667-7 5"/></g>`), g.Group(children),
	)
}

func FlowerRoseSingle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 9.52a4.04 4.04 0 1 1 2-3.47"/><circle cx="17" cy="7.8" r="2"/><path d="m14 2.5l-2 1.3a6 6 0 1 0 6 10.4l2-1.2a4 4 0 0 0-4-6.95"/><path d="M9.77 12C4 15 2 22 2 22"/><path d="M13 20s-5 3-9.2-2c0 0 5.2-3 9.2 2"/></g>`), g.Group(children),
	)
}

func FlowerStem(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 8h1m2-3v1m3 2h-1"/><circle cx="12" cy="8" r="2"/><path d="M12 11a3 3 0 1 1-3-3a3 3 0 1 1 3-3a3 3 0 1 1 3 3a3 3 0 1 1-3 3m0-1v12m0 0c-4.2 0-7-1.667-7-5c4.2 0 7 1.667 7 5m0 0c4.2 0 7-1.667 7-5c-4.2 0-7 1.667-7 5"/></g>`), g.Group(children),
	)
}

func FlowerTulip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 7c-2-3.2-6-4-6-4v5c0 3.3 2.7 6 6 6v8"/><path d="M9.5 4.5C10 3 12 2 12 2s2 1 2.5 2.5"/><path d="M12 14c3.3 0 6-2.7 6-6V3c-6.2.9-10.8 11-6 11m0 8c-4.2 0-7-1.7-7-5c4.2 0 7 1.7 7 5m0 0c4.2 0 7-1.7 7-5c-4.2 0-7 1.7-7 5"/></g>`), g.Group(children),
	)
}

func Football(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 3c-.8-.8-3-1.2-5.8-.9s-6 1.6-8.8 4.4s-4 6-4.4 8.8s.1 5 .9 5.8s3 1.2 5.8.9s6-1.6 8.8-4.4s4-6 4.4-8.8s-.1-5-.9-5.8M6.4 17.6L9 15"/><path d="M8.7 21.9c-.8-3.3-3.4-5.8-6.7-6.7m6.1-1.3l2 2M11 11l2 2m.9-4.9l2 2m-.6-8c.8 3.3 3.4 5.8 6.6 6.6M15 9l2.6-2.6"/></g>`), g.Group(children),
	)
}

func FootballGoal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.7 2.3c-.2-.2-.9-.4-1.7-.3a4.6 4.6 0 0 0-3.7 5.7c.3.2.9.4 1.7.3a4.6 4.6 0 0 0 3.7-5.7"/><path d="M20 2v9c0 .6-.4 1-1 1H5c-.6 0-1-.4-1-1V2m10 14a4 4 0 0 0-4-4m3 4c-.6 0-1 .4-1 1v4c0 .6.4 1 1 1h2c.6 0 1-.4 1-1v-4c0-.6-.4-1-1-1Z"/></g>`), g.Group(children),
	)
}

func FootballHelmet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 14h.01M21.6 9c-1.3-4-5.1-7-9.6-7C6.5 2 2 6.5 2 12c0 2.6 1 5 3 7c1.4 1.3 3.6 1.4 4.9 0c.7-.7 1-1.6 1-2.5V13c0-1.7 1.3-3 3-3h6.8c.7 0 1-.4.9-1m.4 9H10.7"/><path d="M11 14h9a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2c-2.8 0-5-2.2-5-5v-3"/></g>`), g.Group(children),
	)
}

func ForkKnife(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 2v7c0 1.1.9 2 2 2h4a2 2 0 0 0 2-2V2M7 2v20m14-7V2v0a5 5 0 0 0-5 5v6c0 1.1.9 2 2 2zm0 0v7"/>`), g.Group(children),
	)
}

func ForkKnifeCrossed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 2l-2.3 2.3a3 3 0 0 0 0 4.2l1.8 1.8a3 3 0 0 0 4.2 0L22 8m-7 7L3.3 3.3a4.2 4.2 0 0 0 0 6l7.3 7.3c.7.7 2 .7 2.8 0zm0 0l7 7m-19.9-.2l6.4-6.3M19 5l-7 7"/>`), g.Group(children),
	)
}

func FoxFaceTail(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.9 8.3C20.6 7 21 5.6 21 4c0-.6-.4-1-1-1c-2.3 0-4.3.8-5.9 2.2a15 15 0 0 0-4.2 0A8.78 8.78 0 0 0 4 3c-.6 0-1 .4-1 1c0 1.6.4 3 1.1 4.3c-.6.7-1.1 1.4-1.4 2.2C4 13 11 16 12 16s8-3 9.3-5.5c-.3-.8-.8-1.5-1.4-2.2M9 9v.5m4 3.5h-2m1 3v-3m3-4v.5"/><path d="M6.3 20.5A6.87 6.87 0 0 0 9 15H2.2c.8 4 4.9 7 9.8 7c5.5 0 10-3.8 10-8.5c0-1.1-.2-2.1-.7-3"/></g>`), g.Group(children),
	)
}

func FrogFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 7h.01"/><circle cx="6" cy="7" r="4"/><path d="M14.4 5.3a10 10 0 0 0-4.8 0"/><circle cx="18" cy="7" r="4"/><path d="M18 7h.01M22 13.5C22 16 17.5 18 12 18S2 16 2 13.5m8 .5h.01M14 14h.01"/><path d="M3.1 9.75A7 7 0 0 0 2 13.5C2 18.2 6.5 22 12 22s10-3.8 10-8.5a7 7 0 0 0-1.1-3.75"/></g>`), g.Group(children),
	)
}

func Fruit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m10 10l4-3m-4 0l4 3"/></g>`), g.Group(children),
	)
}

func Garlic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 4a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v1c0 1-.5 2-1.4 2.5L5.1 9.1A7 7 0 0 0 9 22h6a7 7 0 0 0 3.8-12.8l-2.5-1.6A3.32 3.32 0 0 1 15 5Z"/><path d="M9 5c0 4-2 4-2 9c0 4.4 2.2 8 5 8s5-3.6 5-8c0-5-2-5-2-9"/></g>`), g.Group(children),
	)
}

func Gearbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4v16m7-16v16m7-16v8H5"/>`), g.Group(children),
	)
}

func GearboxSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 7v10m5-10v10m5-10v5H7"/></g>`), g.Group(children),
	)
}

func GemRing(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13.2 8.1L16 4.4L14.4 2H9.6L8 4.4l2.8 3.7"/><circle cx="12" cy="15" r="7"/></g>`), g.Group(children),
	)
}

func GlassesSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2.2 13.1l3.1-6.4C5.8 5.7 6.8 5 8 5"/><rect width="8" height="6" x="2" y="12" rx="2"/><path d="M14 15a2 2 0 0 0-4 0"/><rect width="8" height="6" x="14" y="12" rx="2"/><path d="m21.8 13.1l-3.1-6.4C18.2 5.7 17.2 5 16 5"/></g>`), g.Group(children),
	)
}

func GlassesSun(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2.16 14.2l3.1-7.4C5.7 5.8 6.8 5 8 5"/><path d="M4 13a2 2 0 0 0-2 2v1c0 1.7 1.3 3 3 3h1c3.3 0 6-2.7 6-6c0 3.3 2.7 6 6 6h1c1.7 0 3-1.3 3-3v-1a2 2 0 0 0-2-2Z"/><path d="m21.83 14.2l-3.1-7.4C18.3 5.8 17.2 5 16 5"/></g>`), g.Group(children),
	)
}

func GoalNet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v14M8 8v12m4-12v12m4-12v12M6 10h12M6 14h12M6 18h12"/>`), g.Group(children),
	)
}

func Goblet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 2c-1.78 2.72-3 6.65-3 9a7 7 0 1 0 14 0c0-2.35-1.22-6.28-3-9Zm4 16v4m-4 0h8"/><path d="M5 11c.84-.5 1.68-1 3.5-1c3.5 0 3.5 2 7 2c1.82 0 2.66-.5 3.5-1"/></g>`), g.Group(children),
	)
}

func GobletCrack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 2c-1.78 2.72-3 6.65-3 9a7 7 0 1 0 14 0c0-2.35-1.22-6.28-3-9Zm4 16v4m-4 0h8"/><path d="m13 11l-1-1l2-2l-3-3l3-3"/></g>`), g.Group(children),
	)
}

func GolfDriver(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="9" r="2"/><path d="M6 11v2M22 2l-9.3 14.1c-.4.6-1 .9-1.7.9H4a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h2c1.6 0 3.1-.7 4.1-2.1l2.6-3.8"/></g>`), g.Group(children),
	)
}

func Grapes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 5V2l-5.89 5.89"/><circle cx="16.6" cy="15.89" r="3"/><circle cx="8.11" cy="7.4" r="3"/><circle cx="12.35" cy="11.65" r="3"/><circle cx="13.91" cy="5.85" r="3"/><circle cx="18.15" cy="10.09" r="3"/><circle cx="6.56" cy="13.2" r="3"/><circle cx="10.8" cy="17.44" r="3"/><circle cx="5" cy="19" r="3"/></g>`), g.Group(children),
	)
}

func GridLines(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3v18m6-18v18M3 9h18M3 15h18"/>`), g.Group(children),
	)
}

func GridLinesOffset(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 3v18m6-18v18M3 11h18M3 17h18"/>`), g.Group(children),
	)
}

func Hairdryer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="8" cy="8" r="2"/><path d="M18 11s-7 3-10 3A6 6 0 0 1 8 2c3 0 10 3 10 3Zm0-6l4-2v10l-4-2"/><path d="m7 13.9l.8 5.1c.1.5.6 1 1.2 1h2c.6 0 .9-.4.8-1l-.9-5.5m.74 4.5s3.3-2 7.3-2a2 2 0 0 1 0 4H17a2 2 0 0 0-2 2"/></g>`), g.Group(children),
	)
}

func HatBaseball(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 3v1m0 10c2.8 0 5.5.3 8 .9V12a8 8 0 0 0-16 0v2.9c2.5-.6 5.2-.9 8-.9"/><path d="M9 14.1V10h6v4.1"/><path d="M2.3 18A2 2 0 0 0 4 21h.4l1.6-.4a26.44 26.44 0 0 1 12 0l1.6.4h.4a2 2 0 0 0 1.7-3l-1.8-3.2a39.9 39.9 0 0 0-15.8 0Z"/></g>`), g.Group(children),
	)
}

func HatBeanie(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.4 6.2C6.7 6.9 4 10.1 4 14v1"/><circle cx="12" cy="5" r="2"/><path d="M20 15v-1c0-3.9-2.7-7.1-6.4-7.8"/><rect width="20" height="5" x="2" y="15" rx="1"/><path d="M6 15v5m4-5v5m4-5v5m4-5v5"/></g>`), g.Group(children),
	)
}

func HatBowler(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 13c0 1.7 2.7 3 6 3s6-1.3 6-3v-3a6 6 0 0 0-12 0Z"/><path d="M6 9c0 1.7 2.7 3 6 3s6-1.3 6-3"/><path d="M6 9.2C3.6 10.3 2 12 2 14c0 3.3 4.5 6 10 6s10-2.7 10-6c0-2-1.6-3.7-4-4.8"/></g>`), g.Group(children),
	)
}

func HatChef(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 13.85A4 4 0 0 1 7.4 6a5 5 0 0 1 9.2 0a4 4 0 0 1 1.4 7.85V21H6ZM6 17h12"/>`), g.Group(children),
	)
}

func HatHard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 6.3c-3.4.9-6 4-6 7.7v2m6-6V5c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v5m6 6v-2c0-3.7-2.6-6.8-6-7.7"/><rect width="20" height="4" x="2" y="16" rx="1"/></g>`), g.Group(children),
	)
}

func HatTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="5" rx="7" ry="3"/><path d="M5 5c0 1 1 4 1 6v4c0 1.7 2.7 3 6 3s6-1.3 6-3v-4c0-2 1-5 1-6"/><path d="M18 11c0 1.7-2.7 3-6 3s-6-1.3-6-3"/><path d="M6 11.2C3.6 12.3 2 14 2 16c0 3.3 4.5 6 10 6s10-2.7 10-6c0-2-1.6-3.7-4-4.8"/></g>`), g.Group(children),
	)
}

func HeadingCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9 16V8m0 4h6m0 4V8"/></g>`), g.Group(children),
	)
}

func HeadingSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 17V7m0 5h8m0 5V7"/></g>`), g.Group(children),
	)
}

func Hedgehog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 11L3 7.7L5.7 7l.1-2.8l2.7.6l1.3-2.6L12 4l2.2-1.8l1.3 2.6l2.7-.6l.1 2.8l2.7.7l-1.2 2.5L22 12l-2.2 1.8l1.2 2.5l-3 .7m-8 0h.01"/><path d="M3 16c2.8 0 5-2.2 5-5c3.3 0 6 2.7 6 6a4 4 0 0 0 4 4h-8c-1.1 0-2.6-.6-3.4-1.4zv-1"/></g>`), g.Group(children),
	)
}

func HelmetDiving(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 17.9c1.5-.9 2.7-2.2 3.4-3.9h.8c.4 0 .8-.4.8-1V9c0-.6-.4-1-.8-1h-.8A7.92 7.92 0 0 0 15 3.6v-.8c0-.4-.4-.8-1-.8h-4c-.6 0-1 .4-1 .8v.8A7.92 7.92 0 0 0 4.6 8h-.8c-.4 0-.8.4-.8 1v4c0 .6.4 1 .8 1h.8c.7 1.7 1.9 3 3.4 3.9"/><circle cx="12" cy="11" r="4"/><path d="M8 11h8m-4-4v8m-5.3 2c-1 .6-1.7 1.2-1.7 2c0 1.7 3.1 3 7 3s7-1.3 7-3c0-.8-.7-1.4-1.7-2"/></g>`), g.Group(children),
	)
}

func HexagonsSeven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.3 4.3v3.9L2 10.1v3.8l3.3 1.9v3.9l3.4 1.9l3.3-1.9l3.3 1.9l3.4-1.9v-3.9l3.3-1.9v-3.8l-3.3-1.9V4.3l-3.4-1.9L12 4.3L8.7 2.4ZM12 8.2V4.3m6.7 3.9l-3.4 1.9m0 3.8l3.4 1.9M12 19.7v-3.9m-3.3-1.9l-3.4 1.9m0-7.6l3.4 1.9"/><path d="m8.7 13.9l3.3 1.9l3.3-1.9v-3.8L12 8.2l-3.3 1.9Z"/></g>`), g.Group(children),
	)
}

func HexagonsThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 4.4a2 2 0 0 0-1 1.73v4.37l-4 2.4a2 2 0 0 0-1 1.73v3.27a2 2 0 0 0 .97 1.68L6 21.4a2 2 0 0 0 2.03-.02L12 19l4 2.4a2 2 0 0 0 2.03-.02L21 19.6a2 2 0 0 0 1-1.73V14.6a2 2 0 0 0-.97-1.68L17 10.5V6.1a2 2 0 0 0-.97-1.68L13 2.6a2 2 0 0 0-2.03.02Z"/><path d="m7 10.5l5 3l5-3m-5 3V19"/></g>`), g.Group(children),
	)
}

func HighHeel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 3c6 6 8.4 10.5 9.8 12c.9 1 2.5 1.3 3.7.6c.3-.2.5-.3.7-.6c.6.3 3.8 3.1 3.8 5c0 1-1 1-1 1h-7c-1 0-2-.5-2.6-1.5L10.1 17c-.9-1.6-2.2-3-3.7-4.2L4 11a5 5 0 0 1 0-8"/><path d="m2.56 9.3l.6 1.1C4.2 12.6 5 16.5 5 21"/></g>`), g.Group(children),
	)
}

func Hockey(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="17" cy="19" r="3"/><path d="M2.8 13a5.95 5.95 0 1 0 10.4 6l8.5-14a1.94 1.94 0 1 0-3.4-2L9.7 17a1.88 1.88 0 1 1-3.4-2a1.94 1.94 0 1 0-3.5-2m17.8-6.2l-3.3-2.1m-2.1 3.4l3.3 2.1"/></g>`), g.Group(children),
	)
}

func HockeyMask(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12a10 10 0 1 0 20 0c0-4.1-.4-6.6-1.9-8.1S16.1 2 12 2s-6.6.4-8.1 1.9S2 7.9 2 12m10-6h.01"/><circle cx="8" cy="10.5" r="2"/><circle cx="16" cy="10.5" r="2"/><path d="M8.5 17h.01M12 15h.01m3.49 2h.01"/></g>`), g.Group(children),
	)
}

func HorseHead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11.5 12H11m-6 3a4 4 0 0 0 4 4h7.8l.3.3a3 3 0 0 0 4-4.46L12 7c0-3-1-5-1-5S8 3 8 7c-4 1-6 3-6 3"/><path d="M6.14 17.8S4 19 2 22"/></g>`), g.Group(children),
	)
}

func HotDog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.1 3.5a4 4 0 0 0-5.9-.3l-8 8a4 4 0 0 0 .2 5.9m3.5 3.6a4.07 4.07 0 0 0 5.9.1l8-8a4 4 0 0 0-.1-5.9"/><path d="M21.3 6.3a2.5 2.5 0 0 0-3.5-3.5l-15 15a2.5 2.5 0 0 0 3.5 3.5Z"/></g>`), g.Group(children),
	)
}

func House(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 9l9-7l9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/><path d="M9 22V12h6v10"/></g>`), g.Group(children),
	)
}

func HouseManor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 6V2H5v4m14 0V2h-4v4"/><rect width="20" height="16" x="2" y="6" rx="2"/><path d="M2 12h4m0 10V12l5.5-6m1 0l5.5 6v10m0-10h4m-10-1h.01M10 22v-5a2 2 0 1 1 4 0v5"/></g>`), g.Group(children),
	)
}

func HouseOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 15.3V9l-9-7l-2.4 1.9M2 2l20 20M6.4 6.4L3 9v11a2 2 0 0 0 2 2h14a2 2 0 0 0 1.8-1.2"/><path d="M12 12H9v10m6 0v-7"/></g>`), g.Group(children),
	)
}

func HouseRoof(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2.6 10.4a2.12 2.12 0 1 0 3.02 2.98L12 7l6.4 6.4a2.12 2.12 0 1 0 2.979-3.021L13.7 2.7a2.4 2.4 0 0 0-3.404.004Z"/><path d="M20 14v6a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-6"/><path d="M14 22v-6a2 2 0 0 0-4 0v6"/></g>`), g.Group(children),
	)
}

func HouseRoofOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 6.8l6.6 6.6a2 2 0 0 0 2.8-2.8l-8-8c-.8-.8-2-.8-2.8 0L9.2 4M2 2l20 20M6.6 6.6l-4 4a2 2 0 0 0 2.8 2.8l4-4M14 22v-6a2 2 0 0 0-4 0v6"/><path d="M4 14v6a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2"/></g>`), g.Group(children),
	)
}

func Houses(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 17H3c-.6 0-1-.4-1-1V8.5L8 4l10 7.5V19c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1v-7.5L16 4l6 4.5V16c0 .6-.4 1-1 1h-3"/><path d="M10 20v-6h4v6"/></g>`), g.Group(children),
	)
}

func IceHockey(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 4v4c0 1.1-1.8 2-4 2s-4-.9-4-2V4"/><ellipse cx="6" cy="4" rx="4" ry="2"/><path d="M4 17a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h4a6 6 0 0 0 5.2-3l8.5-14a1.94 1.94 0 1 0-3.4-2l-7.9 13c-.4.6-1 1-1.7 1ZM20.6 6.8l-3.3-2.1m-2.1 3.4l3.3 2.1M6 17v5"/></g>`), g.Group(children),
	)
}

func IceSkate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 2v9m0-4L8 8m3-5L4 5v11a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2a3.08 3.08 0 0 0-1.8-2.8L11 11l-3 1m-1 6v4m8-4v4M4 22h12c2.1 0 3.9-1.1 5-2.7"/>`), g.Group(children),
	)
}

func Igloo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 20.2c2.4-.7 4-1.9 4-3.2v-5a10 10 0 1 0-20 0v5c0 1.3 1.6 2.5 4 3.2"/><path d="M6.5 3.65C7.5 5 9.6 6 12 6s4.5-1 5.5-2.35"/><path d="M10.1 5.8c-1 .9-1.8 2.6-2 4.6m7.8 0c-.3-2-1-3.6-2-4.6"/><path d="M3.3 7.1C5.3 9.5 8.5 11 12 11s6.7-1.5 8.7-3.9M2 12c.9 1.2 2.4 2.4 4.3 3.1"/><path d="M6 21c0 .6.4 1 1 1h10c.6 0 1-.4 1-1v-4a6 6 0 1 0-12 0Z"/><path d="M17.7 15.1c1.9-.7 3.4-1.9 4.3-3.1M10 22v-5a2 2 0 1 1 4 0v5"/></g>`), g.Group(children),
	)
}

func IndianRupeeCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9 11h6.5m0-4H9a4 4 0 0 1 0 8l3 3"/></g>`), g.Group(children),
	)
}

func IndianRupeeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 7h8m-8 4h8m-4 6l-4-2h1a4 4 0 0 0 0-8"/></g>`), g.Group(children),
	)
}

func Intercom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="2"/><path d="M6 9v6m4-9v12m4-12v12m4-9v6"/></g>`), g.Group(children),
	)
}

func Iron(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h.01M6 11h.01M10 11h.01M6 15h.01M10 15h.01M14 19v-7C14 6 9 2 8 2S2 6 2 12v7h14a2 2 0 0 0 2-2V8a2 2 0 0 1 4 0v9M3 22h10"/>`), g.Group(children),
	)
}

func IronOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.9 7.3C11.4 4 8.7 2 8 2m14 14.3V8a2 2 0 0 0-4 0v4.3M2 2l20 20M6 11h.01M10 11h.01M6 15h.01M10 15h.01M4.7 4.7C3.3 6.4 2 9 2 12v7h12v-5M3 22h10"/>`), g.Group(children),
	)
}

func IroningBoard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 3a4 4 0 0 0 0 8h14a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2Zm0 18l12-10M6 11l12 10"/>`), g.Group(children),
	)
}

func ItalicSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M16 7h-6m3.5 0l-4 10m3.5 0H7"/></g>`), g.Group(children),
	)
}

func Jacket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 4c0 1.1 1.8 2 4 2s4-.9 4-2V3c0-.6-.4-1-1-1H9c-.6 0-1 .4-1 1Z"/><path d="M8 4c0 2 4 5 4 10v8m0-8c0-5 4-8 4-10M6 19H3c-.6 0-1-.4-1-1V7c0-1.1.8-2.3 1.9-2.6L8 3"/><path d="M18 9v12c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1V9"/><path d="m16 3l4.1 1.4C21.2 4.7 22 5.9 22 7v11c0 .6-.4 1-1 1h-3M6 15l2-2m10 2l-2-2"/></g>`), g.Group(children),
	)
}

func JacketSports(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 4c0 1.1 1.8 2 4 2s4-.9 4-2V3c0-.6-.4-1-1-1H9c-.6 0-1 .4-1 1Z"/><path d="M8 4c0 2 4 5 4 10v8m0-8c0-5 4-8 4-10M6 19H3c-.6 0-1-.4-1-1V7c0-1.1.8-2.3 1.9-2.6L8 3"/><path d="M18 9v12c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1V9"/><path d="m16 3l4.1 1.4C21.2 4.7 22 5.9 22 7v11c0 .6-.4 1-1 1h-3M2 15h4l2-2"/><path d="M22 15h-4l-2-2M6 18h12"/></g>`), g.Group(children),
	)
}

func JapaneseYenCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m9 7l3 3v8m0-8l3-3m-6 4h6m-6 4h6"/></g>`), g.Group(children),
	)
}

func JapaneseYenSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m9 7l3 3v7m0-7l3-3m-6 4h6m-6 4h6"/></g>`), g.Group(children),
	)
}

func Jar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 3h16M5 3v1.6c0 .8-.2 1.6-.7 2.2l-.7 1C3.2 8.4 3 9.2 3 10v8c0 1.7 1.3 3 3 3h12c1.7 0 3-1.3 3-3v-8c0-.8-.2-1.6-.7-2.2l-.7-1c-.4-.7-.6-1.4-.6-2.2V3M3 13h4"/><rect width="10" height="7" x="7" y="10" rx="1"/><path d="M17 13h4"/></g>`), g.Group(children),
	)
}

func Jug(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m19 3l-2 5H5L3 3zm0 0c1.7 0 3 1.3 3 3v4M2 16c1.08-.5 2.16-1 4.5-1c4.5 0 4.5 2 9 2c2.34 0 3.42-.5 4.5-1"/><path d="M15 21a5 5 0 0 0 4.48-7.22L17 8H5l-2.5 5.8A5 5 0 0 0 7 21Z"/></g>`), g.Group(children),
	)
}

func Kebab(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 12l4.2-4.2c.4-.4.4-1 .1-1.5a2.9 2.9 0 1 1 4.8.8"/><path d="M15.3 11.3c.9.9.9 2.5 0 3.4l-1.6 1.6c-.9.9-2.5.9-3.4 0c.9.9.9 2.5 0 3.4l-1.6 1.6c-.9.9-2.5.9-3.4 0l-2.6-2.6c-.9-.9-.9-2.5 0-3.4l1.6-1.6c.9-.9 2.5-.9 3.4 0c-.9-.9-.9-2.5 0-3.4l1.6-1.6c.9-.9 2.5-.9 3.4 0Zm-5 5l-2.6-2.6M9 15l-2 2m-5 5l2-2"/></g>`), g.Group(children),
	)
}

func Kettle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 6v1M2 22h16M3 18c-.6 0-1-.4-1-1v-2a8 8 0 0 1 16 0v2c0 .6-.4 1-1 1Z"/><path d="M5 8.8V7a5 5 0 0 1 10 0v1.8"/><path d="M18 14.5A9.06 9.06 0 0 0 22 7l-3-1c-1 2-3.5 5-9 5c-2.5 0-4.4-.6-5.8-1.5"/></g>`), g.Group(children),
	)
}

func KettleElectric(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 14v-4c0-1.7 1.3-3 3-3h16v2c0 2-1.5 3.7-3.5 3.9M11 7v7"/><path d="M12 2C9.2 2 7 4.2 7 7l-.8 9c-.1 1.1.7 2 1.8 2h8c1.1 0 1.9-.9 1.8-2a1607 1607 0 0 1-.8-9c0-2.8-2.2-5-5-5M6 22h12"/></g>`), g.Group(children),
	)
}

func Kiwi(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v1m3 2l1-1m1 4h1m-3 3l1 1m-4 1v1m-4-2l1-1m-3-3h1m1-4l1 1"/></g>`), g.Group(children),
	)
}

func LayoutGridMoveHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="14" y="3" rx="1"/><path d="m7 14l-4 4l4 4m14-4H3m14-4l4 4l-4 4"/></g>`), g.Group(children),
	)
}

func LayoutGridMoveVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="7" x="14" y="3" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/><path d="m2 7l4-4l4 4M6 3v18m-4-4l4 4l4-4"/></g>`), g.Group(children),
	)
}

func LayoutGridPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/><path d="M17.5 3v7M14 6.5h7"/></g>`), g.Group(children),
	)
}

func LayoutListMove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 7L6 3L2 7m4-4v18m-4-4l4 4l4-4m4-13h7m-7 5h7m-7 6h7m-7 5h7"/>`), g.Group(children),
	)
}

func Lemon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.6 2.4c-.5.2-1.3.6-1.8.6H12a9 9 0 0 0-9 9v3.8c0 .6-.4 1.3-.6 1.8a2.95 2.95 0 0 0 4 4c.5-.2 1.3-.6 1.8-.6H12a9 9 0 0 0 9-9V8.2c0-.6.4-1.3.6-1.8a2.95 2.95 0 0 0-4-4"/><path d="M7 12c0-2.8 2.2-5 5-5"/></g>`), g.Group(children),
	)
}

func LifeJacket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14a2.5 2.5 0 0 0-.8-1.9a3.5 3.5 0 1 1 5.6 0l-.3.4A2.5 2.5 0 0 0 14 14v5a3 3 0 1 0 6 0v-9a8 8 0 0 0-16 0v9a3 3 0 1 0 6 0Z"/>`), g.Group(children),
	)
}

func LigatureSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 12h3m4.4-4a3.5 3.5 0 0 0-5.9 2.5V17m1.5 0H7m7-5h1.5v5m1.5 0h-3"/></g>`), g.Group(children),
	)
}

func LightSwitch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="2"/><path d="M10 8h4v8h-4zm0 4h4"/></g>`), g.Group(children),
	)
}

func Lingerie(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 2v2a2 2 0 0 0-2 2v2c0 1.7 1.3 3 3 3h2a2 2 0 0 0 2-2h4a2 2 0 0 0 2 2h2c1.7 0 3-1.3 3-3V6a2 2 0 0 0-2-2"/><path d="M10 9c0-2.8-2.2-5-5-5m14-2v2c-2.8 0-5 2.2-5 5M3 15a7 7 0 0 1 7 7h4a7 7 0 0 1 7-7Z"/></g>`), g.Group(children),
	)
}

func LocateSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="14" x="5" y="5" rx="2"/><path d="M12 5V2m7 10h3M12 22v-3M2 12h3"/></g>`), g.Group(children),
	)
}

func LuggageCabin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 2h8m-6 0v5m4-5v5"/><rect width="12" height="14" x="6" y="7" rx="2"/><path d="M14 21v-8a2 2 0 1 0-4 0v8m-2 0v1m8-1v1"/></g>`), g.Group(children),
	)
}

func LunchBox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 7V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/><path d="M8 21a2 2 0 0 0 2-2v-8a4 4 0 0 0-8 0v8a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-8a4 4 0 0 0-4-4H6m-4 6h20"/></g>`), g.Group(children),
	)
}

func MailboxFlag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 5.5A4 4 0 0 1 22 9v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9a4 4 0 0 1 8 0v8a2 2 0 0 1-2 2M6 5h4"/><path d="M14 9V5h2v1h-2"/></g>`), g.Group(children),
	)
}

func MaskSnorkel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13.5 14a2 2 0 0 1-1.4-.6l-.7-.8c-.8-.8-2-.8-2.8 0l-.7.8a2 2 0 0 1-1.4.6H6a4 4 0 0 1 0-8h8a4 4 0 0 1 0 8ZM12 18a2 2 0 0 1-4 0"/><path d="M10 20a2 2 0 0 0 2 2h4c3.3 0 6-2.7 6-6V2h-4v14a2 2 0 0 1-2 2m2-8h4"/><circle cx="4.5" cy="21.5" r=".5"/><path d="M3 17.5h.01"/></g>`), g.Group(children),
	)
}

func MealBox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="6"/><path d="M22 12c0 3.3-2.7 6-6 6H8c-3.3 0-6-2.7-6-6m5-6h10"/><rect width="4" height="4" x="6" y="10" rx="1"/><rect width="4" height="4" x="14" y="10" rx="1"/></g>`), g.Group(children),
	)
}

func MortarPestle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 4a2 2 0 0 0-3-1.7l-8.5 5.1A3 3 0 0 0 12 13c.8 0 1.5-.3 2-.8l7.3-6.7c.4-.4.7-.9.7-1.5"/><path d="M22 12a10 10 0 0 1-20 0"/><path d="M11.1 7C6 7.2 2 9.4 2 12c0 2.8 4.5 5 10 5s10-2.2 10-5c0-1.5-1.4-2.9-3.6-3.8"/></g>`), g.Group(children),
	)
}

func MotorRacingHelmet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12.2a10 10 0 1 0-19.4 3.2c.2.5.8 1.1 1.3 1.3l13.2 5.1c.5.2 1.2 0 1.6-.3l2.6-2.6c.4-.4.7-1.2.7-1.7Z"/><path d="m21.8 18l-10.5-4a2 2.06 0 0 1 .7-4h9.8"/></g>`), g.Group(children),
	)
}

func Mug(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8h1a4 4 0 1 1 0 8h-1M3 8h14v9a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4Zm1-4a1 1 0 0 1 1-1a1 1 0 0 0 1-1m4 2a1 1 0 0 1 1-1a1 1 0 0 0 1-1m4 2a1 1 0 0 1 1-1a1 1 0 0 0 1-1"/>`), g.Group(children),
	)
}

func MugTeabag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 8h1a4 4 0 1 1 0 8h-1M3 8h14v9a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4Zm1-4a1 1 0 0 1 1-1a1 1 0 0 0 1-1m4 2a1 1 0 0 1 1-1a1 1 0 0 0 1-1m4 2a1 1 0 0 1 1-1a1 1 0 0 0 1-1M9 8v3"/><path d="M11 16v-3.5L9 11l-2 1.5V16Z"/></g>`), g.Group(children),
	)
}

func Mustache(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.2 8.6a3.9 3.9 0 0 0-6.2-.2a3.75 3.75 0 0 0-6.2.2l-.6.8C4.5 10.4 3.3 11 2 11a5.55 5.55 0 0 0 10 3.2A5.45 5.45 0 0 0 22 11c-1.3 0-2.5-.6-3.2-1.6Z"/>`), g.Group(children),
	)
}

func Olive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 10l4-3m-4 0l4 3"/><ellipse cx="12" cy="12" rx="9" ry="10"/><path d="m2 22l5-5M18.69 5.31L22 2"/></g>`), g.Group(children),
	)
}

func Onion(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="2"/><circle cx="12" cy="12" r="6"/><path d="M2.8 8.1a10 10 0 1 0 5.3-5.3C5 4 3 2 3 2L2 3s2 2 .8 5.1M18 20v2m3-1l-1.9-1.9M22 18h-2"/></g>`), g.Group(children),
	)
}

func Owl(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="9" rx="8" ry="7"/><path d="M12 9a4 4 0 1 1 8 0v12h-4C9.4 21 4 15.6 4 9a4 4 0 1 1 8 0v1M8 9h.01M16 9h.01"/><path d="M20 21a3.9 3.9 0 1 1 0-7.8m-10 6.2V22m4-1.15V22"/></g>`), g.Group(children),
	)
}

func PacMan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 12l7.4 6.7a10 10 0 1 1 0-13.4Zm6 0h.01M22 12h.01"/>`), g.Group(children),
	)
}

func PacManGhost(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 10h.01M15 10h.01M12 2a8 8 0 0 0-8 8v12l3-3l2.5 2.5L12 19l2.5 2.5L17 19l3 3V10a8 8 0 0 0-8-8"/>`), g.Group(children),
	)
}

func PalmtreeIslandSun(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="7" r="3"/><path d="M16 14s1-3 1-8V4s-1-2-3-2c-1 0-2 .5-2 .5"/><path d="M13 8a4 4 0 0 1 8 0"/><path d="M17 4s1-2 3-2c1 0 2 .5 2 .5M19.75 19A8 8 0 0 0 4 21"/><path d="M2 20c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c2.6 0 2.4 2 5 2c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1"/></g>`), g.Group(children),
	)
}

func Pancakes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 11.6c3.5-.8 6-2.5 6-4.6c0-2.8-4.5-5-10-5S2 4.2 2 7c0 2.5 3.7 4.6 8.4 5"/><path d="M3.3 9.5C2.5 10.2 2 11.1 2 12c0 2.8 4.5 5 10 5h.3m3.6-.4c3.6-.8 6.1-2.5 6.1-4.6c0-.9-.5-1.8-1.3-2.5"/><path d="M3.3 14.5C2.5 15.2 2 16.1 2 17c0 2.8 4.5 5 10 5s10-2.2 10-5c0-.9-.5-1.8-1.3-2.5"/><path d="M16 16a2 2 0 0 1-4 0v-2c0-1.1-.9-2-2-2.2c-1.8-.5-3-1.6-3-2.8c0-1.7 2.2-3 5-3s5 1.3 5 3c0 .4-.1.7-.3 1.1c-.3.5-.7 1.2-.7 1.7Z"/></g>`), g.Group(children),
	)
}

func Peace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 2v20m7.1-2.9L12 12l-7 7"/></g>`), g.Group(children),
	)
}

func Peach(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 2a2 2 0 0 0-2 2v2"/><path d="M12 6.5A6 6 0 0 1 22 11c0 6.1-4.5 11-10 11S2 17.1 2 11a6 6 0 0 1 12 0"/></g>`), g.Group(children),
	)
}

func Pear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 7a4.95 4.95 0 0 0-8.6-3.4c-1.5 1.6-1.6 1.8-5 2.6a8 8 0 1 0 9.4 9.5c.7-3.4 1-3.6 2.6-5c1-1 1.6-2.3 1.6-3.7m-3-2l3-3"/>`), g.Group(children),
	)
}

func Penguin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.9 15a5.87 5.87 0 0 0-1.7-3.3l-.2-.2c-.6-.6-1-1.5-1-2.5V5a3 3 0 1 0-6 0v4a3.74 3.74 0 0 1-1.2 2.8a6.2 6.2 0 0 0-1.7 3.3"/><path d="M9 10c-2 4-4-1-7 2m7-3.1c3-1.9 6 0 6 0s-2 3.1-3 4c-1-.9-3-4-3-4m6 1.1c2 4 4-1 7 2M2 19c0-1 1-1 1-2c0-.6.4-1 1-1c1 0 1-1 2-1c.4 0 .7.2.9.5L8.8 19a2 2 0 0 1-2.7 2.7l-3.5-1.9c-.4-.1-.6-.4-.6-.8"/><path d="M8.7 21a6.07 6.07 0 0 0 6.6 0"/><path d="M22 19c0-1-1-1-1-2c0-.6-.4-1-1-1c-1 0-1-1-2-1c-.4 0-.7.2-.9.5L15.2 19a2 2 0 0 0 2.7 2.7l3.5-1.9c.4-.1.6-.4.6-.8"/></g>`), g.Group(children),
	)
}

func PepperChilli(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 7V4a2 2 0 0 0-4 0m0 6s2 0 4 2c2-2 4-2 4-2"/><path d="M22 10c0 6.6-5.4 12-12 12c-4.4 0-8-2.7-8-6v-.4C3.3 17.1 5 18 7 18c3.9 0 7-3.6 7-8c0-1.7 1.3-3 3-3h2c1.7 0 3 1.3 3 3"/></g>`), g.Group(children),
	)
}

func Pie(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 2C5.5 4 8.5 5 7 7m5-5c-1.5 2 1.5 3 0 5m5-5c-1.5 2 1.5 3 0 5m4 9s-2-5-9-5s-9 5-9 5l1.7 5.1c.2.5.7.9 1.3.9h12c.5 0 1.1-.4 1.3-.9Z"/><path d="M2 16c1.7 0 1.6 1 3.3 1s1.7-1 3.4-1s1.6 1 3.3 1s1.7-1 3.3-1c1.7 0 1.6 1 3.3 1s1.7-1 3.3-1M8.5 16l1.5 6m5.5-6L14 22"/></g>`), g.Group(children),
	)
}

func Pig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 4.5a4.12 4.12 0 0 0-5.5 1.6C13 6 12.5 6 12 6c-4.4 0-8 2.7-8 6c0 1.5.8 2.9 2 4v2a2 2 0 0 0 2 2h2v-2.2a12.3 12.3 0 0 0 4 0V19c0 .6.4 1 1 1h3v-4c.7-.6 1.2-1.2 1.5-2H21c.6 0 1-.4 1-1v-3h-2.5c-.4-1-1.2-1.8-2.2-2.5ZM16 11h.01"/><path d="M2.3 7a2 2 0 0 0 2.2 2.9"/></g>`), g.Group(children),
	)
}

func PigHead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 17.9c1.8-.9 3-2.5 3-5.1c0-1.8-.5-3.4-1.5-4.9c1.5-.3 2.5-1.5 2.5-3V3h-3c-1.3 0-2.4.8-2.8 2a10 10 0 0 0-8.4 0C7.4 3.8 6.3 3 5 3H2v2a3 3 0 0 0 2.5 2.9C3.5 9.3 3 11 3 12.8c0 2.6 1.2 4.2 3 5.1m4-3.9v-2m4 2v-2"/><path d="M14 22a4 4 0 1 0-2-7.45A4 4 0 1 0 10 22Zm-4-4h.01M14 18h.01"/></g>`), g.Group(children),
	)
}

func Pillow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21.3 7.5a2 2 0 1 0-2.9-2.7C17 4.3 14.6 4 12 4s-4.9.3-6.4.8a2 2 0 1 0-2.9 2.7a14 14 0 0 0 0 9a2 2 0 1 0 2.9 2.7c1.5.5 3.8.8 6.4.8s5-.3 6.4-.8a2 2 0 1 0 2.9-2.7a14 14 0 0 0 0-9"/>`), g.Group(children),
	)
}

func PinSafety(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.8 3.2c-1.6-1.6-4.1-1.6-5.7 0L12.3 6S15 9 18 6c-3 3 0 5.7 0 5.7l2.8-2.8c1.6-1.6 1.6-4.2 0-5.7M7.1 21.1l10.3-10.2"/><circle cx="5" cy="19" r="3"/><path d="M2.9 16.9L13.1 6.6"/></g>`), g.Group(children),
	)
}

func PinSafetyOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.8 3.2c-1.6-1.6-4.1-1.6-5.7 0L12.3 6S15 9 18 6c-3 3 0 5.7 0 5.7l2.8-2.8c1.6-1.6 1.6-4.2 0-5.7M7.1 21.1l10.3-10.2"/><circle cx="5" cy="19" r="3"/><path d="M9 2s-4.1 9.5-6.755 15.8"/></g>`), g.Group(children),
	)
}

func PineappleRing(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="10" rx="10" ry="8"/><ellipse cx="12" cy="10" rx="3" ry="2"/><path d="m6 4l1.5 1.5m9.2-2.1L15.5 5M2 10v4c0 4.4 4.5 8 10 8s10-3.6 10-8v-4h-3"/><path d="m8 15l-1 2v3.9m5-4.9v6"/></g>`), g.Group(children),
	)
}

func Planet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="8"/><path d="M4.05 13c-1.7 1.8-2.5 3.5-1.8 4.5c1.1 1.9 6.4 1 11.8-2s8.9-7.1 7.7-9c-.6-1-2.4-1.2-4.7-.7"/></g>`), g.Group(children),
	)
}

func Pond(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 3v2"/><rect width="4" height="7" x="10" y="4" rx="2"/><path d="M4 12v10m8-20v2"/><rect width="4" height="7" x="2" y="5" rx="2"/><path d="M12 11v4.35m3 3.15V22c-3.8 0-7-1.6-7-3.5s3.2-3.5 7-3.5s7 1.6 7 3.5c0 1.3-1.5 2.5-3.9 3.1Z"/></g>`), g.Group(children),
	)
}

func PoundSterlingCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M10 17V9.5a2.5 2.5 0 0 1 5 0M8 13h5m-5 4h7"/></g>`), g.Group(children),
	)
}

func PoundSterlingSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M10 17V9.5a2.5 2.5 0 0 1 5 0M8 13h5m-5 4h7"/></g>`), g.Group(children),
	)
}

func Pram(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18.7 4.4L14.5 10M13 10V2a8.1 8.1 0 0 1 8 8v1c0 1.7-1.3 3-3 3H6c-1.7 0-3-1.3-3-3v-1h18M8.2 18.4l3.3-4.4"/><circle cx="7" cy="20" r="2"/><path d="M15.8 18.4L5.6 4.8A1.94 1.94 0 0 0 2 6"/><circle cx="17" cy="20" r="2"/></g>`), g.Group(children),
	)
}

func Pretzel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 19l-6.5-6.5A4.9 4.9 0 0 1 12 9a5 5 0 0 1 10 0A10 10 0 0 1 2 9a5 5 0 1 1 10 0c0 1.4-.6 2.6-1.5 3.5L4 19"/>`), g.Group(children),
	)
}

func Pumpkin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 2c-1 1-1 2-1 2"/><path d="M17 4c-.9 0-1.8.4-2.5 1.2a3.32 3.32 0 0 0-5 0C8.8 4.4 7.9 4 7 4c-2.8 0-5 4-5 9s2.2 9 5 9c.9 0 1.8-.4 2.5-1.2a3.32 3.32 0 0 0 5 0c.7.8 1.6 1.2 2.5 1.2c2.8 0 5-4 5-9s-2.2-9-5-9"/><path d="M10 11L8 9l-2 2m12 0l-2-2l-2 2m-8 4l2 2l2-2l2 2l2-2l2 2l2-2"/></g>`), g.Group(children),
	)
}

func Razor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m22 11l-1.6 1.6c-.8.8-2 .8-2.8 0l-6.2-6.2c-.8-.8-.8-2 0-2.8L13 2m2.8 2.8l3.4 3.4"/><path d="M17 12c-1.4 1.4-3.6 1.4-4.9 0s-1.4-3.6-.1-5"/><path d="m11.1 10.1l-8.5 8.5a1.95 1.95 0 1 0 2.8 2.8l8.4-8.4"/></g>`), g.Group(children),
	)
}

func RazorBlade(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 8h-2V6H4v2H2v8h2v2h16v-2h2ZM6 11v2m4-1H6"/><circle cx="12" cy="12" r="2"/><path d="M18 12h-4m4-1v2"/></g>`), g.Group(children),
	)
}

func ReelThread(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 6L4.4 4.6A1.5 1.5 0 0 1 5.5 2h13a1.5 1.5 0 0 1 1.1 2.5L18 6M6 6h12v12H6zm0 5l10-5"/><path d="M22 16v-3a4 4 0 0 0-4-4L6 15m2 3l10-5m0 5l1.6 1.4a1.45 1.45 0 0 1-1.1 2.5h-13a1.5 1.5 0 0 1-1.1-2.5L6 18"/></g>`), g.Group(children),
	)
}

func RefrigeratorFreezer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="20" x="5" y="2" rx="2"/><path d="M9 6h.01M5 10h14M9 14h.01"/></g>`), g.Group(children),
	)
}

func RemoveFormattingSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 9V7h10v2m-4-2L8 17m-1 0h3m7-3l-3 3m0-3l3 3"/></g>`), g.Group(children),
	)
}

func Rugby(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.7 2.3c-.2-.2-.9-.4-1.7-.3a4.6 4.6 0 0 0-3.7 5.7c.3.2.9.4 1.7.3a4.6 4.6 0 0 0 3.7-5.7M20 12H4"/><rect width="4" height="6" x="2" y="16" rx="1"/><path d="M4 2v14"/><rect width="4" height="6" x="18" y="16" rx="1"/><path d="M20 2v14"/></g>`), g.Group(children),
	)
}

func RussianRubleCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 11h5a2 2 0 1 0 0-4h-3v10m-2-2h5"/></g>`), g.Group(children),
	)
}

func RussianRubleSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 15h5m-5-4h5a2 2 0 1 0 0-4h-3v10"/></g>`), g.Group(children),
	)
}

func Sausage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 19a3 3 0 1 1-6 0A11 11 0 0 0 5 8a3 3 0 1 1 0-6a17 17 0 0 1 17 17m-9.2-7.8L2 22"/><path d="m9.2 8.8l-2.5 2.5a3.1 3.1 0 0 0 0 4.2l1.8 1.8a3.1 3.1 0 0 0 4.2 0l2.5-2.5"/></g>`), g.Group(children),
	)
}

func Scarf(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.5 2.5L7 15c-.5.5-.6 1.5-.2 2L9 20L21.6 7.6a2 1.7 0 0 0 .1-1.9l-2-3c-.2-.4-.7-.7-1.2-.7h-13c-.5 0-1 .3-1.2.7l-2 3a2 1.7 0 0 0 .2 2l6 5.8M12 10L4.5 2.5M13 20v2m3-16H8m9 6.1V22m0-4h4m-4 2H9v2M21 8.2V20"/>`), g.Group(children),
	)
}

func ScissorsHairComb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 2C5 5 7 5 6 8m4-6c-1 3 1 3 0 6"/><circle cx="4" cy="20" r="2"/><path d="M5.4 18.6L8 16m2.8-2.8L14 10"/><circle cx="12" cy="20" r="2"/><path d="m2 10l8.6 8.6M18 2h2a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2h-2m0-16h4m-4 4h4m-4 4h4m-4 4h4"/></g>`), g.Group(children),
	)
}

func Shark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.6 15a9.07 9.07 0 0 0 11.7 5.3S19 22 22 22c0 0-1-3-3-4.5c1.1-1.5 1.9-3.3 2-5.3l-8 4.6a1.94 1.94 0 1 1-2-3.4l6-3.5s5-2.8 5-6.8c0-.6-.4-1-1-1h-9q-2.7 0-4.8 1.5C5.7 2.5 3.9 2 2 2c0 0 1.4 2.1 2.3 4.5A10.63 10.63 0 0 0 3.1 13m10.7-6L13 6"/><path d="M21.12 6h-3.5c-1.1 0-2.8.5-3.82 1L9 9.8C3 11 2 15 2 15h4"/></g>`), g.Group(children),
	)
}

func ShaveFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 20a7 7 0 0 1-7-7V4c0-.6.4-1 1-1h6M7 7h.01M11 13h3V4c0-.6.4-1 1-1h6m-3 4h.01M14 19v2m4-4l1.5 1.5M19 13h2"/>`), g.Group(children),
	)
}

func ShirtFoldedButtons(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 21H5a2 2 0 0 1-2-2V4c0-.6.4-1 1-1h12c.6 0 1 .4 1 1v15a2 2 0 1 0 4 0V7c0-.6-.4-1-1-1h-3"/><path d="M7 3v1a3 3 0 1 0 6 0V3m-3 8h.01M10 15h.01"/></g>`), g.Group(children),
	)
}

func ShirtLongSleeve(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 19H3c-.6 0-1-.4-1-1V6c0-1.1.8-2.3 1.9-2.6L8 2a4 4 0 0 0 8 0l4.1 1.4C21.2 3.7 22 4.9 22 6v12c0 .6-.4 1-1 1h-3"/><path d="M18 8v13c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1V8"/></g>`), g.Group(children),
	)
}

func ShirtT(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 11H3c-.6 0-1-.4-1-1V6c0-1.1.8-2.3 1.9-2.6L8 2a4 4 0 0 0 8 0l4.1 1.4C21.2 3.7 22 4.9 22 6v4c0 .6-.4 1-1 1h-3"/><path d="M18 8v13c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1V8"/></g>`), g.Group(children),
	)
}

func ShirtTRuler(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 11H3c-.6 0-1-.4-1-1V6c0-1.1.8-2.3 1.9-2.6L8 2a4 4 0 0 0 8 0l4.1 1.4C21.2 3.7 22 4.9 22 6v4c0 .6-.4 1-1 1h-3M6 18V8m12 0v10"/><rect width="20" height="6" x="2" y="16" rx="2"/><path d="M10 16v2m4-2v2"/></g>`), g.Group(children),
	)
}

func ShirtTVNeck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 11H3c-.6 0-1-.4-1-1V6c0-1.1.8-2.3 1.9-2.6L8 2c0 2.2 3 5 4 5s4-2.8 4-5l4.1 1.4C21.2 3.7 22 4.9 22 6v4c0 .6-.4 1-1 1h-3"/><path d="M18 8v13c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1V8"/></g>`), g.Group(children),
	)
}

func Shorts(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 8h20M9 20H4a2 2 0 0 1-2-2V5c0-.6.4-1 1-1h18c.6 0 1 .4 1 1v13a2 2 0 0 1-2 2h-5l-3-5Zm0-8V8m6 0v4M5 13l-3 2m20 0l-3-2"/>`), g.Group(children),
	)
}

func ShortsBoxer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.7 15.8L9 20H4a2 2 0 0 1-2-2V5c0-.6.4-1 1-1h18c.6 0 1 .4 1 1v13a2 2 0 0 1-2 2h-5l-1.7-4.2M2 8h20"/><path d="M16 8v4a4 4 0 0 1-8 0V8"/></g>`), g.Group(children),
	)
}

func ShovelDig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.5c-1.7 0-3-1.3-3-3V2h6v1.5c0 1.7-1.3 3-3 3m0 9.5V6.5M8 22v-4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v4M6 22h12"/>`), g.Group(children),
	)
}

func Shower(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 10V8a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2M7 10h14"/><path d="M3 22V4a2 2 0 0 1 2-2h7a2 2 0 0 1 2 2v2m-4 8h.01M14 14h.01M18 14h.01M9 18h.01M14 18h.01M19 18h.01M8 22h.01M14 22h.01M20 22h.01"/></g>`), g.Group(children),
	)
}

func Skirt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3h12v4H6zm0 4c0 1.7-.4 3.3-1 4.4C3.8 13.6 2 17 2 17s1.8 1.2 4.5 2.1"/><path d="m8 16l-2 4s2.7 1 6 1s6-1 6-1l-2-4"/><path d="M17.5 19.1C20.2 18.2 22 17 22 17s-1.8-3.4-3-5.6c-.6-1.1-1-2.7-1-4.4"/></g>`), g.Group(children),
	)
}

func Skis(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 4l3-1M3 2l7 20m0-20L3 22m-1-2l3 1m17 1V6c0-2.2-2-4-2-4s-2 1.8-2 4c0-2.2-2-4-2-4s-2 1.8-2 4v16ZM18 6v16"/>`), g.Group(children),
	)
}

func SlotCard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 13H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-3M6 9h12"/><path d="m13 9l4 4v4a2 2 0 0 1-2 2H9a2 2 0 0 1-2-2V9"/></g>`), g.Group(children),
	)
}

func SlotCardCredit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 13H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-3M6 9h12"/><path d="M17 9v8.3c0 .9-.9 1.7-2 1.7H9c-1.1 0-2-.7-2-1.7V9m4 0v10"/></g>`), g.Group(children),
	)
}

func SlotDisc(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 13H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2M6 9h12"/><circle cx="12" cy="14" r=".5"/><path d="M8.7 9a6.07 6.07 0 1 0 6.6 0"/></g>`), g.Group(children),
	)
}

func Sneaker(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.1 7.9L12.5 10m4.9.1L16 12M2 16a2 2 0 0 0 2 2h13c2.8 0 5-2.2 5-5a2 2 0 0 0-2-2c-.8 0-1.6-.2-2.2-.7l-6.2-4.2c-.4-.3-.9-.2-1.3.1c0 0-.6.8-1.2 1.1a3.5 3.5 0 0 1-4.2.1C4.4 7 3.7 6.3 3.7 6.3A.92.92 0 0 0 2 7Z"/><path d="M2 11c0 1.7 1.3 3 3 3h7"/></g>`), g.Group(children),
	)
}

func Snowboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 6a4 4 0 0 0-7.2-2.3c-4.2 5.8-5.3 6.9-11.1 11.1a4 4 0 1 0 5.5 5.5c4.2-5.8 5.3-6.9 11.1-11.1c1-.7 1.7-1.9 1.7-3.2"/><path d="M6.15 13H11v4.85"/></g>`), g.Group(children),
	)
}

func Snowman(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="6" r="4"/><path d="M12 14h.01M12 18h.01M2 9h2V7m3 5L4 9m13.8 2.1L20 9m0-2v2h2M9 8.7a7 7 0 1 0 6 0M5 22h14"/></g>`), g.Group(children),
	)
}

func SoapBar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11.3 2.7c.9-.9 2.5-.9 3.4 0l5.6 5.6c.9.9.9 2.5 0 3.4l-8.6 8.6c-.9.9-2.5.9-3.4 0l-5.6-5.6c-.9-.9-.9-2.5 0-3.4Z"/><path d="m13 7l-6 6l3 3l6-6Z"/><circle cx="20.5" cy="17.5" r=".5"/><circle cx="17.5" cy="21.5" r=".5"/><path d="M22 22h.01"/></g>`), g.Group(children),
	)
}

func SoccerBall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M11.9 6.7s-3 1.3-5 3.6c0 0 0 3.6 1.9 5.9c0 0 3.1.7 6.2 0c0 0 1.9-2.3 1.9-5.9c0 .1-2-2.3-5-3.6m0 0V2m5 8.4s3-1.4 4.5-1.6M15 16.3s1.9 2.7 2.9 3.7m-9.1-3.7S6.9 19 6 20"/><path d="M2.6 8.7C4 9 7 10.4 7 10.4"/></g>`), g.Group(children),
	)
}

func SoccerPitch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 5v5m0 4v5"/><circle cx="12" cy="12" r="2"/><path d="M2 9h4v6H2"/><path d="M3 19c-.6 0-1-.4-1-1V6c0-.6.4-1 1-1h18c.6 0 1 .4 1 1v12c0 .6-.4 1-1 1Z"/><path d="M22 15h-4V9h4"/></g>`), g.Group(children),
	)
}

func SocketEu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="2"/><circle cx="12" cy="12" r="6"/><path d="M10 12h.01M14 12h.01"/></g>`), g.Group(children),
	)
}

func SocketUk(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="2"/><path d="M12 8v2m-2 5H8m6 0h2"/></g>`), g.Group(children),
	)
}

func SocketUsa(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="2"/><circle cx="12" cy="12" r="6"/><path d="M10 11v2m4-2v2"/></g>`), g.Group(children),
	)
}

func Socks(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.6 20.4L9 21a3.38 3.38 0 1 1-4.9-4.9l3.5-3.5C8.4 11.6 9 10.4 9 9V3c0-.6.4-1 1-1h10c.6 0 1 .4 1 1v10a5.15 5.15 0 0 1-1.5 3.6L15 21a3.38 3.38 0 1 1-4.9-4.9l3.5-3.5c.8-1 1.4-2.2 1.4-3.6V2M9 6h12"/>`), g.Group(children),
	)
}

func Spider(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 5v1m4 0V5m-4 5.4V8a2 2 0 1 1 4 0v2.4M7 15H4l-2 2.5m5.42-.5L5 20l1 2m2-10l-4-1l-2-3m7 3L5.5 6L7 2"/><path d="M8 18a5 5 0 1 1 8 0s-2 3-4 4c-2-1-4-4-4-4"/><path d="m15 11l3.5-5L17 2m-1 10l4-1l2-3m-5 7h3l2 2.5m-5.43-.5L19 20l-1 2"/></g>`), g.Group(children),
	)
}

func SpiderWeb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 17.2V6.8L12 2L3 6.8v10.4l9 4.8Zm-19 .6L22 6.2m-20 0l20 11.6M12 2v20"/><path d="M17 14.9V9.1l-5-2.6l-5 2.6v5.8l5 2.6Z"/></g>`), g.Group(children),
	)
}

func Stairs(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 16h10v4H2zm2-4h10v4H4zm2-4h10v4H6zm2-4h10v4H8z"/><path d="M12 20h10V4h-4"/></g>`), g.Group(children),
	)
}

func StairsArch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 21H3V11a9 9 0 1 1 18 0Z"/><path d="M20.77 9H12v4m-4 4v-4h13M3 17h18"/></g>`), g.Group(children),
	)
}

func StairsArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 2l-9 9m0-4v4h4m-5 9h5v-5h5v-5h5V5h5"/>`), g.Group(children),
	)
}

func StairsArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 11l9-9M8 2h4v4M2 20h5v-5h5v-5h5V5h5"/>`), g.Group(children),
	)
}

func StarNorth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h18m-9-9v18m5-14L7 17M7 7l10 10"/>`), g.Group(children),
	)
}

func SteeringWheel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m3.3 7l7 4m3.4 0l7-4M12 14v8"/><circle cx="12" cy="12" r="2"/></g>`), g.Group(children),
	)
}

func Strawberry(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 7l3.5-3.5M17 2v5h5M2.1 17.1a4 4 0 0 0 4.8 4.8l9-2.1a6.32 6.32 0 0 0 2.9-10.9L15 5.2A6.5 6.5 0 0 0 4.1 8.3Zm6.4-7.6h.01m3.99-1h.01m-5.01 5h.01m3.99-1h.01m3.99-1h.01m-9.01 6h.01m3.99-1h.01m3.99-1h.01"/>`), g.Group(children),
	)
}

func StrikethroughSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M10.5 12a2.5 2.5 0 0 1 0-5H15m-8 5h10m-9 5h5.5a2.5 2.5 0 0 0 0-5"/></g>`), g.Group(children),
	)
}

func Stroller(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 12.95c1.6-1.6 4.1-1.6 5.7.05"/><circle cx="11" cy="6.5" r="2.5"/><path d="M18.3 17.2L5.45 4.5M19.7 17L13 18.1c-2.7.5-5.5-1-5.7-4.1c-.4-2.6-.9-5.7-1.3-8.3A2 2 0 0 0 2 6"/><circle cx="8" cy="19" r="2"/><circle cx="20" cy="19" r="2"/></g>`), g.Group(children),
	)
}

func SunloungerParasolSun(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="20" cy="4" r="2"/><path d="M2.4 14.4a7 7 0 0 1 13.2-4.8ZM9 12l3 9m-9 0l.7-2.1c.2-.5.7-.9 1.3-.9h12c.5 0 1.3-.4 1.6-.8L22 13m-1 8l-3-3"/></g>`), g.Group(children),
	)
}

func SunloungerParasolSunPalmtree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="4" cy="4" r="2"/><path d="M16 14s1-3 1-8V4s-1-2-3-2c-1 0-2 .5-2 .5"/><path d="M13 8a4 4 0 0 1 8 0"/><path d="M17 4s1-2 3-2c1 0 2 .5 2 .5M4 14l3-5l5 3Zm4-1l2 8m-7 0l.7-2.1c.2-.5.7-.9 1.3-.9h12c.5 0 1.3-.4 1.6-.8L22 13m-1 8l-3-3"/></g>`), g.Group(children),
	)
}

func SunloungerParasolTable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 8H3l9-6Zm-9 0v13m-4-8h8M3 21l.7-2.1c.2-.5.7-.9 1.3-.9h12c.5 0 1.3-.4 1.6-.8L22 13m-1 8l-3.2-3.2"/>`), g.Group(children),
	)
}

func Surfboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 2L2.6 21.4m11.2-2.2A18 18 0 0 0 22 4V2h-2C10.1 2 2 10.1 2 20a2 2 0 0 0 2 2a17 17 0 0 0 7.63-1.7"/><path d="M7 17c2.7 0 4.9 2.3 5 5a6.7 6.7 0 0 0-.1-9.9"/></g>`), g.Group(children),
	)
}

func Sushi(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="7"/><rect width="8" height="8" x="8" y="8" rx="2"/><path d="M12 8v8m-4-4h4"/></g>`), g.Group(children),
	)
}

func SushiChopsticks(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 18V2m20 16V2M6 11c0-2.8 2.2-5 5-5h2c2.8 0 5 2.2 5 5v6c0 2.8-2.2 5-5 5h-2c-2.8 0-5-2.2-5-5Z"/><path d="M18 13c0 2.8-2.2 5-5 5h-2c-2.8 0-5-2.2-5-5"/><path d="M11 14c-.6 0-1-.4-1-1v-2c0-.6.4-1 1-1h2c.6 0 1 .4 1 1v2c0 .6-.4 1-1 1Z"/></g>`), g.Group(children),
	)
}

func SushiThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 10a2 2 0 0 1-2 2h-.5l-5.6-1.4c-1.1-.3-2.8-.3-3.9 0L4.4 12H4a2 2 0 0 1-2-2a4 4 0 0 1 4-4h12a4 4 0 0 1 4 4M6 11l1-5m3 4l1-4m3 4l1-4m3 5l1-4m1 5v4a2 2 0 0 1-4 0a2 2 0 0 1-4 0a2 2 0 0 1-4 0a2 2 0 0 1-4 0v-4"/>`), g.Group(children),
	)
}

func SushiTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16.4 3.3a8.23 8.23 0 0 0-8.8 0L3.8 6c-2.4 1.7-2.4 4.4 0 6.1l3.9 2.7c2.4 1.7 6.3 1.7 8.7 0l3.9-2.7c2.4-1.7 2.4-4.4 0-6.1Z"/><path d="M2 9v6c0 1.1.6 2.2 1.8 3l3.9 2.7c2.4 1.7 6.3 1.7 8.7 0l3.9-2.7c1.2-.8 1.8-1.9 1.8-3V9"/><path d="M7.7 10.1c-.9-.6-.9-1.6 0-2.2l2.7-1.8c.9-.6 2.4-.6 3.3 0l2.7 1.8c.9.6.9 1.6 0 2.2l-2.7 1.8c-.9.6-2.4.6-3.3 0Z"/><path d="M15 11c-2-3-5-2-6 0"/></g>`), g.Group(children),
	)
}

func Sweater(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 19H3c-.6 0-1-.4-1-1V6c0-1.1.8-2.3 1.9-2.6L8 2a4 4 0 0 0 8 0l4.1 1.4C21.2 3.7 22 4.9 22 6v12c0 .6-.4 1-1 1h-3"/><path d="M18 8v13c0 .6-.4 1-1 1H7c-.6 0-1-.4-1-1V8"/><path d="m6 10l2 2l2-2l2 2l2-2l2 2l2-2M6 16l2 2l2-2l2 2l2-2l2 2l2-2"/></g>`), g.Group(children),
	)
}

func SwissFrancCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M10 17V7h5m-5 4h4m-6 4h5"/></g>`), g.Group(children),
	)
}

func SwissFrancSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M10 17V7h5m-5 4h4m-6 4h5"/></g>`), g.Group(children),
	)
}

func Tab(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v14m2 0H2"/>`), g.Group(children),
	)
}

func TabArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 16V8m4 4l-4 4l-4-4"/><path d="M4 20V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v14m2 0H2"/></g>`), g.Group(children),
	)
}

func TabArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 9l-6 6m0-6h6v6"/><path d="M4 20V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v14m2 0H2"/></g>`), g.Group(children),
	)
}

func TabCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 12l2 2l4-4"/><path d="M4 20V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v14m2 0H2"/></g>`), g.Group(children),
	)
}

func TabDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="1"/><path d="M4 20V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v14m2 0H2"/></g>`), g.Group(children),
	)
}

func TabPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12H9m3-3v6m-8 5V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v14m2 0H2"/>`), g.Group(children),
	)
}

func TabSlash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.5 9.5l-5 5M4 20V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v14m2 0H2"/>`), g.Group(children),
	)
}

func TabText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8h6m-6 4h8m-8 4h6M4 20V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v14m2 0H2"/>`), g.Group(children),
	)
}

func TabX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.5 9.5l-5 5m5 0l-5-5M4 20V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v14m2 0H2"/>`), g.Group(children),
	)
}

func TargetArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 2v3h3m-8.6 5.6L22 2"/><circle cx="12" cy="12" r="2"/><path d="M12.3 6H12a6 6 0 1 0 6 6v-.3"/><path d="M15 2.5A9.93 9.93 0 1 0 21.5 9M5.3 19.4L4 22m14.7-2.6L20 22"/></g>`), g.Group(children),
	)
}

func TennisBall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12c5.5 0 10-4.5 10-10"/><circle cx="12" cy="12" r="10"/><path d="M22 12c-5.5 0-10 4.5-10 10"/></g>`), g.Group(children),
	)
}

func TennisRacket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.7 4.7c3-3 7.4-3.6 9.8-1.2s1.8 6.8-1.2 9.8a9.5 9.5 0 0 1-4.3 2.5c-2.1.5-4.1.1-5.5-1.3S7.7 11.1 8.2 9a9.5 9.5 0 0 1 2.5-4.3"/><path d="M8.2 9L6 18l9-2.2M2 22l4-4"/><circle cx="20" cy="20" r="2"/></g>`), g.Group(children),
	)
}

func TextSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 12h10M7 16h6M7 8h8"/></g>`), g.Group(children),
	)
}

func Tie(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.9 3c.1-.6.5-1 1.1-1h4c.6 0 1 .4 1.1 1l.9 15l-4 4l-4-4Zm5.95-.6L16 11.8m-6.3 1.35l6.5 8.5M22 5v16c0 .6-.4 1-1 1h-4c-.6 0-1-.4-1-1V5l3-3Z"/>`), g.Group(children),
	)
}

func TieBow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 10h4v4h-4zm-2 2h2"/><path d="M10 10C8.8 8.5 6.6 7 4 7c-1.1 0-2 2.2-2 5s.9 5 2 5c2.6 0 4.8-1.5 6-3m4-2h2m-2 2c1.2 1.5 3.4 3 6 3c1.1 0 2-2.2 2-5s-.9-5-2-5c-2.6 0-4.8 1.5-6 3"/></g>`), g.Group(children),
	)
}

func TieBowRibbon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 7h4v4h-4zM8 9h2"/><path d="M10 7C8.8 5.5 6.6 4 4 4C2.9 4 2 6.2 2 9s.9 5 2 5c2.6 0 4.8-1.5 6-3m4-2h2m-2 2c1.2 1.5 3.4 3 6 3c1.1 0 2-2.2 2-5s-.9-5-2-5c-2.6 0-4.8 1.5-6 3"/><path d="M5.5 13.83L4 20l3-1l2 2l2.5-10m7 2.83L20 20l-3-1l-2 2l-2.5-10"/></g>`), g.Group(children),
	)
}

func Tire(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="2"/><circle cx="12" cy="12" r="6"/><path d="M12 14v4m-1.9-5.38l-3.8 1.23m4.52-3.47L8.47 7.15m5.43 5.47l3.8 1.23m-4.52-3.47l2.35-3.23"/></g>`), g.Group(children),
	)
}

func Toast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.5 3A3.5 3.5 0 0 0 3 8.9V19a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8.9A3.5 3.5 0 0 0 18.5 3Z"/><path d="M7.5 10c0-1.8 1.2-3 3.2-3c2.5 0 2.4 1.5 3.8 2.5s2.5 1 2.5 3c0 2.2-1.2 3.2-3.5 3.2c-1.2 0-1.2 1.2-3 1.2S7 16 7 14.2c0-1.5.8-1.5.8-2.5c0-.7-.3-1.2-.3-1.7"/></g>`), g.Group(children),
	)
}

func Toaster(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 10V5.7A2 2 0 0 0 15 2H9a2 2 0 0 0-1 3.7V10m-2 0h12"/><path d="M4 7a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2"/><circle cx="8" cy="16" r="2"/><path d="M14 16h4m-2-2v8"/></g>`), g.Group(children),
	)
}

func ToiletRoll(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="10" cy="8" rx="3" ry="2"/><ellipse cx="10" cy="8" rx="7" ry="6"/><path d="M3 8v8c0 3.3 3.1 6 7 6s7-2.7 7-6V8c0 2.2 2.2 4 5 4v8c-2.8 0-5-1.8-5-4m-7-2v2m0 4v2"/></g>`), g.Group(children),
	)
}

func Toolbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 7V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2"/><path d="M8 21a2 2 0 0 0 2-2v-8a4 4 0 0 0-8 0v8a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-8a4 4 0 0 0-4-4H6m-4 6h20m-8 2v-4m4 4v-4"/></g>`), g.Group(children),
	)
}

func ToolboxTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2M4 21a2 2 0 0 1-2-2v-7c0-.6.3-1.3.7-1.7l2.6-2.6C5.7 7.3 6.4 7 7 7h10c.6 0 1.3.3 1.7.7l2.6 2.6c.4.4.7 1.2.7 1.7v7a2 2 0 0 1-2 2Zm-2-7h20M9 16v-4m6 4v-4"/>`), g.Group(children),
	)
}

func TopCrop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 17a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5c-1.7 0-3-1.3-3-3V5h-4v1a3 3 0 1 1-6 0V5H5v4c0 1.7-1.3 3-3 3Z"/>`), g.Group(children),
	)
}

func TowelFolded(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 13h10a4 4 0 0 1 0 8H7a4 4 0 0 1-4-4V7a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4v10"/><path d="M17 17H7a4 4 0 0 1-4-4"/></g>`), g.Group(children),
	)
}

func TowelRack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6H2m4-4h12a2 2 0 0 1 2 2v18H8V4a2 2 0 0 0-4 0v15h4M22 6h-2M8 18h12"/>`), g.Group(children),
	)
}

func TreesForest(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 5l3-3l3 3m-6 5l3-3l3 3m-3 2V2M2 15l3-3l3 3m-6 5l3-3l3 3m-3 2V12m11 3l3-3l3 3m-6 5l3-3l3 3m-3 2V12"/>`), g.Group(children),
	)
}

func TriangleStripes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.75 4a2 2 0 0 0-3.5 0L2.2 18A2 2.1 0 0 0 4 21h16a2 2 0 0 0 1.75-3ZM7.5 9h9m-11 4h13M3 17h18"/>`), g.Group(children),
	)
}

func Trousers(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M6 22a2 2 0 0 1-2-2V3c0-.6.4-1 1-1h14c.6 0 1 .4 1 1v17a2 2 0 0 1-2 2h-3l-3-10l-3 10Zm0-11l-2 1m5-3.5V6m6 0v2.5m5 3.5l-2-1M4 18h6m4 0h6"/>`), g.Group(children),
	)
}

func Tuxedo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 3v2l4-2v2Z"/><path d="M18 3h1a2 2 0 0 1 1.7 3A5271 5271 0 0 0 12 21S6.8 12 3.3 6A2 2 0 0 1 5 3h1m6 6h.01M12 13h.01"/><path d="M21 5v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5"/></g>`), g.Group(children),
	)
}

func TypeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 9V7h10v2m-5-2v10m-2 0h4"/></g>`), g.Group(children),
	)
}

func Ufo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 8c0 1-3 2-6 2S6 9 6 8a6 6 0 0 1 12 0M7 13h.01M12 14h.01M17 13h.01"/><path d="M6 8.1c-2.4 1-4 2.6-4 4.4c0 3 4.5 5.5 10 5.5s10-2.5 10-5.5c0-1.8-1.6-3.4-4-4.4M7 22l2-4m8 4l-2-4"/></g>`), g.Group(children),
	)
}

func UnderlineSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M15 7v3a3 3 0 1 1-6 0V7M7 17h10"/></g>`), g.Group(children),
	)
}

func UnicornHead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15.6 4.8l2.7 2.3M15.5 10S19 7 22 2c-6 2-10 5-10 5m-.5 5H11"/><path d="M5 15a4 4 0 0 0 4 4h7.8l.3.3a3 3 0 0 0 4-4.46L12 7c0-3-1-5-1-5S8 3 8 7c-4 1-6 3-6 3"/><path d="M2 4.5C4 3 6 3 6 3l2 4M6.14 17.8S4 19 2 22"/></g>`), g.Group(children),
	)
}

func Venn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="8" cy="12" r="6"/><circle cx="16" cy="12" r="6"/></g>`), g.Group(children),
	)
}

func Vest(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 4a2 2 0 0 0 4 0V3h4v3c0 1.7 1.3 3 3 3v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9c1.7 0 3-1.3 3-3V3h4Z"/>`), g.Group(children),
	)
}

func Volleyball(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M6.3 3.8a16.55 16.55 0 0 0 1.9 11.5M20.7 17a12.8 12.8 0 0 0-8.7-5a13.3 13.3 0 0 1 0-10"/><path d="M22 11.1c-.8-.6-1.7-1.3-2.6-1.8c-3-1.7-6.1-2.5-8.3-2.2m-3.3 14c1-.4 1.9-.8 2.9-1.4c3-1.7 5.2-4 6.1-6.1M12 12a12.6 12.6 0 0 1-8.7 5"/></g>`), g.Group(children),
	)
}

func Waffle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m7 14l7-7m-4 10l7-7M7 10l7 7M10 7l7 7"/></g>`), g.Group(children),
	)
}

func Wardrobe(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="20" x="3" y="2" rx="2"/><path d="M8 10h.01M12 2v15m4-7h.01M3 17h18"/></g>`), g.Group(children),
	)
}

func WatchActivity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15.8 6l-.5-2.4c-.2-1-1-1.6-2-1.6h-2.7a2 2 0 0 0-2 1.6L8.2 6"/><rect width="12" height="12" x="6" y="6" rx="2"/><path d="m8.2 18l.5 2.4c.2 1 1 1.6 2 1.6h2.7a2 2 0 0 0 2-1.6l.5-2.4M6 12h3l2 2l2.2-4l1.8 2h3"/></g>`), g.Group(children),
	)
}

func WatchAlarm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2.5 9a9.93 9.93 0 0 0 0 6m19 0a9.93 9.93 0 0 0 0-6"/><circle cx="12" cy="12" r="6"/><path d="M12 10v2l1 1m3.13-5.34l-.81-4.05a2 2 0 0 0-2-1.61h-2.68a2 2 0 0 0-2 1.61l-.78 4.05m.02 8.7l.8 4a2 2 0 0 0 2 1.61h2.72a2 2 0 0 0 2-1.61l.81-4.05"/></g>`), g.Group(children),
	)
}

func WatchBars(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15.8 6l-.5-2.4c-.2-1-1-1.6-2-1.6h-2.7a2 2 0 0 0-2 1.6L8.2 6"/><rect width="12" height="12" x="6" y="6" rx="2"/><path d="m8.2 18l.5 2.4c.2 1 1 1.6 2 1.6h2.7a2 2 0 0 0 2-1.6l.5-2.4M10 12v2m4-4v4"/></g>`), g.Group(children),
	)
}

func WatchCharging(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15.8 6l-.5-2.4c-.2-1-1-1.6-2-1.6h-2.7a2 2 0 0 0-2 1.6L8.2 6"/><rect width="12" height="12" x="6" y="6" rx="2"/><path d="m8.2 18l.5 2.4c.2 1 1 1.6 2 1.6h2.7a2 2 0 0 0 2-1.6l.5-2.4M12 10l-1 2h2l-1 2"/></g>`), g.Group(children),
	)
}

func WatchCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15.8 6l-.5-2.4c-.2-1-1-1.6-2-1.6h-2.7a2 2 0 0 0-2 1.6L8.2 6"/><rect width="12" height="12" x="6" y="6" rx="2"/><path d="m8.2 18l.5 2.4c.2 1 1 1.6 2 1.6h2.7a2 2 0 0 0 2-1.6l.5-2.4M14 11l-2.5 2.5L10 12"/></g>`), g.Group(children),
	)
}

func WatchLoader(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15.8 6l-.5-2.4c-.2-1-1-1.6-2-1.6h-2.7a2 2 0 0 0-2 1.6L8.2 6"/><rect width="12" height="12" x="6" y="6" rx="2"/><path d="m8.2 18l.5 2.4c.2 1 1 1.6 2 1.6h2.7a2 2 0 0 0 2-1.6l.5-2.4M14 12a2 2 0 1 1-2-2"/></g>`), g.Group(children),
	)
}

func WatchMusic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15.8 6l-.5-2.4c-.2-1-1-1.6-2-1.6h-2.7a2 2 0 0 0-2 1.6L8.2 6"/><rect width="12" height="12" x="6" y="6" rx="2"/><path d="m8.2 18l.5 2.4c.2 1 1 1.6 2 1.6h2.7a2 2 0 0 0 2-1.6l.5-2.4"/><circle cx="11.5" cy="13.5" r=".5"/><path d="m14 11l-2-1v3.5"/></g>`), g.Group(children),
	)
}

func WatchSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15.8 6l-.5-2.4c-.2-1-1-1.6-2-1.6h-2.7a2 2 0 0 0-2 1.6L8.2 6"/><rect width="12" height="12" x="6" y="6" rx="2"/><path d="m8.2 18l.5 2.4c.2 1 1 1.6 2 1.6h2.7a2 2 0 0 0 2-1.6l.5-2.4M12 10v2l1 1"/></g>`), g.Group(children),
	)
}

func WatchSquareAlarm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15.8 6l-.5-2.4c-.2-1-1-1.6-2-1.6h-2.7a2 2 0 0 0-2 1.6L8.2 6"/><rect width="12" height="12" x="6" y="6" rx="2"/><path d="m8.2 18l.5 2.4c.2 1 1 1.6 2 1.6h2.7a2 2 0 0 0 2-1.6l.5-2.4M12 10v2l1 1M2 16c0 2.1 1.1 4 2.7 5M22 8c0-2.1-1.1-4-2.7-5"/></g>`), g.Group(children),
	)
}

func WatchText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15.8 6l-.5-2.4c-.2-1-1-1.6-2-1.6h-2.7a2 2 0 0 0-2 1.6L8.2 6"/><rect width="12" height="12" x="6" y="6" rx="2"/><path d="m8.2 18l.5 2.4c.2 1 1 1.6 2 1.6h2.7a2 2 0 0 0 2-1.6l.5-2.4M10 10h2m2 4h-4"/></g>`), g.Group(children),
	)
}

func Watermelon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.7 17.7a1.08 1.08 0 0 1-.08 1.57A12 12 0 0 1 4.73 2.38a1.1 1.1 0 0 1 1.61-.04Z"/><path d="M19.7 15.7A8 8 0 0 1 8.35 4.34M10 11h.01M13 14h.01"/></g>`), g.Group(children),
	)
}

func WaveCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M16.5 20.93a5 5 0 1 1-.6-9a7 7 0 0 0-13.9.6"/></g>`), g.Group(children),
	)
}

func WavesBirds(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 3c3-1 5 2 5 2s2-2.1 5-1.2M10 8c3-1 5 2 5 2s2-3 5-2M2 15c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c2.6 0 2.4 2 5 2c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1M2 21c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c2.6 0 2.4 2 5 2c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1"/>`), g.Group(children),
	)
}

func WavesSharkFin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.3 14.8C15.3 11 15.8 6.2 19 3C11.6 3 5.6 8.7 5.1 16"/><path d="M2 15c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c2.6 0 2.4 2 5 2c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1M2 21c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c2.6 0 2.4 2 5 2c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1"/></g>`), g.Group(children),
	)
}

func Whale(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 9.1V5a2 2 0 0 0-4 0m4 0a2 2 0 0 1 4 0"/><path d="M6 9.7L3.9 8.4C2.7 7.7 2 6.4 2 5V3c2 0 4 2 4 2s2-2 4-2v2c0 1.4-.7 2.7-1.9 3.4l-3.8 2.4A5 5 0 0 0 7 20h12c1.7 0 3-1.3 3-3v-3c0-2.8-2.2-5-5-5c-2.7 0-5.1 1.4-6.4 3.6L9.7 14A2 2 0 0 1 6 13Zm9 5.3h.01"/></g>`), g.Group(children),
	)
}

func WhaleNarwhal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 9.98s1-3 1-7c-3 2-5 6-5 6.08"/><path d="M6 9.7L3.9 8.4C2.7 7.7 2 6.4 2 5V3c2 0 4 2 4 2s2-2 4-2v2c0 1.4-.7 2.7-1.9 3.4l-3.8 2.4A5 5 0 0 0 7 20h12c1.7 0 3-1.3 3-3v-3c0-2.8-2.2-5-5-5c-2.7 0-5.1 1.4-6.4 3.6L9.7 14A2 2 0 0 1 6 13Zm9 5.3h.01"/></g>`), g.Group(children),
	)
}

func Wheel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="2.5"/><path d="M12 2v7.5M19 5l-5.23 5.23M22 12h-7.5m4.5 7l-5.23-5.23M12 14.5V22m-1.77-8.23L5 19m4.5-7H2m8.23-1.77L5 5"/></g>`), g.Group(children),
	)
}

func Whisk(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 2L3.45 20.55m.05-7.05a5 5 0 1 0 7.1 7.1C12.6 18.6 15 9 15 9s-9.6 2.5-11.5 4.5"/>`), g.Group(children),
	)
}

func WhiskForkKnife(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 2C3.8 2 2 3.8 2 6s4 8 4 8s4-5.8 4-8s-1.8-4-4-4m0 20V2m12 20v-4a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v4m4 0V2m8 11h-2a2 2 0 0 1-2-2V6a4 4 0 0 1 4-4v20"/>`), g.Group(children),
	)
}

func Whisks(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 2v20m0-12s-4 5.8-4 8s1.8 4 4 4s4-1.8 4-4s-4-8-4-8m12-8v20m0-12s-4 5.8-4 8s1.8 4 4 4s4-1.8 4-4s-4-8-4-8"/>`), g.Group(children),
	)
}

func Windmill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 14l8 4l2-4L4 6l2-4l8 4"/><path d="m8 8l-4 8l4 2m8-6l4-8l-4-2L6 22m13 0l-2.4-4.6M12.5 20v2M4 22h17"/></g>`), g.Group(children),
	)
}

func WineGlassBottle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 13h8M5 7s-2 3-2 6a4 4 0 0 0 8 0c0-3-2-6-2-6Zm2 10v5m-3 0h6m8-18c0 3-3 3-3 6v11c0 .6.4 1 1 1h4c.6 0 1-.4 1-1V10c0-3-3-3-3-6m0 0V2"/>`), g.Group(children),
	)
}

func YarnBall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M10 6h10m-6 4h7.8M7.2 3.2l13.6 13.6M4 6l15.3 15.3c.4.4 1.2.7 1.7.7h1M2.2 10.2l11.6 11.6"/></g>`), g.Group(children),
	)
}

func YinYang(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="7" r=".5"/><path d="M12 22a5 5 0 1 0 0-10a5 5 0 1 1 0-10"/><circle cx="12" cy="17" r=".5"/></g>`), g.Group(children),
	)
}
