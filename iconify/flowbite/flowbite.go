package flowbite

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

const IconifyVersion = ""

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 24 24")
)

func AddressBookOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6H5m2 3H5m2 3H5m2 3H5m2 3H5m11-1a2 2 0 0 0-2-2h-2a2 2 0 0 0-2 2M7 3h11a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1m8 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/>`), g.Group(children),
	)
}

func AddressBookSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 2a2 2 0 0 0-2 2v1a1 1 0 0 0 0 2v1a1 1 0 0 0 0 2v1a1 1 0 1 0 0 2v1a1 1 0 1 0 0 2v1a1 1 0 1 0 0 2v1a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm3 8a3 3 0 1 1 6 0a3 3 0 0 1-6 0m-1 7a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3a1 1 0 0 1-1 1h-6a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func AdjustmentsHorizontalOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M20 6H10m0 0a2 2 0 1 0-4 0m4 0a2 2 0 1 1-4 0m0 0H4m16 6h-2m0 0a2 2 0 1 0-4 0m4 0a2 2 0 1 1-4 0m0 0H4m16 6H10m0 0a2 2 0 1 0-4 0m4 0a2 2 0 1 1-4 0m0 0H4"/>`), g.Group(children),
	)
}

func AdjustmentsHorizontalSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.83 5a3.001 3.001 0 0 0-5.66 0H4a1 1 0 1 0 0 2h1.17a3.001 3.001 0 0 0 5.66 0H20a1 1 0 1 0 0-2zM4 11h9.17a3.001 3.001 0 0 1 5.66 0H20a1 1 0 1 1 0 2h-1.17a3.001 3.001 0 0 1-5.66 0H4a1 1 0 1 1 0-2m1.17 6H4a1 1 0 1 0 0 2h1.17a3.001 3.001 0 0 0 5.66 0H20a1 1 0 1 0 0-2h-9.17a3.001 3.001 0 0 0-5.66 0"/>`), g.Group(children),
	)
}

func AdjustmentsVerticalOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M6 4v10m0 0a2 2 0 1 0 0 4m0-4a2 2 0 1 1 0 4m0 0v2m6-16v2m0 0a2 2 0 1 0 0 4m0-4a2 2 0 1 1 0 4m0 0v10m6-16v10m0 0a2 2 0 1 0 0 4m0-4a2 2 0 1 1 0 4m0 0v2"/>`), g.Group(children),
	)
}

func AdjustmentsVerticalSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 13.17a3.001 3.001 0 0 0 0 5.66V20a1 1 0 1 0 2 0v-1.17a3.001 3.001 0 0 0 0-5.66V4a1 1 0 0 0-2 0zM11 20v-9.17a3.001 3.001 0 0 1 0-5.66V4a1 1 0 1 1 2 0v1.17a3.001 3.001 0 0 1 0 5.66V20a1 1 0 1 1-2 0m6-1.17V20a1 1 0 1 0 2 0v-1.17a3.001 3.001 0 0 0 0-5.66V4a1 1 0 1 0-2 0v9.17a3.001 3.001 0 0 0 0 5.66"/>`), g.Group(children),
	)
}

func AlignCenterOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6h8M6 10h12M8 14h8M6 18h12"/>`), g.Group(children),
	)
}

func AlignCenterSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 6c0-.6.4-1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1m-2 4c0-.6.4-1 1-1h12a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1m2 4c0-.6.4-1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1m-2 4c0-.6.4-1 1-1h12a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func AngleDownOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 9l-7 7l-7-7"/>`), g.Group(children),
	)
}

func AngleLeftOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 19l-7-7l7-7"/>`), g.Group(children),
	)
}

func AngleRightOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 5l7 7l-7 7"/>`), g.Group(children),
	)
}

func AngleUpOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 15l7-7l7 7"/>`), g.Group(children),
	)
}

func AnnotationOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.556 8.5h8m-8 3.5H12m7.111-7H4.89a.9.9 0 0 0-.629.256a.87.87 0 0 0-.26.619v9.25c0 .232.094.455.26.619A.9.9 0 0 0 4.89 16H9l3 4l3-4h4.111a.9.9 0 0 0 .629-.256a.87.87 0 0 0 .26-.619v-9.25a.87.87 0 0 0-.26-.619a.9.9 0 0 0-.63-.256Z"/>`), g.Group(children),
	)
}

func AnnotationSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.559 4.544c.355-.35.834-.544 1.33-.544H19.11c.496 0 .975.194 1.33.544s.559.829.559 1.331v9.25c0 .502-.203.981-.559 1.331c-.355.35-.834.544-1.33.544H15.5l-2.7 3.6a1 1 0 0 1-1.6 0L8.5 17H4.889c-.496 0-.975-.194-1.33-.544A1.87 1.87 0 0 1 3 15.125v-9.25c0-.502.203-.981.559-1.331M7.556 7.5a1 1 0 1 0 0 2h8a1 1 0 0 0 0-2zm0 3.5a1 1 0 1 0 0 2H12a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func AppleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.537 12.625a4.42 4.42 0 0 0 2.684 4.047a11 11 0 0 1-1.384 2.845c-.834 1.218-1.7 2.432-3.062 2.457c-1.34.025-1.77-.794-3.3-.794c-1.531 0-2.01.769-3.275.82c-1.316.049-2.317-1.318-3.158-2.532c-1.72-2.484-3.032-7.017-1.27-10.077A4.9 4.9 0 0 1 8.91 6.884c1.292-.025 2.51.869 3.3.869c.789 0 2.27-1.075 3.828-.917a4.67 4.67 0 0 1 3.66 1.984a4.52 4.52 0 0 0-2.16 3.805m-2.52-7.432A4.4 4.4 0 0 0 16.06 2a4.48 4.48 0 0 0-2.945 1.516a4.18 4.18 0 0 0-1.061 3.093a3.7 3.7 0 0 0 2.967-1.416Z"/>`), g.Group(children),
	)
}

func ArchiveArrowDownOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 11v5m0 0l2-2m-2 2l-2-2M3 6v1a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1m2 2v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V8z"/>`), g.Group(children),
	)
}

func ArchiveArrowDownSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4a2 2 0 1 0 0 4h16a2 2 0 1 0 0-4zm0 6h16v8a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10.707 5.707a1 1 0 0 0-1.414-1.414l-.293.293V12a1 1 0 1 0-2 0v2.586l-.293-.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArchiveOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M10 12v1h4v-1m4 7H6a1 1 0 0 1-1-1V9h14v9a1 1 0 0 1-1 1ZM4 5h16a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1Z"/>`), g.Group(children),
	)
}

func ArchiveSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M20 10H4v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2zM9 13v-1h6v1a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1" clip-rule="evenodd"/><path d="M2 6a2 2 0 0 1 2-2h16a2 2 0 1 1 0 4H4a2 2 0 0 1-2-2"/></g>`), g.Group(children),
	)
}

func ArrowDownOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19V5m0 14l-4-4m4 4l4-4"/>`), g.Group(children),
	)
}

func ArrowDownToBracketOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 15v2a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-2m-8 1V4m0 12l-4-4m4 4l4-4"/>`), g.Group(children),
	)
}

func ArrowDownToBracketSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M4 15v2a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-2"/><path d="M12 15.5V4"/><path stroke-linejoin="round" d="m8 12l4 4l4-4"/></g>`), g.Group(children),
	)
}

func ArrowLeftOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14M5 12l4-4m-4 4l4 4"/>`), g.Group(children),
	)
}

func ArrowLeftSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M6 12h13"/><path stroke-linejoin="round" d="m9 8l-4 4l4 4"/></g>`), g.Group(children),
	)
}

func ArrowLeftToBracketOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12H4m12 0l-4 4m4-4l-4-4m3-4h2a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3h-2"/>`), g.Group(children),
	)
}

func ArrowLeftToBracketSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M15.5 12H4"/><path stroke-linejoin="round" d="M15 4h2a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3h-2"/><path stroke-linejoin="round" d="m12 16l4-4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowRightAltOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.153 19L21 12l-4.847-7H3l4.848 7L3 19z"/>`), g.Group(children),
	)
}

func ArrowRightAltSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 0-.822 1.57L6.632 12l-4.454 6.43A1 1 0 0 0 3 20h13.153a1 1 0 0 0 .822-.43l4.847-7a1 1 0 0 0 0-1.14l-4.847-7a1 1 0 0 0-.822-.43z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowRightArrowLeftSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M4 16h13m3-8H7"/><path stroke-linejoin="round" d="m8 12l-4 4l4 4m8-8l4-4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowRightOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 12H5m14 0l-4 4m4-4l-4-4"/>`), g.Group(children),
	)
}

func ArrowRightSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M18 12H5"/><path stroke-linejoin="round" d="m15 16l4-4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowRightToBracketOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H8m12 0l-4 4m4-4l-4-4M9 4H7a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h2"/>`), g.Group(children),
	)
}

func ArrowRightToBracketSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M19.5 12H8"/><path stroke-linejoin="round" d="M9 4H7a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h2m7-4l4-4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowSortLettersOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 4v16M7 4l3 3M7 4L4 7m9-3h6l-6 6h6m-6.5 10l3.5-7l3.5 7M14 18h4"/>`), g.Group(children),
	)
}

func ArrowSortLettersSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M7 5v15"/><path stroke-linejoin="round" d="M10 7L7 4L4 7m9-3h6l-6 6h6m-6.5 10l3.5-7l3.5 7M14 18h4"/></g>`), g.Group(children),
	)
}

func ArrowUpDownOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 20V7m0 13l-4-4m4 4l4-4m4-12v13m0-13l4 4m-4-4l-4 4"/>`), g.Group(children),
	)
}

func ArrowUpDownSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M8 20V7m8-3v13"/><path stroke-linejoin="round" d="m4 16l4 4l4-4m8-8l-4-4l-4 4"/></g>`), g.Group(children),
	)
}

func ArrowUpFromBracketOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 15v2a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-2M12 4v12m0-12l4 4m-4-4L8 8"/>`), g.Group(children),
	)
}

func ArrowUpOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v13m0-13l4 4m-4-4l-4 4"/>`), g.Group(children),
	)
}

func ArrowUpRightDownLeftOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 4h4m0 0v4m0-4l-5 5M8 20H4m0 0v-4m0 4l5-5"/>`), g.Group(children),
	)
}

func ArrowUpRightDownLeftSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 4h4v4m-.5-3.5L15 9M8 20H4v-4m.5 3.5L9 15"/>`), g.Group(children),
	)
}

func ArrowUpRightFromSquareOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 14v4.833A1.166 1.166 0 0 1 16.833 20H5.167A1.167 1.167 0 0 1 4 18.833V7.167A1.166 1.166 0 0 1 5.167 6h4.618m4.447-2H20v5.768m-7.889 2.121l7.778-7.778"/>`), g.Group(children),
	)
}

func ArrowUpRightFromSquareSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.403 5H5a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-6.403a3 3 0 0 1-1.743-1.612l-3.025 3.025A3 3 0 1 1 9.99 9.768l3.025-3.025A3 3 0 0 1 11.403 5"/><path d="M13.232 4a1 1 0 0 1 1-1H20a1 1 0 0 1 1 1v5.768a1 1 0 1 1-2 0V6.414l-6.182 6.182a1 1 0 0 1-1.414-1.414L17.586 5h-3.354a1 1 0 0 1-1-1"/></g>`), g.Group(children),
	)
}

func ArrowUpSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M12 18V5"/><path stroke-linejoin="round" d="m8 15l4 4l4-4"/></g>`), g.Group(children),
	)
}

func ArrowsRepeatCountOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 4l3 3H5v3m3 10l-3-3h14v-3m-9-2.5l2-1.5v4"/>`), g.Group(children),
	)
}

func ArrowsRepeatOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 10l3-3m0 0l-3-3m3 3H5v3m3 4l-3 3m0 0l3 3m-3-3h14v-3"/>`), g.Group(children),
	)
}

func AtomOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M8.737 8.737a21.5 21.5 0 0 1 3.308-2.724m0 0c3.063-2.026 5.99-2.641 7.331-1.3c1.827 1.828.026 6.591-4.023 10.64s-8.812 5.85-10.64 4.023c-1.33-1.33-.736-4.218 1.249-7.253m6.083-6.11c-3.063-2.026-5.99-2.641-7.331-1.3c-1.827 1.828-.026 6.591 4.023 10.64m3.308-9.34a21.5 21.5 0 0 1 3.308 2.724m2.775 3.386c1.985 3.035 2.579 5.923 1.248 7.253c-1.336 1.337-4.245.732-7.295-1.275M14 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/>`), g.Group(children),
	)
}

func AtomSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M10.2 6L8 8a1 1 0 0 0 1.4 1.4A21 21 0 0 1 12 7.2a21 21 0 0 1 2.6 2.2A1 1 0 0 0 16.1 8l-2.2-2l2.6-1c1.2-.1 1.8 0 2.2.4c.4.5.6 1.6 0 3.4c-.7 1.8-2.1 3.9-4 5.8c-2 2-4 3.4-5.9 4c-1.8.7-3 .5-3.4 0c-.3-.3-.5-1-.3-2a9 9 0 0 1 1-2.7L8 16a1 1 0 0 0 1.3-1.5c-1.9-1.9-3.3-4-4-5.8c-.6-1.8-.4-3 0-3.4c.4-.3 1-.5 2.2-.3c.7.1 1.6.5 2.6 1ZM12 4.9c1.5-.8 2.9-1.4 4.2-1.7C17.6 3 19 3 20 4.1c1.3 1.3 1.2 3.5.4 5.5a15 15 0 0 1-1.2 2.4c.8 1.5 1.4 3 1.7 4.2c.2 1.4 0 2.9-1 3.9s-2.4 1.1-3.8.9c-1.3-.3-2.7-.9-4.2-1.7l-2.4 1.2c-2 .8-4.2 1-5.6-.4c-1-1-1.1-2.5-.9-3.9A12 12 0 0 1 4.7 12a15 15 0 0 1-1.2-2.4c-.8-2-1-4.2.4-5.6C5 3 6.5 3 8 3.1c1.2.3 2.6.9 4 1.7ZM14 18a9 9 0 0 0 2.7 1c1 .2 1.7 0 2-.3c.4-.4.6-1 .4-2.1a9 9 0 0 0-1-2.7A23.4 23.4 0 0 1 14 18" clip-rule="evenodd"/><path fill="currentColor" d="M14 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path stroke="currentColor" d="M10.2 6L8 8a1 1 0 0 0 1.4 1.4A21 21 0 0 1 12 7.2a21 21 0 0 1 2.6 2.2A1 1 0 0 0 16.1 8l-2.2-2l2.6-1c1.2-.1 1.8 0 2.2.4c.4.5.6 1.6 0 3.4c-.7 1.8-2.1 3.9-4 5.8c-2 2-4 3.4-5.9 4c-1.8.7-3 .5-3.4 0c-.3-.3-.5-1-.3-2a9 9 0 0 1 1-2.7L8 16a1 1 0 0 0 1.3-1.5c-1.9-1.9-3.3-4-4-5.8c-.6-1.8-.4-3 0-3.4c.4-.3 1-.5 2.2-.3c.7.1 1.6.5 2.6 1ZM12 4.9c1.5-.8 2.9-1.4 4.2-1.7C17.6 3 19 3 20 4.1c1.3 1.3 1.2 3.5.4 5.5a15 15 0 0 1-1.2 2.4c.8 1.5 1.4 3 1.7 4.2c.2 1.4 0 2.9-1 3.9s-2.4 1.1-3.8.9c-1.3-.3-2.7-.9-4.2-1.7l-2.4 1.2c-2 .8-4.2 1-5.6-.4c-1-1-1.1-2.5-.9-3.9A12 12 0 0 1 4.7 12a15 15 0 0 1-1.2-2.4c-.8-2-1-4.2.4-5.6C5 3 6.5 3 8 3.1c1.2.3 2.6.9 4 1.7ZM14 18a9 9 0 0 0 2.7 1c1 .2 1.7 0 2-.3c.4-.4.6-1 .4-2.1a9 9 0 0 0-1-2.7A23.4 23.4 0 0 1 14 18Z" clip-rule="evenodd"/><path stroke="currentColor" d="M14 12a2 2 0 1 1-4 0a2 2 0 0 1 4 0Z"/></g>`), g.Group(children),
	)
}

func AwardOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7.171 12.906l-2.153 6.411l2.672-.89l1.568 2.34l1.825-5.183m5.73-2.678l2.154 6.411l-2.673-.89l-1.568 2.34l-1.825-5.183M9.165 4.3c.58.068 1.153-.17 1.515-.628a1.68 1.68 0 0 1 2.64 0a1.68 1.68 0 0 0 1.515.628a1.68 1.68 0 0 1 1.866 1.866c-.068.58.17 1.154.628 1.516a1.68 1.68 0 0 1 0 2.639a1.68 1.68 0 0 0-.628 1.515a1.68 1.68 0 0 1-1.866 1.866a1.68 1.68 0 0 0-1.516.628a1.68 1.68 0 0 1-2.639 0a1.68 1.68 0 0 0-1.515-.628a1.68 1.68 0 0 1-1.867-1.866a1.68 1.68 0 0 0-.627-1.515a1.68 1.68 0 0 1 0-2.64c.458-.361.696-.935.627-1.515A1.68 1.68 0 0 1 9.165 4.3M14 9a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/>`), g.Group(children),
	)
}

func AwardSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M11 9a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/><path fill-rule="evenodd" d="M9.896 3.051a2.68 2.68 0 0 1 4.208 0c.147.186.38.282.615.255a2.68 2.68 0 0 1 2.976 2.975a.68.68 0 0 0 .254.615a2.68 2.68 0 0 1 0 4.208a.68.68 0 0 0-.254.615a2.68 2.68 0 0 1-2.976 2.976a.68.68 0 0 0-.615.254a2.682 2.682 0 0 1-4.208 0a.68.68 0 0 0-.614-.255a2.68 2.68 0 0 1-2.976-2.975a.68.68 0 0 0-.255-.615a2.68 2.68 0 0 1 0-4.208a.68.68 0 0 0 .255-.615a2.68 2.68 0 0 1 2.976-2.975a.68.68 0 0 0 .614-.255M12 6a3 3 0 1 0 0 6a3 3 0 0 0 0-6" clip-rule="evenodd"/><path d="M5.395 15.055L4.07 19a1 1 0 0 0 1.264 1.267l1.95-.65l1.144 1.707A1 1 0 0 0 10.2 21.1l1.12-3.18a4.64 4.64 0 0 1-2.515-1.208a4.67 4.67 0 0 1-3.411-1.656Zm7.269 2.867l1.12 3.177a1 1 0 0 0 1.773.224l1.144-1.707l1.95.65A1 1 0 0 0 19.915 19l-1.32-3.93a4.67 4.67 0 0 1-3.4 1.642a4.64 4.64 0 0 1-2.53 1.21Z"/></g>`), g.Group(children),
	)
}

func BackwardStepOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 6v12m8-12v12l-8-6z"/>`), g.Group(children),
	)
}

func BackwardStepSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 6a1 1 0 0 1 2 0v4l6.4-4.8A1 1 0 0 1 17 6v12a1 1 0 0 1-1.6.8L9 14v4a1 1 0 1 1-2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BadgeCheckOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8.032 12l1.984 1.984l4.96-4.96m4.55 5.272l.893-.893a1.984 1.984 0 0 0 0-2.806l-.893-.893a1.98 1.98 0 0 1-.581-1.403V7.04a1.984 1.984 0 0 0-1.984-1.984h-1.262a1.98 1.98 0 0 1-1.403-.581l-.893-.893a1.984 1.984 0 0 0-2.806 0l-.893.893a1.98 1.98 0 0 1-1.403.581H7.04A1.984 1.984 0 0 0 5.055 7.04v1.262c0 .527-.209 1.031-.581 1.403l-.893.893a1.984 1.984 0 0 0 0 2.806l.893.893c.372.372.581.876.581 1.403v1.262a1.984 1.984 0 0 0 1.984 1.984h1.262c.527 0 1.031.209 1.403.581l.893.893a1.984 1.984 0 0 0 2.806 0l.893-.893a2 2 0 0 1 1.403-.581h1.262a1.984 1.984 0 0 0 1.984-1.984V15.7c0-.527.209-1.031.581-1.403Z"/>`), g.Group(children),
	)
}

func BadgeCheckSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2c-.791 0-1.55.314-2.11.874l-.893.893a1 1 0 0 1-.696.288H7.04A2.984 2.984 0 0 0 4.055 7.04v1.262a1 1 0 0 1-.288.696l-.893.893a2.984 2.984 0 0 0 0 4.22l.893.893a1 1 0 0 1 .288.696v1.262a2.984 2.984 0 0 0 2.984 2.984h1.262c.261 0 .512.104.696.288l.893.893a2.984 2.984 0 0 0 4.22 0l.893-.893a1 1 0 0 1 .696-.288h1.262a2.984 2.984 0 0 0 2.984-2.984V15.7c0-.261.104-.512.288-.696l.893-.893a2.984 2.984 0 0 0 0-4.22l-.893-.893a1 1 0 0 1-.288-.696V7.04a2.984 2.984 0 0 0-2.984-2.984h-1.262a1 1 0 0 1-.696-.288l-.893-.893A2.98 2.98 0 0 0 12 2m3.683 7.73a1 1 0 1 0-1.414-1.413l-4.253 4.253l-1.277-1.277a1 1 0 0 0-1.415 1.414l1.985 1.984a1 1 0 0 0 1.414 0l4.96-4.96Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BanOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m6 6l12 12m3-6a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`), g.Group(children),
	)
}

func BarsFromLeftOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 7h14M5 12h14M5 17h10"/>`), g.Group(children),
	)
}

func BarsOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 7h14M5 12h14M5 17h14"/>`), g.Group(children),
	)
}

func BellActiveAltOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5.365V3m0 2.365a5.34 5.34 0 0 1 5.133 5.368v1.8c0 2.386 1.867 2.982 1.867 4.175c0 .593 0 1.193-.538 1.193H5.538c-.538 0-.538-.6-.538-1.193c0-1.193 1.867-1.789 1.867-4.175v-1.8A5.34 5.34 0 0 1 12 5.365m-8.134 5.368a8.46 8.46 0 0 1 2.252-5.714m14.016 5.714a8.46 8.46 0 0 0-2.252-5.714M8.54 17.901a3.48 3.48 0 0 0 6.92 0z"/>`), g.Group(children),
	)
}

func BellActiveAltSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.133 12.632v-1.8a5.41 5.41 0 0 0-4.154-5.262A1 1 0 0 0 13 5.464V3.1a1 1 0 0 0-2 0v2.364a1 1 0 0 0 .021.106a5.406 5.406 0 0 0-4.154 5.262v1.8C6.867 15.018 5 15.614 5 16.807C5 17.4 5 18 5.538 18h12.924C19 18 19 17.4 19 16.807c0-1.193-1.867-1.789-1.867-4.175m-13.267-.8a1 1 0 0 1-1-1a9.42 9.42 0 0 1 2.517-6.391A1.001 1.001 0 1 1 6.854 5.8a7.43 7.43 0 0 0-1.988 5.037a1 1 0 0 1-1 .995m16.268 0a1 1 0 0 1-1-1A7.43 7.43 0 0 0 17.146 5.8a1 1 0 0 1 1.471-1.354a9.42 9.42 0 0 1 2.517 6.391a1 1 0 0 1-1 .995M8.823 19a3.453 3.453 0 0 0 6.354 0z"/>`), g.Group(children),
	)
}

func BellActiveOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.827 5.465l-.435-2.324m.435 2.324a5.34 5.34 0 0 1 6.033 4.333l.331 1.769c.44 2.345 2.383 2.588 2.6 3.761c.11.586.22 1.171-.31 1.271l-12.7 2.377c-.529.099-.639-.488-.749-1.074C5.813 16.73 7.538 15.8 7.1 13.455c-.219-1.169.218 1.162-.33-1.769a5.34 5.34 0 0 1 4.058-6.221Zm-7.046 4.41c.143-1.877.822-3.461 2.086-4.856m2.646 13.633a3.472 3.472 0 0 0 6.728-.777l.09-.5z"/>`), g.Group(children),
	)
}

func BellActiveSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M11.209 3.816a1 1 0 0 0-1.966.368l.325 1.74a5.34 5.34 0 0 0-2.8 5.762l.276 1.473l.055.296c.258 1.374-.228 2.262-.63 2.998c-.285.52-.527.964-.437 1.449c.11.586.22 1.173.75 1.074l12.7-2.377c.528-.1.418-.685.308-1.27c-.103-.564-.636-1.123-1.195-1.711c-.606-.636-1.243-1.306-1.404-2.051c-.233-1.085-.275-1.387-.303-1.587c-.009-.063-.016-.117-.028-.182a5.34 5.34 0 0 0-5.353-4.39z"/><path fill-rule="evenodd" d="M6.539 4.278a1 1 0 0 1 .07 1.412c-1.115 1.23-1.705 2.605-1.83 4.26a1 1 0 0 1-1.995-.15c.16-2.099.929-3.893 2.342-5.453a1 1 0 0 1 1.413-.069" clip-rule="evenodd"/><path d="M8.95 19.7c.7.8 1.7 1.3 2.8 1.3c1.6 0 2.9-1.1 3.3-2.5z"/></g>`), g.Group(children),
	)
}

func BellOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5.365V3m0 2.365a5.34 5.34 0 0 1 5.133 5.368v1.8c0 2.386 1.867 2.982 1.867 4.175c0 .593 0 1.292-.538 1.292H5.538C5 18 5 17.301 5 16.708c0-1.193 1.867-1.789 1.867-4.175v-1.8A5.34 5.34 0 0 1 12 5.365M8.733 18c.094.852.306 1.54.944 2.112a3.48 3.48 0 0 0 4.646 0c.638-.572 1.236-1.26 1.33-2.112z"/>`), g.Group(children),
	)
}

func BellRingOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5.464V3.099m0 2.365a5.34 5.34 0 0 1 5.133 5.368v1.8c0 2.386 1.867 2.982 1.867 4.175C19 17.4 19 18 18.462 18H5.538C5 18 5 17.4 5 16.807c0-1.193 1.867-1.789 1.867-4.175v-1.8A5.34 5.34 0 0 1 12 5.464M6 5L5 4M4 9H3m15-4l1-1m1 5h1M8.54 18a3.48 3.48 0 0 0 6.92 0z"/>`), g.Group(children),
	)
}

func BellRingSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.133 12.632v-1.8a5.406 5.406 0 0 0-4.154-5.262A1 1 0 0 0 13 5.464V3.1a1 1 0 0 0-2 0v2.364a1 1 0 0 0 .021.106a5.406 5.406 0 0 0-4.154 5.262v1.8C6.867 15.018 5 15.614 5 16.807C5 17.4 5 18 5.538 18h12.924C19 18 19 17.4 19 16.807c0-1.193-1.867-1.789-1.867-4.175M6 6a1 1 0 0 1-.707-.293l-1-1a1 1 0 0 1 1.414-1.414l1 1A1 1 0 0 1 6 6m-2 4H3a1 1 0 0 1 0-2h1a1 1 0 1 1 0 2m14-4a1 1 0 0 1-.707-1.707l1-1a1 1 0 1 1 1.414 1.414l-1 1A1 1 0 0 1 18 6m3 4h-1a1 1 0 1 1 0-2h1a1 1 0 1 1 0 2M8.823 19a3.453 3.453 0 0 0 6.354 0z"/>`), g.Group(children),
	)
}

func BellSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.133 12.632v-1.8a5.406 5.406 0 0 0-4.154-5.262A1 1 0 0 0 13 5.464V3.1a1 1 0 0 0-2 0v2.364a1 1 0 0 0 .021.106a5.406 5.406 0 0 0-4.154 5.262v1.8C6.867 15.018 5 15.614 5 16.807C5 17.4 5 18 5.538 18h12.924C19 18 19 17.4 19 16.807c0-1.193-1.867-1.789-1.867-4.175M8.823 19a3.453 3.453 0 0 0 6.354 0z"/>`), g.Group(children),
	)
}

func BlenderPhoneOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20.283 8h-4.285m3.85 3h-3.85m4.061-6H11v11h8.27l1.715-9.847A.983.983 0 0 0 20.059 5M6.581 13.23h-.838A14 14 0 0 1 5.622 11q-.03-1.119.12-2.23h1.04c.252 0 .496-.088.683-.245a.93.93 0 0 0 .329-.61l.2-1.872a.9.9 0 0 0-.045-.39a.9.9 0 0 0-.212-.34a1 1 0 0 0-.341-.231A1.1 1.1 0 0 0 6.983 5h-2.06a1.27 1.27 0 0 0-.699.204a1.14 1.14 0 0 0-.442.543A15.1 15.1 0 0 0 3.007 11a15.7 15.7 0 0 0 .795 5.229c.165.462 1.342.771 1.864.771h1.116c.142 0 .283-.028.413-.082q.197-.081.341-.23a.9.9 0 0 0 .212-.34a.9.9 0 0 0 .046-.391l-.201-1.873a.93.93 0 0 0-.33-.609a1.06 1.06 0 0 0-.682-.245M10 18v1h10v-1a2 2 0 0 0-2-2h-6a2 2 0 0 0-2 2"/>`), g.Group(children),
	)
}

func BlenderPhoneSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 4a1 1 0 0 0-1 1v10h10.459l.522-3H16a1 1 0 1 1 0-2h5.33l.174-1H16a1 1 0 1 1 0-2h5.852l.117-.67v-.003A1.983 1.983 0 0 0 20.06 4zM9 18c0-.35.06-.687.17-1h11.66c.11.313.17.65.17 1v1a1 1 0 0 1-1 1H10a1 1 0 0 1-1-1zm-6.991-7a17.8 17.8 0 0 0 .953 6.1c.198.54 1.61.9 2.237.9h1.34c.17 0 .339-.032.495-.095a1.2 1.2 0 0 0 .41-.27c.114-.114.2-.25.254-.396a1 1 0 0 0 .055-.456l-.242-2.185a1.07 1.07 0 0 0-.395-.71a1.3 1.3 0 0 0-.819-.286H5.291Q5.11 12.306 5.146 11q-.036-1.307.145-2.602H6.54c.302 0 .594-.102.818-.286a1.07 1.07 0 0 0 .396-.71l.24-2.185a1 1 0 0 0-.054-.456a1.1 1.1 0 0 0-.254-.397a1.2 1.2 0 0 0-.41-.269A1.3 1.3 0 0 0 6.78 4H4.307c-.3-.001-.592.082-.838.238a1.34 1.34 0 0 0-.531.634A17.1 17.1 0 0 0 2.008 11Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BookOpenOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6.03v13m0-13c-2.819-.831-4.715-1.076-8.029-1.023A.99.99 0 0 0 3 6v11c0 .563.466 1.014 1.03 1.007c3.122-.043 5.018.212 7.97 1.023m0-13c2.819-.831 4.715-1.076 8.029-1.023A.99.99 0 0 1 21 6v11c0 .563-.466 1.014-1.03 1.007c-3.122-.043-5.018.212-7.97 1.023"/>`), g.Group(children),
	)
}

func BookOpenSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 4.717c-2.286-.58-4.16-.756-7.045-.71A1.99 1.99 0 0 0 2 6v11c0 1.133.934 2.022 2.044 2.007c2.759-.038 4.5.16 6.956.791zm2 15.081c2.456-.631 4.198-.829 6.956-.791A2.013 2.013 0 0 0 22 16.999V6a1.99 1.99 0 0 0-1.955-1.993c-2.885-.046-4.76.13-7.045.71z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BookOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 19V4a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v13H7a2 2 0 0 0-2 2m0 0a2 2 0 0 0 2 2h12M9 3v14m7 0v4"/>`), g.Group(children),
	)
}

func BookSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 2a2 2 0 0 0-2 2v15a3 3 0 0 0 3 3h12a1 1 0 1 0 0-2h-2v-2h2a1 1 0 0 0 1-1V4a2 2 0 0 0-2-2h-8v16h5v2H7a1 1 0 1 1 0-2h1V2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BookmarkOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 21l-5-4l-5 4V3.889a.92.92 0 0 1 .244-.629a.8.8 0 0 1 .59-.26h8.333a.8.8 0 0 1 .589.26a.92.92 0 0 1 .244.63z"/>`), g.Group(children),
	)
}

func BookmarkSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.833 2c-.507 0-.98.216-1.318.576A1.92 1.92 0 0 0 6 3.89V21a1 1 0 0 0 1.625.78L12 18.28l4.375 3.5A1 1 0 0 0 18 21V3.889c0-.481-.178-.954-.515-1.313A1.8 1.8 0 0 0 16.167 2z"/>`), g.Group(children),
	)
}

func BrainOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18.5A2.493 2.493 0 0 1 7.51 20H7.5a2.468 2.468 0 0 1-2.4-3.154a2.98 2.98 0 0 1-.85-5.274a2.47 2.47 0 0 1 .92-3.182a2.477 2.477 0 0 1 1.876-3.344a2.5 2.5 0 0 1 3.41-1.856A2.5 2.5 0 0 1 12 5.5m0 13v-13m0 13a2.493 2.493 0 0 0 4.49 1.5h.01a2.468 2.468 0 0 0 2.403-3.154a2.98 2.98 0 0 0 .847-5.274a2.47 2.47 0 0 0-.921-3.182a2.477 2.477 0 0 0-1.875-3.344A2.5 2.5 0 0 0 14.5 3A2.5 2.5 0 0 0 12 5.5m-8 5a2.5 2.5 0 0 1 3.48-2.3m-.28 8.551a3 3 0 0 1-2.953-5.185M20 10.5a2.5 2.5 0 0 0-3.481-2.3m.28 8.551a3 3 0 0 0 2.954-5.185"/>`), g.Group(children),
	)
}

func BrainSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 21V2.352A3.45 3.45 0 0 0 9.5 2a3.5 3.5 0 0 0-3.261 2.238A3.5 3.5 0 0 0 4.04 8.015a3.5 3.5 0 0 0-.766 1.128c-.042.1-.064.209-.1.313a3 3 0 0 0-.106.344a3.5 3.5 0 0 0 .02 1.468A4 4 0 0 0 2.3 12.5l-.015.036a4 4 0 0 0-.216.779A4 4 0 0 0 2 14q.005.361.072.716a4 4 0 0 0 .235.832l.021.041a4 4 0 0 0 .417.727q.158.22.342.415q.109.113.225.216q.15.137.315.26c.11.081.2.14.308.2l.059.04v.053a3.506 3.506 0 0 0 3.03 3.469a3.426 3.426 0 0 0 4.154.577A.97.97 0 0 1 11 21m10.934-7.68a4 4 0 0 0-.215-.779l-.017-.038a4 4 0 0 0-.79-1.235a3.4 3.4 0 0 0 .017-1.468a3 3 0 0 0-.1-.333c-.034-.108-.057-.22-.1-.324a3.5 3.5 0 0 0-.766-1.128a3.5 3.5 0 0 0-2.202-3.777A3.5 3.5 0 0 0 14.5 2a3.45 3.45 0 0 0-1.5.352V21a.97.97 0 0 1-.184.546a3.426 3.426 0 0 0 4.154-.577A3.506 3.506 0 0 0 20 17.5v-.049l.059-.04q.159-.096.308-.2c.149-.104.214-.169.315-.26q.116-.104.225-.216a4 4 0 0 0 .459-.588q.173-.264.3-.554l.021-.041q.131-.32.205-.659q.019-.086.035-.173q.069-.356.073-.72a4 4 0 0 0-.066-.68"/>`), g.Group(children),
	)
}

func BriefcaseOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7H5a2 2 0 0 0-2 2v4m5-6h8M8 7V5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2v2m0 0h3a2 2 0 0 1 2 2v4m0 0v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-6m18 0s-4 2-9 2s-9-2-9-2m9-2h.01"/>`), g.Group(children),
	)
}

func BriefcaseSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 2a3 3 0 0 0-3 3v1H5a3 3 0 0 0-3 3v2.382l1.447.723l.005.003l.027.013l.12.056q.163.077.486.212c.429.177 1.056.416 1.834.655C7.481 13.524 9.63 14 12 14c2.372 0 4.52-.475 6.08-.956c.78-.24 1.406-.478 1.835-.655a14 14 0 0 0 .606-.268l.027-.013l.005-.002L22 11.381V9a3 3 0 0 0-3-3h-2V5a3 3 0 0 0-3-3zm5 4V5a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v1zm6.447 7.894l.553-.276V19a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3v-5.382l.553.276l.002.002l.004.002l.013.006l.041.02l.151.07c.13.06.318.144.557.242c.478.198 1.163.46 2.01.72C7.019 15.476 9.37 16 12 16c2.628 0 4.98-.525 6.67-1.044a23 23 0 0 0 2.01-.72a16 16 0 0 0 .707-.312l.041-.02l.013-.006l.004-.002zl-.431-.866zM12 10a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BugOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 5L9 4V3m5 2l1-1V3m-3 6v11m0-11a5 5 0 0 1 5 5m-5-5a5 5 0 0 0-5 5m5-5a4.96 4.96 0 0 1 2.973 1H15V8a3 3 0 0 0-6 0v2h.027A4.96 4.96 0 0 1 12 9m-5 5H5m2 0v2a5 5 0 0 0 10 0v-2m2.025 0H17m-9.975 4H6a1 1 0 0 0-1 1v2m12-3h1.025a1 1 0 0 1 1 1v2M16 11h1a1 1 0 0 0 1-1V8m-9.975 3H7a1 1 0 0 1-1-1V8"/>`), g.Group(children),
	)
}

func BugSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 17h-.09q.087-.496.09-1v-1h1a1 1 0 0 0 0-2h-1.09a6 6 0 0 0-.26-1H17a2 2 0 0 0 2-2V8a1 1 0 1 0-2 0v2h-.54a6 6 0 0 0-.46-.46V8a3.96 3.96 0 0 0-.986-2.6l.693-.693A1 1 0 0 0 16 4V3a1 1 0 1 0-2 0v.586l-.661.661a3.75 3.75 0 0 0-2.678 0L10 3.586V3a1 1 0 1 0-2 0v1a1 1 0 0 0 .293.707l.693.693A3.96 3.96 0 0 0 8 8v1.54a6 6 0 0 0-.46.46H7V8a1 1 0 0 0-2 0v2a2 2 0 0 0 2 2h-.65a6 6 0 0 0-.26 1H5a1 1 0 0 0 0 2h1v1a6 6 0 0 0 .09 1H6a2 2 0 0 0-2 2v2a1 1 0 1 0 2 0v-2h.812A6.01 6.01 0 0 0 11 21.907V12a1 1 0 0 1 2 0v9.907A6.01 6.01 0 0 0 17.188 19H18v2a1 1 0 0 0 2 0v-2a2 2 0 0 0-2-2m-4-8.65a6 6 0 0 0-.941-.251l-.111-.017a5.5 5.5 0 0 0-1.9 0l-.111.017A6 6 0 0 0 10 8.35V8a2 2 0 1 1 4 0z"/>`), g.Group(children),
	)
}

func BuildingOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 4h12M6 4v16M6 4H5m13 0v16m0-16h1m-1 16H6m12 0h1M6 20H5M9 7h1v1H9zm5 0h1v1h-1zm-5 4h1v1H9zm5 0h1v1h-1zm-3 4h2a1 1 0 0 1 1 1v4h-4v-4a1 1 0 0 1 1-1"/>`), g.Group(children),
	)
}

func BuildingSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4a1 1 0 0 1 1-1h14a1 1 0 1 1 0 2v14a1 1 0 1 1 0 2H5a1 1 0 1 1 0-2V5a1 1 0 0 1-1-1m5 2a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1zm5 0a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1zm-5 4a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1zm5 0a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1zm-3 4a2 2 0 0 0-2 2v3h2v-3h2v3h2v-3a2 2 0 0 0-2-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BullhornOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 9H5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h6m0-6v6m0-6l5.419-3.87A1 1 0 0 1 18 5.942v12.114a1 1 0 0 1-1.581.814L11 15m7 0a3 3 0 0 0 0-6M6 15h3v5H6z"/>`), g.Group(children),
	)
}

func BullhornSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18.458 3.11A1 1 0 0 1 19 4v16a1 1 0 0 1-1.581.814L12 16.944V7.056l5.419-3.87a1 1 0 0 1 1.039-.076M22 12c0 1.48-.804 2.773-2 3.465v-6.93c1.196.692 2 1.984 2 3.465M10 8H4a1 1 0 0 0-1 1v6a1 1 0 0 0 1 1h6zm0 9H5v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CalendarEditOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11.5 11.5l2.071 1.994M4 10h5m11 0h-1.5M12 7V4M7 7V4m10 3V4m-7 13H8v-2l5.227-5.292a1.46 1.46 0 0 1 2.065 2.065zm-5 3h14a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1"/>`), g.Group(children),
	)
}

func CalendarEditSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12.512 8.72a2.46 2.46 0 0 1 3.479 0a2.46 2.46 0 0 1 0 3.479l-.004.005l-1.094 1.08a1 1 0 0 0-.194-.272l-3-3a1 1 0 0 0-.272-.193zm-2.415 2.445L7.28 14.017a1 1 0 0 0-.289.702v2a1 1 0 0 0 1 1h2a1 1 0 0 0 .703-.288l2.851-2.816a1 1 0 0 1-.26-.189l-3-3a1 1 0 0 1-.19-.26Z"/><path d="M7 3a1 1 0 0 1 1 1v1h3V4a1 1 0 1 1 2 0v1h3V4a1 1 0 1 1 2 0v1h1a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h1V4a1 1 0 0 1 1-1m10.67 8H19v8H5v-8h3.855l.53-.537a1 1 0 0 1 .87-.285c.097.015.233.13.277.087s-.073-.18-.09-.276a1 1 0 0 1 .274-.873l1.09-1.104a3.46 3.46 0 0 1 4.892 0l.001.002A3.46 3.46 0 0 1 17.67 11"/></g>`), g.Group(children),
	)
}

func CalendarMonthOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 10h16m-8-3V4M7 7V4m10 3V4M5 20h14a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1m3-7h.01v.01H8zm4 0h.01v.01H12zm4 0h.01v.01H16zm-8 4h.01v.01H8zm4 0h.01v.01H12zm4 0h.01v.01H16z"/>`), g.Group(children),
	)
}

func CalendarMonthSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 5a1 1 0 0 0 1-1a1 1 0 1 1 2 0a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1a1 1 0 1 1 2 0a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1a1 1 0 1 1 2 0a1 1 0 0 0 1 1a2 2 0 0 1 2 2v1a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V7a2 2 0 0 1 2-2M3 19v-7a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m6.01-6a1 1 0 1 0-2 0a1 1 0 0 0 2 0m2 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0m6 0a1 1 0 1 0-2 0a1 1 0 0 0 2 0m-10 4a1 1 0 1 1 2 0a1 1 0 0 1-2 0m6 0a1 1 0 1 0-2 0a1 1 0 0 0 2 0m2 0a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CalendarPlusOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 9.05H3v2h1zm16 2h1v-2h-1zM10 14a1 1 0 1 0 0 2zm4 2a1 1 0 1 0 0-2zm-3 1a1 1 0 1 0 2 0zm2-4a1 1 0 1 0-2 0zm-2-5.95a1 1 0 1 0 2 0zm2-3a1 1 0 1 0-2 0zm-7 3a1 1 0 0 0 2 0zm2-3a1 1 0 1 0-2 0zm8 3a1 1 0 1 0 2 0zm2-3a1 1 0 1 0-2 0zm-13 3h14v-2H5zm14 0v12h2v-12zm0 12H5v2h14zm-14 0v-12H3v12zm0 0H3a2 2 0 0 0 2 2zm14 0v2a2 2 0 0 0 2-2zm0-12h2a2 2 0 0 0-2-2zm-14-2a2 2 0 0 0-2 2h2zm-1 6h16v-2H4zM10 16h4v-2h-4zm3 1v-4h-2v4zm0-9.95v-3h-2v3zm-5 0v-3H6v3zm10 0v-3h-2v3z"/>`), g.Group(children),
	)
}

func CalendarPlusSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 5.05h1a2 2 0 0 1 2 2v2H3v-2a2 2 0 0 1 2-2h1v-1a1 1 0 1 1 2 0v1h3v-1a1 1 0 1 1 2 0v1h3v-1a1 1 0 1 1 2 0zm-15 6v8a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-8zM11 18a1 1 0 1 0 2 0v-1h1a1 1 0 1 0 0-2h-1v-1a1 1 0 1 0-2 0v1h-1a1 1 0 1 0 0 2h1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CalendarWeekOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 10h16M8 14h8m-4-7V4M7 7V4m10 3V4M5 20h14a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1"/>`), g.Group(children),
	)
}

func CalendarWeekSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 5V4a1 1 0 1 1 2 0v1h3V4a1 1 0 1 1 2 0v1h3V4a1 1 0 1 1 2 0v1h1a2 2 0 0 1 2 2v2H3V7a2 2 0 0 1 2-2zM3 19v-8h18v8a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m5-6a1 1 0 1 0 0 2h8a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CameraFotoOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M4 18V8c0-.6.4-1 1-1h1.5l1.7-1.7c.2-.2.4-.3.7-.3h6.2c.3 0 .5.1.7.3L17.5 7H19c.6 0 1 .4 1 1v10c0 .6-.4 1-1 1H5a1 1 0 0 1-1-1Z"/><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g>`), g.Group(children),
	)
}

func CameraFotoSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 4.6A2 2 0 0 1 8.9 4h6.2c.5 0 1 .2 1.4.6L17.9 6H19a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8c0-1.1.9-2 2-2h1zM10 12a2 2 0 1 1 4 0a2 2 0 0 1-4 0m2-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CameraPhotoOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M4 18V8a1 1 0 0 1 1-1h1.5l1.707-1.707A1 1 0 0 1 8.914 5h6.172a1 1 0 0 1 .707.293L17.5 7H19a1 1 0 0 1 1 1v10a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1Z"/><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g>`), g.Group(children),
	)
}

func CameraPhotoSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.5 4.586A2 2 0 0 1 8.914 4h6.172a2 2 0 0 1 1.414.586L17.914 6H19a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h1.086zM10 12a2 2 0 1 1 4 0a2 2 0 0 1-4 0m2-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CaptionOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.855 14.322a2.475 2.475 0 1 1 .133-4.241m6.053 4.241a2.475 2.475 0 1 1 .133-4.241M4 5h16a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1"/>`), g.Group(children),
	)
}

func CaptionSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm6.962 4.856a1.48 1.48 0 0 1 1.484.066A1 1 0 1 0 11.53 9.24a3.475 3.475 0 1 0-.187 5.955a1 1 0 1 0-.976-1.746a1.474 1.474 0 1 1-1.405-2.593m6.186 0a1.48 1.48 0 0 1 1.484.066a1 1 0 1 0 1.084-1.682a3.475 3.475 0 1 0-.187 5.955a1 1 0 1 0-.976-1.746a1.474 1.474 0 1 1-1.405-2.593" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CaptioningOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.9 14.3a2.5 2.5 0 1 1 0-4.2m6.1 4.2a2.5 2.5 0 1 1 .2-4.2M4 5h16c.6 0 1 .4 1 1v12c0 .6-.4 1-1 1H4a1 1 0 0 1-1-1V6c0-.6.4-1 1-1"/>`), g.Group(children),
	)
}

func CaptioningSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 6c0-1.1.9-2 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm7 4.9a1.5 1.5 0 0 1 1.4 0a1 1 0 1 0 1.1-1.7a3.5 3.5 0 1 0-.2 6a1 1 0 1 0-1-1.8A1.5 1.5 0 1 1 9 11Zm6.1 0a1.5 1.5 0 0 1 1.5 0a1 1 0 1 0 1.1-1.7a3.5 3.5 0 1 0-.2 6a1 1 0 1 0-1-1.8a1.5 1.5 0 1 1-1.4-2.5" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CaretDownOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.119 8h9.762a1 1 0 0 1 .772 1.636l-4.881 5.927a1 1 0 0 1-1.544 0l-4.88-5.927A1 1 0 0 1 7.118 8Z"/>`), g.Group(children),
	)
}

func CaretDownSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18.425 10.271C19.499 8.967 18.57 7 16.88 7H7.12c-1.69 0-2.618 1.967-1.544 3.271l4.881 5.927a2 2 0 0 0 3.088 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CaretLeftOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 16.881V7.119a1 1 0 0 0-1.636-.772l-5.927 4.881a1 1 0 0 0 0 1.544l5.927 4.88a1 1 0 0 0 1.636-.77Z"/>`), g.Group(children),
	)
}

func CaretLeftSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.729 5.575c1.304-1.074 3.27-.146 3.27 1.544v9.762c0 1.69-1.966 2.618-3.27 1.544l-5.927-4.881a2 2 0 0 1 0-3.088z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CaretRightOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16.881V7.119a1 1 0 0 1 1.636-.772l5.927 4.881a1 1 0 0 1 0 1.544l-5.927 4.88A1 1 0 0 1 8 16.882Z"/>`), g.Group(children),
	)
}

func CaretRightSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.271 5.575C8.967 4.501 7 5.43 7 7.12v9.762c0 1.69 1.967 2.618 3.271 1.544l5.927-4.881a2 2 0 0 0 0-3.088l-5.927-4.88Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CaretSortOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 10l4-6l4 6zm8 4l-4 6l-4-6z"/>`), g.Group(children),
	)
}

func CaretSortSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.832 3.445a1 1 0 0 0-1.664 0l-4 6A1 1 0 0 0 8 11h8a1 1 0 0 0 .832-1.555zm-1.664 17.11a1 1 0 0 0 1.664 0l4-6A1 1 0 0 0 16 13H8a1 1 0 0 0-.832 1.555z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CaretUpOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.881 16H7.119a1 1 0 0 1-.772-1.636l4.881-5.927a1 1 0 0 1 1.544 0l4.88 5.927a1 1 0 0 1-.77 1.636Z"/>`), g.Group(children),
	)
}

func CaretUpSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.575 13.729C4.501 15.033 5.43 17 7.12 17h9.762c1.69 0 2.618-1.967 1.544-3.271l-4.881-5.927a2 2 0 0 0-3.088 0l-4.88 5.927Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CartOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4h1.5L9 16m0 0h8m-8 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4m8 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-8.5-3h9.25L19 7H7.312"/>`), g.Group(children),
	)
}

func CartPlusAltOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h1.5L8 16m0 0h8m-8 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4m8 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4m.75-3H7.5M11 7H6.312M17 4v6m-3-3h6"/>`), g.Group(children),
	)
}

func CartPlusAltSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M12.268 6A2 2 0 0 0 14 9h1v1a2 2 0 0 0 3.04 1.708l-.311 1.496a1 1 0 0 1-.979.796H8.605l.208 1H16a3 3 0 1 1-2.83 2h-2.34a3 3 0 1 1-4.009-1.76L4.686 5H4a1 1 0 0 1 0-2h1.5a1 1 0 0 1 .979.796L6.939 6z"/><path d="M18 4a1 1 0 1 0-2 0v2h-2a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0V8h2a1 1 0 1 0 0-2h-2z"/></g>`), g.Group(children),
	)
}

func CartPlusOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 4h1.5L9 16m0 0h8m-8 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4m8 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-8.5-3h9.25L19 7h-1M8 7h-.688M13 5v4m-2-2h4"/>`), g.Group(children),
	)
}

func CartPlusSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5 3a1 1 0 0 0 0 2h.687L7.82 15.24A3 3 0 1 0 11.83 17h2.34A3 3 0 1 0 17 15H9.813l-.208-1h8.145a1 1 0 0 0 .979-.796l1.25-6A1 1 0 0 0 19 6h-2.268A2 2 0 0 1 15 9a2 2 0 1 1-4 0a2 2 0 0 1-1.732-3h-1.33L7.48 3.796A1 1 0 0 0 6.5 3z"/><path d="M14 5a1 1 0 1 0-2 0v1h-1a1 1 0 1 0 0 2h1v1a1 1 0 1 0 2 0V8h1a1 1 0 1 0 0-2h-1z"/></g>`), g.Group(children),
	)
}

func CartSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4a1 1 0 0 1 1-1h1.5a1 1 0 0 1 .979.796L7.939 6H19a1 1 0 0 1 .979 1.204l-1.25 6a1 1 0 0 1-.979.796H9.605l.208 1H17a3 3 0 1 1-2.83 2h-2.34a3 3 0 1 1-4.009-1.76L5.686 5H5a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CashOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M8 7V6a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-1M3 18v-7a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1Zm8-3.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0Z"/>`), g.Group(children),
	)
}

func CashSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M7 6a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2h-2v-4a3 3 0 0 0-3-3H7z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M2 11a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm7.5 1a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5" clip-rule="evenodd"/><path d="M10.5 14.5a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/></g>`), g.Group(children),
	)
}

func ChartLineDownOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.5V19a1 1 0 0 0 1 1h15M7 10l4 4l4-4l5 5m0 0h-3.207M20 15v-3.207"/>`), g.Group(children),
	)
}

func ChartLineUpOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.5V19a1 1 0 0 0 1 1h15M7 14l4-4l4 4l5-5m0 0h-3.207M20 9v3.207"/>`), g.Group(children),
	)
}

func ChartLineUpSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4.5V19c0 .6.4 1 1 1h15"/><path d="m7 10l4 4l4-4l5 5m0 0h-3.2m3.2 0v-3.2"/></g>`), g.Group(children),
	)
}

func ChartMixedDollarOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.6 16.733c.234.269.548.456.895.534a1.4 1.4 0 0 0 1.75-.762c.172-.615-.446-1.287-1.242-1.481s-1.41-.861-1.241-1.481a1.4 1.4 0 0 1 1.75-.762c.343.077.654.26.888.524m-1.358 4.017v.617m0-5.939v.725M4 15v4m3-6v6M6 8.5L10.5 5L14 7.5L18 4m0 0h-3.5M18 4v3m2 8a5 5 0 1 1-10 0a5 5 0 0 1 10 0"/>`), g.Group(children),
	)
}

func ChartMixedDollarSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 15a6 6 0 1 1 12 0a6 6 0 0 1-12 0m3.845-1.855a2.4 2.4 0 0 1 1.2-1.226a1 1 0 0 1 1.992-.026c.426.15.809.408 1.111.749a1 1 0 1 1-1.496 1.327a.7.7 0 0 0-.36-.213a1 1 0 0 1-.113-.032a.4.4 0 0 0-.394.074a.93.93 0 0 0 .455.254a2.9 2.9 0 0 1 1.504.9c.373.433.669 1.092.464 1.823a1 1 0 0 1-.046.129c-.226.519-.627.94-1.132 1.192a1 1 0 0 1-1.956.093a2.7 2.7 0 0 1-1.227-.798a1 1 0 1 1 1.506-1.315a.7.7 0 0 0 .363.216q.057.014.111.032a.4.4 0 0 0 .395-.074a.93.93 0 0 0-.455-.254a2.9 2.9 0 0 1-1.503-.9c-.375-.433-.666-1.089-.466-1.817a1 1 0 0 1 .047-.134m1.884.573l.003.008zm.55 2.613s-.002-.002-.003-.007zM4 14a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0v-4a1 1 0 0 1 1-1m3-2a1 1 0 0 1 1 1v6a1 1 0 1 1-2 0v-6a1 1 0 0 1 1-1m6.5-8a1 1 0 0 1 1-1H18a1 1 0 0 1 1 1v3a1 1 0 1 1-2 0v-.796l-2.341 2.049a1 1 0 0 1-1.24.06l-2.894-2.066L6.614 9.29a1 1 0 1 1-1.228-1.578l4.5-3.5a1 1 0 0 1 1.195-.025l2.856 2.04L15.34 5h-.84a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ChartMixedOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 15v4m6-6v6m6-4v4m6-6v6M3 11l6-5l6 5l5.5-5.5"/>`), g.Group(children),
	)
}

func ChartOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4v15a1 1 0 0 0 1 1h15M8 16l2.5-5.5l3 3L17.273 7L20 9.667"/>`), g.Group(children),
	)
}

func ChartPieOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M10 6.025A7.5 7.5 0 1 0 17.975 14H10z"/><path d="M13.5 3c-.169 0-.334.014-.5.025V11h7.975c.011-.166.025-.331.025-.5A7.5 7.5 0 0 0 13.5 3"/></g>`), g.Group(children),
	)
}

func ChartPieSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M13.5 2c-.178 0-.356.013-.492.022l-.074.005a1 1 0 0 0-.934.998V11a1 1 0 0 0 1 1h7.975a1 1 0 0 0 .998-.934l.005-.074A7 7 0 0 0 22 10.5A8.5 8.5 0 0 0 13.5 2"/><path d="M11 6.025a1 1 0 0 0-1.065-.998a8.5 8.5 0 1 0 9.038 9.039A1 1 0 0 0 17.975 13H11z"/></g>`), g.Group(children),
	)
}

func ChartSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M4 4v15c0 .6.4 1 1 1h15"/><path stroke-linejoin="round" d="m8 16l2.5-5.5l3 3L17.3 7L20 9.7"/></g>`), g.Group(children),
	)
}

func CheckCircleOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.5 11.5L11 14l4-4m6 2a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`), g.Group(children),
	)
}

func CheckCircleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m13.707-1.293a1 1 0 0 0-1.414-1.414L11 12.586l-1.793-1.793a1 1 0 0 0-1.414 1.414l2.5 2.5a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CheckOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 11.917L9.724 16.5L19 7.5"/>`), g.Group(children),
	)
}

func CheckPlusCircleOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 1 0-18c1.052 0 2.062.18 3 .512M7 9.577l3.923 3.923l8.5-8.5M17 14v6m-3-3h6"/>`), g.Group(children),
	)
}

func CheckPlusCircleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M18 14a1 1 0 1 0-2 0v2h-2a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2z"/><path d="M15.026 21.534A10 10 0 0 1 12 22C6.477 22 2 17.523 2 12S6.477 2 12 2c2.51 0 4.802.924 6.558 2.45l-7.635 7.636L7.707 8.87a1 1 0 0 0-1.414 1.414l3.923 3.923a1 1 0 0 0 1.414 0l8.3-8.3A9.96 9.96 0 0 1 22 12a10 10 0 0 1-.466 3.026A2.5 2.5 0 0 0 20 14.5h-.5V14a2.5 2.5 0 0 0-5 0v.5H14a2.5 2.5 0 0 0 0 5h.5v.5c0 .578.196 1.11.526 1.534"/></g>`), g.Group(children),
	)
}

func CheckSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 12l4.7 4.5l9.3-9"/>`), g.Group(children),
	)
}

func ChevronDoubleDownOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 7l4 4l4-4m-8 6l4 4l4-4"/>`), g.Group(children),
	)
}

func ChevronDoubleLeftOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 16l-4-4l4-4m-6 8l-4-4l4-4"/>`), g.Group(children),
	)
}

func ChevronDoubleRightOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 16l4-4l-4-4m6 8l4-4l-4-4"/>`), g.Group(children),
	)
}

func ChevronDoubleUpOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 17l-4-4l-4 4m8-6l-4-4l-4 4"/>`), g.Group(children),
	)
}

func ChevronDownOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 10l4 4l4-4"/>`), g.Group(children),
	)
}

func ChevronLeftOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 8l-4 4l4 4"/>`), g.Group(children),
	)
}

func ChevronRightOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 16l4-4l-4-4"/>`), g.Group(children),
	)
}

func ChevronSortOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 15l4 4l4-4m0-6l-4-4l-4 4"/>`), g.Group(children),
	)
}

func ChevronUpOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m16 14l-4-4l-4 4"/>`), g.Group(children),
	)
}

func CircleMinusOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.757 12h8.486M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`), g.Group(children),
	)
}

func CircleMinusSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m5.757-1a1 1 0 1 0 0 2h8.486a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CirclePauseOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 9v6m4-6v6m7-3a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`), g.Group(children),
	)
}

func CirclePauseSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m9-3a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0zm4 0a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CirclePlusOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7.757v8.486M7.757 12h8.486M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`), g.Group(children),
	)
}

func CirclePlusSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m11-4.243a1 1 0 1 0-2 0V11H7.757a1 1 0 1 0 0 2H11v3.243a1 1 0 1 0 2 0V13h3.243a1 1 0 1 0 0-2H13z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ClapperboardPlayOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 4H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1m0 0l-4 4m5 0H4m1 0l4-4m1 4l4-4m-4 7v6l4-3z"/>`), g.Group(children),
	)
}

func ClapperboardPlaySolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19.003 3A2 2 0 0 1 21 5v2h-2V5.414L17.414 7h-2.828l2-2h-2.172l-2 2H9.586l2-2H9.414l-2 2H3V5a2 2 0 0 1 2-2zM3 9v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V9zm2-2.414L6.586 5H5zm4.553 4.52a1 1 0 0 1 1.047.094l4 3a1 1 0 0 1 0 1.6l-4 3A1 1 0 0 1 9 18v-6a1 1 0 0 1 .553-.894" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ClipboardCheckOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 4h3a1 1 0 0 1 1 1v15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3m0 3h6m-6 7l2 2l4-4m-5-9v4h4V3z"/>`), g.Group(children),
	)
}

func ClipboardCheckSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2a1 1 0 0 0-1 1H6a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-2a1 1 0 0 0-1-1zm1 2h4v2h1a1 1 0 1 1 0 2H9a1 1 0 0 1 0-2h1zm5.707 8.707a1 1 0 0 0-1.414-1.414L11 14.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ClipboardCleanOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 4h3a1 1 0 0 1 1 1v15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3m0 3h6m-5-4v4h4V3z"/>`), g.Group(children),
	)
}

func ClipboardCleanSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 3a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1h2a2 2 0 0 1 2 2v15a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2zm6 1h-4v2H9a1 1 0 0 0 0 2h6a1 1 0 1 0 0-2h-1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ClipboardListOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 4h3a1 1 0 0 1 1 1v15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3m0 3h6m-3 5h3m-6 0h.01M12 16h3m-6 0h.01M10 3v4h4V3z"/>`), g.Group(children),
	)
}

func ClipboardListSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 3a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1h2a2 2 0 0 1 2 2v15a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2zm6 1h-4v2H9a1 1 0 0 0 0 2h6a1 1 0 1 0 0-2h-1zm-3 8a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1m-2-1a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2zm2 5a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1m-2-1a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ClipboardOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 4h3a1 1 0 0 1 1 1v15a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h3m0 3h6m-6 5h6m-6 4h6M10 3v4h4V3z"/>`), g.Group(children),
	)
}

func ClipboardSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 3a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1h2a2 2 0 0 1 2 2v15a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2zm6 1h-4v2H9a1 1 0 0 0 0 2h6a1 1 0 1 0 0-2h-1zm-6 8a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1m1 3a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ClockOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 8v4l3 3m6-3a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`), g.Group(children),
	)
}

func ClockSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m11-4a1 1 0 1 0-2 0v4a1 1 0 0 0 .293.707l3 3a1 1 0 0 0 1.414-1.414L13 11.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloseCircleOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 9l-6 6m0-6l6 6m6-3a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`), g.Group(children),
	)
}

func CloseCircleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7.707-3.707a1 1 0 0 0-1.414 1.414L10.586 12l-2.293 2.293a1 1 0 1 0 1.414 1.414L12 13.414l2.293 2.293a1 1 0 0 0 1.414-1.414L13.414 12l2.293-2.293a1 1 0 0 0-1.414-1.414L12 10.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloseOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L17.94 6M18 18L6.06 6"/>`), g.Group(children),
	)
}

func CloseSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 18L18 6m0 12L6 6"/>`), g.Group(children),
	)
}

func CloudArrowUpOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 17h3a3 3 0 0 0 0-6h-.025a6 6 0 0 0 .025-.5A5.5 5.5 0 0 0 7.207 9.021C7.137 9.017 7.071 9 7 9a4 4 0 1 0 0 8h2.167M12 19v-9m0 0l-2 2m2-2l2 2"/>`), g.Group(children),
	)
}

func CloudArrowUpSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M13.383 4.076a6.5 6.5 0 0 0-6.887 3.95A5 5 0 0 0 7 18h3v-4a2 2 0 0 1-1.414-3.414l2-2a2 2 0 0 1 2.828 0l2 2A2 2 0 0 1 14 14v4h4a4 4 0 0 0 .988-7.876a6.5 6.5 0 0 0-5.605-6.048"/><path d="M12.707 9.293a1 1 0 0 0-1.414 0l-2 2a1 1 0 1 0 1.414 1.414l.293-.293V19a1 1 0 1 0 2 0v-6.586l.293.293a1 1 0 0 0 1.414-1.414z"/></g>`), g.Group(children),
	)
}

func CodeBranchOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8v8m0-8a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4m8-8a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0 0a4 4 0 0 1-4 4h-1a3 3 0 0 0-3 3"/>`), g.Group(children),
	)
}

func CodeBranchSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 3a3 3 0 0 0-1 5.83v6.34a3.001 3.001 0 1 0 2 0V15a2 2 0 0 1 2-2h1a5 5 0 0 0 4.927-4.146A3.001 3.001 0 0 0 16 3a3 3 0 0 0-1.105 5.79A3 3 0 0 1 12 11h-1c-.729 0-1.412.195-2 .535V8.83A3.001 3.001 0 0 0 8 3"/>`), g.Group(children),
	)
}

func CodeForkOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12v4m0 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4M8 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0 0v2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V8m0 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`), g.Group(children),
	)
}

func CodeForkSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 6a3 3 0 1 1 4 2.83V10a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1V8.83a3.001 3.001 0 1 1 2 0V10a3 3 0 0 1-3 3h-1v2.17a3.001 3.001 0 1 1-2 0V13h-1a3 3 0 0 1-3-3V8.83A3 3 0 0 1 5 6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CodeMergeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8v8m0-8a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4m6-2a2 2 0 1 1 4 0a2 2 0 0 1-4 0m0 0h-1a5 5 0 0 1-5-5v-.5"/>`), g.Group(children),
	)
}

func CodeMergeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 6a3 3 0 1 1 4 2.83V9a4 4 0 0 0 4 4h.17a3.001 3.001 0 1 1 0 2H13a5.98 5.98 0 0 1-4-1.528v1.699a3.001 3.001 0 1 1-2 0V8.829A3 3 0 0 1 5 6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CodeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 8l-4 4l4 4m8 0l4-4l-4-4m-2-3l-4 14"/>`), g.Group(children),
	)
}

func CodePullRequestOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8v8m0-8a2 2 0 1 0 0-4a2 2 0 0 0 0 4m0 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4m12 0a2 2 0 1 0 0 4a2 2 0 0 0 0-4m0 0V9a3 3 0 0 0-3-3h-3m1.5-2l-2 2l2 2"/>`), g.Group(children),
	)
}

func CodePullRequestSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 6a3 3 0 1 1 4 2.83v6.34a3.001 3.001 0 1 1-2 0V8.83A3 3 0 0 1 3 6m11.207-2.707a1 1 0 0 1 0 1.414L13.914 5H15a4 4 0 0 1 4 4v6.17a3.001 3.001 0 1 1-2 0V9a2 2 0 0 0-2-2h-1.086l.293.293a1 1 0 0 1-1.414 1.414l-2-2a1 1 0 0 1 0-1.414l2-2a1 1 0 0 1 1.414 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CodeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="m8 8l-4 4l4 4m8 0l4-4l-4-4"/><path d="m14 5l-4 14"/></g>`), g.Group(children),
	)
}

func CogOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M21 13v-2a1 1 0 0 0-1-1h-.757l-.707-1.707l.535-.536a1 1 0 0 0 0-1.414l-1.414-1.414a1 1 0 0 0-1.414 0l-.536.535L14 4.757V4a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v.757l-1.707.707l-.536-.535a1 1 0 0 0-1.414 0L4.929 6.343a1 1 0 0 0 0 1.414l.536.536L4.757 10H4a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h.757l.707 1.707l-.535.536a1 1 0 0 0 0 1.414l1.414 1.414a1 1 0 0 0 1.414 0l.536-.535l1.707.707V20a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-.757l1.707-.708l.536.536a1 1 0 0 0 1.414 0l1.414-1.414a1 1 0 0 0 0-1.414l-.535-.536l.707-1.707H20a1 1 0 0 0 1-1"/><path d="M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/></g>`), g.Group(children),
	)
}

func CogSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.586 2.586A2 2 0 0 1 11 2h2a2 2 0 0 1 2 2v.089l.473.196l.063-.063a2 2 0 0 1 2.828 0l1.414 1.414a2 2 0 0 1 0 2.827l-.063.064l.196.473H20a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-.089l-.196.473l.063.063a2 2 0 0 1 0 2.828l-1.414 1.414a2 2 0 0 1-2.828 0l-.063-.063l-.473.196V20a2 2 0 0 1-2 2h-2a2 2 0 0 1-2-2v-.089l-.473-.196l-.063.063a2 2 0 0 1-2.828 0l-1.414-1.414a2 2 0 0 1 0-2.827l.063-.064L4.089 15H4a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2h.09l.195-.473l-.063-.063a2 2 0 0 1 0-2.828l1.414-1.414a2 2 0 0 1 2.827 0l.064.063L9 4.089V4a2 2 0 0 1 .586-1.414M8 12a4 4 0 1 1 8 0a4 4 0 0 1-8 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ColumnOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 5v14M9 5v14M4 5h16a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1"/>`), g.Group(children),
	)
}

func ColumnSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 4H9v16h6zm2 16h3a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-3zM4 4h3v16H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CommandOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8v8m0-8h8M8 8H6a2 2 0 1 1 2-2zm0 8h8m-8 0H6a2 2 0 1 0 2 2zm8 0V8m0 8h2a2 2 0 1 1-2 2zm0-8h2a2 2 0 1 0-2-2z"/>`), g.Group(children),
	)
}

func CommandSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 16V8h8v8zm0-8H6a2 2 0 1 1 2-2zm8 0h2a2 2 0 1 0-2-2zm-8 8H6a2 2 0 1 0 2 2zm8 0h2a2 2 0 1 1-2 2z"/>`), g.Group(children),
	)
}

func CompressOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h4V4m12 4h-4V4M4 16h4v4m12-4h-4v4"/>`), g.Group(children),
	)
}

func ComputerSpeakerOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 16H5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v1M9 12H4m8 8V9h8v11zm0 0H9m8-4a1 1 0 1 0-2 0a1 1 0 0 0 2 0"/>`), g.Group(children),
	)
}

func ComputerSpeakerSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12 8a1 1 0 0 0-1 1v10H9a1 1 0 1 0 0 2h11a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1zm4 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/><path d="M5 3a2 2 0 0 0-2 2v6h6V9a3 3 0 0 1 3-3h8c.35 0 .687.06 1 .17V5a2 2 0 0 0-2-2zm4 10H3v2a2 2 0 0 0 2 2h4z"/></g>`), g.Group(children),
	)
}

func CreditCardOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 10h18M6 14h2m3 0h5M3 7v10a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1"/>`), g.Group(children),
	)
}

func CreditCardPlusOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" stroke-width="2" d="M16.5 15v1.5m0 0V18m0-1.5H15m1.5 0H18M3 9V6a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v3M3 9v6a1 1 0 0 0 1 1h5M3 9h16m0 0v1M6 12h3m12 4.5a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0Z"/>`), g.Group(children),
	)
}

func CreditCardPlusSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M11 16.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m4.5 2.5v-1.5H14v-2h1.5V14h2v1.5H19v2h-1.5V19z" clip-rule="evenodd"/><path d="M3.987 4A2 2 0 0 0 2 6v9a2 2 0 0 0 2 2h5v-2H4v-5h16V6a2 2 0 0 0-2-2z"/><path fill-rule="evenodd" d="M5 12a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func CreditCardSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4 5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2zm0 6h16v6H4z"/><path d="M5 14a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1m5 0a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2h-5a1 1 0 0 1-1-1"/></g>`), g.Group(children),
	)
}

func CssSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m3 2l1.578 17.834L12 22l7.468-2.165L21 2zm13.3 14.722l-4.293 1.204H12l-4.297-1.204l-.297-3.167h2.108l.15 1.526l2.335.639l2.34-.64l.245-3.05h-7.27l-.187-2.006h7.64l.174-2.006H6.924l-.176-2.006h10.506z"/>`), g.Group(children),
	)
}

func DatabaseOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 6c0 1.657-3.134 3-7 3S5 7.657 5 6m14 0c0-1.657-3.134-3-7-3S5 4.343 5 6m14 0v6M5 6v6m0 0c0 1.657 3.134 3 7 3s7-1.343 7-3M5 12v6c0 1.657 3.134 3 7 3s7-1.343 7-3v-6"/>`), g.Group(children),
	)
}

func DatabaseSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 7.205c4.418 0 8-1.165 8-2.602C20 3.165 16.418 2 12 2S4 3.165 4 4.603c0 1.437 3.582 2.602 8 2.602M12 22c4.963 0 8-1.686 8-2.603v-4.404c-.052.032-.112.06-.165.09a8 8 0 0 1-.745.387q-.29.132-.6.253q-.093.037-.189.073a19 19 0 0 1-6.3.998c-2.135.027-4.26-.31-6.3-.998l-.189-.073a10 10 0 0 1-.852-.373a8 8 0 0 1-.493-.267c-.053-.03-.113-.058-.165-.09v4.404C4 20.315 7.037 22 12 22m7.09-13.928a10 10 0 0 1-.6.253q-.093.038-.189.074a19 19 0 0 1-6.3.998c-2.135.027-4.26-.31-6.3-.998l-.189-.074a10 10 0 0 1-.852-.372a8 8 0 0 1-.493-.268c-.055-.03-.115-.058-.167-.09V12c0 .917 3.037 2.603 8 2.603s8-1.686 8-2.603V7.596c-.052.031-.112.059-.165.09a8 8 0 0 1-.745.386"/>`), g.Group(children),
	)
}

func DesktopPcOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15v5m-3 0h6M4 11h16M5 15h14a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1"/>`), g.Group(children),
	)
}

func DesktopPcSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v5h18V5a2 2 0 0 0-2-2zM3 14v-2h18v2a2 2 0 0 1-2 2h-6v3h2a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h2v-3H5a2 2 0 0 1-2-2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DiscordSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.942 5.556a16.3 16.3 0 0 0-4.126-1.3a12 12 0 0 0-.529 1.1a15.2 15.2 0 0 0-4.573 0a12 12 0 0 0-.535-1.1a16.3 16.3 0 0 0-4.129 1.3a17.4 17.4 0 0 0-2.868 11.662a15.8 15.8 0 0 0 4.963 2.521q.616-.847 1.084-1.785a10.6 10.6 0 0 1-1.706-.83q.215-.16.418-.331a11.66 11.66 0 0 0 10.118 0q.206.172.418.331q-.817.492-1.71.832a12.6 12.6 0 0 0 1.084 1.785a16.5 16.5 0 0 0 5.064-2.595a17.3 17.3 0 0 0-2.973-11.59M8.678 14.813a1.94 1.94 0 0 1-1.8-2.045a1.93 1.93 0 0 1 1.8-2.047a1.92 1.92 0 0 1 1.8 2.047a1.93 1.93 0 0 1-1.8 2.045m6.644 0a1.94 1.94 0 0 1-1.8-2.045a1.93 1.93 0 0 1 1.8-2.047a1.92 1.92 0 0 1 1.8 2.047a1.93 1.93 0 0 1-1.8 2.045"/>`), g.Group(children),
	)
}

func DnaOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.041 13.862A5 5 0 0 1 17 17.831V21M7 3v3.169a5 5 0 0 0 1.891 3.916M17 3v3.169a5 5 0 0 1-2.428 4.288l-5.144 3.086A5 5 0 0 0 7 17.831V21M7 5h10M7.399 8h9.252M8 16h8.652M7 19h10"/>`), g.Group(children),
	)
}

func DollarOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 17.345a4.76 4.76 0 0 0 2.558 1.618c2.274.589 4.512-.446 4.999-2.31c.487-1.866-1.273-3.9-3.546-4.49S7.977 9.54 8.464 7.675s2.724-2.899 4.998-2.31c.982.236 1.87.793 2.538 1.592m-3.879 12.171V21m0-18v2.2"/>`), g.Group(children),
	)
}

func DollarSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.1 2c.6 0 1 .4 1 1v1.3a6 6 0 0 1 3.7 2a1 1 0 1 1-1.6 1.3c-.5-.6-1.2-1-2-1.3a4.2 4.2 0 0 0-1.1-.1c-1.4 0-2.4.8-2.7 1.7c-.1.6 0 1.2.6 1.9c.5.6 1.3 1.2 2.3 1.4c1.3.3 2.5 1.1 3.3 2.1a4 4 0 0 1 1 3.6a4.3 4.3 0 0 1-3.5 3V21a1 1 0 1 1-2 0v-1a6 6 0 0 1-3.9-2a1 1 0 1 1 1.6-1.3c.5.7 1.2 1.1 2 1.3c2 .5 3.5-.4 3.8-1.6a2 2 0 0 0-.6-1.8a4.4 4.4 0 0 0-2.2-1.5a6.4 6.4 0 0 1-3.4-2a4 4 0 0 1-.9-3.7c.4-1.7 2-2.8 3.6-3.1V3c0-.6.5-1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DotsHorizontalOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="3" d="M6 12h.01m6 0h.01m5.99 0h.01"/>`), g.Group(children),
	)
}

func DotsHorizontalSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="3" d="M6 12h0m6 0h0m6 0h0"/>`), g.Group(children),
	)
}

func DotsVerticalOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="3" d="M12 6h.01M12 12h.01M12 18h.01"/>`), g.Group(children),
	)
}

func DotsVerticalSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="3" d="M12 6h0m0 6h0m0 6h0"/>`), g.Group(children),
	)
}

func DownloadOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 13V4M7 14H5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-2m-1-5l-4 5l-4-5m9 8h.01"/>`), g.Group(children),
	)
}

func DownloadSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 11.15V4a1 1 0 1 0-2 0v7.15L8.78 8.374a1 1 0 1 0-1.56 1.25l4 5a1 1 0 0 0 1.56 0l4-5a1 1 0 1 0-1.56-1.25z"/><path d="M9.657 15.874L7.358 13H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-2.358l-2.3 2.874a3 3 0 0 1-4.685 0M17 16a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2z"/></g>`), g.Group(children),
	)
}

func DrawSquareOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 6.5h2M11 18h2m-7-5v-2m12 2v-2M5 8h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1m0 12h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1m12 0h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1m0-12h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1"/>`), g.Group(children),
	)
}

func DrawSquareSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M5 3a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm0 12a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2zm12 0a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2zm0-12a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2z"/><path fill-rule="evenodd" d="M10 6.5a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1M10 18a1 1 0 0 1 1-1h2a1 1 0 1 1 0 2h-2a1 1 0 0 1-1-1m-4-4a1 1 0 0 1-1-1v-2a1 1 0 1 1 2 0v2a1 1 0 0 1-1 1m12 0a1 1 0 0 1-1-1v-2a1 1 0 1 1 2 0v2a1 1 0 0 1-1 1" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func DribbbleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2a10 10 0 1 0 10 10A10.01 10.01 0 0 0 12 2m6.613 4.614a8.52 8.52 0 0 1 1.93 5.32a20.1 20.1 0 0 0-5.949-.274c-.059-.149-.122-.292-.184-.441a24 24 0 0 0-.566-1.239a11.4 11.4 0 0 0 4.769-3.366M10 3.707a8.8 8.8 0 0 1 2-.238a8.5 8.5 0 0 1 5.664 2.152a9.6 9.6 0 0 1-4.476 3.087A46 46 0 0 0 10 3.707m-6.358 6.555a8.57 8.57 0 0 1 4.73-5.981a54 54 0 0 1 3.168 4.941a32 32 0 0 1-7.9 1.04zm2.01 7.46a8.5 8.5 0 0 1-2.2-5.707v-.262a31.6 31.6 0 0 0 8.777-1.219c.243.477.477.964.692 1.449q-.172.05-.336.1a13.57 13.57 0 0 0-6.942 5.636zM12 20.556a8.5 8.5 0 0 1-5.243-1.8a11.72 11.72 0 0 1 6.7-5.332l.055-.02a35.7 35.7 0 0 1 1.819 6.476a8.5 8.5 0 0 1-3.331.676m4.772-1.462A37 37 0 0 0 15.113 13a12.5 12.5 0 0 1 5.321.364a8.56 8.56 0 0 1-3.66 5.73z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DropboxSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M12.013 6.175L7.006 9.369l5.007 3.194l-5.007 3.193L2 12.545l5.006-3.193L2 6.175l5.006-3.194zM6.981 17.806l5.006-3.193l5.006 3.193L11.987 21z"/><path d="m12.013 12.545l5.006-3.194l-5.006-3.176l4.98-3.194L22 6.175l-5.007 3.194L22 12.562l-5.007 3.194z"/></g>`), g.Group(children),
	)
}

func EditOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.304 4.844l2.852 2.852M7 7H4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-4.5m2.409-9.91a2.017 2.017 0 0 1 0 2.853l-6.844 6.844L8 14l.713-3.565l6.844-6.844a2.015 2.015 0 0 1 2.852 0Z"/>`), g.Group(children),
	)
}

func EditSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M11.32 6.176H5c-1.105 0-2 .949-2 2.118v10.588C3 20.052 3.895 21 5 21h11c1.105 0 2-.948 2-2.118v-7.75l-3.914 4.144A2.46 2.46 0 0 1 12.81 16l-2.681.568c-1.75.37-3.292-1.263-2.942-3.115l.536-2.839c.097-.512.335-.983.684-1.352z"/><path d="M19.846 4.318a2.2 2.2 0 0 0-.437-.692a2 2 0 0 0-.654-.463a1.92 1.92 0 0 0-1.544 0a2 2 0 0 0-.654.463l-.546.578l2.852 3.02l.546-.579a2.1 2.1 0 0 0 .437-.692a2.24 2.24 0 0 0 0-1.635M17.45 8.721L14.597 5.7L9.82 10.76a.54.54 0 0 0-.137.27l-.536 2.84c-.07.37.239.696.588.622l2.682-.567a.5.5 0 0 0 .255-.145l4.778-5.06Z"/></g>`), g.Group(children),
	)
}

func EnvelopeOpenOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 8v11a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V8m18 0l-8.029-4.46a2 2 0 0 0-1.942 0L3 8m18 0l-9 6.5L3 8"/>`), g.Group(children),
	)
}

func EnvelopeOpenSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="m3.62 6.389l8.396 6.724l8.638-6.572l-7.69-4.29a1.98 1.98 0 0 0-1.928 0z"/><path d="m22 8.053l-8.784 6.683a1.98 1.98 0 0 1-2.44-.031L2.02 7.693a1 1 0 0 0-.019.199v11.065C2 20.637 3.343 22 5 22h14c1.657 0 3-1.362 3-3.043z"/></g>`), g.Group(children),
	)
}

func EnvelopeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m3.5 5.5l7.893 6.036a1 1 0 0 0 1.214 0L20.5 5.5M4 19h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Z"/>`), g.Group(children),
	)
}

func EnvelopeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M2.038 5.61A2 2 0 0 0 2 6v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6q0-.18-.03-.352l-.866.65l-7.89 6.032a2 2 0 0 1-2.429 0L2.884 6.288l-.846-.677Z"/><path d="M20.677 4.117A2 2 0 0 0 20 4H4q-.338.002-.642.105l.758.607L12 10.742L19.9 4.7z"/></g>`), g.Group(children),
	)
}

func EuroOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 10h9.231M6 14h9.231M18 5.086A5.95 5.95 0 0 0 14.615 4c-3.738 0-6.769 3.582-6.769 8s3.031 8 6.769 8A5.94 5.94 0 0 0 18 18.916"/>`), g.Group(children),
	)
}

func EuroSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.3 9c1-3.4 3.8-6 7.3-6a7 7 0 0 1 4 1.3a1 1 0 1 1-1.2 1.6a5 5 0 0 0-2.8-.9c-2.2 0-4.3 1.6-5.2 4h5.8a1 1 0 1 1 0 2H9a8.7 8.7 0 0 0 0 2h6.3a1 1 0 1 1 0 2H9.4c1 2.4 3 4 5.2 4c1 0 2-.3 2.8-.9a1 1 0 1 1 1.2 1.6a7 7 0 0 1-4 1.3c-3.5 0-6.3-2.6-7.3-6H6a1 1 0 1 1 0-2h.9a10.4 10.4 0 0 1 0-2H6a1 1 0 1 1 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ExclamationCircleOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 13V8m0 8h.01M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`), g.Group(children),
	)
}

func ExclamationCircleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m11-4a1 1 0 1 0-2 0v5a1 1 0 1 0 2 0zm-1 7a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ExpandOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4H4m0 0v4m0-4l5 5m7-5h4m0 0v4m0-4l-5 5M8 20H4m0 0v-4m0 4l5-5m7 5h4m0 0v-4m0 4l-5-5"/>`), g.Group(children),
	)
}

func ExpandSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4H4v4m.5-3.5L9 9m7-5h4v4m-.5-3.5L15 9M8 20H4v-4m.5 3.5L9 15m7 5h4v-4m-.5 3.5L15 15"/>`), g.Group(children),
	)
}

func EyeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path d="M21 12c0 1.2-4.03 6-9 6s-9-4.8-9-6s4.03-6 9-6s9 4.8 9 6Z"/><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g>`), g.Group(children),
	)
}

func EyeSlashOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3.933 13.909A4.36 4.36 0 0 1 3 12c0-1 4-6 9-6m7.6 3.8A5.07 5.07 0 0 1 21 12c0 1-3 6-9 6q-.471 0-.918-.04M5 19L19 5m-4 7a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/>`), g.Group(children),
	)
}

func EyeSlashSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="m4 15.6l3.055-3.056A5 5 0 0 1 7 12.012a5.006 5.006 0 0 1 5-5q.268.014.532.054l1.744-1.744A9 9 0 0 0 12 5.012c-5.388 0-10 5.336-10 7A6.5 6.5 0 0 0 4 15.6"/><path d="m14.7 10.726l4.995-5.007A.998.998 0 0 0 18.99 4a1 1 0 0 0-.71.305l-4.995 5.007a3 3 0 0 0-.588-.21l-.035-.01a2.98 2.98 0 0 0-3.584 3.583c0 .012.008.022.01.033q.075.307.211.59l-4.995 4.983a1 1 0 1 0 1.414 1.414l4.995-4.983q.284.137.59.211c.011 0 .021.007.033.01a2.982 2.982 0 0 0 3.584-3.584c0-.012-.008-.023-.011-.035a3 3 0 0 0-.21-.588Z"/><path d="m19.821 8.605l-2.857 2.857a4.952 4.952 0 0 1-5.514 5.514l-1.785 1.785c.767.166 1.55.25 2.335.251c6.453 0 10-5.258 10-7c0-1.166-1.637-2.874-2.179-3.407"/></g>`), g.Group(children),
	)
}

func EyeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.998 7.78C6.729 6.345 9.198 5 12 5s5.27 1.345 7.002 2.78a12.7 12.7 0 0 1 2.096 2.183c.253.344.465.682.618.997c.14.286.284.658.284 1.04s-.145.754-.284 1.04a6.6 6.6 0 0 1-.618.997a12.7 12.7 0 0 1-2.096 2.183C17.271 17.655 14.802 19 12 19s-5.27-1.345-7.002-2.78a12.7 12.7 0 0 1-2.096-2.183a6.6 6.6 0 0 1-.618-.997C2.144 12.754 2 12.382 2 12s.145-.754.284-1.04c.153-.315.365-.653.618-.997A12.7 12.7 0 0 1 4.998 7.78M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FaceExplodeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 17a2 2 0 0 1 2 2h-4a2 2 0 0 1 2-2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.815 9H16.5a2 2 0 1 0-1.03-3.707A2 2 0 0 0 15.5 5A1.992 1.992 0 0 0 12 3.69A1.992 1.992 0 0 0 8.5 5q.003.147.03.293A2 2 0 1 0 7.5 9h3.388m2.927-.985v3.604M10.228 9v2.574M15 16h.01M9 16h.01m11.962-4.426a1.805 1.805 0 0 1-1.74 1.326a1.89 1.89 0 0 1-1.811-1.326a1.9 1.9 0 0 1-3.621 0a1.8 1.8 0 0 1-1.749 1.326a1.98 1.98 0 0 1-1.87-1.326A1.76 1.76 0 0 1 8.46 12.9a2.035 2.035 0 0 1-1.905-1.326A1.9 1.9 0 0 1 4.74 12.9A1.805 1.805 0 0 1 3 11.574V12a9 9 0 0 0 18 0z"/></g>`), g.Group(children),
	)
}

func FaceExplodeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21.972 11.517a.527.527 0 0 0-1.034-.105a1.38 1.38 0 0 1-1.324 1.01a1.47 1.47 0 0 1-1.4-1.009a.526.526 0 0 0-1.015 0a1.467 1.467 0 0 1-2.737.143l-.049-.204l.021-.146V9.369h2.304a2.63 2.63 0 0 0 2.631-2.632a2.68 2.68 0 0 0-2.654-2.632l-.526.022l-.13-.369A2.63 2.63 0 0 0 13.579 2c-.461 0-.915.124-1.313.358L12 2.513l-.266-.155A2.6 2.6 0 0 0 10.422 2a2.63 2.63 0 0 0-2.483 1.759l-.13.37l-.518-.024a2.68 2.68 0 0 0-2.66 2.632A2.63 2.63 0 0 0 7.264 9.37H9.61v1.887l-.007.09l-.028.08a1.33 1.33 0 0 1-1.301.996a1.63 1.63 0 0 1-1.502-1.024a.526.526 0 0 0-1.01.013a1.47 1.47 0 0 1-1.404 1.01a1.38 1.38 0 0 1-1.325-1.01a.55.55 0 0 0-.569-.382h-.008a.526.526 0 0 0-.456.526v.446a10.01 10.01 0 0 0 10 10a9.9 9.9 0 0 0 7.067-2.94A10.02 10.02 0 0 0 22 11.966zM8.316 15.685a1.053 1.053 0 1 1 2.105 0a1.053 1.053 0 0 1-2.105 0m1.58 3.684a2.105 2.105 0 0 1 4.21 0zm4.736-2.631a1.052 1.052 0 1 1 0-2.105a1.052 1.052 0 0 1 0 2.105"/>`), g.Group(children),
	)
}

func FaceGrinOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.99 9H15M8.99 9H9m12 3a9 9 0 1 1-18 0a9 9 0 0 1 18 0M7 13c0 1 .507 2.397 1.494 3.216a5.5 5.5 0 0 0 7.022 0C16.503 15.397 17 14 17 13c0 0-1.99 1-4.995 1S7 13 7 13"/>`), g.Group(children),
	)
}

func FaceGrinSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m5.495.93A.5.5 0 0 0 6.5 13c0 1.19.644 2.438 1.618 3.375C9.099 17.319 10.469 18 12 18s2.9-.681 3.882-1.625c.974-.937 1.618-2.184 1.618-3.375a.5.5 0 0 0-.995-.07a.8.8 0 0 1-.156.096c-.214.106-.554.208-1.006.295c-.896.173-2.111.262-3.343.262s-2.447-.09-3.343-.262c-.452-.087-.792-.19-1.005-.295a.8.8 0 0 1-.157-.096M8.99 8a1 1 0 0 0 0 2H9a1 1 0 1 0 0-2zm6 0a1 1 0 1 0 0 2H15a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FaceGrinStarsOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 13c0 2.038-2.239 4.5-5 4.5S7 15.038 7 13c0 1.444 10 1.444 10 0"/><path fill="currentColor" d="m9 6.811l.618 1.253l1.382.2l-1 .975l.236 1.377L9 9.966l-1.236.65L8 9.239l-1-.975l1.382-.2zm6 0l.618 1.253l1.382.2l-1 .975l.236 1.377L15 9.966l-1.236.65L14 9.239l-1-.975l1.382-.2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m9 6.811l.618 1.253l1.382.2l-1 .975l.236 1.377L9 9.966l-1.236.65L8 9.239l-1-.975l1.382-.2zm6 0l.618 1.253l1.382.2l-1 .975l.236 1.377L15 9.966l-1.236.65L14 9.239l-1-.975l1.382-.2z"/></g>`), g.Group(children),
	)
}

func FaceGrinStarsSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2M7 12.5a.5.5 0 0 1 .495.43a.8.8 0 0 0 .157.096c.213.106.553.208 1.005.295c.896.173 2.111.262 3.343.262s2.447-.09 3.343-.262c.452-.087.792-.19 1.006-.295a.8.8 0 0 0 .156-.096a.5.5 0 0 1 .995.07c0 1.19-.644 2.438-1.618 3.375C14.9 17.319 13.531 18 12 18s-2.9-.681-3.882-1.625C7.144 15.438 6.5 14.19 6.5 13a.5.5 0 0 1 .5-.5m9.519.417l.003-.004zm-9.038 0l-.003-.004zm.901-4.853L9 6.81l.619 1.253l1.381.2l-1 .976l.236 1.376l-1.237-.65l-1.235.65L8 9.239l-1-.975zm6 0L15 6.81l.619 1.253l1.381.2l-1 .976l.236 1.376l-1.237-.65l-1.235.65L14 9.239l-1-.975z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FaceLaughOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 9h.01M8.99 9H9m12 3a9 9 0 1 1-18 0a9 9 0 0 1 18 0M6.6 13a5.5 5.5 0 0 0 10.81 0z"/>`), g.Group(children),
	)
}

func FaceLaughSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2M7.99 9a1 1 0 0 1 1-1H9a1 1 0 0 1 0 2h-.01a1 1 0 0 1-1-1M14 9a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H15a1 1 0 0 1-1-1m-5.506 7.216A5.5 5.5 0 0 1 6.6 13h10.81a5.5 5.5 0 0 1-8.916 3.216" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FacebookSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.135 6H15V3h-1.865a4.147 4.147 0 0 0-4.142 4.142V9H7v3h2v9.938h3V12h2.021l.592-3H12V6.591A.6.6 0 0 1 12.592 6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileChartBarOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m4 10v-2m3 2v-6m3 6v-3m4-11v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1"/>`), g.Group(children),
	)
}

func FileChartBarSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5zm2 0V2h7a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2m-1 9a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0zm2-5a1 1 0 0 1 1 1v6a1 1 0 1 1-2 0v-6a1 1 0 0 1 1-1m4 4a1 1 0 1 0-2 0v3a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileCheckOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m4 6l2 2l4-4m4-8v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1"/>`), g.Group(children),
	)
}

func FileCheckSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5z"/><path fill-rule="evenodd" d="M11 7V2h7a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2m4.707 5.707a1 1 0 0 0-1.414-1.414L11 14.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func FileCirclePlusOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9V4a1 1 0 0 0-1-1H8.914a1 1 0 0 0-.707.293L4.293 7.207A1 1 0 0 0 4 7.914V20a1 1 0 0 0 1 1h4M9 3v4a1 1 0 0 1-1 1H4m11 6v4m-2-2h4m3 0a5 5 0 1 1-10 0a5 5 0 0 1 10 0"/>`), g.Group(children),
	)
}

func FileCirclePlusSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5zm2 0V2h7a2 2 0 0 1 2 2v6.41A7.5 7.5 0 1 0 10.5 22H6a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2"/><path d="M9 16a6 6 0 1 1 12 0a6 6 0 0 1-12 0m6-3a1 1 0 0 1 1 1v1h1a1 1 0 1 1 0 2h-1v1a1 1 0 1 1-2 0v-1h-1a1 1 0 1 1 0-2h1v-1a1 1 0 0 1 1-1"/></g>`), g.Group(children),
	)
}

func FileCloneOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M15 4v3a1 1 0 0 1-1 1h-3m2 10v1a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-7.13a1 1 0 0 1 .24-.65L6.7 8.35A1 1 0 0 1 7.46 8H9m-1 4H4m16-7v10a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1V7.87a1 1 0 0 1 .24-.65l2.46-2.87a1 1 0 0 1 .76-.35H19a1 1 0 0 1 1 1Z"/>`), g.Group(children),
	)
}

func FileCloneSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 12.732A2 2 0 0 1 7 13H3v6a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2h-2a4 4 0 0 1-4-4zM7 11V7.054a2 2 0 0 0-1.059.644l-2.46 2.87A2 2 0 0 0 3.2 11z"/><path d="M14 3.054V7h-3.8q.111-.232.282-.432l2.46-2.87A2 2 0 0 1 14 3.054M16 3v4a2 2 0 0 1-2 2h-4v6a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2z"/></g>`), g.Group(children),
	)
}

func FileCodeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m5 4l-2 2l2 2m4-4l2 2l-2 2m5-12v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1"/>`), g.Group(children),
	)
}

func FileCodeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.221V7H4.221a2 2 0 0 1 .365-.5L8.5 2.586A2 2 0 0 1 9 2.22ZM11 2v5a2 2 0 0 1-2 2H4v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm-.293 9.293a1 1 0 0 1 0 1.414L9.414 14l1.293 1.293a1 1 0 0 1-1.414 1.414l-2-2a1 1 0 0 1 0-1.414l2-2a1 1 0 0 1 1.414 0m2.586 1.414a1 1 0 0 1 1.414-1.414l2 2a1 1 0 0 1 0 1.414l-2 2a1 1 0 0 1-1.414-1.414L14.586 14z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileCopyAltOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M14 4v3a1 1 0 0 1-1 1h-3m4 10v1a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h2m11-3v10a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1V7.87a1 1 0 0 1 .24-.65l2.46-2.87a1 1 0 0 1 .76-.35H18a1 1 0 0 1 1 1Z"/>`), g.Group(children),
	)
}

func FileCopyAltSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M7 9v6a4 4 0 0 0 4 4h4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9a2 2 0 0 1 2-2h1z"/><path d="M13 3.054V7H9.2a2 2 0 0 1 .281-.432l2.46-2.87A2 2 0 0 1 13 3.054M15 3v4a2 2 0 0 1-2 2H9v6a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2z"/></g>`), g.Group(children),
	)
}

func FileCopyOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M9 8v3a1 1 0 0 1-1 1H5m11 4h2a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1h-7a1 1 0 0 0-1 1v1m4 3v10a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-7.13a1 1 0 0 1 .24-.65L7.7 8.35A1 1 0 0 1 8.46 8H13a1 1 0 0 1 1 1Z"/>`), g.Group(children),
	)
}

func FileCopySolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M18 3a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1V9a4 4 0 0 0-4-4h-3a2 2 0 0 0-1 .267V5a2 2 0 0 1 2-2z"/><path d="M8 7.054V11H4.2a2 2 0 0 1 .281-.432l2.46-2.87A2 2 0 0 1 8 7.054M10 7v4a2 2 0 0 1-2 2H4v6a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z"/></g>`), g.Group(children),
	)
}

func FileCsvOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1v6M5 19v1a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-1M10 3v4a1 1 0 0 1-1 1H5m2.665 9H6.647A1.647 1.647 0 0 1 5 15.353v-1.706A1.647 1.647 0 0 1 6.647 12h1.018M16 12l1.443 4.773L19 12m-6.057-.152l-.943-.02a1.34 1.34 0 0 0-1.359 1.22a1.32 1.32 0 0 0 1.172 1.421l.536.059a1.273 1.273 0 0 1 1.226 1.718c-.2.571-.636.754-1.337.754h-1.13"/>`), g.Group(children),
	)
}

func FileCsvSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.221V7H4.221a2 2 0 0 1 .365-.5L8.5 2.586A2 2 0 0 1 9 2.22ZM11 2v5a2 2 0 0 1-2 2H4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2V4a2 2 0 0 0-2-2zm1.018 8.828a2.34 2.34 0 0 0-2.373 2.13v.008a2.32 2.32 0 0 0 2.06 2.497l.535.059a1 1 0 0 0 .136.006a.272.272 0 0 1 .263.367l-.008.02l-.018.044l-.078.02a2 2 0 0 1-.297.021h-1.13a1 1 0 1 0 0 2h1.13c.417 0 .892-.05 1.324-.279c.47-.248.78-.648.953-1.134a2.272 2.272 0 0 0-2.115-3.06l-.478-.052a.32.32 0 0 1-.285-.341a.34.34 0 0 1 .344-.306l.94.02a1 1 0 1 0 .043-2l-.943-.02zm7.933 1.482a1 1 0 1 0-1.902-.62l-.57 1.747l-.522-1.726a1 1 0 0 0-1.914.578l1.443 4.773a1 1 0 0 0 1.908.021zm-13.762.88a.65.65 0 0 1 .458-.19h1.018a1 1 0 1 0 0-2H6.647A2.647 2.647 0 0 0 4 13.647v1.706A2.647 2.647 0 0 0 6.647 18h1.018a1 1 0 1 0 0-2H6.647A.647.647 0 0 1 6 15.353v-1.706c0-.172.068-.336.19-.457Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileDocOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 10V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1v6M5 19v1a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-1M10 3v4a1 1 0 0 1-1 1H5m14 9.006h-.335a1.647 1.647 0 0 1-1.647-1.647v-1.706a1.647 1.647 0 0 1 1.647-1.647L19 12M5 12v5h1.375A1.626 1.626 0 0 0 8 15.375v-1.75A1.626 1.626 0 0 0 6.375 12zm9 1.5v2a1.5 1.5 0 0 1-1.5 1.5v0a1.5 1.5 0 0 1-1.5-1.5v-2a1.5 1.5 0 0 1 1.5-1.5v0a1.5 1.5 0 0 1 1.5 1.5"/>`), g.Group(children),
	)
}

func FileDocSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M6 16v-3h.375a.626.626 0 0 1 .625.626v1.749a.626.626 0 0 1-.626.625zm6-2.5a.5.5 0 1 1 1 0v2a.5.5 0 0 1-1 0z"/><path fill-rule="evenodd" d="M11 7V2h7a2 2 0 0 1 2 2v5h1a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-1a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2H3a1 1 0 0 1-1-1v-9a1 1 0 0 1 1-1h6a2 2 0 0 0 2-2m7.683 6.006l1.335-.024l-.037-2l-1.327.024a2.647 2.647 0 0 0-2.636 2.647v1.706a2.647 2.647 0 0 0 2.647 2.647H20v-2h-1.335a.647.647 0 0 1-.647-.647v-1.706a.647.647 0 0 1 .647-.647zM5 11a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h1.376A2.626 2.626 0 0 0 9 15.375v-1.75A2.626 2.626 0 0 0 6.375 11zm7.5 0a2.5 2.5 0 0 0-2.5 2.5v2a2.5 2.5 0 0 0 5 0v-2a2.5 2.5 0 0 0-2.5-2.5" clip-rule="evenodd"/><path d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5z"/></g>`), g.Group(children),
	)
}

func FileExportOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 10V4a1 1 0 0 0-1-1H9.914a1 1 0 0 0-.707.293L5.293 7.207A1 1 0 0 0 5 7.914V20a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-2M10 3v4a1 1 0 0 1-1 1H5m5 6h9m0 0l-2-2m2 2l-2 2"/>`), g.Group(children),
	)
}

func FileExportSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5zm2 0V2h7a2 2 0 0 1 2 2v9.293l-2-2a1 1 0 0 0-1.414 1.414l.293.293h-6.586a1 1 0 1 0 0 2h6.586l-.293.293A1 1 0 0 0 18 16.707l2-2V20a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileIcvoiceSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7V2.2a2 2 0 0 0-.5.4l-4 3.9a2 2 0 0 0-.3.5zm2 0V2h7a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2m2-2a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2zm0 3a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2zm-6 4c0-.6.4-1 1-1h8c.6 0 1 .4 1 1v6c0 .6-.4 1-1 1H8a1 1 0 0 1-1-1zm8 1v1h-2v-1zm0 3h-2v1h2zm-4-3v1H9v-1zm0 3H9v1h2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileImageOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M16 18H8l2.5-6l2 4l1.5-2zm-1-8.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m14-4v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1M8 18h8l-2-4l-1.5 2l-2-4zm7-8.5a.5.5 0 1 1-1 0a.5.5 0 0 1 1 0"/></g>`), g.Group(children),
	)
}

func FileImageSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.221V7H4.221a2 2 0 0 1 .365-.5L8.5 2.586A2 2 0 0 1 9 2.22ZM11 2v5a2 2 0 0 1-2 2H4v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm.394 9.553a1 1 0 0 0-1.817.062l-2.5 6A1 1 0 0 0 8 19h8a1 1 0 0 0 .894-1.447l-2-4A1 1 0 0 0 13.2 13.4l-.53.706zM13 9.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileImportOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-4m5-13v4a1 1 0 0 1-1 1H5m0 6h9m0 0l-2-2m2 2l-2 2"/>`), g.Group(children),
	)
}

func FileImportSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5zm2 0V2h7a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2v-5h7.586l-.293.293a1 1 0 0 0 1.414 1.414l2-2a1 1 0 0 0 0-1.414l-2-2a1 1 0 0 0-1.414 1.414l.293.293H4V9h5a2 2 0 0 0 2-2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileInvoiceOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m8-2h3m-3 3h3m-4 3v6m4-3H8M19 4v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1M8 12v6h8v-6z"/>`), g.Group(children),
	)
}

func FileInvoiceSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5zm2 0V2h7a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2m2-2a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2zm0 3a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2zm-6 4a1 1 0 0 1 1-1h8a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1zm8 1v1h-2v-1zm0 3h-2v1h2zm-4-3v1H9v-1zm0 3H9v1h2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileLinesOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m4 8h6m-6-4h6m4-8v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1"/>`), g.Group(children),
	)
}

func FileLinesSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.221V7H4.221a2 2 0 0 1 .365-.5L8.5 2.586A2 2 0 0 1 9 2.22ZM11 2v5a2 2 0 0 1-2 2H4v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zM8 16a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1m1-5a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileMusicOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m8 7.5V8s3 1 3 4m3-8v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1m-6 12c0 1.105-1.12 2-2.5 2S8 17.105 8 16s1.12-2 2.5-2s2.5.895 2.5 2"/>`), g.Group(children),
	)
}

func FileMusicSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5zm2 0V2h7a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2m2.318.052h-.002A1 1 0 0 0 12 8v5.293A4 4 0 0 0 10.5 13C8.787 13 7 14.146 7 16s1.787 3 3.5 3s3.5-1.146 3.5-3q0-.16-.017-.313A1 1 0 0 0 14 15.5V9.766c.538.493 1 1.204 1 2.234a1 1 0 1 0 2 0c0-1.881-.956-3.14-1.86-3.893a6.4 6.4 0 0 0-1.636-.985l-.165-.063l-.014-.005l-.005-.001zM9 16c0-.356.452-1 1.5-1s1.5.644 1.5 1s-.452 1-1.5 1S9 16.356 9 16" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m14-4v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1Z"/>`), g.Group(children),
	)
}

func FilePasteOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 20H5a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h2.429M7 8h3M8 8V4h4v2m4 0V5h-4m3 4v3a1 1 0 0 1-1 1h-3m9-3v9a1 1 0 0 1-1 1h-7a1 1 0 0 1-1-1v-6.397a1 1 0 0 1 .27-.683l2.434-2.603a1 1 0 0 1 .73-.317H19a1 1 0 0 1 1 1"/>`), g.Group(children),
	)
}

func FilePasteSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.5 3.85c0-.47.392-.85.875-.85h5.25c.483 0 .875.38.875.85h1.75c.966 0 1.75.761 1.75 1.7V6h-1c-.728 0-1.732-.06-2.434.095a4 4 0 0 0-.88.307l-.061-.002h-.875V4.7h-3.5v1.7h-.875a.863.863 0 0 0-.875.85c0 .47.392.85.875.85h3.36L9.077 9.871a4 4 0 0 0-.892 1.526C7.97 12.083 8 13.268 8 14v5c0 .729.195 1.412.535 2H4.75C3.784 21 3 20.239 3 19.3V5.55c0-.939.784-1.7 1.75-1.7z"/><path d="M14 8.048V12h-3.907a2 2 0 0 1 .446-.763l2.434-2.603A2 2 0 0 1 14 8.048M16 8v4a2 2 0 0 1-2 2h-4v5a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2v-9a2 2 0 0 0-2-2z"/></g>`), g.Group(children),
	)
}

func FilePdfOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17v-5h1.5a1.5 1.5 0 1 1 0 3H5m12 2v-5h2m-2 3h2M5 10V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1v6M5 19v1a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-1M10 3v4a1 1 0 0 1-1 1H5m6 4v5h1.375A1.627 1.627 0 0 0 14 15.375v-1.75A1.627 1.627 0 0 0 12.375 12z"/>`), g.Group(children),
	)
}

func FilePdfSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.221V7H4.221a2 2 0 0 1 .365-.5L8.5 2.586A2 2 0 0 1 9 2.22ZM11 2v5a2 2 0 0 1-2 2H4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2V4a2 2 0 0 0-2-2zm-6 9a1 1 0 0 0-1 1v5a1 1 0 1 0 2 0v-1h.5a2.5 2.5 0 0 0 0-5zm1.5 3H6v-1h.5a.5.5 0 0 1 0 1m4.5-3a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h1.376A2.626 2.626 0 0 0 15 15.375v-1.75A2.626 2.626 0 0 0 12.375 11zm1 5v-3h.375a.626.626 0 0 1 .625.626v1.748a.625.625 0 0 1-.626.626zm5-5a1 1 0 0 0-1 1v5a1 1 0 1 0 2 0v-1h1a1 1 0 1 0 0-2h-1v-1h1a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FilePenOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 5V4a1 1 0 0 0-1-1H8.914a1 1 0 0 0-.707.293L4.293 7.207A1 1 0 0 0 4 7.914V20a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-5M9 3v4a1 1 0 0 1-1 1H4m11.383.772l2.745 2.746m1.215-3.906a2.09 2.09 0 0 1 0 2.953l-6.65 6.646L9 17.95l.739-3.692l6.646-6.646a2.087 2.087 0 0 1 2.958 0"/>`), g.Group(children),
	)
}

func FilePenSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 7V2.221a2 2 0 0 0-.5.365L3.586 6.5a2 2 0 0 0-.365.5zm2 0V2h7a2 2 0 0 1 2 2v.126a5.09 5.09 0 0 0-4.74 1.368v.001l-6.642 6.642a3 3 0 0 0-.82 1.532l-.74 3.692a3 3 0 0 0 3.53 3.53l3.694-.738a3 3 0 0 0 1.532-.82L19 15.149V20a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2"/><path d="M17.447 8.08a1.09 1.09 0 0 1 1.187.238l.002.001a1.09 1.09 0 0 1 0 1.539l-.377.377l-1.54-1.542l.373-.374l.002-.001q.152-.154.353-.237Zm-2.143 2.027l-4.644 4.644l-.385 1.924l1.925-.385l4.644-4.642l-1.54-1.54Zm2.56-4.11a3.1 3.1 0 0 0-2.187.909l-6.645 6.645a1 1 0 0 0-.274.51l-.739 3.693a1 1 0 0 0 1.177 1.176l3.693-.738a1 1 0 0 0 .51-.274l6.65-6.646a3.088 3.088 0 0 0-2.185-5.275"/></g>`), g.Group(children),
	)
}

func FilePptOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17v-5h1.5a1.5 1.5 0 1 1 0 3H5m6 2v-5h1.5a1.5 1.5 0 1 1 0 3H11m7-3v5m-1-5h2M5 10V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1v6M5 19v1a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-1M10 3v4a1 1 0 0 1-1 1H5"/>`), g.Group(children),
	)
}

func FilePptSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.221V7H4.221a2 2 0 0 1 .365-.5L8.5 2.586A2 2 0 0 1 9 2.22ZM11 2v5a2 2 0 0 1-2 2H4a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2V4a2 2 0 0 0-2-2zm-6 9a1 1 0 0 0-1 1v5a1 1 0 1 0 2 0v-1h.5a2.5 2.5 0 0 0 0-5zm1.5 3H6v-1h.5a.5.5 0 0 1 0 1m4.5-3a1 1 0 0 0-1 1v5a1 1 0 1 0 2 0v-1h.5a2.5 2.5 0 0 0 0-5zm1.5 3H12v-1h.5a.5.5 0 0 1 0 1m4.5-3a1 1 0 1 0 0 2v4a1 1 0 1 0 2 0v-4a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileSearchOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m8 7.5l2.5 2.5M19 4v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1m-5 9.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0"/>`), g.Group(children),
	)
}

func FileSearchSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5zm2 0V2h7a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2m.5 5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m0 5c.47 0 .917-.092 1.326-.26l1.967 1.967a1 1 0 0 0 1.414-1.414l-1.817-1.818A3.5 3.5 0 1 0 11.5 17" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileShieldOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 9V4a1 1 0 0 0-1-1H8.914a1 1 0 0 0-.707.293L4.293 7.207A1 1 0 0 0 4 7.914V20a1 1 0 0 0 1 1h6M9 3v4a1 1 0 0 1-1 1H4m11 13a11.4 11.4 0 0 1-3.637-3.99A11.14 11.14 0 0 1 10 11.833L15 10l5 1.833a11.14 11.14 0 0 1-1.363 5.176A11.4 11.4 0 0 1 15.001 21Z"/>`), g.Group(children),
	)
}

func FileShieldSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5zm2 0V2h7a2 2 0 0 1 2 2v5.703l-4.311-1.58a2 2 0 0 0-1.377 0l-5 1.832A2 2 0 0 0 8 11.861c.03 2.134.582 4.228 1.607 6.106c.848 1.555 2 2.924 3.382 4.033H6a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2"/><path d="M15.345 9.061a1 1 0 0 0-.689 0l-5 1.833a1 1 0 0 0-.656.953c.028 1.97.538 3.905 1.485 5.641a12.4 12.4 0 0 0 3.956 4.34a1 1 0 0 0 1.12 0a12.4 12.4 0 0 0 3.954-4.34A12.14 12.14 0 0 0 21 11.848a1 1 0 0 0-.656-.954zM15 19.765a10.4 10.4 0 0 0 2.76-3.235a10.15 10.15 0 0 0 1.206-4.011L15 11.065z"/></g>`), g.Group(children),
	)
}

func FileSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.221V7H4.221a2 2 0 0 1 .365-.5L8.5 2.586A2 2 0 0 1 9 2.22ZM11 2v5a2 2 0 0 1-2 2H4v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileVideoOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m14-4v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1ZM9 12h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1Zm5.697 2.395v-.733l1.269-1.219v2.984z"/>`), g.Group(children),
	)
}

func FileVideoSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5zm2 0V2h7a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2m-2 4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h2a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2zm0 2h2v2H9zm7.965-.557a1 1 0 0 0-1.692-.72l-1.268 1.218a1 1 0 0 0-.308.721v.733a1 1 0 0 0 .37.776l1.267 1.032a1 1 0 0 0 1.631-.776z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileWordOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m4 4l1 5l2-3.333L14 17l1-5m4-8v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1"/>`), g.Group(children),
	)
}

func FileWordSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7V2.221a2 2 0 0 0-.5.365L4.586 6.5a2 2 0 0 0-.365.5zm2 0V2h7a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V9h5a2 2 0 0 0 2-2m-1.02 4.804a1 1 0 1 0-1.96.392l1 5a1 1 0 0 0 1.838.319L12 15.61l1.143 1.905a1 1 0 0 0 1.838-.319l1-5a1 1 0 0 0-1.962-.392l-.492 2.463l-.67-1.115a1 1 0 0 0-1.714 0l-.67 1.116l-.492-2.464Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileZipOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M10 3v4a1 1 0 0 1-1 1H5m14-4v16a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V7.914a1 1 0 0 1 .293-.707l3.914-3.914A1 1 0 0 1 9.914 3H18a1 1 0 0 1 1 1Zm-4 1h.01v.01H15zm-2 2h.01v.01H13zm2 2h.01v.01H15zm-2 2h.01v.01H13zm2 2h.01v.01H15zm-2 2h.01v.01H13zm2 2h.01v.01H15zm-2 2h.01v.01H13z"/>`), g.Group(children),
	)
}

func FileZipSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2.221V7H4.221a2 2 0 0 1 .365-.5L8.5 2.586A2 2 0 0 1 9 2.22ZM11 2v5a2 2 0 0 1-2 2H4v11a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2zm3 2h2.01v2.01h-2V8h2v2.01h-2V12h2v2.01h-2V16h2v2.01h-2v2H12V18h2v-1.99h-2V14h2v-1.99h-2V10h2V8.01h-2V6h2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FilterOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M18.796 4H5.204a1 1 0 0 0-.753 1.659l5.302 6.058a1 1 0 0 1 .247.659v4.874a.5.5 0 0 0 .2.4l3 2.25a.5.5 0 0 0 .8-.4v-7.124a1 1 0 0 1 .247-.659l5.302-6.059c.566-.646.106-1.658-.753-1.658Z"/>`), g.Group(children),
	)
}

func FilterSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.05 3C3.291 3 2.352 5.024 3.51 6.317l5.422 6.059v4.874c0 .472.227.917.613 1.2l3.069 2.25c1.01.742 2.454.036 2.454-1.2v-7.124l5.422-6.059C21.647 5.024 20.708 3 18.95 3z"/>`), g.Group(children),
	)
}

func FingerprintOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a28.1 28.1 0 0 1-1.091 9M7.231 4.37a8.994 8.994 0 0 1 12.88 3.73M2.958 15S3 14.577 3 12a8.95 8.95 0 0 1 1.735-5.307m12.84 3.088A6 6 0 0 1 18 12a30 30 0 0 1-.464 6.232M6 12a6 6 0 0 1 9.352-4.974M4 21a5.96 5.96 0 0 1 1.01-3.328a5.15 5.15 0 0 0 .786-1.926m8.66 2.486a14 14 0 0 1-.962 2.683M7.5 19.336C9 17.092 9 14.845 9 12a3 3 0 1 1 6 0c0 .749 0 1.521-.031 2.311M12 12c0 3 0 6-2 9"/>`), g.Group(children),
	)
}

func FingerprintSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12c.1 3-.2 6-1 9M7.1 4.4a9 9 0 0 1 13 3.7M3 15v-3a9 9 0 0 1 1.7-5.3m12.9 3c.3.8.4 1.5.4 2.3c0 2 0 4.2-.5 6.2M6 12a6 6 0 0 1 9.4-5M4 21a6 6 0 0 1 1-3.3a5 5 0 0 0 .8-2m8.7 2.5a14 14 0 0 1-1 2.7m-6-1.6C9 17.1 9 14.8 9 12a3 3 0 1 1 6 0v2.3M12 12c0 3 0 6-2 9"/>`), g.Group(children),
	)
}

func FireOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.122 17.645a7.2 7.2 0 0 1-2.656 2.495a7.06 7.06 0 0 1-3.52.853a6.6 6.6 0 0 1-3.306-.718a6.73 6.73 0 0 1-2.54-2.266c-2.672-4.57.287-8.846.887-9.668A4.45 4.45 0 0 0 8.07 6.31A4.5 4.5 0 0 0 7.997 4c1.284.965 6.43 3.258 5.525 10.631c1.496-1.136 2.7-3.046 2.846-6.216c1.43 1.061 3.985 5.462 1.754 9.23"/>`), g.Group(children),
	)
}

func FireSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.597 3.2A1 1 0 0 0 7.04 4.289a3.5 3.5 0 0 1 .057 1.795a3.45 3.45 0 0 1-.84 1.575a1 1 0 0 0-.077.094c-.596.817-3.96 5.6-.941 10.762l.03.049a7.73 7.73 0 0 0 2.917 2.602a7.6 7.6 0 0 0 3.772.829a8.06 8.06 0 0 0 3.986-.975a8.2 8.2 0 0 0 3.04-2.864c1.301-2.2 1.184-4.556.588-6.441c-.583-1.848-1.68-3.414-2.607-4.102a1 1 0 0 0-1.594.757c-.067 1.431-.363 2.551-.794 3.431c-.222-2.407-1.127-4.196-2.224-5.524c-1.147-1.39-2.564-2.3-3.323-2.788a9 9 0 0 1-.432-.287Z"/>`), g.Group(children),
	)
}

func FlagOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 14v7M5 4.971v9.541c5.6-5.538 8.4 2.64 14-.086v-9.54C13.4 7.61 10.6-.568 5 4.97Z"/>`), g.Group(children),
	)
}

func FlagSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.09 3.294c1.924.95 3.422 1.69 5.472.692a1 1 0 0 1 1.438.9v9.54a1 1 0 0 1-.562.9c-2.981 1.45-5.382.24-7.25-.701a39 39 0 0 0-.622-.31c-1.033-.497-1.887-.812-2.756-.77c-.76.036-1.672.357-2.81 1.396V21a1 1 0 1 1-2 0V4.971a1 1 0 0 1 .297-.71c1.522-1.506 2.967-2.185 4.417-2.255c1.407-.068 2.653.453 3.72.967q.337.163.655.32Z"/>`), g.Group(children),
	)
}

func FloppyDiskAltOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M11 16h2m6.707-9.293l-2.414-2.414A1 1 0 0 0 16.586 4H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V7.414a1 1 0 0 0-.293-.707ZM16 20v-6a1 1 0 0 0-1-1H9a1 1 0 0 0-1 1v6zM9 4h6v3a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1z"/>`), g.Group(children),
	)
}

func FloppyDiskAltSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V7.414A2 2 0 0 0 20.414 6L18 3.586A2 2 0 0 0 16.586 3zm3 11a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v6H8zm1-7V5h6v2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1"/><path d="M14 17h-4v-2h4z"/></g>`), g.Group(children),
	)
}

func FloppyDiskOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M4 5a1 1 0 0 1 1-1h11.586a1 1 0 0 1 .707.293l2.414 2.414a1 1 0 0 1 .293.707V19a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1z"/><path d="M8 4h8v4H8zm7 10a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/></g>`), g.Group(children),
	)
}

func FloppyDiskSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V7.414A2 2 0 0 0 20.414 6L18 3.586A2 2 0 0 0 16.586 3zm10 11a3 3 0 1 1-6 0a3 3 0 0 1 6 0M8 7V5h8v2a1 1 0 0 1-1 1H9a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FlowbiteSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M15.907 11.998L10.332 9.23a1 1 0 0 1-.16-.037l-.018-.007v6.554c0 .017.008.034.01.051l2.388-2.974z"/><path d="m11.463 4.054l5.579 3.323A4 4 0 0 1 18.525 9c.332.668.47 1.414.398 2.155a3.07 3.07 0 0 1-.745 1.65a3.1 3.1 0 0 1-1.55.951c-.022.007-.045.005-.07.01q-.093.045-.191.08l-2.72.667l-1.992 2.48c-.18.227-.41.409-.67.534c.047.034.085.077.137.107a2.05 2.05 0 0 0 1.995.035c.592-.33 2.15-1.201 4.636-2.892l.28-.19c1.328-.895 3.616-2.442 3.967-4.215a9.94 9.94 0 0 0-1.713-4.154a10 10 0 0 0-3.375-2.989a10.1 10.1 0 0 0-8.802-.418c1.162.287 2.287.704 3.354 1.243Z"/><path d="M5.382 17.082v-6.457a3.7 3.7 0 0 1 .45-1.761a3.7 3.7 0 0 1 1.238-1.34a3.92 3.92 0 0 1 3.433-.245q.265.045.508.161l5.753 2.856q.123.075.236.165a2.13 2.13 0 0 0-.953-1.455l-5.51-3.284c-1.74-.857-3.906-1.523-5.244-1.097a10 10 0 0 0-2.5 3.496a9.9 9.9 0 0 0 .283 8.368a10 10 0 0 0 2.73 3.322a17 17 0 0 1-.424-2.729"/><path d="m19.102 16.163l-.272.183c-2.557 1.74-4.169 2.64-4.698 2.935a4.1 4.1 0 0 1-2 .53a3.95 3.95 0 0 1-1.983-.535a3.8 3.8 0 0 1-1.36-1.361a3.75 3.75 0 0 1-.51-1.85a2 2 0 0 1-.043-.26V9.143c0-.024.009-.046.01-.07q-.084.03-.162.07a1.8 1.8 0 0 0-.787 1.516v6.377a10.7 10.7 0 0 0 1.113 4.27a10.11 10.11 0 0 0 8.505-.53a10 10 0 0 0 3.282-2.858a9.9 9.9 0 0 0 1.75-3.97a19.6 19.6 0 0 1-2.845 2.216Z"/></g>`), g.Group(children),
	)
}

func FolderArrowRightOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 8H4m4 6h8m0 0l-2-2m2 2l-2 2M4 6v13a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-5.032a1 1 0 0 1-.768-.36l-1.9-2.28a1 1 0 0 0-.768-.36H5a1 1 0 0 0-1 1"/>`), g.Group(children),
	)
}

func FolderArrowRightSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 4a2 2 0 0 0-2 2v1h10.968l-1.9-2.28A2 2 0 0 0 10.532 4zM3 19V9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m11.707-7.707a1 1 0 0 0-1.414 1.414l.293.293H8a1 1 0 1 0 0 2h5.586l-.293.293a1 1 0 0 0 1.414 1.414l2-2a1 1 0 0 0 0-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderDuplicateOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 11H4m15.5 5a.5.5 0 0 0 .5-.5V8a1 1 0 0 0-1-1h-3.75a1 1 0 0 1-.829-.44l-1.436-2.12a1 1 0 0 0-.828-.44H8a1 1 0 0 0-1 1M4 9v10a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-7a1 1 0 0 0-1-1h-3.75a1 1 0 0 1-.829-.44L9.985 8.44A1 1 0 0 0 9.157 8H5a1 1 0 0 0-1 1"/>`), g.Group(children),
	)
}

func FolderDuplicateSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M6 5a2 2 0 0 1 2-2h4.157a2 2 0 0 1 1.656.879L15.249 6H19a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2v-5a3 3 0 0 0-3-3h-3.22l-1.14-1.682A3 3 0 0 0 9.157 6H6z"/><path d="M3 9a2 2 0 0 1 2-2h4.157a2 2 0 0 1 1.656.879L12.249 10H3zm0 3v7a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2v-7z"/></g>`), g.Group(children),
	)
}

func FolderOpenOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 19V6a1 1 0 0 1 1-1h4.032a1 1 0 0 1 .768.36l1.9 2.28a1 1 0 0 0 .768.36H16a1 1 0 0 1 1 1v1M3 19l3-8h15l-3 8z"/>`), g.Group(children),
	)
}

func FolderOpenSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 .087.586l2.977-7.937A1 1 0 0 1 6 10h12V9a2 2 0 0 0-2-2h-4.532l-1.9-2.28A2 2 0 0 0 8.032 4zm2.693 8H6.5l-3 8H18l3-8z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.5 8H4m0-2v13a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-5.032a1 1 0 0 1-.768-.36l-1.9-2.28a1 1 0 0 0-.768-.36H5a1 1 0 0 0-1 1"/>`), g.Group(children),
	)
}

func FolderPlusOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 8H4m8 3.5v5M9.5 14h5M4 6v13a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V9a1 1 0 0 0-1-1h-5.032a1 1 0 0 1-.768-.36l-1.9-2.28a1 1 0 0 0-.768-.36H5a1 1 0 0 0-1 1"/>`), g.Group(children),
	)
}

func FolderPlusSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 4a2 2 0 0 0-2 2v1h10.968l-1.9-2.28A2 2 0 0 0 10.532 4zM3 19V9h18v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m9-8.5a1 1 0 0 1 1 1V13h1.5a1 1 0 1 1 0 2H13v1.5a1 1 0 1 1-2 0V15H9.5a1 1 0 1 1 0-2H11v-1.5a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 6a2 2 0 0 1 2-2h5.532a2 2 0 0 1 1.536.72l1.9 2.28H3zm0 3v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V9z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ForwardOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.248 19C3.22 15.77 5.275 8.232 12.466 8.232V6.079a1.025 1.025 0 0 1 1.644-.862l5.479 4.307a1.108 1.108 0 0 1 0 1.723l-5.48 4.307a1.026 1.026 0 0 1-1.643-.861v-2.154C5.275 13.616 4.248 19 4.248 19"/>`), g.Group(children),
	)
}

func ForwardSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.027 10.9a8.73 8.73 0 0 1 6.422-3.62v-1.2A2.06 2.06 0 0 1 12.61 4.2a1.99 1.99 0 0 1 2.104.23l5.491 4.308a2.11 2.11 0 0 1 .588 2.566a2.1 2.1 0 0 1-.588.734l-5.489 4.308a1.98 1.98 0 0 1-2.104.228a2.07 2.07 0 0 1-1.16-1.876v-.942c-5.33 1.284-6.212 5.251-6.25 5.441a1 1 0 0 1-.923.806h-.06a1 1 0 0 1-.955-.7A10.22 10.22 0 0 1 5.027 10.9"/>`), g.Group(children),
	)
}

func ForwardStepOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 6v12M8 6v12l8-6z"/>`), g.Group(children),
	)
}

func ForwardStepSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 6a1 1 0 1 0-2 0v4L8.6 5.2A1 1 0 0 0 7 6v12a1 1 0 0 0 1.6.8L15 14v4a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func GiftBoxOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 21v-9m3-4H7.5a2.5 2.5 0 1 1 0-5c1.5 0 2.875 1.25 3.875 2.5M14 21v-9m-9 0h14v8a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1zM4 8h16a1 1 0 0 1 1 1v3H3V9a1 1 0 0 1 1-1m12.155-5c-3 0-5.5 5-5.5 5h5.5a2.5 2.5 0 0 0 0-5"/>`), g.Group(children),
	)
}

func GiftBoxSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 7h-.7c.229-.467.349-.98.351-1.5a3.5 3.5 0 0 0-3.5-3.5c-1.717 0-3.215 1.2-4.331 2.481C10.4 2.842 8.949 2 7.5 2A3.5 3.5 0 0 0 4 5.5c.003.52.123 1.033.351 1.5H4a2 2 0 0 0-2 2v2a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V9a2 2 0 0 0-2-2m-9.942 0H7.5a1.5 1.5 0 0 1 0-3c.9 0 2 .754 3.092 2.122c-.219.337-.392.635-.534.878m6.1 0h-3.742c.933-1.368 2.371-3 3.739-3a1.5 1.5 0 0 1 0 3zM13 14h-2v8h2zm-4 0H4v6a2 2 0 0 0 2 2h3zm6 0v8h3a2 2 0 0 0 2-2v-6z"/>`), g.Group(children),
	)
}

func GithubSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.006 2a9.85 9.85 0 0 0-6.484 2.44a10.32 10.32 0 0 0-3.393 6.17a10.48 10.48 0 0 0 1.317 6.955a10.05 10.05 0 0 0 5.4 4.418c.504.095.683-.223.683-.494c0-.245-.01-1.052-.014-1.908c-2.78.62-3.366-1.21-3.366-1.21a2.7 2.7 0 0 0-1.11-1.5c-.907-.637.07-.621.07-.621c.317.044.62.163.885.346c.266.183.487.426.647.71c.135.253.318.476.538.655a2.08 2.08 0 0 0 2.37.196c.045-.52.27-1.006.635-1.37c-2.219-.259-4.554-1.138-4.554-5.07a4.02 4.02 0 0 1 1.031-2.75a3.77 3.77 0 0 1 .096-2.713s.839-.275 2.749 1.05a9.26 9.26 0 0 1 5.004 0c1.906-1.325 2.74-1.05 2.74-1.05c.37.858.406 1.828.101 2.713a4.02 4.02 0 0 1 1.029 2.75c0 3.939-2.339 4.805-4.564 5.058a2.47 2.47 0 0 1 .679 1.897c0 1.372-.012 2.477-.012 2.814c0 .272.18.592.687.492a10.05 10.05 0 0 0 5.388-4.421a10.47 10.47 0 0 0 1.313-6.948a10.32 10.32 0 0 0-3.39-6.165A9.85 9.85 0 0 0 12.007 2Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func GlobeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M4.37 7.657c2.063.528 2.396 2.806 3.202 3.87c1.07 1.413 2.075 1.228 3.192 2.644c1.805 2.289 1.312 5.705 1.312 6.705M20 15h-1a4 4 0 0 0-4 4v1M8.587 3.992c0 .822.112 1.886 1.515 2.58c1.402.693 2.918.351 2.918 2.334c0 .276 0 2.008 1.972 2.008c2.026.031 2.026-1.678 2.026-2.008c0-.65.527-.9 1.177-.9H20M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0Z"/>`), g.Group(children),
	)
}

func GlobeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.64 4.737A8 8 0 0 1 12 4a8 8 0 0 1 6.933 4.006h-.738c-.65 0-1.177.25-1.177.9c0 .33 0 2.04-2.026 2.008c-1.972 0-1.972-1.732-1.972-2.008c0-1.429-.787-1.65-1.752-1.923c-.374-.105-.774-.218-1.166-.411c-1.004-.497-1.347-1.183-1.461-1.835ZM6 4a10.1 10.1 0 0 0-2.812 3.27A9.96 9.96 0 0 0 2 12c0 5.289 4.106 9.619 9.304 9.976l.054.004a10 10 0 0 0 1.155.007h.002a10 10 0 0 0 1.5-.19a10 10 0 0 0 2.259-.754a10.04 10.04 0 0 0 4.987-5.263A9.9 9.9 0 0 0 22 12a10 10 0 0 0-.315-2.5A10 10 0 0 0 12 2a9.96 9.96 0 0 0-6 2m13.372 11.113a2.6 2.6 0 0 0-.75-.112h-.217A3.405 3.405 0 0 0 15 18.405v1.014a8.03 8.03 0 0 0 4.372-4.307ZM12.114 20H12A8 8 0 0 1 5.1 7.95c.95.541 1.421 1.537 1.835 2.415c.209.441.403.853.637 1.162c.54.712 1.063 1.019 1.591 1.328c.52.305 1.047.613 1.6 1.316c1.44 1.825 1.419 4.366 1.35 5.828Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func GoogleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.037 21.998a10.3 10.3 0 0 1-7.168-3.049a9.9 9.9 0 0 1-2.868-7.118a9.95 9.95 0 0 1 3.064-6.949A10.37 10.37 0 0 1 12.212 2h.176a9.94 9.94 0 0 1 6.614 2.564L16.457 6.88a6.2 6.2 0 0 0-4.131-1.566a6.9 6.9 0 0 0-4.794 1.913a6.62 6.62 0 0 0-2.045 4.657a6.6 6.6 0 0 0 1.882 4.723a6.9 6.9 0 0 0 4.725 2.07h.143c1.41.072 2.8-.354 3.917-1.2a5.77 5.77 0 0 0 2.172-3.41l.043-.117H12.22v-3.41h9.678q.113.927.1 1.859c-.099 5.741-4.017 9.6-9.746 9.6l-.215-.002Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func GridOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.143 4H4.857A.857.857 0 0 0 4 4.857v4.286c0 .473.384.857.857.857h4.286A.857.857 0 0 0 10 9.143V4.857A.857.857 0 0 0 9.143 4m10 0h-4.286a.857.857 0 0 0-.857.857v4.286c0 .473.384.857.857.857h4.286A.857.857 0 0 0 20 9.143V4.857A.857.857 0 0 0 19.143 4m-10 10H4.857a.857.857 0 0 0-.857.857v4.286c0 .473.384.857.857.857h4.286a.857.857 0 0 0 .857-.857v-4.286A.857.857 0 0 0 9.143 14m10 0h-4.286a.857.857 0 0 0-.857.857v4.286c0 .473.384.857.857.857h4.286a.857.857 0 0 0 .857-.857v-4.286a.857.857 0 0 0-.857-.857"/>`), g.Group(children),
	)
}

func GridPlusOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 17h6m-3 3v-6M4.857 4h4.286c.473 0 .857.384.857.857v4.286a.857.857 0 0 1-.857.857H4.857A.857.857 0 0 1 4 9.143V4.857C4 4.384 4.384 4 4.857 4m10 0h4.286c.473 0 .857.384.857.857v4.286a.857.857 0 0 1-.857.857h-4.286A.857.857 0 0 1 14 9.143V4.857c0-.473.384-.857.857-.857m-10 10h4.286c.473 0 .857.384.857.857v4.286a.857.857 0 0 1-.857.857H4.857A.857.857 0 0 1 4 19.143v-4.286c0-.473.384-.857.857-.857"/>`), g.Group(children),
	)
}

func GridPlusSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.857 3A1.857 1.857 0 0 0 3 4.857v4.286C3 10.169 3.831 11 4.857 11h4.286A1.857 1.857 0 0 0 11 9.143V4.857A1.857 1.857 0 0 0 9.143 3zm10 0A1.857 1.857 0 0 0 13 4.857v4.286c0 1.026.831 1.857 1.857 1.857h4.286A1.857 1.857 0 0 0 21 9.143V4.857A1.857 1.857 0 0 0 19.143 3zm-10 10A1.857 1.857 0 0 0 3 14.857v4.286C3 20.169 3.831 21 4.857 21h4.286A1.857 1.857 0 0 0 11 19.143v-4.286A1.857 1.857 0 0 0 9.143 13zM18 14a1 1 0 1 0-2 0v2h-2a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func GridSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.857 3A1.857 1.857 0 0 0 3 4.857v4.286C3 10.169 3.831 11 4.857 11h4.286A1.857 1.857 0 0 0 11 9.143V4.857A1.857 1.857 0 0 0 9.143 3zm10 0A1.857 1.857 0 0 0 13 4.857v4.286c0 1.026.831 1.857 1.857 1.857h4.286A1.857 1.857 0 0 0 21 9.143V4.857A1.857 1.857 0 0 0 19.143 3zm-10 10A1.857 1.857 0 0 0 3 14.857v4.286C3 20.169 3.831 21 4.857 21h4.286A1.857 1.857 0 0 0 11 19.143v-4.286A1.857 1.857 0 0 0 9.143 13zm10 0A1.857 1.857 0 0 0 13 14.857v4.286c0 1.026.831 1.857 1.857 1.857h4.286A1.857 1.857 0 0 0 21 19.143v-4.286A1.857 1.857 0 0 0 19.143 13z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func HeadphonesOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M20 16v-4a8 8 0 1 0-16 0v4m16 0v2a2 2 0 0 1-2 2h-2v-6h2a2 2 0 0 1 2 2ZM4 16v2a2 2 0 0 0 2 2h2v-6H6a2 2 0 0 0-2 2Z"/>`), g.Group(children),
	)
}

func HeadphonesSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 5a7 7 0 0 0-7 7v1.17c.313-.11.65-.17 1-.17h2a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H6a3 3 0 0 1-3-3v-6a9 9 0 0 1 18 0v6a3 3 0 0 1-3 3h-2a1 1 0 0 1-1-1v-6a1 1 0 0 1 1-1h2c.35 0 .687.06 1 .17V12a7 7 0 0 0-7-7" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func HeartOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.01 6.001C6.5 1 1 8 5.782 13.001L12.011 20l6.23-7C23 8 17.5 1 12.01 6.002Z"/>`), g.Group(children),
	)
}

func HeartSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12.75 20.66l6.184-7.098c2.677-2.884 2.559-6.506.754-8.705c-.898-1.095-2.206-1.816-3.72-1.855c-1.293-.034-2.652.43-3.963 1.442c-1.315-1.012-2.678-1.476-3.973-1.442c-1.515.04-2.825.76-3.724 1.855c-1.806 2.201-1.915 5.823.772 8.706l6.183 7.097c.19.216.46.34.743.34a1 1 0 0 0 .743-.34Z"/>`), g.Group(children),
	)
}

func HomeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4 12l8-8l8 8M6 10.5V19a1 1 0 0 0 1 1h3v-3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v3h3a1 1 0 0 0 1-1v-8.5"/>`), g.Group(children),
	)
}

func HomeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.293 3.293a1 1 0 0 1 1.414 0l6 6l2 2a1 1 0 0 1-1.414 1.414L19 12.414V19a2 2 0 0 1-2 2h-3a1 1 0 0 1-1-1v-3h-2v3a1 1 0 0 1-1 1H7a2 2 0 0 1-2-2v-6.586l-.293.293a1 1 0 0 1-1.414-1.414l2-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func HourglassOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.5 4h-13m13 16h-13M8 20v-3.333a2 2 0 0 1 .4-1.2L10 12.6a1 1 0 0 0 0-1.2L8.4 8.533a2 2 0 0 1-.4-1.2V4h8v3.333a2 2 0 0 1-.4 1.2L13.957 11.4a1 1 0 0 0 0 1.2l1.643 2.867a2 2 0 0 1 .4 1.2V20z"/>`), g.Group(children),
	)
}

func HourglassSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 3a1 1 0 0 0 0 2H7v2.333a3 3 0 0 0 .556 1.74l1.57 2.814A1 1 0 0 0 9.2 12a1 1 0 0 0-.073.113l-1.57 2.814A3 3 0 0 0 7 16.667V19H5.5a1 1 0 1 0 0 2h13a1 1 0 1 0 0-2H17v-2.333a3 3 0 0 0-.56-1.745l-1.616-2.82a1 1 0 0 0-.067-.102a1 1 0 0 0 .067-.103l1.616-2.819A3 3 0 0 0 17 7.333V5h1.5a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func HtmlSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m3 2l1.578 17.824L12 22l7.467-2.175L21 2zm14.049 6.048H9.075l.172 2.016h7.697l-.626 6.565l-4.246 1.381l-4.281-1.455l-.288-2.932h2.024l.16 1.411l2.4.815l2.346-.763l.297-3.005H7.416l-.562-6.05h10.412z"/>`), g.Group(children),
	)
}

func ImageOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 16l5-7l6 6.5m6.5 2.5L16 13l-4.286 6M14 10h.01M4 19h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1"/>`), g.Group(children),
	)
}

func ImageSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13 10a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H14a1 1 0 0 1-1-1"/><path d="M2 6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12c0 .556-.227 1.06-.593 1.422A1 1 0 0 1 20.5 20H4a2 2 0 0 1-2-2zm6.892 12l3.833-5.356l-3.99-4.322a1 1 0 0 0-1.549.097L4 12.879V6h16v9.95l-3.257-3.619a1 1 0 0 0-1.557.088L11.2 18z"/></g>`), g.Group(children),
	)
}

func InboxFullOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 13h3.439a.99.99 0 0 1 .908.6a3.978 3.978 0 0 0 7.306 0a.99.99 0 0 1 .908-.6H20M4 13v6a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-6M4 13l2-9h12l2 9M9 7h6m-7 3h8"/>`), g.Group(children),
	)
}

func InboxFullSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.024 3.783A1 1 0 0 1 6 3h12a1 1 0 0 1 .976.783L20.802 12h-4.244a1.99 1.99 0 0 0-1.824 1.205a2.978 2.978 0 0 1-5.468 0A1.99 1.99 0 0 0 7.442 12H3.198zM3 14v5a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5h-4.43a4.978 4.978 0 0 1-9.14 0zm5-7a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1m0 2a1 1 0 0 0 0 2h8a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func InboxOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 13h3.439a.99.99 0 0 1 .908.6a3.978 3.978 0 0 0 7.306 0a.99.99 0 0 1 .908-.6H20M4 13v6a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-6M4 13l2-9h12l2 9"/>`), g.Group(children),
	)
}

func InboxSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.024 3.783A1 1 0 0 1 6 3h12a1 1 0 0 1 .976.783L20.802 12h-4.244a1.99 1.99 0 0 0-1.824 1.205a2.978 2.978 0 0 1-5.468 0A1.99 1.99 0 0 0 7.442 12H3.198zM3 14v5a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-5h-4.43a4.978 4.978 0 0 1-9.14 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func IndentOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 6h12M6 18h12m-5-8h5m-5 4h5M6 9v6l3.5-3z"/>`), g.Group(children),
	)
}

func IndentSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 6a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1m0 12a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1m1.65-9.76A1 1 0 0 0 5 9v6a1 1 0 0 0 1.65.76l3.5-3a1 1 0 0 0 0-1.52zM12 10a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2h-5a1 1 0 0 1-1-1m0 4a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2h-5a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func InfoCircleOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 11h2v5m-2 0h4m-2.592-8.5h.01M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`), g.Group(children),
	)
}

func InfoCircleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m9.408-5.5a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2zM10 10a1 1 0 1 0 0 2h1v3h-1a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2h-1v-4a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func InstagramSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 8a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5v8a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5zm5-3a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3zm7.597 2.214a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2h-.01a1 1 0 0 1-1-1M12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func KeyboardOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="square" stroke-width="2"><path d="M8 15h7.01v.01H15z"/><path d="M20 6H4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1Z"/><path d="M6 9h.01v.01H6zm0 3h.01v.01H6zm0 3h.01v.01H6zm3-6h.01v.01H9zm0 3h.01v.01H9zm3-3h.01v.01H12zm0 3h.01v.01H12zm3 0h.01v.01H15zm3 0h.01v.01H18zm0 3h.01v.01H18zm-3-6h.01v.01H15zm3 0h.01v.01H18z"/></g>`), g.Group(children),
	)
}

func KeyboardSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 7a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm5.01 1H5v2.01h2.01zm3 0H8v2.01h2.01zm3 0H11v2.01h2.01zm3 0H14v2.01h2.01zm3 0H17v2.01h2.01zm-12 3H5v2.01h2.01zm3 0H8v2.01h2.01zm3 0H11v2.01h2.01zm3 0H14v2.01h2.01zm3 0H17v2.01h2.01zm-12 3H5v2.01h2.01zM8 14l-.001 2l8.011.01V14zm11.01 0H17v2.01h2.01z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LabelOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.2 6H4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h11.2a1 1 0 0 0 .747-.334l4.46-5a1 1 0 0 0 0-1.332l-4.46-5A1 1 0 0 0 15.2 6"/>`), g.Group(children),
	)
}

func LabelSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 6a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h11.613a2 2 0 0 0 1.346-.52l4.4-4a2 2 0 0 0 0-2.96l-4.4-4A2 2 0 0 0 15.613 6z"/>`), g.Group(children),
	)
}

func LandmarkOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M3 21h18M4 18h16M6 10v8m4-8v8m4-8v8m4-8v8M4 9.5v-.955a1 1 0 0 1 .458-.84l7-4.52a1 1 0 0 1 1.084 0l7 4.52a1 1 0 0 1 .458.84V9.5a.5.5 0 0 1-.5.5h-15a.5.5 0 0 1-.5-.5Z"/>`), g.Group(children),
	)
}

func LandmarkSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M10.915 2.345a2 2 0 0 1 2.17 0l7 4.52A2 2 0 0 1 21 8.544V9.5a1.5 1.5 0 0 1-1.5 1.5H19v6h1a1 1 0 1 1 0 2H4a1 1 0 1 1 0-2h1v-6h-.5A1.5 1.5 0 0 1 3 9.5v-.955a2 2 0 0 1 .915-1.68zM17 17v-6h-2v6zm-6-6h2v6h-2zm-2 6v-6H7v6z" clip-rule="evenodd"/><path d="M2 21a1 1 0 0 1 1-1h18a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1"/></g>`), g.Group(children),
	)
}

func LanguageOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 19l3.5-9l3.5 9m-6.125-2h5.25M3 7h7m0 0h2m-2 0c0 1.63-.793 3.926-2.239 5.655M7.5 6.818V5m.261 7.655C6.79 13.82 5.521 14.725 4 15m3.761-2.345L5 10m2.761 2.655L10.2 15"/>`), g.Group(children),
	)
}

func LayersOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.005 11.19V12l6.998 4.042L19 12v-.81M5 16.15v.81L11.997 21l6.998-4.042v-.81M12.003 3L5.005 7.042l6.998 4.042L19 7.042z"/>`), g.Group(children),
	)
}

func LayersSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5.005 10.19a1 1 0 0 1 1 1v.233l5.998 3.464L18 11.423v-.232a1 1 0 1 1 2 0V12a1 1 0 0 1-.5.866l-6.997 4.042a1 1 0 0 1-1 0l-6.998-4.042a1 1 0 0 1-.5-.866v-.81a1 1 0 0 1 1-1M5 15.15a1 1 0 0 1 1 1v.232l5.997 3.464l5.998-3.464v-.232a1 1 0 1 1 2 0v.81a1 1 0 0 1-.5.865l-6.998 4.042a1 1 0 0 1-1 0L4.5 17.824a1 1 0 0 1-.5-.866v-.81a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path d="M12.503 2.134a1 1 0 0 0-1 0L4.501 6.17A1 1 0 0 0 4.5 7.902l7.002 4.047a1 1 0 0 0 1 0l6.998-4.04a1 1 0 0 0 0-1.732z"/></g>`), g.Group(children),
	)
}

func LetterBoldOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 5h4.5a3.5 3.5 0 1 1 0 7H8m0-7v7m0-7H6m2 7h6.5a3.5 3.5 0 1 1 0 7H8m0-7v7m0 0H6"/>`), g.Group(children),
	)
}

func LetterBoldSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 5c0-.6.4-1 1-1h6.5a4.5 4.5 0 0 1 3.5 7.3a4.5 4.5 0 0 1-1.5 8.7H6a1 1 0 1 1 0-2h1V6H6a1 1 0 0 1-1-1m4 1v5h3.5a2.5 2.5 0 0 0 0-5zm0 7v5h5.5a2.5 2.5 0 0 0 0-5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LetterItalicOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8.874 19l6.143-14M6 19h6.33m-.66-14H18"/>`), g.Group(children),
	)
}

func LetterItalicSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 4h-3.3a1 1 0 1 0 0 2h1.8L8.2 18H6a1 1 0 1 0 0 2h6.3a1 1 0 1 0 0-2h-1.9l5.3-12H18a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LetterUnderlineOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M6 19h12M8 5v9a4 4 0 0 0 8 0V5M6 5h4m4 0h4"/>`), g.Group(children),
	)
}

func LetterUnderlineSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 6H9v8a3 3 0 1 0 6 0V6h-1a1 1 0 1 1 0-2h4a1 1 0 1 1 0 2h-1v8a5 5 0 0 1-2 4h3a1 1 0 1 1 0 2H6a1 1 0 1 1 0-2h3a5 5 0 0 1-2-4V6H6a1 1 0 0 1 0-2h4a1 1 0 1 1 0 2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LifeSaverOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13.46 8.291l3.849-3.849a1.5 1.5 0 0 1 2.122 0l.127.127a1.5 1.5 0 0 1 0 2.122l-3.84 3.838a4 4 0 0 0-2.258-2.238m0 0a4 4 0 0 1 2.263 2.238l3.662-3.662a8.96 8.96 0 0 1 0 10.27l-3.676-3.676m-2.25-5.17l3.678-3.676a8.96 8.96 0 0 0-10.27 0l3.662 3.662a4 4 0 0 0-2.238 2.258L4.615 6.863a8.96 8.96 0 0 0 0 10.27l3.662-3.662a4 4 0 0 0 2.258 2.238l-3.672 3.676a8.96 8.96 0 0 0 10.27 0l-3.662-3.662a4 4 0 0 0 2.238-2.262m0 0l3.849 3.848a1.5 1.5 0 0 1 0 2.122l-.127.126a1.5 1.5 0 0 1-2.122 0l-3.838-3.838a4 4 0 0 0 2.238-2.258m.29-1.461a4 4 0 1 1-8 0a4 4 0 0 1 8 0m-7.718 1.471l-3.84 3.838a1.5 1.5 0 0 0 0 2.122l.128.126a1.5 1.5 0 0 0 2.122 0l3.848-3.848a4 4 0 0 1-2.258-2.238m2.248-5.19L6.69 4.442a1.5 1.5 0 0 0-2.122 0l-.127.127a1.5 1.5 0 0 0 0 2.122l3.849 3.848a4 4 0 0 1 2.238-2.258Z"/>`), g.Group(children),
	)
}

func LifeSaverSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7.4 3.736l3.43 3.429A5 5 0 0 1 12.133 7q.535.016 1.056.147l3.41-3.412a2.3 2.3 0 0 1 .451-.344A9.9 9.9 0 0 0 12.268 2a10 10 0 0 0-5.322 1.392q.249.143.454.344m11.451 1.54l-.127-.127a.5.5 0 0 0-.706 0l-2.932 2.932c.03.023.05.054.078.077q.356.292.651.645c.033.038.077.067.11.107l2.926-2.927a.5.5 0 0 0 0-.707m-2.931 9.81c-.025.03-.058.052-.082.082a5 5 0 0 1-.633.639c-.04.036-.072.083-.115.117l2.927 2.927a.5.5 0 0 0 .707 0l.127-.127a.5.5 0 0 0 0-.707l-2.932-2.931Zm-1.443-4.763a3.04 3.04 0 0 0-1.383-1.1l-.012-.007a3 3 0 0 0-1-.213H12a2.96 2.96 0 0 0-2.122.893c-.285.29-.509.634-.657 1.013l-.009.016a3 3 0 0 0-.21 1a3 3 0 0 0 .488 1.716l.032.04a3.04 3.04 0 0 0 1.384 1.1l.012.007c.319.129.657.2 1 .213c.393.015.784-.05 1.15-.192l.033-.018a3 3 0 0 0 1.676-1.7v-.007a2.9 2.9 0 0 0 0-2.207a3 3 0 0 0-.27-.515c-.007-.012-.02-.025-.03-.039m6.137-3.373a2.5 2.5 0 0 1-.349.447l-3.426 3.426c.112.428.166.869.161 1.311a5 5 0 0 1-.148 1.054l3.413 3.412q.2.202.347.444A9.9 9.9 0 0 0 22 12.269a9.9 9.9 0 0 0-1.386-5.319M16.6 20.264l-3.42-3.421c-.386.1-.782.152-1.18.157h-.135q-.535-.016-1.056-.147L7.4 20.265a2.5 2.5 0 0 1-.444.347A9.9 9.9 0 0 0 11.732 22H12a9.9 9.9 0 0 0 5.044-1.388a2.5 2.5 0 0 1-.444-.348M3.735 16.6l3.426-3.426a4.6 4.6 0 0 1-.013-2.367L3.735 7.4a2.5 2.5 0 0 1-.349-.447a9.89 9.89 0 0 0 0 10.1a2.5 2.5 0 0 1 .35-.453Zm5.101-.758a5 5 0 0 1-.65-.645c-.034-.038-.078-.067-.11-.107L5.15 18.017a.5.5 0 0 0 0 .707l.127.127a.5.5 0 0 0 .706 0l2.932-2.933c-.029-.018-.049-.053-.078-.076Zm-.755-6.928c.03-.037.07-.063.1-.1q.274-.33.6-.609c.046-.04.081-.092.128-.13L5.983 5.149a.5.5 0 0 0-.707 0l-.127.127a.5.5 0 0 0 0 .707z"/>`), g.Group(children),
	)
}

func LightbulbOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 9a3 3 0 0 1 3-3m-2 15h4m0-3c0-4.1 4-4.9 4-9A6 6 0 1 0 6 9c0 4 4 5 4 9z"/>`), g.Group(children),
	)
}

func LightbulbSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.05 4.05A7 7 0 0 1 19 9c0 2.407-1.197 3.874-2.186 5.084l-.04.048C15.77 15.362 15 16.34 15 18a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1c0-1.612-.77-2.613-1.78-3.875l-.045-.056C6.193 12.842 5 11.352 5 9a7 7 0 0 1 2.05-4.95M9 21a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1m1.586-13.414A2 2 0 0 1 12 7a1 1 0 1 0 0-2a4 4 0 0 0-4 4a1 1 0 0 0 2 0a2 2 0 0 1 .586-1.414" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LinkOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.213 9.787a3.39 3.39 0 0 0-4.795 0l-3.425 3.426a3.39 3.39 0 0 0 4.795 4.794l.321-.304m-.321-4.49a3.39 3.39 0 0 0 4.795 0l3.424-3.426a3.39 3.39 0 0 0-4.794-4.795l-1.028.961"/>`), g.Group(children),
	)
}

func LinkSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.2 9.8a3.4 3.4 0 0 0-4.8 0L5 13.2A3.4 3.4 0 0 0 9.8 18l.3-.3m-.3-4.5a3.4 3.4 0 0 0 4.8 0L18 9.8A3.4 3.4 0 0 0 13.2 5l-1 1"/>`), g.Group(children),
	)
}

func LinkedinSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12.51 8.796v1.697a3.74 3.74 0 0 1 3.288-1.684c3.455 0 4.202 2.16 4.202 4.97V19.5h-3.2v-5.072c0-1.21-.244-2.766-2.128-2.766c-1.827 0-2.139 1.317-2.139 2.676V19.5h-3.19V8.796h3.168ZM7.2 6.106a1.61 1.61 0 0 1-.988 1.483a1.595 1.595 0 0 1-1.743-.348A1.607 1.607 0 0 1 5.6 4.5a1.6 1.6 0 0 1 1.6 1.606" clip-rule="evenodd"/><path d="M7.2 8.809H4V19.5h3.2z"/></g>`), g.Group(children),
	)
}

func ListMusicOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 15.5V5s3 1 3 4m-7-3H4m9 4H4m4 4H4m13 2.4c0 1.326-1.343 2.4-3 2.4s-3-1.075-3-2.4s1.343-2.4 3-2.4s3 1.075 3 2.4"/>`), g.Group(children),
	)
}

func ListMusicSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.316 4.052a.99.99 0 0 0-.9.14c-.262.19-.416.495-.416.82v8.566a4.6 4.6 0 0 0-2-.464c-1.99 0-4 1.342-4 3.443S12.01 20 14 20s4-1.342 4-3.443V6.801c.538.5 1 1.219 1 2.262c0 .56.448 1.013 1 1.013s1-.453 1-1.013c0-1.905-.956-3.18-1.86-3.942a6.4 6.4 0 0 0-1.636-.998l-.166-.063l-.013-.005l-.005-.002h-.002zM4 5.012c-.552 0-1 .454-1 1.013c0 .56.448 1.013 1 1.013h9c.552 0 1-.453 1-1.013s-.448-1.012-1-1.012H4Zm0 4.051c-.552 0-1 .454-1 1.013c0 .56.448 1.013 1 1.013h9c.552 0 1-.454 1-1.013c0-.56-.448-1.013-1-1.013zm0 4.05c-.552 0-1 .454-1 1.014c0 .559.448 1.012 1 1.012h4c.552 0 1-.453 1-1.012c0-.56-.448-1.013-1-1.013z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ListOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M9 8h10M9 12h10M9 16h10M4.99 8H5m-.02 4h.01m0 4H5"/>`), g.Group(children),
	)
}

func ListSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 8c0-.6.4-1 1-1a1 1 0 1 1 0 2a1 1 0 0 1-1-1m4 0c0-.6.4-1 1-1h10a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1m-4 4c0-.6.4-1 1-1a1 1 0 1 1 0 2a1 1 0 0 1-1-1m4 0c0-.6.4-1 1-1h10a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1m-4 4c0-.6.4-1 1-1a1 1 0 1 1 0 2a1 1 0 0 1-1-1m4 0c0-.6.4-1 1-1h10a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LockOpenOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14v3m4-6V7a3 3 0 1 1 6 0v4M5 11h10a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1"/>`), g.Group(children),
	)
}

func LockOpenSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 7a2 2 0 1 1 4 0v4a1 1 0 1 0 2 0V7a4 4 0 0 0-8 0v3H5a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2zm-5 6a1 1 0 0 1 1 1v3a1 1 0 1 1-2 0v-3a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LockOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 14v3m-3-6V7a3 3 0 1 1 6 0v4m-8 0h10a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1v-7a1 1 0 0 1 1-1"/>`), g.Group(children),
	)
}

func LockSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 10V7a4 4 0 1 1 8 0v3h1a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2zm2-3a2 2 0 1 1 4 0v3h-4zm2 6a1 1 0 0 1 1 1v3a1 1 0 1 1-2 0v-3a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LockTimeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 11H5a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h4.5M7 11V7a3 3 0 0 1 6 0v1.5m2.5 5.5v1.5l1 1m3.5-1a4.5 4.5 0 1 1-9 0a4.5 4.5 0 0 1 9 0"/>`), g.Group(children),
	)
}

func LockTimeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M10 5a2 2 0 0 0-2 2v3h2.4A7.48 7.48 0 0 0 8 15.5a7.48 7.48 0 0 0 2.4 5.5H5a2 2 0 0 1-2-2v-7a2 2 0 0 1 2-2h1V7a4 4 0 1 1 8 0v1.15a7.5 7.5 0 0 0-1.943.685A1 1 0 0 1 12 8.5V7a2 2 0 0 0-2-2"/><path d="M10 15.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m6.5-1.5a1 1 0 1 0-2 0v1.5a1 1 0 0 0 .293.707l1 1a1 1 0 0 0 1.414-1.414l-.707-.707z"/></g>`), g.Group(children),
	)
}

func MagicWandOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.9 9.7L20 6.6L17.4 4L4 17.4L6.6 20zm0 0L14.3 7M6 7v2m0 0v2m0-2H4m2 0h2m7 7v2m0 0v2m0-2h-2m2 0h2m3-2"/>`), g.Group(children),
	)
}

func MagicWandSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.4 3c.3 0 .6.1.7.3l2.6 2.6c.4.3.4 1 0 1.4l-2.5 2.5l-4-4l2.5-2.5c.2-.2.5-.3.7-.3m-4.6 4.2l-9.5 9.5a1 1 0 0 0 0 1.4l2.6 2.6c.3.4 1 .4 1.4 0l9.5-9.5zM6 6c.6 0 1 .4 1 1v1h1a1 1 0 0 1 0 2H7v1a1 1 0 1 1-2 0v-1H4a1 1 0 0 1 0-2h1V7c0-.6.4-1 1-1m9 9c.6 0 1 .4 1 1v1h1a1 1 0 1 1 0 2h-1v1a1 1 0 1 1-2 0v-1h-1a1 1 0 1 1 0-2h1v-1c0-.6.4-1 1-1" clip-rule="evenodd"/><path d="M19 13h-2v2h2zM13 3h-2v2h2zm-2 2H9v2h2zM9 3H7v2h2zm12 8h-2v2h2zm0 4h-2v2h2z"/></g>`), g.Group(children),
	)
}

func MailBoxOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 16v-5.5A3.5 3.5 0 0 0 7.5 7m3.5 9H4v-5.5A3.5 3.5 0 0 1 7.5 7m3.5 9v4M7.5 7H14m0 0V4h2.5M14 7v3m-3.5 6H20v-6a3 3 0 0 0-3-3m-2 9v4m-8-6.5h1"/>`), g.Group(children),
	)
}

func MailBoxSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 6h-2V5h1a1 1 0 1 0 0-2h-2a1 1 0 0 0-1 1v2h-.541A5.97 5.97 0 0 1 14 10v4a1 1 0 1 1-2 0v-4c0-2.206-1.794-4-4-4q-.112.002-.22.028C7.686 6.022 7.596 6 7.5 6A4.505 4.505 0 0 0 3 10.5V16a1 1 0 0 0 1 1h7v3a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-3h5a1 1 0 0 0 1-1v-6c0-2.206-1.794-4-4-4m-9 8.5H7a1 1 0 1 1 0-2h1a1 1 0 1 1 0 2"/>`), g.Group(children),
	)
}

func MapLocationOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 13a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path d="M17.8 14a7 7 0 1 0-11.5 0h0l.1.3l.3.3L12 21l5.1-6.2l.6-.7l.1-.2Z"/></g>`), g.Group(children),
	)
}

func MapPinAltOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 13a3 3 0 1 0 0-6a3 3 0 0 0 0 6"/><path d="M17.8 13.938h-.011a7 7 0 1 0-11.464.144h-.016l.14.171q.15.19.3.371L12 21l5.13-6.248q.291-.314.54-.659z"/></g>`), g.Group(children),
	)
}

func MapPinAltSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.906 1.994a8 8 0 0 1 8.09 8.421a8 8 0 0 1-1.297 3.957a1 1 0 0 1-.133.204l-.108.129q-.268.365-.573.699l-5.112 6.224a1 1 0 0 1-1.545 0L5.982 15.26l-.002-.002a18 18 0 0 1-.309-.38l-.133-.163a1 1 0 0 1-.13-.202a7.995 7.995 0 0 1 6.498-12.518ZM15 9.997a3 3 0 1 1-5.999 0a3 3 0 0 1 5.999 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MapPinOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 15a6 6 0 1 0 0-12a6 6 0 0 0 0 12m0 0v6M9.5 9A2.5 2.5 0 0 1 12 6.5"/>`), g.Group(children),
	)
}

func MapPinSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 9a7 7 0 1 1 8 6.93V21a1 1 0 1 1-2 0v-5.07A7 7 0 0 1 5 9m5.94-1.06A1.5 1.5 0 0 1 12 7.5a1 1 0 1 0 0-2A3.5 3.5 0 0 0 8.5 9a1 1 0 0 0 2 0c0-.398.158-.78.44-1.06" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MastercardSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zm5.178 12.137a4.137 4.137 0 1 1 1.036-8.144A6.1 6.1 0 0 0 8.726 12c0 1.531.56 2.931 1.488 4.006a4 4 0 0 1-1.036.131M10.726 12c0-1.183.496-2.252 1.294-3.006A4.13 4.13 0 0 1 13.315 12a4.13 4.13 0 0 1-1.294 3.006A4.13 4.13 0 0 1 10.726 12m4.59 0a6.1 6.1 0 0 1-1.489 4.006a4.137 4.137 0 1 0 0-8.013A6.1 6.1 0 0 1 15.315 12Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MessageCaptionOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 9h5m3 0h2M7 12h2m3 0h5M5 5h14a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-6.616a1 1 0 0 0-.67.257l-2.88 2.592A.5.5 0 0 1 8 18.477V17a1 1 0 0 0-1-1H5a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1"/>`), g.Group(children),
	)
}

func MessageCaptionSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 6a2 2 0 0 1 2-2h14a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2h-6.616l-2.88 2.592C8.537 20.461 7 19.776 7 18.477V17H5a2 2 0 0 1-2-2zm4 2a1 1 0 0 0 0 2h5a1 1 0 1 0 0-2zm8 0a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2zm-8 3a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2zm5 0a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MessageDotsOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 10.5h.01m-4.01 0h.01M8 10.5h.01M5 5h14a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-6.6a1 1 0 0 0-.69.275l-2.866 2.723A.5.5 0 0 1 8 18.635V17a1 1 0 0 0-1-1H5a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1"/>`), g.Group(children),
	)
}

func MessageDotsSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 5.983C3 4.888 3.895 4 5 4h14c1.105 0 2 .888 2 1.983v8.923a1.99 1.99 0 0 1-2 1.983h-6.6l-2.867 2.7c-.955.899-2.533.228-2.533-1.08v-1.62H5c-1.105 0-2-.888-2-1.983zm5.706 3.809a1 1 0 1 0-1.412 1.417a1 1 0 1 0 1.412-1.417m2.585.002a1 1 0 1 1 .003 1.414a1 1 0 0 1-.003-1.414m5.415-.002a1 1 0 1 0-1.412 1.417a1 1 0 1 0 1.412-1.417" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MessagesOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17h6l3 3v-3h2V9h-2M4 4h11v8H9l-3 3v-3H4z"/>`), g.Group(children),
	)
}

func MessagesSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M4 3a1 1 0 0 0-1 1v8a1 1 0 0 0 1 1h1v2a1 1 0 0 0 1.707.707L9.414 13H15a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1z"/><path d="M8.023 17.215q.05-.046.098-.094L10.243 15H15a3 3 0 0 0 3-3V8h2a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-1v2a1 1 0 0 1-1.707.707L14.586 18H9a1 1 0 0 1-.977-.785"/></g>`), g.Group(children),
	)
}

func MicrophoneOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 9v3a5.006 5.006 0 0 1-5 5h-4a5.006 5.006 0 0 1-5-5V9m7 9v3m-3 0h6M11 3h2a3 3 0 0 1 3 3v5a3 3 0 0 1-3 3h-2a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3"/>`), g.Group(children),
	)
}

func MicrophoneSlashOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.97 9.012a1 1 0 1 0-2 0zm-1 2.988l1 .001V12zm-8.962 4.98l-.001 1h.001zm-3.52-1.46l.708-.708zM5.029 12h-1v.001zm3.984 7.963a1 1 0 1 0 0 2zm5.975 2a1 1 0 0 0 0-2zM7.017 8.017a1 1 0 1 0 2 0zm6.641 4.862a1 1 0 1 0 .667 1.886zm-7.63-2.87a1 1 0 1 0-2 0zm9.953 5.435a1 1 0 1 0 1 1.731zM12 16.979h1a1 1 0 0 0-1-1zM5.736 4.322a1 1 0 0 0-1.414 1.414zm12.528 15.356a1 1 0 0 0 1.414-1.414zM17.97 9.012V12h2V9.012zm0 2.987a4 4 0 0 1-1.168 2.813l1.415 1.414a6 6 0 0 0 1.753-4.225zm-7.962 3.98a4 4 0 0 1-2.813-1.167l-1.414 1.414a6 6 0 0 0 4.225 1.753zm-2.813-1.167a4 4 0 0 1-1.167-2.813l-2 .002a6 6 0 0 0 1.753 4.225zm3.808-10.775h1.992v-2h-1.992zm1.992 0c1.097 0 1.987.89 1.987 1.988h2a3.99 3.99 0 0 0-3.987-3.988zm1.987 1.988v4.98h2v-4.98zm-5.967 0c0-1.098.89-1.988 1.988-1.988v-2a3.99 3.99 0 0 0-3.988 3.988zm-.004 15.938H12v-2H9.012v2Zm2.988 0h2.987v-2H12zM9.016 8.017V6.025h-2v1.992zm5.967 2.987a1.99 1.99 0 0 1-1.325 1.875l.667 1.886a3.99 3.99 0 0 0 2.658-3.76zM6.03 12v-1.992h-2V12zm10.774 2.812a4 4 0 0 1-.823.632l1.002 1.731a6 6 0 0 0 1.236-.949zM4.322 5.736l13.942 13.942l1.414-1.414L5.736 4.322zM12 15.98h-1.992v2H12zm-1 1v3.984h2V16.98z"/>`), g.Group(children),
	)
}

func MicrophoneSlashSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15.506 14.097l.994.995A3.99 3.99 0 0 0 17.975 12V9.011a.996.996 0 0 1 1.992 0v2.99a5.98 5.98 0 0 1-2.054 4.503l1.762 1.762a.996.996 0 1 1-1.408 1.408L4.325 5.733a.996.996 0 0 1 1.408-1.408L7.04 5.632a3.984 3.984 0 0 1 3.964-3.59h1.992c2.2 0 3.983 1.783 3.983 3.983v4.98a3.98 3.98 0 0 1-1.473 3.092M4.033 10.008a.996.996 0 1 1 1.992 0V12a3.99 3.99 0 0 0 3.984 3.984H12c.55 0 .996.446.996.996v2.988h1.992a.996.996 0 0 1 0 1.992H9.012a.996.996 0 0 1 0-1.992h1.992v-1.992h-.997a5.98 5.98 0 0 1-5.974-5.974z"/>`), g.Group(children),
	)
}

func MicrophoneSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M5 8a1 1 0 0 1 1 1v3a4.006 4.006 0 0 0 4 4h4a4.006 4.006 0 0 0 4-4V9a1 1 0 1 1 2 0v3.001A6.006 6.006 0 0 1 14.001 18H13v2h2a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h2v-2H9.999A6.006 6.006 0 0 1 4 12.001V9a1 1 0 0 1 1-1" clip-rule="evenodd"/><path d="M7 6a4 4 0 0 1 4-4h2a4 4 0 0 1 4 4v5a4 4 0 0 1-4 4h-2a4 4 0 0 1-4-4z"/></g>`), g.Group(children),
	)
}

func MinimizeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9h4m0 0V5m0 4L4 4m15 5h-4m0 0V5m0 4l5-5M5 15h4m0 0v4m0-4l-5 5m15-5h-4m0 0v4m0-4l5 5"/>`), g.Group(children),
	)
}

func MinimizeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 9h4V5m-.5 3.5L4 4m15 5h-4V5m.5 3.5L20 4M5 15h4v4m-.5-3.5L4 20m15-5h-4v4m.5-3.5L20 20"/>`), g.Group(children),
	)
}

func MinusOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14"/>`), g.Group(children),
	)
}

func MobilePhoneOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 15h12M6 6h12m-6 12h.01M7 21h10a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1"/>`), g.Group(children),
	)
}

func MobilePhoneSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 4a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v16a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zm12 12V5H7v11zm-5 1a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MoonOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 0 1-.5-17.986V3c-.354.966-.5 1.911-.5 3a9 9 0 0 0 9 9c.239 0 .254.018.488 0A9 9 0 0 1 12 21"/>`), g.Group(children),
	)
}

func MoonSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.675 2.015a1 1 0 0 0-.403.011C6.09 2.4 2 6.722 2 12c0 5.523 4.477 10 10 10c4.356 0 8.058-2.784 9.43-6.667a1 1 0 0 0-1.02-1.33c-.08.006-.105.005-.127.005h-.001l-.028-.002A5 5 0 0 0 20 14a8 8 0 0 1-8-8c0-.952.121-1.752.404-2.558a1 1 0 0 0 .096-.428V3a1 1 0 0 0-.825-.985" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NewspaperOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7h1v12a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h11.5M7 14h6m-6 3h6m0-10h.5m-.5 3h.5M7 7h3v3H7z"/>`), g.Group(children),
	)
}

func NewspaperSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h11.5q.106 0 .207-.021Q16.85 21 17 21h2a2 2 0 0 0 2-2V7a1 1 0 0 0-1-1h-1a1 1 0 1 0 0 2v11h-2V5a2 2 0 0 0-2-2zm7 4a1 1 0 0 1 1-1h.5a1 1 0 1 1 0 2H13a1 1 0 0 1-1-1m0 3a1 1 0 0 1 1-1h.5a1 1 0 1 1 0 2H13a1 1 0 0 1-1-1m-6 4a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1m0 3a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1M7 6a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1zm1 3V8h1v1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NewspapperOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 7h1v12c0 .6-.4 1-1 1h-2a1 1 0 0 1-1-1V5c0-.6-.4-1-1-1H5a1 1 0 0 0-1 1v14c0 .6.4 1 1 1h11.5M7 14h6m-6 3h6m0-10h.5m-.5 3h.5M7 7h3v3H7z"/>`), g.Group(children),
	)
}

func NewspapperSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14c0 1.1.9 2 2 2h14a2 2 0 0 0 2-2V7c0-.6-.4-1-1-1h-1a1 1 0 1 0 0 2v11h-2V5a2 2 0 0 0-2-2zm7 4c0-.6.4-1 1-1h.5a1 1 0 1 1 0 2H13a1 1 0 0 1-1-1m0 3c0-.6.4-1 1-1h.5a1 1 0 1 1 0 2H13a1 1 0 0 1-1-1m-6 4c0-.6.4-1 1-1h6a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1m0 3c0-.6.4-1 1-1h6a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1M7 6a1 1 0 0 0-1 1v3c0 .6.4 1 1 1h3c.6 0 1-.4 1-1V7c0-.6-.4-1-1-1zm1 3V8h1v1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NpmSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3.87H4v16h8v-13h5v13h3v-16z"/>`), g.Group(children),
	)
}

func ObjectsColumnOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M4 5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1zm16 14a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1v-2a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1zM4 13a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1zm16-2a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1z"/>`), g.Group(children),
	)
}

func ObjectsColumnSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 3a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2zm14 18a2 2 0 0 0 2-2v-2a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2zM5 11a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2zm14 2a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2z"/>`), g.Group(children),
	)
}

func OrderedListOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6h8m-8 6h8m-8 6h8M4 16a2 2 0 1 1 3.321 1.5L4 20h5M4 5l2-1v6m-2 0h4"/>`), g.Group(children),
	)
}

func OrdoredListOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6h8m-8 6h8m-8 6h8M4 16a2 2 0 1 1 3.3 1.5L4 20h5M4 5l2-1v6m-2 0h4"/>`), g.Group(children),
	)
}

func OrdoredListSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 3.1c.3.2.5.6.5.9v5h1a1 1 0 0 1 0 2H4a1 1 0 1 1 0-2h1V5.6l-.6.3A1 1 0 0 1 3.6 4l2-1a1 1 0 0 1 1 0ZM11 6c0-.6.4-1 1-1h8a1 1 0 1 1 0 2h-8a1 1 0 0 1-1-1m0 6c0-.6.4-1 1-1h8a1 1 0 1 1 0 2h-8a1 1 0 0 1-1-1m-4.6 3A1 1 0 0 0 5 16a1 1 0 0 1-2 0a3 3 0 1 1 5 2.3L7 19h2a1 1 0 1 1 0 2H4a1 1 0 0 1-.6-1.8l3.3-2.5a1 1 0 0 0-.3-1.6Zm4.6 3c0-.6.4-1 1-1h8a1 1 0 1 1 0 2h-8a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func OutdentOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 6h12M6 18h12m-5-8h5m-5 4h5M9.5 9v6L6 12z"/>`), g.Group(children),
	)
}

func OutdentSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 6a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1m0 12a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1m3.85-9.76A1 1 0 0 1 10.5 9v6a1 1 0 0 1-1.65.76l-3.5-3a1 1 0 0 1 0-1.52zM12 10a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2h-5a1 1 0 0 1-1-1m0 4a1 1 0 0 1 1-1h5a1 1 0 1 1 0 2h-5a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PaletteOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7h.01m3.486 1.513h.01m-6.978 0h.01M6.99 12H7m9 4h2.706a1.96 1.96 0 0 0 1.883-1.325A9 9 0 1 0 3.043 12.89A9.1 9.1 0 0 0 8.2 20.1a8.6 8.6 0 0 0 3.769.9a2.013 2.013 0 0 0 2.03-2v-.857A2.036 2.036 0 0 1 16 16"/>`), g.Group(children),
	)
}

func PaletteSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-3a2 2 0 0 0-2-2h-1v3a1 1 0 1 1-2 0v-3h-1v3a1 1 0 1 1-2 0v-3h-1v3a1 1 0 1 1-2 0v-3H7a1 1 0 1 1 0-2h3v-1H7a1 1 0 1 1 0-2h3V8H7a1 1 0 0 1 0-2h3V5a2 2 0 0 0-2-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PaperClipOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 8v8a5 5 0 1 0 10 0V6.5a3.5 3.5 0 1 0-7 0V15a2 2 0 0 0 4 0V8"/>`), g.Group(children),
	)
}

func PaperPlaneOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 18l-7 3l7-18l7 18zm0 0v-5"/>`), g.Group(children),
	)
}

func PaperPlaneSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2a1 1 0 0 1 .932.638l7 18a1 1 0 0 1-1.326 1.281L13 19.517V13a1 1 0 1 0-2 0v6.517l-5.606 2.402a1 1 0 0 1-1.326-1.281l7-18A1 1 0 0 1 12 2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PapperPlaneSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2c.4 0 .8.3 1 .6l7 18a1 1 0 0 1-1.4 1.3L13 19.5V13a1 1 0 1 0-2 0v6.5L5.4 22A1 1 0 0 1 4 20.6l7-18a1 1 0 0 1 1-.6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ParagraphOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v7m0 7v-7m4-7v14m3-14H8.5A3.5 3.5 0 0 0 5 8.5v0A3.5 3.5 0 0 0 8.5 12H12"/>`), g.Group(children),
	)
}

func ParagraphSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 4a4.5 4.5 0 0 0 0 9H11v6a1 1 0 1 0 2 0V6h2v13a1 1 0 1 0 2 0V6h2a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PauseOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 6H8a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1m7 0h-1a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1"/>`), g.Group(children),
	)
}

func PauseSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2zm7 0a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h1a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PenNibOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m4.988 19.012l5.41-5.41m2.366-6.424l4.058 4.058l-2.03 5.41L5.3 20L4 18.701l3.355-9.494l5.41-2.029Zm4.626 4.625L12.197 6.61L14.807 4L20 9.194z"/>`), g.Group(children),
	)
}

func PenNibSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.514 3.293a1 1 0 0 0-1.415 0L12.151 5.24l.056.052l6.5 6.5l.052.056L20.707 9.9a1 1 0 0 0 0-1.415l-5.193-5.193ZM7.004 8.27l3.892-1.46l6.293 6.293l-1.46 3.893a1 1 0 0 1-.603.591l-9.494 3.355a1 1 0 0 1-.98-.18l6.452-6.453a1 1 0 0 0-1.414-1.414l-6.453 6.452a1 1 0 0 1-.18-.98l3.355-9.494a1 1 0 0 1 .591-.603Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PenOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.779 17.779L4.36 19.918L6.5 13.5m4.279 4.279l8.364-8.643a3.027 3.027 0 0 0-2.14-5.165a3.03 3.03 0 0 0-2.14.886L6.5 13.5m4.279 4.279L6.499 13.5m2.14 2.14l6.213-6.504M12.75 7.04L17 11.28"/>`), g.Group(children),
	)
}

func PenSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 4.182A4.14 4.14 0 0 1 16.9 3c1.087 0 2.13.425 2.899 1.182A4 4 0 0 1 21 7.037c0 1.068-.43 2.092-1.194 2.849L18.5 11.214l-5.8-5.71l1.287-1.31l.012-.012Zm-2.717 2.763L6.186 12.13l2.175 2.141l5.063-5.218zm-6.25 6.886l-1.98 5.849a.99.99 0 0 0 .245 1.026a1.03 1.03 0 0 0 1.043.242L10.282 19l-5.25-5.168Zm6.954 4.01l5.096-5.186l-2.218-2.183l-5.063 5.218l2.185 2.15Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PhoneHangupOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.693 16.013H7.31a1.685 1.685 0 0 0 1.685-1.684v-.645A1.684 1.684 0 0 1 10.679 12h2.647a1.686 1.686 0 0 1 1.686 1.686v.646c0 .446.178.875.494 1.19c.316.317.693.495 1.14.495h1.685a1.556 1.556 0 0 0 1.597-1.016c.078-.214.107-.776.088-1.002c.014-4.415-3.571-6.003-8-6.004c-4.427 0-8.014 1.585-8.01 5.996c-.02.227.009.79.087 1.003a1.56 1.56 0 0 0 1.6 1.02Z"/>`), g.Group(children),
	)
}

func PhoneHangupSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.017 6.995c-2.306 0-4.534.408-6.215 1.507c-1.737 1.135-2.788 2.944-2.797 5.451a5 5 0 0 0 .01.62c.015.193.047.512.138.763a2.56 2.56 0 0 0 2.579 1.677H7.31a2.685 2.685 0 0 0 2.685-2.684v-.645a.684.684 0 0 1 .684-.684h2.647a.686.686 0 0 1 .686.687v.645c0 .712.284 1.395.787 1.898c.478.478 1.101.787 1.847.787h1.647a2.555 2.555 0 0 0 2.575-1.674c.09-.25.123-.57.137-.763c.015-.2.022-.433.01-.617c-.002-2.508-1.049-4.32-2.785-5.458c-1.68-1.1-3.907-1.51-6.213-1.51"/>`), g.Group(children),
	)
}

func PhoneOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.427 14.768L17.2 13.542a1.733 1.733 0 0 0-2.45 0l-.613.613a1.73 1.73 0 0 1-2.45 0l-1.838-1.84a1.735 1.735 0 0 1 0-2.452l.612-.613a1.735 1.735 0 0 0 0-2.452L9.237 5.572a1.6 1.6 0 0 0-2.45 0c-3.223 3.2-1.702 6.896 1.519 10.117s6.914 4.745 10.12 1.535a1.6 1.6 0 0 0 0-2.456Z"/>`), g.Group(children),
	)
}

func PhoneSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.978 4a2.55 2.55 0 0 0-1.926.877C4.233 6.7 3.699 8.751 4.153 10.814c.44 1.995 1.778 3.893 3.456 5.572c1.68 1.679 3.577 3.018 5.57 3.459c2.062.456 4.115-.073 5.94-1.885a2.556 2.556 0 0 0 .001-3.861l-1.21-1.21a2.69 2.69 0 0 0-3.802 0l-.617.618a.806.806 0 0 1-1.14 0l-1.854-1.855a.807.807 0 0 1 0-1.14l.618-.62a2.69 2.69 0 0 0 0-3.803l-1.21-1.211A2.56 2.56 0 0 0 7.978 4"/>`), g.Group(children),
	)
}

func PieChartSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M13.5 2a7 7 0 0 0-.5 0a1 1 0 0 0-1 1v8c0 .6.4 1 1 1h8c.5 0 1-.4 1-1v-.5A8.5 8.5 0 0 0 13.5 2"/><path d="M11 6a1 1 0 0 0-1-1a8.5 8.5 0 1 0 9 9a1 1 0 0 0-1-1h-7z"/></g>`), g.Group(children),
	)
}

func PlayOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 18V6l8 6z"/>`), g.Group(children),
	)
}

func PlaySolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.6 5.2A1 1 0 0 0 7 6v12a1 1 0 0 0 1.6.8l8-6a1 1 0 0 0 0-1.6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlusOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14m-7 7V5"/>`), g.Group(children),
	)
}

func PrinterOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M16.444 18H19a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h2.556M17 11V5a1 1 0 0 0-1-1H8a1 1 0 0 0-1 1v6zM7 15h10v4a1 1 0 0 1-1 1H8a1 1 0 0 1-1-1z"/>`), g.Group(children),
	)
}

func PrinterSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 3a2 2 0 0 0-2 2v3h12V5a2 2 0 0 0-2-2zm-3 7a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h1v-4a1 1 0 0 1 1-1h10a1 1 0 0 1 1 1v4h1a2 2 0 0 0 2-2v-5a2 2 0 0 0-2-2zm4 11a1 1 0 0 1-1-1v-4h8v4a1 1 0 0 1-1 1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ProfileCardOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 9h3m-3 3h3m-3 3h3m-6 1c-.306-.613-.933-1-1.618-1H7.618c-.685 0-1.312.387-1.618 1M4 5h16a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1m7 5a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/>`), g.Group(children),
	)
}

func ProfileCardSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zm10 5a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1m0 3a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1m0 3a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1m-8-5a3 3 0 1 1 6 0a3 3 0 0 1-6 0m1.942 4a3 3 0 0 0-2.847 2.051l-.044.133l-.004.012c-.042.126-.055.167-.042.195c.006.013.02.023.038.039c.032.025.08.064.146.155A1 1 0 0 0 6 17h6a1 1 0 0 0 .811-.415a.7.7 0 0 1 .146-.155c.019-.016.031-.026.038-.04c.014-.027 0-.068-.042-.194l-.004-.012l-.044-.133A3 3 0 0 0 10.059 14z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func QrCodeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M4 4h6v6H4zm10 10h6v6h-6zm0-10h6v6h-6zm-4 10h.01v.01H10zm0 4h.01v.01H10zm-3 2h.01v.01H7zm0-4h.01v.01H7zm-3 2h.01v.01H4zm0-4h.01v.01H4z"/><path d="M7 7h.01v.01H7zm10 10h.01v.01H17z"/></g>`), g.Group(children),
	)
}

func QuestionCircleOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.529 9.988a2.502 2.502 0 1 1 5 .191A2.44 2.44 0 0 1 12 12.582V14m-.01 3.008H12M21 12a9 9 0 1 1-18 0a9 9 0 0 1 18 0"/>`), g.Group(children),
	)
}

func QuestionCircleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m9.008-3.018a1.502 1.502 0 0 1 2.522 1.159v.024a1.44 1.44 0 0 1-1.493 1.418a1 1 0 0 0-1.037.999V14a1 1 0 1 0 2 0v-.539a3.44 3.44 0 0 0 2.529-3.256a3.502 3.502 0 0 0-7-.255a1 1 0 0 0 2 .076c.014-.398.187-.774.48-1.044Zm.982 7.026a1 1 0 1 0 0 2H12a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func QuoteOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 11V8a1 1 0 0 0-1-1H6a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1m0 0v2a4 4 0 0 1-4 4H5m14-6V8a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1m0 0v2a4 4 0 0 1-4 4h-1"/>`), g.Group(children),
	)
}

func QuoteSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 6a2 2 0 0 0-2 2v3a2 2 0 0 0 2 2h3a3 3 0 0 1-3 3H5a1 1 0 1 0 0 2h1a5 5 0 0 0 5-5V8a2 2 0 0 0-2-2zm9 0a2 2 0 0 0-2 2v3a2 2 0 0 0 2 2h3a3 3 0 0 1-3 3h-1a1 1 0 1 0 0 2h1a5 5 0 0 0 5-5V8a2 2 0 0 0-2-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ReactSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M21.718 12c0-1.429-1.339-2.681-3.467-3.5c.029-.18.077-.37.1-.545c.217-2.058-.273-3.543-1.379-4.182c-1.235-.714-2.983-.186-4.751 1.239C10.45 3.589 8.7 3.061 7.468 3.773c-1.107.639-1.6 2.124-1.379 4.182c.018.175.067.365.095.545c-2.127.819-3.466 2.071-3.466 3.5s1.339 2.681 3.466 3.5c-.028.18-.077.37-.095.545c-.218 2.058.272 3.543 1.379 4.182c.376.213.803.322 1.235.316a6 6 0 0 0 3.514-1.56a6 6 0 0 0 3.515 1.56a2.44 2.44 0 0 0 1.236-.316c1.106-.639 1.6-2.124 1.379-4.182c-.019-.175-.067-.365-.1-.545c2.132-.819 3.471-2.071 3.471-3.5m-6.01-7.548a1.5 1.5 0 0 1 .76.187c.733.424 1.055 1.593.884 3.212c-.012.106-.043.222-.058.33q-1.263-.365-2.57-.523a16 16 0 0 0-1.747-1.972a4.9 4.9 0 0 1 2.731-1.234m-7.917 8.781c.172.34.335.68.529 1.017s.395.656.6.969a14 14 0 0 1-1.607-.376a14 14 0 0 1 .478-1.61m-.479-4.076a14 14 0 0 1 1.607-.376q-.308.468-.6.969c-.195.335-.357.677-.529 1.017q-.286-.79-.478-1.61M8.3 12a19 19 0 0 1 .888-1.75q.496-.852 1.076-1.65c.619-.061 1.27-.1 1.954-.1q1.025.001 1.952.1a20 20 0 0 1 1.079 1.654q.488.851.887 1.746a19 19 0 0 1-1.953 3.403a19.2 19.2 0 0 1-3.931 0a20 20 0 0 1-1.066-1.653A19 19 0 0 1 8.3 12m7.816 2.25c.2-.337.358-.677.53-1.017q.286.791.478 1.611a15 15 0 0 1-1.607.376c.202-.314.404-.635.597-.97zm.53-3.483c-.172-.34-.335-.68-.53-1.017a20 20 0 0 0-.6-.97q.814.142 1.606.376a14 14 0 0 1-.478 1.611zM12.217 6.34q.6.563 1.13 1.193q-.555-.031-1.129-.033c-.574-.002-.76.013-1.131.033q.53-.63 1.13-1.193m-4.249-1.7a1.5 1.5 0 0 1 .76-.187a4.9 4.9 0 0 1 2.729 1.233A16 16 0 0 0 9.71 7.658q-1.306.158-2.569.524c-.015-.109-.047-.225-.058-.331c-.171-1.619.151-2.787.885-3.211M3.718 12c0-.9.974-1.83 2.645-2.506c.218.857.504 1.695.856 2.506c-.352.811-.638 1.65-.856 2.506C4.692 13.83 3.718 12.9 3.718 12m4.25 7.361c-.734-.423-1.056-1.593-.885-3.212c.011-.106.043-.222.058-.331q1.262.365 2.564.524a16.4 16.4 0 0 0 1.757 1.982c-1.421 1.109-2.714 1.488-3.494 1.037m3.11-2.895q.56.033 1.14.034q.58-.001 1.139-.034a14 14 0 0 1-1.14 1.215a14 14 0 0 1-1.139-1.215m5.39 2.895c-.782.451-2.075.072-3.5-1.038a16 16 0 0 0 1.757-1.981a16.4 16.4 0 0 0 2.565-.523c.015.108.046.224.058.33c.175 1.619-.148 2.789-.88 3.212m1.6-4.854A16.6 16.6 0 0 0 17.216 12q.529-1.22.856-2.507c1.671.677 2.646 1.607 2.646 2.507s-.975 1.83-2.646 2.507z"/><path d="M12.215 13.773a1.792 1.792 0 1 0-1.786-1.8v.006a1.787 1.787 0 0 0 1.786 1.794"/></g>`), g.Group(children),
	)
}

func ReceiptOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 8h6m-6 4h6m-6 4h6M6 3v18l2-2l2 2l2-2l2 2l2-2l2 2V3l-2 2l-2-2l-2 2l-2-2l-2 2z"/>`), g.Group(children),
	)
}

func ReceiptSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.617 2.076a1 1 0 0 1 1.09.217L8 3.586l1.293-1.293a1 1 0 0 1 1.414 0L12 3.586l1.293-1.293a1 1 0 0 1 1.414 0L16 3.586l1.293-1.293A1 1 0 0 1 19 3v18a1 1 0 0 1-1.707.707L16 20.414l-1.293 1.293a1 1 0 0 1-1.414 0L12 20.414l-1.293 1.293a1 1 0 0 1-1.414 0L8 20.414l-1.293 1.293A1 1 0 0 1 5 21V3a1 1 0 0 1 .617-.924M9 7a1 1 0 0 0 0 2h6a1 1 0 1 0 0-2zm0 4a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2zm0 4a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func RectangleListOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 9h6m-6 3h6m-6 3h6M6.996 9h.01m-.01 3h.01m-.01 3h.01M4 5h16a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1"/>`), g.Group(children),
	)
}

func RectangleListSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm4.996 2a1 1 0 0 0 0 2h.01a1 1 0 1 0 0-2zM11 8a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2zm-4.004 3a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2zM11 11a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2zm-4.004 3a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2zM11 14a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func RedditSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M12.008 16.521a3.84 3.84 0 0 0 2.47-.77v.04a.28.28 0 0 0 .005-.396a.28.28 0 0 0-.395-.005a3.3 3.3 0 0 1-2.09.61a3.27 3.27 0 0 1-2.081-.63a.27.27 0 0 0-.38.381a3.84 3.84 0 0 0 2.47.77Z"/><path fill-rule="evenodd" d="M22 12c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2s10 4.477 10 10m-4.845-1.407A1.463 1.463 0 0 1 18.67 12a1.46 1.46 0 0 1-.808 1.33q.015.22 0 .44c0 2.242-2.61 4.061-5.829 4.061s-5.83-1.821-5.83-4.061a3 3 0 0 1 0-.44a1.458 1.458 0 0 1-.457-2.327a1.46 1.46 0 0 1 2.063-.064a7.16 7.16 0 0 1 3.9-1.23l.738-3.47v-.006a.31.31 0 0 1 .37-.236l2.452.49a1 1 0 1 1-.132.611l-2.14-.45l-.649 3.12a7.1 7.1 0 0 1 3.85 1.23c.259-.246.6-.393.957-.405" clip-rule="evenodd"/><path d="M15.305 13a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-4.625 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/></g>`), g.Group(children),
	)
}

func RedoOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 9H8a5 5 0 0 0 0 10h9m4-10l-4-4m4 4l-4 4"/>`), g.Group(children),
	)
}

func RedoSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M4 9h12a5 5 0 0 1 5 5v0a5 5 0 0 1-5 5H7"/><path stroke-linejoin="round" d="M7 5L3 9l4 4"/></g>`), g.Group(children),
	)
}

func RefreshOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.651 7.65a7.131 7.131 0 0 0-12.68 3.15M18.001 4v4h-4m-7.652 8.35a7.13 7.13 0 0 0 12.68-3.15M6 20v-4h4"/>`), g.Group(children),
	)
}

func ReplyAllOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.757 6L3.24 10.95a1.05 1.05 0 0 0 0 1.549l5.611 5.088m5.73-3.214v1.615a.948.948 0 0 1-1.524.845l-5.108-4.251a1.1 1.1 0 0 1 0-1.646l5.108-4.251a.95.95 0 0 1 1.524.846v1.7c3.312 0 6 2.979 6 6.654v1.329a.7.7 0 0 1-1.345.353a5.17 5.17 0 0 0-4.652-3.191z"/>`), g.Group(children),
	)
}

func ReplyAllSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.484 6.743c.41-.368.443-1 .077-1.41a.99.99 0 0 0-1.405-.078L2.67 10.203l-.007.006A2.05 2.05 0 0 0 2 11.721a2.06 2.06 0 0 0 .662 1.51l5.584 5.09a.99.99 0 0 0 1.405-.07a1.003 1.003 0 0 0-.07-1.412l-5.577-5.082a.05.05 0 0 1 0-.072zm6.543 9.199v-.42a4.17 4.17 0 0 1 2.715 2.415c.154.382.44.695.806.88a1.683 1.683 0 0 0 2.167-.571c.214-.322.312-.707.279-1.092V15.88c0-3.77-2.526-7.039-5.966-7.573V7.57a1.96 1.96 0 0 0-.994-1.838a1.93 1.93 0 0 0-2.153.184L7.8 10.164l-.012.011l-.011.01a2.1 2.1 0 0 0-.703 1.57a2.1 2.1 0 0 0 .726 1.59l5.08 4.25a1.933 1.933 0 0 0 2.929-.614c.167-.32.242-.68.218-1.04Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ReplyOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.5 8.046H11V6.119c0-.921-.9-1.446-1.524-.894l-5.108 4.49a1.2 1.2 0 0 0 0 1.739l5.108 4.49c.624.556 1.524.027 1.524-.893v-1.928h2a3.023 3.023 0 0 1 3 3.046V19a5.593 5.593 0 0 0-1.5-10.954"/>`), g.Group(children),
	)
}

func ReplySolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.502 7.046h-2.5v-.928a2.12 2.12 0 0 0-1.199-1.954a1.83 1.83 0 0 0-1.984.311L3.71 8.965a2.2 2.2 0 0 0 0 3.24L8.82 16.7a1.83 1.83 0 0 0 1.985.31a2.12 2.12 0 0 0 1.199-1.959v-.928h1a2.025 2.025 0 0 1 1.999 2.047V19a1 1 0 0 0 1.275.961a6.59 6.59 0 0 0 4.662-7.22a6.59 6.59 0 0 0-6.437-5.695Z"/>`), g.Group(children),
	)
}

func RestoreWindowOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 11.5h13m-13 0V18a1 1 0 0 0 1 1h11a1 1 0 0 0 1-1v-6.5m-13 0V9a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v2.5M9 5h11a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1h-1"/>`), g.Group(children),
	)
}

func RocketOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.051 8.102l-3.778.322l-1.994 1.994a.94.94 0 0 0 .533 1.6l2.698.316m8.39 1.617l-.322 3.78l-1.994 1.994a.94.94 0 0 1-1.595-.533l-.4-2.652m8.166-11.174a1.37 1.37 0 0 0-1.12-1.12c-1.616-.279-4.906-.623-6.38.853c-1.671 1.672-5.211 8.015-6.31 10.023a.93.93 0 0 0 .162 1.111l.828.835l.833.832a.93.93 0 0 0 1.111.163c2.008-1.102 8.35-4.642 10.021-6.312c1.475-1.478 1.133-4.77.855-6.385m-2.961 3.722a1.88 1.88 0 1 1-3.76 0a1.88 1.88 0 0 1 3.76 0"/>`), g.Group(children),
	)
}

func RocketSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20.337 3.664c.213.212.354.486.404.782c.294 1.711.657 5.195-.906 6.76c-1.77 1.768-8.485 5.517-10.611 6.683a.99.99 0 0 1-1.176-.173l-.882-.88l-.877-.884a.99.99 0 0 1-.173-1.177c1.165-2.126 4.913-8.841 6.682-10.611c1.562-1.563 5.046-1.198 6.757-.904c.296.05.57.191.782.404M5.407 7.576l4-.341l-2.69 4.48l-2.857-.334a.996.996 0 0 1-.565-1.694zm11.357 7.02l-.34 4l-2.111 2.113a.996.996 0 0 1-1.69-.565l-.422-2.807zm.84-6.21a1.99 1.99 0 1 1-3.98 0a1.99 1.99 0 0 1 3.98 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func RotateSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17.7 7.7A7.1 7.1 0 0 0 5 10.8"/><path d="M18 4v4h-4m-7.7 8.3A7.1 7.1 0 0 0 19 13.2"/><path d="M6 20v-4h4"/></g>`), g.Group(children),
	)
}

func RulerCombinedOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7H7m2 3H7m2 3H7m4 2v2m3-2v2m3-2v2M4 5v14a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1h-9a1 1 0 0 1-1-1V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1"/>`), g.Group(children),
	)
}

func SalePercentOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.891 15.107L15.11 8.89m-5.183-.52h.01m3.089 7.254h.01M14.08 3.902a2.85 2.85 0 0 0 2.176.902a2.845 2.845 0 0 1 2.94 2.94a2.85 2.85 0 0 0 .901 2.176a2.847 2.847 0 0 1 0 4.16a2.85 2.85 0 0 0-.901 2.175a2.843 2.843 0 0 1-2.94 2.94a2.85 2.85 0 0 0-2.176.902a2.847 2.847 0 0 1-4.16 0a2.85 2.85 0 0 0-2.176-.902a2.845 2.845 0 0 1-2.94-2.94a2.85 2.85 0 0 0-.901-2.176a2.85 2.85 0 0 1 0-4.16a2.85 2.85 0 0 0 .901-2.176a2.845 2.845 0 0 1 2.941-2.94a2.85 2.85 0 0 0 2.176-.901a2.847 2.847 0 0 1 4.159 0"/>`), g.Group(children),
	)
}

func SalePercentSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20.29 8.567c.133.323.334.613.59.85v.002a3.536 3.536 0 0 1 0 5.166a2.44 2.44 0 0 0-.776 1.868a3.534 3.534 0 0 1-3.651 3.653a2.48 2.48 0 0 0-1.87.776a3.537 3.537 0 0 1-5.164 0a2.44 2.44 0 0 0-1.87-.776a3.533 3.533 0 0 1-3.653-3.654a2.44 2.44 0 0 0-.775-1.868a3.537 3.537 0 0 1 0-5.166a2.44 2.44 0 0 0 .775-1.87a3.55 3.55 0 0 1 1.033-2.62a3.6 3.6 0 0 1 2.62-1.032a2.4 2.4 0 0 0 1.87-.775a3.535 3.535 0 0 1 5.165 0a2.44 2.44 0 0 0 1.869.775a3.53 3.53 0 0 1 3.652 3.652c-.012.35.051.697.184 1.02ZM9.927 7.371a1 1 0 1 0 0 2h.01a1 1 0 0 0 0-2zm5.889 2.226a1 1 0 0 0-1.414-1.415L8.184 14.4a1 1 0 0 0 1.414 1.414zm-2.79 5.028a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ScaleBalancedOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.5 21h13M12 21V7m0 0a2 2 0 1 0 0-4a2 2 0 0 0 0 4m2-1.8c3.073.661 2.467 2.8 5 2.8M5 8c3.359 0 2.192-2.115 5.012-2.793M7 9.556V7.75m0 1.806l-1.95 4.393a.773.773 0 0 0 .37.962a.8.8 0 0 0 .362.089h2.436a.79.79 0 0 0 .643-.335a.78.78 0 0 0 .09-.716zm10 0V7.313m0 2.243l-1.95 4.393a.773.773 0 0 0 .37.962a.8.8 0 0 0 .362.089h2.436a.79.79 0 0 0 .643-.335a.78.78 0 0 0 .09-.716z"/>`), g.Group(children),
	)
}

func ScaleBalancedSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-2.952.462c-.483.19-.868.432-1.19.71c-.363.315-.638.677-.831.93l-.106.14c-.21.268-.36.418-.574.527C6.125 6.883 5.74 7 5 7a1 1 0 0 0 0 2q.545 0 1-.067v.41l-1.864 4.2a1.774 1.774 0 0 0 .821 2.255c.255.133.538.202.825.202h2.436a1.786 1.786 0 0 0 1.768-1.558a1.8 1.8 0 0 0-.122-.899L8 9.343V8.028c.2-.188.36-.38.495-.553q.093-.118.168-.217c.185-.24.311-.406.503-.571a2 2 0 0 1 .24-.177A3 3 0 0 0 11 7.829V20H5.5a1 1 0 1 0 0 2h13a1 1 0 1 0 0-2H13V7.83a3 3 0 0 0 1.63-1.387c.206.091.373.19.514.29c.31.219.532.465.811.78l.025.027l.02.023v1.78l-1.864 4.2a1.774 1.774 0 0 0 .821 2.255c.255.133.538.202.825.202h2.436a1.785 1.785 0 0 0 1.768-1.558a1.8 1.8 0 0 0-.122-.899L18 9.343v-.452q.451.108 1 .109a1 1 0 1 0 0-2c-.48 0-.731-.098-.899-.2c-.2-.12-.363-.293-.651-.617l-.024-.026c-.267-.3-.622-.7-1.127-1.057a5.2 5.2 0 0 0-1.355-.678a3.001 3.001 0 0 0-5.896.04" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SearchOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m21 21l-3.5-3.5M17 10a7 7 0 1 1-14 0a7 7 0 0 1 14 0Z"/>`), g.Group(children),
	)
}

func SearchSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M10 2a8 8 0 1 0 0 16a8 8 0 0 0 0-16"/><path fill-rule="evenodd" d="M21.707 21.707a1 1 0 0 1-1.414 0l-3.5-3.5a1 1 0 0 1 1.414-1.414l3.5 3.5a1 1 0 0 1 0 1.414" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func ServerOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1M5 12h14M5 12a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1m-2 3h.01M14 15h.01M17 9h.01M14 9h.01"/>`), g.Group(children),
	)
}

func ServerSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 5a2 2 0 0 0-2 2v3a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V7a2 2 0 0 0-2-2zm9 2a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2zm3 0a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2zM3 17v-3a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v3a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m11-2a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2zm3 0a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShareAllOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15.141 6l5.518 4.95a1.05 1.05 0 0 1 0 1.549l-5.612 5.088m-6.154-3.214v1.615a.95.95 0 0 0 1.525.845l5.108-4.251a1.1 1.1 0 0 0 0-1.646l-5.108-4.251a.95.95 0 0 0-1.525.846v1.7c-3.312 0-6 2.979-6 6.654v1.329a.7.7 0 0 0 1.344.353a5.17 5.17 0 0 1 4.652-3.191z"/>`), g.Group(children),
	)
}

func ShareAllSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.516 6.743c-.41-.368-.443-1-.077-1.41a.99.99 0 0 1 1.405-.078l5.487 4.948l.007.006A2.05 2.05 0 0 1 22 11.721a2.06 2.06 0 0 1-.662 1.51l-5.584 5.09a.99.99 0 0 1-1.404-.07a1.003 1.003 0 0 1 .068-1.412l5.578-5.082a.05.05 0 0 0 .015-.036a.05.05 0 0 0-.015-.036zm-6.543 9.199v-.42a4.17 4.17 0 0 0-2.715 2.415c-.154.382-.44.695-.806.88a1.683 1.683 0 0 1-2.167-.571a1.7 1.7 0 0 1-.279-1.092V15.88c0-3.77 2.526-7.039 5.967-7.573V7.57a1.96 1.96 0 0 1 .993-1.838a1.93 1.93 0 0 1 2.153.184l5.08 4.248l.012.011l.011.01a2.1 2.1 0 0 1 .703 1.57a2.1 2.1 0 0 1-.726 1.59l-5.08 4.25a1.933 1.933 0 0 1-2.929-.614a1.96 1.96 0 0 1-.217-1.04Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShareNodesOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M7.926 10.898L15 7.727m-7.074 5.39L15 16.29M8 12a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm12 5.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Zm0-11a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/>`), g.Group(children),
	)
}

func ShareNodesSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.5 3a3.5 3.5 0 0 0-3.456 4.06L8.143 9.704a3.5 3.5 0 1 0-.01 4.6l5.91 2.65a3.5 3.5 0 1 0 .863-1.805l-5.94-2.662a3.5 3.5 0 0 0 .002-.961l5.948-2.667A3.5 3.5 0 1 0 17.5 3"/>`), g.Group(children),
	)
}

func ShieldCheckOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 11.5L11 13l4-3.5M12 20a16.4 16.4 0 0 1-5.092-5.804A16.7 16.7 0 0 1 5 6.666L12 4l7 2.667a16.7 16.7 0 0 1-1.908 7.529A16.4 16.4 0 0 1 12 20"/>`), g.Group(children),
	)
}

func ShieldCheckSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.644 3.066a1 1 0 0 1 .712 0l7 2.666A1 1 0 0 1 20 6.68a17.7 17.7 0 0 1-2.023 7.98a17.4 17.4 0 0 1-5.402 6.158a1 1 0 0 1-1.15 0a17.4 17.4 0 0 1-5.403-6.157A17.7 17.7 0 0 1 4 6.68a1 1 0 0 1 .644-.949zm4.014 7.187a1 1 0 0 0-1.316-1.506l-3.296 2.884l-.839-.838a1 1 0 0 0-1.414 1.414l1.5 1.5a1 1 0 0 0 1.366.046l4-3.5Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShieldOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20a16.4 16.4 0 0 1-5.092-5.804A16.7 16.7 0 0 1 5 6.666L12 4l7 2.667a16.7 16.7 0 0 1-1.908 7.529A16.4 16.4 0 0 1 12 20"/>`), g.Group(children),
	)
}

func ShieldSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.356 3.066a1 1 0 0 0-.712 0l-7 2.666A1 1 0 0 0 4 6.68a17.7 17.7 0 0 0 2.022 7.98a17.4 17.4 0 0 0 5.403 6.158a1 1 0 0 0 1.15 0a17.4 17.4 0 0 0 5.402-6.157A17.7 17.7 0 0 0 20 6.68a1 1 0 0 0-.644-.949z"/>`), g.Group(children),
	)
}

func ShoppingBagOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 10V6a3 3 0 0 1 3-3v0a3 3 0 0 1 3 3v4m3-2l.917 11.923A1 1 0 0 1 17.92 21H6.08a1 1 0 0 1-.997-1.077L6 8z"/>`), g.Group(children),
	)
}

func ShoppingBagSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 7h-4v3a1 1 0 0 1-2 0V7H6a1 1 0 0 0-.997.923l-.917 11.924A2 2 0 0 0 6.08 22h11.84a2 2 0 0 0 1.994-2.153l-.917-11.924A1 1 0 0 0 18 7h-2v3a1 1 0 1 1-2 0zm-2-3a2 2 0 0 0-2 2v1H8V6a4 4 0 0 1 8 0v1h-2V6a2 2 0 0 0-2-2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShuffleOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13.484 9.166L15 7h5m0 0l-3-3m3 3l-3 3M4 17h4l1.577-2.253M4 7h4l7 10h5m0 0l-3 3m3-3l-3-3"/>`), g.Group(children),
	)
}

func ShuffleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.3 3.3a1 1 0 0 1 1.4 0l3 3c.4.4.4 1 0 1.4l-3 3a1 1 0 0 1-1.4-1.4L17.6 8h-2l-1.3 1.7a1 1 0 1 1-1.6-1.1l1.5-2.2c.2-.2.5-.4.8-.4h2.6l-1.3-1.3a1 1 0 0 1 0-1.4M3 7c0-.6.4-1 1-1h4c.3 0 .6.2.8.4l6.7 9.6h2l-1.2-1.3a1 1 0 0 1 1.4-1.4l3 3c.4.4.4 1 0 1.4l-3 3a1 1 0 0 1-1.4-1.4l1.3-1.3H15a1 1 0 0 1-.8-.4L7.5 8H4a1 1 0 0 1-1-1m7.2 7c.4.2.5.9.2 1.3l-1.6 2.3a1 1 0 0 1-.8.4H4a1 1 0 1 1 0-2h3.5l1.3-1.8a1 1 0 0 1 1.4-.3Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SortHorizontalOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 16h13M4 16l4-4m-4 4l4 4M20 8H7m13 0l-4 4m4-4l-4-4"/>`), g.Group(children),
	)
}

func SortOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 20V10m0 10l-3-3m3 3l3-3m5-13v10m0-10l3 3m-3-3l-3 3"/>`), g.Group(children),
	)
}

func SortSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M8 19.5V10m8-5.5V14"/><path stroke-linejoin="round" d="m5 17l3 3l3-3m8-10l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func StackoverflowSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M17 20v-5h2v6.988H3V15h1.98v5z"/><path d="m6.84 14.522l8.73 1.825l.369-1.755l-8.73-1.825zm1.155-4.323l8.083 3.764l.739-1.617l-8.083-3.787zm3.372-5.481L10.235 6.08l6.859 5.704l1.132-1.362zM15.57 17H6.655v2h8.915zM12.861 3.111l6.193 6.415l1.414-1.415l-6.43-6.177z"/></g>`), g.Group(children),
	)
}

func StarHalfOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="m12.25 20.825l4.247-2.436a1 1 0 0 0 .503-.867V4.03c0-.405-2.062 3.38-2.8 4.747a1 1 0 0 1-.807.523l-4.87.367c-.903.068-1.258 1.208-.55 1.776l3.576 2.878a1 1 0 0 1 .343 1.025l-1.11 4.366c-.217.856.701 1.553 1.468 1.113Z"/>`), g.Group(children),
	)
}

func StarHalfSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 4.024v-.005c0-.053.002-.353-.217-.632a1.01 1.01 0 0 0-1.176-.315c-.192.076-.315.193-.35.225c-.052.05-.094.1-.122.134a4 4 0 0 0-.31.457c-.207.343-.484.84-.773 1.375a169 169 0 0 0-1.606 3.074h-.002l-4.599.367c-1.775.14-2.495 2.339-1.143 3.488L6.17 15.14l-1.06 4.406c-.412 1.72 1.472 3.078 2.992 2.157l3.94-2.388c.592-.359.958-.996.958-1.692v-13.6Zm-2.002 0v.025z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func StarHalfStrokeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 4.392v14.832M8.476 9.38l-4.553.36c-.888.07-1.248 1.165-.572 1.737l3.47 2.934a.98.98 0 0 1 .322.98l-1.06 4.388c-.206.855.736 1.531 1.497 1.073l3.898-2.351c.32-.193.723-.193 1.044 0l3.898 2.351c.76.458 1.703-.218 1.497-1.073l-1.06-4.388a.98.98 0 0 1 .322-.98l3.47-2.934c.676-.572.316-1.667-.572-1.737l-4.553-.36a1 1 0 0 1-.845-.606l-1.754-4.165c-.342-.812-1.508-.812-1.85 0L9.321 8.774a1 1 0 0 1-.845.606Z"/>`), g.Group(children),
	)
}

func StarHalfStrokeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m13.001 19.927l2.896 1.773c1.52.93 3.405-.442 2.992-2.179l-1.06-4.452l3.468-2.978c1.353-1.162.633-3.382-1.142-3.525L15.603 8.2l-1.754-4.226A1.97 1.97 0 0 0 13 3zM10.999 3c-.36.205-.663.53-.848.974L8.397 8.2l-4.552.366c-1.775.143-2.495 2.363-1.142 3.525l3.468 2.978l-1.06 4.452c-.413 1.737 1.472 3.11 2.992 2.178l2.896-1.773z"/>`), g.Group(children),
	)
}

func StarOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M11.083 5.104c.35-.8 1.485-.8 1.834 0l1.752 4.022a1 1 0 0 0 .84.597l4.463.342c.9.069 1.255 1.2.556 1.771l-3.33 2.723a1 1 0 0 0-.337 1.016l1.03 4.119c.214.858-.71 1.552-1.474 1.106l-3.913-2.281a1 1 0 0 0-1.008 0L7.583 20.8c-.764.446-1.688-.248-1.474-1.106l1.03-4.119A1 1 0 0 0 6.8 14.56l-3.33-2.723c-.698-.571-.342-1.702.557-1.771l4.462-.342a1 1 0 0 0 .84-.597z"/>`), g.Group(children),
	)
}

func StarSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.849 4.22c-.684-1.626-3.014-1.626-3.698 0L8.397 8.387l-4.552.361c-1.775.14-2.495 2.331-1.142 3.477l3.468 2.937l-1.06 4.392c-.413 1.713 1.472 3.067 2.992 2.149L12 19.35l3.897 2.354c1.52.918 3.405-.436 2.992-2.15l-1.06-4.39l3.468-2.938c1.353-1.146.633-3.336-1.142-3.477l-4.552-.36z"/>`), g.Group(children),
	)
}

func StopOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="12" height="12" x="6" y="6" fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" rx="1"/>`), g.Group(children),
	)
}

func StopSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 5a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2z"/>`), g.Group(children),
	)
}

func StoreOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12c.263 0 .524-.06.767-.175a2 2 0 0 0 .65-.491c.186-.21.333-.46.433-.734s.15-.568.15-.864a2.4 2.4 0 0 0 .586 1.591c.375.422.884.659 1.414.659s1.04-.237 1.414-.659A2.4 2.4 0 0 0 12 9.736a2.4 2.4 0 0 0 .586 1.591c.375.422.884.659 1.414.659s1.04-.237 1.414-.659A2.4 2.4 0 0 0 16 9.736c0 .295.052.588.152.861s.248.521.434.73a2 2 0 0 0 .649.488a1.8 1.8 0 0 0 1.53 0a2 2 0 0 0 .65-.488c.185-.209.332-.457.433-.73s.152-.566.152-.861c0-.974-1.108-3.85-1.618-5.121A.98.98 0 0 0 17.466 4H6.456a.99.99 0 0 0-.93.645C5.045 5.962 4 8.905 4 9.736c.023.59.241 1.148.611 1.567c.37.418.865.667 1.389.697m0 0c.328 0 .651-.091.94-.266A2.1 2.1 0 0 0 7.66 11h.681a2.1 2.1 0 0 0 .718.734c.29.175.613.266.942.266s.651-.091.94-.266c.29-.174.537-.427.719-.734h.681a2.1 2.1 0 0 0 .719.734c.289.175.612.266.94.266c.329 0 .652-.091.942-.266c.29-.174.536-.427.718-.734h.681c.183.307.43.56.719.734c.29.174.613.266.941.266a1.8 1.8 0 0 0 1.06-.351M6 12a1.77 1.77 0 0 1-1.163-.476M5 12v7a1 1 0 0 0 1 1h2v-5h3v5h7a1 1 0 0 0 1-1v-7m-5 3v2h2v-2z"/>`), g.Group(children),
	)
}

func StoreSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.535 7.677c.313-.98.687-2.023.926-2.677H17.46c.253.63.646 1.64.977 2.61c.166.487.312.953.416 1.347c.11.42.148.675.148.779c0 .18-.032.355-.09.515c-.06.161-.144.3-.243.412q-.152.166-.324.245a.8.8 0 0 1-.686 0a1 1 0 0 1-.324-.245q-.152-.169-.242-.412a1.5 1.5 0 0 1-.091-.515a1 1 0 1 0-2 0a1.4 1.4 0 0 1-.333.927a.9.9 0 0 1-.667.323a.9.9 0 0 1-.667-.323A1.4 1.4 0 0 1 13 9.736a1 1 0 1 0-2 0a1.4 1.4 0 0 1-.333.927a.9.9 0 0 1-.667.323a.9.9 0 0 1-.667-.323A1.4 1.4 0 0 1 9 9.74v-.008a1 1 0 0 0-2 .003v.008a1.5 1.5 0 0 1-.18.712a1.2 1.2 0 0 1-.146.209l-.007.007a1 1 0 0 1-.325.248a.8.8 0 0 1-.316.08a.97.97 0 0 1-.563-.256a1 1 0 0 1-.102-.103A1.52 1.52 0 0 1 5 9.724v-.006a3 3 0 0 1 .029-.207q.035-.198.11-.49c.098-.385.237-.85.395-1.344ZM4 12.112a3.52 3.52 0 0 1-1-2.376c0-.349.098-.8.202-1.208c.112-.441.264-.95.428-1.46c.327-1.024.715-2.104.958-2.767A1.985 1.985 0 0 1 6.456 3h11.01c.803 0 1.539.481 1.844 1.243c.258.641.67 1.697 1.019 2.72a22 22 0 0 1 .457 1.487c.114.433.214.903.214 1.286c0 .412-.072.821-.214 1.207A3.3 3.3 0 0 1 20 12.16V19a2 2 0 0 1-2 2h-6a1 1 0 0 1-1-1v-4H8v4a1 1 0 0 1-1 1H6a2 2 0 0 1-2-2zM13 15a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SunOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5V3m0 18v-2M7.05 7.05L5.636 5.636m12.728 12.728L16.95 16.95M5 12H3m18 0h-2M7.05 16.95l-1.414 1.414M18.364 5.636L16.95 7.05M16 12a4 4 0 1 1-8 0a4 4 0 0 1 8 0"/>`), g.Group(children),
	)
}

func SunSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 3a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0zM6.343 4.929A1 1 0 0 0 4.93 6.343l1.414 1.414a1 1 0 0 0 1.414-1.414zm12.728 1.414a1 1 0 0 0-1.414-1.414l-1.414 1.414a1 1 0 0 0 1.414 1.414zM12 7a5 5 0 1 0 0 10a5 5 0 0 0 0-10m-9 4a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2zm16 0a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2zM7.757 17.657a1 1 0 1 0-1.414-1.414l-1.414 1.414a1 1 0 1 0 1.414 1.414zm9.9-1.414a1 1 0 0 0-1.414 1.414l1.414 1.414a1 1 0 0 0 1.414-1.414zM13 19a1 1 0 1 0-2 0v2a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SwatchbookOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M7.111 20A3.11 3.11 0 0 1 4 16.889v-12C4 4.398 4.398 4 4.889 4h4.444a.89.89 0 0 1 .89.889v12A3.11 3.11 0 0 1 7.11 20Zm0 0h12a.89.89 0 0 0 .889-.889v-4.444a.89.89 0 0 0-.889-.89h-4.389a.9.9 0 0 0-.62.253l-3.767 3.665a1 1 0 0 0-.146.185c-.868 1.433-1.581 1.858-3.078 2.12Zm0-3.556h.009m7.933-10.927l3.143 3.143a.89.89 0 0 1 0 1.257l-7.974 7.974v-8.8l3.574-3.574a.89.89 0 0 1 1.257 0Z"/>`), g.Group(children),
	)
}

func SwatchbookSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 14h-2.722L11 20.278a5.5 5.5 0 0 1-.9.722H20a1 1 0 0 0 1-1v-5a1 1 0 0 0-1-1M9 3H4a1 1 0 0 0-1 1v13.5a3.5 3.5 0 1 0 7 0V4a1 1 0 0 0-1-1M6.5 18.5a1 1 0 1 1 0-2a1 1 0 0 1 0 2M19.132 7.9L15.6 4.368a1 1 0 0 0-1.414 0L12 6.55v9.9l7.132-7.132a1 1 0 0 0 0-1.418"/>`), g.Group(children),
	)
}

func TableColumnOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 11h18m-9 0v8m-8 0h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Z"/>`), g.Group(children),
	)
}

func TableColumnSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zm0 8v6h7v-6zm16 6h-7v-6h7z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TableRowOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M3 11h18M3 15h18m-9-4v8m-8 0h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1Z"/>`), g.Group(children),
	)
}

func TableRowSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 6a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm2 8v-2h7v2zm0 2v2h7v-2zm9 2h7v-2h-7zm7-4v-2h-7v2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TabletOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 18h2M5.875 3h12.25c.483 0 .875.448.875 1v16c0 .552-.392 1-.875 1H5.875C5.392 21 5 20.552 5 20V4c0-.552.392-1 .875-1"/>`), g.Group(children),
	)
}

func TabletSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4c0-.975.718-2 1.875-2h12.25C19.282 2 20 3.025 20 4v16c0 .975-.718 2-1.875 2H5.875C4.718 22 4 20.975 4 20zm7 13a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TagOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.583 8.445h.01M10.86 19.71l-6.573-6.63a.993.993 0 0 1 0-1.4l7.329-7.394A.98.98 0 0 1 12.31 4l5.734.007A1.97 1.97 0 0 1 20 5.983v5.5a1 1 0 0 1-.316.727l-7.44 7.5a.974.974 0 0 1-1.384.001Z"/>`), g.Group(children),
	)
}

func TagSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.045 3.007L12.31 3a1.97 1.97 0 0 0-1.4.585l-7.33 7.394a2 2 0 0 0 0 2.805l6.573 6.631a1.96 1.96 0 0 0 1.4.585a1.97 1.97 0 0 0 1.4-.585l7.409-7.477A2 2 0 0 0 21 11.479v-5.5a2.97 2.97 0 0 0-2.955-2.972m-2.452 6.438a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`), g.Group(children),
	)
}

func TailwindSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.782 5.72a4.773 4.773 0 0 0-4.8 4.173a3.43 3.43 0 0 1 2.741-1.687c1.689 0 2.974 1.972 3.758 2.587a5.73 5.73 0 0 0 5.382.935c2-.638 2.934-2.865 3.137-3.921c-.969 1.379-2.44 2.207-4.259 1.231c-1.253-.673-2.19-3.438-5.959-3.318M6.8 11.979A4.77 4.77 0 0 0 2 16.151a3.43 3.43 0 0 1 2.745-1.687c1.689 0 2.974 1.972 3.758 2.587a5.73 5.73 0 0 0 5.382.935c2-.638 2.933-2.865 3.137-3.921c-.97 1.379-2.44 2.208-4.259 1.231c-1.253-.673-2.19-3.443-5.963-3.317"/>`), g.Group(children),
	)
}

func TerminalOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m8 9l3 3l-3 3m5 0h3M4 19h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1"/>`), g.Group(children),
	)
}

func TerminalSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1zm4.293 5.707a1 1 0 0 1 1.414-1.414l3 3a1 1 0 0 1 0 1.414l-3 3a1 1 0 0 1-1.414-1.414L9.586 12zM13 14a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TextSizeOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6.2V5h11v1.2M8 5v14m-3 0h6m2-6.8V11h8v1.2M17 11v8m-1.5 0h3"/>`), g.Group(children),
	)
}

func TextSizeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5c0-.6.4-1 1-1h11c.6 0 1 .4 1 1v1.2a1 1 0 1 1-2 0V6H9v12h2a1 1 0 1 1 0 2H5a1 1 0 1 1 0-2h2V6H4v.2a1 1 0 1 1-2 0zm10 6c0-.6.4-1 1-1h8c.6 0 1 .4 1 1v1.2a1 1 0 1 1-2 0V12h-2v6h.5a1 1 0 1 1 0 2h-3a1 1 0 1 1 0-2h.5v-6h-2v.2a1 1 0 1 1-2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TextSlashOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 6.2V5h12v1.2M7 19h6m.2-14l-1.677 6.523M9.6 19l1.029-4M5 5l6.523 6.523M19 19l-7.477-7.477"/>`), g.Group(children),
	)
}

func TextSlashSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.2 4H7a1 1 0 0 0-1 .6l-.3-.3a1 1 0 0 0-1.4 1.4l14 14a1 1 0 0 0 1.4-1.4l-7-7L14 6h4v.2a1 1 0 1 0 2 0V5c0-.6-.4-1-1-1zM11 9.6L12 6H8v.5l3 3Zm-.1 4.4c.5.2.8.7.7 1.2l-.7 2.8H13a1 1 0 1 1 0 2H7a1 1 0 0 1 0-2h1.8l.9-3.2a1 1 0 0 1 1.2-.8" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ThumbsDownOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 13c-.889.086-1.416.543-2.156 1.057a22.3 22.3 0 0 0-3.958 5.084a1.6 1.6 0 0 1-.582.628a1.55 1.55 0 0 1-1.466.087a1.6 1.6 0 0 1-.537-.406a1.67 1.67 0 0 1-.384-1.279l1.389-4.114M17 13h3V6.5A1.5 1.5 0 0 0 18.5 5v0A1.5 1.5 0 0 0 17 6.5zm-6.5 1H5.585c-.286 0-.372-.014-.626-.15a1.8 1.8 0 0 1-.637-.572a1.87 1.87 0 0 1-.215-1.673l2.098-6.4C6.462 4.48 6.632 4 7.88 4c2.302 0 4.79.943 6.67 1.475"/>`), g.Group(children),
	)
}

func ThumbsDownSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.97 14.316H5.004c-.322 0-.64-.08-.925-.232a2 2 0 0 1-.717-.645a2.1 2.1 0 0 1-.242-1.883l2.36-7.201C5.769 3.54 5.96 3 7.365 3c2.072 0 4.276.678 6.156 1.256c.473.145.925.284 1.35.404h.114v9.862a25.5 25.5 0 0 0-4.238 5.514c-.197.376-.516.67-.901.83a1.74 1.74 0 0 1-1.21.048a1.8 1.8 0 0 1-.96-.757a1.87 1.87 0 0 1-.269-1.211l1.562-4.63ZM19.822 14H17V6a2 2 0 1 1 4 0v6.823c0 .65-.527 1.177-1.177 1.177Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ThumbsUpOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 11c.889-.086 1.416-.543 2.156-1.057a22.3 22.3 0 0 0 3.958-5.084a1.6 1.6 0 0 1 .582-.628a1.55 1.55 0 0 1 1.466-.087c.205.095.388.233.537.406a1.64 1.64 0 0 1 .384 1.279l-1.388 4.114M7 11H4v6.5A1.5 1.5 0 0 0 5.5 19v0A1.5 1.5 0 0 0 7 17.5zm6.5-1h4.915c.286 0 .372.014.626.15s.472.332.637.572a1.87 1.87 0 0 1 .215 1.673l-2.098 6.4C17.538 19.52 17.368 20 16.12 20c-2.303 0-4.79-.943-6.67-1.475"/>`), g.Group(children),
	)
}

func ThumbsUpSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.03 9.684h3.965c.322 0 .64.08.925.232s.532.374.717.645a2.11 2.11 0 0 1 .242 1.883l-2.36 7.201c-.288.814-.48 1.355-1.884 1.355c-2.072 0-4.276-.677-6.157-1.256c-.472-.145-.924-.284-1.348-.404h-.115V9.478a25.5 25.5 0 0 0 4.238-5.514a1.8 1.8 0 0 1 .901-.83a1.74 1.74 0 0 1 1.21-.048c.396.13.736.397.96.757c.225.36.32.788.269 1.211zM4.177 10H7v8a2 2 0 1 1-4 0v-6.823C3 10.527 3.527 10 4.176 10Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TicketOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18.5 12A2.5 2.5 0 0 1 21 9.5V7a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v2.5a2.5 2.5 0 0 1 0 5V17a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-2.5a2.5 2.5 0 0 1-2.5-2.5"/>`), g.Group(children),
	)
}

func TicketSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 5a2 2 0 0 0-2 2v2.5a1 1 0 0 0 1 1a1.5 1.5 0 1 1 0 3a1 1 0 0 0-1 1V17a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2v-2.5a1 1 0 0 0-1-1a1.5 1.5 0 1 1 0-3a1 1 0 0 0 1-1V7a2 2 0 0 0-2-2z"/>`), g.Group(children),
	)
}

func TrashBinOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7h14m-9 3v8m4-8v8M10 3h4a1 1 0 0 1 1 1v3H9V4a1 1 0 0 1 1-1M6 7h12v13a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1z"/>`), g.Group(children),
	)
}

func TrashBinSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.586 2.586A2 2 0 0 1 10 2h4a2 2 0 0 1 2 2v2h3a1 1 0 1 1 0 2v12a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2V8a1 1 0 0 1 0-2h3V4a2 2 0 0 1 .586-1.414M10 6h4V4h-4zm1 4a1 1 0 1 0-2 0v8a1 1 0 1 0 2 0zm4 0a1 1 0 1 0-2 0v8a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TruckOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 7h6l2 4m-8-4v8m0-8V6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v9h2m8 0H9m4 0h2m4 0h2v-4m0 0h-5m3.5 5.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0m-10 0a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0"/>`), g.Group(children),
	)
}

func TruckSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4a2 2 0 0 0-2 2v9a1 1 0 0 0 1 1h.535a3.5 3.5 0 1 0 6.93 0h3.07a3.5 3.5 0 1 0 6.93 0H21a1 1 0 0 0 1-1v-4a1 1 0 0 0-.106-.447l-2-4A1 1 0 0 0 19 6h-5a2 2 0 0 0-2-2zm14.192 11.59l.016.02a1.5 1.5 0 1 1-.016-.021Zm-10 0l.016.02a1.5 1.5 0 1 1-.016-.021Zm5.806-5.572v-2.02h4.396l1 2.02z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TwitterSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 5.892a8.2 8.2 0 0 1-2.355.635a4.07 4.07 0 0 0 1.8-2.235a8.3 8.3 0 0 1-2.605.981A4.13 4.13 0 0 0 15.85 4a4.07 4.07 0 0 0-4.1 4.038q0 .466.105.919A11.7 11.7 0 0 1 3.4 4.734a4.006 4.006 0 0 0 1.268 5.392a4.2 4.2 0 0 1-1.859-.5v.05A4.06 4.06 0 0 0 6.1 13.635a4.2 4.2 0 0 1-1.856.07a4.11 4.11 0 0 0 3.831 2.807A8.36 8.36 0 0 1 2 18.184A11.73 11.73 0 0 0 8.291 20A11.5 11.5 0 0 0 19.964 8.5c0-.177 0-.349-.012-.523A8.1 8.1 0 0 0 22 5.892" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UndoOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9h13a5 5 0 0 1 0 10H7M3 9l4-4M3 9l4 4"/>`), g.Group(children),
	)
}

func UndoSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="M20 9H8a5 5 0 0 0-5 5v0a5 5 0 0 0 5 5h9"/><path stroke-linejoin="round" d="m17 5l4 4l-4 4"/></g>`), g.Group(children),
	)
}

func UploadOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v9m-5 0H5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-2M8 9l4-5l4 5m1 8h.01"/>`), g.Group(children),
	)
}

func UploadSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 3a1 1 0 0 1 .78.375l4 5a1 1 0 1 1-1.56 1.25L13 6.85V14a1 1 0 1 1-2 0V6.85L8.78 9.626a1 1 0 1 1-1.56-1.25l4-5A1 1 0 0 1 12 3M9 14v-1H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2h-4v1a3 3 0 1 1-6 0m8 2a1 1 0 1 0 0 2h.01a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UserAddOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12h4m-2 2v-4M4 18v-1a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v1a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1m8-10a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/>`), g.Group(children),
	)
}

func UserAddSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 4a4 4 0 1 0 0 8a4 4 0 0 0 0-8m-2 9a4 4 0 0 0-4 4v1a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-1a4 4 0 0 0-4-4zm8-1a1 1 0 0 1 1-1h1v-1a1 1 0 1 1 2 0v1h1a1 1 0 1 1 0 2h-1v1a1 1 0 1 1-2 0v-1h-1a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UserCircleOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m0 0a8.95 8.95 0 0 0 4.951-1.488A3.987 3.987 0 0 0 13 16h-2a3.987 3.987 0 0 0-3.951 3.512A8.95 8.95 0 0 0 12 21m3-11a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/>`), g.Group(children),
	)
}

func UserCircleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 20a7.97 7.97 0 0 1-5.002-1.756l.002.001v-.683c0-1.794 1.492-3.25 3.333-3.25h3.334c1.84 0 3.333 1.456 3.333 3.25v.683A7.97 7.97 0 0 1 12 20M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10c0 5.5-4.44 9.963-9.932 10h-.138C6.438 21.962 2 17.5 2 12m10-5c-1.84 0-3.333 1.455-3.333 3.25S10.159 13.5 12 13.5c1.84 0 3.333-1.455 3.333-3.25S13.841 7 12 7" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UserEditOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" stroke-width="2" d="M7 19H5a1 1 0 0 1-1-1v-1a3 3 0 0 1 3-3h1m4-6a3 3 0 1 1-6 0a3 3 0 0 1 6 0Zm7.441 1.559a1.907 1.907 0 0 1 0 2.698l-6.069 6.069L10 19l.674-3.372l6.07-6.07a1.907 1.907 0 0 1 2.697 0Z"/>`), g.Group(children),
	)
}

func UserEditSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 8a4 4 0 1 1 7.796 1.263l-2.533 2.534A4 4 0 0 1 5 8m4.06 5H7a4 4 0 0 0-4 4v1a2 2 0 0 0 2 2h2.172a3 3 0 0 1-.114-1.588l.674-3.372a3 3 0 0 1 .82-1.533zm9.032-5a2.9 2.9 0 0 0-2.056.852L9.967 14.92a1 1 0 0 0-.273.51l-.675 3.373a1 1 0 0 0 1.177 1.177l3.372-.675a1 1 0 0 0 .511-.273l6.07-6.07a2.91 2.91 0 0 0-.944-4.742A2.9 2.9 0 0 0 18.092 8" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UserHeadsetOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14.079 6.839a3 3 0 0 0-4.255.1M13 20h1.083A3.916 3.916 0 0 0 18 16.083V9A6 6 0 1 0 6 9v7m7 4v-1a1 1 0 0 0-1-1h-1a1 1 0 0 0-1 1v1a1 1 0 0 0 1 1h1a1 1 0 0 0 1-1m-7-4v-6H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2zm12-6h1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-1z"/>`), g.Group(children),
	)
}

func UserHeadsetSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2a7 7 0 0 0-7 7a3 3 0 0 0-3 3v2a3 3 0 0 0 3 3h1a1 1 0 0 0 1-1V9a5 5 0 1 1 10 0v7.083A2.92 2.92 0 0 1 14.083 19H14a2 2 0 0 0-2-2h-1a2 2 0 0 0-2 2v1a2 2 0 0 0 2 2h1a2 2 0 0 0 1.732-1h.351a4.92 4.92 0 0 0 4.83-4H19a3 3 0 0 0 3-3v-2a3 3 0 0 0-3-3a7 7 0 0 0-7-7m1.45 3.275a4 4 0 0 0-4.352.976a1 1 0 0 0 1.452 1.376a2 2 0 0 1 2.836-.067a1 1 0 1 0 1.386-1.442a4 4 0 0 0-1.321-.843Z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UserOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M7 17v1a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-1a3 3 0 0 0-3-3h-4a3 3 0 0 0-3 3Zm8-9a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/>`), g.Group(children),
	)
}

func UserRemoveOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 12h4M4 18v-1a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v1a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1m8-10a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/>`), g.Group(children),
	)
}

func UserRemoveSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 8a4 4 0 1 1 8 0a4 4 0 0 1-8 0m-2 9a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v1a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm13-6a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UserSettingsOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="square" stroke-linejoin="round" stroke-width="2" d="M10 19H5a1 1 0 0 1-1-1v-1a3 3 0 0 1 3-3h2m10 1a3 3 0 0 1-3 3m3-3a3 3 0 0 0-3-3m3 3h1m-4 3a3 3 0 0 1-3-3m3 3v1m-3-4a3 3 0 0 1 3-3m-3 3h-1m4-3v-1m-2.121 1.879l-.707-.707m5.656 5.656l-.707-.707m-4.242 0l-.707.707m5.656-5.656l-.707.707M12 8a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/>`), g.Group(children),
	)
}

func UserSettingsSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 10v1.126c.367.095.714.24 1.032.428l.796-.797l1.415 1.415l-.797.796c.188.318.333.665.428 1.032H21v2h-1.126c-.095.367-.24.714-.428 1.032l.797.796l-1.415 1.415l-.796-.797a4 4 0 0 1-1.032.428V20h-2v-1.126a4 4 0 0 1-1.032-.428l-.796.797l-1.415-1.415l.797-.796A4 4 0 0 1 12.126 16H11v-2h1.126c.095-.367.24-.714.428-1.032l-.797-.796l1.415-1.415l.796.797A4 4 0 0 1 15 11.126V10zm.406 3.578l.016.016c.354.358.574.85.578 1.392v.028a2 2 0 0 1-3.409 1.406l-.01-.012a2 2 0 0 1 2.826-2.83ZM5 8a4 4 0 1 1 7.938.703a7.03 7.03 0 0 0-3.235 3.235A4 4 0 0 1 5 8m4.29 5H7a4 4 0 0 0-4 4v1a2 2 0 0 0 2 2h6.101A6.98 6.98 0 0 1 9 15c0-.695.101-1.366.29-2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UserSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 4a4 4 0 1 0 0 8a4 4 0 0 0 0-8m-2 9a4 4 0 0 0-4 4v1a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-1a4 4 0 0 0-4-4z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UsersGroupOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M4.5 17H4a1 1 0 0 1-1-1a3 3 0 0 1 3-3h1m0-3.05A2.5 2.5 0 1 1 9 5.5M19.5 17h.5a1 1 0 0 0 1-1a3 3 0 0 0-3-3h-1m0-3.05a2.5 2.5 0 1 0-2-4.45m.5 13.5h-7a1 1 0 0 1-1-1a3 3 0 0 1 3-3h3a3 3 0 0 1 3 3a1 1 0 0 1-1 1Zm-1-9.5a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0Z"/>`), g.Group(children),
	)
}

func UsersGroupSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 6a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7m-1.5 8a4 4 0 0 0-4 4a2 2 0 0 0 2 2h7a2 2 0 0 0 2-2a4 4 0 0 0-4-4zm6.82-3.096a5.51 5.51 0 0 0-2.797-6.293a3.5 3.5 0 1 1 2.796 6.292ZM19.5 18h.5a2 2 0 0 0 2-2a4 4 0 0 0-4-4h-1.1a5.5 5.5 0 0 1-.471.762A6 6 0 0 1 19.5 18M4 7.5a3.5 3.5 0 0 1 5.477-2.889a5.5 5.5 0 0 0-2.796 6.293A3.5 3.5 0 0 1 4 7.5M7.1 12H6a4 4 0 0 0-4 4a2 2 0 0 0 2 2h.5a6 6 0 0 1 3.071-5.238A5.5 5.5 0 0 1 7.1 12" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UsersOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M16 19h4a1 1 0 0 0 1-1v-1a3 3 0 0 0-3-3h-2m-2.236-4a3 3 0 1 0 0-4M3 18v-1a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v1a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1Zm8-10a3 3 0 1 1-6 0a3 3 0 0 1 6 0Z"/>`), g.Group(children),
	)
}

func UsersSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 4a4 4 0 1 0 0 8a4 4 0 0 0 0-8m-2 9a4 4 0 0 0-4 4v1a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2v-1a4 4 0 0 0-4-4zm7.25-2.095c.478-.86.75-1.85.75-2.905a6 6 0 0 0-.75-2.906a4 4 0 1 1 0 5.811M15.466 20c.34-.588.535-1.271.535-2v-1a5.98 5.98 0 0 0-1.528-4H18a4 4 0 0 1 4 4v1a2 2 0 0 1-2 2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func VideoCameraOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 6H4a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1m7 11l-6-2V9l6-2z"/>`), g.Group(children),
	)
}

func VideoCameraSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 7a2 2 0 0 0-2-2H4a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2zm2 9.387l4.684 1.562A1 1 0 0 0 22 17V7a1 1 0 0 0-1.316-.949L16 7.613z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func VolumeDownOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17.5 8.43A5 5 0 0 1 19 12a4.98 4.98 0 0 1-1.43 3.5M14 6.135v11.73a1 1 0 0 1-1.64.768L8 15H6a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h2l4.36-3.633a1 1 0 0 1 1.64.768"/>`), g.Group(children),
	)
}

func VolumeDownSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M15 6.037c0-1.724-1.978-2.665-3.28-1.562L7.638 7.933H6c-1.105 0-2 .91-2 2.034v4.066c0 1.123.895 2.034 2 2.034h1.638l4.082 3.458c1.302 1.104 3.28.162 3.28-1.562z"/><path fill-rule="evenodd" d="M16.786 7.658a.99.99 0 0 1 1.414-.014A6.14 6.14 0 0 1 20 12c0 1.662-.655 3.17-1.715 4.27a.99.99 0 0 1-1.414.014a1.03 1.03 0 0 1-.014-1.437A4.1 4.1 0 0 0 18 12a4.1 4.1 0 0 0-1.2-2.904a1.03 1.03 0 0 1-.014-1.438" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func VolumeMuteOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.5 8.43A5 5 0 0 1 17 12c0 1.126-.5 2.5-1.5 3.5m2.864-9.864A8.97 8.97 0 0 1 21 12c0 2.023-.5 4.5-2.5 6M7.8 7.5l2.56-2.133a1 1 0 0 1 1.64.768V12m0 4.5v1.365a1 1 0 0 1-1.64.768L6 15H4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1m1-4l14 14"/>`), g.Group(children),
	)
}

func VolumeMuteSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.707 4.293a1 1 0 0 0-1.414 1.414l14 14a1 1 0 0 0 1.414-1.414l-.004-.005C21.57 16.498 22 13.938 22 12a9.97 9.97 0 0 0-2.929-7.071a1 1 0 1 0-1.414 1.414A7.97 7.97 0 0 1 20 12c0 1.752-.403 3.636-1.712 4.873l-1.433-1.433C17.616 14.37 18 13.107 18 12c0-1.678-.69-3.197-1.8-4.285a1 1 0 1 0-1.4 1.428A4 4 0 0 1 16 12c0 .606-.195 1.335-.59 1.996L13 11.586V6.135c0-1.696-1.978-2.622-3.28-1.536L7.698 6.284l-1.99-1.991ZM4 8h.586L13 16.414v1.451c0 1.696-1.978 2.622-3.28 1.536L5.638 16H4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2"/>`), g.Group(children),
	)
}

func VolumeUpOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.5 8.43A5 5 0 0 1 17 12a4.98 4.98 0 0 1-1.43 3.5m2.794 2.864A8.97 8.97 0 0 0 21 12a8.97 8.97 0 0 0-2.636-6.364M12 6.135v11.73a1 1 0 0 1-1.64.768L6 15H4a1 1 0 0 1-1-1v-4a1 1 0 0 1 1-1h2l4.36-3.633a1 1 0 0 1 1.64.768"/>`), g.Group(children),
	)
}

func VolumeUpSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M13 6.037c0-1.724-1.978-2.665-3.28-1.562L5.638 7.933H4c-1.105 0-2 .91-2 2.034v4.066c0 1.123.895 2.034 2 2.034h1.638l4.082 3.458c1.302 1.104 3.28.162 3.28-1.562z"/><path fill-rule="evenodd" d="M14.786 7.658a.99.99 0 0 1 1.414-.014A6.14 6.14 0 0 1 18 12c0 1.662-.655 3.17-1.715 4.27a.99.99 0 0 1-1.414.014a1.03 1.03 0 0 1-.014-1.437A4.1 4.1 0 0 0 16 12a4.1 4.1 0 0 0-1.2-2.904a1.03 1.03 0 0 1-.014-1.438" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17.657 4.811a.99.99 0 0 1 1.414 0A10.22 10.22 0 0 1 22 12c0 2.807-1.12 5.35-2.929 7.189a.99.99 0 0 1-1.414 0a1.03 1.03 0 0 1 0-1.438A8.17 8.17 0 0 0 20 12a8.17 8.17 0 0 0-2.343-5.751a1.03 1.03 0 0 1 0-1.438" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func VueSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.5 3L12 7.156L9.857 3H2l10 18L22 3zM4.486 4.5h2.4L12 13.8l5.107-9.3h2.4L12 18.021z"/>`), g.Group(children),
	)
}

func WalletOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 8H5m12 0a1 1 0 0 1 1 1v2.6M17 8l-4-4M5 8a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1v-2.6M5 8l4-4l4 4m6 4h-4a2 2 0 1 0 0 4h4a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1"/>`), g.Group(children),
	)
}

func WalletSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M12 14a3 3 0 0 1 3-3h4a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-4a3 3 0 0 1-3-3m3-1a1 1 0 1 0 0 2h4v-2z"/><path d="M12.293 3.293a1 1 0 0 1 1.414 0L16.414 6h-2.828l-1.293-1.293a1 1 0 0 1 0-1.414M12.414 6L9.707 3.293a1 1 0 0 0-1.414 0L5.586 6zM4.586 7l-.056.055A2 2 0 0 0 3 9v10a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2h-4a5 5 0 0 1 0-10h4a2 2 0 0 0-1.53-1.945L17.414 7z"/></g>`), g.Group(children),
	)
}

func WandMagicSparklesOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.872 9.687L20 6.56L17.44 4L4 17.44L6.56 20L16.873 9.687Zm0 0l-2.56-2.56M6 7v2m0 0v2m0-2H4m2 0h2m7 7v2m0 0v2m0-2h-2m2 0h2M8 4h.01v.01H8zm2 2h.01v.01H10zm2-2h.01v.01H12zm8 8h.01v.01H20zm-2 2h.01v.01H18zm2 2h.01v.01H20z"/>`), g.Group(children),
	)
}

func WandMagicSparklesSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M17.44 3a1 1 0 0 1 .707.293l2.56 2.56a1 1 0 0 1 0 1.414L18.194 9.78L14.22 5.806l2.513-2.513A1 1 0 0 1 17.44 3m-4.634 4.22l-9.513 9.513a1 1 0 0 0 0 1.414l2.56 2.56a1 1 0 0 0 1.414 0l9.513-9.513zM6 6a1 1 0 0 1 1 1v1h1a1 1 0 0 1 0 2H7v1a1 1 0 1 1-2 0v-1H4a1 1 0 0 1 0-2h1V7a1 1 0 0 1 1-1m9 9a1 1 0 0 1 1 1v1h1a1 1 0 1 1 0 2h-1v1a1 1 0 1 1-2 0v-1h-1a1 1 0 1 1 0-2h1v-1a1 1 0 0 1 1-1" clip-rule="evenodd"/><path d="M19 13h-2v2h2zM13 3h-2v2h2zm-2 2H9v2h2zM9 3H7v2h2zm12 8h-2v2h2zm0 4h-2v2h2z"/></g>`), g.Group(children),
	)
}

func WhatsappSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M12 4a8 8 0 0 0-6.895 12.06l.569.718l-.697 2.359l2.32-.648l.379.243A8 8 0 1 0 12 4M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10a9.96 9.96 0 0 1-5.016-1.347l-4.948 1.382l1.426-4.829l-.006-.007l-.033-.055A9.96 9.96 0 0 1 2 12" clip-rule="evenodd"/><path d="M16.735 13.492c-.038-.018-1.497-.736-1.756-.83a1 1 0 0 0-.34-.075c-.196 0-.362.098-.49.291c-.146.217-.587.732-.723.886c-.018.02-.042.045-.057.045c-.013 0-.239-.093-.307-.123c-1.564-.68-2.751-2.313-2.914-2.589c-.023-.04-.024-.057-.024-.057c.005-.021.058-.074.085-.101c.08-.079.166-.182.249-.283l.117-.14c.121-.14.175-.25.237-.375l.033-.066a.68.68 0 0 0-.02-.64c-.034-.069-.65-1.555-.715-1.711c-.158-.377-.366-.552-.655-.552c-.027 0 0 0-.112.005c-.137.005-.883.104-1.213.311c-.35.22-.94.924-.94 2.16c0 1.112.705 2.162 1.008 2.561l.041.06c1.161 1.695 2.608 2.951 4.074 3.537c1.412.564 2.081.63 2.461.63c.16 0 .288-.013.4-.024l.072-.007c.488-.043 1.56-.599 1.804-1.276c.192-.534.243-1.117.115-1.329c-.088-.144-.239-.216-.43-.308"/></g>`), g.Group(children),
	)
}

func WindowOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 8h.01M9 8h.01M12 8h.01M4 11h16M4 19h16a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1H4a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1"/>`), g.Group(children),
	)
}

func WindowRestoreSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M8 5a1 1 0 0 1 1-1h11a2 2 0 0 1 2 2v9a2 2 0 0 1-2 2h-1a1 1 0 1 1 0-2h1V6H9a1 1 0 0 1-1-1"/><path d="M4 7a2 2 0 0 0-2 2v9a2 2 0 0 0 2 2h11a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2zm0 11v-5.5h11V18z"/></g>`), g.Group(children),
	)
}

func WindowSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 4a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2zm16 7H4v7h16zM5 8a1 1 0 0 1 1-1h.01a1 1 0 0 1 0 2H6a1 1 0 0 1-1-1m4-1a1 1 0 0 0 0 2h.01a1 1 0 0 0 0-2zm2 1a1 1 0 0 1 1-1h.01a1 1 0 1 1 0 2H12a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func WindowsSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.005 12L3 6.408l6.8-.923v6.517H3.005ZM11 5.32L19.997 4v8H11zM20.067 13l-.069 8l-9.065-1.275L11 13zM9.8 19.58l-6.795-.931V13H9.8z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func XCircleSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12a10 10 0 1 1 20 0a10 10 0 0 1-20 0m7.7-3.7a1 1 0 0 0-1.4 1.4l2.3 2.3l-2.3 2.3a1 1 0 1 0 1.4 1.4l2.3-2.3l2.3 2.3a1 1 0 0 0 1.4-1.4L13.4 12l2.3-2.3a1 1 0 0 0-1.4-1.4L12 10.6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func XCompanySolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.8 10.5L20.7 2h-3l-5.3 6.5L7.7 2H1l7.8 11l-7.3 9h3l5.7-7l5.1 7H22zm-2.4 3l-1.4-2l-5.6-7.9h2.3l4.5 6.3l1.4 2l6 8.5h-2.3l-4.9-7Z"/>`), g.Group(children),
	)
}

func XSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.795 10.533L20.68 2h-3.073l-5.255 6.517L7.69 2H1l7.806 10.91L1.47 22h3.074l5.705-7.07L15.31 22H22zm-2.38 2.95L9.97 11.464L4.36 3.627h2.31l4.528 6.317l1.443 2.02l6.018 8.409h-2.31z"/>`), g.Group(children),
	)
}

func YoutubeSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21.7 8.037a4.26 4.26 0 0 0-.789-1.964a2.84 2.84 0 0 0-1.984-.839c-2.767-.2-6.926-.2-6.926-.2s-4.157 0-6.928.2a2.84 2.84 0 0 0-1.983.839a4.2 4.2 0 0 0-.79 1.965a30 30 0 0 0-.2 3.206v1.5a30 30 0 0 0 .2 3.206c.094.712.364 1.39.784 1.972c.604.536 1.38.837 2.187.848c1.583.151 6.731.2 6.731.2s4.161 0 6.928-.2a2.84 2.84 0 0 0 1.985-.84a4.3 4.3 0 0 0 .787-1.965a30 30 0 0 0 .2-3.206v-1.516a31 31 0 0 0-.202-3.206m-11.692 6.554v-5.62l5.4 2.819z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ZoomInOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m21 21l-3.5-3.5M10 7v6m-3-3h6m4 0a7 7 0 1 1-14 0a7 7 0 0 1 14 0Z"/>`), g.Group(children),
	)
}

func ZoomInSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21.707 21.707a1 1 0 0 1-1.414 0l-3.5-3.5a1 1 0 0 1 1.414-1.414l3.5 3.5a1 1 0 0 1 0 1.414M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0m9-3a1 1 0 1 0-2 0v2H7a1 1 0 0 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ZoomOutOutline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m21 21l-3.5-3.5M7 10h6m4 0a7 7 0 1 1-14 0a7 7 0 0 1 14 0Z"/>`), g.Group(children),
	)
}

func ZoomOutSolid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M21.707 21.707a1 1 0 0 1-1.414 0l-3.5-3.5a1 1 0 0 1 1.414-1.414l3.5 3.5a1 1 0 0 1 0 1.414M2 10a8 8 0 1 1 16 0a8 8 0 0 1-16 0m4 0a1 1 0 0 0 1 1h6a1 1 0 1 0 0-2H7a1 1 0 0 0-1 1" clip-rule="evenodd"/>`), g.Group(children),
	)
}
