package proicons

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

const IconifyVersion = "4.10.0"

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 24 24")
)

func Accessibility(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.867 5.543a1.58 1.58 0 0 0-2.155.576a1.57 1.57 0 0 0 .577 2.15l2.63 1.515a1 1 0 0 1 .5.866v3.536a1 1 0 0 1-.134.5l-2.408 4.162a1.57 1.57 0 0 0 .577 2.15a1.58 1.58 0 0 0 2.156-.576l3.258-5.629h.258l3.258 5.629a1.58 1.58 0 0 0 2.156.576a1.57 1.57 0 0 0 .577-2.15l-2.402-4.15a1 1 0 0 1-.135-.502V10.65a1 1 0 0 1 .501-.866l2.63-1.514a1.57 1.57 0 0 0 .577-2.15a1.58 1.58 0 0 0-2.155-.577l-3.636 2.094a3 3 0 0 1-2.994 0z"/><path d="M14.623 5.414a2.623 2.623 0 1 1-5.246 0a2.623 2.623 0 0 1 5.246 0"/></g>`), g.Group(children),
	)
}

func Add(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 12h8m0 0h8m-8 0V4m0 8v8"/>`), g.Group(children),
	)
}

func AddCircular(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9.25"/><path d="M12 8.5v7M8.5 12h7"/></g>`), g.Group(children),
	)
}

func AddRhombus(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.775 14.475a3.5 3.5 0 0 1 0-4.95l5.75-5.75a3.5 3.5 0 0 1 4.95 0l5.75 5.75a3.5 3.5 0 0 1 0 4.95l-5.75 5.75a3.5 3.5 0 0 1-4.95 0zM8.25 12H12m0 0h3.75M12 12V8.25M12 12v3.75"/>`), g.Group(children),
	)
}

func AddSquare(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 8.5v7M8.5 12h7"/><rect width="16.5" height="16.5" x="3.75" y="3.75" rx="4"/></g>`), g.Group(children),
	)
}

func AddSquareMultiple(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.63 7.66v5.94m-2.97-2.97h5.94"/><rect width="14" height="14" x="3.63" y="3.63" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="3"/><path fill="currentColor" d="M20.971 7.928c.227.444.321.924.366 1.47c.043.531.043 1.187.043 2v.466c0 1.511 0 2.692-.077 3.64c-.08.963-.242 1.752-.604 2.463a6.25 6.25 0 0 1-2.732 2.732c-.711.362-1.5.525-2.463.604c-.948.077-2.129.077-3.64.077h-.466c-.813 0-1.469 0-2-.043c-.546-.045-1.026-.14-1.47-.366A3.75 3.75 0 0 1 6.63 19.88h5.2c1.553 0 2.672 0 3.551-.072c.871-.072 1.44-.209 1.906-.446a4.75 4.75 0 0 0 2.075-2.075c.237-.465.374-1.035.446-1.906c.071-.88.072-1.998.072-3.551v-5.2c.456.342.83.785 1.091 1.298"/></g>`), g.Group(children),
	)
}

func AddSquareMultipleVar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.63 7.66v5.94m-2.97-2.97h5.94"/><rect width="14" height="14" x="3.63" y="3.63" rx="3"/><path d="M20.63 7.63v7a6 6 0 0 1-6 6h-7"/></g>`), g.Group(children),
	)
}

func Airplane(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.238 4.512a1.762 1.762 0 1 1 3.524 0V8.9l6.733 3.535a1 1 0 0 1 .535.885v.431a.6.6 0 0 1-.712.59l-6.556-1.24v4.107a.6.6 0 0 0 .317.53l1.862.996a1.5 1.5 0 0 1 .792 1.322v.447a.6.6 0 0 1-.73.586L12 20.204l-4.003.885a.6.6 0 0 1-.73-.586v-.447a1.5 1.5 0 0 1 .792-1.322l1.862-.997a.6.6 0 0 0 .317-.529v-4.106L3.682 14.34a.6.6 0 0 1-.712-.59v-.43a1 1 0 0 1 .535-.886L10.238 8.9z"/>`), g.Group(children),
	)
}

func AirplaneLanding(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 20.75h18.5m-2.05-7.453a1.783 1.783 0 1 1-.923 3.445L8.185 14.04a4 4 0 0 1-1.32-.628l-2.14-1.543A3.04 3.04 0 0 1 3.47 9.271l.11-2.508a.607.607 0 0 1 .765-.56l.436.117a1.52 1.52 0 0 1 1.086 1.121l.486 2.082c.051.218.218.39.434.448l4.015 1.076l.506-6.735a.607.607 0 0 1 .763-.541l.422.113c.363.097.643.388.725.755l1.692 7.509z"/>`), g.Group(children),
	)
}

func AirplaneTakeoff(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 20.75h18.5M18.575 6.299a1.783 1.783 0 0 1 1.783 3.089L11.31 14.61a4 4 0 0 1-1.377.49l-2.604.422a3.04 3.04 0 0 1-2.725-.948L2.91 12.723a.607.607 0 0 1 .145-.936l.391-.226a1.52 1.52 0 0 1 1.56.025l1.816 1.128c.19.118.43.122.624.01l3.6-2.078l-4.404-5.12a.607.607 0 0 1 .156-.922l.378-.218c.326-.188.73-.18 1.047.02l6.506 4.113z"/>`), g.Group(children),
	)
}

func AlertCircle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9.25" stroke="currentColor" stroke-width="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 12.438v-5"/><circle cx="1.25" cy="1.25" r="1.25" fill="currentColor" transform="matrix(1 0 0 -1 10.75 17.063)"/></g>`), g.Group(children),
	)
}

func AlertTriangle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-width="1.5" d="M5.732 20.5c-2.29 0-3.723-2.498-2.581-4.5L9.419 5.006c1.144-2.008 4.018-2.008 5.163 0L20.849 16c1.142 2.002-.291 4.5-2.581 4.5z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 13.375V9"/><circle cx="1.25" cy="1.25" r="1.25" fill="currentColor" transform="matrix(1 0 0 -1 10.75 17.938)"/></g>`), g.Group(children),
	)
}

func AlignBottom(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 21.25h18.5"/><rect width="6" height="10" rx="2" transform="matrix(-1 0 0 1 19.75 7.75)"/><rect width="6" height="15" rx="2" transform="matrix(-1 0 0 1 10.25 2.75)"/></g>`), g.Group(children),
	)
}

func AlignHorizontalCenters(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19.75 12h1.5m-11 0h3.5m-11 0h1.5"/><rect width="6" height="10" rx="2" transform="matrix(-1 0 0 1 19.75 7)"/><rect width="6" height="15" rx="2" transform="matrix(-1 0 0 1 10.25 4.5)"/></g>`), g.Group(children),
	)
}

func AlignLeft(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 2.75v18.5"/><rect width="6" height="10" rx="2" transform="matrix(0 -1 -1 0 16.25 19.75)"/><rect width="6" height="15" rx="2" transform="matrix(0 -1 -1 0 21.25 10.25)"/></g>`), g.Group(children),
	)
}

func AlignRight(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 2.75v18.5"/><rect width="6" height="10" x="7.75" y="19.75" rx="2" transform="rotate(-90 7.75 19.75)"/><rect width="6" height="15" x="2.75" y="10.25" rx="2" transform="rotate(-90 2.75 10.25)"/></g>`), g.Group(children),
	)
}

func AlignTop(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 2.75h18.5"/><rect width="6" height="10" x="19.75" y="16.25" rx="2" transform="rotate(180 19.75 16.25)"/><rect width="6" height="15" x="10.25" y="21.25" rx="2" transform="rotate(180 10.25 21.25)"/></g>`), g.Group(children),
	)
}

func AlignVerticalCenters(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 19.75v1.5m0-11v3.5m0-11v1.5"/><rect width="6" height="10" x="7" y="19.75" rx="2" transform="rotate(-90 7 19.75)"/><rect width="6" height="15" x="4.5" y="10.25" rx="2" transform="rotate(-90 4.5 10.25)"/></g>`), g.Group(children),
	)
}

func Amazon(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.3" d="M17.284 17.724h1.967c.543 0 .983.44.983.983v1.967"/><path fill="currentColor" fill-rule="evenodd" d="m6.901 7.111l1.976.2a.57.57 0 0 0 .558-.34c.268-.573.924-1.59 2.209-1.59c.258 0 .478.041.665.108c1.04.373 1.02 1.77 1.02 2.874c0 0-2.415-.058-3.992.476c-.465.157-.934.333-1.36.578c-1.223.705-1.77 2.276-1.969 3.022c-.071.27-.098.549-.08.828c.08 1.178.558 3.184 2.968 3.599c2.247.386 3.749-.64 4.43-1.267c.23-.213.612-.22.817.018l.88 1.017a.59.59 0 0 0 .825.066l1.676-1.402a.59.59 0 0 0 .101-.796l-.267-.374c-.215-.3-.422-.613-.517-.97c-.099-.372-.09-.763-.09-1.148V6.704c0-2.372-2.295-3.681-3.94-3.85l-.592-.057a6 6 0 0 0-.972-.034a6 6 0 0 0-1.162.196c-1.537.399-3.348 1.765-3.597 3.664c-.033.25.162.463.413.488m6.428 4.576v-1.432s-1.905-.047-2.664.363c-.312.168-.574.491-.775.813a2.54 2.54 0 0 0-.291 1.94c.066.275.158.54.282.695c.259.323.876.439 1.166.476q.156.022.311.003c.157-.022.413-.08.67-.23c.39-.229.695-.595.903-.995c.118-.226.208-.471.273-.676a3.2 3.2 0 0 0 .125-.957" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M4.5 18.707c1.933 1.619 4.584 2.548 7.375 2.548c1.853 0 3.643-.41 5.203-1.158a11 11 0 0 0 .573-.295"/></g>`), g.Group(children),
	)
}

func Anchor(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M20.693 13.5a8.75 8.75 0 0 1-17.386 0M12 21.25V7.75"/><circle cx="12" cy="5.25" r="2.5"/><path stroke-linecap="round" d="M9 11.246h6M3.307 13.5h3m11.386 0h3"/></g>`), g.Group(children),
	)
}

func Angle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.25 4.25v12.5a3 3 0 0 0 3 3h12.5"/><path d="M4.25 10.356h.394a9 9 0 0 1 9 9v.394"/></g>`), g.Group(children),
	)
}

func AppRemove(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="6.5" height="6.5" rx="2" transform="matrix(1 0 0 -1 3.75 20.25)"/><path d="M20.25 20.25L17 17m0 0l-3.25-3.25M17 17l-3.25 3.25M17 17l3.25-3.25"/><rect width="6.5" height="6.5" rx="2" transform="matrix(1 0 0 -1 3.75 10.25)"/><rect width="6.5" height="6.5" rx="2" transform="matrix(1 0 0 -1 13.75 10.25)"/></g>`), g.Group(children),
	)
}

func AppStore(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M10.904 13.794h2.024l-1.012-1.755z"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="m8.843 17.367l-1.58 2.74A1.784 1.784 0 1 1 4.17 18.32l.55-.953m4.122 0H4.72m4.122 0h6.146m-10.268 0h-.466a1.786 1.786 0 0 1 0-3.573h2.527l3.073-5.327l-1.03-1.787a1.784 1.784 0 1 1 3.092-1.786a1.784 1.784 0 1 1 3.09 1.786l-1.03 1.786m-1.049 5.328h-2.024l1.012-1.755m1.012 1.755l2.061 3.573m-2.06-3.573l-1.013-1.755m3.073 5.328l1.58 2.74a1.784 1.784 0 1 0 3.092-1.787l-.55-.953h.634a1.786 1.786 0 0 0 0-3.573H17.05l-3.073-5.328m0 0l-2.06 3.573"/></g>`), g.Group(children),
	)
}

func Apple(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round"><path stroke-width="1.5" d="M10.102 7.78a4 4 0 0 0-1.185-.161c-1.918 0-4.167 1.695-4.167 5.359c0 3.36 2.358 7.21 3.801 7.822c.328.14.731.227 1.084.193c.707-.07 1.348-.421 2.023-.622c.228-.038.5-.074.719-.075c.24-.002.51.06.774.12c.715.164 1.505.65 2.255.577c.267-.026.633-.187.746-.239l.07-.03c.85-.343 1.988-1.826 2.759-3.703c.162-.396-.052-.833-.404-1.077c-.999-.694-1.663-1.934-1.663-3.195c0-.969.313-1.854.907-2.523l.02-.023a8 8 0 0 1 .45-.5c.255-.248.326-.645.082-.904c-.808-.854-1.832-1.277-2.764-1.277a3.6 3.6 0 0 0-1.106.156c-.138.033-.778.192-1.222.355c-.257.071-.635.167-.81.21a.7.7 0 0 1-.192.023c-.735-.024-1.47-.303-2.177-.487Z" clip-rule="evenodd"/><path fill="currentColor" d="M13.86 5.178c-.733.878-1.975 1.047-1.975 1.047s-.05-1.256.684-2.135s1.976-1.046 1.976-1.046s.05 1.256-.684 2.134Z"/></g>`), g.Group(children),
	)
}

func Apps(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="6.5" height="6.5" x="3.75" y="3.75" rx="2"/><path d="M15.586 3.818a2 2 0 0 1 2.828 0l1.768 1.768a2 2 0 0 1 0 2.828l-1.768 1.768a2 2 0 0 1-2.828 0l-1.768-1.768a2 2 0 0 1 0-2.828z"/><rect width="6.5" height="6.5" x="3.75" y="13.75" rx="1.5"/><rect width="6.5" height="6.5" x="13.75" y="13.75" rx="2"/></g>`), g.Group(children),
	)
}

func AppsAdd(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="6.5" height="6.5" x="3.75" y="3.75" rx="2"/><path d="M17 3.75V7m0 0v3.25M17 7h-3.25M17 7h3.25"/><rect width="6.5" height="6.5" x="3.75" y="13.75" rx="2"/><rect width="6.5" height="6.5" x="13.75" y="13.75" rx="2"/></g>`), g.Group(children),
	)
}

func AppsList(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.25 4.5h11"/><rect width="3.5" height="3.5" x="2.75" y="2.75" rx="1.2"/><path d="M10.25 12h11"/><rect width="3.5" height="3.5" x="2.75" y="10.25" rx="1.2"/><path d="M10.25 19.5h11"/><rect width="3.5" height="3.5" x="2.75" y="17.75" rx="1.2"/></g>`), g.Group(children),
	)
}

func Arc(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.25 4.75h-.5c-7.732 0-14 6.268-14 14v.5"/>`), g.Group(children),
	)
}

func Archive(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.25 12.5h3.5m-10-4.75a4 4 0 0 1 4-4h8.5a4 4 0 0 1 4 4v8.5a4 4 0 0 1-4 4h-8.5a4 4 0 0 1-4-4zm0 1h16.5"/>`), g.Group(children),
	)
}

func ArchiveAddTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.25 12.25h2m-8.5-4h16.5m0 2v-1.7c0-1.68 0-2.52-.327-3.162a3 3 0 0 0-1.311-1.311c-.642-.327-1.482-.327-3.162-.327h-6.9c-1.68 0-2.52 0-3.162.327a3 3 0 0 0-1.311 1.311C3.75 6.03 3.75 6.87 3.75 8.55v6.9c0 1.68 0 2.52.327 3.162a3 3 0 0 0 1.311 1.311c.642.327 1.482.327 3.162.327h1.7"/><path fill="currentColor" fill-rule="evenodd" d="M17.5 23a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-8.993a.5.5 0 0 1 .5.5V17h2.493a.5.5 0 1 1 0 1H18v2.493a.5.5 0 1 1-1 0V18h-2.493a.5.5 0 1 1 0-1H17v-2.493a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func ArrowClockwise(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.719 14.385a8.25 8.25 0 1 1-.824-6.26l.42.908m.58-4.658v3.75a1 1 0 0 1-.58.908m-4.17.092h3.75c.15 0 .293-.033.42-.092"/>`), g.Group(children),
	)
}

func ArrowCounterclockwise(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.281 14.385a8.25 8.25 0 1 0 .824-6.26l-.477.88m-.523-4.63v3.75a1 1 0 0 0 .523.88m4.227.12h-3.75a1 1 0 0 1-.477-.12"/>`), g.Group(children),
	)
}

func ArrowDown(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 4v15.879M5.25 13.75l5.69 5.69c.292.292.676.439 1.06.439m6.75-6.129l-5.69 5.69a1.5 1.5 0 0 1-1.06.439"/>`), g.Group(children),
	)
}

func ArrowDownload(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.074 3.25v12.478M6.19 10.465l4.822 4.822c.293.293.677.44 1.06.44m5.883-5.262l-4.822 4.822c-.293.293-.677.44-1.06.44m8.677.788v.935a3.3 3.3 0 0 1-3.3 3.3H6.55a3.3 3.3 0 0 1-3.3-3.3v-.935"/>`), g.Group(children),
	)
}

func ArrowEnter(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.75 5.623V9.52a4 4 0 0 1-4 4H3.871m4.236 4.857L4.31 14.58a1.5 1.5 0 0 1-.44-1.061m4.236-4.857L4.31 12.46c-.293.293-.44.677-.44 1.061"/>`), g.Group(children),
	)
}

func ArrowExport(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.798 12H2.871m5.157-5.778l-4.717 4.717c-.293.293-.44.677-.44 1.061m5.157 5.778l-4.717-4.717A1.5 1.5 0 0 1 2.87 12m17.88-7.905v15.81"/>`), g.Group(children),
	)
}

func ArrowFoward(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.25 18.219c0-2.352 0-3.527.383-4.455a5.06 5.06 0 0 1 2.743-2.743c.928-.383 2.103-.383 4.455-.383h8.298m-4.236-4.857l3.796 3.796c.293.293.44.677.44 1.061m-4.236 4.857l3.796-3.796c.293-.293.44-.677.44-1.061"/>`), g.Group(children),
	)
}

func ArrowImport(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.25 12h12.927M10.02 6.222l4.717 4.717c.293.293.44.677.44 1.061m-5.157 5.778l4.717-4.717c.293-.293.44-.677.44-1.061m5.573-7.905v15.81"/>`), g.Group(children),
	)
}

func ArrowLeft(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20 12H4.121m6.129 6.75l-5.69-5.69A1.5 1.5 0 0 1 4.122 12m6.129-6.75l-5.69 5.69A1.5 1.5 0 0 0 4.122 12"/>`), g.Group(children),
	)
}

func ArrowLeftRight(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m16.25 17l3.94-3.94c.292-.292.439-.676.439-1.06M16.25 7l3.94 3.94c.292.292.439.676.439 1.06M7.75 17l-3.94-3.94A1.5 1.5 0 0 1 3.372 12M7.75 7l-3.94 3.94A1.5 1.5 0 0 0 3.371 12m0 0H20.63"/>`), g.Group(children),
	)
}

func ArrowMaximize(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 12.75v6c0 .414.168.79.44 1.06m7.06.44h-6a1.5 1.5 0 0 1-1.06-.44m16.06-8.56v-6c0-.414-.168-.79-.44-1.06m-7.06-.44h6c.414 0 .79.168 1.06.44M4.19 19.81l.56-.56l14.5-14.5l.56-.56"/>`), g.Group(children),
	)
}

func ArrowMinimize(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.75 21.25l6.5-6.5l.56-.56m.44 7.06v-6c0-.414-.168-.79-.44-1.06m-7.06-.44h6c.414 0 .79.168 1.06.44m3.94-11.44v6c0 .414.168.79.44 1.06m7.06.44h-6a1.5 1.5 0 0 1-1.06-.44m7.06-7.06l-6.5 6.5l-.56.56"/>`), g.Group(children),
	)
}

func ArrowMove(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9.005V2.75M9 5.324l2.273-2.273c.201-.2.464-.301.727-.301m3 2.574l-2.273-2.273c-.2-.2-.464-.301-.727-.301M14.995 12h6.255m-2.574-3l2.273 2.273c.2.201.301.464.301.727m-2.574 3l2.273-2.273c.2-.2.301-.464.301-.727M12 14.995v6.255m-3-2.574l2.273 2.273c.201.2.464.301.727.301m3-2.574l-2.273 2.273c-.2.2-.464.301-.727.301M9.005 12H2.75m2.574-3l-2.273 2.273c-.2.201-.301.464-.301.727m2.574 3l-2.273-2.273c-.2-.2-.301-.464-.301-.727"/>`), g.Group(children),
	)
}

func ArrowRedo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.03 21.25l-8.26-8.26a5.999 5.999 0 0 1 8.483-8.483l4.734 4.734l.622.622m-7.199.505h6.077c.447 0 .848-.195 1.122-.505m.378-7.072v6.077c0 .382-.142.73-.378.995"/>`), g.Group(children),
	)
}

func ArrowRedoTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.75 12l3.293-3.293A1 1 0 0 0 18.336 8M14.75 4l3.293 3.293a1 1 0 0 1 .293.707M16.75 19.5H10.5a5.75 5.75 0 0 1 0-11.5h7.836"/>`), g.Group(children),
	)
}

func ArrowReply(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.75 18.219c0-2.352 0-3.527-.383-4.455a5.06 5.06 0 0 0-2.743-2.743c-.928-.383-2.103-.383-4.455-.383H3.871m4.236-4.857L4.31 9.577c-.293.293-.44.677-.44 1.061m4.236 4.857L4.31 11.699a1.5 1.5 0 0 1-.44-1.061"/>`), g.Group(children),
	)
}

func ArrowRight(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4 12h15.879m-6.129 6.75l5.69-5.69c.292-.292.439-.676.439-1.06M13.75 5.25l5.69 5.69c.292.292.439.676.439 1.06"/>`), g.Group(children),
	)
}

func ArrowRotateClockwise(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M3.75 12a8.25 8.25 0 1 1 12.375 7.145l-1.199.421m-.051-3.316v3q0 .166.051.316m3.949.684h-3a1 1 0 0 1-.949-.684"/><path d="M14.5 11.75a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0Z"/></g>`), g.Group(children),
	)
}

func ArrowRotateCounterclockwise(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" stroke-linejoin="round" d="M20.25 12a8.25 8.25 0 1 0-12.375 7.145l1.168.503m.082-3.398v3q-.002.213-.082.398m-3.918.602h3a1 1 0 0 0 .918-.602"/><path d="M14.5 11.75a2.25 2.25 0 1 1-4.5 0a2.25 2.25 0 0 1 4.5 0Z"/></g>`), g.Group(children),
	)
}

func ArrowSort(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 9.867L8.186 6.053a1.5 1.5 0 0 0-1.061-.44M2.25 9.868l3.814-3.814c.293-.293.677-.44 1.061-.44m0 13.395V5.614m9.75-.124v13.394m4.875-4.253l-3.814 3.814c-.293.293-.677.44-1.061.44M12 14.63l3.814 3.814c.293.293.677.44 1.061.44"/>`), g.Group(children),
	)
}

func ArrowSwap(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.618 12.249l-3.814 3.814c-.293.293-.44.677-.44 1.06M9.619 22l-3.814-3.814a1.5 1.5 0 0 1-.44-1.061m13.395 0H5.365m-.124-9.751h13.394m-4.253-4.875l3.814 3.814c.293.293.44.677.44 1.06m-4.254 4.876l3.814-3.814c.293-.293.44-.677.44-1.061"/>`), g.Group(children),
	)
}

func ArrowSync(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.548 9.735a8.75 8.75 0 0 1 16.03-2.11l.335.759m.837-5.134v4.147a1 1 0 0 1-.837.987m-4.31.013h4.147q.083 0 .163-.013M3.25 20.75v-4.147a1 1 0 0 1 1-1m0 0h4.147m-4.147 0l.172.772a8.75 8.75 0 0 0 16.03-2.11"/>`), g.Group(children),
	)
}

func ArrowSyncTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.915 18.532A7.83 7.83 0 0 1 11.831 3.92h1.043m-1.755-2.17l1.462 1.462a1 1 0 0 1 .293.708m-1.755 2.169l1.462-1.462a1 1 0 0 0 .293-.707m3.211 1.299a7.83 7.83 0 0 1-3.916 14.612h-1.043M12.881 22l-1.462-1.462a1 1 0 0 1-.293-.707m1.755-2.17l-1.462 1.462a1 1 0 0 0-.293.708"/>`), g.Group(children),
	)
}

func ArrowTrending(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.25 16.75l4.793-4.793a1 1 0 0 1 1.414 0l2.586 2.586a1 1 0 0 0 1.414 0L19.75 8.25l.56-.56m-5.56-.44h4.5c.414 0 .79.168 1.06.44m.44 5.56v-4.5c0-.414-.168-.79-.44-1.06"/>`), g.Group(children),
	)
}

func ArrowUndo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.97 21.25l8.26-8.26a5.999 5.999 0 0 0-8.483-8.483L5.013 9.24l-.622.622m7.199.505H5.513c-.447 0-.848-.195-1.122-.505M4.013 2.79v6.077c0 .382.143.73.378.995"/>`), g.Group(children),
	)
}

func ArrowUndoTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.25 12L5.957 8.707A1 1 0 0 1 5.664 8M9.25 4L5.957 7.293A1 1 0 0 0 5.664 8M7.25 19.5h6.25a5.75 5.75 0 0 0 0-11.5H5.664"/>`), g.Group(children),
	)
}

func ArrowUp(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 20V4.121M5.25 10.25l5.69-5.69A1.5 1.5 0 0 1 12 4.121m6.75 6.129l-5.69-5.69A1.5 1.5 0 0 0 12 4.122"/>`), g.Group(children),
	)
}

func ArrowUpDown(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7 16.25l3.94 3.94c.292.292.676.439 1.06.439m5-4.379l-3.94 3.94a1.5 1.5 0 0 1-1.06.439M7 7.75l3.94-3.94A1.5 1.5 0 0 1 12 3.371m5 4.379l-3.94-3.94A1.5 1.5 0 0 0 12 3.372m0 0V20.63"/>`), g.Group(children),
	)
}

func ArrowUpload(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 16.349V3.87M6.118 9.132l4.821-4.821c.293-.293.677-.44 1.061-.44m5.882 5.261l-4.821-4.821A1.5 1.5 0 0 0 12 3.87m8.75 12.645v.935a3.3 3.3 0 0 1-3.3 3.3H6.55a3.3 3.3 0 0 1-3.3-3.3v-.935"/>`), g.Group(children),
	)
}

func Attatch(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.861 13.625l9.342-9.342a5.235 5.235 0 1 1 7.403 7.403L10.88 20.41a2.867 2.867 0 0 1-4.054-4.054l8.372-8.373"/>`), g.Group(children),
	)
}

func BackgroundColor(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M13.04 4.75h3.66c1.68 0 2.52 0 3.162.327a3 3 0 0 1 1.311 1.311c.327.642.327 1.482.327 3.162v5.4c0 1.68 0 2.52-.327 3.162a3 3 0 0 1-1.311 1.311c-.642.327-1.482.327-3.162.327H7.8c-1.68 0-2.52 0-3.162-.327a3 3 0 0 1-1.311-1.311C3 17.47 3 16.63 3 14.95v-.93"/><path d="m6.407 2.818l5.384 5.385M6.407 2.818L3.01 6.214c-.696.696-1.044 1.044-1.174 1.446c-.057.176-.086.36-.086.543m4.657-5.385L5.589 2m6.202 6.203l-3.396 3.396c-.696.696-1.044 1.044-1.445 1.174a1.76 1.76 0 0 1-1.086 0c-.401-.13-.75-.478-1.445-1.174L3.01 10.191c-.696-.696-1.044-1.044-1.174-1.445a1.8 1.8 0 0 1-.086-.543m10.042 0H1.75m10.672 4.094l1.485-2.448l1.744 2.313a1.824 1.824 0 0 1-.668 2.491c-1.546.893-3.322-.74-2.562-2.356"/></g>`), g.Group(children),
	)
}

func BackgroundColorAccent(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3 14.95v-1.68h.633c.168.163.33.33.486.462c.307.261.696.54 1.204.706a3.5 3.5 0 0 0 2.168 0c.507-.165.896-.445 1.203-.706c.28-.236.583-.54.895-.852l1.185-1.186c-1.322 3.12 2.099 6.198 5.084 4.475a3.574 3.574 0 0 0 1.308-4.882a2 2 0 0 0-.118-.179l-1.744-2.313a1.75 1.75 0 0 0-1.764-.657a1.75 1.75 0 0 0-.511-1.173l-.739-.716V4.75h4.41c1.68 0 2.52 0 3.162.327a3 3 0 0 1 1.311 1.311c.327.642.327 1.482.327 3.162v5.4c0 1.68 0 2.52-.327 3.162a3 3 0 0 1-1.311 1.311c-.642.327-1.482.327-3.162.327H7.8c-1.68 0-2.52 0-3.162-.327a3 3 0 0 1-1.311-1.311C3 17.47 3 16.63 3 14.95" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func Backspace(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.71 6.084a3.5 3.5 0 0 1 2.572-1.126h7.968a3.5 3.5 0 0 1 3.5 3.5v7.084a3.5 3.5 0 0 1-3.5 3.5H9.282a3.5 3.5 0 0 1-2.571-1.126l-3.27-3.542a3.5 3.5 0 0 1 0-4.748zm4.11 3.281l5.271 5.27m0-5.27l-5.27 5.27"/>`), g.Group(children),
	)
}

func Badge(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 3.75H8.75a5 5 0 0 0-5 5v6.5a5 5 0 0 0 5 5h6.5a5 5 0 0 0 5-5V12"/><circle cx="18.25" cy="5.75" r="3"/></g>`), g.Group(children),
	)
}

func BarGraph(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 21h18.5"/><rect width="6" height="10" x="4" y="7.5" rx="2"/><rect width="6" height="15" x="14" y="2.5" rx="2"/></g>`), g.Group(children),
	)
}

func Battery(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="16" height="12" x="2.75" y="6" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="3.25"/><rect width="7" height="7" x="5.25" y="8.5" fill="currentColor" rx="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.25 10v4"/></g>`), g.Group(children),
	)
}

func BatteryFull(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="16" height="12" x="2.75" y="6" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="3.25"/><rect width="11" height="7" x="5.25" y="8.5" fill="currentColor" rx="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.25 10v4"/></g>`), g.Group(children),
	)
}

func Beach(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.135 4.027c1.969 1.137 3.369 3.074 3.89 5.386a9.4 9.4 0 0 1 .005 4.066c-.361 1.67-2.311 2.182-3.791 1.328l-9.492-5.48c-1.48-.855-2.012-2.8-.746-3.948a9.4 9.4 0 0 1 3.523-2.03c2.263-.703 4.641-.46 6.61.678m0 0c-1.368-.79-4.554 2.17-7.118 6.61m7.118-6.61c1.367.789.397 5.028-2.167 9.469m-2.475-1.43l-3.61 6.254m7.027 2.725a9.306 9.306 0 0 0-13.159 0"/>`), g.Group(children),
	)
}

func Beaker(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.456 3.75v5.09a3 3 0 0 1-.557 1.742l-1.736 2.436M9.456 3.75h-1.65m1.65 0h5.088m0 0v5.09a3 3 0 0 0 .557 1.742l1.736 2.436M14.544 3.75h1.65m-9.031 9.268l-2.378 3.337a2.465 2.465 0 0 0 2.007 3.895h10.416a2.465 2.465 0 0 0 2.007-3.895l-2.378-3.337m-9.674 0h9.674"/>`), g.Group(children),
	)
}

func Bell(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="M18.934 14.98a3 3 0 0 1-.457-1.59V9.226a6.477 6.477 0 0 0-12.954 0v4.162a3 3 0 0 1-.457 1.592l-1.088 1.74a1 1 0 0 0 .848 1.53h14.348a1 1 0 0 0 .848-1.53z"/><path d="M10 21.25h4"/></g>`), g.Group(children),
	)
}

func Bluetooth(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.25 7.5l5.454 4.5m0 0l5.246 3.982a.65.65 0 0 1 0 1.036l-4.202 3.19a.65.65 0 0 1-1.043-.517V4.31a.65.65 0 0 1 1.043-.518l4.202 3.19a.65.65 0 0 1 0 1.036zm0 0L6.25 16.5"/>`), g.Group(children),
	)
}

func Board(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.75 4H6.25a3.5 3.5 0 0 0-3.5 3.5v9a3.5 3.5 0 0 0 3.5 3.5h11.5a3.5 3.5 0 0 0 3.5-3.5v-9a3.5 3.5 0 0 0-3.5-3.5"/><path fill="currentColor" d="M17.1 6.95H6.9a1.2 1.2 0 0 0-1.2 1.2v.483a1.2 1.2 0 0 0 1.2 1.2h10.2a1.2 1.2 0 0 0 1.2-1.2V8.15a1.2 1.2 0 0 0-1.2-1.2m0 5.64h-2.9a1.2 1.2 0 0 0-1.2 1.2v2.06a1.2 1.2 0 0 0 1.2 1.2h2.9a1.2 1.2 0 0 0 1.2-1.2v-2.06a1.2 1.2 0 0 0-1.2-1.2m-8.1 0H6.9a1.2 1.2 0 0 0-1.2 1.2v2.06a1.2 1.2 0 0 0 1.2 1.2H9a1.2 1.2 0 0 0 1.2-1.2v-2.06a1.2 1.2 0 0 0-1.2-1.2"/></g>`), g.Group(children),
	)
}

func Book(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.5 16.625v-10.5a3 3 0 0 1 3-3h11a1 1 0 0 1 1 1v12.5H7.375M4.5 16.62V19"/><path d="M18.5 21.625H7a2.5 2.5 0 0 1 0-5h12.5v4a1 1 0 0 1-1 1"/></g>`), g.Group(children),
	)
}

func BookAdd(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M17.5 23a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-8.993a.5.5 0 0 1 .5.5V17h2.493a.5.5 0 1 1 0 1H18v2.493a.5.5 0 1 1-1 0V18h-2.493a.5.5 0 1 1 0-1H17v-2.493a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 19.125a2.5 2.5 0 0 0 2.5 2.5h4.61m-7.11-2.5a2.5 2.5 0 0 1 2.5-2.5h3.32m-5.82 2.5v-13a3 3 0 0 1 3-3h11a1 1 0 0 1 1 1v6.495"/></g>`), g.Group(children),
	)
}

func BookAddTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 18.749v-12.8c0-1.12 0-1.68.218-2.108a2 2 0 0 1 .874-.874c.428-.218.988-.218 2.108-.218h8.6c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.427.218.987.218 2.108v4.301m-15 8.499c0 .467 0 .7.039.895a2 2 0 0 0 1.568 1.569c.195.039.429.039.896.039h3.247M4.5 18.749h4.75M8.96 9.25h6.08c.336 0 .504 0 .632-.065a.6.6 0 0 0 .263-.263C16 8.794 16 8.626 16 8.29V7.21c0-.336 0-.504-.065-.632a.6.6 0 0 0-.263-.263c-.128-.065-.296-.065-.632-.065H8.96c-.336 0-.504 0-.632.065a.6.6 0 0 0-.263.263C8 6.706 8 6.874 8 7.21v1.08c0 .336 0 .504.065.632a.6.6 0 0 0 .263.263c.128.065.296.065.632.065"/><path fill="currentColor" fill-rule="evenodd" d="M17.5 23a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-8.993a.5.5 0 0 1 .5.5V17h2.493a.5.5 0 1 1 0 1H18v2.493a.5.5 0 1 1-1 0V18h-2.493a.5.5 0 1 1 0-1H17v-2.493a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func BookInfo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 16.625v-10.5a3 3 0 0 1 3-3h11a1 1 0 0 1 1 1v12.5H7.375M4.5 16.62V19"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.5 21.625H7a2.5 2.5 0 0 1 0-5h12.5v4a1 1 0 0 1-1 1"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 9.825v3.79"/><circle cx="12" cy="6.592" r="1.197" fill="currentColor"/></g>`), g.Group(children),
	)
}

func BookInfoTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 4.749a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v13h-15z"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 10.07v4.79"/><circle cx="12" cy="6.837" r="1.197" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 17.749v1.503a2 2 0 0 0 2 2h13"/></g>`), g.Group(children),
	)
}

func BookOpen(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 7.21a2 2 0 0 1 2-2H8.5a3.5 3.5 0 0 1 3.5 3.5v10.885l-1.015-.721a4 4 0 0 0-2.318-.74H4.75a2 2 0 0 1-2-2zm18.5 0a2 2 0 0 0-2-2H15.5a3.5 3.5 0 0 0-3.5 3.5v10.885l1.015-.721a4 4 0 0 1 2.317-.74h3.918a2 2 0 0 0 2-2z"/>`), g.Group(children),
	)
}

func BookTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.5 4.749a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v13h-15zm0 13v1.503a2 2 0 0 0 2 2h13"/><rect width="8" height="3" x="8" y="6.25" rx=".6"/></g>`), g.Group(children),
	)
}

func Bookmark(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M4.421 6.75v13a1 1 0 0 0 1.619.786l5.342-4.205a1 1 0 0 1 1.236 0l5.342 4.205a1 1 0 0 0 1.619-.786v-13a3.5 3.5 0 0 0-3.5-3.5H7.921a3.5 3.5 0 0 0-3.5 3.5Z"/>`), g.Group(children),
	)
}

func BookmarkAdd(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M17.5 12a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-8.993a.5.5 0 0 1 .5.5V6h2.493a.5.5 0 1 1 0 1H18v2.493a.5.5 0 0 1-1 0V7h-2.493a.5.5 0 1 1 0-1H17V3.507a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.25 3.25H7.921a3.5 3.5 0 0 0-3.5 3.5v13a1 1 0 0 0 1.619.786l5.342-4.205a1 1 0 0 1 1.236 0l5.342 4.205a1 1 0 0 0 1.619-.786v-6"/></g>`), g.Group(children),
	)
}

func BookmarkMultiple(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M4.421 8.45v11.884c0 .763.88 1.19 1.48.718l4.883-3.843a.914.914 0 0 1 1.13 0l4.884 3.843a.914.914 0 0 0 1.48-.718V8.45a3.2 3.2 0 0 0-3.2-3.2H7.621a3.2 3.2 0 0 0-3.2 3.2Z"/><path fill="currentColor" fill-rule="evenodd" d="M7.421 3a1.5 1.5 0 0 1 1.5-1.5h7.357a5.75 5.75 0 0 1 5.75 5.75v8.584a1.5 1.5 0 0 1-1.5 1.5V7.25A4.25 4.25 0 0 0 16.278 3z" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func BookmarkMultipleVar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M4.421 8.45v11.884c0 .763.88 1.19 1.48.718l4.883-3.843a.914.914 0 0 1 1.13 0l4.884 3.843a.914.914 0 0 0 1.48-.718V8.45a3.2 3.2 0 0 0-3.2-3.2H7.621a3.2 3.2 0 0 0-3.2 3.2Z"/><path stroke-linecap="round" d="M21.278 16.334V8.25a6 6 0 0 0-6-6H8.421"/></g>`), g.Group(children),
	)
}

func BorderAll(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 20.25h4.25a4 4 0 0 0 4-4V12M12 20.25H7.75a4 4 0 0 1-4-4V12M12 20.25V3.75m0 0H7.75a4 4 0 0 0-4 4V12M12 3.75h4.25a4 4 0 0 1 4 4V12m-16.5 0h16.5"/>`), g.Group(children),
	)
}

func Box(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="m7.687 9.687l2.66 1.426A3.5 3.5 0 0 0 12 11.53M7.687 9.687L3.884 7.65m3.803 2.038l8.496-4.555l.128-.07M3.884 7.648a3.5 3.5 0 0 0-.51 1.82v5.061a3.5 3.5 0 0 0 1.845 3.085l5.127 2.748A3.5 3.5 0 0 0 12 20.78M3.884 7.649a3.5 3.5 0 0 1 1.335-1.264l5.127-2.748a3.5 3.5 0 0 1 3.308 0L16.31 5.06M12 11.53a3.5 3.5 0 0 0 1.654-.416l6.462-3.464M12 11.529v9.25m0 0a3.5 3.5 0 0 0 1.654-.416l5.127-2.748a3.5 3.5 0 0 0 1.846-3.085V9.47a3.5 3.5 0 0 0-.511-1.821m0 0a3.5 3.5 0 0 0-1.335-1.264l-2.47-1.324"/>`), g.Group(children),
	)
}

func BoxDrag(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.073 17.533v-3.864m2.544-2.937c0-.699-.57-1.265-1.272-1.265s-1.272.566-1.272 1.265v2.937m2.544-.444v-4.21c0-.698.57-1.265 1.272-1.265s1.272.567 1.272 1.266v1.716m0 0v2.493m0-2.493a1.273 1.273 0 0 1 2.545 0v1.717m0 0v.776m0-.776c0-.7.57-1.266 1.272-1.266s1.272.567 1.272 1.266V17.1a5.15 5.15 0 0 1-5.15 5.15h-2.31a5.55 5.55 0 0 1-5.541-5.872l.012-.201a2.4 2.4 0 0 1 1.67-2.146l1.142-.362"/><path d="M6.75 18.25a3 3 0 0 1-3-3v-8.5a3 3 0 0 1 3-3h8.5a3 3 0 0 1 2.959 2.5"/></g>`), g.Group(children),
	)
}

func BoxMargins(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="16.5" height="16.5" x="3.75" y="3.75" rx="3"/><path d="M7.75 3.75v16.5m8.5 0V3.75m-12.5 4h16.5m0 8.5H3.75"/></g>`), g.Group(children),
	)
}

func Braces(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 20.25a2 2 0 0 1-2-2v-2.343a4 4 0 0 0-1.172-2.829L3.75 12l1.078-1.078A4 4 0 0 0 6 8.093V5.75a2 2 0 0 1 2-2m8 16.5a2 2 0 0 0 2-2v-2.343a4 4 0 0 1 1.172-2.829L20.25 12l-1.078-1.078A4 4 0 0 1 18 8.093V5.75a2 2 0 0 0-2-2"/>`), g.Group(children),
	)
}

func BracesVariable(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 20.25a2 2 0 0 1-2-2v-2.343a4 4 0 0 0-1.172-2.829L2.75 12l1.078-1.078A4 4 0 0 0 5 8.093V5.75a2 2 0 0 1 2-2m10 16.5a2 2 0 0 0 2-2v-2.343a4 4 0 0 1 1.172-2.829L21.25 12l-1.078-1.078A4 4 0 0 1 19 8.093V5.75a2 2 0 0 0-2-2M9 8.143l6 7.714m0-7.714l-6 7.714"/>`), g.Group(children),
	)
}

func Branch(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 8.25a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5m0 0v7.5m0-7.5l.321.566A9.42 9.42 0 0 0 14.25 13.5v0M7 15.75a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5m7.25-2.25a2.75 2.75 0 1 0 5.5 0a2.75 2.75 0 0 0-5.5 0"/>`), g.Group(children),
	)
}

func BranchCompare(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.25 5.5a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0m13 13a2.75 2.75 0 1 1-5.5 0a2.75 2.75 0 0 1 5.5 0"/><path d="M5.5 8.25v7.25a3 3 0 0 0 3 3h3.75"/><path d="m10.75 16l1.793 1.793a1 1 0 0 1 0 1.414L10.75 21m7.75-5.25V8.5a3 3 0 0 0-3-3h-3.75"/><path d="m13.25 8l-1.793-1.793a1 1 0 0 1 0-1.414L13.25 3"/></g>`), g.Group(children),
	)
}

func BranchFork(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7 8.25a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5m0 0V12m0 3.75a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5m0 0V12m10-3.75a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5m0 0V9a3 3 0 0 1-3 3H7"/>`), g.Group(children),
	)
}

func BranchForkTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 8.25a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5m0 0V10a2 2 0 0 0 2 2h3m6-3.75a2.75 2.75 0 1 0 0-5.5a2.75 2.75 0 0 0 0 5.5m0 0V10a2 2 0 0 1-2 2h-4m0 0v3.75m0 0a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5"/>`), g.Group(children),
	)
}

func Briefcase(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.75 9.894a2.5 2.5 0 0 1 2.5-2.5h11.5a2.5 2.5 0 0 1 2.5 2.5V17.5a2.5 2.5 0 0 1-2.5 2.5H6.25a2.5 2.5 0 0 1-2.5-2.5z"/><path d="M17.75 7.394H6.25a2.5 2.5 0 0 0-2.5 2.5v.303a3 3 0 0 0 3 3h10.5a3 3 0 0 0 3-3v-.303a2.5 2.5 0 0 0-2.5-2.5M8.603 5.5a1.5 1.5 0 0 1 1.5-1.5h3.794a1.5 1.5 0 0 1 1.5 1.5v1.894H8.603z"/></g>`), g.Group(children),
	)
}

func BriefcaseTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 10.5A2.5 2.5 0 0 1 6.25 8h11.5a2.5 2.5 0 0 1 2.5 2.5v7a2.5 2.5 0 0 1-2.5 2.5H6.25a2.5 2.5 0 0 1-2.5-2.5zm4.853-5a1.5 1.5 0 0 1 1.5-1.5h3.794a1.5 1.5 0 0 1 1.5 1.5V8H8.603z"/>`), g.Group(children),
	)
}

func Brightness(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="4.25"/><path stroke-linecap="round" d="M12 2.75v1.5M5.46 5.46l1.06 1.06M12 19.75v1.5m5.48-3.77l1.06 1.06M2.75 12h1.5m1.21 6.54l1.06-1.06M19.75 12h1.5m-3.77-5.48l1.06-1.06"/></g>`), g.Group(children),
	)
}

func Broom(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13.444 17.626l.707-.707a5 5 0 0 0 0-7.071m-.707 7.778l-7.071-7.071m7.07 7.07l-2.828 4.243l-8.485-8.485l4.243-2.828m0 0l.707-.707a5 5 0 0 1 7.07 0m0 0l6.718-6.718"/>`), g.Group(children),
	)
}

func Bug(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.5 9.08a3.23 3.23 0 0 1 3.23-3.23h2.54a3.23 3.23 0 0 1 3.23 3.23v6.27a4.5 4.5 0 0 1-4.5 4.5v0a4.5 4.5 0 0 1-4.5-4.5zm9 3.77h4.75m-18.5 0H7.5m2.25-9.7v.45A2.25 2.25 0 0 0 12 5.85v0a2.25 2.25 0 0 0 2.25-2.25v-.45M16.5 16.6h1.253a2.5 2.5 0 0 1 2.5 2.5v1.75M7.5 16.6H6.247a2.5 2.5 0 0 0-2.5 2.5v1.75M16.5 9.1h1.253a2.5 2.5 0 0 0 2.5-2.5V4.85M7.5 9.1H6.247a2.5 2.5 0 0 1-2.5-2.5V4.85"/>`), g.Group(children),
	)
}

func BuildingMultiple(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.75 5a1.5 1.5 0 0 0-1.5-1.5H9.288a1 1 0 0 0-1 1v1H4.75a2 2 0 0 0-2 2v11a2 2 0 0 0 2 2h10zm0 3.5h4.5a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2h-4.5zm0 4h2.5m-2.5 4h2.5"/><circle cx="6.75" cy="9.5" r="1" fill="currentColor"/><circle cx="6.75" cy="13" r="1" fill="currentColor"/><circle cx="6.75" cy="16.5" r="1" fill="currentColor"/><circle cx="10.75" cy="9.5" r="1" fill="currentColor"/><circle cx="10.75" cy="13" r="1" fill="currentColor"/><circle cx="10.75" cy="16.5" r="1" fill="currentColor"/></g>`), g.Group(children),
	)
}

func Button(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="18.5" height="11" x="2.75" y="6.5" rx="4"/><path d="M7 12h10"/></g>`), g.Group(children),
	)
}

func Cake(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4.25 13.25a2.5 2.5 0 0 1 2.5-2.5h10.5a2.5 2.5 0 0 1 2.5 2.5v8H4.25z"/><path d="m4.25 14.87l2.249 1.45a3 3 0 0 0 3.252 0l.623-.4a3 3 0 0 1 3.252 0l.623.4a3 3 0 0 0 3.252 0l2.249-1.45m1.5 6.375H2.75M12 2.75l1.414 1.414a2 2 0 1 1-2.828 0zm0 4.83v3.17"/></g>`), g.Group(children),
	)
}

func Calculator(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="14.5" height="18.5" x="4.75" y="2.75" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="3"/><rect width="7.5" height="3.75" x="8.25" y="6.25" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="1"/><circle cx="8.5" cy="13.5" r="1" fill="currentColor"/><circle cx="12" cy="13.5" r="1" fill="currentColor"/><circle cx="15.5" cy="13.5" r="1" fill="currentColor"/><circle cx="8.5" cy="17.5" r="1" fill="currentColor"/><circle cx="12" cy="17.5" r="1" fill="currentColor"/><circle cx="15.5" cy="17.5" r="1" fill="currentColor"/></g>`), g.Group(children),
	)
}

func Calendar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.25 4.75H6.75a3.5 3.5 0 0 0-3.5 3.5v9.5a3.5 3.5 0 0 0 3.5 3.5h10.5a3.5 3.5 0 0 0 3.5-3.5v-9.5a3.5 3.5 0 0 0-3.5-3.5m-14 4.5h17.5M7.361 4.75v-2m9.25 2v-2"/>`), g.Group(children),
	)
}

func CalligraphyPen(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.16 2.75v1.049c0 1.187 0 1.78-.231 2.233a2.12 2.12 0 0 1-.926.927c-.454.23-1.047.23-2.234.23h-.423M4.84 2.75v1.049c0 1.187 0 1.78.231 2.233c.203.4.528.724.926.927c.454.23 1.047.23 2.234.23h.423m0 0L6.319 11.47c-.367.673-.55 1.01-.61 1.362a2.1 2.1 0 0 0 .054.935c.1.344.32.657.762 1.283l2.704 3.83c.934 1.323 1.4 1.983 1.98 2.217a2.12 2.12 0 0 0 1.582 0c.58-.234 1.047-.894 1.98-2.216l2.704-3.83c.442-.627.662-.94.761-1.284a2.1 2.1 0 0 0 .055-.935c-.06-.352-.243-.689-.61-1.362l-2.335-4.28m-6.692 0h6.692M12 13.325v7.522"/><circle cx="12" cy="13.325" r="1.673" fill="currentColor"/></g>`), g.Group(children),
	)
}

func Camera(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12.857 3.189h-1.714c-.681 0-1.022 0-1.331.094c-.274.083-.529.22-.75.401c-.25.205-.438.489-.816 1.056L7.103 6.454c-1.524 0-2.286 0-2.868.296a2.72 2.72 0 0 0-1.188 1.19c-.297.581-.297 1.343-.297 2.867v5.651c0 1.524 0 2.286.297 2.868c.26.512.677.928 1.188 1.189c.582.296 1.344.296 2.868.296h9.794c1.524 0 2.286 0 2.868-.296a2.72 2.72 0 0 0 1.188-1.19c.297-.581.297-1.343.297-2.867v-5.651c0-1.524 0-2.286-.297-2.868a2.72 2.72 0 0 0-1.188-1.189c-.582-.296-1.344-.296-2.868-.296L15.754 4.74c-.378-.567-.567-.85-.816-1.056a2.2 2.2 0 0 0-.75-.401c-.309-.094-.65-.094-1.331-.094"/><path d="M15.775 13.212a3.775 3.775 0 1 1-7.55 0a3.775 3.775 0 0 1 7.55 0"/></g>`), g.Group(children),
	)
}

func Cancel(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5 19l7-7m0 0l7-7m-7 7L5 5m7 7l7 7"/>`), g.Group(children),
	)
}

func CancelCircle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="9.25"/><path stroke-linecap="round" d="m8.875 8.875l6.25 6.25m0-6.25l-6.25 6.25"/></g>`), g.Group(children),
	)
}

func CancelSquare(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M8.25 3.75a4.5 4.5 0 0 0-4.5 4.5v7.5a4.5 4.5 0 0 0 4.5 4.5h7.5a4.5 4.5 0 0 0 4.5-4.5v-7.5a4.5 4.5 0 0 0-4.5-4.5z"/><path stroke-linecap="round" d="m8.655 8.655l6.69 6.69m0-6.69l-6.69 6.69"/></g>`), g.Group(children),
	)
}

func Cart(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="9.549" cy="19.049" r="1.701"/><circle cx="16.96" cy="19.049" r="1.701"/><path d="m5.606 5.555l2.01 6.364c.309.978.463 1.467.76 1.829c.26.32.599.567.982.72c.435.173.947.173 1.973.173h3.855c1.026 0 1.538 0 1.972-.173c.384-.153.722-.4.983-.72c.296-.362.45-.851.76-1.829l.409-1.296l.24-.766l.331-1.05a2.5 2.5 0 0 0-2.384-3.252zm0 0l-.011-.037a7 7 0 0 0-.14-.42a2.92 2.92 0 0 0-2.512-1.84C2.84 3.25 2.727 3.25 2.5 3.25"/></g>`), g.Group(children),
	)
}

func Cent(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.085 3.875v2.417m0 14.083v-2.417m4.382-1.705a5.84 5.84 0 0 1-4.382 1.705m4.382-9.961a5.84 5.84 0 0 0-4.382-1.705m0 11.666A5.836 5.836 0 0 1 6.5 12.125a5.84 5.84 0 0 1 5.585-5.833"/>`), g.Group(children),
	)
}

func CenterHorizontal(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.25 20.25V3.75m-16.5 16.5V3.75"/><rect width="6" height="13" rx="2" transform="matrix(-1 0 0 1 15 5.5)"/></g>`), g.Group(children),
	)
}

func CenterVertical(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.75 20.25h16.5M3.75 3.75h16.5"/><rect width="6" height="13" rx="2" transform="matrix(0 -1 -1 0 18.5 15)"/></g>`), g.Group(children),
	)
}

func Chat(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21.25a9.25 9.25 0 1 0-8.307-5.177c.108.22.144.468.089.706l-.816 3.536a.6.6 0 0 0 .72.72l3.535-.817a1.06 1.06 0 0 1 .706.09A9.2 9.2 0 0 0 12 21.25M7.97 9.886h8.06m-8.06 4.228h5.748"/>`), g.Group(children),
	)
}

func CheckboxIntermediateTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="18.5" height="18.5" x="2.75" y="2.75" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="4"/><path fill="currentColor" d="M7.113 6.25a.86.86 0 0 0-.863.862v9.775c0 .477.386.863.862.863h9.775a.863.863 0 0 0 .863-.863V7.114a.863.863 0 0 0-.863-.863z"/></g>`), g.Group(children),
	)
}

func CheckboxUnchecked(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="16.5" height="16.5" x="3.75" y="3.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="4"/>`), g.Group(children),
	)
}

func Checkmark(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.75 7.018l-9.257 9.257a1 1 0 0 1-1.414 0L4.25 11.446"/>`), g.Group(children),
	)
}

func CheckmarkChecked(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="16.5" height="16.5" x="3.75" y="3.75" rx="4"/><path d="m16.512 9.107l-5.787 5.786l-3.237-3.232"/></g>`), g.Group(children),
	)
}

func CheckmarkCircle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><circle cx="12" cy="12" r="9.25"/><path stroke-linejoin="round" d="m16.375 9.194l-5.611 5.612l-3.139-3.134"/></g>`), g.Group(children),
	)
}

func CheckmarkIntermediate(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="16.5" height="16.5" x="3.75" y="3.75" rx="4"/><path d="M16.19 12H7.81"/></g>`), g.Group(children),
	)
}

func CheckmarkStarburst(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.072 4.069a2.17 2.17 0 0 1 2.804-1.162l1.315.529c.52.208 1.099.208 1.618 0l1.315-.529a2.17 2.17 0 0 1 2.804 1.162l.556 1.303c.22.515.63.925 1.144 1.144l1.303.556a2.17 2.17 0 0 1 1.162 2.804l-.529 1.315a2.17 2.17 0 0 0 0 1.618l.529 1.315a2.17 2.17 0 0 1-1.162 2.804l-1.303.556a2.17 2.17 0 0 0-1.144 1.144l-.556 1.303a2.17 2.17 0 0 1-2.804 1.162l-1.315-.529a2.17 2.17 0 0 0-1.618 0l-1.315.529a2.17 2.17 0 0 1-2.804-1.162l-.556-1.303a2.17 2.17 0 0 0-1.144-1.144l-1.303-.556a2.17 2.17 0 0 1-1.162-2.804l.529-1.315a2.17 2.17 0 0 0 0-1.618l-.529-1.315A2.17 2.17 0 0 1 4.07 7.072l1.303-.556a2.17 2.17 0 0 0 1.144-1.144z"/><path d="m15.899 9.5l-5 5l-2.797-2.793"/></g>`), g.Group(children),
	)
}

func ChevronDown(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.75 9.75l6.19 6.19a1.5 1.5 0 0 0 2.12 0l6.19-6.19"/>`), g.Group(children),
	)
}

func ChevronLeft(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m14.25 4.75l-6.19 6.19a1.5 1.5 0 0 0 0 2.12l6.19 6.19"/>`), g.Group(children),
	)
}

func ChevronRight(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.75 4.75l6.19 6.19a1.5 1.5 0 0 1 0 2.12l-6.19 6.19"/>`), g.Group(children),
	)
}

func ChevronUp(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4.75 14.25l6.19-6.19a1.5 1.5 0 0 1 2.12 0l6.19 6.19"/>`), g.Group(children),
	)
}

func ChromeRestore(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 9.75a4 4 0 0 1 4-4h7.5a4 4 0 0 1 4 4v7.5a4 4 0 0 1-4 4h-7.5a4 4 0 0 1-4-4z"/><path fill="currentColor" d="M15.25 2h-5.5a4.74 4.74 0 0 0-3.464 1.5h8.964c2.9 0 5.25 2.35 5.25 5.25v8.964A4.74 4.74 0 0 0 22 14.25v-5.5A6.75 6.75 0 0 0 15.25 2"/></g>`), g.Group(children),
	)
}

func ChromeRestoreVar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 9.75a4 4 0 0 1 4-4h7.5a4 4 0 0 1 4 4v7.5a4 4 0 0 1-4 4h-7.5a4 4 0 0 1-4-4z"/><path d="M6.75 2.75h7.5a7 7 0 0 1 7 7v7.5"/></g>`), g.Group(children),
	)
}

func Circle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<circle cx="12" cy="12" r="9.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/>`), g.Group(children),
	)
}

func CircleSmall(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<circle cx="12" cy="12" r="3.25" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/>`), g.Group(children),
	)
}

func Clipboard(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="16.5" height="18.5" x="3.75" y="2.75" rx="3.5"/><path d="M8.25 2.75h7.5v2.5a2 2 0 0 1-2 2h-3.5a2 2 0 0 1-2-2z"/></g>`), g.Group(children),
	)
}

func ClipboardPaste(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M9.75 21.25h-3a3.5 3.5 0 0 1-3.5-3.5V6.25a3.5 3.5 0 0 1 3.5-3.5h9.5a3.5 3.5 0 0 1 3.5 3.5v2"/><path d="M7.75 2.75h7.5v2.5a2 2 0 0 1-2 2h-3.5a2 2 0 0 1-2-2z"/><path stroke-linecap="round" stroke-linejoin="round" d="M12.25 13.25a2.5 2.5 0 0 1 2.5-2.5h3.5a2.5 2.5 0 0 1 2.5 2.5v5.5a2.5 2.5 0 0 1-2.5 2.5h-3.5a2.5 2.5 0 0 1-2.5-2.5z"/></g>`), g.Group(children),
	)
}

func ClipboardSearch(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path stroke-linecap="round" d="M14.25 21.25h3a3.5 3.5 0 0 0 3.5-3.5V6.25a3.5 3.5 0 0 0-3.5-3.5h-9.5a3.5 3.5 0 0 0-3.5 3.5v4"/><path d="M8.75 2.75h7.5v2.5a2 2 0 0 1-2 2h-3.5a2 2 0 0 1-2-2z"/><path stroke-linecap="round" stroke-linejoin="round" d="M9.17 19.67a4.054 4.054 0 1 0-5.733-5.733A4.054 4.054 0 0 0 9.17 19.67m0 0l2.58 2.58"/></g>`), g.Group(children),
	)
}

func Clock(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 12a9.25 9.25 0 1 1-18.5 0a9.25 9.25 0 0 1 18.5 0"/><path d="M11.25 6.75v6h4"/></g>`), g.Group(children),
	)
}

func ClosedCaptions(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="18.5" height="14.5" x="2.75" y="4.75" rx="4"/><path d="M10.5 14.382a2.75 2.75 0 1 1 0-4.764m7.125 4.764a2.75 2.75 0 1 1 0-4.764"/></g>`), g.Group(children),
	)
}

func Cloud(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.81 10.045a.35.35 0 0 1-.341-.356v0a4.27 4.27 0 0 0-4.27-4.267A4.27 4.27 0 0 0 6.93 9.69v0c0 .193-.15.358-.341.377a4.268 4.268 0 0 0 .43 8.512h9.962a4.268 4.268 0 1 0 0-8.533z"/>`), g.Group(children),
	)
}

func Code(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.75 6.5L3.25 12l5.5 5.5m6.5-11l5.5 5.5l-5.5 5.5"/>`), g.Group(children),
	)
}

func CoffeeHot(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.5 10.682c0-.875.71-1.585 1.585-1.585H16.52c.876 0 1.585.71 1.585 1.585v3.265a7.303 7.303 0 0 1-7.302 7.303v0A7.303 7.303 0 0 1 3.5 13.947z"/><path d="M18.105 10.556h1.464A2.43 2.43 0 0 1 22 12.986v0a2.43 2.43 0 0 1-2.43 2.43h-1.465M6.421 3.75v2.43m4.382-2.43v2.43m4.381-2.43v2.43"/></g>`), g.Group(children),
	)
}

func ColorPalette(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="1.25" cy="1.25" r="1.25" fill="currentColor" transform="matrix(-1 0 0 1 16.654 6.034)"/><circle cx="1.25" cy="1.25" r="1.25" fill="currentColor" transform="matrix(-1 0 0 1 12.156 5.221)"/><circle cx="1.25" cy="1.25" r="1.25" fill="currentColor" transform="matrix(-1 0 0 1 8.654 7.94)"/><circle cx="1.25" cy="1.25" r="1.25" fill="currentColor" transform="matrix(-1 0 0 1 7.685 12.156)"/><circle cx="1.25" cy="1.25" r="1.25" fill="currentColor" transform="matrix(-1 0 0 1 9.904 15.948)"/><path stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M21.25 12A9.25 9.25 0 1 0 12 21.25c1.318 0 2.224-1.28 2.329-2.594l.117-1.473a3 3 0 0 1 2.758-2.752l1.651-.129c1.28-.1 2.395-1.019 2.395-2.302Z"/></g>`), g.Group(children),
	)
}

func Comment(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.183 16.958h.75a.75.75 0 0 0-.75-.75zm.839 4.16l.508.552zm4.52-4.16v-.75a.75.75 0 0 0-.508.199zM3.839 6.75A3.25 3.25 0 0 1 7.09 3.5V2a4.75 4.75 0 0 0-4.75 4.75zm0 6.208V6.75h-1.5v6.208zm3.25 3.25a3.25 3.25 0 0 1-3.25-3.25h-1.5a4.75 4.75 0 0 0 4.75 4.75zm.072 0H7.09v1.5h.072zm.022 0h-.022v1.5h.022zm.75 4.542v-3.792h-1.5v3.792zm-.419-.184a.25.25 0 0 1 .42.184h-1.5c0 1.09 1.295 1.657 2.096.92zm4.52-4.16l-4.52 4.16L8.53 21.67l4.52-4.16zm4.877-.198h-4.37v1.5h4.37zm3.25-3.25a3.25 3.25 0 0 1-3.25 3.25v1.5a4.75 4.75 0 0 0 4.75-4.75zm0-6.208v6.208h1.5V6.75zM16.91 3.5a3.25 3.25 0 0 1 3.25 3.25h1.5A4.75 4.75 0 0 0 16.91 2zm-9.822 0h9.822V2H7.089z"/>`), g.Group(children),
	)
}

func CommentMultiple(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M6.435 17.735h.75a.75.75 0 0 0-.75-.75zm.839 3.384l.508.552zm3.677-3.384v-.75a.75.75 0 0 0-.508.198zM3.735 8.76a2.25 2.25 0 0 1 2.25-2.25v-1.5a3.75 3.75 0 0 0-3.75 3.75zm0 5.975V8.76h-1.5v5.975zm2.25 2.25a2.25 2.25 0 0 1-2.25-2.25h-1.5a3.75 3.75 0 0 0 3.75 3.75zm.431 0h-.431v1.5h.431zm.02 0h-.02v1.5h.02zm.75 3.766v-3.016h-1.5v3.016zm-.42-.184a.25.25 0 0 1 .42.184h-1.5c0 1.089 1.294 1.657 2.096.92zm3.677-3.384l-3.677 3.384l1.016 1.104l3.677-3.384zm4.561-.198h-4.053v1.5h4.053zm2.25-2.25a2.25 2.25 0 0 1-2.25 2.25v1.5a3.75 3.75 0 0 0 3.75-3.75zm0-5.975v5.975h1.5V8.76zm-2.25-2.25a2.25 2.25 0 0 1 2.25 2.25h1.5a3.75 3.75 0 0 0-3.75-3.75zm-9.02 0h9.02v-1.5h-9.02z"/><path d="M13.015 3.499c1.413 0 2.427 0 3.223.065c.788.065 1.295.188 1.707.398a4.25 4.25 0 0 1 1.857 1.857c.21.412.333.92.398 1.707c.065.796.065 1.81.065 3.223v3.974a1.5 1.5 0 0 0 1.5-1.5v-2.508c0-1.372 0-2.447-.07-3.311c-.072-.88-.221-1.608-.556-2.266a5.75 5.75 0 0 0-2.513-2.513c-.658-.335-1.386-.484-2.266-.556C15.496 2 14.42 2 13.05 2H7.496a1.5 1.5 0 0 0-1.5 1.5z"/></g>`), g.Group(children),
	)
}

func CommentMultipleVar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M6.43 17.742h.75a.75.75 0 0 0-.75-.75zm.84 3.384l-.509-.55zm3.67-3.385v-.75a.75.75 0 0 0-.508.2zM3.736 8.76a2.25 2.25 0 0 1 2.25-2.25v-1.5a3.75 3.75 0 0 0-3.75 3.75zm0 5.982V8.76h-1.5v5.982zm2.25 2.25a2.25 2.25 0 0 1-2.25-2.25h-1.5a3.75 3.75 0 0 0 3.75 3.75zm.427 0h-.427v1.5h.427zm.019 0h-.02v1.5h.02zm.75 3.767V17.74h-1.5v3.018zm-.42-.184a.25.25 0 0 1 .42.184h-1.5c0 1.09 1.296 1.657 2.097.919zm3.671-3.385l-3.67 3.385l1.016 1.103l3.671-3.385zm4.553-.199H10.94v1.5h4.044zm2.25-2.25a2.25 2.25 0 0 1-2.25 2.25v1.5a3.75 3.75 0 0 0 3.75-3.75zm0-5.981v5.982h1.5V8.76zm-2.25-2.25a2.25 2.25 0 0 1 2.25 2.25h1.5a3.75 3.75 0 0 0-3.75-3.75zm-9 0h9v-1.5h-9z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.985 2.76h8a6 6 0 0 1 6 6v4.982"/></g>`), g.Group(children),
	)
}

func CompareSize(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.75 8.75a2.5 2.5 0 0 1 2.5 2.5v6.5a2.5 2.5 0 0 1-2.5 2.5H5.25a2.5 2.5 0 0 1-2.5-2.5v-6.5a2.5 2.5 0 0 1 2.5-2.5zm-2.5 7.5v2.5m0-8v2.5m-8-9.5h2.5m3 0h.5a2 2 0 0 1 2 2v.5m-11-2.5h-.5a2 2 0 0 0-2 2v.5"/>`), g.Group(children),
	)
}

func Compass(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9.25"/><path d="M10.195 10.195q.221-.22.475-.404c.382-.275.835-.456 1.74-.818l2.357-.943c.632-.252.947-.379 1.148-.313c.174.058.31.194.368.368c.066.2-.06.517-.313 1.148l-.943 2.357c-.362.905-.543 1.358-.818 1.74q-.183.255-.404.475m-3.61-3.61a4 4 0 0 0-.404.475c-.275.382-.456.835-.818 1.74l-.943 2.357c-.252.632-.379.947-.313 1.148c.058.174.194.31.368.368c.2.066.516-.06 1.148-.313l2.357-.943c.905-.362 1.358-.543 1.74-.818q.255-.183.475-.404m-3.61-3.61l3.61 3.61"/></g>`), g.Group(children),
	)
}

func Component(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="5.671" height="5.671" x="7.99" y="17.86" rx="1.5" transform="rotate(-45 7.99 17.86)"/><rect width="5.671" height="5.671" x="13.851" y="12" rx="1.5" transform="rotate(-45 13.85 12)"/><rect width="5.671" height="5.671" x="2.13" y="12" rx="1.5" transform="rotate(-45 2.13 12)"/><rect width="5.671" height="5.671" x="7.99" y="6.14" rx="1.5" transform="rotate(-45 7.99 6.14)"/></g>`), g.Group(children),
	)
}

func Compose(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M10.371 4.25H8.25a5 5 0 0 0-5 5v6.5a5 5 0 0 0 5 5h6.5a5 5 0 0 0 5-5v-2.121"/><path d="M12.299 14.75a1.86 1.86 0 0 0 1.316-.545l6.59-6.59a1.86 1.86 0 0 0 0-2.633l-1.187-1.187a1.86 1.86 0 0 0-2.633 0l-6.59 6.59a1.86 1.86 0 0 0-.545 1.316v3.049z"/></g>`), g.Group(children),
	)
}

func Computer(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="18.5" height="15.031" x="2.75" y="2.75" rx="3.5"/><path d="M9.11 17.781v3.469m5.78-3.469v3.469m-8.382 0h10.984"/></g>`), g.Group(children),
	)
}

func ComputerMac(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 6.25a3.5 3.5 0 0 1 3.5-3.5h11.5a3.5 3.5 0 0 1 3.5 3.5v8.031a3.5 3.5 0 0 1-3.5 3.5H6.25a3.5 3.5 0 0 1-3.5-3.5zm0 7.75h18.5M9.11 17.781v1.469a2 2 0 0 1-2 2h-.6m8.38-3.469v1.469a2 2 0 0 0 2 2h.6m-10.982 0h10.984"/>`), g.Group(children),
	)
}

func ContractDown(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 20.25h4.25a4 4 0 0 0 4-4v-2.5M12 20.25H7.75a4 4 0 0 1-4-4V12M12 20.25V15a3 3 0 0 0-3-3H3.75m0 0V7.75a4 4 0 0 1 4-4h2.5m10 6.5h-5.5a1 1 0 0 1-.707-.293M13.75 3.75v5.5c0 .276.112.526.293.707M20.25 3.75l-5.5 5.5l-.707.707"/>`), g.Group(children),
	)
}

func Cookies(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.076 10.535a2.75 2.75 0 0 0 3.89 0v0c.127-.128.362-.076.393.102a9.25 9.25 0 0 1-15.65 8.154a9.25 9.25 0 0 1 8.154-15.65c.178.031.23.266.102.394v0a2.75 2.75 0 0 0 2.333 4.667a2.75 2.75 0 0 0 .778 2.333"/><circle cx="8.5" cy="15.5" r="1.25" fill="currentColor"/><circle cx="7.5" cy="9.5" r="1.25" fill="currentColor"/><circle cx="12.5" cy="12.5" r="1.25" fill="currentColor"/><circle cx="15.5" cy="16.5" r="1.25" fill="currentColor"/></g>`), g.Group(children),
	)
}

func Copy(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.869 5.75a3 3 0 0 0-3-3H7.21a3 3 0 0 0-3 3v9.579a3 3 0 0 0 3 3h6.658a3 3 0 0 0 3-3z"/><path fill="currentColor" fill-rule="evenodd" d="M8.414 21.763a3.25 3.25 0 0 1-1.52-1.263h3.961c1.377 0 2.374 0 3.162-.053c.781-.052 1.304-.153 1.74-.33a5.25 5.25 0 0 0 2.9-2.9c.177-.435.278-.959.33-1.74c.053-.788.053-1.785.053-3.162V5.432c.486.311.882.746 1.146 1.264c.199.39.28.809.317 1.272c.037.448.037.998.037 1.672v2.702c0 1.345 0 2.39-.056 3.235c-.058.856-.175 1.557-.436 2.203a6.75 6.75 0 0 1-3.728 3.728c-.646.261-1.347.379-2.203.436c-.844.056-1.89.056-3.235.056h-.053c-.548 0-.996 0-1.362-.024c-.379-.026-.723-.08-1.053-.213" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func CopyVar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.869 5.75a3 3 0 0 0-3-3H7.211a3 3 0 0 0-3 3v9.579a3 3 0 0 0 3 3h6.658a3 3 0 0 0 3-3z"/><path d="M19.79 6.67v8.579a6 6 0 0 1-6 6H8.132"/></g>`), g.Group(children),
	)
}

func CornerRadius(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.25 4.75h-7.5a7 7 0 0 0-7 7v7.5"/>`), g.Group(children),
	)
}

func CreditCard(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="18.5" height="14" x="2.75" y="5" stroke-linecap="round" stroke-linejoin="round" rx="3"/><path d="M2.75 9.5h18.5"/><path stroke-linecap="round" d="M14.75 14.25h3"/></g>`), g.Group(children),
	)
}

func Crop(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.25 17.653H9.347a3 3 0 0 1-3-3V2.75M2.75 6.347h3.597m11.306 11.306v3.597M8.917 6.347h5.736a3 3 0 0 1 3 3v5.736"/>`), g.Group(children),
	)
}

func Css(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m20.754 4.792l-2.84 14.31a1.5 1.5 0 0 1-1.099 1.161l-4.442 1.141a1.5 1.5 0 0 1-.746 0l-4.442-1.14a1.5 1.5 0 0 1-1.098-1.162L3.246 4.792A1.5 1.5 0 0 1 4.717 3h14.566a1.5 1.5 0 0 1 1.471 1.792"/><path d="M7.77 6.881h9.53l-8.63 5.577h6.225a1 1 0 0 1 .978 1.205l-.562 2.683a1 1 0 0 1-.653.74l-2.283.786a1 1 0 0 1-.66-.003L8.99 16.9"/></g>`), g.Group(children),
	)
}

func Cursor(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M8.084 20.276c-1.06 1.38-3.264.66-3.306-1.079L4.443 5.392C4.407 3.932 6 3.012 7.247 3.773l11.788 7.192c1.485.906 1.006 3.176-.719 3.403l-5.581.738a1.84 1.84 0 0 0-1.221.705z"/>`), g.Group(children),
	)
}

func CursorClick(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M12.394 20.734c-.757.985-2.332.471-2.362-.771l-.239-9.86a1.317 1.317 0 0 1 2.003-1.157l8.42 5.137c1.06.647.718 2.268-.513 2.431l-3.987.527c-.346.046-.66.227-.872.503z"/><path stroke-linecap="round" d="M3.797 8.75h2.5m3.75-3.502v-2.5M6.815 5.765L5.047 3.998m8.232 1.767l1.768-1.767"/></g>`), g.Group(children),
	)
}

func CursorDrag(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.315 11.055v4.176m3.246-6.593V6.555c0-.892-.727-1.615-1.623-1.615s-1.623.723-1.623 1.615v3.747m3.246-.566V4.365a1.62 1.62 0 0 1 1.623-1.615c.897 0 1.623.723 1.623 1.615V9.73m0 .005v-3.18a1.624 1.624 0 0 1 3.246 0v2.19m0 0v.99m0-.99a1.62 1.62 0 0 1 1.624-1.615A1.62 1.62 0 0 1 20.3 8.745v5.935a6.57 6.57 0 0 1-6.57 6.57h-2.95a7.08 7.08 0 0 1-7.069-7.492l.015-.256a3.06 3.06 0 0 1 2.13-2.738l1.458-.462m0 0V11.4"/>`), g.Group(children),
	)
}

func Cut(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.651 14.43a3.75 3.75 0 1 0-4.302 6.143a3.75 3.75 0 0 0 4.302-6.144m0 0l3.35-4.446m5.45-7.235l-3.82 5.069m1.718 6.611a3.75 3.75 0 1 1 4.302 6.144a3.75 3.75 0 0 1-4.302-6.144m0 0L12 9.984M6.55 2.749L12 9.984"/>`), g.Group(children),
	)
}

func DarkTheme(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2.75 12A9.25 9.25 0 0 0 12 21.25V2.75A9.25 9.25 0 0 0 2.75 12"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21.25a9.25 9.25 0 0 0 0-18.5m0 18.5a9.25 9.25 0 0 1 0-18.5m0 18.5V2.75"/></g>`), g.Group(children),
	)
}

func Database(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.25 6.25c0 1.933-3.246 3.5-7.25 3.5s-7.25-1.567-7.25-3.5m14.5 0c0-1.933-3.246-3.5-7.25-3.5s-7.25 1.567-7.25 3.5m14.5 0V12M4.75 6.25V12m0 0v5.75c0 1.933 3.246 3.5 7.25 3.5s7.25-1.567 7.25-3.5V12m-14.5 0c0 1.933 3.246 3.5 7.25 3.5s7.25-1.567 7.25-3.5"/>`), g.Group(children),
	)
}

func DatabaseAdd(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.25 6.25c0 1.933-3.246 3.5-7.25 3.5s-7.25-1.567-7.25-3.5m14.5 0c0-1.933-3.246-3.5-7.25-3.5s-7.25 1.567-7.25 3.5m14.5 0v3.53M4.75 6.25V12m0 0v5.75c0 1.756 2.678 3.21 6.17 3.461M4.75 12c0 1.577 2.16 2.91 5.13 3.348"/><path fill="currentColor" fill-rule="evenodd" d="M17.5 23a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-8.993a.5.5 0 0 1 .5.5V17h2.493a.5.5 0 1 1 0 1H18v2.493a.5.5 0 1 1-1 0V18h-2.493a.5.5 0 1 1 0-1H17v-2.493a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func Delete(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.687 6.213L6.8 18.976a2.5 2.5 0 0 0 2.466 2.092h3.348m6.698-14.855L17.2 18.976a2.5 2.5 0 0 1-2.466 2.092h-3.348m-1.364-9.952v5.049m3.956-5.049v5.049M2.75 6.213h18.5m-6.473 0v-1.78a1.5 1.5 0 0 0-1.5-1.5h-2.554a1.5 1.5 0 0 0-1.5 1.5v1.78z"/>`), g.Group(children),
	)
}

func Diamond(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.69 9.915h8.62m-8.62 0H2.75m4.94 0l3.65 10.051M7.69 9.915l2.4-3.933l.55-1.012m5.67 4.945h4.94m-4.94 0l-3.65 10.051m3.65-10.051l-2.4-3.933l-.566-1.012M2.75 9.915c0 .42.079.84.236 1.236c.212.535.642 1.028 1.5 2.013l3.467 3.976c1.397 1.602 2.095 2.404 2.923 2.698q.229.081.464.128M2.75 9.915c0-.42.079-.84.236-1.237c.212-.535.642-1.027 1.5-2.013c.467-.534.7-.801.97-1.008a3.36 3.36 0 0 1 1.361-.619c.334-.068.688-.068 1.397-.068h2.427M21.25 9.915c0 .42-.079.84-.236 1.236c-.212.535-.642 1.028-1.5 2.013l-3.467 3.976c-1.397 1.602-2.095 2.404-2.923 2.698a3.4 3.4 0 0 1-.464.128m8.59-10.051c0-.42-.079-.84-.236-1.237c-.212-.535-.642-1.027-1.5-2.013c-.467-.534-.7-.801-.97-1.008a3.36 3.36 0 0 0-1.361-.619c-.334-.068-.688-.068-1.397-.068h-2.442m-.684 14.996a3.4 3.4 0 0 1-1.32 0M13.344 4.97H10.64"/>`), g.Group(children),
	)
}

func Directions(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.507 14.007a2.84 2.84 0 0 1 0-4.014l6.486-6.486a2.84 2.84 0 0 1 4.014 0l6.486 6.486a2.84 2.84 0 0 1 0 4.014l-6.486 6.486a2.84 2.84 0 0 1-4.014 0z"/><path d="m14.46 9.02l1.394 1.395a1 1 0 0 1 .293.707m-1.688 2.102l1.395-1.395a1 1 0 0 0 .293-.707m-7.46 4.031v-2.53a1.5 1.5 0 0 1 1.5-1.5h5.96"/></g>`), g.Group(children),
	)
}

func DoNotDisturb(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="9.25"/><path stroke-linecap="round" d="M7.5 12h9"/></g>`), g.Group(children),
	)
}

func Document(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="15" height="18.5" x="4.5" y="2.75" rx="3.5"/><path d="M8.5 6.755h7m-7 4h7m-7 4H12"/></g>`), g.Group(children),
	)
}

func Dollar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M16.441 8.26S15.884 5.621 12 5.334m-4.376 10.52S8.75 18.518 12 18.684M12 2.75v2.584m0 15.916v-2.567m0-13.35a10 10 0 0 0-.704-.024c-1.688 0-3.881 1.405-3.881 3.367c0 1.963 1.589 2.732 4.388 3.21s4.782 1.531 4.782 3.696s-2.32 3.11-4.266 3.11a6 6 0 0 1-.319-.009"/>`), g.Group(children),
	)
}

func DrawText(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.719 16.25l1.92-4.404m0 0h6.91m-6.91 0l2.94-6.747a.553.553 0 0 1 1.029 0l2.941 6.747m0 0l.337.774"/><path fill="currentColor" d="M15.68 20.936a2.5 2.5 0 0 0 1.218-.673l5.455-5.45a2.526 2.526 0 1 0-3.57-3.573l-5.453 5.452c-.335.336-.57.76-.675 1.222l-.535 2.354a1.007 1.007 0 0 0 1.206 1.206z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.75 19.687l.568.234c.638.263 1.364.175 1.956-.18c.69-.411 1.649-.915 2.483-1.1c.583-.13 1.243.199 1.091.776c-.17.642-.69 1.396-.192 1.745c.75.525 5.031-.818 5.031-.818"/></g>`), g.Group(children),
	)
}

func Drop(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M12 21c-1.804 0-3.246-.566-4.397-1.446C2.38 15.557 5.832 8.09 10.801 3.522a1.767 1.767 0 0 1 2.398 0c4.97 4.568 8.42 12.035 3.198 16.032C15.246 20.434 13.804 21 12 21Z"/>`), g.Group(children),
	)
}

func Emoji(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9.25" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/><circle cx="9" cy="9.5" r="1.25" fill="currentColor"/><circle cx="15" cy="9.5" r="1.25" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.464 14.25a4 4 0 0 1-6.928 0"/></g>`), g.Group(children),
	)
}

func EmojiFrown(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9.25" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/><circle cx="9" cy="9.5" r="1.25" fill="currentColor"/><circle cx="15" cy="9.5" r="1.25" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.464 15.813a4 4 0 0 0-6.928 0"/></g>`), g.Group(children),
	)
}

func EmojiGrin(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9.25" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/><circle cx="9" cy="9" r="1.25" fill="currentColor"/><circle cx="15" cy="9" r="1.25" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.5 12.75c.276 0 .503.224.475.5a4.999 4.999 0 0 1-9.594 1.413a5 5 0 0 1-.356-1.414c-.028-.275.199-.499.475-.499z"/></g>`), g.Group(children),
	)
}

func EmojiLaughter(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="12" r="9.25"/><path d="M16.5 12.75c.276 0 .503.224.475.5a4.999 4.999 0 0 1-9.594 1.413a5 5 0 0 1-.356-1.414c-.028-.275.199-.499.475-.499zM7.264 9.082a1.797 1.797 0 0 1 3.472 0m2.528 0a1.796 1.796 0 0 1 3.472 0"/></g>`), g.Group(children),
	)
}

func Eraser(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.788 20.5h9.02m-9.02 0a3.47 3.47 0 0 0 2.486-1.02l1.29-1.29M9.788 20.5a3.47 3.47 0 0 1-2.438-1.02l-3.33-3.33a3.48 3.48 0 0 1 0-4.923l1.29-1.29m0 0l5.417-5.417a3.48 3.48 0 0 1 4.923 0l3.33 3.33a3.48 3.48 0 0 1 0 4.924l-5.417 5.416M5.31 9.936l.367.368l7.585 7.585l.301.301"/>`), g.Group(children),
	)
}

func Expand(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 20.25h4.25a4 4 0 0 0 4-4v-2.5M12 20.25H7.75a4 4 0 0 1-4-4V12M12 20.25V15a3 3 0 0 0-3-3H3.75m0 0V7.75a4 4 0 0 1 4-4h2.5m3.5 0h5.5c.276 0 .526.112.707.293m.293 6.207v-5.5a1 1 0 0 0-.293-.707M13.75 10.25l5.5-5.5l.707-.707"/>`), g.Group(children),
	)
}

func Extension(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 9a.75.75 0 0 0 0 1.5zm4 10a.75.75 0 0 0 1.5 0zm8.5-10a.75.75 0 0 0 0 1.5zM16 5a.75.75 0 0 0-1.5 0zm-.267.738l-.146.736zm-.472-.471l-.735.146zm3.977 9.466l-.735.146zm-.738-.483V15zm.267.011l.146-.735zM6.738 9.268l.736.146zm-.471.471l-.146-.735zm8.994 8.995l.736.146zm.472-.471l-.146-.736zm-9.466-4l-.146.735zm.471.471l-.735.146zm12.029-4.995l.146.736zm.471-.471l-.735-.146zm-8.971 8.995l-.146.735zm.471.47l.736-.145zm0-13.465l.736.146zm-.471.471l.146.736zM11.5 5A1.5 1.5 0 0 1 13 3.5V2a3 3 0 0 0-3 3zM9.75 6.5H10V5h-.25zM7.5 8.75A2.25 2.25 0 0 1 9.75 6.5V5A3.75 3.75 0 0 0 6 8.75zm0 .25v-.25H6V9zm-3 3A1.5 1.5 0 0 1 6 10.5V9a3 3 0 0 0-3 3zM6 13.5A1.5 1.5 0 0 1 4.5 12H3a3 3 0 0 0 3 3zm1.5 1.75V15H6v.25zm2.25 2.25a2.25 2.25 0 0 1-2.25-2.25H6A3.75 3.75 0 0 0 9.75 19zm.25 0h-.25V19H10zm3 3a1.5 1.5 0 0 1-1.5-1.5H10a3 3 0 0 0 3 3zm1.5-1.5a1.5 1.5 0 0 1-1.5 1.5V22a3 3 0 0 0 3-3zm1.75-1.5H16V19h.25zm2.25-2.25a2.25 2.25 0 0 1-2.25 2.25V19A3.75 3.75 0 0 0 20 15.25zm0-.25v.25H20V15zm0-1.5A1.5 1.5 0 0 1 17 12h-1.5a3 3 0 0 0 3 3zM17 12a1.5 1.5 0 0 1 1.5-1.5V9a3 3 0 0 0-3 3zm1.5-3.25V9H20v-.25zM16.25 6.5a2.25 2.25 0 0 1 2.25 2.25H20A3.75 3.75 0 0 0 16.25 5zm-.25 0h.25V5H16zm-3-3A1.5 1.5 0 0 1 14.5 5H16a3 3 0 0 0-3-3zM16 5h-.087l-.051-.001c-.024 0-.01-.002.017.004l-.292 1.471c.148.03.31.026.413.026zm-1.5 0c0 .103-.004.265.026.413l1.471-.292c.006.027.005.04.004.017v-.05L16 5zm1.38.003a.15.15 0 0 1 .117.118l-1.471.292a1.35 1.35 0 0 0 1.06 1.061zM20 15c0-.104.003-.265-.026-.413l-1.471.292c-.006-.027-.005-.04-.004-.017V15zm-1.5 0l.138.001c.024 0 .01.002-.017-.004l.292-1.471c-.148-.03-.31-.026-.413-.026zm1.474-.413a1.35 1.35 0 0 0-1.06-1.061l-.293 1.471a.15.15 0 0 1-.118-.118zM6 9l-.001.138c0 .024-.002.01.004-.017l1.471.292c.03-.148.026-.31.026-.413zm0 1.5c.103 0 .265.004.413-.026l-.292-1.471c.027-.006.04-.005.017-.004A2 2 0 0 1 6 9zm.003-1.38a.15.15 0 0 1 .118-.117l.292 1.471a1.35 1.35 0 0 0 1.061-1.06zM16 19v-.087l.001-.051c0-.024.002-.01-.004.017l-1.471-.292c-.03.148-.026.31-.026.413zm0-1.5c-.104 0-.265-.004-.413.026l.292 1.471c-.027.006-.04.005-.017.004h.05L16 19zm-.003 1.38a.15.15 0 0 1-.118.117l-.292-1.471a1.35 1.35 0 0 0-1.061 1.06zM6 15l.138.001c.024 0 .01.002-.017-.004l.292-1.471c-.148-.03-.31-.026-.413-.026zm1.5 0c0-.104.004-.265-.026-.413l-1.471.292c-.006-.027-.005-.04-.004-.017V15zm-1.38-.003a.15.15 0 0 1-.117-.118l1.471-.292a1.35 1.35 0 0 0-1.06-1.061zM18.5 10.5c.103 0 .265.004.413-.026l-.292-1.471c.027-.006.04-.005.017-.004A2 2 0 0 1 18.5 9zm0-1.5v.087l-.001.051c0 .024-.002.01.004-.017l1.471.292c.03-.148.026-.31.026-.413zm.413 1.474a1.35 1.35 0 0 0 1.061-1.06l-1.471-.293a.15.15 0 0 1 .118-.118zM10 19l.138.001c.024 0 .01.002-.017-.004l.292-1.471c-.148-.03-.31-.026-.413-.026zm1.5 0c0-.103.004-.265-.026-.413l-1.471.292c-.006-.027-.005-.04-.004-.017A2 2 0 0 1 10 19zm-1.38-.003a.15.15 0 0 1-.117-.118l1.471-.292a1.35 1.35 0 0 0-1.06-1.061zM10 5l-.001.138c0 .024-.002.01.004-.017l1.471.292c.03-.148.026-.31.026-.413zm0 1.5c.103 0 .265.004.413-.026l-.292-1.471c.027-.006.04-.005.017-.004h-.05L10 5zm.003-1.38a.15.15 0 0 1 .118-.117l.292 1.471a1.35 1.35 0 0 0 1.061-1.06z"/>`), g.Group(children),
	)
}

func Eye(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M3.182 12.808C4.233 14.613 7.195 18.81 12 18.81c4.813 0 7.77-4.199 8.82-6.002a1.6 1.6 0 0 0-.001-1.615C19.769 9.389 16.809 5.19 12 5.19s-7.768 4.197-8.818 6.001a1.6 1.6 0 0 0 0 1.617Z"/><path d="M12 14.625a2.625 2.625 0 1 0 0-5.25a2.625 2.625 0 0 0 0 5.25Z"/></g>`), g.Group(children),
	)
}

func EyeOff(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.282 21.782l4.278-4.278M21.782 3.282L17.673 7.39m-3.363 3.363a2.64 2.64 0 0 0-1.063-1.063a2.625 2.625 0 1 0-2.494 4.62m3.557-3.557l-3.557 3.557m3.557-3.557l3.363-3.363m-6.92 6.92L7.56 17.504M17.673 7.39c-.38-.319-.791-.621-1.232-.894C15.2 5.726 13.717 5.19 12 5.19c-4.956 0-7.948 4.459-8.91 6.16c-.11.196-.165.293-.197.446a1.2 1.2 0 0 0 0 .408c.032.152.088.25.198.445c.51.903 1.593 2.582 3.237 3.96c.38.319.791.621 1.232.895m12.18-7.925c.528.694.919 1.328 1.17 1.773c.11.194.165.292.197.444c.023.112.023.296 0 .408c-.032.152-.087.25-.197.444c-.96 1.702-3.95 6.162-8.91 6.162q-.714-.002-1.374-.117"/>`), g.Group(children),
	)
}

func Eyedropper(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.047 7.819L20.92 9.69m-1.872-1.87l-2.864-2.864m2.864 2.864l1.496-1.496a1 1 0 0 0 0-1.414l-1.45-1.45a1 1 0 0 0-1.414 0l-1.496 1.496m-1.872-1.872l1.872 1.872m-3.796 2.751a1.5 1.5 0 0 1 2.121 0l1.787 1.786a1.5 1.5 0 0 1 0 2.12l-8.562 8.563a1.5 1.5 0 0 1-.829.421l-2.12.332a1.5 1.5 0 0 1-1.714-1.715l.334-2.118a1.5 1.5 0 0 1 .42-.827z"/>`), g.Group(children),
	)
}

func EyedropperColor(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.047 7.819L20.92 9.69m-1.872-1.87l-2.864-2.864m2.864 2.864l1.496-1.496a1 1 0 0 0 0-1.414l-1.45-1.45a1 1 0 0 0-1.414 0l-1.496 1.496m-1.872-1.872l1.872 1.872M6.235 13.86l-2.202 2.202c-.205.205-.308.307-.387.425q-.107.157-.171.335c-.048.134-.07.277-.116.563l-.096.608c-.173 1.1-.26 1.65-.087 2.05c.151.35.43.629.78.78c.399.174.949.087 2.05-.085l.61-.096c.286-.044.43-.067.563-.115q.18-.065.337-.17c.117-.08.22-.183.425-.389l7.718-7.718c.594-.594.891-.89 1.002-1.233a1.5 1.5 0 0 0 0-.927c-.11-.343-.408-.64-1.002-1.234l-.513-.513c-.594-.594-.891-.891-1.234-1.002a1.5 1.5 0 0 0-.927 0c-.342.11-.64.408-1.233 1.002zm0 0h7.815"/>`), g.Group(children),
	)
}

func EyedropperColorAccent(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7.941 19.968l6.109-6.108H6.235l-2.202 2.202c-.205.205-.308.307-.387.425a1.5 1.5 0 0 0-.171.335c-.048.134-.07.277-.116.563l-.096.608c-.173 1.1-.26 1.65-.087 2.05c.151.35.43.629.78.78c.399.174.949.087 2.05-.085l.61-.096c.286-.044.43-.067.563-.115q.18-.065.337-.17c.117-.08.22-.183.425-.389"/>`), g.Group(children),
	)
}

func Facebook(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M10.478 21.125a9.3 9.3 0 0 0 3.037.002m-3.038-.002A9.25 9.25 0 0 1 2.75 12a9.25 9.25 0 1 1 10.765 9.127m-3.038-.002V16.12H8.58a.6.6 0 0 1-.6-.6v-1.838a.6.6 0 0 1 .6-.6h1.897V9.95a3 3 0 0 1 3-3h1.81a1 1 0 0 1 1 1v1.04a1 1 0 0 1-1 1h-.772a1 1 0 0 0-1 1v2.092h2.297a.6.6 0 0 1 .592.698l-.25 1.504a1 1 0 0 1-.986.836h-1.653v5.007"/>`), g.Group(children),
	)
}

func FastFoward(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M5.996 18.323c-1.02.784-2.496.057-2.496-1.229V6.906c0-1.286 1.476-2.013 2.496-1.229l6.224 5.192a1.473 1.473 0 0 1 0 2.262z"/><path d="M15.246 18.323c-1.02.784-2.496.057-2.496-1.229V6.906c0-1.286 1.476-2.013 2.496-1.229l6.224 5.192a1.473 1.473 0 0 1 0 2.262z"/></g>`), g.Group(children),
	)
}

func Figma(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M5.833 5.833A3.083 3.083 0 0 1 8.917 2.75H12v6.167H8.917a3.083 3.083 0 0 1-3.084-3.084m12.334 0a3.083 3.083 0 0 0-3.084-3.083H12v6.167h3.083a3.083 3.083 0 0 0 3.084-3.084"/><rect width="6.167" height="6.167" rx="3.083" transform="matrix(-1 0 0 1 18.167 8.917)"/><path d="M5.833 12a3.083 3.083 0 0 1 3.084-3.083H12v6.166H8.917A3.083 3.083 0 0 1 5.833 12"/><path d="M5.833 18.167a3.083 3.083 0 0 1 3.084-3.084H12v3.084a3.083 3.083 0 0 1-3.083 3.083v0a3.083 3.083 0 0 1-3.084-3.083"/></g>`), g.Group(children),
	)
}

func File(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.688 3.063a3.5 3.5 0 0 1 1.027.712l5.968 5.97c.3.3.54.647.711 1.026m-7.706-7.708a3.5 3.5 0 0 0-1.448-.313H7.792a3.5 3.5 0 0 0-3.5 3.5v11.5a3.5 3.5 0 0 0 3.5 3.5h8.416a3.5 3.5 0 0 0 3.5-3.5v-5.53c0-.505-.109-.999-.314-1.45m-7.706-7.707V8.77a2 2 0 0 0 2 2h5.706"/>`), g.Group(children),
	)
}

func FileAdd(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M6.5 23a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-8.993a.5.5 0 0 1 .5.5V17h2.493a.5.5 0 1 1 0 1H7v2.493a.5.5 0 1 1-1 0V18H3.507a.5.5 0 0 1 0-1H6v-2.493a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.292 10.25v-4a3.5 3.5 0 0 1 3.5-3.5h2.448a3.5 3.5 0 0 1 1.447.313M13.75 21.25h2.458a3.5 3.5 0 0 0 3.5-3.5v-5.53c0-.505-.109-.999-.314-1.45m-7.706-7.707a3.5 3.5 0 0 1 1.027.712l5.968 5.97c.3.3.54.647.711 1.026m-7.706-7.708V8.77a2 2 0 0 0 2 2h5.706"/></g>`), g.Group(children),
	)
}

func FileSync(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.292 10.25v-4a3.5 3.5 0 0 1 3.5-3.5h2.448a3.5 3.5 0 0 1 1.447.313M13.75 21.25h2.458a3.5 3.5 0 0 0 3.5-3.5v-5.53c0-.505-.109-.999-.314-1.45m-7.706-7.707a3.5 3.5 0 0 1 1.027.712l5.968 5.97c.3.3.54.647.711 1.026m-7.706-7.708V8.77a2 2 0 0 0 2 2h5.706"/><path fill="currentColor" fill-rule="evenodd" d="M6.5 23a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m1.548-8.64a3.5 3.5 0 0 0-4.929 2.234a.5.5 0 0 0 .966.259A2.5 2.5 0 0 1 8.3 15.765h-.565a.5.5 0 0 0 0 1H9.5a.5.5 0 0 0 .5-.5V14.5a.5.5 0 0 0-1 0v.55a3.5 3.5 0 0 0-.952-.69m1.833 4.046a.5.5 0 0 0-.966-.259A2.5 2.5 0 0 1 4.7 19.235h.565a.5.5 0 0 0 0-1H3.5a.5.5 0 0 0-.5.5V20.5a.5.5 0 0 0 1 0v-.55a3.5 3.5 0 0 0 5.88-1.544" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func Filter(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 7.25h15M7.385 12h9.23m-6.345 4.75h3.46"/>`), g.Group(children),
	)
}

func FilterCancel(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 7.25h5.75M7.385 12H12m-1.73 4.75h3.46"/><path fill="currentColor" fill-rule="evenodd" d="M17.5 12a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m-2.352-7.852a.5.5 0 0 1 .707 0L17.5 5.793l1.645-1.645a.5.5 0 1 1 .707.707L18.207 6.5l1.645 1.645a.5.5 0 0 1-.707.707L17.5 7.207l-1.645 1.645a.5.5 0 0 1-.707-.707L16.793 6.5l-1.645-1.645a.5.5 0 0 1 0-.707" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func FilterCancelTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m17.52 10.25l2.919-2.814a2.664 2.664 0 0 0-1.852-4.579H5.397a2.648 2.648 0 0 0-1.82 4.572l4.817 4.555c.277.262.415.393.515.548a1.5 1.5 0 0 1 .192.446c.044.178.044.369.044.75v3.976c0 1.268 0 1.902.204 2.318a2 2 0 0 0 .901.91"/><path fill="currentColor" fill-rule="evenodd" d="M17.5 23a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m-2.352-7.852a.5.5 0 0 1 .707 0l1.645 1.645l1.645-1.645a.5.5 0 1 1 .707.707L18.207 17.5l1.645 1.645a.5.5 0 0 1-.707.707L17.5 18.207l-1.645 1.645a.5.5 0 1 1-.707-.707l1.645-1.645l-1.645-1.645a.5.5 0 0 1 0-.707" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func FilterTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.145 17.704v-3.976c0-.381 0-.572-.044-.75a1.5 1.5 0 0 0-.192-.446c-.1-.155-.238-.286-.515-.548L3.578 7.43a2.648 2.648 0 0 1 1.82-4.572h13.189a2.664 2.664 0 0 1 1.852 4.579l-4.765 4.607c-.27.261-.405.392-.501.545a1.5 1.5 0 0 0-.187.441c-.044.176-.044.364-.044.74v3.935c0 .542 0 .813-.062 1.057a2 2 0 0 1-.641 1.027c-.192.163-.436.282-.923.52c-1.14.557-1.709.835-2.172.835a2 2 0 0 1-1.795-1.121c-.204-.416-.204-1.05-.204-2.318"/>`), g.Group(children),
	)
}

func Flag(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.25 21.25v-6m0 0v-10a1.5 1.5 0 0 1 1.5-1.5h11.086a1 1 0 0 1 .821 1.571l-2.311 3.322a1.5 1.5 0 0 0 0 1.714l2.311 3.322a1 1 0 0 1-.82 1.571z"/>`), g.Group(children),
	)
}

func Flashlight(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M6.512 7.71h10.976m-8.73 3.78h6.477v8.11a1.9 1.9 0 0 1-1.9 1.9h-2.677a1.9 1.9 0 0 1-1.9-1.9z"/><path d="m5.806 7.99l2.952 3.5h6.477l2.959-3.5a1 1 0 0 0 .236-.645V4.85A1.85 1.85 0 0 0 16.58 3H7.42a1.85 1.85 0 0 0-1.85 1.85v2.495a1 1 0 0 0 .236.645m6.191 7.039v1.766"/></g>`), g.Group(children),
	)
}

func FoldableHorizontal(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 5.25L5.223 3.601A2 2 0 0 0 2.75 5.545v12.91A2 2 0 0 0 5.223 20.4L12 18.75m0-13.5v13.5m0-13.5l6.777-1.649a2 2 0 0 1 2.473 1.944v12.91a2 2 0 0 1-2.473 1.944L12 18.75"/>`), g.Group(children),
	)
}

func FoldableHorizontalHalf(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m12 5.25l6.777-1.649a2 2 0 0 1 2.473 1.944v12.91a2 2 0 0 1-2.473 1.944L12 18.75z"/><path stroke-dasharray="2.5 3" d="m9.5 19.358l-4.895 1.19a1.5 1.5 0 0 1-1.855-1.457V4.91a1.5 1.5 0 0 1 1.855-1.46L9.5 4.641"/></g>`), g.Group(children),
	)
}

func FoldableVertical(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.375 12l-1.39 6.853a2 2 0 0 0 1.961 2.397h10.108a2 2 0 0 0 1.96-2.397L17.625 12m-11.25 0h11.25m-11.25 0l-1.39-6.853A2 2 0 0 1 6.947 2.75h10.108a2 2 0 0 1 1.96 2.397L17.625 12"/>`), g.Group(children),
	)
}

func FoldableVerticalHalf(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m17.25 12l1.648 6.777a2 2 0 0 1-1.943 2.473h-9.91A2 2 0 0 1 5.1 18.777L6.75 12z"/><path stroke-dasharray="2 3" d="m6.142 9.5l-1.19-4.895A1.5 1.5 0 0 1 6.408 2.75H17.59a1.5 1.5 0 0 1 1.458 1.855L17.859 9.5"/></g>`), g.Group(children),
	)
}

func Folder(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M2.75 8.623v7.379a4 4 0 0 0 4 4h10.5a4 4 0 0 0 4-4v-5.69a4 4 0 0 0-4-4H12M2.75 8.624V6.998a3 3 0 0 1 3-3h2.9a2.5 2.5 0 0 1 1.768.732L12 6.313m-9.25 2.31h5.904a2.5 2.5 0 0 0 1.768-.732L12 6.313"/>`), g.Group(children),
	)
}

func FolderAdd(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M17.5 23a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-8.993a.5.5 0 0 1 .5.5V17h2.493a.5.5 0 1 1 0 1H18v2.494a.5.5 0 0 1-1 0V18h-2.493a.5.5 0 1 1 0-1H17v-2.493a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 8.623v6.579c0 1.68 0 2.52.327 3.162a3 3 0 0 0 1.311 1.31c.642.328 1.482.328 3.162.328h2.7M2.75 8.623V7.598c0-1.26 0-1.89.245-2.371a2.25 2.25 0 0 1 .984-.984c.48-.245 1.11-.245 2.371-.245h1.679c.611 0 .917 0 1.205.07a2.5 2.5 0 0 1 .722.299c.253.154.469.37.901.803L12 6.313m-9.25 2.31h5.283c.611 0 .917 0 1.205-.069a2.5 2.5 0 0 0 .722-.3c.252-.154.469-.37.901-.802L12 6.312m0 0h4.449c1.68 0 2.52 0 3.163.326a3 3 0 0 1 1.31 1.311c.267.521.316 1.174.326 2.301"/></g>`), g.Group(children),
	)
}

func FolderOpen(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="m3.882 18.043l4.041-5.623a4 4 0 0 1 3.249-1.665h8.752M3.882 18.043a3.65 3.65 0 0 0 2.777 1.277h8.343a4 4 0 0 0 3.405-1.9l2.918-4.734a1.287 1.287 0 0 0-1.115-1.931h-.286M3.882 18.043A3.65 3.65 0 0 1 3 15.661V7.424A2.744 2.744 0 0 1 5.744 4.68h2.653c.607 0 1.189.24 1.618.67l.911.91a1.83 1.83 0 0 0 1.294.537l4.044-.001a3.66 3.66 0 0 1 3.66 3.66v.299"/>`), g.Group(children),
	)
}

func FullScreenMaximize(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 8.345V6.25a2.5 2.5 0 0 1 2.5-2.5h2.095M3.75 15.655v2.095a2.5 2.5 0 0 0 2.5 2.5h2.095M20.25 8.345V6.25a2.5 2.5 0 0 0-2.5-2.5h-2.095m4.595 11.905v2.095a2.5 2.5 0 0 1-2.5 2.5h-2.095"/>`), g.Group(children),
	)
}

func FullScreenMinimize(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.345 3.75v2.095a2.5 2.5 0 0 1-2.5 2.5H3.75M8.345 20.25v-2.095a2.5 2.5 0 0 0-2.5-2.5H3.75M15.655 3.75v2.095a2.5 2.5 0 0 0 2.5 2.5h2.095M15.655 20.25v-2.095a2.5 2.5 0 0 1 2.5-2.5h2.095"/>`), g.Group(children),
	)
}

func Game(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.031 3.887H7.97a5.22 5.22 0 0 0-5.219 5.22v8.265c0 2.075 2.533 3.085 3.962 1.581l2.976-3.134h4.624l2.875 3.46c1.374 1.654 4.063.682 4.063-1.467V9.106a5.22 5.22 0 0 0-5.219-5.219M8.138 8.39v4m-2-2h4"/><circle cx="14.662" cy="9.39" r="1" fill="currentColor"/><circle cx="16.862" cy="11.59" r="1" fill="currentColor"/></g>`), g.Group(children),
	)
}

func Gift(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.5 12.5H12V21H7a2.5 2.5 0 0 1-2.5-2.5zm-1.75-3A1.5 1.5 0 0 1 4.25 8H12v4.5H4.25a1.5 1.5 0 0 1-1.5-1.5zm9.25 3h7.5v6A2.5 2.5 0 0 1 17 21h-5zM12 8h7.75a1.5 1.5 0 0 1 1.5 1.5V11a1.5 1.5 0 0 1-1.5 1.5H12zM7 5.5A2.5 2.5 0 0 1 9.5 3v0A2.5 2.5 0 0 1 12 5.5V8H9.5A2.5 2.5 0 0 1 7 5.5m10 0A2.5 2.5 0 0 0 14.5 3v0A2.5 2.5 0 0 0 12 5.5V8h2.5A2.5 2.5 0 0 0 17 5.5"/>`), g.Group(children),
	)
}

func GitCommit(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.75 12a3.75 3.75 0 1 1-7.5 0a3.75 3.75 0 0 1 7.5 0m0 0h5.5m-18.5 0h5.5"/>`), g.Group(children),
	)
}

func Github(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M9.096 21.25v-3.146a3.33 3.33 0 0 1 .758-2.115c-3.005-.4-5.28-1.859-5.28-5.798c0-1.666 1.432-3.89 1.432-3.89c-.514-1.13-.5-3.084.06-3.551c0 0 1.95.175 3.847 1.75c1.838-.495 3.764-.554 5.661 0c1.897-1.575 3.848-1.75 3.848-1.75c.558.467.573 2.422.06 3.551c0 0 1.432 2.224 1.432 3.89c0 3.94-2.276 5.398-5.28 5.798a3.33 3.33 0 0 1 .757 2.115v3.146"/><path d="M3.086 16.57c.163.554.463 1.066.878 1.496c.414.431.932.77 1.513.988a4.46 4.46 0 0 0 3.62-.216"/></g>`), g.Group(children),
	)
}

func Globe(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.25 12A9.25 9.25 0 0 0 12 2.75M21.25 12H2.75m18.5 0A9.25 9.25 0 0 1 12 21.25m0-18.5A9.25 9.25 0 0 0 2.75 12M12 2.75c-.5 0-4 4.141-4 9.25s3.5 9.25 4 9.25m0-18.5c.5 0 4 4.141 4 9.25s-3.5 9.25-4 9.25M2.75 12A9.25 9.25 0 0 0 12 21.25"/>`), g.Group(children),
	)
}

func Google(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M13.365 2.83a9.25 9.25 0 0 1 4.744 2.089c.338.284.336.794.024 1.106l-1.616 1.616c-.312.312-.816.306-1.171.044a5.365 5.365 0 1 0 1.615 6.705h-3.91a.8.8 0 0 1-.8-.8V11.3a.8.8 0 0 1 .8-.8h7.493c.316 0 .61.186.681.495c.313 1.362-.125 3.246-.158 3.384l-.004.016c-.528 1.963-1.661 3.706-3.274 4.944a9.25 9.25 0 1 1-4.424-16.51"/>`), g.Group(children),
	)
}

func GoogleChrome(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8.25a3.75 3.75 0 0 0-3.747 3.904M12 8.25a3.75 3.75 0 0 1 3.608 4.775M12 8.25h8.458m-4.85 4.775a3.752 3.752 0 0 1-7.355-.871m7.355.871l-3.08 8.21m7.93-12.985A9.252 9.252 0 0 0 4.6 6.45m15.858 1.8q.085.19.161.386a9.25 9.25 0 0 1-8.09 12.599m0 0A9.25 9.25 0 0 1 2.75 12c0-2.083.688-4.004 1.85-5.55m3.653 5.704L4.6 6.45"/>`), g.Group(children),
	)
}

func GooglePlay(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="m16.28 16.311l3.29-1.861A2.79 2.79 0 0 0 21 12c0-.955-.476-1.91-1.43-2.45l-3.29-1.86m0 8.622L8.209 20.88a2.8 2.8 0 0 1-2.779 0a2.9 2.9 0 0 1-.7-.557m11.552-4.012L4.729 3.677m0 16.646A2.78 2.78 0 0 1 4 18.43V5.57a2.78 2.78 0 0 1 1.061-2.202a2.81 2.81 0 0 1 3.147-.248l8.073 4.569M4.729 20.323L16.281 7.69"/>`), g.Group(children),
	)
}

func GoogleTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.201 9.947a5.365 5.365 0 0 1 8.145-2.262c.355.262.858.268 1.17-.044l1.617-1.616c.312-.312.314-.822-.024-1.106A9.25 9.25 0 0 0 3.612 8.46M7.2 9.947a5.365 5.365 0 0 0 1.69 6.31M7.2 9.947L3.611 8.46m5.28 7.796a5.365 5.365 0 0 0 6.533 0m-6.532 0l-2.365 3.082m8.897-3.082a5.4 5.4 0 0 0 1.537-1.866h-3.91a.8.8 0 0 1-.8-.8V11.3a.8.8 0 0 1 .8-.8h7.493c.316 0 .61.186.681.495c.313 1.362-.125 3.246-.158 3.384l-.004.016c-.528 1.963-1.661 3.706-3.274 4.944m-2.365-3.083l2.365 3.082m0 0a9.25 9.25 0 0 1-11.262 0m0 0A9.25 9.25 0 0 1 3.612 8.46"/>`), g.Group(children),
	)
}

func Graph(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M3.5 4v12.5a4 4 0 0 0 4 4H20"/><path d="m7 14l3.293-3.293a1 1 0 0 1 1.414 0l1.336 1.336a1 1 0 0 0 1.414 0L19 7.5l.648-.649M15 6.5h3.8c.331 0 .631.134.848.351M20 11.5V7.7c0-.331-.134-.631-.352-.849"/></g>`), g.Group(children),
	)
}

func Grid(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="6.5" height="6.5" x="3.75" y="13.75" rx="2"/><rect width="6.5" height="6.5" x="13.75" y="13.75" rx="2"/><rect width="6.5" height="6.5" x="3.75" y="3.75" rx="2"/><rect width="6.5" height="6.5" x="13.75" y="3.75" rx="2"/></g>`), g.Group(children),
	)
}

func GridDots(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><circle cx="5" cy="5" r="1.5"/><circle cx="12" cy="5" r="1.5"/><circle cx="19" cy="5" r="1.5"/><circle cx="5" cy="12" r="1.5"/><circle cx="12" cy="12" r="1.5"/><circle cx="19" cy="12" r="1.5"/><circle cx="5" cy="19" r="1.5"/><circle cx="12" cy="19" r="1.5"/><circle cx="19" cy="19" r="1.5"/></g>`), g.Group(children),
	)
}

func Hamburger(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M12 3.724c-4.302 0-7.79 3.051-7.79 6.816h15.58c0-3.765-3.488-6.816-7.79-6.816"/><rect width="18.5" height="5.355" x="2.75" y="10.54" rx="2"/><path d="M4.21 15.895h15.58l-.278 1.249a4 4 0 0 1-3.905 3.132H8.393a4 4 0 0 1-3.905-3.132zm10.926-2.833l-2.162-2.522h5.842l-2.162 2.522a1 1 0 0 1-1.518 0"/></g>`), g.Group(children),
	)
}

func Hand(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M4 10.059v3.424c0 1.853 0 2.78.221 3.536c.527 1.8 1.935 3.216 3.735 3.846c.6.176 1.196.363 2.344.532a5.8 5.8 0 0 0 2.014-.066c.303-.062.55-.115.758-.16c.49-.106.98-.233 1.43-.454c.508-.248.903-.506 1.475-.933c.342-.255.655-.566 1.28-1.188l3.247-3.23a1.68 1.68 0 0 0 0-2.384a1.7 1.7 0 0 0-2.396 0l-2.25 2.239v-5.162"/><path d="M12.893 7.852V5.95c0-.815.664-1.475 1.483-1.475c.818 0 1.482.66 1.482 1.475v4.424m-5.929-.319V3.95c0-.815.664-1.475 1.482-1.475c.819 0 1.482.66 1.482 1.475v6.109M6.964 7.32v2.739v-5.104a1.483 1.483 0 0 1 2.965 0v5.104M6.964 8.854V7.95c0-.815-.663-1.475-1.482-1.475C4.664 6.475 4 7.135 4 7.95v2.738"/></g>`), g.Group(children),
	)
}

func HardDrive(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M21.25 16.75v2.5a2 2 0 0 1-2 2H4.75a2 2 0 0 1-2-2v-2.5m18.5 0a2 2 0 0 0-2-2H4.75a2 2 0 0 0-2 2m18.5 0v-1.63a3 3 0 0 0-.09-.728l-2.342-9.37a3 3 0 0 0-2.91-2.272H8.092a3 3 0 0 0-2.91 2.272l-2.342 9.37a3 3 0 0 0-.09.727v1.631"/><circle cx="18" cy="18" r="1" fill="currentColor"/></g>`), g.Group(children),
	)
}

func Hash(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.778 8.395H21.25m-18.5 7.21h17.472M6.282 21.13L9.495 2.87m5.01 18.26l3.212-18.26"/>`), g.Group(children),
	)
}

func HatGraduation(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.652 14.714V9.78m3.18 2.07l4.049 2.667a4 4 0 0 0 4.402 0l4.049-2.668m-12.5 0L3.099 10.05a.99.99 0 0 1-.45-.815m3.183 2.616v5.061c0 .495.119.987.44 1.364c.747.877 2.514 2.39 5.81 2.39s5.063-1.513 5.81-2.39c.32-.377.44-.869.44-1.364V11.85m0 0l2.48-1.634a1.2 1.2 0 0 0 0-2.004l-6.53-4.302a4 4 0 0 0-4.401 0L3.099 8.379a.99.99 0 0 0-.45.855m0 0v.547"/>`), g.Group(children),
	)
}

func Headphones(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M21.25 17.368V12A9.25 9.25 0 0 0 12 2.75v0A9.25 9.25 0 0 0 2.75 12v5.368"/><path d="M2.75 13.321h4a1.5 1.5 0 0 1 1.5 1.5v4.429a2 2 0 0 1-2 2h-1.5a2 2 0 0 1-2-2zm13 1.5a1.5 1.5 0 0 1 1.5-1.5h4v5.929a2 2 0 0 1-2 2h-1.5a2 2 0 0 1-2-2z"/></g>`), g.Group(children),
	)
}

func Heart(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M4.087 6.477a4.565 4.565 0 0 1 6.456 0L12 7.934l1.457-1.457a4.565 4.565 0 0 1 6.456 6.457l-1.457 1.456l.013.013l-6.456 6.457l-.013-.013l-.013.013l-6.456-6.457l.013-.013l-1.457-1.456a4.565 4.565 0 0 1 0-6.457Z"/>`), g.Group(children),
	)
}

func HeartStylistic(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="m12 7.934l-1.457-1.457a4.565 4.565 0 1 0-6.456 6.457l1.457 1.456M12 7.934l1.457-1.457a4.565 4.565 0 0 1 6.456 6.457l-1.457 1.456l.013.013l-6.456 6.457l-.013-.013l-.013.013l-6.456-6.457l.013-.013M12 7.934L5.544 14.39"/>`), g.Group(children),
	)
}

func Hexagon(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.42 3.173a3.16 3.16 0 0 1 3.16 0l5.275 3.046a3.16 3.16 0 0 1 1.579 2.735v6.092a3.16 3.16 0 0 1-1.58 2.735l-5.275 3.046a3.16 3.16 0 0 1-3.158 0L5.145 17.78a3.16 3.16 0 0 1-1.579-2.735V8.954c0-1.128.602-2.17 1.58-2.735z"/>`), g.Group(children),
	)
}

func Highlighter(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 2.75v3.5a3.5 3.5 0 0 0 3.5 3.5h9.5a3.5 3.5 0 0 0 3.5-3.5v-3.5m-14.5 7h12.5v1.8a2.2 2.2 0 0 1-2.2 2.2h-8.1a2.2 2.2 0 0 1-2.2-2.2zm10.5 4h-8.5v5.663a1.3 1.3 0 0 0 1.733 1.226l5.433-1.918a2 2 0 0 0 1.334-1.886z"/>`), g.Group(children),
	)
}

func HighlighterAccent(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.25 6.25v-3.5H3.75v3.5a3.5 3.5 0 0 0 3.5 3.5h9.5a3.5 3.5 0 0 0 3.5-3.5m-4 7.5h-8.5v5.663a1.3 1.3 0 0 0 1.733 1.226l5.433-1.918a2 2 0 0 0 1.334-1.886z"/>`), g.Group(children),
	)
}

func History(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M11.25 7.75v5h3"/><path d="M4.855 7.875a8.25 8.25 0 1 1-.824 6.26m-.176-5.26v-4.75m0 4.75h4.75"/></g>`), g.Group(children),
	)
}

func Home(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M15.29 20.663h3.017a2.194 2.194 0 0 0 2.193-2.194v-6.454a3.3 3.3 0 0 0-1.13-2.48l-5.93-5.166a2.194 2.194 0 0 0-2.88 0L4.63 9.534a3.3 3.3 0 0 0-1.13 2.481v6.454c0 1.212.982 2.194 2.194 2.194h3.29m6.306 0v-6.581c0-.908-.736-1.645-1.645-1.645H10.63c-.909 0-1.645.737-1.645 1.645v6.581m6.306 0H8.984"/>`), g.Group(children),
	)
}

func HomeTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.71 18.65v-7.622a3 3 0 0 0-1.151-2.362l-6.326-4.951a2 2 0 0 0-2.466 0l-6.326 4.95a3 3 0 0 0-1.15 2.363v7.622c0 1.16.94 2.1 2.1 2.1h3.97v-7.965h5.278v7.965h3.97a2.1 2.1 0 0 0 2.1-2.1"/>`), g.Group(children),
	)
}

func Hourglass(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="m4.095 3.298l15.81-.048m-15.81 17.5l15.81-.048"/><path d="M18.426 3.31H5.574l.079 1.449a7.38 7.38 0 0 0 2.251 4.913l1.242 1.195a1.58 1.58 0 0 1 0 2.279L7.904 14.34a7.38 7.38 0 0 0-2.251 4.913l-.08 1.448h12.853l-.079-1.445a7.38 7.38 0 0 0-2.256-4.917l-1.242-1.194a1.58 1.58 0 0 1 0-2.28l1.242-1.193a7.38 7.38 0 0 0 2.256-4.918z"/></g>`), g.Group(children),
	)
}

func Html(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m20.754 4.792l-2.84 14.31a1.5 1.5 0 0 1-1.099 1.161l-4.442 1.141a1.5 1.5 0 0 1-.746 0l-4.442-1.14a1.5 1.5 0 0 1-1.098-1.162L3.246 4.792A1.5 1.5 0 0 1 4.717 3h14.566a1.5 1.5 0 0 1 1.471 1.792"/><path d="M17.3 6.881H8.944a1 1 0 0 0-.987 1.16l.577 3.576a1 1 0 0 0 .987.841h5.374a1 1 0 0 1 .978 1.205l-.562 2.683a1 1 0 0 1-.653.74l-2.283.786a1 1 0 0 1-.66-.003L8.99 16.9"/></g>`), g.Group(children),
	)
}

func Info(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9.25" stroke="currentColor" stroke-width="1.5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M12 11.813v5"/><circle cx="12" cy="8.438" r="1.25" fill="currentColor"/></g>`), g.Group(children),
	)
}

func Instagram(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="17" height="17" x="3.5" y="3.5" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="5.5"/><circle cx="12" cy="12" r="3.606" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/><circle cx="16.894" cy="7.106" r="1.03" fill="currentColor"/></g>`), g.Group(children),
	)
}

func Javascript(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="16.5" height="16.5" x="3.75" y="3.75" rx="3"/><path d="M11.5 11.25v4.5a1.5 1.5 0 0 1-1.5 1.5H9m8.25-6h-1.5a1.5 1.5 0 0 0-1.5 1.5v0a1.5 1.5 0 0 0 1.5 1.5v0a1.5 1.5 0 0 1 1.5 1.5v0a1.5 1.5 0 0 1-1.5 1.5h-1.5"/></g>`), g.Group(children),
	)
}

func Keyboard(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="18.5" height="13.5" x="2.75" y="5.25" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="3"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8 15.38h8"/><circle cx="7.5" cy="8.875" r="1" fill="currentColor"/><circle cx="10.5" cy="8.875" r="1" fill="currentColor"/><circle cx="13.5" cy="8.875" r="1" fill="currentColor"/><circle cx="16.5" cy="8.875" r="1" fill="currentColor"/><circle cx="7.5" cy="11.875" r="1" fill="currentColor"/><circle cx="10.5" cy="11.875" r="1" fill="currentColor"/><circle cx="13.5" cy="11.875" r="1" fill="currentColor"/><circle cx="16.5" cy="11.875" r="1" fill="currentColor"/></g>`), g.Group(children),
	)
}

func KeyboardShift(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.503 11.913l6.601-8.252a2.428 2.428 0 0 1 3.792 0l6.6 8.251c.83 1.037.092 2.573-1.235 2.573h-4.095v3.725c0 1.064 0 1.596-.207 2.003a1.9 1.9 0 0 1-.83.83c-.406.207-.938.207-2.002.207h-.254c-1.064 0-1.596 0-2.002-.207a1.9 1.9 0 0 1-.83-.83c-.207-.407-.207-.939-.207-2.003v-3.725H4.739c-1.327 0-2.065-1.536-1.236-2.572"/>`), g.Group(children),
	)
}

func Laptop(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M4.75 7a2 2 0 0 1 2-2h10.5a2 2 0 0 1 2 2v9H4.75zm-2 10a1 1 0 0 1 1-1h16.5a1 1 0 0 1 1 1v1a2 2 0 0 1-2 2H4.75a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func Layers(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M12.991 3.066a2 2 0 0 0-1.982 0L2.75 7.778l8.259 4.712a2 2 0 0 0 1.982 0l8.259-4.712z"/><path stroke-linecap="round" d="m2.75 12l7.268 4.147a4 4 0 0 0 3.964 0L21.25 12"/><path stroke-linecap="round" d="m2.75 16.222l7.268 4.147a4 4 0 0 0 3.964 0l7.268-4.147"/></g>`), g.Group(children),
	)
}

func Layout(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 8.75v8a3 3 0 0 0 3 3H10m-7.25-11v-1.5a3 3 0 0 1 3-3h12.5a3 3 0 0 1 3 3v1.5m-18.5 0H10m11.25 0v8a3 3 0 0 1-3 3H10m11.25-11H10m0 0v11"/>`), g.Group(children),
	)
}

func Leaf(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.508 16.223a6.38 6.38 0 0 1-9.016 0a6.367 6.367 0 0 1 0-9.009l3.094-3.091a2 2 0 0 1 2.828 0l3.094 3.091a6.367 6.367 0 0 1 0 9.01M12 12.265v9.025"/>`), g.Group(children),
	)
}

func LeafThree(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.526 7.499a4.75 4.75 0 0 1 4.75-4.75h2.217c.887 0 1.33 0 1.669.172c.298.152.54.394.692.692c.172.339.172.782.172 1.669v2.217c0 .433-.058.853-.167 1.252m-10.716 2.84l2.3 2.301m5.624 1.288l-6.071 6.071M12.012 9.81a6.03 6.03 0 0 1 3.423-1.059h2.424M12.012 9.81a4.85 4.85 0 0 0-4.163-2.361H5.586c-.905 0-1.358 0-1.704.176a1.6 1.6 0 0 0-.706.706C3 8.677 3 9.13 3 10.035v2.262a4.85 4.85 0 0 0 6.7 4.483m2.312-6.97a6.1 6.1 0 0 0-1.54 1.523a6 6 0 0 0-1.029 2.56M17.86 8.75h.404c1.132 0 1.698 0 2.13.22c.38.194.689.503.883.884c.22.432.22.997.22 2.129v2.828A6.06 6.06 0 0 1 9.7 16.78m-.258-2.888A6.1 6.1 0 0 0 9.7 16.78"/>`), g.Group(children),
	)
}

func LeafTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.35 8.1l1.93 1.93m5.277 3.663L8 20.25m4.872-13.162a6.5 6.5 0 0 1 2.082-.338h3.055c1.222 0 1.833 0 2.3.238c.41.209.744.543.953.953c.238.467.238 1.078.238 2.3v3.055a6.546 6.546 0 0 1-13.062.625m4.434-6.833a5.09 5.09 0 0 0-4.78-3.338H5.714c-.95 0-1.425 0-1.788.185a1.7 1.7 0 0 0-.742.742C3 5.04 3 5.515 3 6.465v2.376a5.09 5.09 0 0 0 5.438 5.08m4.434-6.833A6.57 6.57 0 0 0 9.28 10.03m-.842 3.89a6.5 6.5 0 0 1 .842-3.89"/>`), g.Group(children),
	)
}

func Library(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="3.998" height="15" x="2.75" y="4.504" rx="1.5"/><rect width="3.998" height="15" x="9.201" y="4.504" rx="1.5"/><path d="M15.267 8.378c-.165-.615.2-1.247.814-1.411l1.038-.278c.614-.165 1.245.2 1.41.814l2.681 10.014a1.15 1.15 0 0 1-.814 1.41l-1.038.279a1.15 1.15 0 0 1-1.41-.815z"/></g>`), g.Group(children),
	)
}

func Lightbulb(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.013 17.104c.126-.958.736-1.764 1.464-2.4a6.816 6.816 0 1 0-8.955 0c.729.636 1.34 1.442 1.465 2.4l.084.633l.233 1.774a2 2 0 0 0 1.983 1.739h1.426a2 2 0 0 0 1.983-1.739l.233-1.774zm-5.943.633h5.86"/>`), g.Group(children),
	)
}

func LineDiagonal(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="m21.25 2.75l-18.5 18.5"/>`), g.Group(children),
	)
}

func Link(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.621 7.5H7.25a4.5 4.5 0 0 0-4.5 4.5v0a4.5 4.5 0 0 0 4.5 4.5h2.371m4.758-9h2.371a4.5 4.5 0 0 1 4.5 4.5v0a4.5 4.5 0 0 1-4.5 4.5h-2.371M7.243 12h9.514"/>`), g.Group(children),
	)
}

func Location(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M6.4 4.76a7.92 7.92 0 0 1 11.2 11.2l-4.186 4.186a2 2 0 0 1-2.828 0L6.4 15.96a7.92 7.92 0 0 1 0-11.2Z"/><circle cx="12" cy="10.36" r="3" stroke-linecap="round"/></g>`), g.Group(children),
	)
}

func Lock(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="14.478" height="12.87" x="4.761" y="8.38" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="3"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.217 8.38V5.967A3.217 3.217 0 0 0 12 2.75v0a3.217 3.217 0 0 0-3.217 3.217V8.38"/><circle cx="12" cy="14.815" r="1.5" fill="currentColor"/></g>`), g.Group(children),
	)
}

func LockOpen(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="14.478" height="12.87" x="4.761" y="8.38" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="3"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.783 8.38V5.967a3.217 3.217 0 0 1 6.132-1.363"/><circle cx="12" cy="14.815" r="1.5" fill="currentColor"/></g>`), g.Group(children),
	)
}

func Mail(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="18.5" height="15.5" x="2.75" y="4.25" rx="3"/><path d="m2.75 8l8.415 3.866a2 2 0 0 0 1.67 0L21.25 8"/></g>`), g.Group(children),
	)
}

func MailOpen(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.59 8.949l-7.755 3.562a2 2 0 0 1-1.67 0L3.41 8.95m14.84 11.3H5.75a3 3 0 0 1-3-3v-7.215A2.5 2.5 0 0 1 3.93 7.91l7.014-4.36a2 2 0 0 1 2.112 0l7.014 4.36a2.5 2.5 0 0 1 1.18 2.124v7.215a3 3 0 0 1-3 3"/>`), g.Group(children),
	)
}

func Map(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M8.496 4.439L4.247 6.91a1 1 0 0 0-.497.864V18.26a1 1 0 0 0 1.503.865l3.243-1.887a1.5 1.5 0 0 1 1.508 0l3.992 2.322a1.5 1.5 0 0 0 1.508 0l4.249-2.472a1 1 0 0 0 .497-.864V5.739a1 1 0 0 0-1.503-.865l-3.243 1.887a1.5 1.5 0 0 1-1.508 0L10.004 4.44a1.5 1.5 0 0 0-1.508 0Zm.754.311v11.8m5.5-9.1v11.8"/>`), g.Group(children),
	)
}

func Mask(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="12" cy="12" r="9.25" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/><path fill="currentColor" d="M8.844 20.698a9.254 9.254 0 0 1 0-17.396a9.254 9.254 0 0 1 0 17.396"/></g>`), g.Group(children),
	)
}

func Megaphone(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="m7.142 15.9l-2.5-.627a2.5 2.5 0 0 1-1.892-2.425V10.78a2.5 2.5 0 0 1 1.891-2.424l13.5-3.39a2.5 2.5 0 0 1 3.109 2.425v8.847a2.5 2.5 0 0 1-3.109 2.425l-5.19-1.304m-5.81-1.458a3 3 0 1 0 5.809 1.459M7.143 15.9l5.809 1.46"/>`), g.Group(children),
	)
}

func MegaphoneLoud(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="m8.784 19.283l-1.251.75c-.632.378-.948.568-1.272.637c-.44.094-.897.037-1.3-.16c-.298-.147-.558-.408-1.08-.93c-.52-.52-.781-.781-.927-1.079a2 2 0 0 1-.16-1.3c.068-.325.258-.642.636-1.274l4.356-7.282c.718-1.2 1.077-1.8 1.57-2.061c.433-.23.935-.291 1.41-.175c.544.133 1.038.628 2.026 1.617l2.636 2.638c.988.989 1.482 1.483 1.616 2.026c.116.476.054.979-.175 1.412c-.261.494-.86.853-2.06 1.572l-1.27.761m-4.755 2.848a2.778 2.778 0 0 0 4.463.744a2.784 2.784 0 0 0 .291-3.592m-4.754 2.848l4.754-2.848"/><path stroke-linecap="round" stroke-linejoin="round" d="M20.007 4.404L17.542 6.87m3.709 3.007h-2.49m-4.224-6.719V5.65"/></g>`), g.Group(children),
	)
}

func Mention(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M16.625 20.01A9.25 9.25 0 1 1 21.25 12v1.5a2.5 2.5 0 0 1-2.5 2.5v0a2.5 2.5 0 0 1-2.5-2.5V12m0 0a4.25 4.25 0 1 1-8.5 0a4.25 4.25 0 0 1 8.5 0m0 0V7.75"/>`), g.Group(children),
	)
}

func Menu(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 4.75h18.5M2.75 12h18.5m-18.5 7.25h18.5"/>`), g.Group(children),
	)
}

func Microphone(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.801 6.449a3.199 3.199 0 1 1 6.398 0v4.95a3.199 3.199 0 0 1-6.398 0zM12 18.181a6.78 6.78 0 0 1-6.779-6.779M12 18.182a6.78 6.78 0 0 0 6.779-6.78M12 18.182v2.568"/>`), g.Group(children),
	)
}

func MicrophoneOff(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12 14.597a3.2 3.2 0 0 0 3.199-3.199V9.864l-4.47 4.47c.39.17.82.263 1.271.263"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 18.181a6.78 6.78 0 0 0 6.779-6.779M12 18.182v2.568m0-2.569a6.75 6.75 0 0 1-3.89-1.227m-2.89-5.552c0 1.1.262 2.14.727 3.058m-2.666 7.322l4.827-4.828M21.782 3.282l-6.583 6.582m0 0v1.534a3.2 3.2 0 0 1-4.47 2.937m4.47-4.47l-4.47 4.47m0 0l-2.62 2.62m.734-5.035a3 3 0 0 1-.042-.522v-4.95a3.199 3.199 0 0 1 6.391-.208"/></g>`), g.Group(children),
	)
}

func Microsoft(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 3.75H5.75a2 2 0 0 0-2 2V12M12 3.75h6.25a2 2 0 0 1 2 2V12M12 3.75v16.5m0 0h6.25a2 2 0 0 0 2-2V12M12 20.25H5.75a2 2 0 0 1-2-2V12m0 0h16.5"/>`), g.Group(children),
	)
}

func MicrosoftEdge(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.76 12a9.25 9.25 0 0 0 11.527 8.968M2.761 12a9.25 9.25 0 0 1 9.25-9.25c3.795 0 7.996 1.777 9.017 6.798c.223 1.1.09 2.258-.463 3.235c-.582 1.025-1.559 2.182-3.01 2.182c-.55.088-4.164.176-3.979-1.312A2.27 2.27 0 0 0 14.287 12M2.761 12s.444-4.849 5.78-4.833C14.117 7.184 14.288 12 14.288 12m0 0a2.277 2.277 0 0 0-2.277-2.277c-.642 0-1.636.694-1.636.694m3.913 10.55a9.26 9.26 0 0 0 5.406-3.814c.185-.275.058-.694-.26-.6a.5.5 0 0 0-.174.092c-1.364.617-4.22 1.12-6.685-.257c-1.5-.838-2.374-2.135-2.639-3.202c-.099-.255-.201-.896-.201-1.186c0-.615.243-1.173.64-1.583m3.913 10.55s-6.012.49-6.546-5.053c-.297-3.09 1.307-4.702 2.633-5.497"/>`), g.Group(children),
	)
}

func Moon(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M11.578 3.512a6.307 6.307 0 0 0 8.91 8.91a.45.45 0 0 1 .466-.095c.176.067.29.24.275.428A9.255 9.255 0 1 1 5.461 5.45a9.22 9.22 0 0 1 5.784-2.68a.42.42 0 0 1 .428.275c.06.16.02.34-.095.466Z"/>`), g.Group(children),
	)
}

func More(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><circle cx="6" cy="12" r="1.75"/><circle cx="12" cy="12" r="1.75"/><circle cx="18" cy="12" r="1.75"/></g>`), g.Group(children),
	)
}

func MoreVertical(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><circle cx="12" cy="18" r="1.75" transform="rotate(-90 12 18)"/><circle cx="12" cy="12" r="1.75" transform="rotate(-90 12 12)"/><circle cx="12" cy="6" r="1.75" transform="rotate(-90 12 6)"/></g>`), g.Group(children),
	)
}

func Motherboard(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="13.5" height="13.5" x="5.25" y="5.25" rx="3"/><path d="M18.75 8.75h2.5M18.75 12h2.5m-2.5 3.25h2.5m-6 3.5v2.5M12 18.75v2.5m-3.25-2.5v2.5m-6-12.5h2.5M2.75 12h2.5m-2.5 3.25h2.5m10-12.5v2.5M12 2.75v2.5m-3.25-2.5v2.5"/><rect width="5" height="5" x="9.5" y="9.5" rx="1"/></g>`), g.Group(children),
	)
}

func Movie(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 7.5a3 3 0 0 1 3-3h12.5a3 3 0 0 1 3 3v9a3 3 0 0 1-3 3H5.75a3 3 0 0 1-3-3zM7 5v14M17 5v14M2.75 9.5H7m-4.25 5H7m10-5h4.25m-4.25 5h4.25"/>`), g.Group(children),
	)
}

func Museum(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="18.5" height="3" x="2.75" y="18.376" rx="1"/><path d="M11.04 3.15L3.27 7.4a1 1 0 0 0-.52.877v.997a.6.6 0 0 0 .6.6h17.3a.6.6 0 0 0 .6-.6v-.997a1 1 0 0 0-.52-.877l-7.77-4.25a2 2 0 0 0-1.92 0M5.25 9.874v8.51m13.5-8.51v8.51m-4.25-8.51v8.51m-5-8.51v8.51"/></g>`), g.Group(children),
	)
}

func MusicNote(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M8.962 17.968V6.696a1.5 1.5 0 0 1 1.106-1.447l8.15-2.223a1.5 1.5 0 0 1 1.895 1.447v11.468M8.963 9.92l11.15-3.04M8.962 17.968a3.041 3.041 0 1 1-6.082 0a3.041 3.041 0 0 1 6.082 0"/><path d="M20.113 15.94a3.041 3.041 0 1 1-6.082 0a3.041 3.041 0 0 1 6.082 0"/></g>`), g.Group(children),
	)
}

func MusicNoteTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.5 17.25a4 4 0 1 1-8 0a4 4 0 0 1 8 0m0 0v-9m0 0l4.83 2.415a1.5 1.5 0 0 0 2.17-1.342V7.177a1.5 1.5 0 0 0-.83-1.342l-4.723-2.361a1 1 0 0 0-1.447.894z"/>`), g.Group(children),
	)
}

func Narrator(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7.183 16.958h.75a.75.75 0 0 0-.75-.75zm.839 4.16l.508.552zm4.52-4.16v-.75a.75.75 0 0 0-.508.199zM3.839 6.75A3.25 3.25 0 0 1 7.09 3.5V2a4.75 4.75 0 0 0-4.75 4.75zm0 6.208V6.75h-1.5v6.208zm3.25 3.25a3.25 3.25 0 0 1-3.25-3.25h-1.5a4.75 4.75 0 0 0 4.75 4.75zm.072 0H7.09v1.5h.072zm.022 0h-.022v1.5h.022zm.75 4.542v-3.792h-1.5v3.792zm-.419-.184a.25.25 0 0 1 .42.184h-1.5c0 1.09 1.295 1.657 2.096.92zm4.52-4.16l-4.52 4.16L8.53 21.67l4.52-4.16zm4.877-.198h-4.37v1.5h4.37zm3.25-3.25a3.25 3.25 0 0 1-3.25 3.25v1.5a4.75 4.75 0 0 0 4.75-4.75zm0-6.208v6.208h1.5V6.75zM16.91 3.5a3.25 3.25 0 0 1 3.25 3.25h1.5A4.75 4.75 0 0 0 16.91 2zm-9.822 0h9.822V2H7.089z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 14V6m3.25 5.788V8.212m-6.5 3.576V8.212"/></g>`), g.Group(children),
	)
}

func Note(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M16.5 4H8a4 4 0 0 0-4 4v8.5a4 4 0 0 0 4 4h6.843a4 4 0 0 0 2.829-1.172l1.656-1.656a4 4 0 0 0 1.172-2.829V8a4 4 0 0 0-4-4"/><path d="M20.5 14H17a3 3 0 0 0-3 3v3.5M8 8h7.5M8 12h5"/></g>`), g.Group(children),
	)
}

func NoteAdd(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M6.5 12a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-8.993a.5.5 0 0 1 .5.5V6h2.493a.5.5 0 0 1 0 1H7v2.493a.5.5 0 1 1-1 0V7H3.507a.5.5 0 0 1 0-1H6V3.507a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.5 14h-1.7c-1.68 0-2.52 0-3.162.327a3 3 0 0 0-1.311 1.311C14 16.28 14 17.12 14 18.8v1.7M15.5 8H14m-2.5 4H13m.75-8h.35c2.24 0 3.36 0 4.216.436a4 4 0 0 1 1.748 1.748c.436.856.436 1.976.436 4.216v3.449c0 .978 0 1.468-.11 1.928c-.099.408-.26.798-.48 1.156c-.247.404-.593.75-1.285 1.442l-.25.25c-.692.692-1.038 1.038-1.442 1.286a4 4 0 0 1-1.156.479c-.46.11-.95.11-1.928.11H10.4c-2.24 0-3.36 0-4.216-.436a4 4 0 0 1-1.748-1.748C4 17.46 4 16.34 4 14.1v-.35"/></g>`), g.Group(children),
	)
}

func Octagon(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.29 3.629a3 3 0 0 1 2.12-.879h5.178a3 3 0 0 1 2.121.879l3.661 3.66a3 3 0 0 1 .879 2.122v5.178a3 3 0 0 1-.879 2.121l-3.66 3.661a3 3 0 0 1-2.122.879H9.41a3 3 0 0 1-2.121-.879l-3.661-3.66a3 3 0 0 1-.879-2.122V9.41a3 3 0 0 1 .879-2.121z"/>`), g.Group(children),
	)
}

func Open(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.75 4.25h-3.5a4 4 0 0 0-4 4v8.5a4 4 0 0 0 4 4h8.5a4 4 0 0 0 4-4v-3.5m-5.5-10h5.5c.276 0 .526.112.707.293m.293 6.207v-5.5a1 1 0 0 0-.293-.707M13.25 10.75l6.5-6.5l.707-.707"/>`), g.Group(children),
	)
}

func PageMargins(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="15.5" height="18.5" x="4.25" y="2.75" rx="2.5"/><path d="M8.25 2.75v18.5m-4-14.5h15.5m-15.5 10.5h15.5m-4-14.5v18.5"/></g>`), g.Group(children),
	)
}

func PaintBucket(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="m18.677 13.35l-1.923 3.17c-.985 2.093 1.314 4.206 3.317 3.05a2.36 2.36 0 0 0 .864-3.225z"/><path stroke-linecap="round" d="m8.778 3.558l6.972 6.972M8.778 3.558L4.38 7.956c-.901.9-1.352 1.351-1.52 1.87a2.3 2.3 0 0 0-.112.704m6.03-6.972l-1.06-1.059m8.032 8.031l-4.398 4.398c-.9.9-1.351 1.351-1.871 1.52c-.457.149-.95.149-1.406 0c-.52-.169-.97-.62-1.871-1.52L4.38 13.104c-.901-.9-1.352-1.351-1.52-1.87a2.3 2.3 0 0 1-.112-.704m13.002 0H2.748m10.188 9.971H3.748"/></g>`), g.Group(children),
	)
}

func PaintBucketAccent(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m18.677 13.35l-1.923 3.17c-.985 2.092 1.314 4.206 3.317 3.05a2.36 2.36 0 0 0 .864-3.225zm-7.325 1.578l4.398-4.398H2.748c0 .237.037.475.112.703c.168.52.619.97 1.52 1.871l1.824 1.824c.9.9 1.351 1.351 1.87 1.52c.458.149.95.149 1.407 0c.52-.169.97-.62 1.871-1.52"/>`), g.Group(children),
	)
}

func Paintbrush(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.32 5.75a3 3 0 0 1 3-3h7.36a3 3 0 0 1 3 3V12H5.32z"/><path fill="currentColor" d="M5.32 12v-.75a.75.75 0 0 0-.75.75zm13.36 0h.75a.75.75 0 0 0-.75-.75zm-8.479 4.111h.75a.75.75 0 0 0-.75-.75zm3.598 0v-.75a.75.75 0 0 0-.75.75zm-8.48-3.361h13.362v-1.5H5.319zm.75 1.861V12h-1.5v2.611zm.75.75a.75.75 0 0 1-.75-.75h-1.5a2.25 2.25 0 0 0 2.25 2.25zm3.382 0H6.82v1.5h3.382zm.75 4.09v-3.34h-1.5v3.34zM12 20.5c-.58 0-1.049-.47-1.049-1.049h-1.5A2.55 2.55 0 0 0 12 22zm1.049-1.049c0 .58-.47 1.049-1.049 1.049V22a2.55 2.55 0 0 0 2.549-2.549zm0-3.34v3.34h1.5v-3.34zm4.132-.75h-3.382v1.5h3.382zm.75-.75a.75.75 0 0 1-.75.75v1.5a2.25 2.25 0 0 0 2.25-2.25zm0-2.611v2.611h1.5V12z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.5 2.75V6.5m-3-3.75v2.5"/></g>`), g.Group(children),
	)
}

func PaintbrushTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M9.781 10.832A4.79 4.79 0 0 0 3.75 15.46v3.79a1 1 0 0 0 1 1h3.79a4.79 4.79 0 0 0 4.628-6.03m-3.387-3.388a4.8 4.8 0 0 1 3.387 3.387m-3.387-3.387l3.19-3.19m.197 6.577l3.19-3.19m-3.387-3.387l3.19-3.19a2.395 2.395 0 0 1 3.387 3.387l-3.19 3.19m-3.387-3.387l3.387 3.387"/>`), g.Group(children),
	)
}

func PanelLeft(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.547 3.75H7.25a3.5 3.5 0 0 0-3.5 3.5v9.5a3.5 3.5 0 0 0 3.5 3.5h2.297m0-16.5h7.203a3.5 3.5 0 0 1 3.5 3.5v9.5a3.5 3.5 0 0 1-3.5 3.5H9.547m0-16.5v16.5"/>`), g.Group(children),
	)
}

func PanelLeftContract(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.25 4.75h-3.5a3 3 0 0 0-3 3v8.5a3 3 0 0 0 3 3h3.5m0-14.5h9a3 3 0 0 1 3 3v8.5a3 3 0 0 1-3 3h-9m0-14.5v14.5m5.86-5L12.9 12m0 0l2.21-2.25M12.9 12h4.7"/>`), g.Group(children),
	)
}

func PanelLeftExpand(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.25 4.75h-3.5a3 3 0 0 0-3 3v8.5a3 3 0 0 0 3 3h3.5m0-14.5h9a3 3 0 0 1 3 3v8.5a3 3 0 0 1-3 3h-9m0-14.5v14.5m6.14-5L17.6 12m0 0l-2.21-2.25M17.6 12h-4.7"/>`), g.Group(children),
	)
}

func PanelLeftOpen(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M16.75 3.75a3.5 3.5 0 0 1 3.5 3.5v9.5a3.5 3.5 0 0 1-3.5 3.5h-2.297V3.75z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.453 3.75h2.297a3.5 3.5 0 0 1 3.5 3.5v9.5a3.5 3.5 0 0 1-3.5 3.5h-2.297m0-16.5H7.25a3.5 3.5 0 0 0-3.5 3.5v9.5a3.5 3.5 0 0 0 3.5 3.5h7.203m0-16.5v16.5"/></g>`), g.Group(children),
	)
}

func PanelRight(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.453 3.75h2.297a3.5 3.5 0 0 1 3.5 3.5v9.5a3.5 3.5 0 0 1-3.5 3.5h-2.297m0-16.5H7.25a3.5 3.5 0 0 0-3.5 3.5v9.5a3.5 3.5 0 0 0 3.5 3.5h7.203m0-16.5v16.5"/>`), g.Group(children),
	)
}

func PanelRightContract(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.75 4.75h3.5a3 3 0 0 1 3 3v8.5a3 3 0 0 1-3 3h-3.5m0-14.5h-9a3 3 0 0 0-3 3v8.5a3 3 0 0 0 3 3h9m0-14.5v14.5m-5.86-5L11.1 12m0 0L8.89 9.75M11.1 12H6.4"/>`), g.Group(children),
	)
}

func PanelRightExpand(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.75 4.75h3.5a3 3 0 0 1 3 3v8.5a3 3 0 0 1-3 3h-3.5m0-14.5h-9a3 3 0 0 0-3 3v8.5a3 3 0 0 0 3 3h9m0-14.5v14.5m-6.14-5L6.4 12m0 0l2.21-2.25M6.4 12h4.7"/>`), g.Group(children),
	)
}

func PanelRightOpen(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M7.25 3.75a3.5 3.5 0 0 0-3.5 3.5v9.5a3.5 3.5 0 0 0 3.5 3.5h2.297V3.75z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.547 3.75H7.25a3.5 3.5 0 0 0-3.5 3.5v9.5a3.5 3.5 0 0 0 3.5 3.5h2.297m0-16.5h7.203a3.5 3.5 0 0 1 3.5 3.5v9.5a3.5 3.5 0 0 1-3.5 3.5H9.547m0-16.5v16.5"/></g>`), g.Group(children),
	)
}

func Pause(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><rect width="5" height="16.5" x="5" y="3.75" rx="2"/><rect width="5" height="16.5" x="14" y="3.75" rx="2"/></g>`), g.Group(children),
	)
}

func Pdf(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor"><path d="m11.79 10.673l-.058.265a9.8 9.8 0 0 1-1.368 3.286m1.425-3.551l.467-2.136c.162-.738-.556-1.316-1.11-.894c-.297.226-.407.665-.26 1.037l.246.617q.286.719.657 1.376Zm0 0a10.4 10.4 0 0 0 2.064 2.596m0 0l2.255-.286c.632-.08 1.09.671.806 1.32c-.207.474-.721.649-1.121.382l-.851-.568a9.4 9.4 0 0 1-1.089-.848Zm0 0l-.095.013a12.3 12.3 0 0 0-3.394.942m0 0q-.626.274-1.228.618l-1.706.975c-.475.271-.577.994-.202 1.423c.332.379.88.338 1.165-.087l1.91-2.837z"/><path stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.792 21.25h8.416a3.5 3.5 0 0 0 3.5-3.5v-5.53a3.5 3.5 0 0 0-1.024-2.475l-5.969-5.97A3.5 3.5 0 0 0 10.24 2.75H7.792a3.5 3.5 0 0 0-3.5 3.5v11.5a3.5 3.5 0 0 0 3.5 3.5"/></g>`), g.Group(children),
	)
}

func PdfTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="1.5" d="M7.792 21.25h8.416a3.5 3.5 0 0 0 3.5-3.5v-5.53a3.5 3.5 0 0 0-1.024-2.475l-5.969-5.97A3.5 3.5 0 0 0 10.24 2.75H7.792a3.5 3.5 0 0 0-3.5 3.5v11.5a3.5 3.5 0 0 0 3.5 3.5"/><path stroke-width="1.5" d="M11.688 3.11v5.66a2 2 0 0 0 2 2h5.662"/><path d="M7.25 16.5v-1m0 0v-2h1a1 1 0 0 1 1 1v0a1 1 0 0 1-1 1zm4 1v-3h.5a1.5 1.5 0 0 1 0 3zm4 0v-1.25m1.5-1.75h-1.5v1.75m0 0h1.5"/></g>`), g.Group(children),
	)
}

func Pencil(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.44 5.78L4.198 16.02a2 2 0 0 0-.565 1.125l-.553 3.774l3.775-.553A2 2 0 0 0 7.98 19.8L18.22 9.56m-3.78-3.78l2.229-2.23a1.6 1.6 0 0 1 2.263 0l1.518 1.518a1.6 1.6 0 0 1 0 2.263l-2.23 2.23M14.44 5.78l3.78 3.78"/>`), g.Group(children),
	)
}

func Pentagon(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.13 3.655a3.18 3.18 0 0 1 3.74 0l6.069 4.409a3.18 3.18 0 0 1 1.155 3.557l-2.318 7.134a3.18 3.18 0 0 1-3.025 2.198H8.249a3.18 3.18 0 0 1-3.025-2.198L2.906 11.62A3.18 3.18 0 0 1 4.06 8.063z"/>`), g.Group(children),
	)
}

func Person(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19 20.75a1 1 0 0 0 1-1v-1.246c.004-2.806-3.974-5.004-8-5.004s-8 2.198-8 5.004v1.246a1 1 0 0 0 1 1zM15.604 6.854a3.604 3.604 0 1 1-7.208 0a3.604 3.604 0 0 1 7.208 0"/>`), g.Group(children),
	)
}

func PersonAdd(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.104 6.854a3.604 3.604 0 1 1-7.208 0a3.604 3.604 0 0 1 7.208 0M10.87 20.75H3.5a1 1 0 0 1-1-1v-1.246c0-2.806 3.974-5.004 8-5.004q.387 0 .77.027"/><path fill="currentColor" fill-rule="evenodd" d="M17.5 23a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-8.993a.5.5 0 0 1 .5.5V17h2.493a.5.5 0 1 1 0 1H18v2.494a.5.5 0 0 1-1 0V18h-2.493a.5.5 0 1 1 0-1H17v-2.493a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func PersonAddTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M19.25 1.75v6m-3-3h6m-11.814 8.814a3.907 3.907 0 1 0 0-7.814a3.907 3.907 0 0 0 0 7.814m0 0a6.686 6.686 0 0 1 6.685 6.686m-6.685-6.686A6.686 6.686 0 0 0 3.75 20.25"/>`), g.Group(children),
	)
}

func PersonCircle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 13.826a3.506 3.506 0 1 0 0-7.013a3.506 3.506 0 0 0 0 7.013m0 0a6 6 0 0 1 5.953 5.254M12 13.826a6 6 0 0 0-5.953 5.254m0 0A9.2 9.2 0 0 0 12 21.25a9.2 9.2 0 0 0 5.953-2.17m-11.906 0a9.25 9.25 0 1 1 11.907 0"/>`), g.Group(children),
	)
}

func PersonMultiple(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12.933 7.28a3.3 3.3 0 0 1-3.305 3.305a3.305 3.305 0 1 1 3.305-3.305m2.714 12.745c.475 0 .86-.41.86-.917v-1.143c.002-1.547-1.234-2.893-2.972-3.72a9.2 9.2 0 0 0-3.906-.87c-3.462 0-6.879 2.016-6.879 4.59v1.143c0 .506.385.917.86.917zm3.353 0h1.333c.506 0 .917-.41.917-.917V18.2c.002-1.562-1.514-2.892-3.521-3.575M15.434 5.89a2.929 2.929 0 1 1-1.807 5.551"/>`), g.Group(children),
	)
}

func PersonTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><circle cx="12" cy="8.196" r="4.446"/><path d="M19.608 20.25a7.608 7.608 0 0 0-15.216 0"/></g>`), g.Group(children),
	)
}

func Phone(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="12.5" height="18.5" x="5.75" y="2.75" rx="3"/><path d="M11 17.75h2"/></g>`), g.Group(children),
	)
}

func PhoneAccept(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M8.14 15.733c2.158 2.158 4.278 3.28 5.89 3.864c1.768.64 3.606-.117 4.935-1.446l.459-.458a1.5 1.5 0 0 0 0-2.122l-1.149-1.149a1.5 1.5 0 0 0-2.121 0l-.387.387a2 2 0 0 1-2.828 0l-3.713-3.712a2 2 0 0 1 0-2.829l.387-.387a1.5 1.5 0 0 0 0-2.12l-1.15-1.15a1.5 1.5 0 0 0-2.12 0l-.572.572c-1.262 1.262-2.013 2.99-1.438 4.68c.538 1.58 1.622 3.685 3.806 5.87Z"/>`), g.Group(children),
	)
}

func PhoneHangUp(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M12.116 7.953c-3.053 0-5.346.706-6.899 1.433c-1.702.797-2.467 2.632-2.467 4.512v.649a1.5 1.5 0 0 0 1.5 1.5h1.625a1.5 1.5 0 0 0 1.5-1.5V14a2 2 0 0 1 2-2h5.25a2 2 0 0 1 2 2v.547a1.5 1.5 0 0 0 1.5 1.5h1.625a1.5 1.5 0 0 0 1.5-1.5v-.81c0-1.784-.691-3.537-2.293-4.325c-1.496-.736-3.752-1.459-6.841-1.459Z"/>`), g.Group(children),
	)
}

func Photo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.33 17.657c.11-.366.17-.755.17-1.157v-9a4 4 0 0 0-4-4h-9a4 4 0 0 0-4 4v9.07m16.83 1.087l-.088-.104l-2.466-2.976a2 2 0 0 0-3.073-.008l-1.312 1.566l-.214.261m7.153 1.26a4 4 0 0 1-3.713 2.842m0 0l-.117.002h-9a4 4 0 0 1-4-3.93m13.117 3.928l-.093-.106l-3.347-3.996m-9.676.175l.177-.201l3.206-3.827a2 2 0 0 1 3.066 0l3.227 3.853"/><circle cx="15.091" cy="8.909" r="1.5" fill="currentColor"/></g>`), g.Group(children),
	)
}

func PhotoFilter(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="8.73" r="5.98"/><circle cx="8.729" cy="15.27" r="5.98"/><circle cx="15.271" cy="15.27" r="5.98"/></g>`), g.Group(children),
	)
}

func PictureInPicture(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.25 18.25h-3.5a3 3 0 0 1-3-3v-8.5a3 3 0 0 1 3-3h12.5a3 3 0 0 1 3 3v3.5"/><rect width="12" height="10" x="11" y="12" fill="currentColor" rx="2"/></g>`), g.Group(children),
	)
}

func PictureInPictureEnter(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.25 18.25h-3.5a3 3 0 0 1-3-3v-8.5a3 3 0 0 1 3-3h12.5a3 3 0 0 1 3 3v3.5"/><rect width="12" height="10" x="11" y="12" fill="currentColor" rx="2"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 7.667V10.4a.6.6 0 0 1-.176.424M6.667 11H9.4a.6.6 0 0 0 .424-.176M6 7l3 3l.824.824"/></g>`), g.Group(children),
	)
}

func PictureInPictureExit(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.75 5.75h3.5a3 3 0 0 1 3 3v8.5a3 3 0 0 1-3 3H5.75a3 3 0 0 1-3-3v-3.5"/><rect width="12" height="10" x="13" y="12" fill="currentColor" rx="2" transform="rotate(180 13 12)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18 13.667V16.4a.6.6 0 0 1-.176.424M14.667 17H17.4a.6.6 0 0 0 .424-.176M14 13l3 3l.824.824"/></g>`), g.Group(children),
	)
}

func PictureInPictureTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.25 4H5.75a3 3 0 0 0-3 3v10a3 3 0 0 0 3 3h12.5a3 3 0 0 0 3-3V7a3 3 0 0 0-3-3"/><rect width="8.5" height="7.083" x="10.25" y="10.42" fill="currentColor" rx="1.5"/></g>`), g.Group(children),
	)
}

func Pin(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.77 16.233l-4.02 4.02M14.976 3.336l5.69 5.691a2 2 0 0 1-.698 3.282L16.595 13.6a4 4 0 0 0-2.426 2.674l-.689 2.5a1.5 1.5 0 0 1-2.507.662L4.568 13.03a1.5 1.5 0 0 1 .662-2.507l2.5-.688a4 4 0 0 0 2.673-2.427l1.291-3.372a2 2 0 0 1 3.282-.7"/>`), g.Group(children),
	)
}

func PinOff(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m7.77 16.233l-4.02 4.02m16.968 1.529l-6.258-6.26M2.217 3.283L8.48 9.545m5.979 5.979q-.183.357-.292.75l-.688 2.5a1.5 1.5 0 0 1-2.507.663L4.568 13.03a1.5 1.5 0 0 1 .662-2.507l2.5-.688a4 4 0 0 0 .75-.291m5.98 5.979l-5.98-5.98m1.765-1.775q.087-.175.158-.36l1.291-3.372a2 2 0 0 1 3.282-.7l5.69 5.691a2 2 0 0 1-.698 3.282L16.595 13.6q-.184.07-.36.159"/>`), g.Group(children),
	)
}

func Play(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M5.5 12V5.624c0-1.974 2.18-3.17 3.844-2.108l10 6.376c1.541.983 1.541 3.233 0 4.216l-10 6.376C7.68 21.545 5.5 20.35 5.5 18.376z"/>`), g.Group(children),
	)
}

func PlayCircular(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><circle cx="12" cy="12" r="9.25" stroke-linecap="round" stroke-linejoin="round"/><path d="M8.93 13.29c0 1.098 0 1.646.23 1.964c.202.277.51.456.85.492c.391.041.867-.232 1.818-.779l2.244-1.29c.957-.55 1.435-.825 1.595-1.185c.14-.313.14-.671 0-.984c-.16-.36-.639-.635-1.595-1.184l-2.244-1.291c-.951-.547-1.427-.82-1.817-.779c-.34.036-.65.215-.85.492c-.23.318-.23.866-.23 1.963z"/></g>`), g.Group(children),
	)
}

func Printer(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.75 17v1.05c0 1.12 0 1.68.218 2.108a2 2 0 0 0 .874.874c.428.218.988.218 2.108.218h4.1c1.12 0 1.68 0 2.108-.218a2 2 0 0 0 .874-.874c.218-.428.218-.988.218-2.108V17m-10.5 0v-1.05c0-1.12 0-1.68.218-2.108a2 2 0 0 1 .874-.874c.428-.218.988-.218 2.108-.218h4.1c1.12 0 1.68 0 2.108.218a2 2 0 0 1 .874.874c.218.428.218.988.218 2.108V17m-10.5 0h-.8c-1.12 0-1.68 0-2.108-.218a2 2 0 0 1-.874-.874c-.218-.428-.218-.988-.218-2.108v-3c0-1.68 0-2.52.327-3.162a3 3 0 0 1 1.311-1.311C5.03 6 5.87 6 7.55 6h8.9c1.68 0 2.52 0 3.162.327a3 3 0 0 1 1.311 1.311c.327.642.327 1.482.327 3.162v3c0 1.12 0 1.68-.218 2.108a2 2 0 0 1-.874.874C19.73 17 19.17 17 18.05 17h-.8M6.75 4.25a1.5 1.5 0 0 1 1.5-1.5h7.5a1.5 1.5 0 0 1 1.5 1.5V6H6.75z"/>`), g.Group(children),
	)
}

func Prohibited(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.46 18.54A9.25 9.25 0 0 0 18.54 5.46M5.459 18.541A9.25 9.25 0 0 1 18.54 5.46M5.46 18.54L18.54 5.46"/>`), g.Group(children),
	)
}

func Python(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 7.75H7.5m4.5 9h4.5m0 0h1.521c.807 0 1.634-.188 2.13-.824c.531-.679 1.099-1.835 1.099-3.676c0-1.84-.568-2.997-1.098-3.676c-.497-.636-1.324-.824-2.13-.824H16.5m0 9v1.521c0 .807-.188 1.634-.824 2.13c-.679.531-1.835 1.099-3.676 1.099c-1.84 0-2.997-.568-3.676-1.098c-.636-.497-.824-1.324-.824-2.13V16.75m0-9H5.978c-.807 0-1.633.188-2.13.824c-.53.679-1.098 1.835-1.098 3.676c0 1.84.568 2.997 1.098 3.676c.497.636 1.323.824 2.13.824H7.5m0-9V6.228c0-.807.188-1.633.824-2.13C9.003 3.568 10.159 3 12 3c1.84 0 2.997.568 3.676 1.098c.636.497.824 1.323.824 2.13V7.75m-9 9v-2.5a2 2 0 0 1 2-2h5a2 2 0 0 0 2-2v-2.5"/><path fill="currentColor" d="M15 18.75a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0m-6-13a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0"/></g>`), g.Group(children),
	)
}

func QrCode(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="7.5" height="7.5" x="2.75" y="2.75" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="2.5"/><rect width="7.5" height="7.5" x="13.75" y="2.75" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="2.5"/><rect width="7.5" height="7.5" x="2.75" y="13.75" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="2.5"/><rect width="3" height="3" x="5" y="5" fill="currentColor" rx="1.5"/><rect width="3" height="3" x="16" y="5" fill="currentColor" rx="1.5"/><rect width="3" height="3" x="5" y="16" fill="currentColor" rx="1.5"/><rect width="3" height="3" x="13" y="13" fill="currentColor" rx="1.5"/><rect width="3" height="3" x="16" y="16" fill="currentColor" rx="1.5"/><rect width="3" height="3" x="19" y="19" fill="currentColor" rx="1.5"/><rect width="3" height="3" x="19" y="13" fill="currentColor" rx="1.5"/><rect width="3" height="3" x="13" y="19" fill="currentColor" rx="1.5"/></g>`), g.Group(children),
	)
}

func Question(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.817 8.808a4.183 4.183 0 1 1 7.518 2.526l-.133.145c-.065.07-.29.286-.363.347a4 4 0 0 1-.353.266l-1.517 1.045a2.81 2.81 0 0 0-1.215 2.315"/><circle cx="11.754" cy="19.141" r=".984" fill="currentColor"/></g>`), g.Group(children),
	)
}

func QuestionCircle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.301 9.709a2.699 2.699 0 1 1 4.85 1.63a4 4 0 0 1-.32.317c-.092.078-.137.11-.227.171l-.979.675a1.81 1.81 0 0 0-.784 1.493"/><circle cx="11.828" cy="16.74" r="1" fill="currentColor"/><circle cx="12" cy="12" r="9.25" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"/></g>`), g.Group(children),
	)
}

func Quote(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.937 10.453l-.01.13c0 3.35-2.038 5.115-4.63 6.058m4.64-6.188a3.093 3.093 0 1 1-6.187.001a3.093 3.093 0 0 1 6.187-.001m10.313 0l-.01.13c0 3.35-2.038 5.115-4.63 6.058m4.64-6.188a3.093 3.093 0 1 1-6.187 0a3.093 3.093 0 0 1 6.187 0"/>`), g.Group(children),
	)
}

func RectangleWide(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="18.5" height="14.5" x="2.75" y="4.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="4"/>`), g.Group(children),
	)
}

func Reverse(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M18.005 18.323c1.019.784 2.495.057 2.495-1.229V6.906c0-1.286-1.476-2.013-2.495-1.229L11.78 10.87a1.473 1.473 0 0 0 0 2.262z"/><path d="M8.754 18.323c1.02.784 2.496.057 2.496-1.229V6.906c0-1.286-1.476-2.013-2.496-1.229L2.53 10.87a1.473 1.473 0 0 0 0 2.262z"/></g>`), g.Group(children),
	)
}

func Rhombus(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.775 14.475a3.5 3.5 0 0 1 0-4.95l5.75-5.75a3.5 3.5 0 0 1 4.95 0l5.75 5.75a3.5 3.5 0 0 1 0 4.95l-5.75 5.75a3.5 3.5 0 0 1-4.95 0z"/>`), g.Group(children),
	)
}

func Ribbon(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M18.41 9.16a6.4 6.4 0 0 1-2.426 5.02A6.38 6.38 0 0 1 12 15.57c-1.506 0-2.89-.52-3.984-1.388A6.41 6.41 0 1 1 18.41 9.16"/><path d="M15.984 14.18v7.07L12 18.267L8.016 21.25v-7.07"/></g>`), g.Group(children),
	)
}

func RibbonStar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.41 9.16a6.4 6.4 0 0 1-2.426 5.02A6.38 6.38 0 0 1 12 15.57c-1.506 0-2.89-.52-3.984-1.388A6.41 6.41 0 1 1 18.41 9.16"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.984 14.18v7.07L12 18.267L8.016 21.25v-7.07"/><path fill="currentColor" d="M11.455 6.74c.179-.307.268-.461.385-.513a.4.4 0 0 1 .32 0c.117.052.206.206.385.513l.488.838a1 1 0 0 0 .112.168q.045.047.106.076a1 1 0 0 0 .194.055l.947.205c.348.075.522.113.607.208c.075.083.11.193.1.304c-.014.127-.133.26-.37.525l-.646.723a1 1 0 0 0-.125.159a.4.4 0 0 0-.04.123a1 1 0 0 0 .008.202l.098.964c.036.355.054.532-.01.642a.4.4 0 0 1-.26.188c-.124.027-.287-.045-.612-.188l-.887-.391a1 1 0 0 0-.19-.07a.4.4 0 0 0-.13 0a1 1 0 0 0-.19.07l-.886.39c-.326.144-.49.216-.614.19a.4.4 0 0 1-.259-.189c-.064-.11-.046-.287-.01-.642l.098-.964c.01-.102.015-.153.008-.202a.4.4 0 0 0-.04-.123a1 1 0 0 0-.125-.159L9.27 9.12c-.238-.265-.356-.398-.37-.525a.4.4 0 0 1 .1-.304c.085-.095.259-.133.607-.208l.947-.205c.1-.022.15-.033.194-.055a.4.4 0 0 0 .106-.076a1 1 0 0 0 .112-.168z"/></g>`), g.Group(children),
	)
}

func Roblox(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="16.02" height="16.02" x="6.336" y="2.19" rx="2.5" transform="rotate(15 6.336 2.19)"/><rect width="5.34" height="5.34" x="10.112" y="8.73" rx=".8" transform="rotate(15 10.112 8.73)"/></g>`), g.Group(children),
	)
}

func Ruler(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><rect width="9.5" height="18.5" x="6.75" y="2.75" stroke-linejoin="round" rx="2"/><path d="M6.75 12h4.5m-4.5-4.5h4.5m-4.5 9h4.5"/></g>`), g.Group(children),
	)
}

func Save(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.558 3.75H7.25a3.5 3.5 0 0 0-3.5 3.5v9.827a3.173 3.173 0 0 0 3.173 3.173v0m.635-16.5v2.442a2 2 0 0 0 2 2h2.346a2 2 0 0 0 2-2V3.75m-6.346 0h6.346m0 0h.026a3 3 0 0 1 2.122.879l3.173 3.173a3.5 3.5 0 0 1 1.025 2.475v6.8a3.173 3.173 0 0 1-3.173 3.173v0m-10.154 0V15a3 3 0 0 1 3-3h4.154a3 3 0 0 1 3 3v5.25m-10.154 0h10.154"/>`), g.Group(children),
	)
}

func SaveAs(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.558 3.75v1.27c0 1.094 0 1.641.208 2.061a2 2 0 0 0 .903.903c.42.208.967.208 2.062.208s1.642 0 2.062-.208a2 2 0 0 0 .902-.903c.209-.42.209-.967.209-2.062V3.75m-6.346 0h6.346m-6.346 0c-.751 0-1.126 0-1.438.067A3 3 0 0 0 3.817 6.12c-.067.312-.067.687-.067 1.438v9.519c0 .16 0 .241.003.309a3 3 0 0 0 2.861 2.86c.068.004.148.004.31.004m6.98-16.5h.052a3 3 0 0 1 2.078.86l.037.037l2.773 2.774c.519.519.778.778.964 1.081a3 3 0 0 1 .36.867c.028.12.047.242.059.381M6.923 20.25V16c0-1.4 0-2.1.273-2.635a2.5 2.5 0 0 1 1.092-1.092C8.823 12 9.523 12 10.923 12h2.154c1.168 0 1.849 0 2.353.158M6.923 20.25H9.75"/><path fill="currentColor" d="M21.796 12.204a2.4 2.4 0 0 0-1.546-.7a2.4 2.4 0 0 0-1.854.7l-1.47 1.47l-4.305 4.306a2.3 2.3 0 0 0-.635 1.172l-.21 1.098l-.261 1.371a.74.74 0 0 0 .195.67a.74.74 0 0 0 .669.194l2.469-.47a2.3 2.3 0 0 0 1.172-.636l5.776-5.776a2.404 2.404 0 0 0 0-3.399"/></g>`), g.Group(children),
	)
}

func ScreenSize(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="18.5" height="15.5" x="2.75" y="4.25" rx="3"/><path d="M6.75 12.25v-4h4m6.5 3.5v4h-4"/></g>`), g.Group(children),
	)
}

func Script(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.25 21.25H4A2.25 2.25 0 0 1 1.75 19v-1.25a1.5 1.5 0 0 1 1.5-1.5h1.5m10.5 5a2.5 2.5 0 0 0 2.5-2.5v-11m-2.5 13.5a2.5 2.5 0 0 1-2.5-2.5v-1.5a1 1 0 0 0-1-1h-7m15.376-13.5H8.25a3.5 3.5 0 0 0-3.5 3.5v10m13-8.5h3.5a1 1 0 0 0 1-1V5a2.25 2.25 0 0 0-4.5 0zm-9.25-.5h6m-6 4h4"/>`), g.Group(children),
	)
}

func ScriptTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.25 21.25H4A2.25 2.25 0 0 1 1.75 19v-1.25a1.5 1.5 0 0 1 1.5-1.5h1.5m10.5 5a2.5 2.5 0 0 0 2.5-2.5v-11m-2.5 13.5a2.5 2.5 0 0 1-2.5-2.5v-1.5a1 1 0 0 0-1-1h-7m15.376-13.5H8.25a3.5 3.5 0 0 0-3.5 3.5v10m13-8.5h3.5a1 1 0 0 0 1-1V5a2.25 2.25 0 0 0-4.5 0z"/>`), g.Group(children),
	)
}

func Search(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M15.553 15.553a7.06 7.06 0 1 0-9.985-9.985a7.06 7.06 0 0 0 9.985 9.985m0 0L20 20"/>`), g.Group(children),
	)
}

func SearchCancel(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m8.44 12.682l2.12-2.121m0 0l2.122-2.122m-2.121 2.122l2.12 2.12m-2.12-2.12L8.439 8.439m7.114 7.114a7.06 7.06 0 1 0-9.985-9.985a7.06 7.06 0 0 0 9.985 9.985m0 0L20 20"/><path d="M15.553 5.568a7.06 7.06 0 1 1-9.985 9.985a7.06 7.06 0 0 1 9.985-9.985"/></g>`), g.Group(children),
	)
}

func SectionBreak(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 12h18.5M4 2.75V5.5a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2.75M4 21.25V18.5a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v2.75"/>`), g.Group(children),
	)
}

func Send(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.76 12H6.832m0 0c0-.275-.057-.55-.17-.808L4.285 5.814c-.76-1.72 1.058-3.442 2.734-2.591L20.8 10.217c1.46.74 1.46 2.826 0 3.566L7.02 20.777c-1.677.851-3.495-.872-2.735-2.591l2.375-5.378A2 2 0 0 0 6.83 12"/>`), g.Group(children),
	)
}

func Settings(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.05 6.462a2 2 0 0 0 2.63-1.519l.32-1.72a9 9 0 0 1 3.998 0l.322 1.72a2 2 0 0 0 2.63 1.519l1.649-.58a9 9 0 0 1 2.001 3.46l-1.33 1.14a2 2 0 0 0 0 3.037l1.33 1.139a9 9 0 0 1-2.001 3.46l-1.65-.58a2 2 0 0 0-2.63 1.519L14 20.777a9 9 0 0 1-3.998 0l-.322-1.72a2 2 0 0 0-2.63-1.519l-1.649.58a9 9 0 0 1-2.001-3.46l1.33-1.14a2 2 0 0 0 0-3.036L3.4 9.342a9 9 0 0 1 2-3.46zM12 9a3 3 0 1 1 0 6a3 3 0 0 1 0-6" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ShapeDifference(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor"><path d="M1.998 5.75A3.75 3.75 0 0 1 5.748 2h7a3.75 3.75 0 0 1 3.75 3.75V7.5h-5.246a3.75 3.75 0 0 0-3.75 3.75v5.25H5.748a3.75 3.75 0 0 1-3.75-3.75z"/><path d="M7.502 16.5h5.246a3.75 3.75 0 0 0 3.75-3.75V7.5h1.754a3.75 3.75 0 0 1 3.75 3.75v7a3.75 3.75 0 0 1-3.75 3.75h-7a3.75 3.75 0 0 1-3.75-3.75z"/></g>`), g.Group(children),
	)
}

func ShapeIntersect(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12.748 15.75a3 3 0 0 0 3-3v-4.5h-4.496a3 3 0 0 0-3 3v4.5z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.252 15.75v2.5a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3h-2.504m-7.496 7.5v-4.5a3 3 0 0 1 3-3h4.496m-7.496 7.5h4.496a3 3 0 0 0 3-3v-4.5m-7.496 7.5H5.748a3 3 0 0 1-3-3v-7a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v2.5"/></g>`), g.Group(children),
	)
}

func ShapeSubtract(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M12.748 2.75h-7a3 3 0 0 0-3 3v7a3 3 0 0 0 3 3h2.504v-4.5a3 3 0 0 1 3-3h4.496v-2.5a3 3 0 0 0-3-3"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.252 15.75v2.5a3 3 0 0 0 3 3h7a3 3 0 0 0 3-3v-7a3 3 0 0 0-3-3h-2.504m-7.496 7.5v-4.5a3 3 0 0 1 3-3h4.496m-7.496 7.5H5.748a3 3 0 0 1-3-3v-7a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v2.5"/></g>`), g.Group(children),
	)
}

func ShapeUnion(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path fill="currentColor" d="M2.748 5.75a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3z"/><path fill="currentColor" d="M8.252 11.25a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.748 5.75a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M8.252 11.25a3 3 0 0 1 3-3h7a3 3 0 0 1 3 3v7a3 3 0 0 1-3 3h-7a3 3 0 0 1-3-3z"/></g>`), g.Group(children),
	)
}

func Shield(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.25 10.907V7.272c0-.829-.633-1.521-1.453-1.644c-.951-.142-2.18-.376-3.078-.722c-.907-.349-1.997-1.007-2.762-1.505a1.76 1.76 0 0 0-1.914 0c-.764.498-1.855 1.156-2.762 1.505c-.899.346-2.127.58-3.078.722c-.82.123-1.453.815-1.453 1.644v3.635a10.13 10.13 0 0 0 5.363 8.939l.23.123l1.962.946a1.6 1.6 0 0 0 1.39 0l1.961-.946l.23-.123a10.13 10.13 0 0 0 5.364-8.939"/>`), g.Group(children),
	)
}

func ShieldCancel(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.25 10.907V7.272c0-.829-.633-1.521-1.453-1.644c-.951-.142-2.18-.376-3.078-.722c-.907-.349-1.997-1.007-2.762-1.505a1.76 1.76 0 0 0-1.914 0c-.764.498-1.855 1.156-2.762 1.505c-.899.346-2.127.58-3.078.722c-.82.123-1.453.815-1.453 1.644v3.635a10.13 10.13 0 0 0 5.363 8.939l.23.123l1.962.946a1.6 1.6 0 0 0 1.39 0l1.961-.946l.23-.123a10.13 10.13 0 0 0 5.364-8.939M9.5 9.5l5 5m0-5l-5 5"/>`), g.Group(children),
	)
}

func ShieldCheckmark(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M20.25 10.907V7.272c0-.829-.633-1.521-1.453-1.644c-.951-.142-2.18-.376-3.078-.722c-.907-.349-1.997-1.007-2.762-1.505a1.76 1.76 0 0 0-1.914 0c-.764.498-1.855 1.156-2.762 1.505c-.899.346-2.127.58-3.078.722c-.82.123-1.453.815-1.453 1.644v3.635a10.13 10.13 0 0 0 5.363 8.939l.23.123l1.962.946a1.6 1.6 0 0 0 1.39 0l1.961-.946l.23-.123a10.13 10.13 0 0 0 5.364-8.939"/><path d="m15.509 10l-4.076 4.076a.6.6 0 0 1-.849 0l-2.093-2.09"/></g>`), g.Group(children),
	)
}

func ShieldKeyhole(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.25 10.907V7.272c0-.829-.633-1.521-1.453-1.644c-.951-.142-2.18-.376-3.078-.722c-.907-.349-1.997-1.007-2.762-1.505a1.76 1.76 0 0 0-1.914 0c-.764.498-1.855 1.156-2.762 1.505c-.899.346-2.127.58-3.078.722c-.82.123-1.453.815-1.453 1.644v3.635a10.13 10.13 0 0 0 5.363 8.939l.23.123l1.962.946a1.6 1.6 0 0 0 1.39 0l1.961-.946l.23-.123a10.13 10.13 0 0 0 5.364-8.939"/><circle cx="12" cy="10.5" r="2" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 14.75v-2.5"/></g>`), g.Group(children),
	)
}

func Skull(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M17.99 17.284c2.225-1.528 3.26-3.442 3.26-6.164c0-4.62-4.141-8.37-9.25-8.37S2.75 6.5 2.75 11.12c0 2.722 1.437 4.636 3.663 6.164c-.175.61-.132 2.251.187 2.78c.414.687 1.219 1.186 1.87 1.186c.743 0 1.396-.474 1.765-1.186c.37.712 1.022 1.186 1.765 1.186s1.396-.473 1.765-1.186c.37.713 1.022 1.186 1.765 1.186c.651 0 1.454-.499 1.94-1.186c.374-.529.52-2.17.52-2.78"/><path fill="currentColor" d="M5.75 10.25a2.5 2.5 0 0 1 5 0v1a1.5 1.5 0 0 1-1.5 1.5h-1a2.5 2.5 0 0 1-2.5-2.5m7.5 0a2.5 2.5 0 1 1 2.5 2.5h-1a1.5 1.5 0 0 1-1.5-1.5zm-1.773 2.93l-1.224 2.176a.6.6 0 0 0 .523.894h2.448a.6.6 0 0 0 .523-.894l-1.224-2.176a.6.6 0 0 0-1.046 0"/></g>`), g.Group(children),
	)
}

func Soundwave(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 20.75V3.25m8 11.71V9.04M4 14.96V9.04m12 8.872V6.088M8 17.912V6.088"/>`), g.Group(children),
	)
}

func Spacebar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 10v2.5a2 2 0 0 0 2 2h12.5a2 2 0 0 0 2-2V10"/>`), g.Group(children),
	)
}

func Sparkle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M11.256 3.567c.266-.675 1.222-.675 1.488 0l2.047 5.19a.8.8 0 0 0 .451.452l5.191 2.047c.675.266.675 1.222 0 1.488l-5.19 2.047a.8.8 0 0 0-.452.451l-2.047 5.191c-.266.675-1.222.675-1.488 0l-2.047-5.19a.8.8 0 0 0-.451-.452l-5.191-2.047c-.675-.266-.675-1.222 0-1.488l5.19-2.047a.8.8 0 0 0 .452-.451z"/><circle cx="5.25" cy="5.25" r="1.25" fill="currentColor"/><circle cx="19" cy="19" r="1" fill="currentColor"/></g>`), g.Group(children),
	)
}

func Spinner(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21.25A9.25 9.25 0 1 0 2.75 12"/>`), g.Group(children),
	)
}

func Square(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<rect width="16.5" height="16.5" x="3.75" y="3.75" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="4"/>`), g.Group(children),
	)
}

func Star(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.704 4.325a1.5 1.5 0 0 1 2.592 0l1.818 3.12a1.5 1.5 0 0 0 .978.712l3.53.764a1.5 1.5 0 0 1 .8 2.465l-2.405 2.693a1.5 1.5 0 0 0-.374 1.15l.363 3.593a1.5 1.5 0 0 1-2.097 1.524l-3.304-1.456a1.5 1.5 0 0 0-1.21 0l-3.304 1.456a1.5 1.5 0 0 1-2.097-1.524l.363-3.593a1.5 1.5 0 0 0-.373-1.15l-2.406-2.693a1.5 1.5 0 0 1 .8-2.465l3.53-.764a1.5 1.5 0 0 0 .979-.711z"/>`), g.Group(children),
	)
}

func StrokeThickness(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 4.5h18.5M20.5 10h-17a.75.75 0 0 0 0 1.5h17a.75.75 0 0 0 0-1.5"/><path fill="currentColor" d="M19.75 17H4.25a1.5 1.5 0 0 0 0 3h15.5a1.5 1.5 0 0 0 0-3"/></g>`), g.Group(children),
	)
}

func Subtract(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 12h16.5"/>`), g.Group(children),
	)
}

func Symbols(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M13.9 3.918a1.955 1.955 0 0 1 2.765 0l.624.624l.624-.624a1.955 1.955 0 1 1 2.764 2.765l-.623.624l.005.005l-2.764 2.765l-.006-.006l-.006.006l-2.764-2.765l.005-.005l-.623-.624a1.955 1.955 0 0 1 0-2.765Z"/><path stroke-linecap="round" d="M3.736 9.091a3.366 3.366 0 1 0 4.76-4.76m-4.76 4.76a3.366 3.366 0 1 1 4.76-4.76m-4.76 4.76l4.76-4.76M2.75 17.289h3.366m0 0H9.48m-3.365 0v-3.366m0 3.366v3.366m7.328-4.237a3.365 3.365 0 0 1 6.165-.812l.189.297m.262-1.98v1.98h-.262m-1.718 0h1.718m-6.47 4.752v-1.98h.267m1.714 0h-1.714m6.35-.515a3.366 3.366 0 0 1-6.165.812l-.185-.297"/></g>`), g.Group(children),
	)
}

func Table(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="18.5" height="16.5" x="2.75" y="3.75" rx="3"/><path d="M2.75 7.75h18.5M2.75 14h18.5M8.92 7.75v12.5m6.17-12.5v12.5"/></g>`), g.Group(children),
	)
}

func TableSimple(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="16.5" height="16.5" x="3.75" y="3.75" rx="3"/><path d="M3.75 9.25h16.5m-16.5 5.5h16.5m-11-11v16.5m5.5-16.5v16.5"/></g>`), g.Group(children),
	)
}

func Tablet(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="18.5" height="15" x="2.75" y="4.5" rx="3"/><path d="M10 16h4"/></g>`), g.Group(children),
	)
}

func Tag(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M20.162 10.926L13.716 4.48a2.5 2.5 0 0 0-1.767-.732h-5.2a3 3 0 0 0-3 3v5.2a2.5 2.5 0 0 0 .731 1.768l6.445 6.446a4 4 0 0 0 5.657 0l1.79-1.79l1.79-1.79a4 4 0 0 0 0-5.657"/><circle cx="7.738" cy="7.738" r="1.277" fill="currentColor" transform="rotate(-45 7.738 7.738)"/></g>`), g.Group(children),
	)
}

func TagAccent(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.162 10.925L13.716 4.48a2.5 2.5 0 0 0-1.767-.732h-5.2a3 3 0 0 0-3 3v5.2a2.5 2.5 0 0 0 .731 1.768l6.445 6.446a4 4 0 0 0 5.657 0l3.58-3.58a4 4 0 0 0 0-5.657"/>`), g.Group(children),
	)
}

func TagMultiple(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.14 10.015l-6.209-6.21a2 2 0 0 0-1.414-.585H6.115a2 2 0 0 0-2 2v5.402a2 2 0 0 0 .586 1.414l6.21 6.209a3 3 0 0 0 4.242 0l3.987-3.988a3 3 0 0 0 0-4.242"/><circle cx="8.562" cy="7.667" r="1.138" fill="currentColor" transform="rotate(-45 8.562 7.667)"/><path fill="currentColor" d="m5.232 14.926l4.794 4.794a4.25 4.25 0 0 0 6.01 0l2.574-2.573l2.121-2.122a3.74 3.74 0 0 1-1.06 3.183l-2.574 2.573a5.75 5.75 0 0 1-8.131 0L4.17 15.986a2.75 2.75 0 0 1-.806-1.944v-.983z"/></g>`), g.Group(children),
	)
}

func TagMultipleVar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.14 10.015l-6.209-6.21a2 2 0 0 0-1.414-.585H6.115a2 2 0 0 0-2 2v5.402a2 2 0 0 0 .586 1.414l6.209 6.209a3 3 0 0 0 4.243 0l3.987-3.988a3 3 0 0 0 0-4.242"/><circle cx="8.562" cy="7.667" r="1.138" fill="currentColor" transform="rotate(-45 8.562 7.667)"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.701 14.456l5.795 5.794a5 5 0 0 0 7.07 0l3.696-3.694"/></g>`), g.Group(children),
	)
}

func TaskList(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10.25 4.5h11m-14-1.446L4.357 5.946L2.75 4.34m7.5 7.66h11m-14-1.446l-2.893 2.892L2.75 11.84m7.5 7.66h11m-14-1.446l-2.893 2.892L2.75 19.34"/>`), g.Group(children),
	)
}

func Terminal(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="18.5" height="15.5" x="2.75" y="4.25" rx="3.5"/><path d="m7.25 9l3 3l-3 3m5.5 0h4"/></g>`), g.Group(children),
	)
}

func Text(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.25 8.389l-.62-1.235A3 3 0 0 0 15.95 5.5h-7.9a3 3 0 0 0-2.68 1.654L4.75 8.39M12 5.5v13m0 0h-1.45m1.45 0h1.45"/>`), g.Group(children),
	)
}

func TextAdd(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m19.25 8.389l-.62-1.235A3 3 0 0 0 15.95 5.5h-7.9a3 3 0 0 0-2.68 1.654L4.75 8.39M12 5.5v6"/><path fill="currentColor" fill-rule="evenodd" d="M17.5 23a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-8.993a.5.5 0 0 1 .5.5V17h2.493a.5.5 0 1 1 0 1H18v2.493a.5.5 0 1 1-1 0V18h-2.493a.5.5 0 1 1 0-1H17v-2.493a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func TextAlignCenter(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M5.25 5.25h13.5M2.75 12h18.5M7 18.75h10"/>`), g.Group(children),
	)
}

func TextAlignJustify(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M2.75 5.25h18.5M2.75 12h18.5m-18.5 6.75h18.5"/>`), g.Group(children),
	)
}

func TextAlignLeft(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M2.75 5.25h13.5M2.75 12h18.5m-18.5 6.75h10"/>`), g.Group(children),
	)
}

func TextAlignRight(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M7.75 5.25h13.5M2.75 12h18.5m-10 6.75h10"/>`), g.Group(children),
	)
}

func TextBold(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="3" d="M7.5 11.776h4.551c1.712 0 3.099-1.405 3.099-3.138S13.763 5.5 12.051 5.5H8.5a1 1 0 0 0-1 1zm0 0h5.625a3.37 3.37 0 0 1 3.375 3.362a3.37 3.37 0 0 1-3.375 3.362H8.5a1 1 0 0 1-1-1z"/>`), g.Group(children),
	)
}

func TextBulletList(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="4.443" cy="5.081" r="1.331" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.123 5.08h11.765"/><circle cx="4.443" cy="12" r="1.331" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.123 12h11.765"/><circle cx="4.443" cy="18.919" r="1.331" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.123 18.92h11.765"/></g>`), g.Group(children),
	)
}

func TextBulletListSquare(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="7.877" cy="8.25" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M11.062 8.25h5.31"/><circle cx="7.877" cy="12" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M11.062 12h5.31"/><circle cx="7.877" cy="15.75" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M11.062 15.75h5.31"/><rect width="16.5" height="16.5" x="3.75" y="3.75" stroke="currentColor" stroke-width="1.5" rx="4"/></g>`), g.Group(children),
	)
}

func TextBulletListSquareAdd(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><circle cx="7.877" cy="8.25" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M11.062 8.25h5.31"/><circle cx="7.877" cy="12" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M11.062 12h1.31"/><circle cx="7.877" cy="15.75" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M20.25 10.25v-2.5a4 4 0 0 0-4-4h-8.5a4 4 0 0 0-4 4v8.5a4 4 0 0 0 4 4h2.5"/><path fill="currentColor" fill-rule="evenodd" d="M17.5 23a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11m0-8.993a.5.5 0 0 1 .5.5V17h2.493a.5.5 0 1 1 0 1H18v2.493a.5.5 0 1 1-1 0V18h-2.493a.5.5 0 1 1 0-1H17v-2.493a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func TextCaseLowercase(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m9.81 12.919l-.074-.015c-3.049-.578-6.676-.49-6.676 2.555c0 .854.491 1.658 1.206 2.095c.71.434 1.604.446 2.44.446c1.15 0 2.41-.65 2.914-1.662q.146-.29.19-.657v-1.167m0-1.595v-.609c0-.762-.09-1.768-.547-2.384c-.497-.67-1.38-1.426-2.828-1.426c-2.393 0-3.166 1.654-3.166 1.654M9.81 12.92v1.595m0 3.383v-3.383M13.56 6.5V18m8-4a4 4 0 1 1-8 0a4 4 0 0 1 8 0"/>`), g.Group(children),
	)
}

func TextCaseTitle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m1.88 18l1.83-4.404m0 0h6.59m-6.59 0L6.515 6.85c.193-.465.787-.465.98 0l2.805 6.747m0 0L12.13 18m2.66-11.5V18m7.59-3.795a3.795 3.795 0 1 1-7.59 0a3.795 3.795 0 0 1 7.59 0"/>`), g.Group(children),
	)
}

func TextCaseUppercase(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m1.01 18l1.92-4.404m0 0h6.91m-6.91 0L5.87 6.85a.553.553 0 0 1 1.028 0l2.941 6.747m0 0L11.76 18m3-5.948h4.046a2.765 2.765 0 0 0 2.754-2.776A2.765 2.765 0 0 0 18.806 6.5H14.76zm0 0h5c1.657 0 3 1.331 3 2.974c0 1.642-1.343 2.974-3 2.974h-5z"/>`), g.Group(children),
	)
}

func TextClearFormatting(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M14.77 22.252h6.476m-9.69-7.733l5.925 5.925m-6.195-.808l1.078 1.078c.849.848 1.273 1.272 1.762 1.431c.43.14.894.14 1.324 0c.49-.159.913-.583 1.762-1.431l3.502-3.502c.848-.849 1.272-1.273 1.431-1.762c.14-.43.14-.894 0-1.324c-.159-.49-.583-.913-1.431-1.762l-1.078-1.078c-.849-.848-1.273-1.272-1.762-1.431a2.14 2.14 0 0 0-1.324 0c-.49.159-.913.583-1.762 1.431l-3.502 3.502c-.848.849-1.272 1.273-1.431 1.762c-.14.43-.14.894 0 1.324c.159.49.583.913 1.431 1.762M2.75 12.427L4.328 8.72m0 0h5.68m-5.68 0l2.417-5.678a.45.45 0 0 1 .845 0l2.418 5.678m0 0l1.137 2.673M14.09 2.75v5.739l.112-.096a3.195 3.195 0 0 1 6.202.158"/>`), g.Group(children),
	)
}

func TextCollapse(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.125 4.5h14.5m-14.5 15h14.5m-7.5-10h7.5m-7.5 5h7.5"/><path fill="currentColor" fill-rule="evenodd" d="M6.875 16.5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9m-2-5a.5.5 0 1 0 0 1h4a.5.5 0 0 0 0-1z" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func TextColor(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="m7.154 13.088l1.73-3.959m0 0h6.231m-6.23 0l2.652-6.065a.5.5 0 0 1 .926 0l2.652 6.065m0 0l1.731 3.96"/><rect width="15.5" height="4.353" x="4.25" y="16.897" rx="1.5"/></g>`), g.Group(children),
	)
}

func TextColorAccent(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.25 18.397a1.5 1.5 0 0 1 1.5-1.5h12.5a1.5 1.5 0 0 1 1.5 1.5v1.353a1.5 1.5 0 0 1-1.5 1.5H5.75a1.5 1.5 0 0 1-1.5-1.5z"/>`), g.Group(children),
	)
}

func TextEditStyle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.25 14.25l1.875-4.403m0 0h6.75m-6.75 0l2.873-6.748a.536.536 0 0 1 1.004 0l2.873 6.748m0 0l.575 1.349m.886 9.49a2.5 2.5 0 0 0 1.219-.673l5.454-5.45a2.526 2.526 0 1 0-3.57-3.573l-5.453 5.452c-.335.336-.569.76-.674 1.222l-.536 2.354a1.007 1.007 0 0 0 1.206 1.206z"/>`), g.Group(children),
	)
}

func TextEffects(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M14.453 12.928H9.547L12 7.63zM12 2.75a2.5 2.5 0 0 1 2.269 1.45l6.25 13.5a2.5 2.5 0 0 1-4.538 2.1l-1.56-3.372H9.58L8.019 19.8a2.5 2.5 0 0 1-4.538-2.1m5.14-2.772h6.758l1.964 4.242a1 1 0 1 0 1.814-.84l-6.25-13.5a1 1 0 0 0-1.815 0l-6.25 13.5a1 1 0 0 0 1.815.84zM12 2.75A2.5 2.5 0 0 0 9.731 4.2zM9.731 4.2l-6.25 13.5z"/><path d="M14.269 4.2a2.5 2.5 0 0 0-4.538 0l-6.25 13.5a2.5 2.5 0 0 0 4.538 2.1l1.56-3.372h4.841l1.561 3.372a2.5 2.5 0 0 0 4.538-2.1zM8.62 14.927h6.758l1.964 4.242a1 1 0 1 0 1.814-.84l-6.25-13.5a1 1 0 0 0-1.815 0l-6.25 13.5a1 1 0 0 0 1.815.84z"/></g>`), g.Group(children),
	)
}

func TextExpand(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M7.125 4.5h14.5m-14.5 15h14.5m-7.5-10h7.5m-7.5 5h7.5"/><path fill="currentColor" fill-rule="evenodd" d="M6.875 16.5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9m0-7a.5.5 0 0 1 .5.5v1.5h1.5a.5.5 0 1 1 0 1h-1.5V14a.5.5 0 1 1-1 0v-1.5h-1.5a.5.5 0 0 1 0-1h1.5V10a.5.5 0 0 1 .5-.5" clip-rule="evenodd"/></g>`), g.Group(children),
	)
}

func TextFont(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m20.08 21.25l-2.133-4.842m0 0H10.27m7.677 0l-3.268-7.42a.617.617 0 0 0-1.142 0l-3.267 7.42m0 0L8.137 21.25m-1.174 0h2.812m8.663 0h2.812M4.423 8.82L2.75 12.584M4.423 8.82h6.021L7.881 3.05a.485.485 0 0 0-.895 0z"/>`), g.Group(children),
	)
}

func TextFontSize(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.75 17.75l1.875-4.404m0 0h6.75m-6.75 0L7.498 6.6a.536.536 0 0 1 1.004 0l2.873 6.747m0 0l1.875 4.404m0-.001l1.429-3.277m0 0h5.142m-5.142 0l2.188-5.022a.412.412 0 0 1 .765 0l2.19 5.022m0 0l1.428 3.277"/>`), g.Group(children),
	)
}

func TextFootnote(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.75 17.248l1.44-3.384m0 0h5.188m-5.187 0l2.208-5.186a.412.412 0 0 1 .771 0l2.208 5.186m0 0l1.441 3.384m2.398-8.838v8.838m5.833-2.916a2.916 2.916 0 1 1-5.833 0a2.916 2.916 0 0 1 5.833 0m0-7.707l2.2-1.375v5.5"/>`), g.Group(children),
	)
}

func TextHighlightColor(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="16.5" height="16.5" x="3.75" y="3.75" rx="4"/><path d="m8.25 16l1.34-3.063m0 0h4.82m-4.82 0l2.051-4.694a.386.386 0 0 1 .718 0l2.052 4.694m0 0L15.75 16"/></g>`), g.Group(children),
	)
}

func TextHighlightColorAccent(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.75 3.75a4 4 0 0 0-4 4v8.5a4 4 0 0 0 4 4h8.5a4 4 0 0 0 4-4v-8.5a4 4 0 0 0-4-4zm2.986 8.437L12 9.295l1.264 2.892zm-.002 2.5h2.532l.88 2.014a1.75 1.75 0 0 0 3.207-1.402l-3.391-7.757c-.753-1.723-3.171-1.723-3.924 0l-3.391 7.757a1.75 1.75 0 0 0 3.206 1.402zm2.312-6.745l3.391 7.758a.75.75 0 1 1-1.374.6l-1.143-2.613h-3.84L8.937 16.3a.75.75 0 1 1-1.374-.6l3.391-7.758a1.136 1.136 0 0 1 2.092 0" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TextIndentDecrease(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M8 17.502h12M8 6.498h12m-9 5.504h9"/><path stroke-linejoin="round" d="m5.81 9.195l-2.807 2.807l2.807 2.807"/></g>`), g.Group(children),
	)
}

func TextIndentIncrease(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="M8 17.502h12M8 6.498h12m-9 5.504h9"/><path stroke-linejoin="round" d="m4 9.195l2.807 2.807L4 14.809"/></g>`), g.Group(children),
	)
}

func TextItalic(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M9.25 4.75h9m-13.5 14.5h9m-4.75 0l5.263-14.5"/>`), g.Group(children),
	)
}

func TextLarge(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m5.5 19.207l2.321-5.325m0 0h8.358m-8.358 0l3.558-8.16a.669.669 0 0 1 1.242 0l3.558 8.16m0 0l2.321 5.325"/>`), g.Group(children),
	)
}

func TextLetterSpacing(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7.6 16.937l1.571-3.781m0 0h5.658m-5.658 0l2.408-5.793c.166-.4.676-.4.842 0l2.408 5.793m0 0l1.571 3.78"/><path d="M21.25 20.25V3.75m-18.5 16.5V3.75"/></g>`), g.Group(children),
	)
}

func TextLineHeight(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m2.75 18.345l1.992 2.037c.235.24.547.361.858.361m2.85-2.398l-1.992 2.037c-.235.24-.547.361-.858.361M2.75 5.61l2.002-2.002c.234-.234.54-.351.848-.351M8.45 5.61l-2-2.002a1.2 1.2 0 0 0-.849-.351m0 17.486V3.257"/><path d="M11.55 4.25h9.7m-9.7 15.5h9.7"/><path stroke-linejoin="round" d="m13.12 15.594l1.171-2.752m0 0h4.219m-4.219 0l1.796-4.218a.335.335 0 0 1 .627 0l1.796 4.217m0 0l1.172 2.753"/></g>`), g.Group(children),
	)
}

func TextLineSpacing(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m2.75 18.345l1.992 2.037c.235.24.547.361.858.361m2.85-2.398l-1.992 2.037c-.235.24-.547.361-.858.361M2.75 5.61l2.002-2.002c.234-.234.54-.351.848-.351M8.45 5.61l-2-2.002a1.2 1.2 0 0 0-.849-.351m0 0v17.486"/><path d="M11.55 4.25h9.7m-9.7 5.17h9.7m-9.7 5.17h9.7m-9.7 5.17h6.7"/></g>`), g.Group(children),
	)
}

func TextNumberList(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M9.75 5.247h10M9.75 12h10m-10 6.753h10M3.694 4.123L5.494 3v4.493m-2.223 3.582A1.32 1.32 0 0 1 4.465 9.76a1.325 1.325 0 0 1 1.428 1.057a1.32 1.32 0 0 1-.363 1.192l-1.994 1.973v.265H5.86m-2.265 2.26h2.344l-2.103 1.927h1.037c.589 0 1.066.477 1.066 1.065s-.172 1.242-1.241 1.483c-.587.132-1.448-.482-1.448-.964"/>`), g.Group(children),
	)
}

func TextPositionBottom(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 2.75v13.379m-6.287-5.666l5.226 5.226c.293.293.677.44 1.061.44m6.287-5.666l-5.226 5.226c-.293.293-.677.44-1.061.44M2.75 21.25h18.5"/>`), g.Group(children),
	)
}

func TextPositionMiddle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 2.75v6.086M9.081 6.33l2.212 2.212a1 1 0 0 0 .707.293m2.919-2.505l-2.212 2.212a1 1 0 0 1-.707.293m0 12.415v-6.086M9.081 17.67l2.212-2.212a1 1 0 0 1 .707-.293m2.919 2.505l-2.212-2.212a1 1 0 0 0-.707-.293M2.75 12h18.5"/>`), g.Group(children),
	)
}

func TextPositionTop(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 21.25V7.871m-6.287 5.666l5.226-5.226c.293-.293.677-.44 1.061-.44m6.287 5.666l-5.226-5.226A1.5 1.5 0 0 0 12 7.87M2.75 2.75h18.5"/>`), g.Group(children),
	)
}

func TextSmall(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m6.5 18.138l1.964-4.507m0 0h7.072m-7.072 0l3.01-6.904a.566.566 0 0 1 1.052 0l3.01 6.905m0 0l1.964 4.506"/>`), g.Group(children),
	)
}

func TextStrikethrough(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M17.059 8.113s-.71-3.363-5.861-3.363c-1.923 0-4.42 1.601-4.42 3.838c0 .493.087.92.259 1.292q.194.418.526.75m-.526 6.365S8.382 20 12.364 20c2.217 0 4.859-1.078 4.859-3.544c0-1.571-.918-2.63-2.408-3.331m-10.065 0h10.065m4.435 0h-4.435"/>`), g.Group(children),
	)
}

func TextSubscript(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m3.5 19.25l12-14.5m-12 0l8.586 10.375l2.04 2.465m2.051-1.928a1.912 1.912 0 1 1 3.263 1.352l-2.881 2.854v.382h3.359"/>`), g.Group(children),
	)
}

func TextSuperscript(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m4 4.75l12 14.5m-12 0L14.626 6.41m2.051-1.748a1.912 1.912 0 1 1 3.263 1.352l-2.88 2.854v.382h3.359"/>`), g.Group(children),
	)
}

func TextTypography(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m2.25 18.25l1.875-4.404m0 0h6.75m-6.75 0L6.998 7.1a.536.536 0 0 1 1.004 0l2.873 6.747m0 0l1.875 4.404m9-4.948l-.07-.014c-2.88-.557-6.308-.472-6.308 2.462c0 .824.464 1.598 1.14 2.02c.67.417 1.515.429 2.305.429c1.086 0 2.277-.626 2.754-1.602q.137-.279.179-.633V14.84m0-1.537v-.587c0-.734-.085-1.704-.517-2.297c-.47-.646-1.303-1.375-2.672-1.375c-2.261 0-2.991 1.595-2.991 1.595m6.18 2.664v1.537m0 3.26v-3.26"/>`), g.Group(children),
	)
}

func TextUnderline(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M6.5 4.75v6.5a5.5 5.5 0 0 0 5.5 5.5v0a5.5 5.5 0 0 0 5.5-5.5v-6.5m-11 15h11"/>`), g.Group(children),
	)
}

func ThumbsDown(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="4.2" height="13.296" x="21.25" y="16.799" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="1.5" transform="rotate(180 21.25 16.8)"/><path fill="currentColor" d="M17.05 11.447h.75a.75.75 0 0 0-.75-.75zm0 3.028l.627.412a.75.75 0 0 0 .123-.412zm0-3.028h-.75c0 .415.336.75.75.75zm-12.285-6.1l.723.196zm-1.958 7.187l.723.197zm6.507 3.157h.75a.75.75 0 0 0-.75-.75zM13.02 20.6l-.626-.412zm3.28-9.152v3.028h1.5v-3.028zm.75.75v-1.5zm-.75-5.694v4.944h1.5V6.503zm-2.25-2.25a2.25 2.25 0 0 1 2.25 2.25h1.5a3.75 3.75 0 0 0-3.75-3.75zm-6.873 0h6.873v-1.5H7.177zm-1.689 1.29a1.75 1.75 0 0 1 1.689-1.29v-1.5A3.25 3.25 0 0 0 4.04 5.15zM3.53 12.731l1.958-7.188l-1.447-.394l-1.958 7.188zm1.689 2.21a1.75 1.75 0 0 1-1.689-2.21l-1.447-.394a3.25 3.25 0 0 0 3.136 4.104zm4.095 0H5.219v1.5h4.095zm.75 1.972V15.69h-1.5v1.222zm0 0h-1.5zm0 2.587v-2.587h-1.5V19.5zm1.25 1.25c-.69 0-1.25-.56-1.25-1.25h-1.5a2.75 2.75 0 0 0 2.75 2.75zm.036 0h-.036v1.5h.036zm1.044-.563a1.25 1.25 0 0 1-1.044.563v1.5a2.75 2.75 0 0 0 2.297-1.238zm4.03-6.125l-4.03 6.125l1.253.825l4.03-6.125z"/></g>`), g.Group(children),
	)
}

func ThumbsUp(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><rect width="4.2" height="13.296" x="3.25" y="7.201" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" rx="1.5"/><path fill="currentColor" d="M7.45 12.553H6.7c0 .414.336.75.75.75zm0-3.027l-.627-.413a.75.75 0 0 0-.123.413zm0 3.027h.75a.75.75 0 0 0-.75-.75zm12.285 6.1l-.723-.196zm1.958-7.187l-.723-.197zm-6.507-3.157h-.75c0 .414.336.75.75.75zM11.48 3.4l.626.412zM8.2 12.553V9.526H6.7v3.027zm-.75-.75v1.5zm.75 5.694v-4.944H6.7v4.944zm2.25 2.25a2.25 2.25 0 0 1-2.25-2.25H6.7a3.75 3.75 0 0 0 3.75 3.75zm6.873 0H10.45v1.5h6.873zm1.689-1.29a1.75 1.75 0 0 1-1.689 1.29v1.5a3.25 3.25 0 0 0 3.136-2.396zm1.958-7.188l-1.958 7.188l1.447.394l1.958-7.188zm-1.689-2.21a1.75 1.75 0 0 1 1.689 2.21l1.447.394a3.25 3.25 0 0 0-3.136-4.104zm-4.095 0h4.095v-1.5h-4.095zm-.75-1.972V8.31h1.5V7.087zm0 0h1.5zm0-2.587v2.587h1.5V4.5zm-1.25-1.25c.69 0 1.25.56 1.25 1.25h1.5a2.75 2.75 0 0 0-2.75-2.75zm-.036 0h.036v-1.5h-.036zm-1.044.563a1.25 1.25 0 0 1 1.044-.563v-1.5a2.75 2.75 0 0 0-2.297 1.238zm-4.03 6.125l4.03-6.125l-1.253-.825l-4.03 6.125z"/></g>`), g.Group(children),
	)
}

func Tiktok(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.139 9.891l-.287.693zM15.8 15.5h-.75zm-1.011 3.333l-.624-.416zm-2.693 2.21l-.287-.693zm-3.467.342l.147-.736zm-3.072-1.642l.53-.53zM3.915 16.67l.736-.147zm.342-3.467l.693.287zm2.21-2.693l-.417-.623zm2 2.994l-.417-.624zm-.884 1.077l.693.287zm-.137 1.386l-.735.146zm.657 1.23l-.53.53zm2.615.52l.287.692zm1.078-.885l.623.417zM12.2 15.5h-.75zm3.935-11.316l.693-.287zm2.381 2.381l-.287.693zm1.285.317l.068-.747zm0 3.608l.037-.75zM9.4 13.133l-.124-.74zm6.399-3.21h-.75zM12.2 2.9h.75zm3.618 0l-.747.068zM9.4 9.512l-.05-.748zm7.452 1.07a8.8 8.8 0 0 0 2.91.656l.076-1.498a7.3 7.3 0 0 1-2.413-.543zm-.767-.36q.374.199.767.36l.573-1.385a7 7 0 0 1-.634-.3zm-1.035-.3V15.5h1.5V9.924zm.362 9.327a6.75 6.75 0 0 0 1.138-3.75h-1.5a5.25 5.25 0 0 1-.885 2.917zm-3.029 2.486a6.75 6.75 0 0 0 3.03-2.486l-1.248-.833a5.25 5.25 0 0 1-2.356 1.933zm-3.9.384a6.75 6.75 0 0 0 3.9-.384l-.574-1.386a5.25 5.25 0 0 1-3.033.3zm-3.456-1.847a6.75 6.75 0 0 0 3.456 1.847l.293-1.47a5.25 5.25 0 0 1-2.688-1.438zM3.18 16.817a6.75 6.75 0 0 0 1.847 3.456l1.06-1.06a5.25 5.25 0 0 1-1.436-2.689zm.384-3.9a6.75 6.75 0 0 0-.384 3.9l1.47-.293a5.25 5.25 0 0 1 .3-3.033zm2.486-3.03a6.75 6.75 0 0 0-2.486 3.03l1.386.574a5.25 5.25 0 0 1 1.933-2.356zm3.3-1.122a6.75 6.75 0 0 0-3.3 1.123l.833 1.247a5.25 5.25 0 0 1 2.567-.873zm1.2 3.935V9.9h-1.5v2.8zm-1.667 1.428a1.65 1.65 0 0 1 .643-.255l-.249-1.48a3.15 3.15 0 0 0-1.227.488zm-.607.74a1.65 1.65 0 0 1 .607-.74l-.833-1.247a3.15 3.15 0 0 0-1.16 1.413zm-.094.954a1.65 1.65 0 0 1 .094-.953l-1.386-.575a3.15 3.15 0 0 0-.18 1.82zm.451.845a1.65 1.65 0 0 1-.451-.845l-1.471.293a3.15 3.15 0 0 0 .862 1.612zm.845.451a1.65 1.65 0 0 1-.845-.451l-1.06 1.06c.44.44 1.001.74 1.612.863zm.953-.094a1.65 1.65 0 0 1-.953.094l-.293 1.472a3.15 3.15 0 0 0 1.82-.18zm.74-.607a1.65 1.65 0 0 1-.74.607l.574 1.386a3.15 3.15 0 0 0 1.414-1.16zm.279-.917c0 .326-.097.645-.278.917l1.247.833a3.15 3.15 0 0 0 .531-1.75zm1.5 0V2.9h-1.5v12.6zM12.6 3.25h2.8v-1.5h-2.8zm4.228.647a3.7 3.7 0 0 1-.263-1.066l-1.494.137c.047.516.172 1.022.371 1.503zm.791 1.184a3.7 3.7 0 0 1-.791-1.184l-1.386.574a5.2 5.2 0 0 0 1.116 1.67zm1.184.791a3.7 3.7 0 0 1-1.184-.791l-1.06 1.06a5.2 5.2 0 0 0 1.67 1.117zm1.066.263a3.7 3.7 0 0 1-1.066-.263l-.574 1.386c.48.199.987.324 1.503.37zM20.95 10.1V7.3h-1.5v2.8zm-1.218-2.471a.32.32 0 0 1-.282-.329h1.5c0-.65-.518-1.114-1.081-1.165zm.03 3.61A1.135 1.135 0 0 0 20.95 10.1h-1.5c0-.183.155-.37.388-.36zM9.05 12.7c0-.102.044-.183.089-.23a.25.25 0 0 1 .138-.076l.25 1.479c.501-.085 1.023-.513 1.023-1.173zm7.74-3.802c-.729-.388-1.74.084-1.74 1.026h1.5c0 .143-.086.25-.17.297a.31.31 0 0 1-.295.001zM12.95 2.9a.35.35 0 0 1-.35.35v-1.5a1.15 1.15 0 0 0-1.15 1.15zm2.45.35a.32.32 0 0 1-.329-.282l1.494-.137c-.051-.563-.514-1.081-1.165-1.081zm-5.95 7.012a.37.37 0 0 1-.4-.362h1.5c0-.62-.513-1.18-1.2-1.135z"/>`), g.Group(children),
	)
}

func Timer(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.022 8.862A7.716 7.716 0 1 0 5.74 18.206a7.716 7.716 0 0 0 12.282-9.344m0 0l1.813-1.813M9.28 2.75h5.2m-2.6 7.215v3.57"/>`), g.Group(children),
	)
}

func Toolbox(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 10a3 3 0 0 1 3-3h12.5a3 3 0 0 1 3 3v7.5a3 3 0 0 1-3 3H5.75a3 3 0 0 1-3-3zM8 5.5a2 2 0 0 1 2-2h4a2 2 0 0 1 2 2V7H8zm.5 5.5v3m7-3v3M2.75 12.5h18.5"/>`), g.Group(children),
	)
}

func Triangle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M5.732 20.5c-2.29 0-3.723-2.498-2.581-4.5L9.419 5.006c1.144-2.008 4.018-2.008 5.163 0L20.849 16c1.142 2.002-.291 4.5-2.581 4.5z"/>`), g.Group(children),
	)
}

func Tune(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 12h10.5m0 0a2.5 2.5 0 0 0 5 0m-5 0a2.5 2.5 0 0 1 5 0m0 0h3M7.75 5.25h13.5m-13.5 0a2.5 2.5 0 1 1-5 0a2.5 2.5 0 0 1 5 0m-5 13.5h3m0 0a2.5 2.5 0 0 0 5 0m-5 0a2.5 2.5 0 0 1 5 0m0 0h10.5"/>`), g.Group(children),
	)
}

func Tv(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="18.5" height="13" x="2.75" y="4" rx="3"/><path d="M7 20h10"/></g>`), g.Group(children),
	)
}

func VehicleCar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 11.75a2 2 0 0 1 2-2h12.5a2 2 0 0 1 2 2v6H3.75z"/><circle cx="7" cy="13" r="1" fill="currentColor"/><circle cx="17" cy="13" r="1" fill="currentColor"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M10 14.75h4M6.271 5.567a2 2 0 0 1 1.88-1.317h7.698a2 2 0 0 1 1.88 1.317L19.25 9.75H4.75zM3.75 17.75h3.438v1.7a1.3 1.3 0 0 1-1.3 1.3H5.05a1.3 1.3 0 0 1-1.3-1.3zm13.063 0h3.437v1.7a1.3 1.3 0 0 1-1.3 1.3h-.837a1.3 1.3 0 0 1-1.3-1.3z"/></g>`), g.Group(children),
	)
}

func Video(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><rect width="13.5" height="12" x="2.75" y="6" rx="3.5"/><path d="m16.25 9.74l3.554-1.77a1 1 0 0 1 1.446.895v6.268a1 1 0 0 1-1.447.895l-3.553-1.773z"/></g>`), g.Group(children),
	)
}

func VideoClip(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><rect width="18.5" height="15.5" x="2.75" y="4.25" stroke-linecap="round" stroke-linejoin="round" rx="3.5"/><path d="M8.83 13.29c0 1.098 0 1.646.23 1.964c.201.277.51.456.85.492c.391.041.867-.232 1.818-.779l2.244-1.29c.956-.55 1.435-.825 1.595-1.185a1.2 1.2 0 0 0 0-.984c-.16-.36-.639-.635-1.595-1.184l-2.244-1.291c-.951-.547-1.427-.82-1.817-.779c-.34.036-.65.215-.85.492c-.231.318-.231.866-.231 1.963z"/></g>`), g.Group(children),
	)
}

func VisualStudioCode(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M15.695 16.318L10.333 12m5.362 4.318V7.682m0 8.636v2.18c0 .68-.27 1.332-.75 1.813M10.332 12l5.362-4.318M10.333 12L7.828 9.983m7.867-2.3v-2.18c0-.68-.27-1.333-.75-1.814M7.827 14.017l-3.516 2.831a.61.61 0 0 1-.813-.044l-1.071-1.076a.61.61 0 0 1 .022-.881L5.598 12m2.23 2.017l6.456 5.84q.304.273.66.454m-7.116-6.294L5.598 12m0 0L2.45 9.153a.61.61 0 0 1-.022-.881l1.071-1.076c.22-.22.57-.24.813-.044l3.516 2.83m0 0l6.456-5.839q.304-.273.66-.454m0 16.622a3.04 3.04 0 0 0 2.517.113l2.334-.937c.577-.232.955-.791.955-1.413V5.926c0-.622-.378-1.181-.955-1.413l-2.334-.937a3.04 3.04 0 0 0-2.517.113"/>`), g.Group(children),
	)
}

func Volume(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M12.253 19.4L7.99 15.782a1 1 0 0 0-.647-.238H4.75a2 2 0 0 1-2-2v-3.086a2 2 0 0 1 2-2h2.594a1 1 0 0 0 .647-.238l4.262-3.62a1 1 0 0 1 1.647.762V18.64a1 1 0 0 1-1.647.762Z"/><path stroke-linecap="round" d="M16.664 8.542c.48.35.88.854 1.158 1.462c.277.607.424 1.295.424 1.996c0 .7-.147 1.39-.424 1.996c-.278.607-.677 1.112-1.158 1.462M18.7 6.424c.775.565 1.42 1.378 1.867 2.357c.447.978.683 2.089.683 3.219s-.236 2.24-.683 3.22c-.448.978-1.092 1.791-1.867 2.356"/></g>`), g.Group(children),
	)
}

func VolumeMute(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M12.253 19.4L7.99 15.782a1 1 0 0 0-.647-.238H4.75a2 2 0 0 1-2-2v-3.086a2 2 0 0 1 2-2h2.594a1 1 0 0 0 .647-.238l4.262-3.62a1 1 0 0 1 1.647.762V18.64a1 1 0 0 1-1.647.762Z"/><path stroke-linecap="round" d="m16.53 9.64l4.72 4.72m0-4.72l-4.72 4.72"/></g>`), g.Group(children),
	)
}

func VolumeOne(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M12.253 19.4L7.99 15.782a1 1 0 0 0-.647-.238H4.75a2 2 0 0 1-2-2v-3.086a2 2 0 0 1 2-2h2.594a1 1 0 0 0 .647-.238l4.262-3.62a1 1 0 0 1 1.647.762V18.64a1 1 0 0 1-1.647.762Z"/><path stroke-linecap="round" d="M16.664 8.542c.48.35.88.854 1.158 1.462c.277.607.423 1.295.423 1.996c0 .7-.146 1.39-.423 1.996c-.278.607-.677 1.112-1.158 1.462"/></g>`), g.Group(children),
	)
}

func VolumeZero(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path d="M12.253 19.4L7.99 15.782a1 1 0 0 0-.647-.238H4.75a2 2 0 0 1-2-2v-3.086a2 2 0 0 1 2-2h2.594a1 1 0 0 0 .647-.238l4.262-3.62a1 1 0 0 1 1.647.762V18.64a1 1 0 0 1-1.647.762Z"/><path stroke-linecap="round" d="M16.664 8.542c.48.35.88.854 1.158 1.462c.277.607.424 1.295.424 1.996c0 .7-.147 1.39-.424 1.996c-.278.607-.677 1.112-1.158 1.462M18.7 6.424c.775.565 1.42 1.378 1.867 2.357c.447.978.683 2.089.683 3.219s-.236 2.24-.683 3.22c-.448.978-1.092 1.791-1.867 2.356"/></g>`), g.Group(children),
	)
}

func Watch(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M18.314 12a6.3 6.3 0 0 0-2.623-5.123v-2.92c0-.805-.652-1.457-1.457-1.457H9.766c-.805 0-1.457.652-1.457 1.457v2.92A6.3 6.3 0 0 0 5.686 12m12.628 0a6.3 6.3 0 0 1-2.623 5.123v2.92c0 .805-.652 1.457-1.457 1.457H9.766a1.457 1.457 0 0 1-1.457-1.457v-2.92A6.3 6.3 0 0 1 5.686 12m12.628 0a6.314 6.314 0 1 1-12.628 0m12.628 0a6.314 6.314 0 1 0-12.628 0"/>`), g.Group(children),
	)
}

func WeatherCloudy(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.465 13.796a3.43 3.43 0 0 1-.715-2.101a3.457 3.457 0 0 1 3.114-3.437a.31.31 0 0 0 .276-.305A3.46 3.46 0 0 1 9.603 4.5c.88 0 1.682.327 2.293.866m4.665 6.083a3.46 3.46 0 0 0-3.493-3.453a3.46 3.46 0 0 0-3.432 3.453c0 .157-.12.29-.276.305a3.46 3.46 0 0 0-3.114 3.437a3.46 3.46 0 0 0 3.462 3.453h8.08a3.46 3.46 0 0 0 3.462-3.453a3.46 3.46 0 0 0-3.462-3.454h-.95a.283.283 0 0 1-.277-.288"/>`), g.Group(children),
	)
}

func WiFi(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 8.988A12.068 12.068 0 0 1 21.25 9"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M5.64 11.964a8.297 8.297 0 0 1 12.72.01m-9.805 3.029a4.495 4.495 0 0 1 6.89.005"/><circle cx="12" cy="17.878" r="1.445" fill="currentColor"/></g>`), g.Group(children),
	)
}

func Window(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M3.75 7.75a4 4 0 0 1 4-4h8.5a4 4 0 0 1 4 4v8.5a4 4 0 0 1-4 4h-8.5a4 4 0 0 1-4-4zm0 .5h16.5"/>`), g.Group(children),
	)
}

func WindowMultiple(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 6.75a4 4 0 0 1 4-4h7.5a4 4 0 0 1 4 4v7.5a4 4 0 0 1-4 4h-7.5a4 4 0 0 1-4-4z"/><path fill="currentColor" d="M21.952 9.486c-.05-.655-.152-1.21-.388-1.725A4.8 4.8 0 0 0 20.5 6.286v5.664c0 1.55 0 2.671-.067 3.556c-.066.876-.193 1.46-.415 1.942a5.25 5.25 0 0 1-2.57 2.57c-.483.222-1.066.35-1.942.415c-.885.067-2.006.067-3.556.067H6.286a4.8 4.8 0 0 0 1.475 1.064c.514.236 1.07.339 1.725.388c.64.048 1.433.048 2.434.048h.061c1.513 0 2.69 0 3.637-.071c.96-.072 1.745-.22 2.458-.549a6.75 6.75 0 0 0 3.304-3.304c.329-.713.477-1.497.549-2.458C22 14.67 22 13.494 22 11.98v-.061c0-1 0-1.793-.048-2.434"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M2.75 6.977h15.5"/></g>`), g.Group(children),
	)
}

func WindowMultipleVar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M2.75 6.75a4 4 0 0 1 4-4h7.5a4 4 0 0 1 4 4v7.5a4 4 0 0 1-4 4h-7.5a4 4 0 0 1-4-4z"/><path d="M21.25 6.75v7.5a7 7 0 0 1-7 7h-7.5m-4-14.273h15.5"/></g>`), g.Group(children),
	)
}

func Wrench(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M10.691 4.562a6.19 6.19 0 0 1 6.545-1.42c.378.141.45.62.165.906l-2.787 2.787a1.037 1.037 0 0 0 0 1.467l1.084 1.084a1.037 1.037 0 0 0 1.467 0L19.953 6.6c.285-.285.764-.212.905.165a6.187 6.187 0 0 1-7.696 8.058c-.396-.128-.84-.054-1.134.24L6.481 20.61a2.186 2.186 0 1 1-3.09-3.09l5.547-5.548c.294-.294.368-.738.24-1.134a6.19 6.19 0 0 1 1.513-6.276Z"/>`), g.Group(children),
	)
}

func XTwitter(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="m13.081 10.712l-4.786-6.71a.6.6 0 0 0-.489-.252H5.28a.6.6 0 0 0-.488.948l6.127 8.59m2.162-2.576l6.127 8.59a.6.6 0 0 1-.488.948h-2.526a.6.6 0 0 1-.489-.252l-4.786-6.71m2.162-2.576l5.842-6.962m-8.004 9.538L5.077 20.25"/>`), g.Group(children),
	)
}

func Youtube(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="1.5"><path stroke-linecap="round" d="M2.45 11.419c0-3.017.3-4.526 1.237-5.463s2.446-.937 5.463-.937h5.7c3.017 0 4.525 0 5.463.937s1.237 2.446 1.237 5.463v1.162c0 3.017-.3 4.526-1.237 5.463s-2.446.937-5.463.937h-5.7c-3.017 0-4.526 0-5.463-.937S2.45 15.598 2.45 12.581z"/><path d="m14.686 11.491l-4.268-2.667a.6.6 0 0 0-.918.509v5.335a.6.6 0 0 0 .918.508l4.268-2.667a.6.6 0 0 0 0-1.018Z"/></g>`), g.Group(children),
	)
}

func YoutubeShorts(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none"><path stroke="currentColor" stroke-linejoin="round" stroke-width="1.5" d="M12.834 3.186a3.627 3.627 0 0 1 3.627 6.282l-.74.426a3.626 3.626 0 0 1 1.935 6.766l-7.02 4.053a3.626 3.626 0 1 1-3.627-6.28l.739-.428A3.627 3.627 0 0 1 5.814 7.24z"/><path fill="currentColor" d="M13.992 11.016L11.2 9.271c-.74-.463-1.7.07-1.7.942v3.49c0 .873.96 1.405 1.7.943l2.792-1.746a1.11 1.11 0 0 0 0-1.884"/></g>`), g.Group(children),
	)
}

func ZoomIn(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.56 10.56h6m-3 3v-6m4.993 7.993a7.06 7.06 0 1 0-9.985-9.985a7.06 7.06 0 0 0 9.985 9.985m0 0L20 20"/><path d="M15.553 5.568a7.06 7.06 0 1 1-9.985 9.985a7.06 7.06 0 0 1 9.985-9.985"/></g>`), g.Group(children),
	)
}

func ZoomOut(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M7.56 10.56h6m1.993 4.993a7.06 7.06 0 1 0-9.985-9.985a7.06 7.06 0 0 0 9.985 9.985m0 0L20 20"/><path d="M15.553 5.568a7.06 7.06 0 1 1-9.985 9.985a7.06 7.06 0 0 1 9.985-9.985"/></g>`), g.Group(children),
	)
}
