package rivet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const IconifyVersion = ""

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 0 0")
)

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M4.855.642A8.03 8.03 0 0 0 .642 4.855A3 3 0 0 1 4.855.642m10.503 4.213A3 3 0 0 0 11.145.642a8.03 8.03 0 0 1 4.213 4.213M9 7.586V5H7v3.414l2 2L10.414 9z"/><path d="M1 8a7 7 0 1 1 12.606 4.192l2.308 2.308l-1.414 1.414l-2.308-2.308A6.97 6.97 0 0 1 8 15a6.97 6.97 0 0 1-4.192-1.394L1.5 15.914L.086 14.5l2.308-2.308A6.97 6.97 0 0 1 1 8m7-5a5 5 0 1 0 0 10A5 5 0 0 0 8 3"/></g>`), g.Group(children),
	)
}

func AlarmSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M.642 4.855A8.03 8.03 0 0 1 4.855.642A3 3 0 0 0 .642 4.855m14.716 0A3 3 0 0 0 11.145.642a8.03 8.03 0 0 1 4.213 4.213"/><path d="M8 15a6.97 6.97 0 0 0 4.192-1.394l2.308 2.308l1.414-1.414l-2.308-2.308a7 7 0 1 0-11.213 0L.087 14.5L1.5 15.914l2.308-2.308A6.97 6.97 0 0 0 8 15m2.414-6L9 10.414l-2-2V5h2v2.586z"/></g>`), g.Group(children),
	)
}

func ArrowAnchorDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 10V2h2v10H4.414l1.793 1.793l-1.414 1.414L.586 11l4.207-4.207l1.414 1.414L4.414 10z"/>`), g.Group(children),
	)
}

func ArrowAnchorDownRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2h2v8h8.586L9.793 8.207l1.414-1.414L15.414 11l-4.207 4.207l-1.414-1.414L11.586 12H1z"/>`), g.Group(children),
	)
}

func ArrowAnchorUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m4.414 6l1.793 1.793l-1.414 1.414L.586 5L4.793.793l1.414 1.414L4.414 4H15v10h-2V6z"/>`), g.Group(children),
	)
}

func ArrowAnchorUpRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.207.793L15.414 5l-4.207 4.207l-1.414-1.414L11.586 6H3v8H1V4h10.586L9.793 2.207z"/>`), g.Group(children),
	)
}

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 1v10.844L2.146 7.737L.854 9.263L8 15.31l7.146-6.047l-1.292-1.526L9 11.844V1z"/>`), g.Group(children),
	)
}

func ArrowDownLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.293 2.293L4 10.586V4H2v10h10v-2H5.414l8.293-8.293z"/>`), g.Group(children),
	)
}

func ArrowDownRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.707 2.293L12 10.586V4h2v10H4v-2h6.586L2.293 3.707z"/>`), g.Group(children),
	)
}

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 7H4.156l4.107-4.854L6.737.854L.69 8l6.047 7.146l1.526-1.292L4.156 9H15z"/>`), g.Group(children),
	)
}

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 7h10.844L7.737 2.146L9.263.854L15.31 8l-6.047 7.146l-1.526-1.292L11.844 9H1z"/>`), g.Group(children),
	)
}

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 15V4.156l4.854 4.107l1.292-1.526L8 .69L.854 6.737l1.292 1.526L7 4.156V15z"/>`), g.Group(children),
	)
}

func ArrowUpLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h10v2H5.414l8.293 8.293l-1.414 1.414L4 5.414V12H2z"/>`), g.Group(children),
	)
}

func ArrowUpRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.586 4H4V2h10v10h-2V5.414l-8.293 8.293l-1.414-1.414z"/>`), g.Group(children),
	)
}

func Audio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M10 0v5.401a2.999 2.999 0 0 1 0 5.198V16l-5.333-4H0V4h4.667zM8 4L5.333 6H2v4h3.333L8 12z"/><path d="m13.6 3.2l-.799-.6L11.6 4.199l.8.6a4 4 0 0 1 .8.802c.503.668.8 1.497.8 2.399s-.297 1.73-.8 2.4a4 4 0 0 1-.8.8l-.8.601l1.201 1.6l.8-.601a6 6 0 0 0 1.198-1.2A5.98 5.98 0 0 0 16 8c0-1.35-.447-2.598-1.2-3.6a6 6 0 0 0-1.2-1.2"/></g>`), g.Group(children),
	)
}

func AudioOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 16l-5.333-4H0V4h3.667L9 0zM4.333 6H2v4h2.333L7 12V4zm7.253 2l-1.5 1.5l1.414 1.414l1.5-1.5l1.5 1.5L15.914 9.5l-1.5-1.5l1.5-1.5L14.5 5.086l-1.5 1.5l-1.5-1.5L10.086 6.5z"/>`), g.Group(children),
	)
}

func AudioOffSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 0v16l-5.333-4H0V4h3.667zm2.586 8l-1.5 1.5l1.414 1.414l1.5-1.5l1.5 1.5L15.914 9.5l-1.5-1.5l1.5-1.5L14.5 5.086l-1.5 1.5l-1.5-1.5L10.086 6.5z"/>`), g.Group(children),
	)
}

func AudioSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M10 16v-5.401A2.999 2.999 0 0 0 10 5.4V0L4.667 4H0v8h4.667z"/><path d="m13.6 3.2l-.799-.6L11.6 4.199l.8.6a4 4 0 0 1 .8.802c.503.668.8 1.497.8 2.399s-.297 1.73-.8 2.4a4 4 0 0 1-.8.8l-.8.601l1.201 1.6l.8-.601a6 6 0 0 0 1.198-1.2A5.98 5.98 0 0 0 16 8c0-1.35-.447-2.598-1.2-3.6a6 6 0 0 0-1.2-1.2"/></g>`), g.Group(children),
	)
}

func Ban(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.109 4.523a6.002 6.002 0 0 0 8.368 8.368zM4.523 3.11l8.368 8.368A6.002 6.002 0 0 0 4.523 3.11m-2.18 10.548A8 8 0 1 1 13.657 2.343A8 8 0 0 1 2.343 13.657Z"/>`), g.Group(children),
	)
}

func BanSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 8a7.97 7.97 0 0 1-1.68 4.906a8 8 0 0 1-1.414 1.414A8 8 0 1 1 16 8m-3.11 3.476L4.524 3.11A6 6 0 0 0 3.11 4.523l8.367 8.368a6 6 0 0 0 1.415-1.415Z"/>`), g.Group(children),
	)
}

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 14h5.5v-2h-1.382l-.366-.733a2.9 2.9 0 0 1-.298-1.12l-.234-4.209A5.23 5.23 0 0 0 9 1.096V0H7v1.096a5.23 5.23 0 0 0-4.22 4.842l-.234 4.209a2.9 2.9 0 0 1-.298 1.12L1.882 12H.5v2H6a2 2 0 1 0 4 0m1.886-2H4.114a4.9 4.9 0 0 0 .429-1.742l.234-4.209a3.228 3.228 0 0 1 6.446 0l.234 4.209A4.9 4.9 0 0 0 11.886 12"/>`), g.Group(children),
	)
}

func BellSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M.5 14v-2h1.382l.366-.733c.175-.349.276-.73.298-1.12l.234-4.209A5.23 5.23 0 0 1 7 1.096V0h2v1.096a5.23 5.23 0 0 1 4.22 4.842l.234 4.209c.022.39.123.771.298 1.12l.366.733H15.5v2H10a2 2 0 1 1-4 0z"/>`), g.Group(children),
	)
}

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m13 15.563l-5-2.308l-5 2.308V1h10zM11 3H5v9.437l3-1.384l3 1.384z"/>`), g.Group(children),
	)
}

func BookmarkSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 15.563V1H3v14.563l5-2.308z"/>`), g.Group(children),
	)
}

func BrowserWindow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M6 5a1 1 0 1 1-2 0a1 1 0 0 1 2 0m3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/><path d="M15 1H1v14h14zM3 7V3h10v4zm0 6V9h10v4z"/></g>`), g.Group(children),
	)
}

func BrowserWindowSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h14v14H1zm5 3a1 1 0 1 0-2 0a1 1 0 0 0 2 0m3 0a1 1 0 1 0-2 0a1 1 0 0 0 2 0m4 3H3v2h10z"/>`), g.Group(children),
	)
}

func Building(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M5 4h2v2H5zm6 0H9v2h2zM5 7h2v2H5zm6 0H9v2h2z"/><path d="M2 0h12v14h1v2H1v-2h1zm2 14h3v-2a1 1 0 1 1 2 0v2h3V2H4z"/></g>`), g.Group(children),
	)
}

func BuildingSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 0h12v14h1v2H1v-2h1zm5 4H5v2h2zm4 0H9v2h2zM7 7H5v2h2zm4 0H9v2h2zm-2 5a1 1 0 1 0-2 0v2h2z"/>`), g.Group(children),
	)
}

func Bus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 0H1v14h2v1.5a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5V14h6v1.5a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5V14h2zM3.5 2h9a.5.5 0 0 1 .5.5v5a.5.5 0 0 1-.5.5h-9a.5.5 0 0 1-.5-.5v-5a.5.5 0 0 1 .5-.5M5 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0m7 1a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`), g.Group(children),
	)
}

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M7 9.5v2H5v-2zm4 2v-2H9v2z"/><path d="M5 0a1 1 0 0 1 1 1v1h4V1a1 1 0 1 1 2 0v1h3v13H1V2h3V1a1 1 0 0 1 1-1M3 4v2h10V4zm10 4H3v5h10z"/></g>`), g.Group(children),
	)
}

func CalendarSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 1a1 1 0 0 0-2 0v1H1v13h14V2h-3V1a1 1 0 1 0-2 0v1H6zM3 6h10v2H3zm2 6.5v-2h2v2zm4 0v-2h2v2z"/>`), g.Group(children),
	)
}

func Caution(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M7 10V6h2v4zm1 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M6.382 0h3.236L16 12.764V16H0v-3.236zm1.236 2L2 13.236V14h12v-.764L8.382 2z"/></g>`), g.Group(children),
	)
}

func CautionSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.382 0h3.236L16 12.764V16H0v-3.236zM7 6v4h2V6zm1 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`), g.Group(children),
	)
}

func Chat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M5 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2m4-1a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M16 1H0v12h3.586L6 15.414L8.414 13H16zM2 11V3h12v8H7.586L6 12.586L4.414 11z"/></g>`), g.Group(children),
	)
}

func ChatSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 1h16v12H8.414L6 15.414L3.586 13H0zm6 6a1 1 0 1 0-2 0a1 1 0 0 0 2 0m3 0a1 1 0 1 0-2 0a1 1 0 0 0 2 0m3 0a1 1 0 1 0-2 0a1 1 0 0 0 2 0"/>`), g.Group(children),
	)
}

func Check(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m14.914 4l-9.47 9.47L1.09 8.393L2.608 7.09l2.948 3.44L13.5 2.585z"/>`), g.Group(children),
	)
}

func CheckAll(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m4.261 13.49l6.614-10.822l-1.707-1.043l-5.43 8.885l-2.394-1.916l-1.25 1.562zM16 3h-4v2h4zm0 4h-6v2h6zm0 4v2H8v-2z"/>`), g.Group(children),
	)
}

func CheckCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M7 11.414L11.914 6.5L10.5 5.086L7 8.586l-1.5-1.5L4.086 8.5z"/><path d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8"/></g>`), g.Group(children),
	)
}

func CheckCircleBreakout(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M2 8a6 6 0 0 1 8.57-5.423L11.43.77A8 8 0 1 0 16 8h-2A6 6 0 0 1 2 8"/><path d="M8 11.414L15.914 3.5L14.5 2.086L8 8.586l-2-2L4.586 8z"/></g>`), g.Group(children),
	)
}

func CheckCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m7 3.414L11.914 6.5L10.5 5.086L7 8.586l-1.5-1.5L4.086 8.5z"/>`), g.Group(children),
	)
}

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15.146 6.263l-1.292-1.526L8 9.69L2.146 4.737L.854 6.263L8 12.31z"/>`), g.Group(children),
	)
}

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.737.854L3.69 8l6.047 7.146l1.526-1.292L6.31 8l4.953-5.854z"/>`), g.Group(children),
	)
}

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.263 15.146L12.31 8L6.263.854L4.737 2.146L9.69 8l-4.953 5.854z"/>`), g.Group(children),
	)
}

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.146 10.737L8 4.69L.854 10.737l1.292 1.526L8 7.31l5.854 4.953z"/>`), g.Group(children),
	)
}

func ChevronsLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M.586 8L7 14.414L8.414 13l-5-5l5-5L7 1.586z"/><path d="M6.586 8L13 14.414L14.414 13l-5-5l5-5L13 1.586z"/></g>`), g.Group(children),
	)
}

func ChevronsRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M9.414 8L3 1.586L1.586 3l5 5l-5 5L3 14.414z"/><path d="M15.414 8L9 1.586L7.586 3l5 5l-5 5L9 14.414z"/></g>`), g.Group(children),
	)
}

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 2a6 6 0 1 0 0 12A6 6 0 0 0 8 2M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8"/>`), g.Group(children),
	)
}

func CircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0"/>`), g.Group(children),
	)
}

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M5 7h6v2H5zm0 3h6v2H5z"/><path d="M6 0a2 2 0 0 0-2 2H1v14h14V2h-3a2 2 0 0 0-2-2zM3 4h1v1h8V4h1v10H3zm3-1V2h4v1z"/></g>`), g.Group(children),
	)
}

func ClipboardSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2h3v14H1V2zm6 1V2H6v1zM5 7v2h6V7zm0 3v2h6v-2z"/>`), g.Group(children),
	)
}

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M9 7.586V4H7v4.414l2.5 2.5L10.914 9.5z"/><path d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8"/></g>`), g.Group(children),
	)
}

func ClockSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m9-.414V4H7v4.414l2.5 2.5L10.914 9.5z"/>`), g.Group(children),
	)
}

func Close(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m3.5 2.086l4.5 4.5l4.5-4.5L13.914 3.5L9.414 8l4.5 4.5l-1.414 1.414l-4.5-4.5l-4.5 4.5L2.086 12.5l4.5-4.5l-4.5-4.5z"/>`), g.Group(children),
	)
}

func CloseCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="m8 6.586l-2-2L4.586 6l2 2l-2 2L6 11.414l2-2l2 2L11.414 10l-2-2l2-2L10 4.586z"/><path d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8"/></g>`), g.Group(children),
	)
}

func CloseCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m8-1.414l-2-2L4.586 6l2 2l-2 2L6 11.414l2-2l2 2L11.414 10l-2-2l2-2L10 4.586z"/>`), g.Group(children),
	)
}

func Code(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.926 14.412L.155 8l5.77-6.412l1.487 1.338L2.845 8l4.567 5.074zm4.148-12.824L15.845 8l-5.77 6.412l-1.487-1.338L13.155 8L8.588 2.926z"/>`), g.Group(children),
	)
}

func Collapse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m.586 14l3-3H1V9h6v6H5v-2.586l-3 3zM15.414 2l-3 3H15v2H9V1h2v2.586l3-3z"/>`), g.Group(children),
	)
}

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M11 0H0v11h3V9H2V2h7v1h2z"/><path d="M16 5H5v11h11zm-9 9V7h7v7z"/></g>`), g.Group(children),
	)
}

func CopySolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M11 0H0v11h3V9H2V2h7v1h2z"/><path d="M5 5h11v11H5z"/></g>`), g.Group(children),
	)
}

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 2h16v12H0zm2 2v2h10v2H2v4h12V4z"/>`), g.Group(children),
	)
}

func CreditCardSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 2h16v12H0V8h12V6H0z"/>`), g.Group(children),
	)
}

func Css(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 1a3 3 0 0 0-3 3v1.146A2 2 0 0 1 .894 6.935l-.341.17v1.79l.341.17A2 2 0 0 1 2 10.855V12a3 3 0 0 0 3 3h2v-2H5a1 1 0 0 1-1-1v-1.146A4 4 0 0 0 2.803 8A4 4 0 0 0 4 5.146V4a1 1 0 0 1 1-1h2V1zm6 14a3 3 0 0 0 3-3v-1.146a2 2 0 0 1 1.106-1.789l.341-.17v-1.79l-.341-.17A2 2 0 0 1 14 5.145V4a3 3 0 0 0-3-3H9v2h2a1 1 0 0 1 1 1v1.146A4 4 0 0 0 13.197 8A4 4 0 0 0 12 10.854V12a1 1 0 0 1-1 1H9v2z"/>`), g.Group(children),
	)
}

func Data(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 4C1 2.589 2.118 1.58 3.31.984C4.562.358 6.222 0 8 0s3.438.358 4.69.984C13.882 1.58 15 2.59 15 4v8c0 1.411-1.118 2.42-2.31 3.016C11.438 15.642 9.778 16 8 16s-3.438-.358-4.69-.984C2.118 14.42 1 13.41 1 12zm2 0c0 .246.225.737 1.205 1.227C5.125 5.687 6.465 6 8 6s2.876-.313 3.795-.773C12.775 4.737 13 4.246 13 4s-.225-.737-1.205-1.227C10.875 2.313 9.535 2 8 2s-2.876.313-3.795.773C3.225 3.263 3 3.754 3 4m0 6.85V12c0 .246.225.737 1.205 1.227c.92.46 2.26.773 3.795.773s2.876-.313 3.795-.773C12.775 12.737 13 12.246 13 12v-1.15a6 6 0 0 1-.31.166C11.438 11.642 9.778 12 8 12s-3.438-.358-4.69-.984A6 6 0 0 1 3 10.85M13 8V6.85a7 7 0 0 1-.31.166C11.438 7.642 9.778 8 8 8s-3.438-.358-4.69-.984A6 6 0 0 1 3 6.85V8c0 .246.225.737 1.205 1.227c.92.46 2.26.773 3.795.773s2.876-.313 3.795-.773C12.775 8.737 13 8.246 13 8"/>`), g.Group(children),
	)
}

func DataSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 12c0 1.411-1.118 2.42-2.31 3.016C11.438 15.642 9.778 16 8 16s-3.438-.358-4.69-.984C2.118 14.42 1 13.41 1 12V4C1 2.589 2.118 1.58 3.31.984C4.562.358 6.222 0 8 0s3.438.358 4.69.984C13.882 1.58 15 2.59 15 4zM3 6.85q.153.087.31.166C4.562 7.642 6.222 8 8 8s3.438-.358 4.69-.984q.157-.078.31-.166V4c0 .246-.225.737-1.205 1.227C10.875 5.687 9.535 6 8 6s-2.876-.313-3.795-.773C3.225 4.737 3 4.246 3 4zm8.795 2.377c-.92.46-2.26.773-3.795.773s-2.876-.313-3.795-.773C3.225 8.737 3 8.246 3 8v2.85q.153.088.31.166C4.562 11.642 6.222 12 8 12s3.438-.358 4.69-.984q.157-.079.31-.166V8c0 .246-.225.737-1.205 1.227"/>`), g.Group(children),
	)
}

func Device(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M1 0a1 1 0 0 0-1 1v9a1 1 0 0 0 1 1h5v1H2v2h6V9H2V2h10v4h2V1a1 1 0 0 0-1-1z"/><path d="M10 7a1 1 0 0 0-1 1v7a1 1 0 0 0 1 1h5a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1zm1 7V9h3v5z"/></g>`), g.Group(children),
	)
}

func DeviceSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M0 1a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v5h-4a2 2 0 0 0-2 2v6H2v-2h4v-1H1a1 1 0 0 1-1-1z"/><path d="M9 8a1 1 0 0 1 1-1h5a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-5a1 1 0 0 1-1-1z"/></g>`), g.Group(children),
	)
}

func Download(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M9 7.826V1H7v6.826L4.392 5.59L3.09 7.108L8 11.318l4.91-4.21l-1.302-1.518z"/><path d="M3 13v-3H1v3a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-3h-2v3z"/></g>`), g.Group(children),
	)
}

func Ellipsis(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 8a2 2 0 1 1-4 0a2 2 0 0 1 4 0m6 0a2 2 0 1 1-4 0a2 2 0 0 1 4 0m4 2a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`), g.Group(children),
	)
}

func EllipsisVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 2a2 2 0 1 0-4 0a2 2 0 0 0 4 0M8 6a2 2 0 1 1 0 4a2 2 0 0 1 0-4m0 6a2 2 0 1 1 0 4a2 2 0 0 1 0-4"/>`), g.Group(children),
	)
}

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 2v12h16V2zm2 2.023V4h12v.023L8 7.356zM2 6.31l6 3.333l6-3.333V12H2z"/>`), g.Group(children),
	)
}

func EnvelopeSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M0 2h16v.912L8 7.356L0 2.912z"/><path d="M0 5.2V14h16V5.2L8 9.644L0 5.199Z"/></g>`), g.Group(children),
	)
}

func ExclamationMark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 11V1H7v10zm-1 4a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`), g.Group(children),
	)
}

func ExclamationMarkCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M7 9V4h2v5zm1 3a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m8-6a6 6 0 1 0 0 12A6 6 0 0 0 8 2"/></g>`), g.Group(children),
	)
}

func ExclamationMarkCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m7 1h2V4H7zm2 2a1 1 0 1 0-2 0a1 1 0 0 0 2 0"/>`), g.Group(children),
	)
}

func Expand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.586 3l-3 3L10 7.414l3-3V7h2V1H9v2zM4.414 13l3-3L6 8.586l-3 3V9H1v6h6v-2z"/>`), g.Group(children),
	)
}

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M8 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path d="m15.356 7.478l.027.051l.235.471l-.236.47l-.026.052a13 13 0 0 1-.464.794a14 14 0 0 1-1.399 1.853C12.303 12.492 10.427 14 8 14s-4.302-1.508-5.493-2.831A14 14 0 0 1 .644 8.522l-.027-.051L.382 8l.235-.47a6 6 0 0 1 .125-.232a14 14 0 0 1 1.765-2.467C3.697 3.508 5.573 2 8 2s4.302 1.508 5.493 2.831a14 14 0 0 1 1.863 2.647m-12.558.768c.276.436.68 1.013 1.195 1.585C5.053 11.008 6.427 12 8 12s2.948-.992 4.007-2.169A12 12 0 0 0 13.354 8a12 12 0 0 0-1.347-1.831C10.947 4.992 9.573 4 8 4s-2.948.992-4.007 2.169A12 12 0 0 0 2.646 8q.068.113.152.246"/></g>`), g.Group(children),
	)
}

func EyeOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M.456 14.456L14.414.414l1.293 1.293l-2.657 2.657q.236.236.443.467a14 14 0 0 1 1.863 2.647q.094.18.183.363l.079.159l-.079.159a9 9 0 0 1-.28.543a14 14 0 0 1-1.765 2.467C12.301 12.492 10.426 14 8 14c-1.34 0-2.511-.46-3.492-1.094l-2.8 2.801l-1.252-1.25Zm11.18-8.678L9.932 7.482a2.002 2.002 0 0 1-2.45 2.45l-1.516 1.516C6.594 11.79 7.276 12 8 12c1.573 0 2.948-.992 4.007-2.169A12 12 0 0 0 13.354 8a12 12 0 0 0-1.347-1.831a10 10 0 0 0-.371-.39ZM1.78 10.287A14 14 0 0 1 .743 8.702l-.098-.18q-.094-.18-.183-.363L.382 8l.079-.159a16 16 0 0 1 .28-.543a14 14 0 0 1 1.765-2.467C3.699 3.508 5.574 2 8 2c.62 0 1.206.099 1.752.268L8.03 4H8c-1.573 0-2.948.992-4.007 2.169A12 12 0 0 0 2.646 8a12 12 0 0 0 .561.852z"/>`), g.Group(children),
	)
}

func EyeOffSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m14.293.293l1.414 1.414l-2.657 2.657q.236.236.443.467a14 14 0 0 1 1.863 2.647l.027.051l.007.016l.003.005l.001.002l.224.448l-.224.449l-.001.001l-.003.005l-.007.016a6 6 0 0 1-.125.23a14 14 0 0 1-1.765 2.467C12.303 12.492 10.427 14 8 14c-1.34 0-2.511-.46-3.492-1.094l-2.8 2.801l-1.415-1.414zm-4.36 7.19l-2.451 2.45a2.003 2.003 0 0 0 2.45-2.45Zm-.131-5.2A5.9 5.9 0 0 0 8 2C5.573 2 3.698 3.508 2.507 4.831A14 14 0 0 0 .644 7.478l-.026.051l-.008.016l-.003.005v.001c-.232.463-.091.182-.069.137L.382 8l.224.447l.001.003l.003.005l.008.016a6 6 0 0 0 .124.23a14 14 0 0 0 1.047 1.596z"/>`), g.Group(children),
	)
}

func EyeSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15.356 8.522l.027-.051l.235-.471l-.236-.47l-.026-.052a13 13 0 0 0-.464-.794a14 14 0 0 0-1.399-1.853C12.303 3.508 10.427 2 8 2S3.698 3.508 2.507 4.831A14 14 0 0 0 .644 7.478l-.027.051l-.16.322L.381 8l.235.47a5 5 0 0 0 .125.232a14 14 0 0 0 1.765 2.467C3.697 12.492 5.573 14 8 14s4.302-1.508 5.493-2.831a14 14 0 0 0 1.863-2.647M8 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/>`), g.Group(children),
	)
}

func File(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 1h8.414L14 4.586V15H2zm2 2v10h8V7.5H7.5V3zm5.5 0v2.5H12v-.086L9.586 3z"/>`), g.Group(children),
	)
}

func FileSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.414 1H2v14h12V4.586zM9.5 5.5V3h.086L12 5.414V5.5z"/>`), g.Group(children),
	)
}

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 1a3 3 0 0 0-2.83 2H0v2h3.17a3.001 3.001 0 0 0 5.66 0H16V3H8.83A3 3 0 0 0 6 1M5 4a1 1 0 1 1 2 0a1 1 0 0 1-2 0m5 5a3 3 0 0 0-2.83 2H0v2h7.17a3.001 3.001 0 0 0 5.66 0H16v-2h-3.17A3 3 0 0 0 10 9m-1 3a1 1 0 1 1 2 0a1 1 0 0 1-2 0"/>`), g.Group(children),
	)
}

func FilterSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 1a3 3 0 0 0-2.83 2H0v2h3.17a3.001 3.001 0 0 0 5.66 0H16V3H8.83A3 3 0 0 0 6 1m4 8a3 3 0 0 0-2.83 2H0v2h7.17a3.001 3.001 0 0 0 5.66 0H16v-2h-3.17A3 3 0 0 0 10 9"/>`), g.Group(children),
	)
}

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 1v1h9.387l-1.333 4l1.333 4H5v5H3V1zm0 3v4h6.613l-.667-2l.667-2z"/>`), g.Group(children),
	)
}

func FlagSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 2V1H3v14h2v-5h9.387l-1.333-4l1.333-4z"/>`), g.Group(children),
	)
}

func Gear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M8 10a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path d="M5 1.5A1.5 1.5 0 0 1 6.5 0h3A1.5 1.5 0 0 1 11 1.5v1.304l1.13-.652a1.5 1.5 0 0 1 2.048.549l1.5 2.598a1.5 1.5 0 0 1-.549 2.05L14 8l1.13.652a1.5 1.5 0 0 1 .548 2.049l-1.5 2.598a1.5 1.5 0 0 1-2.049.55L11 13.195V14.5A1.5 1.5 0 0 1 9.5 16h-3A1.5 1.5 0 0 1 5 14.5v-1.304l-1.13.652a1.5 1.5 0 0 1-2.048-.549l-1.5-2.598a1.5 1.5 0 0 1 .549-2.05L2 8L.87 7.348A1.5 1.5 0 0 1 .323 5.3l1.5-2.598a1.5 1.5 0 0 1 2.049-.55L5 2.805zM7 2v1.67q-.001.255-.076.476a4 4 0 0 0-1.722.996a1.5 1.5 0 0 1-.452-.173l-1.446-.835l-1 1.732l1.446.835c.146.084.271.188.375.304a4 4 0 0 0 0 1.99q-.156.176-.375.304l-1.446.835l1 1.732l1.446-.835a1.5 1.5 0 0 1 .452-.173a4 4 0 0 0 1.722.996q.075.222.076.476V14h2v-1.67q.001-.255.076-.476a4 4 0 0 0 1.722-.996q.23.047.452.173l1.446.835l1-1.732l-1.446-.835a1.5 1.5 0 0 1-.375-.304a4 4 0 0 0 0-1.99c.104-.116.229-.22.375-.304l1.446-.835l-1-1.732l-1.446.835a1.5 1.5 0 0 1-.452.173a4 4 0 0 0-1.722-.996A1.5 1.5 0 0 1 9 3.67V2z"/></g>`), g.Group(children),
	)
}

func GearSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.5 0A1.5 1.5 0 0 0 5 1.5v1.304l-1.13-.652a1.5 1.5 0 0 0-2.048.549l-1.5 2.598a1.5 1.5 0 0 0 .549 2.05L2 8l-1.13.652A1.5 1.5 0 0 0 .323 10.7l1.5 2.598a1.5 1.5 0 0 0 2.049.55L5 13.195V14.5A1.5 1.5 0 0 0 6.5 16h3a1.5 1.5 0 0 0 1.5-1.5v-1.304l1.13.652a1.5 1.5 0 0 0 2.048-.549l1.5-2.598a1.5 1.5 0 0 0-.549-2.05L14 8l1.13-.652a1.5 1.5 0 0 0 .548-2.049l-1.5-2.598a1.5 1.5 0 0 0-2.049-.55L11 2.805V1.5A1.5 1.5 0 0 0 9.5 0zM8 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/>`), g.Group(children),
	)
}

func Gears(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M3.83 1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v.245a3 3 0 0 0-1.485 1.392A3 3 0 0 0 3.59 4.134l-1 1.732A3 3 0 0 0 2.672 9a3 3 0 0 0-.47 1.346a1 1 0 0 1-1.068-.48l-1-1.732A1 1 0 0 1 .5 6.768l.202-.117a.26.26 0 0 0 .126-.245a5 5 0 0 1 0-.812a.26.26 0 0 0-.126-.245L.5 5.232a1 1 0 0 1-.366-1.366l1-1.732A1 1 0 0 1 2.5 1.768l.21.12c.085.05.192.044.274-.012q.33-.227.697-.4a.26.26 0 0 0 .15-.232z"/><path d="M9.019 3a1 1 0 0 0-1 1v.244c0 .1-.06.19-.15.232a5 5 0 0 0-.697.4a.26.26 0 0 1-.274.013l-.21-.121a1 1 0 0 0-1.366.366l-1 1.732a1 1 0 0 0 .366 1.366l.202.117c.086.05.134.145.126.245a5 5 0 0 0 0 .812c.008.1-.04.196-.126.245l-.202.117a1 1 0 0 0-.366 1.366l1 1.732a1 1 0 0 0 1.366.366l.21-.12a.26.26 0 0 1 .274.012q.33.227.697.4a.26.26 0 0 1 .15.232V14a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-.261c0-.1.058-.189.147-.231q.355-.171.678-.395a.26.26 0 0 1 .275-.013l.23.132a1 1 0 0 0 1.366-.366l1-1.732a1 1 0 0 0-.366-1.366l-.238-.137a.26.26 0 0 1-.126-.244a5 5 0 0 0 0-.774a.26.26 0 0 1 .126-.244l.238-.137a1 1 0 0 0 .366-1.366l-1-1.732a1 1 0 0 0-1.366-.366l-.23.132a.26.26 0 0 1-.275-.013a5 5 0 0 0-.678-.395a.26.26 0 0 1-.148-.23V4a1 1 0 0 0-1-1h-2ZM12 9a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/></g>`), g.Group(children),
	)
}

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m4.17-1a17 17 0 0 1 .624-3.67l.14-.488A6 6 0 0 0 2.083 7zm0 2H2.083a6 6 0 0 0 2.85 4.159l-.139-.489A17 17 0 0 1 4.17 9m3.07 4.952a6 6 0 0 0 1.52 0l.523-1.831c.292-1.023.473-2.07.543-3.121H6.174c.07 1.052.251 2.098.543 3.12l.524 1.832ZM6.175 7h3.652a15 15 0 0 0-.543-3.12L8.76 2.048a6 6 0 0 0-1.518 0l-.524 1.831A15 15 0 0 0 6.174 7Zm5.032 5.67l-.14.489A6 6 0 0 0 13.917 9H11.83a17 17 0 0 1-.624 3.67Zm0-9.34c.343 1.202.551 2.433.624 3.67h2.087a6 6 0 0 0-2.85-4.158z"/>`), g.Group(children),
	)
}

func GlobeSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.76.004L6.652 3.985c-.275.99-.446 2-.514 3.015H9.86a15 15 0 0 0-.514-3.015L8.24.004a8 8 0 0 0-.482 0ZM5.58.372L4.727 3.45A17 17 0 0 0 4.136 7H.061A8.01 8.01 0 0 1 5.58.372M.062 9a8.01 8.01 0 0 0 5.519 6.628l-.855-3.078A17 17 0 0 1 4.136 9H.061Zm7.698 6.996a8 8 0 0 0 .48 0l1.107-3.981A15 15 0 0 0 9.86 9H6.139a15 15 0 0 0 .514 3.015zm2.659-.368A8.01 8.01 0 0 0 15.939 9h-4.074a17 17 0 0 1-.591 3.55zM15.938 7A8.01 8.01 0 0 0 10.42.372l.855 3.078c.323 1.164.52 2.354.59 3.55z"/>`), g.Group(children),
	)
}

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 0H0v7h7zM2 5V2h3v3zm14-5H9v7h7zm-5 5V2h3v3zM9 9h7v7H9zm2 2v3h3v-3zM7 9H0v7h7zm-5 5v-3h3v3z"/>`), g.Group(children),
	)
}

func GridHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 3H1v2h2zm0 4H1v2h2zm-2 4h2v2H1zm6-8H5v2h2zM5 7h2v2H5zm2 4H5v2h2zm2-8h2v2H9zm6 0h-2v2h2zM9 7h2v2H9zm6 0h-2v2h2zm-6 4h2v2H9zm6 0h-2v2h2z"/>`), g.Group(children),
	)
}

func GridSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 0h7v7H0zm9 0h7v7H9zm7 9H9v7h7zM0 9h7v7H0z"/>`), g.Group(children),
	)
}

func GridVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 1v2H3V1zm4 2V1H7v2zm4 0V1h-2v2zm0 4V5h-2v2zM9 5v2H7V5zM5 7V5H3v2zm8 2v2h-2V9zm0 6v-2h-2v2zM9 9v2H7V9zm0 6v-2H7v2zM5 9v2H3V9zm0 6v-2H3v2z"/>`), g.Group(children),
	)
}

func Happy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M10 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2M7 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0M4.86 8.732a1 1 0 0 1 1.398.116l.005.005q.014.016.062.06c.064.057.169.142.315.23c.287.172.736.357 1.36.357s1.073-.185 1.36-.357a2 2 0 0 0 .377-.29l.005-.005a1 1 0 0 1 1.526 1.292l-.001.002l-.001.001l-.003.004l-.007.008l-.018.02a2.4 2.4 0 0 1-.225.224c-.145.13-.352.296-.623.459A4.6 4.6 0 0 1 8 11.5a4.6 4.6 0 0 1-2.39-.643a4 4 0 0 1-.795-.624l-.053-.057l-.018-.021l-.007-.008l-.003-.004v-.001l-.002-.001a1 1 0 0 1 .128-1.41Z"/><path d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8"/></g>`), g.Group(children),
	)
}

func HappySolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m11-2a1 1 0 1 0-2 0a1 1 0 0 0 2 0M6.258 8.848a1 1 0 0 0-1.526 1.292l.001.002l.001.001l.003.004l.007.008l.018.02l.053.058a4 4 0 0 0 .796.625A4.6 4.6 0 0 0 8 11.5a4.6 4.6 0 0 0 2.39-.643a4 4 0 0 0 .795-.624l.053-.057l.018-.021l.007-.008l.003-.004l.002-.002a1 1 0 0 0-1.526-1.293l-.005.005q-.014.016-.062.06a2 2 0 0 1-.314.23A2.6 2.6 0 0 1 8 9.5a2.6 2.6 0 0 1-1.36-.357a2 2 0 0 1-.377-.29zM7 6a1 1 0 1 0-2 0a1 1 0 0 0 2 0"/>`), g.Group(children),
	)
}

func Headphones(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0v8h-4a2 2 0 0 1-2-2v-4a2 2 0 0 1 2-2h2A6 6 0 0 0 2 8h2a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2H0zm2 2v4h2v-4zm12 0h-2v4h2z"/>`), g.Group(children),
	)
}

func HeadphonesSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 0a8 8 0 0 0-8 8v8h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2H2a6 6 0 1 1 12 0h-2a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4V8a8 8 0 0 0-8-8"/>`), g.Group(children),
	)
}

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 5.5L6.32 3.82a2.375 2.375 0 0 0-3.358 3.36L8 12.216l5.038-5.038A2.375 2.375 0 1 0 9.68 3.821zm0 9.546L1.548 8.594a4.375 4.375 0 0 1 6.187-6.188L8 2.672l.265-.266a4.375 4.375 0 1 1 6.187 6.188z"/>`), g.Group(children),
	)
}

func HeartSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" stroke="currentColor" stroke-width="2" d="M13.745 7.886L8 13.632L2.255 7.886a3.375 3.375 0 0 1 4.773-4.772l.265.265l.707.707l.707-.707l.265-.265a3.375 3.375 0 1 1 4.773 4.772Z"/>`), g.Group(children),
	)
}

func Home(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m8 .798l7 4.667V15H9v-4.444H7V15H1V5.465zM3 6.535V13h2V8.556h6V13h2V6.535L8 3.202z"/>`), g.Group(children),
	)
}

func HomeSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 5.465L8 .798L1 5.465V15h6v-4.444h2V15h6z"/>`), g.Group(children),
	)
}

func Image(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M10 8a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/><path d="M15 1H1v14h14zM3 6.92V3h10v7.132l-1.928-1.286l-1.914.957l-4.624-4.11L3 6.918Zm7.928 4.234L13 12.535V13H3V9.48l1.466-1.172l4.376 3.89z"/></g>`), g.Group(children),
	)
}

func ImageSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h14v14H1zm11 5a2 2 0 1 0-4 0a2 2 0 0 0 4 0m-9 7h10v-.465l-2.072-1.381l-2.086 1.043l-4.376-3.89L3 9.482z"/>`), g.Group(children),
	)
}

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 4.086V0H7v4.086l-1.5-1.5L4.086 4L8 7.914L11.914 4L10.5 2.586zM6 8H0v8h16V8h-6v2H6zm6 4v-2h2v4H2v-4h2v2z"/>`), g.Group(children),
	)
}

func InboxComplete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 6.914L13.414 1.5L12 .086l-4 4l-2-2L4.586 3.5zM6 8H0v8h16V8h-6v2H6zm6 4v-2h2v4H2v-4h2v2z"/>`), g.Group(children),
	)
}

func InboxCompleteSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.414 1.5L12 .086l-4 4l-2-2L4.586 3.5L8 6.914zM0 8h6v2h4V8h6v8H0z"/>`), g.Group(children),
	)
}

func InboxSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 4.086V0H7v4.086l-1.5-1.5L4.086 4L8 7.914L11.914 4L10.5 2.586zM0 8h6v2h4V8h6v8H0z"/>`), g.Group(children),
	)
}

func InfoCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M9 7v5H7V7zM8 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/><path d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m8-6a6 6 0 1 0 0 12A6 6 0 0 0 8 2"/></g>`), g.Group(children),
	)
}

func InfoCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m7-3a1 1 0 1 0 2 0a1 1 0 0 0-2 0m2 2H7v5h2z"/>`), g.Group(children),
	)
}

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3a2 2 0 0 1 2-2h10a2 2 0 0 1 2 2v6h1v6H0V9h1zm1 8v2h12v-2h-2v1H4v-1zm4-1h4V9h3V3H3v6h3z"/>`), g.Group(children),
	)
}

func LaptopSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1a2 2 0 0 0-2 2v6H0v6h16V9h-1V3a2 2 0 0 0-2-2zM2 13v-2h2v1h8v-1h2v2z"/>`), g.Group(children),
	)
}

func Lightning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.5 6.5L10 0H9L2 9.5h4.5L6 16h1l7-9.5z"/>`), g.Group(children),
	)
}

func LightningBox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M8.5 4L5 9h2v3h.5L11 7H9V4z"/><path d="M15 1H1v14h14zM3 13V3h10v10z"/></g>`), g.Group(children),
	)
}

func LightningBoxSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h14v14H1zm7.5 3L5 9h2v3h.5L11 7H9V4z"/>`), g.Group(children),
	)
}

func Link(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M12.243 3.757a2 2 0 0 0-2.829 0L7.293 5.88L5.879 4.464L8 2.344a4 4 0 0 1 5.657 0l.707.706l-.09.09A4 4 0 0 1 13.658 8l-2.121 2.121l-1.415-1.414l2.122-2.121a2 2 0 0 0 0-2.829Zm-8.486 8.486a2 2 0 0 0 2.829 0l2.121-2.122l1.414 1.415L8 13.655a4 4 0 0 1-5.657 0l-.707-.706l.09-.09A4 4 0 0 1 2.342 8l2.121-2.121L5.88 7.293L3.757 9.414a2 2 0 0 0 0 2.829"/><path d="M10.828 6.586L9.414 5.172L5.172 9.414l1.414 1.414z"/></g>`), g.Group(children),
	)
}

func LinkExternal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M15 1H9v2h2.586l-3 3L10 7.414l3-3V7h2z"/><path d="M7 3H1v12h12V9h-2v4H3V5h4z"/></g>`), g.Group(children),
	)
}

func List(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3a1 1 0 1 0 2 0a1 1 0 0 0-2 0m3 1h11V2H4zm11 5H4V7h11zM1 8a1 1 0 1 0 2 0a1 1 0 0 0-2 0m14 6H4v-2h11zM1 13a1 1 0 1 0 2 0a1 1 0 0 0-2 0"/>`), g.Group(children),
	)
}

func LockClosed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M9 10a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0z"/><path d="M7 0a3 3 0 0 0-3 3v2a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2V3a3 3 0 0 0-3-3zm3 5H6V3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1zm-6 9V7h8v7z"/></g>`), g.Group(children),
	)
}

func LockClosedSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 3a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v2a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2zm2 2h4V3a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1zm3 5a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0z"/>`), g.Group(children),
	)
}

func LockOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M9 10a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0z"/><path d="M7 0a3 3 0 0 0-3 3v2a2 2 0 0 0-2 2v7a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2H6V3a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3zM4 14V7h8v7z"/></g>`), g.Group(children),
	)
}

func LockOpenSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 3a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3h-2a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v2h6a2 2 0 0 1 2 2v7a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2zm5 7a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0z"/>`), g.Group(children),
	)
}

func MagnifyingGlass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 2a5 5 0 1 0 0 10A5 5 0 0 0 7 2M0 7a7 7 0 1 1 12.606 4.192l3.101 3.1l-1.414 1.415l-3.1-3.1A7 7 0 0 1 0 7"/>`), g.Group(children),
	)
}

func MapPin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M10 6.5a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/><path d="M7.193 15.162a18 18 0 0 1-2.666-2.283C3.1 11.386 1.5 9.145 1.5 6.5C1.5 3.245 4.141 0 8 0s6.5 3.245 6.5 6.5c0 2.645-1.6 4.886-3.027 6.379a18 18 0 0 1-2.666 2.283c-.263.183-.536.351-.807.523c-.27-.172-.544-.34-.807-.523M8 2C5.359 2 3.5 4.232 3.5 6.5c0 1.855 1.15 3.614 2.473 4.996A16 16 0 0 0 8 13.28a16 16 0 0 0 2.027-1.783C11.35 10.114 12.5 8.356 12.5 6.5C12.5 4.232 10.641 2 8 2"/></g>`), g.Group(children),
	)
}

func MapPinSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7.385 15.293l-.192-.13a18 18 0 0 1-2.666-2.283C3.1 11.385 1.5 9.144 1.5 6.499C1.5 3.245 4.141 0 8 0s6.5 3.245 6.5 6.5c0 2.645-1.6 4.886-3.027 6.379a18 18 0 0 1-2.666 2.283q-.122.085-.192.13c-.203.135-.41.263-.615.393c-.205-.13-.412-.258-.615-.392M8 8.5a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`), g.Group(children),
	)
}

func Megaphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 5.764V1.382L9.975 3.894A1 1 0 0 1 9.528 4H3a3 3 0 0 0-3 3v2a3 3 0 0 0 2.83 2.995L.59 14.61l1.518 1.301L5.46 12h4.068a1 1 0 0 1 .447.106L15 14.618v-4.382c.614-.55 1-1.347 1-2.236c0-.888-.386-1.687-1-2.236m-4.13-.08L13 4.617v6.764l-2.13-1.065A3 3 0 0 0 9.527 10H7V6h2.528a3 3 0 0 0 1.341-.317ZM5 6v4H3a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1z"/>`), g.Group(children),
	)
}

func MegaphoneSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 1.382v4.382c.614.55 1 1.348 1 2.236a3 3 0 0 1-1 2.236v4.382l-5.025-2.512A1 1 0 0 0 9.528 12H5.46l-3.352 3.91L.59 14.608l2.24-2.613A3 3 0 0 1 0 9V7a3 3 0 0 1 3-3h6.528a1 1 0 0 0 .447-.106zM7 10V6H5v4z"/>`), g.Group(children),
	)
}

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 4H1V2h14zm0 5H1V7h14zM1 14h14v-2H1z"/>`), g.Group(children),
	)
}

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M8 0a3 3 0 0 0-3 3v4a3 3 0 0 0 6 0V3a3 3 0 0 0-3-3M7 3a1 1 0 0 1 2 0v4a1 1 0 0 1-2 0z"/><path d="M8 11a4 4 0 0 1-4-4H2a6 6 0 0 0 5 5.917V14H5v2h6v-2H9v-1.083A6 6 0 0 0 14 7h-2a4 4 0 0 1-4 4"/></g>`), g.Group(children),
	)
}

func MicrophoneOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 0a3 3 0 0 1 2.761 1.825L9 3.585V3a1 1 0 0 0-2 0v2.586L5.048 7.538A3 3 0 0 1 5 7V3a3 3 0 0 1 3-3m-.538 9.952l-.814.814A4 4 0 0 0 12 7h2a6 6 0 0 1-5 5.917V14h2v2H5v-2h2v-1.083a6 6 0 0 1-1.861-.642l-3.432 3.432l-1.414-1.414l14-14l1.414 1.414L11 6.414V7a3 3 0 0 1-3.538 2.952M4 7c0 .474.083.93.234 1.352L2.724 9.86A6 6 0 0 1 2 7z"/>`), g.Group(children),
	)
}

func MicrophoneOffSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m14.293.293l1.414 1.414L11 6.414V7a3 3 0 0 1-3.538 2.952l-.814.814A4 4 0 0 0 12 7h2a6 6 0 0 1-5 5.917V14h2v2H5v-2h2v-1.083a6 6 0 0 1-1.861-.642l-3.432 3.432l-1.414-1.414zm-3.532 1.532A3 3 0 0 0 5 3v4q0 .276.048.538zM4.234 8.352A4 4 0 0 1 4 7H2c0 1.036.263 2.01.725 2.861l1.51-1.51Z"/>`), g.Group(children),
	)
}

func MicrophoneSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M5 3a3 3 0 0 1 6 0v4a3 3 0 0 1-6 0z"/><path d="M8 11a4 4 0 0 1-4-4H2a6 6 0 0 0 5 5.917V14H5v2h6v-2H9v-1.083A6 6 0 0 0 14 7h-2a4 4 0 0 1-4 4"/></g>`), g.Group(children),
	)
}

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 9H1V7h14z"/>`), g.Group(children),
	)
}

func MinusCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M12 7H4v2h8z"/><path d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8"/></g>`), g.Group(children),
	)
}

func MinusCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m12-1H4v2h8z"/>`), g.Group(children),
	)
}

func Money(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 2H5a3.5 3.5 0 1 0 0 7h2v3H2.5v2H7v2h2v-2h2a3.5 3.5 0 1 0 0-7H9V4h4.5V2H9V0H7zm2 7h2a1.5 1.5 0 0 1 0 3H9zM7 7H5a1.5 1.5 0 1 1 0-3h2z"/>`), g.Group(children),
	)
}

func Neutral(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M5.5 9a1 1 0 1 0 0 2h5a1 1 0 1 0 0-2zM11 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0M6 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8"/></g>`), g.Group(children),
	)
}

func NeutralSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m11-2a1 1 0 1 0-2 0a1 1 0 0 0 2 0m-6.5 4a1 1 0 0 0 1 1h5a1 1 0 1 0 0-2h-5a1 1 0 0 0-1 1M6 7a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`), g.Group(children),
	)
}

func Newspaper(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M8 5h4v2H8zm4 3H8v2h4z"/><path d="M16 1H4v2H0v9a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3zM5.83 13c.11-.313.17-.65.17-1V3h8v9a1 1 0 0 1-1 1zM2 12V5h2v7a1 1 0 1 1-2 0"/></g>`), g.Group(children),
	)
}

func NewspaperSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 1H4v2H0v9a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3zM4 12a1 1 0 1 1-2 0V5h2zm8-8v2H8V4zm0 3v2H8V7z"/>`), g.Group(children),
	)
}

func Note(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M10 8H6V6h4zm0 3H6V9h4z"/><path d="M10.414 1H2v14h12V4.586zM4 13V3h5.586L12 5.414V13z"/></g>`), g.Group(children),
	)
}

func NoteSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 1h8.414L14 4.586V15H2zm8 7V6H6v2zm0 3V9H6v2z"/>`), g.Group(children),
	)
}

func OrderedList(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 1h3v6H3V3H2zm6 4h7V3H8zm0 8h7v-2H8zM2 9h4v1.236L4.618 13H6v2H2v-1.236L3.382 11H2z"/>`), g.Group(children),
	)
}

func PageBottom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h14V1H1zm8 8.826V5H7v6.826L4.392 9.59L3.09 11.108L8 15.319l4.91-4.21l-1.302-1.518z"/>`), g.Group(children),
	)
}

func PageTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h14V1H1zm6 5.174V15h2V8.174l2.608 2.236l1.302-1.518L8 4.682l-4.91 4.21l1.302 1.518z"/>`), g.Group(children),
	)
}

func Parking(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M9 4H5v8h2v-2h2a3 3 0 1 0 0-6m0 4H7V6h2a1 1 0 0 1 0 2"/><path d="M15 1H1v14h14zM3 13V3h10v10z"/></g>`), g.Group(children),
	)
}

func ParkingSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M9 8H7V6h2a1 1 0 0 1 0 2"/><path d="M1 1h14v14H1zm8 3H5v8h2v-2h2a3 3 0 1 0 0-6"/></g>`), g.Group(children),
	)
}

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 15V1H4v14zm6 0V1h-2v14z"/>`), g.Group(children),
	)
}

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.353 2.938L13.06.646a1.5 1.5 0 0 0-2.122 0L1 10.586v4.359h4.413l9.937-9.882a1.5 1.5 0 0 0 .003-2.125m-4.439.562L12 2.414l1.584 1.584l-1.089 1.083zM9.5 4.914l1.577 1.577l-6.49 6.454H3v-1.53l6.5-6.5Z"/>`), g.Group(children),
	)
}

func PencilSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.707 1L15 3.292a1 1 0 0 1-.002 1.416l-1.441 1.434l-3.702-3.703L11.293 1a1 1 0 0 1 1.414 0M8.44 3.854L1.5 10.793v3.652h3.706l6.932-6.893z"/>`), g.Group(children),
	)
}

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 1h1l1.5 5L5 7c0 2 2 4 4 4l1-1.5l5 1.5v1a3 3 0 0 1-3 3C5 15 1 10 1 4a3 3 0 0 1 3-3"/>`), g.Group(children),
	)
}

func PhoneMobile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M6 13h4v-2H6z"/><path d="M4 0a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2zm3.5 3h1a2 2 0 0 0 1.732-1H12v12H4V2h1.768A2 2 0 0 0 7.5 3"/></g>`), g.Group(children),
	)
}

func PhoneMobileSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2zm4 11h4v-2H6z"/>`), g.Group(children),
	)
}

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 0h10v1c0 1.173-.294 2.075-.606 2.697a5 5 0 0 1-.394.656v2.175A6 6 0 0 1 14 11v1H9v3a1 1 0 1 1-2 0v-3H2v-1c0-1.778.774-3.375 2-4.472V4.353a5 5 0 0 1-.394-.656C3.294 3.075 3 2.173 3 1zm2.11 2c.076.332.183.6.284.803a2.7 2.7 0 0 0 .316.492l.006.006l.284.285v3.913l-.4.3A4 4 0 0 0 4.127 10h7.748A4 4 0 0 0 10.4 7.8l-.4-.3V3.585l.284-.284l.006-.007a2.7 2.7 0 0 0 .315-.492c.102-.204.209-.47.285-.803z"/>`), g.Group(children),
	)
}

func PinSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 0H3v1c0 1.173.294 2.075.606 2.697c.137.275.277.493.394.656v2.175A6 6 0 0 0 2 11v1h5v3a1 1 0 1 0 2 0v-3h5v-1a6 6 0 0 0-2-4.472V4.353c.117-.163.257-.381.394-.656C12.706 3.075 13 2.173 13 1z"/>`), g.Group(children),
	)
}

func Plane(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.662 8L.684 15.815L1.986 8L.684.185zM3.847 9l-.53 3.185L11.337 8L3.316 3.815L3.847 7H7v2z"/>`), g.Group(children),
	)
}

func PlaneSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.662 8L.684 15.815L1.82 9H7V7H1.82L.684.185z"/>`), g.Group(children),
	)
}

func Play(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.662 8L1 15.65V.35zm-4.324 0L3 3.65v8.7z"/>`), g.Group(children),
	)
}

func PlaySolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 15.65L15.662 8L1 .35z"/>`), g.Group(children),
	)
}

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 7V.5h2V7h6.5v2H9v6.5H7V9H.5V7z"/>`), g.Group(children),
	)
}

func PlusCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M7 4v3H4v2h3v3h2V9h3V7H9V4z"/><path d="M8 0a8 8 0 1 0 0 16A8 8 0 0 0 8 0M2 8a6 6 0 1 1 12 0A6 6 0 0 1 2 8"/></g>`), g.Group(children),
	)
}

func PlusCircleSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m7-4v3H4v2h3v3h2V9h3V7H9V4z"/>`), g.Group(children),
	)
}

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M12 9a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path d="M13 0H3v4H0v9h3v3h10v-3h3V4h-3zm-2 4H5V2h6zm0 10H5v-2h6zm-9-3V6h12v5h-1v-1H3v1z"/></g>`), g.Group(children),
	)
}

func PrinterSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 0H3v4H0v9h3v3h10v-3h3V4h-3zm0 8a1 1 0 1 1-2 0a1 1 0 0 1 2 0m-2-4H5V2h6zm0 10H5v-2h6z"/>`), g.Group(children),
	)
}

func QuestionMark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M9 11a1 1 0 1 1-2 0a1 1 0 0 1 2 0M7.5 4A2.5 2.5 0 0 0 5 6.5h2a.5.5 0 0 1 .5-.5h.646a.382.382 0 0 1 .17.724L7 7.382V9h2v-.382l.211-.106A2.382 2.382 0 0 0 8.146 4z"/><path d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m8-6a6 6 0 1 0 0 12A6 6 0 0 0 8 2"/></g>`), g.Group(children),
	)
}

func QuestionMarkSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m5-1.5h2a.5.5 0 0 1 .5-.5h.646a.382.382 0 0 1 .17.724L7 7.382V9h2v-.382l.211-.106A2.382 2.382 0 0 0 8.146 4H7.5A2.5 2.5 0 0 0 5 6.5M8 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`), g.Group(children),
	)
}

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 10.5v-1A4.5 4.5 0 0 1 4.5 5h7.586l-2-2L11.5 1.586L15.914 6L11.5 10.414L10.086 9l2-2H4.5a2.5 2.5 0 0 0 0 5H12v2H4.5a4.5 4.5 0 0 1-4.388-3.5z"/>`), g.Group(children),
	)
}

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M2 3c6.075 0 11 4.925 11 11h2C15 6.82 9.18 1 2 1z"/><path d="M2 5a9 9 0 0 1 9 9H9a7 7 0 0 0-7-7zm4 7a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/></g>`), g.Group(children),
	)
}

func Sad(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M8 2a6 6 0 1 0 0 12A6 6 0 0 0 8 2M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8"/><path d="M11 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0M7 6a1 1 0 1 1-2 0a1 1 0 0 1 2 0m2.742 5.152a1 1 0 0 0 1.526-1.292l-.001-.002l-.001-.001l-.003-.004l-.007-.008l-.018-.02a3 3 0 0 0-.225-.224a4 4 0 0 0-.623-.458A4.6 4.6 0 0 0 8 8.5a4.6 4.6 0 0 0-2.39.643a4 4 0 0 0-.795.624l-.053.057l-.018.021l-.007.008l-.003.004v.001l-.002.001a1 1 0 0 0 1.526 1.293l.005-.005q.014-.016.062-.06a2 2 0 0 1 .314-.23A2.6 2.6 0 0 1 8 10.5c.624 0 1.073.185 1.36.357a2 2 0 0 1 .382.295"/></g>`), g.Group(children),
	)
}

func SadSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8a8 8 0 1 1 16 0A8 8 0 0 1 0 8m6-1a1 1 0 1 0 0-2a1 1 0 0 0 0 2m5-1a1 1 0 1 0-2 0a1 1 0 0 0 2 0m-1.258 5.152a1 1 0 0 0 1.526-1.293h-.001l-.001-.002l-.003-.004l-.007-.008l-.018-.02a3 3 0 0 0-.225-.225a4 4 0 0 0-.623-.457A4.6 4.6 0 0 0 8 8.5a4.6 4.6 0 0 0-2.39.643a4 4 0 0 0-.795.624l-.053.057l-.018.021l-.007.008l-.003.004v.001l-.002.002a1 1 0 0 0 1.526 1.292l.005-.005q.014-.016.062-.06a2 2 0 0 1 .315-.23A2.6 2.6 0 0 1 8 10.5c.624 0 1.073.185 1.36.357a2 2 0 0 1 .377.29z"/>`), g.Group(children),
	)
}

func Save(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M9 4H4v2h5z"/><path d="M10.414 1H1v14h14V5.586zM3 13V3h6.586L13 6.414V13h-1V9H4v4zm7-2v2H6v-2z"/></g>`), g.Group(children),
	)
}

func SaveSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.414 1H1v14h14V5.586zM4 4h5v2H4zm8 5v4h-2v-2H6v2H4V9z"/>`), g.Group(children),
	)
}

func Settings(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 0v7.268a2 2 0 1 1-2 0V0zm0 12v4H2v-4zm5-9.732V0H7v2.268a2 2 0 1 0 2 0M9 16V7H7v9zm5-8.732V0h-2v7.268a2 2 0 1 0 2 0M14 16v-4h-2v4z"/>`), g.Group(children),
	)
}

func Share(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 2a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 .917 2.16l-4.94 2.47a3 3 0 0 1 0 .74l4.94 2.47a3 3 0 1 1-.895 1.789l-4.94-2.47a3 3 0 1 1 0-4.319l4.94-2.47A3 3 0 0 1 10 3M3 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2m10 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`), g.Group(children),
	)
}

func ShareSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 0a3 3 0 0 0-2.977 3.37l-4.94 2.47a3 3 0 1 0 0 4.319l4.94 2.47a3 3 0 1 0 .895-1.789l-4.94-2.47a3 3 0 0 0 0-.74l4.94-2.47A3 3 0 1 0 13 0"/>`), g.Group(children),
	)
}

func Shirt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m5 .882l.447.224c1.719.859 3.387.859 5.106 0L11 .882l5 2.5V9h-3v6H3V9H0V3.382zm.016 2.228L2 4.618V7h3v6h6V7h3V4.618L10.984 3.11c-1.979.853-3.99.853-5.968 0"/>`), g.Group(children),
	)
}

func ShirtSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.447 1.106L5 .882l-5 2.5V9h3v6h10V9h3V3.382l-5-2.5l-.447.224c-1.719.859-3.387.859-5.106 0"/>`), g.Group(children),
	)
}

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 4a4 4 0 1 1 8 0h4v12H0V4zm0 2H2v8h12V6h-2v2h-2V6H6v2H4zm6-2a2 2 0 1 0-4 0z"/>`), g.Group(children),
	)
}

func ShoppingBagSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 4a4 4 0 1 1 8 0h4v12H0V4zm4-2a2 2 0 0 0-2 2h4a2 2 0 0 0-2-2M4 6v2h2V6zm6 0v2h2V6z"/>`), g.Group(children),
	)
}

func ShoppingCart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 2h2.117l.845 6.755l-.24.358C1.616 10.775 2.807 13 4.803 13H14v-2H4.803a.5.5 0 0 1-.416-.777L4.535 10H14.82l1.16-5.804a1 1 0 0 0-.738-1.166L3.123 0H0zm13.832 2.739L13.18 8H4.883l-.71-5.676zM5 15a1 1 0 1 1-2 0a1 1 0 0 1 2 0m8 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0"/>`), g.Group(children),
	)
}

func ShoppingCartSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 2h2.117l.845 6.755l-.24.358C1.616 10.775 2.807 13 4.803 13H14v-2H4.803a.5.5 0 0 1-.416-.777L4.535 10H14.82l1.16-5.804a1 1 0 0 0-.738-1.166L3.123 0H0zm4 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2m8 0a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/>`), g.Group(children),
	)
}

func Sidebyside(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h14v14H1zm2 2v10h4V3zm6 0v10h4V3z"/>`), g.Group(children),
	)
}

func SidebysideSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 15V1h14v14zM9 3H7v10h2z"/>`), g.Group(children),
	)
}

func Star(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m8 .684l2.231 4.66l5.629.266l-4.18 4.04l1.336 5.254L8 12.324l-5.016 2.58L4.32 9.65L.14 5.61l5.629-.265zm0 4.632l-.943 1.97l-2.197.104l1.684 1.629l-.528 2.077L8 10.076l1.984 1.02l-.528-2.077l1.684-1.63l-2.197-.103z"/>`), g.Group(children),
	)
}

func StarSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.231 5.345L8 .684l-2.231 4.66l-5.628.266l4.18 4.04l-1.337 5.254L8 12.324l5.016 2.58L11.68 9.65l4.18-4.04z"/>`), g.Group(children),
	)
}

func Sync(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 8v1h.1A5 5 0 0 0 5 13v-2a3 3 0 0 1-.5-5.959v1.827L8.803 4L4.5 1.131v1.894A5 5 0 0 0 0 8m16 0V7h-.1A5 5 0 0 0 11 3h-1v2h1a3 3 0 0 1 .5 5.959V9.13L7.197 12l4.303 2.868v-1.893A5 5 0 0 0 16 8"/>`), g.Group(children),
	)
}

func ThumbsDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.617 1H16v9h-4.323l-2 5H8a3 3 0 0 1-3-3v-1H2.633A2 2 0 0 1 .648 8.752L1.618 1ZM12 8h2V3h-2zm-2-5H3.383l-.75 6H7v3a1 1 0 0 0 1 1h.323L10 8.807z"/>`), g.Group(children),
	)
}

func ThumbsDownSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.5 10H16V1h-3.5zm-2-9H1.617L.648 8.752A2 2 0 0 0 2.633 11H5v1a3 3 0 0 0 3 3h.677l1.823-4.557z"/>`), g.Group(children),
	)
}

func ThumbsUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.323 1H8a3 3 0 0 1 3 3v1h2.367a2 2 0 0 1 1.985 2.248L14.382 15H0V6h4.323zM4 8H2v5h2zm2 5h6.617l.75-6H9V4a1 1 0 0 0-1-1h-.323L6 7.193z"/>`), g.Group(children),
	)
}

func ThumbsUpSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.323 1H8a3 3 0 0 1 3 3v1h2.367a2 2 0 0 1 1.985 2.248L14.382 15H5.5V5.557zM3.5 6H0v9h3.5z"/>`), g.Group(children),
	)
}

func Transfer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.586 3H1v2h10.586l-1.5 1.5L11.5 7.914L15.414 4L11.5.086L10.086 1.5zM4.414 13H15v-2H4.414l1.5-1.5L4.5 8.086L.586 12L4.5 15.914L5.914 14.5z"/>`), g.Group(children),
	)
}

func TransferAlt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 4.414V15h2V4.414l1.5 1.5L7.914 4.5L4 .586L.086 4.5L1.5 5.914zm10 7.172V1h-2v10.586l-1.5-1.5L8.086 11.5L12 15.414l3.914-3.914l-1.414-1.414z"/>`), g.Group(children),
	)
}

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M7 7v5h2V7z"/><path d="M11 0H5v3H0v2h1.165l1.553 8.537A3 3 0 0 0 5.669 16h4.662a3 3 0 0 0 2.952-2.463L14.835 5H16V3h-5zM9 3H7V2h2zM4.685 13.179L3.198 5h9.604l-1.487 8.179a1 1 0 0 1-.984.821H5.669a1 1 0 0 1-.984-.821"/></g>`), g.Group(children),
	)
}

func TrashSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 0h6v3h5v2h-1.165l-1.553 8.537A3 3 0 0 1 10.332 16H5.669a3 3 0 0 1-2.951-2.463L1.165 5H0V3h5zm2 3h2V2H7zm0 4v5h2V7z"/>`), g.Group(children),
	)
}

func Undo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 10.5v-1A4.5 4.5 0 0 0 11.5 5H3.914l2-2L4.5 1.586L.086 6L4.5 10.414L5.914 9l-2-2H11.5a2.5 2.5 0 0 1 0 5H4v2h7.5a4.5 4.5 0 0 0 4.389-3.5z"/>`), g.Group(children),
	)
}

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M7 4.174V11h2V4.174l2.608 2.236l1.302-1.518L8 .682l-4.91 4.21L4.392 6.41z"/><path d="M3 13v-3H1v3a2 2 0 0 0 2 2h10a2 2 0 0 0 2-2v-3h-2v3z"/></g>`), g.Group(children),
	)
}

func User(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 0a4 4 0 1 0 0 8a4 4 0 0 0 0-8M6 4a2 2 0 1 1 4 0a2 2 0 0 1-4 0m0 5a3 3 0 0 0-3 3v4h10v-4a3 3 0 0 0-3-3zm-1 3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v2H5z"/>`), g.Group(children),
	)
}

func UserAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 4a4 4 0 1 1 8 0a4 4 0 0 1-8 0m4-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4M0 12a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v4H0zm3-1a1 1 0 0 0-1 1v2h6v-2a1 1 0 0 0-1-1zm13-3h-2v2h-2V8h-2V6h2V4h2v2h2z"/>`), g.Group(children),
	)
}

func UserAddSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 4a4 4 0 1 1 8 0a4 4 0 0 1-8 0m-1 8a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v4H0zm16-4h-2v2h-2V8h-2V6h2V4h2v2h2z"/>`), g.Group(children),
	)
}

func UserGroup(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 0a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-1 3a1 1 0 1 1 2 0a1 1 0 0 1-2 0M5 3a3 3 0 1 0 0 6a3 3 0 0 0 0-6M4 6a1 1 0 1 1 2 0a1 1 0 0 1-2 0m-4 7a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v3H0zm3-1a1 1 0 0 0-1 1v1h6v-1a1 1 0 0 0-1-1zm6-5v2h4a1 1 0 0 1 1 1v1h-3v2h5v-3a3 3 0 0 0-3-3z"/>`), g.Group(children),
	)
}

func UserGroupSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M8 3a3 3 0 1 1 6 0a3 3 0 0 1-6 0M0 13a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v3H0z"/><path d="M8.873 7.003a4 4 0 0 1-1.278 2.041A4 4 0 0 1 11 13h5v-3a3 3 0 0 0-3-3H9q-.063 0-.127.003M5 3a3 3 0 1 0 0 6a3 3 0 0 0 0-6"/></g>`), g.Group(children),
	)
}

func UserRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 4a4 4 0 1 1 8 0a4 4 0 0 1-8 0m4-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4M0 12a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v4H0zm3-1a1 1 0 0 0-1 1v2h6v-2a1 1 0 0 0-1-1zm7-3h6V6h-6z"/>`), g.Group(children),
	)
}

func UserRemoveSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 4a4 4 0 1 1 8 0a4 4 0 0 1-8 0m-1 8a3 3 0 0 1 3-3h4a3 3 0 0 1 3 3v4H0zm10-4h6V6h-6z"/>`), g.Group(children),
	)
}

func UserSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 0a4 4 0 1 0 0 8a4 4 0 0 0 0-8M6 9a3 3 0 0 0-3 3v4h10v-4a3 3 0 0 0-3-3z"/>`), g.Group(children),
	)
}

func Utensils(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9.975 3.44l1.94-1.94L10.5.086l-1.94 1.94a3.5 3.5 0 0 0-.429 4.428L8 6.586l-5.5-5.5l-.94.94a3.5 3.5 0 0 0 0 4.949L4.587 10l-4 4L2 15.414l4-4l2 2l2-2l4 4L15.414 14l-6-6l.132-.131a3.5 3.5 0 0 0 4.429-.43l1.94-1.939L14.5 4.086l-1.94 1.94a1.5 1.5 0 0 1-1.534.362L13.914 3.5L12.5 2.086L9.612 4.974a1.5 1.5 0 0 1 .363-1.535ZM8.585 10L8 10.586L2.975 5.56a1.5 1.5 0 0 1-.363-1.535z"/>`), g.Group(children),
	)
}

func UtensilsSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9.975 3.44l1.94-1.94L10.5.086l-1.94 1.94a3.5 3.5 0 0 0-.429 4.428L8 6.586l-5.5-5.5l-.94.94a3.5 3.5 0 0 0 0 4.949L4.587 10l-4 4L2 15.414l4-4l2 2l2-2l4 4L15.414 14l-6-6l.132-.131a3.5 3.5 0 0 0 4.429-.43l1.94-1.939L14.5 4.086l-1.94 1.94a1.5 1.5 0 0 1-1.534.362L13.914 3.5L12.5 2.086L9.612 4.974a1.5 1.5 0 0 1 .363-1.535Z"/>`), g.Group(children),
	)
}

func Video(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 2h12v2.367l4-.5v8.266l-4-.5V14H0zm2 2v8h8V9.367l4 .5V6.133l-4 .5V4z"/>`), g.Group(children),
	)
}

func VideoOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m13.197 4.218l2.51-2.51L14.293.292l-14 14l1.414 1.414L3.414 14H12v-2.367l4 .5V3.867zM10.893 6.52L14 6.133v3.734l-4-.5V12H5.414l5.48-5.479ZM0 2h10.586l-2 2H2v6.586l-2 2z"/>`), g.Group(children),
	)
}

func VideoOffSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m14.293.293l1.414 1.414l-2.51 2.51L16 3.868v8.266l-4-.5V14H3.414l-1.707 1.707l-1.414-1.414zM10.086 2H0v10.086z"/>`), g.Group(children),
	)
}

func VideoSolid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2H0v12h12v-2.367l4 .5V3.867l-4 .5z"/>`), g.Group(children),
	)
}
