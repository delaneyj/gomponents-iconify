package lucide

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const IconifyVersion = ""

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 24 24")
)

func Accessibility(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="16" cy="4" r="1"/><path d="m18 19l1-7l-6 1M5 8l3-3l5.5 3l-2.36 3.5m-6.9 3a5 5 0 0 0 6.88 6"/><path d="M13.76 17.5a5 5 0 0 0-6.88-6"/></g>`), g.Group(children),
	)
}

func Activity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12h-4l-3 9L9 3l-3 9H2"/>`), g.Group(children),
	)
}

func ActivitySquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M17 12h-2l-2 5l-2-10l-2 5H7"/></g>`), g.Group(children),
	)
}

func AirVent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2M6 8h12m.3 9.7a2.5 2.5 0 0 1-3.16 3.83a2.53 2.53 0 0 1-1.14-2V12m-7.4 3.6A2 2 0 1 0 10 17v-5"/>`), g.Group(children),
	)
}

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 17H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1"/><path d="m12 15l5 6H7l5-6z"/></g>`), g.Group(children),
	)
}

func AlarmCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="13" r="8"/><path d="M5 3L2 6m20 0l-3-3M6.38 18.7L4 21m13.64-2.33L20 21M9 13l2 2l4-4"/></g>`), g.Group(children),
	)
}

func AlarmClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="13" r="8"/><path d="M12 9v4l2 2M5 3L2 6m20 0l-3-3M6.38 18.7L4 21m13.64-2.33L20 21"/></g>`), g.Group(children),
	)
}

func AlarmClockOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.87 6.87a8 8 0 1 0 11.26 11.26m1.77-3.88a8 8 0 0 0-9.15-9.15M22 6l-3-3M6.26 18.67L4 21M2 2l20 20M4 4L2 6"/>`), g.Group(children),
	)
}

func AlarmMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="13" r="8"/><path d="M5 3L2 6m20 0l-3-3M6.38 18.7L4 21m13.64-2.33L20 21M9 13h6"/></g>`), g.Group(children),
	)
}

func AlarmPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="13" r="8"/><path d="M5 3L2 6m20 0l-3-3M6.38 18.7L4 21m13.64-2.33L20 21m-8-11v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func Album(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M11 3v8l3-3l3 3V3"/></g>`), g.Group(children),
	)
}

func AlertCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 8v4m0 4h.01"/></g>`), g.Group(children),
	)
}

func AlertOctagon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.86 2h8.28L22 7.86v8.28L16.14 22H7.86L2 16.14V7.86L7.86 2zM12 8v4m0 4h.01"/>`), g.Group(children),
	)
}

func AlertTriangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21.73 18l-8-14a2 2 0 0 0-3.48 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3ZM12 9v4m0 4h.01"/>`), g.Group(children),
	)
}

func AlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 6H3m14 6H7m12 6H5"/>`), g.Group(children),
	)
}

func AlignCenterHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h20m-12 4v4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-4m6-8V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v4m16 8v1a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-1m0-8V7c0-1.1.9-2 2-2h2a2 2 0 0 1 2 2v1"/>`), g.Group(children),
	)
}

func AlignCenterVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v20M8 10H4a2 2 0 0 1-2-2V6c0-1.1.9-2 2-2h4m8 6h4a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-4M8 20H7a2 2 0 0 1-2-2v-2c0-1.1.9-2 2-2h1m8 0h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1"/>`), g.Group(children),
	)
}

func AlignEndHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="16" x="4" y="2" rx="2"/><rect width="6" height="9" x="14" y="9" rx="2"/><path d="M22 22H2"/></g>`), g.Group(children),
	)
}

func AlignEndVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="6" x="2" y="4" rx="2"/><rect width="9" height="6" x="9" y="14" rx="2"/><path d="M22 22V2"/></g>`), g.Group(children),
	)
}

func AlignHorizontalDistributeCenter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="4" y="5" rx="2"/><rect width="6" height="10" x="14" y="7" rx="2"/><path d="M17 22v-5m0-10V2M7 22v-3M7 5V2"/></g>`), g.Group(children),
	)
}

func AlignHorizontalDistributeEnd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="4" y="5" rx="2"/><rect width="6" height="10" x="14" y="7" rx="2"/><path d="M10 2v20M20 2v20"/></g>`), g.Group(children),
	)
}

func AlignHorizontalDistributeStart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="4" y="5" rx="2"/><rect width="6" height="10" x="14" y="7" rx="2"/><path d="M4 2v20M14 2v20"/></g>`), g.Group(children),
	)
}

func AlignHorizontalJustifyCenter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="2" y="5" rx="2"/><rect width="6" height="10" x="16" y="7" rx="2"/><path d="M12 2v20"/></g>`), g.Group(children),
	)
}

func AlignHorizontalJustifyEnd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="2" y="5" rx="2"/><rect width="6" height="10" x="12" y="7" rx="2"/><path d="M22 2v20"/></g>`), g.Group(children),
	)
}

func AlignHorizontalJustifyStart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="6" y="5" rx="2"/><rect width="6" height="10" x="16" y="7" rx="2"/><path d="M2 2v20"/></g>`), g.Group(children),
	)
}

func AlignHorizontalSpaceAround(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="10" x="9" y="7" rx="2"/><path d="M4 22V2m16 20V2"/></g>`), g.Group(children),
	)
}

func AlignHorizontalSpaceBetween(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="14" x="3" y="5" rx="2"/><rect width="6" height="10" x="15" y="7" rx="2"/><path d="M3 2v20M21 2v20"/></g>`), g.Group(children),
	)
}

func AlignJustify(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18M3 12h18M3 18h18"/>`), g.Group(children),
	)
}

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 6H3m12 6H3m14 6H3"/>`), g.Group(children),
	)
}

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 6H3m18 6H9m12 6H7"/>`), g.Group(children),
	)
}

func AlignStartHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="16" x="4" y="6" rx="2"/><rect width="6" height="9" x="14" y="6" rx="2"/><path d="M22 2H2"/></g>`), g.Group(children),
	)
}

func AlignStartVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="9" height="6" x="6" y="14" rx="2"/><rect width="16" height="6" x="6" y="4" rx="2"/><path d="M2 2v20"/></g>`), g.Group(children),
	)
}

func AlignVerticalDistributeCenter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="14" rx="2"/><rect width="10" height="6" x="7" y="4" rx="2"/><path d="M22 7h-5M7 7H1m21 10h-3M5 17H2"/></g>`), g.Group(children),
	)
}

func AlignVerticalDistributeEnd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="14" rx="2"/><rect width="10" height="6" x="7" y="4" rx="2"/><path d="M2 20h20M2 10h20"/></g>`), g.Group(children),
	)
}

func AlignVerticalDistributeStart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="14" rx="2"/><rect width="10" height="6" x="7" y="4" rx="2"/><path d="M2 14h20M2 4h20"/></g>`), g.Group(children),
	)
}

func AlignVerticalJustifyCenter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="16" rx="2"/><rect width="10" height="6" x="7" y="2" rx="2"/><path d="M2 12h20"/></g>`), g.Group(children),
	)
}

func AlignVerticalJustifyEnd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="12" rx="2"/><rect width="10" height="6" x="7" y="2" rx="2"/><path d="M2 22h20"/></g>`), g.Group(children),
	)
}

func AlignVerticalJustifyStart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="16" rx="2"/><rect width="10" height="6" x="7" y="6" rx="2"/><path d="M2 2h20"/></g>`), g.Group(children),
	)
}

func AlignVerticalSpaceAround(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="10" height="6" x="7" y="9" rx="2"/><path d="M22 20H2M22 4H2"/></g>`), g.Group(children),
	)
}

func AlignVerticalSpaceBetween(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="6" x="5" y="15" rx="2"/><rect width="10" height="6" x="7" y="3" rx="2"/><path d="M2 21h20M2 3h20"/></g>`), g.Group(children),
	)
}

func Ampersand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 12c0 4.4-3.6 8-8 8A4.5 4.5 0 0 1 5 15.5c0-6 8-4 8-8.5a3 3 0 1 0-6 0c0 3 2.5 8.5 12 13m-3-8h3"/>`), g.Group(children),
	)
}

func Ampersands(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 17c-5-3-7-7-7-9a2 2 0 0 1 4 0c0 2.5-5 2.5-5 6c0 1.7 1.3 3 3 3c2.8 0 5-2.2 5-5m12 5c-5-3-7-7-7-9a2 2 0 0 1 4 0c0 2.5-5 2.5-5 6c0 1.7 1.3 3 3 3c2.8 0 5-2.2 5-5"/>`), g.Group(children),
	)
}

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="5" r="3"/><path d="M12 22V8m-7 4H2a10 10 0 0 0 20 0h-3"/></g>`), g.Group(children),
	)
}

func Angry(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M16 16s-1.5-2-4-2s-4 2-4 2m-.5-8L10 9m4 0l2.5-1M9 10h0m6 0h0"/></g>`), g.Group(children),
	)
}

func Annoyed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 15h8M8 9h2m4 0h2"/></g>`), g.Group(children),
	)
}

func Antenna(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12L7 2m0 10l5-10m0 10l5-10m0 10l5-10M4.5 7h15M12 16v6"/>`), g.Group(children),
	)
}

func Aperture(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m14.31 8l5.74 9.94M9.69 8h11.48M7.38 12l5.74-9.94M9.69 16L3.95 6.06M14.31 16H2.83m13.79-4l-5.74 9.94"/></g>`), g.Group(children),
	)
}

func AppWindow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="M10 4v4M2 8h20M6 4v4"/></g>`), g.Group(children),
	)
}

func Apple(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 20.94c1.5 0 2.75 1.06 4 1.06c3 0 6-8 6-12.22A4.91 4.91 0 0 0 17 5c-2.22 0-4 1.44-5 2c-1-.56-2.78-2-5-2a4.9 4.9 0 0 0-5 4.78C2 14 5 22 8 22c1.25 0 2.5-1.06 4-1.06Z"/><path d="M10 2c1 .5 2 2 2 5"/></g>`), g.Group(children),
	)
}

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="5" x="2" y="3" rx="1"/><path d="M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8m-10 4h4"/></g>`), g.Group(children),
	)
}

func ArchiveRestore(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="5" x="2" y="3" rx="1"/><path d="M4 8v11a2 2 0 0 0 2 2h2M20 8v11a2 2 0 0 1-2 2h-2m-7-6l3-3l3 3m-3-3v9"/></g>`), g.Group(children),
	)
}

func ArchiveX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="5" x="2" y="3" rx="1"/><path d="M4 8v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8M9.5 17l5-5m-5 0l5 5"/></g>`), g.Group(children),
	)
}

func AreaChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3v18h18"/><path d="M7 12v5h12V8l-5 5l-4-4Z"/></g>`), g.Group(children),
	)
}

func Armchair(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 9V6a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v3"/><path d="M3 16a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5a2 2 0 0 0-4 0v2H7v-2a2 2 0 0 0-4 0Zm2 2v2m14-2v2"/></g>`), g.Group(children),
	)
}

func ArrowBigDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 6v6h4l-7 7l-7-7h4V6h6z"/>`), g.Group(children),
	)
}

func ArrowBigDownDash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 5H9m6 4v3h4l-7 7l-7-7h4V9h6z"/>`), g.Group(children),
	)
}

func ArrowBigLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 15h-6v4l-7-7l7-7v4h6v6z"/>`), g.Group(children),
	)
}

func ArrowBigLeftDash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 15V9m-4 6h-3v4l-7-7l7-7v4h3v6z"/>`), g.Group(children),
	)
}

func ArrowBigRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 9h6V5l7 7l-7 7v-4H6V9z"/>`), g.Group(children),
	)
}

func ArrowBigRightDash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9v6m4-6h3V5l7 7l-7 7v-4H9V9z"/>`), g.Group(children),
	)
}

func ArrowBigUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 18v-6H5l7-7l7 7h-4v6H9z"/>`), g.Group(children),
	)
}

func ArrowBigUpDash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 19h6m-6-4v-3H5l7-7l7 7h-4v3H9z"/>`), g.Group(children),
	)
}

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v14m7-7l-7 7l-7-7"/>`), g.Group(children),
	)
}

func ArrowDownAZ(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 16l4 4l4-4m-4 4V4m13 4h-5m0 2V6.5a2.5 2.5 0 0 1 5 0V10m-5 4h5l-5 6h5"/>`), g.Group(children),
	)
}

func ArrowDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 8v8m-4-4l4 4l4-4"/></g>`), g.Group(children),
	)
}

func ArrowDownFromLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 3H5m7 18V7m-6 8l6 6l6-6"/>`), g.Group(children),
	)
}

func ArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 7L7 17m10 0H7V7"/>`), g.Group(children),
	)
}

func ArrowDownLeftFromCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12a10 10 0 1 1 10 10M2 22l10-10M8 22H2v-6"/>`), g.Group(children),
	)
}

func ArrowDownLeftSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m16 8l-8 8m8 0H8V8"/></g>`), g.Group(children),
	)
}

func ArrowDownNarrowWide(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 16l4 4l4-4m-4 4V4m4 0h4m-4 4h7m-7 4h10"/>`), g.Group(children),
	)
}

func ArrowDownOneZero(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 16l4 4l4-4m-4 4V4m10 6V4h-2m0 6h4"/><rect width="4" height="6" x="15" y="14" ry="2"/></g>`), g.Group(children),
	)
}

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10m0-10v10H7"/>`), g.Group(children),
	)
}

func ArrowDownRightFromCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22a10 10 0 1 1 10-10m0 10L12 12m10 4v6h-6"/>`), g.Group(children),
	)
}

func ArrowDownRightSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m8 8l8 8m0-8v8H8"/></g>`), g.Group(children),
	)
}

func ArrowDownSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 8v8m-4-4l4 4l4-4"/></g>`), g.Group(children),
	)
}

func ArrowDownToDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 2v14m7-7l-7 7l-7-7"/><circle cx="12" cy="21" r="1"/></g>`), g.Group(children),
	)
}

func ArrowDownToLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17V3m-6 8l6 6l6-6m1 10H5"/>`), g.Group(children),
	)
}

func ArrowDownUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 16l4 4l4-4m-4 4V4m14 4l-4-4l-4 4m4-4v16"/>`), g.Group(children),
	)
}

func ArrowDownWideNarrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 16l4 4l4-4m-4 4V4m4 0h10M11 8h7m-7 4h4"/>`), g.Group(children),
	)
}

func ArrowDownZA(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 16l4 4l4-4M7 4v16m8-16h5l-5 6h5m-5 10v-3.5a2.5 2.5 0 0 1 5 0V20m0-2h-5"/>`), g.Group(children),
	)
}

func ArrowDownZeroOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 16l4 4l4-4m-4 4V4"/><rect width="4" height="6" x="15" y="4" ry="2"/><path d="M17 20v-6h-2m0 6h4"/></g>`), g.Group(children),
	)
}

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 19l-7-7l7-7m7 7H5"/>`), g.Group(children),
	)
}

func ArrowLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M16 12H8m4-4l-4 4l4 4"/></g>`), g.Group(children),
	)
}

func ArrowLeftFromLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 6l-6 6l6 6m-6-6h14m4 7V5"/>`), g.Group(children),
	)
}

func ArrowLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3L4 7l4 4M4 7h16m-4 14l4-4l-4-4m4 4H4"/>`), g.Group(children),
	)
}

func ArrowLeftSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m12 8l-4 4l4 4m4-4H8"/></g>`), g.Group(children),
	)
}

func ArrowLeftToLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19V5m10 1l-6 6l6 6m-6-6h14"/>`), g.Group(children),
	)
}

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14m-7-7l7 7l-7 7"/>`), g.Group(children),
	)
}

func ArrowRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 12h8m-4 4l4-4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowRightFromLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5v14m18-7H7m8 6l6-6l-6-6"/>`), g.Group(children),
	)
}

func ArrowRightLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 3l4 4l-4 4m4-4H4m4 14l-4-4l4-4m-4 4h16"/>`), g.Group(children),
	)
}

func ArrowRightSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 12h8m-4 4l4-4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowRightToLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 12H3m8 6l6-6l-6-6m10-1v14"/>`), g.Group(children),
	)
}

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 12l7-7l7 7m-7 7V5"/>`), g.Group(children),
	)
}

func ArrowUpAZ(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 8l4-4l4 4M7 4v16M20 8h-5m0 2V6.5a2.5 2.5 0 0 1 5 0V10m-5 4h5l-5 6h5"/>`), g.Group(children),
	)
}

func ArrowUpCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m16 12l-4-4l-4 4m4 4V8"/></g>`), g.Group(children),
	)
}

func ArrowUpDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 16l-4 4l-4-4m4 4V4M3 8l4-4l4 4M7 4v16"/>`), g.Group(children),
	)
}

func ArrowUpFromDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m5 9l7-7l7 7m-7 7V2"/><circle cx="12" cy="21" r="1"/></g>`), g.Group(children),
	)
}

func ArrowUpFromLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 9l-6-6l-6 6m6-6v14m-7 4h14"/>`), g.Group(children),
	)
}

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 17V7h10m0 10L7 7"/>`), g.Group(children),
	)
}

func ArrowUpLeftFromCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 8V2h6M2 2l10 10m0-10A10 10 0 1 1 2 12"/>`), g.Group(children),
	)
}

func ArrowUpLeftSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 16V8h8m0 8L8 8"/></g>`), g.Group(children),
	)
}

func ArrowUpNarrowWide(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 8l4-4l4 4M7 4v16m4-8h4m-4 4h7m-7 4h10"/>`), g.Group(children),
	)
}

func ArrowUpOneZero(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 8l4-4l4 4M7 4v16m10-10V4h-2m0 6h4"/><rect width="4" height="6" x="15" y="14" ry="2"/></g>`), g.Group(children),
	)
}

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h10v10M7 17L17 7"/>`), g.Group(children),
	)
}

func ArrowUpRightFromCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12A10 10 0 1 1 12 2m10 0L12 12m4-10h6v6"/>`), g.Group(children),
	)
}

func ArrowUpRightSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 8h8v8m-8 0l8-8"/></g>`), g.Group(children),
	)
}

func ArrowUpSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m16 12l-4-4l-4 4m4 4V8"/></g>`), g.Group(children),
	)
}

func ArrowUpToLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3h14m-1 10l-6-6l-6 6m6-6v14"/>`), g.Group(children),
	)
}

func ArrowUpWideNarrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 8l4-4l4 4M7 4v16m4-8h10m-10 4h7m-7 4h4"/>`), g.Group(children),
	)
}

func ArrowUpZA(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 8l4-4l4 4M7 4v16m8-16h5l-5 6h5m-5 10v-3.5a2.5 2.5 0 0 1 5 0V20m0-2h-5"/>`), g.Group(children),
	)
}

func ArrowUpZeroOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 8l4-4l4 4M7 4v16"/><rect width="4" height="6" x="15" y="4" ry="2"/><path d="M17 20v-6h-2m0 6h4"/></g>`), g.Group(children),
	)
}

func ArrowsUpFromLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 6l3-3l3 3M7 17V3m7 3l3-3l3 3m-3 11V3M4 21h16"/>`), g.Group(children),
	)
}

func Asterisk(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v12m5.196-9L6.804 15m0-6l10.392 6"/>`), g.Group(children),
	)
}

func AtSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M16 8v5a3 3 0 0 0 6 0v-1a10 10 0 1 0-4 8"/></g>`), g.Group(children),
	)
}

func Atom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="1"/><path d="M20.2 20.2c2.04-2.03.02-7.36-4.5-11.9c-4.54-4.52-9.87-6.54-11.9-4.5c-2.04 2.03-.02 7.36 4.5 11.9c4.54 4.52 9.87 6.54 11.9 4.5Z"/><path d="M15.7 15.7c4.52-4.54 6.54-9.87 4.5-11.9c-2.03-2.04-7.36-.02-11.9 4.5c-4.52 4.54-6.54 9.87-4.5 11.9c2.03 2.04 7.36.02 11.9-4.5Z"/></g>`), g.Group(children),
	)
}

func Award(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="8" r="6"/><path d="M15.477 12.89L17 22l-5-3l-5 3l1.523-9.11"/></g>`), g.Group(children),
	)
}

func Axe(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14 12l-8.5 8.5a2.12 2.12 0 1 1-3-3L11 9"/><path d="M15 13L9 7l4-4l6 6h3a8 8 0 0 1-7 7z"/></g>`), g.Group(children),
	)
}

func AxisThreeD(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16h16M4 20l7-7"/>`), g.Group(children),
	)
}

func Baby(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 12h.01M15 12h.01M10 16c.5.3 1.2.5 2 .5s1.5-.2 2-.5"/><path d="M19 6.3a9 9 0 0 1 1.8 3.9a2 2 0 0 1 0 3.6a9 9 0 0 1-17.6 0a2 2 0 0 1 0-3.6A9 9 0 0 1 12 3c2 0 3.5 1.1 3.5 2.5s-.9 2.5-2 2.5c-.8 0-1.5-.4-1.5-1"/></g>`), g.Group(children),
	)
}

func Backpack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 10a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2Zm5-4V4a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2"/><path d="M8 21v-5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v5M8 10h8m-8 8h8"/></g>`), g.Group(children),
	)
}

func Badge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/>`), g.Group(children),
	)
}

func BadgeAlert(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM12 8v4m0 4h.01"/>`), g.Group(children),
	)
}

func BadgeCent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM12 7v10"/><path d="M15.4 10a4 4 0 1 0 0 4"/></g>`), g.Group(children),
	)
}

func BadgeCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/><path d="m9 12l2 2l4-4"/></g>`), g.Group(children),
	)
}

func BadgeDollarSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/><path d="M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8m4 2V6"/></g>`), g.Group(children),
	)
}

func BadgeEuro(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM7 12h5"/><path d="M15 9.4a4 4 0 1 0 0 5.2"/></g>`), g.Group(children),
	)
}

func BadgeHelp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3m.08 4h.01"/></g>`), g.Group(children),
	)
}

func BadgeIndianRupee(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM8 8h8m-8 4h8"/><path d="m13 17l-5-1h1a4 4 0 0 0 0-8"/></g>`), g.Group(children),
	)
}

func BadgeInfo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM12 16v-4m0-4h.01"/>`), g.Group(children),
	)
}

func BadgeJapaneseYen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/><path d="m9 8l3 3v7m0-7l3-3m-6 4h6m-6 4h6"/></g>`), g.Group(children),
	)
}

func BadgeMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM8 12h8"/>`), g.Group(children),
	)
}

func BadgePercent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM15 9l-6 6m0-6h.01M15 15h.01"/>`), g.Group(children),
	)
}

func BadgePlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM12 8v8m-4-4h8"/>`), g.Group(children),
	)
}

func BadgePoundSterling(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM8 12h4"/><path d="M10 16V9.5a2.5 2.5 0 0 1 5 0M8 16h7"/></g>`), g.Group(children),
	)
}

func BadgeRussianRuble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM9 16h5"/><path d="M9 12h5a2 2 0 1 0 0-4h-3v9"/></g>`), g.Group(children),
	)
}

func BadgeSwissFranc(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76Z"/><path d="M11 17V8h4m-4 4h3m-5 4h4"/></g>`), g.Group(children),
	)
}

func BadgeX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.85 8.62a4 4 0 0 1 4.78-4.77a4 4 0 0 1 6.74 0a4 4 0 0 1 4.78 4.78a4 4 0 0 1 0 6.74a4 4 0 0 1-4.77 4.78a4 4 0 0 1-6.75 0a4 4 0 0 1-4.78-4.77a4 4 0 0 1 0-6.76ZM15 9l-6 6m0-6l6 6"/>`), g.Group(children),
	)
}

func BaggageClaim(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 18H6a2 2 0 0 1-2-2V7a2 2 0 0 0-2-2"/><path d="M17 14V4a2 2 0 0 0-2-2h-1a2 2 0 0 0-2 2v10"/><rect width="13" height="8" x="8" y="6" rx="1"/><circle cx="18" cy="20" r="2"/><circle cx="9" cy="20" r="2"/></g>`), g.Group(children),
	)
}

func Ban(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m4.9 4.9l14.2 14.2"/></g>`), g.Group(children),
	)
}

func Banana(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 13c3.5-2 8-2 10 2a5.5 5.5 0 0 1 8 5"/><path d="M5.15 17.89c5.52-1.52 8.65-6.89 7-12C11.55 4 11.5 2 13 2c3.22 0 5 5.5 5 8c0 6.5-4.2 12-10.49 12C5.11 22 2 22 2 20c0-1.5 1.14-1.55 3.15-2.11Z"/></g>`), g.Group(children),
	)
}

func Banknote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="12" x="2" y="6" rx="2"/><circle cx="12" cy="12" r="2"/><path d="M6 12h.01M18 12h.01"/></g>`), g.Group(children),
	)
}

func BarChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20V10m6 10V4M6 20v-4"/>`), g.Group(children),
	)
}

func BarChartBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3v18h18"/><rect width="4" height="7" x="7" y="10" rx="1"/><rect width="4" height="12" x="15" y="5" rx="1"/></g>`), g.Group(children),
	)
}

func BarChartFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3v18h18m-8-4V9m5 8V5M8 17v-3"/>`), g.Group(children),
	)
}

func BarChartHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3v18h18M7 16h8m-8-5h12M7 6h3"/>`), g.Group(children),
	)
}

func BarChartHorizontalBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3v18h18"/><rect width="12" height="4" x="7" y="5" rx="1"/><rect width="7" height="4" x="7" y="13" rx="1"/></g>`), g.Group(children),
	)
}

func BarChartThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3v18h18m-3-4V9m-5 8V5M8 17v-3"/>`), g.Group(children),
	)
}

func BarChartTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 20V10m-6 10V4M6 20v-6"/>`), g.Group(children),
	)
}

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5v14M8 5v14m4-14v14m5-14v14m4-14v14"/>`), g.Group(children),
	)
}

func Baseline(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16M6 16l6-12l6 12M8 12h8"/>`), g.Group(children),
	)
}

func Bath(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6L6.5 3.5a1.5 1.5 0 0 0-1-.5C4.683 3 4 3.683 4 4.5V17a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5M10 5L8 7m-6 5h20M7 19v2m10-2v2"/>`), g.Group(children),
	)
}

func Battery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="10" x="2" y="7" rx="2" ry="2"/><path d="M22 11v2"/></g>`), g.Group(children),
	)
}

func BatteryCharging(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7h1a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2h-2M6 7H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h1m6-10l-3 5h4l-3 5m13-6v2"/>`), g.Group(children),
	)
}

func BatteryFull(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="10" x="2" y="7" rx="2" ry="2"/><path d="M22 11v2M6 11v2m4-2v2m4-2v2"/></g>`), g.Group(children),
	)
}

func BatteryLow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="10" x="2" y="7" rx="2" ry="2"/><path d="M22 11v2M6 11v2"/></g>`), g.Group(children),
	)
}

func BatteryMedium(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="10" x="2" y="7" rx="2" ry="2"/><path d="M22 11v2M6 11v2m4-2v2"/></g>`), g.Group(children),
	)
}

func BatteryWarning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 7h2a2 2 0 0 1 2 2v6c0 1-1 2-2 2h-2M6 7H4a2 2 0 0 0-2 2v6c0 1 1 2 2 2h2m16-6v2M10 7v6m0 4v.01"/>`), g.Group(children),
	)
}

func Beaker(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.5 3h15M6 3v16a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V3M6 14h12"/>`), g.Group(children),
	)
}

func Bean(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.165 6.598C9.954 7.478 9.64 8.36 9 9c-.64.64-1.521.954-2.402 1.165A6 6 0 0 0 8 22c7.732 0 14-6.268 14-14a6 6 0 0 0-11.835-1.402Z"/><path d="M5.341 10.62a4 4 0 1 0 5.279-5.28"/></g>`), g.Group(children),
	)
}

func BeanOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 9c-.64.64-1.521.954-2.402 1.165A6 6 0 0 0 8 22a13.96 13.96 0 0 0 9.9-4.1M10.75 5.093A6 6 0 0 1 22 8c0 2.411-.61 4.68-1.683 6.66"/><path d="M5.341 10.62a4 4 0 0 0 6.487 1.208M10.62 5.341a4.015 4.015 0 0 1 2.039 2.04M2 2l20 20"/></g>`), g.Group(children),
	)
}

func Bed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 4v16M2 8h18a2 2 0 0 1 2 2v10M2 17h20M6 8v9"/>`), g.Group(children),
	)
}

func BedDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20v-8a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v8M4 10V6a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4m-8-6v6M2 18h20"/>`), g.Group(children),
	)
}

func BedSingle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 20v-8a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v8M5 10V6a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v4M3 18h18"/>`), g.Group(children),
	)
}

func Beef(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12.5" cy="8.5" r="2.5"/><path d="M12.5 2a6.5 6.5 0 0 0-6.22 4.6c-1.1 3.13-.78 3.9-3.18 6.08A3 3 0 0 0 5 18c4 0 8.4-1.8 11.4-4.3A6.5 6.5 0 0 0 12.5 2Z"/><path d="m18.5 6l2.19 4.5a6.48 6.48 0 0 1 .31 2a6.49 6.49 0 0 1-2.6 5.2C15.4 20.2 11 22 7 22a3 3 0 0 1-2.68-1.66L2.4 16.5"/></g>`), g.Group(children),
	)
}

func Beer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 11h1a3 3 0 0 1 0 6h-1m-8-5v6m4-6v6m1-10.5c-1 0-1.44.5-3 .5s-2-.5-3-.5s-1.72.5-2.5.5a2.5 2.5 0 0 1 0-5c.78 0 1.57.5 2.5.5S9.44 2 11 2s2 1.5 3 1.5s1.72-.5 2.5-.5a2.5 2.5 0 0 1 0 5c-.78 0-1.5-.5-2.5-.5Z"/><path d="M5 8v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V8"/></g>`), g.Group(children),
	)
}

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9m4.3 13a1.94 1.94 0 0 0 3.4 0"/>`), g.Group(children),
	)
}

func BellDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.4 14.9C20.2 16.4 21 17 21 17H3s3-2 3-9c0-3.3 2.7-6 6-6c.7 0 1.3.1 1.9.3M10.3 21a1.94 1.94 0 0 0 3.4 0"/><circle cx="18" cy="8" r="3"/></g>`), g.Group(children),
	)
}

func BellMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.4 12c.8 3.8 2.6 5 2.6 5H3s3-2 3-9c0-3.3 2.7-6 6-6c1.8 0 3.4.8 4.5 2m-6.2 17a1.94 1.94 0 0 0 3.4 0M15 8h6"/>`), g.Group(children),
	)
}

func BellOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.7 3A6 6 0 0 1 18 8a21.3 21.3 0 0 0 .6 5M17 17H3s3-2 3-9a4.67 4.67 0 0 1 .3-1.7m4 14.7a1.94 1.94 0 0 0 3.4 0M2 2l20 20"/>`), g.Group(children),
	)
}

func BellPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.3 14.8C20.1 16.4 21 17 21 17H3s3-2 3-9c0-3.3 2.7-6 6-6c1 0 1.9.2 2.8.7M10.3 21a1.94 1.94 0 0 0 3.4 0M15 8h6m-3-3v6"/>`), g.Group(children),
	)
}

func BellRing(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8a6 6 0 0 1 12 0c0 7 3 9 3 9H3s3-2 3-9m4.3 13a1.94 1.94 0 0 0 3.4 0M4 2C2.8 3.7 2 5.7 2 8m20 0c0-2.3-.8-4.3-2-6"/>`), g.Group(children),
	)
}

func Bike(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18.5" cy="17.5" r="3.5"/><circle cx="5.5" cy="17.5" r="3.5"/><circle cx="15" cy="5" r="1"/><path d="M12 17.5V14l-3-3l4-3l2 3h2"/></g>`), g.Group(children),
	)
}

func Binary(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="4" height="6" x="14" y="14" rx="2"/><rect width="4" height="6" x="6" y="4" rx="2"/><path d="M6 20h4m4-10h4M6 14h2v6m6-16h2v6"/></g>`), g.Group(children),
	)
}

func Biohazard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="11.9" r="2"/><path d="M6.7 3.4c-.9 2.5 0 5.2 2.2 6.7C6.5 9 3.7 9.6 2 11.6m6.9-1.5l1.4.8m7-7.5c.9 2.5 0 5.2-2.2 6.7c2.4-1.2 5.2-.6 6.9 1.5m-6.9-1.5l-1.4.8m3 9.9c-2.6-.4-4.6-2.6-4.7-5.3c-.2 2.6-2.1 4.8-4.7 5.2m4.7-6.8v1.6m1.5-10.1c-1-.2-2-.2-3 0m6.5 11c.7-.7 1.2-1.6 1.5-2.5m-13 0c.3.9.8 1.8 1.5 2.5"/></g>`), g.Group(children),
	)
}

func Bird(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 7h.01M3.4 18H12a8 8 0 0 0 8-8V7a4 4 0 0 0-7.28-2.3L2 20"/><path d="m20 7l2 .5l-2 .5M10 18v3m4-3.25V21m-7-3a6 6 0 0 0 3.84-10.61"/></g>`), g.Group(children),
	)
}

func Bitcoin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.767 19.089c4.924.868 6.14-6.025 1.216-6.894m-1.216 6.894L5.86 18.047m5.908 1.042l-.347 1.97m1.563-8.864c4.924.869 6.14-6.025 1.215-6.893m-1.215 6.893l-3.94-.694m5.155-6.2L8.29 4.26m5.908 1.042l.348-1.97M7.48 20.364l3.126-17.727"/>`), g.Group(children),
	)
}

func Blinds(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3h18m-1 4H8m12 4H8m2 8h10M8 15h12M4 3v14"/><circle cx="4" cy="19" r="2"/></g>`), g.Group(children),
	)
}

func Blocks(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="7" x="14" y="3" rx="1"/><path d="M10 21V8a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1H3"/></g>`), g.Group(children),
	)
}

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10l-5 5V2l5 5L7 17"/>`), g.Group(children),
	)
}

func BluetoothConnected(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10l-5 5V2l5 5L7 17m11-5h3M3 12h3"/>`), g.Group(children),
	)
}

func BluetoothOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 17l-5 5V12l-5 5M2 2l20 20M14.5 9.5L17 7l-5-5v4.5"/>`), g.Group(children),
	)
}

func BluetoothSearching(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10l-5 5V2l5 5L7 17m13.83-2.17a4 4 0 0 0 0-5.66M18 12h.01"/>`), g.Group(children),
	)
}

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 12a4 4 0 0 0 0-8H6v8m9 8a4 4 0 0 0 0-8H6v8Z"/>`), g.Group(children),
	)
}

func Bomb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="13" r="9"/><path d="M14.35 4.65L16.3 2.7a2.41 2.41 0 0 1 3.4 0l1.6 1.6a2.4 2.4 0 0 1 0 3.4l-1.95 1.95M22 2l-1.5 1.5"/></g>`), g.Group(children),
	)
}

func Bone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 10c.7-.7 1.69 0 2.5 0a2.5 2.5 0 1 0 0-5a.5.5 0 0 1-.5-.5a2.5 2.5 0 1 0-5 0c0 .81.7 1.8 0 2.5l-7 7c-.7.7-1.69 0-2.5 0a2.5 2.5 0 0 0 0 5c.28 0 .5.22.5.5a2.5 2.5 0 1 0 5 0c0-.81-.7-1.8 0-2.5Z"/>`), g.Group(children),
	)
}

func Book(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/>`), g.Group(children),
	)
}

func BookA(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/><path d="m8 13l4-7l4 7m-6.9-2h5.7"/></g>`), g.Group(children),
	)
}

func BookAudio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20M8 8v3m4-5v7m4-5v3"/>`), g.Group(children),
	)
}

func BookCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/><path d="m9 9.5l2 2l4-4"/></g>`), g.Group(children),
	)
}

func BookCopy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 16V4a2 2 0 0 1 2-2h11"/><path d="M5 14H4a2 2 0 1 0 0 4h1m17 0H11a2 2 0 1 0 0 4h11V6H11a2 2 0 0 0-2 2v12"/></g>`), g.Group(children),
	)
}

func BookDashed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 22h-2m2-7v2h-2M4 19.5V15m16-7v3m-2-9h2v2M4 11V9m8-7h2m-2 20h2m-2-5h2m-6 5H6.5a2.5 2.5 0 0 1 0-5H8M4 5v-.5A2.5 2.5 0 0 1 6.5 2H8"/>`), g.Group(children),
	)
}

func BookDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20m-8-4V7"/><path d="m9 10l3 3l3-3"/></g>`), g.Group(children),
	)
}

func BookHeadphones(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/><circle cx="9" cy="12" r="1"/><path d="M8 12v-2a4 4 0 0 1 8 0v2"/><circle cx="15" cy="12" r="1"/></g>`), g.Group(children),
	)
}

func BookHeart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/><path d="M16 8.2C16 7 15 6 13.8 6c-.8 0-1.4.3-1.8.9c-.4-.6-1-.9-1.8-.9C9 6 8 7 8 8.2c0 .6.3 1.2.7 1.6h0C10 11.1 12 13 12 13s2-1.9 3.3-3.1h0c.4-.4.7-1 .7-1.7z"/></g>`), g.Group(children),
	)
}

func BookImage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/><circle cx="10" cy="8" r="2"/><path d="m20 13.7l-2.1-2.1c-.8-.8-2-.8-2.8 0L9.7 17"/></g>`), g.Group(children),
	)
}

func BookKey(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H14"/><path d="M20 8v14H6.5a2.5 2.5 0 0 1 0-5H20"/><circle cx="14" cy="8" r="2"/><path d="m20 2l-4.5 4.5M19 3l1 1"/></g>`), g.Group(children),
	)
}

func BookLock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H10"/><path d="M20 15v7H6.5a2.5 2.5 0 0 1 0-5H20"/><rect width="8" height="5" x="12" y="6" rx="1"/><path d="M18 6V4a2 2 0 1 0-4 0v2"/></g>`), g.Group(children),
	)
}

func BookMarked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/><path d="M10 2v8l3-3l3 3V2"/></g>`), g.Group(children),
	)
}

func BookMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20M9 10h6"/>`), g.Group(children),
	)
}

func BookOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2zm20 0h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7z"/>`), g.Group(children),
	)
}

func BookOpenCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 3H2v15h7c1.7 0 3 1.3 3 3V7c0-2.2-1.8-4-4-4Zm8 9l2 2l4-4"/><path d="M22 6V3h-6c-2.2 0-4 1.8-4 4v14c0-1.7 1.3-3 3-3h7v-2.3"/></g>`), g.Group(children),
	)
}

func BookOpenText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 3h6a4 4 0 0 1 4 4v14a3 3 0 0 0-3-3H2zm20 0h-6a4 4 0 0 0-4 4v14a3 3 0 0 1 3-3h7zM6 8h2m-2 4h2m8-4h2m-2 4h2"/>`), g.Group(children),
	)
}

func BookPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20M9 10h6m-3-3v6"/>`), g.Group(children),
	)
}

func BookText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20M8 7h6m-6 4h8"/>`), g.Group(children),
	)
}

func BookType(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/><path d="M16 8V6H8v2m4-2v7m-2 0h4"/></g>`), g.Group(children),
	)
}

func BookUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20m-8-4V7"/><path d="m9 10l3-3l3 3"/></g>`), g.Group(children),
	)
}

func BookUpTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2"/><path d="M18 2h2v20H6.5a2.5 2.5 0 0 1 0-5H20m-8-4V7"/><path d="m9 10l3-3l3 3M9 5l3-3l3 3"/></g>`), g.Group(children),
	)
}

func BookUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20"/><circle cx="12" cy="8" r="2"/><path d="M15 13a3 3 0 1 0-6 0"/></g>`), g.Group(children),
	)
}

func BookX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19.5v-15A2.5 2.5 0 0 1 6.5 2H20v20H6.5a2.5 2.5 0 0 1 0-5H20M14.5 7l-5 5m0-5l5 5"/>`), g.Group(children),
	)
}

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 21l-7-4l-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16z"/>`), g.Group(children),
	)
}

func BookmarkCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m19 21l-7-4l-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2Z"/><path d="m9 10l2 2l4-4"/></g>`), g.Group(children),
	)
}

func BookmarkMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 21l-7-4l-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16zm-4-11H9"/>`), g.Group(children),
	)
}

func BookmarkPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 21l-7-4l-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16zM12 7v6m3-3H9"/>`), g.Group(children),
	)
}

func BookmarkX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 21l-7-4l-7 4V5a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2ZM14.5 7.5l-5 5m0-5l5 5"/>`), g.Group(children),
	)
}

func BoomBox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 9V5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v4M8 8v1m4-1v1m4-1v1"/><rect width="20" height="12" x="2" y="9" rx="2"/><circle cx="8" cy="15" r="2"/><circle cx="16" cy="15" r="2"/></g>`), g.Group(children),
	)
}

func Bot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 8V4H8"/><rect width="16" height="12" x="4" y="8" rx="2"/><path d="M2 14h2m16 0h2m-7-1v2m-6-2v2"/></g>`), g.Group(children),
	)
}

func Box(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z"/><path d="m3.3 7l8.7 5l8.7-5M12 22V12"/></g>`), g.Group(children),
	)
}

func BoxSelect(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3a2 2 0 0 0-2 2m16-2a2 2 0 0 1 2 2m0 14a2 2 0 0 1-2 2M5 21a2 2 0 0 1-2-2M9 3h1M9 21h1m4-18h1m-1 18h1M3 9v1m18-1v1M3 14v1m18-1v1"/>`), g.Group(children),
	)
}

func Boxes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2.97 12.92A2 2 0 0 0 2 14.63v3.24a2 2 0 0 0 .97 1.71l3 1.8a2 2 0 0 0 2.06 0L12 19v-5.5l-5-3l-4.03 2.42ZM7 16.5l-4.74-2.85M7 16.5l5-3m-5 3v5.17m5-8.17V19l3.97 2.38a2 2 0 0 0 2.06 0l3-1.8a2 2 0 0 0 .97-1.71v-3.24a2 2 0 0 0-.97-1.71L17 10.5l-5 3Zm5 3l-5-3m5 3l4.74-2.85M17 16.5v5.17"/><path d="M7.97 4.42A2 2 0 0 0 7 6.13v4.37l5 3l5-3V6.13a2 2 0 0 0-.97-1.71l-3-1.8a2 2 0 0 0-2.06 0l-3 1.8ZM12 8L7.26 5.15M12 8l4.74-2.85M12 13.5V8"/></g>`), g.Group(children),
	)
}

func Braces(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3H7a2 2 0 0 0-2 2v5a2 2 0 0 1-2 2a2 2 0 0 1 2 2v5c0 1.1.9 2 2 2h1m8 0h1a2 2 0 0 0 2-2v-5c0-1.1.9-2 2-2a2 2 0 0 1-2-2V5a2 2 0 0 0-2-2h-1"/>`), g.Group(children),
	)
}

func Brackets(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 3h3v18h-3m-8 0H5V3h3"/>`), g.Group(children),
	)
}

func Brain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 2A2.5 2.5 0 0 1 12 4.5v15a2.5 2.5 0 0 1-4.96.44a2.5 2.5 0 0 1-2.96-3.08a3 3 0 0 1-.34-5.58a2.5 2.5 0 0 1 1.32-4.24a2.5 2.5 0 0 1 1.98-3A2.5 2.5 0 0 1 9.5 2Zm5 0A2.5 2.5 0 0 0 12 4.5v15a2.5 2.5 0 0 0 4.96.44a2.5 2.5 0 0 0 2.96-3.08a3 3 0 0 0 .34-5.58a2.5 2.5 0 0 0-1.32-4.24a2.5 2.5 0 0 0-1.98-3A2.5 2.5 0 0 0 14.5 2Z"/>`), g.Group(children),
	)
}

func BrainCircuit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4.5a2.5 2.5 0 0 0-4.96-.46a2.5 2.5 0 0 0-1.98 3a2.5 2.5 0 0 0-1.32 4.24a3 3 0 0 0 .34 5.58a2.5 2.5 0 0 0 2.96 3.08a2.5 2.5 0 0 0 4.91.05L12 20V4.5ZM16 8V5c0-1.1.9-2 2-2m-6 10h4"/><path d="M12 18h6a2 2 0 0 1 2 2v1M12 8h8m.5 0a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-4 5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/><path d="M20.5 21a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Zm-2-18a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0Z"/></g>`), g.Group(children),
	)
}

func BrainCog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="3"/><path d="M12 4.5a2.5 2.5 0 0 0-4.96-.46a2.5 2.5 0 0 0-1.98 3a2.5 2.5 0 0 0-1.32 4.24a3 3 0 0 0 .34 5.58a2.5 2.5 0 0 0 2.96 3.08A2.5 2.5 0 0 0 12 19.5a2.5 2.5 0 0 0 4.96.44a2.5 2.5 0 0 0 2.96-3.08a3 3 0 0 0 .34-5.58a2.5 2.5 0 0 0-1.32-4.24a2.5 2.5 0 0 0-1.98-3A2.5 2.5 0 0 0 12 4.5m3.7 5.9l-.9.4m-5.6 2.4l-.9.4m5.3 2.1l-.4-.9m-2.4-5.6l-.4-.9m5.3 5.2l-.9-.4m-5.6-2.2l-.9-.4m2.2 5.2l.4-.9m2.2-5.6l.4-.9"/></g>`), g.Group(children),
	)
}

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="14" x="2" y="7" rx="2" ry="2"/><path d="M16 21V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v16"/></g>`), g.Group(children),
	)
}

func BringToFront(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="8" x="8" y="8" rx="2"/><path d="M4 10a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2m4 16a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2"/></g>`), g.Group(children),
	)
}

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9.06 11.9l8.07-8.06a2.85 2.85 0 1 1 4.03 4.03l-8.06 8.08m-6.03-1.01c-1.66 0-3 1.35-3 3.02c0 1.33-2.5 1.52-2 2.02c1.08 1.1 2.49 2.02 4 2.02c2.2 0 4-1.8 4-4.04a3.01 3.01 0 0 0-3-3.02z"/>`), g.Group(children),
	)
}

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 2l1.88 1.88m4.24 0L16 2M9 7.13v-1a3.003 3.003 0 1 1 6 0v1"/><path d="M12 20c-3.3 0-6-2.7-6-6v-3a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v3c0 3.3-2.7 6-6 6m0 0v-9"/><path d="M6.53 9C4.6 8.8 3 7.1 3 5m3 8H2m1 8c0-2.1 1.7-3.9 3.8-4M20.97 5c0 2.1-1.6 3.8-3.5 4M22 13h-4m-.8 4c2.1.1 3.8 1.9 3.8 4"/></g>`), g.Group(children),
	)
}

func BugOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 7.13V6a3 3 0 0 0-5.14-2.1L8 2m6.12 1.88L16 2"/><path d="M22 13h-4v-2a4 4 0 0 0-4-4h-1.3"/><path d="M20.97 5c0 2.1-1.6 3.8-3.5 4M2 2l20 20M7.7 7.7A4 4 0 0 0 6 11v3a6 6 0 0 0 11.13 3.13M12 20v-8m-6 1H2"/><path d="M3 21c0-2.1 1.7-3.9 3.8-4"/></g>`), g.Group(children),
	)
}

func BugPlay(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 2l1.88 1.88m4.24 0L16 2M9 7.13v-1a3.003 3.003 0 1 1 6 0v1"/><path d="M18 11a4 4 0 0 0-4-4h-4a4 4 0 0 0-4 4v3a6.1 6.1 0 0 0 2 4.5"/><path d="M6.53 9C4.6 8.8 3 7.1 3 5m3 8H2m1 8c0-2.1 1.7-3.9 3.8-4M20.97 5c0 2.1-1.6 3.8-3.5 4M12 12l8 5l-8 5Z"/></g>`), g.Group(children),
	)
}

func Building(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="20" x="4" y="2" rx="2" ry="2"/><path d="M9 22v-4h6v4M8 6h.01M16 6h.01M12 6h.01M12 10h.01M12 14h.01M16 10h.01M16 14h.01M8 10h.01M8 14h.01"/></g>`), g.Group(children),
	)
}

func BuildingTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 22V4a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v18Zm0-10H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h2M18 9h2a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2h-2M10 6h4m-4 4h4m-4 4h4m-4 4h4"/>`), g.Group(children),
	)
}

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 6v6m7-6v6M2 12h19.6M18 18h3s.5-1.7.8-2.8c.1-.4.2-.8.2-1.2c0-.4-.1-.8-.2-1.2l-1.4-5C20.1 6.8 19.1 6 18 6H4a2 2 0 0 0-2 2v10h3"/><circle cx="7" cy="18" r="2"/><path d="M9 18h5"/><circle cx="16" cy="18" r="2"/></g>`), g.Group(children),
	)
}

func BusFront(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 6L2 7m8-1h4m8 1l-2-1"/><rect width="16" height="16" x="4" y="3" rx="2"/><path d="M4 11h16M8 15h.01M16 15h.01M6 19v2m12 0v-2"/></g>`), g.Group(children),
	)
}

func Cable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 9a2 2 0 0 1-2-2V5h6v2a2 2 0 0 1-2 2ZM3 5V3m4 2V3"/><path d="M19 15V6.5a3.5 3.5 0 0 0-7 0v11a3.5 3.5 0 0 1-7 0V9m12 12v-2m4 2v-2"/><path d="M22 19h-6v-2a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2Z"/></g>`), g.Group(children),
	)
}

func CableCar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 3h.01M14 2h.01M2 9l20-5m-10 8V6.5"/><rect width="16" height="10" x="4" y="12" rx="3"/><path d="M9 12v5m6-5v5M4 17h16"/></g>`), g.Group(children),
	)
}

func Cake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 21v-8a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v8"/><path d="M4 16s.5-1 2-1s2.5 2 4 2s2.5-2 4-2s2.5 2 4 2s2-1 2-1M2 21h20M7 8v2m5-2v2m5-2v2M7 4h.01M12 4h.01M17 4h.01"/></g>`), g.Group(children),
	)
}

func CakeSlice(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="7" r="2"/><path d="M7.2 7.9L3 11v9c0 .6.4 1 1 1h16c.6 0 1-.4 1-1v-9c0-2-3-6-7-8l-3.6 2.6M16 13H3m13 4H3"/></g>`), g.Group(children),
	)
}

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="20" x="4" y="2" rx="2"/><path d="M8 6h8m0 8v4m0-8h.01M12 10h.01M8 10h.01M12 14h.01M8 14h.01M12 18h.01M8 18h.01"/></g>`), g.Group(children),
	)
}

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><path d="M16 2v4M8 2v4m-5 4h18"/></g>`), g.Group(children),
	)
}

func CalendarCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><path d="M16 2v4M8 2v4m-5 4h18M9 16l2 2l4-4"/></g>`), g.Group(children),
	)
}

func CalendarCheckTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 14V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8m3-20v4M8 2v4m-5 4h18m-5 10l2 2l4-4"/>`), g.Group(children),
	)
}

func CalendarClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 7.5V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h3.5M16 2v4M8 2v4m-5 4h5m9.5 7.5L16 16.25V14"/><path d="M22 16a6 6 0 1 1-12 0a6 6 0 0 1 12 0Z"/></g>`), g.Group(children),
	)
}

func CalendarDays(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><path d="M16 2v4M8 2v4m-5 4h18M8 14h.01M12 14h.01M16 14h.01M8 18h.01M12 18h.01M16 18h.01"/></g>`), g.Group(children),
	)
}

func CalendarHeart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 10V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h7m4-20v4M8 2v4m-5 4h18"/><path d="M21.29 14.7a2.43 2.43 0 0 0-2.65-.52c-.3.12-.57.3-.8.53l-.34.34l-.35-.34a2.43 2.43 0 0 0-2.65-.53c-.3.12-.56.3-.79.53c-.95.94-1 2.53.2 3.74L17.5 22l3.6-3.55c1.2-1.21 1.14-2.8.19-3.74Z"/></g>`), g.Group(children),
	)
}

func CalendarMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8m3-20v4M8 2v4m-5 4h18m-5 9h6"/>`), g.Group(children),
	)
}

func CalendarOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.18 4.18A2 2 0 0 0 3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 1.82-1.18M21 15.5V6a2 2 0 0 0-2-2H9.5M16 2v4M3 10h7m11 0h-5.5M2 2l20 20"/>`), g.Group(children),
	)
}

func CalendarPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8m3-20v4M8 2v4m-5 4h18m-2 6v6m-3-3h6"/>`), g.Group(children),
	)
}

func CalendarRange(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><path d="M16 2v4M8 2v4m-5 4h18m-4 4h-6m2 4H7m0-4h.01M17 18h.01"/></g>`), g.Group(children),
	)
}

func CalendarSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h7.5M16 2v4M8 2v4m-5 4h18"/><path d="M18 21a3 3 0 1 0 0-6a3 3 0 0 0 0 6v0Zm4 1l-1.5-1.5"/></g>`), g.Group(children),
	)
}

func CalendarX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="4" rx="2" ry="2"/><path d="M16 2v4M8 2v4m-5 4h18m-11 4l4 4m0-4l-4 4"/></g>`), g.Group(children),
	)
}

func CalendarXTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 13V6a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h8m3-20v4M8 2v4m-5 4h18m-4 7l5 5m-5 0l5-5"/>`), g.Group(children),
	)
}

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 4h-5L7 7H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-3l-2.5-3z"/><circle cx="12" cy="13" r="3"/></g>`), g.Group(children),
	)
}

func CameraOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 2l20 20M7 7H4a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h16M9.5 4h5L17 7h3a2 2 0 0 1 2 2v7.5"/><path d="M14.121 15.121A3 3 0 1 1 9.88 10.88"/></g>`), g.Group(children),
	)
}

func CandlestickChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 5v4"/><rect width="4" height="6" x="7" y="9" rx="1"/><path d="M9 15v2m8-14v2"/><rect width="4" height="8" x="15" y="5" rx="1"/><path d="M17 13v3M3 3v18h18"/></g>`), g.Group(children),
	)
}

func Candy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9.5 7.5l-2 2a4.95 4.95 0 1 0 7 7l2-2a4.95 4.95 0 1 0-7-7Zm4.5-1v10m-4-9v10"/><path d="m16 7l1-5l1.37.68A3 3 0 0 0 19.7 3H21v1.3c0 .46.1.92.32 1.33L22 7l-5 1m-9 9l-1 5l-1.37-.68A3 3 0 0 0 4.3 21H3v-1.3a3 3 0 0 0-.32-1.33L2 17l5-1"/></g>`), g.Group(children),
	)
}

func CandyCane(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.7 21a2 2 0 0 1-3.5-2l8.6-14a6 6 0 0 1 10.4 6a2 2 0 1 1-3.464-2a2 2 0 1 0-3.464-2ZM17.75 7L15 2.1m-4.1 2.7L13 9m-5.1.7l2 4.4m-5 .6L7 18.9"/>`), g.Group(children),
	)
}

func CandyOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8.5 8.5l-1 1a4.95 4.95 0 0 0 7 7l1-1m-3.657-9.313A4.947 4.947 0 0 1 16.5 7.5a4.947 4.947 0 0 1 1.313 4.657M14 16.5V14m0-7.5v1.843M10 10v7.5"/><path d="m16 7l1-5l1.367.683A3 3 0 0 0 19.708 3H21v1.292a3 3 0 0 0 .317 1.341L22 7l-5 1m-9 9l-1 5l-1.367-.683A3 3 0 0 0 4.292 21H3v-1.292a3 3 0 0 0-.317-1.341L2 17l5-1M2 2l20 20"/></g>`), g.Group(children),
	)
}

func Car(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 17h2c.6 0 1-.4 1-1v-3c0-.9-.7-1.7-1.5-1.9C18.7 10.6 16 10 16 10s-1.3-1.4-2.2-2.3c-.5-.4-1.1-.7-1.8-.7H5c-.6 0-1.1.4-1.4.9l-1.4 2.9A3.7 3.7 0 0 0 2 12v4c0 .6.4 1 1 1h2"/><circle cx="7" cy="17" r="2"/><path d="M9 17h6"/><circle cx="17" cy="17" r="2"/></g>`), g.Group(children),
	)
}

func CarFront(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m21 8l-2 2l-1.5-3.7A2 2 0 0 0 15.646 5H8.4a2 2 0 0 0-1.903 1.257L5 10L3 8m4 6h.01M17 14h.01"/><rect width="18" height="8" x="3" y="10" rx="2"/><path d="M5 18v2m14-2v2"/></g>`), g.Group(children),
	)
}

func CarTaxiFront(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 2h4m7 6l-2 2l-1.5-3.7A2 2 0 0 0 15.646 5H8.4a2 2 0 0 0-1.903 1.257L5 10L3 8m4 6h.01M17 14h.01"/><rect width="18" height="8" x="3" y="10" rx="2"/><path d="M5 18v2m14-2v2"/></g>`), g.Group(children),
	)
}

func Caravan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 9h4v4H2zm8 0h4v10h-4z"/><path d="M18 19V9a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v8a2 2 0 0 0 2 2h2"/><circle cx="8" cy="19" r="2"/><path d="M10 19h12v-2"/></g>`), g.Group(children),
	)
}

func Carrot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2.27 21.7s9.87-3.5 12.73-6.36a4.5 4.5 0 0 0-6.36-6.37C5.77 11.84 2.27 21.7 2.27 21.7zM8.64 14l-2.05-2.04M15.34 15l-2.46-2.46"/><path d="M22 9s-1.33-2-3.5-2C16.86 7 15 9 15 9s1.33 2 3.5 2S22 9 22 9z"/><path d="M15 2s-2 1.33-2 3.5S15 9 15 9s2-1.84 2-3.5C17 3.33 15 2 15 2z"/></g>`), g.Group(children),
	)
}

func CaseLower(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7" cy="12" r="3"/><path d="M10 9v6"/><circle cx="17" cy="12" r="3"/><path d="M14 7v8"/></g>`), g.Group(children),
	)
}

func CaseSensitive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 15l4-8l4 8m-7-2h6"/><circle cx="18" cy="12" r="3"/><path d="M21 9v6"/></g>`), g.Group(children),
	)
}

func CaseUpper(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 15l4-8l4 8m-7-2h6m5-2h4.5a2 2 0 0 1 0 4H15V7h4a2 2 0 0 1 0 4"/>`), g.Group(children),
	)
}

func CassetteTape(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><circle cx="8" cy="10" r="2"/><path d="M8 12h8"/><circle cx="16" cy="10" r="2"/><path d="m6 20l.7-2.9A1.4 1.4 0 0 1 8.1 16h7.8a1.4 1.4 0 0 1 1.4 1l.7 3"/></g>`), g.Group(children),
	)
}

func Cast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 8V6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2h-6M2 12a9 9 0 0 1 8 8m-8-4a5 5 0 0 1 4 4m-4 0h.01"/>`), g.Group(children),
	)
}

func Castle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 20v-9H2v9a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2Zm-4-9V4H6v7"/><path d="M15 22v-4a3 3 0 0 0-3-3v0a3 3 0 0 0-3 3v4m13-11V9M2 11V9m4-5V2m12 2V2m-8 2V2m4 2V2"/></g>`), g.Group(children),
	)
}

func Cat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 5c.67 0 1.35.09 2 .26c1.78-2 5.03-2.84 6.42-2.26c1.4.58-.42 7-.42 7c.57 1.07 1 2.24 1 3.44C21 17.9 16.97 21 12 21s-9-3-9-7.56c0-1.25.5-2.4 1-3.44c0 0-1.89-6.42-.5-7c1.39-.58 4.72.23 6.5 2.23A9.04 9.04 0 0 1 12 5Zm-4 9v.5m8-.5v.5"/><path d="M11.25 16.25h1.5L12 17l-.75-.75Z"/></g>`), g.Group(children),
	)
}

func Check(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 6L9 17l-5-5"/>`), g.Group(children),
	)
}

func CheckCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 6L7 17l-5-5m20-2l-7.5 7.5L13 16"/>`), g.Group(children),
	)
}

func CheckCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 11.08V12a10 10 0 1 1-5.93-9.14"/><path d="M22 4L12 14.01l-3-3"/></g>`), g.Group(children),
	)
}

func CheckCircleTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22c5.523 0 10-4.477 10-10S17.523 2 12 2S2 6.477 2 12s4.477 10 10 10z"/><path d="m9 12l2 2l4-4"/></g>`), g.Group(children),
	)
}

func CheckSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 11l3 3L22 4"/><path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11"/></g>`), g.Group(children),
	)
}

func ChefHat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 13.87A4 4 0 0 1 7.41 6a5.11 5.11 0 0 1 1.05-1.54a5 5 0 0 1 7.08 0A5.11 5.11 0 0 1 16.59 6A4 4 0 0 1 18 13.87V21H6ZM6 17h12"/>`), g.Group(children),
	)
}

func Cherry(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 17a5 5 0 0 0 10 0c0-2.76-2.5-5-5-3c-2.5-2-5 .24-5 3Zm10 0a5 5 0 0 0 10 0c0-2.76-2.5-5-5-3c-2.5-2-5 .24-5 3Z"/><path d="M7 14c3.22-2.91 4.29-8.75 5-12c1.66 2.38 4.94 9 5 12"/><path d="M22 9c-4.29 0-7.14-2.33-10-7c5.71 0 10 4.67 10 7Z"/></g>`), g.Group(children),
	)
}

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 9l6 6l6-6"/>`), g.Group(children),
	)
}

func ChevronDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m16 10l-4 4l-4-4"/></g>`), g.Group(children),
	)
}

func ChevronDownSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m16 10l-4 4l-4-4"/></g>`), g.Group(children),
	)
}

func ChevronFirst(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 18l-6-6l6-6M7 6v12"/>`), g.Group(children),
	)
}

func ChevronLast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 18l6-6l-6-6m10 0v12"/>`), g.Group(children),
	)
}

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 18l-6-6l6-6"/>`), g.Group(children),
	)
}

func ChevronLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m14 16l-4-4l4-4"/></g>`), g.Group(children),
	)
}

func ChevronLeftSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m14 16l-4-4l4-4"/></g>`), g.Group(children),
	)
}

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 18l6-6l-6-6"/>`), g.Group(children),
	)
}

func ChevronRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m10 8l4 4l-4 4"/></g>`), g.Group(children),
	)
}

func ChevronRightSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m10 8l4 4l-4 4"/></g>`), g.Group(children),
	)
}

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 15l-6-6l-6 6"/>`), g.Group(children),
	)
}

func ChevronUpCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m8 14l4-4l4 4"/></g>`), g.Group(children),
	)
}

func ChevronUpSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m8 14l4-4l4 4"/></g>`), g.Group(children),
	)
}

func ChevronsDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 6l5 5l5-5M7 13l5 5l5-5"/>`), g.Group(children),
	)
}

func ChevronsDownUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 20l5-5l5 5M7 4l5 5l5-5"/>`), g.Group(children),
	)
}

func ChevronsLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 17l-5-5l5-5m7 10l-5-5l5-5"/>`), g.Group(children),
	)
}

func ChevronsLeftRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 7l-5 5l5 5m6-10l5 5l-5 5"/>`), g.Group(children),
	)
}

func ChevronsRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 17l5-5l-5-5m7 10l5-5l-5-5"/>`), g.Group(children),
	)
}

func ChevronsRightLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m20 17l-5-5l5-5M4 17l5-5l-5-5"/>`), g.Group(children),
	)
}

func ChevronsUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 11l-5-5l-5 5m10 7l-5-5l-5 5"/>`), g.Group(children),
	)
}

func ChevronsUpDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 15l5 5l5-5M7 9l5-5l5 5"/>`), g.Group(children),
	)
}

func Chrome(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="4"/><path d="M21.17 8H12M3.95 6.06L8.54 14m2.34 7.94L15.46 14"/></g>`), g.Group(children),
	)
}

func Church(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m18 7l4 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9l4-2"/><path d="M14 22v-4a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v4"/><path d="M18 22V5l-6-3l-6 3v17m6-15v5m-2-3h4"/></g>`), g.Group(children),
	)
}

func Cigarette(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 12H2v4h16m4-4v4M7 12v4m11-8c0-2.5-2-2.5-2-5m6 5c0-2.5-2-2.5-2-5"/>`), g.Group(children),
	)
}

func CigaretteOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20M12 12H2v4h14m6-4v4m-4-4h-.5M7 12v4m11-8c0-2.5-2-2.5-2-5m6 5c0-2.5-2-2.5-2-5"/>`), g.Group(children),
	)
}

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<circle cx="12" cy="12" r="10" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>`), g.Group(children),
	)
}

func CircleDashed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.1 2.18a9.93 9.93 0 0 1 3.8 0m3.7 1.53a9.95 9.95 0 0 1 2.69 2.7m1.53 3.69a9.93 9.93 0 0 1 0 3.8m-1.53 3.7a9.95 9.95 0 0 1-2.7 2.69m-3.69 1.53a9.94 9.94 0 0 1-3.8 0m-3.7-1.53a9.95 9.95 0 0 1-2.69-2.7M2.18 13.9a9.93 9.93 0 0 1 0-3.8m1.53-3.7a9.95 9.95 0 0 1 2.7-2.69"/>`), g.Group(children),
	)
}

func CircleDollarSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8m4 2V6"/></g>`), g.Group(children),
	)
}

func CircleDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="1"/></g>`), g.Group(children),
	)
}

func CircleDotDashed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10.1 2.18a9.93 9.93 0 0 1 3.8 0m3.7 1.53a9.95 9.95 0 0 1 2.69 2.7m1.53 3.69a9.93 9.93 0 0 1 0 3.8m-1.53 3.7a9.95 9.95 0 0 1-2.7 2.69m-3.69 1.53a9.94 9.94 0 0 1-3.8 0m-3.7-1.53a9.95 9.95 0 0 1-2.69-2.7M2.18 13.9a9.93 9.93 0 0 1 0-3.8m1.53-3.7a9.95 9.95 0 0 1 2.7-2.69"/><circle cx="12" cy="12" r="1"/></g>`), g.Group(children),
	)
}

func CircleEllipsis(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M17 12h.01M12 12h.01M7 12h.01"/></g>`), g.Group(children),
	)
}

func CircleEqual(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 10h10M7 14h10"/><circle cx="12" cy="12" r="10"/></g>`), g.Group(children),
	)
}

func CircleOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20M8.35 2.69A10 10 0 0 1 21.3 15.65m-2.22 3.43A10 10 0 1 1 4.92 4.92"/>`), g.Group(children),
	)
}

func CircleSlash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 15l6-6"/><circle cx="12" cy="12" r="10"/></g>`), g.Group(children),
	)
}

func CircleSlashTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M22 2L2 22"/></g>`), g.Group(children),
	)
}

func CircuitBoard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M11 9h4a2 2 0 0 0 2-2V3"/><circle cx="9" cy="9" r="2"/><path d="M7 21v-4a2 2 0 0 1 2-2h4"/><circle cx="15" cy="15" r="2"/></g>`), g.Group(children),
	)
}

func Citrus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.66 17.67a1.08 1.08 0 0 1-.04 1.6A12 12 0 0 1 4.73 2.38a1.1 1.1 0 0 1 1.61-.04z"/><path d="M19.65 15.66A8 8 0 0 1 8.35 4.34M14 10l-5.5 5.5"/><path d="M14 17.85V10H6.15"/></g>`), g.Group(children),
	)
}

func Clapperboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.2 6L3 11l-.9-2.4c-.3-1.1.3-2.2 1.3-2.5l13.5-4c1.1-.3 2.2.3 2.5 1.3Zm-14-.7l3.1 3.9m3.1-5.8l3.1 4M3 11h18v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2Z"/>`), g.Group(children),
	)
}

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/></g>`), g.Group(children),
	)
}

func ClipboardCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><path d="m9 14l2 2l4-4"/></g>`), g.Group(children),
	)
}

func ClipboardCopy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-2M16 4h2a2 2 0 0 1 2 2v4m1 4H11"/><path d="m15 10l-4 4l4 4"/></g>`), g.Group(children),
	)
}

func ClipboardEdit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M10.42 12.61a2.1 2.1 0 1 1 2.97 2.97L7.95 21L4 22l.99-3.95l5.43-5.44Z"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-5.5M4 13.5V6a2 2 0 0 1 2-2h2"/></g>`), g.Group(children),
	)
}

func ClipboardList(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2m4 7h4m-4 5h4m-8-5h.01M8 16h.01"/></g>`), g.Group(children),
	)
}

func ClipboardPaste(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 2H9a1 1 0 0 0-1 1v2c0 .6.4 1 1 1h6c.6 0 1-.4 1-1V3c0-.6-.4-1-1-1Z"/><path d="M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2M16 4h2a2 2 0 0 1 2 2v2m-9 6h10"/><path d="m17 10l4 4l-4 4"/></g>`), g.Group(children),
	)
}

func ClipboardSignature(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M8 4H6a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-.5M16 4h2a2 2 0 0 1 1.73 1"/><path d="M18.42 9.61a2.1 2.1 0 1 1 2.97 2.97L16.95 17L13 18l.99-3.95l4.43-4.44ZM8 18h1"/></g>`), g.Group(children),
	)
}

func ClipboardType(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2"/><path d="M9 12v-1h6v1m-4 5h2m-1-6v6"/></g>`), g.Group(children),
	)
}

func ClipboardX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="4" x="8" y="2" rx="1" ry="1"/><path d="M16 4h2a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h2m7 7l-6 6m0-6l6 6"/></g>`), g.Group(children),
	)
}

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l4 2"/></g>`), g.Group(children),
	)
}

func ClockEight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l-4 2"/></g>`), g.Group(children),
	)
}

func ClockEleven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6L9.5 8"/></g>`), g.Group(children),
	)
}

func ClockFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l2.5 4"/></g>`), g.Group(children),
	)
}

func ClockFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l4 2"/></g>`), g.Group(children),
	)
}

func ClockNine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6H7.5"/></g>`), g.Group(children),
	)
}

func ClockOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l2.5-4"/></g>`), g.Group(children),
	)
}

func ClockSeven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l-2.5 4"/></g>`), g.Group(children),
	)
}

func ClockSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v10.5"/></g>`), g.Group(children),
	)
}

func ClockTen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l-4-2"/></g>`), g.Group(children),
	)
}

func ClockThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6h4.5"/></g>`), g.Group(children),
	)
}

func ClockTwelve(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6"/></g>`), g.Group(children),
	)
}

func ClockTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 6v6l4-2"/></g>`), g.Group(children),
	)
}

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 19H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z"/>`), g.Group(children),
	)
}

func CloudCog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="17" r="3"/><path d="M4.2 15.1A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.2m-4.3 2.2l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.2-5.6l.3-.9m.2 7.4l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4"/></g>`), g.Group(children),
	)
}

func CloudDrizzle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M8 19v1m0-6v1m8 4v1m0-6v1m-4 6v1m0-6v1"/>`), g.Group(children),
	)
}

func CloudFog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M16 17H7m10 4H9"/>`), g.Group(children),
	)
}

func CloudHail(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M16 14v2m-8-2v2m8 4h.01M8 20h.01M12 16v2m0 4h.01"/>`), g.Group(children),
	)
}

func CloudLightning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 16.326A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 .5 8.973"/><path d="m13 12l-3 5h4l-3 5"/></g>`), g.Group(children),
	)
}

func CloudMoon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 16a3 3 0 1 1 0 6H7a5 5 0 1 1 4.9-6Zm-2.9-7A6 6 0 0 1 16 4a4.24 4.24 0 0 0 6 6a6 6 0 0 1-3 5.197"/>`), g.Group(children),
	)
}

func CloudMoonRain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.083 9A6.002 6.002 0 0 1 16 4a4.243 4.243 0 0 0 6 6c0 2.22-1.206 4.16-3 5.197M3 20a5 5 0 1 1 8.9-4H13a3 3 0 0 1 2 5.24M11 20v2m-4-3v2"/>`), g.Group(children),
	)
}

func CloudOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20M5.782 5.782A7 7 0 0 0 9 19h8.5a4.5 4.5 0 0 0 1.307-.193m2.725-2.307A4.5 4.5 0 0 0 17.5 10h-1.79A7.008 7.008 0 0 0 10 5.07"/>`), g.Group(children),
	)
}

func CloudRain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M16 14v6m-8-6v6m4-4v6"/>`), g.Group(children),
	)
}

func CloudRainWind(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M9.2 22l3-7M9 13l-3 7m11-7l-3 7"/>`), g.Group(children),
	)
}

func CloudSnow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M8 15h.01M8 19h.01M12 17h.01M12 21h.01M16 15h.01M16 19h.01"/>`), g.Group(children),
	)
}

func CloudSun(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v2m-7.07.93l1.41 1.41M20 12h2m-2.93-7.07l-1.41 1.41m-1.713 6.31a4 4 0 0 0-5.925-4.128M13 22H7a5 5 0 1 1 4.9-6H13a3 3 0 0 1 0 6Z"/>`), g.Group(children),
	)
}

func CloudSunRain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v2m-7.07.93l1.41 1.41M20 12h2m-2.93-7.07l-1.41 1.41m-1.713 6.31a4 4 0 0 0-5.925-4.128M3 20a5 5 0 1 1 8.9-4H13a3 3 0 0 1 2 5.24M11 20v2m-4-3v2"/>`), g.Group(children),
	)
}

func Cloudy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 21H9a7 7 0 1 1 6.71-9h1.79a4.5 4.5 0 1 1 0 9Z"/><path d="M22 10a3 3 0 0 0-3-3h-2.207a5.502 5.502 0 0 0-10.702.5"/></g>`), g.Group(children),
	)
}

func Clover(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16.2 3.8a2.7 2.7 0 0 0-3.81 0l-.4.38l-.4-.4a2.7 2.7 0 0 0-3.82 0C6.73 4.85 6.67 6.64 8 8l4 4l4-4c1.33-1.36 1.27-3.15.2-4.2zM8 8c-1.36-1.33-3.15-1.27-4.2-.2a2.7 2.7 0 0 0 0 3.81l.38.4l-.4.4a2.7 2.7 0 0 0 0 3.82C4.85 17.27 6.64 17.33 8 16m8 0c1.36 1.33 3.15 1.27 4.2.2a2.7 2.7 0 0 0 0-3.81l-.38-.4l.4-.4a2.7 2.7 0 0 0 0-3.82C19.15 6.73 17.36 6.67 16 8"/><path d="M7.8 20.2a2.7 2.7 0 0 0 3.81 0l.4-.38l.4.4a2.7 2.7 0 0 0 3.82 0c1.06-1.06 1.12-2.85-.21-4.21l-4-4l-4 4c-1.33 1.36-1.27 3.15-.2 4.2zM7 17l-5 5"/></g>`), g.Group(children),
	)
}

func Club(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.28 9.05a5.5 5.5 0 1 0-10.56 0A5.5 5.5 0 1 0 12 17.66a5.5 5.5 0 1 0 5.28-8.6ZM12 17.66V22"/>`), g.Group(children),
	)
}

func Code(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 18l6-6l-6-6M8 6l-6 6l6 6"/>`), g.Group(children),
	)
}

func CodeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 16l4-4l-4-4M6 8l-4 4l4 4m8.5-12l-5 16"/>`), g.Group(children),
	)
}

func Codepen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 2l10 6.5v7L12 22L2 15.5v-7L12 2zm0 20v-6.5"/><path d="m22 8.5l-10 7l-10-7"/><path d="m2 15.5l10-7l10 7M12 2v6.5"/></g>`), g.Group(children),
	)
}

func Codesandbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/><path d="m7.5 4.21l4.5 2.6l4.5-2.6m-9 15.58V14.6L3 12m18 0l-4.5 2.6v5.19M3.27 6.96L12 12.01l8.73-5.05M12 22.08V12"/></g>`), g.Group(children),
	)
}

func Coffee(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8h1a4 4 0 1 1 0 8h-1M3 8h14v9a4 4 0 0 1-4 4H7a4 4 0 0 1-4-4Zm3-6v2m4-2v2m4-2v2"/>`), g.Group(children),
	)
}

func Cog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 20a8 8 0 1 0 0-16a8 8 0 0 0 0 16Z"/><path d="M12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4Zm0-12v2m0 18v-2m5 .66l-1-1.73m-5-8.66L7 3.34M20.66 17l-1.73-1M3.34 7l1.73 1M14 12h8M2 12h2m16.66-5l-1.73 1M3.34 17l1.73-1M17 3.34l-1 1.73m-5 8.66l-4 6.93"/></g>`), g.Group(children),
	)
}

func Coins(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="8" cy="8" r="6"/><path d="M18.09 10.37A6 6 0 1 1 10.34 18M7 6h1v4"/><path d="m16.71 13.88l.7.71l-2.82 2.82"/></g>`), g.Group(children),
	)
}

func Columns(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M12 3v18"/></g>`), g.Group(children),
	)
}

func Combine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="8" x="2" y="2" rx="2"/><path d="M14 2c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2m6-8c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2m-10 8H5c-1.7 0-3-1.3-3-3v-1"/><path d="m7 21l3-3l-3-3"/><rect width="8" height="8" x="14" y="14" rx="2"/></g>`), g.Group(children),
	)
}

func Command(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 6v12a3 3 0 1 0 3-3H6a3 3 0 1 0 3 3V6a3 3 0 1 0-3 3h12a3 3 0 1 0-3-3"/>`), g.Group(children),
	)
}

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m16.24 7.76l-2.12 6.36l-6.36 2.12l2.12-6.36l6.36-2.12z"/></g>`), g.Group(children),
	)
}

func Component(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.5 8.5L9 12l-3.5 3.5L2 12l3.5-3.5ZM12 2l3.5 3.5L12 9L8.5 5.5L12 2Zm6.5 6.5L22 12l-3.5 3.5L15 12l3.5-3.5ZM12 15l3.5 3.5L12 22l-3.5-3.5L12 15Z"/>`), g.Group(children),
	)
}

func Computer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="8" x="5" y="2" rx="2"/><rect width="20" height="8" x="2" y="14" rx="2"/><path d="M6 18h2m4 0h6"/></g>`), g.Group(children),
	)
}

func ConciergeBell(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 18a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v2H2v-2Zm18-2a8 8 0 1 0-16 0m8-12v4m-2-4h4"/>`), g.Group(children),
	)
}

func Cone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m20.9 18.55l-8-15.98a1 1 0 0 0-1.8 0l-8 15.98"/><ellipse cx="12" cy="19" rx="9" ry="3"/></g>`), g.Group(children),
	)
}

func Construction(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="8" x="2" y="6" rx="1"/><path d="M17 14v7M7 14v7M17 3v3M7 3v3m3 8L2.3 6.3M14 6l7.7 7.7M8 6l8 8"/></g>`), g.Group(children),
	)
}

func Contact(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 18a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2"/><rect width="18" height="18" x="3" y="4" rx="2"/><circle cx="12" cy="10" r="2"/><path d="M8 2v2m8-2v2"/></g>`), g.Group(children),
	)
}

func ContactTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 18a4 4 0 0 0-8 0"/><circle cx="12" cy="11" r="3"/><rect width="18" height="18" x="3" y="4" rx="2"/><path d="M8 2v2m8-2v2"/></g>`), g.Group(children),
	)
}

func Container(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 7.7c0-.6-.4-1.2-.8-1.5l-6.3-3.9a1.72 1.72 0 0 0-1.7 0l-10.3 6c-.5.2-.9.8-.9 1.4v6.6c0 .5.4 1.2.8 1.5l6.3 3.9a1.72 1.72 0 0 0 1.7 0l10.3-6c.5-.3.9-1 .9-1.5Z"/><path d="M10 21.9V14L2.1 9.1M10 14l11.9-6.9M14 19.8v-8.1m4 5.8V9.4"/></g>`), g.Group(children),
	)
}

func Contrast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 18a6 6 0 0 0 0-12v12z"/></g>`), g.Group(children),
	)
}

func Cookie(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2a10 10 0 1 0 10 10a4 4 0 0 1-5-5a4 4 0 0 1-5-5M8.5 8.5v.01M16 15.5v.01M12 12v.01M11 17v.01M7 14v.01"/>`), g.Group(children),
	)
}

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`), g.Group(children),
	)
}

func CopyCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 15l2 2l4-4"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`), g.Group(children),
	)
}

func CopyMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 15h6"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`), g.Group(children),
	)
}

func CopyPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 12v6m-3-3h6"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`), g.Group(children),
	)
}

func CopySlash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 18l6-6"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`), g.Group(children),
	)
}

func CopyX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 12l6 6m-6 0l6-6"/><rect width="14" height="14" x="8" y="8" rx="2" ry="2"/><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"/></g>`), g.Group(children),
	)
}

func Copyleft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9.17 14.83a4 4 0 1 0 0-5.66"/></g>`), g.Group(children),
	)
}

func Copyright(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M14.83 14.83a4 4 0 1 1 0-5.66"/></g>`), g.Group(children),
	)
}

func CornerDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 10l-5 5l5 5"/><path d="M20 4v7a4 4 0 0 1-4 4H4"/></g>`), g.Group(children),
	)
}

func CornerDownRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 10l5 5l-5 5"/><path d="M4 4v7a4 4 0 0 0 4 4h12"/></g>`), g.Group(children),
	)
}

func CornerLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14 15l-5 5l-5-5"/><path d="M20 4h-7a4 4 0 0 0-4 4v12"/></g>`), g.Group(children),
	)
}

func CornerLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 9L9 4L4 9"/><path d="M20 20h-7a4 4 0 0 1-4-4V4"/></g>`), g.Group(children),
	)
}

func CornerRightDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 15l5 5l5-5"/><path d="M4 4h7a4 4 0 0 1 4 4v12"/></g>`), g.Group(children),
	)
}

func CornerRightUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 9l5-5l5 5"/><path d="M4 20h7a4 4 0 0 0 4-4V4"/></g>`), g.Group(children),
	)
}

func CornerUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 14L4 9l5-5"/><path d="M20 20v-7a4 4 0 0 0-4-4H4"/></g>`), g.Group(children),
	)
}

func CornerUpRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 14l5-5l-5-5"/><path d="M4 20v-7a4 4 0 0 1 4-4h12"/></g>`), g.Group(children),
	)
}

func Cpu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="16" x="4" y="4" rx="2"/><path d="M9 9h6v6H9zm6-7v2m0 16v2M2 15h2M2 9h2m16 6h2m-2-6h2M9 2v2m0 16v2"/></g>`), g.Group(children),
	)
}

func CreativeCommons(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M10 9.3a2.8 2.8 0 0 0-3.5 1a3.1 3.1 0 0 0 0 3.4a2.7 2.7 0 0 0 3.5 1m7-5.4a2.8 2.8 0 0 0-3.5 1a3.1 3.1 0 0 0 0 3.4a2.7 2.7 0 0 0 3.5 1"/></g>`), g.Group(children),
	)
}

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="14" x="2" y="5" rx="2"/><path d="M2 10h20"/></g>`), g.Group(children),
	)
}

func Croissant(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4.6 13.11l5.79-3.21c1.89-1.05 4.79 1.78 3.71 3.71l-3.22 5.81C8.8 23.16.79 15.23 4.6 13.11Zm5.9-3.61l-1-2.29C9.2 6.48 8.8 6 8 6H4.5C2.79 6 2 6.5 2 8.5a7.71 7.71 0 0 0 2 4.83M8 6c0-1.55.24-4-2-4c-2 0-2.5 2.17-2.5 4m11 7.5l2.29 1c.73.3 1.21.7 1.21 1.5v3.5c0 1.71-.5 2.5-2.5 2.5a7.71 7.71 0 0 1-4.83-2M18 16c1.55 0 4-.24 4 2c0 2-2.17 2.5-4 2.5"/>`), g.Group(children),
	)
}

func Crop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 2v14a2 2 0 0 0 2 2h14"/><path d="M18 22V8a2 2 0 0 0-2-2H2"/></g>`), g.Group(children),
	)
}

func Cross(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 2a2 2 0 0 0-2 2v5H4a2 2 0 0 0-2 2v2c0 1.1.9 2 2 2h5v5c0 1.1.9 2 2 2h2a2 2 0 0 0 2-2v-5h5a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-5V4a2 2 0 0 0-2-2h-2z"/>`), g.Group(children),
	)
}

func Crosshair(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M22 12h-4M6 12H2m10-6V2m0 20v-4"/></g>`), g.Group(children),
	)
}

func Crown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 4l3 12h14l3-12l-6 7l-4-7l-4 7l-6-7zm3 16h14"/>`), g.Group(children),
	)
}

func Cuboid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m21.12 6.4l-6.05-4.06a2 2 0 0 0-2.17-.05L2.95 8.41a2 2 0 0 0-.95 1.7v5.82a2 2 0 0 0 .88 1.66l6.05 4.07a2 2 0 0 0 2.17.05l9.95-6.12a2 2 0 0 0 .95-1.7V8.06a2 2 0 0 0-.88-1.66Z"/><path d="M10 22v-8L2.25 9.15M10 14l11.77-6.87"/></g>`), g.Group(children),
	)
}

func CupSoda(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m6 8l1.75 12.28a2 2 0 0 0 2 1.72h4.54a2 2 0 0 0 2-1.72L18 8M5 8h14"/><path d="M7 15a6.47 6.47 0 0 1 5 0a6.47 6.47 0 0 0 5 0m-5-7l1-6h2"/></g>`), g.Group(children),
	)
}

func Currency(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="8"/><path d="m3 3l3 3m15-3l-3 3M3 21l3-3m15 3l-3-3"/></g>`), g.Group(children),
	)
}

func Cylinder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 5v14a9 3 0 0 0 18 0V5"/></g>`), g.Group(children),
	)
}

func Database(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 5v14a9 3 0 0 0 18 0V5"/><path d="M3 12a9 3 0 0 0 18 0"/></g>`), g.Group(children),
	)
}

func DatabaseBackup(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 12a9 3 0 0 0 5 2.69M21 9.3V5"/><path d="M3 5v14a9 3 0 0 0 6.47 2.88M12 12v4h4"/><path d="M13 20a5 5 0 0 0 9-3a4.5 4.5 0 0 0-4.5-4.5c-1.33 0-2.54.54-3.41 1.41L12 16"/></g>`), g.Group(children),
	)
}

func DatabaseZap(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="5" rx="9" ry="3"/><path d="M3 5v14a9 3 0 0 0 12 2.84M21 5v3m0 4l-3 5h4l-3 5"/><path d="M3 12a9 3 0 0 0 11.59 2.87"/></g>`), g.Group(children),
	)
}

func Delete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 5H9l-7 7l7 7h11a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2Zm-2 4l-6 6m0-6l6 6"/>`), g.Group(children),
	)
}

func Dessert(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="4" r="2"/><path d="M10.2 3.2C5.5 4 2 8.1 2 13a2 2 0 0 0 4 0v-1a2 2 0 0 1 4 0v4a2 2 0 0 0 4 0v-4a2 2 0 0 1 4 0v1a2 2 0 0 0 4 0c0-4.9-3.5-9-8.2-9.8"/><path d="M3.2 14.8a9 9 0 0 0 17.6 0"/></g>`), g.Group(children),
	)
}

func Diameter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="19" cy="19" r="2"/><circle cx="5" cy="5" r="2"/><path d="M6.48 3.66a10 10 0 0 1 13.86 13.86M6.41 6.41l11.18 11.18M3.66 6.48a10 10 0 0 0 13.86 13.86"/></g>`), g.Group(children),
	)
}

func Diamond(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41L13.7 2.71a2.41 2.41 0 0 0-3.41 0Z"/>`), g.Group(children),
	)
}

func DiceFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01M8 8h.01M8 16h.01M16 16h.01M12 12h.01"/></g>`), g.Group(children),
	)
}

func DiceFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01M8 8h.01M8 16h.01M16 16h.01"/></g>`), g.Group(children),
	)
}

func DiceOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M12 12h.01"/></g>`), g.Group(children),
	)
}

func DiceSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01M16 12h.01M16 16h.01M8 8h.01M8 12h.01M8 16h.01"/></g>`), g.Group(children),
	)
}

func DiceThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M16 8h.01M12 12h.01M8 16h.01"/></g>`), g.Group(children),
	)
}

func DiceTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M15 9h.01M9 15h.01"/></g>`), g.Group(children),
	)
}

func Dices(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="12" height="12" x="2" y="10" rx="2" ry="2"/><path d="m17.92 14l3.5-3.5a2.24 2.24 0 0 0 0-3l-5-4.92a2.24 2.24 0 0 0-3 0L10 6M6 18h.01M10 14h.01M15 6h.01M18 9h.01"/></g>`), g.Group(children),
	)
}

func Diff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v14m-7-7h14M5 21h14"/>`), g.Group(children),
	)
}

func Disc(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="2"/></g>`), g.Group(children),
	)
}

func DiscThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M6 12c0-1.7.7-3.2 1.8-4.2"/><circle cx="12" cy="12" r="2"/><path d="M18 12c0 1.7-.7 3.2-1.8 4.2"/></g>`), g.Group(children),
	)
}

func DiscTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="4"/><path d="M12 12h.01"/></g>`), g.Group(children),
	)
}

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="6" r="1"/><path d="M5 12h14"/><circle cx="12" cy="18" r="1"/></g>`), g.Group(children),
	)
}

func DivideCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 12h8m-4 4h0m0-8h0"/><circle cx="12" cy="12" r="10"/></g>`), g.Group(children),
	)
}

func DivideSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M8 12h8m-4 4h0m0-8h0"/></g>`), g.Group(children),
	)
}

func Dna(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 15c6.667-6 13.333 0 20-6M9 22c1.798-1.998 2.518-3.995 2.807-5.993M15 2c-1.798 1.998-2.518 3.995-2.807 5.993M17 6l-2.5-2.5M14 8l-1-1M7 18l2.5 2.5m-6-6l.5.5m16-6l.5.5m-14 3l1 1m9-3l1 1M10 16l1.5 1.5"/>`), g.Group(children),
	)
}

func DnaOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 2c-1.35 1.5-2.092 3-2.5 4.5M9 22c1.35-1.5 2.092-3 2.5-4.5M2 15c3.333-3 6.667-3 10-3m10-3c-1.5 1.35-3 2.092-4.5 2.5M17 6l-2.5-2.5M14 8l-1.5-1.5M7 18l2.5 2.5m-6-6l.5.5m16-6l.5.5m-14 3l1 1m9-3l1 1M10 16l1.5 1.5M2 2l20 20"/>`), g.Group(children),
	)
}

func Dog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 5.172C10 3.782 8.423 2.679 6.5 3c-2.823.47-4.113 6.006-4 7c.08.703 1.725 1.722 3.656 1c1.261-.472 1.96-1.45 2.344-2.5m5.767-3.328c0-1.39 1.577-2.493 3.5-2.172c2.823.47 4.113 6.006 4 7c-.08.703-1.725 1.722-3.656 1c-1.261-.472-1.855-1.45-2.239-2.5M8 14v.5m8-.5v.5m-4.75 1.75h1.5L12 17l-.75-.75Z"/><path d="M4.42 11.247A13.152 13.152 0 0 0 4 14.556C4 18.728 7.582 21 12 21s8-2.272 8-6.444c0-1.061-.162-2.2-.493-3.309m-9.243-6.082A8.801 8.801 0 0 1 12 5c.78 0 1.5.108 2.161.306"/></g>`), g.Group(children),
	)
}

func DollarSign(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v20m5-17H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/>`), g.Group(children),
	)
}

func Donut(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.5 10a2.5 2.5 0 0 1-2.4-3H18a2.95 2.95 0 0 1-2.6-4.4a10 10 0 1 0 6.3 7.1c-.3.2-.8.3-1.2.3"/><circle cx="12" cy="12" r="3"/></g>`), g.Group(children),
	)
}

func DoorClosed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 20V6a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v14m-4 0h20m-8-8v.01"/>`), g.Group(children),
	)
}

func DoorOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 4h3a2 2 0 0 1 2 2v14M2 20h3m8 0h9m-12-8v.01m3-7.448v16.157a1 1 0 0 1-1.242.97L5 20V5.562a2 2 0 0 1 1.515-1.94l4-1A2 2 0 0 1 13 4.561Z"/>`), g.Group(children),
	)
}

func Dot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<circle cx="12.1" cy="12.1" r="1" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/>`), g.Group(children),
	)
}

func Download(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4m4-5l5 5l5-5m-5 5V3"/>`), g.Group(children),
	)
}

func DownloadCloud(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M12 12v9m-4-4l4 4l4-4"/>`), g.Group(children),
	)
}

func DraftingCompass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="5" r="2"/><path d="m3 21l8.02-14.26m1.97 0l1.93 3.44M19 12c-3.87 4-10.13 4-14 0m16 9l-2.16-3.84"/></g>`), g.Group(children),
	)
}

func Drama(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 11h.01M14 6h.01M18 6h.01M6.5 13.1h.01M22 5c0 9-4 12-6 12s-6-3-6-12c0-2 2-3 6-3s6 1 6 3"/><path d="M17.4 9.9c-.8.8-2 .8-2.8 0m-4.5-2.8C9 7.2 7.7 7.7 6 8.6c-3.5 2-4.7 3.9-3.7 5.6c4.5 7.8 9.5 8.4 11.2 7.4c.9-.5 1.9-2.1 1.9-4.7"/><path d="M9.1 16.5c.3-1.1 1.4-1.7 2.4-1.4"/></g>`), g.Group(children),
	)
}

func Dribbble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M19.13 5.09C15.22 9.14 10 10.44 2.25 10.94m19.5 1.9c-6.62-1.41-12.14 1-16.38 6.32"/><path d="M8.56 2.75c4.37 6 6 9.42 8 17.72"/></g>`), g.Group(children),
	)
}

func Droplet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22a7 7 0 0 0 7-7c0-2-1-3.9-3-5.5s-3.5-4-4-6.5c-.5 2.5-2 4.9-4 6.5C6 11.1 5 13 5 15a7 7 0 0 0 7 7z"/>`), g.Group(children),
	)
}

func Droplets(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 16.3c2.2 0 4-1.83 4-4.05c0-1.16-.57-2.26-1.71-3.19S7.29 6.75 7 5.3c-.29 1.45-1.14 2.84-2.29 3.76S3 11.1 3 12.25c0 2.22 1.8 4.05 4 4.05z"/><path d="M12.56 6.6A10.97 10.97 0 0 0 14 3.02c.5 2.5 2 4.9 4 6.5s3 3.5 3 5.5a6.98 6.98 0 0 1-11.91 4.97"/></g>`), g.Group(children),
	)
}

func Drumstick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.45 15.4c-2.13.65-4.3.32-5.7-1.1c-2.29-2.27-1.76-6.5 1.17-9.42c2.93-2.93 7.15-3.46 9.43-1.18c1.41 1.41 1.74 3.57 1.1 5.71c-1.4-.51-3.26-.02-4.64 1.36c-1.38 1.38-1.87 3.23-1.36 4.63z"/><path d="m11.25 15.6l-2.16 2.16a2.5 2.5 0 1 1-4.56 1.73a2.49 2.49 0 0 1-1.41-4.24a2.5 2.5 0 0 1 3.14-.32l2.16-2.16"/></g>`), g.Group(children),
	)
}

func Dumbbell(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6.5 6.5l11 11M21 21l-1-1M3 3l1 1m14 18l4-4M2 6l4-4m-3 8l7-7m4 18l7-7"/>`), g.Group(children),
	)
}

func Ear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 8.5a6.5 6.5 0 1 1 13 0c0 6-6 6-6 10a3.5 3.5 0 1 1-7 0"/><path d="M15 8.5a2.5 2.5 0 0 0-5 0v1a2 2 0 1 1 0 4"/></g>`), g.Group(children),
	)
}

func EarOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 18.5a3.5 3.5 0 1 0 7 0c0-1.57.92-2.52 2.04-3.46M6 8.5c0-.75.13-1.47.36-2.14M8.8 3.15A6.5 6.5 0 0 1 19 8.5c0 1.63-.44 2.81-1.09 3.76"/><path d="M12.5 6A2.5 2.5 0 0 1 15 8.5M10 13a2 2 0 0 0 1.82-1.18M2 2l20 20"/></g>`), g.Group(children),
	)
}

func Edit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M18.5 2.5a2.121 2.121 0 0 1 3 3L12 15l-4 1l1-4l9.5-9.5z"/></g>`), g.Group(children),
	)
}

func EditThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20h9M16.5 3.5a2.121 2.121 0 0 1 3 3L7 19l-4 1l1-4L16.5 3.5z"/>`), g.Group(children),
	)
}

func EditTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3a2.828 2.828 0 1 1 4 4L7.5 20.5L2 22l1.5-5.5L17 3z"/>`), g.Group(children),
	)
}

func Egg(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22c6.23-.05 7.87-5.57 7.5-10c-.36-4.34-3.95-9.96-7.5-10c-3.55.04-7.14 5.66-7.5 10c-.37 4.43 1.27 9.95 7.5 10z"/>`), g.Group(children),
	)
}

func EggFried(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11.5" cy="12.5" r="3.5"/><path d="M3 8c0-3.5 2.5-6 6.5-6c5 0 4.83 3 7.5 5s5 2 5 6c0 4.5-2.5 6.5-7 6.5c-2.5 0-2.5 2.5-6 2.5s-7-2-7-5.5c0-3 1.5-3 1.5-5C3.5 10 3 9 3 8Z"/></g>`), g.Group(children),
	)
}

func EggOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.399 6.399C5.362 8.157 4.65 10.189 4.5 12c-.37 4.43 1.27 9.95 7.5 10c3.256-.026 5.259-1.547 6.375-3.625m1.157-4.5A14.07 14.07 0 0 0 19.5 12c-.36-4.34-3.95-9.96-7.5-10c-1.04.012-2.082.502-3.046 1.297M2 2l20 20"/>`), g.Group(children),
	)
}

func Equal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9h14M5 15h14"/>`), g.Group(children),
	)
}

func EqualNot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9h14M5 15h14m0-10L5 19"/>`), g.Group(children),
	)
}

func Eraser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 21l-4.3-4.3c-1-1-1-2.5 0-3.4l9.6-9.6c1-1 2.5-1 3.4 0l5.6 5.6c1 1 1 2.5 0 3.4L13 21m9 0H7M5 11l9 9"/>`), g.Group(children),
	)
}

func Euro(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 10h12M4 14h9m6-8a7.7 7.7 0 0 0-5.2-2A7.9 7.9 0 0 0 6 12c0 4.4 3.5 8 7.8 8c2 0 3.8-.8 5.2-2"/>`), g.Group(children),
	)
}

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 21l-6-6m6 6v-4.8m0 4.8h-4.8M3 16.2V21m0 0h4.8M3 21l6-6m12-7.2V3m0 0h-4.8M21 3l-6 6M3 7.8V3m0 0h4.8M3 3l6 6"/>`), g.Group(children),
	)
}

func ExternalLink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 13v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h6m4-3h6v6m-11 5L21 3"/>`), g.Group(children),
	)
}

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12s3-7 10-7s10 7 10 7s-3 7-10 7s-10-7-10-7Z"/><circle cx="12" cy="12" r="3"/></g>`), g.Group(children),
	)
}

func EyeOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.88 9.88a3 3 0 1 0 4.24 4.24m-3.39-9.04A10.43 10.43 0 0 1 12 5c7 0 10 7 10 7a13.16 13.16 0 0 1-1.67 2.68"/><path d="M6.61 6.61A13.526 13.526 0 0 0 2 12s3 7 10 7a9.74 9.74 0 0 0 5.39-1.61M2 2l20 20"/></g>`), g.Group(children),
	)
}

func Facebook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 2h-3a5 5 0 0 0-5 5v3H7v4h3v8h4v-8h3l1-4h-4V7a1 1 0 0 1 1-1h3z"/>`), g.Group(children),
	)
}

func Factory(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8l-7 5V8l-7 5V4a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Zm15-2h1m-6 0h1m-6 0h1"/>`), g.Group(children),
	)
}

func Fan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.827 16.379a6.082 6.082 0 0 1-8.618-7.002l5.412 1.45a6.082 6.082 0 0 1 7.002-8.618l-1.45 5.412a6.082 6.082 0 0 1 8.618 7.002l-5.412-1.45a6.082 6.082 0 0 1-7.002 8.618l1.45-5.412ZM12 12v.01"/>`), g.Group(children),
	)
}

func FastForward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 19l9-7l-9-7v14zM2 19l9-7l-9-7v14z"/>`), g.Group(children),
	)
}

func Feather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.24 12.24a6 6 0 0 0-8.49-8.49L5 10.5V19h8.5zM16 8L2 22m15.5-7H9"/>`), g.Group(children),
	)
}

func FerrisWheel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="2"/><path d="M12 2v4m-5.2 9l-3.5 2M20.7 7l-3.5 2M6.8 9L3.3 7m17.4 10l-3.5-2M9 22l3-8l3 8m-7 0h8"/><path d="M18 18.7a9 9 0 1 0-12 0"/></g>`), g.Group(children),
	)
}

func Figma(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 5.5A3.5 3.5 0 0 1 8.5 2H12v7H8.5A3.5 3.5 0 0 1 5 5.5zM12 2h3.5a3.5 3.5 0 1 1 0 7H12V2z"/><path d="M12 12.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 1 1-7 0zm-7 7A3.5 3.5 0 0 1 8.5 16H12v3.5a3.5 3.5 0 1 1-7 0zm0-7A3.5 3.5 0 0 1 8.5 9H12v7H8.5A3.5 3.5 0 0 1 5 12.5z"/></g>`), g.Group(children),
	)
}

func File(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6"/></g>`), g.Group(children),
	)
}

func FileArchive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22V4c0-.5.2-1 .6-1.4C5 2.2 5.5 2 6 2h8.5L20 7.5V20c0 .5-.2 1-.6 1.4c-.4.4-.9.6-1.4.6h-2"/><path d="M14 2v6h6"/><circle cx="10" cy="20" r="2"/><path d="M10 7V6m0 6v-1m0 7v-2"/></g>`), g.Group(children),
	)
}

func FileAudio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 22h.5c.5 0 1-.2 1.4-.6c.4-.4.6-.9.6-1.4V7.5L14.5 2H6c-.5 0-1 .2-1.4.6C4.2 3 4 3.5 4 4v3"/><path d="M14 2v6h6M10 20v-1a2 2 0 1 1 4 0v1a2 2 0 1 1-4 0Zm-4 0v-1a2 2 0 1 0-4 0v1a2 2 0 1 0 4 0Z"/><path d="M2 19v-3a6 6 0 0 1 12 0v3"/></g>`), g.Group(children),
	)
}

func FileAudioTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v2"/><path d="M14 2v6h6M2 17v-3a4 4 0 0 1 8 0v3"/><circle cx="9" cy="17" r="1"/><circle cx="3" cy="17" r="1"/></g>`), g.Group(children),
	)
}

func FileAxisThreeD(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M8 10v8h8m-8 0l4-4"/></g>`), g.Group(children),
	)
}

func FileBadge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 7V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2h-6"/><path d="M14 2v6h6M5 17a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="M7 16.5L8 22l-3-1l-3 1l1-5.5"/></g>`), g.Group(children),
	)
}

func FileBadgeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M12 13a3 3 0 1 0 0-6a3 3 0 0 0 0 6Z"/><path d="m14 12.5l1 5.5l-3-1l-3 1l1-5.5"/></g>`), g.Group(children),
	)
}

func FileBarChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8 10v-4m-4 4v-2m8 2v-6"/></g>`), g.Group(children),
	)
}

func FileBarChartTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8 10v-6m-4 6v-1m8 1v-3"/></g>`), g.Group(children),
	)
}

func FileBox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 22H18a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M2.97 13.12c-.6.36-.97 1.02-.97 1.74v3.28c0 .72.37 1.38.97 1.74l3 1.83c.63.39 1.43.39 2.06 0l3-1.83c.6-.36.97-1.02.97-1.74v-3.28c0-.72-.37-1.38-.97-1.74l-3-1.83a1.97 1.97 0 0 0-2.06 0l-3 1.83ZM7 17l-4.74-2.85M7 17l4.74-2.85M7 17v5"/></g>`), g.Group(children),
	)
}

func FileCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M9 15l2 2l4-4"/></g>`), g.Group(children),
	)
}

func FileCheckTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M3 15l2 2l4-4"/></g>`), g.Group(children),
	)
}

func FileClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 22h2c.5 0 1-.2 1.4-.6c.4-.4.6-.9.6-1.4V7.5L14.5 2H6c-.5 0-1 .2-1.4.6C4.2 3 4 3.5 4 4v3"/><path d="M14 2v6h6"/><circle cx="8" cy="16" r="6"/><path d="M9.5 17.5L8 16.25V14"/></g>`), g.Group(children),
	)
}

func FileCode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-10 5l-2 2l2 2m4 0l2-2l-2-2"/></g>`), g.Group(children),
	)
}

func FileCodeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M9 18l3-3l-3-3m-4 0l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func FileCog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="13" r="3"/><path d="m9.7 14.4l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.7.9l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4M7.4 9.3l-.3.9M14 2v6h6"/><path d="M4 5.5V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-1.5"/></g>`), g.Group(children),
	)
}

func FileCogTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6"/><circle cx="12" cy="15" r="2"/><path d="M12 12v1m0 4v1m2.6-4.5l-.87.5m-3.46 2l-.87.5m5.2 0l-.87-.5m-3.46-2l-.87-.5"/></g>`), g.Group(children),
	)
}

func FileDiff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2zM12 13V7m-3 3h6m-6 7h6"/>`), g.Group(children),
	)
}

func FileDigit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="4" height="6" x="2" y="12" rx="2"/><path d="M14 2v6h6"/><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M10 12h2v6m-2 0h4"/></g>`), g.Group(children),
	)
}

func FileDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8 10v-6m-3 3l3 3l3-3"/></g>`), g.Group(children),
	)
}

func FileEdit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 13.5V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2h-5.5"/><path d="M14 2v6h6m-9.58 4.61a2.1 2.1 0 1 1 2.97 2.97L7.95 21L4 22l.99-3.95l5.43-5.44Z"/></g>`), g.Group(children),
	)
}

func FileHeart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 6V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H4"/><path d="M14 2v6h6m-9.71 2.7a2.43 2.43 0 0 0-2.66-.52c-.29.12-.56.3-.78.53l-.35.34l-.35-.34a2.43 2.43 0 0 0-2.65-.53c-.3.12-.56.3-.79.53c-.95.94-1 2.53.2 3.74L6.5 18l3.6-3.55c1.2-1.21 1.14-2.8.19-3.74Z"/></g>`), g.Group(children),
	)
}

func FileImage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6"/><circle cx="10" cy="13" r="2"/><path d="m20 17l-1.09-1.09a2 2 0 0 0-2.82 0L10 22"/></g>`), g.Group(children),
	)
}

func FileInput(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M2 15h10m-3 3l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func FileJson(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-10 4a1 1 0 0 0-1 1v1a1 1 0 0 1-1 1a1 1 0 0 1 1 1v1a1 1 0 0 0 1 1m4 0a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1a1 1 0 0 1-1-1v-1a1 1 0 0 0-1-1"/></g>`), g.Group(children),
	)
}

func FileJsonTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M4 12a1 1 0 0 0-1 1v1a1 1 0 0 1-1 1a1 1 0 0 1 1 1v1a1 1 0 0 0 1 1m4 0a1 1 0 0 0 1-1v-1a1 1 0 0 1 1-1a1 1 0 0 1-1-1v-1a1 1 0 0 0-1-1"/></g>`), g.Group(children),
	)
}

func FileKey(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><circle cx="10" cy="16" r="2"/><path d="m16 10l-4.5 4.5M15 11l1 1"/></g>`), g.Group(children),
	)
}

func FileKeyTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 10V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H4"/><path d="M14 2v6h6"/><circle cx="4" cy="16" r="2"/><path d="m10 10l-4.5 4.5M9 11l1 1"/></g>`), g.Group(children),
	)
}

func FileLineChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-4 5l-3.5 3.5l-2-2L8 17"/></g>`), g.Group(children),
	)
}

func FileLock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><rect width="8" height="6" x="8" y="12" rx="1"/><path d="M15 12v-2a3 3 0 1 0-6 0v2"/></g>`), g.Group(children),
	)
}

func FileLockTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 5V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H4"/><path d="M14 2v6h6"/><rect width="8" height="5" x="2" y="13" rx="1"/><path d="M8 13v-2a2 2 0 1 0-4 0v2"/></g>`), g.Group(children),
	)
}

func FileMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M9 15h6"/></g>`), g.Group(children),
	)
}

func FileMinusTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M3 15h6"/></g>`), g.Group(children),
	)
}

func FileOutput(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M2 15h10m-7-3l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func FilePieChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 22h2a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v3"/><path d="M14 2v6h6M4.04 11.71a5.84 5.84 0 1 0 8.2 8.29"/><path d="M13.83 16A5.83 5.83 0 0 0 8 10.17V16h5.83Z"/></g>`), g.Group(children),
	)
}

func FilePlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8 10v-6m-3 3h6"/></g>`), g.Group(children),
	)
}

func FilePlusTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M3 15h6m-3-3v6"/></g>`), g.Group(children),
	)
}

func FileQuestion(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M10 10.3c.2-.4.5-.8.9-1a2.1 2.1 0 0 1 2.6.4c.3.4.5.8.5 1.3c0 1.3-2 2-2 2m0 4h.01"/></g>`), g.Group(children),
	)
}

func FileScan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 10V7.5L14.5 2H6a2 2 0 0 0-2 2v16c0 1.1.9 2 2 2h4.5"/><path d="M14 2v6h6m-4 14a2 2 0 0 1-2-2m6 2a2 2 0 0 0 2-2m-2-6a2 2 0 0 1 2 2m-6-2a2 2 0 0 0-2 2"/></g>`), g.Group(children),
	)
}

func FileSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v3"/><path d="M14 2v6h6M5 17a3 3 0 1 0 0-6a3 3 0 0 0 0 6zm4 1l-1.5-1.5"/></g>`), g.Group(children),
	)
}

func FileSearchTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6"/><circle cx="11.5" cy="14.5" r="2.5"/><path d="M13.25 16.25L15 18"/></g>`), g.Group(children),
	)
}

func FileSignature(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 19.5v.5a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h8.5L18 5.5M8 18h1"/><path d="M18.42 9.61a2.1 2.1 0 1 1 2.97 2.97L16.95 17L13 18l.99-3.95l4.43-4.44Z"/></g>`), g.Group(children),
	)
}

func FileSpreadsheet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M8 13h2m-2 4h2m4-4h2m-2 4h2"/></g>`), g.Group(children),
	)
}

func FileStack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 2v5h5"/><path d="M21 6v6.5c0 .8-.7 1.5-1.5 1.5h-7c-.8 0-1.5-.7-1.5-1.5v-9c0-.8.7-1.5 1.5-1.5H17l4 4z"/><path d="M7 8v8.8c0 .3.2.6.4.8c.2.2.5.4.8.4H15"/><path d="M3 12v8.8c0 .3.2.6.4.8c.2.2.5.4.8.4H11"/></g>`), g.Group(children),
	)
}

func FileSymlink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v7"/><path d="M14 2v6h6M10 18l3-3l-3-3"/><path d="M4 18v-1a2 2 0 0 1 2-2h6"/></g>`), g.Group(children),
	)
}

func FileTerminal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M8 16l2-2l-2-2m4 6h4"/></g>`), g.Group(children),
	)
}

func FileText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-4 5H8m8 4H8m2-8H8"/></g>`), g.Group(children),
	)
}

func FileType(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M9 13v-1h6v1m-4 5h2m-1-6v6"/></g>`), g.Group(children),
	)
}

func FileTypeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M2 13v-1h6v1m-4 5h2m-1-6v6"/></g>`), g.Group(children),
	)
}

func FileUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8 4v6m3-3l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func FileVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-10 3l5 3l-5 3v-6Z"/></g>`), g.Group(children),
	)
}

func FileVideoTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 8V4a2 2 0 0 1 2-2h8.5L20 7.5V20a2 2 0 0 1-2 2H4"/><path d="M14 2v6h6m-10 7.5l4 2.5v-6l-4 2.5"/><rect width="8" height="6" x="2" y="12" rx="1"/></g>`), g.Group(children),
	)
}

func FileVolume(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v3"/><path d="M14 2v6h6M7 10l-3 2H2v4h2l3 2v-8Zm4 1c.64.8 1 1.87 1 3s-.36 2.2-1 3"/></g>`), g.Group(children),
	)
}

func FileVolumeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6m-8.5 5.5c.32.4.5.94.5 1.5s-.18 1.1-.5 1.5M15 12c.64.8 1 1.87 1 3s-.36 2.2-1 3m-7-3h.01"/></g>`), g.Group(children),
	)
}

func FileWarning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2zM12 9v4m0 4h.01"/>`), g.Group(children),
	)
}

func FileX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14.5 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V7.5L14.5 2z"/><path d="M14 2v6h6M9.5 12.5l5 5m0-5l-5 5"/></g>`), g.Group(children),
	)
}

func FileXTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h14a2 2 0 0 0 2-2V7.5L14.5 2H6a2 2 0 0 0-2 2v4"/><path d="M14 2v6h6M3 12.5l5 5m0-5l-5 5"/></g>`), g.Group(children),
	)
}

func Files(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.5 2H8.6c-.4 0-.8.2-1.1.5c-.3.3-.5.7-.5 1.1v12.8c0 .4.2.8.5 1.1c.3.3.7.5 1.1.5h9.8c.4 0 .8-.2 1.1-.5c.3-.3.5-.7.5-1.1V6.5L15.5 2z"/><path d="M3 7.6v12.8c0 .4.2.8.5 1.1c.3.3.7.5 1.1.5h9.8M15 2v5h5"/></g>`), g.Group(children),
	)
}

func Film(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 3v18M3 7.5h4M3 12h18M3 16.5h4M17 3v18m0-13.5h4m-4 9h4"/></g>`), g.Group(children),
	)
}

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 3H2l8 9.46V19l4 2v-8.54L22 3z"/>`), g.Group(children),
	)
}

func FilterX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.013 3H2l8 9.46V19l4 2v-8.54l.9-1.055M22 3l-5 5m0-5l5 5"/>`), g.Group(children),
	)
}

func Fingerprint(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12C2 6.5 6.5 2 12 2a10 10 0 0 1 8 4"/><path d="M5 19.5C5.5 18 6 15 6 12c0-.7.12-1.37.34-2m10.95 11.02c.12-.6.43-2.3.5-3.02M12 10a2 2 0 0 0-2 2c0 1.02-.1 2.51-.26 4m-1.09 6c.21-.66.45-1.32.57-2M14 13.12c0 2.38 0 6.38-1 8.88M2 16h.01m19.79 0c.2-2 .131-5.354 0-6M9 6.8a6 6 0 0 1 9 5.2c0 .47 0 1.17-.02 2"/></g>`), g.Group(children),
	)
}

func Fish(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6.5 12c.94-3.46 4.94-6 8.5-6c3.56 0 6.06 2.54 7 6c-.94 3.47-3.44 6-7 6s-7.56-2.53-8.5-6ZM18 12v.5"/><path d="M16 17.93a9.77 9.77 0 0 1 0-11.86m-9 4.6C7 8 5.58 5.97 2.73 5.5c-1 1.5-1 5 .23 6.5c-1.24 1.5-1.24 5-.23 6.5C5.58 18.03 7 16 7 13.33"/><path d="M10.46 7.26C10.2 5.88 9.17 4.24 8 3h5.8a2 2 0 0 1 1.98 1.67l.23 1.4m0 11.86l-.23 1.4A2 2 0 0 1 13.8 21H9.5a5.96 5.96 0 0 0 1.49-3.98"/></g>`), g.Group(children),
	)
}

func FishOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 12.47v.03m0-.5v.47m-.475 5.056A6.744 6.744 0 0 1 15 18c-3.56 0-7.56-2.53-8.5-6c.348-1.28 1.114-2.433 2.121-3.38m3.444-2.088A8.802 8.802 0 0 1 15 6c3.56 0 6.06 2.54 7 6c-.309 1.14-.786 2.177-1.413 3.058"/><path d="M7 10.67C7 8 5.58 5.97 2.73 5.5c-1 1.5-1 5 .23 6.5c-1.24 1.5-1.24 5-.23 6.5C5.58 18.03 7 16 7 13.33m7.48-4.372A9.77 9.77 0 0 1 16 6.07m0 11.86a9.77 9.77 0 0 1-1.728-3.618"/><path d="m16.01 17.93l-.23 1.4A2 2 0 0 1 13.8 21H9.5a5.96 5.96 0 0 0 1.49-3.98M8.53 3h5.27a2 2 0 0 1 1.98 1.67l.23 1.4M2 2l20 20"/></g>`), g.Group(children),
	)
}

func FishSymbol(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 16s9-15 20-4C11 23 2 8 2 8"/>`), g.Group(children),
	)
}

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 15s1-1 4-1s5 2 8 2s4-1 4-1V3s-1 1-4 1s-5-2-8-2s-4 1-4 1zm0 7v-7"/>`), g.Group(children),
	)
}

func FlagOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 2c3 0 5 2 8 2s4-1 4-1v11M4 22V4m0 11s1-1 4-1s5 2 8 2M2 2l20 20"/>`), g.Group(children),
	)
}

func FlagTriangleLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 22V2L7 7l10 5"/>`), g.Group(children),
	)
}

func FlagTriangleRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 22V2l10 5l-10 5"/>`), g.Group(children),
	)
}

func Flame(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.5 14.5A2.5 2.5 0 0 0 11 12c0-1.38-.5-2-1-3c-1.072-2.143-.224-4.054 2-6c.5 2.5 2 4.9 4 6.5c2 1.6 3 3.5 3 5.5a7 7 0 1 1-14 0c0-1.153.433-2.294 1-3a2.5 2.5 0 0 0 2.5 2.5z"/>`), g.Group(children),
	)
}

func FlameKindling(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2c1 3 2.5 3.5 3.5 4.5A5 5 0 0 1 17 10a5 5 0 1 1-10 0c0-.3 0-.6.1-.9a2 2 0 1 0 3.3-2C8 4.5 11 2 12 2ZM5 22l14-4M5 18l14 4"/>`), g.Group(children),
	)
}

func Flashlight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 6c0 2-2 2-2 4v10a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V10c0-2-2-2-2-4V2h12zM6 6h12m-6 6h0"/>`), g.Group(children),
	)
}

func FlashlightOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 16v4a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V10c0-2-2-2-2-4m1-4h11v4c0 2-2 2-2 4v1m-5-5h7M2 2l20 20"/>`), g.Group(children),
	)
}

func FlaskConical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 2v7.527a2 2 0 0 1-.211.896L4.72 20.55a1 1 0 0 0 .9 1.45h12.76a1 1 0 0 0 .9-1.45l-5.069-10.127A2 2 0 0 1 14 9.527V2M8.5 2h7M7 16h10"/>`), g.Group(children),
	)
}

func FlaskConicalOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 10L4.72 20.55a1 1 0 0 0 .9 1.45h12.76a1 1 0 0 0 .9-1.45l-1.272-2.542M10 2v2.343M14 2v6.343M8.5 2h7M7 16h9M2 2l20 20"/>`), g.Group(children),
	)
}

func FlaskRound(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 2v7.31m4-.01V1.99M8.5 2h7M14 9.3a6.5 6.5 0 1 1-4 0M5.52 16h12.96"/>`), g.Group(children),
	)
}

func FlipHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h3m8-18h3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-3m-4-1v2m0-8v2m0-8v2m0-8v2"/>`), g.Group(children),
	)
}

func FlipHorizontalTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 7l5 5l-5 5V7m18 0l-5 5l5 5V7m-9 13v2m0-8v2m0-8v2m0-8v2"/>`), g.Group(children),
	)
}

func FlipVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 8V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v3m18 8v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-3m1-4H2m8 0H8m8 0h-2m8 0h-2"/>`), g.Group(children),
	)
}

func FlipVerticalTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 3l-5 5l-5-5h10m0 18l-5-5l-5 5h10M4 12H2m8 0H8m8 0h-2m8 0h-2"/>`), g.Group(children),
	)
}

func Flower(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 7.5a4.5 4.5 0 1 1 4.5 4.5M12 7.5A4.5 4.5 0 1 0 7.5 12M12 7.5V9m-4.5 3a4.5 4.5 0 1 0 4.5 4.5M7.5 12H9m7.5 0a4.5 4.5 0 1 1-4.5 4.5m4.5-4.5H15m-3 4.5V15"/><circle cx="12" cy="12" r="3"/><path d="m8 16l1.5-1.5m5-5L16 8M8 8l1.5 1.5m5 5L16 16"/></g>`), g.Group(children),
	)
}

func FlowerTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 5a3 3 0 1 1 3 3m-3-3a3 3 0 1 0-3 3m3-3v1M9 8a3 3 0 1 0 3 3M9 8h1m5 0a3 3 0 1 1-3 3m3-3h-1m-2 3v-1"/><circle cx="12" cy="8" r="2"/><path d="M12 10v12m0 0c4.2 0 7-1.667 7-5c-4.2 0-7 1.667-7 5Zm0 0c-4.2 0-7-1.667-7-5c4.2 0 7 1.667 7 5Z"/></g>`), g.Group(children),
	)
}

func Focus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="3"/><path d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2"/></g>`), g.Group(children),
	)
}

func FoldHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h6m14 0h-6M12 2v2m0 4v2m0 4v2m0 4v2m7-13l-3 3l3 3M5 15l3-3l-3-3"/>`), g.Group(children),
	)
}

func FoldVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22v-6m0-8V2M4 12H2m8 0H8m8 0h-2m8 0h-2m-5 7l-3-3l-3 3m6-14l-3 3l-3-3"/>`), g.Group(children),
	)
}

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/>`), g.Group(children),
	)
}

func FolderArchive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="15" cy="19" r="2"/><path d="M20.9 19.8A2 2 0 0 0 22 18V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2h5.1m5.9-9v-1m0 7v-2"/></g>`), g.Group(children),
	)
}

func FolderCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/><path d="m9 13l2 2l4-4"/></g>`), g.Group(children),
	)
}

func FolderClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="16" cy="16" r="6"/><path d="M7 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2m-6 6v2l1 1"/></g>`), g.Group(children),
	)
}

func FolderClosed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2ZM2 10h20"/>`), g.Group(children),
	)
}

func FolderCog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><path d="M10.3 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v3.3m-.3 8.1l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.2-5.6l.3-.9m.2 7.4l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4"/></g>`), g.Group(children),
	)
}

func FolderCogTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><circle cx="12" cy="13" r="2"/><path d="M12 10v1m0 4v1m2.6-4.5l-.87.5m-3.46 2l-.87.5m5.2 0l-.87-.5m-3.46-2l-.87-.5"/></g>`), g.Group(children),
	)
}

func FolderDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><circle cx="12" cy="13" r="1"/></g>`), g.Group(children),
	)
}

func FolderDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Zm-8-10v6"/><path d="m15 13l-3 3l-3-3"/></g>`), g.Group(children),
	)
}

func FolderEdit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8.4 10.6a2.1 2.1 0 1 1 2.99 2.98L6 19l-4 1l1-3.9Z"/><path d="M2 11.5V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-9.5"/></g>`), g.Group(children),
	)
}

func FolderGit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="13" r="2"/><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Zm-6-7h3M7 13h3"/></g>`), g.Group(children),
	)
}

func FolderGitTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v5"/><circle cx="13" cy="12" r="2"/><path d="M18 19c-2.8 0-5-2.2-5-5v8"/><circle cx="20" cy="19" r="2"/></g>`), g.Group(children),
	)
}

func FolderHeart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v1.5"/><path d="M13.9 17.45c-1.2-1.2-1.14-2.8-.2-3.73a2.43 2.43 0 0 1 3.44 0l.36.34l.34-.34a2.43 2.43 0 0 1 3.45-.01v0c.95.95 1 2.53-.2 3.74L17.5 21Z"/></g>`), g.Group(children),
	)
}

func FolderInput(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 9V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-1m0-4h10"/><path d="m9 16l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func FolderKanban(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Zm4-10v4m4-4v2m4-2v6"/>`), g.Group(children),
	)
}

func FolderKey(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="16" cy="20" r="2"/><path d="M10 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v2m0 4l-4.5 4.5M21 15l1 1"/></g>`), g.Group(children),
	)
}

func FolderLock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="5" x="14" y="17" rx="1"/><path d="M10 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v2.5"/><path d="M20 17v-2a2 2 0 1 0-4 0v2"/></g>`), g.Group(children),
	)
}

func FolderMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 13h6m5 7a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/>`), g.Group(children),
	)
}

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 14l1.5-2.9A2 2 0 0 1 9.24 10H20a2 2 0 0 1 1.94 2.5l-1.54 6a2 2 0 0 1-1.95 1.5H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H18a2 2 0 0 1 2 2v2"/>`), g.Group(children),
	)
}

func FolderOpenDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m6 14l1.45-2.9A2 2 0 0 1 9.24 10H20a2 2 0 0 1 1.94 2.5l-1.55 6a2 2 0 0 1-1.94 1.5H4a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H18a2 2 0 0 1 2 2v2"/><circle cx="14" cy="15" r="1"/></g>`), g.Group(children),
	)
}

func FolderOutput(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 7.5V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H2m0-7h10"/><path d="m5 10l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func FolderPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v6m-3-3h6m5 7a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Z"/>`), g.Group(children),
	)
}

func FolderRoot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 20h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.93a2 2 0 0 1-1.66-.9l-.82-1.2A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13c0 1.1.9 2 2 2Z"/><circle cx="12" cy="13" r="2"/><path d="M12 15v5"/></g>`), g.Group(children),
	)
}

func FolderSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="17" cy="17" r="3"/><path d="M10.7 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v4.1M21 21l-1.5-1.5"/></g>`), g.Group(children),
	)
}

func FolderSearchTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11.5" cy="12.5" r="2.5"/><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Zm-6.7-5.7L15 16"/></g>`), g.Group(children),
	)
}

func FolderSymlink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 9V5c0-1.1.9-2 2-2h3.93a2 2 0 0 1 1.66.9l.82 1.2a2 2 0 0 0 1.66.9H20a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H2"/><path d="m8 16l3-3l-3-3"/><path d="M2 16v-1a2 2 0 0 1 2-2h6"/></g>`), g.Group(children),
	)
}

func FolderSync(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 20H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h3.9a2 2 0 0 1 1.69.9l.81 1.2a2 2 0 0 0 1.67.9H20a2 2 0 0 1 2 2v1"/><path d="M12 10v4h4"/><path d="m12 14l1.5-1.5c.9-.9 2.2-1.5 3.5-1.5s2.6.6 3.5 1.5c.4.4.8 1 1 1.5m.5 8v-4h-4"/><path d="m22 18l-1.5 1.5c-.9.9-2.1 1.5-3.5 1.5s-2.6-.6-3.5-1.5c-.4-.4-.8-1-1-1.5"/></g>`), g.Group(children),
	)
}

func FolderTree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 10a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1h-2.5a1 1 0 0 1-.8-.4l-.9-1.2A1 1 0 0 0 15 3h-2a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1Zm0 11a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-2.9a1 1 0 0 1-.88-.55l-.42-.85a1 1 0 0 0-.92-.6H13a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1ZM3 5a2 2 0 0 0 2 2h3"/><path d="M3 3v13a2 2 0 0 0 2 2h3"/></g>`), g.Group(children),
	)
}

func FolderUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2Zm-8-10v6"/><path d="m9 13l3-3l3 3"/></g>`), g.Group(children),
	)
}

func FolderX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 20a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2h-7.9a2 2 0 0 1-1.69-.9L9.6 3.9A2 2 0 0 0 7.93 3H4a2 2 0 0 0-2 2v13a2 2 0 0 0 2 2ZM9.5 10.5l5 5m0-5l-5 5"/>`), g.Group(children),
	)
}

func Folders(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 17a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2h-3.9a2 2 0 0 1-1.69-.9l-.81-1.2a2 2 0 0 0-1.67-.9H8a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2Z"/><path d="M2 8v11a2 2 0 0 0 2 2h14"/></g>`), g.Group(children),
	)
}

func Footprints(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16v-2.38C4 11.5 2.97 10.5 3 8c.03-2.72 1.49-6 4.5-6C9.37 2 10 3.8 10 5.5c0 3.11-2 5.66-2 8.68V16a2 2 0 1 1-4 0Zm16 4v-2.38c0-2.12 1.03-3.12 1-5.62c-.03-2.72-1.49-6-4.5-6C14.63 6 14 7.8 14 9.5c0 3.11 2 5.66 2 8.68V20a2 2 0 1 0 4 0Zm-4-3h4M4 13h4"/>`), g.Group(children),
	)
}

func Forklift(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 12H5a2 2 0 0 0-2 2v5"/><circle cx="13" cy="19" r="2"/><circle cx="5" cy="19" r="2"/><path d="M8 19h3m5-17v17h6M6 12V7c0-1.1.9-2 2-2h3l5 5"/></g>`), g.Group(children),
	)
}

func FormInput(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="12" x="2" y="6" rx="2"/><path d="M12 12h.01M17 12h.01M7 12h.01"/></g>`), g.Group(children),
	)
}

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 17l5-5l-5-5"/><path d="M4 18v-2a4 4 0 0 1 4-4h12"/></g>`), g.Group(children),
	)
}

func Frame(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 6H2m20 12H2M6 2v20M18 2v20"/>`), g.Group(children),
	)
}

func Framer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 16V9h14V2H5l14 14h-7m-7 0l7 7v-7m-7 0h7"/>`), g.Group(children),
	)
}

func Frown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M16 16s-1.5-2-4-2s-4 2-4 2m1-7h.01M15 9h.01"/></g>`), g.Group(children),
	)
}

func Fuel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 22h12M4 9h10m0 13V4a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v18m10-9h2a2 2 0 0 1 2 2v2a2 2 0 0 0 2 2h0a2 2 0 0 0 2-2V9.83a2 2 0 0 0-.59-1.42L18 5"/>`), g.Group(children),
	)
}

func Fullscreen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2"/><rect width="10" height="8" x="7" y="8" rx="1"/></g>`), g.Group(children),
	)
}

func FunctionSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M9 17c2 0 2.8-1 2.8-2.8V10c0-2 1-3.3 3.2-3m-6 4.2h5.7"/></g>`), g.Group(children),
	)
}

func GalleryHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 3v18"/><rect width="12" height="18" x="6" y="3" rx="2"/><path d="M22 3v18"/></g>`), g.Group(children),
	)
}

func GalleryHorizontalEnd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 7v10M6 5v14"/><rect width="12" height="18" x="10" y="3" rx="2"/></g>`), g.Group(children),
	)
}

func GalleryThumbnails(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="14" x="3" y="3" rx="2"/><path d="M4 21h1m4 0h1m4 0h1m4 0h1"/></g>`), g.Group(children),
	)
}

func GalleryVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 2h18"/><rect width="18" height="12" x="3" y="6" rx="2"/><path d="M3 22h18"/></g>`), g.Group(children),
	)
}

func GalleryVerticalEnd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 2h10M5 6h14"/><rect width="18" height="12" x="3" y="10" rx="2"/></g>`), g.Group(children),
	)
}

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 12h4m-2-2v4m7-1h.01M18 11h.01"/><rect width="20" height="12" x="2" y="6" rx="2"/></g>`), g.Group(children),
	)
}

func GamepadTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 11h4M8 9v4m7-1h.01M18 10h.01m-.69-5H6.68a4 4 0 0 0-3.978 3.59c-.006.052-.01.101-.017.152C2.604 9.416 2 14.456 2 16a3 3 0 0 0 3 3c1 0 1.5-.5 2-1l1.414-1.414A2 2 0 0 1 9.828 16h4.344a2 2 0 0 1 1.414.586L17 18c.5.5 1 1 2 1a3 3 0 0 0 3-3c0-1.545-.604-6.584-.685-7.258c-.007-.05-.011-.1-.017-.151A4 4 0 0 0 17.32 5z"/>`), g.Group(children),
	)
}

func GanttChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6h10M6 12h9m-4 6h7"/>`), g.Group(children),
	)
}

func GanttChartSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 8h7m-8 4h6m-3 4h5"/></g>`), g.Group(children),
	)
}

func Gauge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 14l4-4M3.34 19a10 10 0 1 1 17.32 0"/>`), g.Group(children),
	)
}

func GaugeCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.6 2.7a10 10 0 1 0 5.7 5.7"/><circle cx="12" cy="12" r="2"/><path d="M13.4 10.6L19 5"/></g>`), g.Group(children),
	)
}

func Gavel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 13l-7.5 7.5c-.83.83-2.17.83-3 0c0 0 0 0 0 0a2.12 2.12 0 0 1 0-3L11 10m5 6l6-6M8 8l6-6M9 7l8 8m4-4l-8-8"/>`), g.Group(children),
	)
}

func Gem(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3h12l4 6l-10 13L2 9Z"/><path d="M11 3L8 9l4 13l4-13l-3-6M2 9h20"/></g>`), g.Group(children),
	)
}

func Ghost(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 10h.01M15 10h.01M12 2a8 8 0 0 0-8 8v12l3-3l2.5 2.5L12 19l2.5 2.5L17 19l3 3V10a8 8 0 0 0-8-8z"/>`), g.Group(children),
	)
}

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="4" x="3" y="8" rx="1"/><path d="M12 8v13m7-9v7a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-7m2.5-4a2.5 2.5 0 0 1 0-5A4.8 8 0 0 1 12 8a4.8 8 0 0 1 4.5-5a2.5 2.5 0 0 1 0 5"/></g>`), g.Group(children),
	)
}

func GitBranch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3v12"/><circle cx="18" cy="6" r="3"/><circle cx="6" cy="18" r="3"/><path d="M18 9a9 9 0 0 1-9 9"/></g>`), g.Group(children),
	)
}

func GitBranchPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 3v12m12-6a3 3 0 1 0 0-6a3 3 0 0 0 0 6zM6 21a3 3 0 1 0 0-6a3 3 0 0 0 0 6z"/><path d="M15 6a9 9 0 0 0-9 9m12 0v6m3-3h-6"/></g>`), g.Group(children),
	)
}

func GitCommit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="3"/><path d="M3 12h6m6 0h6"/></g>`), g.Group(children),
	)
}

func GitCompare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M13 6h3a2 2 0 0 1 2 2v7m-7 3H8a2 2 0 0 1-2-2V9"/></g>`), g.Group(children),
	)
}

func GitFork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><circle cx="18" cy="6" r="3"/><path d="M18 9v1a2 2 0 0 1-2 2H8a2 2 0 0 1-2-2V9m6 3v3"/></g>`), g.Group(children),
	)
}

func GitMerge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M6 21V9a9 9 0 0 0 9 9"/></g>`), g.Group(children),
	)
}

func GitPullRequest(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M13 6h3a2 2 0 0 1 2 2v7M6 9v12"/></g>`), g.Group(children),
	)
}

func GitPullRequestClosed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M18 11.5V15m3-12l-6 6m6 0l-6-6M6 9v12"/></g>`), g.Group(children),
	)
}

func GitPullRequestDraft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="18" r="3"/><circle cx="6" cy="6" r="3"/><path d="M18 6V5m0 6v-1M6 9v12"/></g>`), g.Group(children),
	)
}

func Github(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 22v-4a4.8 4.8 0 0 0-1-3.5c3 0 6-2 6-5.5c.08-1.25-.27-2.48-1-3.5c.28-1.15.28-2.35 0-3.5c0 0-1 0-3 1.5c-2.64-.5-5.36-.5-8 0C6 2 5 2 5 2c-.3 1.15-.3 2.35 0 3.5A5.403 5.403 0 0 0 4 9c0 3.5 3 5.5 6 5.5c-.39.49-.68 1.05-.85 1.65c-.17.6-.22 1.23-.15 1.85v4"/><path d="M9 18c-4.51 2-5-2-7-2"/></g>`), g.Group(children),
	)
}

func Gitlab(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m22 13.29l-3.33-10a.42.42 0 0 0-.14-.18a.38.38 0 0 0-.22-.11a.39.39 0 0 0-.23.07a.42.42 0 0 0-.14.18l-2.26 6.67H8.32L6.1 3.26a.42.42 0 0 0-.1-.18a.38.38 0 0 0-.26-.08a.39.39 0 0 0-.23.07a.42.42 0 0 0-.14.18L2 13.29a.74.74 0 0 0 .27.83L12 21l9.69-6.88a.71.71 0 0 0 .31-.83Z"/>`), g.Group(children),
	)
}

func GlassWater(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.2 22H8.8a2 2 0 0 1-2-1.79L5 3h14l-1.81 17.21A2 2 0 0 1 15.2 22Z"/><path d="M6 12a5 5 0 0 1 6 0a5 5 0 0 0 6 0"/></g>`), g.Group(children),
	)
}

func Glasses(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="15" r="4"/><circle cx="18" cy="15" r="4"/><path d="M14 15a2 2 0 0 0-2-2a2 2 0 0 0-2 2m-7.5-2L5 7c.7-1.3 1.4-2 3-2m13.5 8L19 7c-.7-1.3-1.5-2-3-2"/></g>`), g.Group(children),
	)
}

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 2a14.5 14.5 0 0 0 0 20a14.5 14.5 0 0 0 0-20M2 12h20"/></g>`), g.Group(children),
	)
}

func GlobeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.54 15H17a2 2 0 0 0-2 2v4.54M7 3.34V5a3 3 0 0 0 3 3v0a2 2 0 0 1 2 2v0c0 1.1.9 2 2 2v0a2 2 0 0 0 2-2v0c0-1.1.9-2 2-2h3.17M11 21.95V18a2 2 0 0 0-2-2v0a2 2 0 0 1-2-2v-1a2 2 0 0 0-2-2H2.05"/><circle cx="12" cy="12" r="10"/></g>`), g.Group(children),
	)
}

func Goal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 13V2l8 4l-8 4"/><path d="M20.55 10.23A9 9 0 1 1 8 4.94"/><path d="M8 10a5 5 0 1 0 8.9 2.02"/></g>`), g.Group(children),
	)
}

func Grab(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 11.5V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v1.4m0-.4V8a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v2m0-.1V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v5m0 0v0a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v0"/><path d="M18 11v0a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-4a8 8 0 0 1-8-8a2 2 0 1 1 4 0"/></g>`), g.Group(children),
	)
}

func GraduationCap(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 10v6M2 10l10-5l10 5l-10 5z"/><path d="M6 12v5c3 3 9 3 12 0v-5"/></g>`), g.Group(children),
	)
}

func Grape(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 5V2l-5.89 5.89"/><circle cx="16.6" cy="15.89" r="3"/><circle cx="8.11" cy="7.4" r="3"/><circle cx="12.35" cy="11.65" r="3"/><circle cx="13.91" cy="5.85" r="3"/><circle cx="18.15" cy="10.09" r="3"/><circle cx="6.56" cy="13.2" r="3"/><circle cx="10.8" cy="17.44" r="3"/><circle cx="5" cy="19" r="3"/></g>`), g.Group(children),
	)
}

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18M3 15h18M9 3v18m6-18v18"/></g>`), g.Group(children),
	)
}

func GridThreeXThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9h18M3 15h18M9 3v18m6-18v18"/></g>`), g.Group(children),
	)
}

func GridTwoXTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 12h18m-9-9v18"/></g>`), g.Group(children),
	)
}

func Grip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="5" r="1"/><circle cx="19" cy="5" r="1"/><circle cx="5" cy="5" r="1"/><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/><circle cx="12" cy="19" r="1"/><circle cx="19" cy="19" r="1"/><circle cx="5" cy="19" r="1"/></g>`), g.Group(children),
	)
}

func GripHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="9" r="1"/><circle cx="19" cy="9" r="1"/><circle cx="5" cy="9" r="1"/><circle cx="12" cy="15" r="1"/><circle cx="19" cy="15" r="1"/><circle cx="5" cy="15" r="1"/></g>`), g.Group(children),
	)
}

func GripVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="12" r="1"/><circle cx="9" cy="5" r="1"/><circle cx="9" cy="19" r="1"/><circle cx="15" cy="12" r="1"/><circle cx="15" cy="5" r="1"/><circle cx="15" cy="19" r="1"/></g>`), g.Group(children),
	)
}

func Group(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7V5c0-1.1.9-2 2-2h2m10 0h2c1.1 0 2 .9 2 2v2m0 10v2c0 1.1-.9 2-2 2h-2M7 21H5c-1.1 0-2-.9-2-2v-2"/><rect width="7" height="5" x="7" y="7" rx="1"/><rect width="7" height="5" x="10" y="12" rx="1"/></g>`), g.Group(children),
	)
}

func Hammer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 12l-8.5 8.5c-.83.83-2.17.83-3 0c0 0 0 0 0 0a2.12 2.12 0 0 1 0-3L12 9m5.64 6L22 10.64"/><path d="m20.91 11.7l-1.25-1.25c-.6-.6-.93-1.4-.93-2.25v-.86L16.01 4.6a5.56 5.56 0 0 0-3.94-1.64H9l.92.82A6.18 6.18 0 0 1 12 8.4v1.56l2 2h2.47l2.26 1.91"/></g>`), g.Group(children),
	)
}

func Hand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 11V6a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v0m0 4V4a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v2m0 4.5V6a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v8"/><path d="M18 8a2 2 0 1 1 4 0v6a8 8 0 0 1-8 8h-2c-2.8 0-4.5-.86-5.99-2.34l-3.6-3.6a2 2 0 0 1 2.83-2.82L7 15"/></g>`), g.Group(children),
	)
}

func HandMetal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 12.5V10a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v1.4m0-.4V9a2 2 0 1 0-4 0v2m0-.5V5a2 2 0 1 0-4 0v9"/><path d="m7 15l-1.76-1.76a2 2 0 0 0-2.83 2.82l3.6 3.6C7.5 21.14 9.2 22 12 22h2a8 8 0 0 0 8-8V7a2 2 0 1 0-4 0v5"/></g>`), g.Group(children),
	)
}

func HardDrive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12H2m3.45-6.89L2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11zM6 16h.01M10 16h.01"/>`), g.Group(children),
	)
}

func HardDriveDownload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 2v8m4-4l-4 4l-4-4"/><rect width="20" height="8" x="2" y="14" rx="2"/><path d="M6 18h.01M10 18h.01"/></g>`), g.Group(children),
	)
}

func HardDriveUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16 6l-4-4l-4 4m4-4v8"/><rect width="20" height="8" x="2" y="14" rx="2"/><path d="M6 18h.01M10 18h.01"/></g>`), g.Group(children),
	)
}

func HardHat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 18a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v2zm8-8V5a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v5M4 15v-3a6 6 0 0 1 6-6h0m4 0h0a6 6 0 0 1 6 6v3"/>`), g.Group(children),
	)
}

func Hash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9h16M4 15h16M10 3L8 21m8-18l-2 18"/>`), g.Group(children),
	)
}

func Haze(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.2 6.2l1.4 1.4M2 13h2m16 0h2m-4.6-5.4l1.4-1.4M22 17H2m20 4H2m14-8a4 4 0 0 0-8 0m4-8V2.5"/>`), g.Group(children),
	)
}

func HdmiPort(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 9a1 1 0 0 0-1-1H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h1l2 2h12l2-2h1a1 1 0 0 0 1-1ZM7.5 12h9"/>`), g.Group(children),
	)
}

func Heading(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12h12M6 20V4m12 16V4"/>`), g.Group(children),
	)
}

func HeadingFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h8m-8 6V6m8 12V6m5 7v-3h4m-4 7.7c.4.2.8.3 1.3.3c1.5 0 2.7-1.1 2.7-2.5S19.8 13 18.3 13H17"/>`), g.Group(children),
	)
}

func HeadingFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h8m-8 6V6m8 12V6m5 4v4h4m0-4v8"/>`), g.Group(children),
	)
}

func HeadingOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h8m-8 6V6m8 12V6m5 6l3-2v8"/>`), g.Group(children),
	)
}

func HeadingSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 12h8m-8 6V6m8 12V6"/><circle cx="19" cy="16" r="2"/><path d="M20 10c-2 2-3 3.5-3 6"/></g>`), g.Group(children),
	)
}

func HeadingThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h8m-8 6V6m8 12V6m5.5 4.5c1.7-1 3.5 0 3.5 1.5a2 2 0 0 1-2 2m-2 3.5c2 1.5 4 .3 4-1.5a2 2 0 0 0-2-2"/>`), g.Group(children),
	)
}

func HeadingTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h8m-8 6V6m8 12V6m9 12h-4c0-4 4-3 4-6c0-1.5-2-2.5-4-1"/>`), g.Group(children),
	)
}

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 14h3a2 2 0 0 1 2 2v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7a9 9 0 0 1 18 0v7a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2v-3a2 2 0 0 1 2-2h3"/>`), g.Group(children),
	)
}

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2c-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/>`), g.Group(children),
	)
}

func HeartCrack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2c-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/><path d="m12 13l-1-1l2-2l-3-3l2-2"/></g>`), g.Group(children),
	)
}

func HeartHandshake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2c-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/><path d="M12 5L9.04 7.96a2.17 2.17 0 0 0 0 3.08v0c.82.82 2.13.85 3 .07l2.07-1.9a2.82 2.82 0 0 1 3.79 0l2.96 2.66M18 15l-2-2m-1 5l-2-2"/></g>`), g.Group(children),
	)
}

func HeartOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20m-5.5-5.5L12 21l-7-7c-1.5-1.45-3-3.2-3-5.5a5.5 5.5 0 0 1 2.14-4.35M8.76 3.1c1.15.22 2.13.78 3.24 1.9c1.5-1.5 2.74-2 4.5-2A5.5 5.5 0 0 1 22 8.5c0 2.12-1.3 3.78-2.67 5.17"/>`), g.Group(children),
	)
}

func HeartPulse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 14c1.49-1.46 3-3.21 3-5.5A5.5 5.5 0 0 0 16.5 3c-1.76 0-3 .5-4.5 2c-1.5-1.5-2.74-2-4.5-2A5.5 5.5 0 0 0 2 8.5c0 2.3 1.5 4.05 3 5.5l7 7Z"/><path d="M3.22 12H9.5l.5-1l2 4.5l2-7l1.5 3.5h5.27"/></g>`), g.Group(children),
	)
}

func HelpCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9.09 9a3 3 0 0 1 5.83 1c0 2-3 3-3 3m.08 4h.01"/></g>`), g.Group(children),
	)
}

func HelpingHand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 15l5.12-5.12A3 3 0 0 1 10.24 9H13a2 2 0 1 1 0 4h-2.5m4-.68l4.17-4.89a1.88 1.88 0 0 1 2.92 2.36l-4.2 5.94A3 3 0 0 1 14.96 17H9.83a2 2 0 0 0-1.42.59L7 19m-5-5l6 6"/>`), g.Group(children),
	)
}

func Hexagon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 16V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16z"/>`), g.Group(children),
	)
}

func Highlighter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 11l-6 6v3h9l3-3"/><path d="m22 12l-4.6 4.6a2 2 0 0 1-2.8 0l-5.2-5.2a2 2 0 0 1 0-2.8L14 4"/></g>`), g.Group(children),
	)
}

func History(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 9-9a9.75 9.75 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5m4-1v5l4 2"/></g>`), g.Group(children),
	)
}

func Home(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 9l9-7l9 7v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><path d="M9 22V12h6v10"/></g>`), g.Group(children),
	)
}

func Hop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 5.5C19 7 20.5 9 21 11c-2.5.5-5 .5-8.5-1m-7 7.5C7 19 9 20.5 11 21c.5-2.5.5-5-1-8.5m6.5-1c1 2 1 3.5 1 6c-2.5 0-4 0-6-1m8.5-5c1 1.5 2 3.5 2 4.5c-1.5.5-3 0-4.5-.5m-6 4.5c1.5 1 3.5 2 4.5 2c.5-1.5 0-3-.5-4.5m5-1c1 2 1.5 3.5 1.5 5.5c-2 0-3.5-.5-5.5-1.5"/><path d="M4.783 4.782C8.493 1.072 14.5 1 18 5c-1 1-4.5 2-6.5 1.5c1 1.5 1 4 .5 5.5c-1.5.5-4 .5-5.5-.5C7 13.5 6 17 5 18c-4-3.5-3.927-9.508-.217-13.218ZM4.5 4.5L3 3c-.184-.185-.184-.816 0-1"/></g>`), g.Group(children),
	)
}

func HopOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.5 5.5C19 7 20.5 9 21 11c-1.323.265-2.646.39-4.118.226M5.5 17.5C7 19 9 20.5 11 21c.5-2.5.5-5-1-8.5m7.5 5c-2.5 0-4 0-6-1m8.5-5c1 1.5 2 3.5 2 4.5m-10.5 4c1.5 1 3.5 2 4.5 2c.5-1.5 0-3-.5-4.5M22 22c-2 0-3.5-.5-5.5-1.5"/><path d="M4.783 4.782C1.073 8.492 1 14.5 5 18c1-1 2-4.5 1.5-6.5c1.5 1 4 1 5.5.5M8.227 2.57C11.578 1.335 15.453 2.089 18 5c-.88.88-3.7 1.761-5.726 1.618M2 2l20 20"/></g>`), g.Group(children),
	)
}

func Hotel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 2H6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2Z"/><path d="m9 16l.348-.24c1.465-1.013 3.84-1.013 5.304 0L15 16M8 7h.01M16 7h.01M12 7h.01M12 11h.01M16 11h.01M8 11h.01M10 22v-6.5m4 0V22"/></g>`), g.Group(children),
	)
}

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 22h14M5 2h14m-2 20v-4.172a2 2 0 0 0-.586-1.414L12 12l-4.414 4.414A2 2 0 0 0 7 17.828V22M7 2v4.172a2 2 0 0 0 .586 1.414L12 12l4.414-4.414A2 2 0 0 0 17 6.172V2"/>`), g.Group(children),
	)
}

func IceCream(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 11l4.08 10.35a1 1 0 0 0 1.84 0L17 11m0-4A5 5 0 0 0 7 7m10 0a2 2 0 0 1 0 4H7a2 2 0 0 1 0-4"/>`), g.Group(children),
	)
}

func IceCreamTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 17c5 0 8-2.69 8-6H4c0 3.31 3 6 8 6Zm-4 4h8m-4-3v3M5.14 11a3.5 3.5 0 1 1 6.71 0"/><path d="M12.14 11a3.5 3.5 0 1 1 6.71 0"/><path d="M15.5 6.5a3.5 3.5 0 1 0-7 0"/></g>`), g.Group(children),
	)
}

func Image(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><circle cx="9" cy="9" r="2"/><path d="m21 15l-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/></g>`), g.Group(children),
	)
}

func ImageDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="9" r="2"/><path d="M10.3 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v10.8"/><path d="m21 15l-3.1-3.1a2 2 0 0 0-2.814.014L6 21"/><path d="m14 19.5l3 3v-6m0 6l3-3"/></g>`), g.Group(children),
	)
}

func ImageMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7m4 2h6"/><circle cx="9" cy="9" r="2"/><path d="m21 15l-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/></g>`), g.Group(children),
	)
}

func ImageOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 2l20 20M10.41 10.41a2 2 0 1 1-2.83-2.83m5.92 5.92L6 21m12-9l3 3"/><path d="M3.59 3.59A1.99 1.99 0 0 0 3 5v14a2 2 0 0 0 2 2h14c.55 0 1.052-.22 1.41-.59M21 15V5a2 2 0 0 0-2-2H9"/></g>`), g.Group(children),
	)
}

func ImagePlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h7m4 2h6m-3-3v6"/><circle cx="9" cy="9" r="2"/><path d="m21 15l-3.086-3.086a2 2 0 0 0-2.828 0L6 21"/></g>`), g.Group(children),
	)
}

func Import(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 3v12m-4-4l4 4l4-4"/><path d="M8 5H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2h-4"/></g>`), g.Group(children),
	)
}

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12h-6l-2 3h-4l-2-3H2"/><path d="M5.45 5.11L2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0 0-1.79 1.11z"/></g>`), g.Group(children),
	)
}

func Indent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 8l4 4l-4 4m18-4H11m10-6H11m10 12H11"/>`), g.Group(children),
	)
}

func IndianRupee(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 3h12M6 8h12M6 13l8.5 8M6 13h3m0 0c6.667 0 6.667-10 0-10"/>`), g.Group(children),
	)
}

func Infinity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12c-2-2.67-4-4-6-4a4 4 0 1 0 0 8c2 0 4-1.33 6-4Zm0 0c2 2.67 4 4 6 4a4 4 0 0 0 0-8c-2 0-4 1.33-6 4Z"/>`), g.Group(children),
	)
}

func Info(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 16v-4m0-4h.01"/></g>`), g.Group(children),
	)
}

func Instagram(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="5" ry="5"/><path d="M16 11.37A4 4 0 1 1 12.63 8A4 4 0 0 1 16 11.37zm1.5-4.87h.01"/></g>`), g.Group(children),
	)
}

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 4h-9m4 16H5M15 4L9 20"/>`), g.Group(children),
	)
}

func IterationCcw(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 10c0-4.4-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8h8"/><path d="m16 14l4 4l-4 4"/></g>`), g.Group(children),
	)
}

func IterationCw(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 10c0-4.4 3.6-8 8-8s8 3.6 8 8s-3.6 8-8 8H4"/><path d="m8 22l-4-4l4-4"/></g>`), g.Group(children),
	)
}

func JapaneseYen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9.5V21m0-11.5L6 3m6 6.5L18 3M6 15h12M6 11h12"/>`), g.Group(children),
	)
}

func Joystick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 17a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-2ZM6 15v-2m6 2V9"/><circle cx="12" cy="6" r="3"/></g>`), g.Group(children),
	)
}

func Kanban(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 5v11m6-11v6m6-6v14"/>`), g.Group(children),
	)
}

func KanbanSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 7v7m4-7v4m4-4v9"/></g>`), g.Group(children),
	)
}

func KanbanSquareDashed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7v7m4-7v4m4-4v9M5 3a2 2 0 0 0-2 2m6-2h1m4 0h1m4 0a2 2 0 0 1 2 2m0 4v1m0 4v1m0 4a2 2 0 0 1-2 2m-5 0h1m-6 0h1m-5 0a2 2 0 0 1-2-2m0-5v1m0-6v1"/>`), g.Group(children),
	)
}

func Key(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7.5" cy="15.5" r="5.5"/><path d="m21 2l-9.6 9.6m4.1-4.1l3 3L22 7l-3-3"/></g>`), g.Group(children),
	)
}

func KeyRound(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 18v3c0 .6.4 1 1 1h4v-3h3v-3h2l1.4-1.4a6.5 6.5 0 1 0-4-4Z"/><circle cx="16.5" cy="7.5" r=".5"/></g>`), g.Group(children),
	)
}

func KeySquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.4 2.7c.9-.9 2.5-.9 3.4 0l5.5 5.5c.9.9.9 2.5 0 3.4l-3.7 3.7c-.9.9-2.5.9-3.4 0L8.7 9.8c-.9-.9-.9-2.5 0-3.4ZM14 7l3 3m-7.6.6L2 18v3c0 .6.4 1 1 1h4v-3h3v-3h2l1.4-1.4"/>`), g.Group(children),
	)
}

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2" ry="2"/><path d="M6 8h.001M10 8h.001M14 8h.001M18 8h.001M8 12h.001M12 12h.001M16 12h.001M7 16h10"/></g>`), g.Group(children),
	)
}

func Lamp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 2h8l4 10H4L8 2Zm4 10v6m-4 4v-2c0-1.1.9-2 2-2h4a2 2 0 0 1 2 2v2H8Z"/>`), g.Group(children),
	)
}

func LampCeiling(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v5M6 7h12l4 9H2l4-9Zm3.17 9a3 3 0 1 0 5.66 0"/>`), g.Group(children),
	)
}

func LampDesk(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14 5l-3 3l2 7l8-8l-7-2Z"/><path d="m14 5l-3 3l-3-3l3-3l3 3Z"/><path d="M9.5 6.5L4 12l3 6m-4 4v-2c0-1.1.9-2 2-2h4a2 2 0 0 1 2 2v2H3Z"/></g>`), g.Group(children),
	)
}

func LampFloor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 2h6l3 7H6l3-7Zm3 7v13m-3 0h6"/>`), g.Group(children),
	)
}

func LampWallDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 13h6l3 7H8l3-7Zm3 0V8a2 2 0 0 0-2-2H8M4 9h2a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2H4v6Z"/>`), g.Group(children),
	)
}

func LampWallUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4h6l3 7H8l3-7Zm3 7v5a2 2 0 0 1-2 2H8m-4-3h2a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H4v-6Z"/>`), g.Group(children),
	)
}

func LandPlot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 8l6-3l-6-3v10"/><path d="m8 11.99l-5.5 3.14a1 1 0 0 0 0 1.74l8.5 4.86a2 2 0 0 0 2 0l8.5-4.86a1 1 0 0 0 0-1.74L16 12m-9.51.85l11.02 6.3m0-6.3L6.5 19.15"/></g>`), g.Group(children),
	)
}

func Landmark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 22h18M6 18v-7m4 7v-7m4 7v-7m4 7v-7m-6-9l8 5H4z"/>`), g.Group(children),
	)
}

func Languages(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 8l6 6m-7 0l6-6l2-3M2 5h12M7 2h1m14 20l-5-10l-5 10m2-4h6"/>`), g.Group(children),
	)
}

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 16V7a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v9m16 0H4m16 0l1.28 2.55a1 1 0 0 1-.9 1.45H3.62a1 1 0 0 1-.9-1.45L4 16"/>`), g.Group(children),
	)
}

func LaptopTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="3" y="4" rx="2" ry="2"/><path d="M2 20h20"/></g>`), g.Group(children),
	)
}

func Lasso(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 22a5 5 0 0 1-2-4m-1.7-4A6.8 6.8 0 0 1 2 10c0-4.4 4.5-8 10-8s10 3.6 10 8s-4.5 8-10 8a12 12 0 0 1-5-1"/><path d="M5 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4z"/></g>`), g.Group(children),
	)
}

func LassoSelect(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 22a5 5 0 0 1-2-4m2-1.07c.96.43 1.96.74 2.99.91M3.34 14A6.8 6.8 0 0 1 2 10c0-4.42 4.48-8 10-8s10 3.58 10 8a7.19 7.19 0 0 1-.33 2"/><path d="M5 18a2 2 0 1 0 0-4a2 2 0 0 0 0 4zm9.33 4h-.09a.35.35 0 0 1-.24-.32v-10a.34.34 0 0 1 .33-.34c.08 0 .15.03.21.08l7.34 6a.33.33 0 0 1-.21.59h-4.49l-2.57 3.85a.35.35 0 0 1-.28.14v0z"/></g>`), g.Group(children),
	)
}

func Laugh(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M18 13a6 6 0 0 1-6 5a6 6 0 0 1-6-5h12ZM9 9h.01M15 9h.01"/></g>`), g.Group(children),
	)
}

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.83 2.18a2 2 0 0 0-1.66 0L2.6 6.08a1 1 0 0 0 0 1.83l8.58 3.91a2 2 0 0 0 1.66 0l8.58-3.9a1 1 0 0 0 0-1.83ZM22 17.65l-9.17 4.16a2 2 0 0 1-1.66 0L2 17.65m20-5l-9.17 4.16a2 2 0 0 1-1.66 0L2 12.65"/>`), g.Group(children),
	)
}

func LayersThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.83 2.18a2 2 0 0 0-1.66 0L2.6 6.08a1 1 0 0 0 0 1.83l8.58 3.91a2 2 0 0 0 1.66 0l8.58-3.9a1 1 0 0 0 0-1.83Z"/><path d="m6.08 9.5l-3.5 1.6a1 1 0 0 0 0 1.81l8.6 3.91a2 2 0 0 0 1.65 0l8.58-3.9a1 1 0 0 0 0-1.83l-3.5-1.59"/><path d="m6.08 14.5l-3.5 1.6a1 1 0 0 0 0 1.81l8.6 3.91a2 2 0 0 0 1.65 0l8.58-3.9a1 1 0 0 0 0-1.83l-3.5-1.59"/></g>`), g.Group(children),
	)
}

func LayersTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16.02 12l5.48 3.13a1 1 0 0 1 0 1.74L13 21.74a2 2 0 0 1-2 0l-8.5-4.87a1 1 0 0 1 0-1.74L7.98 12"/><path d="M13 13.74a2 2 0 0 1-2 0L2.5 8.87a1 1 0 0 1 0-1.74L11 2.26a2 2 0 0 1 2 0l8.5 4.87a1 1 0 0 1 0 1.74Z"/></g>`), g.Group(children),
	)
}

func Layout(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18M9 21V9"/></g>`), g.Group(children),
	)
}

func LayoutDashboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="9" x="3" y="3" rx="1"/><rect width="7" height="5" x="14" y="3" rx="1"/><rect width="7" height="9" x="14" y="12" rx="1"/><rect width="7" height="5" x="3" y="16" rx="1"/></g>`), g.Group(children),
	)
}

func LayoutGrid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="14" y="3" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/></g>`), g.Group(children),
	)
}

func LayoutList(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/><path d="M14 4h7m-7 5h7m-7 6h7m-7 5h7"/></g>`), g.Group(children),
	)
}

func LayoutPanelLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="18" x="3" y="3" rx="1"/><rect width="7" height="7" x="14" y="3" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/></g>`), g.Group(children),
	)
}

func LayoutPanelTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="7" x="3" y="3" rx="1"/><rect width="7" height="7" x="3" y="14" rx="1"/><rect width="7" height="7" x="14" y="14" rx="1"/></g>`), g.Group(children),
	)
}

func LayoutTemplate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="7" x="3" y="3" rx="1"/><rect width="9" height="7" x="3" y="14" rx="1"/><rect width="5" height="7" x="16" y="14" rx="1"/></g>`), g.Group(children),
	)
}

func Leaf(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 20A7 7 0 0 1 9.8 6.1C15.5 5 17 4.48 19 2c1 2 2 4.18 2 8c0 5.5-4.78 10-10 10Z"/><path d="M2 21c0-3 1.85-5.36 5.08-6C9.5 14.52 12 13 13 12"/></g>`), g.Group(children),
	)
}

func LeafyGreen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 22c1.25-.987 2.27-1.975 3.9-2.2a5.56 5.56 0 0 1 3.8 1.5a4 4 0 0 0 6.187-2.353a3.5 3.5 0 0 0 3.69-5.116A3.5 3.5 0 0 0 20.95 8A3.5 3.5 0 1 0 16 3.05a3.5 3.5 0 0 0-5.831 1.373a3.5 3.5 0 0 0-5.116 3.69a4 4 0 0 0-2.348 6.155C3.499 15.42 4.409 16.712 4.2 18.1C3.926 19.743 3.014 20.732 2 22m0 0L17 7"/>`), g.Group(children),
	)
}

func Library(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 6l4 14M12 6v14M8 8v12M4 4v16"/>`), g.Group(children),
	)
}

func LibraryBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="18" x="3" y="3" rx="1"/><path d="M7 3v18m13.4-2.1c.2.5-.1 1.1-.6 1.3l-1.9.7c-.5.2-1.1-.1-1.3-.6L11.1 5.1c-.2-.5.1-1.1.6-1.3l1.9-.7c.5-.2 1.1.1 1.3.6Z"/></g>`), g.Group(children),
	)
}

func LibrarySquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 7v10m4-10v10m4-10l2 10"/></g>`), g.Group(children),
	)
}

func LifeBuoy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m4.93 4.93l4.24 4.24m5.66 0l4.24-4.24m-4.24 9.9l4.24 4.24m-9.9-4.24l-4.24 4.24"/><circle cx="12" cy="12" r="4"/></g>`), g.Group(children),
	)
}

func Ligature(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 20V8c0-2.2 1.8-4 4-4c1.5 0 2.8.8 3.5 2M6 12h4m4 0h2v8M6 20h4m4 0h4"/>`), g.Group(children),
	)
}

func Lightbulb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 14c.2-1 .7-1.7 1.5-2.5c1-.9 1.5-2.2 1.5-3.5A6 6 0 0 0 6 8c0 1 .2 2.2 1.5 3.5c.7.7 1.3 1.5 1.5 2.5m0 4h6m-5 4h4"/>`), g.Group(children),
	)
}

func LightbulbOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.8 11.2c.8-.9 1.2-2 1.2-3.2a6 6 0 0 0-9.3-5M2 2l20 20M6.3 6.3a4.67 4.67 0 0 0 1.2 5.2c.7.7 1.3 1.5 1.5 2.5m0 4h6m-5 4h4"/>`), g.Group(children),
	)
}

func LineChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 3v18h18"/><path d="m19 9l-5 5l-4-4l-3 3"/></g>`), g.Group(children),
	)
}

func Link(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 13a5 5 0 0 0 7.54.54l3-3a5 5 0 0 0-7.07-7.07l-1.72 1.71"/><path d="M14 11a5 5 0 0 0-7.54-.54l-3 3a5 5 0 0 0 7.07 7.07l1.71-1.71"/></g>`), g.Group(children),
	)
}

func LinkTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17H7A5 5 0 0 1 7 7h2m6 0h2a5 5 0 1 1 0 10h-2m-7-5h8"/>`), g.Group(children),
	)
}

func LinkTwoOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17H7A5 5 0 0 1 7 7m8 0h2a5 5 0 0 1 4 8M8 12h4M2 2l20 20"/>`), g.Group(children),
	)
}

func Linkedin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 8a6 6 0 0 1 6 6v7h-4v-7a2 2 0 0 0-2-2a2 2 0 0 0-2 2v7h-4v-7a6 6 0 0 1 6-6zM2 9h4v12H2z"/><circle cx="4" cy="4" r="2"/></g>`), g.Group(children),
	)
}

func List(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6h13M8 12h13M8 18h13M3 6h.01M3 12h.01M3 18h.01"/>`), g.Group(children),
	)
}

func ListChecks(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 17l2 2l4-4M3 7l2 2l4-4m4 1h8m-8 6h8m-8 6h8"/>`), g.Group(children),
	)
}

func ListEnd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 12H3m13-6H3m7 12H3M21 6v10a2 2 0 0 1-2 2h-5"/><path d="m16 16l-2 2l2 2"/></g>`), g.Group(children),
	)
}

func ListFilter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18M7 12h10m-7 6h4"/>`), g.Group(children),
	)
}

func ListMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 12H3m13-6H3m13 12H3m18-6h-6"/>`), g.Group(children),
	)
}

func ListMusic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15V6m-2.5 12a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5ZM12 12H3m13-6H3m9 12H3"/>`), g.Group(children),
	)
}

func ListOrdered(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 6h11m-11 6h11m-11 6h11M4 6h1v4m-1 0h2m0 8H4c0-1 2-2 2-3s-1-1.5-2-1"/>`), g.Group(children),
	)
}

func ListPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 12H3m13-6H3m13 12H3m15-9v6m3-3h-6"/>`), g.Group(children),
	)
}

func ListRestart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 6H3m4 6H3m4 6H3m9 0a5 5 0 0 0 9-3a4.5 4.5 0 0 0-4.5-4.5c-1.33 0-2.54.54-3.41 1.41L11 14"/><path d="M11 10v4h4"/></g>`), g.Group(children),
	)
}

func ListStart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 12H3m13 6H3m7-12H3m18 12V8a2 2 0 0 0-2-2h-5"/><path d="m16 8l-2-2l2-2"/></g>`), g.Group(children),
	)
}

func ListTodo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="6" x="3" y="5" rx="1"/><path d="m3 17l2 2l4-4m4-9h8m-8 6h8m-8 6h8"/></g>`), g.Group(children),
	)
}

func ListTree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12h-8m8-6H8m13 12h-8M3 6v4c0 1.1.9 2 2 2h3"/><path d="M3 10v6c0 1.1.9 2 2 2h3"/></g>`), g.Group(children),
	)
}

func ListVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12H3m13-6H3m9 12H3m13-6l5 3l-5 3v-6Z"/>`), g.Group(children),
	)
}

func ListX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 12H3m13-6H3m13 12H3m16-8l-4 4m0-4l4 4"/>`), g.Group(children),
	)
}

func Loader(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v4m0 12v4M4.93 4.93l2.83 2.83m8.48 8.48l2.83 2.83M2 12h4m12 0h4M4.93 19.07l2.83-2.83m8.48-8.48l2.83-2.83"/>`), g.Group(children),
	)
}

func LoaderTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 1 1-6.219-8.56"/>`), g.Group(children),
	)
}

func Locate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12h3m14 0h3M12 2v3m0 14v3"/><circle cx="12" cy="12" r="7"/></g>`), g.Group(children),
	)
}

func LocateFixed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12h3m14 0h3M12 2v3m0 14v3"/><circle cx="12" cy="12" r="7"/><circle cx="12" cy="12" r="3"/></g>`), g.Group(children),
	)
}

func LocateOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h3m14 0h3M12 2v3m0 14v3M7.11 7.11C5.83 8.39 5 10.1 5 12c0 3.87 3.13 7 7 7c1.9 0 3.61-.83 4.89-2.11m1.82-2.93c.19-.63.29-1.29.29-1.96c0-3.87-3.13-7-7-7c-.67 0-1.33.1-1.96.29M2 2l20 20"/>`), g.Group(children),
	)
}

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 10 0v4"/></g>`), g.Group(children),
	)
}

func LockKeyhole(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="16" r="1"/><rect width="18" height="12" x="3" y="10" rx="2"/><path d="M7 10V7a5 5 0 0 1 10 0v3"/></g>`), g.Group(children),
	)
}

func LogIn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3h4a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2h-4m-5-4l5-5l-5-5m5 5H3"/>`), g.Group(children),
	)
}

func LogOut(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4m7 14l5-5l-5-5m5 5H9"/>`), g.Group(children),
	)
}

func Lollipop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3M11 11a2 2 0 0 0 4 0a4 4 0 0 0-8 0a6 6 0 0 0 12 0"/></g>`), g.Group(children),
	)
}

func Luggage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 20h0a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h0"/><path d="M8 18V4a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v14m-6 2h4"/><circle cx="16" cy="20" r="2"/><circle cx="8" cy="20" r="2"/></g>`), g.Group(children),
	)
}

func MSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 16V8l4 4l4-4v8"/></g>`), g.Group(children),
	)
}

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 15l-4-4l6.75-6.77a7.79 7.79 0 0 1 11 11L13 22l-4-4l6.39-6.36a2.14 2.14 0 0 0-3-3L6 15M5 8l4 4m3 3l4 4"/>`), g.Group(children),
	)
}

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7"/></g>`), g.Group(children),
	)
}

func MailCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 13V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m14 12l2 2l4-4"/></g>`), g.Group(children),
	)
}

func MailMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 15V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m14 12h6"/></g>`), g.Group(children),
	)
}

func MailOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.2 8.4c.5.38.8.97.8 1.6v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V10a2 2 0 0 1 .8-1.6l8-6a2 2 0 0 1 2.4 0l8 6Z"/><path d="m22 10l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 10"/></g>`), g.Group(children),
	)
}

func MailPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 13V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h8"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m17 9v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func MailQuestion(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 10.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h12.5"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m16 8.28c.2-.4.5-.8.9-1a2.1 2.1 0 0 1 2.6.4c.3.4.5.8.5 1.3c0 1.3-2 2-2 2M20 22v.01"/></g>`), g.Group(children),
	)
}

func MailSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 12.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h7.5"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m16 14a3 3 0 1 0 0-6a3 3 0 0 0 0 6v0Z"/><circle cx="18" cy="18" r="3"/><path d="m22 22l-1.5-1.5"/></g>`), g.Group(children),
	)
}

func MailWarning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 10.5V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h12.5"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m18 7v4m0 4v.01"/></g>`), g.Group(children),
	)
}

func MailX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 13V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v12c0 1.1.9 2 2 2h9"/><path d="m22 7l-8.97 5.7a1.94 1.94 0 0 1-2.06 0L2 7m15 10l4 4m0-4l-4 4"/></g>`), g.Group(children),
	)
}

func Mailbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V9.5C2 7 4 5 6.5 5H18c2.2 0 4 1.8 4 4v8Z"/><path d="M15 9h3v2M6.5 5C9 5 11 7 11 9.5V17a2 2 0 0 1-2 2v0m-3-9h1"/></g>`), g.Group(children),
	)
}

func Mails(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="13" x="6" y="4" rx="2"/><path d="m22 7l-7.1 3.78c-.57.3-1.23.3-1.8 0L6 7M2 8v11c0 1.1.9 2 2 2h14"/></g>`), g.Group(children),
	)
}

func Map(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 6l6-3l6 3l6-3v15l-6 3l-6-3l-6 3zm6-3v15m6-12v15"/>`), g.Group(children),
	)
}

func MapPin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 10c0 6-8 12-8 12s-8-6-8-12a8 8 0 0 1 16 0Z"/><circle cx="12" cy="10" r="3"/></g>`), g.Group(children),
	)
}

func MapPinOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.43 5.43A8.06 8.06 0 0 0 4 10c0 6 8 12 8 12a29.94 29.94 0 0 0 5-5m2.18-3.48A8.66 8.66 0 0 0 20 10a8 8 0 0 0-8-8a7.88 7.88 0 0 0-3.52.82"/><path d="M9.13 9.13A2.78 2.78 0 0 0 9 10a3 3 0 0 0 3 3a2.78 2.78 0 0 0 .87-.13m2.03-3.62a3 3 0 0 0-2.15-2.16M2 2l20 20"/></g>`), g.Group(children),
	)
}

func MapPinned(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 8c0 4.5-6 9-6 9s-6-4.5-6-9a6 6 0 0 1 12 0"/><circle cx="12" cy="8" r="2"/><path d="M8.835 14H5a1 1 0 0 0-.9.7l-2 6c-.1.1-.1.2-.1.3c0 .6.4 1 1 1h18c.6 0 1-.4 1-1c0-.1 0-.2-.1-.3l-2-6a1 1 0 0 0-.9-.7h-3.835"/></g>`), g.Group(children),
	)
}

func Martini(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 22h8m-4-11v11m7-19l-7 8l-7-8Z"/>`), g.Group(children),
	)
}

func Maximize(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3H5a2 2 0 0 0-2 2v3m18 0V5a2 2 0 0 0-2-2h-3M3 16v3a2 2 0 0 0 2 2h3m8 0h3a2 2 0 0 0 2-2v-3"/>`), g.Group(children),
	)
}

func MaximizeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 3h6v6M9 21H3v-6M21 3l-7 7M3 21l7-7"/>`), g.Group(children),
	)
}

func Medal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7.21 15L2.66 7.14a2 2 0 0 1 .13-2.2L4.4 2.8A2 2 0 0 1 6 2h12a2 2 0 0 1 1.6.8l1.6 2.14a2 2 0 0 1 .14 2.2L16.79 15M11 12L5.12 2.2M13 12l5.88-9.8M8 7h8"/><circle cx="12" cy="17" r="5"/><path d="M12 18v-2h-.5"/></g>`), g.Group(children),
	)
}

func Megaphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 11l18-5v12L3 14v-3zm8.6 5.8a3 3 0 1 1-5.8-1.6"/>`), g.Group(children),
	)
}

func MegaphoneOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.26 9.26L3 11v3l14.14 3.14m3.86-1.8V6l-7.31 2.03M11.6 16.8a3 3 0 1 1-5.8-1.6M2 2l20 20"/>`), g.Group(children),
	)
}

func Meh(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 15h8M9 9h.01M15 9h.01"/></g>`), g.Group(children),
	)
}

func MemoryStick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 19v-3m4 3v-3m4 3v-3m4 3v-3M8 11V9m8 2V9m-4 2V9M2 15h20M2 7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v1.1a2 2 0 0 0 0 3.837V17a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-5.1a2 2 0 0 0 0-3.837Z"/>`), g.Group(children),
	)
}

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12h16M4 6h16M4 18h16"/>`), g.Group(children),
	)
}

func MenuSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 8h10M7 12h10M7 16h10"/></g>`), g.Group(children),
	)
}

func Merge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 6l4-4l4 4"/><path d="M12 2v10.3a4 4 0 0 1-1.172 2.872L4 22m16 0l-5-5"/></g>`), g.Group(children),
	)
}

func MessageCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 21l1.9-5.7a8.5 8.5 0 1 1 3.8 3.8z"/>`), g.Group(children),
	)
}

func MessageSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2z"/>`), g.Group(children),
	)
}

func MessageSquareDashed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6V5c0-1.1.9-2 2-2h2m4 0h3m4 0h1c1.1 0 2 .9 2 2m0 4v2m0 4c0 1.1-.9 2-2 2h-1m-4 0h-3m-4 0l-4 4v-5m0-4v-2"/>`), g.Group(children),
	)
}

func MessageSquarePlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2zM9 10h6m-3-3v6"/>`), g.Group(children),
	)
}

func MessagesSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 9a2 2 0 0 1-2 2H6l-4 4V4c0-1.1.9-2 2-2h8a2 2 0 0 1 2 2v5Zm4 0h2a2 2 0 0 1 2 2v11l-4-4h-6a2 2 0 0 1-2-2v-1"/>`), g.Group(children),
	)
}

func Mic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 2a3 3 0 0 0-3 3v7a3 3 0 0 0 6 0V5a3 3 0 0 0-3-3Z"/><path d="M19 10v2a7 7 0 0 1-14 0v-2m7 9v3"/></g>`), g.Group(children),
	)
}

func MicOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 2l20 20m-3.11-8.77A7.12 7.12 0 0 0 19 12v-2M5 10v2a7 7 0 0 0 12 5m-2-7.66V5a3 3 0 0 0-5.68-1.33"/><path d="M9 9v3a3 3 0 0 0 5.12 2.12M12 19v3"/></g>`), g.Group(children),
	)
}

func MicTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 8l-9.04 9.06a2.82 2.82 0 1 0 3.98 3.98L16 12"/><circle cx="17" cy="7" r="5"/></g>`), g.Group(children),
	)
}

func Microscope(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18h8M3 22h18m-7 0a7 7 0 1 0 0-14h-1m-4 6h2m-2-2a2 2 0 0 1-2-2V6h6v4a2 2 0 0 1-2 2Zm3-6V3a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v3"/>`), g.Group(children),
	)
}

func Microwave(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="15" x="2" y="4" rx="2"/><rect width="8" height="7" x="6" y="8" rx="1"/><path d="M18 8v7M6 19v2m12-2v2"/></g>`), g.Group(children),
	)
}

func Milestone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 6H5a2 2 0 0 0-2 2v3a2 2 0 0 0 2 2h13l4-3.5L18 6Zm-6 7v8m0-18v3"/>`), g.Group(children),
	)
}

func Milk(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 2h8M9 2v2.789a4 4 0 0 1-.672 2.219l-.656.984A4 4 0 0 0 7 10.212V20a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-9.789a4 4 0 0 0-.672-2.219l-.656-.984A4 4 0 0 1 15 4.788V2"/><path d="M7 15a6.472 6.472 0 0 1 5 0a6.47 6.47 0 0 0 5 0"/></g>`), g.Group(children),
	)
}

func MilkOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 2h8M9 2v1.343M15 2v2.789a4 4 0 0 0 .672 2.219l.656.984a4 4 0 0 1 .672 2.22v1.131M7.8 7.8l-.128.192A4 4 0 0 0 7 10.212V20a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-3"/><path d="M7 15a6.47 6.47 0 0 1 5 0a6.472 6.472 0 0 0 3.435.435M2 2l20 20"/></g>`), g.Group(children),
	)
}

func Minimize(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3v3a2 2 0 0 1-2 2H3m18 0h-3a2 2 0 0 1-2-2V3M3 16h3a2 2 0 0 1 2 2v3m8 0v-3a2 2 0 0 1 2-2h3"/>`), g.Group(children),
	)
}

func MinimizeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 14h6v6m10-10h-6V4m0 6l7-7M3 21l7-7"/>`), g.Group(children),
	)
}

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14"/>`), g.Group(children),
	)
}

func MinusCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 12h8"/></g>`), g.Group(children),
	)
}

func MinusSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 12h8"/></g>`), g.Group(children),
	)
}

func Monitor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M8 21h8m-4-4v4"/></g>`), g.Group(children),
	)
}

func MonitorCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 10l2 2l4-4"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`), g.Group(children),
	)
}

func MonitorDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="19" cy="6" r="3"/><path d="M22 12v3a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h9m-1 14v4m-4 0h8"/></g>`), g.Group(children),
	)
}

func MonitorDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 13V7m3 3l-3 3l-3-3"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`), g.Group(children),
	)
}

func MonitorOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17H4a2 2 0 0 1-2-2V5c0-1.5 1-2 1-2m19 12V5a2 2 0 0 0-2-2H9M8 21h8m-4-4v4M2 2l20 20"/>`), g.Group(children),
	)
}

func MonitorPause(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 13V7m4 6V7"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`), g.Group(children),
	)
}

func MonitorPlay(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 7l5 3l-5 3Z"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`), g.Group(children),
	)
}

func MonitorSmartphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 8V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h8m-2 4v-3.96v3.15M7 19h5"/><rect width="6" height="10" x="16" y="12" rx="2"/></g>`), g.Group(children),
	)
}

func MonitorSpeaker(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.5 20H8m9-11h.01"/><rect width="10" height="16" x="12" y="4" rx="2"/><path d="M8 6H4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h4"/><circle cx="17" cy="15" r="1"/></g>`), g.Group(children),
	)
}

func MonitorStop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 7h6v6H9z"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`), g.Group(children),
	)
}

func MonitorUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 10l3-3l3 3m-3 3V7"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`), g.Group(children),
	)
}

func MonitorX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m14.5 12.5l-5-5m0 5l5-5"/><rect width="20" height="14" x="2" y="3" rx="2"/><path d="M12 17v4m-4 0h8"/></g>`), g.Group(children),
	)
}

func Moon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3a6 6 0 0 0 9 9a9 9 0 1 1-9-9Z"/>`), g.Group(children),
	)
}

func MoonStar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3a6 6 0 0 0 9 9a9 9 0 1 1-9-9Zm7 0v4m2-2h-4"/>`), g.Group(children),
	)
}

func MoreHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="1"/><circle cx="19" cy="12" r="1"/><circle cx="5" cy="12" r="1"/></g>`), g.Group(children),
	)
}

func MoreVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="1"/><circle cx="12" cy="5" r="1"/><circle cx="12" cy="19" r="1"/></g>`), g.Group(children),
	)
}

func Mountain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 3l4 8l5-5l5 15H2L8 3z"/>`), g.Group(children),
	)
}

func MountainSnow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 3l4 8l5-5l5 15H2L8 3z"/><path d="M4.14 15.08c2.62-1.57 5.24-1.43 7.86.42c2.74 1.94 5.49 2 8.23.19"/></g>`), g.Group(children),
	)
}

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="20" x="5" y="2" rx="7"/><path d="M12 6v4"/></g>`), g.Group(children),
	)
}

func MousePointer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l7.07 16.97l2.51-7.39l7.39-2.51L3 3zm10 10l6 6"/>`), g.Group(children),
	)
}

func MousePointerClick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 9l5 12l1.8-5.2L21 14ZM7.2 2.2L8 5.1M5.1 8l-2.9-.8M14 4.1L12 6m-6 6l-1.9 2"/>`), g.Group(children),
	)
}

func MousePointerSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 11V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h6"/><path d="m12 12l4 10l1.7-4.3L22 16Z"/></g>`), g.Group(children),
	)
}

func MousePointerSquareDashed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3a2 2 0 0 0-2 2m16-2a2 2 0 0 1 2 2m-9 7l4 10l1.7-4.3L22 16Zm-7 9a2 2 0 0 1-2-2M9 3h1M9 21h2m3-18h1M3 9v1m18-1v2M3 14v1"/>`), g.Group(children),
	)
}

func MousePointerTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l7.07 17l2.51-7.39L21 11.07z"/>`), g.Group(children),
	)
}

func Move(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 9l-3 3l3 3M9 5l3-3l3 3m0 14l-3 3l-3-3M19 9l3 3l-3 3M2 12h20M12 2v20"/>`), g.Group(children),
	)
}

func MoveDiagonal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 5h6v6m-8 8H5v-6m14-8L5 19"/>`), g.Group(children),
	)
}

func MoveDiagonalTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 11V5h6m8 8v6h-6M5 5l14 14"/>`), g.Group(children),
	)
}

func MoveDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 18l4 4l4-4M12 2v20"/>`), g.Group(children),
	)
}

func MoveDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 19H5v-6m14-8L5 19"/>`), g.Group(children),
	)
}

func MoveDownRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 13v6h-6M5 5l14 14"/>`), g.Group(children),
	)
}

func MoveHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 8l4 4l-4 4M6 8l-4 4l4 4m-4-4h20"/>`), g.Group(children),
	)
}

func MoveLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 8l-4 4l4 4m-4-4h20"/>`), g.Group(children),
	)
}

func MoveRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 8l4 4l-4 4M2 12h20"/>`), g.Group(children),
	)
}

func MoveThreeD(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 3v16h16M5 19l6-6"/><path d="m2 6l3-3l3 3m10 10l3 3l-3 3"/></g>`), g.Group(children),
	)
}

func MoveUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 6l4-4l4 4m-4-4v20"/>`), g.Group(children),
	)
}

func MoveUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 11V5h6M5 5l14 14"/>`), g.Group(children),
	)
}

func MoveUpRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 5h6v6m0-6L5 19"/>`), g.Group(children),
	)
}

func MoveVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 18l4 4l4-4M8 6l4-4l4 4m-4-4v20"/>`), g.Group(children),
	)
}

func Music(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 18V5l12-2v13"/><circle cx="6" cy="18" r="3"/><circle cx="18" cy="16" r="3"/></g>`), g.Group(children),
	)
}

func MusicFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 18V5l12-2v13M9 9l12-2"/><circle cx="6" cy="18" r="3"/><circle cx="18" cy="16" r="3"/></g>`), g.Group(children),
	)
}

func MusicThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="18" r="4"/><path d="M16 18V2"/></g>`), g.Group(children),
	)
}

func MusicTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="8" cy="18" r="4"/><path d="M12 18V2l7 4"/></g>`), g.Group(children),
	)
}

func Navigation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 11l19-9l-9 19l-2-8l-8-2z"/>`), g.Group(children),
	)
}

func NavigationOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.43 8.43L3 11l8 2l2 8l2.57-5.43m1.82-3.84L22 2l-9.73 4.61M2 2l20 20"/>`), g.Group(children),
	)
}

func NavigationTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 2l7 19l-7-4l-7 4l7-19z"/>`), g.Group(children),
	)
}

func NavigationTwoOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.31 9.31L5 21l7-4l7 4l-1.17-3.17m-3.3-8.95L12 2l-1.17 3.17M2 2l20 20"/>`), g.Group(children),
	)
}

func Network(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="6" x="16" y="16" rx="1"/><rect width="6" height="6" x="2" y="16" rx="1"/><rect width="6" height="6" x="9" y="2" rx="1"/><path d="M5 16v-3a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v3m-7-4V8"/></g>`), g.Group(children),
	)
}

func Newspaper(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22h16a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2H8a2 2 0 0 0-2 2v16a2 2 0 0 1-2 2Zm0 0a2 2 0 0 1-2-2v-9c0-1.1.9-2 2-2h2m12 5h-8m5 4h-5"/><path d="M10 6h8v4h-8V6Z"/></g>`), g.Group(children),
	)
}

func Nfc(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8.32a7.43 7.43 0 0 1 0 7.36m3.46-9.47a11.76 11.76 0 0 1 0 11.58M12.91 4.1a15.91 15.91 0 0 1 .01 15.8M16.37 2a20.16 20.16 0 0 1 0 20"/>`), g.Group(children),
	)
}

func Nut(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4V2m-7 8v4a7.004 7.004 0 0 0 5.277 6.787c.412.104.802.292 1.102.592L12 22l.621-.621c.3-.3.69-.488 1.102-.592A7.003 7.003 0 0 0 19 14v-4"/><path d="M12 4C8 4 4.5 6 4 8c-.243.97-.919 1.952-2 3c1.31-.082 1.972-.29 3-1c.54.92.982 1.356 2 2c1.452-.647 1.954-1.098 2.5-2c.595.995 1.151 1.427 2.5 2c1.31-.621 1.862-1.058 2.5-2c.629.977 1.162 1.423 2.5 2c1.209-.548 1.68-.967 2-2c1.032.916 1.683 1.157 3 1c-1.297-1.036-1.758-2.03-2-3c-.5-2-4-4-8-4Z"/></g>`), g.Group(children),
	)
}

func NutOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 4V2m-7 8v4a7.004 7.004 0 0 0 5.277 6.787c.412.104.802.292 1.102.592L12 22l.621-.621c.3-.3.69-.488 1.102-.592a7.01 7.01 0 0 0 4.125-2.939M19 10v3.343"/><path d="M12 12c-1.349-.573-1.905-1.005-2.5-2c-.546.902-1.048 1.353-2.5 2c-1.018-.644-1.46-1.08-2-2c-1.028.71-1.69.918-3 1c1.081-1.048 1.757-2.03 2-3c.194-.776.84-1.551 1.79-2.21m11.654 5.997c.887-.457 1.28-.891 1.556-1.787c1.032.916 1.683 1.157 3 1c-1.297-1.036-1.758-2.03-2-3c-.5-2-4-4-8-4c-.74 0-1.461.068-2.15.192M2 2l20 20"/></g>`), g.Group(children),
	)
}

func Octagon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.86 2h8.28L22 7.86v8.28L16.14 22H7.86L2 16.14V7.86L7.86 2z"/>`), g.Group(children),
	)
}

func Option(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h6l6 18h6M14 3h7"/>`), g.Group(children),
	)
}

func Orbit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="3"/><circle cx="19" cy="5" r="2"/><circle cx="5" cy="19" r="2"/><path d="M10.4 21.9a10 10 0 0 0 9.941-15.416M13.5 2.1a10 10 0 0 0-9.841 15.416"/></g>`), g.Group(children),
	)
}

func Outdent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 8l-4 4l4 4m14-4H11m10-6H11m10 12H11"/>`), g.Group(children),
	)
}

func Package(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m7.5 4.27l9 5.15M21 8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l7-4A2 2 0 0 0 21 16Z"/><path d="m3.3 7l8.7 5l8.7-5M12 22V12"/></g>`), g.Group(children),
	)
}

func PackageCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m16 16l2 2l4-4"/><path d="M21 10V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14M7.5 4.27l9 5.15"/><path d="M3.29 7L12 12l8.71-5M12 22V12"/></g>`), g.Group(children),
	)
}

func PackageMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 16h6m-1-6V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14M7.5 4.27l9 5.15"/><path d="M3.29 7L12 12l8.71-5M12 22V12"/></g>`), g.Group(children),
	)
}

func PackageOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.91 8.84L8.56 2.23a1.93 1.93 0 0 0-1.81 0L3.1 4.13a2.12 2.12 0 0 0-.05 3.69l12.22 6.93a2 2 0 0 0 1.94 0L21 12.51a2.12 2.12 0 0 0-.09-3.67Z"/><path d="m3.09 8.84l12.35-6.61a1.93 1.93 0 0 1 1.81 0l3.65 1.9a2.12 2.12 0 0 1 .1 3.69L8.73 14.75a2 2 0 0 1-1.94 0L3 12.51a2.12 2.12 0 0 1 .09-3.67ZM12 22v-9"/><path d="M20 13.5v3.37a2.06 2.06 0 0 1-1.11 1.83l-6 3.08a1.93 1.93 0 0 1-1.78 0l-6-3.08A2.06 2.06 0 0 1 4 16.87V13.5"/></g>`), g.Group(children),
	)
}

func PackagePlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 16h6m-3-3v6m2-9V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14M7.5 4.27l9 5.15"/><path d="M3.29 7L12 12l8.71-5M12 22V12"/></g>`), g.Group(children),
	)
}

func PackageSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 10V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14M7.5 4.27l9 5.15"/><path d="M3.29 7L12 12l8.71-5M12 22V12"/><circle cx="18.5" cy="15.5" r="2.5"/><path d="M20.27 17.27L22 19"/></g>`), g.Group(children),
	)
}

func PackageTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9Zm0 0l2.45-4.9A2 2 0 0 1 7.24 3h9.52a2 2 0 0 1 1.8 1.1L21 9m-9-6v6"/>`), g.Group(children),
	)
}

func PackageX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 10V8a2 2 0 0 0-1-1.73l-7-4a2 2 0 0 0-2 0l-7 4A2 2 0 0 0 3 8v8a2 2 0 0 0 1 1.73l7 4a2 2 0 0 0 2 0l2-1.14M7.5 4.27l9 5.15"/><path d="M3.29 7L12 12l8.71-5M12 22V12m5 1l5 5m-5 0l5-5"/></g>`), g.Group(children),
	)
}

func PaintBucket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 11l-8-8l-8.6 8.6a2 2 0 0 0 0 2.8l5.2 5.2c.8.8 2 .8 2.8 0L19 11ZM5 2l5 5m-8 6h15m5 7a2 2 0 1 1-4 0c0-1.6 1.7-2.4 2-4c.3 1.6 2 2.4 2 4Z"/>`), g.Group(children),
	)
}

func Paintbrush(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18.37 2.63L14 7l-1.59-1.59a2 2 0 0 0-2.82 0L8 7l9 9l1.59-1.59a2 2 0 0 0 0-2.82L17 10l4.37-4.37a2.12 2.12 0 1 0-3-3Z"/><path d="M9 8c-2 3-4 3.5-7 4l8 10c2-1 6-5 6-7m-1.5 2.5L4.5 15"/></g>`), g.Group(children),
	)
}

func PaintbrushTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 19.9V16h3a2 2 0 0 0 2-2v-2H5v2c0 1.1.9 2 2 2h3v3.9a2 2 0 1 0 4 0ZM6 12V2h12v10M14 2v4m-4-4v2"/>`), g.Group(children),
	)
}

func Palette(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="13.5" cy="6.5" r=".5"/><circle cx="17.5" cy="10.5" r=".5"/><circle cx="8.5" cy="7.5" r=".5"/><circle cx="6.5" cy="12.5" r=".5"/><path d="M12 2C6.5 2 2 6.5 2 12s4.5 10 10 10c.926 0 1.648-.746 1.648-1.688c0-.437-.18-.835-.437-1.125c-.29-.289-.438-.652-.438-1.125a1.64 1.64 0 0 1 1.668-1.668h1.996c3.051 0 5.555-2.503 5.555-5.554C21.965 6.012 17.461 2 12 2z"/></g>`), g.Group(children),
	)
}

func Palmtree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 8c0-2.76-2.46-5-5.5-5S2 5.24 2 8h2l1-1l1 1h4m3-.86A5.82 5.82 0 0 1 16.5 6c3.04 0 5.5 2.24 5.5 5h-3l-1-1l-1 1h-3"/><path d="M5.89 9.71c-2.15 2.15-2.3 5.47-.35 7.43l4.24-4.25l.7-.7l.71-.71l2.12-2.12c-1.95-1.96-5.27-1.8-7.42.35z"/><path d="M11 15.5c.5 2.5-.17 4.5-1 6.5h4c2-5.5-.5-12-1-14"/></g>`), g.Group(children),
	)
}

func PanelBottom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 15h18"/></g>`), g.Group(children),
	)
}

func PanelBottomClose(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 15h18m-6-7l-3 3l-3-3"/></g>`), g.Group(children),
	)
}

func PanelBottomInactive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M14 15h1m4 0h2M3 15h2m4 0h1"/></g>`), g.Group(children),
	)
}

func PanelBottomOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 15h18M9 10l3-3l3 3"/></g>`), g.Group(children),
	)
}

func PanelLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M9 3v18"/></g>`), g.Group(children),
	)
}

func PanelLeftClose(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M9 3v18m7-6l-3-3l3-3"/></g>`), g.Group(children),
	)
}

func PanelLeftInactive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 14v1m0 4v2M9 3v2m0 4v1"/></g>`), g.Group(children),
	)
}

func PanelLeftOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M9 3v18m5-12l3 3l-3 3"/></g>`), g.Group(children),
	)
}

func PanelRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M15 3v18"/></g>`), g.Group(children),
	)
}

func PanelRightClose(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M15 3v18M8 9l3 3l-3 3"/></g>`), g.Group(children),
	)
}

func PanelRightInactive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M15 14v1m0 4v2m0-18v2m0 4v1"/></g>`), g.Group(children),
	)
}

func PanelRightOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M15 3v18m-5-6l-3-3l3-3"/></g>`), g.Group(children),
	)
}

func PanelTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18"/></g>`), g.Group(children),
	)
}

func PanelTopClose(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18M9 16l3-3l3 3"/></g>`), g.Group(children),
	)
}

func PanelTopInactive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M14 9h1m4 0h2M3 9h2m4 0h1"/></g>`), g.Group(children),
	)
}

func PanelTopOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18m-6 5l-3 3l-3-3"/></g>`), g.Group(children),
	)
}

func Paperclip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21.44 11.05l-9.19 9.19a6 6 0 0 1-8.49-8.49l8.57-8.57A4 4 0 1 1 18 8.84l-8.59 8.57a2 2 0 0 1-2.83-2.83l8.49-8.48"/>`), g.Group(children),
	)
}

func Parentheses(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 21s-4-3-4-9s4-9 4-9m8 0s4 3 4 9s-4 9-4 9"/>`), g.Group(children),
	)
}

func ParkingCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9 17V7h4a3 3 0 0 1 0 6H9"/></g>`), g.Group(children),
	)
}

func ParkingCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m5 5l14 14m-6-6a3 3 0 1 0 0-6H9v2m0 8v-2.34"/></g>`), g.Group(children),
	)
}

func ParkingMeter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 9a3 3 0 1 1 6 0m-3 3v3m-1 0h2"/><path d="M19 9a7 7 0 1 0-13.6 2.3C6.4 14.4 8 19 8 19h8s1.6-4.6 2.6-7.7c.3-.8.4-1.5.4-2.3m-7 10v3"/></g>`), g.Group(children),
	)
}

func ParkingSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M9 17V7h4a3 3 0 0 1 0 6H9"/></g>`), g.Group(children),
	)
}

func ParkingSquareOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3.6 3.6A2 2 0 0 1 5 3h14a2 2 0 0 1 2 2v14a2 2 0 0 1-.59 1.41M3 8.7V19a2 2 0 0 0 2 2h10.3M2 2l20 20"/><path d="M13 13a3 3 0 1 0 0-6H9v2m0 8v-2.3"/></g>`), g.Group(children),
	)
}

func PartyPopper(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.8 11.3L2 22l10.7-3.79M4 3h.01M22 8h.01M15 2h.01M22 20h.01M22 2l-2.24.75a2.9 2.9 0 0 0-1.96 3.12v0c.1.86-.57 1.63-1.45 1.63h-.38c-.86 0-1.6.6-1.76 1.44L14 10m8 3l-.82-.33c-.86-.34-1.82.2-1.98 1.11v0c-.11.7-.72 1.22-1.43 1.22H17M11 2l.33.82c.34.86-.2 1.82-1.11 1.98v0C9.52 4.9 9 5.52 9 6.23V7"/><path d="M11 13c1.93 1.93 2.83 4.17 2 5c-.83.83-3.07-.07-5-2c-1.93-1.93-2.83-4.17-2-5c.83-.83 3.07.07 5 2Z"/></g>`), g.Group(children),
	)
}

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 4h4v16H6zm8 0h4v16h-4z"/>`), g.Group(children),
	)
}

func PauseCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M10 15V9m4 6V9"/></g>`), g.Group(children),
	)
}

func PauseOctagon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 15V9m4 6V9M7.714 2h8.572L22 7.714v8.572L16.286 22H7.714L2 16.286V7.714L7.714 2z"/>`), g.Group(children),
	)
}

func PawPrint(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="4" r="2"/><circle cx="18" cy="8" r="2"/><circle cx="20" cy="16" r="2"/><path d="M9 10a5 5 0 0 1 5 5v3.5a3.5 3.5 0 0 1-6.84 1.045q-.64-2.065-2.7-2.705A3.5 3.5 0 0 1 5.5 10Z"/></g>`), g.Group(children),
	)
}

func PcCase(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="20" x="5" y="2" rx="2"/><path d="M15 14h.01M9 6h6m-6 4h6"/></g>`), g.Group(children),
	)
}

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5L2 22l1.5-5.5Z"/>`), g.Group(children),
	)
}

func PenLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20h9M16.5 3.5a2.12 2.12 0 0 1 3 3L7 19l-4 1l1-4Z"/>`), g.Group(children),
	)
}

func PenSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 4H4a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7"/><path d="M18.5 2.5a2.12 2.12 0 0 1 3 3L12 15l-4 1l1-4Z"/></g>`), g.Group(children),
	)
}

func PenTool(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 19l7-7l3 3l-7 7l-3-3z"/><path d="m18 13l-1.5-7.5L2 2l3.5 14.5L13 18l5-5zM2 2l7.586 7.586"/><circle cx="11" cy="11" r="2"/></g>`), g.Group(children),
	)
}

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3a2.85 2.83 0 1 1 4 4L7.5 20.5L2 22l1.5-5.5Zm-2 2l4 4"/>`), g.Group(children),
	)
}

func PencilLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20h9M16.5 3.5a2.12 2.12 0 0 1 3 3L7 19l-4 1l1-4ZM15 5l3 3"/>`), g.Group(children),
	)
}

func PencilRuler(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 5l4 4m-6-2L8.7 2.7a2.41 2.41 0 0 0-3.4 0L2.7 5.3a2.41 2.41 0 0 0 0 3.4L7 13m1-7l2-2M2 22l5.5-1.5L21.17 6.83a2.82 2.82 0 0 0-4-4L3.5 16.5Zm16-6l2-2"/><path d="m17 11l4.3 4.3c.94.94.94 2.46 0 3.4l-2.6 2.6c-.94.94-2.46.94-3.4 0L11 17"/></g>`), g.Group(children),
	)
}

func Pentagon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.5 8.7c-.7.5-1 1.4-.7 2.2l2.8 8.7c.3.8 1 1.4 1.9 1.4h9.1c.9 0 1.6-.6 1.9-1.4l2.8-8.7c.3-.8 0-1.7-.7-2.2l-7.4-5.3a2.1 2.1 0 0 0-2.4 0Z"/>`), g.Group(children),
	)
}

func Percent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 5L5 19"/><circle cx="6.5" cy="6.5" r="2.5"/><circle cx="17.5" cy="17.5" r="2.5"/></g>`), g.Group(children),
	)
}

func PercentCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m15 9l-6 6m0-6h.01M15 15h.01"/></g>`), g.Group(children),
	)
}

func PercentDiamond(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.7 10.3a2.41 2.41 0 0 0 0 3.41l7.59 7.59a2.41 2.41 0 0 0 3.41 0l7.59-7.59a2.41 2.41 0 0 0 0-3.41L13.7 2.71a2.41 2.41 0 0 0-3.41 0Zm6.5-1.1h.01m5.29.3l-5 5m5.2.3h.01"/>`), g.Group(children),
	)
}

func PercentSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m15 9l-6 6m0-6h.01M15 15h.01"/></g>`), g.Group(children),
	)
}

func PersonStanding(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="5" r="1"/><path d="m9 20l3-6l3 6M6 8l6 2l6-2m-6 2v4"/></g>`), g.Group(children),
	)
}

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 16.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`), g.Group(children),
	)
}

func PhoneCall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 16.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92zM14.05 2a9 9 0 0 1 8 7.94m-8-3.94A5 5 0 0 1 18 10"/>`), g.Group(children),
	)
}

func PhoneForwarded(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 2l4 4l-4 4m-4-4h8m0 10.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`), g.Group(children),
	)
}

func PhoneIncoming(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 2v6h6m0-6l-6 6m6 8.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`), g.Group(children),
	)
}

func PhoneMissed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m22 2l-6 6m0-6l6 6m0 8.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`), g.Group(children),
	)
}

func PhoneOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.68 13.31a16 16 0 0 0 3.41 2.6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7a2 2 0 0 1 1.72 2v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.42 19.42 0 0 1-3.33-2.67m-2.67-3.34a19.79 19.79 0 0 1-3.07-8.63A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91M22 2L2 22"/>`), g.Group(children),
	)
}

func PhoneOutgoing(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 8V2h-6m0 6l6-6m0 14.92v3a2 2 0 0 1-2.18 2a19.79 19.79 0 0 1-8.63-3.07a19.5 19.5 0 0 1-6-6a19.79 19.79 0 0 1-3.07-8.67A2 2 0 0 1 4.11 2h3a2 2 0 0 1 2 1.72a12.84 12.84 0 0 0 .7 2.81a2 2 0 0 1-.45 2.11L8.09 9.91a16 16 0 0 0 6 6l1.27-1.27a2 2 0 0 1 2.11-.45a12.84 12.84 0 0 0 2.81.7A2 2 0 0 1 22 16.92z"/>`), g.Group(children),
	)
}

func Pi(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 4v16M4 7c0-1.7 1.3-3 3-3h13"/><path d="M18 20c-1.7 0-3-1.3-3-3V4"/></g>`), g.Group(children),
	)
}

func PiSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 7h10m-7 0v10m6 0a2 2 0 0 1-2-2V7"/></g>`), g.Group(children),
	)
}

func PictureInPicture(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 4.5v5H3m-1-6l6 6m13 0v-3c0-1.16-.84-2-2-2h-7m-9 9v2c0 1.05.95 2 2 2h3"/><rect width="10" height="7" x="12" y="13.5" ry="2"/></g>`), g.Group(children),
	)
}

func PictureInPictureTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 9V6a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10c0 1.1.9 2 2 2h4"/><rect width="10" height="7" x="12" y="13" rx="2"/></g>`), g.Group(children),
	)
}

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21.21 15.89A10 10 0 1 1 8 2.83"/><path d="M22 12A10 10 0 0 0 12 2v10z"/></g>`), g.Group(children),
	)
}

func PiggyBank(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 5c-1.5 0-2.8 1.4-3 2c-3.5-1.5-11-.3-11 5c0 1.8 0 3 2 4.5V20h4v-2h3v2h4v-4c1-.5 1.7-1 2-2h2v-4h-2c0-1-.5-1.5-1-2h0V5zM2 9v1c0 1.1.9 2 2 2h1m11-1h0"/>`), g.Group(children),
	)
}

func Pilcrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 4v16m4-16v16m2-16H9.5a4.5 4.5 0 0 0 0 9H13"/>`), g.Group(children),
	)
}

func PilcrowSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 12H9.5a2.5 2.5 0 0 1 0-5H17m-5 0v10m4-10v10"/></g>`), g.Group(children),
	)
}

func Pill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.5 20.5l10-10a4.95 4.95 0 1 0-7-7l-10 10a4.95 4.95 0 1 0 7 7Zm-2-12l7 7"/>`), g.Group(children),
	)
}

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17v5m-7-5h14v-1.76a2 2 0 0 0-1.11-1.79l-1.78-.9A2 2 0 0 1 15 10.76V6h1a2 2 0 0 0 0-4H8a2 2 0 0 0 0 4h1v4.76a2 2 0 0 1-1.11 1.79l-1.78.9A2 2 0 0 0 5 15.24Z"/>`), g.Group(children),
	)
}

func PinOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20m-10-5v5M9 9v1.76a2 2 0 0 1-1.11 1.79l-1.78.9A2 2 0 0 0 5 15.24V17h12m-2-7.66V6h1a2 2 0 0 0 0-4H7.89"/>`), g.Group(children),
	)
}

func Pipette(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 22l1-1h3l9-9M3 21v-3l9-9"/><path d="m15 6l3.4-3.4a2.1 2.1 0 1 1 3 3L18 9l.4.4a2.1 2.1 0 1 1-3 3l-3.8-3.8a2.1 2.1 0 1 1 3-3l.4.4Z"/></g>`), g.Group(children),
	)
}

func Pizza(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 11h.01M11 15h.01M16 16h.01M2 16l20 6l-6-20A20 20 0 0 0 2 16"/><path d="M5.71 17.11a17.04 17.04 0 0 1 11.4-11.4"/></g>`), g.Group(children),
	)
}

func Plane(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.8 19.2L16 11l3.5-3.5C21 6 21.5 4 21 3c-1-.5-3 0-4.5 1.5L13 8L4.8 6.2c-.5-.1-.9.1-1.1.5l-.3.5c-.2.5-.1 1 .3 1.3L9 12l-2 3H4l-1 1l3 2l2 3l1-1v-3l3-2l3.5 5.3c.3.4.8.5 1.3.3l.5-.2c.4-.3.6-.7.5-1.2z"/>`), g.Group(children),
	)
}

func PlaneLanding(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 22h20M3.77 10.77L2 9l2-4.5l1.1.55c.55.28.9.84.9 1.45s.35 1.17.9 1.45L8 8.5l3-6l1.05.53a2 2 0 0 1 1.09 1.52l.72 5.4a2 2 0 0 0 1.09 1.52l4.4 2.2c.42.22.78.55 1.01.96l.6 1.03c.49.88-.06 1.98-1.06 2.1l-1.18.15c-.47.06-.95-.02-1.37-.24L4.29 11.15a2 2 0 0 1-.52-.38Z"/>`), g.Group(children),
	)
}

func PlaneTakeoff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 22h20M6.36 17.4L4 17l-2-4l1.1-.55a2 2 0 0 1 1.8 0l.17.1a2 2 0 0 0 1.8 0L8 12L5 6l.9-.45a2 2 0 0 1 2.09.2l4.02 3a2 2 0 0 0 2.1.2l4.19-2.06a2.41 2.41 0 0 1 1.73-.17L21 7a1.4 1.4 0 0 1 .87 1.99l-.38.76c-.23.46-.6.84-1.07 1.08L7.58 17.2a2 2 0 0 1-1.22.18Z"/>`), g.Group(children),
	)
}

func Play(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 3l14 9l-14 9V3z"/>`), g.Group(children),
	)
}

func PlayCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m10 8l6 4l-6 4V8z"/></g>`), g.Group(children),
	)
}

func PlaySquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m9 8l6 4l-6 4Z"/></g>`), g.Group(children),
	)
}

func Plug(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22v-5M9 8V2m6 6V2m3 6v5a4 4 0 0 1-4 4h-4a4 4 0 0 1-4-4V8Z"/>`), g.Group(children),
	)
}

func PlugTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 2v6m6-6v6m-3 9v5M5 8h14M6 11V8h12v3a6 6 0 1 1-12 0v0Z"/>`), g.Group(children),
	)
}

func PlugZap(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.3 20.3a2.4 2.4 0 0 0 3.4 0L12 18l-6-6l-2.3 2.3a2.4 2.4 0 0 0 0 3.4ZM2 22l3-3m2.5-5.5L10 11m.5 5.5L13 14m5-11l-4 4h6l-4 4"/>`), g.Group(children),
	)
}

func PlugZapTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 2l-2 2.5h3L12 7m-2 7v-3m4 3v-3m-3 8c-1.7 0-3-1.3-3-3v-2h8v2c0 1.7-1.3 3-3 3Zm1 3v-3"/>`), g.Group(children),
	)
}

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14m-7-7v14"/>`), g.Group(children),
	)
}

func PlusCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 12h8m-4-4v8"/></g>`), g.Group(children),
	)
}

func PlusSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M8 12h8m-4-4v8"/></g>`), g.Group(children),
	)
}

func Pocket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 3h16a2 2 0 0 1 2 2v6a10 10 0 0 1-10 10A10 10 0 0 1 2 11V5a2 2 0 0 1 2-2z"/><path d="m8 10l4 4l4-4"/></g>`), g.Group(children),
	)
}

func PocketKnife(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 2v1c0 1 2 1 2 2S3 6 3 7s2 1 2 2s-2 1-2 2s2 1 2 2m13-7h.01M6 18h.01m14.82-9.17a4 4 0 0 0-5.66-5.66l-12 12a4 4 0 1 0 5.66 5.66Z"/><path d="M18 11.66V22a4 4 0 0 0 4-4V6"/></g>`), g.Group(children),
	)
}

func Podcast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="11" r="1"/><path d="M11 17a1 1 0 0 1 2 0c0 .5-.34 3-.5 4.5a.5.5 0 0 1-1 0c-.16-1.5-.5-4-.5-4.5Zm-3-3a5 5 0 1 1 8 0"/><path d="M17 18.5a9 9 0 1 0-10 0"/></g>`), g.Group(children),
	)
}

func Pointer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 14a8 8 0 0 1-8 8m4-11v-1a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v0m0 0V9a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v1m0-.5V4a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v10"/><path d="M18 11a2 2 0 1 1 4 0v3a8 8 0 0 1-8 8h-2c-2.8 0-4.5-.86-5.99-2.34l-3.6-3.6a2 2 0 0 1 2.83-2.82L7 15"/></g>`), g.Group(children),
	)
}

func Popcorn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 8a2 2 0 0 0 0-4a2 2 0 0 0-4 0a2 2 0 0 0-4 0a2 2 0 0 0-4 0a2 2 0 0 0 0 4m4 14L9 8m5 14l1-14"/><path d="M20 8c.5 0 .9.4.8 1l-2.6 12c-.1.5-.7 1-1.2 1H7c-.6 0-1.1-.4-1.2-1L3.2 9c-.1-.6.3-1 .8-1Z"/></g>`), g.Group(children),
	)
}

func Popsicle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.6 14.4c.8-.8.8-2 0-2.8l-8.1-8.1a4.95 4.95 0 1 0-7.1 7.1l8.1 8.1c.9.7 2.1.7 2.9-.1ZM22 22l-5.5-5.5"/>`), g.Group(children),
	)
}

func PoundSterling(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 7c0-5.333-8-5.333-8 0m0 0v14m-4 0h12M6 13h10"/>`), g.Group(children),
	)
}

func Power(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v10m6.4-5.4a9 9 0 1 1-12.77.04"/>`), g.Group(children),
	)
}

func PowerCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M12 12V6M8 7.5A6.1 6.1 0 0 0 12 18a6 6 0 0 0 4-10.5"/></g>`), g.Group(children),
	)
}

func PowerOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.36 6.64A9 9 0 0 1 20.77 15M6.16 6.16a9 9 0 1 0 12.68 12.68M12 2v4M2 2l20 20"/>`), g.Group(children),
	)
}

func PowerSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 7v5M8 9a5.14 5.14 0 0 0 4 8a4.95 4.95 0 0 0 4-8"/></g>`), g.Group(children),
	)
}

func Presentation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 3h20m-1 0v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V3m4 18l5-5l5 5"/>`), g.Group(children),
	)
}

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 9V2h12v7M6 18H4a2 2 0 0 1-2-2v-5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2h-2"/><path d="M6 14h12v8H6z"/></g>`), g.Group(children),
	)
}

func Projector(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 7L3 5m6 1V3m4 4l2-2"/><circle cx="9" cy="13" r="3"/><path d="M11.83 12H20a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h2.17M16 16h2"/></g>`), g.Group(children),
	)
}

func Puzzle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.439 7.85c-.049.322.059.648.289.878l1.568 1.568c.47.47.706 1.087.706 1.704s-.235 1.233-.706 1.704l-1.611 1.611a.98.98 0 0 1-.837.276c-.47-.07-.802-.48-.968-.925a2.501 2.501 0 1 0-3.214 3.214c.446.166.855.497.925.968a.979.979 0 0 1-.276.837l-1.61 1.61a2.404 2.404 0 0 1-1.705.707a2.402 2.402 0 0 1-1.704-.706l-1.568-1.568a1.026 1.026 0 0 0-.877-.29c-.493.074-.84.504-1.02.968a2.5 2.5 0 1 1-3.237-3.237c.464-.18.894-.527.967-1.02a1.026 1.026 0 0 0-.289-.877l-1.568-1.568A2.402 2.402 0 0 1 1.998 12c0-.617.236-1.234.706-1.704L4.23 8.77c.24-.24.581-.353.917-.303c.515.077.877.528 1.073 1.01a2.5 2.5 0 1 0 3.259-3.259c-.482-.196-.933-.558-1.01-1.073c-.05-.336.062-.676.303-.917l1.525-1.525A2.402 2.402 0 0 1 12 1.998c.617 0 1.234.236 1.704.706l1.568 1.568c.23.23.556.338.877.29c.493-.074.84-.504 1.02-.968a2.5 2.5 0 1 1 3.237 3.237c-.464.18-.894.527-.967 1.02Z"/>`), g.Group(children),
	)
}

func Pyramid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2.5 16.88a1 1 0 0 1-.32-1.43l9-13.02a1 1 0 0 1 1.64 0l9 13.01a1 1 0 0 1-.32 1.44l-8.51 4.86a2 2 0 0 1-1.98 0ZM12 2v20"/>`), g.Group(children),
	)
}

func QrCode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="5" height="5" x="3" y="3" rx="1"/><rect width="5" height="5" x="16" y="3" rx="1"/><rect width="5" height="5" x="3" y="16" rx="1"/><path d="M21 16h-3a2 2 0 0 0-2 2v3m5 0v.01M12 7v3a2 2 0 0 1-2 2H7m-4 0h.01M12 3h.01M12 16v.01M16 12h1m4 0v.01M12 21v-1"/></g>`), g.Group(children),
	)
}

func Quote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21c3 0 7-1 7-8V5c0-1.25-.756-2.017-2-2H4c-1.25 0-2 .75-2 1.972V11c0 1.25.75 2 2 2c1 0 1 0 1 1v1c0 1-1 2-2 2s-1 .008-1 1.031V20c0 1 0 1 1 1zm12 0c3 0 7-1 7-8V5c0-1.25-.757-2.017-2-2h-4c-1.25 0-2 .75-2 1.972V11c0 1.25.75 2 2 2h.75c0 2.25.25 4-2.75 4v3c0 1 0 1 1 1z"/>`), g.Group(children),
	)
}

func Rabbit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 8.54V4a2 2 0 1 0-4 0v3"/><path d="M18 21h-8a4 4 0 0 1-4-4a7 7 0 0 1 7-7h.2L9.6 6.4a1.93 1.93 0 1 1 2.8-2.8L15.8 7h.2c3.3 0 6 2.7 6 6v1a2 2 0 0 1-2 2h-1c-1.7 0-3 1.3-3 3"/><path d="M7.61 12.53a3 3 0 1 0-1.6 4.3M13 16a3 3 0 0 1 2.24 5M18 12h.01"/></g>`), g.Group(children),
	)
}

func Radar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19.07 4.93A10 10 0 0 0 6.99 3.34M4 6h.01M2.29 9.62a10 10 0 1 0 19.02-1.27"/><path d="M16.24 7.76a6 6 0 1 0-8.01 8.91M12 18h.01m5.98-6.34a6 6 0 0 1-2.22 5.01"/><circle cx="12" cy="12" r="2"/><path d="m13.41 10.59l5.66-5.66"/></g>`), g.Group(children),
	)
}

func Radiation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12h0M7.5 4.2c-.3-.5-.9-.7-1.3-.4C3.9 5.5 2.3 8.1 2 11c-.1.5.4 1 1 1h5c0-1.5.8-2.8 2-3.4c-1.1-1.9-2-3.5-2.5-4.4zM21 12c.6 0 1-.4 1-1c-.3-2.9-1.8-5.5-4.1-7.1c-.4-.3-1.1-.2-1.3.3c-.6.9-1.5 2.5-2.6 4.3c1.2.7 2 2 2 3.5h5zM7.5 19.8c-.3.5-.1 1.1.4 1.3c2.6 1.2 5.6 1.2 8.2 0c.5-.2.7-.8.4-1.3c-.5-.9-1.4-2.5-2.5-4.3c-1.2.7-2.8.7-4 0c-1.1 1.8-2 3.4-2.5 4.3z"/>`), g.Group(children),
	)
}

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.9 19.1C1 15.2 1 8.8 4.9 4.9m2.9 11.3c-2.3-2.3-2.3-6.1 0-8.5"/><circle cx="12" cy="12" r="2"/><path d="M16.2 7.8c2.3 2.3 2.3 6.1 0 8.5m2.9-11.4C23 8.8 23 15.1 19.1 19"/></g>`), g.Group(children),
	)
}

func RadioReceiver(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 16v2m14-2v2"/><rect width="20" height="8" x="2" y="8" rx="2"/><path d="M18 12h0"/></g>`), g.Group(children),
	)
}

func RadioTower(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.9 16.1C1 12.2 1 5.8 4.9 1.9m2.9 2.8a6.14 6.14 0 0 0-.8 7.5"/><circle cx="12" cy="9" r="2"/><path d="M16.2 4.8c2 2 2.26 5.11.8 7.47M19.1 1.9a9.96 9.96 0 0 1 0 14.1m-9.6 2h5M8 22l4-11l4 11"/></g>`), g.Group(children),
	)
}

func Radius(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20.34 17.52a10 10 0 1 0-2.82 2.82"/><circle cx="19" cy="19" r="2"/><path d="m13.41 13.41l4.18 4.18"/><circle cx="12" cy="12" r="2"/></g>`), g.Group(children),
	)
}

func RailSymbol(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 15h14M5 9h14m-5 11l-5-5l6-6l-5-5"/>`), g.Group(children),
	)
}

func Rainbow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 17a10 10 0 0 0-20 0"/><path d="M6 17a6 6 0 0 1 12 0"/><path d="M10 17a2 2 0 0 1 4 0"/></g>`), g.Group(children),
	)
}

func Rat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 5c0-1.7-1.3-3-3-3s-3 1.3-3 3c0 .8.3 1.5.8 2H11c-3.9 0-7 3.1-7 7v0c0 2.2 1.8 4 4 4"/><path d="M16.8 3.9c.3-.3.6-.5 1-.7c1.5-.6 3.3.1 3.9 1.6c.6 1.5-.1 3.3-1.6 3.9l1.6 2.8c.2.3.2.7.2 1c-.2.8-.9 1.2-1.7 1.1c0 0-1.6-.3-2.7-.6H17c-1.7 0-3 1.3-3 3"/><path d="M13.2 18a3 3 0 0 0-2.2-5m2 9H4a2 2 0 0 1 0-4h12m0-9h.01"/></g>`), g.Group(children),
	)
}

func Ratio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="12" height="20" x="6" y="2" rx="2"/><rect width="20" height="12" x="2" y="6" rx="2"/></g>`), g.Group(children),
	)
}

func Receipt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 2v20l2-1l2 1l2-1l2 1l2-1l2 1l2-1l2 1V2l-2 1l-2-1l-2 1l-2-1l-2 1l-2-1l-2 1l-2-1Z"/><path d="M16 8h-6a2 2 0 1 0 0 4h4a2 2 0 1 1 0 4H8m4 1V7"/></g>`), g.Group(children),
	)
}

func RectangleHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="20" height="12" x="2" y="6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="2"/>`), g.Group(children),
	)
}

func RectangleVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="12" height="20" x="6" y="2" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="2"/>`), g.Group(children),
	)
}

func Recycle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 19H4.815a1.83 1.83 0 0 1-1.57-.881a1.785 1.785 0 0 1-.004-1.784L7.196 9.5M11 19h8.203a1.83 1.83 0 0 0 1.556-.89a1.784 1.784 0 0 0 0-1.775l-1.226-2.12"/><path d="m14 16l-3 3l3 3m-5.707-8.404L7.196 9.5L3.1 10.598m6.244-4.787l1.093-1.892A1.83 1.83 0 0 1 11.985 3a1.784 1.784 0 0 1 1.546.888l3.943 6.843"/><path d="m13.378 9.633l4.096 1.098l1.097-4.096"/></g>`), g.Group(children),
	)
}

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 7v6h-6"/><path d="M3 17a9 9 0 0 1 9-9a9 9 0 0 1 6 2.3l3 2.7"/></g>`), g.Group(children),
	)
}

func RedoDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="17" r="1"/><path d="M21 7v6h-6"/><path d="M3 17a9 9 0 0 1 9-9a9 9 0 0 1 6 2.3l3 2.7"/></g>`), g.Group(children),
	)
}

func RedoTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 14l5-5l-5-5"/><path d="M20 9H9.5A5.5 5.5 0 0 0 4 14.5v0A5.5 5.5 0 0 0 9.5 20H13"/></g>`), g.Group(children),
	)
}

func RefreshCcw(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 0 0-9-9a9.75 9.75 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5m-5 4a9 9 0 0 0 9 9a9.75 9.75 0 0 0 6.74-2.74L21 16"/><path d="M16 16h5v5"/></g>`), g.Group(children),
	)
}

func RefreshCcwDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 2v6h6"/><path d="M21 12A9 9 0 0 0 6 5.3L3 8m18 14v-6h-6"/><path d="M3 12a9 9 0 0 0 15 6.7l3-2.7"/><circle cx="12" cy="12" r="1"/></g>`), g.Group(children),
	)
}

func RefreshCw(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 0 1 9-9a9.75 9.75 0 0 1 6.74 2.74L21 8"/><path d="M21 3v5h-5m5 4a9 9 0 0 1-9 9a9.75 9.75 0 0 1-6.74-2.74L3 16"/><path d="M8 16H3v5"/></g>`), g.Group(children),
	)
}

func RefreshCwOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m21 8l-2.26-2.26A9.75 9.75 0 0 0 12 3c-1 0-1.97.16-2.87.47M8 16H3v5m0-9c0-2.49 1-4.74 2.64-6.36"/><path d="m3 16l2.26 2.26A9.75 9.75 0 0 0 12 21c2.49 0 4.74-1 6.36-2.64M21 12c0 1-.16 1.97-.47 2.87M21 3v5h-5m6 14L2 2"/></g>`), g.Group(children),
	)
}

func Refrigerator(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 6a4 4 0 0 1 4-4h6a4 4 0 0 1 4 4v14a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V6Zm0 4h14m-4-3v6"/>`), g.Group(children),
	)
}

func Regex(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 3v10m-4.33-7.5l8.66 5m-8.66 0l8.66-5M9 17a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2z"/>`), g.Group(children),
	)
}

func RemoveFormatting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7V4h16v3M5 20h6m2-16L8 20m7-5l5 5m0-5l-5 5"/>`), g.Group(children),
	)
}

func Repeat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 2l4 4l-4 4"/><path d="M3 11v-1a4 4 0 0 1 4-4h14M7 22l-4-4l4-4"/><path d="M21 13v1a4 4 0 0 1-4 4H3"/></g>`), g.Group(children),
	)
}

func RepeatOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m17 2l4 4l-4 4"/><path d="M3 11v-1a4 4 0 0 1 4-4h14M7 22l-4-4l4-4"/><path d="M21 13v1a4 4 0 0 1-4 4H3m8-8h1v4"/></g>`), g.Group(children),
	)
}

func RepeatTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 9l3-3l3 3"/><path d="M13 18H7a2 2 0 0 1-2-2V6m17 9l-3 3l-3-3"/><path d="M11 6h6a2 2 0 0 1 2 2v10"/></g>`), g.Group(children),
	)
}

func Replace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 4c0-1.1.9-2 2-2m4 0c1.1 0 2 .9 2 2m0 4c0 1.1-.9 2-2 2m-4 0c-1.1 0-2-.9-2-2M3 7l3 3l3-3"/><path d="M6 10V5c0-1.7 1.3-3 3-3h1"/><rect width="8" height="8" x="2" y="14" rx="2"/></g>`), g.Group(children),
	)
}

func ReplaceAll(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 4c0-1.1.9-2 2-2m4 0c1.1 0 2 .9 2 2m0 4c0 1.1-.9 2-2 2m-4 0c-1.1 0-2-.9-2-2M3 7l3 3l3-3"/><path d="M6 10V5c0-1.7 1.3-3 3-3h1"/><rect width="8" height="8" x="2" y="14" rx="2"/><path d="M14 14c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2m6-8c1.1 0 2 .9 2 2v4c0 1.1-.9 2-2 2"/></g>`), g.Group(children),
	)
}

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 17l-5-5l5-5"/><path d="M20 18v-2a4 4 0 0 0-4-4H4"/></g>`), g.Group(children),
	)
}

func ReplyAll(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m7 17l-5-5l5-5m5 10l-5-5l5-5"/><path d="M22 18v-2a4 4 0 0 0-4-4H7"/></g>`), g.Group(children),
	)
}

func Rewind(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 19l-9-7l9-7v14zm11 0l-9-7l9-7v14z"/>`), g.Group(children),
	)
}

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.5 16.5c-1.5 1.26-2 5-2 5s3.74-.5 5-2c.71-.84.7-2.13-.09-2.91a2.18 2.18 0 0 0-2.91-.09zM12 15l-3-3a22 22 0 0 1 2-3.95A12.88 12.88 0 0 1 22 2c0 2.72-.78 7.5-6 11a22.35 22.35 0 0 1-4 2z"/><path d="M9 12H4s.55-3.03 2-4c1.62-1.08 5 0 5 0m1 7v5s3.03-.55 4-2c1.08-1.62 0-5 0-5"/></g>`), g.Group(children),
	)
}

func RockingChair(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3.5 2l3 10.5H18m-8.5 0l-4 7.5m9.5-7.5l3.5 7.5M2.75 18a13 13 0 0 0 18.5 0"/>`), g.Group(children),
	)
}

func RollerCoaster(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 19V5m4 14V6.8M14 19v-7.8M18 5v4m0 10v-6m4 6V9M2 19V9a4 4 0 0 1 4-4c2 0 4 1.33 6 4s4 4 6 4a4 4 0 1 0-3-6.65"/>`), g.Group(children),
	)
}

func RotateCcw(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 1 0 9-9a9.75 9.75 0 0 0-6.74 2.74L3 8"/><path d="M3 3v5h5"/></g>`), g.Group(children),
	)
}

func RotateCw(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12a9 9 0 1 1-9-9c2.52 0 4.93 1 6.74 2.74L21 8"/><path d="M21 3v5h-5"/></g>`), g.Group(children),
	)
}

func RotateThreeD(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16.466 7.5C15.643 4.237 13.952 2 12 2C9.239 2 7 6.477 7 12s2.239 10 5 10c.342 0 .677-.069 1-.2m2.194-8.093l3.814 1.86l-1.86 3.814"/><path d="M19 15.57c-1.804.885-4.274 1.43-7 1.43c-5.523 0-10-2.239-10-5s4.477-5 10-5c4.838 0 8.873 1.718 9.8 4"/></g>`), g.Group(children),
	)
}

func Route(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="19" r="3"/><path d="M9 19h8.5a3.5 3.5 0 0 0 0-7h-11a3.5 3.5 0 0 1 0-7H15"/><circle cx="18" cy="5" r="3"/></g>`), g.Group(children),
	)
}

func RouteOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="19" r="3"/><path d="M9 19h8.5c.4 0 .9-.1 1.3-.2M5.2 5.2A3.5 3.53 0 0 0 6.5 12H12M2 2l20 20m-1-6.7a3.5 3.5 0 0 0-3.3-3.3M15 5h-4.3"/><circle cx="18" cy="5" r="3"/></g>`), g.Group(children),
	)
}

func Router(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="8" x="2" y="14" rx="2"/><path d="M6.01 18H6m4.01 0H10m5-8v4m2.84-6.83a4 4 0 0 0-5.66 0m8.48-2.83a8 8 0 0 0-11.31 0"/></g>`), g.Group(children),
	)
}

func Rows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 12h18"/></g>`), g.Group(children),
	)
}

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 11a9 9 0 0 1 9 9M4 4a16 16 0 0 1 16 16"/><circle cx="5" cy="19" r="1"/></g>`), g.Group(children),
	)
}

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21.3 15.3a2.4 2.4 0 0 1 0 3.4l-2.6 2.6a2.4 2.4 0 0 1-3.4 0L2.7 8.7a2.41 2.41 0 0 1 0-3.4l2.6-2.6a2.41 2.41 0 0 1 3.4 0Zm-6.8-2.8l2-2m-5-1l2-2m-5-1l2-2m7 11l2-2"/>`), g.Group(children),
	)
}

func RussianRuble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 11h8a4 4 0 0 0 0-8H9v18m-3-6h8"/>`), g.Group(children),
	)
}

func Sailboat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 18H2a4 4 0 0 0 4 4h12a4 4 0 0 0 4-4Zm-1-4L10 2L3 14h18ZM10 2v16"/>`), g.Group(children),
	)
}

func Salad(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 21h10m-5 0a9 9 0 0 0 9-9H3a9 9 0 0 0 9 9Z"/><path d="M11.38 12a2.4 2.4 0 0 1-.4-4.77a2.4 2.4 0 0 1 3.2-2.77a2.4 2.4 0 0 1 3.47-.63a2.4 2.4 0 0 1 3.37 3.37a2.4 2.4 0 0 1-1.1 3.7a2.51 2.51 0 0 1 .03 1.1M13 12l4-4"/><path d="M10.9 7.25A3.99 3.99 0 0 0 4 10c0 .73.2 1.41.54 2"/></g>`), g.Group(children),
	)
}

func Sandwich(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 11v3a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-3m-9 8H4a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-3.83M3 11l7.77-6.04a2 2 0 0 1 2.46 0L21 11H3Z"/><path d="M12.97 19.77L7 15h12.5l-3.75 4.5a2 2 0 0 1-2.78.27Z"/></g>`), g.Group(children),
	)
}

func Satellite(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M13 7L9 3L5 7l4 4m8 0l4 4l-4 4l-4-4"/><path d="m8 12l4 4l6-6l-4-4Zm8-4l3-3M9 21a6 6 0 0 0-6-6"/></g>`), g.Group(children),
	)
}

func SatelliteDish(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 10a7.31 7.31 0 0 0 10 10Zm5 5l3-3m5 1a6 6 0 0 0-6-6m10 6A10 10 0 0 0 11 3"/>`), g.Group(children),
	)
}

func Save(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h11l5 5v11a2 2 0 0 1-2 2z"/><path d="M17 21v-8H7v8M7 3v5h8"/></g>`), g.Group(children),
	)
}

func SaveAll(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 4a2 2 0 0 1 2-2h10l4 4v10.2a2 2 0 0 1-2 1.8H8a2 2 0 0 1-2-2Z"/><path d="M10 2v4h6m2 12v-7h-8v7"/><path d="M18 22H4a2 2 0 0 1-2-2V6"/></g>`), g.Group(children),
	)
}

func Scale(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 16l3-8l3 8c-.87.65-1.92 1-3 1s-2.13-.35-3-1ZM2 16l3-8l3 8c-.87.65-1.92 1-3 1s-2.13-.35-3-1Zm5 5h10M12 3v18M3 7h2c2 0 5-1 7-2c2 1 5 2 7 2h2"/>`), g.Group(children),
	)
}

func ScaleThreeD(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="19" cy="19" r="2"/><circle cx="5" cy="5" r="2"/><path d="M5 7v12h12M5 19l6-6"/></g>`), g.Group(children),
	)
}

func Scaling(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 3L9 15m3-12H3v18h18v-9m-5-9h5v5"/><path d="M14 15H9v-5"/></g>`), g.Group(children),
	)
}

func Scan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2"/>`), g.Group(children),
	)
}

func ScanBarcode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2M8 7v10m4-10v10m5-10v10"/>`), g.Group(children),
	)
}

func ScanEye(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2"/><circle cx="12" cy="12" r="1"/><path d="M5 12s2.5-5 7-5s7 5 7 5s-2.5 5-7 5s-7-5-7-5"/></g>`), g.Group(children),
	)
}

func ScanFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2m5-3s1.5 2 4 2s4-2 4-2M9 9h.01M15 9h.01"/>`), g.Group(children),
	)
}

func ScanLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2m4-5h10"/>`), g.Group(children),
	)
}

func ScanSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2"/><circle cx="12" cy="12" r="3"/><path d="m16 16l-1.9-1.9"/></g>`), g.Group(children),
	)
}

func ScanText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7V5a2 2 0 0 1 2-2h2m10 0h2a2 2 0 0 1 2 2v2m0 10v2a2 2 0 0 1-2 2h-2M7 21H5a2 2 0 0 1-2-2v-2m4-9h8m-8 4h10M7 16h6"/>`), g.Group(children),
	)
}

func ScatterChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7.5" cy="7.5" r=".5"/><circle cx="18.5" cy="5.5" r=".5"/><circle cx="11.5" cy="11.5" r=".5"/><circle cx="7.5" cy="16.5" r=".5"/><circle cx="17.5" cy="14.5" r=".5"/><path d="M3 3v18h18"/></g>`), g.Group(children),
	)
}

func School(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m4 6l8-4l8 4m-2 4l4 2v8a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2v-8l4-2"/><path d="M14 22v-4a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v4m8-17v17M6 5v17"/><circle cx="12" cy="9" r="2"/></g>`), g.Group(children),
	)
}

func SchoolTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="10" r="1"/><path d="M22 20V8h-4l-6-4l-6 4H2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2ZM6 17v.01M6 13v.01M18 17v.01M18 13v.01"/><path d="M14 22v-5a2 2 0 0 0-2-2v0a2 2 0 0 0-2 2v5"/></g>`), g.Group(children),
	)
}

func Scissors(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="6" r="3"/><path d="M8.12 8.12L12 12m8-8L8.12 15.88"/><circle cx="6" cy="18" r="3"/><path d="M14.8 14.8L20 20"/></g>`), g.Group(children),
	)
}

func ScissorsLineDashed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.42 9.42L8 12"/><circle cx="4" cy="8" r="2"/><path d="m14 6l-8.58 8.58"/><circle cx="4" cy="16" r="2"/><path d="M10.8 14.8L14 18m2-6h-2m8 0h-2"/></g>`), g.Group(children),
	)
}

func ScissorsSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="20" x="2" y="2" rx="2"/><circle cx="8" cy="8" r="2"/><path d="M9.414 9.414L12 12m2.8 2.8L18 18"/><circle cx="8" cy="16" r="2"/><path d="m18 6l-8.586 8.586"/></g>`), g.Group(children),
	)
}

func ScissorsSquareDashedBottom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 22a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2m-10 0H8m8 0h-2"/><circle cx="8" cy="8" r="2"/><path d="M9.414 9.414L12 12m2.8 2.8L18 18"/><circle cx="8" cy="16" r="2"/><path d="m18 6l-8.586 8.586"/></g>`), g.Group(children),
	)
}

func ScreenShare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 3H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-3M8 21h8m-4-4v4m5-13l5-5m-5 0h5v5"/>`), g.Group(children),
	)
}

func ScreenShareOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 3H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-3M8 21h8m-4-4v4M22 3l-5 5m0-5l5 5"/>`), g.Group(children),
	)
}

func Scroll(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 21h12a2 2 0 0 0 2-2v-2H10v2a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v3h4"/><path d="M19 17V5a2 2 0 0 0-2-2H4"/></g>`), g.Group(children),
	)
}

func ScrollText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 21h12a2 2 0 0 0 2-2v-2H10v2a2 2 0 1 1-4 0V5a2 2 0 1 0-4 0v3h4"/><path d="M19 17V5a2 2 0 0 0-2-2H4m11 5h-5m5 4h-5"/></g>`), g.Group(children),
	)
}

func Search(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3"/></g>`), g.Group(children),
	)
}

func SearchCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 11l2 2l4-4"/><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3"/></g>`), g.Group(children),
	)
}

func SearchCode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 9l-2 2l2 2m4 0l2-2l-2-2"/><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3"/></g>`), g.Group(children),
	)
}

func SearchLarge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path clip-rule="evenodd" d="M18.874 19.581a6 6 0 1 1 .707-.707l4.273 4.272l-.708.708zM20 15a5 5 0 1 1-10 0a5 5 0 0 1 10 0z" fill="currentColor" fill-rule="evenodd"/></g>`), g.Group(children),
	)
}

func SearchSlash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13.5 8.5l-5 5"/><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3"/></g>`), g.Group(children),
	)
}

func SearchX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m13.5 8.5l-5 5m0-5l5 5"/><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.3-4.3"/></g>`), g.Group(children),
	)
}

func Send(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m22 2l-7 20l-4-9l-9-4Zm0 0L11 13"/>`), g.Group(children),
	)
}

func SendHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l3 9l-3 9l19-9Zm3 9h16"/>`), g.Group(children),
	)
}

func SendToBack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="8" x="14" y="14" rx="2"/><rect width="8" height="8" x="2" y="2" rx="2"/><path d="M7 14v1a2 2 0 0 0 2 2h1m4-10h1a2 2 0 0 1 2 2v1"/></g>`), g.Group(children),
	)
}

func SeparatorHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12h18M8 8l4-4l4 4m0 8l-4 4l-4-4"/>`), g.Group(children),
	)
}

func SeparatorVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v18M8 8l-4 4l4 4m8 0l4-4l-4-4"/>`), g.Group(children),
	)
}

func Server(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="8" x="2" y="2" rx="2" ry="2"/><rect width="20" height="8" x="2" y="14" rx="2" ry="2"/><path d="M6 6h.01M6 18h.01"/></g>`), g.Group(children),
	)
}

func ServerCog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="3"/><path d="M4.5 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-.5m-15 4H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-.5M6 6h.01M6 18h.01m9.69-4.6l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.7.9l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4m-2.3-2.1l-.3.9"/></g>`), g.Group(children),
	)
}

func ServerCrash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 10H4a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-2M6 14H4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-2M6 6h.01M6 18h.01"/><path d="m13 6l-4 6h6l-4 6"/></g>`), g.Group(children),
	)
}

func ServerOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 2h13a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-5m-5 0L2.5 2.5C2 2 2 2.5 2 5v3a2 2 0 0 0 2 2h6zm12 7v-1a2 2 0 0 0-2-2h-1M4 14a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h16.5l1-.5l.5.5l-8-8H4zm2 4h.01M2 2l20 20"/>`), g.Group(children),
	)
}

func Settings(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.22 2h-.44a2 2 0 0 0-2 2v.18a2 2 0 0 1-1 1.73l-.43.25a2 2 0 0 1-2 0l-.15-.08a2 2 0 0 0-2.73.73l-.22.38a2 2 0 0 0 .73 2.73l.15.1a2 2 0 0 1 1 1.72v.51a2 2 0 0 1-1 1.74l-.15.09a2 2 0 0 0-.73 2.73l.22.38a2 2 0 0 0 2.73.73l.15-.08a2 2 0 0 1 2 0l.43.25a2 2 0 0 1 1 1.73V20a2 2 0 0 0 2 2h.44a2 2 0 0 0 2-2v-.18a2 2 0 0 1 1-1.73l.43-.25a2 2 0 0 1 2 0l.15.08a2 2 0 0 0 2.73-.73l.22-.39a2 2 0 0 0-.73-2.73l-.15-.08a2 2 0 0 1-1-1.74v-.5a2 2 0 0 1 1-1.74l.15-.09a2 2 0 0 0 .73-2.73l-.22-.38a2 2 0 0 0-2.73-.73l-.15.08a2 2 0 0 1-2 0l-.43-.25a2 2 0 0 1-1-1.73V4a2 2 0 0 0-2-2z"/><circle cx="12" cy="12" r="3"/></g>`), g.Group(children),
	)
}

func SettingsTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 7h-9m3 10H5"/><circle cx="17" cy="17" r="3"/><circle cx="7" cy="7" r="3"/></g>`), g.Group(children),
	)
}

func Shapes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8.3 10a.7.7 0 0 1-.626-1.079L11.4 3a.7.7 0 0 1 1.198-.043L16.3 8.9a.7.7 0 0 1-.572 1.1Z"/><rect width="7" height="7" x="3" y="14" rx="1"/><circle cx="17.5" cy="17.5" r="3.5"/></g>`), g.Group(children),
	)
}

func Share(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8m-4-6l-4-4l-4 4m4-4v13"/>`), g.Group(children),
	)
}

func ShareTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="5" r="3"/><circle cx="6" cy="12" r="3"/><circle cx="18" cy="19" r="3"/><path d="m8.59 13.51l6.83 3.98m-.01-10.98l-6.82 3.98"/></g>`), g.Group(children),
	)
}

func Sheet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M3 9h18M3 15h18M9 9v12m6-12v12"/></g>`), g.Group(children),
	)
}

func Shell(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 11a2 2 0 1 1-4 0a4 4 0 0 1 8 0a6 6 0 0 1-12 0a8 8 0 0 1 16 0a10 10 0 1 1-20 0a11.93 11.93 0 0 1 2.42-7.22a2 2 0 1 1 3.16 2.44"/>`), g.Group(children),
	)
}

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10"/>`), g.Group(children),
	)
}

func ShieldAlert(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10m0-14v4m0 4h.01"/>`), g.Group(children),
	)
}

func ShieldBan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10M4 5l14 12"/>`), g.Group(children),
	)
}

func ShieldCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10"/><path d="m9 12l2 2l4-4"/></g>`), g.Group(children),
	)
}

func ShieldClose(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10zM9.5 9l5 5m0-5l-5 5"/>`), g.Group(children),
	)
}

func ShieldEllipsis(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10M8 11h.01M12 11h.01M16 11h.01"/>`), g.Group(children),
	)
}

func ShieldHalf(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10m0 0V2"/>`), g.Group(children),
	)
}

func ShieldMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10M8 11h8"/>`), g.Group(children),
	)
}

func ShieldOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.7 14a6.9 6.9 0 0 0 .3-2V5l-8-3l-3.2 1.2M2 2l20 20M4.7 4.7L4 5v7c0 6 8 10 8 10a20.3 20.3 0 0 0 5.62-4.38"/>`), g.Group(children),
	)
}

func ShieldPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10M8 11h8m-4 4V7"/>`), g.Group(children),
	)
}

func ShieldQuestion(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10"/><path d="M9.1 9a3 3 0 0 1 5.82 1c0 2-3 3-3 3m.08 4h.01"/></g>`), g.Group(children),
	)
}

func ShieldX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22s8-4 8-10V5l-8-3l-8 3v7c0 6 8 10 8 10m2.5-13l-5 5m0-5l5 5"/>`), g.Group(children),
	)
}

func Ship(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 21c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1"/><path d="M19.38 20A11.6 11.6 0 0 0 21 14l-9-4l-9 4c0 2.9.94 5.34 2.81 7.76"/><path d="M19 13V7a2 2 0 0 0-2-2H7a2 2 0 0 0-2 2v6m7-3v4m0-12v3"/></g>`), g.Group(children),
	)
}

func ShipWheel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="8"/><path d="M12 2v7.5M19 5l-5.23 5.23M22 12h-7.5m4.5 7l-5.23-5.23M12 14.5V22m-1.77-8.23L5 19m4.5-7H2m8.23-1.77L5 5"/><circle cx="12" cy="12" r="2.5"/></g>`), g.Group(children),
	)
}

func Shirt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.38 3.46L16 2a4 4 0 0 1-8 0L3.62 3.46a2 2 0 0 0-1.34 2.23l.58 3.47a1 1 0 0 0 .99.84H6v10c0 1.1.9 2 2 2h8a2 2 0 0 0 2-2V10h2.15a1 1 0 0 0 .99-.84l.58-3.47a2 2 0 0 0-1.34-2.23z"/>`), g.Group(children),
	)
}

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 2L3 6v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V6l-3-4ZM3 6h18"/><path d="M16 10a4 4 0 0 1-8 0"/></g>`), g.Group(children),
	)
}

func ShoppingBasket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 11l4-7m10 7l-4-7M2 11h20M3.5 11l1.6 7.4a2 2 0 0 0 2 1.6h9.8c.9 0 1.8-.7 2-1.6l1.7-7.4M9 11l1 9m-5.5-4.5h15M15 11l-1 9"/>`), g.Group(children),
	)
}

func ShoppingCart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="8" cy="21" r="1"/><circle cx="19" cy="21" r="1"/><path d="M2.05 2.05h2l2.66 12.42a2 2 0 0 0 2 1.58h9.78a2 2 0 0 0 1.95-1.57l1.65-7.43H5.12"/></g>`), g.Group(children),
	)
}

func Shovel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 22v-5l5-5l5 5l-5 5zm7.5-7.5L16 8m1-6l5 5l-.5.5a3.53 3.53 0 0 1-5 0s0 0 0 0a3.53 3.53 0 0 1 0-5L17 2"/>`), g.Group(children),
	)
}

func ShowerHead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 4l2.5 2.5m7 0a4.95 4.95 0 0 0-7 7M15 5L5 15m9 2v.01M10 16v.01M13 13v.01M16 10v.01M11 20v.01M17 14v.01M20 11v.01"/>`), g.Group(children),
	)
}

func Shrink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 15l6 6m-6-6v4.8m0-4.8h4.8M9 19.8V15m0 0H4.2M9 15l-6 6M15 4.2V9m0 0h4.8M15 9l6-6M9 4.2V9m0 0H4.2M9 9L3 3"/>`), g.Group(children),
	)
}

func Shrub(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 22v-7l-2-2"/><path d="M17 8v.8A6 6 0 0 1 13.8 20v0H10v0A6.5 6.5 0 0 1 7 8h0a5 5 0 0 1 10 0Zm-3 6l-2 2"/></g>`), g.Group(children),
	)
}

func Shuffle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 18h1.4c1.3 0 2.5-.6 3.3-1.7l6.1-8.6c.7-1.1 2-1.7 3.3-1.7H22"/><path d="m18 2l4 4l-4 4M2 6h1.9c1.5 0 2.9.9 3.6 2.2M22 18h-5.9c-1.3 0-2.6-.7-3.3-1.8l-.5-.8"/><path d="m18 14l4 4l-4 4"/></g>`), g.Group(children),
	)
}

func Sigma(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 7V4H6l6 8l-6 8h12v-3"/>`), g.Group(children),
	)
}

func SigmaSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M16 8.9V7H8l4 5l-4 5h8v-1.9"/></g>`), g.Group(children),
	)
}

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h.01M7 20v-4m5 4v-8m5 8V8m5-4v16"/>`), g.Group(children),
	)
}

func SignalHigh(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h.01M7 20v-4m5 4v-8m5 8V8"/>`), g.Group(children),
	)
}

func SignalLow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h.01M7 20v-4"/>`), g.Group(children),
	)
}

func SignalMedium(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h.01M7 20v-4m5 4v-8"/>`), g.Group(children),
	)
}

func SignalZero(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 20h.01"/>`), g.Group(children),
	)
}

func Signpost(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3v3m6.5 7h-13L2 9.5L5.5 6h13L22 9.5ZM12 13v8"/>`), g.Group(children),
	)
}

func SignpostBig(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9H4L2 7l2-2h6m4 0h6l2 2l-2 2h-6m-4 13V4a2 2 0 1 1 4 0v18m-6 0h8"/>`), g.Group(children),
	)
}

func Siren(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 12a5 5 0 0 1 5-5v0a5 5 0 0 1 5 5v6H7v-6Zm-2 8a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v2H5v-2Zm16-8h1m-3.5-7.5L18 5M2 12h1m9-10v1M4.929 4.929l.707.707M12 12v6"/>`), g.Group(children),
	)
}

func SkipBack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 20L9 12l10-8v16zM5 19V5"/>`), g.Group(children),
	)
}

func SkipForward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 4l10 8l-10 8V4zm14 1v14"/>`), g.Group(children),
	)
}

func Skull(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="12" r="1"/><circle cx="15" cy="12" r="1"/><path d="M8 20v2h8v-2m-3.5-3l-.5-1l-.5 1h1z"/><path d="M16 20a2 2 0 0 0 1.56-3.25a8 8 0 1 0-11.12 0A2 2 0 0 0 8 20"/></g>`), g.Group(children),
	)
}

func Slack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="3" height="8" x="13" y="2" rx="1.5"/><path d="M19 8.5V10h1.5A1.5 1.5 0 1 0 19 8.5"/><rect width="3" height="8" x="8" y="14" rx="1.5"/><path d="M5 15.5V14H3.5A1.5 1.5 0 1 0 5 15.5"/><rect width="8" height="3" x="14" y="13" rx="1.5"/><path d="M15.5 19H14v1.5a1.5 1.5 0 1 0 1.5-1.5"/><rect width="8" height="3" x="2" y="8" rx="1.5"/><path d="M8.5 5H10V3.5A1.5 1.5 0 1 0 8.5 5"/></g>`), g.Group(children),
	)
}

func Slash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 2L2 22"/>`), g.Group(children),
	)
}

func Slice(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m8 14l-6 6h9v-3"/><path d="M18.37 3.63L8 14l3 3L21.37 6.63a2.12 2.12 0 1 0-3-3Z"/></g>`), g.Group(children),
	)
}

func Sliders(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 21v-7m0-4V3m8 18v-9m0-4V3m8 18v-5m0-4V3M2 14h4m4-6h4m4 8h4"/>`), g.Group(children),
	)
}

func SlidersHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 4h-7m-4 0H3m18 8h-9m-4 0H3m18 8h-5m-4 0H3M14 2v4m-6 4v4m8 4v4"/>`), g.Group(children),
	)
}

func Smartphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="20" x="5" y="2" rx="2" ry="2"/><path d="M12 18h.01"/></g>`), g.Group(children),
	)
}

func SmartphoneCharging(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="14" height="20" x="5" y="2" rx="2" ry="2"/><path d="M12.667 8L10 12h4l-2.667 4"/></g>`), g.Group(children),
	)
}

func SmartphoneNfc(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="7" height="12" x="2" y="6" rx="1"/><path d="M13 8.32a7.43 7.43 0 0 1 0 7.36m3.46-9.47a11.76 11.76 0 0 1 0 11.58M19.91 4.1a15.91 15.91 0 0 1 .01 15.8"/></g>`), g.Group(children),
	)
}

func Smile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M8 14s1.5 2 4 2s4-2 4-2M9 9h.01M15 9h.01"/></g>`), g.Group(children),
	)
}

func SmilePlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 11v1a10 10 0 1 1-9-10"/><path d="M8 14s1.5 2 4 2s4-2 4-2M9 9h.01M15 9h.01M16 5h6m-3-3v6"/></g>`), g.Group(children),
	)
}

func Snail(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 13a6 6 0 1 0 12 0a4 4 0 1 0-8 0a2 2 0 0 0 4 0"/><circle cx="10" cy="13" r="8"/><path d="M2 21h12c4.4 0 8-3.6 8-8V7a2 2 0 1 0-4 0v6m0-10l1.1 2.2M22 3l-1.1 2.2"/></g>`), g.Group(children),
	)
}

func Snowflake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h20M12 2v20m8-6l-4-4l4-4M4 8l4 4l-4 4M16 4l-4 4l-4-4m0 16l4-4l4 4"/>`), g.Group(children),
	)
}

func Sofa(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M20 9V6a2 2 0 0 0-2-2H6a2 2 0 0 0-2 2v3"/><path d="M2 11v5a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-5a2 2 0 0 0-4 0v2H6v-2a2 2 0 0 0-4 0Zm2 7v2m16-2v2M12 4v9"/></g>`), g.Group(children),
	)
}

func SortAsc(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 11h4m-4 4h7m-7 4h10M9 7L6 4L3 7m3-1v14"/>`), g.Group(children),
	)
}

func SortDesc(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5h10M11 9h7m-7 4h4M3 17l3 3l3-3m-3 1V4"/>`), g.Group(children),
	)
}

func Soup(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 0 0 9-9H3a9 9 0 0 0 9 9Zm-5 0h10m2.5-9L22 6m-5.75-3c.27.1.8.53.75 1.36c-.06.83-.93 1.2-1 2.02c-.05.78.34 1.24.73 1.62m-5.48-5c.27.1.8.53.74 1.36c-.05.83-.93 1.2-.98 2.02c-.06.78.33 1.24.72 1.62M6.25 3c.27.1.8.53.75 1.36c-.06.83-.93 1.2-1 2.02c-.05.78.34 1.24.74 1.62"/>`), g.Group(children),
	)
}

func Space(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 17v1c0 .5-.5 1-1 1H3c-.5 0-1-.5-1-1v-1"/>`), g.Group(children),
	)
}

func Spade(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9c-1.5 1.5-3 3.2-3 5.5A5.5 5.5 0 0 0 7.5 20c1.8 0 3-.5 4.5-2c1.5 1.5 2.7 2 4.5 2a5.5 5.5 0 0 0 5.5-5.5c0-2.3-1.5-4-3-5.5l-7-7l-7 7Zm7 9v4"/>`), g.Group(children),
	)
}

func Sparkle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 3l-1.9 5.8a2 2 0 0 1-1.287 1.288L3 12l5.8 1.9a2 2 0 0 1 1.288 1.287L12 21l1.9-5.8a2 2 0 0 1 1.287-1.288L21 12l-5.8-1.9a2 2 0 0 1-1.288-1.287Z"/>`), g.Group(children),
	)
}

func Sparkles(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 3l-1.912 5.813a2 2 0 0 1-1.275 1.275L3 12l5.813 1.912a2 2 0 0 1 1.275 1.275L12 21l1.912-5.813a2 2 0 0 1 1.275-1.275L21 12l-5.813-1.912a2 2 0 0 1-1.275-1.275L12 3ZM5 3v4m14 10v4M3 5h4m10 14h4"/>`), g.Group(children),
	)
}

func Speaker(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="20" x="4" y="2" rx="2" ry="2"/><circle cx="12" cy="14" r="4"/><path d="M12 6h.01"/></g>`), g.Group(children),
	)
}

func Speech(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.8 20v-4.1l1.9.2a2.3 2.3 0 0 0 2.164-2.1V8.3A5.37 5.37 0 0 0 2 8.25c0 2.8.656 3.054 1 4.55a5.77 5.77 0 0 1 .029 2.758L2 20m17.8-2.2a7.5 7.5 0 0 0 .003-10.603M17 15a3.5 3.5 0 0 0-.025-4.975"/>`), g.Group(children),
	)
}

func SpellCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 16l6-12l6 12M8 12h8m0 8l2 2l4-4"/>`), g.Group(children),
	)
}

func SpellCheckTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 16l6-12l6 12M8 12h8M4 21c1.1 0 1.1-1 2.3-1s1.1 1 2.3 1c1.1 0 1.1-1 2.3-1c1.1 0 1.1 1 2.3 1c1.1 0 1.1-1 2.3-1c1.1 0 1.1 1 2.3 1c1.1 0 1.1-1 2.3-1"/>`), g.Group(children),
	)
}

func Spline(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="19" cy="5" r="2"/><circle cx="5" cy="19" r="2"/><path d="M5 17A12 12 0 0 1 17 5"/></g>`), g.Group(children),
	)
}

func Split(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 3h5v5M8 3H3v5"/><path d="M12 22v-8.3a4 4 0 0 0-1.172-2.872L3 3m12 6l6-6"/></g>`), g.Group(children),
	)
}

func SplitSquareHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 19H5c-1 0-2-1-2-2V7c0-1 1-2 2-2h3m8 0h3c1 0 2 1 2 2v10c0 1-1 2-2 2h-3M12 4v16"/>`), g.Group(children),
	)
}

func SplitSquareVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 8V5c0-1 1-2 2-2h10c1 0 2 1 2 2v3m0 8v3c0 1-1 2-2 2H7c-1 0-2-1-2-2v-3m-1-4h16"/>`), g.Group(children),
	)
}

func SprayCan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 3h.01M7 5h.01M11 7h.01M3 7h.01M7 9h.01M3 11h.01M15 5h4v4h-4zm4 4l2 2v10c0 .6-.4 1-1 1h-6c-.6 0-1-.4-1-1V11l2-2m-2 5l8-2m-8 7l8-2"/>`), g.Group(children),
	)
}

func Sprout(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 20h10m-7 0c5.5-2.5.8-6.4 3-10"/><path d="M9.5 9.4c1.1.8 1.8 2.2 2.3 3.7c-2 .4-3.5.4-4.8-.3c-1.2-.6-2.3-1.9-3-4.2c2.8-.5 4.4 0 5.5.8zM14.1 6a7 7 0 0 0-1.1 4c1.9-.1 3.3-.6 4.3-1.4c1-1 1.6-2.3 1.7-4.6c-2.7.1-4 1-4.9 2z"/></g>`), g.Group(children),
	)
}

func Square(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="18" height="18" x="3" y="3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="2"/>`), g.Group(children),
	)
}

func SquareAsterisk(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M12 8v8m-3.5-2l7-4m-7 0l7 4"/></g>`), g.Group(children),
	)
}

func SquareCode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m10 10l-2 2l2 2m4 0l2-2l-2-2"/></g>`), g.Group(children),
	)
}

func SquareDashedBottom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2M9 21h1m4 0h1"/>`), g.Group(children),
	)
}

func SquareDashedBottomCode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m10 10l-2 2l2 2m4 0l2-2l-2-2"/><path d="M5 21a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2M9 21h1m4 0h1"/></g>`), g.Group(children),
	)
}

func SquareDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><circle cx="12" cy="12" r="1"/></g>`), g.Group(children),
	)
}

func SquareEqual(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M7 10h10M7 14h10"/></g>`), g.Group(children),
	)
}

func SquareSlash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="m9 15l6-6"/></g>`), g.Group(children),
	)
}

func SquareStack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 10c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h4c1.1 0 2 .9 2 2m0 12c-1.1 0-2-.9-2-2v-4c0-1.1.9-2 2-2h4c1.1 0 2 .9 2 2"/><rect width="8" height="8" x="14" y="14" rx="2"/></g>`), g.Group(children),
	)
}

func Squirrel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 6a4 4 0 0 0-4 4a7 7 0 0 0-7 7c0-5 4-5 4-10.5a4.5 4.5 0 1 0-9 0a2.5 2.5 0 0 0 5 0C7 10 3 11 3 17c0 2.8 2.2 5 5 5h10"/><path d="M16 20c0-1.7 1.3-3 3-3h1a2 2 0 0 0 2-2v-2a4 4 0 0 0-4-4V4"/><path d="M15.2 22a3 3 0 0 0-2.2-5m5-4h.01"/></g>`), g.Group(children),
	)
}

func Stamp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 22h14m.27-8.27A2.5 2.5 0 0 0 17.5 13h-11A2.5 2.5 0 0 0 4 15.5V17a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-1.5c0-.66-.26-1.3-.73-1.77Z"/><path d="M14 13V8.5C14 7 15 7 15 5a3 3 0 0 0-3-3c-1.66 0-3 1-3 3s1 2 1 3.5V13"/></g>`), g.Group(children),
	)
}

func Star(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 2l3.09 6.26L22 9.27l-5 4.87l1.18 6.88L12 17.77l-6.18 3.25L7 14.14L2 9.27l6.91-1.01L12 2z"/>`), g.Group(children),
	)
}

func StarHalf(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 17.8L5.8 21L7 14.1L2 9.3l7-1L12 2"/>`), g.Group(children),
	)
}

func StarOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.34 8.34L2 9.27l5 4.87L5.82 21L12 17.77L18.18 21l-.59-3.43m.83-4.81L22 9.27l-6.91-1L12 2l-1.44 2.91M2 2l20 20"/>`), g.Group(children),
	)
}

func StepBack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 20V4m-4 16L4 12l10-8z"/>`), g.Group(children),
	)
}

func StepForward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 4v16m4-16l10 8l-10 8z"/>`), g.Group(children),
	)
}

func Stethoscope(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4.8 2.3A.3.3 0 1 0 5 2H4a2 2 0 0 0-2 2v5a6 6 0 0 0 6 6v0a6 6 0 0 0 6-6V4a2 2 0 0 0-2-2h-1a.2.2 0 1 0 .3.3"/><path d="M8 15v1a6 6 0 0 0 6 6v0a6 6 0 0 0 6-6v-4"/><circle cx="20" cy="10" r="2"/></g>`), g.Group(children),
	)
}

func Sticker(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.5 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h14a2 2 0 0 0 2-2V8.5L15.5 3Z"/><path d="M15 3v6h6m-11 7s.8 1 2 1c1.3 0 2-1 2-1m-6-3h0m8 0h0"/></g>`), g.Group(children),
	)
}

func StickyNote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15.5 3H5a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h14a2 2 0 0 0 2-2V8.5L15.5 3Z"/><path d="M15 3v6h6"/></g>`), g.Group(children),
	)
}

func StopCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="M9 9h6v6H9z"/></g>`), g.Group(children),
	)
}

func Store(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 7l4.41-4.41A2 2 0 0 1 7.83 2h8.34a2 2 0 0 1 1.42.59L22 7M4 12v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/><path d="M15 22v-4a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2v4M2 7h20m0 0v3a2 2 0 0 1-2 2v0a2.7 2.7 0 0 1-1.59-.63a.7.7 0 0 0-.82 0A2.7 2.7 0 0 1 16 12a2.7 2.7 0 0 1-1.59-.63a.7.7 0 0 0-.82 0A2.7 2.7 0 0 1 12 12a2.7 2.7 0 0 1-1.59-.63a.7.7 0 0 0-.82 0A2.7 2.7 0 0 1 8 12a2.7 2.7 0 0 1-1.59-.63a.7.7 0 0 0-.82 0A2.7 2.7 0 0 1 4 12v0a2 2 0 0 1-2-2V7"/></g>`), g.Group(children),
	)
}

func StretchHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="6" x="2" y="4" rx="2"/><rect width="20" height="6" x="2" y="14" rx="2"/></g>`), g.Group(children),
	)
}

func StretchVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="20" x="4" y="2" rx="2"/><rect width="6" height="20" x="14" y="2" rx="2"/></g>`), g.Group(children),
	)
}

func Strikethrough(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 4H9a3 3 0 0 0-2.83 4M14 12a4 4 0 0 1 0 8H6m-2-8h16"/>`), g.Group(children),
	)
}

func Subscript(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 5l8 8m0-8l-8 8m16 6h-4c0-1.5.44-2 1.5-2.5S20 15.33 20 14c0-.47-.17-.93-.48-1.29a2.11 2.11 0 0 0-2.62-.44c-.42.24-.74.62-.9 1.07"/>`), g.Group(children),
	)
}

func Subtitles(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 13h4m4 0h2M7 9h2m4 0h4m4 6a2 2 0 0 1-2 2H7l-4 4V5a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2Z"/>`), g.Group(children),
	)
}

func Sun(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M12 2v2m0 16v2M4.93 4.93l1.41 1.41m11.32 11.32l1.41 1.41M2 12h2m16 0h2M6.34 17.66l-1.41 1.41M19.07 4.93l-1.41 1.41"/></g>`), g.Group(children),
	)
}

func SunDim(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M12 4h.01M20 12h.01M12 20h.01M4 12h.01m13.647-5.657h.01m-.01 11.314h.01m-11.324 0h.01m-.01-11.314h.01"/></g>`), g.Group(children),
	)
}

func SunMedium(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="4"/><path d="M12 3v1m0 16v1m-9-9h1m16 0h1m-2.636-6.364l-.707.707M6.343 17.657l-.707.707m0-12.728l.707.707m11.314 11.314l.707.707"/></g>`), g.Group(children),
	)
}

func SunMoon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8a2.83 2.83 0 0 0 4 4a4 4 0 1 1-4-4m0-6v2m0 16v2M4.9 4.9l1.4 1.4m11.4 11.4l1.4 1.4M2 12h2m16 0h2M6.3 17.7l-1.4 1.4M19.1 4.9l-1.4 1.4"/>`), g.Group(children),
	)
}

func SunSnow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9a3 3 0 1 0 0 6m-8-3h1m11 9V3m-4 1V3m0 18v-1m-6.36-1.64l.7-.7m0-11.32l-.7-.7M14 12h8m-5-8l-3 3m0 10l3 3m4-5l-3-3l3-3"/>`), g.Group(children),
	)
}

func Sunrise(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v8m-7.07.93l1.41 1.41M2 18h2m16 0h2m-2.93-7.07l-1.41 1.41M22 22H2M8 6l4-4l4 4m0 12a4 4 0 0 0-8 0"/>`), g.Group(children),
	)
}

func Sunset(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10V2m-7.07 8.93l1.41 1.41M2 18h2m16 0h2m-2.93-7.07l-1.41 1.41M22 22H2M16 6l-4 4l-4-4m8 12a4 4 0 0 0-8 0"/>`), g.Group(children),
	)
}

func Superscript(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 19l8-8m0 8l-8-8m16 1h-4c0-1.5.442-2 1.5-2.5S20 8.334 20 7.002c0-.472-.17-.93-.484-1.29a2.105 2.105 0 0 0-2.617-.436c-.42.239-.738.614-.899 1.06"/>`), g.Group(children),
	)
}

func SwissFranc(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 21V3h8M6 16h9m-5-6.5h7"/>`), g.Group(children),
	)
}

func SwitchCamera(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M11 19H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h5m4 0h7a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-5"/><circle cx="12" cy="12" r="3"/><path d="m18 22l-3-3l3-3M6 2l3 3l-3 3"/></g>`), g.Group(children),
	)
}

func Sword(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 17.5L3 6V3h3l11.5 11.5M13 19l6-6m-3 3l4 4m-1 1l2-2"/>`), g.Group(children),
	)
}

func Swords(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 17.5L3 6V3h3l11.5 11.5M13 19l6-6m-3 3l4 4m-1 1l2-2M14.5 6.5L18 3h3v3l-3.5 3.5M5 14l4 4m-2-1l-3 3m-1-1l2 2"/>`), g.Group(children),
	)
}

func Syringe(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18 2l4 4m-5 1l3-3m-1 5L8.7 19.3c-1 1-2.5 1-3.4 0l-.6-.6c-1-1-1-2.5 0-3.4L15 5m-6 6l4 4m-8 4l-3 3M14 4l6 6"/>`), g.Group(children),
	)
}

func Table(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 3v18"/><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9h18M3 15h18"/></g>`), g.Group(children),
	)
}

func TableProperties(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 3v18"/><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M21 9H3m18 6H3"/></g>`), g.Group(children),
	)
}

func TableTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 3H5a2 2 0 0 0-2 2v4m6-6h10a2 2 0 0 1 2 2v4M9 3v18m0 0h10a2 2 0 0 0 2-2V9M9 21H5a2 2 0 0 1-2-2V9m0 0h18"/>`), g.Group(children),
	)
}

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="20" x="4" y="2" rx="2" ry="2"/><path d="M12 18h.01"/></g>`), g.Group(children),
	)
}

func TabletSmartphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="10" height="14" x="3" y="8" rx="2"/><path d="M5 4a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2h-2.4M8 18h.01"/></g>`), g.Group(children),
	)
}

func Tablets(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7" cy="7" r="5"/><circle cx="17" cy="17" r="5"/><path d="M12 17h10M3.46 10.54l7.08-7.08"/></g>`), g.Group(children),
	)
}

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2H2v10l9.29 9.29c.94.94 2.48.94 3.42 0l6.58-6.58c.94-.94.94-2.48 0-3.42L12 2ZM7 7h.01"/>`), g.Group(children),
	)
}

func Tags(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 5H2v7l6.29 6.29c.94.94 2.48.94 3.42 0l3.58-3.58c.94-.94.94-2.48 0-3.42L9 5ZM6 9.01V9"/><path d="m15 5l6.3 6.3a2.4 2.4 0 0 1 0 3.4L17 19"/></g>`), g.Group(children),
	)
}

func TallyFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16M9 4v16m5-16v16m5-16v16m3-14L2 18"/>`), g.Group(children),
	)
}

func TallyFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16M9 4v16m5-16v16m5-16v16"/>`), g.Group(children),
	)
}

func TallyOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16"/>`), g.Group(children),
	)
}

func TallyThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16M9 4v16m5-16v16"/>`), g.Group(children),
	)
}

func TallyTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v16M9 4v16"/>`), g.Group(children),
	)
}

func Tangent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="17" cy="4" r="2"/><path d="M15.59 5.41L5.41 15.59"/><circle cx="4" cy="17" r="2"/><path d="M12 22s-4-9-1.5-11.5S22 12 22 12"/></g>`), g.Group(children),
	)
}

func Target(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="12" r="6"/><circle cx="12" cy="12" r="2"/></g>`), g.Group(children),
	)
}

func Tent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.5 21L14 3m6.5 18L10 3m5.5 18L12 15l-3.5 6M2 21h20"/>`), g.Group(children),
	)
}

func TentTree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="4" cy="4" r="2"/><path d="m14 5l3-3l3 3m-6 5l3-3l3 3m-3 4V2m0 12H7l-5 8h20Zm-9 0v8m1-8l5 8"/></g>`), g.Group(children),
	)
}

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 17l6-6l-6-6m8 14h8"/>`), g.Group(children),
	)
}

func TerminalSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m7 11l2-2l-2-2m4 6h4"/><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/></g>`), g.Group(children),
	)
}

func TestTube(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 2v17.5c0 1.4-1.1 2.5-2.5 2.5h0c-1.4 0-2.5-1.1-2.5-2.5V2m-1 0h7m-1 14h-5"/>`), g.Group(children),
	)
}

func TestTubeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 7L6.82 21.18a2.83 2.83 0 0 1-3.99-.01v0a2.83 2.83 0 0 1 0-4L17 3m-1-1l6 6m-10 8H4"/>`), g.Group(children),
	)
}

func TestTubes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 2v17.5A2.5 2.5 0 0 1 6.5 22v0A2.5 2.5 0 0 1 4 19.5V2m16 0v17.5a2.5 2.5 0 0 1-2.5 2.5v0a2.5 2.5 0 0 1-2.5-2.5V2M3 2h7m4 0h7M9 16H4m16 0h-5"/>`), g.Group(children),
	)
}

func Text(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 6.1H3m18 6H3M15.1 18H3"/>`), g.Group(children),
	)
}

func TextCursor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 22h-1a4 4 0 0 1-4-4V6a4 4 0 0 1 4-4h1M7 22h1a4 4 0 0 0 4-4v-1M7 2h1a4 4 0 0 1 4 4v1"/>`), g.Group(children),
	)
}

func TextCursorInput(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4h1a3 3 0 0 1 3 3a3 3 0 0 1 3-3h1m0 16h-1a3 3 0 0 1-3-3a3 3 0 0 1-3 3H5m0-4H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h1m8 0h7a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-7M9 7v10"/>`), g.Group(children),
	)
}

func TextQuote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 6H3m18 6H8m13 6H8m-5-6v6"/>`), g.Group(children),
	)
}

func TextSelect(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3a2 2 0 0 0-2 2m16-2a2 2 0 0 1 2 2m0 14a2 2 0 0 1-2 2M5 21a2 2 0 0 1-2-2M9 3h1M9 21h1m4-18h1m-1 18h1M3 9v1m18-1v1M3 14v1m18-1v1M7 8h8m-8 4h10M7 16h6"/>`), g.Group(children),
	)
}

func Theater(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 10s3-3 3-8m17 8s-3-3-3-8"/><path d="M10 2c0 4.4-3.6 8-8 8m12-8c0 4.4 3.6 8 8 8M2 10s2 2 2 5m18-5s-2 2-2 5M8 15h8M2 22v-1a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1m4 0v-1a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v1"/></g>`), g.Group(children),
	)
}

func Thermometer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0Z"/>`), g.Group(children),
	)
}

func ThermometerSnowflake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12h10M9 4v16M3 9l3 3l-3 3m9-9L9 9L6 6m0 12l3-3l1.5 1.5M20 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0Z"/>`), g.Group(children),
	)
}

func ThermometerSun(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9a4 4 0 0 0-2 7.5M12 3v2M6.6 18.4l-1.4 1.4M20 4v10.54a4 4 0 1 1-4 0V4a2 2 0 0 1 4 0ZM4 13H2m4.34-5.66L4.93 5.93"/>`), g.Group(children),
	)
}

func ThumbsDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 14V2M9 18.12L10 14H4.17a2 2 0 0 1-1.92-2.56l2.33-8A2 2 0 0 1 6.5 2H20a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-2.76a2 2 0 0 0-1.79 1.11L12 22h0a3.13 3.13 0 0 1-3-3.88Z"/>`), g.Group(children),
	)
}

func ThumbsUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 10v12m8-16.12L14 10h5.83a2 2 0 0 1 1.92 2.56l-2.33 8A2 2 0 0 1 17.5 22H4a2 2 0 0 1-2-2v-8a2 2 0 0 1 2-2h2.76a2 2 0 0 0 1.79-1.11L12 2h0a3.13 3.13 0 0 1 3 3.88Z"/>`), g.Group(children),
	)
}

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 9a3 3 0 0 1 0 6v2a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2a3 3 0 0 1 0-6V7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2Zm11-4v2m0 10v2m0-8v2"/>`), g.Group(children),
	)
}

func Timer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 2h4m-2 12l3-3"/><circle cx="12" cy="14" r="8"/></g>`), g.Group(children),
	)
}

func TimerOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 2h4m-9.4 9a8 8 0 0 0 1.7 8.7a8 8 0 0 0 8.7 1.7m-7.6-14a8 8 0 0 1 10.3 1a8 8 0 0 1 .9 10.2M2 2l20 20M12 12v-2"/>`), g.Group(children),
	)
}

func TimerReset(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 2h4m-2 12v-4m-8 3a8 8 0 0 1 8-7a8 8 0 1 1-5.3 14L4 17.6"/><path d="M9 17H4v5"/></g>`), g.Group(children),
	)
}

func ToggleLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="12" x="2" y="6" rx="6" ry="6"/><circle cx="8" cy="12" r="2"/></g>`), g.Group(children),
	)
}

func ToggleRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="12" x="2" y="6" rx="6" ry="6"/><circle cx="16" cy="12" r="2"/></g>`), g.Group(children),
	)
}

func Tornado(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 4H3m15 4H6m13 4H9m7 4h-6m1 4H9"/>`), g.Group(children),
	)
}

func Torus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><ellipse cx="12" cy="11" rx="3" ry="2"/><ellipse cx="12" cy="12.5" rx="10" ry="8.5"/></g>`), g.Group(children),
	)
}

func Touchpad(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="M2 14h20m-10 6v-6"/></g>`), g.Group(children),
	)
}

func TouchpadOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16M2 14h12m8 0h-2m-8 6v-6M2 2l20 20m0-6V6a2 2 0 0 0-2-2H10"/>`), g.Group(children),
	)
}

func TowerControl(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.2 12.27L20 6H4l1.8 6.27a1 1 0 0 0 .95.73h10.5a1 1 0 0 0 .96-.73ZM8 13v9m8 0v-9M9 6l1 7m5-7l-1 7m-2-7V2m1 0h-2"/>`), g.Group(children),
	)
}

func ToyBrick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="12" x="3" y="8" rx="1"/><path d="M10 8V5c0-.6-.4-1-1-1H6a1 1 0 0 0-1 1v3m14 0V5c0-.6-.4-1-1-1h-3a1 1 0 0 0-1 1v3"/></g>`), g.Group(children),
	)
}

func Tractor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 4h9l1 7m-9 0V4m4 6V4m10 1c-.6 0-1 .4-1 1v5.6"/><path d="m10 11l11 .9c.6 0 .9.5.8 1.1l-.8 5h-1"/><circle cx="7" cy="15" r=".5"/><circle cx="7" cy="15" r="5"/><path d="M16 18h-5"/><circle cx="18" cy="18" r="2"/></g>`), g.Group(children),
	)
}

func TrafficCone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9.3 6.2a4.55 4.55 0 0 0 5.4 0m-6.8 4.5c.9.8 2.4 1.3 4.1 1.3s3.2-.5 4.1-1.3"/><path d="M13.9 3.5a1.93 1.93 0 0 0-3.8-.1l-3 10c-.1.2-.1.4-.1.6c0 1.7 2.2 3 5 3s5-1.3 5-3c0-.2 0-.4-.1-.5Z"/><path d="m7.5 12.2l-4.7 2.7c-.5.3-.8.7-.8 1.1s.3.8.8 1.1l7.6 4.5c.9.5 2.1.5 3 0l7.6-4.5c.7-.3 1-.7 1-1.1s-.3-.8-.8-1.1l-4.7-2.8"/></g>`), g.Group(children),
	)
}

func TrainFront(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M8 3.1V7a4 4 0 0 0 8 0V3.1M9 15l-1-1m7 1l1-1"/><path d="M9 19c-2.8 0-5-2.2-5-5v-4a8 8 0 0 1 16 0v4c0 2.8-2.2 5-5 5Zm-1 0l-2 3m10-3l2 3"/></g>`), g.Group(children),
	)
}

func TrainFrontTunnel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 22V12a10 10 0 1 1 20 0v10"/><path d="M15 6.8v1.4a3 2.8 0 1 1-6 0V6.8m1 8.2h.01M14 15h.01"/><path d="M10 19a4 4 0 0 1-4-4v-3a6 6 0 1 1 12 0v3a4 4 0 0 1-4 4Zm-1 0l-2 3m8-3l2 3"/></g>`), g.Group(children),
	)
}

func TrainTrack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 17L17 2M2 14l8 8M5 11l8 8M8 8l8 8M11 5l8 8M14 2l8 8M7 22L22 7"/>`), g.Group(children),
	)
}

func TramFront(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="16" height="16" x="4" y="3" rx="2"/><path d="M4 11h16m-8-8v8m-4 8l-2 3m12 0l-2-3m-8-4h0m8 0h0"/></g>`), g.Group(children),
	)
}

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18m-2 0v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6m3 0V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2"/>`), g.Group(children),
	)
}

func TrashTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18m-2 0v14c0 1-1 2-2 2H7c-1 0-2-1-2-2V6m3 0V4c0-1 1-2 2-2h4c1 0 2 1 2 2v2m-6 5v6m4-6v6"/>`), g.Group(children),
	)
}

func TreeDeciduous(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 19h8a4 4 0 0 0 3.8-2.8a4 4 0 0 0-1.6-4.5c1-1.1 1-2.7.4-4c-.7-1.2-2.2-2-3.6-1.7a3 3 0 0 0-3-3a3 3 0 0 0-3 3c-1.4-.2-2.9.5-3.6 1.7c-.7 1.3-.5 2.9.4 4a4 4 0 0 0-1.6 4.5A4 4 0 0 0 8 19Zm4 0v3"/>`), g.Group(children),
	)
}

func TreePine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 14l3 3.3a1 1 0 0 1-.7 1.7H4.7a1 1 0 0 1-.7-1.7L7 14h-.3a1 1 0 0 1-.7-1.7L9 9h-.2A1 1 0 0 1 8 7.3L12 3l4 4.3a1 1 0 0 1-.8 1.7H15l3 3.3a1 1 0 0 1-.7 1.7H17Zm-5 8v-3"/>`), g.Group(children),
	)
}

func Trees(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 10v.2A3 3 0 0 1 8.9 16v0H5v0h0a3 3 0 0 1-1-5.8V10a3 3 0 0 1 6 0Zm-3 6v6m6-3v3"/><path d="M12 19h8.3a1 1 0 0 0 .7-1.7L18 14h.3a1 1 0 0 0 .7-1.7L16 9h.2a1 1 0 0 0 .8-1.7L13 3l-1.4 1.5"/></g>`), g.Group(children),
	)
}

func Trello(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="M7 7h3v9H7zm7 0h3v5h-3z"/></g>`), g.Group(children),
	)
}

func TrendingDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m22 17l-8.5-8.5l-5 5L2 7"/><path d="M16 17h6v-6"/></g>`), g.Group(children),
	)
}

func TrendingUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m22 7l-8.5 8.5l-5-5L2 17"/><path d="M16 7h6v6"/></g>`), g.Group(children),
	)
}

func Triangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.73 4a2 2 0 0 0-3.46 0l-8 14A2 2 0 0 0 4 21h16a2 2 0 0 0 1.73-3Z"/>`), g.Group(children),
	)
}

func TriangleRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 18a2 2 0 0 1-2 2H3c-1.1 0-1.3-.6-.4-1.3L20.4 4.3c.9-.7 1.6-.4 1.6.7Z"/>`), g.Group(children),
	)
}

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M6 9H4.5a2.5 2.5 0 0 1 0-5H6m12 5h1.5a2.5 2.5 0 0 0 0-5H18M4 22h16m-10-7.34V17c0 .55-.47.98-.97 1.21C7.85 18.75 7 20.24 7 22m7-7.34V17c0 .55.47.98.97 1.21C16.15 18.75 17 20.24 17 22"/><path d="M18 2H6v7a6 6 0 0 0 12 0V2Z"/></g>`), g.Group(children),
	)
}

func Truck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 18H3c-.6 0-1-.4-1-1V7c0-.6.4-1 1-1h10c.6 0 1 .4 1 1v11m0-9h4l4 4v4c0 .6-.4 1-1 1h-2"/><circle cx="7" cy="18" r="2"/><path d="M15 18H9"/><circle cx="17" cy="18" r="2"/></g>`), g.Group(children),
	)
}

func Turtle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 10l2 4v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3a8 8 0 1 0-16 0v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3l2-4h4ZM4.82 7.9L8 10m7.18-2.1L12 10"/><path d="M16.93 10H20a2 2 0 0 1 0 4H2"/></g>`), g.Group(children),
	)
}

func Tv(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="15" x="2" y="7" rx="2" ry="2"/><path d="m17 2l-5 5l-5-5"/></g>`), g.Group(children),
	)
}

func TvTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 21h10"/><rect width="20" height="14" x="2" y="3" rx="2"/></g>`), g.Group(children),
	)
}

func Twitch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 2H3v16h5v4l4-4h5l4-4V2zm-10 9V7m5 4V7"/>`), g.Group(children),
	)
}

func Twitter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 4s-.7 2.1-2 3.4c1.6 10-9.4 17.3-18 11.6c2.2.1 4.4-.6 6-2C3 15.5.5 9.6 3 5c2.2 2.6 5.6 4.1 9 4c-.9-4.2 4-6.6 7-3.8c1.1 0 3-1.2 3-1.2z"/>`), g.Group(children),
	)
}

func Type(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 7V4h16v3M9 20h6M12 4v16"/>`), g.Group(children),
	)
}

func Umbrella(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 12a10.06 10.06 1 0 0-20 0Zm-10 0v8a2 2 0 0 0 4 0M12 2v1"/>`), g.Group(children),
	)
}

func UmbrellaOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v1m3.5 18a1.85 1.85 0 0 1-3.5-1v-8H2a10 10 0 0 1 3.428-6.575M17.5 12H22A10 10 0 0 0 9.004 3.455M2 2l20 20"/>`), g.Group(children),
	)
}

func Underline(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 4v6a6 6 0 0 0 12 0V4M4 20h16"/>`), g.Group(children),
	)
}

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 7v6h6"/><path d="M21 17a9 9 0 0 0-9-9a9 9 0 0 0-6 2.3L3 13"/></g>`), g.Group(children),
	)
}

func UndoDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="17" r="1"/><path d="M3 7v6h6"/><path d="M21 17a9 9 0 0 0-9-9a9 9 0 0 0-6 2.3L3 13"/></g>`), g.Group(children),
	)
}

func UndoTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M9 14L4 9l5-5"/><path d="M4 9h10.5a5.5 5.5 0 0 1 5.5 5.5v0a5.5 5.5 0 0 1-5.5 5.5H11"/></g>`), g.Group(children),
	)
}

func UnfoldHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12h6M8 12H2M12 2v2m0 4v2m0 4v2m0 4v2m7-7l3-3l-3-3M5 9l-3 3l3 3"/>`), g.Group(children),
	)
}

func UnfoldVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 22v-6m0-8V2M4 12H2m8 0H8m8 0h-2m8 0h-2m-5 7l-3 3l-3-3m6-14l-3-3l-3 3"/>`), g.Group(children),
	)
}

func Ungroup(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="6" x="5" y="4" rx="1"/><rect width="8" height="6" x="11" y="14" rx="1"/></g>`), g.Group(children),
	)
}

func Unlink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m18.84 12.25l1.72-1.71h-.02a5.004 5.004 0 0 0-.12-7.07a5.006 5.006 0 0 0-6.95 0l-1.72 1.71m-6.58 6.57l-1.71 1.71a5.004 5.004 0 0 0 .12 7.07a5.006 5.006 0 0 0 6.95 0l1.71-1.71M8 2v3M2 8h3m11 11v3m3-6h3"/>`), g.Group(children),
	)
}

func UnlinkTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 7h2a5 5 0 0 1 0 10h-2m-6 0H7A5 5 0 0 1 7 7h2"/>`), g.Group(children),
	)
}

func Unlock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="11" x="3" y="11" rx="2" ry="2"/><path d="M7 11V7a5 5 0 0 1 9.9-1"/></g>`), g.Group(children),
	)
}

func UnlockKeyhole(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="16" r="1"/><rect width="18" height="12" x="3" y="10" rx="2"/><path d="M7 10V7a5 5 0 0 1 9.33-2.5"/></g>`), g.Group(children),
	)
}

func Unplug(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 5l3-3M2 22l3-3m1.3 1.3a2.4 2.4 0 0 0 3.4 0L12 18l-6-6l-2.3 2.3a2.4 2.4 0 0 0 0 3.4Zm1.2-6.8L10 11m.5 5.5L13 14m-1-8l6 6l2.3-2.3a2.4 2.4 0 0 0 0-3.4l-2.6-2.6a2.4 2.4 0 0 0-3.4 0Z"/>`), g.Group(children),
	)
}

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 15v4a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-4m14-7l-5-5l-5 5m5-5v12"/>`), g.Group(children),
	)
}

func UploadCloud(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 14.899A7 7 0 1 1 15.71 8h1.79a4.5 4.5 0 0 1 2.5 8.242M12 12v9"/><path d="m16 16l-4-4l-4 4"/></g>`), g.Group(children),
	)
}

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="10" cy="7" r="1"/><circle cx="4" cy="20" r="1"/><path d="M4.7 19.3L19 5m2-2l-3 1l2 2ZM9.26 7.68L5 12l2 5m3-3l5 2l3.5-3.5"/><path d="m18 12l1-1l1 1l-1 1Z"/></g>`), g.Group(children),
	)
}

func User(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M19 21v-2a4 4 0 0 0-4-4H9a4 4 0 0 0-4 4v2"/><circle cx="12" cy="7" r="4"/></g>`), g.Group(children),
	)
}

func UserCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="m16 11l2 2l4-4"/></g>`), g.Group(children),
	)
}

func UserCheckTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 19a6 6 0 0 0-12 0"/><circle cx="8" cy="9" r="4"/><path d="m16 11l2 2l4-4"/></g>`), g.Group(children),
	)
}

func UserCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><circle cx="12" cy="10" r="3"/><path d="M7 20.662V19a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v1.662"/></g>`), g.Group(children),
	)
}

func UserCircleTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 20a6 6 0 0 0-12 0"/><circle cx="12" cy="10" r="4"/><circle cx="12" cy="12" r="10"/></g>`), g.Group(children),
	)
}

func UserCog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="15" r="3"/><circle cx="9" cy="7" r="4"/><path d="M10 15H6a4 4 0 0 0-4 4v2m19.7-4.6l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.2-5.6l.3-.9m.2 7.4l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4"/></g>`), g.Group(children),
	)
}

func UserCogTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="18" cy="15" r="3"/><circle cx="8" cy="9" r="4"/><path d="M10.5 13.5A6 6 0 0 0 2 19m19.7-2.6l-.9-.3m-5.6-2.2l-.9-.3m2.3 5.1l.3-.9m2.2-5.6l.3-.9m.2 7.4l-.4-1m-2.4-5.4l-.4-1m-2.1 5.3l1-.4m5.4-2.4l1-.4"/></g>`), g.Group(children),
	)
}

func UserMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M22 11h-6"/></g>`), g.Group(children),
	)
}

func UserMinusTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 19a6 6 0 0 0-12 0"/><circle cx="8" cy="9" r="4"/><path d="M22 11h-6"/></g>`), g.Group(children),
	)
}

func UserPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M19 8v6m3-3h-6"/></g>`), g.Group(children),
	)
}

func UserPlusTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 19a6 6 0 0 0-12 0"/><circle cx="8" cy="9" r="4"/><path d="M19 8v6m3-3h-6"/></g>`), g.Group(children),
	)
}

func UserSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><circle cx="12" cy="10" r="3"/><path d="M7 21v-2a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v2"/></g>`), g.Group(children),
	)
}

func UserSquareTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 21a6 6 0 0 0-12 0"/><circle cx="12" cy="11" r="4"/><rect width="18" height="18" x="3" y="3" rx="2"/></g>`), g.Group(children),
	)
}

func UserTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="8" r="5"/><path d="M20 21a8 8 0 1 0-16 0"/></g>`), g.Group(children),
	)
}

func UserX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="m17 8l5 5m0-5l-5 5"/></g>`), g.Group(children),
	)
}

func UserXTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 19a6 6 0 0 0-12 0"/><circle cx="8" cy="9" r="4"/><path d="m17 8l5 5m0-5l-5 5"/></g>`), g.Group(children),
	)
}

func Users(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 21v-2a4 4 0 0 0-4-4H6a4 4 0 0 0-4 4v2"/><circle cx="9" cy="7" r="4"/><path d="M22 21v-2a4 4 0 0 0-3-3.87m-3-12a4 4 0 0 1 0 7.75"/></g>`), g.Group(children),
	)
}

func UsersTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M14 19a6 6 0 0 0-12 0"/><circle cx="8" cy="9" r="4"/><path d="M22 19a6 6 0 0 0-6-6a4 4 0 1 0 0-8"/></g>`), g.Group(children),
	)
}

func Utensils(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 2v7c0 1.1.9 2 2 2h4a2 2 0 0 0 2-2V2M7 2v20m14-7V2v0a5 5 0 0 0-5 5v6c0 1.1.9 2 2 2h3Zm0 0v7"/>`), g.Group(children),
	)
}

func UtensilsCrossed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 2l-2.3 2.3a3 3 0 0 0 0 4.2l1.8 1.8a3 3 0 0 0 4.2 0L22 8m-7 7L3.3 3.3a4.2 4.2 0 0 0 0 6l7.3 7.3c.7.7 2 .7 2.8 0L15 15Zm0 0l7 7m-19.9-.2l6.4-6.3M19 5l-7 7"/>`), g.Group(children),
	)
}

func UtilityPole(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 2v20M2 5h20M3 3v2m4-2v2m10-2v2m4-2v2m-2 0l-7 7l-7-7"/>`), g.Group(children),
	)
}

func Variable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 21s-4-3-4-9s4-9 4-9m8 0s4 3 4 9s-4 9-4 9M15 9l-6 6m0-6l6 6"/>`), g.Group(children),
	)
}

func Vegan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 2a26.6 26.6 0 0 1 10 20c.9-6.82 1.5-9.5 4-14m0 0c4 0 6-2 6-6c-4 0-6 2-6 6"/><path d="M17.41 3.6a10 10 0 1 0 3 3"/></g>`), g.Group(children),
	)
}

func VenetianMask(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 12a5 5 0 0 0 5 5a8 8 0 0 1 5 2a8 8 0 0 1 5-2a5 5 0 0 0 5-5V7h-5a8 8 0 0 0-5 2a8 8 0 0 0-5-2H2Z"/><path d="M6 11c1.5 0 3 .5 3 2c-2 0-3 0-3-2Zm12 0c-1.5 0-3 .5-3 2c2 0 3 0 3-2Z"/></g>`), g.Group(children),
	)
}

func Verified(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 3c-1.2 0-2.4.6-3 1.7A3.6 3.6 0 0 0 4.6 9c-1 .6-1.7 1.8-1.7 3s.7 2.4 1.7 3c-.3 1.2 0 2.5 1 3.4c.8.8 2.1 1.2 3.3 1c.6 1 1.8 1.6 3 1.6s2.4-.6 3-1.7c1.2.3 2.5 0 3.4-1c.8-.8 1.2-2 1-3.3c1-.6 1.6-1.8 1.6-3s-.6-2.4-1.7-3c.3-1.2 0-2.5-1-3.4a3.7 3.7 0 0 0-3.3-1c-.6-1-1.8-1.6-3-1.6Z"/><path d="m9 12l2 2l4-4"/></g>`), g.Group(children),
	)
}

func Vibrate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 8l2 2l-2 2l2 2l-2 2m20-8l-2 2l2 2l-2 2l2 2"/><rect width="8" height="14" x="8" y="5" rx="1"/></g>`), g.Group(children),
	)
}

func VibrateOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 8l2 2l-2 2l2 2l-2 2m20-8l-2 2l2 2l-2 2l2 2M8 8v10c0 .55.45 1 1 1h6c.55 0 1-.45 1-1v-2m0-5.66V6c0-.55-.45-1-1-1h-4.34M2 2l20 20"/>`), g.Group(children),
	)
}

func Video(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m22 8l-6 4l6 4V8Z"/><rect width="14" height="12" x="2" y="6" rx="2" ry="2"/></g>`), g.Group(children),
	)
}

func VideoOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.66 6H14a2 2 0 0 1 2 2v2.34l1 1L22 8v8m-6 0a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h2l10 10ZM2 2l20 20"/>`), g.Group(children),
	)
}

func Videotape(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="20" height="16" x="2" y="4" rx="2"/><path d="M2 8h20"/><circle cx="8" cy="14" r="2"/><path d="M8 12h8"/><circle cx="16" cy="14" r="2"/></g>`), g.Group(children),
	)
}

func View(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5 12s2.545-5 7-5c4.454 0 7 5 7 5s-2.546 5-7 5c-4.455 0-7-5-7-5z"/><path d="M12 13a1 1 0 1 0 0-2a1 1 0 0 0 0 2zm9 4v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2M21 7V5a2 2 0 0 0-2-2H5a2 2 0 0 0-2 2v2"/></g>`), g.Group(children),
	)
}

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="6" cy="12" r="4"/><circle cx="18" cy="12" r="4"/><path d="M6 16h12"/></g>`), g.Group(children),
	)
}

func Volume(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5z"/>`), g.Group(children),
	)
}

func VolumeOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5zm4.54 3.46a5 5 0 0 1 0 7.07"/>`), g.Group(children),
	)
}

func VolumeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5zm4.54 3.46a5 5 0 0 1 0 7.07m3.53-10.6a10 10 0 0 1 0 14.14"/>`), g.Group(children),
	)
}

func VolumeX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 5L6 9H2v6h4l5 4V5zm11 4l-6 6m0-6l6 6"/>`), g.Group(children),
	)
}

func Vote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m9 12l2 2l4-4"/><path d="M5 7c0-1.1.9-2 2-2h10a2 2 0 0 1 2 2v12H5V7Zm17 12H2"/></g>`), g.Group(children),
	)
}

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 12V7H5a2 2 0 0 1 0-4h14v4"/><path d="M3 5v14a2 2 0 0 0 2 2h16v-5"/><path d="M18 12a2 2 0 0 0 0 4h4v-4Z"/></g>`), g.Group(children),
	)
}

func WalletCards(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2"/><path d="M3 9a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2M3 11h3c.8 0 1.6.3 2.1.9l1.1.9c1.6 1.6 4.1 1.6 5.7 0l1.1-.9c.5-.5 1.3-.9 2.1-.9H21"/></g>`), g.Group(children),
	)
}

func WalletTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 14h.01M7 7h12a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h14"/>`), g.Group(children),
	)
}

func Wallpaper(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="8" cy="9" r="2"/><path d="m9 17l6.1-6.1a2 2 0 0 1 2.81.01L22 15V5a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2M8 21h8m-4-4v4"/></g>`), g.Group(children),
	)
}

func Wand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 4V2m0 14v-2M8 9h2m10 0h2m-4.2 2.8L19 13m-4-4h0m2.8-2.8L19 5M3 21l9-9m.2-5.8L11 5"/>`), g.Group(children),
	)
}

func WandTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21.64 3.64l-1.28-1.28a1.21 1.21 0 0 0-1.72 0L2.36 18.64a1.21 1.21 0 0 0 0 1.72l1.28 1.28a1.2 1.2 0 0 0 1.72 0L21.64 5.36a1.2 1.2 0 0 0 0-1.72ZM14 7l3 3M5 6v4m14 4v4M10 2v2M7 8H3m18 8h-4M11 3H9"/>`), g.Group(children),
	)
}

func Warehouse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M22 8.35V20a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8.35A2 2 0 0 1 3.26 6.5l8-3.2a2 2 0 0 1 1.48 0l8 3.2A2 2 0 0 1 22 8.35ZM6 18h12M6 14h12"/><path d="M6 10h12v12H6z"/></g>`), g.Group(children),
	)
}

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="6"/><path d="M12 10v2l1 1m3.13-5.34l-.81-4.05a2 2 0 0 0-2-1.61h-2.68a2 2 0 0 0-2 1.61l-.78 4.05m.02 8.7l.8 4a2 2 0 0 0 2 1.61h2.72a2 2 0 0 0 2-1.61l.81-4.05"/></g>`), g.Group(children),
	)
}

func Waves(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 6c.6.5 1.2 1 2.5 1C7 7 7 5 9.5 5c2.6 0 2.4 2 5 2c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1M2 12c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c2.6 0 2.4 2 5 2c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1M2 18c.6.5 1.2 1 2.5 1c2.5 0 2.5-2 5-2c2.6 0 2.4 2 5 2c2.5 0 2.5-2 5-2c1.3 0 1.9.5 2.5 1"/>`), g.Group(children),
	)
}

func Waypoints(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="4.5" r="2.5"/><path d="m10.2 6.3l-3.9 3.9"/><circle cx="4.5" cy="12" r="2.5"/><path d="M7 12h10"/><circle cx="19.5" cy="12" r="2.5"/><path d="m13.8 17.7l3.9-3.9"/><circle cx="12" cy="19.5" r="2.5"/></g>`), g.Group(children),
	)
}

func Webcam(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="10" r="8"/><circle cx="12" cy="10" r="3"/><path d="M7 22h10m-5 0v-4"/></g>`), g.Group(children),
	)
}

func Webhook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M18 16.98h-5.99c-1.1 0-1.95.94-2.48 1.9A4 4 0 0 1 2 17c.01-.7.2-1.4.57-2"/><path d="m6 17l3.13-5.78c.53-.97.1-2.18-.5-3.1a4 4 0 1 1 6.89-4.06"/><path d="m12 6l3.13 5.73C15.66 12.7 16.9 13 18 13a4 4 0 0 1 0 8"/></g>`), g.Group(children),
	)
}

func Weight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="5" r="3"/><path d="M6.5 8a2 2 0 0 0-1.905 1.46L2.1 18.5A2 2 0 0 0 4 21h16a2 2 0 0 0 1.925-2.54L19.4 9.5A2 2 0 0 0 17.48 8Z"/></g>`), g.Group(children),
	)
}

func Wheat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2 22L16 8M3.47 12.53L5 11l1.53 1.53a3.5 3.5 0 0 1 0 4.94L5 19l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Zm4-4L9 7l1.53 1.53a3.5 3.5 0 0 1 0 4.94L9 15l-1.53-1.53a3.5 3.5 0 0 1 0-4.94Zm4-4L13 3l1.53 1.53a3.5 3.5 0 0 1 0 4.94L13 11l-1.53-1.53a3.5 3.5 0 0 1 0-4.94ZM20 2h2v2a4 4 0 0 1-4 4h-2V6a4 4 0 0 1 4-4Z"/><path d="M11.47 17.47L13 19l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L5 19l1.53-1.53a3.5 3.5 0 0 1 4.94 0Zm4-4L17 15l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L9 15l1.53-1.53a3.5 3.5 0 0 1 4.94 0Zm4-4L21 11l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L13 11l1.53-1.53a3.5 3.5 0 0 1 4.94 0Z"/></g>`), g.Group(children),
	)
}

func WheatOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m2 22l10-10m4-4l-1.17 1.17M3.47 12.53L5 11l1.53 1.53a3.5 3.5 0 0 1 0 4.94L5 19l-1.53-1.53a3.5 3.5 0 0 1 0-4.94ZM8 8l-.53.53a3.5 3.5 0 0 0 0 4.94L9 15l1.53-1.53c.55-.55.88-1.25.98-1.97m-.6-6.24c.15-.26.34-.51.56-.73L13 3l1.53 1.53a3.5 3.5 0 0 1 .28 4.62M20 2h2v2a4 4 0 0 1-4 4h-2V6a4 4 0 0 1 4-4Z"/><path d="M11.47 17.47L13 19l-1.53 1.53a3.5 3.5 0 0 1-4.94 0L5 19l1.53-1.53a3.5 3.5 0 0 1 4.94 0ZM16 16l-.53.53a3.5 3.5 0 0 1-4.94 0L9 15l1.53-1.53a3.49 3.49 0 0 1 1.97-.98m6.24.6c.26-.15.51-.34.73-.56L21 11l-1.53-1.53a3.5 3.5 0 0 0-4.62-.28M2 2l20 20"/></g>`), g.Group(children),
	)
}

func WholeWord(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="7" cy="12" r="3"/><path d="M10 9v6"/><circle cx="17" cy="12" r="3"/><path d="M14 7v8m8 2v1c0 .5-.5 1-1 1H3c-.5 0-1-.5-1-1v-1"/></g>`), g.Group(children),
	)
}

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 13a10 10 0 0 1 14 0M8.5 16.5a5 5 0 0 1 7 0M2 8.82a15 15 0 0 1 20 0M12 20h.01"/>`), g.Group(children),
	)
}

func WifiOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m2 2l20 20M8.5 16.5a5 5 0 0 1 7 0M2 8.82a15 15 0 0 1 4.17-2.65M10.66 5c4.01-.36 8.14.9 11.34 3.76m-5.15 2.49a10 10 0 0 1 2.22 1.68M5 13a10 10 0 0 1 5.24-2.76M12 20h.01"/>`), g.Group(children),
	)
}

func Wind(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.7 7.7a2.5 2.5 0 1 1 1.8 4.3H2m7.6-7.4A2 2 0 1 1 11 8H2m10.6 11.4A2 2 0 1 0 14 16H2"/>`), g.Group(children),
	)
}

func Wine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 22h8M7 10h10m-5 5v7m0-7a5 5 0 0 0 5-5c0-2-.5-4-2-8H9c-1.5 4-2 6-2 8a5 5 0 0 0 5 5Z"/>`), g.Group(children),
	)
}

func WineOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 22h8M7 10h3m7 0h-1.343M12 15v7M7.307 7.307A12.33 12.33 0 0 0 7 10a5 5 0 0 0 7.391 4.391M8.638 2.981C8.75 2.668 8.872 2.34 9 2h6c1.5 4 2 6 2 8c0 .407-.05.809-.145 1.198M2 2l20 20"/>`), g.Group(children),
	)
}

func Workflow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="8" height="8" x="3" y="3" rx="2"/><path d="M7 11v4a2 2 0 0 0 2 2h4"/><rect width="8" height="8" x="13" y="13" rx="2"/></g>`), g.Group(children),
	)
}

func WrapText(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 6h18M3 12h15a3 3 0 1 1 0 6h-4"/><path d="m16 16l-2 2l2 2M3 18h7"/></g>`), g.Group(children),
	)
}

func Wrench(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.7 6.3a1 1 0 0 0 0 1.4l1.6 1.6a1 1 0 0 0 1.4 0l3.77-3.77a6 6 0 0 1-7.94 7.94l-6.91 6.91a2.12 2.12 0 0 1-3-3l6.91-6.91a6 6 0 0 1 7.94-7.94l-3.76 3.76z"/>`), g.Group(children),
	)
}

func X(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 6L6 18M6 6l12 12"/>`), g.Group(children),
	)
}

func XCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="10"/><path d="m15 9l-6 6m0-6l6 6"/></g>`), g.Group(children),
	)
}

func XOctagon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.86 2h8.28L22 7.86v8.28L16.14 22H7.86L2 16.14V7.86L7.86 2zM15 9l-6 6m0-6l6 6"/>`), g.Group(children),
	)
}

func XSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="18" height="18" x="3" y="3" rx="2" ry="2"/><path d="m15 9l-6 6m0-6l6 6"/></g>`), g.Group(children),
	)
}

func Youtube(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M2.5 17a24.12 24.12 0 0 1 0-10a2 2 0 0 1 1.4-1.4a49.56 49.56 0 0 1 16.2 0A2 2 0 0 1 21.5 7a24.12 24.12 0 0 1 0 10a2 2 0 0 1-1.4 1.4a49.55 49.55 0 0 1-16.2 0A2 2 0 0 1 2.5 17"/><path d="m10 15l5-3l-5-3z"/></g>`), g.Group(children),
	)
}

func Zap(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 2L3 14h9l-1 8l10-12h-9l1-8z"/>`), g.Group(children),
	)
}

func ZapOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.41 6.75L13 2l-2.43 2.92m8 7.99L21 10h-5.34M8 8l-5 6h9l-1 8l5-6M2 2l20 20"/>`), g.Group(children),
	)
}

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.35-4.35M11 8v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func ZoomOut(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="11" cy="11" r="8"/><path d="m21 21l-4.35-4.35M8 11h6"/></g>`), g.Group(children),
	)
}
