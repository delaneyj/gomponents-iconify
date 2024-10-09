package iconamoon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const IconifyVersion = ""

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 24 24")
)

func Apps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h2v2H4zm0 14h2v2H4zM18 4h2v2h-2zm0 7h2v2h-2zm-7 0h2v2h-2zm-7 0h2v2H4zm7-7h2v2h-2zm0 14h2v2h-2zm7 0h2v2h-2z"/>`), g.Group(children),
	)
}

func AppsBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h2v2H4zm0 14h2v2H4zM18 4h2v2h-2zm0 7h2v2h-2zm-7 0h2v2h-2zm-7 0h2v2H4zm7-7h2v2h-2zm0 14h2v2h-2zm7 0h2v2h-2z"/>`), g.Group(children),
	)
}

func AppsFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1zm0 14a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1zM18 3a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zm-1 8a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1zm-6-1a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1zm-8 1a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1zm8-8a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zm-1 15a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-2a1 1 0 0 1-1-1zm8-1a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func AppsLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h2v2H4zm0 14h2v2H4zM18 4h2v2h-2zm0 7h2v2h-2zm-7 0h2v2h-2zm-7 0h2v2H4zm7-7h2v2h-2zm0 14h2v2h-2zm7 0h2v2h-2z"/>`), g.Group(children),
	)
}

func AppsThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h2v2H4zm0 14h2v2H4zM18 4h2v2h-2zm0 7h2v2h-2zm-7 0h2v2h-2zm-7 0h2v2H4zm7-7h2v2h-2zm0 14h2v2h-2zm7 0h2v2h-2z"/>`), g.Group(children),
	)
}

func ArrowBottomLeftFiveCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M10 10v4m0 0h4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftFiveCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M10 10v4m0 0h4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftFiveCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 10v4m0 0h4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftFiveCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m8-3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h4a1 1 0 1 0 0-2h-1.586l2.293-2.293a1 1 0 0 0-1.414-1.414L11 11.586V10a1 1 0 0 0-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomLeftFiveCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M10 10v4m0 0h4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftFiveCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M10 10v4m0 0h4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftFourSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 10v4h4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftFourSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 10v4h4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftFourSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 10v4h4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftFourSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm6 10a1 1 0 0 0 1 1h4a1 1 0 1 0 0-2h-3v-3a1 1 0 1 0-2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomLeftFourSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 10v4h4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftFourSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 10v4h4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 17H7m0 0V7m0 10L17 7"/>`), g.Group(children),
	)
}

func ArrowBottomLeftOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M17 17H7m0 0V7m0 10L17 7"/>`), g.Group(children),
	)
}

func ArrowBottomLeftOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 7v10a1 1 0 0 0 1 1h10a1 1 0 0 0 .707-1.707L13.414 12l4.293-4.293a1 1 0 0 0-1.414-1.414L12 10.586L7.707 6.293A1 1 0 0 0 6 7" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomLeftOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 17H7m0 0V7m0 10L17 7"/>`), g.Group(children),
	)
}

func ArrowBottomLeftOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 17H7m0 0V7m0 10L17 7"/>`), g.Group(children),
	)
}

func ArrowBottomLeftSixCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M10 10v4h4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftSixCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M10 10v4h4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftSixCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 10v4h4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftSixCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7 2a1 1 0 0 0 1 1h4a1 1 0 1 0 0-2h-3v-3a1 1 0 1 0-2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomLeftSixCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M10 10v4h4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftSixCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M10 10v4h4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftThreeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 6v4m0 0h4m-4 0l4-4"/>`), g.Group(children),
	)
}

func ArrowBottomLeftThreeSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 6v4m0 0h4m-4 0l4-4"/>`), g.Group(children),
	)
}

func ArrowBottomLeftThreeSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 6v4m0 0h4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomLeftThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm6 6v4a1 1 0 0 0 1 1h4a1 1 0 1 0 0-2h-1.586l2.293-2.293a1 1 0 0 0-1.414-1.414L11 11.586V10a1 1 0 1 0-2 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomLeftThreeSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 6v4m0 0h4m-4 0l4-4"/>`), g.Group(children),
	)
}

func ArrowBottomLeftThreeSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 6v4m0 0h4m-4 0l4-4"/>`), g.Group(children),
	)
}

func ArrowBottomLeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8v8h8"/>`), g.Group(children),
	)
}

func ArrowBottomLeftTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M8 8v8h8"/>`), g.Group(children),
	)
}

func ArrowBottomLeftTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 8v8a1 1 0 0 0 1 1h8a1 1 0 0 0 .707-1.707l-8-8A1 1 0 0 0 7 8" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomLeftTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 8v8h8"/>`), g.Group(children),
	)
}

func ArrowBottomLeftTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 8v8h8"/>`), g.Group(children),
	)
}

func ArrowBottomRightFiveCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M10 14h4m0 0v-4m0 4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightFiveCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M10 14h4m0 0v-4m0 4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightFiveCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14h4m0 0v-4m0 4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightFiveCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7 2a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1v-4a1 1 0 1 0-2 0v1.586l-2.293-2.293a1 1 0 0 0-1.414 1.414L11.586 13H10a1 1 0 0 0-1 1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomRightFiveCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M10 14h4m0 0v-4m0 4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightFiveCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M10 14h4m0 0v-4m0 4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightFourSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 14h4v-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightFourSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 14h4v-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightFourSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14h4v-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightFourSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm11 11a1 1 0 0 0 1-1v-4a1 1 0 1 0-2 0v3h-3a1 1 0 1 0 0 2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomRightFourSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 14h4v-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightFourSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 14h4v-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 7v10m0 0H7m10 0L7 7"/>`), g.Group(children),
	)
}

func ArrowBottomRightOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M17 7v10m0 0H7m10 0L7 7"/>`), g.Group(children),
	)
}

func ArrowBottomRightOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 18h10a1 1 0 0 0 1-1V7a1 1 0 0 0-1.707-.707L12 10.586L7.707 6.293a1 1 0 0 0-1.414 1.414L10.586 12l-4.293 4.293A1 1 0 0 0 7 18" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomRightOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 7v10m0 0H7m10 0L7 7"/>`), g.Group(children),
	)
}

func ArrowBottomRightOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 7v10m0 0H7m10 0L7 7"/>`), g.Group(children),
	)
}

func ArrowBottomRightSixCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M10 14h4v-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightSixCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M10 14h4v-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightSixCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 14h4v-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightSixCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m12 3a1 1 0 0 0 1-1v-4a1 1 0 1 0-2 0v3h-3a1 1 0 1 0 0 2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomRightSixCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M10 14h4v-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightSixCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M10 14h4v-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightThreeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 10h4m0 0v-4m0 4l-4-4"/>`), g.Group(children),
	)
}

func ArrowBottomRightThreeSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 10h4m0 0v-4m0 4l-4-4"/>`), g.Group(children),
	)
}

func ArrowBottomRightThreeSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 10h4m0 0v-4m0 4l-4-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomRightThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm7 11h4a1 1 0 0 0 1-1v-4a1 1 0 1 0-2 0v1.586l-2.293-2.293a1 1 0 0 0-1.414 1.414L11.586 13H10a1 1 0 1 0 0 2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomRightThreeSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 10h4m0 0v-4m0 4l-4-4"/>`), g.Group(children),
	)
}

func ArrowBottomRightThreeSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 10h4m0 0v-4m0 4l-4-4"/>`), g.Group(children),
	)
}

func ArrowBottomUpFourSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M14 14v-4h-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomUpFourSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M14 14v-4h-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomUpFourSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 14v-4h-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomUpFourSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm12 6a1 1 0 0 0-1-1h-4a1 1 0 1 0 0 2h3v3a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowBottomUpFourSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M14 14v-4h-4"/></g>`), g.Group(children),
	)
}

func ArrowBottomUpFourSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M14 14v-4h-4"/></g>`), g.Group(children),
	)
}

func ArrowDownFiveCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m9 12l3 3m0 0l3-3m-3 3V9"/></g>`), g.Group(children),
	)
}

func ArrowDownFiveCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m9 12l3 3m0 0l3-3m-3 3V9"/></g>`), g.Group(children),
	)
}

func ArrowDownFiveCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 12l3 3m0 0l3-3m-3 3V9"/></g>`), g.Group(children),
	)
}

func ArrowDownFiveCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m11-3a1 1 0 1 0-2 0v3.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0-1.414-1.414L13 12.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowDownFiveCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m9 12l3 3m0 0l3-3m-3 3V9"/></g>`), g.Group(children),
	)
}

func ArrowDownFiveCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m9 12l3 3m0 0l3-3m-3 3V9"/></g>`), g.Group(children),
	)
}

func ArrowDownFourSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m9 11l3 3l3-3"/></g>`), g.Group(children),
	)
}

func ArrowDownFourSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m9 11l3 3l3-3"/></g>`), g.Group(children),
	)
}

func ArrowDownFourSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 11l3 3l3-3"/></g>`), g.Group(children),
	)
}

func ArrowDownFourSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm6.707 6.293a1 1 0 1 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0-1.414-1.414L12 12.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowDownFourSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m9 11l3 3l3-3"/></g>`), g.Group(children),
	)
}

func ArrowDownFourSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m9 11l3 3l3-3"/></g>`), g.Group(children),
	)
}

func ArrowDownOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 19V5m-7 7l7 7l7-7"/>`), g.Group(children),
	)
}

func ArrowDownOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 19V5m-7 7l7 7l7-7"/>`), g.Group(children),
	)
}

func ArrowDownOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 12l7 7m0 0l7-7m-7 7V5"/>`), g.Group(children),
	)
}

func ArrowDownOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 11a1 1 0 0 0-.707 1.707l7 7a1 1 0 0 0 1.414 0l7-7A1 1 0 0 0 19 11h-6V5a1 1 0 1 0-2 0v6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowDownOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 12l7 7m0 0l7-7m-7 7V5"/>`), g.Group(children),
	)
}

func ArrowDownOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5 12l7 7m0 0l7-7m-7 7V5"/>`), g.Group(children),
	)
}

func ArrowDownSixCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m9 11l3 3l3-3"/></g>`), g.Group(children),
	)
}

func ArrowDownSixCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m9 11l3 3l3-3"/></g>`), g.Group(children),
	)
}

func ArrowDownSixCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m9 11l3 3l3-3"/></g>`), g.Group(children),
	)
}

func ArrowDownSixCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7.707-1.707a1 1 0 1 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0-1.414-1.414L12 12.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowDownSixCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m9 11l3 3l3-3"/></g>`), g.Group(children),
	)
}

func ArrowDownSixCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m9 11l3 3l3-3"/></g>`), g.Group(children),
	)
}

func ArrowDownThreeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm5 8l3 3m0 0l3-3m-3 3V9"/>`), g.Group(children),
	)
}

func ArrowDownThreeSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm5 8l3 3m0 0l3-3m-3 3V9"/>`), g.Group(children),
	)
}

func ArrowDownThreeSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm5 8l3 3m0 0l3-3m-3 3V9"/></g>`), g.Group(children),
	)
}

func ArrowDownThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm10 5a1 1 0 1 0-2 0v3.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0-1.414-1.414L13 12.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowDownThreeSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm5 8l3 3m0 0l3-3m-3 3V9"/>`), g.Group(children),
	)
}

func ArrowDownThreeSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm5 8l3 3m0 0l3-3m-3 3V9"/>`), g.Group(children),
	)
}

func ArrowDownTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 10l5 5m0 0l5-5"/>`), g.Group(children),
	)
}

func ArrowDownTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m7 10l5 5m0 0l5-5"/>`), g.Group(children),
	)
}

func ArrowDownTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 10l5 5l5-5"/>`), g.Group(children),
	)
}

func ArrowDownTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 9a1 1 0 0 0-.707 1.707l5 5a1 1 0 0 0 1.414 0l5-5A1 1 0 0 0 17 9z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowDownTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7 10l5 5l5-5"/>`), g.Group(children),
	)
}

func ArrowDownTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7 10l5 5l5-5"/>`), g.Group(children),
	)
}

func ArrowLeftFiveCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m12 9l-3 3m0 0l3 3m-3-3h6"/></g>`), g.Group(children),
	)
}

func ArrowLeftFiveCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m12 9l-3 3m0 0l3 3m-3-3h6"/></g>`), g.Group(children),
	)
}

func ArrowLeftFiveCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 9l-3 3m0 0l3 3m-3-3h6"/></g>`), g.Group(children),
	)
}

func ArrowLeftFiveCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m10.707-2.293a1 1 0 0 0-1.414-1.414l-3 3a1 1 0 0 0 0 1.414l3 3a1 1 0 0 0 1.414-1.414L11.414 13H15a1 1 0 1 0 0-2h-3.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowLeftFiveCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m12 9l-3 3m0 0l3 3m-3-3h6"/></g>`), g.Group(children),
	)
}

func ArrowLeftFiveCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m12 9l-3 3m0 0l3 3m-3-3h6"/></g>`), g.Group(children),
	)
}

func ArrowLeftFourSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m13 9l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ArrowLeftFourSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m13 9l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ArrowLeftFourSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 9l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ArrowLeftFourSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm10.707 5.707a1 1 0 0 0-1.414-1.414l-3 3a1 1 0 0 0 0 1.414l3 3a1 1 0 0 0 1.414-1.414L11.414 12z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowLeftFourSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m13 9l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ArrowLeftFourSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m13 9l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ArrowLeftOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 12h14m-7-7l-7 7l7 7"/>`), g.Group(children),
	)
}

func ArrowLeftOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 12h14m-7-7l-7 7l7 7"/>`), g.Group(children),
	)
}

func ArrowLeftOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 5l-7 7m0 0l7 7m-7-7h14"/>`), g.Group(children),
	)
}

func ArrowLeftOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 5a1 1 0 0 0-1.707-.707l-7 7a1 1 0 0 0 0 1.414l7 7A1 1 0 0 0 13 19v-6h6a1 1 0 1 0 0-2h-6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowLeftOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 5l-7 7m0 0l7 7m-7-7h14"/>`), g.Group(children),
	)
}

func ArrowLeftOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12 5l-7 7m0 0l7 7m-7-7h14"/>`), g.Group(children),
	)
}

func ArrowLeftSixCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m13 9l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ArrowLeftSixCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m13 9l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ArrowLeftSixCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 9l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ArrowLeftSixCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m11.707-2.293a1 1 0 0 0-1.414-1.414l-3 3a1 1 0 0 0 0 1.414l3 3a1 1 0 0 0 1.414-1.414L11.414 12z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowLeftSixCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m13 9l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ArrowLeftSixCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m13 9l-3 3l3 3"/></g>`), g.Group(children),
	)
}

func ArrowLeftThreeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 5l-3 3m0 0l3 3m-3-3h6"/>`), g.Group(children),
	)
}

func ArrowLeftThreeSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 5l-3 3m0 0l3 3m-3-3h6"/>`), g.Group(children),
	)
}

func ArrowLeftThreeSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 5l-3 3m0 0l3 3m-3-3h6"/></g>`), g.Group(children),
	)
}

func ArrowLeftThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm9.707 5.707a1 1 0 0 0-1.414-1.414l-3 3a1 1 0 0 0 0 1.414l3 3a1 1 0 0 0 1.414-1.414L11.414 13H15a1 1 0 1 0 0-2h-3.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowLeftThreeSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 5l-3 3m0 0l3 3m-3-3h6"/>`), g.Group(children),
	)
}

func ArrowLeftThreeSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 5l-3 3m0 0l3 3m-3-3h6"/>`), g.Group(children),
	)
}

func ArrowLeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 7l-5 5m0 0l5 5"/>`), g.Group(children),
	)
}

func ArrowLeftTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m14 7l-5 5m0 0l5 5"/>`), g.Group(children),
	)
}

func ArrowLeftTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 7l-5 5l5 5"/>`), g.Group(children),
	)
}

func ArrowLeftTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 7a1 1 0 0 0-1.707-.707l-5 5a1 1 0 0 0 0 1.414l5 5A1 1 0 0 0 15 17z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowLeftTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14 7l-5 5l5 5"/>`), g.Group(children),
	)
}

func ArrowLeftTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14 7l-5 5l5 5"/>`), g.Group(children),
	)
}

func ArrowRightFiveCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m12 15l3-3m0 0l-3-3m3 3H9"/></g>`), g.Group(children),
	)
}

func ArrowRightFiveCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m12 15l3-3m0 0l-3-3m3 3H9"/></g>`), g.Group(children),
	)
}

func ArrowRightFiveCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 15l3-3m0 0l-3-3m3 3H9"/></g>`), g.Group(children),
	)
}

func ArrowRightFiveCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m10.707-3.707a1 1 0 1 0-1.414 1.414L12.586 11H9a1 1 0 1 0 0 2h3.586l-1.293 1.293a1 1 0 0 0 1.414 1.414l3-3a1 1 0 0 0 0-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowRightFiveCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m12 15l3-3m0 0l-3-3m3 3H9"/></g>`), g.Group(children),
	)
}

func ArrowRightFiveCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m12 15l3-3m0 0l-3-3m3 3H9"/></g>`), g.Group(children),
	)
}

func ArrowRightFourSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m11 15l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func ArrowRightFourSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m11 15l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func ArrowRightFourSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 15l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func ArrowRightFourSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm7.293 10.293a1 1 0 1 0 1.414 1.414l3-3a1 1 0 0 0 0-1.414l-3-3a1 1 0 1 0-1.414 1.414L12.586 12z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowRightFourSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m11 15l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func ArrowRightFourSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m11 15l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func ArrowRightOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 12H5m7 7l7-7l-7-7"/>`), g.Group(children),
	)
}

func ArrowRightOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M19 12H5m7 7l7-7l-7-7"/>`), g.Group(children),
	)
}

func ArrowRightOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 19l7-7l-7-7m7 7H5"/>`), g.Group(children),
	)
}

func ArrowRightOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 19a1 1 0 0 0 1.707.707l7-7a1 1 0 0 0 0-1.414l-7-7A1 1 0 0 0 11 5v6H5a1 1 0 1 0 0 2h6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowRightOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m12 19l7-7l-7-7m7 7H5"/>`), g.Group(children),
	)
}

func ArrowRightOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12 19l7-7l-7-7m7 7H5"/>`), g.Group(children),
	)
}

func ArrowRightSixCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m11 15l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func ArrowRightSixCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m11 15l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func ArrowRightSixCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m11 15l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func ArrowRightSixCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m8.293 2.293a1 1 0 1 0 1.414 1.414l3-3a1 1 0 0 0 0-1.414l-3-3a1 1 0 1 0-1.414 1.414L12.586 12z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowRightSixCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m11 15l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func ArrowRightSixCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m11 15l3-3l-3-3"/></g>`), g.Group(children),
	)
}

func ArrowRightThreeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 11l3-3m0 0l-3-3m3 3H9"/>`), g.Group(children),
	)
}

func ArrowRightThreeSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 11l3-3m0 0l-3-3m3 3H9"/>`), g.Group(children),
	)
}

func ArrowRightThreeSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 11l3-3m0 0l-3-3m3 3H9"/></g>`), g.Group(children),
	)
}

func ArrowRightThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm9.707 4.293a1 1 0 1 0-1.414 1.414L12.586 11H9a1 1 0 1 0 0 2h3.586l-1.293 1.293a1 1 0 0 0 1.414 1.414l3-3a1 1 0 0 0 0-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowRightThreeSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 11l3-3m0 0l-3-3m3 3H9"/>`), g.Group(children),
	)
}

func ArrowRightThreeSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 11l3-3m0 0l-3-3m3 3H9"/>`), g.Group(children),
	)
}

func ArrowRightTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 17l5-5m0 0l-5-5"/>`), g.Group(children),
	)
}

func ArrowRightTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m10 17l5-5m0 0l-5-5"/>`), g.Group(children),
	)
}

func ArrowRightTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 17l5-5l-5-5"/>`), g.Group(children),
	)
}

func ArrowRightTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 17a1 1 0 0 0 1.707.707l5-5a1 1 0 0 0 0-1.414l-5-5A1 1 0 0 0 9 7z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowRightTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10 17l5-5l-5-5"/>`), g.Group(children),
	)
}

func ArrowRightTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10 17l5-5l-5-5"/>`), g.Group(children),
	)
}

func ArrowTopLeftFiveCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M14 10h-4m0 0v4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftFiveCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M14 10h-4m0 0v4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftFiveCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10h-4m0 0v4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftFiveCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m13-2a1 1 0 0 0-1-1h-4a1 1 0 0 0-1 1v4a1 1 0 1 0 2 0v-1.586l2.293 2.293a1 1 0 0 0 1.414-1.414L12.414 11H14a1 1 0 0 0 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowTopLeftFiveCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14 10h-4m0 0v4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftFiveCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M14 10h-4m0 0v4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftFourSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M14 10h-4v4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftFourSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M14 10h-4v4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftFourSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10h-4v4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftFourSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm7 5a1 1 0 0 0-1 1v4a1 1 0 1 0 2 0v-3h3a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowTopLeftFourSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M14 10h-4v4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftFourSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M14 10h-4v4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 7H7m0 0v10M7 7l10 10"/>`), g.Group(children),
	)
}

func ArrowTopLeftOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M17 7H7m0 0v10M7 7l10 10"/>`), g.Group(children),
	)
}

func ArrowTopLeftOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 6H7a1 1 0 0 0-1 1v10a1 1 0 0 0 1.707.707L12 13.414l4.293 4.293a1 1 0 0 0 1.414-1.414L13.414 12l4.293-4.293A1 1 0 0 0 17 6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowTopLeftOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 7H7m0 0v10M7 7l10 10"/>`), g.Group(children),
	)
}

func ArrowTopLeftOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 7H7m0 0v10M7 7l10 10"/>`), g.Group(children),
	)
}

func ArrowTopLeftSixCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M14 10h-4v4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftSixCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M14 10h-4v4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftSixCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10h-4v4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftSixCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m8-3a1 1 0 0 0-1 1v4a1 1 0 1 0 2 0v-3h3a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowTopLeftSixCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14 10h-4v4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftSixCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M14 10h-4v4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftThreeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 6h-4m0 0v4m0-4l4 4"/>`), g.Group(children),
	)
}

func ArrowTopLeftThreeSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 6h-4m0 0v4m0-4l4 4"/>`), g.Group(children),
	)
}

func ArrowTopLeftThreeSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 6h-4m0 0v4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopLeftThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm11 5h-4a1 1 0 0 0-1 1v4a1 1 0 1 0 2 0v-1.586l2.293 2.293a1 1 0 0 0 1.414-1.414L12.414 11H14a1 1 0 1 0 0-2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowTopLeftThreeSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 6h-4m0 0v4m0-4l4 4"/>`), g.Group(children),
	)
}

func ArrowTopLeftThreeSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 6h-4m0 0v4m0-4l4 4"/>`), g.Group(children),
	)
}

func ArrowTopLeftTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 8H8v8"/>`), g.Group(children),
	)
}

func ArrowTopLeftTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M16 8H8v8"/>`), g.Group(children),
	)
}

func ArrowTopLeftTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 7H8a1 1 0 0 0-1 1v8a1 1 0 0 0 1.707.707l8-8A1 1 0 0 0 16 7" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowTopLeftTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 8H8v8"/>`), g.Group(children),
	)
}

func ArrowTopLeftTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 8H8v8"/>`), g.Group(children),
	)
}

func ArrowTopRightFiveCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M14 14v-4m0 0h-4m4 0l-4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopRightFiveCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M14 14v-4m0 0h-4m4 0l-4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopRightFiveCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 14v-4m0 0h-4m4 0l-4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopRightFiveCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m12 3a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1h-4a1 1 0 1 0 0 2h1.586l-2.293 2.293a1 1 0 1 0 1.414 1.414L13 12.414V14a1 1 0 0 0 1 1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowTopRightFiveCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14 14v-4m0 0h-4m4 0l-4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopRightFiveCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M14 14v-4m0 0h-4m4 0l-4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopRightOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 7h10m0 0v10m0-10L7 17"/>`), g.Group(children),
	)
}

func ArrowTopRightOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M7 7h10m0 0v10m0-10L7 17"/>`), g.Group(children),
	)
}

func ArrowTopRightOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 17V7a1 1 0 0 0-1-1H7a1 1 0 0 0-.707 1.707L10.586 12l-4.293 4.293a1 1 0 1 0 1.414 1.414L12 13.414l4.293 4.293A1 1 0 0 0 18 17" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowTopRightOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 7h10m0 0v10m0-10L7 17"/>`), g.Group(children),
	)
}

func ArrowTopRightOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 7h10m0 0v10m0-10L7 17"/>`), g.Group(children),
	)
}

func ArrowTopRightSixCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M14 14v-4h-4"/></g>`), g.Group(children),
	)
}

func ArrowTopRightSixCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M14 14v-4h-4"/></g>`), g.Group(children),
	)
}

func ArrowTopRightSixCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 14v-4h-4"/></g>`), g.Group(children),
	)
}

func ArrowTopRightSixCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m13-2a1 1 0 0 0-1-1h-4a1 1 0 1 0 0 2h3v3a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowTopRightSixCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14 14v-4h-4"/></g>`), g.Group(children),
	)
}

func ArrowTopRightSixCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M14 14v-4h-4"/></g>`), g.Group(children),
	)
}

func ArrowTopRightThreeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 10v-4m0 0h-4m4 0l-4 4"/>`), g.Group(children),
	)
}

func ArrowTopRightThreeSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 10v-4m0 0h-4m4 0l-4 4"/>`), g.Group(children),
	)
}

func ArrowTopRightThreeSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 10v-4m0 0h-4m4 0l-4 4"/></g>`), g.Group(children),
	)
}

func ArrowTopRightThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm12 10v-4a1 1 0 0 0-1-1h-4a1 1 0 1 0 0 2h1.586l-2.293 2.293a1 1 0 1 0 1.414 1.414L13 12.414V14a1 1 0 1 0 2 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowTopRightThreeSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 10v-4m0 0h-4m4 0l-4 4"/>`), g.Group(children),
	)
}

func ArrowTopRightThreeSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 10v-4m0 0h-4m4 0l-4 4"/>`), g.Group(children),
	)
}

func ArrowTopRightTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 16V8H8"/>`), g.Group(children),
	)
}

func ArrowTopRightTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M16 16V8H8"/>`), g.Group(children),
	)
}

func ArrowTopRightTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 16V8a1 1 0 0 0-1-1H8a1 1 0 0 0-.707 1.707l8 8A1 1 0 0 0 17 16" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowTopRightTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 16V8H8"/>`), g.Group(children),
	)
}

func ArrowTopRightTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 16V8H8"/>`), g.Group(children),
	)
}

func ArrowUpFiveCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m15 12l-3-3m0 0l-3 3m3-3v6"/></g>`), g.Group(children),
	)
}

func ArrowUpFiveCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m15 12l-3-3m0 0l-3 3m3-3v6"/></g>`), g.Group(children),
	)
}

func ArrowUpFiveCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 12l-3-3m0 0l-3 3m3-3v6"/></g>`), g.Group(children),
	)
}

func ArrowUpFiveCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m13.707-.707l-3-3a1 1 0 0 0-1.414 0l-3 3a1 1 0 1 0 1.414 1.414L11 11.414V15a1 1 0 1 0 2 0v-3.586l1.293 1.293a1 1 0 0 0 1.414-1.414" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowUpFiveCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m15 12l-3-3m0 0l-3 3m3-3v6"/></g>`), g.Group(children),
	)
}

func ArrowUpFiveCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m15 12l-3-3m0 0l-3 3m3-3v6"/></g>`), g.Group(children),
	)
}

func ArrowUpFourSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m15 13l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func ArrowUpFourSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m15 13l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func ArrowUpFourSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 13l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func ArrowUpFourSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm11.293 9.707a1 1 0 0 0 1.414-1.414l-3-3a1 1 0 0 0-1.414 0l-3 3a1 1 0 1 0 1.414 1.414L12 11.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowUpFourSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m15 13l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func ArrowUpFourSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m15 13l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func ArrowUpOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 5v14m7-7l-7-7l-7 7"/>`), g.Group(children),
	)
}

func ArrowUpOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 5v14m7-7l-7-7l-7 7"/>`), g.Group(children),
	)
}

func ArrowUpOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m19 12l-7-7l-7 7m7-7v14"/>`), g.Group(children),
	)
}

func ArrowUpOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 13a1 1 0 0 0 .707-1.707l-7-7a1 1 0 0 0-1.414 0l-7 7A1 1 0 0 0 5 13h6v6a1 1 0 1 0 2 0v-6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowUpOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19 12l-7-7l-7 7m7-7v14"/>`), g.Group(children),
	)
}

func ArrowUpOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19 12l-7-7l-7 7m7-7v14"/>`), g.Group(children),
	)
}

func ArrowUpSixCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m15 13l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func ArrowUpSixCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m15 13l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func ArrowUpSixCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 13l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func ArrowUpSixCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m12.293 1.707a1 1 0 0 0 1.414-1.414l-3-3a1 1 0 0 0-1.414 0l-3 3a1 1 0 1 0 1.414 1.414L12 11.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowUpSixCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m15 13l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func ArrowUpSixCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m15 13l-3-3l-3 3"/></g>`), g.Group(children),
	)
}

func ArrowUpThreeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm11 8l-3-3m0 0l-3 3m3-3v6"/>`), g.Group(children),
	)
}

func ArrowUpThreeSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm11 8l-3-3m0 0l-3 3m3-3v6"/>`), g.Group(children),
	)
}

func ArrowUpThreeSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm11 8l-3-3m0 0l-3 3m3-3v6"/></g>`), g.Group(children),
	)
}

func ArrowUpThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm9.707 4.293a1 1 0 0 0-1.414 0l-3 3a1 1 0 1 0 1.414 1.414L11 11.414V15a1 1 0 1 0 2 0v-3.586l1.293 1.293a1 1 0 0 0 1.414-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowUpThreeSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm11 8l-3-3m0 0l-3 3m3-3v6"/>`), g.Group(children),
	)
}

func ArrowUpThreeSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm11 8l-3-3m0 0l-3 3m3-3v6"/>`), g.Group(children),
	)
}

func ArrowUpTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 14l-5-5m0 0l-5 5"/>`), g.Group(children),
	)
}

func ArrowUpTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m17 14l-5-5m0 0l-5 5"/>`), g.Group(children),
	)
}

func ArrowUpTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 14l-5-5l-5 5"/>`), g.Group(children),
	)
}

func ArrowUpTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 15a1 1 0 0 0 .707-1.707l-5-5a1 1 0 0 0-1.414 0l-5 5A1 1 0 0 0 7 15z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ArrowUpTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 14l-5-5l-5 5"/>`), g.Group(children),
	)
}

func ArrowUpTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17 14l-5-5l-5 5"/>`), g.Group(children),
	)
}

func Attachment(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 15V9a6 6 0 1 1 12 0v8a4 4 0 0 1-8 0V9a2 2 0 1 1 4 0v8"/>`), g.Group(children),
	)
}

func AttachmentBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M6 15V9a6 6 0 1 1 12 0v8a4 4 0 0 1-8 0V9a2 2 0 1 1 4 0v8"/>`), g.Group(children),
	)
}

func AttachmentDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 15V9a6 6 0 1 1 12 0v8a4 4 0 1 1-8 0V9a2 2 0 1 1 4 0v8"/>`), g.Group(children),
	)
}

func AttachmentFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 4a5 5 0 0 0-5 5v6a1 1 0 1 1-2 0V9a7 7 0 1 1 14 0v8a5 5 0 0 1-10 0V9a3 3 0 0 1 6 0v8a1 1 0 1 1-2 0V9a1 1 0 1 0-2 0v8a3 3 0 0 0 6 0V9a5 5 0 0 0-5-5" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func AttachmentLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 15V9a6 6 0 1 1 12 0v8a4 4 0 1 1-8 0V9a2 2 0 1 1 4 0v8"/>`), g.Group(children),
	)
}

func AttachmentThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 15V9a6 6 0 1 1 12 0v8a4 4 0 1 1-8 0V9a2 2 0 1 1 4 0v8"/>`), g.Group(children),
	)
}

func AttentionCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M12 12V8"/></g>`), g.Group(children),
	)
}

func AttentionCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M12 12V8"/></g>`), g.Group(children),
	)
}

func AttentionCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 16h.01v.01H12z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12V8"/></g>`), g.Group(children),
	)
}

func AttentionCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m10 1a1 1 0 0 1-1-1V8a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m-1.5 3a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func AttentionCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M12 12V8"/></g>`), g.Group(children),
	)
}

func AttentionCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" d="M12 12V8"/></g>`), g.Group(children),
	)
}

func AttentionSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M12 12V8"/></g>`), g.Group(children),
	)
}

func AttentionSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3.75" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M12 12V8"/></g>`), g.Group(children),
	)
}

func AttentionSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 16h.01v.01H12z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12V8"/></g>`), g.Group(children),
	)
}

func AttentionSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zM12 13a1 1 0 0 1-1-1V8a1 1 0 1 1 2 0v4a1 1 0 0 1-1 1m-1.5 3a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func AttentionSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="1.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="2.25" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M12 12V8"/></g>`), g.Group(children),
	)
}

func AttentionSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="1.5" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" d="M12 12V8"/></g>`), g.Group(children),
	)
}

func Backspace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.92 5a2 2 0 0 0-1.519.698l-4.285 5a2 2 0 0 0 0 2.604l4.285 5A2 2 0 0 0 7.92 19H19a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2zM15 10l-4 4m0-4l4 4"/>`), g.Group(children),
	)
}

func BackspaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M7.92 5a2 2 0 0 0-1.519.698l-4.285 5a2 2 0 0 0 0 2.604l4.285 5A2 2 0 0 0 7.92 19H19a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2zM15 10l-4 4m0-4l4 4"/>`), g.Group(children),
	)
}

func BackspaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7.92 5a2 2 0 0 0-1.519.698l-4.285 5a2 2 0 0 0 0 2.604l4.285 5A2 2 0 0 0 7.92 19H19a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.92 5a2 2 0 0 0-1.519.698l-4.285 5a2 2 0 0 0 0 2.604l4.285 5A2 2 0 0 0 7.92 19H19a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2zM15 10l-4 4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func BackspaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.92 4H19a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H7.92a3 3 0 0 1-2.278-1.048l-4.285-5a3 3 0 0 1 0-3.904l4.285-5A3 3 0 0 1 7.92 4m2.373 5.293a1 1 0 0 1 1.414 0L13 10.586l1.293-1.293a1 1 0 1 1 1.414 1.414L14.414 12l1.293 1.293a1 1 0 0 1-1.414 1.414L13 13.414l-1.293 1.293a1 1 0 1 1-1.414-1.414L11.586 12l-1.293-1.293a1 1 0 0 1 0-1.414" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BackspaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.92 5a2 2 0 0 0-1.519.698l-4.285 5a2 2 0 0 0 0 2.604l4.285 5A2 2 0 0 0 7.92 19H19a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2zM15 10l-4 4m0-4l4 4"/>`), g.Group(children),
	)
}

func BackspaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.92 5a2 2 0 0 0-1.519.698l-4.285 5a2 2 0 0 0 0 2.604l4.285 5A2 2 0 0 0 7.92 19H19a2 2 0 0 0 2-2V7a2 2 0 0 0-2-2zM15 10l-4 4m0-4l4 4"/>`), g.Group(children),
	)
}

func Badge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2" d="M12.586 4.586A2 2 0 0 0 11.172 4H4v7.172a2 2 0 0 0 .586 1.414l7 7a2 2 0 0 0 2.828 0l5.172-5.172a2 2 0 0 0 0-2.828z"/><path stroke-width="3" d="M9 9h.01v.01H9z"/></g>`), g.Group(children),
	)
}

func BadgeBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M12.586 4.586A2 2 0 0 0 11.172 4H4v7.172a2 2 0 0 0 .586 1.414l7 7a2 2 0 0 0 2.828 0l5.172-5.172a2 2 0 0 0 0-2.828z"/><path stroke-width="3.75" d="M9 9h.01v.01H9z"/></g>`), g.Group(children),
	)
}

func BadgeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12.586 4.586A2 2 0 0 0 11.172 4H4v7.172a2 2 0 0 0 .586 1.414l7 7a2 2 0 0 0 2.828 0l5.172-5.172a2 2 0 0 0 0-2.828z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.586 4.586A2 2 0 0 0 11.172 4H4v7.172a2 2 0 0 0 .586 1.414l7 7a2 2 0 0 0 2.828 0l5.172-5.172a2 2 0 0 0 0-2.828z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9 9h.01v.01H9z"/></g>`), g.Group(children),
	)
}

func BadgeFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h7.172a3 3 0 0 1 2.12.879l7 7a3 3 0 0 1 0 4.242l-5.17 5.172a3 3 0 0 1-4.243 0l-7-7A3 3 0 0 1 3 11.172zm6 3.5A1.5 1.5 0 0 0 7.5 9v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9a1.5 1.5 0 0 0-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BadgeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="1.5" d="M12.586 4.586A2 2 0 0 0 11.172 4H4v7.172a2 2 0 0 0 .586 1.414l7 7a2 2 0 0 0 2.828 0l5.172-5.172a2 2 0 0 0 0-2.828z"/><path stroke-width="2.25" d="M9 9h.01v.01H9z"/></g>`), g.Group(children),
	)
}

func BadgeThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M12.586 4.586A2 2 0 0 0 11.172 4H4v7.172a2 2 0 0 0 .586 1.414l7 7a2 2 0 0 0 2.828 0l5.172-5.172a2 2 0 0 0 0-2.828z"/><path stroke-width="1.5" d="M9 9h.01v.01H9z"/></g>`), g.Group(children),
	)
}

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17L19 7l-7-4v18l7-4L5 7"/>`), g.Group(children),
	)
}

func BluetoothBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 17L19 7l-7-4v18l7-4L5 7"/>`), g.Group(children),
	)
}

func BluetoothDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l7 4l-7 5l7 5l-7 4z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 17L19 7l-7-4v18l7-4L5 7"/></g>`), g.Group(children),
	)
}

func BluetoothFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.498 2.135a1 1 0 0 1 .998-.003l7 4a1 1 0 0 1 .085 1.682L13.721 12l5.86 4.186a1 1 0 0 1-.085 1.682l-7 4A1 1 0 0 1 11 21v-7.057l-5.419 3.87a1 1 0 1 1-1.162-1.627L10.279 12L4.42 7.814a1 1 0 1 1 1.16-1.628l5.42 3.87V3a1 1 0 0 1 .498-.865M13 13.943l4.148 2.963L13 19.276zm0-3.886V4.723l4.148 2.37z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BluetoothLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 17L19 7l-7-4v18l7-4L5 7"/>`), g.Group(children),
	)
}

func BluetoothThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 17L19 7l-7-4v18l7-4L5 7"/>`), g.Group(children),
	)
}

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 3H8a2 2 0 0 0-2 2v16l6-3l6 3V5a2 2 0 0 0-2-2"/>`), g.Group(children),
	)
}

func BookmarkBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M16 3H8a2 2 0 0 0-2 2v16l6-3l6 3V5a2 2 0 0 0-2-2"/>`), g.Group(children),
	)
}

func BookmarkDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M16 3H8a2 2 0 0 0-2 2v16l6-3l6 3V5a2 2 0 0 0-2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 3H8a2 2 0 0 0-2 2v16l6-3l6 3V5a2 2 0 0 0-2-2"/></g>`), g.Group(children),
	)
}

func BookmarkFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 2a3 3 0 0 0-3 3v16a1 1 0 0 0 1.447.894L12 19.118l5.553 2.776A1 1 0 0 0 19 21V5a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BookmarkLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 3H8a2 2 0 0 0-2 2v16l6-3l6 3V5a2 2 0 0 0-2-2"/>`), g.Group(children),
	)
}

func BookmarkOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M8.657 3H16a2 2 0 0 1 2 2v7.343M6 6v15l6-3l6 3v-3"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func BookmarkOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><path stroke-linejoin="round" d="M8.657 3H16a2 2 0 0 1 2 2v7.343M6 6v15l6-3l6 3v-3"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func BookmarkOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 18l6 3v-3L6 6v15z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.657 3H16a2 2 0 0 1 2 2v7.343M6 6v15l6-3l6 3v-3"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func BookmarkOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.657 3a1 1 0 0 1 1-1H16a3 3 0 0 1 3 3v7.343a1 1 0 0 1-2 0V5a1 1 0 0 0-1-1H8.657a1 1 0 0 1-1-1m11.061 14.304l-.023-.023L6.72 5.304l-.023-.023l-1.989-1.988a1 1 0 0 0-1.414 1.414L5 6.414V21a1 1 0 0 0 1.447.894L12 19.118l5.553 2.776A1 1 0 0 0 19 21v-.586l.293.293a1 1 0 0 0 1.414-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BookmarkOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M8.657 3H16a2 2 0 0 1 2 2v7.343M6 6v15l6-3l6 3v-3"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func BookmarkOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" d="M8.657 3H16a2 2 0 0 1 2 2v7.343M6 6v15l6-3l6 3v-3"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func BookmarkThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 3H8a2 2 0 0 0-2 2v16l6-3l6 3V5a2 2 0 0 0-2-2"/>`), g.Group(children),
	)
}

func Box(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm4-4h8l4 4H4zm0 8h4"/>`), g.Group(children),
	)
}

func BoxBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 8h16v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm4-4h8l4 4H4zm0 8h4"/>`), g.Group(children),
	)
}

func BoxDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 8h16v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h16v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm4-4h8l4 4H4zm0 8h4"/></g>`), g.Group(children),
	)
}

func BoxFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 3a1 1 0 0 0-.707.293l-4 4A1 1 0 0 0 3 7.954m0 .054V18a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V8.008l-.001-.054a1 1 0 0 0-.292-.661l-4-4A1 1 0 0 0 16 3H8m9.586 4l-2-2H8.414l-2 2zM7 12a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BoxLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 8h16v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm4-4h8l4 4H4zm0 8h4"/>`), g.Group(children),
	)
}

func BoxThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 8h16v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm4-4h8l4 4H4zm0 8h4"/>`), g.Group(children),
	)
}

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="2" d="M5 16h14v3a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zM4 7h16v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3" d="M12 12h.01v.01H12z"/><path stroke-width="2" d="M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2H9z"/></g>`), g.Group(children),
	)
}

func BriefcaseBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="2.5" d="M5 16h14v3a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zM4 7h16v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3.75" d="M12 12h.01v.01H12z"/><path stroke-width="2.5" d="M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2H9z"/></g>`), g.Group(children),
	)
}

func BriefcaseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 7h16v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M5 16h14v3a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zM4 7h16v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 12h.01v.01H12z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2H9z"/></g>`), g.Group(children),
	)
}

func BriefcaseFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 2a3 3 0 0 0-3 3v1H4a1 1 0 0 0-1 1v7a3 3 0 0 0 1 2.236V19a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-2.764c.614-.55 1-1.348 1-2.236V7a1 1 0 0 0-1-1h-4V5a3 3 0 0 0-3-3zm3 4V5a1 1 0 0 0-1-1h-2a1 1 0 0 0-1 1v1zM6 19v-2h12v2a1 1 0 0 1-1 1H7a1 1 0 0 1-1-1m6-8.5a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V12a1.5 1.5 0 0 0-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func BriefcaseLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="1.5" d="M5 16h14v3a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zM4 7h16v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="2.25" d="M12 12h.01v.01H12z"/><path stroke-width="1.5" d="M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2H9z"/></g>`), g.Group(children),
	)
}

func BriefcaseThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M5 16h14v3a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zM4 7h16v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="1.5" d="M12 12h.01v.01H12z"/><path d="M9 5a2 2 0 0 1 2-2h2a2 2 0 0 1 2 2v2H9z"/></g>`), g.Group(children),
	)
}

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M5 3h14v16a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linecap="round" d="M5 9h14M9 13h1m4 0h1m-6 4h1m4 0h1"/></g>`), g.Group(children),
	)
}

func CalculatorBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><path stroke-linejoin="round" d="M5 3h14v16a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linecap="round" d="M5 9h14M9 13h1m4 0h1m-6 4h1m4 0h1"/></g>`), g.Group(children),
	)
}

func CalculatorDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M5 9h14v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M5 9h14"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M5 3h14v16a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M9 13h1m4 0h1m-6 4h1m4 0h1"/></g>`), g.Group(children),
	)
}

func CalculatorFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 3a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v16a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3zm14 1v4H6V4zM8 13a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1m5 0a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1m-4 3a1 1 0 1 0 0 2h1a1 1 0 1 0 0-2zm4 1a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CalculatorLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M5 9h14"/><path stroke-linejoin="round" d="M5 3h14v16a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linecap="round" d="M9 13h1m4 0h1m-6 4h1m4 0h1"/></g>`), g.Group(children),
	)
}

func CalculatorThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" d="M5 9h14"/><path stroke-linejoin="round" d="M5 3h14v16a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linecap="round" d="M9 13h1m4 0h1m-6 4h1m4 0h1"/></g>`), g.Group(children),
	)
}

func CalendarAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 8v4m-2-2h4M4 8h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarAddBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 8v4m-2-2h4M4 8h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarAddDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 8h16v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 8v4m-2-2h4M4 8h16m-4-5v2M8 3v2"/></g>`), g.Group(children),
	)
}

func CalendarAddFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 3a1 1 0 1 0-2 0H9a1 1 0 0 0-2 0H4a1 1 0 0 0-1 1v14a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V4a1 1 0 0 0-1-1zM8 6a1 1 0 0 1-1-1H5v2h14V5h-2a1 1 0 1 1-2 0H9a1 1 0 0 1-1 1m4 5a1 1 0 0 1 1 1v1h1a1 1 0 1 1 0 2h-1v1a1 1 0 1 1-2 0v-1h-1a1 1 0 1 1 0-2h1v-1a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CalendarAddLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 8v4m-2-2h4M4 8h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarAddThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm8 8v4m-2-2h4M4 8h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm0 4h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm0 4h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 8h16v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm0 4h16m-4-5v2M8 3v2"/></g>`), g.Group(children),
	)
}

func CalendarOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 2a1 1 0 0 1 1 1h3a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3V4a1 1 0 0 1 1-1h3a1 1 0 0 1 2 0h6a1 1 0 0 1 1-1M7 5a1 1 0 0 0 2 0h6a1 1 0 1 0 2 0h2v2H5V5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CalendarOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm0 4h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm0 4h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 10h4M4 8h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarRemoveBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 10h4M4 8h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarRemoveDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 8h16v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 10h4M4 8h16m-4-5v2M8 3v2"/></g>`), g.Group(children),
	)
}

func CalendarRemoveFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 3a1 1 0 1 0-2 0H9a1 1 0 0 0-2 0H4a1 1 0 0 0-1 1v14a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V4a1 1 0 0 0-1-1zM8 6a1 1 0 0 1-1-1H5v2h14V5h-2a1 1 0 1 1-2 0H9a1 1 0 0 1-1 1m2 7a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CalendarRemoveLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 10h4M4 8h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarRemoveThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm6 10h4M4 8h16m-4-5v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm0 4h16M8 12h4m4-9v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm0 4h16M8 12h4m4-9v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 8h16v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm0 4h16M8 12h4m4-9v2M8 3v2"/></g>`), g.Group(children),
	)
}

func CalendarTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 3a1 1 0 1 0-2 0H9a1 1 0 0 0-2 0H4a1 1 0 0 0-1 1v14a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V4a1 1 0 0 0-1-1zM8 6a1 1 0 0 1-1-1H5v2h14V5h-2a1 1 0 1 1-2 0H9a1 1 0 0 1-1 1m0 5a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CalendarTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm0 4h16M8 12h4m4-9v2M8 3v2"/>`), g.Group(children),
	)
}

func CalendarTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm0 4h16M8 12h4m4-9v2M8 3v2"/>`), g.Group(children),
	)
}

func CameraImage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><circle cx="12" cy="12" r="3" stroke-linecap="round"/><path stroke-linecap="round" d="M17 2h2"/></g>`), g.Group(children),
	)
}

func CameraImageBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><circle cx="12" cy="12" r="3" stroke-linecap="round"/><path stroke-linecap="round" d="M17 2h2"/></g>`), g.Group(children),
	)
}

func CameraImageDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M21 6H3v10a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2zm-9 9a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><circle cx="12" cy="12" r="3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 2h2"/></g>`), g.Group(children),
	)
}

func CameraImageFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 3a1 1 0 1 0 0-2h-2a1 1 0 1 0 0 2zM2 5a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm10 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CameraImageLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><circle cx="12" cy="12" r="3" stroke-linecap="round"/><path stroke-linecap="round" d="M17 2h2"/></g>`), g.Group(children),
	)
}

func CameraImageThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z"/><circle cx="12" cy="12" r="3" stroke-linecap="round"/><path stroke-linecap="round" d="M17 2h2"/></g>`), g.Group(children),
	)
}

func CameraVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3 5h14v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm14 4l2.758-.69A1 1 0 0 1 21 9.28v5.44a1 1 0 0 1-1.242.97L17 15z"/>`), g.Group(children),
	)
}

func CameraVideoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M3 5h14v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm14 4l2.758-.69A1 1 0 0 1 21 9.28v5.44a1 1 0 0 1-1.242.97L17 15z"/>`), g.Group(children),
	)
}

func CameraVideoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 6h14v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3 5h14v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm14 4l2.758-.69A1 1 0 0 1 21 9.28v5.44a1 1 0 0 1-1.242.97L17 15z"/></g>`), g.Group(children),
	)
}

func CameraVideoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v2.72l1.515-.38A2 2 0 0 1 22 9.28v5.44a2 2 0 0 1-2.485 1.94L18 16.28V17a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm16 9.22l2 .5V9.28l-2 .5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CameraVideoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M3 5h14v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm14 4l2.758-.69A1 1 0 0 1 21 9.28v5.44a1 1 0 0 1-1.242.97L17 15z"/>`), g.Group(children),
	)
}

func CameraVideoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M3 5h14v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm14 4l2.758-.69A1 1 0 0 1 21 9.28v5.44a1 1 0 0 1-1.242.97L17 15z"/>`), g.Group(children),
	)
}

func Category(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="17" cy="7" r="3"/><circle cx="7" cy="17" r="3"/><path d="M14 14h6v5a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zM4 4h6v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1z"/></g>`), g.Group(children),
	)
}

func CategoryBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="17" cy="7" r="3"/><circle cx="7" cy="17" r="3"/><path d="M14 14h6v5a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zM4 4h6v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1z"/></g>`), g.Group(children),
	)
}

func CategoryDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="17" cy="7" r="3" fill="currentColor" opacity=".16"/><circle cx="7" cy="17" r="3" fill="currentColor" opacity=".16"/><path fill="currentColor" d="M14 14h6v5a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zM4 4h6v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1z" opacity=".16"/><circle cx="17" cy="7" r="3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="7" cy="17" r="3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 14h6v5a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zM4 4h6v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1z"/></g>`), g.Group(children),
	)
}

func CategoryFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 3a4 4 0 1 0 0 8a4 4 0 0 0 0-8M3 17a4 4 0 1 1 8 0a4 4 0 0 1-8 0m10-3a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v5a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2zM3 4a1 1 0 0 1 1-1h6a1 1 0 0 1 1 1v5a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CategoryLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="17" cy="7" r="3"/><circle cx="7" cy="17" r="3"/><path d="M14 14h6v5a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zM4 4h6v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1z"/></g>`), g.Group(children),
	)
}

func CategoryThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="17" cy="7" r="3"/><circle cx="7" cy="17" r="3"/><path d="M14 14h6v5a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zM4 4h6v5a1 1 0 0 1-1 1H5a1 1 0 0 1-1-1z"/></g>`), g.Group(children),
	)
}

func CertificateBadge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 22l-.555.832A1 1 0 0 0 16 22zm-3-2l.555-.832a1 1 0 0 0-1.11 0zm-3 2H8a1 1 0 0 0 1.555.832zM8.75 3.537l-.08.997zm1.685-.697l-.762-.648zM6.532 5.686l-.997.08zm2.154-2.154l.08-.997zM5.84 7.435l.648.761zm.697-1.684l.997-.08zm-.747 4.772l-.648.762zm0-3.046l-.648-.762zm.747 4.772l-.997-.08zm-.697-1.684l.648-.761zm2.846 3.903l.08.997zm-2.154-2.154l.997.08zm3.903 2.846l.761-.648zm-1.684-.697l-.08-.997zm4.772.747l.762.648zm-3.046 0l-.762.648zm4.772-.747l.08-.997zm-1.684.697l-.761-.648zm3.903-2.846l.997-.08zm-2.154 2.154l-.08.997zm2.846-3.903l.648.762zm-.697 1.684l-.997.08zm.747-4.772l.648-.762zm0 3.046l-.648-.761zm-.747-4.772l-.997-.08zm.697 1.684l-.648.761zm-2.846-3.903l-.08-.997zm2.154 2.154l.997.08zM13.565 2.84l.762-.648zm1.684.697l.08.997zm-1.726-.747l-.761.648zm-3.046 0l.761.648zM9 14.458l.044-.999zm6.555 6.71l-3-2l-1.11 1.664l3 2zm-4.11-2l-3 2l1.11 1.664l3-2zm1.317-15.73l.042.05l1.523-1.296l-.042-.05zm2.567 1.096l.065-.005l-.16-1.994l-.065.005zm1.142 1.072l-.005.065l1.994.16l.005-.065zm1.041 2.59l.05.042l1.296-1.523l-.05-.042zm.05 1.566l-.05.042l1.296 1.523l.05-.042zm-1.096 2.567l.005.065l1.994-.16l-.005-.065zm-1.072 1.142l-.065-.005l-.16 1.994l.065.005zm-2.59 1.041l-.042.05l1.523 1.296l.042-.05zm-1.566.05l-.042-.05l-1.523 1.296l.042.05zm-2.567-1.096l-.065.005l.16 1.994l.065-.005zm-1.142-1.072l.005-.065l-1.994-.16l-.005.065zm-1.041-2.59l-.05-.042l-1.296 1.523l.05.042zm-.05-1.566l.05-.042l-1.296-1.523l-.05.042zm1.096-2.567l-.005-.065l-1.994.16l.005.065zm1.072-1.142l.065.005l.16-1.994l-.065-.005zm2.59-1.041l.042-.05l-1.523-1.296l-.042.05zM8.671 4.534a3 3 0 0 0 2.525-1.046L9.673 2.192a1 1 0 0 1-.842.348zM7.529 5.606a1 1 0 0 1 1.077-1.077l.16-1.994a3 3 0 0 0-3.23 3.231zm-1.041 2.59a3 3 0 0 0 1.046-2.525l-1.994.16a1 1 0 0 1-.348.842zm-.05 1.566a1 1 0 0 1 0-1.524L5.142 6.715a3 3 0 0 0 0 4.57zm1.096 2.567a3 3 0 0 0-1.046-2.525l-1.296 1.523a1 1 0 0 1 .348.842zm1.072 1.142a1 1 0 0 1-1.077-1.077l-1.994-.16a3 3 0 0 0 3.231 3.23zm4.156 1.09a1 1 0 0 1-1.524 0l-1.523 1.297a3 3 0 0 0 4.57 0zm3.71-2.167a1 1 0 0 1-1.078 1.077l-.16 1.994a3 3 0 0 0 3.23-3.231zm1.04-2.59a3 3 0 0 0-1.046 2.525l1.994-.16a1 1 0 0 1 .348-.842zm.05-1.566a1 1 0 0 1 0 1.524l1.296 1.523a3 3 0 0 0 0-4.57zm-1.096-2.567a3 3 0 0 0 1.046 2.525l1.296-1.523a1 1 0 0 1-.348-.842zm-1.072-1.142a1 1 0 0 1 1.077 1.077l1.994.16a3 3 0 0 0-3.231-3.23zm-2.59-1.041a3 3 0 0 0 2.525 1.046l-.16-1.994a1 1 0 0 1-.842-.348zm1.48-1.346a3 3 0 0 0-4.569 0l1.523 1.296a1 1 0 0 1 1.524 0zm-3.088 12.37a3 3 0 0 0-2.152-1.053l-.088 1.998a1 1 0 0 1 .717.351zM9.044 13.46a3 3 0 0 0-.373.007l.16 1.994a1 1 0 0 1 .125-.003zM10 22v-7.542H8V22zm5.33-8.534a3 3 0 0 0-.374-.007l.088 1.998a1 1 0 0 1 .125.003zm-.374-.007a3 3 0 0 0-2.152 1.053l1.523 1.296a1 1 0 0 1 .717-.35zm-.956 1V22h2v-7.542z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 8l-3 3l-1-1"/></g>`), g.Group(children),
	)
}

func CertificateBadgeBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 22l-.693 1.04A1.25 1.25 0 0 0 16.25 22zm-3-2l.693-1.04a1.25 1.25 0 0 0-1.386 0zm-3 2H7.75a1.25 1.25 0 0 0 1.943 1.04zM8.75 3.537l-.1 1.246zm1.685-.697l-.952-.81zM6.532 5.686l-1.246.1zm2.154-2.154l.1-1.246zM5.84 7.435l.81.952zm.697-1.684l1.246-.1zm-.747 4.772l-.81.952zm0-3.046l-.81-.952zm.747 4.772l-1.246-.1zm-.697-1.684l.81-.952zm2.846 3.903l.1 1.246zm-2.154-2.154l1.246.1zm3.903 2.846l.952-.81zm-1.684-.697l-.1-1.246zm4.772.747l.952.81zm-3.046 0l-.952.81zm4.772-.747l.1-1.246zm-1.684.697l-.952-.81zm3.903-2.846l1.246-.1zm-2.154 2.154l-.1 1.246zm2.846-3.903l.81.952zm-.697 1.684l-1.246.1zm.747-4.772l.81-.952zm0 3.046l-.81-.952zm-.747-4.772l-1.246-.1zm.697 1.684l-.81.952zm-2.846-3.903l-.1-1.246zm2.154 2.154l1.246.1zM13.565 2.84l.952-.81zm1.684.697l.1 1.246zm-1.726-.747l-.952.81zm-3.046 0l.952.81zM9 14.458l.055-1.248zm6.693 6.502l-3-2l-1.386 2.08l3 2zm-4.386-2l-3 2l1.386 2.08l3-2zM12.57 3.6l.042.05l1.904-1.62l-.042-.05zm2.779 1.183l.064-.005l-.2-2.492l-.065.005zm.872.803l-.005.064l2.492.201l.005-.065zm1.128 2.8l.05.043l1.62-1.904l-.05-.042zm.05 1.185l-.05.042l1.62 1.904l.05-.042zm-1.183 2.779l.005.064l2.492-.2l-.005-.065zm-.803.872l-.064-.005l-.201 2.492l.065.005zm-2.8 1.128l-.043.05l1.904 1.62l.042-.05zm-1.185.05l-.042-.05l-1.904 1.62l.042.05zM8.65 13.217l-.064.005l.2 2.492l.065-.005zm-.872-.803l.005-.064l-2.492-.201l-.005.065zm-1.128-2.8L6.6 9.57l-1.62 1.904l.05.042zM6.6 8.428l.05-.042l-1.62-1.904l-.05.042zM7.783 5.65l-.005-.064l-2.492.2l.005.065zm.803-.872l.064.005l.201-2.492l-.065-.005zm2.8-1.128l.043-.05l-1.904-1.62l-.042.05zM8.65 4.783a3.25 3.25 0 0 0 2.737-1.133L9.483 2.03a.75.75 0 0 1-.632.261zm-.872.803a.75.75 0 0 1 .808-.808l.2-2.492a3.25 3.25 0 0 0-3.5 3.5zm-1.128 2.8A3.25 3.25 0 0 0 7.783 5.65l-2.492.201a.75.75 0 0 1-.261.632zM6.6 9.572a.75.75 0 0 1 0-1.142L4.98 6.525a3.25 3.25 0 0 0 0 4.95zm1.183 2.779A3.25 3.25 0 0 0 6.65 9.613l-1.62 1.904a.75.75 0 0 1 .261.632zm.803.872a.75.75 0 0 1-.808-.808l-2.492-.2a3.25 3.25 0 0 0 3.5 3.5zM12.57 14.4a.75.75 0 0 1-1.142 0l-1.904 1.62a3.25 3.25 0 0 0 4.95 0zm3.651-1.986a.75.75 0 0 1-.808.808l-.2 2.492a3.25 3.25 0 0 0 3.5-3.5zm1.128-2.8a3.25 3.25 0 0 0-1.133 2.736l2.492-.201a.75.75 0 0 1 .261-.632zm.05-1.185a.75.75 0 0 1 0 1.142l1.62 1.904a3.25 3.25 0 0 0 0-4.95zM16.217 5.65a3.25 3.25 0 0 0 1.133 2.737l1.62-1.904a.75.75 0 0 1-.261-.632zm-.803-.872a.75.75 0 0 1 .808.808l2.492.2a3.25 3.25 0 0 0-3.5-3.5zm-2.8-1.128a3.25 3.25 0 0 0 2.736 1.133l-.201-2.492a.75.75 0 0 1-.632-.261zm1.861-1.67a3.25 3.25 0 0 0-4.95 0l1.904 1.62a.75.75 0 0 1 1.142 0zm-3.088 12.37a3.25 3.25 0 0 0-2.332-1.14l-.11 2.497a.75.75 0 0 1 .538.263zm-2.332-1.14a3 3 0 0 0-.405.007l.201 2.492a1 1 0 0 1 .094-.002zM10.25 22v-7.542h-2.5V22zm5.1-8.783a3 3 0 0 0-.405-.007l.11 2.497q.046-.001.094.002zm-.405-.007a3.25 3.25 0 0 0-2.332 1.14l1.904 1.62a.75.75 0 0 1 .538-.263zm-1.195 1.248V22h2.5v-7.542z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m14 8l-3 3l-1-1"/></g>`), g.Group(children),
	)
}

func CertificateBadgeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M10.477 2.79a2 2 0 0 1 3.046 0l.042.05a2 2 0 0 0 1.684.697l.065-.005a2 2 0 0 1 2.154 2.154l-.005.065a2 2 0 0 0 .697 1.684l.05.042a2 2 0 0 1 0 3.046l-.05.042a2 2 0 0 0-.697 1.684l.005.065a2 2 0 0 1-2.154 2.154l-.065-.005a2 2 0 0 0-1.684.697l-.042.05a2 2 0 0 1-3.046 0l-.042-.05a2 2 0 0 0-1.684-.697l-.065.005a2 2 0 0 1-2.154-2.154l.005-.065a2 2 0 0 0-.697-1.684l-.05-.042a2 2 0 0 1 0-3.046l.05-.042a2 2 0 0 0 .697-1.684l-.005-.065a2 2 0 0 1 2.154-2.154l.065.005a2 2 0 0 0 1.684-.697z" opacity=".16"/><path fill="currentColor" d="m15 22l-.555.832A1 1 0 0 0 16 22zm-3-2l.555-.832a1 1 0 0 0-1.11 0zm-3 2H8a1 1 0 0 0 1.555.832zM8.75 3.537l-.08.997zm1.685-.697l-.762-.648zM6.532 5.686l-.997.08zm2.154-2.154l.08-.997zM5.84 7.435l.648.761zm.697-1.684l.997-.08zm-.747 4.772l-.648.762zm0-3.046l-.648-.762zm.747 4.772l-.997-.08zm-.697-1.684l.648-.761zm2.846 3.903l.08.997zm-2.154-2.154l.997.08zm3.903 2.846l.761-.648zm-1.684-.697l-.08-.997zm4.772.747l.762.648zm-3.046 0l-.762.648zm4.772-.747l.08-.997zm-1.684.697l-.761-.648zm3.903-2.846l.997-.08zm-2.154 2.154l-.08.997zm2.846-3.903l.648.762zm-.697 1.684l-.997.08zm.747-4.772l.648-.762zm0 3.046l-.648-.761zm-.747-4.772l-.997-.08zm.697 1.684l-.648.761zm-2.846-3.903l-.08-.997zm2.154 2.154l.997.08zM13.565 2.84l.762-.648zm1.684.697l.08.997zm-1.726-.747l-.761.648zm-3.046 0l.761.648zM9 14.458l.044-.999zm6.555 6.71l-3-2l-1.11 1.664l3 2zm-4.11-2l-3 2l1.11 1.664l3-2zm1.317-15.73l.042.05l1.523-1.296l-.042-.05zm2.567 1.096l.065-.005l-.16-1.994l-.065.005zm1.142 1.072l-.005.065l1.994.16l.005-.065zm1.041 2.59l.05.042l1.296-1.523l-.05-.042zm.05 1.566l-.05.042l1.296 1.523l.05-.042zm-1.096 2.567l.005.065l1.994-.16l-.005-.065zm-1.072 1.142l-.065-.005l-.16 1.994l.065.005zm-2.59 1.041l-.042.05l1.523 1.296l.042-.05zm-1.566.05l-.042-.05l-1.523 1.296l.042.05zm-2.567-1.096l-.065.005l.16 1.994l.065-.005zm-1.142-1.072l.005-.065l-1.994-.16l-.005.065zm-1.041-2.59l-.05-.042l-1.296 1.523l.05.042zm-.05-1.566l.05-.042l-1.296-1.523l-.05.042zm1.096-2.567l-.005-.065l-1.994.16l.005.065zm1.072-1.142l.065.005l.16-1.994l-.065-.005zm2.59-1.041l.042-.05l-1.523-1.296l-.042.05zM8.671 4.534a3 3 0 0 0 2.525-1.046L9.673 2.192a1 1 0 0 1-.842.348zM7.529 5.606a1 1 0 0 1 1.077-1.077l.16-1.994a3 3 0 0 0-3.23 3.231zm-1.041 2.59a3 3 0 0 0 1.046-2.525l-1.994.16a1 1 0 0 1-.348.842zm-.05 1.566a1 1 0 0 1 0-1.524L5.142 6.715a3 3 0 0 0 0 4.57zm1.096 2.567a3 3 0 0 0-1.046-2.525l-1.296 1.523a1 1 0 0 1 .348.842zm1.072 1.142a1 1 0 0 1-1.077-1.077l-1.994-.16a3 3 0 0 0 3.231 3.23zm4.156 1.09a1 1 0 0 1-1.524 0l-1.523 1.297a3 3 0 0 0 4.57 0zm3.71-2.167a1 1 0 0 1-1.078 1.077l-.16 1.994a3 3 0 0 0 3.23-3.231zm1.04-2.59a3 3 0 0 0-1.046 2.525l1.994-.16a1 1 0 0 1 .348-.842zm.05-1.566a1 1 0 0 1 0 1.524l1.296 1.523a3 3 0 0 0 0-4.57zm-1.096-2.567a3 3 0 0 0 1.046 2.525l1.296-1.523a1 1 0 0 1-.348-.842zm-1.072-1.142a1 1 0 0 1 1.077 1.077l1.994.16a3 3 0 0 0-3.231-3.23zm-2.59-1.041a3 3 0 0 0 2.525 1.046l-.16-1.994a1 1 0 0 1-.842-.348zm1.48-1.346a3 3 0 0 0-4.569 0l1.523 1.296a1 1 0 0 1 1.524 0zm-3.088 12.37a3 3 0 0 0-2.152-1.053l-.088 1.998a1 1 0 0 1 .717.351zM9.044 13.46a3 3 0 0 0-.373.007l.16 1.994a1 1 0 0 1 .125-.003zM10 22v-7.542H8V22zm5.33-8.534a3 3 0 0 0-.374-.007l.088 1.998a1 1 0 0 1 .125.003zm-.374-.007a3 3 0 0 0-2.152 1.053l1.523 1.296a1 1 0 0 1 .717-.35zm-.956 1V22h2v-7.542z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 8l-3 3l-1-1"/></g>`), g.Group(children),
	)
}

func CertificateBadgeFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.285 2.142a3 3 0 0 0-4.57 0l-.042.05a1 1 0 0 1-.842.348l-.065-.005a3 3 0 0 0-3.23 3.231l.004.065a1 1 0 0 1-.348.842l-.05.042a3 3 0 0 0 0 4.57l.05.042a1 1 0 0 1 .348.842l-.005.065A3 3 0 0 0 8 15.429V22a1 1 0 0 0 1.555.832L12 21.202l2.445 1.63A1 1 0 0 0 16 22v-6.57a3 3 0 0 0 2.465-3.196l-.005-.065a1 1 0 0 1 .348-.842l.05-.042a3 3 0 0 0 0-4.57l-.05-.042a1 1 0 0 1-.348-.842l.005-.065a3 3 0 0 0-3.231-3.23l-.065.004a1 1 0 0 1-.842-.348zM10 20.132V16.15a3 3 0 0 0 4 0v3.98l-1.445-.963a1 1 0 0 0-1.11 0zm4.707-11.425a1 1 0 0 0-1.414-1.414L11 9.586l-.293-.293a1 1 0 0 0-1.414 1.414l1 1a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CertificateBadgeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 22l-.416.624A.75.75 0 0 0 15.75 22zm-3-2l.416-.624a.75.75 0 0 0-.832 0zm-3 2h-.75a.75.75 0 0 0 1.166.624zM8.75 3.537l-.06.748zm1.685-.697l-.572-.486zM6.532 5.686l-.748.06zm2.154-2.154l.06-.748zM5.84 7.435l.486.57zm.697-1.684l.748-.06zm-.747 4.772l-.486.571zm0-3.046l-.486-.571zm.747 4.772l-.747-.06zm-.697-1.684l.486-.57zm2.846 3.903l.06.748zm-2.154-2.154l.747.06zm3.903 2.846l.57-.486zm-1.684-.697l-.06-.748zm4.772.747l.571.486zm-3.046 0l-.571.486zm4.772-.747l.06-.748zm-1.684.697l-.57-.486zm3.903-2.846l.748-.06zm-2.154 2.154l-.06.748zm2.846-3.903l.486.572zm-.697 1.684l-.748.06zm.747-4.772l.486-.571zm0 3.046l-.486-.571zm-.747-4.772l-.748-.06zm.697 1.684l-.486.57zm-2.846-3.903l-.06-.748zm2.154 2.154l.748.06zM13.565 2.84l.572-.486zm1.684.697l.06.748zm-1.726-.747l-.571.486zm-3.046 0l.571.486zM9 14.458l.033-.749zm6.416 6.918l-3-2l-.832 1.248l3 2zm-3.832-2l-3 2l.832 1.248l3-2zm1.368-16.1l.042.05l1.143-.972l-.043-.05zm2.357 1.009l.065-.006l-.12-1.495l-.065.006zm1.412 1.34l-.006.066l1.495.12l.006-.065zm.953 2.38l.05.043l.972-1.142l-.05-.043zm.05 1.947l-.05.042l.972 1.143l.05-.043zm-1.009 2.357l.006.065l1.495-.12l-.006-.065zm-1.34 1.412l-.066-.006l-.12 1.495l.065.006zm-2.38.953l-.043.05l1.142.972l.043-.05zm-1.947.05l-.042-.05l-1.143.972l.043.05zm-2.357-1.009l-.065.005l.12 1.496l.065-.005zm-1.412-1.34l.006-.066l-1.495-.12l-.006.065zm-.953-2.38l-.05-.043l-.972 1.142l.05.043zm-.05-1.947l.05-.042l-.972-1.143l-.05.043zm1.009-2.357l-.006-.065l-1.495.12l.006.065zm1.34-1.412l.066.006l.12-1.495l-.065-.006zm2.38-.953l.043-.05l-1.142-.972l-.043.05zm-2.314.959a2.75 2.75 0 0 0 2.315-.96l-1.143-.971a1.25 1.25 0 0 1-1.052.436zm-1.412 1.34A1.25 1.25 0 0 1 8.626 4.28l.12-1.495a2.75 2.75 0 0 0-2.962 2.962zm-.953 2.38a2.75 2.75 0 0 0 .959-2.314l-1.495.12c.032.4-.13.792-.436 1.052zm-.05 1.947a1.25 1.25 0 0 1 0-1.904l-.972-1.142a2.75 2.75 0 0 0 0 4.188zm1.009 2.357a2.75 2.75 0 0 0-.96-2.315l-.971 1.143c.306.26.468.652.436 1.052zm1.34 1.412a1.25 1.25 0 0 1-1.346-1.347l-1.495-.12a2.75 2.75 0 0 0 2.962 2.962zm4.327 1.003a1.25 1.25 0 0 1-1.904 0l-1.142.972a2.75 2.75 0 0 0 4.188 0zm3.769-2.35a1.25 1.25 0 0 1-1.347 1.346l-.12 1.496a2.75 2.75 0 0 0 2.962-2.962zm.953-2.38a2.75 2.75 0 0 0-.959 2.315l1.495-.12c-.032-.4.13-.792.436-1.052zm.05-1.946a1.25 1.25 0 0 1 0 1.904l.972 1.142a2.75 2.75 0 0 0 0-4.188zm-1.009-2.357a2.75 2.75 0 0 0 .96 2.315l.971-1.143a1.25 1.25 0 0 1-.436-1.052zm-1.34-1.412a1.25 1.25 0 0 1 1.346 1.347l1.495.12a2.75 2.75 0 0 0-2.962-2.962zm-2.38-.953a2.75 2.75 0 0 0 2.314.959l-.12-1.495a1.25 1.25 0 0 1-1.052-.436zm1.1-1.022a2.75 2.75 0 0 0-4.19 0l1.143.972a1.25 1.25 0 0 1 1.904 0zm-3.09 12.37a2.75 2.75 0 0 0-1.972-.965l-.066 1.499c.344.015.67.172.896.438zm-1.972-.965a3 3 0 0 0-.342.006l.12 1.495q.078-.006.156-.002zM9.75 22v-7.542h-1.5V22zm5.56-8.285a3 3 0 0 0-.343-.006l.066 1.499q.078-.004.156.002zm-.343-.006a2.75 2.75 0 0 0-1.973.965l1.143.972a1.25 1.25 0 0 1 .896-.438zm-.717.75V22h1.5v-7.542z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14 8l-3 3l-1-1"/></g>`), g.Group(children),
	)
}

func CertificateBadgeThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 22l-.277.416A.5.5 0 0 0 15.5 22zm-3-2l.277-.416a.5.5 0 0 0-.554 0zm-3 2h-.5a.5.5 0 0 0 .777.416zM8.75 3.537l-.04.499zm1.685-.697l-.381-.324zM6.532 5.686l-.498.04zm2.154-2.154l.04-.498zM5.84 7.435l.324.38zm.697-1.684l.499-.04zm-.747 4.772l-.324.381zm0-3.046l-.324-.381zm.747 4.772l-.498-.04zm-.697-1.684l.324-.38zm2.846 3.903l.04.498zm-2.154-2.154l.498.04zm3.903 2.846l.38-.324zm-1.684-.697l-.04-.498zm4.772.747l.381.324zm-3.046 0l-.381.324zm4.772-.747l.04-.498zm-1.684.697l-.38-.324zm3.903-2.846l.498-.04zm-2.154 2.154l-.04.498zm2.846-3.903l.324.381zm-.697 1.684l-.498.04zm.747-4.772l.324-.381zm0 3.046l-.324-.38zm-.747-4.772l-.498-.04zm.697 1.684l-.324.38zm-2.846-3.903l-.04-.498zm2.154 2.154l.498.04zM13.565 2.84l.381-.324zm1.684.697l.04.499zm-1.726-.747l-.38.324zm-3.046 0l.38.324zM9 14.458l.022-.5zm6.277 7.126l-3-2l-.554.832l3 2zm-3.554-2l-3 2l.554.832l3-2zm1.42-16.47l.042.05l.761-.648l-.042-.05zm2.146.922l.065-.006l-.08-.996l-.065.005zm1.68 1.61l-.005.065l.997.08l.006-.065zm.867 2.17l.05.042l.648-.762l-.05-.042zm.05 2.326l-.05.043l.648.761l.05-.042zm-.921 2.147l.005.065l.997-.08l-.006-.065zm-1.61 1.68l-.066-.005l-.08.997l.065.005zm-2.17.867l-.043.05l.762.648l.042-.05zm-2.327.05l-.043-.05l-.761.648l.042.05zm-2.147-.921l-.065.005l.08.996l.065-.005zm-1.68-1.61l.005-.066l-.997-.08l-.005.065zm-.867-2.17l-.05-.043l-.648.762l.05.042zm-.05-2.327l.05-.043l-.648-.761l-.05.042zm.922-2.147l-.006-.065l-.996.08l.005.065zm1.61-1.68l.065.005l.08-.997l-.065-.005zm2.17-.867l.042-.05l-.762-.648l-.042.05zm-2.105.872a2.5 2.5 0 0 0 2.105-.872l-.762-.648a1.5 1.5 0 0 1-1.263.523zm-1.68 1.61A1.5 1.5 0 0 1 8.645 4.03l.08-.996a2.5 2.5 0 0 0-2.692 2.692zm-.867 2.17a2.5 2.5 0 0 0 .872-2.105l-.997.08a1.5 1.5 0 0 1-.523 1.263zm-.05 2.326a1.5 1.5 0 0 1 0-2.284l-.648-.762a2.5 2.5 0 0 0 0 3.808zm.922 2.147a2.5 2.5 0 0 0-.872-2.104l-.648.761a1.5 1.5 0 0 1 .523 1.263zm1.61 1.68a1.5 1.5 0 0 1-1.616-1.615l-.996-.08a2.5 2.5 0 0 0 2.692 2.693zm4.496.917a1.5 1.5 0 0 1-2.284 0l-.762.648a2.5 2.5 0 0 0 3.808 0zm3.828-2.532a1.5 1.5 0 0 1-1.616 1.616l-.08.996a2.5 2.5 0 0 0 2.693-2.692zm.866-2.17a2.5 2.5 0 0 0-.871 2.105l.996-.08a1.5 1.5 0 0 1 .523-1.263zm.05-2.326a1.5 1.5 0 0 1 0 2.284l.648.762a2.5 2.5 0 0 0 0-3.808zm-.921-2.147a2.5 2.5 0 0 0 .871 2.104l.648-.761a1.5 1.5 0 0 1-.523-1.263zm-1.61-1.68a1.5 1.5 0 0 1 1.615 1.615l.997.08a2.5 2.5 0 0 0-2.693-2.692zm-2.17-.867a2.5 2.5 0 0 0 2.104.872l-.08-.997a1.5 1.5 0 0 1-1.263-.523zm.719-.698a2.5 2.5 0 0 0-3.808 0l.762.648a1.5 1.5 0 0 1 2.284 0zm-3.088 12.37a2.5 2.5 0 0 0-1.794-.877l-.044.999a1.5 1.5 0 0 1 1.076.526zm-1.794-.877a3 3 0 0 0-.311.005l.08.997q.094-.008.187-.003zM9.5 22v-7.542h-1V22zm5.79-8.036a3 3 0 0 0-.312-.005l.044.999q.093-.005.187.003zm-.312-.005a2.5 2.5 0 0 0-1.793.877l.761.648a1.5 1.5 0 0 1 1.076-.526zm-.478.5V22h1v-7.542z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14 8l-3 3l-1-1"/></g>`), g.Group(children),
	)
}

func Check(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7L10 17l-5-5"/>`), g.Group(children),
	)
}

func CheckBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M20 7L10 17l-5-5"/>`), g.Group(children),
	)
}

func CheckCircleOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CheckCircleOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CheckCircleOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CheckCircleOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m13.707-1.293a1 1 0 0 0-1.414-1.414L11 12.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CheckCircleOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CheckCircleOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CheckCircleTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 7l1 1l3-3"/><path d="M21 12a9 9 0 1 1-9-9"/></g>`), g.Group(children),
	)
}

func CheckCircleTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="m15 7l1 1l3-3"/><path d="M21 12a9 9 0 1 1-9-9"/></g>`), g.Group(children),
	)
}

func CheckCircleTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m15 7l1 1l3-3"/><path d="M21 12a9 9 0 1 1-9-9"/></g>`), g.Group(children),
	)
}

func CheckCircleTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 12a8 8 0 0 1 8-8a1 1 0 1 0 0-2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10a1 1 0 1 0-2 0a8 8 0 1 1-16 0m15.707-6.293a1 1 0 0 0-1.414-1.414L16 6.586l-.293-.293a1 1 0 1 0-1.414 1.414l1 1a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CheckCircleTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m15 7l1 1l3-3"/><path d="M21 12a9 9 0 1 1-9-9"/></g>`), g.Group(children),
	)
}

func CheckCircleTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m15 7l1 1l3-3"/><path d="M21 12a9 9 0 1 1-9-9"/></g>`), g.Group(children),
	)
}

func CheckDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 7L10 17l-5-5"/>`), g.Group(children),
	)
}

func CheckFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20.207 6.793a1 1 0 0 1 0 1.414l-9.5 9.5a1 1 0 0 1-1.414 0l-4.5-4.5a1 1 0 0 1 1.414-1.414L10 15.586l8.793-8.793a1 1 0 0 1 1.414 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CheckLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 7L10 17l-5-5"/>`), g.Group(children),
	)
}

func CheckSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CheckSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CheckSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CheckSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm12.707 6.707a1 1 0 0 0-1.414-1.414L11 12.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CheckSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CheckSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CheckThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20 7L10 17l-5-5"/>`), g.Group(children),
	)
}

func Cheque(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm4 9h3m-3-4h10"/>`), g.Group(children),
	)
}

func ChequeBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm4 9h3m-3-4h10"/>`), g.Group(children),
	)
}

func ChequeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm4 9h3m-3-4h10"/></g>`), g.Group(children),
	)
}

func ChequeFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm4 5a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1m0 4a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ChequeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm4 9h3m-3-4h10"/>`), g.Group(children),
	)
}

func ChequeThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm4 9h3m-3-4h10"/>`), g.Group(children),
	)
}

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M11 8v5h5"/></g>`), g.Group(children),
	)
}

func ClockBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M11 8v5h5"/></g>`), g.Group(children),
	)
}

func ClockDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 8v5h5"/></g>`), g.Group(children),
	)
}

func ClockFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m10-4a1 1 0 1 0-2 0v5a1 1 0 0 0 1 1h5a1 1 0 1 0 0-2h-4z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ClockLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M11 8v5h5"/></g>`), g.Group(children),
	)
}

func ClockThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M11 8v5h5"/></g>`), g.Group(children),
	)
}

func Close(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10M7 17L17 7"/>`), g.Group(children),
	)
}

func CloseBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m7 7l10 10M7 17L17 7"/>`), g.Group(children),
	)
}

func CloseCircleOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m14 10l-4 4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func CloseCircleOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m14 10l-4 4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func CloseCircleOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 10l-4 4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func CloseCircleOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7.293-2.707a1 1 0 0 1 1.414 0L12 10.586l1.293-1.293a1 1 0 1 1 1.414 1.414L13.414 12l1.293 1.293a1 1 0 0 1-1.414 1.414L12 13.414l-1.293 1.293a1 1 0 0 1-1.414-1.414L10.586 12l-1.293-1.293a1 1 0 0 1 0-1.414" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloseCircleOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m14 10l-4 4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func CloseCircleOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m14 10l-4 4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func CloseCircleTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 1 1-9-9m7 2l-3 3m0-3l3 3"/>`), g.Group(children),
	)
}

func CloseCircleTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M21 12a9 9 0 1 1-9-9m7 2l-3 3m0-3l3 3"/>`), g.Group(children),
	)
}

func CloseCircleTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 12a9 9 0 1 1-9-9m7 2l-3 3m0-3l3 3"/>`), g.Group(children),
	)
}

func CloseCircleTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 4a8 8 0 1 0 8 8a1 1 0 1 1 2 0c0 5.523-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2a1 1 0 1 1 0 2m7.707.293a1 1 0 0 1 0 1.414l-.793.793l.793.793a1 1 0 0 1-1.414 1.414l-.793-.793l-.793.793a1 1 0 1 1-1.414-1.414l.793-.793l-.793-.793a1 1 0 0 1 1.414-1.414l.793.793l.793-.793a1 1 0 0 1 1.414 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloseCircleTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21 12a9 9 0 1 1-9-9m7 2l-3 3m0-3l3 3"/>`), g.Group(children),
	)
}

func CloseCircleTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21 12a9 9 0 1 1-9-9m7 2l-3 3m0-3l3 3"/>`), g.Group(children),
	)
}

func CloseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 7l10 10M7 17L17 7"/>`), g.Group(children),
	)
}

func CloseFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.793 6.793a1 1 0 0 1 1.414 0L12 10.586l3.793-3.793a1 1 0 1 1 1.414 1.414L13.414 12l3.793 3.793a1 1 0 0 1-1.414 1.414L12 13.414l-3.793 3.793a1 1 0 0 1-1.414-1.414L10.586 12L6.793 8.207a1 1 0 0 1 0-1.414" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloseLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7 7l10 10M7 17L17 7"/>`), g.Group(children),
	)
}

func CloseSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zM14 10l-4 4m0-4l4 4"/>`), g.Group(children),
	)
}

func CloseSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zM14 10l-4 4m0-4l4 4"/>`), g.Group(children),
	)
}

func CloseSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 6l-4 4m0-4l4 4"/></g>`), g.Group(children),
	)
}

func CloseSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm6.293 5.293a1 1 0 0 1 1.414 0L12 10.586l1.293-1.293a1 1 0 1 1 1.414 1.414L13.414 12l1.293 1.293a1 1 0 0 1-1.414 1.414L12 13.414l-1.293 1.293a1 1 0 0 1-1.414-1.414L10.586 12l-1.293-1.293a1 1 0 0 1 0-1.414" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloseSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 6l-4 4m0-4l4 4"/>`), g.Group(children),
	)
}

func CloseSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm10 6l-4 4m0-4l4 4"/>`), g.Group(children),
	)
}

func CloseThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7 7l10 10M7 17L17 7"/>`), g.Group(children),
	)
}

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/>`), g.Group(children),
	)
}

func CloudAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M12 11v4m-2-2h4"/></g>`), g.Group(children),
	)
}

func CloudAddBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M12 11v4m-2-2h4"/></g>`), g.Group(children),
	)
}

func CloudAddDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 11v4m-2-2h4"/></g>`), g.Group(children),
	)
}

func CloudAddFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.697 6.697a7.5 7.5 0 0 1 12.794 4.927A4.002 4.002 0 0 1 18.5 19.5h-12a5 5 0 0 1-1.667-9.715a7.5 7.5 0 0 1 1.864-3.088M12 10a1 1 0 0 1 1 1v1h1a1 1 0 1 1 0 2h-1v1a1 1 0 1 1-2 0v-1h-1a1 1 0 0 1 0-2h1v-1a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloudAddLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M12 11v4m-2-2h4"/></g>`), g.Group(children),
	)
}

func CloudAddThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M12 11v4m-2-2h4"/></g>`), g.Group(children),
	)
}

func CloudBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/>`), g.Group(children),
	)
}

func CloudClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M11 11v3h3"/></g>`), g.Group(children),
	)
}

func CloudClockBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M11 11v3h3"/></g>`), g.Group(children),
	)
}

func CloudClockDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 11v3h3"/></g>`), g.Group(children),
	)
}

func CloudClockFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.697 6.697a7.5 7.5 0 0 1 12.794 4.927A4.002 4.002 0 0 1 18.5 19.5h-12a5 5 0 0 1-1.667-9.715a7.5 7.5 0 0 1 1.864-3.088M12 11v2h2a1 1 0 0 1 0 2h-3a1 1 0 0 1-1-1v-3a1 1 0 1 1 2 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloudClockLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M11 11v3h3"/></g>`), g.Group(children),
	)
}

func CloudClockThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M11 11v3h3"/></g>`), g.Group(children),
	)
}

func CloudDownload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m10 13l2 2m0 0l2-2m-2 2V9"/></g>`), g.Group(children),
	)
}

func CloudDownloadBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m10 13l2 2m0 0l2-2m-2 2V9"/></g>`), g.Group(children),
	)
}

func CloudDownloadDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 13l2 2m0 0l2-2m-2 2V9"/></g>`), g.Group(children),
	)
}

func CloudDownloadFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.697 6.697a7.5 7.5 0 0 1 12.794 4.927A4.002 4.002 0 0 1 18.5 19.5h-12a5 5 0 0 1-1.667-9.715a7.5 7.5 0 0 1 1.864-3.088m4.596 9.01a1 1 0 0 0 1.414 0l2-2a1 1 0 0 0-1.414-1.414l-.293.293V9a1 1 0 1 0-2 0v3.586l-.293-.293a1 1 0 0 0-1.414 1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloudDownloadLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m10 13l2 2m0 0l2-2m-2 2V9"/></g>`), g.Group(children),
	)
}

func CloudDownloadThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m10 13l2 2m0 0l2-2m-2 2V9"/></g>`), g.Group(children),
	)
}

func CloudDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/></g>`), g.Group(children),
	)
}

func CloudError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="2" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-width="3" d="M12 15.5h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M12 12V9"/></g>`), g.Group(children),
	)
}

func CloudErrorBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="2.5" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-width="3.75" d="M12 15.5h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M12 12V9"/></g>`), g.Group(children),
	)
}

func CloudErrorDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 15.5h.01v.01H12z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12V9"/></g>`), g.Group(children),
	)
}

func CloudErrorFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.697 6.697a7.5 7.5 0 0 1 12.794 4.927A4.002 4.002 0 0 1 18.5 19.5h-12a5 5 0 0 1-1.667-9.715a7.5 7.5 0 0 1 1.864-3.088M12 13a1 1 0 0 1-1-1V9a1 1 0 0 1 2 0v3a1 1 0 0 1-1 1m-1.5 2.5a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloudErrorLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="1.5" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-width="2.25" d="M12 15.5h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M12 12V9"/></g>`), g.Group(children),
	)
}

func CloudErrorThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-width="1.5" d="M12 15.5h.01v.01H12z"/><path stroke-linecap="round" d="M12 12V9"/></g>`), g.Group(children),
	)
}

func CloudFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.303 6.697a7.5 7.5 0 0 0-12.47 3.088A5.002 5.002 0 0 0 6.5 19.5h12a4 4 0 0 0 .99-7.876a7.47 7.47 0 0 0-2.187-4.927" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloudLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/>`), g.Group(children),
	)
}

func CloudNo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m10 11l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func CloudNoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m10 11l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func CloudNoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 11l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func CloudNoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.697 6.697a7.5 7.5 0 0 1 12.794 4.927A4.002 4.002 0 0 1 18.5 19.5h-12a5 5 0 0 1-1.667-9.715a7.5 7.5 0 0 1 1.864-3.088m4.01 3.596a1 1 0 0 0-1.414 1.414L10.586 13l-1.293 1.293a1 1 0 1 0 1.414 1.414L12 14.414l1.293 1.293a1 1 0 0 0 1.414-1.414L13.414 13l1.293-1.293a1 1 0 0 0-1.414-1.414L12 11.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloudNoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m10 11l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func CloudNoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m10 11l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func CloudOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M7.05 7.05a6.96 6.96 0 0 0-1.81 3.129A4.502 4.502 0 0 0 6.5 19h12q.239 0 .469-.031m-8.203-13.86A6.99 6.99 0 0 1 16.95 7.05A6.98 6.98 0 0 1 19 12.035a3.5 3.5 0 0 1 2.917 4.225"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func CloudOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><path stroke-linejoin="round" d="M7.05 7.05a6.96 6.96 0 0 0-1.81 3.129A4.502 4.502 0 0 0 6.5 19h12q.239 0 .469-.031m-8.203-13.86A6.99 6.99 0 0 1 16.95 7.05A6.98 6.98 0 0 1 19 12.035a3.5 3.5 0 0 1 2.917 4.225"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func CloudOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path fill="currentColor" stroke-linejoin="round" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12q.239 0 .469-.031L7.05 7.05a6.96 6.96 0 0 0-1.81 3.129A4.5 4.5 0 0 0 2 14.5" opacity=".16"/><path stroke-linejoin="round" d="M7.05 7.05a6.96 6.96 0 0 0-1.81 3.129A4.502 4.502 0 0 0 6.5 19h12q.239 0 .469-.031m-8.203-13.86A6.99 6.99 0 0 1 16.95 7.05A6.98 6.98 0 0 1 19 12.035a3.5 3.5 0 0 1 2.917 4.225"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func CloudOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.707 3.293a1 1 0 0 0-1.414 1.414L5.68 7.094a8 8 0 0 0-1.248 2.308A5.501 5.501 0 0 0 6.5 20h12.085l.708.707a1 1 0 0 0 1.414-1.414l-1.023-1.023l-.015-.015L7.757 6.343zm11.536 4.464a5.99 5.99 0 0 0-5.302-1.663a1 1 0 1 1-.35-1.97a7.99 7.99 0 0 1 7.066 2.22a7.97 7.97 0 0 1 2.307 4.9a4.5 4.5 0 0 1 2.93 5.233a1 1 0 0 1-1.953-.433q.058-.261.059-.544a2.5 2.5 0 0 0-2.142-2.475A1 1 0 0 1 18 12.03a5.98 5.98 0 0 0-1.757-4.273" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloudOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M7.05 7.05a6.96 6.96 0 0 0-1.81 3.129A4.502 4.502 0 0 0 6.5 19h12q.239 0 .469-.031m-8.203-13.86A6.99 6.99 0 0 1 16.95 7.05A6.98 6.98 0 0 1 19 12.035a3.5 3.5 0 0 1 2.917 4.225"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func CloudOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" d="M7.05 7.05a6.96 6.96 0 0 0-1.81 3.129A4.502 4.502 0 0 0 6.5 19h12q.239 0 .469-.031m-8.203-13.86A6.99 6.99 0 0 1 16.95 7.05A6.98 6.98 0 0 1 19 12.035a3.5 3.5 0 0 1 2.917 4.225"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func CloudRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M10 13h4"/></g>`), g.Group(children),
	)
}

func CloudRemoveBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M10 13h4"/></g>`), g.Group(children),
	)
}

func CloudRemoveDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 13h4"/></g>`), g.Group(children),
	)
}

func CloudRemoveFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.697 6.697a7.5 7.5 0 0 1 12.794 4.927A4.002 4.002 0 0 1 18.5 19.5h-12a5 5 0 0 1-1.667-9.715a7.5 7.5 0 0 1 1.864-3.088M10 12a1 1 0 0 0 0 2h4a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloudRemoveLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M10 13h4"/></g>`), g.Group(children),
	)
}

func CloudRemoveThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="M10 13h4"/></g>`), g.Group(children),
	)
}

func CloudThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/>`), g.Group(children),
	)
}

func CloudUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m14 11l-2-2m0 0l-2 2m2-2v6"/></g>`), g.Group(children),
	)
}

func CloudUploadBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m14 11l-2-2m0 0l-2 2m2-2v6"/></g>`), g.Group(children),
	)
}

func CloudUploadDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 11l-2-2m0 0l-2 2m2-2v6"/></g>`), g.Group(children),
	)
}

func CloudUploadFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.697 6.697a7.5 7.5 0 0 1 12.794 4.927A4.002 4.002 0 0 1 18.5 19.5h-12a5 5 0 0 1-1.667-9.715a7.5 7.5 0 0 1 1.864-3.088m6.01 1.596a1 1 0 0 0-1.414 0l-2 2a1 1 0 1 0 1.414 1.414l.293-.293V15a1 1 0 1 0 2 0v-3.586l.293.293a1 1 0 0 0 1.414-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloudUploadLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m14 11l-2-2m0 0l-2 2m2-2v6"/></g>`), g.Group(children),
	)
}

func CloudUploadThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m14 11l-2-2m0 0l-2 2m2-2v6"/></g>`), g.Group(children),
	)
}

func CloudYes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m15 11l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CloudYesBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m15 11l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CloudYesDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 11l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CloudYesFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.697 6.697a7.5 7.5 0 0 1 12.794 4.927A4.002 4.002 0 0 1 18.5 19.5h-12a5 5 0 0 1-1.667-9.715a7.5 7.5 0 0 1 1.864-3.088m9.01 5.01a1 1 0 0 0-1.414-1.414L11 13.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CloudYesLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m15 11l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CloudYesThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M2 14.5A4.5 4.5 0 0 0 6.5 19h12a3.5 3.5 0 0 0 .5-6.965a7 7 0 0 0-13.76-1.857A4.5 4.5 0 0 0 2 14.5Z"/><path stroke-linecap="round" d="m15 11l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func Comment(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.891 1 4.127L3 21l4.873-1c1.236.64 2.64 1 4.127 1"/>`), g.Group(children),
	)
}

func CommentAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.891 1 4.127L3 21l4.873-1c1.236.64 2.64 1 4.127 1m0-11.999v6m-3-3h6"/>`), g.Group(children),
	)
}

func CommentAddBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.891 1 4.127L3 21l4.873-1c1.236.64 2.64 1 4.127 1m0-11.999v6m-3-3h6"/>`), g.Group(children),
	)
}

func CommentAddDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1" opacity=".16"/><path fill="currentColor" d="m4 16.127l.98.201a1 1 0 0 0-.092-.66zM7.873 20l.459-.888a1 1 0 0 0-.66-.092zM3 21l-.98-.201a1 1 0 0 0 1.181 1.18zm17-9a8 8 0 0 1-8 8v2c5.523 0 10-4.477 10-10zM4 12a8 8 0 0 1 8-8V2C6.477 2 2 6.477 2 12zm8-8a8 8 0 0 1 8 8h2c0-5.523-4.477-10-10-10zM4.888 15.668A7.96 7.96 0 0 1 4 12H2c0 1.651.4 3.211 1.112 4.586zM12 20a8 8 0 0 1-3.668-.888l-.918 1.776A9.96 9.96 0 0 0 12 22zm-8.98-4.074l-1 4.873l1.96.402l1-4.873zm.181 6.054l4.873-1l-.402-1.96l-4.873 1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9.001v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func CommentAddFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10a10 10 0 0 1-4.262-.951l-4.537.93a1 1 0 0 1-1.18-1.18l.93-4.537A10 10 0 0 1 2 12m10-4a1 1 0 0 1 1 1v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2H9a1 1 0 1 1 0-2h2V9a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CommentAddLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1m0-11.999v6m-3-3h6"/>`), g.Group(children),
	)
}

func CommentAddThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1m0-11.999v6m-3-3h6"/>`), g.Group(children),
	)
}

func CommentBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.891 1 4.127L3 21l4.873-1c1.236.64 2.64 1 4.127 1"/>`), g.Group(children),
	)
}

func CommentCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.891 1 4.127L3 21l4.873-1c1.236.64 2.64 1 4.127 1"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CommentCheckBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.891 1 4.127L3 21l4.873-1c1.236.64 2.64 1 4.127 1"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CommentCheckDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CommentCheckFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10a10 10 0 0 1-4.262-.951l-4.537.93a1 1 0 0 1-1.18-1.18l.93-4.537A10 10 0 0 1 2 12m13.707-1.293a1 1 0 1 0-1.414-1.414L11 12.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CommentCheckLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CommentCheckThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1"/><path d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func CommentClose(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.891 1 4.127L3 21l4.873-1c1.236.64 2.64 1 4.127 1m-2-11l4 4m-4 0l4-4"/>`), g.Group(children),
	)
}

func CommentCloseBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.891 1 4.127L3 21l4.873-1c1.236.64 2.64 1 4.127 1m-2-11l4 4m-4 0l4-4"/>`), g.Group(children),
	)
}

func CommentCloseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1m-2-11l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func CommentCloseFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10a10 10 0 0 1-4.262-.951l-4.537.93a1 1 0 0 1-1.18-1.18l.93-4.537A10 10 0 0 1 2 12m8.707-2.707a1 1 0 0 0-1.414 1.414L10.586 12l-1.293 1.293a1 1 0 1 0 1.414 1.414L12 13.414l1.293 1.293a1 1 0 1 0 1.414-1.414L13.414 12l1.293-1.293a1 1 0 1 0-1.414-1.414L12 10.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CommentCloseLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1m-2-11l4 4m-4 0l4-4"/>`), g.Group(children),
	)
}

func CommentCloseThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1m-2-11l4 4m-4 0l4-4"/>`), g.Group(children),
	)
}

func CommentDots(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2" d="M12 21a9 9 0 1 0-8-4.873L3 21l4.873-1c1.236.639 2.64 1 4.127 1"/><path stroke-width="3" d="M7.5 12h.01v.01H7.5zm4.5 0h.01v.01H12zm4.5 0h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func CommentDotsBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M12 21a9 9 0 1 0-8-4.873L3 21l4.873-1c1.236.639 2.64 1 4.127 1"/><path stroke-width="3.75" d="M7.5 12h.01v.01H7.5zm4.5 0h.01v.01H12zm4.5 0h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func CommentDotsDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 21a9 9 0 1 0-8-4.873L3 21l4.873-1c1.236.639 2.64 1 4.127 1" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0-8-4.873L3 21l4.873-1c1.236.639 2.64 1 4.127 1"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 12h.01v.01H12zm-4.5 0h.01v.01H7.5zm9 0h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func CommentDotsFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10a10 10 0 0 1-4.262-.951l-4.537.93a1 1 0 0 1-1.18-1.18l.93-4.537A10 10 0 0 1 2 12m5.5-1.5A1.5 1.5 0 0 0 6 12v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V12a1.5 1.5 0 0 0-1.5-1.5zm4.5 0a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V12a1.5 1.5 0 0 0-1.5-1.5zm3 1.5a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5h-.01a1.5 1.5 0 0 1-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CommentDotsLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="1.5" d="M12 21a9 9 0 1 0-8-4.873L3 21l4.873-1c1.236.639 2.64 1 4.127 1"/><path stroke-width="2.25" d="M7.5 12h.01v.01H7.5zm4.5 0h.01v.01H12zm4.5 0h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func CommentDotsThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M12 21a9 9 0 1 0-8-4.873L3 21l4.873-1c1.236.639 2.64 1 4.127 1"/><path stroke-width="1.5" d="M7.5 12h.01v.01H7.5zm4.5 0h.01v.01H12zm4.5 0h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func CommentDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1"/></g>`), g.Group(children),
	)
}

func CommentFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12a10 10 0 0 0 .951 4.262l-.93 4.537a1 1 0 0 0 1.18 1.18l4.537-.93c1.294.61 2.74.95 4.262.95c5.523 0 10-4.476 10-10c0-5.522-4.477-10-10-10" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CommentLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1"/>`), g.Group(children),
	)
}

func CommentRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.891 1 4.127L3 21l4.873-1c1.236.64 2.64 1 4.127 1m-3-8.999h6"/>`), g.Group(children),
	)
}

func CommentRemoveBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.891 1 4.127L3 21l4.873-1c1.236.64 2.64 1 4.127 1m-3-8.999h6"/>`), g.Group(children),
	)
}

func CommentRemoveDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1m-3-8.999h6"/></g>`), g.Group(children),
	)
}

func CommentRemoveFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10a10 10 0 0 1-4.262-.951l-4.537.93a1 1 0 0 1-1.18-1.18l.93-4.537A10 10 0 0 1 2 12m7-1a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CommentRemoveLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1m-3-8.999h6"/>`), g.Group(children),
	)
}

func CommentRemoveThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1m-3-8.999h6"/>`), g.Group(children),
	)
}

func CommentThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 21a9 9 0 1 0-9-9c0 1.488.36 2.89 1 4.127L3 21l4.873-1c1.236.639 2.64 1 4.127 1"/>`), g.Group(children),
	)
}

func Compare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7m4-16h1a2 2 0 0 1 2 2v1m0 10v1a2 2 0 0 1-2 2h-1m3-9v2M12 2v20"/>`), g.Group(children),
	)
}

func CompareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7m4-16h1a2 2 0 0 1 2 2v1m0 10v1a2 2 0 0 1-2 2h-1m3-9v2M12 2v20"/>`), g.Group(children),
	)
}

func CompareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6 4h6v16H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7m4-16h1a2 2 0 0 1 2 2v1m0 10v1a2 2 0 0 1-2 2h-1m3-9v2M12 2v20"/></g>`), g.Group(children),
	)
}

func CompareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 2a1 1 0 1 0-2 0v1H6a3 3 0 0 0-3 3v12a3 3 0 0 0 3 3h5v1a1 1 0 0 0 2 0v-1a1 1 0 0 0 0-2V5a1 1 0 1 0 0-2zm4 1a1 1 0 1 0 0 2h1a1 1 0 0 1 1 1v1a1 1 0 1 0 2 0V6a3 3 0 0 0-3-3zm4 8a1 1 0 1 0-2 0v2a1 1 0 0 0 2 0zm0 6a1 1 0 1 0-2 0v1a1 1 0 0 1-1 1h-1a1 1 0 1 0 0 2h1a3 3 0 0 0 3-3z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CompareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7m4-16h1a2 2 0 0 1 2 2v1m0 10v1a2 2 0 0 1-2 2h-1m3-9v2M12 2v20"/>`), g.Group(children),
	)
}

func CompareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13 4H6a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h7m4-16h1a2 2 0 0 1 2 2v1m0 10v1a2 2 0 0 1-2 2h-1m3-9v2M12 2v20"/>`), g.Group(children),
	)
}

func Component(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m12 3l3 3l-3 3l-3-3zm0 12l3 3l-3 3l-3-3zm6-6l3 3l-3 3l-3-3zM6 9l3 3l-3 3l-3-3z"/>`), g.Group(children),
	)
}

func ComponentBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="m12 3l3 3l-3 3l-3-3zm0 12l3 3l-3 3l-3-3zm6-6l3 3l-3 3l-3-3zM6 9l3 3l-3 3l-3-3z"/>`), g.Group(children),
	)
}

func ComponentDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l3 3l-3 3l-3-3zm0 12l3 3l-3 3l-3-3zm6-6l3 3l-3 3l-3-3zM6 9l3 3l-3 3l-3-3z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m12 3l3 3l-3 3l-3-3zm0 12l3 3l-3 3l-3-3zm6-6l3 3l-3 3l-3-3zM6 9l3 3l-3 3l-3-3z"/></g>`), g.Group(children),
	)
}

func ComponentFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.707 2.293a1 1 0 0 0-1.414 0l-3 3a1 1 0 0 0 0 1.414l3 3a1 1 0 0 0 1.414 0l3-3a1 1 0 0 0 0-1.414zm-1.414 12a1 1 0 0 1 1.414 0l3 3a1 1 0 0 1 0 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 0 1 0-1.414zm6-6a1 1 0 0 1 1.414 0l3 3a1 1 0 0 1 0 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 0 1 0-1.414zm-12 0a1 1 0 0 1 1.414 0l3 3a1 1 0 0 1 0 1.414l-3 3a1 1 0 0 1-1.414 0l-3-3a1 1 0 0 1 0-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ComponentLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="m12 3l3 3l-3 3l-3-3zm0 12l3 3l-3 3l-3-3zm6-6l3 3l-3 3l-3-3zM6 9l3 3l-3 3l-3-3z"/>`), g.Group(children),
	)
}

func ComponentThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m12 3l3 3l-3 3l-3-3zm0 12l3 3l-3 3l-3-3zm6-6l3 3l-3 3l-3-3zM6 9l3 3l-3 3l-3-3z"/>`), g.Group(children),
	)
}

func ConfusedFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2" d="M13 14c-1.48 0-2.773.804-3.465 2"/></g>`), g.Group(children),
	)
}

func ConfusedFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2.5" d="M13 14c-1.48 0-2.773.804-3.465 2"/></g>`), g.Group(children),
	)
}

func ConfusedFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 14c-1.48 0-2.773.804-3.465 2"/></g>`), g.Group(children),
	)
}

func ConfusedFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7-4a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5A1.5 1.5 0 0 0 9.01 8zm6 0a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5a1.5 1.5 0 0 0-1.5-1.5zm-5.966 8.866a1 1 0 0 1-.364-1.367A5 5 0 0 1 13 13a1 1 0 1 1 0 2a3 3 0 0 0-2.6 1.5a1 1 0 0 1-1.366.366" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ConfusedFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/><path stroke-linecap="round" stroke-width="1.5" d="M13 14c-1.48 0-2.773.804-3.465 2"/></g>`), g.Group(children),
	)
}

func ConfusedFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/><path stroke-linecap="round" d="M13 14c-1.48 0-2.773.804-3.465 2"/></g>`), g.Group(children),
	)
}

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M16 3H4v13"/><path d="M8 7h12v12a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2z"/></g>`), g.Group(children),
	)
}

func CopyBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M16 3H4v13"/><path d="M8 7h12v12a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2z"/></g>`), g.Group(children),
	)
}

func CopyDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M8 7h12v12a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 3H4v13"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 7h12v12a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2z"/></g>`), g.Group(children),
	)
}

func CopyFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 3a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H5v12a1 1 0 1 1-2 0zm4 4a1 1 0 0 1 1-1h12a1 1 0 0 1 1 1v12a3 3 0 0 1-3 3h-8a3 3 0 0 1-3-3z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CopyLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16 3H4v13"/><path d="M8 7h12v12a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2z"/></g>`), g.Group(children),
	)
}

func CopyThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M16 3H4v13"/><path d="M8 7h12v12a2 2 0 0 1-2 2h-8a2 2 0 0 1-2-2z"/></g>`), g.Group(children),
	)
}

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9h18M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm4 8h3"/>`), g.Group(children),
	)
}

func CreditCardBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M3 9h18M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm4 8h3"/>`), g.Group(children),
	)
}

func CreditCardDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M5 19h14a2 2 0 0 0 2-2V9H3v8a2 2 0 0 0 2 2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 9h18M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm4 8h3"/></g>`), g.Group(children),
	)
}

func CreditCardFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v3H2zm0 5v7a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-7zm4 3a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CreditCardLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 9h18M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm4 8h3"/>`), g.Group(children),
	)
}

func CreditCardThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 9h18M3 5h18v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2zm4 8h3"/>`), g.Group(children),
	)
}

func Cursor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 21L4 4l17 7l-6.265 2.685a2 2 0 0 0-1.05 1.05z"/>`), g.Group(children),
	)
}

func CursorBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M11 21L4 4l17 7l-6.265 2.685a2 2 0 0 0-1.05 1.05z"/>`), g.Group(children),
	)
}

func CursorDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M11 21L4 4l17 7l-6.265 2.685a2 2 0 0 0-1.05 1.05z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 21L4 4l17 7l-6.265 2.685a2 2 0 0 0-1.05 1.05z"/></g>`), g.Group(children),
	)
}

func CursorFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.38 3.075a1 1 0 0 0-1.305 1.306l7 17a1 1 0 0 0 1.844.013l2.685-6.265a1 1 0 0 1 .525-.525l6.265-2.685a1 1 0 0 0-.013-1.844z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CursorLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 21L4 4l17 7l-6.265 2.685a2 2 0 0 0-1.05 1.05z"/>`), g.Group(children),
	)
}

func CursorThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 21L4 4l17 7l-6.265 2.685a2 2 0 0 0-1.05 1.05z"/>`), g.Group(children),
	)
}

func Delivery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 3V2a1 1 0 0 0-1 1zm11 0h1a1 1 0 0 0-1-1zm0 6V8a1 1 0 0 0-1 1zM2 4h11V2H2zm10-1v16h2V3zM3 17V3H1v14zm10-7h5V8h-5zm8 3v4h2v-4zm-7 6V9h-2v10zm4.707.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zM6.707 19.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zm13.414 0c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM19 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 21 19zm-3-1h-3v2h3zm1.293 1.707A1 1 0 0 1 17 19h-2c0 .766.293 1.536.879 2.121zM17 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 15 19zm-11.707.707A1 1 0 0 1 5 19H3c0 .766.293 1.536.879 2.121zM5 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 3 19zm8-1H8v2h5zm-6.293.293c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM7 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 9 19zm14-2a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-3-7a3 3 0 0 1 3 3h2a5 5 0 0 0-5-5zM1 17a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1z"/>`), g.Group(children),
	)
}

func DeliveryBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 3V1.75C1.31 1.75.75 2.31.75 3zm11 0h1.25c0-.69-.56-1.25-1.25-1.25zm0 6V7.75A1.25 1.25 0 0 0 11.75 9zM2 4.25h11v-2.5H2zM11.75 3v16h2.5V3zm-8.5 14V3H.75v14zM13 10.25h5v-2.5h-5zM20.75 13v4h2.5v-4zm-6.5 6V9h-2.5v10zm4.28.53a.75.75 0 0 1-1.06 0l-1.768 1.768a3.25 3.25 0 0 0 4.596 0zm-1.06-1.06a.75.75 0 0 1 1.06 0l1.768-1.768a3.25 3.25 0 0 0-4.596 0zM6.53 19.53a.75.75 0 0 1-1.06 0l-1.768 1.768a3.25 3.25 0 0 0 4.596 0zm-1.06-1.06a.75.75 0 0 1 1.06 0l1.768-1.768a3.25 3.25 0 0 0-4.596 0zm13.06 0a.74.74 0 0 1 .22.53h2.5a3.24 3.24 0 0 0-.952-2.298zm.22.53a.74.74 0 0 1-.22.53l1.768 1.768A3.24 3.24 0 0 0 21.25 19zM16 17.75h-3v2.5h3zm1.47 1.78a.74.74 0 0 1-.22-.53h-2.5c0 .83.318 1.664.952 2.298zm-.22-.53a.74.74 0 0 1 .22-.53l-1.768-1.768A3.24 3.24 0 0 0 14.75 19zm-11.78.53a.74.74 0 0 1-.22-.53h-2.5c0 .83.318 1.664.952 2.298zM5.25 19a.74.74 0 0 1 .22-.53l-1.768-1.768A3.24 3.24 0 0 0 2.75 19zM13 17.75H8v2.5h5zm-6.47.72a.74.74 0 0 1 .22.53h2.5c0-.83-.318-1.664-.952-2.298zm.22.53a.74.74 0 0 1-.22.53l1.768 1.768A3.24 3.24 0 0 0 9.25 19zm14-2a.75.75 0 0 1-.75.75v2.5A3.25 3.25 0 0 0 23.25 17zM18 10.25A2.75 2.75 0 0 1 20.75 13h2.5c0-2.9-2.35-5.25-5.25-5.25zM.75 17A3.25 3.25 0 0 0 4 20.25v-2.5a.75.75 0 0 1-.75-.75z"/>`), g.Group(children),
	)
}

func DeliveryDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M13 3H2v14a2 2 0 0 0 2 2a2 2 0 1 1 4 0h5z" opacity=".16"/><path d="M2 3V2a1 1 0 0 0-1 1zm11 0h1a1 1 0 0 0-1-1zm0 6V8a1 1 0 0 0-1 1zM2 4h11V2H2zm10-1v16h2V3zM3 17V3H1v14zm10-7h5V8h-5zm8 3v4h2v-4zm-7 6V9h-2v10zm4.707.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zM6.707 19.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zm13.414 0c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM19 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 21 19zm-3-1h-3v2h3zm1.293 1.707A1 1 0 0 1 17 19h-2c0 .766.293 1.536.879 2.121zM17 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 15 19zm-11.707.707A1 1 0 0 1 5 19H3c0 .766.293 1.536.879 2.121zM5 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 3 19zm8-1H8v2h5zm-6.293.293c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM7 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 9 19zm14-2a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-3-7a3 3 0 0 1 3 3h2a5 5 0 0 0-5-5zM1 17a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1z"/></g>`), g.Group(children),
	)
}

func DeliveryFast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 3V2a1 1 0 0 0-1 1zm11 0h1a1 1 0 0 0-1-1zm0 6V8a1 1 0 0 0-1 1zM2 4h11V2H2zm10-1v16h2V3zM3 17V3H1v14zm10-7h5V8h-5zm8 3v4h2v-4zm-7 6V9h-2v10zm4.707.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zM6.707 19.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zm13.414 0c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM19 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 21 19zm-3-1h-3v2h3zm1.293 1.707A1 1 0 0 1 17 19h-2c0 .766.293 1.536.879 2.121zM17 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 15 19zm-11.707.707A1 1 0 0 1 5 19H3c0 .766.293 1.536.879 2.121zM5 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 3 19zm8-1H8v2h5zm-6.293.293c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM7 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 9 19zm14-2a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-3-7a3 3 0 0 1 3 3h2a5 5 0 0 0-5-5zM1 17a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 8h3m-3 4h5"/></g>`), g.Group(children),
	)
}

func DeliveryFastBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 3V1.75C1.31 1.75.75 2.31.75 3zm11 0h1.25c0-.69-.56-1.25-1.25-1.25zm0 6V7.75A1.25 1.25 0 0 0 11.75 9zM2 4.25h11v-2.5H2zM11.75 3v16h2.5V3zm-8.5 14V3H.75v14zM13 10.25h5v-2.5h-5zM20.75 13v4h2.5v-4zm-6.5 6V9h-2.5v10zm4.28.53a.75.75 0 0 1-1.06 0l-1.768 1.768a3.25 3.25 0 0 0 4.596 0zm-1.06-1.06a.75.75 0 0 1 1.06 0l1.768-1.768a3.25 3.25 0 0 0-4.596 0zM6.53 19.53a.75.75 0 0 1-1.06 0l-1.768 1.768a3.25 3.25 0 0 0 4.596 0zm-1.06-1.06a.75.75 0 0 1 1.06 0l1.768-1.768a3.25 3.25 0 0 0-4.596 0zm13.06 0a.74.74 0 0 1 .22.53h2.5a3.24 3.24 0 0 0-.952-2.298zm.22.53a.74.74 0 0 1-.22.53l1.768 1.768A3.24 3.24 0 0 0 21.25 19zM16 17.75h-3v2.5h3zm1.47 1.78a.74.74 0 0 1-.22-.53h-2.5c0 .83.318 1.664.952 2.298zm-.22-.53a.74.74 0 0 1 .22-.53l-1.768-1.768A3.24 3.24 0 0 0 14.75 19zm-11.78.53a.74.74 0 0 1-.22-.53h-2.5c0 .83.318 1.664.952 2.298zM5.25 19a.74.74 0 0 1 .22-.53l-1.768-1.768A3.24 3.24 0 0 0 2.75 19zM13 17.75H8v2.5h5zm-6.47.72a.74.74 0 0 1 .22.53h2.5c0-.83-.318-1.664-.952-2.298zm.22.53a.74.74 0 0 1-.22.53l1.768 1.768A3.24 3.24 0 0 0 9.25 19zm14-2a.75.75 0 0 1-.75.75v2.5A3.25 3.25 0 0 0 23.25 17zM18 10.25A2.75 2.75 0 0 1 20.75 13h2.5c0-2.9-2.35-5.25-5.25-5.25zM.75 17A3.25 3.25 0 0 0 4 20.25v-2.5a.75.75 0 0 1-.75-.75z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M2 8h3m-3 4h5"/></g>`), g.Group(children),
	)
}

func DeliveryFastDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M13 3H2v14a2 2 0 0 0 2 2a2 2 0 1 1 4 0h5z" opacity=".16"/><path fill="currentColor" d="M2 3V2a1 1 0 0 0-1 1zm11 0h1a1 1 0 0 0-1-1zm0 6V8a1 1 0 0 0-1 1zM2 4h11V2H2zm10-1v16h2V3zM3 17V3H1v14zm10-7h5V8h-5zm8 3v4h2v-4zm-7 6V9h-2v10zm4.707.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zM6.707 19.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zm13.414 0c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM19 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 21 19zm-3-1h-3v2h3zm1.293 1.707A1 1 0 0 1 17 19h-2c0 .766.293 1.536.879 2.121zM17 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 15 19zm-11.707.707A1 1 0 0 1 5 19H3c0 .766.293 1.536.879 2.121zM5 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 3 19zm8-1H8v2h5zm-6.293.293c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM7 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 9 19zm14-2a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-3-7a3 3 0 0 1 3 3h2a5 5 0 0 0-5-5zM1 17a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 8h3m-3 4h5"/></g>`), g.Group(children),
	)
}

func DeliveryFastFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1 3a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v5h4a5 5 0 0 1 5 5v4a3 3 0 0 1-2.129 2.872a3 3 0 0 1-5.7.128H8.83a3 3 0 0 1-5.7-.128A3 3 0 0 1 1 17v-4h6a1 1 0 1 0 0-2H1V9h4a1 1 0 0 0 0-2H1zm13 15h1.171a3 3 0 0 1 5.536-.293A1 1 0 0 0 21 17v-4a3 3 0 0 0-3-3h-4zm-7 1a1 1 0 1 0-2 0a1 1 0 0 0 2 0m10.293-.707A1 1 0 0 0 17 19a1 1 0 1 0 .293-.707" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DeliveryFastLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 3v-.75a.75.75 0 0 0-.75.75zm11 0h.75a.75.75 0 0 0-.75-.75zm0 6v-.75a.75.75 0 0 0-.75.75zM2 3.75h11v-1.5H2zM12.25 3v16h1.5V3zm-9.5 14V3h-1.5v14zM13 9.75h5v-1.5h-5zM21.25 13v4h1.5v-4zm-7.5 6V9h-1.5v10zm5.134.884a1.25 1.25 0 0 1-1.768 0l-1.06 1.06a2.75 2.75 0 0 0 3.889 0zm-1.768-1.768a1.25 1.25 0 0 1 1.768 0l1.06-1.06a2.75 2.75 0 0 0-3.889 0zM6.884 19.884a1.25 1.25 0 0 1-1.768 0l-1.06 1.06a2.75 2.75 0 0 0 3.889 0zm-1.768-1.768a1.25 1.25 0 0 1 1.768 0l1.06-1.06a2.75 2.75 0 0 0-3.889 0zm13.768 0c.244.244.366.563.366.884h1.5c0-.703-.269-1.408-.805-1.945zm.366.884c0 .321-.122.64-.366.884l1.06 1.06A2.74 2.74 0 0 0 20.75 19zM16 18.25h-3v1.5h3zm1.116 1.634A1.24 1.24 0 0 1 16.75 19h-1.5c0 .703.269 1.408.805 1.945zM16.75 19c0-.321.122-.64.366-.884l-1.06-1.06A2.74 2.74 0 0 0 15.25 19zm-11.634.884A1.24 1.24 0 0 1 4.75 19h-1.5c0 .703.269 1.408.805 1.945zM4.75 19c0-.321.122-.64.366-.884l-1.06-1.06A2.74 2.74 0 0 0 3.25 19zm8.25-.75H8v1.5h5zm-6.116-.134c.244.244.366.563.366.884h1.5c0-.703-.269-1.408-.805-1.945zM7.25 19c0 .321-.122.64-.366.884l1.06 1.06A2.74 2.74 0 0 0 8.75 19zm14-2c0 .69-.56 1.25-1.25 1.25v1.5A2.75 2.75 0 0 0 22.75 17zM18 9.75A3.25 3.25 0 0 1 21.25 13h1.5A4.75 4.75 0 0 0 18 8.25zM1.25 17A2.75 2.75 0 0 0 4 19.75v-1.5c-.69 0-1.25-.56-1.25-1.25z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2 8h3m-3 4h5"/></g>`), g.Group(children),
	)
}

func DeliveryFastThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 3v-.5a.5.5 0 0 0-.5.5zm11 0h.5a.5.5 0 0 0-.5-.5zm0 6v-.5a.5.5 0 0 0-.5.5zM2 3.5h11v-1H2zM12.5 3v16h1V3zm-10 14V3h-1v14zM13 9.5h5v-1h-5zm8.5 3.5v4h1v-4zm-8 6V9h-1v10zm5.56 1.06a1.5 1.5 0 0 1-2.12 0l-.708.708a2.5 2.5 0 0 0 3.536 0zm-2.12-2.12a1.5 1.5 0 0 1 2.12 0l.708-.708a2.5 2.5 0 0 0-3.536 0zm-9.88 2.12a1.5 1.5 0 0 1-2.12 0l-.708.708a2.5 2.5 0 0 0 3.536 0zm-2.12-2.12a1.5 1.5 0 0 1 2.12 0l.708-.708a2.5 2.5 0 0 0-3.536 0zm14.12 0c.294.292.44.675.44 1.06h1c0-.639-.244-1.28-.732-1.768zM19.5 19c0 .385-.146.768-.44 1.06l.708.708A2.5 2.5 0 0 0 20.5 19zm-3.5-.5h-3v1h3zm.94 1.56A1.5 1.5 0 0 1 16.5 19h-1c0 .639.244 1.28.732 1.768zM16.5 19c0-.385.146-.768.44-1.06l-.708-.708A2.5 2.5 0 0 0 15.5 19zM4.94 20.06A1.5 1.5 0 0 1 4.5 19h-1c0 .639.244 1.28.732 1.768zM4.5 19c0-.385.146-.768.44-1.06l-.708-.708A2.5 2.5 0 0 0 3.5 19zm8.5-.5H8v1h5zm-5.94-.56c.294.292.44.675.44 1.06h1c0-.639-.244-1.28-.732-1.768zM7.5 19c0 .385-.146.768-.44 1.06l.708.708A2.5 2.5 0 0 0 8.5 19zm14-2a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5zM18 9.5a3.5 3.5 0 0 1 3.5 3.5h1A4.5 4.5 0 0 0 18 8.5zM1.5 17A2.5 2.5 0 0 0 4 19.5v-1A1.5 1.5 0 0 1 2.5 17z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2 8h3m-3 4h5"/></g>`), g.Group(children),
	)
}

func DeliveryFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 2a1 1 0 0 0-1 1v14c0 1.354.897 2.498 2.129 2.872a3 3 0 0 0 5.7.128h6.342a3 3 0 0 0 5.7-.128A3 3 0 0 0 23 17v-4a5 5 0 0 0-5-5h-4V3a1 1 0 0 0-1-1zm13.171 16H14v-8h4a3 3 0 0 1 3 3v4a1 1 0 0 1-.293.707a3 3 0 0 0-5.536.293m-9.878.293a1 1 0 1 1 1.414 1.414a1 1 0 0 1-1.414-1.414M17 19a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DeliveryFree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 7V6a1 1 0 0 0-1 1zm11 0h1a1 1 0 0 0-1-1zm0 2V8a1 1 0 0 0-1 1zM2 8h11V6H2zm10-1v12h2V7zM3 17V7H1v10zm10-7h5V8h-5zm8 3v4h2v-4zm-7 6V9h-2v10zm4.707.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zM6.707 19.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zm13.414 0c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM19 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 21 19zm-3-1h-3v2h3zm1.293 1.707A1 1 0 0 1 17 19h-2c0 .766.293 1.536.879 2.121zM17 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 15 19zm-11.707.707A1 1 0 0 1 5 19H3c0 .766.293 1.536.879 2.121zM5 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 3 19zm8-1H8v2h5zm-6.293.293c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM7 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 9 19zm14-2a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-3-7a3 3 0 0 1 3 3h2a5 5 0 0 0-5-5zM1 17a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3.5 4a1 1 0 0 1 1-1a3 3 0 0 1 3 3v1h-1a3 3 0 0 1-3-3Zm8 0a1 1 0 0 0-1-1a3 3 0 0 0-3 3v1h1a3 3 0 0 0 3-3Z"/></g>`), g.Group(children),
	)
}

func DeliveryFreeBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 7V5.75C1.31 5.75.75 6.31.75 7zm11 0h1.25c0-.69-.56-1.25-1.25-1.25zm0 2V7.75A1.25 1.25 0 0 0 11.75 9zM2 8.25h11v-2.5H2zM11.75 7v12h2.5V7zm-8.5 10V7H.75v10zM13 10.25h5v-2.5h-5zM20.75 13v4h2.5v-4zm-6.5 6V9h-2.5v10zm4.28.53a.75.75 0 0 1-1.06 0l-1.768 1.768a3.25 3.25 0 0 0 4.596 0zm-1.06-1.06a.75.75 0 0 1 1.06 0l1.768-1.768a3.25 3.25 0 0 0-4.596 0zM6.53 19.53a.75.75 0 0 1-1.06 0l-1.768 1.768a3.25 3.25 0 0 0 4.596 0zm-1.06-1.06a.75.75 0 0 1 1.06 0l1.768-1.768a3.25 3.25 0 0 0-4.596 0zm13.06 0a.74.74 0 0 1 .22.53h2.5a3.24 3.24 0 0 0-.952-2.298zm.22.53a.74.74 0 0 1-.22.53l1.768 1.768A3.24 3.24 0 0 0 21.25 19zM16 17.75h-3v2.5h3zm1.47 1.78a.74.74 0 0 1-.22-.53h-2.5c0 .83.318 1.664.952 2.298zm-.22-.53a.74.74 0 0 1 .22-.53l-1.768-1.768A3.24 3.24 0 0 0 14.75 19zm-11.78.53a.74.74 0 0 1-.22-.53h-2.5c0 .83.318 1.664.952 2.298zM5.25 19a.74.74 0 0 1 .22-.53l-1.768-1.768A3.24 3.24 0 0 0 2.75 19zM13 17.75H8v2.5h5zm-6.47.72a.74.74 0 0 1 .22.53h2.5c0-.83-.318-1.664-.952-2.298zm.22.53a.74.74 0 0 1-.22.53l1.768 1.768A3.24 3.24 0 0 0 9.25 19zm14-2a.75.75 0 0 1-.75.75v2.5A3.25 3.25 0 0 0 23.25 17zM18 10.25A2.75 2.75 0 0 1 20.75 13h2.5c0-2.9-2.35-5.25-5.25-5.25zM.75 17A3.25 3.25 0 0 0 4 20.25v-2.5a.75.75 0 0 1-.75-.75z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M3.5 4a1 1 0 0 1 1-1a3 3 0 0 1 3 3v1h-1a3 3 0 0 1-3-3Zm8 0a1 1 0 0 0-1-1a3 3 0 0 0-3 3v1h1a3 3 0 0 0 3-3Z"/></g>`), g.Group(children),
	)
}

func DeliveryFreeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 7V6a1 1 0 0 0-1 1zm11 0h1a1 1 0 0 0-1-1zm0 2V8a1 1 0 0 0-1 1zM2 8h11V6H2zm10-1v12h2V7zM3 17V7H1v10zm10-7h5V8h-5zm8 3v4h2v-4zm-7 6V9h-2v10zm4.707.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zM6.707 19.707a1 1 0 0 1-1.414 0l-1.414 1.414a3 3 0 0 0 4.242 0zm-1.414-1.414a1 1 0 0 1 1.414 0l1.414-1.414a3 3 0 0 0-4.242 0zm13.414 0c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM19 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 21 19zm-3-1h-3v2h3zm1.293 1.707A1 1 0 0 1 17 19h-2c0 .766.293 1.536.879 2.121zM17 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 15 19zm-11.707.707A1 1 0 0 1 5 19H3c0 .766.293 1.536.879 2.121zM5 19a1 1 0 0 1 .293-.707l-1.414-1.414A3 3 0 0 0 3 19zm8-1H8v2h5zm-6.293.293c.196.195.293.45.293.707h2c0-.766-.293-1.536-.879-2.121zM7 19a1 1 0 0 1-.293.707l1.414 1.414A3 3 0 0 0 9 19zm14-2a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-3-7a3 3 0 0 1 3 3h2a5 5 0 0 0-5-5zM1 17a3 3 0 0 0 3 3v-2a1 1 0 0 1-1-1z"/><path fill="currentColor" d="M13 7H2v10a2 2 0 0 0 2 2a2 2 0 1 1 4 0h5z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3.5 4a1 1 0 0 1 1-1a3 3 0 0 1 3 3v1h-1a3 3 0 0 1-3-3Zm8 0a1 1 0 0 0-1-1a3 3 0 0 0-3 3v1h1a3 3 0 0 0 3-3Z"/></g>`), g.Group(children),
	)
}

func DeliveryFreeFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 4a2 2 0 0 1 2-2a4 4 0 0 1 3 1.354A4 4 0 0 1 10.5 2a2 2 0 0 1 2 2c0 .729-.195 1.412-.535 2H13a1 1 0 0 1 1 1v1h4a5 5 0 0 1 5 5v4a3 3 0 0 1-2.129 2.872a3 3 0 0 1-5.7.128H8.83a3 3 0 0 1-5.7-.128A3 3 0 0 1 1 17V7a1 1 0 0 1 1-1h1.035A4 4 0 0 1 2.5 4m2 0a2 2 0 0 1 2 2a2 2 0 0 1-2-2m4 2a2 2 0 0 0 2-2a2 2 0 0 0-2 2m6.671 12H14v-8h4a3 3 0 0 1 3 3v4a1 1 0 0 1-.293.707a3 3 0 0 0-5.536.293m-9.878.293a1 1 0 1 1 1.414 1.414a1 1 0 0 1-1.414-1.414M17 19a1 1 0 1 1 2 0a1 1 0 0 1-2 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DeliveryFreeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 7v-.75a.75.75 0 0 0-.75.75zm11 0h.75a.75.75 0 0 0-.75-.75zm0 2v-.75a.75.75 0 0 0-.75.75zM2 7.75h11v-1.5H2zM12.25 7v12h1.5V7zm-9.5 10V7h-1.5v10zM13 9.75h5v-1.5h-5zM21.25 13v4h1.5v-4zm-7.5 6V9h-1.5v10zm5.134.884a1.25 1.25 0 0 1-1.768 0l-1.06 1.06a2.75 2.75 0 0 0 3.889 0zm-1.768-1.768a1.25 1.25 0 0 1 1.768 0l1.06-1.06a2.75 2.75 0 0 0-3.889 0zM6.884 19.884a1.25 1.25 0 0 1-1.768 0l-1.06 1.06a2.75 2.75 0 0 0 3.889 0zm-1.768-1.768a1.25 1.25 0 0 1 1.768 0l1.06-1.06a2.75 2.75 0 0 0-3.889 0zm13.768 0c.244.244.366.563.366.884h1.5c0-.703-.269-1.408-.805-1.945zm.366.884c0 .321-.122.64-.366.884l1.06 1.06A2.74 2.74 0 0 0 20.75 19zM16 18.25h-3v1.5h3zm1.116 1.634A1.24 1.24 0 0 1 16.75 19h-1.5c0 .703.269 1.408.805 1.945zM16.75 19c0-.321.122-.64.366-.884l-1.06-1.06A2.74 2.74 0 0 0 15.25 19zm-11.634.884A1.24 1.24 0 0 1 4.75 19h-1.5c0 .703.269 1.408.805 1.945zM4.75 19c0-.321.122-.64.366-.884l-1.06-1.06A2.74 2.74 0 0 0 3.25 19zm8.25-.75H8v1.5h5zm-6.116-.134c.244.244.366.563.366.884h1.5c0-.703-.269-1.408-.805-1.945zM7.25 19c0 .321-.122.64-.366.884l1.06 1.06A2.74 2.74 0 0 0 8.75 19zm14-2c0 .69-.56 1.25-1.25 1.25v1.5A2.75 2.75 0 0 0 22.75 17zM18 9.75A3.25 3.25 0 0 1 21.25 13h1.5A4.75 4.75 0 0 0 18 8.25zM1.25 17A2.75 2.75 0 0 0 4 19.75v-1.5c-.69 0-1.25-.56-1.25-1.25z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M3.5 4a1 1 0 0 1 1-1a3 3 0 0 1 3 3v1h-1a3 3 0 0 1-3-3Zm8 0a1 1 0 0 0-1-1a3 3 0 0 0-3 3v1h1a3 3 0 0 0 3-3Z"/></g>`), g.Group(children),
	)
}

func DeliveryFreeThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2 7v-.5a.5.5 0 0 0-.5.5zm11 0h.5a.5.5 0 0 0-.5-.5zm0 2v-.5a.5.5 0 0 0-.5.5zM2 7.5h11v-1H2zM12.5 7v12h1V7zm-10 10V7h-1v10zM13 9.5h5v-1h-5zm8.5 3.5v4h1v-4zm-8 6V9h-1v10zm5.56 1.06a1.5 1.5 0 0 1-2.12 0l-.708.708a2.5 2.5 0 0 0 3.536 0zm-2.12-2.12a1.5 1.5 0 0 1 2.12 0l.708-.708a2.5 2.5 0 0 0-3.536 0zm-9.88 2.12a1.5 1.5 0 0 1-2.12 0l-.708.708a2.5 2.5 0 0 0 3.536 0zm-2.12-2.12a1.5 1.5 0 0 1 2.12 0l.708-.708a2.5 2.5 0 0 0-3.536 0zm14.12 0c.294.292.44.675.44 1.06h1c0-.639-.244-1.28-.732-1.768zM19.5 19c0 .385-.146.768-.44 1.06l.708.708A2.5 2.5 0 0 0 20.5 19zm-3.5-.5h-3v1h3zm.94 1.56A1.5 1.5 0 0 1 16.5 19h-1c0 .639.244 1.28.732 1.768zM16.5 19c0-.385.146-.768.44-1.06l-.708-.708A2.5 2.5 0 0 0 15.5 19zM4.94 20.06A1.5 1.5 0 0 1 4.5 19h-1c0 .639.244 1.28.732 1.768zM4.5 19c0-.385.146-.768.44-1.06l-.708-.708A2.5 2.5 0 0 0 3.5 19zm8.5-.5H8v1h5zm-5.94-.56c.294.292.44.675.44 1.06h1c0-.639-.244-1.28-.732-1.768zM7.5 19c0 .385-.146.768-.44 1.06l.708.708A2.5 2.5 0 0 0 8.5 19zm14-2a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5zM18 9.5a3.5 3.5 0 0 1 3.5 3.5h1A4.5 4.5 0 0 0 18 8.5zM1.5 17A2.5 2.5 0 0 0 4 19.5v-1A1.5 1.5 0 0 1 2.5 17z"/><path stroke="currentColor" stroke-linejoin="round" d="M3.5 4a1 1 0 0 1 1-1a3 3 0 0 1 3 3v1h-1a3 3 0 0 1-3-3Zm8 0a1 1 0 0 0-1-1a3 3 0 0 0-3 3v1h1a3 3 0 0 0 3-3Z"/></g>`), g.Group(children),
	)
}

func DeliveryLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 3v-.75a.75.75 0 0 0-.75.75zm11 0h.75a.75.75 0 0 0-.75-.75zm0 6v-.75a.75.75 0 0 0-.75.75zM2 3.75h11v-1.5H2zM12.25 3v16h1.5V3zm-9.5 14V3h-1.5v14zM13 9.75h5v-1.5h-5zM21.25 13v4h1.5v-4zm-7.5 6V9h-1.5v10zm5.134.884a1.25 1.25 0 0 1-1.768 0l-1.06 1.06a2.75 2.75 0 0 0 3.889 0zm-1.768-1.768a1.25 1.25 0 0 1 1.768 0l1.06-1.06a2.75 2.75 0 0 0-3.889 0zM6.884 19.884a1.25 1.25 0 0 1-1.768 0l-1.06 1.06a2.75 2.75 0 0 0 3.889 0zm-1.768-1.768a1.25 1.25 0 0 1 1.768 0l1.06-1.06a2.75 2.75 0 0 0-3.889 0zm13.768 0c.244.244.366.563.366.884h1.5c0-.703-.269-1.408-.805-1.945zm.366.884c0 .321-.122.64-.366.884l1.06 1.06A2.74 2.74 0 0 0 20.75 19zM16 18.25h-3v1.5h3zm1.116 1.634A1.24 1.24 0 0 1 16.75 19h-1.5c0 .703.269 1.408.805 1.945zM16.75 19c0-.321.122-.64.366-.884l-1.06-1.06A2.74 2.74 0 0 0 15.25 19zm-11.634.884A1.24 1.24 0 0 1 4.75 19h-1.5c0 .703.269 1.408.805 1.945zM4.75 19c0-.321.122-.64.366-.884l-1.06-1.06A2.74 2.74 0 0 0 3.25 19zm8.25-.75H8v1.5h5zm-6.116-.134c.244.244.366.563.366.884h1.5c0-.703-.269-1.408-.805-1.945zM7.25 19c0 .321-.122.64-.366.884l1.06 1.06A2.74 2.74 0 0 0 8.75 19zm14-2c0 .69-.56 1.25-1.25 1.25v1.5A2.75 2.75 0 0 0 22.75 17zM18 9.75A3.25 3.25 0 0 1 21.25 13h1.5A4.75 4.75 0 0 0 18 8.25zM1.25 17A2.75 2.75 0 0 0 4 19.75v-1.5c-.69 0-1.25-.56-1.25-1.25z"/>`), g.Group(children),
	)
}

func DeliveryThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 3v-.5a.5.5 0 0 0-.5.5zm11 0h.5a.5.5 0 0 0-.5-.5zm0 6v-.5a.5.5 0 0 0-.5.5zM2 3.5h11v-1H2zM12.5 3v16h1V3zm-10 14V3h-1v14zM13 9.5h5v-1h-5zm8.5 3.5v4h1v-4zm-8 6V9h-1v10zm5.56 1.06a1.5 1.5 0 0 1-2.12 0l-.708.708a2.5 2.5 0 0 0 3.536 0zm-2.12-2.12a1.5 1.5 0 0 1 2.12 0l.708-.708a2.5 2.5 0 0 0-3.536 0zm-9.88 2.12a1.5 1.5 0 0 1-2.12 0l-.708.708a2.5 2.5 0 0 0 3.536 0zm-2.12-2.12a1.5 1.5 0 0 1 2.12 0l.708-.708a2.5 2.5 0 0 0-3.536 0zm14.12 0c.294.292.44.675.44 1.06h1c0-.639-.244-1.28-.732-1.768zM19.5 19c0 .385-.146.768-.44 1.06l.708.708A2.5 2.5 0 0 0 20.5 19zm-3.5-.5h-3v1h3zm.94 1.56A1.5 1.5 0 0 1 16.5 19h-1c0 .639.244 1.28.732 1.768zM16.5 19c0-.385.146-.768.44-1.06l-.708-.708A2.5 2.5 0 0 0 15.5 19zM4.94 20.06A1.5 1.5 0 0 1 4.5 19h-1c0 .639.244 1.28.732 1.768zM4.5 19c0-.385.146-.768.44-1.06l-.708-.708A2.5 2.5 0 0 0 3.5 19zm8.5-.5H8v1h5zm-5.94-.56c.294.292.44.675.44 1.06h1c0-.639-.244-1.28-.732-1.768zM7.5 19c0 .385-.146.768-.44 1.06l.708.708A2.5 2.5 0 0 0 8.5 19zm14-2a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5zM18 9.5a3.5 3.5 0 0 1 3.5 3.5h1A4.5 4.5 0 0 0 18 8.5zM1.5 17A2.5 2.5 0 0 0 4 19.5v-1A1.5 1.5 0 0 1 2.5 17z"/>`), g.Group(children),
	)
}

func DisappointedFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M10 15.535A4 4 0 0 1 12 15a4 4 0 0 1 2 .535m3-5.035l-3-1m-4 0l-3 1"/></g>`), g.Group(children),
	)
}

func DisappointedFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M10 15.535A4 4 0 0 1 12 15a4 4 0 0 1 2 .535m3-5.035l-3-1m-4 0l-3 1"/></g>`), g.Group(children),
	)
}

func DisappointedFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 15.535A4 4 0 0 1 12 15a4 4 0 0 1 2 .535m3-5.035l-3-1m-4 0l-3 1"/></g>`), g.Group(children),
	)
}

func DisappointedFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m8.949-2.816a1 1 0 0 1-.633 1.265l-3 1a1 1 0 0 1-.632-1.898l3-1a1 1 0 0 1 1.265.633m3.367-.633a1 1 0 0 0-.632 1.898l3 1a1 1 0 0 0 .632-1.898zM9.5 14.67a1 1 0 1 0 1.002 1.73c.44-.254.95-.4 1.499-.4c.548 0 1.059.146 1.5.4a1 1 0 0 0 1-1.73A5 5 0 0 0 12 14c-.91 0-1.764.243-2.5.67" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DisappointedFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M10 15.535A4 4 0 0 1 12 15a4 4 0 0 1 2 .535m3-5.035l-3-1m-4 0l-3 1"/></g>`), g.Group(children),
	)
}

func DisappointedFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M10 15.535A4 4 0 0 1 12 15a4 4 0 0 1 2 .535m3-5.035l-3-1m-4 0l-3 1"/></g>`), g.Group(children),
	)
}

func Discount(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-width="2" d="M10.51 3.665a2 2 0 0 1 2.98 0l.7.782a2 2 0 0 0 1.601.663l1.05-.058a2 2 0 0 1 2.107 2.108l-.058 1.049a2 2 0 0 0 .663 1.6l.782.7a2 2 0 0 1 0 2.981l-.782.7a2 2 0 0 0-.663 1.601l.058 1.05a2 2 0 0 1-2.108 2.107l-1.049-.058a2 2 0 0 0-1.6.663l-.7.782a2 2 0 0 1-2.981 0l-.7-.782a2 2 0 0 0-1.601-.663l-1.05.058a2 2 0 0 1-2.107-2.108l.058-1.049a2 2 0 0 0-.663-1.6l-.782-.7a2 2 0 0 1 0-2.981l.782-.7a2 2 0 0 0 .663-1.601l-.058-1.05A2 2 0 0 1 7.16 5.053l1.049.058a2 2 0 0 0 1.6-.663z"/><path stroke-linejoin="round" stroke-width="3" d="M9.5 9.5h.01v.01H9.5zm5 5h.01v.01h-.01z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 9l-6 6"/></g>`), g.Group(children),
	)
}

func DiscountBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-width="2.5" d="M10.51 3.665a2 2 0 0 1 2.98 0l.7.782a2 2 0 0 0 1.601.663l1.05-.058a2 2 0 0 1 2.107 2.108l-.058 1.049a2 2 0 0 0 .663 1.6l.782.7a2 2 0 0 1 0 2.981l-.782.7a2 2 0 0 0-.663 1.601l.058 1.05a2 2 0 0 1-2.108 2.107l-1.049-.058a2 2 0 0 0-1.6.663l-.7.782a2 2 0 0 1-2.981 0l-.7-.782a2 2 0 0 0-1.601-.663l-1.05.058a2 2 0 0 1-2.107-2.108l.058-1.049a2 2 0 0 0-.663-1.6l-.782-.7a2 2 0 0 1 0-2.981l.782-.7a2 2 0 0 0 .663-1.601l-.058-1.05A2 2 0 0 1 7.16 5.053l1.049.058a2 2 0 0 0 1.6-.663z"/><path stroke-linejoin="round" stroke-width="3.75" d="M9.5 9.5h.01v.01H9.5zm5 5h.01v.01h-.01z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m15 9l-6 6"/></g>`), g.Group(children),
	)
}

func DiscountDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M10.51 3.665a2 2 0 0 1 2.98 0l.7.782a2 2 0 0 0 1.601.663l1.05-.058a2 2 0 0 1 2.107 2.108l-.058 1.049a2 2 0 0 0 .663 1.6l.782.7a2 2 0 0 1 0 2.981l-.782.7a2 2 0 0 0-.663 1.601l.058 1.05a2 2 0 0 1-2.108 2.107l-1.049-.058a2 2 0 0 0-1.6.663l-.7.782a2 2 0 0 1-2.981 0l-.7-.782a2 2 0 0 0-1.601-.663l-1.05.058a2 2 0 0 1-2.107-2.108l.058-1.049a2 2 0 0 0-.663-1.6l-.782-.7a2 2 0 0 1 0-2.981l.782-.7a2 2 0 0 0 .663-1.601l-.058-1.05A2 2 0 0 1 7.16 5.053l1.049.058a2 2 0 0 0 1.6-.663z" opacity=".16"/><path stroke="currentColor" stroke-width="2" d="M10.51 3.665a2 2 0 0 1 2.98 0l.7.782a2 2 0 0 0 1.601.663l1.05-.058a2 2 0 0 1 2.107 2.108l-.058 1.049a2 2 0 0 0 .663 1.6l.782.7a2 2 0 0 1 0 2.981l-.782.7a2 2 0 0 0-.663 1.601l.058 1.05a2 2 0 0 1-2.108 2.107l-1.049-.058a2 2 0 0 0-1.6.663l-.7.782a2 2 0 0 1-2.981 0l-.7-.782a2 2 0 0 0-1.601-.663l-1.05.058a2 2 0 0 1-2.107-2.108l.058-1.049a2 2 0 0 0-.663-1.6l-.782-.7a2 2 0 0 1 0-2.981l.782-.7a2 2 0 0 0 .663-1.601l-.058-1.05A2 2 0 0 1 7.16 5.053l1.049.058a2 2 0 0 0 1.6-.663z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9.5 9.5h.01v.01H9.5zm5 5h.01v.01h-.01z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 9l-6 6"/></g>`), g.Group(children),
	)
}

func DiscountFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.765 2.998a3 3 0 0 1 4.47 0l.7.782a1 1 0 0 0 .801.332l1.05-.058a3 3 0 0 1 3.16 3.16l-.058 1.05a1 1 0 0 0 .332.8l.783.7a3 3 0 0 1 0 4.471l-.783.7a1 1 0 0 0-.332.801l.058 1.05a3 3 0 0 1-3.16 3.16l-1.05-.058a1 1 0 0 0-.8.332l-.7.783a3 3 0 0 1-4.471 0l-.7-.783a1 1 0 0 0-.801-.332l-1.05.058a3 3 0 0 1-3.16-3.16l.058-1.05a1 1 0 0 0-.332-.8l-.782-.7a3 3 0 0 1 0-4.471l.782-.7a1 1 0 0 0 .332-.801l-.058-1.05a3 3 0 0 1 3.16-3.16l1.05.058a1 1 0 0 0 .8-.332zm5.942 5.295a1 1 0 0 1 0 1.414l-6 6a1 1 0 0 1-1.414-1.414l6-6a1 1 0 0 1 1.414 0M9.5 8A1.5 1.5 0 0 0 8 9.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5A1.5 1.5 0 0 0 9.51 8zm5 5a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DiscountLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-width="1.5" d="M10.51 3.665a2 2 0 0 1 2.98 0l.7.782a2 2 0 0 0 1.601.663l1.05-.058a2 2 0 0 1 2.107 2.108l-.058 1.049a2 2 0 0 0 .663 1.6l.782.7a2 2 0 0 1 0 2.981l-.782.7a2 2 0 0 0-.663 1.601l.058 1.05a2 2 0 0 1-2.108 2.107l-1.049-.058a2 2 0 0 0-1.6.663l-.7.782a2 2 0 0 1-2.981 0l-.7-.782a2 2 0 0 0-1.601-.663l-1.05.058a2 2 0 0 1-2.107-2.108l.058-1.049a2 2 0 0 0-.663-1.6l-.782-.7a2 2 0 0 1 0-2.981l.782-.7a2 2 0 0 0 .663-1.601l-.058-1.05A2 2 0 0 1 7.16 5.053l1.049.058a2 2 0 0 0 1.6-.663z"/><path stroke-linejoin="round" stroke-width="2.25" d="M9.5 9.5h.01v.01H9.5zm5 5h.01v.01h-.01z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 9l-6 6"/></g>`), g.Group(children),
	)
}

func DiscountThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path d="M10.51 3.665a2 2 0 0 1 2.98 0l.7.782a2 2 0 0 0 1.601.663l1.05-.058a2 2 0 0 1 2.107 2.108l-.058 1.049a2 2 0 0 0 .663 1.6l.782.7a2 2 0 0 1 0 2.981l-.782.7a2 2 0 0 0-.663 1.601l.058 1.05a2 2 0 0 1-2.108 2.107l-1.049-.058a2 2 0 0 0-1.6.663l-.7.782a2 2 0 0 1-2.981 0l-.7-.782a2 2 0 0 0-1.601-.663l-1.05.058a2 2 0 0 1-2.107-2.108l.058-1.049a2 2 0 0 0-.663-1.6l-.782-.7a2 2 0 0 1 0-2.981l.782-.7a2 2 0 0 0 .663-1.601l-.058-1.05A2 2 0 0 1 7.16 5.053l1.049.058a2 2 0 0 0 1.6-.663z"/><path stroke-linejoin="round" stroke-width="1.5" d="M9.5 9.5h.01v.01H9.5zm5 5h.01v.01h-.01z"/><path stroke-linecap="round" stroke-linejoin="round" d="m15 9l-6 6"/></g>`), g.Group(children),
	)
}

func Discover(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M11.307 9.739L15 9l-.739 3.693a2 2 0 0 1-1.568 1.569L9 15l.739-3.693a2 2 0 0 1 1.568-1.568"/></g>`), g.Group(children),
	)
}

func DiscoverBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M11.307 9.739L15 9l-.739 3.693a2 2 0 0 1-1.568 1.569L9 15l.739-3.693a2 2 0 0 1 1.568-1.568"/></g>`), g.Group(children),
	)
}

func DiscoverDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m3-12l-3.693.739a2 2 0 0 0-1.568 1.568L9 15l3.693-.739a2 2 0 0 0 1.569-1.568z" clip-rule="evenodd" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.307 9.739L15 9l-.739 3.693a2 2 0 0 1-1.568 1.569L9 15l.739-3.693a2 2 0 0 1 1.568-1.568"/></g>`), g.Group(children),
	)
}

func DiscoverFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m13-3l-3.693.739a2 2 0 0 0-1.568 1.568L9 15l3.693-.739a2 2 0 0 0 1.569-1.568z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DiscoverLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M11.307 9.739L15 9l-.739 3.693a2 2 0 0 1-1.568 1.569L9 15l.739-3.693a2 2 0 0 1 1.568-1.568"/></g>`), g.Group(children),
	)
}

func DiscoverThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M11.307 9.739L15 9l-.739 3.693a2 2 0 0 1-1.568 1.569L9 15l.739-3.693a2 2 0 0 1 1.568-1.568"/></g>`), g.Group(children),
	)
}

func Dislike(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 14l-.986.164A1 1 0 0 1 15 13zM4 14v1a1 1 0 0 1-1-1zm16.522-2.392l.98-.196zM6 3h11.36v2H6zm12.56 12H15v-2h3.56zm-2.573-1.164l.805 4.835L14.82 19l-.806-4.836zM14.82 21h-.214v-2h.214zm-3.543-1.781l-2.515-3.774l1.664-1.11l2.516 3.774zM7.93 15H4v-2h3.93zM3 14V6h2v8zm17.302-8.588l1.2 6l-1.96.392l-1.2-6zM8.762 15.445A1 1 0 0 0 7.93 15v-2a3 3 0 0 1 2.496 1.336zm8.03 3.226A2 2 0 0 1 14.82 21v-2zM18.56 13a1 1 0 0 0 .981-1.196l1.961-.392A3 3 0 0 1 18.561 15zm-1.2-10a3 3 0 0 1 2.942 2.412l-1.96.392A1 1 0 0 0 17.36 5zm-2.754 18a4 4 0 0 1-3.328-1.781l1.664-1.11a2 2 0 0 0 1.664.891zM6 5a1 1 0 0 0-1 1H3a3 3 0 0 1 3-3z"/><path stroke="currentColor" stroke-width="2" d="M8 14V4"/></g>`), g.Group(children),
	)
}

func DislikeBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 14l-1.233.206A1.25 1.25 0 0 1 15 12.75zM4 14v1.25c-.69 0-1.25-.56-1.25-1.25zm16.522-2.392l1.225-.245zM6 2.75h11.36v2.5H6zm12.56 12.5H15v-2.5h3.56zm-2.327-1.456l.806 4.836l-2.466.411l-.806-4.835zM14.82 21.25h-.214v-2.5h.214zm-3.75-1.892l-2.516-3.774l2.08-1.387l2.516 3.774zM7.93 15.25H4v-2.5h3.93zM2.75 14V6h2.5v8zm17.797-8.637l1.2 6l-2.451.49l-1.2-6zM8.554 15.584a.75.75 0 0 0-.624-.334v-2.5a3.25 3.25 0 0 1 2.704 1.447zm8.485 3.046a2.25 2.25 0 0 1-2.22 2.62v-2.5a.25.25 0 0 0-.246.291zm1.521-5.88a.75.75 0 0 0 .736-.897l2.451-.49a3.25 3.25 0 0 1-3.186 3.887zm-1.2-10a3.25 3.25 0 0 1 3.187 2.613l-2.451.49a.75.75 0 0 0-.736-.603zm-2.754 18.5a4.25 4.25 0 0 1-3.536-1.892l2.08-1.387c.324.487.87.779 1.456.779zM6 5.25a.75.75 0 0 0-.75.75h-2.5A3.25 3.25 0 0 1 6 2.75z"/><path stroke="currentColor" stroke-width="2.5" d="M8 14V4"/></g>`), g.Group(children),
	)
}

func DislikeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M17.36 4H8v10c.625 0 1.208.312 1.555.832l2.554 3.832A3 3 0 0 0 14.606 20h.213a1 1 0 0 0 .987-1.164L15 14h3.56a2 2 0 0 0 1.962-2.392l-1.2-6A2 2 0 0 0 17.36 4" opacity=".16"/><path fill="currentColor" d="m15 14l-.986.164A1 1 0 0 1 15 13zM4 14v1a1 1 0 0 1-1-1zm16.522-2.392l.98-.196zM6 3h11.36v2H6zm12.56 12H15v-2h3.56zm-2.574-1.164l.806 4.835L14.82 19l-.805-4.836zM14.82 21h-.213v-2h.213zm-3.542-1.781l-2.515-3.774l1.664-1.11l2.515 3.774zM7.93 15H4v-2h3.93zM3 14V6h2v8zm17.302-8.588l1.2 6l-1.961.392l-1.2-6zM8.762 15.445A1 1 0 0 0 7.93 15v-2a3 3 0 0 1 2.496 1.336zm8.03 3.226A2 2 0 0 1 14.82 21v-2zM18.56 13a1 1 0 0 0 .981-1.196l1.961-.392A3 3 0 0 1 18.56 15zm-1.2-10a3 3 0 0 1 2.942 2.412l-1.961.392a1 1 0 0 0-.98-.804zm-2.754 18a4 4 0 0 1-3.329-1.781l1.665-1.11a2 2 0 0 0 1.664.891zM6 5a1 1 0 0 0-1 1H3a3 3 0 0 1 3-3z"/><path stroke="currentColor" stroke-width="2" d="M8 14V4"/></g>`), g.Group(children),
	)
}

func DislikeFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.277 19.219A4 4 0 0 0 14.606 21h.213a2 2 0 0 0 1.973-2.329L16.18 15h2.38a3 3 0 0 0 2.942-3.588l-1.2-6A3 3 0 0 0 17.36 3H6a3 3 0 0 0-3 3v8a1 1 0 0 0 1 1h3.93a1 1 0 0 1 .832.445zM7 5v8H5V6a1 1 0 0 1 1-1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DislikeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 14l-.74.123a.75.75 0 0 1 .74-.873zM4 14v.75a.75.75 0 0 1-.75-.75zm16.522-2.392l.735-.147zM6 3.25h11.36v1.5H6zm12.56 11.5H15v-1.5h3.56zm-2.82-.873l.806 4.835l-1.48.247l-.806-4.836zm-.92 6.873h-.214v-1.5h.213zm-3.335-1.67L8.97 15.307l1.248-.832l2.515 3.773zM7.93 14.75H4v-1.5h3.93zM3.25 14V6h1.5v8zm16.807-8.54l1.2 6l-1.47.295l-1.2-6zM8.97 15.308a1.25 1.25 0 0 0-1.04-.557v-1.5c.92 0 1.778.46 2.288 1.225zm7.576 3.405a1.75 1.75 0 0 1-1.726 2.038v-1.5a.25.25 0 0 0 .246-.291zm2.014-5.462a1.25 1.25 0 0 0 1.226-1.495l1.471-.294a2.75 2.75 0 0 1-2.697 3.289zm-1.2-10a2.75 2.75 0 0 1 2.697 2.21l-1.47.295A1.25 1.25 0 0 0 17.36 4.75zm-2.754 17.5a3.75 3.75 0 0 1-3.12-1.67l1.247-.832a2.25 2.25 0 0 0 1.873 1.002zM6 4.75c-.69 0-1.25.56-1.25 1.25h-1.5A2.75 2.75 0 0 1 6 3.25z"/><path stroke="currentColor" stroke-width="1.5" d="M8 14V4"/></g>`), g.Group(children),
	)
}

func DislikeThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 14l-.493.082A.5.5 0 0 1 15 13.5zM4 14v.5a.5.5 0 0 1-.5-.5zm16.522-2.392l.49-.098zM6 3.5h11.36v1H6zm12.56 11H15v-1h3.56zm-3.067-.582l.806 4.835l-.986.165l-.806-4.836zM14.82 20.5h-.213v-1h.213zm-3.126-1.558l-2.515-3.774l.832-.555l2.515 3.774zM7.93 14.5H4v-1h3.93zM3.5 14V6h1v8zm16.312-8.49l1.2 6l-.98.196l-1.2-6zM9.178 15.168A1.5 1.5 0 0 0 7.93 14.5v-1a2.5 2.5 0 0 1 2.08 1.113zm7.121 3.585a1.5 1.5 0 0 1-1.48 1.747v-1a.5.5 0 0 0 .494-.582zM18.56 13.5a1.5 1.5 0 0 0 1.471-1.794l.98-.196a2.5 2.5 0 0 1-2.45 2.99zm-1.2-10a2.5 2.5 0 0 1 2.452 2.01l-.98.196A1.5 1.5 0 0 0 17.36 4.5zm-2.754 17a3.5 3.5 0 0 1-2.913-1.558l.832-.555a2.5 2.5 0 0 0 2.08 1.113zM6 4.5A1.5 1.5 0 0 0 4.5 6h-1A2.5 2.5 0 0 1 6 3.5z"/><path stroke="currentColor" d="M8 14V4"/></g>`), g.Group(children),
	)
}

func DoRedo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.707 5.293a1 1 0 1 0-1.414 1.414zM19 9l.707.707a1 1 0 0 0 0-1.414zm-3.707 2.293a1 1 0 0 0 1.414 1.414zM19 18a1 1 0 1 0 0-2zM15.293 6.707l3 3l1.414-1.414l-3-3zm3 1.586l-3 3l1.414 1.414l3-3zM19 8H8v2h11zM3 13a5 5 0 0 0 5 5v-2a3 3 0 0 1-3-3zm5-5a5 5 0 0 0-5 5h2a3 3 0 0 1 3-3zm11 8H8v2h11z"/>`), g.Group(children),
	)
}

func DoRedoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.884 5.116a1.25 1.25 0 0 0-1.768 1.768zM19 9l.884.884a1.25 1.25 0 0 0 0-1.768zm-3.884 2.116a1.25 1.25 0 0 0 1.768 1.768zM19 18.25a1.25 1.25 0 1 0 0-2.5zM15.116 6.884l3 3l1.768-1.768l-3-3zm3 1.232l-3 3l1.768 1.768l3-3zM19 7.75H8v2.5h11zM2.75 13c0 2.9 2.35 5.25 5.25 5.25v-2.5A2.75 2.75 0 0 1 5.25 13zM8 7.75A5.25 5.25 0 0 0 2.75 13h2.5A2.75 2.75 0 0 1 8 10.25zm11 8H8v2.5h11z"/>`), g.Group(children),
	)
}

func DoRedoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.707 5.293a1 1 0 1 0-1.414 1.414zM19 9l.707.707a1 1 0 0 0 0-1.414zm-3.707 2.293a1 1 0 0 0 1.414 1.414zM19 18a1 1 0 1 0 0-2zM15.293 6.707l3 3l1.414-1.414l-3-3zm3 1.586l-3 3l1.414 1.414l3-3zM19 8H8v2h11zM3 13a5 5 0 0 0 5 5v-2a3 3 0 0 1-3-3zm5-5a5 5 0 0 0-5 5h2a3 3 0 0 1 3-3zm11 8H8v2h11z"/>`), g.Group(children),
	)
}

func DoRedoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.617 5.076a1 1 0 0 1 1.09.217l3 3a1 1 0 0 1 0 1.414l-3 3A1 1 0 0 1 15 12v-2H8a3 3 0 1 0 0 6h11a1 1 0 1 1 0 2H8A5 5 0 0 1 8 8h7V6a1 1 0 0 1 .617-.924" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DoRedoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.53 5.47a.75.75 0 1 0-1.06 1.06zM19 9l.53.53a.75.75 0 0 0 0-1.06zm-3.53 2.47a.75.75 0 1 0 1.06 1.06zM19 17.75a.75.75 0 0 0 0-1.5zM15.47 6.53l3 3l1.06-1.06l-3-3zm3 1.94l-3 3l1.06 1.06l3-3zm.53-.22H8v1.5h11zM3.25 13A4.75 4.75 0 0 0 8 17.75v-1.5A3.25 3.25 0 0 1 4.75 13zM8 8.25A4.75 4.75 0 0 0 3.25 13h1.5A3.25 3.25 0 0 1 8 9.75zm11 8H8v1.5h11z"/>`), g.Group(children),
	)
}

func DoRedoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.354 5.646a.5.5 0 0 0-.708.708zM19 9l.354.354a.5.5 0 0 0 0-.708zm-3.354 2.646a.5.5 0 0 0 .708.708zM19 17.5a.5.5 0 0 0 0-1zM15.646 6.354l3 3l.708-.708l-3-3zm3 2.292l-3 3l.708.708l3-3zM19 8.5H8v1h11zM3.5 13A4.5 4.5 0 0 0 8 17.5v-1A3.5 3.5 0 0 1 4.5 13zM8 8.5A4.5 4.5 0 0 0 3.5 13h1A3.5 3.5 0 0 1 8 9.5zm11 8H8v1h11z"/>`), g.Group(children),
	)
}

func DoUndo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.707 6.707a1 1 0 0 0-1.414-1.414zM5 9l-.707-.707a1 1 0 0 0 0 1.414zm2.293 3.707a1 1 0 1 0 1.414-1.414zM5 16a1 1 0 0 0 0 2zM7.293 5.293l-3 3l1.414 1.414l3-3zm-3 4.414l3 3l1.414-1.414l-3-3zM5 10h11V8H5zm14 3a3 3 0 0 1-3 3v2a5 5 0 0 0 5-5zm-3-3a3 3 0 0 1 3 3h2a5 5 0 0 0-5-5zM5 18h11v-2H5z"/>`), g.Group(children),
	)
}

func DoUndoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.884 6.884a1.25 1.25 0 1 0-1.768-1.768zM5 9l-.884-.884a1.25 1.25 0 0 0 0 1.768zm2.116 3.884a1.25 1.25 0 1 0 1.768-1.768zM5 15.75a1.25 1.25 0 0 0 0 2.5zM7.116 5.116l-3 3l1.768 1.768l3-3zm-3 4.768l3 3l1.768-1.768l-3-3zM5 10.25h11v-2.5H5zM18.75 13A2.75 2.75 0 0 1 16 15.75v2.5c2.9 0 5.25-2.35 5.25-5.25zM16 10.25A2.75 2.75 0 0 1 18.75 13h2.5c0-2.9-2.35-5.25-5.25-5.25zm-11 8h11v-2.5H5z"/>`), g.Group(children),
	)
}

func DoUndoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.707 6.707a1 1 0 0 0-1.414-1.414zM5 9l-.707-.707a1 1 0 0 0 0 1.414zm2.293 3.707a1 1 0 0 0 1.414-1.414zM5 16a1 1 0 1 0 0 2zM7.293 5.293l-3 3l1.414 1.414l3-3zm-3 4.414l3 3l1.414-1.414l-3-3zM5 10h11V8H5zm14 3a3 3 0 0 1-3 3v2a5 5 0 0 0 5-5zm-3-3a3 3 0 0 1 3 3h2a5 5 0 0 0-5-5zM5 18h11v-2H5z"/>`), g.Group(children),
	)
}

func DoUndoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.383 5.076A1 1 0 0 1 9 6v2h7a5 5 0 0 1 0 10H5a1 1 0 1 1 0-2h11a3 3 0 1 0 0-6H9v2a1 1 0 0 1-1.707.707l-3-3a1 1 0 0 1 0-1.414l3-3a1 1 0 0 1 1.09-.217" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DoUndoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.53 6.53a.75.75 0 0 0-1.06-1.06zM5 9l-.53-.53a.75.75 0 0 0 0 1.06zm2.47 3.53a.75.75 0 0 0 1.06-1.06zM5 16.25a.75.75 0 0 0 0 1.5zM7.47 5.47l-3 3l1.06 1.06l3-3zm-3 4.06l3 3l1.06-1.06l-3-3zm.53.22h11v-1.5H5zM19.25 13A3.25 3.25 0 0 1 16 16.25v1.5A4.75 4.75 0 0 0 20.75 13zM16 9.75A3.25 3.25 0 0 1 19.25 13h1.5A4.75 4.75 0 0 0 16 8.25zm-11 8h11v-1.5H5z"/>`), g.Group(children),
	)
}

func DoUndoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.354 6.354a.5.5 0 1 0-.708-.708zM5 9l-.354-.354a.5.5 0 0 0 0 .708zm2.646 3.354a.5.5 0 0 0 .708-.708zM5 16.5a.5.5 0 0 0 0 1zM7.646 5.646l-3 3l.708.708l3-3zm-3 3.708l3 3l.708-.708l-3-3zM5 9.5h11v-1H5zM19.5 13a3.5 3.5 0 0 1-3.5 3.5v1a4.5 4.5 0 0 0 4.5-4.5zM16 9.5a3.5 3.5 0 0 1 3.5 3.5h1A4.5 4.5 0 0 0 16 8.5zm-11 8h11v-1H5z"/>`), g.Group(children),
	)
}

func Download(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 9l-5-5m0 0L7 9m5-5v13m-6 3h12"/>`), g.Group(children),
	)
}

func DownloadBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m17 9l-5-5m0 0L7 9m5-5v13m-6 3h12"/>`), g.Group(children),
	)
}

func DownloadFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 18a1 1 0 0 1-1-1v-7H7a1 1 0 0 1-.707-1.707l5-5a1 1 0 0 1 1.414 0l5 5A1 1 0 0 1 17 10h-4v7a1 1 0 0 1-1 1m-6 1a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DownloadLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17 9l-5-5m0 0L7 9m5-5v13m-6 3h12"/>`), g.Group(children),
	)
}

func DownloadThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17 9l-5-5m0 0L7 9m5-5v13m-6 3h12"/>`), g.Group(children),
	)
}

func Edit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 16l-1 4l4-1L19.586 7.414a2 2 0 0 0 0-2.828l-.172-.172a2 2 0 0 0-2.828 0zM15 6l3 3m-5 11h8"/>`), g.Group(children),
	)
}

func EditBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m5 16l-1 4l4-1L19.586 7.414a2 2 0 0 0 0-2.828l-.172-.172a2 2 0 0 0-2.828 0zM15 6l3 3m-5 11h8"/>`), g.Group(children),
	)
}

func EditDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m5 16l-1 4l4-1L18 9l-3-3z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 16l-1 4l4-1L19.586 7.414a2 2 0 0 0 0-2.828l-.172-.172a2 2 0 0 0-2.828 0zM15 6l3 3m-5 11h8"/></g>`), g.Group(children),
	)
}

func EditFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 16l-1 4l4-1L19.586 7.414a2 2 0 0 0 0-2.828l-.172-.172a2 2 0 0 0-2.828 0z"/><path fill="currentColor" d="m5 16l-1 4l4-1L18 9l-3-3z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 6l3 3m-5 11h8"/></g>`), g.Group(children),
	)
}

func EditLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 16l-1 4l4-1L19.586 7.414a2 2 0 0 0 0-2.828l-.172-.172a2 2 0 0 0-2.828 0zM15 6l3 3m-5 11h8"/>`), g.Group(children),
	)
}

func EditThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5 16l-1 4l4-1L19.586 7.414a2 2 0 0 0 0-2.828l-.172-.172a2 2 0 0 0-2.828 0zM15 6l3 3m-5 11h8"/>`), g.Group(children),
	)
}

func Email(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 5V4a1 1 0 0 0-1 1zm18 0h1a1 1 0 0 0-1-1zM3 6h18V4H3zm17-1v12h2V5zm-1 13H5v2h14zM4 17V5H2v12zm1 1a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3zm15-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 5l9 9l9-9"/></g>`), g.Group(children),
	)
}

func EmailBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 5V3.75c-.69 0-1.25.56-1.25 1.25zm18 0h1.25c0-.69-.56-1.25-1.25-1.25zM3 6.25h18v-2.5H3zM19.75 5v12h2.5V5zM19 17.75H5v2.5h14zM4.25 17V5h-2.5v12zm.75.75a.75.75 0 0 1-.75-.75h-2.5A3.25 3.25 0 0 0 5 20.25zM19.75 17a.75.75 0 0 1-.75.75v2.5A3.25 3.25 0 0 0 22.25 17z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m3 5l9 9l9-9"/></g>`), g.Group(children),
	)
}

func EmailDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m3 5l7.586 7.586a2 2 0 0 0 2.828 0L21 5v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2z" opacity=".16"/><path fill="currentColor" d="M3 5V4a1 1 0 0 0-1 1zm18 0h1a1 1 0 0 0-1-1zM3 6h18V4H3zm17-1v12h2V5zm-1 13H5v2h14zM4 17V5H2v12zm1 1a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3zm15-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 5l9 9l9-9"/></g>`), g.Group(children),
	)
}

func EmailFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.234 4.357A1 1 0 0 0 2 5v12a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V5.01a1 1 0 0 0-.364-.781A1 1 0 0 0 21.004 4H3a1 1 0 0 0-.766.357M4 7.414V17a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V7.414l-7.293 7.293a1 1 0 0 1-1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func EmailLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 5v-.75a.75.75 0 0 0-.75.75zm18 0h.75a.75.75 0 0 0-.75-.75zM3 5.75h18v-1.5H3zM20.25 5v12h1.5V5zM19 18.25H5v1.5h14zM3.75 17V5h-1.5v12zM5 18.25c-.69 0-1.25-.56-1.25-1.25h-1.5A2.75 2.75 0 0 0 5 19.75zM20.25 17c0 .69-.56 1.25-1.25 1.25v1.5A2.75 2.75 0 0 0 21.75 17z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 5l9 9l9-9"/></g>`), g.Group(children),
	)
}

func EmailThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 5v-.5a.5.5 0 0 0-.5.5zm18 0h.5a.5.5 0 0 0-.5-.5zM3 5.5h18v-1H3zM20.5 5v12h1V5zM19 18.5H5v1h14zM3.5 17V5h-1v12zM5 18.5A1.5 1.5 0 0 1 3.5 17h-1A2.5 2.5 0 0 0 5 19.5zM20.5 17a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3 5l9 9l9-9"/></g>`), g.Group(children),
	)
}

func Enter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 4h10v14a2 2 0 0 1-2 2H9m3-5l3-3m0 0l-3-3m3 3H5"/>`), g.Group(children),
	)
}

func EnterBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M9 4h10v14a2 2 0 0 1-2 2H9m3-5l3-3m0 0l-3-3m3 3H5"/>`), g.Group(children),
	)
}

func EnterDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 4h10v14a2 2 0 0 1-2 2H9m3-5l3-3m0 0l-3-3m3 3H5"/>`), g.Group(children),
	)
}

func EnterFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 3a1 1 0 0 0 0 2h9v13a1 1 0 0 1-1 1H9a1 1 0 0 0 0 2h8a3 3 0 0 0 3-3V4a1 1 0 0 0-1-1zm3.707 5.293A1 1 0 0 0 11 9v2H5a1 1 0 0 0 0 2h6v2a1 1 0 0 0 1.707.707l3-3a1 1 0 0 0 0-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func EnterLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 4h10v14a2 2 0 0 1-2 2H9m3-5l3-3m0 0l-3-3m3 3H5"/>`), g.Group(children),
	)
}

func EnterThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 4h10v14a2 2 0 0 1-2 2H9m3-5l3-3m0 0l-3-3m3 3H5"/>`), g.Group(children),
	)
}

func Exit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 4.001H5v14a2 2 0 0 0 2 2h8m1-5l3-3m0 0l-3-3m3 3H9"/>`), g.Group(children),
	)
}

func ExitBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M15 4.001H5v14a2 2 0 0 0 2 2h8m1-5l3-3m0 0l-3-3m3 3H9"/>`), g.Group(children),
	)
}

func ExitFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 3.001a1 1 0 1 1 0 2H6v13a1 1 0 0 0 1 1h8a1 1 0 1 1 0 2H7a3 3 0 0 1-3-3v-14a1 1 0 0 1 1-1zm1.707 5.293A1 1 0 0 0 15 9v2H9a1 1 0 1 0 0 2h6v2a1 1 0 0 0 1.707.707l3-3a1 1 0 0 0 0-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ExitLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15 4.001H5v14a2 2 0 0 0 2 2h8m1-5l3-3m0 0l-3-3m3 3H9"/>`), g.Group(children),
	)
}

func ExitThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15 4.001H5v14a2 2 0 0 0 2 2h8m1-5l3-3m0 0l-3-3m3 3H9"/>`), g.Group(children),
	)
}

func ExpressionlessFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="12" r="9" stroke-linejoin="round"/><path d="M9 15h6"/><path stroke-linejoin="round" d="M16.5 9.5H14m-4 0H7.5"/></g>`), g.Group(children),
	)
}

func ExpressionlessFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><circle cx="12" cy="12" r="9" stroke-linejoin="round"/><path d="M9 15h6"/><path stroke-linejoin="round" d="M16.5 9.5H14m-4 0H7.5"/></g>`), g.Group(children),
	)
}

func ExpressionlessFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M9 15h6"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.5 9.5H14m-4 0H7.5"/></g>`), g.Group(children),
	)
}

func ExpressionlessFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m4.5-2.5a1 1 0 0 1 1-1H10a1 1 0 1 1 0 2H7.5a1 1 0 0 1-1-1m7.5-1a1 1 0 1 0 0 2h2.5a1 1 0 1 0 0-2zM9 14a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ExpressionlessFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><circle cx="12" cy="12" r="9" stroke-linejoin="round"/><path d="M9 15h6"/><path stroke-linejoin="round" d="M16.5 9.5H14m-4 0H7.5"/></g>`), g.Group(children),
	)
}

func ExpressionlessFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><circle cx="12" cy="12" r="9" stroke-linejoin="round"/><path d="M9 15h6"/><path stroke-linejoin="round" d="M16.5 9.5H14m-4 0H7.5"/></g>`), g.Group(children),
	)
}

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/><path d="M2 12c1.6-4.097 5.336-7 10-7s8.4 2.903 10 7c-1.6 4.097-5.336 7-10 7s-8.4-2.903-10-7"/></g>`), g.Group(children),
	)
}

func EyeBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/><path d="M2 12c1.6-4.097 5.336-7 10-7s8.4 2.903 10 7c-1.6 4.097-5.336 7-10 7s-8.4-2.903-10-7"/></g>`), g.Group(children),
	)
}

func EyeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M12 5C7.336 5 3.6 7.903 2 12c1.6 4.097 5.336 7 10 7s8.4-2.903 10-7c-1.6-4.097-5.336-7-10-7m0 10a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M2 12c1.6-4.097 5.336-7 10-7s8.4 2.903 10 7c-1.6 4.097-5.336 7-10 7s-8.4-2.903-10-7"/></g>`), g.Group(children),
	)
}

func EyeFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M1.069 11.636C2.803 7.194 6.884 4 12 4s9.197 3.194 10.932 7.636a1 1 0 0 1 0 .728C21.197 16.806 17.116 20 12 20s-9.197-3.194-10.931-7.636a1 1 0 0 1 0-.728M12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func EyeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/><path d="M2 12c1.6-4.097 5.336-7 10-7s8.4 2.903 10 7c-1.6 4.097-5.336 7-10 7s-8.4-2.903-10-7"/></g>`), g.Group(children),
	)
}

func EyeOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M10.73 5.073A11 11 0 0 1 12 5c4.664 0 8.4 2.903 10 7a11.6 11.6 0 0 1-1.555 2.788M6.52 6.519C4.48 7.764 2.9 9.693 2 12c1.6 4.097 5.336 7 10 7a10.44 10.44 0 0 0 5.48-1.52m-7.6-7.6a3 3 0 1 0 4.243 4.243"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func EyeOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><path stroke-linejoin="round" d="M10.73 5.073A11 11 0 0 1 12 5c4.664 0 8.4 2.903 10 7a11.6 11.6 0 0 1-1.555 2.788M6.52 6.519C4.48 7.764 2.9 9.693 2 12c1.6 4.097 5.336 7 10 7a10.44 10.44 0 0 0 5.48-1.52m-7.6-7.6a3 3 0 1 0 4.243 4.243"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func EyeOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M9 12a3 3 0 0 0 5.121 2.121l3.36 3.36A10.44 10.44 0 0 1 12 19c-4.664 0-8.4-2.903-10-7c.901-2.307 2.48-4.236 4.52-5.48l3.359 3.359A3 3 0 0 0 9 12" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.73 5.073A11 11 0 0 1 12 5c4.664 0 8.4 2.903 10 7a11.6 11.6 0 0 1-1.555 2.788M6.52 6.519C4.48 7.764 2.9 9.693 2 12c1.6 4.097 5.336 7 10 7a10.44 10.44 0 0 0 5.48-1.52m-7.6-7.6a3 3 0 1 0 4.243 4.243"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func EyeOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.707 3.293a1 1 0 0 0-1.414 1.414l1.67 1.671C3.23 7.716 1.889 9.538 1.07 11.636a1 1 0 0 0 0 .728C2.803 16.806 6.884 20 12 20c1.935 0 3.73-.459 5.31-1.276l1.983 1.983a1 1 0 0 0 1.414-1.414l-2.501-2.501l-.038-.038l-3.328-3.328l-.011-.012l-.012-.011l-4.22-4.22l-.023-.023l-3.329-3.33l-.038-.037zm4.585 7.414a3 3 0 0 0 4.001 4.001zm1.554-4.64Q11.409 6 12 6c4.074 0 7.38 2.443 8.919 6c-.34.787-.768 1.52-1.271 2.184a1 1 0 1 0 1.594 1.208a12.6 12.6 0 0 0 1.69-3.028a1 1 0 0 0 0-.728C21.197 7.194 17.116 4 12 4q-.705 0-1.386.08a1 1 0 0 0 .232 1.986" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func EyeOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M10.73 5.073A11 11 0 0 1 12 5c4.664 0 8.4 2.903 10 7a11.6 11.6 0 0 1-1.555 2.788M6.52 6.519C4.48 7.764 2.9 9.693 2 12c1.6 4.097 5.336 7 10 7a10.44 10.44 0 0 0 5.48-1.52m-7.6-7.6a3 3 0 1 0 4.243 4.243"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func EyeOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" d="M10.73 5.073A11 11 0 0 1 12 5c4.664 0 8.4 2.903 10 7a11.6 11.6 0 0 1-1.555 2.788M6.52 6.519C4.48 7.764 2.9 9.693 2 12c1.6 4.097 5.336 7 10 7a10.44 10.44 0 0 0 5.48-1.52m-7.6-7.6a3 3 0 1 0 4.243 4.243"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func EyeThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M15 12a3 3 0 1 1-6 0a3 3 0 0 1 6 0"/><path d="M2 12c1.6-4.097 5.336-7 10-7s8.4 2.903 10 7c-1.6 4.097-5.336 7-10 7s-8.4-2.903-10-7"/></g>`), g.Group(children),
	)
}

func FaceWithOpenMouth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15zM12 15h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func FaceWithOpenMouthBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15zM12 15h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func FaceWithOpenMouthDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15zM12 15h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func FaceWithOpenMouthFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7-4a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5A1.5 1.5 0 0 0 9.01 8zm6 0a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5a1.5 1.5 0 0 0-1.5-1.5zm-4.5 7a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FaceWithOpenMouthLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5zm-3 5.5v.01H12V15z"/></g>`), g.Group(children),
	)
}

func FaceWithOpenMouthThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5zm-3 5.5v.01H12V15z"/></g>`), g.Group(children),
	)
}

func FaceWithoutMouth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/></g>`), g.Group(children),
	)
}

func FaceWithoutMouthBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/></g>`), g.Group(children),
	)
}

func FaceWithoutMouthDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/></g>`), g.Group(children),
	)
}

func FaceWithoutMouthFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m5.5-2.5A1.5 1.5 0 0 1 9 8h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H9a1.5 1.5 0 0 1-1.5-1.5zm6 0A1.5 1.5 0 0 1 15 8h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H15a1.5 1.5 0 0 1-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FaceWithoutMouthLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/></g>`), g.Group(children),
	)
}

func FaceWithoutMouthThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/></g>`), g.Group(children),
	)
}

func File(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2zm5-8v4m-2-2h4"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileAddBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2zm5-8v4m-2-2h4"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileAddDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 21a2 2 0 0 1-2-2V3h8v6h6v10a2 2 0 0 1-2 2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2zm5-8v4m-2-2h4"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileAddFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2a1 1 0 0 0-1 1v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-.293-.707l-5-5A1 1 0 0 0 14 2zm9 2.414L17.586 8H14zM12 12a1 1 0 0 1 1 1v1h1a1 1 0 1 1 0 2h-1v1a1 1 0 1 1-2 0v-1h-1a1 1 0 1 1 0-2h1v-1a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileAddLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2zm5-8v4m-2-2h4"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileAddThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2zm5-8v4m-2-2h4"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileAudio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke-linejoin="round" d="M13 3v6h6"/><circle cx="11" cy="17" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12v5m2-3l-2-2"/></g>`), g.Group(children),
	)
}

func FileAudioBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke-linejoin="round" d="M13 3v6h6"/><circle cx="11" cy="17" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12v5m2-3l-2-2"/></g>`), g.Group(children),
	)
}

func FileAudioDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 21a2 2 0 0 1-2-2V3h8v6h6v10a2 2 0 0 1-2 2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v6h6"/><circle cx="11" cy="17" r="1" stroke="currentColor" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12v5m2-3l-2-2"/></g>`), g.Group(children),
	)
}

func FileAudioFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2a1 1 0 0 0-1 1v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-.293-.707l-5-5A1 1 0 0 0 14 2zm9 2.414L17.586 8H14zM11 15v-3a.997.997 0 0 1 1-1a1 1 0 0 1 .707.293l2 2a1 1 0 0 1-1.414 1.414L13 14.414V17a2 2 0 1 1-2-2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileAudioLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke-linejoin="round" d="M13 3v6h6"/><circle cx="11" cy="17" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12v5m2-3l-2-2"/></g>`), g.Group(children),
	)
}

func FileAudioThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke-linejoin="round" d="M13 3v6h6"/><circle cx="11" cy="17" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 12v5m2-3l-2-2"/></g>`), g.Group(children),
	)
}

func FileBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="m15 13l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func FileCheckBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="m15 13l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func FileCheckDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 21a2 2 0 0 1-2-2V3h8v6h6v10a2 2 0 0 1-2 2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v6h6"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 13l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func FileCheckFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2a1 1 0 0 0-1 1v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-.293-.707l-5-5A1 1 0 0 0 14 2zm9 2.414L17.586 8H14zm1.707 9.293a1 1 0 0 0-1.414-1.414L11 15.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileCheckLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="m15 13l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func FileCheckThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="m15 13l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func FileClose(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="m10 13l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func FileCloseBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="m10 13l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func FileCloseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 21a2 2 0 0 1-2-2V3h8v6h6v10a2 2 0 0 1-2 2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v6h6"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 13l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func FileCloseFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2a1 1 0 0 0-1 1v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-.293-.707l-5-5A1 1 0 0 0 14 2zm9 2.414L17.586 8H14zm-3.293 7.879a1 1 0 0 0-1.414 1.414L10.586 15l-1.293 1.293a1 1 0 1 0 1.414 1.414L12 16.414l1.293 1.293a1 1 0 0 0 1.414-1.414L13.414 15l1.293-1.293a1 1 0 0 0-1.414-1.414L12 13.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileCloseLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="m10 13l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func FileCloseThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="m10 13l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func FileDocument(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="M9 13h6m-6 4h6"/></g>`), g.Group(children),
	)
}

func FileDocumentBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="M9 13h6m-6 4h6"/></g>`), g.Group(children),
	)
}

func FileDocumentDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 21a2 2 0 0 1-2-2V3h8v6h6v10a2 2 0 0 1-2 2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v6h6"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 13h6m-6 4h6"/></g>`), g.Group(children),
	)
}

func FileDocumentFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2a1 1 0 0 0-1 1v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-.293-.707l-5-5A1 1 0 0 0 14 2zm9 2.414L17.586 8H14zM8 13a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1m0 4a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileDocumentLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="M9 13h6m-6 4h6"/></g>`), g.Group(children),
	)
}

func FileDocumentThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/><path stroke-linecap="round" d="M9 13h6m-6 4h6"/></g>`), g.Group(children),
	)
}

func FileDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 21a2 2 0 0 1-2-2V3h8v6h6v10a2 2 0 0 1-2 2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 3a1 1 0 0 1 1-1h9a1 1 0 0 1 .707.293l5 5A1 1 0 0 1 20 8v11a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3zm13.586 5L14 4.414V8z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileImage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke-linecap="round" stroke-width="2" d="m6 20l6-6l6 6"/><path stroke-width="3" d="M9.5 10.5h.01v.01H9.5z"/><path stroke-width="2" d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileImageBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke-linecap="round" stroke-width="2.5" d="m6 20l6-6l6 6"/><path stroke-width="3.75" d="M9.5 10.5h.01v.01H9.5z"/><path stroke-width="2.5" d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileImageDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 21a2 2 0 0 1-2-2V3h8v6h6v10a2 2 0 0 1-2 2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m6 20l6-6l6 6"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9.5 10.5h.01v.01H9.5z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileImageFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2a1 1 0 0 0-1 1v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-.293-.707l-5-5A1 1 0 0 0 14 2zm9 2.414L17.586 8H14zM8 10.5A1.5 1.5 0 0 1 9.5 9h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H9.5a1.5 1.5 0 0 1-1.5-1.5zm4.707 2.793a1 1 0 0 0-1.414 0l-4 4a1 1 0 1 0 1.414 1.414L12 15.414l3.293 3.293a1 1 0 0 0 1.414-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileImageLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="1.5" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke-linecap="round" stroke-width="1.5" d="m6 20l6-6l6 6"/><path stroke-width="2.25" d="M9.5 10.5h.01v.01H9.5z"/><path stroke-width="1.5" d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileImageThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke-linecap="round" d="m6 20l6-6l6 6"/><path stroke-width="1.5" d="M9.5 10.5h.01v.01H9.5z"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2zm3-6h4"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileRemoveBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2zm3-6h4"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileRemoveDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 21a2 2 0 0 1-2-2V3h8v6h6v10a2 2 0 0 1-2 2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2zm3-6h4"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileRemoveFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2a1 1 0 0 0-1 1v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-.293-.707l-5-5A1 1 0 0 0 14 2zm9 2.414L17.586 8H14zM10 14a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileRemoveLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2zm3-6h4"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileRemoveThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2zm3-6h4"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6"/></g>`), g.Group(children),
	)
}

func FileVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6m-5 5.5l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func FileVideoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6m-5 5.5l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func FileVideoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 21a2 2 0 0 1-2-2V3h8v6h6v10a2 2 0 0 1-2 2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 3v6h6m-5 5.5l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func FileVideoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 2a1 1 0 0 0-1 1v16a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V8a1 1 0 0 0-.293-.707l-5-5A1 1 0 0 0 14 2zm9 2.414L17.586 8H14zm.5 10.952a1 1 0 0 0 0-1.732l-3-1.732a1 1 0 0 0-1.5.866v3.464a1 1 0 0 0 1.5.866z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FileVideoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6m-5 5.5l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func FileVideoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M7 21a2 2 0 0 1-2-2V3h9l5 5v11a2 2 0 0 1-2 2z"/><path d="M13 3v6h6m-5 5.5l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3v2m0 16v-8m0-8h12l-2 4l2 4H8m0-8v8"/>`), g.Group(children),
	)
}

func FlagBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M8 3v2m0 16v-8m0-8h12l-2 4l2 4H8m0-8v8"/>`), g.Group(children),
	)
}

func FlagDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M20 5H8v8h12l-2-4z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 3v2m0 16v-8m0-8h12l-2 4l2 4H8m0-8v8"/></g>`), g.Group(children),
	)
}

func FlagFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20 14a1 1 0 0 0 .894-1.447L19.118 9l1.776-3.553A1 1 0 0 0 20 4H9V3a1 1 0 1 0-2 0v18a1 1 0 1 0 2 0v-7z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FlagLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 3v2m0 16v-8m0-8h12l-2 4l2 4H8m0-8v8"/>`), g.Group(children),
	)
}

func FlagThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 3v2m0 16v-8m0-8h12l-2 4l2 4H8m0-8v8"/>`), g.Group(children),
	)
}

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/>`), g.Group(children),
	)
}

func FolderAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m9-6v4m-2-2h4"/>`), g.Group(children),
	)
}

func FolderAddBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m9-6v4m-2-2h4"/>`), g.Group(children),
	)
}

func FolderAddDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m9-6v4m-2-2h4"/></g>`), g.Group(children),
	)
}

func FolderAddFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h7a1 1 0 0 1 .707.293L12.414 6H21a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm10 5a1 1 0 0 1 1 1v1h1a1 1 0 1 1 0 2h-1v1a1 1 0 1 1-2 0v-1h-1a1 1 0 1 1 0-2h1v-1a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderAddLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m9-6v4m-2-2h4"/>`), g.Group(children),
	)
}

func FolderAddThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m9-6v4m-2-2h4"/>`), g.Group(children),
	)
}

func FolderBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/>`), g.Group(children),
	)
}

func FolderCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path d="m15 11l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func FolderCheckBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path d="m15 11l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func FolderCheckDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 11l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func FolderCheckFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h7a1 1 0 0 1 .707.293L12.414 6H21a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm13.707 6.707a1 1 0 0 0-1.414-1.414L11 13.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderCheckLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path d="m15 11l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func FolderCheckThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path d="m15 11l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func FolderClose(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m7-6l4 4m-4 0l4-4"/>`), g.Group(children),
	)
}

func FolderCloseBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m7-6l4 4m-4 0l4-4"/>`), g.Group(children),
	)
}

func FolderCloseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m7-6l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func FolderCloseFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h7a1 1 0 0 1 .707.293L12.414 6H21a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm8.707 5.293a1 1 0 0 0-1.414 1.414L10.586 13l-1.293 1.293a1 1 0 1 0 1.414 1.414L12 14.414l1.293 1.293a1 1 0 0 0 1.414-1.414L13.414 13l1.293-1.293a1 1 0 0 0-1.414-1.414L12 11.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderCloseLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m7-6l4 4m-4 0l4-4"/>`), g.Group(children),
	)
}

func FolderCloseThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m7-6l4 4m-4 0l4-4"/>`), g.Group(children),
	)
}

func FolderDocument(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m4-6h2m-2 4h2"/><path d="M13 11h4v4h-4z"/></g>`), g.Group(children),
	)
}

func FolderDocumentBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m4-6h2m-2 4h2"/><path d="M13 11h4v4h-4z"/></g>`), g.Group(children),
	)
}

func FolderDocumentDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m4-6h2m-2 4h2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 11h4v4h-4z"/></g>`), g.Group(children),
	)
}

func FolderDocumentFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h7a1 1 0 0 1 .707.293L12.414 6H21a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm5 5a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2zm0 4a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2zm5-3a1 1 0 0 1 1-1h4a1 1 0 0 1 1 1v4a1 1 0 0 1-1 1h-4a1 1 0 0 1-1-1zm2 1v2h2v-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderDocumentLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m4-6h2m-2 4h2"/><path d="M13 11h4v4h-4z"/></g>`), g.Group(children),
	)
}

func FolderDocumentThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m4-6h2m-2 4h2"/><path d="M13 11h4v4h-4z"/></g>`), g.Group(children),
	)
}

func FolderDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/></g>`), g.Group(children),
	)
}

func FolderFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 0-1 1v12a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3V7a1 1 0 0 0-1-1h-8.586l-1.707-1.707A1 1 0 0 0 10 4z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderImage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path stroke-linecap="round" stroke-width="2" d="m5 19l7-7l7 7"/><path stroke-width="3" d="M7.5 9.5h.01v.01H7.5z"/></g>`), g.Group(children),
	)
}

func FolderImageBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path stroke-linecap="round" stroke-width="2.5" d="m5 19l7-7l7 7"/><path stroke-width="3.75" d="M7.5 9.5h.01v.01H7.5z"/></g>`), g.Group(children),
	)
}

func FolderImageDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 19l7-7l7 7"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M7.5 9.5h.01v.01H7.5z"/></g>`), g.Group(children),
	)
}

func FolderImageFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h7a1 1 0 0 1 .707.293L12.414 6H21a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm4 4.5A1.5 1.5 0 0 1 7.5 8h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H7.5A1.5 1.5 0 0 1 6 9.51zm6.707 1.793a1 1 0 0 0-1.414 0l-4 4a1 1 0 1 0 1.414 1.414L12 13.414l3.293 3.293a1 1 0 0 0 1.414-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderImageLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="1.5" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path stroke-linecap="round" stroke-width="1.5" d="m5 19l7-7l7 7"/><path stroke-width="2.25" d="M7.5 9.5h.01v.01H7.5z"/></g>`), g.Group(children),
	)
}

func FolderImageThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path stroke-linecap="round" d="m5 19l7-7l7 7"/><path stroke-width="1.5" d="M7.5 9.5h.01v.01H7.5z"/></g>`), g.Group(children),
	)
}

func FolderLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/>`), g.Group(children),
	)
}

func FolderMusic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><circle cx="11" cy="15" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 10v5m2-3l-2-2"/></g>`), g.Group(children),
	)
}

func FolderMusicBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><path stroke-linecap="round" stroke-linejoin="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><circle cx="11" cy="15" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 10v5m2-3l-2-2"/></g>`), g.Group(children),
	)
}

func FolderMusicDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><circle cx="11" cy="15" r="1" stroke="currentColor" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 10v5m2-3l-2-2"/></g>`), g.Group(children),
	)
}

func FolderMusicFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h7a1 1 0 0 1 .707.293L12.414 6H21a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm9 8v-3a.997.997 0 0 1 1-1a1 1 0 0 1 .707.293l2 2a1 1 0 0 1-1.414 1.414L13 12.414V15a2 2 0 1 1-2-2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderMusicLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><circle cx="11" cy="15" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 10v5m2-3l-2-2"/></g>`), g.Group(children),
	)
}

func FolderMusicThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><circle cx="11" cy="15" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M12 10v5m2-3l-2-2"/></g>`), g.Group(children),
	)
}

func FolderRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m7-4h4"/>`), g.Group(children),
	)
}

func FolderRemoveBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m7-4h4"/>`), g.Group(children),
	)
}

func FolderRemoveDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m7-4h4"/></g>`), g.Group(children),
	)
}

func FolderRemoveFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h7a1 1 0 0 1 .707.293L12.414 6H21a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm8 7a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderRemoveLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m7-4h4"/>`), g.Group(children),
	)
}

func FolderRemoveThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2m7-4h4"/>`), g.Group(children),
	)
}

func FolderThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/>`), g.Group(children),
	)
}

func FolderVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path d="m14 13l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func FolderVideoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path d="m14 13l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func FolderVideoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m14 13l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func FolderVideoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a1 1 0 0 1 1-1h7a1 1 0 0 1 .707.293L12.414 6H21a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3zm12.5 8.866a1 1 0 0 0 0-1.732l-3-1.732a1 1 0 0 0-1.5.866v3.464a1 1 0 0 0 1.5.866z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FolderVideoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path d="m14 13l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func FolderVideoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M3 17V5h7l2 2h9v10a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2"/><path d="m14 13l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func Frame(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18m-3-3v18M3 18h18M6 3v18"/>`), g.Group(children),
	)
}

func FrameBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M3 6h18m-3-3v18M3 18h18M6 3v18"/>`), g.Group(children),
	)
}

func FrameDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6 6h12v12H6z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18m-3-3v18M3 18h18M6 3v18"/></g>`), g.Group(children),
	)
}

func FrameFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 7v10H7V7zM5 7v10H3a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h10v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2V7h2a1 1 0 1 0 0-2h-2V3a1 1 0 1 0-2 0v2H7V3a1 1 0 0 0-2 0v2H3a1 1 0 0 0 0 2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FrameLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6h18m-3-3v18M3 18h18M6 3v18"/>`), g.Group(children),
	)
}

func FrameThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 6h18m-3-3v18M3 18h18M6 3v18"/>`), g.Group(children),
	)
}

func FrowningFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2" d="M8.535 16A4 4 0 0 1 12 14c1.48 0 2.773.804 3.465 2"/></g>`), g.Group(children),
	)
}

func FrowningFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2.5" d="M8.535 16A4 4 0 0 1 12 14c1.48 0 2.773.804 3.465 2"/></g>`), g.Group(children),
	)
}

func FrowningFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.535 16A4 4 0 0 1 12 14c1.48 0 2.773.804 3.465 2"/></g>`), g.Group(children),
	)
}

func FrowningFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7-4a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5A1.5 1.5 0 0 0 9.01 8zm6 0a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5a1.5 1.5 0 0 0-1.5-1.5zm.966 8.866a1 1 0 0 1-1.367-.365C14.08 15.6 13.11 15 12 15s-2.08.601-2.6 1.5a1 1 0 0 1-1.73-1A5 5 0 0 1 12 13a5 5 0 0 1 4.33 2.5a1 1 0 0 1-.364 1.366" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FrowningFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/><path stroke-linecap="round" stroke-width="1.5" d="M8.535 16A4 4 0 0 1 12 14c1.48 0 2.773.804 3.465 2"/></g>`), g.Group(children),
	)
}

func FrowningFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/><path stroke-linecap="round" d="M8.535 16A4 4 0 0 1 12 14c1.48 0 2.773.804 3.465 2"/></g>`), g.Group(children),
	)
}

func Funnel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M20 4H4l5.6 7.467a2 2 0 0 1 .4 1.2V20l4-2v-5.333a2 2 0 0 1 .4-1.2z"/>`), g.Group(children),
	)
}

func FunnelBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M20 4H4l5.6 7.467a2 2 0 0 1 .4 1.2V20l4-2v-5.333a2 2 0 0 1 .4-1.2z"/>`), g.Group(children),
	)
}

func FunnelDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M20 4H4l5.6 7.467a2 2 0 0 1 .4 1.2V20l4-2v-5.333a2 2 0 0 1 .4-1.2z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M20 4H4l5.6 7.467a2 2 0 0 1 .4 1.2V20l4-2v-5.333a2 2 0 0 1 .4-1.2z"/></g>`), g.Group(children),
	)
}

func FunnelFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 3a1 1 0 0 0-.8 1.6l5.6 7.467a1 1 0 0 1 .2.6V20a1 1 0 0 0 1.447.894l4-2A1 1 0 0 0 15 18v-5.333a1 1 0 0 1 .2-.6L20.8 4.6A1 1 0 0 0 20 3z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func FunnelLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M20 4H4l5.6 7.467a2 2 0 0 1 .4 1.2V20l4-2v-5.333a2 2 0 0 1 .4-1.2z"/>`), g.Group(children),
	)
}

func FunnelThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M20 4H4l5.6 7.467a2 2 0 0 1 .4 1.2V20l4-2v-5.333a2 2 0 0 1 .4-1.2z"/>`), g.Group(children),
	)
}

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M4 11v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/><path d="M6 4.5A2.5 2.5 0 0 1 8.5 2A3.5 3.5 0 0 1 12 5.5V7H8.5A2.5 2.5 0 0 1 6 4.5Zm12 0A2.5 2.5 0 0 0 15.5 2A3.5 3.5 0 0 0 12 5.5V7h3.5A2.5 2.5 0 0 0 18 4.5Z"/><path stroke-linecap="round" d="M3 7h18v4H3zm9 4v10"/></g>`), g.Group(children),
	)
}

func GiftBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M4 11v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/><path d="M6 4.5A2.5 2.5 0 0 1 8.5 2A3.5 3.5 0 0 1 12 5.5V7H8.5A2.5 2.5 0 0 1 6 4.5Zm12 0A2.5 2.5 0 0 0 15.5 2A3.5 3.5 0 0 0 12 5.5V7h3.5A2.5 2.5 0 0 0 18 4.5Z"/><path stroke-linecap="round" d="M3 7h18v4H3zm9 4v10"/></g>`), g.Group(children),
	)
}

func GiftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 19v-7h16v7a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12v8m-8-9v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M6 4.5A2.5 2.5 0 0 1 8.5 2A3.5 3.5 0 0 1 12 5.5V7H8.5A2.5 2.5 0 0 1 6 4.5Zm12 0A2.5 2.5 0 0 0 15.5 2A3.5 3.5 0 0 0 12 5.5V7h3.5A2.5 2.5 0 0 0 18 4.5Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 7h18v4H3zm9 4v10"/></g>`), g.Group(children),
	)
}

func GiftFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.5 1a3.5 3.5 0 0 0-3.163 5H3a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1v7a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-7a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1h-2.337A3.5 3.5 0 0 0 15.5 1A4.5 4.5 0 0 0 12 2.671A4.5 4.5 0 0 0 8.5 1M13 20h5a1 1 0 0 0 1-1v-7h-6zm-2-8v8H6a1 1 0 0 1-1-1v-7zm4.5-6a1.5 1.5 0 0 0 0-3A2.5 2.5 0 0 0 13 5.5V6zM11 6v-.5A2.5 2.5 0 0 0 8.5 3a1.5 1.5 0 1 0 0 3z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func GiftLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M4 11v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/><path d="M6 4.5A2.5 2.5 0 0 1 8.5 2A3.5 3.5 0 0 1 12 5.5V7H8.5A2.5 2.5 0 0 1 6 4.5Zm12 0A2.5 2.5 0 0 0 15.5 2A3.5 3.5 0 0 0 12 5.5V7h3.5A2.5 2.5 0 0 0 18 4.5Z"/><path stroke-linecap="round" d="M3 7h18v4H3zm9 4v10"/></g>`), g.Group(children),
	)
}

func GiftThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M4 11v8a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-8"/><path d="M6 4.5A2.5 2.5 0 0 1 8.5 2A3.5 3.5 0 0 1 12 5.5V7H8.5A2.5 2.5 0 0 1 6 4.5Zm12 0A2.5 2.5 0 0 0 15.5 2A3.5 3.5 0 0 0 12 5.5V7h3.5A2.5 2.5 0 0 0 18 4.5Z"/><path stroke-linecap="round" d="M3 7h18v4H3zm9 4v10"/></g>`), g.Group(children),
	)
}

func Headphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M21 15v-3a9 9 0 1 0-18 0v3"/><path fill="currentColor" d="M3 15v-1H2v1zm0 1h3v-2H3zm3 0v3h2v-3zm-2 3v-4H2v4zm1 1a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3zm1-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm0-3h2a2 2 0 0 0-2-2zm15-1h1v-1h-1zm-3 1h3v-2h-3zm2-1v4h2v-4zm-2 4v-3h-2v3zm1 1a1 1 0 0 1-1-1h-2a3 3 0 0 0 3 3zm1-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-2-5a2 2 0 0 0-2 2h2z"/></g>`), g.Group(children),
	)
}

func HeadphoneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="2.5" d="M21 15v-3a9 9 0 1 0-18 0v3"/><path fill="currentColor" d="M3 15v-1.25H1.75V15zm0 1.25h3v-2.5H3zM5.75 16v3h2.5v-3zm-1.5 3v-4h-2.5v4zm.75.75a.75.75 0 0 1-.75-.75h-2.5A3.25 3.25 0 0 0 5 22.25zm.75-.75a.75.75 0 0 1-.75.75v2.5A3.25 3.25 0 0 0 8.25 19zM6 16.25a.25.25 0 0 1-.25-.25h2.5A2.25 2.25 0 0 0 6 13.75zM21 15h1.25v-1.25H21zm-3 1.25h3v-2.5h-3zM19.75 15v4h2.5v-4zm-1.5 4v-3h-2.5v3zm.75.75a.75.75 0 0 1-.75-.75h-2.5A3.25 3.25 0 0 0 19 22.25zm.75-.75a.75.75 0 0 1-.75.75v2.5A3.25 3.25 0 0 0 22.25 19zM18 13.75A2.25 2.25 0 0 0 15.75 16h2.5a.25.25 0 0 1-.25.25z"/></g>`), g.Group(children),
	)
}

func HeadphoneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 15h3a1 1 0 0 1 1 1v3a2 2 0 1 1-4 0zm14 1a1 1 0 0 1 1-1h3v4a2 2 0 1 1-4 0z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M21 15v-3a9 9 0 1 0-18 0v3"/><path fill="currentColor" d="M3 15v-1H2v1zm0 1h3v-2H3zm3 0v3h2v-3zm-2 3v-4H2v4zm1 1a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3zm1-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm0-3h2a2 2 0 0 0-2-2zm15-1h1v-1h-1zm-3 1h3v-2h-3zm2-1v4h2v-4zm-2 4v-3h-2v3zm1 1a1 1 0 0 1-1-1h-2a3 3 0 0 0 3 3zm1-1a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-2-5a2 2 0 0 0-2 2h2z"/></g>`), g.Group(children),
	)
}

func HeadphoneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 4a8 8 0 0 0-8 8v2h2a2 2 0 0 1 2 2v3a3 3 0 1 1-6 0v-7C2 6.477 6.477 2 12 2s10 4.477 10 10v7a3 3 0 1 1-6 0v-3a2 2 0 0 1 2-2h2v-2a8 8 0 0 0-8-8" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func HeadphoneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M21 15v-3a9 9 0 1 0-18 0v3"/><path fill="currentColor" d="M3 15v-.75h-.75V15zm0 .75h3v-1.5H3zm3.25.25v3h1.5v-3zm-2.5 3v-4h-1.5v4zM5 20.25c-.69 0-1.25-.56-1.25-1.25h-1.5A2.75 2.75 0 0 0 5 21.75zM6.25 19c0 .69-.56 1.25-1.25 1.25v1.5A2.75 2.75 0 0 0 7.75 19zM6 15.75a.25.25 0 0 1 .25.25h1.5A1.75 1.75 0 0 0 6 14.25zM21 15h.75v-.75H21zm-3 .75h3v-1.5h-3zm2.25-.75v4h1.5v-4zm-2.5 4v-3h-1.5v3zM19 20.25c-.69 0-1.25-.56-1.25-1.25h-1.5A2.75 2.75 0 0 0 19 21.75zM20.25 19c0 .69-.56 1.25-1.25 1.25v1.5A2.75 2.75 0 0 0 21.75 19zM18 14.25A1.75 1.75 0 0 0 16.25 16h1.5a.25.25 0 0 1 .25-.25z"/></g>`), g.Group(children),
	)
}

func HeadphoneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" d="M21 15v-3a9 9 0 1 0-18 0v3"/><path fill="currentColor" d="M3 15v-.5h-.5v.5zm0 .5h3v-1H3zm3.5.5v3h1v-3zm-3 3v-4h-1v4zM5 20.5A1.5 1.5 0 0 1 3.5 19h-1A2.5 2.5 0 0 0 5 21.5zM6.5 19A1.5 1.5 0 0 1 5 20.5v1A2.5 2.5 0 0 0 7.5 19zM6 15.5a.5.5 0 0 1 .5.5h1A1.5 1.5 0 0 0 6 14.5zm15-.5h.5v-.5H21zm-3 .5h3v-1h-3zm2.5-.5v4h1v-4zm-3 4v-3h-1v3zm1.5 1.5a1.5 1.5 0 0 1-1.5-1.5h-1a2.5 2.5 0 0 0 2.5 2.5zm1.5-1.5a1.5 1.5 0 0 1-1.5 1.5v1a2.5 2.5 0 0 0 2.5-2.5zM18 14.5a1.5 1.5 0 0 0-1.5 1.5h1a.5.5 0 0 1 .5-.5z"/></g>`), g.Group(children),
	)
}

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.071 13.142L13.414 18.8a2 2 0 0 1-2.828 0l-5.657-5.657A5 5 0 1 1 12 6.072a5 5 0 0 1 7.071 7.07"/>`), g.Group(children),
	)
}

func HeartBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M19.071 13.142L13.414 18.8a2 2 0 0 1-2.828 0l-5.657-5.657A5 5 0 1 1 12 6.072a5 5 0 0 1 7.071 7.07"/>`), g.Group(children),
	)
}

func HeartDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M19.071 13.142L13.414 18.8a2 2 0 0 1-2.828 0l-5.657-5.657A5 5 0 1 1 12 6.072a5 5 0 0 1 7.071 7.07" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.071 13.142L13.414 18.8a2 2 0 0 1-2.828 0l-5.657-5.657a5 5 0 0 1 7.07-7.071a5 5 0 0 1 7.072 7.071"/></g>`), g.Group(children),
	)
}

func HeartFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.222 5.364A6 6 0 0 1 12 4.758a6.002 6.002 0 0 1 7.778 9.091l-5.657 5.657a3 3 0 0 1-4.242 0L4.222 13.85a6 6 0 0 1 0-8.485" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func HeartLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.071 13.142L13.414 18.8a2 2 0 0 1-2.828 0l-5.657-5.657A5 5 0 1 1 12 6.072a5 5 0 0 1 7.071 7.07"/>`), g.Group(children),
	)
}

func HeartOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M5.546 5.546q-.325.233-.617.525a5 5 0 0 0 0 7.071l5.657 5.657a2 2 0 0 0 2.828 0l2.693-2.692M10.89 5.233c.398.22.772.5 1.11.838a5 5 0 0 1 7.071 7.071l-.136.136"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func HeartOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><path stroke-linejoin="round" d="M5.546 5.546q-.325.233-.617.525a5 5 0 0 0 0 7.071l5.657 5.657a2 2 0 0 0 2.828 0l2.693-2.692M10.89 5.233c.398.22.772.5 1.11.838a5 5 0 0 1 7.071 7.071l-.136.136"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func HeartOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m4.929 13.142l5.657 5.657a2 2 0 0 0 2.828 0l2.693-2.692L5.546 5.546q-.325.232-.617.525a5 5 0 0 0 0 7.071" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.546 5.546q-.325.233-.617.525a5 5 0 0 0 0 7.071l5.657 5.657a2 2 0 0 0 2.828 0l2.693-2.692M10.89 5.233c.398.22.772.5 1.11.838a5 5 0 0 1 7.071 7.071l-.136.136"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func HeartOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.707 3.293a1 1 0 0 0-1.414 1.414l.795.795a6 6 0 0 0 .134 8.347l5.657 5.657a3 3 0 0 0 4.242 0l1.986-1.985l3.186 3.186a1 1 0 0 0 1.414-1.414l-3.885-3.885l-.017-.017L6.267 4.853l-.03-.03zM12 4.758a6.002 6.002 0 0 1 7.778 9.091l-.136.136a1 1 0 0 1-1.414-1.414l.136-.136a4 4 0 0 0-5.657-5.657a1 1 0 0 1-1.414 0a4 4 0 0 0-.889-.671a1 1 0 1 1 .971-1.749q.322.18.625.4" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func HeartOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M5.546 5.546q-.325.233-.617.525a5 5 0 0 0 0 7.071l5.657 5.657a2 2 0 0 0 2.828 0l2.693-2.692M10.89 5.233c.398.22.772.5 1.11.838a5 5 0 0 1 7.071 7.071l-.136.136"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func HeartOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" d="M5.546 5.546q-.325.233-.617.525a5 5 0 0 0 0 7.071l5.657 5.657a2 2 0 0 0 2.828 0l2.693-2.692M10.89 5.233c.398.22.772.5 1.11.838a5 5 0 0 1 7.071 7.071l-.136.136"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func HeartThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.071 13.142L13.414 18.8a2 2 0 0 1-2.828 0l-5.657-5.657A5 5 0 1 1 12 6.072a5 5 0 0 1 7.071 7.07"/>`), g.Group(children),
	)
}

func History(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M5.636 18.364A9 9 0 1 0 3 12.004V14"/><path d="m1 12l2 2l2-2m6-4v5h5"/></g>`), g.Group(children),
	)
}

func HistoryBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M5.636 18.364A9 9 0 1 0 3 12.004V14"/><path d="m1 12l2 2l2-2m6-4v5h5"/></g>`), g.Group(children),
	)
}

func HistoryDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5.636 18.364A9 9 0 1 0 3 12.004V14"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 12l2 2l2-2m6-4v5h5"/></g>`), g.Group(children),
	)
}

func HistoryFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.929 17.657a1 1 0 0 0 0 1.414c3.905 3.905 10.237 3.905 14.142 0s3.905-10.237 0-14.142s-10.237-3.905-14.142 0A9.96 9.96 0 0 0 2.049 11H1a1 1 0 0 0-.707 1.707l2 2l.002.002a.997.997 0 0 0 1.413-.003l2-1.999A1 1 0 0 0 5 11h-.938a8 8 0 1 1 2.28 6.657a1 1 0 0 0-1.413 0M10 8a1 1 0 1 1 2 0v4h4a1 1 0 1 1 0 2h-5a1 1 0 0 1-1-1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func HistoryLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.636 18.364A9 9 0 1 0 3 12.004V14"/><path d="m1 12l2 2l2-2m6-4v5h5"/></g>`), g.Group(children),
	)
}

func HistoryThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M5.636 18.364A9 9 0 1 0 3 12.004V14"/><path d="m1 12l2 2l2-2m6-4v5h5"/></g>`), g.Group(children),
	)
}

func Home(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 10a1 1 0 1 0-2 0zM6 10a1 1 0 0 0-2 0zm14.293 2.707a1 1 0 0 0 1.414-1.414zM12 3l.707-.707a1 1 0 0 0-1.414 0zm-9.707 8.293a1 1 0 1 0 1.414 1.414zM7 22h10v-2H7zm13-3v-9h-2v9zM6 19v-9H4v9zm15.707-7.707l-9-9l-1.414 1.414l9 9zm-10.414-9l-9 9l1.414 1.414l9-9zM17 22a3 3 0 0 0 3-3h-2a1 1 0 0 1-1 1zM7 20a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3z"/>`), g.Group(children),
	)
}

func HomeBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.25 10a1.25 1.25 0 1 0-2.5 0zm-14 0a1.25 1.25 0 1 0-2.5 0zm13.866 2.884a1.25 1.25 0 0 0 1.768-1.768zM12 3l.884-.884a1.25 1.25 0 0 0-1.768 0zm-9.884 8.116a1.25 1.25 0 0 0 1.768 1.768zM7 22.25h10v-2.5H7zM20.25 19v-9h-2.5v9zm-14 0v-9h-2.5v9zm15.634-7.884l-9-9l-1.768 1.768l9 9zm-10.768-9l-9 9l1.768 1.768l9-9zM17 22.25A3.25 3.25 0 0 0 20.25 19h-2.5a.75.75 0 0 1-.75.75zm-10-2.5a.75.75 0 0 1-.75-.75h-2.5A3.25 3.25 0 0 0 7 22.25z"/>`), g.Group(children),
	)
}

func HomeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M17 21H7a2 2 0 0 1-2-2v-9l7-7l7 7v9a2 2 0 0 1-2 2" opacity=".16"/><path d="M20 10a1 1 0 1 0-2 0zM6 10a1 1 0 0 0-2 0zm14.293 2.707a1 1 0 0 0 1.414-1.414zM12 3l.707-.707a1 1 0 0 0-1.414 0zm-9.707 8.293a1 1 0 1 0 1.414 1.414zM7 22h10v-2H7zm13-3v-9h-2v9zM6 19v-9H4v9zm15.707-7.707l-9-9l-1.414 1.414l9 9zm-10.414-9l-9 9l1.414 1.414l9-9zM17 22a3 3 0 0 0 3-3h-2a1 1 0 0 1-1 1zM7 20a1 1 0 0 1-1-1H4a3 3 0 0 0 3 3z"/></g>`), g.Group(children),
	)
}

func HomeFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.707 2.293a1 1 0 0 0-1.414 0l-7 7l-2 2a1 1 0 1 0 1.414 1.414L4 12.414V19a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-6.586l.293.293a1 1 0 0 0 1.414-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func HomeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.75 10a.75.75 0 0 0-1.5 0zm-14 0a.75.75 0 0 0-1.5 0zm14.72 2.53a.75.75 0 1 0 1.06-1.06zM12 3l.53-.53a.75.75 0 0 0-1.06 0zm-9.53 8.47a.75.75 0 1 0 1.06 1.06zM7 21.75h10v-1.5H7zM19.75 19v-9h-1.5v9zm-14 0v-9h-1.5v9zm15.78-7.53l-9-9l-1.06 1.06l9 9zm-10.06-9l-9 9l1.06 1.06l9-9zM17 21.75A2.75 2.75 0 0 0 19.75 19h-1.5c0 .69-.56 1.25-1.25 1.25zm-10-1.5c-.69 0-1.25-.56-1.25-1.25h-1.5A2.75 2.75 0 0 0 7 21.75z"/>`), g.Group(children),
	)
}

func HomeThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.5 10a.5.5 0 0 0-1 0zm-14 0a.5.5 0 0 0-1 0zm15.146 2.354a.5.5 0 0 0 .708-.708zM12 3l.354-.354a.5.5 0 0 0-.708 0zm-9.354 8.646a.5.5 0 0 0 .708.708zM7 21.5h10v-1H7zM19.5 19v-9h-1v9zm-14 0v-9h-1v9zm15.854-7.354l-9-9l-.708.708l9 9zm-9.708-9l-9 9l.708.708l9-9zM17 21.5a2.5 2.5 0 0 0 2.5-2.5h-1a1.5 1.5 0 0 1-1.5 1.5zm-10-1A1.5 1.5 0 0 1 5.5 19h-1A2.5 2.5 0 0 0 7 21.5z"/>`), g.Group(children),
	)
}

func InformationCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M12 8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M12 12v4"/></g>`), g.Group(children),
	)
}

func InformationCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M12 8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M12 12v4"/></g>`), g.Group(children),
	)
}

func InformationCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 8h.01v.01H12z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12v4"/></g>`), g.Group(children),
	)
}

func InformationCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m8.5-4A1.5 1.5 0 0 1 12 6.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5zm1.5 3a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0v-4a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func InformationCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M12 8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M12 12v4"/></g>`), g.Group(children),
	)
}

func InformationCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M12 8h.01v.01H12z"/><path stroke-linecap="round" d="M12 12v4"/></g>`), g.Group(children),
	)
}

func InformationSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3" d="M12 8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M12 12v4"/></g>`), g.Group(children),
	)
}

func InformationSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3.75" d="M12 8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M12 12v4"/></g>`), g.Group(children),
	)
}

func InformationSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 8h.01v.01H12z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 12v4"/></g>`), g.Group(children),
	)
}

func InformationSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zM10.5 8A1.5 1.5 0 0 1 12 6.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5zm1.5 3a1 1 0 0 1 1 1v4a1 1 0 1 1-2 0v-4a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func InformationSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="1.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="2.25" d="M12 8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M12 12v4"/></g>`), g.Group(children),
	)
}

func InformationSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="1.5" d="M12 8h.01v.01H12z"/><path stroke-linecap="round" d="M12 12v4"/></g>`), g.Group(children),
	)
}

func Invoice(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3h14v18l-1.032-.884a2 2 0 0 0-2.603 0L14.333 21l-1.031-.884a2 2 0 0 0-2.604 0L9.667 21l-1.032-.884a2 2 0 0 0-2.603 0L5 21zm10 4H9m6 4H9m6 4h-4"/>`), g.Group(children),
	)
}

func InvoiceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 3h14v18l-1.032-.884a2 2 0 0 0-2.603 0L14.333 21l-1.031-.884a2 2 0 0 0-2.604 0L9.667 21l-1.032-.884a2 2 0 0 0-2.603 0L5 21zm10 4H9m6 4H9m6 4h-4"/>`), g.Group(children),
	)
}

func InvoiceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M5 3h14v18l-1.032-.884a2 2 0 0 0-2.603 0L14.333 21l-1.031-.884a2 2 0 0 0-2.604 0L9.667 21l-1.032-.884a2 2 0 0 0-2.603 0L5 21z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 3h14v18l-1.032-.884a2 2 0 0 0-2.603 0L14.333 21l-1.031-.884a2 2 0 0 0-2.604 0L9.667 21l-1.032-.884a2 2 0 0 0-2.603 0L5 21zm10 4H9m6 4H9m6 4h-4"/></g>`), g.Group(children),
	)
}

func InvoiceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 3a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1v18a1 1 0 0 1-1.65.76l-1.033-.885a1 1 0 0 0-1.301 0l-1.032.884a1 1 0 0 1-1.302 0l-1.031-.884a1 1 0 0 0-1.302 0l-1.031.884a1 1 0 0 1-1.302 0l-1.032-.884a1 1 0 0 0-1.301 0l-1.032.884A1 1 0 0 1 4 21zm5 3a1 1 0 0 0 0 2h6a1 1 0 1 0 0-2zm0 4a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2zm1 5a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2h-4a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func InvoiceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 3h14v18l-1.032-.884a2 2 0 0 0-2.603 0L14.333 21l-1.031-.884a2 2 0 0 0-2.604 0L9.667 21l-1.032-.884a2 2 0 0 0-2.603 0L5 21zm10 4H9m6 4H9m6 4h-4"/>`), g.Group(children),
	)
}

func InvoiceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 3h14v18l-1.032-.884a2 2 0 0 0-2.603 0L14.333 21l-1.031-.884a2 2 0 0 0-2.604 0L9.667 21l-1.032-.884a2 2 0 0 0-2.603 0L5 21zm10 4H9m6 4H9m6 4h-4"/>`), g.Group(children),
	)
}

func KissingFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-linecap="round" stroke-width="2" d="m13 16l-2-1l2-1"/><path stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/></g>`), g.Group(children),
	)
}

func KissingFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-linecap="round" stroke-width="2.5" d="m13 16l-2-1l2-1"/><path stroke-width="3.75" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/></g>`), g.Group(children),
	)
}

func KissingFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m13 16l-2-1l2-1"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/></g>`), g.Group(children),
	)
}

func KissingFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7-4a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5A1.5 1.5 0 0 0 9.01 8zm4.447 6.894a1 1 0 1 0-.894-1.788l-2 1a1 1 0 0 0 0 1.788l2 1a1 1 0 1 0 .894-1.788l-.21-.106zM13.5 9.5A1.5 1.5 0 0 1 15 8h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H15a1.5 1.5 0 0 1-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func KissingFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-linecap="round" stroke-width="1.5" d="m13 16l-2-1l2-1"/><path stroke-width="2.25" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/></g>`), g.Group(children),
	)
}

func KissingFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-linecap="round" d="m13 16l-2-1l2-1"/><path stroke-width="1.5" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/></g>`), g.Group(children),
	)
}

func KissingFaceWithSmilingEyes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m14 10l1-1l1 1m-6 0L9 9l-1 1m5 6l-2-1l2-1"/></g>`), g.Group(children),
	)
}

func KissingFaceWithSmilingEyesBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m14 10l1-1l1 1m-6 0L9 9l-1 1m5 6l-2-1l2-1"/></g>`), g.Group(children),
	)
}

func KissingFaceWithSmilingEyesDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 10l1-1l1 1m-6 0L9 9l-1 1m5 6l-2-1l2-1"/></g>`), g.Group(children),
	)
}

func KissingFaceWithSmilingEyesFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m6.707-1.293a1 1 0 0 1-1.414-1.414l1-1a1 1 0 0 1 1.414 0l1 1a1 1 0 0 1-1.414 1.414L9 10.414zm6.586 0a1 1 0 0 0 1.414-1.414l-1-1a1 1 0 0 0-1.414 0l-1 1a1 1 0 0 0 1.414 1.414l.293-.293zm-1.399 2.846a1 1 0 0 1-.447 1.341l-.21.106l.21.106a1 1 0 1 1-.894 1.788l-2-1a1 1 0 0 1 0-1.788l2-1a1 1 0 0 1 1.341.447" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func KissingFaceWithSmilingEyesLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m14 10l1-1l1 1m-6 0L9 9l-1 1m5 6l-2-1l2-1"/></g>`), g.Group(children),
	)
}

func KissingFaceWithSmilingEyesThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m14 10l1-1l1 1m-6 0L9 9l-1 1m5 6l-2-1l2-1"/></g>`), g.Group(children),
	)
}

func LightningOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m6 14l7-12v8h5l-7 12v-8z"/>`), g.Group(children),
	)
}

func LightningOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="m6 14l7-12v8h5l-7 12v-8z"/>`), g.Group(children),
	)
}

func LightningOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m6 14l7-12v8h5l-7 12v-8z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m6 14l7-12v8h5l-7 12v-8z"/></g>`), g.Group(children),
	)
}

func LightningOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 2a1 1 0 0 0-1.864-.504l-7 12A1 1 0 0 0 6 15h4v7a1 1 0 0 0 1.864.504l7-12A1 1 0 0 0 18 9h-4z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LightningOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="m6 14l7-12v8h5l-7 12v-8z"/>`), g.Group(children),
	)
}

func LightningOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m6 14l7-12v8h5l-7 12v-8z"/>`), g.Group(children),
	)
}

func LightningTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M11 14H6L9.5 2H16l-3 8h5l-8 12z"/>`), g.Group(children),
	)
}

func LightningTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M11 14H6L9.5 2H16l-3 8h5l-8 12z"/>`), g.Group(children),
	)
}

func LightningTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M11 14H6L9.5 2H16l-3 8h5l-8 12z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M11 14H6L9.5 2H16l-3 8h5l-8 12z"/></g>`), g.Group(children),
	)
}

func LightningTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.5 1a1 1 0 0 0-.96.72l-3.5 12A1 1 0 0 0 6 15h3.867l-.86 6.876a1 1 0 0 0 1.825.679l8-12A1 1 0 0 0 18 9h-3.557l2.493-6.649A1 1 0 0 0 16 1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LightningTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M11 14H6L9.5 2H16l-3 8h5l-8 12z"/>`), g.Group(children),
	)
}

func LightningTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M11 14H6L9.5 2H16l-3 8h5l-8 12z"/>`), g.Group(children),
	)
}

func Like(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 10l-.986-.164A1 1 0 0 0 15 11zM4 10V9a1 1 0 0 0-1 1zm16.522 2.392l.98.196zM6 21h11.36v-2H6zM18.56 9H15v2h3.56zm-2.573 1.164l.805-4.835L14.82 5l-.806 4.836zM14.82 3h-.214v2h.214zm-3.543 1.781L8.763 8.555l1.664 1.11l2.516-3.774zM7.93 9H4v2h3.93zM3 10v8h2v-8zm17.302 8.588l1.2-6l-1.96-.392l-1.2 6zM8.762 8.555A1 1 0 0 1 7.93 9v2a3 3 0 0 0 2.496-1.336zm8.03-3.226A2 2 0 0 0 14.82 3v2zM18.56 11a1 1 0 0 1 .981 1.196l1.961.392A3 3 0 0 0 18.561 9zm-1.2 10a3 3 0 0 0 2.942-2.412l-1.96-.392a1 1 0 0 1-.982.804zM14.606 3a4 4 0 0 0-3.328 1.781l1.664 1.11A2 2 0 0 1 14.606 5zM6 19a1 1 0 0 1-1-1H3a3 3 0 0 0 3 3z"/><path stroke="currentColor" stroke-width="2" d="M8 10v10"/></g>`), g.Group(children),
	)
}

func LikeBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 10l-1.233-.206A1.25 1.25 0 0 0 15 11.25zM4 10V8.75c-.69 0-1.25.56-1.25 1.25zm16.522 2.392l1.225.245zM6 21.25h11.36v-2.5H6zm12.56-12.5H15v2.5h3.56zm-2.327 1.456l.806-4.836l-2.466-.411l-.806 4.836zM14.82 2.75h-.214v2.5h.214zm-3.75 1.893L8.554 8.416l2.08 1.387l2.516-3.774zM7.93 8.75H4v2.5h3.93zM2.75 10v8h2.5v-8zm17.797 8.637l1.2-6l-2.451-.49l-1.2 6zM8.554 8.416a.75.75 0 0 1-.624.334v2.5a3.25 3.25 0 0 0 2.704-1.447zm8.485-3.046a2.25 2.25 0 0 0-2.22-2.62v2.5a.25.25 0 0 1-.246-.291zm1.521 5.88a.75.75 0 0 1 .736.897l2.451.49a3.25 3.25 0 0 0-3.186-3.887zm-1.2 10a3.25 3.25 0 0 0 3.187-2.613l-2.451-.49a.75.75 0 0 1-.736.603zm-2.754-18.5a4.25 4.25 0 0 0-3.536 1.893l2.08 1.386a1.75 1.75 0 0 1 1.456-.779zM6 18.75a.75.75 0 0 1-.75-.75h-2.5A3.25 3.25 0 0 0 6 21.25z"/><path stroke="currentColor" stroke-width="2.5" d="M8 10v10"/></g>`), g.Group(children),
	)
}

func LikeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M17.36 20H8V10c.625 0 1.208-.312 1.555-.832l2.554-3.832A3 3 0 0 1 14.606 4h.213a1 1 0 0 1 .987 1.164L15 10h3.56a2 2 0 0 1 1.962 2.392l-1.2 6A2 2 0 0 1 17.36 20" opacity=".16"/><path fill="currentColor" d="m15 10l-.986-.164A1 1 0 0 0 15 11zM4 10V9a1 1 0 0 0-1 1zm16.522 2.392l.98.196zM6 21h11.36v-2H6zM18.56 9H15v2h3.56zm-2.574 1.164l.806-4.835L14.82 5l-.805 4.836zM14.82 3h-.213v2h.213zm-3.542 1.781L8.762 8.555l1.664 1.11L12.94 5.89zM7.93 9H4v2h3.93zM3 10v8h2v-8zm17.302 8.588l1.2-6l-1.961-.392l-1.2 6zM8.762 8.555A1 1 0 0 1 7.93 9v2a3 3 0 0 0 2.496-1.336zm8.03-3.226A2 2 0 0 0 14.82 3v2zM18.56 11a1 1 0 0 1 .981 1.196l1.961.392A3 3 0 0 0 18.56 9zm-1.2 10a3 3 0 0 0 2.942-2.412l-1.961-.392a1 1 0 0 1-.98.804zM14.606 3a4 4 0 0 0-3.329 1.781l1.665 1.11A2 2 0 0 1 14.606 5zM6 19a1 1 0 0 1-1-1H3a3 3 0 0 0 3 3z"/><path stroke="currentColor" stroke-width="2" d="M8 10v10"/></g>`), g.Group(children),
	)
}

func LikeFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.277 4.781A4 4 0 0 1 14.606 3h.213a2 2 0 0 1 1.973 2.329L16.18 9h2.38a3 3 0 0 1 2.942 3.588l-1.2 6A3 3 0 0 1 17.36 21H6a3 3 0 0 1-3-3v-8a1 1 0 0 1 1-1h3.93a1 1 0 0 0 .832-.445zM7 11H5v7a1 1 0 0 0 1 1h1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LikeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 10l-.74-.123a.75.75 0 0 0 .74.873zM4 10v-.75a.75.75 0 0 0-.75.75zm16.522 2.392l.735.147zM6 20.75h11.36v-1.5H6zm12.56-11.5H15v1.5h3.56zm-2.82.873l.806-4.835l-1.48-.247l-.806 4.836zm-.92-6.873h-.214v1.5h.213zm-3.335 1.67L8.97 8.693l1.248.832l2.515-3.773zM7.93 9.25H4v1.5h3.93zM3.25 10v8h1.5v-8zm16.807 8.54l1.2-6l-1.47-.295l-1.2 6zM8.97 8.692a1.25 1.25 0 0 1-1.04.557v1.5c.92 0 1.778-.46 2.288-1.225zm7.576-3.405A1.75 1.75 0 0 0 14.82 3.25v1.5a.25.25 0 0 1 .246.291zm2.014 5.462c.79 0 1.38.722 1.226 1.495l1.471.294A2.75 2.75 0 0 0 18.56 9.25zm-1.2 10a2.75 2.75 0 0 0 2.697-2.21l-1.47-.295a1.25 1.25 0 0 1-1.227 1.005zm-2.754-17.5a3.75 3.75 0 0 0-3.12 1.67l1.247.832a2.25 2.25 0 0 1 1.873-1.002zM6 19.25c-.69 0-1.25-.56-1.25-1.25h-1.5A2.75 2.75 0 0 0 6 20.75z"/><path stroke="currentColor" stroke-width="1.5" d="M8 10v10"/></g>`), g.Group(children),
	)
}

func LikeThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m15 10l-.493-.082A.5.5 0 0 0 15 10.5zM4 10v-.5a.5.5 0 0 0-.5.5zm16.522 2.392l.49.098zM6 20.5h11.36v-1H6zm12.56-11H15v1h3.56zm-3.067.582l.806-4.835l-.986-.165l-.806 4.836zM14.82 3.5h-.213v1h.213zm-3.126 1.559L9.178 8.832l.832.555l2.515-3.774zM7.93 9.5H4v1h3.93zM3.5 10v8h1v-8zm16.312 8.49l1.2-6l-.98-.196l-1.2 6zM9.178 8.832A1.5 1.5 0 0 1 7.93 9.5v1a2.5 2.5 0 0 0 2.08-1.113zm7.121-3.585A1.5 1.5 0 0 0 14.82 3.5v1a.5.5 0 0 1 .494.582zM18.56 10.5a1.5 1.5 0 0 1 1.471 1.794l.98.196a2.5 2.5 0 0 0-2.45-2.99zm-1.2 10a2.5 2.5 0 0 0 2.452-2.01l-.98-.196A1.5 1.5 0 0 1 17.36 19.5zm-2.754-17a3.5 3.5 0 0 0-2.913 1.559l.832.554a2.5 2.5 0 0 1 2.08-1.113zM6 19.5A1.5 1.5 0 0 1 4.5 18h-1A2.5 2.5 0 0 0 6 20.5z"/><path stroke="currentColor" d="M8 10v10"/></g>`), g.Group(children),
	)
}

func Link(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h8M9 8H6a4 4 0 1 0 0 8h3m6-8h3a4 4 0 0 1 0 8h-3"/>`), g.Group(children),
	)
}

func LinkBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M8 12h8M9 8H6a4 4 0 1 0 0 8h3m6-8h3a4 4 0 0 1 0 8h-3"/>`), g.Group(children),
	)
}

func LinkDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 12h8M9 8H6a4 4 0 1 0 0 8h3m6-8h3a4 4 0 0 1 0 8h-3"/>`), g.Group(children),
	)
}

func LinkExternal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4H4v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5M9 15L20 4m-5 0h5v5"/>`), g.Group(children),
	)
}

func LinkExternalBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M11 4H4v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5M9 15L20 4m-5 0h5v5"/>`), g.Group(children),
	)
}

func LinkExternalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 4H4v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5M9 15L20 4m-5 0h5v5"/></g>`), g.Group(children),
	)
}

func LinkExternalFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.707 15.707L17.5 7.914l1.793 1.793A1 1 0 0 0 21 9V4l-.001-.048A.996.996 0 0 0 20 3h-5a1 1 0 0 0-.707 1.707L16.086 6.5l-7.793 7.793a1 1 0 1 0 1.414 1.414M4 3a1 1 0 0 0-1 1v14a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3v-5a1 1 0 1 0-2 0v5a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V5h6a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LinkExternalLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 4H4v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5M9 15L20 4m-5 0h5v5"/>`), g.Group(children),
	)
}

func LinkExternalThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 4H4v14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2v-5M9 15L20 4m-5 0h5v5"/>`), g.Group(children),
	)
}

func LinkFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 9a3 3 0 1 0 0 6h3a1 1 0 1 1 0 2H6A5 5 0 0 1 6 7h3a1 1 0 0 1 0 2zm1 3a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1m8-5a1 1 0 1 0 0 2h3a3 3 0 1 1 0 6h-3a1 1 0 1 0 0 2h3a5 5 0 0 0 0-10z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LinkLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 12h8M9 8H6a4 4 0 1 0 0 8h3m6-8h3a4 4 0 0 1 0 8h-3"/>`), g.Group(children),
	)
}

func LinkThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 12h8M9 8H6a4 4 0 1 0 0 8h3m6-8h3a4 4 0 0 1 0 8h-3"/>`), g.Group(children),
	)
}

func Location(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.252 9.975l11.66-5.552c1.7-.81 3.474.965 2.665 2.666l-5.552 11.659c-.759 1.593-3.059 1.495-3.679-.158L9.32 15.851a2 2 0 0 0-1.17-1.17l-2.74-1.027c-1.652-.62-1.75-2.92-.157-3.679"/>`), g.Group(children),
	)
}

func LocationBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m5.252 9.975l11.66-5.552c1.7-.81 3.474.965 2.665 2.666l-5.552 11.659c-.759 1.593-3.059 1.495-3.679-.158L9.32 15.851a2 2 0 0 0-1.17-1.17l-2.74-1.027c-1.652-.62-1.75-2.92-.157-3.679"/>`), g.Group(children),
	)
}

func LocationDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m5.252 9.975l11.66-5.552c1.7-.81 3.474.965 2.665 2.666l-5.552 11.659c-.759 1.593-3.059 1.495-3.679-.158L9.32 15.851a2 2 0 0 0-1.17-1.17l-2.74-1.027c-1.652-.62-1.75-2.92-.157-3.679" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.252 9.975l11.66-5.552c1.7-.81 3.474.965 2.665 2.666l-5.552 11.659c-.759 1.593-3.059 1.495-3.679-.158L9.32 15.851a2 2 0 0 0-1.17-1.17l-2.74-1.027c-1.652-.62-1.75-2.92-.157-3.679"/></g>`), g.Group(children),
	)
}

func LocationFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20.48 7.519c1.214-2.55-1.448-5.213-3.999-3.999L4.822 9.072c-2.39 1.138-2.242 4.589.237 5.518l2.739 1.027a1 1 0 0 1 .585.585l1.027 2.74c.93 2.478 4.38 2.626 5.518.236z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LocationLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.252 9.975l11.66-5.552c1.7-.81 3.474.965 2.665 2.666l-5.552 11.659c-.759 1.593-3.059 1.495-3.679-.158L9.32 15.851a2 2 0 0 0-1.17-1.17l-2.74-1.027c-1.652-.62-1.75-2.92-.157-3.679"/>`), g.Group(children),
	)
}

func LocationPin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="3" d="M12 11h.01v.01H12z"/><path stroke-width="2" d="m12 22l5.5-5.5a7.778 7.778 0 1 0-11 0z"/></g>`), g.Group(children),
	)
}

func LocationPinBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="3.75" d="M12 11h.01v.01H12z"/><path stroke-width="2.5" d="m12 22l5.5-5.5a7.778 7.778 0 1 0-11 0z"/></g>`), g.Group(children),
	)
}

func LocationPinDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M17.5 16.5L12 22l-5.5-5.5a7.778 7.778 0 1 1 11 0M12 12a1 1 0 1 0 0-2a1 1 0 0 0 0 2" clip-rule="evenodd" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 11h.01v.01H12z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m12 22l5.5-5.5a7.778 7.778 0 1 0-11 0z"/></g>`), g.Group(children),
	)
}

func LocationPinFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18.207 4.793A8.778 8.778 0 0 0 5.793 17.207l5.5 5.5a1 1 0 0 0 1.414 0l5.5-5.5a8.78 8.78 0 0 0 0-12.414M10.5 11A1.5 1.5 0 0 1 12 9.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LocationPinLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="2.25" d="M12 11h.01v.01H12z"/><path stroke-width="1.5" d="m12 22l5.5-5.5a7.778 7.778 0 1 0-11 0z"/></g>`), g.Group(children),
	)
}

func LocationPinOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path d="m4 4l16 16"/><path stroke-linejoin="round" d="M6.023 6.022A7.78 7.78 0 0 0 6.5 16.5L12 22l5-5M9.344 3.687a7.78 7.78 0 0 1 9.969 9.969"/></g>`), g.Group(children),
	)
}

func LocationPinOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><path d="m4 4l16 16"/><path stroke-linejoin="round" d="M6.023 6.022A7.78 7.78 0 0 0 6.5 16.5L12 22l5-5M9.344 3.687a7.78 7.78 0 0 1 9.969 9.969"/></g>`), g.Group(children),
	)
}

func LocationPinOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6.5 16.5L12 22l5-5L6.022 6.022A7.78 7.78 0 0 0 6.5 16.5" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6.022 6.022A7.78 7.78 0 0 0 6.5 16.5L12 22l5-5M9.344 3.687a7.78 7.78 0 0 1 9.969 9.969"/></g>`), g.Group(children),
	)
}

func LocationPinOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.685 4.627a6.78 6.78 0 0 1 7.108 1.58a6.78 6.78 0 0 1 1.58 7.108a1 1 0 1 0 1.88.681a8.78 8.78 0 0 0-2.046-9.203a8.78 8.78 0 0 0-9.204-2.046a1 1 0 1 0 .682 1.88m-4.982 1.49l-1.41-1.41a1 1 0 0 1 1.414-1.414L6.72 5.306l.02.02l10.96 10.96l.007.007l.007.006l2.993 2.994a1 1 0 0 1-1.414 1.414L17 18.414l-4.293 4.293a1 1 0 0 1-1.414 0l-5.5-5.5a8.78 8.78 0 0 1-1.09-11.09" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LocationPinOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="m4 4l16 16"/><path stroke-linejoin="round" d="M6.022 6.022A7.78 7.78 0 0 0 6.5 16.5L12 22l5-5M9.344 3.687a7.78 7.78 0 0 1 9.969 9.969"/></g>`), g.Group(children),
	)
}

func LocationPinOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path d="m4 4l16 16"/><path stroke-linejoin="round" d="M6.022 6.022A7.78 7.78 0 0 0 6.5 16.5L12 22l5-5M9.344 3.687a7.78 7.78 0 0 1 9.969 9.969"/></g>`), g.Group(children),
	)
}

func LocationPinThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="1.5" d="M12 11h.01v.01H12z"/><path d="m12 22l5.5-5.5a7.778 7.778 0 1 0-11 0z"/></g>`), g.Group(children),
	)
}

func LocationThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.252 9.975l11.66-5.552c1.7-.81 3.474.965 2.665 2.666l-5.552 11.659c-.759 1.593-3.059 1.495-3.679-.158L9.32 15.851a2 2 0 0 0-1.17-1.17l-2.74-1.027c-1.652-.62-1.75-2.92-.157-3.679"/>`), g.Group(children),
	)
}

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-width="2" d="M8 10V7a4 4 0 1 1 8 0v3"/><path stroke-linejoin="round" stroke-width="2" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linejoin="round" stroke-width="3" d="M14.5 15.5h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func LockBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-width="2.5" d="M8 10V7a4 4 0 1 1 8 0v3"/><path stroke-linejoin="round" stroke-width="2.5" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linejoin="round" stroke-width="3.75" d="M14.5 15.5h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func LockDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M8 10V7a4 4 0 1 1 8 0v3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M14.5 15.5h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func LockFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7a3 3 0 1 1 6 0v2H9zM7 9V7a5 5 0 0 1 10 0v2h2a1 1 0 0 1 1 1v9a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3v-9a1 1 0 0 1 1-1zm6 6.5a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5h-.01a1.5 1.5 0 0 1-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LockLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-width="1.5" d="M8 10V7a4 4 0 1 1 8 0v3"/><path stroke-linejoin="round" stroke-width="1.5" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linejoin="round" stroke-width="2.25" d="M14.5 15.5h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func LockOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-width="2" d="M8 10V7a4 4 0 0 1 7.874-1"/><path stroke-linejoin="round" stroke-width="2" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linejoin="round" stroke-width="3" d="M14.5 15.5h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func LockOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-width="2.5" d="M8 10V7a4 4 0 0 1 7.874-1"/><path stroke-linejoin="round" stroke-width="2.5" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linejoin="round" stroke-width="3.75" d="M14.5 15.5h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func LockOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M8 10V7a4 4 0 0 1 7.874-1"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M14.5 15.5h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func LockOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 7a3 3 0 0 1 5.906-.75a1 1 0 0 0 1.936-.5A5.002 5.002 0 0 0 7 7v2H5a1 1 0 0 0-1 1v9a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3v-9a1 1 0 0 0-1-1H9zm4 8.5a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5h-.01a1.5 1.5 0 0 1-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LockOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-width="1.5" d="M8 10V7a4 4 0 0 1 7.874-1"/><path stroke-linejoin="round" stroke-width="1.5" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linejoin="round" stroke-width="2.25" d="M14.5 15.5h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func LockOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" d="M8 10V7a4 4 0 0 1 7.874-1"/><path stroke-linejoin="round" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linejoin="round" stroke-width="1.5" d="M14.5 15.5h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func LockThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" d="M8 10V7a4 4 0 1 1 8 0v3"/><path stroke-linejoin="round" d="M5 10h14v9a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2z"/><path stroke-linejoin="round" stroke-width="1.5" d="M14.5 15.5h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func MenuBurgerHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6h18M3 12h18M3 18h18"/>`), g.Group(children),
	)
}

func MenuBurgerHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M3 6h18M3 12h18M3 18h18"/>`), g.Group(children),
	)
}

func MenuBurgerHorizontalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 6.001h18m-18 6h18m-18 6h18"/>`), g.Group(children),
	)
}

func MenuBurgerHorizontalFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 5a1 1 0 0 0 0 2h17a1 1 0 1 0 0-2zm-1 7a1 1 0 0 1 1-1h17a1 1 0 1 1 0 2h-17a1 1 0 0 1-1-1m0 6.001a1 1 0 0 1 1-1h17a1 1 0 1 1 0 2h-17a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MenuBurgerHorizontalLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 6.001h18m-18 6h18m-18 6h18"/>`), g.Group(children),
	)
}

func MenuBurgerHorizontalThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 6.001h18m-18 6h18m-18 6h18"/>`), g.Group(children),
	)
}

func MenuBurgerVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3v18M12 3v18M6 3v18"/>`), g.Group(children),
	)
}

func MenuBurgerVerticalBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M18 3v18M12 3v18M6 3v18"/>`), g.Group(children),
	)
}

func MenuBurgerVerticalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M18 3.001v18m-6-18v18m-6-18v18"/>`), g.Group(children),
	)
}

func MenuBurgerVerticalFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 3.5a1 1 0 1 0-2 0v17a1 1 0 1 0 2 0zm-7-1a1 1 0 0 1 1 1v17a1 1 0 1 1-2 0v-17a1 1 0 0 1 1-1m-6 0a1 1 0 0 1 1 1v17a1 1 0 1 1-2 0v-17a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MenuBurgerVerticalLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 3.001v18m-6-18v18m-6-18v18"/>`), g.Group(children),
	)
}

func MenuBurgerVerticalThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18 3.001v18m-6-18v18m-6-18v18"/>`), g.Group(children),
	)
}

func MenuKebabHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12.01 12v.01H12V12zm7 0v.01H19V12zm-14 0v.01H5V12z"/>`), g.Group(children),
	)
}

func MenuKebabHorizontalBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="3.75" d="M12.01 12v.01H12V12zm7 0v.01H19V12zm-14 0v.01H5V12z"/>`), g.Group(children),
	)
}

func MenuKebabHorizontalCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M12.01 12v.01H12V12zm4.5 0v.01h-.01V12zm-9 0v.01H7.5V12z"/></g>`), g.Group(children),
	)
}

func MenuKebabHorizontalCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M12.01 12v.01H12V12zm4.5 0v.01h-.01V12zm-9 0v.01H7.5V12z"/></g>`), g.Group(children),
	)
}

func MenuKebabHorizontalCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M7.51 12v.01H7.5V12zm4.5 0v.01H12V12zm4.5 0v.01h-.01V12z"/></g>`), g.Group(children),
	)
}

func MenuKebabHorizontalCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7.01 0a1.5 1.5 0 0 0-1.5-1.5H7.5A1.5 1.5 0 0 0 6 12v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5zm4.5 0a1.5 1.5 0 0 0-1.5-1.5H12a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5zm3-1.5a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5h-.01a1.5 1.5 0 0 1-1.5-1.5V12a1.5 1.5 0 0 1 1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MenuKebabHorizontalCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M12.01 12v.01H12V12zm4.5 0v.01h-.01V12zm-9 0v.01H7.5V12z"/></g>`), g.Group(children),
	)
}

func MenuKebabHorizontalCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M12.01 12v.01H12V12zm4.5 0v.01h-.01V12zm-9 0v.01H7.5V12z"/></g>`), g.Group(children),
	)
}

func MenuKebabHorizontalFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.51 12a1.5 1.5 0 0 0-1.5-1.5H5A1.5 1.5 0 0 0 3.5 12v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5zm5.5-1.5a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5V12a1.5 1.5 0 0 1 1.5-1.5zm7 0a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H19a1.5 1.5 0 0 1-1.5-1.5V12a1.5 1.5 0 0 1 1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MenuKebabHorizontalLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.25" d="M12.01 12v.01H12V12zm7 0v.01H19V12zm-14 0v.01H5V12z"/>`), g.Group(children),
	)
}

func MenuKebabHorizontalSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3" d="M12.01 12v.01H12V12zm4.5 0v.01h-.01V12zm-9 0v.01H7.5V12z"/></g>`), g.Group(children),
	)
}

func MenuKebabHorizontalSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3.75" d="M12.01 12v.01H12V12zm4.5 0v.01h-.01V12zm-9 0v.01H7.5V12z"/></g>`), g.Group(children),
	)
}

func MenuKebabHorizontalSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M7.51 12v.01H7.5V12zm4.5 0v.01H12V12zm4.5 0v.01h-.01V12z"/></g>`), g.Group(children),
	)
}

func MenuKebabHorizontalSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zM9.01 12a1.5 1.5 0 0 0-1.5-1.5H7.5A1.5 1.5 0 0 0 6 12v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5zm4.5 0a1.5 1.5 0 0 0-1.5-1.5H12a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5zm3-1.5a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5h-.01a1.5 1.5 0 0 1-1.5-1.5V12a1.5 1.5 0 0 1 1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MenuKebabHorizontalSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="1.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="2.25" d="M12.01 12v.01H12V12zm4.5 0v.01h-.01V12zm-9 0v.01H7.5V12z"/></g>`), g.Group(children),
	)
}

func MenuKebabHorizontalSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="1.5" d="M12.01 12v.01H12V12zm4.5 0v.01h-.01V12zm-9 0v.01H7.5V12z"/></g>`), g.Group(children),
	)
}

func MenuKebabHorizontalThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M12.01 12v.01H12V12zm7 0v.01H19V12zm-14 0v.01H5V12z"/>`), g.Group(children),
	)
}

func MenuKebabVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 12h.01v.01H12zm0-7h.01v.01H12zm0 14h.01v.01H12z"/>`), g.Group(children),
	)
}

func MenuKebabVerticalBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="3.75" d="M12 12h.01v.01H12zm0-7h.01v.01H12zm0 14h.01v.01H12z"/>`), g.Group(children),
	)
}

func MenuKebabVerticalCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M12 12h.01v.01H12zm0-4.5h.01v.01H12zm0 9h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func MenuKebabVerticalCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M12 12h.01v.01H12zm0-4.5h.01v.01H12zm0 9h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func MenuKebabVerticalCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 16.5h.01v.01H12zm0-4.5h.01v.01H12zm0-4.5h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func MenuKebabVerticalCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m8.5-4.5A1.5 1.5 0 0 1 12 6h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5zm1.5 3a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V12a1.5 1.5 0 0 0-1.5-1.5zm0 4.5a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MenuKebabVerticalCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M12 12h.01v.01H12zm0-4.5h.01v.01H12zm0 9h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func MenuKebabVerticalCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M12 12h.01v.01H12zm0-4.5h.01v.01H12zm0 9h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func MenuKebabVerticalFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 5A1.5 1.5 0 0 1 12 3.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5zm0 7a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5zm1.5 5.5a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V19a1.5 1.5 0 0 0-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MenuKebabVerticalLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.25" d="M12 12h.01v.01H12zm0-7h.01v.01H12zm0 14h.01v.01H12z"/>`), g.Group(children),
	)
}

func MenuKebabVerticalSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3" d="M12 12h.01v.01H12zm0-4.5h.01v.01H12zm0 9h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func MenuKebabVerticalSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3.75" d="M12 12h.01v.01H12zm0-4.5h.01v.01H12zm0 9h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func MenuKebabVerticalSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 16.5h.01v.01H12zm0-4.5h.01v.01H12zm0-4.5h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func MenuKebabVerticalSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zM10.5 7.5A1.5 1.5 0 0 1 12 6h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5zm1.5 3a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V12a1.5 1.5 0 0 0-1.5-1.5zm0 4.5a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MenuKebabVerticalSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="1.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="2.25" d="M12 12h.01v.01H12zm0-4.5h.01v.01H12zm0 9h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func MenuKebabVerticalSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="1.5" d="M12 12h.01v.01H12zm0-4.5h.01v.01H12zm0 9h.01v.01H12z"/></g>`), g.Group(children),
	)
}

func MenuKebabVerticalThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M12 12h.01v.01H12zm0-7h.01v.01H12zm0 14h.01v.01H12z"/>`), g.Group(children),
	)
}

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><rect width="6" height="11" x="9" y="3" rx="3"/><path d="M19 11a7 7 0 1 1-14 0m7 7v3"/></g>`), g.Group(children),
	)
}

func MicrophoneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><rect width="6" height="11" x="9" y="3" rx="3"/><path d="M19 11a7 7 0 1 1-14 0m7 7v3"/></g>`), g.Group(children),
	)
}

func MicrophoneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="6" height="11" x="9" y="3" fill="currentColor" opacity=".16" rx="3"/><rect width="6" height="11" x="9" y="3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="3"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11a7 7 0 1 1-14 0m7 7v3"/></g>`), g.Group(children),
	)
}

func MicrophoneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2a4 4 0 0 0-4 4v5a4 4 0 0 0 8 0V6a4 4 0 0 0-4-4m-7 8a1 1 0 0 1 1 1a6 6 0 0 0 12 0a1 1 0 1 1 2 0a8 8 0 0 1-7 7.938V21a1 1 0 1 1-2 0v-2.062A8 8 0 0 1 4 11a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MicrophoneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="6" height="11" x="9" y="3" rx="3"/><path d="M19 11a7 7 0 1 1-14 0m7 7v3"/></g>`), g.Group(children),
	)
}

func MicrophoneOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M19 11c0 .71-.106 1.395-.302 2.04M12 18a7 7 0 0 1-7-7m7 7v3m0-3a6.97 6.97 0 0 0 4.425-1.576M9.714 4.057A3 3 0 0 1 15 6v3.343M9 9v2a3 3 0 0 0 4.562 2.562"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func MicrophoneOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><path stroke-linejoin="round" d="M19 11c0 .71-.106 1.395-.302 2.04M12 18a7 7 0 0 1-7-7m7 7v3m0-3a6.97 6.97 0 0 0 4.425-1.576M9.714 4.057A3 3 0 0 1 15 6v3.343M9 9v2a3 3 0 0 0 4.562 2.562"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func MicrophoneOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M9 11a3 3 0 0 0 4.562 2.562L9 9z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19 11c0 .71-.106 1.395-.302 2.04M12 18a7 7 0 0 1-7-7m7 7v3m0-3a6.97 6.97 0 0 0 4.425-1.576M9.714 4.057A3 3 0 0 1 15 6v3.343M9 9v2a3 3 0 0 0 4.562 2.562"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func MicrophoneOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.953 3.409A4 4 0 0 1 16 6v3.343a1 1 0 1 1-2 0V6a2 2 0 0 0-3.524-1.295a1 1 0 0 1-1.524-1.296m5.336 9.465l-.04-.04L9.72 8.306l-.026-.026l-4.987-4.987a1 1 0 0 0-1.414 1.414L8 9.414V11a4 4 0 0 0 5.351 3.766l1.51 1.51A6 6 0 0 1 6 11a1 1 0 1 0-2-.001a8 8 0 0 0 7 7.938V21a1 1 0 1 0 2 0v-2.062a7.96 7.96 0 0 0 3.32-1.204l2.973 2.973a1 1 0 0 0 1.414-1.414l-3.559-3.56l-.034-.033zM19 10a1 1 0 0 1 1 1c0 .81-.12 1.593-.346 2.332a1 1 0 1 1-1.913-.582c.168-.553.259-1.14.259-1.75a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MicrophoneOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M19 11c0 .71-.106 1.395-.302 2.04M12 18a7 7 0 0 1-7-7m7 7v3m0-3a6.97 6.97 0 0 0 4.425-1.576M9.714 4.057A3 3 0 0 1 15 6v3.343M9 9v2a3 3 0 0 0 4.562 2.562"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func MicrophoneOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" d="M19 11c0 .71-.106 1.395-.302 2.04M12 18a7 7 0 0 1-7-7m7 7v3m0-3a6.97 6.97 0 0 0 4.425-1.576M9.714 4.057A3 3 0 0 1 15 6v3.343M9 9v2a3 3 0 0 0 4.562 2.562"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func MicrophoneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><rect width="6" height="11" x="9" y="3" rx="3"/><path d="M19 11a7 7 0 1 1-14 0m7 7v3"/></g>`), g.Group(children),
	)
}

func ModeDark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m20.996 11.712l.999-.032a1 1 0 0 0-1.284-.927zm-8.707-8.707l.958.284a1 1 0 0 0-.927-1.284zm8.422 7.748A6 6 0 0 1 19 11v2a8 8 0 0 0 2.28-.33zM19 11a6 6 0 0 1-6-6h-2a8 8 0 0 0 8 8zm-6-6c0-.596.087-1.17.247-1.71l-1.917-.57A8 8 0 0 0 11 5zm-1-1q.13 0 .257.004l.063-1.999A10 10 0 0 0 12 2zm-8 8a8 8 0 0 1 8-8V2C6.477 2 2 6.477 2 12zm8 8a8 8 0 0 1-8-8H2c0 5.523 4.477 10 10 10zm8-8a8 8 0 0 1-8 8v2c5.523 0 10-4.477 10-10zm-.004-.257Q20 11.87 20 12h2q0-.16-.005-.32z"/>`), g.Group(children),
	)
}

func ModeDarkBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m20.996 11.712l1.249-.04a1.25 1.25 0 0 0-1.605-1.159zm-8.707-8.707l1.198.355a1.25 1.25 0 0 0-1.16-1.605zm8.35 7.508A5.8 5.8 0 0 1 19 10.75v2.5a8.3 8.3 0 0 0 2.351-.34zM19 10.75A5.75 5.75 0 0 1 13.25 5h-2.5A8.25 8.25 0 0 0 19 13.25zM13.25 5c0-.572.083-1.122.237-1.64l-2.397-.71A8.3 8.3 0 0 0 10.75 5zM12 4.25q.124 0 .25.004l.078-2.499A10 10 0 0 0 12 1.75zM4.25 12A7.75 7.75 0 0 1 12 4.25v-2.5C6.34 1.75 1.75 6.34 1.75 12zM12 19.75A7.75 7.75 0 0 1 4.25 12h-2.5c0 5.66 4.59 10.25 10.25 10.25zM19.75 12A7.75 7.75 0 0 1 12 19.75v2.5c5.66 0 10.25-4.59 10.25-10.25zm-.004-.25q.004.126.004.25h2.5q0-.165-.005-.328z"/>`), g.Group(children),
	)
}

func ModeDarkDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 21a9 9 0 0 0 8.997-9.252a7 7 0 0 1-10.371-8.643A9 9 0 0 0 12 21" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 21a9 9 0 0 0 8.997-9.252a7 7 0 0 1-10.371-8.643A9 9 0 0 0 12 21"/></g>`), g.Group(children),
	)
}

func ModeDarkFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.535 3.519a1 1 0 0 0-1.061-1.403C5.675 2.852 2 6.996 2 12c0 5.523 4.477 10 10 10s10-4.477 10-10l-.004-.28a1 1 0 0 0-1.571-.793a6 6 0 0 1-8.89-7.409" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ModeDarkLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21a9 9 0 0 0 8.997-9.252a7 7 0 0 1-10.371-8.643A9 9 0 0 0 12 21"/>`), g.Group(children),
	)
}

func ModeDarkThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 21a9 9 0 0 0 8.997-9.252a7 7 0 0 1-10.371-8.643A9 9 0 0 0 12 21"/>`), g.Group(children),
	)
}

func ModeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="12" r="4" stroke-linejoin="round"/><path stroke-linecap="round" d="M20 12h1M3 12h1m8 8v1m0-18v1m5.657 13.657l.707.707M5.636 5.636l.707.707m0 11.314l-.707.707M18.364 5.636l-.707.707"/></g>`), g.Group(children),
	)
}

func ModeLightBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="12" cy="12" r="4" stroke-linejoin="round"/><path stroke-linecap="round" d="M20 12h1M3 12h1m8 8v1m0-18v1m5.657 13.657l.707.707M5.636 5.636l.707.707m0 11.314l-.707.707M18.364 5.636l-.707.707"/></g>`), g.Group(children),
	)
}

func ModeLightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="4" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="4" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M20 12h1M3 12h1m8 8v1m0-18v1m5.657 13.657l.707.707M5.636 5.636l.707.707m0 11.314l-.707.707M18.364 5.636l-.707.707"/></g>`), g.Group(children),
	)
}

func ModeLightFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0V3a1 1 0 0 1 1-1M2 12a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2H3a1 1 0 0 1-1-1m17 0a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1m-6 8a1 1 0 1 0-2 0v1a1 1 0 1 0 2 0zm5.364-3.05a1 1 0 1 0-1.414 1.414l.707.707a1 1 0 0 0 1.414-1.414zM4.929 4.929a1 1 0 0 1 1.414 0l.707.707A1 1 0 0 1 5.636 7.05l-.707-.707a1 1 0 0 1 0-1.414M7.05 18.364a1 1 0 1 0-1.414-1.414l-.707.707a1 1 0 1 0 1.414 1.414zM19.071 4.929a1 1 0 0 1 0 1.414l-.707.707a1 1 0 1 1-1.414-1.414l.707-.707a1 1 0 0 1 1.414 0M7 12a5 5 0 1 1 10 0a5 5 0 0 1-10 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ModeLightLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="4" stroke-linejoin="round"/><path stroke-linecap="round" d="M20 12h1M3 12h1m8 8v1m0-18v1m5.657 13.657l.707.707M5.636 5.636l.707.707m0 11.314l-.707.707M18.364 5.636l-.707.707"/></g>`), g.Group(children),
	)
}

func ModeLightThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><circle cx="12" cy="12" r="4" stroke-linejoin="round"/><path stroke-linecap="round" d="M20 12h1M3 12h1m8 8v1m0-18v1m5.657 13.657l.707.707M5.636 5.636l.707.707m0 11.314l-.707.707M18.364 5.636l-.707.707"/></g>`), g.Group(children),
	)
}

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M6 9a6 6 0 0 1 12 0v6a6 6 0 0 1-12 0z"/><path stroke-linecap="round" d="M12 7v4"/></g>`), g.Group(children),
	)
}

func MouseBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M6 9a6 6 0 0 1 12 0v6a6 6 0 0 1-12 0z"/><path stroke-linecap="round" d="M12 7v4"/></g>`), g.Group(children),
	)
}

func MouseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6 9a6 6 0 1 1 12 0v6a6 6 0 0 1-12 0z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M6 9a6 6 0 0 1 12 0v6a6 6 0 0 1-12 0z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7v4"/></g>`), g.Group(children),
	)
}

func MouseFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 9a7 7 0 0 1 14 0v6a7 7 0 1 1-14 0zm8-2a1 1 0 1 0-2 0v4a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MouseLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M6 9a6 6 0 0 1 12 0v6a6 6 0 0 1-12 0z"/><path stroke-linecap="round" d="M12 7v4"/></g>`), g.Group(children),
	)
}

func MouseThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M6 9a6 6 0 0 1 12 0v6a6 6 0 0 1-12 0z"/><path stroke-linecap="round" d="M12 7v4"/></g>`), g.Group(children),
	)
}

func Move(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 5l-2-2m0 0l-2 2m2-2v18m0 0l2-2m-2 2l-2-2m9-5l2-2m0 0l-2-2m2 2H3m0 0l2 2m-2-2l2-2"/>`), g.Group(children),
	)
}

func MoveBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m14 5l-2-2m0 0l-2 2m2-2v18m0 0l2-2m-2 2l-2-2m9-5l2-2m0 0l-2-2m2 2H3m0 0l2 2m-2-2l2-2"/>`), g.Group(children),
	)
}

func MoveDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14 5l-2-2m0 0l-2 2m2-2v18m0 0l2-2m-2 2l-2-2m9-5l2-2m0 0l-2-2m2 2H3m0 0l2 2m-2-2l2-2"/>`), g.Group(children),
	)
}

func MoveFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.707 2.293a1 1 0 0 0-1.414 0l-2 2A1 1 0 0 0 10 6h1v5H6v-1a1 1 0 0 0-1.707-.707l-2 2a1 1 0 0 0 0 1.414l2 2A1 1 0 0 0 6 14v-1h5v5h-1a1 1 0 0 0-.707 1.707l2 2a1 1 0 0 0 1.414 0l2-2A1 1 0 0 0 14 18h-1v-5h5v1a1 1 0 0 0 1.707.707l2-2a1 1 0 0 0 0-1.414l-2-2A1 1 0 0 0 18 10v1h-5V6h1a1 1 0 0 0 .707-1.707z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MoveLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14 5l-2-2m0 0l-2 2m2-2v18m0 0l2-2m-2 2l-2-2m9-5l2-2m0 0l-2-2m2 2H3m0 0l2 2m-2-2l2-2"/>`), g.Group(children),
	)
}

func MoveThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14 5l-2-2m0 0l-2 2m2-2v18m0 0l2-2m-2 2l-2-2m9-5l2-2m0 0l-2-2m2 2H3m0 0l2 2m-2-2l2-2"/>`), g.Group(children),
	)
}

func MusicAlbum(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="2"/><circle cx="18" cy="9" r="2"/><path d="M15.318 3.631a9 9 0 1 0 5.368 10.736M20 9V2l2 2"/></g>`), g.Group(children),
	)
}

func MusicAlbumBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="2"/><circle cx="18" cy="9" r="2"/><path d="M15.318 3.631a9 9 0 1 0 5.368 10.736M20 9V2l2 2"/></g>`), g.Group(children),
	)
}

func MusicAlbumDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="2" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="2" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="18" cy="9" r="2" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.318 3.631a9 9 0 1 0 5.368 10.736M20 9V2l2 2"/></g>`), g.Group(children),
	)
}

func MusicAlbumFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20.707 1.293A1 1 0 0 0 19 2v4.17A3 3 0 1 0 21 9V4.414l.293.293a1 1 0 1 0 1.414-1.414zM12 4a8 8 0 1 0 7.72 10.105a1 1 0 1 1 1.93.524C20.497 18.876 16.615 22 12 22C6.477 22 2 17.523 2 12S6.477 2 12 2c1.3 0 2.545.249 3.687.702a1 1 0 0 1-.738 1.859A8 8 0 0 0 12 4m0 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MusicAlbumLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="2"/><circle cx="18" cy="9" r="2"/><path d="M15.318 3.631a9 9 0 1 0 5.368 10.736M20 9V2l2 2"/></g>`), g.Group(children),
	)
}

func MusicAlbumThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="2"/><circle cx="18" cy="9" r="2"/><path d="M15.318 3.631a9 9 0 1 0 5.368 10.736M20 9V2l2 2"/></g>`), g.Group(children),
	)
}

func MusicArtist(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="12" cy="7" r="3"/><circle cx="18" cy="18" r="2" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M12.341 20H6a2 2 0 0 1-2-2a4 4 0 0 1 4-4h5.528M20 18v-7l2 2"/></g>`), g.Group(children),
	)
}

func MusicArtistBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="12" cy="7" r="3"/><circle cx="18" cy="18" r="2" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M12.341 20H6a2 2 0 0 1-2-2a4 4 0 0 1 4-4h5.528M20 18v-7l2 2"/></g>`), g.Group(children),
	)
}

func MusicArtistDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="7" r="3" fill="currentColor" opacity=".16"/><circle cx="12" cy="7" r="3" stroke="currentColor" stroke-width="2"/><circle cx="18" cy="18" r="2" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.341 20H6a2 2 0 0 1-2-2a4 4 0 0 1 4-4h5.528M20 18v-7l2 2"/></g>`), g.Group(children),
	)
}

func MusicArtistFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 3a4 4 0 1 0 0 8a4 4 0 0 0 0-8m3 15a3 3 0 0 1 4-2.83V11a1 1 0 0 1 1.707-.707l2 2a1 1 0 0 1-1.414 1.414L21 13.414V18a3 3 0 1 1-6 0M3 18a5 5 0 0 1 5-5h5.528a1 1 0 1 1 0 2H8a3 3 0 0 0-3 3a1 1 0 0 0 1 1h6.341a1 1 0 1 1 0 2H6a3 3 0 0 1-3-3" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MusicArtistLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="7" r="3"/><circle cx="18" cy="18" r="2" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M12.341 20H6a2 2 0 0 1-2-2a4 4 0 0 1 4-4h5.528M20 18v-7l2 2"/></g>`), g.Group(children),
	)
}

func MusicArtistThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><circle cx="12" cy="7" r="3"/><circle cx="18" cy="18" r="2" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linecap="round" stroke-linejoin="round" d="M12.341 20H6a2 2 0 0 1-2-2a4 4 0 0 1 4-4h5.528M20 18v-7l2 2"/></g>`), g.Group(children),
	)
}

func MusicOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="9" cy="18" r="3"/><path d="M12 18V3m0 0l6 1v4l-6-1z"/></g>`), g.Group(children),
	)
}

func MusicOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="9" cy="18" r="3"/><path d="M12 18V3m0 0l6 1v4l-6-1z"/></g>`), g.Group(children),
	)
}

func MusicOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="9" cy="18" r="3" fill="currentColor" opacity=".16"/><circle cx="9" cy="18" r="3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 18V3m0 0l6 1v4l-6-1z"/></g>`), g.Group(children),
	)
}

func MusicOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.972 2A.996.996 0 0 0 11 3v11.535A4 4 0 1 0 13 18V8.18l4.836.806A1 1 0 0 0 19 8V4a1 1 0 0 0-.836-.986l-5.981-.997A1 1 0 0 0 11.972 2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MusicOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="9" cy="18" r="3"/><path d="M12 18V3m0 0l6 1v4l-6-1z"/></g>`), g.Group(children),
	)
}

func MusicOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="9" cy="18" r="3"/><path d="M12 18V3m0 0l6 1v4l-6-1z"/></g>`), g.Group(children),
	)
}

func MusicTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="6" cy="18" r="3" stroke-linejoin="round"/><path stroke-linejoin="round" d="M9 18V5"/><path d="M21 3L9 5m12 2L9 9"/><circle cx="18" cy="16" r="3" stroke-linejoin="round"/><path stroke-linejoin="round" d="M21 16V3"/></g>`), g.Group(children),
	)
}

func MusicTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><circle cx="6" cy="18" r="3" stroke-linejoin="round"/><path stroke-linejoin="round" d="M9 18V5"/><path d="M21 3L9 5m12 2L9 9"/><circle cx="18" cy="16" r="3" stroke-linejoin="round"/><path stroke-linejoin="round" d="M21 16V3"/></g>`), g.Group(children),
	)
}

func MusicTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M9 9V5l12-2v4z" opacity=".16"/><circle cx="6" cy="18" r="3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 18V5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M21 3L9 5m12 2L9 9"/><circle cx="18" cy="16" r="3" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M21 16V3"/></g>`), g.Group(children),
	)
}

func MusicTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 7.019v-4a1 1 0 0 0-.354-.782a1 1 0 0 0-.829-.22L8.852 4.01A1 1 0 0 0 8 5.017v9.519A4 4 0 1 0 10 18V9.846L20 8.18v4.355A4 4 0 1 0 22 16z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func MusicTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><circle cx="6" cy="18" r="3" stroke-linejoin="round"/><path stroke-linejoin="round" d="M9 18V5"/><path d="M21 3L9 5m12 2L9 9"/><circle cx="18" cy="16" r="3" stroke-linejoin="round"/><path stroke-linejoin="round" d="M21 16V3"/></g>`), g.Group(children),
	)
}

func MusicTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><circle cx="6" cy="18" r="3" stroke-linejoin="round"/><path stroke-linejoin="round" d="M9 18V5"/><path d="M21 3L9 5m12 2L9 9"/><circle cx="18" cy="16" r="3" stroke-linejoin="round"/><path stroke-linejoin="round" d="M21 16V3"/></g>`), g.Group(children),
	)
}

func NeutralFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke-linejoin="round" stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2" d="M9 15h6"/></g>`), g.Group(children),
	)
}

func NeutralFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"/><path stroke-linejoin="round" stroke-width="3.75" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2.5" d="M9 15h6"/></g>`), g.Group(children),
	)
}

func NeutralFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M9 15h6"/></g>`), g.Group(children),
	)
}

func NeutralFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7-4a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5A1.5 1.5 0 0 0 9.01 8zm6 0a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5a1.5 1.5 0 0 0-1.5-1.5zm-7 7a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NeutralFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/><path stroke-linejoin="round" stroke-width="2.25" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/><path stroke-linecap="round" stroke-width="1.5" d="M9 15h6"/></g>`), g.Group(children),
	)
}

func NeutralFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-linejoin="round"/><path stroke-linejoin="round" stroke-width="1.5" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/><path stroke-linecap="round" d="M9 15h6"/></g>`), g.Group(children),
	)
}

func News(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 4v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8h-4"/><path d="M3 4h14v14a2 2 0 0 0 2 2v0M13 8H7m6 4H9"/></g>`), g.Group(children),
	)
}

func NewsBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M3 4v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8h-4"/><path d="M3 4h14v14a2 2 0 0 0 2 2v0M13 8H7m6 4H9"/></g>`), g.Group(children),
	)
}

func NewsDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 4h14v16H5a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8h-4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4h14v14a2 2 0 0 0 2 2v0M13 8H7m6 4H9"/></g>`), g.Group(children),
	)
}

func NewsFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 4v3h3a1 1 0 0 1 1 1v10a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V4a1 1 0 0 1 1-1h14a1 1 0 0 1 1 1m2 14a1 1 0 1 1-2 0V9h2zM6 8a1 1 0 0 1 1-1h6a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1m2 4a1 1 0 0 1 1-1h4a1 1 0 1 1 0 2H9a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NewsLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 4v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8h-4"/><path d="M3 4h14v14a2 2 0 0 0 2 2v0M13 8H7m6 4H9"/></g>`), g.Group(children),
	)
}

func NewsThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 4v14a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2V8h-4"/><path d="M3 4h14v14a2 2 0 0 0 2 2v0M13 8H7m6 4H9"/></g>`), g.Group(children),
	)
}

func Notification(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M6 19v-9a6 6 0 0 1 6-6v0a6 6 0 0 1 6 6v9M6 19h12M6 19H4m14 0h2m-9 3h2"/><circle cx="12" cy="3" r="1"/></g>`), g.Group(children),
	)
}

func NotificationBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6 19v-9a6 6 0 0 1 6-6v0a6 6 0 0 1 6 6v9M6 19h12M6 19H4m14 0h2m-9 3h2"/><circle cx="12" cy="3" r="1"/></g>`), g.Group(children),
	)
}

func NotificationDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6 10v9h12v-9a6 6 0 0 0-12 0" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 19v-9a6 6 0 0 1 6-6v0a6 6 0 0 1 6 6v9M6 19h12M6 19H4m14 0h2m-9 3h2"/><circle cx="12" cy="3" r="1" stroke="currentColor" stroke-width="2"/></g>`), g.Group(children),
	)
}

func NotificationFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1a2 2 0 0 0-1.98 2.284A7 7 0 0 0 5 10v8H4a1 1 0 1 0 0 2h16a1 1 0 1 0 0-2h-1v-8a7 7 0 0 0-5.02-6.716Q14 3.144 14 3a2 2 0 0 0-2-2m2 21a1 1 0 0 1-1 1h-2a1 1 0 1 1 0-2h2a1 1 0 0 1 1 1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NotificationLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M6 19v-9a6 6 0 0 1 6-6v0a6 6 0 0 1 6 6v9M6 19h12M6 19H4m14 0h2m-9 3h2"/><circle cx="12" cy="3" r="1"/></g>`), g.Group(children),
	)
}

func NotificationOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M11 22h2"/><circle cx="12" cy="3" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 19v-9c0-1.144.32-2.214.876-3.124M6 19h12M6 19H4m14 0v-1m0 1h1M9.999 4.342A6 6 0 0 1 18 10v2.343"/><path stroke-linecap="round" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func NotificationOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><path stroke-linecap="round" stroke-linejoin="round" d="M11 22h2"/><circle cx="12" cy="3" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 19v-9c0-1.144.32-2.214.876-3.124M6 19h12M6 19H4m14 0v-1m0 1h1M9.999 4.342A6 6 0 0 1 18 10v2.343"/><path stroke-linecap="round" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func NotificationOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6 10v9h12v-1L6.876 6.876A5.97 5.97 0 0 0 6 10" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 22h2"/><circle cx="12" cy="3" r="1" stroke="currentColor" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.999 4.342A6 6 0 0 1 18 10v2.343M6 19v-9c0-1.144.32-2.214.876-3.124M6 19h12M6 19H4m14 0v-1m0 1h1"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func NotificationOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 3a2 2 0 1 1 3.98.284A7 7 0 0 1 19 10v2.343a1 1 0 0 1-2 0V10a5 5 0 0 0-6.668-4.715a1 1 0 0 1-.667-1.886q.175-.062.355-.115A2 2 0 0 1 10 3M7.604 6.19l11.09 11.09l.026.026l.981.98l.012.013l.994.994a1 1 0 0 1-1.414 1.414L18.586 20H4a1 1 0 1 1 0-2h1v-8a7 7 0 0 1 .646-2.94L3.293 4.707a1 1 0 0 1 1.414-1.414L7.564 6.15zM13 23a1 1 0 1 0 0-2h-2a1 1 0 1 0 0 2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NotificationOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M11 22h2"/><circle cx="12" cy="3" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 19v-9c0-1.144.32-2.214.876-3.124M6 19h12M6 19H4m14 0v-1m0 1h1M9.999 4.342A6 6 0 0 1 18 10v2.343"/><path stroke-linecap="round" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func NotificationOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" d="M11 22h2"/><circle cx="12" cy="3" r="1"/><path stroke-linecap="round" stroke-linejoin="round" d="M6 19v-9c0-1.144.32-2.214.876-3.124M6 19h12M6 19H4m14 0v-1m0 1h1M9.999 4.342A6 6 0 0 1 18 10v2.343"/><path stroke-linecap="round" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func NotificationThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" d="M6 19v-9a6 6 0 0 1 6-6v0a6 6 0 0 1 6 6v9M6 19h12M6 19H4m14 0h2m-9 3h2"/><circle cx="12" cy="3" r="1"/></g>`), g.Group(children),
	)
}

func NumberEight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="15" r="5"/><circle cx="12" cy="7" r="3"/></g>`), g.Group(children),
	)
}

func NumberEightBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="15" r="5"/><circle cx="12" cy="7" r="3"/></g>`), g.Group(children),
	)
}

func NumberEightCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><circle cx="12" cy="14" r="3"/><circle cx="12" cy="9" r="2"/></g>`), g.Group(children),
	)
}

func NumberEightCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><circle cx="12" cy="14" r="3"/><circle cx="12" cy="9" r="2"/></g>`), g.Group(children),
	)
}

func NumberEightCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><circle cx="12" cy="14" r="3" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/><circle cx="12" cy="9" r="2" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/></g>`), g.Group(children),
	)
}

func NumberEightCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m10-4a1 1 0 1 0 0 2a1 1 0 0 0 0-2m3 1c0 .675-.223 1.299-.6 1.8a4 4 0 1 1-4.8 0A3 3 0 1 1 15 9m-5 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberEightCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><circle cx="12" cy="14" r="3"/><circle cx="12" cy="9" r="2"/></g>`), g.Group(children),
	)
}

func NumberEightCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><circle cx="12" cy="14" r="3"/><circle cx="12" cy="9" r="2"/></g>`), g.Group(children),
	)
}

func NumberEightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M17 15a5 5 0 1 1-10 0a5 5 0 0 1 10 0" opacity=".16"/><circle cx="12" cy="15" r="5" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/><circle cx="12" cy="7" r="3" fill="currentColor" opacity=".16"/><circle cx="12" cy="7" r="3" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/></g>`), g.Group(children),
	)
}

func NumberEightFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 7a2 2 0 1 1 4 0a2 2 0 0 1-4 0m4.906 2.75a4 4 0 1 0-5.811 0a6 6 0 1 0 5.811 0M12 11a4 4 0 1 0 0 8a4 4 0 0 0 0-8" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberEightLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="15" r="5"/><circle cx="12" cy="7" r="3"/></g>`), g.Group(children),
	)
}

func NumberEightSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><circle cx="12" cy="14" r="3"/><circle cx="12" cy="9" r="2"/></g>`), g.Group(children),
	)
}

func NumberEightSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><circle cx="12" cy="14" r="3"/><circle cx="12" cy="9" r="2"/></g>`), g.Group(children),
	)
}

func NumberEightSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><circle cx="12" cy="14" r="3" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/><circle cx="12" cy="9" r="2" stroke="currentColor" stroke-linejoin="round" stroke-width="2"/></g>`), g.Group(children),
	)
}

func NumberEightSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zM12 8a1 1 0 1 0 0 2a1 1 0 0 0 0-2m3 1c0 .675-.223 1.299-.6 1.8a4 4 0 1 1-4.8 0A3 3 0 1 1 15 9m-5 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberEightSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><circle cx="12" cy="14" r="3"/><circle cx="12" cy="9" r="2"/></g>`), g.Group(children),
	)
}

func NumberEightSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><circle cx="12" cy="14" r="3"/><circle cx="12" cy="9" r="2"/></g>`), g.Group(children),
	)
}

func NumberEightThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="15" r="5"/><circle cx="12" cy="7" r="3"/></g>`), g.Group(children),
	)
}

func NumberFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 19a5 5 0 1 0 0-8l1-7h7"/>`), g.Group(children),
	)
}

func NumberFiveBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M8 19a5 5 0 1 0 0-8l1-7h7"/>`), g.Group(children),
	)
}

func NumberFiveCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M9.5 16.236a3 3 0 1 0 0-4.472L10.5 7h4"/></g>`), g.Group(children),
	)
}

func NumberFiveCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 16.236a3 3 0 1 0 0-4.472L10.5 7h4"/></g>`), g.Group(children),
	)
}

func NumberFiveCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 16.236a3 3 0 1 0 0-4.472L10.5 7h4"/></g>`), g.Group(children),
	)
}

func NumberFiveCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m8.5-6a1 1 0 0 0-.979.795l-1 4.764a1 1 0 0 0 1.646.95a2 2 0 1 1 0 2.982a1 1 0 1 0-1.334 1.49a4 4 0 1 0 2.049-6.934L11.312 8H14.5a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberFiveCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 16.236a3 3 0 1 0 0-4.472L10.5 7h4"/></g>`), g.Group(children),
	)
}

func NumberFiveCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M9.5 16.236a3 3 0 1 0 0-4.472L10.5 7h4"/></g>`), g.Group(children),
	)
}

func NumberFiveFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.01 3.859A1 1 0 0 1 9 3h7a1 1 0 1 1 0 2H9.867L9.26 9.257A6 6 0 1 1 7.4 19.8a1 1 0 1 1 1.202-1.6a4 4 0 1 0 0-6.402a1 1 0 0 1-1.59-.94z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberFiveLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 19a5 5 0 1 0 0-8l1-7h7"/>`), g.Group(children),
	)
}

func NumberFiveSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9.5 16.236a3 3 0 1 0 0-4.472L10.5 7h4"/></g>`), g.Group(children),
	)
}

func NumberFiveSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9.5 16.236a3 3 0 1 0 0-4.472L10.5 7h4"/></g>`), g.Group(children),
	)
}

func NumberFiveSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 16.236a3 3 0 1 0 0-4.472L10.5 7h4"/></g>`), g.Group(children),
	)
}

func NumberFiveSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zM10.5 6a1 1 0 0 0-.979.795l-1 4.764a1 1 0 0 0 1.646.95a2 2 0 1 1 0 2.982a1 1 0 1 0-1.334 1.49a4 4 0 1 0 2.049-6.934L11.312 8H14.5a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberFiveSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9.5 16.236a3 3 0 1 0 0-4.472L10.5 7h4"/></g>`), g.Group(children),
	)
}

func NumberFiveSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9.5 16.236a3 3 0 1 0 0-4.472L10.5 7h4"/></g>`), g.Group(children),
	)
}

func NumberFiveThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 19a5 5 0 1 0 0-8l1-7h7"/>`), g.Group(children),
	)
}

func NumberFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 4l-1.522 7.608A2 2 0 0 0 10.44 14H16m0 0V8m0 6v6"/>`), g.Group(children),
	)
}

func NumberFourBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m10 4l-1.522 7.608A2 2 0 0 0 10.44 14H16m0 0V8m0 6v6"/>`), g.Group(children),
	)
}

func NumberFourCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="m10.5 7l-1.272 4.45A2 2 0 0 0 11.152 14H14.5m0 0v-4m0 4v3"/></g>`), g.Group(children),
	)
}

func NumberFourCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="m10.5 7l-1.272 4.45A2 2 0 0 0 11.152 14H14.5m0 0v-4m0 4v3"/></g>`), g.Group(children),
	)
}

func NumberFourCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.5 7l-1.272 4.45A2 2 0 0 0 11.152 14H14.5m0 0v-4m0 4v3"/></g>`), g.Group(children),
	)
}

func NumberFourCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m9.462-4.725a1 1 0 1 0-1.924-.55l-1.271 4.45A3 3 0 0 0 11.152 15H13.5v2a1 1 0 1 0 2 0v-7a1 1 0 1 0-2 0v3h-2.348a1 1 0 0 1-.962-1.275z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberFourCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="m10.5 7l-1.272 4.45A2 2 0 0 0 11.152 14H14.5m0 0v-4m0 4v3"/></g>`), g.Group(children),
	)
}

func NumberFourCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="m10.5 7l-1.272 4.45A2 2 0 0 0 11.152 14H14.5m0 0v-4m0 4v3"/></g>`), g.Group(children),
	)
}

func NumberFourFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.196 3.02a1 1 0 0 1 .785 1.176l-1.522 7.608a1 1 0 0 0 .98 1.196H15V8a1 1 0 1 1 2 0v12a1 1 0 1 1-2 0v-5h-4.56a3 3 0 0 1-2.942-3.588l1.521-7.608a1 1 0 0 1 1.177-.785" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberFourLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10 4l-1.522 7.608A2 2 0 0 0 10.44 14H16m0 0V8m0 6v6"/>`), g.Group(children),
	)
}

func NumberFourSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m10.5 7l-1.272 4.45A2 2 0 0 0 11.152 14H14.5m0 0v-4m0 4v3"/></g>`), g.Group(children),
	)
}

func NumberFourSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m10.5 7l-1.272 4.45A2 2 0 0 0 11.152 14H14.5m0 0v-4m0 4v3"/></g>`), g.Group(children),
	)
}

func NumberFourSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.5 7l-1.272 4.45A2 2 0 0 0 11.152 14H14.5m0 0v-4m0 4v3"/></g>`), g.Group(children),
	)
}

func NumberFourSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm8.462 3.274a1 1 0 1 0-1.924-.55l-1.271 4.45A3 3 0 0 0 11.152 15H13.5v2a1 1 0 1 0 2 0v-7a1 1 0 1 0-2 0v3h-2.348a1 1 0 0 1-.962-1.275z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberFourSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m10.5 7l-1.272 4.45A2 2 0 0 0 11.152 14H14.5m0 0v-4m0 4v3"/></g>`), g.Group(children),
	)
}

func NumberFourSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="m10.5 7l-1.272 4.45A2 2 0 0 0 11.152 14H14.5m0 0v-4m0 4v3"/></g>`), g.Group(children),
	)
}

func NumberFourThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10 4l-1.522 7.608A2 2 0 0 0 10.44 14H16m0 0V8m0 6v6"/>`), g.Group(children),
	)
}

func NumberNine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.153 19.468a1 1 0 1 0 1.694 1.064zm6.92-7.264a1 1 0 1 0-1.693-1.064zM8 9a4 4 0 0 1 4-4V3a6 6 0 0 0-6 6zm4-4a4 4 0 0 1 4 4h2a6 6 0 0 0-6-6zm4 4a4 4 0 0 1-4 4v2a6 6 0 0 0 6-6zm-4 4a4 4 0 0 1-4-4H6a6 6 0 0 0 6 6zm-.153 7.532l5.227-8.328l-1.694-1.064l-5.227 8.328z"/>`), g.Group(children),
	)
}

func NumberNineBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.941 19.336a1.25 1.25 0 1 0 2.118 1.329zm7.345-7a1.25 1.25 0 1 0-2.118-1.328zM8.25 9A3.75 3.75 0 0 1 12 5.25v-2.5A6.25 6.25 0 0 0 5.75 9zM12 5.25A3.75 3.75 0 0 1 15.75 9h2.5A6.25 6.25 0 0 0 12 2.75zM15.75 9A3.75 3.75 0 0 1 12 12.75v2.5A6.25 6.25 0 0 0 18.25 9zM12 12.75A3.75 3.75 0 0 1 8.25 9h-2.5A6.25 6.25 0 0 0 12 15.25zm.059 7.915l5.227-8.328l-2.118-1.33l-5.227 8.329z"/>`), g.Group(children),
	)
}

func NumberNineCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="currentColor" d="M10.164 16.452a1 1 0 0 0 1.672 1.096zm5.197-4.283a1 1 0 0 0-1.672-1.096zM10 10a2 2 0 0 1 2-2V6a4 4 0 0 0-4 4zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4zm2 2a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4zm-2 2a2 2 0 0 1-2-2H8a4 4 0 0 0 4 4zm-.164 5.548l3.525-5.38l-1.672-1.095l-3.525 5.379z"/></g>`), g.Group(children),
	)
}

func NumberNineCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"/><path fill="currentColor" d="M9.955 16.315a1.25 1.25 0 1 0 2.09 1.37zm5.615-4.01a1.25 1.25 0 1 0-2.09-1.37zM10.25 10c0-.966.784-1.75 1.75-1.75v-2.5A4.25 4.25 0 0 0 7.75 10zM12 8.25c.966 0 1.75.784 1.75 1.75h2.5A4.25 4.25 0 0 0 12 5.75zM13.75 10A1.75 1.75 0 0 1 12 11.75v2.5A4.25 4.25 0 0 0 16.25 10zM12 11.75A1.75 1.75 0 0 1 10.25 10h-2.5A4.25 4.25 0 0 0 12 14.25zm.046 5.935l3.524-5.38l-2.09-1.37l-3.525 5.38z"/></g>`), g.Group(children),
	)
}

func NumberNineCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="currentColor" d="M10.164 16.452a1 1 0 0 0 1.672 1.096zm5.197-4.283a1 1 0 0 0-1.672-1.096zM10 10a2 2 0 0 1 2-2V6a4 4 0 0 0-4 4zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4zm2 2a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4zm-2 2a2 2 0 0 1-2-2H8a4 4 0 0 0 4 4zm-.164 5.548l3.525-5.38l-1.672-1.095l-3.525 5.379z"/></g>`), g.Group(children),
	)
}

func NumberNineCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m8.164 4.452a1 1 0 0 0 1.672 1.096l3.489-5.323a4 4 0 1 0-3.55 1.769zM14 10a2 2 0 1 1-4 0a2 2 0 0 1 4 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberNineCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/><path fill="currentColor" d="M10.373 16.589a.75.75 0 0 0 1.254.822zm4.78-4.557a.75.75 0 0 0-1.255-.822zM9.75 10A2.25 2.25 0 0 1 12 7.75v-1.5A3.75 3.75 0 0 0 8.25 10zM12 7.75A2.25 2.25 0 0 1 14.25 10h1.5A3.75 3.75 0 0 0 12 6.25zM14.25 10A2.25 2.25 0 0 1 12 12.25v1.5A3.75 3.75 0 0 0 15.75 10zM12 12.25A2.25 2.25 0 0 1 9.75 10h-1.5A3.75 3.75 0 0 0 12 13.75zm-.373 5.161l3.525-5.38l-1.254-.821l-3.525 5.379z"/></g>`), g.Group(children),
	)
}

func NumberNineCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="currentColor" d="M10.582 16.726a.5.5 0 1 0 .836.548zm4.361-4.831a.5.5 0 1 0-.836-.548zM9.5 10A2.5 2.5 0 0 1 12 7.5v-1A3.5 3.5 0 0 0 8.5 10zM12 7.5a2.5 2.5 0 0 1 2.5 2.5h1A3.5 3.5 0 0 0 12 6.5zm2.5 2.5a2.5 2.5 0 0 1-2.5 2.5v1a3.5 3.5 0 0 0 3.5-3.5zM12 12.5A2.5 2.5 0 0 1 9.5 10h-1a3.5 3.5 0 0 0 3.5 3.5zm-.582 4.774l3.525-5.38l-.836-.547l-3.525 5.379z"/></g>`), g.Group(children),
	)
}

func NumberNineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M17 9A5 5 0 1 1 7 9a5 5 0 0 1 10 0" opacity=".16"/><path d="M10.153 19.468a1 1 0 1 0 1.694 1.064zm6.92-7.264a1 1 0 1 0-1.693-1.064zM8 9a4 4 0 0 1 4-4V3a6 6 0 0 0-6 6zm4-4a4 4 0 0 1 4 4h2a6 6 0 0 0-6-6zm4 4a4 4 0 0 1-4 4v2a6 6 0 0 0 6-6zm-4 4a4 4 0 0 1-4-4H6a6 6 0 0 0 6 6zm-.153 7.532l5.227-8.328l-1.694-1.064l-5.227 8.328z"/></g>`), g.Group(children),
	)
}

func NumberNineFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.468 20.847a1 1 0 0 1-.315-1.379l2.858-4.553a6 6 0 1 1 4.05-2.691l-5.214 8.308a1 1 0 0 1-1.379.315M12 13a4 4 0 1 0 0-8a4 4 0 0 0 0 8" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberNineLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.365 19.601a.75.75 0 0 0 1.27.798zm6.497-7.53a.75.75 0 1 0-1.27-.798zM7.75 9A4.25 4.25 0 0 1 12 4.75v-1.5A5.75 5.75 0 0 0 6.25 9zM12 4.75A4.25 4.25 0 0 1 16.25 9h1.5A5.75 5.75 0 0 0 12 3.25zM16.25 9A4.25 4.25 0 0 1 12 13.25v1.5A5.75 5.75 0 0 0 17.75 9zM12 13.25A4.25 4.25 0 0 1 7.75 9h-1.5A5.75 5.75 0 0 0 12 14.75zm-.365 7.149l5.227-8.328l-1.27-.798l-5.227 8.328z"/>`), g.Group(children),
	)
}

func NumberNineSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M10.164 16.452a1 1 0 0 0 1.672 1.096zm5.197-4.283a1 1 0 0 0-1.672-1.096zM10 10a2 2 0 0 1 2-2V6a4 4 0 0 0-4 4zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4zm2 2a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4zm-2 2a2 2 0 0 1-2-2H8a4 4 0 0 0 4 4zm-.164 5.548l3.525-5.38l-1.672-1.095l-3.525 5.379z"/></g>`), g.Group(children),
	)
}

func NumberNineSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M9.955 16.315a1.25 1.25 0 1 0 2.09 1.37zm5.615-4.01a1.25 1.25 0 1 0-2.09-1.37zM10.25 10c0-.966.784-1.75 1.75-1.75v-2.5A4.25 4.25 0 0 0 7.75 10zM12 8.25c.966 0 1.75.784 1.75 1.75h2.5A4.25 4.25 0 0 0 12 5.75zM13.75 10A1.75 1.75 0 0 1 12 11.75v2.5A4.25 4.25 0 0 0 16.25 10zM12 11.75A1.75 1.75 0 0 1 10.25 10h-2.5A4.25 4.25 0 0 0 12 14.25zm.046 5.935l3.524-5.38l-2.09-1.37l-3.525 5.38z"/></g>`), g.Group(children),
	)
}

func NumberNineSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M10.164 16.452a1 1 0 0 0 1.672 1.096zm5.197-4.283a1 1 0 0 0-1.672-1.096zM10 10a2 2 0 0 1 2-2V6a4 4 0 0 0-4 4zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4zm2 2a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4zm-2 2a2 2 0 0 1-2-2H8a4 4 0 0 0 4 4zm-.164 5.548l3.525-5.38l-1.672-1.095l-3.525 5.379z"/></g>`), g.Group(children),
	)
}

func NumberNineSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm7.164 12.45a1 1 0 0 0 1.672 1.097l3.489-5.323a4 4 0 1 0-3.55 1.769zM14 10a2 2 0 1 1-4 0a2 2 0 0 1 4 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberNineSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M10.373 16.589a.75.75 0 0 0 1.254.822zm4.78-4.557a.75.75 0 0 0-1.255-.822zM9.75 10A2.25 2.25 0 0 1 12 7.75v-1.5A3.75 3.75 0 0 0 8.25 10zM12 7.75A2.25 2.25 0 0 1 14.25 10h1.5A3.75 3.75 0 0 0 12 6.25zM14.25 10A2.25 2.25 0 0 1 12 12.25v1.5A3.75 3.75 0 0 0 15.75 10zM12 12.25A2.25 2.25 0 0 1 9.75 10h-1.5A3.75 3.75 0 0 0 12 13.75zm-.373 5.161l3.525-5.38l-1.254-.821l-3.525 5.379z"/></g>`), g.Group(children),
	)
}

func NumberNineSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M10.582 16.726a.5.5 0 1 0 .836.548zm4.361-4.831a.5.5 0 1 0-.836-.548zM9.5 10A2.5 2.5 0 0 1 12 7.5v-1A3.5 3.5 0 0 0 8.5 10zM12 7.5a2.5 2.5 0 0 1 2.5 2.5h1A3.5 3.5 0 0 0 12 6.5zm2.5 2.5a2.5 2.5 0 0 1-2.5 2.5v1a3.5 3.5 0 0 0 3.5-3.5zM12 12.5A2.5 2.5 0 0 1 9.5 10h-1a3.5 3.5 0 0 0 3.5 3.5zm-.582 4.774l3.525-5.38l-.836-.547l-3.525 5.379z"/></g>`), g.Group(children),
	)
}

func NumberNineThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.577 19.734a.5.5 0 1 0 .847.532zm6.073-7.796a.5.5 0 1 0-.847-.532zM7.5 9A4.5 4.5 0 0 1 12 4.5v-1A5.5 5.5 0 0 0 6.5 9zM12 4.5A4.5 4.5 0 0 1 16.5 9h1A5.5 5.5 0 0 0 12 3.5zM16.5 9a4.5 4.5 0 0 1-4.5 4.5v1A5.5 5.5 0 0 0 17.5 9zM12 13.5A4.5 4.5 0 0 1 7.5 9h-1a5.5 5.5 0 0 0 5.5 5.5zm-.576 6.766l5.226-8.328l-.847-.532l-5.227 8.328z"/>`), g.Group(children),
	)
}

func NumberOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 20V4L9 7"/>`), g.Group(children),
	)
}

func NumberOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 20V4L9 7"/>`), g.Group(children),
	)
}

func NumberOneCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12.5 17V7l-2 2"/><circle cx="12" cy="12" r="9"/></g>`), g.Group(children),
	)
}

func NumberOneCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M12.5 17V7l-2 2"/><circle cx="12" cy="12" r="9"/></g>`), g.Group(children),
	)
}

func NumberOneCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.5 17V7l-2 2"/></g>`), g.Group(children),
	)
}

func NumberOneCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m11.5-5a1 1 0 0 0-1.707-.707l-2 2a1 1 0 0 0 1.414 1.414l.293-.293V17a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberOneCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.5 17V7l-2 2"/><circle cx="12" cy="12" r="9"/></g>`), g.Group(children),
	)
}

func NumberOneCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12.5 17V7l-2 2"/><circle cx="12" cy="12" r="9"/></g>`), g.Group(children),
	)
}

func NumberOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.383 3.076A1 1 0 0 1 13 4v16a1 1 0 1 1-2 0V6.414L9.707 7.707a1 1 0 0 1-1.414-1.414l3-3a1 1 0 0 1 1.09-.217" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 20V4L9 7"/>`), g.Group(children),
	)
}

func NumberOneSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M12.5 17V7l-2 2"/></g>`), g.Group(children),
	)
}

func NumberOneSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M12.5 17V7l-2 2"/></g>`), g.Group(children),
	)
}

func NumberOneSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12.5 17V7l-2 2"/></g>`), g.Group(children),
	)
}

func NumberOneSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zM13.5 7a1 1 0 0 0-1.707-.707l-2 2a1 1 0 0 0 1.414 1.414l.293-.293V17a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberOneSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M12.5 17V7l-2 2"/></g>`), g.Group(children),
	)
}

func NumberOneSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M12.5 17V7l-2 2"/></g>`), g.Group(children),
	)
}

func NumberOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 20V4L9 7"/>`), g.Group(children),
	)
}

func NumberSeven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 4h8l-6 16"/>`), g.Group(children),
	)
}

func NumberSevenBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M8 4h8l-6 16"/>`), g.Group(children),
	)
}

func NumberSevenCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M9 7h6l-4 10"/></g>`), g.Group(children),
	)
}

func NumberSevenCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M9 7h6l-4 10"/></g>`), g.Group(children),
	)
}

func NumberSevenCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6l-4 10"/></g>`), g.Group(children),
	)
}

func NumberSevenCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7-6a1 1 0 0 0 0 2h4.523l-3.451 8.629a1 1 0 0 0 1.857.742l4-10A1 1 0 0 0 15 6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberSevenCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9 7h6l-4 10"/></g>`), g.Group(children),
	)
}

func NumberSevenCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M9 7h6l-4 10"/></g>`), g.Group(children),
	)
}

func NumberSevenFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 4a1 1 0 0 1 1-1h8a1 1 0 0 1 .936 1.351l-6 16a1 1 0 0 1-1.872-.702L14.557 5H8a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberSevenLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 4h8l-6 16"/>`), g.Group(children),
	)
}

func NumberSevenSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9 7h6l-4 10"/></g>`), g.Group(children),
	)
}

func NumberSevenSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9 7h6l-4 10"/></g>`), g.Group(children),
	)
}

func NumberSevenSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 7h6l-4 10"/></g>`), g.Group(children),
	)
}

func NumberSevenSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zM9 6a1 1 0 0 0 0 2h4.523l-3.451 8.629a1 1 0 0 0 1.857.742l4-10A1 1 0 0 0 15 6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberSevenSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9 7h6l-4 10"/></g>`), g.Group(children),
	)
}

func NumberSevenSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9 7h6l-4 10"/></g>`), g.Group(children),
	)
}

func NumberSevenThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 4h8l-6 16"/>`), g.Group(children),
	)
}

func NumberSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.847 4.532a1 1 0 1 0-1.694-1.064zm-6.92 7.264A1 1 0 0 0 8.62 12.86zM16 15a4 4 0 0 1-4 4v2a6 6 0 0 0 6-6zm-4 4a4 4 0 0 1-4-4H6a6 6 0 0 0 6 6zm-4-4a4 4 0 0 1 4-4V9a6 6 0 0 0-6 6zm4-4a4 4 0 0 1 4 4h2a6 6 0 0 0-6-6zm.153-7.532l-5.227 8.328L8.62 12.86l5.227-8.328z"/>`), g.Group(children),
	)
}

func NumberSixBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.059 4.665a1.25 1.25 0 0 0-2.118-1.33zm-7.345 6.998a1.25 1.25 0 0 0 2.118 1.33zM15.75 15A3.75 3.75 0 0 1 12 18.75v2.5A6.25 6.25 0 0 0 18.25 15zM12 18.75A3.75 3.75 0 0 1 8.25 15h-2.5A6.25 6.25 0 0 0 12 21.25zM8.25 15A3.75 3.75 0 0 1 12 11.25v-2.5A6.25 6.25 0 0 0 5.75 15zM12 11.25A3.75 3.75 0 0 1 15.75 15h2.5A6.25 6.25 0 0 0 12 8.75zm-.059-7.914l-5.227 8.327l2.118 1.33l5.227-8.328z"/>`), g.Group(children),
	)
}

func NumberSixCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="currentColor" d="M13.836 7.548a1 1 0 1 0-1.672-1.096zM8.64 11.831a1 1 0 0 0 1.672 1.096zM14 14a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4zm-2 2a2 2 0 0 1-2-2H8a4 4 0 0 0 4 4zm-2-2a2 2 0 0 1 2-2v-2a4 4 0 0 0-4 4zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4zm.164-5.548l-3.525 5.38l1.672 1.095l3.525-5.379z"/></g>`), g.Group(children),
	)
}

func NumberSixCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"/><path fill="currentColor" d="M14.046 7.685a1.25 1.25 0 0 0-2.092-1.37zm-5.617 4.01a1.25 1.25 0 1 0 2.092 1.37zM13.75 14A1.75 1.75 0 0 1 12 15.75v2.5A4.25 4.25 0 0 0 16.25 14zM12 15.75A1.75 1.75 0 0 1 10.25 14h-2.5A4.25 4.25 0 0 0 12 18.25zM10.25 14c0-.966.784-1.75 1.75-1.75v-2.5A4.25 4.25 0 0 0 7.75 14zM12 12.25c.966 0 1.75.784 1.75 1.75h2.5A4.25 4.25 0 0 0 12 9.75zm-.046-5.935l-3.525 5.38l2.092 1.37l3.524-5.38z"/></g>`), g.Group(children),
	)
}

func NumberSixCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path fill="currentColor" d="M13.836 7.548a1 1 0 1 0-1.672-1.096zM8.64 11.831a1 1 0 0 0 1.672 1.096zM14 14a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4zm-2 2a2 2 0 0 1-2-2H8a4 4 0 0 0 4 4zm-2-2a2 2 0 0 1 2-2v-2a4 4 0 0 0-4 4zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4zm.164-5.548l-3.525 5.38l1.672 1.095l3.525-5.379z"/></g>`), g.Group(children),
	)
}

func NumberSixCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m11.836-4.452a1 1 0 1 0-1.672-1.096l-3.489 5.323a4 4 0 1 0 3.55-1.769zM10 14a2 2 0 1 1 4 0a2 2 0 0 1-4 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberSixCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/><path fill="currentColor" d="M13.627 7.411a.75.75 0 0 0-1.254-.822zm-4.78 4.557a.75.75 0 0 0 1.255.822zM14.25 14A2.25 2.25 0 0 1 12 16.25v1.5A3.75 3.75 0 0 0 15.75 14zM12 16.25A2.25 2.25 0 0 1 9.75 14h-1.5A3.75 3.75 0 0 0 12 17.75zM9.75 14A2.25 2.25 0 0 1 12 11.75v-1.5A3.75 3.75 0 0 0 8.25 14zM12 11.75A2.25 2.25 0 0 1 14.25 14h1.5A3.75 3.75 0 0 0 12 10.25zm.373-5.161l-3.525 5.38l1.254.821l3.525-5.379z"/></g>`), g.Group(children),
	)
}

func NumberSixCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="currentColor" d="M13.418 7.274a.5.5 0 0 0-.836-.548zm-4.361 4.831a.5.5 0 1 0 .836.548zM14.5 14a2.5 2.5 0 0 1-2.5 2.5v1a3.5 3.5 0 0 0 3.5-3.5zM12 16.5A2.5 2.5 0 0 1 9.5 14h-1a3.5 3.5 0 0 0 3.5 3.5zM9.5 14a2.5 2.5 0 0 1 2.5-2.5v-1A3.5 3.5 0 0 0 8.5 14zm2.5-2.5a2.5 2.5 0 0 1 2.5 2.5h1a3.5 3.5 0 0 0-3.5-3.5zm.582-4.774l-3.525 5.38l.836.547l3.525-5.379z"/></g>`), g.Group(children),
	)
}

func NumberSixDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M17 15a5 5 0 1 1-10 0a5 5 0 0 1 10 0" opacity=".16"/><path d="M13.847 4.532a1 1 0 1 0-1.694-1.064zm-6.92 7.264A1 1 0 0 0 8.62 12.86zM16 15a4 4 0 0 1-4 4v2a6 6 0 0 0 6-6zm-4 4a4 4 0 0 1-4-4H6a6 6 0 0 0 6 6zm-4-4a4 4 0 0 1 4-4V9a6 6 0 0 0-6 6zm4-4a4 4 0 0 1 4 4h2a6 6 0 0 0-6-6zm.153-7.532l-5.227 8.328L8.62 12.86l5.227-8.328z"/></g>`), g.Group(children),
	)
}

func NumberSixFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.532 3.153a1 1 0 0 1 .315 1.379l-2.858 4.553a6 6 0 1 1-4.05 2.691l5.214-8.308a1 1 0 0 1 1.379-.315M12 11a4 4 0 1 0 0 8a4 4 0 0 0 0-8" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberSixLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.635 4.399a.75.75 0 0 0-1.27-.798zm-6.497 7.53a.75.75 0 0 0 1.27.798zM16.25 15A4.25 4.25 0 0 1 12 19.25v1.5A5.75 5.75 0 0 0 17.75 15zM12 19.25A4.25 4.25 0 0 1 7.75 15h-1.5A5.75 5.75 0 0 0 12 20.75zM7.75 15A4.25 4.25 0 0 1 12 10.75v-1.5A5.75 5.75 0 0 0 6.25 15zM12 10.75A4.25 4.25 0 0 1 16.25 15h1.5A5.75 5.75 0 0 0 12 9.25zm.365-7.149L7.138 11.93l1.27.798L13.635 4.4z"/>`), g.Group(children),
	)
}

func NumberSixSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M13.836 7.548a1 1 0 1 0-1.672-1.096zM8.64 11.831a1 1 0 0 0 1.672 1.096zM14 14a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4zm-2 2a2 2 0 0 1-2-2H8a4 4 0 0 0 4 4zm-2-2a2 2 0 0 1 2-2v-2a4 4 0 0 0-4 4zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4zm.164-5.548l-3.525 5.38l1.672 1.095l3.525-5.379z"/></g>`), g.Group(children),
	)
}

func NumberSixSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M14.046 7.685a1.25 1.25 0 0 0-2.092-1.37zm-5.617 4.01a1.25 1.25 0 1 0 2.092 1.37zM13.75 14A1.75 1.75 0 0 1 12 15.75v2.5A4.25 4.25 0 0 0 16.25 14zM12 15.75A1.75 1.75 0 0 1 10.25 14h-2.5A4.25 4.25 0 0 0 12 18.25zM10.25 14c0-.966.784-1.75 1.75-1.75v-2.5A4.25 4.25 0 0 0 7.75 14zM12 12.25c.966 0 1.75.784 1.75 1.75h2.5A4.25 4.25 0 0 0 12 9.75zm-.046-5.935l-3.525 5.38l2.092 1.37l3.524-5.38z"/></g>`), g.Group(children),
	)
}

func NumberSixSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M13.836 7.548a1 1 0 1 0-1.672-1.096zM8.64 11.831a1 1 0 0 0 1.672 1.096zM14 14a2 2 0 0 1-2 2v2a4 4 0 0 0 4-4zm-2 2a2 2 0 0 1-2-2H8a4 4 0 0 0 4 4zm-2-2a2 2 0 0 1 2-2v-2a4 4 0 0 0-4 4zm2-2a2 2 0 0 1 2 2h2a4 4 0 0 0-4-4zm.164-5.548l-3.525 5.38l1.672 1.095l3.525-5.379z"/></g>`), g.Group(children),
	)
}

func NumberSixSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm10.836 3.547a1 1 0 1 0-1.672-1.096l-3.489 5.323a4 4 0 1 0 3.55-1.769zM10 14a2 2 0 1 1 4 0a2 2 0 0 1-4 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberSixSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M13.627 7.411a.75.75 0 0 0-1.254-.822zm-4.78 4.557a.75.75 0 0 0 1.255.822zM14.25 14A2.25 2.25 0 0 1 12 16.25v1.5A3.75 3.75 0 0 0 15.75 14zM12 16.25A2.25 2.25 0 0 1 9.75 14h-1.5A3.75 3.75 0 0 0 12 17.75zM9.75 14A2.25 2.25 0 0 1 12 11.75v-1.5A3.75 3.75 0 0 0 8.25 14zM12 11.75A2.25 2.25 0 0 1 14.25 14h1.5A3.75 3.75 0 0 0 12 10.25zm.373-5.161l-3.525 5.38l1.254.821l3.525-5.379z"/></g>`), g.Group(children),
	)
}

func NumberSixSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path fill="currentColor" d="M13.418 7.274a.5.5 0 0 0-.836-.548zm-4.361 4.831a.5.5 0 1 0 .836.548zM14.5 14a2.5 2.5 0 0 1-2.5 2.5v1a3.5 3.5 0 0 0 3.5-3.5zM12 16.5A2.5 2.5 0 0 1 9.5 14h-1a3.5 3.5 0 0 0 3.5 3.5zM9.5 14a2.5 2.5 0 0 1 2.5-2.5v-1A3.5 3.5 0 0 0 8.5 14zm2.5-2.5a2.5 2.5 0 0 1 2.5 2.5h1a3.5 3.5 0 0 0-3.5-3.5zm.582-4.774l-3.525 5.38l.836.547l3.525-5.379z"/></g>`), g.Group(children),
	)
}

func NumberSixThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.424 4.266a.5.5 0 1 0-.848-.532zM7.35 12.062a.5.5 0 1 0 .847.532zM16.5 15a4.5 4.5 0 0 1-4.5 4.5v1a5.5 5.5 0 0 0 5.5-5.5zM12 19.5A4.5 4.5 0 0 1 7.5 15h-1a5.5 5.5 0 0 0 5.5 5.5zM7.5 15a4.5 4.5 0 0 1 4.5-4.5v-1A5.5 5.5 0 0 0 6.5 15zm4.5-4.5a4.5 4.5 0 0 1 4.5 4.5h1A5.5 5.5 0 0 0 12 9.5zm.576-6.766L7.35 12.062l.847.532l5.227-8.328z"/>`), g.Group(children),
	)
}

func NumberThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 19a5 5 0 1 0 3-9l5-6H8"/>`), g.Group(children),
	)
}

func NumberThreeBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M8 19a5 5 0 1 0 3-9l5-6H8"/>`), g.Group(children),
	)
}

func NumberThreeCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M10 16.236A3 3 0 1 0 12 11l3-4h-5"/></g>`), g.Group(children),
	)
}

func NumberThreeCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M10 16.236A3 3 0 1 0 12 11l3-4h-5"/></g>`), g.Group(children),
	)
}

func NumberThreeCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 16.236A3 3 0 1 0 12 11l3-4h-5"/></g>`), g.Group(children),
	)
}

func NumberThreeCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m8-6a1 1 0 0 0 0 2h3l-1.8 2.4A1 1 0 0 0 12 12a2 2 0 1 1-1.333 3.491a1 1 0 1 0-1.334 1.49a4 4 0 1 0 4.379-6.597L15.8 7.6A1 1 0 0 0 15 6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberThreeCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M10 16.236A3 3 0 1 0 12 11l3-4h-5"/></g>`), g.Group(children),
	)
}

func NumberThreeCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M10 16.236A3 3 0 1 0 12 11l3-4h-5"/></g>`), g.Group(children),
	)
}

func NumberThreeFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 4a1 1 0 0 1 1-1h8a1 1 0 0 1 .768 1.64l-3.884 4.662A6 6 0 1 1 7.4 19.8a1 1 0 1 1 1.2-1.6A4 4 0 1 0 11 11a1 1 0 0 1-.768-1.64L13.865 5H8a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberThreeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 19a5 5 0 1 0 3-9l5-6H8"/>`), g.Group(children),
	)
}

func NumberThreeSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 16.236A3 3 0 1 0 12 11l3-4h-5"/></g>`), g.Group(children),
	)
}

func NumberThreeSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 16.236A3 3 0 1 0 12 11l3-4h-5"/></g>`), g.Group(children),
	)
}

func NumberThreeSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 16.236A3 3 0 1 0 12 11l3-4h-5"/></g>`), g.Group(children),
	)
}

func NumberThreeSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zM10 6a1 1 0 0 0 0 2h3l-1.8 2.4A1 1 0 0 0 12 12a2 2 0 1 1-1.333 3.491a1 1 0 1 0-1.334 1.49a4 4 0 1 0 4.379-6.597L15.8 7.6A1 1 0 0 0 15 6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberThreeSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 16.236A3 3 0 1 0 12 11l3-4h-5"/></g>`), g.Group(children),
	)
}

func NumberThreeSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M10 16.236A3 3 0 1 0 12 11l3-4h-5"/></g>`), g.Group(children),
	)
}

func NumberThreeThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 19a5 5 0 1 0 3-9l5-6H8"/>`), g.Group(children),
	)
}

func NumberTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8a4 4 0 1 1 6.828 2.828l-5.656 5.657A4 4 0 0 0 8 19.314V20h8"/>`), g.Group(children),
	)
}

func NumberTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M8 8a4 4 0 1 1 6.828 2.828l-5.656 5.657A4 4 0 0 0 8 19.314V20h8"/>`), g.Group(children),
	)
}

func NumberTwoCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M9.5 9.5a2.5 2.5 0 1 1 4.268 1.768l-3.829 3.828a1.5 1.5 0 0 0-.439 1.06V17h5"/></g>`), g.Group(children),
	)
}

func NumberTwoCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 9.5a2.5 2.5 0 1 1 4.268 1.768l-3.829 3.828a1.5 1.5 0 0 0-.439 1.06V17h5"/></g>`), g.Group(children),
	)
}

func NumberTwoCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 9.5a2.5 2.5 0 1 1 4.268 1.768l-3.829 3.828a1.5 1.5 0 0 0-.439 1.06V17h5"/></g>`), g.Group(children),
	)
}

func NumberTwoCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m8.5-2.5a1.5 1.5 0 1 1 2.56 1.06l-3.828 3.83a2.5 2.5 0 0 0-.732 1.767V17a1 1 0 0 0 1 1h5a1 1 0 1 0 0-2h-3.975a.5.5 0 0 1 .121-.197l3.829-3.828A3.5 3.5 0 1 0 8.5 9.5a1 1 0 1 0 2 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberTwoCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M9.5 9.5a2.5 2.5 0 1 1 4.268 1.768l-3.829 3.828a1.5 1.5 0 0 0-.439 1.06V17h5"/></g>`), g.Group(children),
	)
}

func NumberTwoCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M9.5 9.5a2.5 2.5 0 1 1 4.268 1.768l-3.829 3.828a1.5 1.5 0 0 0-.439 1.06V17h5"/></g>`), g.Group(children),
	)
}

func NumberTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8 8a4 4 0 1 1 6.828 2.828l-5.656 5.657A4 4 0 0 0 8 19.314V20h8"/>`), g.Group(children),
	)
}

func NumberTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.121 5.879A3 3 0 0 0 9 8a1 1 0 0 1-2 0a5 5 0 1 1 8.535 3.535L9.88 17.193A3 3 0 0 0 9.016 19H16a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1v-.686a5 5 0 0 1 1.464-3.536l5.657-5.657a3 3 0 0 0 0-4.242" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 8a4 4 0 1 1 6.828 2.828l-5.656 5.657A4 4 0 0 0 8 19.314V20h8"/>`), g.Group(children),
	)
}

func NumberTwoSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9.5 9.5a2.5 2.5 0 1 1 4.268 1.768l-3.829 3.828a1.5 1.5 0 0 0-.439 1.06V17h5"/></g>`), g.Group(children),
	)
}

func NumberTwoSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9.5 9.5a2.5 2.5 0 1 1 4.268 1.768l-3.829 3.828a1.5 1.5 0 0 0-.439 1.06V17h5"/></g>`), g.Group(children),
	)
}

func NumberTwoSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9.5 9.5a2.5 2.5 0 1 1 4.268 1.768l-3.829 3.828a1.5 1.5 0 0 0-.439 1.06V17h5"/></g>`), g.Group(children),
	)
}

func NumberTwoSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm7.94 4.438a1.5 1.5 0 0 1 2.12 2.122l-3.828 3.828a2.5 2.5 0 0 0-.732 1.768V17a1 1 0 0 0 1 1h5a1 1 0 1 0 0-2h-3.975a.5.5 0 0 1 .121-.197l3.829-3.828A3.5 3.5 0 1 0 8.5 9.5a1 1 0 1 0 2 0c0-.385.146-.768.44-1.06" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberTwoSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9.5 9.5a2.5 2.5 0 1 1 4.268 1.768l-3.829 3.828a1.5 1.5 0 0 0-.439 1.06V17h5"/></g>`), g.Group(children),
	)
}

func NumberTwoSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M9.5 9.5a2.5 2.5 0 1 1 4.268 1.768l-3.829 3.828a1.5 1.5 0 0 0-.439 1.06V17h5"/></g>`), g.Group(children),
	)
}

func NumberTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8 8a4 4 0 1 1 6.828 2.828l-5.656 5.657A4 4 0 0 0 8 19.314V20h8"/>`), g.Group(children),
	)
}

func NumberZero(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M12 20a5 5 0 0 0 5-5V9A5 5 0 0 0 7 9v6a5 5 0 0 0 5 5Z"/>`), g.Group(children),
	)
}

func NumberZeroBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M12 20a5 5 0 0 0 5-5V9A5 5 0 0 0 7 9v6a5 5 0 0 0 5 5Z"/>`), g.Group(children),
	)
}

func NumberZeroCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/><circle cx="12" cy="12" r="9" stroke-linecap="round"/></g>`), g.Group(children),
	)
}

func NumberZeroCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/><circle cx="12" cy="12" r="9" stroke-linecap="round"/></g>`), g.Group(children),
	)
}

func NumberZeroCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/></g>`), g.Group(children),
	)
}

func NumberZeroCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m8-1.875C10 8.913 10.934 8 12 8s2 .913 2 2.125v3.75C14 15.088 13.066 16 12 16s-2-.912-2-2.125zM12 6c-2.247 0-4 1.886-4 4.125v3.75C8 16.115 9.753 18 12 18s4-1.886 4-4.125v-3.75C16 7.885 14.247 6 12 6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberZeroCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/><circle cx="12" cy="12" r="9" stroke-linecap="round"/></g>`), g.Group(children),
	)
}

func NumberZeroCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/><circle cx="12" cy="12" r="9" stroke-linecap="round"/></g>`), g.Group(children),
	)
}

func NumberZeroDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 20a5 5 0 0 0 5-5V9A5 5 0 0 0 7 9v6a5 5 0 0 0 5 5" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M12 20a5 5 0 0 0 5-5V9A5 5 0 0 0 7 9v6a5 5 0 0 0 5 5Z"/></g>`), g.Group(children),
	)
}

func NumberZeroFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 5a4 4 0 0 0-4 4v6a4 4 0 0 0 8 0V9a4 4 0 0 0-4-4M6 9a6 6 0 1 1 12 0v6a6 6 0 0 1-12 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberZeroLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M12 20a5 5 0 0 0 5-5V9A5 5 0 0 0 7 9v6a5 5 0 0 0 5 5Z"/>`), g.Group(children),
	)
}

func NumberZeroSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/></g>`), g.Group(children),
	)
}

func NumberZeroSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/></g>`), g.Group(children),
	)
}

func NumberZeroSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/></g>`), g.Group(children),
	)
}

func NumberZeroSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm7 6.124C10 8.913 10.934 8 12 8s2 .913 2 2.125v3.75C14 15.088 13.066 16 12 16s-2-.912-2-2.125zM12 6c-2.247 0-4 1.886-4 4.125v3.75C8 16.115 9.753 18 12 18s4-1.886 4-4.125v-3.75C16 7.885 14.247 6 12 6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func NumberZeroSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/></g>`), g.Group(children),
	)
}

func NumberZeroSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path d="M12 17c1.657 0 3-1.4 3-3.125v-3.75C15 8.399 13.657 7 12 7s-3 1.4-3 3.125v3.75C9 15.601 10.343 17 12 17Z"/></g>`), g.Group(children),
	)
}

func NumberZeroThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M12 20a5 5 0 0 0 5-5V9A5 5 0 0 0 7 9v6a5 5 0 0 0 5 5Z"/>`), g.Group(children),
	)
}

func Options(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M4 8h9m4 0h3m-9 8h9M4 16h3"/><circle cx="9" cy="16" r="2"/><circle cx="15" cy="8" r="2"/></g>`), g.Group(children),
	)
}

func OptionsBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><path stroke-linecap="round" stroke-linejoin="round" d="M4 8h9m4 0h3m-9 8h9M4 16h3"/><circle cx="9" cy="16" r="2"/><circle cx="15" cy="8" r="2"/></g>`), g.Group(children),
	)
}

func OptionsDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="9" cy="16" r="2" fill="currentColor" opacity=".16"/><circle cx="15" cy="8" r="2" fill="currentColor" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 8h9m4 0h3m-9 8h9M4 16h3"/><circle cx="9" cy="16" r="2" stroke="currentColor" stroke-width="2"/><circle cx="15" cy="8" r="2" stroke="currentColor" stroke-width="2"/></g>`), g.Group(children),
	)
}

func OptionsFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 7h8.17a3.001 3.001 0 0 1 5.66 0H20a1 1 0 1 1 0 2h-2.17a3.001 3.001 0 0 1-5.66 0H4a1 1 0 0 1 0-2m0 8h2.17a3.001 3.001 0 0 1 5.66 0H20a1 1 0 1 1 0 2h-8.17a3.001 3.001 0 0 1-5.66 0H4a1 1 0 1 1 0-2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func OptionsLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M4 8h9m4 0h3m-9 8h9M4 16h3"/><circle cx="9" cy="16" r="2"/><circle cx="15" cy="8" r="2"/></g>`), g.Group(children),
	)
}

func OptionsThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" d="M4 8h9m4 0h3m-9 8h9M4 16h3"/><circle cx="9" cy="16" r="2"/><circle cx="15" cy="8" r="2"/></g>`), g.Group(children),
	)
}

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 3V2a1 1 0 0 0-1 1zm15.293 11.293l-.911-.412A1 1 0 0 0 17.586 15zM21 17l.707.707a1 1 0 0 0 0-1.414zm-4 4l-.707.707a1 1 0 0 0 1.414 0zm-2.707-2.707l.707-.707a1 1 0 0 0-1.12-.204zM11 2H3v2h8zm9 9a9 9 0 0 0-9-9v2a7 7 0 0 1 7 7zm-.796 3.705A9 9 0 0 0 20 11h-2a7 7 0 0 1-.618 2.88zM17.586 15l2.707 2.707l1.414-1.414L19 13.586zm2.707 1.293l-4 4l1.414 1.414l4-4zm-2.586 4L15 17.586L13.586 19l2.707 2.707zM11 20a9 9 0 0 0 3.705-.796l-.824-1.822A7 7 0 0 1 11 18zm-9-9a9 9 0 0 0 9 9v-2a7 7 0 0 1-7-7zm0-8v8h2V3z"/><circle cx="11" cy="11" r="2" fill="currentColor" transform="rotate(-180 11 11)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l8 8"/></g>`), g.Group(children),
	)
}

func PenBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 3V1.75c-.69 0-1.25.56-1.25 1.25zm15.293 11.293l-1.139-.515a1.25 1.25 0 0 0 .255 1.399zM21 17l.884.884a1.25 1.25 0 0 0 0-1.768zm-4 4l-.884.884a1.25 1.25 0 0 0 1.768 0zm-2.707-2.707l.884-.884a1.25 1.25 0 0 0-1.4-.255zM11 1.75H3v2.5h8zM20.25 11A9.25 9.25 0 0 0 11 1.75v2.5A6.75 6.75 0 0 1 17.75 11zm-.818 3.808A9.2 9.2 0 0 0 20.25 11h-2.5a6.7 6.7 0 0 1-.596 2.778zm-2.023.369l2.707 2.707l1.768-1.768l-2.707-2.707zm2.707.94l-4 4l1.768 1.767l4-4zm-2.232 4l-2.707-2.708l-1.768 1.768l2.707 2.707zM11 20.25a9.2 9.2 0 0 0 3.808-.818l-1.03-2.278A6.7 6.7 0 0 1 11 17.75zM1.75 11A9.25 9.25 0 0 0 11 20.25v-2.5A6.75 6.75 0 0 1 4.25 11zm0-8v8h2.5V3z"/><circle cx="11" cy="11" r="2" fill="currentColor" transform="rotate(-180 11 11)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m3 3l8 8"/></g>`), g.Group(children),
	)
}

func PenDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M3 3h8a8 8 0 0 1 7.293 11.293L21 17l-4 4l-2.707-2.707A8 8 0 0 1 3 11z" clip-rule="evenodd" opacity=".16"/><path fill="currentColor" d="M3 3V2a1 1 0 0 0-1 1zm15.293 11.293l-.911-.412a1 1 0 0 0 .204 1.12zM21 17l.707.707a1 1 0 0 0 0-1.414zm-4 4l-.707.707a1 1 0 0 0 1.414 0zm-2.707-2.707l.707-.707a1 1 0 0 0-1.12-.204zM11 2H3v2h8zm9 9a9 9 0 0 0-9-9v2a7 7 0 0 1 7 7zm-.796 3.705A9 9 0 0 0 20 11h-2a7 7 0 0 1-.618 2.88zM17.586 15l2.707 2.707l1.414-1.414L19 13.586zm2.707 1.293l-4 4l1.414 1.414l4-4zm-2.586 4L15 17.586L13.586 19l2.707 2.707zM11 20a9 9 0 0 0 3.705-.796l-.824-1.822A7 7 0 0 1 11 18zm-9-9a9 9 0 0 0 9 9v-2a7 7 0 0 1-7-7zm0-8v8h2V3z"/><circle cx="11" cy="11" r="2" fill="currentColor" transform="rotate(-180 11 11)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 3l8 8"/></g>`), g.Group(children),
	)
}

func PenFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 4.621a.5.5 0 0 1 .854-.353l6.01 6.01c.126.126.17.31.15.487a2 2 0 1 0 1.751-1.751a.59.59 0 0 1-.487-.15l-6.01-6.01A.5.5 0 0 1 4.62 2H11a9 9 0 0 1 8.468 12.054l2.24 2.239a1 1 0 0 1 0 1.414l-4 4a1 1 0 0 1-1.415 0l-2.239-2.239A9 9 0 0 1 2 11z"/>`), g.Group(children),
	)
}

func PenLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 3v-.75a.75.75 0 0 0-.75.75zm15.293 11.293l-.683-.31a.75.75 0 0 0 .153.84zM21 17l.53.53a.75.75 0 0 0 0-1.06zm-4 4l-.53.53a.75.75 0 0 0 1.06 0zm-2.707-2.707l.53-.53a.75.75 0 0 0-.84-.153zM11 2.25H3v1.5h8zM19.75 11A8.75 8.75 0 0 0 11 2.25v1.5A7.25 7.25 0 0 1 18.25 11zm-.774 3.602c.498-1.1.774-2.32.774-3.602h-1.5a7.2 7.2 0 0 1-.64 2.984zm-1.213.221l2.707 2.707l1.06-1.06l-2.707-2.707zm2.707 1.647l-4 4l1.06 1.06l4-4zm-2.94 4l-2.707-2.707l-1.06 1.06l2.707 2.707zM11 19.75a8.7 8.7 0 0 0 3.602-.774l-.618-1.366a7.2 7.2 0 0 1-2.984.64zM2.25 11A8.75 8.75 0 0 0 11 19.75v-1.5A7.25 7.25 0 0 1 3.75 11zm0-8v8h1.5V3z"/><circle cx="11" cy="11" r="1" stroke="currentColor" stroke-width="1.5" transform="rotate(-180 11 11)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 3l7 7"/></g>`), g.Group(children),
	)
}

func PenThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 3v-.5a.5.5 0 0 0-.5.5zm15.293 11.293l-.456-.206a.5.5 0 0 0 .102.56zM21 17l.354.354a.5.5 0 0 0 0-.708zm-4 4l-.354.354a.5.5 0 0 0 .708 0zm-2.707-2.707l.354-.354a.5.5 0 0 0-.56-.102zM11 2.5H3v1h8zm8.5 8.5A8.5 8.5 0 0 0 11 2.5v1a7.5 7.5 0 0 1 7.5 7.5zm-.751 3.499A8.5 8.5 0 0 0 19.5 11h-1a7.5 7.5 0 0 1-.663 3.087zm-.81.147l2.707 2.708l.708-.708l-2.707-2.707zm2.707 2l-4 4l.708.708l4-4zm-3.292 4l-2.707-2.707l-.708.707l2.707 2.708zM11 19.5a8.5 8.5 0 0 0 3.499-.751l-.412-.912A7.5 7.5 0 0 1 11 18.5zM2.5 11a8.5 8.5 0 0 0 8.5 8.5v-1A7.5 7.5 0 0 1 3.5 11zm0-8v8h1V3z"/><circle cx="11" cy="11" r="1" stroke="currentColor" transform="rotate(-180 11 11)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3 3l7 7"/></g>`), g.Group(children),
	)
}

func PensiveFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><circle cx="12" cy="12" r="9" stroke-linejoin="round"/><path d="M10 16h4"/><path stroke-linejoin="round" d="M14 10.5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`), g.Group(children),
	)
}

func PensiveFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><circle cx="12" cy="12" r="9" stroke-linejoin="round"/><path d="M10 16h4"/><path stroke-linejoin="round" d="M14 10.5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`), g.Group(children),
	)
}

func PensiveFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M10 16h4"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 10.5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`), g.Group(children),
	)
}

func PensiveFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m4.073-1.875a1 1 0 0 1 1.302-.552c.346.14.725.218 1.125.218s.779-.078 1.125-.218a1 1 0 1 1 .75 1.854a5 5 0 0 1-1.875.364a5 5 0 0 1-1.875-.364a1 1 0 0 1-.552-1.302m8.302-.552a1 1 0 1 0-.75 1.854a5 5 0 0 0 3.75 0a1 1 0 1 0-.75-1.854c-.346.14-.725.218-1.125.218s-.779-.078-1.125-.218M10 15a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PensiveFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><circle cx="12" cy="12" r="9" stroke-linejoin="round"/><path d="M10 16h4"/><path stroke-linejoin="round" d="M14 10.5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`), g.Group(children),
	)
}

func PensiveFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><circle cx="12" cy="12" r="9" stroke-linejoin="round"/><path d="M10 16h4"/><path stroke-linejoin="round" d="M14 10.5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`), g.Group(children),
	)
}

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.506 7.96A16.03 16.03 0 0 1 7.96 19.506C5.819 20.051 4 18.21 4 16v-1c0-.552.449-.995.999-1.05a10 10 0 0 0 2.656-.639l1.519 1.52a12.05 12.05 0 0 0 5.657-5.657l-1.52-1.52a10 10 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.208 0 4.05 1.819 3.505 3.96"/>`), g.Group(children),
	)
}

func PhoneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M19.506 7.96A16.03 16.03 0 0 1 7.96 19.506C5.819 20.051 4 18.21 4 16v-1c0-.552.449-.995.999-1.05a10 10 0 0 0 2.656-.639l1.519 1.52a12.05 12.05 0 0 0 5.657-5.657l-1.52-1.52a10 10 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.208 0 4.05 1.819 3.505 3.96"/>`), g.Group(children),
	)
}

func PhoneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M19.506 7.96A16.03 16.03 0 0 1 7.96 19.506C5.819 20.051 4 18.21 4 16v-1c0-.552.449-.995.998-1.05a10 10 0 0 0 2.656-.639l1.52 1.52a12.05 12.05 0 0 0 5.657-5.657l-1.52-1.52a10 10 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.21 0 4.051 1.819 3.506 3.96" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.506 7.96A16.03 16.03 0 0 1 7.96 19.506C5.819 20.051 4 18.21 4 16v-1c0-.552.449-.995.998-1.05a10 10 0 0 0 2.656-.639l1.52 1.52a12.05 12.05 0 0 0 5.657-5.657l-1.52-1.52a10 10 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.21 0 4.051 1.819 3.506 3.96"/></g>`), g.Group(children),
	)
}

func PhoneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 3c2.669 0 5.226 2.258 4.475 5.206a17.03 17.03 0 0 1-12.269 12.27C5.258 21.225 3 18.668 3 16v-1c0-1.127.901-1.945 1.9-2.044a9 9 0 0 0 2.389-.575a1 1 0 0 1 1.072.223l1.003 1.002a11.06 11.06 0 0 0 4.242-4.242L12.604 8.36a1 1 0 0 1-.223-1.072a9 9 0 0 0 .575-2.39C13.055 3.902 13.873 3 15 3z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PhoneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.506 7.96A16.03 16.03 0 0 1 7.96 19.506C5.819 20.051 4 18.21 4 16v-1c0-.552.449-.995.998-1.05a10 10 0 0 0 2.656-.639l1.52 1.52a12.05 12.05 0 0 0 5.657-5.657l-1.52-1.52a10 10 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.21 0 4.051 1.819 3.506 3.96"/>`), g.Group(children),
	)
}

func PhoneOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="m14.83 9.174l-1.519-1.52a10 10 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.21 0 4.051 1.819 3.506 3.96a16 16 0 0 1-1.721 4.168m-5.3.357a12 12 0 0 1-3.311 2.346l-1.52-1.52a10 10 0 0 1-2.656.64C4.448 14.005 4 14.448 4 15v1c0 2.21 1.819 4.051 3.96 3.506a16 16 0 0 0 7.354-4.192"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func PhoneOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><path stroke-linejoin="round" d="m14.83 9.174l-1.519-1.52a10 10 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.21 0 4.051 1.819 3.506 3.96a16 16 0 0 1-1.721 4.168m-5.3.357a12 12 0 0 1-3.311 2.346l-1.52-1.52a10 10 0 0 1-2.656.64C4.448 14.005 4 14.448 4 15v1c0 2.21 1.819 4.051 3.96 3.506a16 16 0 0 0 7.354-4.192"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func PhoneOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 16v-1c0-.552.449-.995.998-1.05a10 10 0 0 0 2.656-.639l1.52 1.52a12 12 0 0 0 3.311-2.346l2.829 2.829c-2 2-4.527 3.472-7.354 4.192C5.819 20.051 4 18.21 4 16" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m14.83 9.174l-1.519-1.52a10 10 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.21 0 4.051 1.819 3.506 3.96a16 16 0 0 1-1.721 4.168m-5.3.357a12 12 0 0 1-3.311 2.346l-1.52-1.52a10 10 0 0 1-2.656.64C4.448 14.005 4 14.448 4 15v1c0 2.21 1.819 4.051 3.96 3.506a16 16 0 0 0 7.354-4.192"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func PhoneOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4.707 3.293a1 1 0 0 0-1.414 1.414l7.746 7.746q-.781.652-1.675 1.153L8.36 12.604a1 1 0 0 0-1.072-.223a9 9 0 0 1-2.39.575C3.902 13.055 3 13.873 3 15v1c0 2.669 2.258 5.226 5.206 4.475a16.96 16.96 0 0 0 7.087-3.768l4 4a1 1 0 0 0 1.414-1.414l-4.671-4.671l-.015-.016l-.015-.014l-2.799-2.799l-.03-.03zM12.956 4.9C13.055 3.9 13.873 3 15 3h1c2.669 0 5.226 2.258 4.475 5.206a17 17 0 0 1-1.83 4.43a1 1 0 0 1-1.721-1.017a15 15 0 0 0 1.613-3.906C18.877 6.379 17.75 5 16 5h-1l-.015.01a.14.14 0 0 0-.04.087a11 11 0 0 1-.48 2.298l1.073 1.072a1 1 0 0 1-1.415 1.414l-1.519-1.52a1 1 0 0 1-.223-1.072a9 9 0 0 0 .575-2.39" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PhoneOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m14.83 9.174l-1.519-1.52a10 10 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.21 0 4.051 1.819 3.506 3.96a16 16 0 0 1-1.721 4.168m-5.3.357a12 12 0 0 1-3.311 2.346l-1.52-1.52a10 10 0 0 1-2.656.64C4.448 14.005 4 14.448 4 15v1c0 2.21 1.819 4.051 3.96 3.506a16 16 0 0 0 7.354-4.192"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func PhoneOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" d="m14.83 9.174l-1.519-1.52a10 10 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.21 0 4.051 1.819 3.506 3.96a16 16 0 0 1-1.721 4.168m-5.3.357a12 12 0 0 1-3.311 2.346l-1.52-1.52a10 10 0 0 1-2.656.64C4.448 14.005 4 14.448 4 15v1c0 2.21 1.819 4.051 3.96 3.506a16 16 0 0 0 7.354-4.192"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func PhoneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.506 7.96A16.03 16.03 0 0 1 7.96 19.506C5.819 20.051 4 18.21 4 16v-1c0-.552.449-.995.998-1.05a10 10 0 0 0 2.656-.639l1.52 1.52a12.05 12.05 0 0 0 5.657-5.657l-1.52-1.52a10 10 0 0 0 .64-2.656C14.005 4.448 14.448 4 15 4h1c2.21 0 4.051 1.819 3.506 3.96"/>`), g.Group(children),
	)
}

func PlayCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path d="m14 12l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func PlayCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path d="m14 12l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func PlayCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" fill-opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m14 12l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func PlayCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2m3 10a1 1 0 0 1-.5.866l-3 1.732a1 1 0 0 1-1.5-.866v-3.464a1 1 0 0 1 1.5-.866l3 1.732A1 1 0 0 1 15 12" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlayCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path d="m14 12l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func PlayCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path d="m14 12l-3 1.732v-3.464z"/></g>`), g.Group(children),
	)
}

func PlayerEnd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M17 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/><path stroke-linecap="round" d="M20 5v14"/></g>`), g.Group(children),
	)
}

func PlayerEndBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M17 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/><path stroke-linecap="round" d="M20 5v14"/></g>`), g.Group(children),
	)
}

func PlayerEndDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M17 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M17 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 5v14"/></g>`), g.Group(children),
	)
}

func PlayerEndFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.5 9.402c2 1.155 2 4.041 0 5.196l-9 5.196c-2 1.155-4.5-.288-4.5-2.598V6.804c0-2.31 2.5-3.753 4.5-2.598zM21 5a1 1 0 1 0-2 0v14a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlayerEndLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M17 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/><path stroke-linecap="round" d="M20 5v14"/></g>`), g.Group(children),
	)
}

func PlayerEndThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M17 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/><path stroke-linecap="round" d="M20 5v14"/></g>`), g.Group(children),
	)
}

func PlayerNext(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3 6.804v10.392c0 1.54 1.667 2.502 3 1.732l3-1.732V6.804L6 5.072c-1.333-.77-3 .192-3 1.732Zm18 3.464c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/>`), g.Group(children),
	)
}

func PlayerNextBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M3 6.804v10.392c0 1.54 1.667 2.502 3 1.732l3-1.732V6.804L6 5.072c-1.333-.77-3 .192-3 1.732Zm18 3.464c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/>`), g.Group(children),
	)
}

func PlayerNextDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M21 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M3 6.804v10.392c0 1.54 1.667 2.502 3 1.732l3-1.732V6.804L6 5.072c-1.333-.77-3 .192-3 1.732Zm18 3.464c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/></g>`), g.Group(children),
	)
}

func PlayerNextFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.5 5.938a1 1 0 0 0-1.5.866v10.392a1 1 0 0 0 1.5.866L8 16.62V7.38zm2.898-.636L6.5 4.206l-.5.866l.5-.866C4.5 3.05 2 4.494 2 6.804v10.392c0 2.31 2.5 3.753 4.5 2.598l1.898-1.096c.785 1.355 2.587 1.971 4.102 1.096l9-5.196c2-1.155 2-4.041 0-5.196l-9-5.196c-1.515-.875-3.317-.259-4.102 1.096" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlayerNextLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M3 6.804v10.392c0 1.54 1.667 2.502 3 1.732l3-1.732V6.804L6 5.072c-1.333-.77-3 .192-3 1.732Zm18 3.464c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/>`), g.Group(children),
	)
}

func PlayerNextThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M3 6.804v10.392c0 1.54 1.667 2.502 3 1.732l3-1.732V6.804L6 5.072c-1.333-.77-3 .192-3 1.732Zm18 3.464c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/>`), g.Group(children),
	)
}

func PlayerPause(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M5 7a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zm9 0a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func PlayerPauseBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M5 7a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zm9 0a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func PlayerPauseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M5 7a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zm9 0a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M5 7a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zm9 0a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2z"/></g>`), g.Group(children),
	)
}

func PlayerPauseFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 7a3 3 0 0 1 3-3h1a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3zm12-3a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h1a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlayerPauseLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M5 7a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zm9 0a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func PlayerPauseThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M5 7a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2H7a2 2 0 0 1-2-2zm9 0a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v10a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func PlayerPlay(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M19 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/>`), g.Group(children),
	)
}

func PlayerPlayBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M19 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/>`), g.Group(children),
	)
}

func PlayerPlayDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M19 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M19 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/></g>`), g.Group(children),
	)
}

func PlayerPlayFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19.5 14.598c2-1.155 2-4.041 0-5.196l-9-5.196C8.5 3.05 6 4.494 6 6.804v10.392c0 2.31 2.5 3.753 4.5 2.598z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlayerPlayLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M19 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/>`), g.Group(children),
	)
}

func PlayerPlayThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M19 10.268c1.333.77 1.333 2.694 0 3.464l-9 5.196c-1.333.77-3-.192-3-1.732V6.804c0-1.54 1.667-2.502 3-1.732z"/>`), g.Group(children),
	)
}

func PlayerPrevious(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M21 17.196V6.804c0-1.54-1.667-2.502-3-1.732l-3 1.732v10.392l3 1.732c1.333.77 3-.192 3-1.732ZM3 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z"/>`), g.Group(children),
	)
}

func PlayerPreviousBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M21 17.196V6.804c0-1.54-1.667-2.502-3-1.732l-3 1.732v10.392l3 1.732c1.333.77 3-.192 3-1.732ZM3 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z"/>`), g.Group(children),
	)
}

func PlayerPreviousDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M21 17.196V6.804c0-1.54-1.667-2.502-3-1.732l-3 1.732v10.392l3 1.732c1.333.77 3-.192 3-1.732ZM3 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z"/></g>`), g.Group(children),
	)
}

func PlayerPreviousFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.5 9.402c-2 1.155-2 4.041 0 5.196l9 5.196c1.515.875 3.317.259 4.102-1.096l1.898 1.096c2 1.155 4.5-.288 4.5-2.598V6.804c0-2.31-2.5-3.753-4.5-2.598l-1.898 1.096c-.785-1.355-2.587-1.971-4.102-1.096zM16 7.382v9.237l2.5 1.443a1 1 0 0 0 1.5-.866V6.804a1 1 0 0 0-1.5-.866z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlayerPreviousLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M21 17.196V6.804c0-1.54-1.667-2.502-3-1.732l-3 1.732v10.392l3 1.732c1.333.77 3-.192 3-1.732ZM3 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z"/>`), g.Group(children),
	)
}

func PlayerPreviousThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M21 17.196V6.804c0-1.54-1.667-2.502-3-1.732l-3 1.732v10.392l3 1.732c1.333.77 3-.192 3-1.732ZM3 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z"/>`), g.Group(children),
	)
}

func PlayerStart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M7 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z"/><path stroke-linecap="round" d="M4 19V5"/></g>`), g.Group(children),
	)
}

func PlayerStartBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M7 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z"/><path stroke-linecap="round" d="M4 19V5"/></g>`), g.Group(children),
	)
}

func PlayerStartDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M7 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 19V5"/></g>`), g.Group(children),
	)
}

func PlayerStartFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m15.5 19.794l-9-5.196c-2-1.155-2-4.041 0-5.196l9-5.196c2-1.155 4.5.288 4.5 2.598v10.392c0 2.31-2.5 3.753-4.5 2.598M3 19a1 1 0 1 0 2 0V5a1 1 0 1 0-2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlayerStartLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M7 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z"/><path stroke-linecap="round" d="M4 19V5"/></g>`), g.Group(children),
	)
}

func PlayerStartThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M7 13.732c-1.333-.77-1.333-2.694 0-3.464l9-5.196c1.333-.77 3 .192 3 1.732v10.392c0 1.54-1.667 2.502-3 1.732z"/><path stroke-linecap="round" d="M4 19V5"/></g>`), g.Group(children),
	)
}

func PlayerStop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="14" height="14" x="5" y="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="2"/>`), g.Group(children),
	)
}

func PlayerStopBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="14" height="14" x="5" y="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" rx="2"/>`), g.Group(children),
	)
}

func PlayerStopDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="14" height="14" x="5" y="5" fill="currentColor" opacity=".16" rx="2"/><rect width="14" height="14" x="5" y="5" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" rx="2"/></g>`), g.Group(children),
	)
}

func PlayerStopFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 4a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h10a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlayerStopLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="14" height="14" x="5" y="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="2"/>`), g.Group(children),
	)
}

func PlayerStopThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="14" height="14" x="5" y="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`), g.Group(children),
	)
}

func Playlist(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M6 6L3 7.732V4.268z"/><path stroke-linecap="round" d="M3 12h18M10 6h11M3 18h18"/></g>`), g.Group(children),
	)
}

func PlaylistBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M6 6L3 7.732V4.268z"/><path stroke-linecap="round" d="M3 12h18M10 6h11M3 18h18"/></g>`), g.Group(children),
	)
}

func PlaylistFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 6a1 1 0 0 1-.5.866l-3 1.732A1 1 0 0 1 2 7.732V4.268a1 1 0 0 1 1.5-.866l3 1.732A1 1 0 0 1 7 6m-4 5a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2zm6-5a1 1 0 0 1 1-1h11a1 1 0 1 1 0 2H10a1 1 0 0 1-1-1M3 17a1 1 0 1 0 0 2h18a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlaylistLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M6 6L3 7.732V4.268z"/><path stroke-linecap="round" d="M3 12h18M10 6h11M3 18h18"/></g>`), g.Group(children),
	)
}

func PlaylistRepeatList(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="m19 5l2 2m0 0l-2 2m2-2H7M5 19l-2-2m0 0l2-2m-2 2h14"/><path d="M3 11a4 4 0 0 1 4-4m14 6a4 4 0 0 1-4 4"/></g>`), g.Group(children),
	)
}

func PlaylistRepeatListBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><path stroke-linejoin="round" d="m19 5l2 2m0 0l-2 2m2-2H7M5 19l-2-2m0 0l2-2m-2 2h14"/><path d="M3 11a4 4 0 0 1 4-4m14 6a4 4 0 0 1-4 4"/></g>`), g.Group(children),
	)
}

func PlaylistRepeatListFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 5a1 1 0 0 1 1.707-.707l2 2a1 1 0 0 1 0 1.414l-2 2A1 1 0 0 1 18 9V8H7a3 3 0 0 0-3 3a1 1 0 1 1-2 0a5 5 0 0 1 5-5h11zM6 19a1 1 0 0 1-1.707.707l-2-2a1 1 0 0 1 0-1.414l2-2A1 1 0 0 1 6 15v1h11a3 3 0 0 0 3-3a1 1 0 1 1 2 0a5 5 0 0 1-4.998 5H6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlaylistRepeatListLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m19 5l2 2m0 0l-2 2m2-2H7M5 19l-2-2m0 0l2-2m-2 2h14"/><path d="M3 11a4 4 0 0 1 4-4m14 6a4 4 0 0 1-4 4"/></g>`), g.Group(children),
	)
}

func PlaylistRepeatListThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" d="m19 5l2 2m0 0l-2 2m2-2H7M5 19l-2-2m0 0l2-2m-2 2h14"/><path d="M3 11a4 4 0 0 1 4-4m14 6a4 4 0 0 1-4 4"/></g>`), g.Group(children),
	)
}

func PlaylistRepeatSong(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="M21 9V4l-2 1m-4 2H7M5 19l-2-2m0 0l2-2m-2 2h14"/><path d="M3 11a4 4 0 0 1 4-4m14 6a4 4 0 0 1-4 4"/></g>`), g.Group(children),
	)
}

func PlaylistRepeatSongBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><path stroke-linejoin="round" d="M21 9V4l-2 1m-4 2H7M5 19l-2-2m0 0l2-2m-2 2h14"/><path d="M3 11a4 4 0 0 1 4-4m14 6a4 4 0 0 1-4 4"/></g>`), g.Group(children),
	)
}

func PlaylistRepeatSongFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 9a1 1 0 1 1-2 0V5.618l-.553.276a1 1 0 1 1-.894-1.788l2-1A1 1 0 0 1 22 4zM6 19a1 1 0 0 1-1.707.707l-2-2a1 1 0 0 1 0-1.414l2-2A1 1 0 0 1 6 15v1h11a3 3 0 0 0 3-3a1 1 0 1 1 2 0a5 5 0 0 1-4.998 5H6zM16 7a1 1 0 0 0-1-1H7a5 5 0 0 0-5 5a1 1 0 1 0 2 0a3 3 0 0 1 3-3h8a1 1 0 0 0 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlaylistRepeatSongLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M21 9V4l-2 1m-4 2H7M5 19l-2-2m0 0l2-2m-2 2h14"/><path d="M3 11a4 4 0 0 1 4-4m14 6a4 4 0 0 1-4 4"/></g>`), g.Group(children),
	)
}

func PlaylistRepeatSongThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" d="M21 9V4l-2 1m-4 2H7M5 19l-2-2m0 0l2-2m-2 2h14"/><path d="M3 11a4 4 0 0 1 4-4m14 6a4 4 0 0 1-4 4"/></g>`), g.Group(children),
	)
}

func PlaylistShuffle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17h2.735a4 4 0 0 0 3.43-1.942l3.67-6.116A4 4 0 0 1 16.265 7H21m0 0l-2-2m2 2l-2 2M3 7h2.735a4 4 0 0 1 2.871 1.215M21 17h-4.735a4 4 0 0 1-2.871-1.215M21 17l-2 2m2-2l-2-2"/>`), g.Group(children),
	)
}

func PlaylistShuffleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M3 17h2.735a4 4 0 0 0 3.43-1.942l3.67-6.116A4 4 0 0 1 16.265 7H21m0 0l-2-2m2 2l-2 2M3 7h2.735a4 4 0 0 1 2.871 1.215M21 17h-4.735a4 4 0 0 1-2.871-1.215M21 17l-2 2m2-2l-2-2"/>`), g.Group(children),
	)
}

func PlaylistShuffleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 17h2.735a4 4 0 0 0 3.43-1.942l3.67-6.116A4 4 0 0 1 16.265 7H21m0 0l-2-2m2 2l-2 2M3 7h2.735a4 4 0 0 1 2.871 1.215M21 17h-4.735a4 4 0 0 1-2.871-1.215M21 17l-2 2m2-2l-2-2"/>`), g.Group(children),
	)
}

func PlaylistShuffleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19.707 4.293A1 1 0 0 0 18 5.01l.01.989h-1.745a5 5 0 0 0-4.288 2.428l-3.67 6.115A3 3 0 0 1 5.736 16H3a1 1 0 1 0 0 2h2.735a5 5 0 0 0 4.288-2.427l3.67-6.116A3 3 0 0 1 16.264 8h1.767l.008.72a1 1 0 0 0 1.667.987l.042-.042l-.707-.707l.707.707l1.958-1.958a1 1 0 0 0 0-1.414zM3 6a1 1 0 0 0 0 2h2.735a3 3 0 0 1 2.154.911A1 1 0 1 0 9.324 7.52A5 5 0 0 0 5.735 6zm18.707 10.293l-2-2A1 1 0 0 0 18 15v1h-1.735a3 3 0 0 1-2.154-.911a1 1 0 0 0-1.435 1.393A5 5 0 0 0 16.265 18H18v1a1 1 0 0 0 1.707.707l2-2a1 1 0 0 0 0-1.414" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PlaylistShuffleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3 17h2.735a4 4 0 0 0 3.43-1.942l3.67-6.116A4 4 0 0 1 16.265 7H21m0 0l-2-2m2 2l-2 2M3 7h2.735a4 4 0 0 1 2.871 1.215M21 17h-4.735a4 4 0 0 1-2.871-1.215M21 17l-2 2m2-2l-2-2"/>`), g.Group(children),
	)
}

func PlaylistShuffleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M3 17h2.735a4 4 0 0 0 3.43-1.942l3.67-6.116A4 4 0 0 1 16.265 7H21m0 0l-2-2m2 2l-2 2M3 7h2.735a4 4 0 0 1 2.871 1.215M21 17h-4.735a4 4 0 0 1-2.871-1.215M21 17l-2 2m2-2l-2-2"/>`), g.Group(children),
	)
}

func PlaylistThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M6 6L3 7.732V4.268z"/><path stroke-linecap="round" d="M3 12h18M10 6h11M3 18h18"/></g>`), g.Group(children),
	)
}

func Printer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M8 18H6a2 2 0 0 1-2-2V7h16v9a2 2 0 0 1-2 2h-2M8 3h8v4H8z"/><path stroke-linecap="round" d="M12 11H8"/><path d="M8 15h8v6H8z"/></g>`), g.Group(children),
	)
}

func PrinterBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M8 18H6a2 2 0 0 1-2-2V7h16v9a2 2 0 0 1-2 2h-2M8 3h8v4H8z"/><path stroke-linecap="round" d="M12 11H8"/><path d="M8 15h8v6H8z"/></g>`), g.Group(children),
	)
}

func PrinterDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M20 7H4v9a2 2 0 0 0 2 2h2v-3h8v3h2a2 2 0 0 0 2-2z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M8 18H6a2 2 0 0 1-2-2V7h16v9a2 2 0 0 1-2 2h-2M8 3h8v4H8z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 11H8"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M8 15h8v6H8z"/></g>`), g.Group(children),
	)
}

func PrinterFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 2a1 1 0 0 0-1 1v3H4a1 1 0 0 0-1 1v9a3 3 0 0 0 3 3h1v2a1 1 0 0 0 1 1h8a1 1 0 0 0 1-1v-2h1a3 3 0 0 0 3-3V7a1 1 0 0 0-1-1h-3V3a1 1 0 0 0-1-1zm0 12a1 1 0 0 0-1 1v2H6a1 1 0 0 1-1-1V8h14v8a1 1 0 0 1-1 1h-1v-2a1 1 0 0 0-1-1zm4-2a1 1 0 1 0 0-2H8a1 1 0 1 0 0 2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func PrinterLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M8 18H6a2 2 0 0 1-2-2V7h16v9a2 2 0 0 1-2 2h-2M8 3h8v4H8z"/><path stroke-linecap="round" d="M12 11H8"/><path d="M8 15h8v6H8z"/></g>`), g.Group(children),
	)
}

func PrinterThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M8 18H6a2 2 0 0 1-2-2V7h16v9a2 2 0 0 1-2 2h-2M8 3h8v4H8z"/><path stroke-linecap="round" d="M12 11H8"/><path d="M8 15h8v6H8z"/></g>`), g.Group(children),
	)
}

func Profile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M4 18a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2Z"/><circle cx="12" cy="7" r="3"/></g>`), g.Group(children),
	)
}

func ProfileBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><path stroke-linejoin="round" d="M4 18a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2Z"/><circle cx="12" cy="7" r="3"/></g>`), g.Group(children),
	)
}

func ProfileCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="2" d="M21 12a8.96 8.96 0 0 1-1.526 5.016A8.99 8.99 0 0 1 12 21a8.99 8.99 0 0 1-7.474-3.984A9 9 0 1 1 21 12Z"/><path fill="currentColor" d="M13 9a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-1 1a1 1 0 0 1-1-1H9a3 3 0 0 0 3 3zm-1-1a1 1 0 0 1 1-1V6a3 3 0 0 0-3 3zm1-1a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3zm-6.834 9.856l-.959-.285l-.155.523l.355.413zm13.668 0l.76.651l.354-.413l-.155-.523zM9 16h6v-2H9zm0-2a5 5 0 0 0-4.793 3.571l1.917.57A3 3 0 0 1 9 16zm3 6a7.98 7.98 0 0 1-6.075-2.795l-1.518 1.302A9.98 9.98 0 0 0 12 22zm3-4c1.357 0 2.506.902 2.876 2.142l1.916-.571A5 5 0 0 0 15 14zm3.075 1.205A7.98 7.98 0 0 1 12 20v2a9.98 9.98 0 0 0 7.593-3.493z"/></g>`), g.Group(children),
	)
}

func ProfileCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="2.5" d="M21 12a8.96 8.96 0 0 1-1.526 5.016A8.99 8.99 0 0 1 12 21a8.99 8.99 0 0 1-7.474-3.984A9 9 0 1 1 21 12Z"/><path fill="currentColor" d="M12.75 9a.75.75 0 0 1-.75.75v2.5A3.25 3.25 0 0 0 15.25 9zm-.75.75a.75.75 0 0 1-.75-.75h-2.5A3.25 3.25 0 0 0 12 12.25zM11.25 9a.75.75 0 0 1 .75-.75v-2.5A3.25 3.25 0 0 0 8.75 9zm.75-.75a.75.75 0 0 1 .75.75h2.5A3.25 3.25 0 0 0 12 5.75zm-6.834 9.606L3.968 17.5l-.195.653l.444.517zm13.668 0l.949.814l.444-.517l-.195-.653zM9 16.25h6v-2.5H9zm0-2.5a5.25 5.25 0 0 0-5.032 3.75l2.396.713A2.75 2.75 0 0 1 9 16.25zm3 6a7.73 7.73 0 0 1-5.885-2.707L4.217 18.67A10.23 10.23 0 0 0 12 22.25zm3-3.5c1.244 0 2.298.827 2.636 1.963l2.396-.713A5.25 5.25 0 0 0 15 13.75zm2.885.793A7.73 7.73 0 0 1 12 19.75v2.5a10.23 10.23 0 0 0 7.783-3.58z"/></g>`), g.Group(children),
	)
}

func ProfileCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M15 15H9a4 4 0 0 0-3.834 2.856A8.98 8.98 0 0 0 12 21a8.98 8.98 0 0 0 6.834-3.144A4 4 0 0 0 15 15" opacity=".16"/><path stroke="currentColor" stroke-width="2" d="M21 12a8.96 8.96 0 0 1-1.526 5.016A8.99 8.99 0 0 1 12 21a8.99 8.99 0 0 1-7.474-3.984A9 9 0 1 1 21 12Z"/><path fill="currentColor" d="M13 9a1 1 0 0 1-1 1v2a3 3 0 0 0 3-3zm-1 1a1 1 0 0 1-1-1H9a3 3 0 0 0 3 3zm-1-1a1 1 0 0 1 1-1V6a3 3 0 0 0-3 3zm1-1a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3zm-6.834 9.856l-.959-.285l-.155.523l.355.413zm13.668 0l.76.651l.354-.413l-.155-.523zM9 16h6v-2H9zm0-2a5 5 0 0 0-4.793 3.571l1.917.57A3 3 0 0 1 9 16zm3 6a7.98 7.98 0 0 1-6.075-2.795l-1.518 1.302A9.98 9.98 0 0 0 12 22zm3-4c1.357 0 2.506.902 2.876 2.142l1.916-.571A5 5 0 0 0 15 14zm3.075 1.205A7.98 7.98 0 0 1 12 20v2a9.98 9.98 0 0 0 7.593-3.493z"/></g>`), g.Group(children),
	)
}

func ProfileCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 4a8 8 0 0 0-6.96 11.947A4.99 4.99 0 0 1 9 14h6a4.99 4.99 0 0 1 3.96 1.947A8 8 0 0 0 12 4m7.943 14.076q.188-.245.36-.502A9.96 9.96 0 0 0 22 12c0-5.523-4.477-10-10-10S2 6.477 2 12a9.96 9.96 0 0 0 2.057 6.076l-.005.018l.355.413A9.98 9.98 0 0 0 12 22q.324 0 .644-.02a9.95 9.95 0 0 0 5.031-1.745a10 10 0 0 0 1.918-1.728l.355-.413zM12 6a3 3 0 1 0 0 6a3 3 0 0 0 0-6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ProfileCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M21 12a8.96 8.96 0 0 1-1.526 5.016A8.99 8.99 0 0 1 12 21a8.99 8.99 0 0 1-7.474-3.984A9 9 0 1 1 21 12Z"/><path fill="currentColor" d="M13.25 9c0 .69-.56 1.25-1.25 1.25v1.5A2.75 2.75 0 0 0 14.75 9zM12 10.25c-.69 0-1.25-.56-1.25-1.25h-1.5A2.75 2.75 0 0 0 12 11.75zM10.75 9c0-.69.56-1.25 1.25-1.25v-1.5A2.75 2.75 0 0 0 9.25 9zM12 7.75c.69 0 1.25.56 1.25 1.25h1.5A2.75 2.75 0 0 0 12 6.25zM5.166 17.856l-.719-.214l-.117.392l.267.31zm13.668 0l.57.489l.266-.31l-.117-.393zM9 15.75h6v-1.5H9zm0-1.5a4.75 4.75 0 0 0-4.553 3.392l1.438.428A3.25 3.25 0 0 1 9 15.75zm3 6a8.23 8.23 0 0 1-6.265-2.882l-1.138.977A9.73 9.73 0 0 0 12 21.75zm3-4.5c1.47 0 2.715.978 3.115 2.32l1.438-.428A4.75 4.75 0 0 0 15 14.25zm3.265 1.618A8.23 8.23 0 0 1 12 20.25v1.5a9.73 9.73 0 0 0 7.403-3.405z"/></g>`), g.Group(children),
	)
}

func ProfileCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" d="M21 12a8.96 8.96 0 0 1-1.526 5.016A8.99 8.99 0 0 1 12 21a8.99 8.99 0 0 1-7.474-3.984A9 9 0 1 1 21 12Z"/><path fill="currentColor" d="M13.5 9a1.5 1.5 0 0 1-1.5 1.5v1A2.5 2.5 0 0 0 14.5 9zM12 10.5A1.5 1.5 0 0 1 10.5 9h-1a2.5 2.5 0 0 0 2.5 2.5zM10.5 9A1.5 1.5 0 0 1 12 7.5v-1A2.5 2.5 0 0 0 9.5 9zM12 7.5A1.5 1.5 0 0 1 13.5 9h1A2.5 2.5 0 0 0 12 6.5zM5.166 17.856l-.48-.142l-.077.261l.177.207zm13.668 0l.38.326l.177-.207l-.078-.261zM9 15.5h6v-1H9zm0-1a4.5 4.5 0 0 0-4.313 3.214l.958.285A3.5 3.5 0 0 1 9 15.5zm3 6a8.48 8.48 0 0 1-6.455-2.97l-.759.652A9.48 9.48 0 0 0 12 21.5zm3-5a3.5 3.5 0 0 1 3.355 2.5l.958-.286A4.5 4.5 0 0 0 15 14.5zm3.455 2.03A8.48 8.48 0 0 1 12 20.5v1a9.48 9.48 0 0 0 7.214-3.318z"/></g>`), g.Group(children),
	)
}

func ProfileDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 18a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M4 18a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2Z"/><circle cx="12" cy="7" r="3" stroke="currentColor" stroke-width="2"/></g>`), g.Group(children),
	)
}

func ProfileFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 7a4 4 0 1 1 8 0a4 4 0 0 1-8 0m0 6a5 5 0 0 0-5 5a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3a5 5 0 0 0-5-5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ProfileLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linejoin="round" d="M4 18a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2Z"/><circle cx="12" cy="7" r="3"/></g>`), g.Group(children),
	)
}

func ProfileThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linejoin="round" d="M4 18a4 4 0 0 1 4-4h8a4 4 0 0 1 4 4a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2Z"/><circle cx="12" cy="7" r="3"/></g>`), g.Group(children),
	)
}

func QuestionMarkCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M10.586 7.586c.39-.39.9-.585 1.41-.586a2 2 0 0 1 1.418.586c.39.39.586.902.586 1.414s-.195 1.024-.586 1.414a2 2 0 0 1-1.418.586L12 12"/></g>`), g.Group(children),
	)
}

func QuestionMarkCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M10.586 7.586c.39-.39.9-.585 1.41-.586a2 2 0 0 1 1.418.586c.39.39.586.902.586 1.414s-.195 1.024-.586 1.414a2 2 0 0 1-1.418.586L12 12"/></g>`), g.Group(children),
	)
}

func QuestionMarkCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 16h.01v.01H12z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.586 7.586c.39-.39.9-.585 1.41-.586a2 2 0 0 1 1.418.586c.39.39.586.902.586 1.414s-.195 1.024-.586 1.414a2 2 0 0 1-1.418.586L12 12"/></g>`), g.Group(children),
	)
}

func QuestionMarkCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m10.707-3.707A1 1 0 0 0 11.998 8a1 1 0 0 0-.705.293a1 1 0 0 1-1.414-1.414a3 3 0 0 1 2.116-.88a3 3 0 0 1 2.126.88A2.994 2.994 0 0 1 13 11.829l.001.167a1 1 0 0 1-2 .008l-.004-1A1 1 0 0 1 11.999 10a1 1 0 0 0 .71-.293A1 1 0 0 0 13 9a1 1 0 0 0-.293-.707M10.5 16a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func QuestionMarkCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M10.586 7.586c.39-.39.9-.585 1.41-.586a2 2 0 0 1 1.418.586c.39.39.586.902.586 1.414s-.195 1.024-.586 1.414a2 2 0 0 1-1.418.586L12 12"/></g>`), g.Group(children),
	)
}

func QuestionMarkCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" d="M10.586 7.586c.39-.39.9-.585 1.41-.586a2 2 0 0 1 1.418.586c.39.39.586.902.586 1.414s-.195 1.024-.586 1.414a2 2 0 0 1-1.418.586L12 12"/></g>`), g.Group(children),
	)
}

func QuestionMarkSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M10.586 7.586c.39-.39.9-.585 1.41-.586a2 2 0 0 1 1.418.586c.39.39.586.902.586 1.414s-.195 1.024-.586 1.414a2 2 0 0 1-1.418.586L12 12"/></g>`), g.Group(children),
	)
}

func QuestionMarkSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3.75" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M10.586 7.586c.39-.39.9-.585 1.41-.586a2 2 0 0 1 1.418.586c.39.39.586.902.586 1.414s-.195 1.024-.586 1.414a2 2 0 0 1-1.418.586L12 12"/></g>`), g.Group(children),
	)
}

func QuestionMarkSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 16h.01v.01H12z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10.586 7.586c.39-.39.9-.585 1.41-.586a2 2 0 0 1 1.418.586c.39.39.586.902.586 1.414s-.195 1.024-.586 1.414a2 2 0 0 1-1.418.586L12 12"/></g>`), g.Group(children),
	)
}

func QuestionMarkSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4.001a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm9.707 4.292A1 1 0 0 0 11.998 8a1 1 0 0 0-.705.293a1 1 0 0 1-1.414-1.414a3 3 0 0 1 2.116-.88a3 3 0 0 1 2.126.88A2.994 2.994 0 0 1 13 11.829l.001.167a1 1 0 0 1-2 .008l-.004-1A1 1 0 0 1 11.999 10a1 1 0 0 0 .71-.293A1 1 0 0 0 13 9a1 1 0 0 0-.293-.707M10.5 16a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func QuestionMarkSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="1.5" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="2.25" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M10.586 7.586c.39-.39.9-.585 1.41-.586a2 2 0 0 1 1.418.586c.39.39.586.902.586 1.414s-.195 1.024-.586 1.414a2 2 0 0 1-1.418.586L12 12"/></g>`), g.Group(children),
	)
}

func QuestionMarkSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="1.5" d="M12 16h.01v.01H12z"/><path stroke-linecap="round" d="M10.586 7.586c.39-.39.9-.585 1.41-.586a2 2 0 0 1 1.418.586c.39.39.586.902.586 1.414s-.195 1.024-.586 1.414a2 2 0 0 1-1.418.586L12 12"/></g>`), g.Group(children),
	)
}

func RelievedFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M14 15.5a4 4 0 0 1-2 .535a4 4 0 0 1-2-.535m4-5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`), g.Group(children),
	)
}

func RelievedFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M14 15.5a4 4 0 0 1-2 .535a4 4 0 0 1-2-.535m4-5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`), g.Group(children),
	)
}

func RelievedFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 15.5a4 4 0 0 1-2 .535a4 4 0 0 1-2-.535m4-5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`), g.Group(children),
	)
}

func RelievedFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m4.073-1.875a1 1 0 0 1 1.302-.552c.346.14.725.218 1.125.218s.779-.078 1.125-.218a1 1 0 1 1 .75 1.854a5 5 0 0 1-1.875.364a5 5 0 0 1-1.875-.364a1 1 0 0 1-.552-1.302m8.302-.552a1 1 0 1 0-.75 1.854a5 5 0 0 0 3.75 0a1 1 0 1 0-.75-1.854c-.346.14-.725.218-1.125.218s-.779-.078-1.125-.218m.126 6.793a1 1 0 0 0-1.002-1.732c-.44.255-.95.401-1.499.401a3 3 0 0 1-1.5-.4a1 1 0 0 0-1 1.73c.736.427 1.591.67 2.5.67c.91 0 1.764-.243 2.5-.67" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func RelievedFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M14 15.5a4 4 0 0 1-2 .535a4 4 0 0 1-2-.535m4-5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`), g.Group(children),
	)
}

func RelievedFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M14 15.5a4 4 0 0 1-2 .535a4 4 0 0 1-2-.535m4-5c.463.188.97.29 1.5.29s1.037-.102 1.5-.29m-10 0c.463.188.97.29 1.5.29s1.037-.102 1.5-.29"/></g>`), g.Group(children),
	)
}

func Restart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M12 3a9 9 0 1 1-5.657 2"/><path d="M3 4.5h4v4"/></g>`), g.Group(children),
	)
}

func RestartBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M12 3a9 9 0 1 1-5.657 2"/><path d="M3 4.5h4v4"/></g>`), g.Group(children),
	)
}

func RestartDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 3a9 9 0 1 1-5.657 2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 4.5h4v4"/></g>`), g.Group(children),
	)
}

func RestartFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2a1 1 0 1 0 0 2a8 8 0 1 1-6.924 3.99l1.217 1.217A1 1 0 0 0 8 8.5v-4a1 1 0 0 0-1-1H3a1 1 0 0 0-.707 1.707l1.33 1.33A9.96 9.96 0 0 0 2 12c0 5.523 4.477 10 10 10s10-4.477 10-10S17.523 2 12 2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func RestartLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 3a9 9 0 1 1-5.657 2"/><path d="M3 4.5h4v4"/></g>`), g.Group(children),
	)
}

func RestartThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M12 3a9 9 0 1 1-5.657 2"/><path d="M3 4.5h4v4"/></g>`), g.Group(children),
	)
}

func Scanner(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M20 12H4m12-8h2a2 2 0 0 1 2 2v2M8 20H6a2 2 0 0 1-2-2v-2m16 0v2a2 2 0 0 1-2 2h-2M4 8V6a2 2 0 0 1 2-2h2"/>`), g.Group(children),
	)
}

func ScannerBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M20 12H4m12-8h2a2 2 0 0 1 2 2v2M8 20H6a2 2 0 0 1-2-2v-2m16 0v2a2 2 0 0 1-2 2h-2M4 8V6a2 2 0 0 1 2-2h2"/>`), g.Group(children),
	)
}

func ScannerFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15 4a1 1 0 0 1 1-1h2a3 3 0 0 1 3 3v2a1 1 0 1 1-2 0V6a1 1 0 0 0-1-1h-2a1 1 0 0 1-1-1M3 12a1.5 1.5 0 0 1 1.5-1.5h15a1.5 1.5 0 0 1 0 3h-15A1.5 1.5 0 0 1 3 12m5 9a1 1 0 0 0 0-2H6a1 1 0 0 1-1-1v-2a1 1 0 1 0-2 0v2a3 3 0 0 0 3 3zm12-6a1 1 0 0 1 1 1v2a3 3 0 0 1-3 3h-2a1 1 0 0 1 0-2h2a1 1 0 0 0 1-1v-2a1 1 0 0 1 1-1M3 8a1 1 0 0 0 2 0V6a1 1 0 0 1 1-1h2a1 1 0 1 0 0-2H6a3 3 0 0 0-3 3z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ScannerLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 12H4m12-8h2a2 2 0 0 1 2 2v2M8 20H6a2 2 0 0 1-2-2v-2m16 0v2a2 2 0 0 1-2 2h-2M4 8V6a2 2 0 0 1 2-2h2"/>`), g.Group(children),
	)
}

func ScannerThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M20 12H4m12-8h2a2 2 0 0 1 2 2v2M8 20H6a2 2 0 0 1-2-2v-2m16 0v2a2 2 0 0 1-2 2h-2M4 8V6a2 2 0 0 1 2-2h2"/>`), g.Group(children),
	)
}

func ScreenFull(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9V6a2 2 0 0 1 2-2h3m11 11v3a2 2 0 0 1-2 2h-3m0-16h3a2 2 0 0 1 2 2v3M9 20H6a2 2 0 0 1-2-2v-3"/>`), g.Group(children),
	)
}

func ScreenFullBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 9V6a2 2 0 0 1 2-2h3m11 11v3a2 2 0 0 1-2 2h-3m0-16h3a2 2 0 0 1 2 2v3M9 20H6a2 2 0 0 1-2-2v-3"/>`), g.Group(children),
	)
}

func ScreenFullDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 9V6a2 2 0 0 1 2-2h3m11 11v3a2 2 0 0 1-2 2h-3m0-16h3a2 2 0 0 1 2 2v3M9 20H6a2 2 0 0 1-2-2v-3"/>`), g.Group(children),
	)
}

func ScreenFullFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.5 4.5a1 1 0 0 1 1-1h2a3 3 0 0 1 3 3v2a1 1 0 1 1-2 0v-2a1 1 0 0 0-1-1h-2a1 1 0 0 1-1-1m-11 4a1 1 0 1 0 2 0v-2a1 1 0 0 1 1-1h2a1 1 0 0 0 0-2h-2a3 3 0 0 0-3 3zm17 7a1 1 0 1 0-2 0v2a1 1 0 0 1-1 1h-2a1 1 0 1 0 0 2h2a3 3 0 0 0 3-3zm-12 5a1 1 0 1 0 0-2h-2a1 1 0 0 1-1-1v-2a1 1 0 1 0-2 0v2a3 3 0 0 0 3 3z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ScreenFullLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 9V6a2 2 0 0 1 2-2h3m11 11v3a2 2 0 0 1-2 2h-3m0-16h3a2 2 0 0 1 2 2v3M9 20H6a2 2 0 0 1-2-2v-3"/>`), g.Group(children),
	)
}

func ScreenFullThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 9V6a2 2 0 0 1 2-2h3m11 11v3a2 2 0 0 1-2 2h-3m0-16h3a2 2 0 0 1 2 2v3M9 20H6a2 2 0 0 1-2-2v-3"/>`), g.Group(children),
	)
}

func ScreenNormal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 4v3a2 2 0 0 1-2 2H4m11 11v-3a2 2 0 0 1 2-2h3m0-6h-3a2 2 0 0 1-2-2V4M4 15h3a2 2 0 0 1 2 2v3"/>`), g.Group(children),
	)
}

func ScreenNormalBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M9 4v3a2 2 0 0 1-2 2H4m11 11v-3a2 2 0 0 1 2-2h3m0-6h-3a2 2 0 0 1-2-2V4M4 15h3a2 2 0 0 1 2 2v3"/>`), g.Group(children),
	)
}

func ScreenNormalDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 4v3a2 2 0 0 1-2 2H4m11 11v-3a2 2 0 0 1 2-2h3m0-6h-3a2 2 0 0 1-2-2V4M4 15h3a2 2 0 0 1 2 2v3"/>`), g.Group(children),
	)
}

func ScreenNormalFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.5 4.5a1 1 0 1 0-2 0v2a1 1 0 0 1-1 1h-2a1 1 0 0 0 0 2h2a3 3 0 0 0 3-3zm5 15a1 1 0 1 0 2 0v-2a1 1 0 0 1 1-1h2a1 1 0 1 0 0-2h-2a3 3 0 0 0-3 3zm6-11a1 1 0 0 1-1 1h-2a3 3 0 0 1-3-3v-2a1 1 0 1 1 2 0v2a1 1 0 0 0 1 1h2a1 1 0 0 1 1 1m-16 6a1 1 0 1 0 0 2h2a1 1 0 0 1 1 1v2a1 1 0 1 0 2 0v-2a3 3 0 0 0-3-3z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ScreenNormalLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 4v3a2 2 0 0 1-2 2H4m11 11v-3a2 2 0 0 1 2-2h3m0-6h-3a2 2 0 0 1-2-2V4M4 15h3a2 2 0 0 1 2 2v3"/>`), g.Group(children),
	)
}

func ScreenNormalThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 4v3a2 2 0 0 1-2 2H4m11 11v-3a2 2 0 0 1 2-2h3m0-6h-3a2 2 0 0 1-2-2V4M4 15h3a2 2 0 0 1 2 2v3"/>`), g.Group(children),
	)
}

func Search(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314"/>`), g.Group(children),
	)
}

func SearchBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314"/>`), g.Group(children),
	)
}

func SearchDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M19 11a8 8 0 1 1-16 0a8 8 0 0 1 16 0" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314"/></g>`), g.Group(children),
	)
}

func SearchFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 2a9 9 0 1 0 5.618 16.032l3.675 3.675a1 1 0 0 0 1.414-1.414l-3.675-3.675A9 9 0 0 0 11 2m-6 9a6 6 0 1 1 12 0a6 6 0 0 1-12 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SearchLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314"/>`), g.Group(children),
	)
}

func SearchThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314"/>`), g.Group(children),
	)
}

func Send(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 12l-.604-5.437C4.223 5.007 5.825 3.864 7.24 4.535l11.944 5.658c1.525.722 1.525 2.892 0 3.614L7.24 19.466c-1.415.67-3.017-.472-2.844-2.028zm0 0h7"/>`), g.Group(children),
	)
}

func SendBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m5 12l-.604-5.437C4.223 5.007 5.825 3.864 7.24 4.535l11.944 5.658c1.525.722 1.525 2.892 0 3.614L7.24 19.466c-1.415.67-3.017-.472-2.844-2.028zm0 0h7"/>`), g.Group(children),
	)
}

func SendDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m7.24 4.535l11.944 5.658c1.525.722 1.525 2.892 0 3.614L7.24 19.466c-1.415.67-3.017-.472-2.844-2.028l.58-5.216a2 2 0 0 0 0-.442l-.58-5.216c-.173-1.557 1.429-2.7 2.844-2.029" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5 12l-.604-5.437C4.223 5.007 5.825 3.864 7.24 4.535l11.944 5.658c1.525.722 1.525 2.892 0 3.614L7.24 19.466c-1.415.67-3.017-.472-2.844-2.028zm0 0h7"/></g>`), g.Group(children),
	)
}

func SendFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.402 6.673c-.26-2.334 2.143-4.048 4.266-3.042l11.944 5.658c2.288 1.083 2.288 4.339 0 5.422L7.668 20.37c-2.123 1.006-4.525-.708-4.266-3.042L3.882 13H12a1 1 0 1 0 0-2H3.883z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SendLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 12l-.604-5.437C4.223 5.007 5.825 3.864 7.24 4.535l11.944 5.658c1.525.722 1.525 2.892 0 3.614L7.24 19.466c-1.415.67-3.017-.472-2.844-2.028zm0 0h7"/>`), g.Group(children),
	)
}

func SendThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5 12l-.604-5.437C4.223 5.007 5.825 3.864 7.24 4.535l11.944 5.658c1.525.722 1.525 2.892 0 3.614L7.24 19.466c-1.415.67-3.017-.472-2.844-2.028zm0 0h7"/>`), g.Group(children),
	)
}

func Settings(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="2" stroke="currentColor" stroke-width="2"/><path fill="currentColor" d="m5.399 5.88l.5-.867a1 1 0 0 0-1.234.186zM3.4 9.344l-.956-.295a1 1 0 0 0 .456 1.16zm-.002 5.311l-.5-.866a1 1 0 0 0-.455 1.162zm2 3.464l-.734.68a1 1 0 0 0 1.234.186zm4.6 2.655H9a1 1 0 0 0 .778.975zm4.001.002l.223.975a1 1 0 0 0 .777-.975zM18.6 18.12l-.5.866a1 1 0 0 0 1.233-.186zm1.998-3.466l.956.295a1 1 0 0 0-.456-1.16zm.002-5.311l.5.866a1 1 0 0 0 .455-1.162zm-2-3.465l.734-.679a1 1 0 0 0-1.234-.187zM14 3.225h1a1 1 0 0 0-.777-.975zm-4-.002l-.223-.975A1 1 0 0 0 9 3.223zm4 1.849h-1zm5 8.66l-.5.866zm-2 3.464l-.5.866zM5 13.732l.5.866zm2-6.928l-.5.866zM4.356 9.639a8 8 0 0 1 1.776-3.08L4.665 5.2a10 10 0 0 0-2.22 3.85zM5.072 16a8 8 0 0 1-.718-1.64l-1.91.592c.217.701.515 1.388.896 2.048zm1.06 1.441A8 8 0 0 1 5.073 16L3.34 17c.38.66.827 1.261 1.325 1.8zm7.646 2.361a8 8 0 0 1-3.556-.002l-.445 1.95a10 10 0 0 0 4.446.002zm5.866-5.441a8 8 0 0 1-1.776 3.08l1.467 1.36a10 10 0 0 0 2.22-3.85zM18.928 8c.306.53.545 1.08.718 1.64l1.91-.592A10 10 0 0 0 20.66 7zm-1.06-1.441c.397.43.754.91 1.06 1.441l1.732-1a10 10 0 0 0-1.325-1.8zm-7.646-2.361a8 8 0 0 1 3.556.002l.444-1.95a10 10 0 0 0-4.445-.002zm.778.874v-1.85H9v1.85zm-3.5.866l-1.601-.925l-1 1.732l1.6.925zm-3 6.928l-1.601.924l1 1.732l1.6-.924zm1-3.464l-1.6-.923l-1 1.732l1.6.923zM11 20.775v-1.847H9v1.847zM6.5 16.33l-1.601.925l1 1.732l1.6-.925zm12.601.925L17.5 16.33l-1 1.732l1.601.925zM15 20.777v-1.849h-2v1.85zm5.101-12.3l-1.601.925l1 1.732l1.601-.925zm.998 5.312l-1.599-.923l-1 1.732l1.6.923zM15 5.072V3.225h-2v1.847zm3.101-.059l-1.601.925l1 1.732l1.601-.925zM13 5.072c0 2.31 2.5 3.753 4.5 2.598l-1-1.732a1 1 0 0 1-1.5-.866zm5.5 4.33c-2 1.155-2 4.041 0 5.196l1-1.732a1 1 0 0 1 0-1.732zm-1 6.928c-2-1.154-4.5.289-4.5 2.598h2a1 1 0 0 1 1.5-.866zM11 18.928c0-2.31-2.5-3.753-4.5-2.598l1 1.732a1 1 0 0 1 1.5.866zm-5.5-4.33c2-1.155 2-4.041 0-5.196l-1 1.732a1 1 0 0 1 0 1.732zM9 5.072a1 1 0 0 1-1.5.866l-1 1.732c2 1.154 4.5-.289 4.5-2.598z"/></g>`), g.Group(children),
	)
}

func SettingsBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="2" stroke="currentColor" stroke-width="2.5"/><path fill="currentColor" d="m5.399 5.88l.625-1.083a1.25 1.25 0 0 0-1.542.232zM3.4 9.344l-1.195-.369a1.25 1.25 0 0 0 .57 1.451zm-.002 5.311l-.625-1.082a1.25 1.25 0 0 0-.57 1.452zm2 3.464l-.918.85a1.25 1.25 0 0 0 1.543.233zm4.6 2.655H8.75c0 .583.404 1.089.973 1.219zm4.001.002l.278 1.22a1.25 1.25 0 0 0 .972-1.22zM18.6 18.12l-.625 1.082a1.25 1.25 0 0 0 1.542-.232zm1.998-3.466l1.195.369a1.25 1.25 0 0 0-.57-1.451zm.002-5.311l.625 1.082a1.25 1.25 0 0 0 .57-1.452zm-2-3.465l.917-.849a1.25 1.25 0 0 0-1.542-.233zM14 3.225h1.25a1.25 1.25 0 0 0-.972-1.219zm-4-.002l-.278-1.22a1.25 1.25 0 0 0-.972 1.22zm4 1.849h-1.25zm5 8.66l-.625 1.083zm-2 3.464l-.625 1.083zM5 13.732l.625 1.083zm2-6.928l-.625 1.082zm-2.405 2.91a7.74 7.74 0 0 1 1.72-2.985l-1.833-1.7a10.24 10.24 0 0 0-2.276 3.947zm.693 6.161a7.8 7.8 0 0 1-.695-1.589l-2.388.74c.222.719.528 1.422.918 2.099zm1.028 1.396a7.8 7.8 0 0 1-1.028-1.396l-2.165 1.25c.39.676.848 1.293 1.358 1.845zm7.406 2.288a7.74 7.74 0 0 1-3.444-.003l-.556 2.438c1.479.337 3.037.349 4.556.002zm5.683-5.272a7.74 7.74 0 0 1-1.72 2.984l1.833 1.7a10.24 10.24 0 0 0 2.276-3.947zm-.693-6.162c.297.514.527 1.047.695 1.588l2.388-.74a10.3 10.3 0 0 0-.918-2.098zm-1.028-1.397q.58.626 1.028 1.397l2.165-1.25a10.3 10.3 0 0 0-1.359-1.845zm-7.406-2.287a7.74 7.74 0 0 1 3.444.003l.556-2.438a10.24 10.24 0 0 0-4.556-.002zm.972.63V3.224h-2.5v1.849zm-3.625.65l-1.601-.924l-1.25 2.165l1.6.924zm-3.25 6.929l-1.601.924l1.25 2.165l1.6-.924zm1.25-3.465l-1.6-.923l-1.25 2.165l1.6.924zm5.625 11.59v-1.847h-2.5v1.847zm-4.875-4.661l-1.601.924l1.25 2.165l1.6-.924zm12.851.924l-1.601-.924l-1.25 2.165l1.601.924zm-3.976 3.74v-1.85h-2.5v1.85zM19.976 8.26l-1.601.924l1.25 2.165l1.601-.924zm1.248 5.312l-1.599-.924l-1.25 2.166l1.6.923zm-5.974-8.5V3.225h-2.5v1.847zm2.726-.275l-1.601.924l1.25 2.165l1.601-.924zm-5.226.275c0 2.502 2.708 4.065 4.875 2.814l-1.25-2.165a.75.75 0 0 1-1.125-.65zm5.625 4.113c-2.167 1.251-2.167 4.379 0 5.63l1.25-2.165a.75.75 0 0 1 0-1.3zm-.75 6.929c-2.167-1.251-4.875.312-4.875 2.814h2.5a.75.75 0 0 1 1.125-.65zm-6.375 2.814c0-2.502-2.708-4.065-4.875-2.814l1.25 2.165a.75.75 0 0 1 1.125.65zm-5.625-4.113c2.167-1.251 2.167-4.379 0-5.63l-1.25 2.165a.75.75 0 0 1 0 1.3zM8.75 5.072a.75.75 0 0 1-1.125.65l-1.25 2.164c2.167 1.251 4.875-.312 4.875-2.814z"/></g>`), g.Group(children),
	)
}

func SettingsDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M3.4 9.345a9 9 0 0 1 2-3.466l1.6.925c1.333.77 3-.193 3-1.732v-1.85a9 9 0 0 1 4 .003v1.847c0 1.54 1.667 2.502 3 1.732l1.601-.925a9 9 0 0 1 2 3.465L19 10.268c-1.333.77-1.333 2.694 0 3.464l1.6.923a9 9 0 0 1-1.998 3.466L17 17.196c-1.333-.77-3 .193-3 1.732v1.85a9 9 0 0 1-4-.003v-1.847c0-1.54-1.667-2.502-3-1.732l-1.601.925a9 9 0 0 1-2-3.465l1.6-.924c1.334-.77 1.334-2.694 0-3.464zM12 14a2 2 0 1 0 0-4a2 2 0 0 0 0 4" clip-rule="evenodd" opacity=".16"/><circle cx="12" cy="12" r="2" stroke="currentColor" stroke-width="2"/><path fill="currentColor" d="m5.399 5.88l.5-.867a1 1 0 0 0-1.234.186zM3.4 9.344l-.956-.295a1 1 0 0 0 .456 1.16zm-.002 5.311l-.5-.866a1 1 0 0 0-.455 1.162zm2 3.464l-.734.68a1 1 0 0 0 1.234.187zm4.6 2.655H9a1 1 0 0 0 .778.975zm4.001.002l.223.975a1 1 0 0 0 .777-.975zM18.6 18.12l-.5.866a1 1 0 0 0 1.233-.186zm1.998-3.466l.956.295a1 1 0 0 0-.456-1.16zm.002-5.311l.5.866a1 1 0 0 0 .455-1.162zm-2-3.465l.734-.679a1 1 0 0 0-1.234-.187zM14 3.225h1a1 1 0 0 0-.777-.975zm-4-.002l-.223-.975A1 1 0 0 0 9 3.223zm4 1.849h-1zm5 8.66l-.5.866zm-2 3.464l-.5.866zM5 13.732l.5.866zm2-6.928l-.5.866zM4.356 9.639a8 8 0 0 1 1.776-3.08L4.665 5.2a10 10 0 0 0-2.22 3.85zM5.072 16a8 8 0 0 1-.718-1.64l-1.91.592c.217.701.515 1.388.896 2.048zm1.06 1.441A8 8 0 0 1 5.073 16L3.34 17c.38.66.827 1.261 1.325 1.8zm7.646 2.361a8 8 0 0 1-3.556-.002l-.445 1.95a10 10 0 0 0 4.446.002zm5.866-5.441a8 8 0 0 1-1.776 3.08l1.467 1.36a10 10 0 0 0 2.22-3.85zM18.928 8c.306.53.545 1.08.718 1.64l1.91-.592A10 10 0 0 0 20.66 7zm-1.06-1.441c.397.43.754.91 1.06 1.441l1.732-1a10 10 0 0 0-1.325-1.8zm-7.646-2.361a8 8 0 0 1 3.556.002l.444-1.95a10 10 0 0 0-4.445-.002zm.778.874v-1.85H9v1.85zm-3.5.866l-1.601-.925l-1 1.732l1.6.925zm-3 6.928l-1.601.924l1 1.732l1.6-.924zm1-3.464l-1.6-.923l-1 1.732l1.6.923zM11 20.775v-1.847H9v1.847zM6.5 16.33l-1.601.925l1 1.732l1.6-.925zm12.601.925L17.5 16.33l-1 1.732l1.601.925zM15 20.777v-1.849h-2v1.85zm5.101-12.3l-1.601.925l1 1.732l1.601-.924zm.998 5.312l-1.599-.923l-1 1.732l1.6.923zM15 5.072V3.225h-2v1.847zm3.101-.059l-1.601.925l1 1.732l1.601-.925zM13 5.072c0 2.31 2.5 3.753 4.5 2.598l-1-1.732a1 1 0 0 1-1.5-.866zm5.5 4.33c-2 1.155-2 4.041 0 5.196l1-1.732a1 1 0 0 1 0-1.732zm-1 6.928c-2-1.154-4.5.289-4.5 2.598h2a1 1 0 0 1 1.5-.866zM11 18.928c0-2.31-2.5-3.753-4.5-2.598l1 1.732a1 1 0 0 1 1.5.866zm-5.5-4.33c2-1.155 2-4.041 0-5.196l-1 1.732a1 1 0 0 1 0 1.732zM9 5.072a1 1 0 0 1-1.5.866l-1 1.732c2 1.155 4.5-.289 4.5-2.598z"/></g>`), g.Group(children),
	)
}

func SettingsFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 3.223a1 1 0 0 1 .777-.975a10 10 0 0 1 4.445.002a1 1 0 0 1 .778.975v1.847a1 1 0 0 0 1.5.866l1.601-.925a1 1 0 0 1 1.234.187a10.1 10.1 0 0 1 2.221 3.848a1 1 0 0 1-.455 1.162l-1.601.924a1 1 0 0 0 0 1.732l1.6.923a1 1 0 0 1 .455 1.161a10 10 0 0 1-2.22 3.851a1 1 0 0 1-1.234.186l-1.601-.925a1 1 0 0 0-1.5.866v1.85a1 1 0 0 1-.778.974a10 10 0 0 1-4.445-.002A1 1 0 0 1 9 20.775v-1.847a1 1 0 0 0-1.5-.866l-1.601.925a1 1 0 0 1-1.234-.187A10 10 0 0 1 3.34 17a10 10 0 0 1-.896-2.048a1 1 0 0 1 .455-1.162l1.6-.924a1 1 0 0 0 0-1.732l-1.598-.923a1 1 0 0 1-.456-1.161a10 10 0 0 1 2.22-3.85a1 1 0 0 1 1.233-.187l1.602.925A1 1 0 0 0 9 5.072zM12 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SettingsLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="2" stroke="currentColor" stroke-width="1.5"/><path fill="currentColor" d="m5.399 5.88l.375-.65a.75.75 0 0 0-.925.14zM3.4 9.344l-.717-.222a.75.75 0 0 0 .342.871zm-.002 5.311l-.375-.65a.75.75 0 0 0-.341.872zm2 3.464l-.55.51a.75.75 0 0 0 .925.14zm4.6 2.655h-.75c0 .35.243.653.584.731zm4.001.002l.167.732a.75.75 0 0 0 .583-.732zM18.6 18.12l-.375.65a.75.75 0 0 0 .925-.14zm1.998-3.466l.717.222a.75.75 0 0 0-.342-.871zm.002-5.311l.375.65a.75.75 0 0 0 .341-.872zm-2-3.465l.55-.509a.75.75 0 0 0-.925-.14zM14 3.225h.75a.75.75 0 0 0-.583-.731zm-4-.002l-.167-.732a.75.75 0 0 0-.583.732zm4 1.849h-.75zm5 8.66l-.375.65zm-2 3.464l-.375.65zM5 13.732l.375.65zm2-6.928l-.375.65zM4.117 9.566a8.24 8.24 0 0 1 1.831-3.177l-1.1-1.02a9.74 9.74 0 0 0-2.164 3.754zm.738 6.559a8.3 8.3 0 0 1-.74-1.69l-1.432.443a9.8 9.8 0 0 0 .873 1.997zm1.094 1.486a8.3 8.3 0 0 1-1.094-1.486l-1.299.75a9.8 9.8 0 0 0 1.292 1.755zm7.884 2.435a8.24 8.24 0 0 1-3.666-.002l-.334 1.462a9.74 9.74 0 0 0 4.334.003zm6.05-5.612a8.24 8.24 0 0 1-1.831 3.177l1.1 1.02a9.74 9.74 0 0 0 2.164-3.754zm-.738-6.559q.471.823.74 1.69l1.432-.443a9.8 9.8 0 0 0-.873-1.997zM18.05 6.389c.41.443.778.94 1.094 1.486l1.299-.75A9.8 9.8 0 0 0 19.15 5.37zm-7.884-2.435a8.24 8.24 0 0 1 3.666.002l.334-1.462a9.74 9.74 0 0 0-4.334-.003zm.583 1.118v-1.85h-1.5v1.85zM7.375 6.154L5.774 5.23l-.75 1.299l1.6.924zm-2.75 6.929l-1.601.924l.75 1.299l1.6-.924zm.75-3.465l-1.6-.923l-.75 1.3l1.6.923zm5.375 11.157v-1.847h-1.5v1.847zm-4.125-4.228l-1.601.924l.75 1.3l1.6-.925zm12.351.924l-1.601-.924l-.75 1.299l1.601.924zm-4.226 3.306v-1.849h-1.5v1.85zm5.476-12.083l-1.601.924l.75 1.3l1.601-.925zm.748 5.312l-1.599-.924l-.75 1.3l1.6.923zM14.75 5.072V3.225h-1.5v1.847zm3.476.158l-1.601.924l.75 1.3l1.601-.925zm-4.976-.158c0 2.117 2.292 3.44 4.125 2.381l-.75-1.299a1.25 1.25 0 0 1-1.875-1.082zm5.375 4.546c-1.833 1.059-1.833 3.705 0 4.764l.75-1.3a1.25 1.25 0 0 1 0-2.165zm-1.25 6.929c-1.833-1.059-4.125.264-4.125 2.381h1.5a1.25 1.25 0 0 1 1.875-1.082zm-6.625 2.381c0-2.117-2.292-3.44-4.125-2.381l.75 1.299a1.25 1.25 0 0 1 1.875 1.082zm-5.375-4.546c1.833-1.059 1.833-3.705 0-4.764l-.75 1.3a1.25 1.25 0 0 1 0 2.165zm3.875-9.31a1.25 1.25 0 0 1-1.875 1.082l-.75 1.3c1.833 1.058 4.125-.265 4.125-2.382z"/></g>`), g.Group(children),
	)
}

func SettingsThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="2" stroke="currentColor"/><path fill="currentColor" d="m5.399 5.88l.25-.434a.5.5 0 0 0-.617.093zM3.4 9.344l-.478-.148a.5.5 0 0 0 .228.58zm-.002 5.311l-.25-.433a.5.5 0 0 0-.228.581zm2 3.464l-.367.34a.5.5 0 0 0 .617.093zm4.6 2.655h-.5a.5.5 0 0 0 .39.488zm4.001.002l.111.488a.5.5 0 0 0 .389-.488zM18.6 18.12l-.25.433a.5.5 0 0 0 .617-.093zm1.998-3.466l.478.148a.5.5 0 0 0-.228-.58zm.002-5.311l.25.433a.5.5 0 0 0 .228-.581zm-2-3.465l.367-.34a.5.5 0 0 0-.617-.093zM14 3.225h.5a.5.5 0 0 0-.389-.487zm-4-.002l-.111-.488a.5.5 0 0 0-.39.488zm4 1.849h-.5zm5 8.66l-.25.433zm-2 3.464l-.25.433zM5 13.732l.25.433zm2-6.928l-.25.433zM3.878 9.492a8.5 8.5 0 0 1 1.887-3.273l-.733-.68a9.5 9.5 0 0 0-2.11 3.658zm.76 6.758a8.5 8.5 0 0 1-.761-1.742l-.956.296a9.5 9.5 0 0 0 .852 1.946zm1.128 1.53a8.5 8.5 0 0 1-1.127-1.53l-.866.5a9.5 9.5 0 0 0 1.259 1.71zm8.123 2.51a8.5 8.5 0 0 1-3.778-.002l-.222.974a9.5 9.5 0 0 0 4.222.003zm6.233-5.782a8.5 8.5 0 0 1-1.887 3.273l.733.68a9.5 9.5 0 0 0 2.11-3.658zm-.76-6.758c.324.563.577 1.147.762 1.742l.955-.296a9.5 9.5 0 0 0-.852-1.946zm-1.128-1.53a8.5 8.5 0 0 1 1.127 1.53l.866-.5a9.5 9.5 0 0 0-1.259-1.71zm-8.123-2.51a8.5 8.5 0 0 1 3.778.002l.222-.974a9.5 9.5 0 0 0-4.222-.003zm.389 1.362v-1.85h-1v1.85zM7.25 6.37l-1.601-.925l-.5.866l1.6.925zm-2.5 6.928l-1.601.924l.5.866l1.6-.924zm.5-3.464l-1.6-.923l-.5.866l1.6.923zm5.25 10.94v-1.847h-1v1.847zm-3.75-4.012l-1.601.924l.5.866l1.6-.924zm12.101.925l-1.601-.925l-.5.866l1.601.925zm-4.351 3.09v-1.85h-1v1.85zM20.351 8.91l-1.601.924l.5.866l1.601-.924zm.498 5.311L19.25 13.3l-.5.866l1.6.923zM14.5 5.072V3.225h-1v1.847zm3.851.374l-1.601.925l.5.866l1.601-.925zM13.5 5.072c0 1.924 2.083 3.127 3.75 2.165l-.5-.866a1.5 1.5 0 0 1-2.25-1.3zm5.25 4.763c-1.667.962-1.667 3.368 0 4.33l.5-.866a1.5 1.5 0 0 1 0-2.598zm-1.5 6.928c-1.667-.962-3.75.24-3.75 2.165h1a1.5 1.5 0 0 1 2.25-1.299zm-6.75 2.165c0-1.924-2.083-3.127-3.75-2.165l.5.866a1.5 1.5 0 0 1 2.25 1.3zm-5.25-4.763c1.667-.962 1.667-3.368 0-4.33l-.5.866c1 .577 1 2.02 0 2.598zM9.5 5.072A1.5 1.5 0 0 1 7.25 6.37l-.5.866c1.667.962 3.75-.24 3.75-2.165z"/></g>`), g.Group(children),
	)
}

func ShareOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.272 6.008A2 2 0 0 0 20 5a2 2 0 1 0-3.728 1.008m0 0l-8.544 4.984m0 0A1.999 1.999 0 0 0 4 12a2 2 0 0 0 3.728 1.008m0-2.016a2 2 0 0 1 0 2.016m0 0l8.544 4.984m0 0A1.999 1.999 0 0 1 20 19a2 2 0 1 1-3.728-1.008"/>`), g.Group(children),
	)
}

func ShareOneBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M16.272 6.008A2 2 0 0 0 20 5a2 2 0 1 0-3.728 1.008m0 0l-8.544 4.984m0 0A1.999 1.999 0 0 0 4 12a2 2 0 0 0 3.728 1.008m0-2.016a2 2 0 0 1 0 2.016m0 0l8.544 4.984m0 0A1.999 1.999 0 0 1 20 19a2 2 0 1 1-3.728-1.008"/>`), g.Group(children),
	)
}

func ShareOneDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M16 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0m0 14a2 2 0 1 1 3.998-.002A2 2 0 0 1 16 19M4 12a2 2 0 1 1 4 0a2 2 0 0 1-4 0" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.272 6.008A2 2 0 0 0 20 5a2 2 0 1 0-3.728 1.008m0 0l-8.544 4.984m0 0a2 2 0 1 0 0 2.016m0-2.016a2 2 0 0 1 0 2.016m0 0l8.544 4.984m0 0a2 2 0 1 1 3.454 2.014a2 2 0 0 1-3.454-2.014"/></g>`), g.Group(children),
	)
}

func ShareOneFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M18 2a3 3 0 0 0-2.947 3.562l-7.114 4.15a3 3 0 1 0 0 4.578l7.114 4.148Q15 18.712 15 19a3 3 0 1 0 1.06-2.289l-7.113-4.15a3 3 0 0 0 0-1.123l7.114-4.15A3 3 0 1 0 18 2" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShareOneLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.272 6.008A2 2 0 0 0 20 5a2 2 0 1 0-3.728 1.008m0 0l-8.544 4.984m0 0a2 2 0 1 0 0 2.016m0-2.016a2 2 0 0 1 0 2.016m0 0l8.544 4.984m0 0a2 2 0 1 1 3.454 2.014a2 2 0 0 1-3.454-2.014"/>`), g.Group(children),
	)
}

func ShareOneThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.272 6.008A2 2 0 0 0 20 5a2 2 0 1 0-3.728 1.008m0 0l-8.544 4.984m0 0a2 2 0 1 0 0 2.016m0-2.016a2 2 0 0 1 0 2.016m0 0l8.544 4.984m0 0a2 2 0 1 1 3.454 2.014a2 2 0 0 1-3.454-2.014"/>`), g.Group(children),
	)
}

func ShareTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 5l-3-3m0 0L9 5m3-3v12M6 9H4v9a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9h-2"/>`), g.Group(children),
	)
}

func ShareTwoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m15 5l-3-3m0 0L9 5m3-3v12M6 9H4v9a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9h-2"/>`), g.Group(children),
	)
}

func ShareTwoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 5l-3-3m0 0L9 5m3-3v12M6 9H4v9a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9h-2"/>`), g.Group(children),
	)
}

func ShareTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11 14a1 1 0 1 0 2 0V6h2a1 1 0 0 0 .707-1.707l-3-3a1 1 0 0 0-1.414 0l-3 3A1 1 0 0 0 9 6h2zM4 8a1 1 0 0 0-1 1v9a3 3 0 0 0 3 3h12a3 3 0 0 0 3-3V9a1 1 0 0 0-1-1h-2a1 1 0 1 0 0 2h1v8a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1v-8h1a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShareTwoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 5l-3-3m0 0L9 5m3-3v12M6 9H4v9a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9h-2"/>`), g.Group(children),
	)
}

func ShareTwoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15 5l-3-3m0 0L9 5m3-3v12M6 9H4v9a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V9h-2"/>`), g.Group(children),
	)
}

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 3l.394-.92a1 1 0 0 0-.788 0zm0 18l-.496.868a1 1 0 0 0 .992 0zm6.394-15.26L18 6.66zM8.024 18.727l-.497.869zm3.582-16.646L5.212 4.82L6 6.66l6.394-2.74zM4 6.659v6.86h2v-6.86zm3.527 12.937l3.977 2.272l.992-1.736l-3.977-2.273zm4.97 2.272l3.976-2.272l-.992-1.737l-3.977 2.273zM20 13.518V6.66h-2v6.86zm-1.212-8.697l-6.394-2.74l-.788 1.838L18 6.66zM20 6.66a2 2 0 0 0-1.212-1.838L18 6.66zm-3.527 12.937A7 7 0 0 0 20 13.518h-2a5 5 0 0 1-2.52 4.341zM4 13.518a7 7 0 0 0 3.527 6.078l.992-1.737A5 5 0 0 1 6 13.52zm1.212-8.697A2 2 0 0 0 4 6.66h2z"/>`), g.Group(children),
	)
}

func ShieldBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 3l.492-1.149a1.25 1.25 0 0 0-.984 0zm0 18l-.62 1.085c.384.22.856.22 1.24 0zm6.394-15.26l-.492 1.15zM8.024 18.727l-.621 1.086zm3.484-16.876l-6.394 2.74l.984 2.298l6.394-2.74zM3.75 6.66v6.86h2.5V6.66zm3.653 13.154l3.977 2.272l1.24-2.17l-3.977-2.273zm5.217 2.272l3.977-2.272l-1.24-2.17l-3.977 2.272zm7.63-8.567V6.66h-2.5v6.86zm-1.364-8.927l-6.394-2.74l-.984 2.298l6.393 2.74zM20.25 6.66c0-.9-.536-1.713-1.364-2.068l-.985 2.298a.25.25 0 0 1-.151-.23zm-3.653 13.154a7.25 7.25 0 0 0 3.653-6.295h-2.5a4.75 4.75 0 0 1-2.393 4.124zM3.75 13.518a7.25 7.25 0 0 0 3.653 6.295l1.24-2.17a4.75 4.75 0 0 1-2.393-4.125zm1.364-8.927A2.25 2.25 0 0 0 3.75 6.66h2.5c0 .1-.06.19-.152.23z"/>`), g.Group(children),
	)
}

func ShieldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M5.606 5.74L12 3l6.394 2.74a1 1 0 0 1 .606.92v6.858a6 6 0 0 1-3.023 5.21L12 21l-3.977-2.273A6 6 0 0 1 5 13.519V6.66a1 1 0 0 1 .606-.919" opacity=".16"/><path d="m12 3l.394-.92a1 1 0 0 0-.788 0zm0 18l-.496.868a1 1 0 0 0 .992 0zm6.394-15.26L18 6.66zM8.024 18.727l-.497.869zm3.582-16.646L5.212 4.82L6 6.66l6.394-2.74zM4 6.659v6.86h2v-6.86zm3.527 12.937l3.977 2.272l.992-1.736l-3.977-2.273zm4.97 2.272l3.976-2.272l-.992-1.737l-3.977 2.273zM20 13.518V6.66h-2v6.86zm-1.212-8.697l-6.394-2.74l-.788 1.838L18 6.66zM20 6.66a2 2 0 0 0-1.212-1.838L18 6.66zm-3.527 12.937A7 7 0 0 0 20 13.518h-2a5 5 0 0 1-2.52 4.341zM4 13.518a7 7 0 0 0 3.527 6.078l.992-1.737A5 5 0 0 1 6 13.52zm1.212-8.697A2 2 0 0 0 4 6.66h2z"/></g>`), g.Group(children),
	)
}

func ShieldFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.394 2.08a1 1 0 0 0-.788 0L5.212 4.822A2 2 0 0 0 4 6.66v6.86a7 7 0 0 0 3.527 6.077l3.977 2.272a1 1 0 0 0 .992 0l3.977-2.272A7 7 0 0 0 20 13.518V6.66a2 2 0 0 0-1.212-1.838z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShieldLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 3l.295-.69a.75.75 0 0 0-.59 0zm0 18l-.372.651a.75.75 0 0 0 .744 0zm6.394-15.26l-.295.69zM8.024 18.727l-.373.652zm3.68-16.416L5.312 5.05L5.9 6.43l6.394-2.74zM4.25 6.659v6.86h1.5v-6.86zm3.401 12.72l3.977 2.272l.744-1.302l-3.977-2.273zm4.721 2.272l3.977-2.272l-.744-1.303l-3.977 2.273zm7.378-8.133V6.66h-1.5v6.86zm-1.06-8.467l-6.395-2.74l-.59 1.378l6.394 2.74zm1.06 1.608c0-.7-.417-1.332-1.06-1.608l-.591 1.379a.25.25 0 0 1 .151.23zm-3.401 12.72a6.75 6.75 0 0 0 3.401-5.86h-1.5a5.25 5.25 0 0 1-2.645 4.557zM4.25 13.519a6.75 6.75 0 0 0 3.401 5.86l.744-1.303a5.25 5.25 0 0 1-2.645-4.558zM5.31 5.05a1.75 1.75 0 0 0-1.06 1.608h1.5c0-.1.06-.19.152-.23z"/>`), g.Group(children),
	)
}

func ShieldNo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l.394-.92a1 1 0 0 0-.788 0zm0 18l-.496.868a1 1 0 0 0 .992 0zm6.394-15.26L18 6.66zM8.024 18.727l-.497.869zm3.582-16.646L5.212 4.82L6 6.66l6.394-2.74zM4 6.659v6.86h2v-6.86zm3.527 12.937l3.977 2.272l.992-1.736l-3.977-2.273zm4.97 2.272l3.976-2.272l-.992-1.737l-3.977 2.273zM20 13.518V6.66h-2v6.86zm-1.212-8.697l-6.394-2.74l-.788 1.838L18 6.66zM20 6.66a2 2 0 0 0-1.212-1.838L18 6.66zm-3.527 12.937A7 7 0 0 0 20 13.518h-2a5 5 0 0 1-2.52 4.341zM4 13.518a7 7 0 0 0 3.527 6.078l.992-1.737A5 5 0 0 1 6 13.52zm1.212-8.697A2 2 0 0 0 4 6.66h2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 10l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func ShieldNoBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l.492-1.149a1.25 1.25 0 0 0-.984 0zm0 18l-.62 1.085c.384.22.856.22 1.24 0zm6.394-15.26l-.492 1.15zM8.024 18.727l-.621 1.086zm3.484-16.876l-6.394 2.74l.984 2.298l6.394-2.74zM3.75 6.66v6.86h2.5V6.66zm3.653 13.154l3.977 2.272l1.24-2.17l-3.977-2.273zm5.217 2.272l3.977-2.272l-1.24-2.17l-3.977 2.272zm7.63-8.567V6.66h-2.5v6.86zm-1.364-8.927l-6.394-2.74l-.984 2.298l6.393 2.74zM20.25 6.66c0-.9-.536-1.713-1.364-2.068l-.985 2.298a.25.25 0 0 1-.151-.23zm-3.653 13.154a7.25 7.25 0 0 0 3.653-6.295h-2.5a4.75 4.75 0 0 1-2.393 4.124zM3.75 13.518a7.25 7.25 0 0 0 3.653 6.295l1.24-2.17a4.75 4.75 0 0 1-2.393-4.125zm1.364-8.927A2.25 2.25 0 0 0 3.75 6.66h2.5c0 .1-.06.19-.152.23z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m10 10l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func ShieldNoDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M5.606 5.74L12 3l6.394 2.74a1 1 0 0 1 .606.92v6.858a6 6 0 0 1-3.023 5.21L12 21l-3.977-2.273A6 6 0 0 1 5 13.519V6.66a1 1 0 0 1 .606-.919" opacity=".16"/><path fill="currentColor" d="m12 3l.394-.92a1 1 0 0 0-.788 0zm0 18l-.496.868a1 1 0 0 0 .992 0zm6.394-15.26L18 6.66zM8.024 18.727l-.497.869zm3.582-16.646L5.212 4.82L6 6.66l6.394-2.74zM4 6.659v6.86h2v-6.86zm3.527 12.937l3.977 2.272l.992-1.736l-3.977-2.273zm4.97 2.272l3.976-2.272l-.992-1.737l-3.977 2.273zM20 13.518V6.66h-2v6.86zm-1.212-8.697l-6.394-2.74l-.788 1.838L18 6.66zM20 6.66a2 2 0 0 0-1.212-1.838L18 6.66zm-3.527 12.937A7 7 0 0 0 20 13.518h-2a5 5 0 0 1-2.52 4.341zM4 13.518a7 7 0 0 0 3.527 6.078l.992-1.737A5 5 0 0 1 6 13.52zm1.212-8.697A2 2 0 0 0 4 6.66h2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 10l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func ShieldNoFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.606 2.08a1 1 0 0 1 .788 0l6.394 2.741A2 2 0 0 1 20 6.66v6.86a7 7 0 0 1-3.527 6.077l-3.977 2.272a1 1 0 0 1-.992 0l-3.977-2.272A7 7 0 0 1 4 13.518V6.66a2 2 0 0 1 1.212-1.838zm-.899 7.213a1 1 0 0 0-1.414 1.414L10.586 12l-1.293 1.293a1 1 0 1 0 1.414 1.414L12 13.414l1.293 1.293a1 1 0 0 0 1.414-1.414L13.414 12l1.293-1.293a1 1 0 0 0-1.414-1.414L12 10.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShieldNoLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l.295-.69a.75.75 0 0 0-.59 0zm0 18l-.372.651a.75.75 0 0 0 .744 0zm6.394-15.26l-.295.69zM8.024 18.727l-.373.652zm3.68-16.416L5.312 5.05L5.9 6.43l6.394-2.74zM4.25 6.659v6.86h1.5v-6.86zm3.401 12.72l3.977 2.272l.744-1.302l-3.977-2.273zm4.721 2.272l3.977-2.272l-.744-1.303l-3.977 2.273zm7.378-8.133V6.66h-1.5v6.86zm-1.06-8.467l-6.395-2.74l-.59 1.378l6.394 2.74zm1.06 1.608c0-.7-.417-1.332-1.06-1.608l-.591 1.379a.25.25 0 0 1 .151.23zm-3.401 12.72a6.75 6.75 0 0 0 3.401-5.86h-1.5a5.25 5.25 0 0 1-2.645 4.557zM4.25 13.519a6.75 6.75 0 0 0 3.401 5.86l.744-1.303a5.25 5.25 0 0 1-2.645-4.558zM5.31 5.05a1.75 1.75 0 0 0-1.06 1.608h1.5c0-.1.06-.19.152-.23z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m10 10l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func ShieldNoThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l.197-.46a.5.5 0 0 0-.394 0zm0 18l-.248.434a.5.5 0 0 0 .496 0zm6.394-15.26l-.197.46zM8.024 18.727l-.249.435zM11.802 2.54L5.409 5.28l.394.92l6.394-2.74zM4.5 6.66v6.858h1V6.66zm3.275 12.502l3.977 2.272l.496-.868l-3.977-2.273zm4.473 2.272l3.977-2.272l-.496-.869l-3.977 2.273zm7.252-7.916V6.66h-1v6.86zm-.91-8.237l-6.393-2.74l-.394.919l6.394 2.74zm.91 1.378a1.5 1.5 0 0 0-.91-1.378l-.393.919a.5.5 0 0 1 .303.46zm-3.275 12.503a6.5 6.5 0 0 0 3.275-5.644h-1a5.5 5.5 0 0 1-2.771 4.775zM4.5 13.518a6.5 6.5 0 0 0 3.275 5.644l.496-.869A5.5 5.5 0 0 1 5.5 13.518zm.91-8.237a1.5 1.5 0 0 0-.91 1.378h1c0-.2.12-.38.303-.46z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10 10l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func ShieldOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2"><path stroke-linejoin="round" d="m5.7 5.7l-.094.04A1 1 0 0 0 5 6.66v6.858a6 6 0 0 0 3.023 5.21L12 21l3.977-2.273a6 6 0 0 0 1.517-1.233M9.66 4.003L12 3l6.394 2.74a1 1 0 0 1 .606.92v6.683"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func ShieldOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2.5"><path stroke-linejoin="round" d="m5.7 5.7l-.094.04A1 1 0 0 0 5 6.66v6.858a6 6 0 0 0 3.023 5.21L12 21l3.977-2.273a6 6 0 0 0 1.517-1.233M9.66 4.003L12 3l6.394 2.74a1 1 0 0 1 .606.92v6.683"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func ShieldOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M5 13.518V6.66a1 1 0 0 1 .606-.919L5.7 5.7l11.794 11.794a6 6 0 0 1-1.517 1.233L12 21l-3.977-2.273A6 6 0 0 1 5 13.519" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m5.7 5.7l-.094.04A1 1 0 0 0 5 6.66v6.858a6 6 0 0 0 3.023 5.21L12 21l3.977-2.273a6 6 0 0 0 1.517-1.233M9.66 4.003L12 3l6.394 2.74a1 1 0 0 1 .606.92v6.683"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func ShieldOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.606 2.08a1 1 0 0 1 .788 0l6.394 2.741A2 2 0 0 1 20 6.66v6.684a1 1 0 0 1-2 0V6.659l-6-2.571l-1.946.834a1 1 0 1 1-.788-1.838zM4.26 5.675c-.167.294-.26.632-.26.985v6.86a7 7 0 0 0 3.527 6.077l3.977 2.272a1 1 0 0 0 .992 0l3.977-2.272q.54-.31 1.009-.7l1.811 1.811a1 1 0 0 0 1.414-1.414L6.422 5.008l-.03-.03l-1.685-1.685a1 1 0 0 0-1.414 1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShieldOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m5.7 5.7l-.094.04A1 1 0 0 0 5 6.66v6.858a6 6 0 0 0 3.023 5.21L12 21l3.977-2.273a6 6 0 0 0 1.517-1.233M9.66 4.003L12 3l6.394 2.74a1 1 0 0 1 .606.92v6.683"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func ShieldOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round"><path stroke-linejoin="round" d="m5.7 5.7l-.094.04A1 1 0 0 0 5 6.66v6.858a6 6 0 0 0 3.023 5.21L12 21l3.977-2.273a6 6 0 0 0 1.517-1.233M9.66 4.003L12 3l6.394 2.74a1 1 0 0 1 .606.92v6.683"/><path d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func ShieldThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 3l.197-.46a.5.5 0 0 0-.394 0zm0 18l-.248.434a.5.5 0 0 0 .496 0zm6.394-15.26l-.197.46zM8.024 18.727l-.249.435zM11.802 2.54L5.409 5.28l.394.92l6.394-2.74zM4.5 6.66v6.858h1V6.66zm3.275 12.502l3.977 2.272l.496-.868l-3.977-2.273zm4.473 2.272l3.977-2.272l-.496-.869l-3.977 2.273zm7.252-7.916V6.66h-1v6.86zm-.91-8.237l-6.393-2.74l-.394.919l6.394 2.74zm.91 1.378a1.5 1.5 0 0 0-.91-1.378l-.393.919a.5.5 0 0 1 .303.46zm-3.275 12.503a6.5 6.5 0 0 0 3.275-5.644h-1a5.5 5.5 0 0 1-2.771 4.775zM4.5 13.518a6.5 6.5 0 0 0 3.275 5.644l.496-.869A5.5 5.5 0 0 1 5.5 13.518zm.91-8.237a1.5 1.5 0 0 0-.91 1.378h1c0-.2.12-.38.303-.46z"/>`), g.Group(children),
	)
}

func ShieldYes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l.394-.92a1 1 0 0 0-.788 0zm0 18l-.496.868a1 1 0 0 0 .992 0zm6.394-15.26L18 6.66zM8.024 18.727l-.497.869zm3.582-16.646L5.212 4.82L6 6.66l6.394-2.74zM4 6.659v6.86h2v-6.86zm3.527 12.937l3.977 2.272l.992-1.736l-3.977-2.273zm4.97 2.272l3.976-2.272l-.992-1.737l-3.977 2.273zM20 13.518V6.66h-2v6.86zm-1.212-8.697l-6.394-2.74l-.788 1.838L18 6.66zM20 6.66a2 2 0 0 0-1.212-1.838L18 6.66zm-3.527 12.937A7 7 0 0 0 20 13.518h-2a5 5 0 0 1-2.52 4.341zM4 13.518a7 7 0 0 0 3.527 6.078l.992-1.737A5 5 0 0 1 6 13.52zm1.212-8.697A2 2 0 0 0 4 6.66h2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func ShieldYesBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l.492-1.149a1.25 1.25 0 0 0-.984 0zm0 18l-.62 1.085c.384.22.856.22 1.24 0zm6.394-15.26l-.492 1.15zM8.024 18.727l-.621 1.086zm3.484-16.876l-6.394 2.74l.984 2.298l6.394-2.74zM3.75 6.66v6.86h2.5V6.66zm3.653 13.154l3.977 2.272l1.24-2.17l-3.977-2.273zm5.217 2.272l3.977-2.272l-1.24-2.17l-3.977 2.272zm7.63-8.567V6.66h-2.5v6.86zm-1.364-8.927l-6.394-2.74l-.984 2.298l6.393 2.74zM20.25 6.66c0-.9-.536-1.713-1.364-2.068l-.985 2.298a.25.25 0 0 1-.151-.23zm-3.653 13.154a7.25 7.25 0 0 0 3.653-6.295h-2.5a4.75 4.75 0 0 1-2.393 4.124zM3.75 13.518a7.25 7.25 0 0 0 3.653 6.295l1.24-2.17a4.75 4.75 0 0 1-2.393-4.125zm1.364-8.927A2.25 2.25 0 0 0 3.75 6.66h2.5c0 .1-.06.19-.152.23z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func ShieldYesDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M5.606 5.74L12 3l6.394 2.74a1 1 0 0 1 .606.92v6.858a6 6 0 0 1-3.023 5.21L12 21l-3.977-2.273A6 6 0 0 1 5 13.519V6.66a1 1 0 0 1 .606-.919" opacity=".16"/><path fill="currentColor" d="m12 3l.394-.92a1 1 0 0 0-.788 0zm0 18l-.496.868a1 1 0 0 0 .992 0zm6.394-15.26L18 6.66zM8.024 18.727l-.497.869zm3.582-16.646L5.212 4.82L6 6.66l6.394-2.74zM4 6.659v6.86h2v-6.86zm3.527 12.937l3.977 2.272l.992-1.736l-3.977-2.273zm4.97 2.272l3.976-2.272l-.992-1.737l-3.977 2.273zM20 13.518V6.66h-2v6.86zm-1.212-8.697l-6.394-2.74l-.788 1.838L18 6.66zM20 6.66a2 2 0 0 0-1.212-1.838L18 6.66zm-3.527 12.937A7 7 0 0 0 20 13.518h-2a5 5 0 0 1-2.52 4.341zM4 13.518a7 7 0 0 0 3.527 6.078l.992-1.737A5 5 0 0 1 6 13.52zm1.212-8.697A2 2 0 0 0 4 6.66h2z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func ShieldYesFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.606 2.08a1 1 0 0 1 .788 0l6.394 2.741A2 2 0 0 1 20 6.66v6.86a7 7 0 0 1-3.527 6.077l-3.977 2.272a1 1 0 0 1-.992 0l-3.977-2.272A7 7 0 0 1 4 13.518V6.66a2 2 0 0 1 1.212-1.838zm4.101 8.627a1 1 0 0 0-1.414-1.414L11 12.586l-1.293-1.293a1 1 0 0 0-1.414 1.414l2 2a1 1 0 0 0 1.414 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShieldYesLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l.295-.69a.75.75 0 0 0-.59 0zm0 18l-.372.651a.75.75 0 0 0 .744 0zm6.394-15.26l-.295.69zM8.024 18.727l-.373.652zm3.68-16.416L5.312 5.05L5.9 6.43l6.394-2.74zM4.25 6.659v6.86h1.5v-6.86zm3.401 12.72l3.977 2.272l.744-1.302l-3.977-2.273zm4.721 2.272l3.977-2.272l-.744-1.303l-3.977 2.273zm7.378-8.133V6.66h-1.5v6.86zm-1.06-8.467l-6.395-2.74l-.59 1.378l6.394 2.74zm1.06 1.608c0-.7-.417-1.332-1.06-1.608l-.591 1.379a.25.25 0 0 1 .151.23zm-3.401 12.72a6.75 6.75 0 0 0 3.401-5.86h-1.5a5.25 5.25 0 0 1-2.645 4.557zM4.25 13.519a6.75 6.75 0 0 0 3.401 5.86l.744-1.303a5.25 5.25 0 0 1-2.645-4.558zM5.31 5.05a1.75 1.75 0 0 0-1.06 1.608h1.5c0-.1.06-.19.152-.23z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func ShieldYesThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l.197-.46a.5.5 0 0 0-.394 0zm0 18l-.248.434a.5.5 0 0 0 .496 0zm6.394-15.26l-.197.46zM8.024 18.727l-.249.435zM11.802 2.54L5.409 5.28l.394.92l6.394-2.74zM4.5 6.66v6.858h1V6.66zm3.275 12.502l3.977 2.272l.496-.868l-3.977-2.273zm4.473 2.272l3.977-2.272l-.496-.869l-3.977 2.273zm7.252-7.916V6.66h-1v6.86zm-.91-8.237l-6.393-2.74l-.394.919l6.394 2.74zm.91 1.378a1.5 1.5 0 0 0-.91-1.378l-.393.919a.5.5 0 0 1 .303.46zm-3.275 12.503a6.5 6.5 0 0 0 3.275-5.644h-1a5.5 5.5 0 0 1-2.771 4.775zM4.5 13.518a6.5 6.5 0 0 0 3.275 5.644l.496-.869A5.5 5.5 0 0 1 5.5 13.518zm.91-8.237a1.5 1.5 0 0 0-.91 1.378h1c0-.2.12-.38.303-.46z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15 10l-4 4l-2-2"/></g>`), g.Group(children),
	)
}

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M4 9h16l-.835 9.181A2 2 0 0 1 17.174 20H6.826a2 2 0 0 1-1.991-1.819z"/><path stroke-linecap="round" d="M8 11V8a4 4 0 1 1 8 0v3"/></g>`), g.Group(children),
	)
}

func ShoppingBagBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><path stroke-linejoin="round" d="M4 9h16l-.835 9.181A2 2 0 0 1 17.174 20H6.826a2 2 0 0 1-1.991-1.819z"/><path stroke-linecap="round" d="M8 11V8a4 4 0 1 1 8 0v3"/></g>`), g.Group(children),
	)
}

func ShoppingBagDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 9h16l-.835 9.181A2 2 0 0 1 17.174 20H6.826a2 2 0 0 1-1.991-1.819z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M4 9h16l-.835 9.181A2 2 0 0 1 17.174 20H6.826a2 2 0 0 1-1.991-1.819z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M8 11V8a4 4 0 1 1 8 0v3"/></g>`), g.Group(children),
	)
}

func ShoppingBagFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 8a3 3 0 1 1 6 0zM7 8a5 5 0 0 1 10 0h3a1 1 0 0 1 .996 1.09l-.835 9.182A3 3 0 0 1 17.174 21H6.826a3 3 0 0 1-2.987-2.728L3.004 9.09A1 1 0 0 1 4 8z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShoppingBagLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linejoin="round" d="M4 9h16l-.835 9.181A2 2 0 0 1 17.174 20H6.826a2 2 0 0 1-1.991-1.819z"/><path stroke-linecap="round" d="M8 11V8a4 4 0 1 1 8 0v3"/></g>`), g.Group(children),
	)
}

func ShoppingBagThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linejoin="round" d="M4 9h16l-.835 9.181A2 2 0 0 1 17.174 20H6.826a2 2 0 0 1-1.991-1.819z"/><path stroke-linecap="round" d="M8 11V8a4 4 0 1 1 8 0v3"/></g>`), g.Group(children),
	)
}

func ShoppingCard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2a1 1 0 0 0 0 2zm2 1l.997-.077A1 1 0 0 0 5 2zm16 3l.994.11A1 1 0 0 0 21 5zM5.23 6l-.996.077zm13.109 9.119l.07.997zm-10.355.74l-.071-.998zM3 4h2V2H3zm5.055 12.856l10.355-.74l-.143-1.995l-10.354.74zm13.123-3.401l.816-7.345l-1.988-.22l-.816 7.344zM4.003 3.077l.23 3l1.995-.154l-.23-3zm.23 3l.617 8.017l1.995-.154l-.617-8.017zM21 5H5.23v2H21zm-2.59 11.116a3 3 0 0 0 2.768-2.661l-1.988-.22a1 1 0 0 1-.923.886zM7.913 14.861a1 1 0 0 1-1.069-.92l-1.994.152a3 3 0 0 0 3.205 2.763z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M8.5 20.5h.01v.01H8.5zm9 0h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func ShoppingCardAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2a1 1 0 0 0 0 2zm2 1l.997-.077A1 1 0 0 0 5 2zm16 3l.994.11A1 1 0 0 0 21 5zM5.23 6l-.996.077zm13.109 9.119l.07.997zm-10.355.74l-.071-.998zM3 4h2V2H3zm5.055 12.856l10.355-.74l-.143-1.995l-10.354.74zm13.123-3.401l.816-7.345l-1.988-.22l-.816 7.344zM4.003 3.077l.23 3l1.995-.154l-.23-3zm.23 3l.617 8.017l1.995-.154l-.617-8.017zM21 5H5.23v2H21zm-2.59 11.116a3 3 0 0 0 2.768-2.661l-1.988-.22a1 1 0 0 1-.923.886zM7.913 14.861a1 1 0 0 1-1.069-.92l-1.994.152a3 3 0 0 0 3.205 2.763z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 9.5v3M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`), g.Group(children),
	)
}

func ShoppingCardAddBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 1.75a1.25 1.25 0 1 0 0 2.5zM5 3l1.246-.096A1.25 1.25 0 0 0 5 1.75zm16 3l1.242.138A1.25 1.25 0 0 0 21 4.75zM5.23 6l-1.246.096zm13.109 9.119l.089 1.247zm-10.355.74l-.089-1.248zM3 4.25h2v-2.5H3zm5.073 12.855l10.355-.74l-.178-2.493l-10.355.74zm13.353-3.622l.816-7.345l-2.484-.276l-.816 7.345zM3.754 3.096l.23 3l2.493-.192l-.23-3zm.23 3l.617 8.017l2.493-.192l-.617-8.017zM21 4.75H5.23v2.5H21zm-2.572 11.616a3.25 3.25 0 0 0 2.998-2.883l-2.484-.276a.75.75 0 0 1-.692.665zM7.895 14.61a.75.75 0 0 1-.801-.69l-2.493.192a3.25 3.25 0 0 0 3.472 2.992z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M13 9.5v3M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3.75" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`), g.Group(children),
	)
}

func ShoppingCardAddDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m18.339 15.119l-10.355.74a2 2 0 0 1-2.137-1.842L5.231 6H21l-.816 7.345a2 2 0 0 1-1.845 1.774" opacity=".16"/><path fill="currentColor" d="M3 2a1 1 0 0 0 0 2zm2 1l.997-.077A1 1 0 0 0 5 2zm16 3l.994.11A1 1 0 0 0 21 5zM5.23 6l-.996.077zm13.109 9.119l.07.997zm-10.355.74l-.071-.998zM3 4h2V2H3zm5.055 12.856l10.355-.74l-.143-1.995l-10.354.74zm13.123-3.401l.816-7.345l-1.988-.22l-.816 7.344zM4.003 3.077l.23 3l1.995-.154l-.23-3zm.23 3l.617 8.017l1.995-.154l-.617-8.017zM21 5H5.23v2H21zm-2.59 11.116a3 3 0 0 0 2.768-2.661l-1.988-.22a1 1 0 0 1-.923.886zM7.913 14.861a1 1 0 0 1-1.069-.92l-1.994.152a3 3 0 0 0 3.205 2.763z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M13 9.5v3M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`), g.Group(children),
	)
}

func ShoppingCardAddFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 3a1 1 0 0 1 1-1h2a1 1 0 0 1 .997.923L6.157 5H21a1 1 0 0 1 .994 1.11l-.816 7.345a3 3 0 0 1-2.768 2.661l-10.355.74a3 3 0 0 1-3.205-2.763l-.616-8.016L4.074 4H3a1 1 0 0 1-1-1m11 5.5a1 1 0 0 1 1 1v.5h.5a1 1 0 1 1 0 2H14v.5a1 1 0 1 1-2 0V12h-.5a1 1 0 1 1 0-2h.5v-.5a1 1 0 0 1 1-1m3 12a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5h-.01a1.5 1.5 0 0 1-1.5-1.5zM8.5 19A1.5 1.5 0 0 0 7 20.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShoppingCardAddLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2.25a.75.75 0 0 0 0 1.5zM5 3l.748-.058A.75.75 0 0 0 5 2.25zm16 3l.745.083A.75.75 0 0 0 21 5.25zM5.23 6l-.747.058zm13.109 9.119l.053.748zm-10.355.74l-.053-.749zM3 3.75h2v-1.5H3zm5.037 12.856l10.355-.74l-.107-1.495l-10.354.74zm12.892-3.179l.816-7.344l-1.49-.166l-.816 7.345zM4.252 3.057l.231 3l1.496-.115l-.231-3zm.231 3l.617 8.017l1.495-.115l-.616-8.017zM21 5.25H5.23v1.5H21zm-2.608 10.617a2.75 2.75 0 0 0 2.537-2.44l-1.49-.165a1.25 1.25 0 0 1-1.154 1.109zM7.931 15.11a1.25 1.25 0 0 1-1.336-1.15l-1.495.114a2.75 2.75 0 0 0 2.937 2.532z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13 9.5v3M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2.25" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`), g.Group(children),
	)
}

func ShoppingCardAddThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2.5a.5.5 0 0 0 0 1zM5 3l.499-.038A.5.5 0 0 0 5 2.5zm16 3l.497.055A.5.5 0 0 0 21 5.5zM5.23 6l-.498.038zm13.109 9.119l.035.498zm-10.355.74l-.036-.5zM3 3.5h2v-1H3zm5.02 12.857l10.354-.74l-.071-.997l-10.355.74zM20.68 13.4l.817-7.345l-.994-.11l-.816 7.344zM4.502 3.038l.231 3l.997-.076l-.23-3zm.231 3l.617 8.017l.997-.077l-.617-8.016zM21 5.5H5.23v1H21zm-2.626 10.117a2.5 2.5 0 0 0 2.307-2.217l-.994-.11a1.5 1.5 0 0 1-1.384 1.33zM7.948 15.36a1.5 1.5 0 0 1-1.602-1.382l-.997.077a2.5 2.5 0 0 0 2.67 2.302z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13 9.5v3M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`), g.Group(children),
	)
}

func ShoppingCardBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 1.75a1.25 1.25 0 1 0 0 2.5zM5 3l1.246-.096A1.25 1.25 0 0 0 5 1.75zm16 3l1.242.138A1.25 1.25 0 0 0 21 4.75zM5.23 6l-1.246.096zm13.109 9.119l.089 1.247zm-10.355.74l-.089-1.248zM3 4.25h2v-2.5H3zm5.073 12.855l10.355-.74l-.178-2.493l-10.355.74zm13.353-3.622l.816-7.345l-2.484-.276l-.816 7.345zM3.754 3.096l.23 3l2.493-.192l-.23-3zm.23 3l.617 8.017l2.493-.192l-.617-8.017zM21 4.75H5.23v2.5H21zm-2.572 11.616a3.25 3.25 0 0 0 2.998-2.883l-2.484-.276a.75.75 0 0 1-.692.665zM7.895 14.61a.75.75 0 0 1-.801-.69l-2.493.192a3.25 3.25 0 0 0 3.472 2.992z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3.75" d="M8.5 20.5h.01v.01H8.5zm9 0h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func ShoppingCardDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m18.339 15.119l-10.355.74a2 2 0 0 1-2.137-1.842L5.231 6H21l-.816 7.345a2 2 0 0 1-1.845 1.774" opacity=".16"/><path fill="currentColor" d="M3 2a1 1 0 0 0 0 2zm2 1l.997-.077A1 1 0 0 0 5 2zm16 3l.994.11A1 1 0 0 0 21 5zM5.23 6l-.996.077zm13.109 9.119l.07.997zm-10.355.74l-.071-.998zM3 4h2V2H3zm5.055 12.856l10.355-.74l-.143-1.995l-10.354.74zm13.123-3.401l.816-7.345l-1.988-.22l-.816 7.344zM4.003 3.077l.23 3l1.995-.154l-.23-3zm.23 3l.617 8.017l1.995-.154l-.617-8.017zM21 5H5.23v2H21zm-2.59 11.116a3 3 0 0 0 2.768-2.661l-1.988-.22a1 1 0 0 1-.923.886zM7.913 14.861a1 1 0 0 1-1.069-.92l-1.994.152a3 3 0 0 0 3.205 2.763z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M8.5 20.5h.01v.01H8.5zm9 0h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func ShoppingCardFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 2a1 1 0 0 0 0 2h1.074l.16 2.077l.616 8.017a3 3 0 0 0 3.205 2.762l10.355-.74a3 3 0 0 0 2.768-2.661l.816-7.345A1 1 0 0 0 21 5H6.157l-.16-2.077A1 1 0 0 0 5 2zm4 18.5A1.5 1.5 0 0 1 8.5 19h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H8.5a1.5 1.5 0 0 1-1.5-1.5zM17.5 19a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShoppingCardLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2.25a.75.75 0 0 0 0 1.5zM5 3l.748-.058A.75.75 0 0 0 5 2.25zm16 3l.745.083A.75.75 0 0 0 21 5.25zM5.23 6l-.747.058zm13.109 9.119l.053.748zm-10.355.74l-.053-.749zM3 3.75h2v-1.5H3zm5.037 12.856l10.355-.74l-.107-1.495l-10.354.74zm12.892-3.179l.816-7.344l-1.49-.166l-.816 7.345zM4.252 3.057l.231 3l1.496-.115l-.231-3zm.231 3l.617 8.017l1.495-.115l-.616-8.017zM21 5.25H5.23v1.5H21zm-2.608 10.617a2.75 2.75 0 0 0 2.537-2.44l-1.49-.165a1.25 1.25 0 0 1-1.154 1.109zM7.931 15.11a1.25 1.25 0 0 1-1.336-1.15l-1.495.114a2.75 2.75 0 0 0 2.937 2.532z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2.25" d="M8.5 20.5h.01v.01H8.5zm9 0h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func ShoppingCardRemove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2a1 1 0 0 0 0 2zm2 1l.997-.077A1 1 0 0 0 5 2zm16 3l.994.11A1 1 0 0 0 21 5zM5.23 6l-.996.077zm13.109 9.119l.07.997zm-10.355.74l-.071-.998zM3 4h2V2H3zm5.055 12.856l10.355-.74l-.143-1.995l-10.354.74zm13.123-3.401l.816-7.345l-1.988-.22l-.816 7.344zM4.003 3.077l.23 3l1.995-.154l-.23-3zm.23 3l.617 8.017l1.995-.154l-.617-8.017zM21 5H5.23v2H21zm-2.59 11.116a3 3 0 0 0 2.768-2.661l-1.988-.22a1 1 0 0 1-.923.886zM7.913 14.861a1 1 0 0 1-1.069-.92l-1.994.152a3 3 0 0 0 3.205 2.763z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`), g.Group(children),
	)
}

func ShoppingCardRemoveBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 1.75a1.25 1.25 0 1 0 0 2.5zM5 3l1.246-.096A1.25 1.25 0 0 0 5 1.75zm16 3l1.242.138A1.25 1.25 0 0 0 21 4.75zM5.23 6l-1.246.096zm13.109 9.119l.089 1.247zm-10.355.74l-.089-1.248zM3 4.25h2v-2.5H3zm5.073 12.855l10.355-.74l-.178-2.493l-10.355.74zm13.353-3.622l.816-7.345l-2.484-.276l-.816 7.345zM3.754 3.096l.23 3l2.493-.192l-.23-3zm.23 3l.617 8.017l2.493-.192l-.617-8.017zM21 4.75H5.23v2.5H21zm-2.572 11.616a3.25 3.25 0 0 0 2.998-2.883l-2.484-.276a.75.75 0 0 1-.692.665zM7.895 14.61a.75.75 0 0 1-.801-.69l-2.493.192a3.25 3.25 0 0 0 3.472 2.992z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3.75" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`), g.Group(children),
	)
}

func ShoppingCardRemoveDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m18.339 15.119l-10.355.74a2 2 0 0 1-2.137-1.842L5.231 6H21l-.816 7.345a2 2 0 0 1-1.845 1.774" opacity=".16"/><path fill="currentColor" d="M3 2a1 1 0 0 0 0 2zm2 1l.997-.077A1 1 0 0 0 5 2zm16 3l.994.11A1 1 0 0 0 21 5zM5.23 6l-.996.077zm13.109 9.119l.07.997zm-10.355.74l-.071-.998zM3 4h2V2H3zm5.055 12.856l10.355-.74l-.143-1.995l-10.354.74zm13.123-3.401l.816-7.345l-1.988-.22l-.816 7.344zM4.003 3.077l.23 3l1.995-.154l-.23-3zm.23 3l.617 8.017l1.995-.154l-.617-8.017zM21 5H5.23v2H21zm-2.59 11.116a3 3 0 0 0 2.768-2.661l-1.988-.22a1 1 0 0 1-.923.886zM7.913 14.861a1 1 0 0 1-1.069-.92l-1.994.152a3 3 0 0 0 3.205 2.763z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`), g.Group(children),
	)
}

func ShoppingCardRemoveFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 3a1 1 0 0 1 1-1h2a1 1 0 0 1 .997.923L6.157 5H21a1 1 0 0 1 .994 1.11l-.816 7.345a3 3 0 0 1-2.768 2.661l-10.355.74a3 3 0 0 1-3.205-2.763l-.616-8.016L4.074 4H3a1 1 0 0 1-1-1m9.5 7a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2zM16 20.5a1.5 1.5 0 0 1 1.5-1.5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5h-.01a1.5 1.5 0 0 1-1.5-1.5zM8.5 19A1.5 1.5 0 0 0 7 20.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShoppingCardRemoveLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2.25a.75.75 0 0 0 0 1.5zM5 3l.748-.058A.75.75 0 0 0 5 2.25zm16 3l.745.083A.75.75 0 0 0 21 5.25zM5.23 6l-.747.058zm13.109 9.119l.053.748zm-10.355.74l-.053-.749zM3 3.75h2v-1.5H3zm5.037 12.856l10.355-.74l-.107-1.495l-10.354.74zm12.892-3.179l.816-7.344l-1.49-.166l-.816 7.345zM4.252 3.057l.231 3l1.496-.115l-.231-3zm.231 3l.617 8.017l1.495-.115l-.616-8.017zM21 5.25H5.23v1.5H21zm-2.608 10.617a2.75 2.75 0 0 0 2.537-2.44l-1.49-.165a1.25 1.25 0 0 1-1.154 1.109zM7.931 15.11a1.25 1.25 0 0 1-1.336-1.15l-1.495.114a2.75 2.75 0 0 0 2.937 2.532z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2.25" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`), g.Group(children),
	)
}

func ShoppingCardRemoveThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2.5a.5.5 0 0 0 0 1zM5 3l.499-.038A.5.5 0 0 0 5 2.5zm16 3l.497.055A.5.5 0 0 0 21 5.5zM5.23 6l-.498.038zm13.109 9.119l.035.498zm-10.355.74l-.036-.5zM3 3.5h2v-1H3zm5.02 12.857l10.354-.74l-.071-.997l-10.355.74zM20.68 13.4l.817-7.345l-.994-.11l-.816 7.344zM4.502 3.038l.231 3l.997-.076l-.23-3zm.231 3l.617 8.017l.997-.077l-.617-8.016zM21 5.5H5.23v1H21zm-2.626 10.117a2.5 2.5 0 0 0 2.307-2.217l-.994-.11a1.5 1.5 0 0 1-1.384 1.33zM7.948 15.36a1.5 1.5 0 0 1-1.602-1.382l-.997.077a2.5 2.5 0 0 0 2.67 2.302z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.5 11h3"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M17.5 20.5h.01v.01h-.01zm-9 0h.01v.01H8.5z"/></g>`), g.Group(children),
	)
}

func ShoppingCardThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M3 2.5a.5.5 0 0 0 0 1zM5 3l.499-.038A.5.5 0 0 0 5 2.5zm16 3l.497.055A.5.5 0 0 0 21 5.5zM5.23 6l-.498.038zm13.109 9.119l.035.498zm-10.355.74l-.036-.5zM3 3.5h2v-1H3zm5.02 12.857l10.354-.74l-.071-.997l-10.355.74zM20.68 13.4l.817-7.345l-.994-.11l-.816 7.344zM4.502 3.038l.231 3l.997-.076l-.23-3zm.231 3l.617 8.017l.997-.077l-.617-8.016zM21 5.5H5.23v1H21zm-2.626 10.117a2.5 2.5 0 0 0 2.307-2.217l-.994-.11a1.5 1.5 0 0 1-1.384 1.33zM7.948 15.36a1.5 1.5 0 0 1-1.602-1.382l-.997.077a2.5 2.5 0 0 0 2.67 2.302z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M8.5 20.5h.01v.01H8.5zm9 0h.01v.01h-.01z"/></g>`), g.Group(children),
	)
}

func SignDivision(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="3" d="M12 6.5h.01v.01H12zm0 11h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M18 12H6"/></g>`), g.Group(children),
	)
}

func SignDivisionBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="3.75" d="M12 6.5h.01v.01H12zm0 11h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M18 12H6"/></g>`), g.Group(children),
	)
}

func SignDivisionCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="11.999" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M12 8h.01v.01H12zm0 8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M15 12H9"/></g>`), g.Group(children),
	)
}

func SignDivisionCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="11.999" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M12 8h.01v.01H12zm0 8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M15 12H9"/></g>`), g.Group(children),
	)
}

func SignDivisionCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="11.999" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="11.999" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 8h.01v.01H12zm0 8h.01v.01H12z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12H9"/></g>`), g.Group(children),
	)
}

func SignDivisionCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 11.999c0-5.523 4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10s-10-4.477-10-10M12 6.5A1.5 1.5 0 0 0 10.5 8v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V8a1.5 1.5 0 0 0-1.5-1.5zm0 8a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V16a1.5 1.5 0 0 0-1.5-1.5zM8 12a1 1 0 0 1 1-1h6a1 1 0 0 1 0 2H9a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignDivisionCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="11.999" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M12 8h.01v.01H12zm0 8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M15 12H9"/></g>`), g.Group(children),
	)
}

func SignDivisionCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="11.999" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M12 8h.01v.01H12zm0 8h.01v.01H12z"/><path stroke-linecap="round" d="M15 12H9"/></g>`), g.Group(children),
	)
}

func SignDivisionDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="3" d="M12 6.5h.01v.01H12zm0 11h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M18 12H6"/></g>`), g.Group(children),
	)
}

func SignDivisionFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.5 6.5A1.5 1.5 0 0 1 12 5h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5zm0 11A1.5 1.5 0 0 1 12 16h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H12a1.5 1.5 0 0 1-1.5-1.5zM6 11a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignDivisionLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="2.25" d="M12 6.5h.01v.01H12zm0 11h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M18 12H6"/></g>`), g.Group(children),
	)
}

func SignDivisionSlash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 4L8 20"/>`), g.Group(children),
	)
}

func SignDivisionSlashBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M16 4L8 20"/>`), g.Group(children),
	)
}

func SignDivisionSlashFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.447 3.106a1 1 0 0 1 .447 1.341l-8 16a1 1 0 1 1-1.788-.894l8-16a1 1 0 0 1 1.341-.447" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignDivisionSlashLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 4L8 20"/>`), g.Group(children),
	)
}

func SignDivisionSlashThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 4L8 20"/>`), g.Group(children),
	)
}

func SignDivisionSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3" d="M12 16h.01v.01H12zm0-8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M15 12H9"/></g>`), g.Group(children),
	)
}

func SignDivisionSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="3.75" d="M12 16h.01v.01H12zm0-8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M15 12H9"/></g>`), g.Group(children),
	)
}

func SignDivisionSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M12 8h.01v.01H12zm0 8h.01v.01H12z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15 12H9"/></g>`), g.Group(children),
	)
}

func SignDivisionSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm9 2.5A1.5 1.5 0 0 0 10.5 8v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V8a1.5 1.5 0 0 0-1.5-1.5zm0 8a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V16a1.5 1.5 0 0 0-1.5-1.5zM8 12a1 1 0 0 1 1-1h6a1 1 0 0 1 0 2H9a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignDivisionSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="2.25" d="M12 8h.01v.01H12zm0 8h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M15 12H9"/></g>`), g.Group(children),
	)
}

func SignDivisionSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-linecap="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/><path stroke-width="1.5" d="M12 8h.01v.01H12zm0 8h.01v.01H12z"/><path stroke-linecap="round" d="M15 12H9"/></g>`), g.Group(children),
	)
}

func SignDivisionThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="1.5" d="M12 6.5h.01v.01H12zm0 11h.01v.01H12z"/><path stroke-linecap="round" d="M18 12H6"/></g>`), g.Group(children),
	)
}

func SignEqual(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 9h12M6 15h12"/>`), g.Group(children),
	)
}

func SignEqualBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M6 9h12M6 15h12"/>`), g.Group(children),
	)
}

func SignEqualFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 8a1 1 0 0 0 0 2h12a1 1 0 1 0 0-2zm0 6a1 1 0 1 0 0 2h12a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignEqualLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 9h12M6 15h12"/>`), g.Group(children),
	)
}

func SignEqualThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 9h12M6 15h12"/>`), g.Group(children),
	)
}

func SignF(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 5v-.295C17 3.763 16.237 3 15.295 3v0a1.73 1.73 0 0 0-1.66 1.242L9.557 18.105A2.64 2.64 0 0 1 7.025 20H7m3.5-11H14"/>`), g.Group(children),
	)
}

func SignFBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M17 5v-.295C17 3.763 16.237 3 15.295 3v0a1.73 1.73 0 0 0-1.66 1.242L9.557 18.105A2.64 2.64 0 0 1 7.025 20H7m3.5-11H14"/>`), g.Group(children),
	)
}

func SignFFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.295 4c.39 0 .705.316.705.705V5a1 1 0 1 0 2 0v-.295A2.705 2.705 0 0 0 15.295 2a2.73 2.73 0 0 0-2.62 1.96L11.487 8H10.5a1 1 0 0 0 0 2h.399l-2.301 7.823A1.64 1.64 0 0 1 7.025 19H7a1 1 0 1 0 0 2h.025a3.64 3.64 0 0 0 3.492-2.613L12.984 10H14a1 1 0 1 0 0-2h-.428l1.022-3.475A.73.73 0 0 1 15.295 4" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignFLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 5v-.295C17 3.763 16.237 3 15.295 3v0a1.73 1.73 0 0 0-1.66 1.242L9.557 18.105A2.64 2.64 0 0 1 7.025 20H7m3.5-11H14"/>`), g.Group(children),
	)
}

func SignFThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 5v-.295C17 3.763 16.237 3 15.295 3v0a1.73 1.73 0 0 0-1.66 1.242L9.557 18.105A2.64 2.64 0 0 1 7.025 20H7m3.5-11H14"/>`), g.Group(children),
	)
}

func SignFactorial(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="3" d="M12 19.5h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2" d="M11.523 4.5L12 15l.477-10.5a.478.478 0 1 0-.954 0"/></g>`), g.Group(children),
	)
}

func SignFactorialBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="3.75" d="M12 19.5h.01v.01H12z"/><path stroke-linecap="round" stroke-width="2.5" d="M11.523 4.5L12 15l.477-10.5a.478.478 0 1 0-.954 0"/></g>`), g.Group(children),
	)
}

func SignFactorialFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.001 15.046a1 1 0 0 0 1.998 0l.477-10.501a1.478 1.478 0 1 0-2.952 0zM12 18a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignFactorialLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="2.25" d="M12 19.5h.01v.01H12z"/><path stroke-linecap="round" stroke-width="1.5" d="M11.523 4.5L12 15l.477-10.5a.478.478 0 1 0-.954 0"/></g>`), g.Group(children),
	)
}

func SignFactorialThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="1.5" d="M12 19.5h.01v.01H12z"/><path stroke-linecap="round" d="M11.523 4.5L12 15l.477-10.5a.478.478 0 1 0-.954 0"/></g>`), g.Group(children),
	)
}

func SignLemniscate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M17 17c-2.761 0-6-5-6-5s3.239-5 6-5a5 5 0 0 1 0 10ZM6 8c2.21 0 5 4 5 4s-2.79 4-5 4a4 4 0 0 1 0-8Z"/>`), g.Group(children),
	)
}

func SignLemniscateBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="M17 17c-2.761 0-6-5-6-5s3.239-5 6-5a5 5 0 0 1 0 10ZM6 8c2.21 0 5 4 5 4s-2.79 4-5 4a4 4 0 0 1 0-8Z"/>`), g.Group(children),
	)
}

func SignLemniscateDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M17 17c-2.761 0-6-5-6-5s3.239-5 6-5a5 5 0 0 1 0 10" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M17 17c-2.761 0-6-5-6-5s3.239-5 6-5a5 5 0 0 1 0 10ZM6 8c2.21 0 5 4 5 4s-2.79 4-5 4a4 4 0 0 1 0-8Z"/></g>`), g.Group(children),
	)
}

func SignLemniscateFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 8a4 4 0 0 1 0 8c-.402 0-.922-.19-1.542-.61c-.604-.409-1.21-.974-1.759-1.569A20 20 0 0 1 12.217 12a20 20 0 0 1 1.482-1.821c.55-.595 1.155-1.16 1.76-1.57C16.077 8.19 16.598 8 17 8m-6.032 5.672a18 18 0 0 1-.923 1.026c-.514.527-1.127 1.083-1.784 1.513c-.64.42-1.425.789-2.261.789A5 5 0 0 1 6 7c.836 0 1.622.37 2.261.789c.657.43 1.27.985 1.784 1.513c.348.357.66.713.923 1.026c.338-.441.769-.973 1.262-1.507c.605-.655 1.327-1.34 2.107-1.868C15.1 6.435 16.02 6 17 6a6 6 0 0 1 0 12c-.979 0-1.899-.435-2.663-.953c-.78-.528-1.502-1.213-2.107-1.868a21 21 0 0 1-1.262-1.507M6 15a3 3 0 1 1 0-6c.268 0 .66.13 1.165.461c.488.32.987.765 1.447 1.237c.456.468.85.94 1.133 1.297l.004.005l-.004.005c-.282.357-.677.83-1.133 1.297c-.46.473-.959.917-1.447 1.237C6.66 14.869 6.268 15 6 15" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignLemniscateLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M17 17c-2.761 0-6-5-6-5s3.239-5 6-5a5 5 0 0 1 0 10ZM6 8c2.21 0 5 4 5 4s-2.79 4-5 4a4 4 0 0 1 0-8Z"/>`), g.Group(children),
	)
}

func SignLemniscateThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="M17 17c-2.761 0-6-5-6-5s3.239-5 6-5a5 5 0 0 1 0 10ZM6 8c2.21 0 5 4 5 4s-2.79 4-5 4a4 4 0 0 1 0-8Z"/>`), g.Group(children),
	)
}

func SignMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 12h12"/>`), g.Group(children),
	)
}

func SignMinusBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M6 12h12"/>`), g.Group(children),
	)
}

func SignMinusCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="11.999" r="9"/><path d="M9 12h6"/></g>`), g.Group(children),
	)
}

func SignMinusCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="11.999" r="9"/><path d="M9 12h6"/></g>`), g.Group(children),
	)
}

func SignMinusCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="11.999" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="11.999" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6"/></g>`), g.Group(children),
	)
}

func SignMinusCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 11.999c0-5.523 4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10s-10-4.477-10-10M9 11a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignMinusCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="11.999" r="9"/><path d="M9 12h6"/></g>`), g.Group(children),
	)
}

func SignMinusCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="11.999" r="9"/><path d="M9 12h6"/></g>`), g.Group(children),
	)
}

func SignMinusFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 12a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignMinusLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 12h12"/>`), g.Group(children),
	)
}

func SignMinusSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func SignMinusSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M9 12h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func SignMinusSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 12h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/></g>`), g.Group(children),
	)
}

func SignMinusSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm6 7a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignMinusSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9 12h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func SignMinusSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9 12h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func SignMinusThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 12h12"/>`), g.Group(children),
	)
}

func SignPercent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><circle cx="16" cy="18" r="2"/><circle cx="8" cy="6" r="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 4L8 20"/></g>`), g.Group(children),
	)
}

func SignPercentBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><circle cx="16" cy="18" r="2"/><circle cx="8" cy="6" r="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 4L8 20"/></g>`), g.Group(children),
	)
}

func SignPercentFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.894 4.447a1 1 0 1 0-1.788-.894l-8 16a1 1 0 1 0 1.788.894zM8 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2M5 6a3 3 0 1 1 6 0a3 3 0 0 1-6 0m11 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignPercentLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="16" cy="18" r="2"/><circle cx="8" cy="6" r="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 4L8 20"/></g>`), g.Group(children),
	)
}

func SignPercentThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><circle cx="16" cy="18" r="2"/><circle cx="8" cy="6" r="2"/><path stroke-linecap="round" stroke-linejoin="round" d="M16 4L8 20"/></g>`), g.Group(children),
	)
}

func SignPi(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7 8h10M9 8v8m6-8v6.03A1.97 1.97 0 0 0 16.97 16H17"/>`), g.Group(children),
	)
}

func SignPiBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M7 8h10M9 8v8m6-8v6.03A1.97 1.97 0 0 0 16.97 16H17"/>`), g.Group(children),
	)
}

func SignPiFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 8a1 1 0 0 1 1-1h10a1 1 0 1 1 0 2h-1v5.03c0 .536.434.97.97.97H17a1 1 0 1 1 0 2h-.03A2.97 2.97 0 0 1 14 14.03V9h-4v7a1 1 0 1 1-2 0V9H7a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignPiLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 8h10M9 8v8m6-8v6.03A1.97 1.97 0 0 0 16.97 16H17"/>`), g.Group(children),
	)
}

func SignPiThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7 8h10M9 8v8m6-8v6.03A1.97 1.97 0 0 0 16.97 16H17"/>`), g.Group(children),
	)
}

func SignPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 6v12m-6-6h12"/>`), g.Group(children),
	)
}

func SignPlusBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 6v12m-6-6h12"/>`), g.Group(children),
	)
}

func SignPlusCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="11.999" r="9"/><path d="M12 9v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func SignPlusCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="11.999" r="9"/><path d="M12 9v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func SignPlusCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="11.999" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="11.999" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func SignPlusCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 11.999c0-5.523 4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10s-10-4.477-10-10M12 8a1 1 0 0 1 1 1v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2H9a1 1 0 1 1 0-2h2V9a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignPlusCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="11.999" r="9"/><path d="M12 9v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func SignPlusCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="11.999" r="9"/><path d="M12 9v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func SignPlusFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 6a1 1 0 1 0-2 0v5H6a1 1 0 1 0 0 2h5v5a1 1 0 1 0 2 0v-5h5a1 1 0 1 0 0-2h-5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignPlusLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 6v12m-6-6h12"/>`), g.Group(children),
	)
}

func SignPlusMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16 4L8 20M7 4v6M4 7h6m4 10h6"/>`), g.Group(children),
	)
}

func SignPlusMinusBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M16 4L8 20M7 4v6M4 7h6m4 10h6"/>`), g.Group(children),
	)
}

func SignPlusMinusFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.894 4.447a1 1 0 1 0-1.788-.894l-8 16a1 1 0 1 0 1.788.894zM8 4a1 1 0 1 0-2 0v2H4a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0V8h2a1 1 0 1 0 0-2H8zm6 12a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignPlusMinusLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16 4L8 20M7 4v6M4 7h6m4 10h6"/>`), g.Group(children),
	)
}

func SignPlusMinusThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16 4L8 20M7 4v6M4 7h6m4 10h6"/>`), g.Group(children),
	)
}

func SignPlusSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v6m-3-3h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func SignPlusSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M12 9v6m-3-3h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func SignPlusSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 9v6m-3-3h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/></g>`), g.Group(children),
	)
}

func SignPlusSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm10 5a1 1 0 1 0-2 0v2H9a1 1 0 1 0 0 2h2v2a1 1 0 1 0 2 0v-2h2a1 1 0 1 0 0-2h-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignPlusSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9v6m-3-3h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func SignPlusSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 9v6m-3-3h6M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func SignPlusThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12 6v12m-6-6h12"/>`), g.Group(children),
	)
}

func SignRadical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m3 15l1.5 5L8 5h13"/>`), g.Group(children),
	)
}

func SignRadicalBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m3 15l1.5 5L8 5h13"/>`), g.Group(children),
	)
}

func SignRadicalFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.026 4.773A1 1 0 0 1 8 4h13a1 1 0 1 1 0 2H8.794l-3.32 14.227a1 1 0 0 1-1.932.06l-1.5-5a1 1 0 1 1 1.916-.574l.421 1.404z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignRadicalLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3 15l1.5 5L8 5h13"/>`), g.Group(children),
	)
}

func SignRadicalThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m3 15l1.5 5L8 5h13"/>`), g.Group(children),
	)
}

func SignTimes(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 7L7 17M7 7l10 10"/>`), g.Group(children),
	)
}

func SignTimesBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M17 7L7 17M7 7l10 10"/>`), g.Group(children),
	)
}

func SignTimesCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="11.999" r="9"/><path d="m15 9l-6 6m0-6l6 6"/></g>`), g.Group(children),
	)
}

func SignTimesCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="11.999" r="9"/><path d="m15 9l-6 6m0-6l6 6"/></g>`), g.Group(children),
	)
}

func SignTimesCircleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="11.999" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="11.999" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m15 9l-6 6m0-6l6 6"/></g>`), g.Group(children),
	)
}

func SignTimesCircleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 11.999c0-5.523 4.477-10 10-10s10 4.477 10 10s-4.477 10-10 10s-10-4.477-10-10m13.707-3.706a1 1 0 0 1 0 1.414L13.414 12l2.293 2.293a1 1 0 0 1-1.414 1.414L12 13.414l-2.293 2.293a1 1 0 0 1-1.414-1.414L10.586 12L8.293 9.707a1 1 0 0 1 1.414-1.414L12 10.586l2.293-2.293a1 1 0 0 1 1.414 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignTimesCircleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="11.999" r="9"/><path d="m15 9l-6 6m0-6l6 6"/></g>`), g.Group(children),
	)
}

func SignTimesCircleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="11.999" r="9"/><path d="m15 9l-6 6m0-6l6 6"/></g>`), g.Group(children),
	)
}

func SignTimesDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M17 7L7 17M7 7l10 10"/>`), g.Group(children),
	)
}

func SignTimesFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.707 7.707a1 1 0 0 0-1.414-1.414L12 10.586L7.707 6.293a1 1 0 0 0-1.414 1.414L10.586 12l-4.293 4.293a1 1 0 1 0 1.414 1.414L12 13.414l4.293 4.293a1 1 0 1 0 1.414-1.414L13.414 12z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignTimesLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17 7L7 17M7 7l10 10"/>`), g.Group(children),
	)
}

func SignTimesSquare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm11 5l-6 6m0-6l6 6"/>`), g.Group(children),
	)
}

func SignTimesSquareBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm11 5l-6 6m0-6l6 6"/>`), g.Group(children),
	)
}

func SignTimesSquareDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M4 4.001h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm11 5l-6 6m0-6l6 6"/></g>`), g.Group(children),
	)
}

func SignTimesSquareFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 4a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v14a3 3 0 0 1-3 3H6a3 3 0 0 1-3-3zm12.707 4.293a1 1 0 0 1 0 1.414L13.414 12l2.293 2.293a1 1 0 0 1-1.414 1.414L12 13.414l-2.293 2.293a1 1 0 0 1-1.414-1.414L10.586 12L8.293 9.707a1 1 0 0 1 1.414-1.414L12 10.586l2.293-2.293a1 1 0 0 1 1.414 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignTimesSquareLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm11 5l-6 6m0-6l6 6"/>`), g.Group(children),
	)
}

func SignTimesSquareThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 4h16v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2zm11 5l-6 6m0-6l6 6"/>`), g.Group(children),
	)
}

func SignTimesThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17 7L7 17M7 7l10 10"/>`), g.Group(children),
	)
}

func SignX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M17 8h-.68a2.64 2.64 0 0 0-2.112 1.056l-4.416 5.888A2.64 2.64 0 0 1 7.68 16H7"/><path d="M8 8h.368c1 0 1.915.565 2.362 1.46l2.54 5.08A2.64 2.64 0 0 0 15.632 16H16"/></g>`), g.Group(children),
	)
}

func SignXBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M17 8h-.68a2.64 2.64 0 0 0-2.112 1.056l-4.416 5.888A2.64 2.64 0 0 1 7.68 16H7"/><path d="M8 8h.368c1 0 1.915.565 2.362 1.46l2.54 5.08A2.64 2.64 0 0 0 15.632 16H16"/></g>`), g.Group(children),
	)
}

func SignXFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 9a1 1 0 1 0 0-2h-.68a3.64 3.64 0 0 0-2.912 1.456l-1.237 1.65l-.547-1.094A3.64 3.64 0 0 0 8.368 7H8a1 1 0 0 0 0 2h.368c.622 0 1.19.351 1.467.907l.994 1.987l-1.837 2.45A1.64 1.64 0 0 1 7.68 15H7a1 1 0 1 0 0 2h.68a3.64 3.64 0 0 0 2.912-1.456l1.237-1.65l.547 1.094A3.64 3.64 0 0 0 15.632 17H16a1 1 0 1 0 0-2h-.368a1.64 1.64 0 0 1-1.467-.907l-.994-1.987l1.837-2.45A1.64 1.64 0 0 1 16.32 9z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignXLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M17 8h-.68a2.64 2.64 0 0 0-2.112 1.056l-4.416 5.888A2.64 2.64 0 0 1 7.68 16H7"/><path d="M8 8h.368c1 0 1.915.565 2.362 1.46l2.54 5.08A2.64 2.64 0 0 0 15.632 16H16"/></g>`), g.Group(children),
	)
}

func SignXThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M17 8h-.68a2.64 2.64 0 0 0-2.112 1.056l-4.416 5.888A2.64 2.64 0 0 1 7.68 16H7"/><path d="M8 8h.368c1 0 1.915.565 2.362 1.46l2.54 5.08A2.64 2.64 0 0 0 15.632 16H16"/></g>`), g.Group(children),
	)
}

func SignY(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M7 21h.525a2.64 2.64 0 0 0 2.248-1.256l6.29-10.22A1 1 0 0 0 15.21 8H14"/><path d="M12 16L9.487 9.298A2 2 0 0 0 7.614 8H7"/></g>`), g.Group(children),
	)
}

func SignYBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M7 21h.525a2.64 2.64 0 0 0 2.248-1.256l6.29-10.22A1 1 0 0 0 15.21 8H14"/><path d="M12 16L9.487 9.298A2 2 0 0 0 7.614 8H7"/></g>`), g.Group(children),
	)
}

func SignYFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.914 10.048C17.734 8.716 16.775 7 15.21 7H14a1 1 0 1 0 0 2h1.21l-2.962 4.814l-1.825-4.867A3 3 0 0 0 7.614 7H7a1 1 0 0 0 0 2h.614a1 1 0 0 1 .936.649l2.37 6.322l-1.998 3.249a1.64 1.64 0 0 1-1.397.78H7a1 1 0 1 0 0 2h.525a3.64 3.64 0 0 0 3.1-1.732z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SignYLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7 21h.525a2.64 2.64 0 0 0 2.248-1.256l6.29-10.22A1 1 0 0 0 15.21 8H14"/><path d="M12 16L9.487 9.298A2 2 0 0 0 7.614 8H7"/></g>`), g.Group(children),
	)
}

func SignYThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M7 21h.525a2.64 2.64 0 0 0 2.248-1.256l6.29-10.22A1 1 0 0 0 15.21 8H14"/><path d="M12 16L9.487 9.298A2 2 0 0 0 7.614 8H7"/></g>`), g.Group(children),
	)
}

func SlightlySmilingFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2" d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2"/></g>`), g.Group(children),
	)
}

func SlightlySmilingFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2.5" d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2"/></g>`), g.Group(children),
	)
}

func SlightlySmilingFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9 9.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2"/></g>`), g.Group(children),
	)
}

func SlightlySmilingFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m7-4a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5A1.5 1.5 0 0 0 9.01 8zm6 0a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5V9.5a1.5 1.5 0 0 0-1.5-1.5zm-6.966 5.134a1 1 0 0 1 1.367.365C9.92 14.4 10.89 15 12 15s2.08-.601 2.6-1.5a1 1 0 1 1 1.73 1A5 5 0 0 1 12 17a5 5 0 0 1-4.33-2.5a1 1 0 0 1 .364-1.366" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SlightlySmilingFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/><path stroke-linecap="round" stroke-width="1.5" d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2"/></g>`), g.Group(children),
	)
}

func SlightlySmilingFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M9.01 9.5v.01H9V9.5zm6 0v.01H15V9.5z"/><path stroke-linecap="round" d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2"/></g>`), g.Group(children),
	)
}

func SmilingFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2M14 10l1-1l1 1m-6 0L9 9l-1 1"/></g>`), g.Group(children),
	)
}

func SmilingFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2M14 10l1-1l1 1m-6 0L9 9l-1 1"/></g>`), g.Group(children),
	)
}

func SmilingFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2M14 10l1-1l1 1m-6 0L9 9l-1 1"/></g>`), g.Group(children),
	)
}

func SmilingFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m5.293-1.293a1 1 0 0 0 1.414 0L9 10.414l.293.293a1 1 0 0 0 1.414-1.414l-1-1a1 1 0 0 0-1.414 0l-1 1a1 1 0 0 0 0 1.414m8 0a1 1 0 0 0 1.414-1.414l-1-1a1 1 0 0 0-1.414 0l-1 1a1 1 0 0 0 1.414 1.414l.293-.293zM9.4 13.5a1 1 0 0 0-1.731 1.002A5 5 0 0 0 12 17a5 5 0 0 0 4.33-2.5a1 1 0 1 0-1.73-1c-.52.899-1.49 1.5-2.6 1.5a3 3 0 0 1-2.6-1.5" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SmilingFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2M14 10l1-1l1 1m-6 0L9 9l-1 1"/></g>`), g.Group(children),
	)
}

func SmilingFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2M14 10l1-1l1 1m-6 0L9 9l-1 1"/></g>`), g.Group(children),
	)
}

func SortingCenter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M7 12h10m-6 6h2"/>`), g.Group(children),
	)
}

func SortingCenterBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 6h16M7 12h10m-6 6h2"/>`), g.Group(children),
	)
}

func SortingCenterDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M7 12h10m-6 6h2"/>`), g.Group(children),
	)
}

func SortingCenterFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 6a1 1 0 0 1 1-1h15a1 1 0 1 1 0 2h-15a1 1 0 0 1-1-1m3 6a1 1 0 0 1 1-1h9a1 1 0 1 1 0 2h-9a1 1 0 0 1-1-1m4 6a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SortingCenterLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6h16M7 12h10m-6 6h2"/>`), g.Group(children),
	)
}

func SortingCenterThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M7 12h10m-6 6h2"/>`), g.Group(children),
	)
}

func SortingLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h10M4 18h4"/>`), g.Group(children),
	)
}

func SortingLeftBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 6h16M4 12h10M4 18h4"/>`), g.Group(children),
	)
}

func SortingLeftDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16M4 12h10M4 18h4"/>`), g.Group(children),
	)
}

func SortingLeftFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 6a1 1 0 0 1 1-1h15a1 1 0 1 1 0 2h-15a1 1 0 0 1-1-1m0 6a1 1 0 0 1 1-1h9a1 1 0 1 1 0 2h-9a1 1 0 0 1-1-1m0 6a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SortingLeftLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6h16M4 12h10M4 18h4"/>`), g.Group(children),
	)
}

func SortingLeftThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 6h16M4 12h10M4 18h4"/>`), g.Group(children),
	)
}

func SortingRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16m-10 6h10m-4 6h4"/>`), g.Group(children),
	)
}

func SortingRightBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M4 6h16m-10 6h10m-4 6h4"/>`), g.Group(children),
	)
}

func SortingRightDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4 6h16m-10 6h10m-4 6h4"/>`), g.Group(children),
	)
}

func SortingRightFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.5 6a1 1 0 0 1 1-1h15a1 1 0 1 1 0 2h-15a1 1 0 0 1-1-1m6 6a1 1 0 0 1 1-1h9a1 1 0 1 1 0 2h-9a1 1 0 0 1-1-1m6 6a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SortingRightLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 6h16m-10 6h10m-4 6h4"/>`), g.Group(children),
	)
}

func SortingRightThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4 6h16m-10 6h10m-4 6h4"/>`), g.Group(children),
	)
}

func SquintingFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><circle cx="12" cy="12" r="9"/><path fill="currentColor" d="M12 16c1.48 0 2.773-.804 3.465-2h-6.93A4 4 0 0 0 12 16"/><path d="m16 10.5l-2-1l2-1m-8 2l2-1l-2-1"/></g>`), g.Group(children),
	)
}

func SquintingFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><circle cx="12" cy="12" r="9"/><path fill="currentColor" d="M12 16c1.48 0 2.773-.804 3.465-2h-6.93A4 4 0 0 0 12 16"/><path d="m16 10.5l-2-1l2-1m-8 2l2-1l-2-1"/></g>`), g.Group(children),
	)
}

func SquintingFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 16c1.48 0 2.773-.804 3.465-2h-6.93A4 4 0 0 0 12 16m4-5.5l-2-1l2-1m-8 2l2-1l-2-1"/></g>`), g.Group(children),
	)
}

func SquintingFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m5.106-3.947a1 1 0 0 0 .447 1.341l.21.106l-.21.106a1 1 0 1 0 .894 1.788l2-1a1 1 0 0 0 0-1.788l-2-1a1 1 0 0 0-1.341.447m9.341 1.341a1 1 0 1 0-.894-1.788l-2 1a1 1 0 0 0 0 1.788l2 1a1 1 0 1 0 .894-1.788l-.211-.106zM8.535 13a1 1 0 0 0-.865 1.5A5 5 0 0 0 12 17a5 5 0 0 0 4.33-2.5a1 1 0 0 0-.865-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SquintingFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9"/><path d="M12 16c1.48 0 2.773-.804 3.465-2h-6.93A4 4 0 0 0 12 16m4-5.5l-2-1l2-1m-8 2l2-1l-2-1"/></g>`), g.Group(children),
	)
}

func SquintingFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="12" cy="12" r="9"/><path d="M12 16c1.48 0 2.773-.804 3.465-2h-6.93A4 4 0 0 0 12 16m4-5.5l-2-1l2-1m-8 2l2-1l-2-1"/></g>`), g.Group(children),
	)
}

func Star(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m12 2l3.104 6.728l7.358.873l-5.44 5.03l1.444 7.268L12 18.28L5.534 21.9l1.444-7.268L1.538 9.6l7.359-.873z"/>`), g.Group(children),
	)
}

func StarBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5" d="m12 2l3.104 6.728l7.358.873l-5.44 5.03l1.444 7.268L12 18.28L5.534 21.9l1.444-7.268L1.538 9.6l7.359-.873z"/>`), g.Group(children),
	)
}

func StarDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 2l3.104 6.728l7.358.873l-5.44 5.03l1.444 7.268L12 18.28L5.534 21.9l1.444-7.268L1.538 9.6l7.359-.873z" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="m12 2l3.104 6.728l7.358.873l-5.44 5.03l1.444 7.268L12 18.28L5.534 21.9l1.444-7.268L1.538 9.6l7.359-.873z"/></g>`), g.Group(children),
	)
}

func StarFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.908 1.581a1 1 0 0 0-1.816 0l-2.87 6.22l-6.801.807a1 1 0 0 0-.562 1.727l5.03 4.65l-1.335 6.72a1 1 0 0 0 1.469 1.067L12 19.426l5.977 3.346a1 1 0 0 0 1.47-1.068l-1.335-6.718l5.029-4.651a1 1 0 0 0-.562-1.727L15.777 7.8z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func StarLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="m12 2l3.104 6.728l7.358.873l-5.44 5.03l1.444 7.268L12 18.28L5.534 21.9l1.444-7.268L1.538 9.6l7.359-.873z"/>`), g.Group(children),
	)
}

func StarOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 2l.908-.419a1 1 0 0 0-1.816 0zm3.104 6.728l-.909.42a1 1 0 0 0 .79.573zm7.358.873l.679.734a1 1 0 0 0-.562-1.727zm-3.996 12.298l-.489.873a1 1 0 0 0 1.47-1.068zM12 18.28l.489-.873a1 1 0 0 0-.977 0zM5.534 21.9l-.98-.196a1 1 0 0 0 1.469 1.068zm1.444-7.268l.981.194a1 1 0 0 0-.302-.929zM1.538 9.6l-.117-.993a1 1 0 0 0-.562 1.727zm7.326.138a1 1 0 1 0-.235-1.986zm9.731 7.68a1 1 0 1 0-1.962.39zM9.721 4.553a1 1 0 1 0 1.816.838zm8.319 7.775a1 1 0 0 0 1.358 1.468zm-6.948-9.91l3.103 6.73l1.817-.838l-3.104-6.729zm3.894 7.303l7.358.873l.235-1.986l-7.358-.873zm3.968 11.306l-6.466-3.62l-.976 1.746l6.465 3.619zm-7.442-3.62l-6.466 3.62l.977 1.745l6.466-3.62zm-4.997 4.687l1.444-7.268l-1.961-.39l-1.444 7.268zm1.142-8.197l-5.44-5.03L.86 10.335l5.44 5.03zm-6-3.303l7.207-.855l-.235-1.986l-7.208.855zm14.976 7.215l.852 4.285l1.962-.39l-.852-4.285zM11.537 5.391l1.371-2.972l-1.816-.838l-1.37 2.972zm10.246 3.476l-3.743 3.46l1.358 1.47l3.743-3.462z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func StarOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 2l1.135-.524a1.25 1.25 0 0 0-2.27 0zm3.104 6.728l-1.136.524c.182.395.557.667.988.718zm7.358.873l.848.917a1.25 1.25 0 0 0-.701-2.158zm-3.996 12.298l-.61 1.09a1.25 1.25 0 0 0 1.836-1.333zM12 18.28l.61-1.09a1.25 1.25 0 0 0-1.22 0zM5.534 21.9l-1.226-.244a1.25 1.25 0 0 0 1.837 1.334zm1.444-7.268l1.226.243a1.25 1.25 0 0 0-.377-1.161zM1.538 9.6l-.147-1.24a1.25 1.25 0 0 0-.701 2.159zm7.355.386a1.25 1.25 0 0 0-.293-2.48zm9.947 7.384a1.25 1.25 0 1 0-2.452.487zM9.494 4.448a1.25 1.25 0 1 0 2.27 1.047zm8.376 7.695a1.25 1.25 0 1 0 1.698 1.836zm-7.005-9.62l3.103 6.728l2.27-1.047l-3.103-6.729zm4.091 7.446l7.358.872l.295-2.482l-7.358-.873zm4.12 10.838l-6.465-3.619l-1.222 2.182l6.466 3.619zM11.39 17.19l-6.465 3.62l1.22 2.18l6.467-3.618zm-4.63 4.953l1.444-7.268l-2.452-.487l-1.444 7.268zm1.067-8.43l-5.44-5.03L.69 10.518l5.44 5.031zm-6.141-2.87l7.208-.855l-.295-2.483l-7.208.855zm14.702 7.015l.852 4.285l2.452-.487l-.852-4.285zM11.764 5.496l1.371-2.972l-2.27-1.048L9.494 4.45zm9.849 3.187l-3.743 3.461l1.698 1.836l3.742-3.462z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2.5" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func StarOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 18.28l6.466 3.62l-.852-4.286l-8.868-8.868l-7.208.855l5.44 5.03L5.534 21.9z" opacity=".16"/><path fill="currentColor" d="m12 2l.908-.419a1 1 0 0 0-1.816 0zm3.104 6.728l-.909.42a1 1 0 0 0 .79.573zm7.358.873l.679.734a1 1 0 0 0-.562-1.727zm-3.996 12.298l-.489.873a1 1 0 0 0 1.47-1.068zM12 18.28l.489-.873a1 1 0 0 0-.977 0zM5.534 21.9l-.98-.196a1 1 0 0 0 1.469 1.068zm1.444-7.268l.981.194a1 1 0 0 0-.302-.929zM1.538 9.6l-.117-.993a1 1 0 0 0-.562 1.727zm7.326.138a1 1 0 1 0-.235-1.986zm9.731 7.68a1 1 0 1 0-1.962.39zM9.721 4.553a1 1 0 1 0 1.816.838zm8.319 7.775a1 1 0 0 0 1.358 1.468zm-6.948-9.91l3.103 6.73l1.817-.838l-3.104-6.729zm3.894 7.303l7.358.873l.235-1.986l-7.358-.873zm3.968 11.306l-6.466-3.62l-.976 1.746l6.465 3.619zm-7.442-3.62l-6.466 3.62l.977 1.745l6.466-3.62zm-4.997 4.687l1.444-7.268l-1.961-.39l-1.444 7.268zm1.142-8.197l-5.44-5.03L.86 10.335l5.44 5.03zm-6-3.303l7.207-.855l-.235-1.986l-7.208.855zm14.976 7.215l.852 4.285l1.962-.39l-.852-4.285zM11.537 5.391l1.371-2.972l-1.816-.838l-1.37 2.972zm10.246 3.476l-3.743 3.46l1.358 1.47l3.743-3.462z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func StarOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1a1 1 0 0 1 .908.581l2.87 6.22l6.801.807a1 1 0 0 1 .562 1.727l-3.743 3.461a1 1 0 1 1-1.358-1.468l2.151-1.99l-5.205-.617a1 1 0 0 1-.79-.574L12 4.387l-.463 1.004a1 1 0 1 1-1.816-.838l1.371-2.972A1 1 0 0 1 12 1M9.11 7.696l.343.343l8.868 8.868l.215.215l2.171 2.17a1 1 0 1 1-1.414 1.415l-.056-.056l.21 1.053a1 1 0 0 1-1.47 1.068L12 19.426l-5.977 3.346a1 1 0 0 1-1.47-1.068l1.336-6.718l-5.03-4.651a1 1 0 0 1 .562-1.727l5.16-.612l-3.288-3.289a1 1 0 1 1 1.414-1.414z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func StarOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 2l.681-.314a.75.75 0 0 0-1.362 0zm3.104 6.728l-.681.315c.109.236.333.4.592.43zm7.358.873l.509.55a.75.75 0 0 0-.421-1.295zm-3.996 12.298l-.367.655a.75.75 0 0 0 1.102-.801zM12 18.28l.366-.654a.75.75 0 0 0-.732 0zM5.534 21.9l-.735-.147a.75.75 0 0 0 1.102.8zm1.444-7.268l.736.146a.75.75 0 0 0-.226-.697zM1.538 9.6l-.088-.745a.75.75 0 0 0-.42 1.296zm7.297-.11A.75.75 0 1 0 8.658 8zm9.515 7.977a.75.75 0 1 0-1.471.292zM9.948 4.658a.75.75 0 1 0 1.362.628zm8.262 7.853a.75.75 0 1 0 1.018 1.102zM11.319 2.314l3.103 6.729l1.363-.629l-3.104-6.728zm3.696 7.16l7.358.872l.177-1.49l-7.358-.872zm3.817 11.77l-6.466-3.618l-.732 1.308l6.465 3.62zm-7.198-3.618l-6.466 3.619l.733 1.309l6.465-3.62zm-5.364 4.42l1.444-7.268l-1.471-.293l-1.444 7.268zm1.218-7.965l-5.44-5.03l-1.019 1.1l5.44 5.031zm-5.861-3.735l7.208-.855l-.177-1.49l-7.208.855zm15.252 7.414l.851 4.285l1.471-.292l-.851-4.285zM11.31 5.286l1.371-2.972l-1.362-.628l-1.37 2.972zM21.952 9.05l-3.742 3.461l1.018 1.102l3.743-3.461z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func StarOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 2l.454-.21a.5.5 0 0 0-.908 0zm3.104 6.728l-.454.21a.5.5 0 0 0 .395.287zm7.358.873l.34.367a.5.5 0 0 0-.282-.864zm-3.996 12.298l-.245.436a.5.5 0 0 0 .735-.533zM12 18.28l.244-.436a.5.5 0 0 0-.488 0zM5.534 21.9l-.49-.098a.5.5 0 0 0 .735.534zm1.444-7.268l.49.097a.5.5 0 0 0-.15-.464zM1.538 9.6l-.058-.497a.5.5 0 0 0-.281.864zm7.267-.358a.5.5 0 0 0-.118-.993zm9.3 8.274a.5.5 0 0 0-.981.195zm-7.93-12.754a.5.5 0 1 0 .908.419zm8.205 7.932a.5.5 0 0 0 .678.734zM11.546 2.209l3.104 6.729l.907-.419l-3.103-6.728zm3.499 7.016l7.358.872l.117-.993l-7.358-.872zm3.665 12.238l-6.466-3.62l-.488.873l6.465 3.62zm-6.954-3.62l-6.466 3.62l.489.872l6.465-3.619zm-5.731 4.154l1.444-7.268l-.981-.195l-1.444 7.268zm1.293-7.732l-5.44-5.031l-.68.734l5.441 5.03zm-5.72-4.168l7.207-.854l-.118-.993l-7.207.854zm15.526 7.615l.851 4.285l.981-.195l-.851-4.285zm-6.04-12.53l1.37-2.973l-.908-.418l-1.37 2.972zm11.038 4.052l-3.742 3.46l.678.735l3.743-3.461z"/><path stroke="currentColor" stroke-linecap="round" d="m4 4l16 16"/></g>`), g.Group(children),
	)
}

func StarThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" d="m12 2l3.104 6.728l7.358.873l-5.44 5.03l1.444 7.268L12 18.28L5.534 21.9l1.444-7.268L1.538 9.6l7.359-.873z"/>`), g.Group(children),
	)
}

func Store(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" stroke-linejoin="round" d="M3 21h2m16 0h-2M5 21h14M5 21V9.328M19 21V9.328m-14 0a2 2 0 0 1-.11-.068l-.54-.36a2 2 0 0 1-.747-2.407l.894-2.236A2 2 0 0 1 6.354 3h11.292a2 2 0 0 1 1.857 1.257l.894 2.236A2 2 0 0 1 19.65 8.9l-.54.36a2 2 0 0 1-.11.068m-14 0a2 2 0 0 0 2.11-.068L9 8l1.89 1.26a2 2 0 0 0 2.22 0L15 8l1.89 1.26a2 2 0 0 0 2.11.068"/><path d="M14 21v-5a2 2 0 1 0-4 0v5"/></g>`), g.Group(children),
	)
}

func StoreBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><path stroke-linecap="round" stroke-linejoin="round" d="M3 21h2m16 0h-2M5 21h14M5 21V9.328M19 21V9.328m-14 0a2 2 0 0 1-.11-.068l-.54-.36a2 2 0 0 1-.747-2.407l.894-2.236A2 2 0 0 1 6.354 3h11.292a2 2 0 0 1 1.857 1.257l.894 2.236A2 2 0 0 1 19.65 8.9l-.54.36a2 2 0 0 1-.11.068m-14 0a2 2 0 0 0 2.11-.068L9 8l1.89 1.26a2 2 0 0 0 2.22 0L15 8l1.89 1.26a2 2 0 0 0 2.11.068"/><path d="M14 21v-5a2 2 0 1 0-4 0v5"/></g>`), g.Group(children),
	)
}

func StoreDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m3.603 6.493l.894-2.236A2 2 0 0 1 6.354 3h11.292a2 2 0 0 1 1.857 1.257l.894 2.236A2 2 0 0 1 19.65 8.9l-.54.36a2 2 0 0 1-2.22 0L15 8l-1.89 1.26a2 2 0 0 1-2.22 0L9 8L7.11 9.26a2 2 0 0 1-2.22 0l-.54-.36a2 2 0 0 1-.747-2.407" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 21h2m16 0h-2M5 21h14M5 21V9.328M19 21V9.328m-14 0a2 2 0 0 1-.11-.068l-.54-.36a2 2 0 0 1-.747-2.407l.894-2.236A2 2 0 0 1 6.354 3h11.292a2 2 0 0 1 1.857 1.257l.894 2.236A2 2 0 0 1 19.65 8.9l-.54.36a2 2 0 0 1-.11.068m-14 0a2 2 0 0 0 2.11-.068L9 8l1.89 1.26a2 2 0 0 0 2.22 0L15 8l1.89 1.26a2 2 0 0 0 2.11.068"/><path stroke="currentColor" stroke-width="2" d="M14 21v-5a2 2 0 1 0-4 0v5"/></g>`), g.Group(children),
	)
}

func StoreFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.569 3.886A3 3 0 0 1 6.354 2h11.292a3 3 0 0 1 2.785 1.886l.895 2.236a3 3 0 0 1-1.122 3.61L20 9.87V20h1a1 1 0 1 1 0 2H3a1 1 0 1 1 0-2h1V9.869l-.204-.137a3 3 0 0 1-1.122-3.61zM6 10.596V20h3v-4a3 3 0 1 1 6 0v4h3v-9.404c-.58 0-1.16-.168-1.664-.504L15 9.202l-1.336.89a3 3 0 0 1-3.328 0L9 9.202l-1.336.89A3 3 0 0 1 6 10.596M13 20v-4a1 1 0 1 0-2 0v4z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func StoreLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M3 21h2m16 0h-2M5 21h14M5 21V9.328M19 21V9.328m-14 0a2 2 0 0 1-.11-.068l-.54-.36a2 2 0 0 1-.747-2.407l.894-2.236A2 2 0 0 1 6.354 3h11.292a2 2 0 0 1 1.857 1.257l.894 2.236A2 2 0 0 1 19.65 8.9l-.54.36a2 2 0 0 1-.11.068m-14 0a2 2 0 0 0 2.11-.068L9 8l1.89 1.26a2 2 0 0 0 2.22 0L15 8l1.89 1.26a2 2 0 0 0 2.11.068"/><path d="M14 21v-5a2 2 0 1 0-4 0v5"/></g>`), g.Group(children),
	)
}

func StoreThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" stroke-linejoin="round" d="M3 21h2m16 0h-2M5 21h14M5 21V9.328M19 21V9.328m-14 0a2 2 0 0 1-.11-.068l-.54-.36a2 2 0 0 1-.747-2.407l.894-2.236A2 2 0 0 1 6.354 3h11.292a2 2 0 0 1 1.857 1.257l.894 2.236A2 2 0 0 1 19.65 8.9l-.54.36a2 2 0 0 1-.11.068m-14 0a2 2 0 0 0 2.11-.068L9 8l1.89 1.26a2 2 0 0 0 2.22 0L15 8l1.89 1.26a2 2 0 0 0 2.11.068"/><path d="M14 21v-5a2 2 0 1 0-4 0v5"/></g>`), g.Group(children),
	)
}

func Swap(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M11 8L7 4m0 0L3 8m4-4v16m6-4l4 4m0 0l4-4m-4 4V4"/>`), g.Group(children),
	)
}

func SwapBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M11 8L7 4m0 0L3 8m4-4v16m6-4l4 4m0 0l4-4m-4 4V4"/>`), g.Group(children),
	)
}

func SwapFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17.707 20.707a1 1 0 0 1-1.414 0l-4-4A1 1 0 0 1 13 15h3V4a1 1 0 1 1 2 0v11h3a1 1 0 0 1 .707 1.707zm-10-17.414a1 1 0 0 0-1.414 0l-4 4A1 1 0 0 0 3 9h3v11a1 1 0 1 0 2 0V9h3a1 1 0 0 0 .707-1.707z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SwapLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11 8L7 4m0 0L3 8m4-4v16m6-4l4 4m0 0l4-4m-4 4V4"/>`), g.Group(children),
	)
}

func SwapThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11 8L7 4m0 0L3 8m4-4v16m6-4l4 4m0 0l4-4m-4 4V4"/>`), g.Group(children),
	)
}

func Synchronize(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="M3 12a9 9 0 0 1 16-5.658"/><path d="M19.5 3v4h-4m5.5 5a9 9 0 0 1-16 5.657"/><path d="M4.5 21v-4h4"/></g>`), g.Group(children),
	)
}

func SynchronizeBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="M3 12a9 9 0 0 1 16-5.658"/><path d="M19.5 3v4h-4m5.5 5a9 9 0 0 1-16 5.657"/><path d="M4.5 21v-4h4"/></g>`), g.Group(children),
	)
}

func SynchronizeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M3 12a9 9 0 0 1 16-5.658"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M19.5 3v4h-4m5.5 5a9 9 0 0 1-16 5.657"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M4.5 21v-4h4"/></g>`), g.Group(children),
	)
}

func SynchronizeFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20.5 3a1 1 0 0 0-1.707-.707l-1.33 1.33A9.96 9.96 0 0 0 12 2C6.477 2 2 6.477 2 12a1 1 0 0 0 2 0a8 8 0 0 1 12.01-6.924l-1.217 1.217A1 1 0 0 0 15.5 8h4a1 1 0 0 0 1-1zM7.99 18.924l1.217-1.217A1 1 0 0 0 8.5 16h-4a1 1 0 0 0-1 1v4a1 1 0 0 0 1.707.707l1.33-1.33A9.96 9.96 0 0 0 12 22c5.523 0 10-4.477 10-10a1 1 0 1 0-2 0a8 8 0 0 1-12.01 6.924" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SynchronizeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3 12a9 9 0 0 1 16-5.658"/><path d="M19.5 3v4h-4m5.5 5a9 9 0 0 1-16 5.657"/><path d="M4.5 21v-4h4"/></g>`), g.Group(children),
	)
}

func SynchronizeThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M3 12a9 9 0 0 1 16-5.658"/><path d="M19.5 3v4h-4m5.5 5a9 9 0 0 1-16 5.657"/><path d="M4.5 21v-4h4"/></g>`), g.Group(children),
	)
}

func ThreeD(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m12 3l7.794 4.5v7.845a2 2 0 0 1-1 1.732L13 20.423a2 2 0 0 1-2 0l-5.794-3.346a2 2 0 0 1-1-1.732V7.5z"/><path d="M12 7v5l-4.33 2.5M12 12l4.33 2.5"/></g>`), g.Group(children),
	)
}

func ThreeDBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="m12 3l7.794 4.5v7.845a2 2 0 0 1-1 1.732L13 20.423a2 2 0 0 1-2 0l-5.794-3.346a2 2 0 0 1-1-1.732V7.5z"/><path d="M12 7v5l-4.33 2.5M12 12l4.33 2.5"/></g>`), g.Group(children),
	)
}

func ThreeDDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="m12 3l7.794 4.5v7.845a2 2 0 0 1-1 1.732L13 20.423a2 2 0 0 1-2 0l-5.794-3.346a2 2 0 0 1-1-1.732V7.5z" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12 3l7.794 4.5v7.845a2 2 0 0 1-1 1.732L13 20.423a2 2 0 0 1-2 0l-5.794-3.346a2 2 0 0 1-1-1.732V7.5z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 7v5l-4.33 2.5M12 12l4.33 2.5"/></g>`), g.Group(children),
	)
}

func ThreeDFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.5 2.134a1 1 0 0 1 1 0l7.794 4.5a1 1 0 0 1 .5.866v7.845a3 3 0 0 1-1.5 2.598L13.5 21.29a3 3 0 0 1-3 0l-5.794-3.346a3 3 0 0 1-1.5-2.598V7.5a1 1 0 0 1 .5-.866zM12 6a1 1 0 0 1 1 1v4.423l3.83 2.211a1 1 0 1 1-1 1.732L12 13.155l-3.83 2.211a1 1 0 1 1-1-1.732L11 11.423V7a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ThreeDLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m12 3l7.794 4.5v7.845a2 2 0 0 1-1 1.732L13 20.423a2 2 0 0 1-2 0l-5.794-3.346a2 2 0 0 1-1-1.732V7.5z"/><path d="M12 7v5l-4.33 2.5M12 12l4.33 2.5"/></g>`), g.Group(children),
	)
}

func ThreeDThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m12 3l7.794 4.5v7.845a2 2 0 0 1-1 1.732L13 20.423a2 2 0 0 1-2 0l-5.794-3.346a2 2 0 0 1-1-1.732V7.5z"/><path d="M12 7v5l-4.33 2.5M12 12l4.33 2.5"/></g>`), g.Group(children),
	)
}

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 10H2a1 1 0 0 0 1 1zm0 4v-1a1 1 0 0 0-1 1zm18-4v1a1 1 0 0 0 1-1zm0 4h1a1 1 0 0 0-1-1zM5 6h5V4H5zm5 0h9V4h-9zm9 12h-9v2h9zm-9 0H5v2h5zM9 5v14h2V5zm-5.293 6.293a1 1 0 0 1 0 1.414l1.414 1.414a3 3 0 0 0 0-4.242zm16.586 1.414a1 1 0 0 1 0-1.414l-1.414-1.414a3 3 0 0 0 0 4.242zM3 11c.257 0 .512.097.707.293l1.414-1.414A3 3 0 0 0 3 9zm1-1V7H2v3zm0 7v-3H2v3zm-.293-4.293A1 1 0 0 1 3 13v2c.766 0 1.536-.293 2.121-.879zm16.586-1.414A1 1 0 0 1 21 11V9c-.766 0-1.536.293-2.121.879zM20 7v3h2V7zm0 7v3h2v-3zm1-1a1 1 0 0 1-.707-.293l-1.414 1.414A3 3 0 0 0 21 15zM5 18a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3zm14 2a3 3 0 0 0 3-3h-2a1 1 0 0 1-1 1zm0-14a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3zM5 4a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1z"/>`), g.Group(children),
	)
}

func TicketBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 10H1.75c0 .69.56 1.25 1.25 1.25zm0 4v-1.25c-.69 0-1.25.56-1.25 1.25zm18-4v1.25c.69 0 1.25-.56 1.25-1.25zm0 4h1.25c0-.69-.56-1.25-1.25-1.25zM5 6.25h5v-2.5H5zm5 0h9v-2.5h-9zm9 11.5h-9v2.5h9zm-9 0H5v2.5h5zM8.75 5v14h2.5V5zm-5.22 6.47a.75.75 0 0 1 0 1.06l1.768 1.768a3.25 3.25 0 0 0 0-4.596zm16.94 1.06a.75.75 0 0 1 0-1.06l-1.768-1.768a3.25 3.25 0 0 0 0 4.596zM3 11.25a.74.74 0 0 1 .53.22l1.768-1.768A3.24 3.24 0 0 0 3 8.75zM4.25 10V7h-2.5v3zm0 7v-3h-2.5v3zm-.72-4.47a.74.74 0 0 1-.53.22v2.5c.83 0 1.664-.318 2.298-.952zm16.94-1.06a.74.74 0 0 1 .53-.22v-2.5c-.83 0-1.664.318-2.298.952zM19.75 7v3h2.5V7zm0 7v3h2.5v-3zM21 12.75a.74.74 0 0 1-.53-.22l-1.768 1.768A3.24 3.24 0 0 0 21 15.25zm-16 5a.75.75 0 0 1-.75-.75h-2.5A3.25 3.25 0 0 0 5 20.25zm14 2.5A3.25 3.25 0 0 0 22.25 17h-2.5a.75.75 0 0 1-.75.75zm0-14a.75.75 0 0 1 .75.75h2.5A3.25 3.25 0 0 0 19 3.75zM5 3.75A3.25 3.25 0 0 0 1.75 7h2.5A.75.75 0 0 1 5 6.25z"/>`), g.Group(children),
	)
}

func TicketDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M19 5h-9v14h9a2 2 0 0 0 2-2v-3a2 2 0 1 1 0-4V7a2 2 0 0 0-2-2" opacity=".16"/><path d="M3 10H2a1 1 0 0 0 1 1zm0 4v-1a1 1 0 0 0-1 1zm18-4v1a1 1 0 0 0 1-1zm0 4h1a1 1 0 0 0-1-1zM5 6h5V4H5zm5 0h9V4h-9zm9 12h-9v2h9zm-9 0H5v2h5zM9 5v14h2V5zm-5.293 6.293a1 1 0 0 1 0 1.414l1.414 1.414a3 3 0 0 0 0-4.242zm16.586 1.414a1 1 0 0 1 0-1.414l-1.414-1.414a3 3 0 0 0 0 4.242zM3 11c.257 0 .512.097.707.293l1.414-1.414A3 3 0 0 0 3 9zm1-1V7H2v3zm0 7v-3H2v3zm-.293-4.293A1 1 0 0 1 3 13v2c.766 0 1.536-.293 2.121-.879zm16.586-1.414A1 1 0 0 1 21 11V9c-.766 0-1.536.293-2.121.879zM20 7v3h2V7zm0 7v3h2v-3zm1-1a1 1 0 0 1-.707-.293l-1.414 1.414A3 3 0 0 0 21 15zM5 18a1 1 0 0 1-1-1H2a3 3 0 0 0 3 3zm14 2a3 3 0 0 0 3-3h-2a1 1 0 0 1-1 1zm0-14a1 1 0 0 1 1 1h2a3 3 0 0 0-3-3zM5 4a3 3 0 0 0-3 3h2a1 1 0 0 1 1-1z"/></g>`), g.Group(children),
	)
}

func TicketFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 4a3 3 0 0 0-3 3v3a1 1 0 0 0 1 1a1 1 0 0 1 0 2a1 1 0 0 0-1 1v3a3 3 0 0 0 3 3h14a3 3 0 0 0 3-3v-3a1 1 0 0 0-1-1a1 1 0 0 1 0-2a1 1 0 0 0 1-1V7a3 3 0 0 0-3-3zM4 7a1 1 0 0 1 1-1h4v12H5a1 1 0 0 1-1-1v-2.171a3 3 0 0 0 1.121-.708l-.692-.692l.692.692A3 3 0 0 0 4 9.171z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TicketLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 10h-.75c0 .414.336.75.75.75zm0 4v-.75a.75.75 0 0 0-.75.75zm18-4v.75a.75.75 0 0 0 .75-.75zm0 4h.75a.75.75 0 0 0-.75-.75zM5 5.75h5v-1.5H5zm5 0h9v-1.5h-9zm9 12.5h-9v1.5h9zm-9 0H5v1.5h5zM9.25 5v14h1.5V5zm-5.366 6.116a1.25 1.25 0 0 1 0 1.768l1.06 1.06a2.75 2.75 0 0 0 0-3.889zm16.232 1.768a1.25 1.25 0 0 1 0-1.768l-1.06-1.06a2.75 2.75 0 0 0 0 3.889zM3 10.75c.321 0 .64.122.884.366l1.06-1.06A2.74 2.74 0 0 0 3 9.25zm.75-.75V7h-1.5v3zm0 7v-3h-1.5v3zm.134-4.116A1.24 1.24 0 0 1 3 13.25v1.5c.703 0 1.408-.269 1.945-.806zm16.232-1.768c.244-.244.563-.366.884-.366v-1.5c-.703 0-1.408.269-1.945.806zM20.25 7v3h1.5V7zm0 7v3h1.5v-3zm.75-.75c-.321 0-.64-.122-.884-.366l-1.06 1.06A2.74 2.74 0 0 0 21 14.75zm-16 5c-.69 0-1.25-.56-1.25-1.25h-1.5A2.75 2.75 0 0 0 5 19.75zm14 1.5A2.75 2.75 0 0 0 21.75 17h-1.5c0 .69-.56 1.25-1.25 1.25zm0-14c.69 0 1.25.56 1.25 1.25h1.5A2.75 2.75 0 0 0 19 4.25zM5 4.25A2.75 2.75 0 0 0 2.25 7h1.5c0-.69.56-1.25 1.25-1.25z"/>`), g.Group(children),
	)
}

func TicketThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 10h-.5a.5.5 0 0 0 .5.5zm0 4v-.5a.5.5 0 0 0-.5.5zm18-4v.5a.5.5 0 0 0 .5-.5zm0 4h.5a.5.5 0 0 0-.5-.5zM5 5.5h5v-1H5zm5 0h9v-1h-9zm9 13h-9v1h9zm-9 0H5v1h5zM9.5 5v14h1V5zm-5.44 5.94a1.5 1.5 0 0 1 0 2.12l.708.708a2.5 2.5 0 0 0 0-3.536zm15.88 2.12a1.5 1.5 0 0 1 0-2.12l-.708-.708a2.5 2.5 0 0 0 0 3.536zM3 10.5c.385 0 .768.146 1.06.44l.708-.708A2.5 2.5 0 0 0 3 9.5zm.5-.5V7h-1v3zm0 7v-3h-1v3zm.56-3.94c-.292.294-.675.44-1.06.44v1c.639 0 1.28-.244 1.768-.732zm15.88-2.12c.292-.294.675-.44 1.06-.44v-1c-.639 0-1.28.244-1.768.732zM20.5 7v3h1V7zm0 7v3h1v-3zm.5-.5c-.385 0-.768-.146-1.06-.44l-.708.708A2.5 2.5 0 0 0 21 14.5zm-16 5A1.5 1.5 0 0 1 3.5 17h-1A2.5 2.5 0 0 0 5 19.5zm14 1a2.5 2.5 0 0 0 2.5-2.5h-1a1.5 1.5 0 0 1-1.5 1.5zm0-14A1.5 1.5 0 0 1 20.5 7h1A2.5 2.5 0 0 0 19 4.5zm-14-1A2.5 2.5 0 0 0 2.5 7h1A1.5 1.5 0 0 1 5 5.5z"/>`), g.Group(children),
	)
}

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 11v6m-4-6v6M6 7v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7M4 7h16M7 7l2-4h6l2 4"/>`), g.Group(children),
	)
}

func TrashBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M14 11v6m-4-6v6M6 7v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7M4 7h16M7 7l2-4h6l2 4"/>`), g.Group(children),
	)
}

func TrashDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M8 21h8a2 2 0 0 0 2-2V7H6v12a2 2 0 0 0 2 2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M14 11v6m-4-6v6M6 7v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7M4 7h16M7 7l2-4h6l2 4"/></g>`), g.Group(children),
	)
}

func TrashFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8.106 2.553A1 1 0 0 1 9 2h6a1 1 0 0 1 .894.553L17.618 6H20a1 1 0 1 1 0 2h-1v11a3 3 0 0 1-3 3H8a3 3 0 0 1-3-3V8H4a1 1 0 0 1 0-2h2.382zM14.382 4l1 2H8.618l1-2zM11 11a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0zm4 0a1 1 0 1 0-2 0v6a1 1 0 1 0 2 0z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TrashLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14 11v6m-4-6v6M6 7v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7M4 7h16M7 7l2-4h6l2 4"/>`), g.Group(children),
	)
}

func TrashSimple(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 7v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7M4 7h16M7 7l2-4h6l2 4"/>`), g.Group(children),
	)
}

func TrashSimpleBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M6 7v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7M4 7h16M7 7l2-4h6l2 4"/>`), g.Group(children),
	)
}

func TrashSimpleDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M8 21h8a2 2 0 0 0 2-2V7H6v12a2 2 0 0 0 2 2" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M6 7v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7M4 7h16M7 7l2-4h6l2 4"/></g>`), g.Group(children),
	)
}

func TrashSimpleFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 2a1 1 0 0 0-.894.553L6.382 6H4a1 1 0 0 0 0 2h1v11a3 3 0 0 0 3 3h8a3 3 0 0 0 3-3V8h1a1 1 0 1 0 0-2h-2.382l-1.724-3.447A1 1 0 0 0 15 2zm6.382 4l-1-2H9.618l-1 2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TrashSimpleLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6 7v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7M4 7h16M7 7l2-4h6l2 4"/>`), g.Group(children),
	)
}

func TrashSimpleThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6 7v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7M4 7h16M7 7l2-4h6l2 4"/>`), g.Group(children),
	)
}

func TrashThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14 11v6m-4-6v6M6 7v12a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V7M4 7h16M7 7l2-4h6l2 4"/>`), g.Group(children),
	)
}

func TrendDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 7l6 6l4-4l8 8"/><path d="M17 17h4v-4"/></g>`), g.Group(children),
	)
}

func TrendDownBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="m3 7l6 6l4-4l8 8"/><path d="M17 17h4v-4"/></g>`), g.Group(children),
	)
}

func TrendDownFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 17.01V13a1 1 0 0 0-1.707-.707L19 13.586l-5.293-5.293a1 1 0 0 0-1.414 0L9 11.586L3.707 6.293a1 1 0 0 0-1.414 1.414l6 6a1 1 0 0 0 1.414 0L13 10.414L17.586 15l-1.293 1.293A1 1 0 0 0 17 18h4l.048-.001a.997.997 0 0 0 .952-.99" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TrendDownLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3 7l6 6l4-4l8 8"/><path d="M17 17h4v-4"/></g>`), g.Group(children),
	)
}

func TrendDownThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3 7l6 6l4-4l8 8"/><path d="M17 17h4v-4"/></g>`), g.Group(children),
	)
}

func TrendUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m3 17l6-6l4 4l8-8"/><path d="M17 7h4v4"/></g>`), g.Group(children),
	)
}

func TrendUpBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5"><path d="m3 17l6-6l4 4l8-8"/><path d="M17 7h4v4"/></g>`), g.Group(children),
	)
}

func TrendUpFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M22 6.99V11a1 1 0 0 1-1.707.707L19 10.414l-5.293 5.293a1 1 0 0 1-1.414 0L9 12.414l-5.293 5.293a1 1 0 0 1-1.414-1.414l6-6a1 1 0 0 1 1.414 0L13 13.586L17.586 9l-1.293-1.293A1 1 0 0 1 17 6h4l.048.001a.996.996 0 0 1 .952.99" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TrendUpLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m3 17l6-6l4 4l8-8"/><path d="M17 7h4v4"/></g>`), g.Group(children),
	)
}

func TrendUpThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="m3 17l6-6l4 4l8-8"/><path d="M17 7h4v4"/></g>`), g.Group(children),
	)
}

func Type(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7V6a2 2 0 0 1 2-2h5m7 3V6a2 2 0 0 0-2-2h-5m0 0v16m0 0H9m3 0h3"/>`), g.Group(children),
	)
}

func TypeBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="M5 7V6a2 2 0 0 1 2-2h5m7 3V6a2 2 0 0 0-2-2h-5m0 0v16m0 0H9m3 0h3"/>`), g.Group(children),
	)
}

func TypeDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M5 7V6a2 2 0 0 1 2-2h5m7 3V6a2 2 0 0 0-2-2h-5m0 0v16m0 0H9m3 0h3"/>`), g.Group(children),
	)
}

func TypeFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 5a1 1 0 0 0-1 1v1a1 1 0 0 1-2 0V6a3 3 0 0 1 3-3h10a3 3 0 0 1 3 3v1a1 1 0 1 1-2 0V6a1 1 0 0 0-1-1h-4v14h2a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2h2V5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TypeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5 7V6a2 2 0 0 1 2-2h5m7 3V6a2 2 0 0 0-2-2h-5m0 0v16m0 0H9m3 0h3"/>`), g.Group(children),
	)
}

func TypeThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5 7V6a2 2 0 0 1 2-2h5m7 3V6a2 2 0 0 0-2-2h-5m0 0v16m0 0H9m3 0h3"/>`), g.Group(children),
	)
}

func Unavailable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 20a8 8 0 0 1-8-8H2c0 5.523 4.477 10 10 10zm0-16a8 8 0 0 1 8 8h2c0-5.523-4.477-10-10-10zm-8 8a7.97 7.97 0 0 1 2.343-5.657L4.93 4.93A9.97 9.97 0 0 0 2 11.999zm2.343-5.657A7.97 7.97 0 0 1 12 4V2a9.97 9.97 0 0 0-7.071 2.929zm-1.414 0l12.728 12.728l1.414-1.414L6.343 4.929zM20 12a7.97 7.97 0 0 1-2.343 5.657l1.414 1.414A9.97 9.97 0 0 0 22 12zm-2.343 5.657A7.97 7.97 0 0 1 12 20v2a9.97 9.97 0 0 0 7.071-2.929z"/>`), g.Group(children),
	)
}

func UnavailableBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 19.75A7.75 7.75 0 0 1 4.25 12h-2.5c0 5.66 4.59 10.25 10.25 10.25zm0-15.5A7.75 7.75 0 0 1 19.75 12h2.5c0-5.661-4.59-10.25-10.25-10.25zM4.25 12c0-2.14.866-4.076 2.27-5.48L4.752 4.752A10.22 10.22 0 0 0 1.75 12zm2.27-5.48A7.72 7.72 0 0 1 12 4.25v-2.5c-2.83 0-5.394 1.149-7.248 3.002zm-1.768 0L17.48 19.248l1.768-1.768L6.52 4.752zM19.75 12c0 2.14-.866 4.076-2.27 5.48l1.768 1.768A10.22 10.22 0 0 0 22.25 12zm-2.27 5.48A7.72 7.72 0 0 1 12 19.75v2.5c2.83 0 5.394-1.149 7.248-3.002z"/>`), g.Group(children),
	)
}

func UnavailableDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="12" r="9" opacity=".16"/><path d="M12 20a8 8 0 0 1-8-8H2c0 5.523 4.477 10 10 10zm0-16a8 8 0 0 1 8 8h2c0-5.523-4.477-10-10-10zm-8 8a7.97 7.97 0 0 1 2.343-5.657L4.93 4.93A9.97 9.97 0 0 0 2 12zm2.343-5.657A7.97 7.97 0 0 1 12 4V2a9.97 9.97 0 0 0-7.071 2.929zm-1.414 0l12.728 12.728l1.414-1.414L6.343 4.929zM20 12a7.97 7.97 0 0 1-2.343 5.657l1.414 1.414A9.97 9.97 0 0 0 22 12zm-2.343 5.657A7.97 7.97 0 0 1 12 20v2a9.97 9.97 0 0 0 7.071-2.929z"/></g>`), g.Group(children),
	)
}

func UnavailableFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.503 5.382a8 8 0 0 1 11.114 11.114zm-2.12 2.121a8 8 0 0 0 11.114 11.114zM12 2a9.97 9.97 0 0 0-7.071 2.929A9.97 9.97 0 0 0 2 12c0 5.523 4.477 10 10 10a9.97 9.97 0 0 0 7.071-2.929A9.97 9.97 0 0 0 22 12c0-5.523-4.477-10-10-10" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UnavailableLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 20.25A8.25 8.25 0 0 1 3.75 12h-1.5A9.75 9.75 0 0 0 12 21.75zm0-16.5A8.25 8.25 0 0 1 20.25 12h1.5A9.75 9.75 0 0 0 12 2.25zM3.75 12c0-2.278.923-4.34 2.416-5.834l-1.06-1.06A9.72 9.72 0 0 0 2.25 12zm2.416-5.834A8.22 8.22 0 0 1 12 3.75v-1.5a9.72 9.72 0 0 0-6.894 2.856zm-1.06 0l12.728 12.728l1.06-1.06L6.166 5.106zM20.25 12c0 2.278-.923 4.34-2.416 5.834l1.06 1.06A9.72 9.72 0 0 0 21.75 12zm-2.416 5.834A8.22 8.22 0 0 1 12 20.25v1.5a9.72 9.72 0 0 0 6.894-2.856z"/>`), g.Group(children),
	)
}

func UnavailableThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 20.5A8.5 8.5 0 0 1 3.5 12h-1a9.5 9.5 0 0 0 9.5 9.5zm0-17a8.5 8.5 0 0 1 8.5 8.5h1A9.5 9.5 0 0 0 12 2.5zM3.5 12c0-2.347.95-4.472 2.49-6.01l-.708-.708A9.47 9.47 0 0 0 2.5 12zm2.49-6.01A8.47 8.47 0 0 1 12 3.5v-1a9.47 9.47 0 0 0-6.718 2.782zm-.708 0L18.01 18.718l.707-.708L5.99 5.282zM20.5 12c0 2.347-.95 4.472-2.49 6.01l.707.707A9.47 9.47 0 0 0 21.5 12zm-2.49 6.01A8.47 8.47 0 0 1 12 20.5v1a9.47 9.47 0 0 0 6.718-2.782z"/>`), g.Group(children),
	)
}

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m7 12l5 5m0 0l5-5m-5 5V4M6 20h12"/>`), g.Group(children),
	)
}

func UploadBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m7 12l5 5m0 0l5-5m-5 5V4M6 20h12"/>`), g.Group(children),
	)
}

func UploadFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 3a1 1 0 0 1 1 1v7h4a1 1 0 0 1 .707 1.707l-5 5a1 1 0 0 1-1.414 0l-5-5A1 1 0 0 1 7 11h4V4a1 1 0 0 1 1-1M5 20a1 1 0 0 1 1-1h12a1 1 0 1 1 0 2H6a1 1 0 0 1-1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UploadLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7 12l5 5m0 0l5-5m-5 5V4M6 20h12"/>`), g.Group(children),
	)
}

func UploadThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7 12l5 5m0 0l5-5m-5 5V4M6 20h12"/>`), g.Group(children),
	)
}

func UpsideDownFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M9 14.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2" d="M8.535 10A4 4 0 0 1 12 8c1.48 0 2.773.804 3.465 2"/></g>`), g.Group(children),
	)
}

func UpsideDownFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M9 14.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke-linecap="round" stroke-width="2.5" d="M8.535 10A4 4 0 0 1 12 8c1.48 0 2.773.804 3.465 2"/></g>`), g.Group(children),
	)
}

func UpsideDownFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9 14.5h.01v.01H9zm6 0h.01v.01H15z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M8.535 10A4 4 0 0 1 12 8c1.48 0 2.773.804 3.465 2"/></g>`), g.Group(children),
	)
}

func UpsideDownFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m13.966-1.134a1 1 0 0 1-1.367-.365C14.08 9.6 13.11 9 12 9s-2.08.601-2.6 1.5a1 1 0 0 1-1.73-1A5 5 0 0 1 12 7a5 5 0 0 1 4.33 2.5a1 1 0 0 1-.364 1.366M9 13a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5zm6 0a1.5 1.5 0 0 0-1.5 1.5v.01a1.5 1.5 0 0 0 1.5 1.5h.01a1.5 1.5 0 0 0 1.5-1.5v-.01a1.5 1.5 0 0 0-1.5-1.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func UpsideDownFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M9.01 14.5v.01H9v-.01zm6 0v.01H15v-.01z"/><path stroke-linecap="round" stroke-width="1.5" d="M8.535 10A4 4 0 0 1 12 8c1.48 0 2.773.804 3.465 2"/></g>`), g.Group(children),
	)
}

func UpsideDownFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M9.01 14.5v.01H9v-.01zm6 0v.01H15v-.01z"/><path stroke-linecap="round" d="M8.535 10A4 4 0 0 1 12 8c1.48 0 2.773.804 3.465 2"/></g>`), g.Group(children),
	)
}

func VolumeDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linejoin="round" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="M16.877 9a36.5 36.5 0 0 1 0 6"/></g>`), g.Group(children),
	)
}

func VolumeDownBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><path stroke-linejoin="round" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="M16.877 9a36.5 36.5 0 0 1 0 6"/></g>`), g.Group(children),
	)
}

func VolumeDownDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M16.877 9a36.5 36.5 0 0 1 0 6"/></g>`), g.Group(children),
	)
}

func VolumeDownFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.427 5.054c1.15-.816 2.828-.217 3.083 1.25c.322 1.85.49 3.754.49 5.696s-.168 3.845-.49 5.697c-.255 1.466-1.932 2.065-3.083 1.25L7.94 17.183A1 1 0 0 0 7.363 17H5a3 3 0 0 1-3-3v-4a3 3 0 0 1 3-3h2.363a1 1 0 0 0 .578-.184zm7.446 3.864a1 1 0 1 0-1.993.164a35.5 35.5 0 0 1 0 5.836a1 1 0 0 0 1.993.164a37.5 37.5 0 0 0 0-6.164" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func VolumeDownLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linejoin="round" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="M16.877 9a36.5 36.5 0 0 1 0 6"/></g>`), g.Group(children),
	)
}

func VolumeDownThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linejoin="round" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="M16.877 9a36.5 36.5 0 0 1 0 6"/></g>`), g.Group(children),
	)
}

func VolumeOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2"><path d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="m17 10l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func VolumeOffBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="2.5"><path d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="m17 10l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func VolumeOffDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m17 10l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func VolumeOffFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.427 5.054c1.15-.816 2.828-.217 3.083 1.25c.322 1.85.49 3.754.49 5.696s-.168 3.845-.49 5.697c-.255 1.466-1.932 2.065-3.083 1.25L7.94 17.183A1 1 0 0 0 7.363 17H5a3 3 0 0 1-3-3v-4a3 3 0 0 1 3-3h2.363a1 1 0 0 0 .578-.184zm7.28 4.239a1 1 0 0 0-1.414 1.414L17.586 12l-1.293 1.293a1 1 0 0 0 1.414 1.414L19 13.414l1.293 1.293a1 1 0 0 0 1.414-1.414L20.414 12l1.293-1.293a1 1 0 0 0-1.414-1.414L19 10.586z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func VolumeOffLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="m17 10l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func VolumeOffThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="m17 10l4 4m-4 0l4-4"/></g>`), g.Group(children),
	)
}

func VolumeUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2"><path stroke-linecap="round" d="M20.803 8a40.5 40.5 0 0 1 0 8"/><path stroke-linejoin="round" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="M16.877 9a36.5 36.5 0 0 1 0 6"/></g>`), g.Group(children),
	)
}

func VolumeUpBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="2.5"><path stroke-linecap="round" d="M20.803 8a40.5 40.5 0 0 1 0 8"/><path stroke-linejoin="round" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="M16.877 9a36.5 36.5 0 0 1 0 6"/></g>`), g.Group(children),
	)
}

func VolumeUpDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M20.803 8a40.5 40.5 0 0 1 0 8"/><path fill="currentColor" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12" opacity=".16"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="2" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M16.877 9a36.5 36.5 0 0 1 0 6"/></g>`), g.Group(children),
	)
}

func VolumeUpFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.51 6.303c-.255-1.466-1.932-2.065-3.083-1.25L7.94 6.817A1 1 0 0 1 7.363 7H5a3 3 0 0 0-3 3v4a3 3 0 0 0 3 3h2.363a1 1 0 0 1 .578.184l2.486 1.762c1.15.816 2.828.217 3.083-1.25c.322-1.85.49-3.754.49-5.696s-.168-3.845-.49-5.697m8.288 1.598a1 1 0 0 0-1.99.198a39.5 39.5 0 0 1 0 7.802a1 1 0 1 0 1.99.198a41.5 41.5 0 0 0 0-8.198m-3.925 1.017a1 1 0 1 0-1.993.164a35.5 35.5 0 0 1 0 5.836a1 1 0 0 0 1.993.164a37.5 37.5 0 0 0 0-6.164" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func VolumeUpLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M20.803 8a40.5 40.5 0 0 1 0 8"/><path stroke-linejoin="round" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="M16.877 9a36.5 36.5 0 0 1 0 6"/></g>`), g.Group(children),
	)
}

func VolumeUpThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path stroke-linecap="round" d="M20.803 8a40.5 40.5 0 0 1 0 8"/><path stroke-linejoin="round" d="M13 12c0-1.884-.163-3.73-.475-5.525c-.123-.704-.937-1.019-1.52-.605L8.52 7.632A2 2 0 0 1 7.363 8H5a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h2.363a2 2 0 0 1 1.157.368l2.485 1.762c.583.414 1.397.1 1.52-.605A32 32 0 0 0 13 12Z"/><path stroke-linecap="round" d="M16.877 9a36.5 36.5 0 0 1 0 6"/></g>`), g.Group(children),
	)
}

func WinkingFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2"/><path stroke-width="3" d="M9 9.5h.01v.01H9z"/><path stroke-linecap="round" stroke-width="2" d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2M16.5 9.5h-3"/></g>`), g.Group(children),
	)
}

func WinkingFaceBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="2.5"/><path stroke-width="3.75" d="M9 9.5h.01v.01H9z"/><path stroke-linecap="round" stroke-width="2.5" d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2M16.5 9.5h-3"/></g>`), g.Group(children),
	)
}

func WinkingFaceDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9" fill="currentColor" opacity=".16"/><circle cx="12" cy="12" r="9" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="3" d="M9 9.5h.01v.01H9z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2M16.5 9.5h-3"/></g>`), g.Group(children),
	)
}

func WinkingFaceFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m5.5-2.5A1.5 1.5 0 0 1 9 8h.01a1.5 1.5 0 0 1 1.5 1.5v.01a1.5 1.5 0 0 1-1.5 1.5H9a1.5 1.5 0 0 1-1.5-1.5zm6-1a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2zm-4.1 5a1 1 0 0 0-1.73 1A5 5 0 0 0 12 17a5 5 0 0 0 4.33-2.5a1 1 0 1 0-1.73-1c-.52.899-1.49 1.5-2.6 1.5a3 3 0 0 1-2.6-1.5" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func WinkingFaceLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round" stroke-width="1.5"/><path stroke-width="2.25" d="M9.01 9.5v.01H9V9.5z"/><path stroke-linecap="round" stroke-width="1.5" d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2M16.5 9.5h-3"/></g>`), g.Group(children),
	)
}

func WinkingFaceThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><circle cx="12" cy="12" r="9" stroke-linecap="round"/><path stroke-width="1.5" d="M9.01 9.5v.01H9V9.5z"/><path stroke-linecap="round" d="M15.465 14A4 4 0 0 1 12 16a4 4 0 0 1-3.465-2M16.5 9.5h-3"/></g>`), g.Group(children),
	)
}

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314M11 8v6m-3-3h6"/>`), g.Group(children),
	)
}

func ZoomInBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314M11 8v6m-3-3h6"/>`), g.Group(children),
	)
}

func ZoomInDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M19 11a8 8 0 1 1-16 0a8 8 0 0 1 16 0" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314M11 8v6m-3-3h6"/></g>`), g.Group(children),
	)
}

func ZoomInFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 11a9 9 0 1 1 16.032 5.618l3.675 3.675a1 1 0 0 1-1.414 1.414l-3.675-3.675A9 9 0 0 1 2 11m9-4a1 1 0 0 1 1 1v2h2a1 1 0 1 1 0 2h-2v2a1 1 0 1 1-2 0v-2H8a1 1 0 1 1 0-2h2V8a1 1 0 0 1 1-1" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ZoomInLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314M11 8v6m-3-3h6"/>`), g.Group(children),
	)
}

func ZoomInThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314M11 8v6m-3-3h6"/>`), g.Group(children),
	)
}

func ZoomOut(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314M8 11h6"/>`), g.Group(children),
	)
}

func ZoomOutBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.5" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314M8 11h6"/>`), g.Group(children),
	)
}

func ZoomOutDuotone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M19 11a8 8 0 1 1-16 0a8 8 0 0 1 16 0" opacity=".16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314M8 11h6"/></g>`), g.Group(children),
	)
}

func ZoomOutFill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 11a9 9 0 1 1 16.032 5.618l3.675 3.675a1 1 0 0 1-1.414 1.414l-3.675-3.675A9 9 0 0 1 2 11m6-1a1 1 0 1 0 0 2h6a1 1 0 1 0 0-2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ZoomOutLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314M8 11h6"/>`), g.Group(children),
	)
}

func ZoomOutThin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m21 21l-4.343-4.343m0 0A8 8 0 1 0 5.343 5.343a8 8 0 0 0 11.314 11.314M8 11h6"/>`), g.Group(children),
	)
}
