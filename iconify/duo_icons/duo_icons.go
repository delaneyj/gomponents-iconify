package duo_icons

import (
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

const IconifyVersion = "1.1.4"

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 24 24")
)

func AddCircle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M12 7a1 1 0 0 0-.993.883L11 8v3H8a1 1 0 0 0-.117 1.993L8 13h3v3a1 1 0 0 0 1.993.117L13 16v-3h3a1 1 0 0 0 .117-1.993L16 11h-3V8a1 1 0 0 0-1-1" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Airplay(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.514 16.903a.6.6 0 0 1 .976 0l2.766 3.868a.6.6 0 0 1-.488.949H9.237a.6.6 0 0 1-.488-.949z" class="duoicon-primary-layer"/><path fill="currentColor" d="M20 4a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2h-3.71l-3.151-4.407a1.4 1.4 0 0 0-2.278 0L7.711 19H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func AlertOctagon(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.314 2a2 2 0 0 1 1.414.586l4.686 4.686A2 2 0 0 1 22 8.686v6.628a2 2 0 0 1-.586 1.414l-4.686 4.686a2 2 0 0 1-1.414.586H8.686a2 2 0 0 1-1.414-.586l-4.686-4.686A2 2 0 0 1 2 15.314V8.686a2 2 0 0 1 .586-1.414l4.686-4.686A2 2 0 0 1 8.686 2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M12 6a1 1 0 0 0-.993.883L11 7v6a1 1 0 0 0 1.993.117L13 13V7a1 1 0 0 0-1-1m0 9a1 1 0 1 0 0 2a1 1 0 0 0 0-2" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func AlertTriangle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m13.299 3.148l8.634 14.954a1.5 1.5 0 0 1-1.299 2.25H3.366a1.5 1.5 0 0 1-1.299-2.25l8.634-14.954c.577-1 2.02-1 2.598 0" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M12 8a1 1 0 0 0-.993.883L11 9v4a1 1 0 0 0 1.993.117L13 13V9a1 1 0 0 0-1-1m0 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func AlignBottom(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5 3a2 2 0 0 0-2 2v14a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M15 7a2 2 0 0 0-2 2v10a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2V9a2 2 0 0 0-2-2z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func AlignCenter(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 17.5a1.5 1.5 0 0 1 .144 2.993L17 20.5H7a1.5 1.5 0 0 1-.144-2.993L7 17.5z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M20 2.5a1.5 1.5 0 0 1 .144 2.993L20 5.5H4a1.5 1.5 0 0 1-.144-2.993L4 2.5z" class="duoicon-primary-layer"/><path fill="currentColor" d="M17 7.5a1.5 1.5 0 0 1 .144 2.993L17 10.5H7a1.5 1.5 0 0 1-.144-2.993L7 7.5z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M20 12.5a1.5 1.5 0 0 1 0 3H4a1.5 1.5 0 0 1 0-3z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Android(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20 14v5a3 3 0 0 1-3 3H7a3 3 0 0 1-3-3v-5zm-5-7a1 1 0 1 0 0 2a1 1 0 0 0 0-2M9 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2" class="duoicon-primary-layer"/><path fill="currentColor" fill-rule="evenodd" d="M12 3c1.33 0 2.584.324 3.687.899l.606-.606a1 1 0 1 1 1.414 1.414l-.35.35A7.98 7.98 0 0 1 20 11v1H4v-1a7.98 7.98 0 0 1 2.644-5.942l-.351-.35a1 1 0 1 1 1.414-1.415l.606.606A8 8 0 0 1 12 3" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func App(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.75 13A2.25 2.25 0 0 1 11 15.25v3.5A2.25 2.25 0 0 1 8.75 21h-3.5A2.25 2.25 0 0 1 3 18.75v-3.5A2.25 2.25 0 0 1 5.25 13zm10-10A2.25 2.25 0 0 1 21 5.25v3.5A2.25 2.25 0 0 1 18.75 11h-3.5A2.25 2.25 0 0 1 13 8.75v-3.5A2.25 2.25 0 0 1 15.25 3z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M8.75 3A2.25 2.25 0 0 1 11 5.25v3.5A2.25 2.25 0 0 1 8.75 11h-3.5A2.25 2.25 0 0 1 3 8.75v-3.5A2.25 2.25 0 0 1 5.25 3z" class="duoicon-primary-layer"/><path fill="currentColor" d="M18.75 13A2.25 2.25 0 0 1 21 15.25v3.5A2.25 2.25 0 0 1 18.75 21h-3.5A2.25 2.25 0 0 1 13 18.75v-3.5A2.25 2.25 0 0 1 15.25 13z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func AppDots(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 11.5c-3.464 0-5.63-3.75-3.897-6.75A4.5 4.5 0 0 1 7 2.5c3.464 0 5.63 3.75 3.897 6.75A4.5 4.5 0 0 1 7 11.5" class="duoicon-primary-layer"/><path fill="currentColor" d="M17 21.5c-3.464 0-5.63-3.75-3.897-6.75A4.5 4.5 0 0 1 17 12.5c3.464 0 5.63 3.75 3.897 6.75A4.5 4.5 0 0 1 17 21.5m0-10c-3.464 0-5.63-3.75-3.897-6.75A4.5 4.5 0 0 1 17 2.5c3.464 0 5.63 3.75 3.897 6.75A4.5 4.5 0 0 1 17 11.5m-10 10c-3.464 0-5.63-3.75-3.897-6.75A4.5 4.5 0 0 1 7 12.5c3.464 0 5.63 3.75 3.897 6.75A4.5 4.5 0 0 1 7 21.5" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Apple(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m13.064 6.685l.745-.306c.605-.24 1.387-.485 2.31-.33c1.891.318 3.195 1.339 3.972 2.693c.3.522.058 1.21-.502 1.429c-1.793.701-2.154 3.081-.65 4.283c.233.187.499.33.783.423c.518.17.81.745.64 1.263c-.442 1.342-1.078 2.581-1.831 3.581c-.744.988-1.652 1.808-2.663 2.209c-.66.26-1.368.163-2.045-.005l-.402-.107l-.597-.173c-.271-.079-.55-.147-.824-.147c-.275 0-.553.068-.824.147l-.597.173l-.402.107c-.677.168-1.386.266-2.045.005c-1.273-.504-2.396-1.68-3.245-3.067a13.4 13.4 0 0 1-1.784-4.986c-.227-1.554-.104-3.299.615-4.775c.74-1.521 2.096-2.705 4.163-3.053c.84-.141 1.562.048 2.14.265l.331.13l.584.241c.4.157.715.249 1.064.249c.348 0 .664-.092 1.064-.249" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M11.768 2.768c.976-.977 2.475-1.061 2.828-.707c.354.353.27 1.852-.707 2.828c-.976.976-2.475 1.06-2.828.707c-.354-.353-.27-1.852.707-2.828" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Approved(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.586 2.1a2 2 0 0 1 2.7-.116l.128.117L15.314 4H18a2 2 0 0 1 1.994 1.85L20 6v2.686l1.9 1.9a2 2 0 0 1 .116 2.701l-.117.127l-1.9 1.9V18a2 2 0 0 1-1.85 1.995L18 20h-2.685l-1.9 1.9a2 2 0 0 1-2.701.116l-.127-.116l-1.9-1.9H6a2 2 0 0 1-1.995-1.85L4 18v-2.686l-1.9-1.9a2 2 0 0 1-.116-2.701l.116-.127l1.9-1.9V6a2 2 0 0 1 1.85-1.994L6 4h2.686z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="m15.079 8.983l-4.244 4.244l-1.768-1.768a1 1 0 1 0-1.414 1.415l2.404 2.404a1.1 1.1 0 0 0 1.556 0l4.88-4.881a1 1 0 0 0-1.414-1.414" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Appstore(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M11.873 5.511a1 1 0 0 0-1.797.873l.051.105l.727 1.297l-1.807 3.226a1 1 0 0 0 1.683 1.075l.063-.098l3.08-5.5a1 1 0 0 0-1.683-1.076l-.063.098L12 5.74l-.127-.227zm2.52 4.5a1 1 0 0 0-1.797.873l.052.105l3.08 5.5a1 1 0 0 0 1.796-.872l-.052-.106l-.286-.51h.514a1 1 0 0 0 .117-1.994L17.7 13h-1.634zM12.5 13H6.3a1 1 0 1 0 0 2h.514l-.287.511a1.001 1.001 0 0 0 1.746.978L9.106 15H12.5a1 1 0 1 0 0-2" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Award(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2c6.158 0 10.007 6.667 6.928 12A8 8 0 0 1 17 16.245v4.61a1.1 1.1 0 0 1-1.486 1.03L12 20.569l-3.514 1.318A1.1 1.1 0 0 1 7 20.856v-4.61C2.192 12.398 3.352 4.788 9.089 2.548A8 8 0 0 1 12 2" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M12 6c3.079 0 5.004 3.333 3.464 6A4 4 0 0 1 12 14c-3.079 0-5.004-3.333-3.464-6A4 4 0 0 1 12 6" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func BabyCarriage(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.746 2.609c.764-.296 1.566.093 1.877.773L12.643 10H18V8.5A2.5 2.5 0 0 1 20.5 6h.5a1 1 0 1 1 0 2h-.5a.5.5 0 0 0-.5.5V11a9 9 0 0 1-2.489 6.213c1.76.778 2.018 3.17.464 4.305s-3.755.163-3.961-1.751A3 3 0 0 1 14 19.5v-.015a9 9 0 0 1-6 0v.015c0 1.925-2.084 3.127-3.75 2.164c-1.667-.962-1.666-3.368 0-4.33q.117-.067.239-.121C.063 12.574 1.769 4.927 7.746 2.609" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M8.012 4.669A7 7 0 0 0 4.072 10h6.372L8.012 4.67z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Bank(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m12.67 2.217l8.5 4.75A1.5 1.5 0 0 1 22 8.31v1.44c0 .69-.56 1.25-1.25 1.25H20v8h1a1 1 0 1 1 0 2H3a1 1 0 1 1 0-2h1v-8h-.75C2.56 11 2 10.44 2 9.75V8.31c0-.522.27-1.002.706-1.274l8.623-4.819c.422-.211.92-.211 1.342 0z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M12 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2m5 5H7v8h2v-6h2v6h2v-6h2v6h2z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Battery(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 6a2 2 0 0 1 2 2v1a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2v1a2 2 0 0 1-2 2h-5.101l2.664-4.441c.986-1.642.197-3.65-1.405-4.325l.415-.69c.478-.799.543-1.73.255-2.544z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M11.142 6.486a1 1 0 0 1 1.77.925l-.055.104L10.768 11h2.215c.746 0 1.221.773.92 1.427l-.054.103l-2.99 4.985a1 1 0 0 1-1.77-.926l.055-.104L11.233 13H9.019a1.01 1.01 0 0 1-.92-1.427l.054-.103l2.99-4.984z" class="duoicon-primary-layer"/><path fill="currentColor" d="m9.101 6l-2.664 4.441c-.953 1.587-.247 3.516 1.247 4.253l.158.072l-.414.69a3 3 0 0 0-.317 2.354l.06.19H4a2 2 0 0 1-1.995-1.85L2 16V8a2 2 0 0 1 1.85-1.994L4 6z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Bell(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.042 19.003h5.916c-.385 2.277-3.09 3.283-4.87 1.811a3 3 0 0 1-1.046-1.811" class="duoicon-primary-layer"/><path fill="currentColor" d="M12 2.003a7.5 7.5 0 0 1 7.5 7.5v4l1.418 3.16A.95.95 0 0 1 20.052 18h-16.1a.95.95 0 0 1-.867-1.338l1.415-3.16V9.49l.005-.25A7.5 7.5 0 0 1 12 2.004z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func BellBadge(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.5 9q.508-.002.973-.137q.027.317.027.64v4l1.418 3.16A.95.95 0 0 1 20.052 18h-16.1a.95.95 0 0 1-.867-1.338l1.415-3.16V9.49l.005-.25c.201-5.749 6.533-9.14 11.43-6.123c-1.834 1.973-.845 5.193 1.781 5.795A3.5 3.5 0 0 0 18.5 9" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M18.5 8c1.925 0 3.127-2.083 2.165-3.75A2.5 2.5 0 0 0 18.5 3c-1.925 0-3.127 2.083-2.165 3.75A2.5 2.5 0 0 0 18.5 8m-3.542 11.003c-.385 2.277-3.09 3.283-4.87 1.811a3 3 0 0 1-1.046-1.811z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Book(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 2a2 2 0 0 0-2 2v1a1 1 0 1 0 0 2v2a1 1 0 1 0 0 2v2a1 1 0 1 0 0 2v2a1 1 0 1 0 0 2v1a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V4a2 2 0 0 0-2-2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M8.5 6A1.5 1.5 0 0 0 7 7.5v1A1.5 1.5 0 0 0 8.5 10h7A1.5 1.5 0 0 0 17 8.5v-1A1.5 1.5 0 0 0 15.5 6z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func BookThree(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M4 5a3 3 0 0 1 3-3h11a2 2 0 0 1 2 2v12.99c0 .168-.038.322-.113.472l-.545 1.09a1 1 0 0 0 0 .895l.543 1.088A1 1 0 0 1 19 22H7a3 3 0 0 1-3-3z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M10 7a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2zM7 18h10.408a3 3 0 0 0 0 2H7a1 1 0 1 1 0-2" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func BookTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.255 3.667A1.01 1.01 0 0 1 4.022 2H16.5c3.464 0 5.629 3.75 3.897 6.75A4.5 4.5 0 0 1 16.5 11H4.022a1.01 1.01 0 0 1-.767-1.667l.754-.88a3 3 0 0 0 0-3.905l-.754-.88z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M3 16.5A4.5 4.5 0 0 1 7.5 12h12.478a1.01 1.01 0 0 1 .767 1.667l-.755.88a3 3 0 0 0 0 3.905l.755.88A1.01 1.01 0 0 1 19.978 21H7.5A4.5 4.5 0 0 1 3 16.5" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Bookmark(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 2a3 3 0 0 1 3 3v16a1 1 0 0 1-1.625.78l-1.875-1.5l-1.875 1.5a1 1 0 0 1-1.332-.073L12 20.414l-1.293 1.293a1 1 0 0 1-1.332.074L7.5 20.28l-1.875 1.5A1 1 0 0 1 4 21V5a3 3 0 0 1 3-3z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M15 8H9a1 1 0 0 0-.117 1.993L9 10h6a1 1 0 0 0 .117-1.993zm-3 4H9a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Box(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21 10v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-9z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M20 3a2 2 0 0 1 2 2v3H2V5a2 2 0 0 1 2-2zm-6 10h-4a1 1 0 1 0 0 2h4a1 1 0 1 0 0-2" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func BoxTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m20.765 7.982l.022.19l.007.194v7.268a2.5 2.5 0 0 1-1.099 2.07l-.15.095l-6.295 3.634l-.124.067l-.126.06v-8.99z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="m13.25 2.567l6.294 3.634q.076.044.148.092L12 10.838L4.308 6.293a3 3 0 0 1 .148-.092l6.294-3.634a2.5 2.5 0 0 1 2.5 0" class="duoicon-primary-layer"/><path fill="currentColor" d="M3.235 7.982L11 12.571v8.988a2 2 0 0 1-.25-.126l-6.294-3.634a2.5 2.5 0 0 1-1.25-2.165V8.366q0-.195.03-.384z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Bread(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.194 2.885c2.3-.299 4.458.21 5.645 1.397c.79.791 1.18 1.711 1.16 2.72c-.017.961-.405 1.894-.938 2.764c-.762 1.245-1.952 2.563-3.27 3.915l-.729.74l-1.616 1.617l-.691.682c-1.378 1.346-2.722 2.566-3.989 3.341c-.87.533-1.803.92-2.764.939c-1.009.02-1.929-.37-2.72-1.161c-1.187-1.187-1.696-3.345-1.397-5.645c.307-2.36 1.471-5.035 3.872-7.437c2.402-2.4 5.078-3.565 7.437-3.872" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M13.111 5.149c-.618.178-1.259.43-1.907.768c.52.847 1.299 1.452 2.308 1.789a1 1 0 1 0 .633-1.897c-.434-.145-.774-.36-1.034-.66m-4.42 2.53a14 14 0 0 0-1.406 1.468c.194.251.406.489.621.704c.585.584 1.333 1.145 2.07 1.39a1 1 0 0 0 .742-1.854l-.109-.043c-.323-.107-.812-.43-1.288-.907a5 5 0 0 1-.63-.758m-3 3.986a9.8 9.8 0 0 0-.742 2.283q.631.541 1.492.83a1 1 0 1 0 .633-1.898c-.667-.222-1.112-.612-1.383-1.216z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Bridge(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 18a1 1 0 0 1 .117 1.993L15 20h-3a1 1 0 0 1-.117-1.993L12 18z" class="duoicon-primary-layer"/><path fill="currentColor" d="M15 3a1 1 0 0 1 .993.883L16 4v.454c1.59.378 2.85.937 3.775 1.459l.225.13V6a1 1 0 0 1 1.993-.117L22 6v6.9a1.1 1.1 0 0 1-.98 1.094L20.9 14h-4.855c-.539 0-.982-.407-1.039-.93L15 12.954V12.5c.002-2.309-2.496-3.755-4.497-2.602a3 3 0 0 0-1.498 2.426L9 12.5v.455c0 .539-.407.982-.93 1.039L7.954 14H3.1a1.1 1.1 0 0 1-1.094-.98L2 12.9V6a1 1 0 0 1 1.993-.117L4 6v.042a14.3 14.3 0 0 1 3.61-1.49L8 4.455V4a1 1 0 0 1 1.993-.117L10 4v.109a18.3 18.3 0 0 1 3.522-.046l.478.046V4a1 1 0 0 1 1-1" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M13 15a1 1 0 1 1 0 2H9a1 1 0 1 1 0-2z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Briefcase(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 3a3 3 0 0 1 3 3h3a2 2 0 0 1 2 2v11a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V8a2 2 0 0 1 2-2h3a3 3 0 0 1 3-3z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M14 5h-4a1 1 0 0 0-.993.883L9 6h6a1 1 0 0 0-.883-.993zm5 5H5a1 1 0 0 0-.117 1.993L5 12h6v1a1 1 0 0 0 1.993.117L13 13v-1h6a1 1 0 0 0 .117-1.993z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Brush(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 15v5a2 2 0 0 1-2 2h-1v-3a1 1 0 1 0-2 0v3h-4v-3a1 1 0 1 0-2 0v3H7v-3a1 1 0 1 0-2 0v3H4a2 2 0 0 1-2-2v-5z" class="duoicon-primary-layer"/><path fill="currentColor" d="M13 2a2 2 0 0 1 2 2v4a1 1 0 0 0 1 1h4a2 2 0 0 1 2 2v2H2v-2a2 2 0 0 1 2-2h4a1 1 0 0 0 1-1V4a2 2 0 0 1 2-2z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func BrushTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m20.485 10.586l1.06 1.06a2.5 2.5 0 0 1 .123 3.405l-.122.13l-6.364 6.365a2.5 2.5 0 0 1-3.405.122l-.13-.122l-1.061-1.06z" class="duoicon-primary-layer"/><path fill="currentColor" d="M2.808 2.808c.885-.886 2.01-1.33 3.184-1.203c1.161.125 2.225.792 3.056 1.846c.888 1.127 1.7 2.732 2.336 4.174c.311.704 1.16.874 1.676.475l.093-.081l.615-.615a2.5 2.5 0 0 1 3.405-.122l.13.122l1.768 1.768l-9.9 9.899l-1.767-1.768a2.5 2.5 0 0 1-.122-3.405l.122-.13l.615-.615c.5-.5.354-1.439-.394-1.769c-1.442-.636-3.047-1.448-4.174-2.336c-1.053-.831-1.72-1.895-1.846-3.056c-.127-1.174.317-2.3 1.203-3.184" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Bug(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.67 5.5a5 5 0 0 1 8.66 0L17.2 7H6.8z" class="duoicon-primary-layer"/><path fill="currentColor" fill-rule="evenodd" d="M3.553 6.106a1 1 0 0 1 1.341.447c.147.293.5.674.973.99C6.353 7.867 6.781 8 7 8h10c.219 0 .647-.133 1.133-.457c.474-.316.826-.697.973-.99a1 1 0 1 1 1.788.894c-.353.707-1 1.326-1.652 1.76a5.5 5.5 0 0 1-.966.516c.297.731.503 1.496.616 2.277H21a1 1 0 1 1 0 2h-2.012a10 10 0 0 1-.74 3.327c.572.33.963.86 1.209 1.35c.349.725.534 1.518.543 2.323a1 1 0 1 1-2 0c0-.374-.101-.966-.332-1.428c-.13-.26-.26-.409-.385-.49c-1.056 1.486-2.539 2.54-4.283 2.835V13a1 1 0 1 0-2 0v8.917c-1.744-.295-3.227-1.35-4.283-2.834c-.126.08-.255.23-.385.49c-.21.447-.323.933-.332 1.427a1 1 0 1 1-2 0a5.5 5.5 0 0 1 .543-2.322c.246-.492.637-1.02 1.209-1.35A10 10 0 0 1 5.012 14H3a1 1 0 1 1 0-2h2.108c.113-.781.32-1.546.616-2.277a5.5 5.5 0 0 1-.966-.516c-.651-.434-1.3-1.053-1.652-1.76a1 1 0 0 1 .447-1.341" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Building(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13 3a2 2 0 0 1 1.995 1.85L15 5v14h1V9.5a.5.5 0 0 1 .41-.492L16.5 9H18a2 2 0 0 1 1.995 1.85L20 11v8h1a1 1 0 0 1 .117 1.993L21 21H3a1 1 0 0 1-.117-1.993L3 19h1V5a2 2 0 0 1 1.85-1.995L6 3z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M11 7H8v2h3zm0 4H8v2h3zm0 4H8v2h3z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Bus(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M11.022 3H18a3 3 0 0 1 3 3a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1v7c0 .715-.381 1.375-1 1.732v.768a1.5 1.5 0 0 1-3 0V19H7v.5a1.5 1.5 0 0 1-3 0v-.768A2 2 0 0 1 3 17v-7a1 1 0 0 1-1-1V7a1 1 0 0 1 1-1a3 3 0 0 1 3-3z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M15 15a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2h-1a1 1 0 0 1-1-1m-9 0a1 1 0 0 1 1-1h1a1 1 0 1 1 0 2H7a1 1 0 0 1-1-1m5.618-10h2.764l-1.276 2.553a1 1 0 1 0 1.788.894L16.618 5H18a1 1 0 0 1 1 1v5H5V6a1 1 0 0 1 1-1h3.382L8.106 7.553a1 1 0 0 0 1.788.894z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Cake(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.707 15.707a.414.414 0 0 1 .586 0a2.41 2.41 0 0 0 2.707.491V20a1 1 0 1 1 0 2H3a1 1 0 1 1 0-2v-3.802c.89.405 1.975.241 2.707-.49a.414.414 0 0 1 .586 0a2.414 2.414 0 0 0 3.414 0a.414.414 0 0 1 .586 0a2.414 2.414 0 0 0 3.414 0a.414.414 0 0 1 .586 0a2.414 2.414 0 0 0 3.414 0zM16.5 2c-.319.638-.028 1.05.225 1.41c.144.203.275.39.275.59a1 1 0 1 1-2 0c0-.552.5-1.5 1.5-2m-8 0c-.319.638-.028 1.05.225 1.41c.144.203.275.39.275.59a1 1 0 1 1-2 0c0-.552.5-1.5 1.5-2m4 0c-.319.638-.028 1.05.225 1.41c.144.203.275.39.275.59a1 1 0 1 1-2 0c0-.552.5-1.5 1.5-2" class="duoicon-primary-layer"/><path fill="currentColor" d="M16 6a1 1 0 0 1 1 1v2h1a3 3 0 0 1 3 3v1.586l-.707.707a.414.414 0 0 1-.586 0a2.414 2.414 0 0 0-3.414 0a.414.414 0 0 1-.586 0a2.414 2.414 0 0 0-3.414 0a.414.414 0 0 1-.586 0a2.414 2.414 0 0 0-3.414 0a.414.414 0 0 1-.586 0a2.414 2.414 0 0 0-3.414 0a.414.414 0 0 1-.586 0L3 13.586V12a3 3 0 0 1 3-3h1V7a1 1 0 1 1 2 0v2h2V7a1 1 0 1 1 2 0v2h2V7a1 1 0 0 1 1-1" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Calendar(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21 12v7a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-7z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M16 3a1 1 0 0 1 1 1v1h2a2 2 0 0 1 2 2v3H3V7a2 2 0 0 1 2-2h2V4a1 1 0 1 1 2 0v1h6V4a1 1 0 0 1 1-1" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Camera(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.793 3c.346 0 .682.12.95.34l.11.1L17.415 5H20a2 2 0 0 1 1.995 1.85L22 7v12a2 2 0 0 1-1.85 1.995L20 21H4a2 2 0 0 1-1.995-1.85L2 19V7a2 2 0 0 1 1.85-1.995L4 5h2.586l1.56-1.56c.245-.246.568-.399.913-.433L9.207 3z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M12 7.5c-3.849 0-6.255 4.167-4.33 7.5A5 5 0 0 0 12 17.5c3.849 0 6.255-4.167 4.33-7.5A5 5 0 0 0 12 7.5" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func CameraSquare(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 3a3 3 0 0 1 3 3v12a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3V6a3 3 0 0 1 3-3z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M19 6h-1a1 1 0 0 0-.117 1.993L18 8h1a1 1 0 0 0 .117-1.993zm-7 1c-3.849 0-6.255 4.167-4.33 7.5A5 5 0 0 0 12 17c3.849 0 6.255-4.167 4.33-7.5A5 5 0 0 0 12 7" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Campground(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.176 3.434a1 1 0 1 1 1.648 1.133l-1.61 2.342L21.526 19H22a1 1 0 1 1 0 2h-4.634l-.044-.07l-4.492-6.487a1.01 1.01 0 0 0-1.66 0L6.678 20.93l-.044.07H2a1 1 0 1 1 0-2h.474l8.313-12.092l-1.611-2.342a1 1 0 0 1 1.648-1.133L12 5.144z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M12 16.757L14.938 21H9.062z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Candle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.058 2.08a1 1 0 0 0-1.382.766c-.027.228-.1.449-.216.647c-.1.165-.25.334-.514.58c-.52.482-.837 1.059-.922 1.681c-.074.572.04 1.153.326 1.655C9.842 8.296 10.84 9 11.903 9C13.285 9 15 8.034 15 6.194c0-.943-.422-1.829-.946-2.512c-.528-.687-1.243-1.279-1.996-1.601z" class="duoicon-primary-layer"/><path fill="currentColor" fill-rule="evenodd" d="M9 10a2 2 0 0 0-2 2v8a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-8a2 2 0 0 0-2-2z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Car(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M5.553 5.658A3 3 0 0 1 8.236 4h7.528a3 3 0 0 1 2.683 1.658l1.386 2.771q.366-.15.72-.324a1 1 0 0 1 .894 1.79q-.36.16-.725.312l.961 1.923c.209.417.317.877.317 1.343V16a3 3 0 0 1-1 2.236V19.5a1.5 1.5 0 0 1-3 0V19H6v.5a1.5 1.5 0 0 1-3 0v-1.264c-.614-.55-1-1.348-1-2.236v-2.528c0-.466.109-.925.317-1.341l.953-1.908q-.362-.152-.715-.327a1.01 1.01 0 0 1-.45-1.343a1 1 0 0 1 1.342-.448q.355.175.72.324z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M7.342 6.553A1 1 0 0 1 8.236 6h7.528c.379 0 .725.214.894.553l1.27 2.538C16.38 9.555 14.294 10 12 10s-4.38-.445-5.927-.91zM16.5 16a1.5 1.5 0 1 0 0-3a1.5 1.5 0 0 0 0 3M9 14.5a1.5 1.5 0 1 1-3 0a1.5 1.5 0 0 1 3 0" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Certificate(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 11c2.309 0 3.753 2.5 2.598 4.5a3 3 0 0 1-.598.736v4.955a.5.5 0 0 1-.724.447L19 21l-1.276.638a.5.5 0 0 1-.724-.447v-4.955c-1.721-1.54-1.13-4.365 1.064-5.086c.302-.099.618-.15.936-.15m-7-2H6a1 1 0 0 0-.117 1.993L6 11h6a1 1 0 0 0 .117-1.993zm-4 4H6a1 1 0 1 0 0 2h2a1 1 0 1 0 0-2" class="duoicon-primary-layer"/><path fill="currentColor" fill-rule="evenodd" d="M20 4a2 2 0 0 1 2 2v4c-3.079-2.309-7.504-.419-7.964 3.402A5 5 0 0 0 15 17v3H4a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func ChartPie(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M12 4v8h8a8 8 0 0 0-8-8" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func CheckCircle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="m15.535 8.381l-4.95 4.95l-2.12-2.121a1 1 0 1 0-1.415 1.414l2.758 2.758a1.1 1.1 0 0 0 1.556 0l5.586-5.586a1 1 0 1 0-1.415-1.415" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Chip(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 2a1 1 0 0 1 1 1v1h1a3 3 0 0 1 3 3v1h1a1 1 0 1 1 0 2h-1v4h1a1 1 0 1 1 0 2h-1v1a3 3 0 0 1-3 3h-1v1a1 1 0 1 1-2 0v-1h-4v1a1 1 0 1 1-2 0v-1H7a3 3 0 0 1-3-3v-1H3a1 1 0 1 1 0-2h1v-4H3a1 1 0 1 1 0-2h1V7a3 3 0 0 1 3-3h1V3a1 1 0 1 1 2 0v1h4V3a1 1 0 0 1 1-1" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M14 8h-4a2 2 0 0 0-2 2v4a2 2 0 0 0 2 2h4a2 2 0 0 0 2-2v-4a2 2 0 0 0-2-2" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Clapperboard(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 5a2 2 0 0 1 2-2h16a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M13 8h3l1-3h-3zM8 8h3l1-3H9zM4 8h2l1-3H4zm16-3h-1l-1 3h2z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Clipboard(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 3v1a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V3h1a2 2 0 0 1 2 2v11a6 6 0 0 1-6 6H6a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M14 2a1 1 0 0 1 .117 1.993L14 4h-4a1 1 0 0 1-.117-1.993L10 2zm1 8H9a1 1 0 0 0-.117 1.993L9 12h6a1 1 0 1 0 0-2m-3 4H9a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Clock(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4c6.928 0 11.258 7.5 7.794 13.5A9 9 0 0 1 12 22C5.072 22 .742 14.5 4.206 8.5A9 9 0 0 1 12 4" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M7.366 2.971A1 1 0 0 1 7 4.337a10.1 10.1 0 0 0-2.729 2.316a1 1 0 1 1-1.544-1.27a12 12 0 0 1 3.271-2.777a1 1 0 0 1 1.367.365zM18 2.606a12 12 0 0 1 3.272 2.776a1 1 0 0 1-1.544 1.27a10 10 0 0 0-2.729-2.315a1 1 0 0 1 1.002-1.731zM12 8a1 1 0 0 0-.993.883L11 9v3.986c-.003.222.068.44.202.617l.09.104l2.106 2.105a1 1 0 0 0 1.498-1.32l-.084-.094L13 12.586V9a1 1 0 0 0-1-1" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func CloudLightning(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.132 13.004a1 1 0 0 1 1.788.888l-.052.104L11.723 16h2.105a1.1 1.1 0 0 1 1.006 1.545l-.051.1l-1.915 3.351a1 1 0 0 1-1.788-.888l.052-.104L12.277 18h-2.105a1.1 1.1 0 0 1-1.006-1.545l.051-.1z" class="duoicon-primary-layer"/><path fill="currentColor" d="M11.5 2a6.5 6.5 0 0 1 6.086 4.212c4.455 1.225 5.913 6.814 2.624 10.059a6 6 0 0 1-3.397 1.674c.449-1.555-.41-3.198-1.894-3.751c.53-2.249-1.573-4.228-3.785-3.563a3 3 0 0 0-1.739 1.381l-1.914 3.35A3.06 3.06 0 0 0 7.273 18H7c-3.849.003-6.257-4.163-4.335-7.497A5 5 0 0 1 5 8.416A6.5 6.5 0 0 1 11.5 2" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func CloudSnow(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.5 19a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m-4-6a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m4 2a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3" class="duoicon-primary-layer"/><path fill="currentColor" d="M11.5 2a6.5 6.5 0 0 1 6.086 4.212c4.455 1.223 5.916 6.811 2.629 10.058a6 6 0 0 1-2.439 1.462C18.637 15.443 16.945 13 14.5 13a1.52 1.52 0 0 1-1.199-.599c-1.615-2.157-4.959-1.757-6.019.72a3.5 3.5 0 0 0 .007 2.772c.167.388.167.828 0 1.216q-.184.422-.253.892H7c-3.849.003-6.257-4.163-4.335-7.497A5 5 0 0 1 5 8.417A6.5 6.5 0 0 1 11.5 2" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M10.5 17a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func CoinStack(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21 16.143V17.5c0 .814-.381 1.51-.91 2.057c-.523.542-1.233.984-2.032 1.334C16.456 21.591 14.314 22 12 22s-4.456-.408-6.058-1.109c-.799-.35-1.509-.792-2.032-1.334c-.485-.5-.845-1.128-.902-1.856L3 17.5v-1.357q.697.398 1.494.695c2.03.751 4.685 1.17 7.506 1.17s5.476-.419 7.506-1.17q.598-.222 1.139-.503zM12 3c2.314 0 4.456.408 6.058 1.109c.799.35 1.509.792 2.032 1.334c.485.5.845 1.128.902 1.856L21 7.5v.748a8.3 8.3 0 0 1-2.188 1.214c-1.755.65-4.164 1.047-6.812 1.047c-2.647 0-5.056-.397-6.812-1.047a8.3 8.3 0 0 1-1.905-1.006L3 8.248V7.5c0-.814.381-1.51.91-2.057c.523-.542 1.233-.984 2.032-1.334C7.544 3.409 9.686 3 12 3" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M3 10.643q.697.398 1.494.695c2.03.751 4.685 1.17 7.506 1.17s5.476-.419 7.506-1.17A10 10 0 0 0 21 10.643v3.105a8.3 8.3 0 0 1-2.188 1.214c-1.755.65-4.164 1.047-6.812 1.047c-2.647 0-5.056-.397-6.812-1.047A8.3 8.3 0 0 1 3 13.748z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Compass(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M16.243 7.757c-.354-.353-4.95.707-6.364 2.122c-1.414 1.414-2.475 6.01-2.122 6.364c.354.353 4.95-.707 6.364-2.122c1.415-1.414 2.475-6.01 2.122-6.364" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func ComputerCamera(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.26 16.877c.413.402.74.932.74 1.596c0 .753-.426 1.339-.9 1.745c-.48.408-1.114.731-1.81.983c-1.402.505-3.272.799-5.29.799s-3.888-.294-5.29-.8c-.696-.25-1.33-.574-1.809-.982c-.475-.406-.901-.992-.901-1.745c0-.663.327-1.194.74-1.596A9.98 9.98 0 0 0 12 20c2.859 0 5.438-1.2 7.26-3.123M12 6c-3.079 0-5.004 3.333-3.464 6A4 4 0 0 0 12 14c3.079 0 5.004-3.333 3.464-6A4 4 0 0 0 12 6" class="duoicon-primary-layer"/><path fill="currentColor" d="M12 2c6.158 0 10.007 6.667 6.928 12A8 8 0 0 1 12 18C5.842 18 1.993 11.333 5.072 6A8 8 0 0 1 12 2" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func ComputerCameraOff(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.74 16.877A9.98 9.98 0 0 0 12 20a9.96 9.96 0 0 0 4.797-1.224l.302-.172l.77.77a1.7 1.7 0 0 0 1.622.446c-.12.147-.254.28-.392.398c-.479.408-1.113.731-1.81.983c-1.401.505-3.271.799-5.289.799s-3.888-.294-5.29-.8c-.696-.25-1.33-.574-1.809-.982c-.475-.406-.901-.992-.901-1.745c0-.663.327-1.194.74-1.596M12 6c-.768 0-1.486.217-2.095.592l1.499 1.498c1.469-.46 2.885.843 2.549 2.346a2 2 0 0 1-.043.16l1.498 1.499c1.613-2.623-.219-6.008-3.297-6.093zm-7.132.373l3.159 3.16a4 4 0 0 0 4.44 4.44l3.16 3.16C10.14 19.929 3.684 15.736 4.006 9.586a8 8 0 0 1 .861-3.213z" class="duoicon-primary-layer"/><path fill="currentColor" d="M12 2c6.158 0 10.008 6.666 6.929 12a8 8 0 0 1-.468.718l-.186.244l1.503 1.503a1 1 0 0 1-1.32 1.497l-.094-.083L4.222 3.737a1 1 0 0 1 1.32-1.497l.094.083l1.402 1.402A7.96 7.96 0 0 1 12 2" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Confetti(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23 7a8.44 8.44 0 0 0-5 1.31c-.36-.41-.73-.82-1.12-1.21l-.29-.27l.14-.12a3.15 3.15 0 0 0 .9-3.49A3.9 3.9 0 0 0 14 1v2a2 2 0 0 1 1.76 1c.17.4 0 .84-.47 1.31l-.23.21a16.7 16.7 0 0 0-3.41-2.2c-2.53-1.14-3.83-.61-4.47 0a2.2 2.2 0 0 0-.46.68l-.18.53L5.1 8.87C6.24 11.71 9 16.76 15 18.94l5-1.66a1 1 0 0 0 .43-.31l.21-.18c1.43-1.44.51-4.21-1.41-6.9A6.63 6.63 0 0 1 23 9z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="m4.4 11l-2.23 6.7c-.693 2.118.882 4.295 3.11 4.3a3.2 3.2 0 0 0 1-.17l6.52-2.17A18.7 18.7 0 0 1 4.4 11m14.81 4.37h-.06c-.69.37-3.55-.57-6.79-3.81c-.34-.34-.66-.67-.95-1c-.1-.11-.19-.23-.29-.35l-.53-.64l-.28-.39c-.14-.19-.28-.38-.4-.56s-.16-.26-.24-.39s-.22-.34-.31-.51s-.13-.24-.19-.37s-.17-.28-.23-.42s-.09-.23-.14-.34s-.11-.27-.15-.4S8.6 6 8.58 5.9s-.06-.24-.08-.34a2 2 0 0 1 0-.24a1.2 1.2 0 0 1 0-.26l.11-.31c.17-.18.91-.23 2.23.37c.882.424 1.717.94 2.49 1.54c-.422.186-.87.3-1.33.34v2a6.4 6.4 0 0 0 3-.94l.49.46c.44.43.83.86 1.19 1.27A5.3 5.3 0 0 0 16 13.2l2-.39a3.2 3.2 0 0 1 0-1.14c1.29 1.97 1.53 3.39 1.21 3.7" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func CreditCard(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 10v7a3 3 0 0 1-3 3H5a3 3 0 0 1-3-3v-7z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M19 4a3 3 0 0 1 3 3v1H2V7a3 3 0 0 1 3-3zm-1 10h-3a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func CurrencyEuro(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M8.7 13c.104.856.357 1.716.786 2.458c.597 1.031 1.572 1.88 2.96 2.043c1.407.166 2.837-.51 3.7-1.615a1 1 0 0 0-.175-1.404c-.492-.384-1.093-.199-1.46.243c-.948 1.145-2.55 1.016-3.294-.27A4.4 4.4 0 0 1 10.72 13H15a1 1 0 1 0 0-2h-4.254c.1-.517.273-.998.51-1.393c.71-1.185 2.334-1.491 3.243-.34c.352.445.95.646 1.448.282a1 1 0 0 0 .213-1.398c-.845-1.156-2.098-1.721-3.531-1.599s-2.457.973-3.088 2.026c-.436.727-.703 1.572-.823 2.422H8a1 1 0 1 0 0 2z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Dashboard(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 11a2 2 0 0 1 1.995 1.85L21 13v6a2 2 0 0 1-1.85 1.995L19 21h-4a2 2 0 0 1-1.995-1.85L13 19v-6a2 2 0 0 1 1.85-1.995L15 11zm0-8a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M9 3a2 2 0 0 1 2 2v6a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2z" class="duoicon-primary-layer"/><path fill="currentColor" fill-rule="evenodd" d="M9 15a2 2 0 0 1 2 2v2a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2v-2a2 2 0 0 1 2-2z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Discount(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.405 2.897a4 4 0 0 1 5.02-.136l.17.136l.376.32c.274.234.605.389.96.45l.178.022l.493.04a4 4 0 0 1 3.648 3.468l.021.2l.04.494c.028.358.153.702.36.996l.11.142l.322.376a4 4 0 0 1 .136 5.02l-.136.17l-.321.376a2 2 0 0 0-.45.96l-.022.178l-.039.493a4 4 0 0 1-3.468 3.648l-.201.021l-.493.04a2 2 0 0 0-.996.36l-.142.111l-.377.32a4 4 0 0 1-5.02.137l-.169-.136l-.376-.321a2 2 0 0 0-.96-.45l-.178-.021l-.493-.04a4 4 0 0 1-3.648-3.468l-.021-.2l-.04-.494a2 2 0 0 0-.36-.996l-.111-.142l-.321-.377a4 4 0 0 1-.136-5.02l.136-.169l.32-.376c.234-.274.389-.605.45-.96l.022-.178l.04-.493A4 4 0 0 1 7.197 3.75l.2-.021l.494-.04c.358-.028.702-.153.996-.36l.142-.111l.376-.32z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M9.5 8a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m4.793.293l-6 6a1 1 0 1 0 1.414 1.414l6-6a1 1 0 0 0-1.414-1.414M14.5 13a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Disk(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M11.44 6.507a1 1 0 0 0-1.276-.61a7 7 0 0 0-3.73 3.1A1 1 0 1 0 8.166 10a5.02 5.02 0 0 1 2.665-2.216a1 1 0 0 0 .61-1.276zM12 10c-1.54 0-2.502 1.667-1.732 3c.357.619 1.017 1 1.732 1c1.54 0 2.502-1.667 1.732-3A2 2 0 0 0 12 10" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func File(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2v6.5a1.5 1.5 0 0 0 1.5 1.5H20v10a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4a2 2 0 0 1 2-2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M14 2.043c.379.08.726.269 1 .543L19.414 7c.274.274.463.621.543 1H14zm-2.346 9.656l-.894 1.535l-1.737.376a.4.4 0 0 0-.213.658l1.184 1.325l-.18 1.768a.4.4 0 0 0 .56.406L12 17.051l1.626.716a.4.4 0 0 0 .56-.406l-.18-1.768l1.184-1.325a.4.4 0 0 0-.213-.658l-1.737-.376l-.894-1.535a.4.4 0 0 0-.692 0" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Fire(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 12.9l-2.13 2.09c-.56.56-.87 1.29-.87 2.07C9 18.68 10.35 20 12 20s3-1.32 3-2.94c0-.78-.31-1.52-.87-2.07z" class="duoicon-primary-layer"/><path fill="currentColor" d="m16 6l-.44.55C14.38 8.02 12 7.19 12 5.3V2S4 6 4 13c0 2.92 1.56 5.47 3.89 6.86c-.56-.79-.89-1.76-.89-2.8c0-1.32.52-2.56 1.47-3.5L12 10.1l3.53 3.47c.95.93 1.47 2.17 1.47 3.5c0 1.02-.31 1.96-.85 2.75c1.89-1.15 3.29-3.06 3.71-5.3c.66-3.55-1.07-6.9-3.86-8.52" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func FolderOpen(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.82 6a2 2 0 0 1 1.972 2.329l-1.666 10A2 2 0 0 1 18.153 20H5.847a2 2 0 0 1-1.973-1.671l-1.666-10A2 2 0 0 1 4.18 6z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M18 3a1 1 0 1 1 0 2H6a1 1 0 1 1 0-2z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func FolderUpload(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 6h-8l-1.41-1.41C10.21 4.21 9.7 4 9.17 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V8c0-1.1-.9-2-2-2" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M13 13v3c0 .55-.45 1-1 1s-1-.45-1-1v-3H9.21c-.45 0-.67-.54-.35-.85l2.8-2.79c.2-.2.51-.19.71 0l2.79 2.79c.3.31.08.85-.36.85z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func GTranslate(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 22l-1-3H4q-.824 0-1.412-.587A1.93 1.93 0 0 1 2 17V4q0-.825.588-1.413A1.93 1.93 0 0 1 4 2h6l.875 3H20q.875 0 1.438.562T22 7v13q0 .825-.562 1.413Q20.875 22 20 22z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M13 21h7q.45 0 .725-.288A1 1 0 0 0 21 20V7q0-.45-.275-.725T20 6h-8.825l1.175 4.05h1.975V9h1.025v1.05H19v1.025h-1.275a8 8 0 0 1-.75 1.85A9.8 9.8 0 0 1 15.8 14.6l2.725 2.675L17.8 18l-2.7-2.7l-.9.925L15 19z" class="duoicon-primary-layer"/><path fill="currentColor" d="M15.1 13.825q.7-.825 1.063-1.575q.362-.75.487-1.175h-3.975l.3 1.05h1q.2.375.475.813t.65.887M13.85 15.1l.55-.525a16 16 0 0 1-.637-.825q-.288-.4-.563-.85zm-6.7-.5q1.725 0 2.838-1.112T11.1 10.6a5 5 0 0 0-.012-.363a1.7 1.7 0 0 0-.063-.337h-3.95v1.55H9.3a1.9 1.9 0 0 1-.763 1.087q-.562.388-1.362.388q-.975 0-1.675-.7T4.8 10.5q0-1.024.7-1.725a2.3 2.3 0 0 1 1.675-.7q.45 0 .85.162q.4.163.725.488L9.975 7.55Q9.451 7 8.713 6.7a4.1 4.1 0 0 0-1.563-.3q-1.675 0-2.862 1.187T3.1 10.5t1.188 2.912Q5.475 14.6 7.15 14.6" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func IdCard(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M10 9v2H8V9zm7 2h-3a1 1 0 0 0-.117 1.993L14 13h3a1 1 0 0 0 .117-1.993z" class="duoicon-primary-layer"/><path fill="currentColor" d="M10 7H8a2 2 0 0 0-1.995 1.85L6 9v2a2 2 0 0 0 1.85 1.995L8 13h2a2 2 0 0 0 1.995-1.85L12 11V9a2 2 0 0 0-1.85-1.995zm7 8H7a1 1 0 1 0 0 2h10a1 1 0 1 0 0-2" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Info(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 1.999c5.524 0 10.002 4.478 10.002 10.002c0 5.523-4.478 10.001-10.002 10.001S2 17.524 2 12.001C1.999 6.477 6.476 1.999 12 1.999" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M12.001 6.5a1.252 1.252 0 1 0 .002 2.503A1.252 1.252 0 0 0 12 6.5zm-.005 3.749a1 1 0 0 0-.992.885l-.007.116l.004 5.502l.006.117a1 1 0 0 0 1.987-.002L13 16.75l-.004-5.501l-.007-.117a1 1 0 0 0-.994-.882z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Lamp(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 20a1 1 0 0 1 .117 1.993L13 22h-2a1 1 0 0 1-.117-1.993L11 20zm.707-13.707a1 1 0 0 0-1.32-.083l-.094.083L10.3 8.286a1.01 1.01 0 0 0-.084 1.333l.084.095L11.586 11l-1.293 1.293a1 1 0 0 0 1.32 1.497l.094-.083l1.993-1.993c.36-.36.396-.931.084-1.333l-.084-.095L12.414 9l1.293-1.293a1 1 0 0 0 0-1.414" class="duoicon-primary-layer"/><path fill="currentColor" d="M12 2c4.41 0 8 3.543 8 7.933c0 3.006-1.522 5.196-2.78 6.494l-.284.283l-.27.252l-.252.22l-.33.27l-.328.244c-.241.17-.403.419-.55.678l-.205.364c-.238.41-.517.762-1.108.762h-3.786c-.59 0-.87-.351-1.108-.762l-.118-.208c-.172-.312-.348-.63-.637-.834l-.232-.171l-.199-.155l-.227-.188l-.252-.22l-.27-.252l-.285-.283C5.522 15.129 4 12.939 4 9.933C4 5.543 7.59 2 12 2" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func LampTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 19a1 1 0 0 1 1 1v1a1 1 0 1 1-2 0v-1a1 1 0 0 1 1-1m2.83-5H9.17c.77 2.178 3.609 2.706 5.11.95c.2-.233.362-.496.482-.778z" class="duoicon-primary-layer"/><path fill="currentColor" fill-rule="evenodd" d="M12 2a1 1 0 0 1 .993.883L13 3v1.1l8.175 8.176c.609.608.223 1.63-.6 1.718l-.114.006H16.9c-.773 3.772-5.34 5.293-8.22 2.737a5 5 0 0 1-1.532-2.525L7.1 14H3.54c-.861 0-1.31-.995-.791-1.639l.077-.085L11 4.1V3a1 1 0 0 1 1-1" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M7.05 16.95a1 1 0 0 1 0 1.414l-.707.707a1 1 0 0 1-1.414-1.414l.707-.707a1 1 0 0 1 1.414 0m9.9 0a1 1 0 0 1 1.32-.083l.094.083l.707.707a1 1 0 0 1-1.32 1.497l-.094-.083l-.707-.707a1 1 0 0 1 0-1.414" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Location(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.72 16.64a1 1 0 1 1 .56 1.92c-.5.146-.86.3-1.091.44c.238.143.614.303 1.136.452C8.48 19.782 10.133 20 12 20s3.52-.218 4.675-.548c.523-.149.898-.309 1.136-.452c-.23-.14-.59-.294-1.09-.44a1 1 0 0 1 .559-1.92c.668.195 1.28.445 1.75.766c.435.299.97.82.97 1.594c0 .783-.548 1.308-.99 1.607c-.478.322-1.103.573-1.786.768C15.846 21.77 14 22 12 22s-3.846-.23-5.224-.625c-.683-.195-1.308-.446-1.786-.768c-.442-.3-.99-.824-.99-1.607c0-.774.535-1.295.97-1.594c.47-.321 1.082-.571 1.75-.766M12 7.5c-1.54 0-2.502 1.667-1.732 3c.357.619 1.017 1 1.732 1c1.54 0 2.502-1.667 1.732-3A2 2 0 0 0 12 7.5" class="duoicon-primary-layer"/><path fill="currentColor" d="M12 2a7.5 7.5 0 0 1 7.5 7.5c0 2.568-1.4 4.656-2.85 6.14a16.4 16.4 0 0 1-1.853 1.615c-.594.446-1.952 1.282-1.952 1.282a1.71 1.71 0 0 1-1.69 0a21 21 0 0 1-1.952-1.282A16.4 16.4 0 0 1 7.35 15.64C5.9 14.156 4.5 12.068 4.5 9.5A7.5 7.5 0 0 1 12 2" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Marker(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2a9 9 0 0 1 9 9c0 3.074-1.676 5.59-3.442 7.395a20.4 20.4 0 0 1-2.876 2.416l-.426.29l-.2.133l-.377.24l-.336.205l-.416.242a1.88 1.88 0 0 1-1.854 0l-.416-.242l-.52-.32l-.192-.125l-.41-.273a20.7 20.7 0 0 1-3.093-2.566C4.676 16.588 3 14.074 3 11a9 9 0 0 1 9-9" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="m15.08 7.983l-4.245 4.244l-1.768-1.768a1 1 0 1 0-1.414 1.415l2.404 2.404a1.1 1.1 0 0 0 1.556 0l4.88-4.881a1 1 0 0 0-1.414-1.414z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Menu(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 18a.97.97 0 0 1-.712-.288A.97.97 0 0 1 3 17q0-.424.288-.712A.97.97 0 0 1 4 16h11q.425 0 .713.288A.97.97 0 0 1 16 17q0 .424-.287.712A.97.97 0 0 1 15 18z" class="duoicon-primary-layer"/><path fill="currentColor" d="m17.4 12l2.9 2.9a.95.95 0 0 1 .275.7a.95.95 0 0 1-.275.7a.95.95 0 0 1-.7.275a.95.95 0 0 1-.7-.275l-3.6-3.6a.96.96 0 0 1-.3-.7q0-.4.3-.7l3.6-3.6a.95.95 0 0 1 .7-.275q.425 0 .7.275a.95.95 0 0 1 .275.7a.95.95 0 0 1-.275.7z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M4 8a.97.97 0 0 1-.712-.287A.97.97 0 0 1 3 7q0-.425.288-.713A.97.97 0 0 1 4 6h11a.97.97 0 0 1 .713.287A.97.97 0 0 1 16 7a.97.97 0 0 1-.287.713A.97.97 0 0 1 15 8z" class="duoicon-primary-layer"/><path fill="currentColor" d="M4 13a.97.97 0 0 1-.712-.288A.97.97 0 0 1 3 12q0-.425.288-.713A.97.97 0 0 1 4 11h8a.97.97 0 0 1 .713.287A.97.97 0 0 1 13 12q0 .424-.287.712A.97.97 0 0 1 12 13z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Message(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2 6a3 3 0 0 1 3-3h14a3 3 0 0 1 3 3v10a3 3 0 0 1-3 3H7.333L4 21.5c-.824.618-2 .03-2-1z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M8 12a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2zM7 9a1 1 0 0 1 1-1h8a1 1 0 1 1 0 2H8a1 1 0 0 1-1-1" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func MessageThree(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 10a3 3 0 0 1 3 3v3a3 3 0 0 1-3 3v.966c0 1.06-1.236 1.639-2.05.96L14.638 19H12a3 3 0 0 1-3-3v-3a3 3 0 0 1 3-3z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M16 4a3 3 0 0 1 3 3v1h-8a4 4 0 0 0-4 4v4c0 1.044.4 1.996 1.056 2.708L7 19.5c-.824.618-2 .03-2-1V17a3 3 0 0 1-3-3V7a3 3 0 0 1 3-3z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func MessageTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10H4a2 2 0 0 1-2-2v-8C2 6.477 6.477 2 12 2" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M15 10H9a1 1 0 0 0-.117 1.993L9 12h6a1 1 0 0 0 .117-1.993zm-3 4H9a1 1 0 1 0 0 2h3a1 1 0 1 0 0-2" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func MoonStars(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.477 4.546A1.01 1.01 0 0 1 13.5 3.127q.037.003.074.01c6.821 1.213 9.771 9.356 5.31 14.656s-12.988 3.784-15.348-2.73a9 9 0 0 1-.399-1.489a1.01 1.01 0 0 1 1.339-1.125q.036.013.07.028c4.214 1.892 8.895-1.488 8.426-6.083a6 6 0 0 0-.495-1.848" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M8.397 2.857c.04-.09.166-.09.206 0l.102.222a5.2 5.2 0 0 0 1.97 2.172l.157.092c.073.04.075.144.003.187l-.003.002l-.158.092a5.2 5.2 0 0 0-2.07 2.394a.113.113 0 0 1-.195.022q-.007-.01-.012-.022l-.102-.222a5.2 5.2 0 0 0-1.97-2.172l-.158-.092a.108.108 0 0 1-.003-.187l.003-.002l.158-.092a5.2 5.2 0 0 0 1.97-2.172zM5.565 7.716l.064.14a3.26 3.26 0 0 0 1.237 1.363l.1.059a.068.068 0 0 1 0 .118l-.1.058a3.26 3.26 0 0 0-1.237 1.364l-.064.14a.07.07 0 0 1-.122.013l-.008-.013l-.064-.14a3.26 3.26 0 0 0-1.237-1.364l-.1-.058a.068.068 0 0 1 0-.118l.1-.059c.534-.326.964-.8 1.236-1.364l.064-.14a.07.07 0 0 1 .122-.013l.008.013z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func MoonTwo(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.15 18.125L5.875 3.85a9.9 9.9 0 0 1 2.437-1.825A10.3 10.3 0 0 1 11.25 1q-.45 2.475.275 4.837a9.9 9.9 0 0 0 2.5 4.138a9.9 9.9 0 0 0 4.138 2.5q2.362.725 4.837.275a9.6 9.6 0 0 1-1.012 2.938a10.2 10.2 0 0 1-1.838 2.437" class="duoicon-primary-layer"/><path fill="currentColor" d="m19.375 23.05l-2.7-2.7a10 10 0 0 1-1.737.487Q14.05 21 13.1 21a9.8 9.8 0 0 1-3.938-.8a10.3 10.3 0 0 1-3.199-2.162a10.3 10.3 0 0 1-2.163-3.2A9.8 9.8 0 0 1 3 10.9q0-.95.163-1.838q.162-.887.487-1.737L.975 4.65L2.4 3.225l18.4 18.4z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Palette(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 3a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M13.636 4a2 2 0 0 1 2.701-.117l.127.117L20 7.536a2 2 0 0 1 .204 2.589L13 17.357V4.636zM7.5 15a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3" class="duoicon-primary-layer"/><path fill="currentColor" fill-rule="evenodd" d="M19.66 12.111c.731.256 1.27.924 1.334 1.727L21 19a2 2 0 0 1-1.85 1.995L19 21h-6v-2.23z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Rocket(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m18.165 2.765l.255.032c.674.093 1.566.218 2.071.724c.414.413.573 1.085.668 1.685l.056.386c.126.91.159 2.102-.056 3.426c-.424 2.613-1.815 5.731-5.308 8.145c-.019.188-.02.378-.016.568l.01.284c.016.437.032.873-.09 1.298c-.19.66-.867 1.095-1.5 1.407l-.31.147l-.4.176c-.748.318-1.758.644-2.391.01c-.38-.379-.536-.935-.663-1.488l-.047-.207a8 8 0 0 0-.2-.774q-.075-.22-.162-.445a3 3 0 0 1-.203.225c-.345.345-.86.586-1.284.755c-.463.183-.987.343-1.472.475l-.249.066l-.477.119l-.432.1l-.517.11l-.323.063a1.01 1.01 0 0 1-1.177-1.177l.086-.431l.154-.698l.124-.51l.094-.36c.132-.484.292-1.008.476-1.47c.168-.425.409-.94.754-1.285l.08-.077l-.064-.026a8 8 0 0 0-.519-.177l-.277-.085c-.694-.21-1.436-.436-1.897-.898c-.56-.559-.371-1.41-.101-2.118l.11-.274l.177-.4l.147-.31c.312-.632.747-1.309 1.407-1.499c.35-.1.714-.106 1.08-.096l.22.007c.286.01.571.021.85-.006c2.414-3.494 5.532-4.885 8.145-5.309a11.8 11.8 0 0 1 3.171-.088" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M15.536 8.466c-1.088-1.089-2.948-.591-3.346.896a2 2 0 0 0 .517 1.932c1.088 1.089 2.948.591 3.346-.896a2 2 0 0 0-.517-1.932M8.353 15.44a1 1 0 0 0-1.1-.06l-.11.074l-.093.083l-.125.158c-.26.376-.408.896-.523 1.382l-.108.468l-.051.213l.191-.046l.418-.096c.578-.135 1.219-.31 1.613-.665a1 1 0 0 0 .088-1.314l-.082-.094l-.024-.023z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Settings(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.965 2.809a1.51 1.51 0 0 0-1.401-.203a10 10 0 0 0-2.982 1.725a1.51 1.51 0 0 0-.524 1.313c.075.753-.058 1.48-.42 2.106c-.361.627-.925 1.106-1.615 1.417c-.458.203-.786.62-.875 1.113a10 10 0 0 0 0 3.44c.093.537.46.926.875 1.114c.69.31 1.254.79 1.616 1.416c.361.627.494 1.353.419 2.106c-.045.452.107.964.524 1.313a10 10 0 0 0 2.982 1.725c.471.169.996.093 1.4-.203c.615-.442 1.312-.691 2.036-.691s1.42.249 2.035.691c.37.266.89.39 1.401.203a10 10 0 0 0 2.982-1.725c.417-.349.57-.86.524-1.313c-.075-.753.057-1.48.42-2.106c.361-.627.925-1.105 1.615-1.416c.414-.188.782-.577.875-1.114a10.1 10.1 0 0 0 0-3.44a1.51 1.51 0 0 0-.875-1.113c-.69-.311-1.254-.79-1.616-1.417c-.362-.626-.494-1.353-.419-2.106a1.51 1.51 0 0 0-.524-1.313a10 10 0 0 0-2.982-1.725a1.51 1.51 0 0 0-1.4.203C13.42 3.25 12.723 3.5 12 3.5s-1.42-.249-2.035-.691" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M9 12c0-2.309 2.5-3.753 4.5-2.598A3 3 0 0 1 15 12c0 2.309-2.5 3.753-4.5 2.598A3 3 0 0 1 9 12" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func ShoppingBag(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.464 3.282a2 2 0 0 1 2.964-.12l.108.12L17.468 8h2.985a1.49 1.49 0 0 1 1.484 1.655l-.092.766l-.1.74l-.082.554l-.095.595l-.108.625l-.122.648l-.136.661q-.108.5-.232.998a21 21 0 0 1-.832 2.584l-.221.54l-.214.488l-.202.434l-.094.194l-.249.49c-.32.61-.924.97-1.563 1.022l-.16.006H6.555a1.93 1.93 0 0 1-1.71-1.008l-.232-.45l-.18-.37l-.095-.205l-.2-.449a21.5 21.5 0 0 1-1.108-3.276a35 35 0 0 1-.156-.654l-.142-.648l-.127-.634l-.112-.613l-.1-.587l-.087-.554l-.074-.513l-.09-.683l-.066-.556l-.017-.153a1.49 1.49 0 0 1 1.348-1.64L3.543 8h2.989z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M12 4.562L9.135 8h5.73zm3.164 7.452a1 1 0 0 0-1.125.708l-.025.114l-.5 3a1 1 0 0 0 1.947.442l.025-.114l.5-3a1 1 0 0 0-.822-1.15m-5.203.708a1 1 0 0 0-1.96.326l.013.116l.5 3l.025.114a1 1 0 0 0 1.96-.326l-.013-.116l-.5-3z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Slideshow(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21 3a1 1 0 1 1 0 2v11a2 2 0 0 1-2 2h-5.055l2.293 2.293a1 1 0 0 1-1.414 1.414l-2.829-2.828l-2.828 2.828a1 1 0 0 1-1.414-1.414L10.046 18H5a2 2 0 0 1-2-2V5a1 1 0 1 1 0-2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M8 11a1 1 0 0 0-1 1v1a1 1 0 1 0 2 0v-1a1 1 0 0 0-1-1m4-2a1 1 0 0 0-1 1v3a1 1 0 1 0 2 0v-3a1 1 0 0 0-1-1m4-2a1 1 0 0 0-.993.883L15 8v5a1 1 0 0 0 1.993.117L17 13V8a1 1 0 0 0-1-1" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Smartphone(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 2a2 2 0 0 1 1.995 1.85L19 4v16a2 2 0 0 1-1.85 1.995L17 22H7a2 2 0 0 1-1.995-1.85L5 20V4a2 2 0 0 1 1.85-1.995L7 2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M12.5 16h-1a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func SmartphoneVibration(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m15.535 2.808l5.657 5.656a2 2 0 0 1 0 2.829l-9.9 9.9a2 2 0 0 1-2.828 0l-5.657-5.658a2 2 0 0 1 0-2.828l9.9-9.9a2 2 0 0 1 2.828 0z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M8.464 2.808a1 1 0 0 1 0 1.414L4.222 8.464A1 1 0 1 1 2.807 7.05L7.05 2.808a1 1 0 0 1 1.414 0m.354 10.96l-.707.707a.5.5 0 0 0 0 .707l.707.707a.5.5 0 0 0 .707 0l.707-.707a.5.5 0 0 0 0-.707l-.707-.707a.5.5 0 0 0-.707 0m12.374 1.768a1 1 0 0 1 0 1.414l-4.242 4.242a1 1 0 1 1-1.415-1.414l4.243-4.242a1 1 0 0 1 1.414 0" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Smartwatch(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.5 2a2 2 0 0 0-2 2v1.29C7.963 5.104 8.47 5 9 5h6a4 4 0 0 1 1.5.29V4a2 2 0 0 0-2-2zM9 19a4 4 0 0 1-1.5-.29V20a2 2 0 0 0 2 2h5a2 2 0 0 0 2-2v-1.29A4 4 0 0 1 15 19z" class="duoicon-primary-layer"/><path fill="currentColor" d="M9 6a3 3 0 0 0-3 3v6a3 3 0 0 0 3 3h6a3 3 0 0 0 3-3v-2a1 1 0 0 0 1-1v-1a1 1 0 0 0-1-1V9a3 3 0 0 0-3-3z" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Sun(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 18.5a1.5 1.5 0 0 1 1.493 1.356L13.5 20v1a1.5 1.5 0 0 1-2.993.144L10.5 21v-1a1.5 1.5 0 0 1 1.5-1.5m0-17a1.5 1.5 0 0 1 1.493 1.356L13.5 3v1a1.5 1.5 0 0 1-2.993.144L10.5 4V3A1.5 1.5 0 0 1 12 1.5m5.303 3.075a1.5 1.5 0 0 1 2.225 2.008l-.103.114l-.707.707a1.5 1.5 0 0 1-2.225-2.008l.103-.114zm-12.728 0a1.5 1.5 0 0 1 2.008-.103l.114.103l.707.707a1.5 1.5 0 0 1-2.008 2.225l-.114-.103l-.707-.707a1.5 1.5 0 0 1 0-2.122M21 10.5a1.5 1.5 0 0 1 .144 2.993L21 13.5h-1a1.5 1.5 0 0 1-.144-2.993L20 10.5zm-17 0a1.5 1.5 0 0 1 .144 2.993L4 13.5H3a1.5 1.5 0 0 1-.144-2.993L3 10.5z" class="duoicon-primary-layer"/><path fill="currentColor" d="M12 6c4.619 0 7.506 5 5.196 9A6 6 0 0 1 12 18c-4.619 0-7.506-5-5.196-9A6 6 0 0 1 12 6" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M5.282 16.596a1.5 1.5 0 0 1 2.225 2.008l-.103.114l-.707.707a1.5 1.5 0 0 1-2.225-2.008l.103-.114zm11.314 0a1.5 1.5 0 0 1 2.008-.103l.114.103l.707.707a1.5 1.5 0 0 1-2.008 2.225l-.114-.103l-.707-.707a1.5 1.5 0 0 1 0-2.122" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Target(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2a1 1 0 0 1 .993.883L13 3v.055a9.005 9.005 0 0 1 7.911 7.674l.034.271H21a1 1 0 0 1 .117 1.993L21 13h-.055a9.005 9.005 0 0 1-7.674 7.911l-.271.034V21a1 1 0 0 1-1.993.117L11 21v-.055a9.005 9.005 0 0 1-7.911-7.674L3.055 13H3a1 1 0 0 1-.117-1.993L3 11h.055a9.005 9.005 0 0 1 7.674-7.911L11 3.055V3a1 1 0 0 1 1-1" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M12 7c-3.849 0-6.255 4.167-4.33 7.5A5 5 0 0 0 12 17c3.849 0 6.255-4.167 4.33-7.5A5 5 0 0 0 12 7" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func Toggle(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 12c0-1.54 1.667-2.502 3-1.732c.619.357 1 1.017 1 1.732c0 1.54-1.667 2.502-3 1.732A2 2 0 0 1 6 12" class="duoicon-primary-layer"/><path fill="currentColor" fill-rule="evenodd" d="M8 5C2.611 5-.756 10.833 1.938 15.5A7 7 0 0 0 8 19h8c5.389 0 8.756-5.833 6.062-10.5A7 7 0 0 0 16 5zm0 3c-3.079 0-5.004 3.333-3.464 6A4 4 0 0 0 8 16c3.079 0 5.004-3.333 3.464-6A4 4 0 0 0 8 8" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func Translation(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 10.5a1.5 1.5 0 0 1 1.493 1.356L18.5 12v.5h1a2 2 0 0 1 1.995 1.85l.005.15v3a2 2 0 0 1-1.85 1.995l-.15.005h-1v.5a1.5 1.5 0 0 1-2.993.144L15.5 20v-.5h-1a2 2 0 0 1-1.995-1.85l-.005-.15v-3a2 2 0 0 1 1.85-1.995l.15-.005h1V12a1.5 1.5 0 0 1 1.5-1.5" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M17 3.5A3.5 3.5 0 0 1 20.5 7v2a1.5 1.5 0 0 1-3 0V7a.5.5 0 0 0-.5-.5h-3a1.5 1.5 0 0 1 0-3z" class="duoicon-primary-layer"/><path fill="currentColor" d="M9.5 2.5a1.5 1.5 0 0 1 0 3h-4v1H9a1.5 1.5 0 0 1 0 3H5.5v1H10a1.5 1.5 0 0 1 0 3H4.1a1.6 1.6 0 0 1-1.6-1.6V4.1a1.6 1.6 0 0 1 1.6-1.6z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M19.5 15h-1v2h1zm-4 0h-1v2h1zM5 14.5A1.5 1.5 0 0 1 6.5 16v1a.5.5 0 0 0 .5.5h3a1.5 1.5 0 0 1 0 3H7A3.5 3.5 0 0 1 3.5 17v-1A1.5 1.5 0 0 1 5 14.5" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func UploadFile(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2v6.5a1.5 1.5 0 0 0 1.356 1.493L13.5 10H20v10a2 2 0 0 1-1.85 1.995L18 22H6a2 2 0 0 1-1.995-1.85L4 20V4a2 2 0 0 1 1.85-1.995L6 2z" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M14 2.043a2 2 0 0 1 .877.43l.123.113L19.414 7c.234.234.407.523.502.84l.04.16H14zm-2.707 9.13l-2.121 2.121a1 1 0 1 0 1.414 1.414l.414-.414V17a1 1 0 1 0 2 0v-2.706l.414.414a1 1 0 1 0 1.414-1.414l-2.12-2.121a1 1 0 0 0-1.415 0" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func User(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 13c2.396 0 4.575.694 6.178 1.671c.8.49 1.484 1.065 1.978 1.69c.486.616.844 1.352.844 2.139c0 .845-.411 1.511-1.003 1.986c-.56.45-1.299.748-2.084.956c-1.578.417-3.684.558-5.913.558s-4.335-.14-5.913-.558c-.785-.208-1.524-.506-2.084-.956C3.41 20.01 3 19.345 3 18.5c0-.787.358-1.523.844-2.139c.494-.625 1.177-1.2 1.978-1.69C7.425 13.694 9.605 13 12 13" class="duoicon-primary-layer"/><path fill="currentColor" d="M12 2c3.849 0 6.255 4.167 4.33 7.5A5 5 0 0 1 12 12c-3.849 0-6.255-4.167-4.33-7.5A5 5 0 0 1 12 2" class="duoicon-secondary-layer" opacity=".3"/>`), g.Group(children),
	)
}

func UserCard(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14.447 1.106a1 1 0 0 1 .447 1.341L14.118 4H18a2 2 0 0 1 2 2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V6a2 2 0 0 1 2-2h3.882l-.776-1.553a1 1 0 0 1 1.788-.894L12 3.763l1.106-2.21a1 1 0 0 1 1.341-.447" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" fill-rule="evenodd" d="M12 9c-1.54 0-2.502 1.667-1.732 3c.357.619 1.017 1 1.732 1c1.54 0 2.502-1.667 1.732-3A2 2 0 0 0 12 9m1.5 5h-3a2.5 2.5 0 0 0-2.495 2.336L8 16.5v.5a1 1 0 0 0 1.993.117L10 17v-.5a.5.5 0 0 1 .41-.492L10.5 16h3a.5.5 0 0 1 .492.41l.008.09v.5a1 1 0 0 0 1.993.117L16 17v-.5a2.5 2.5 0 0 0-2.336-2.495z" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}

func World(children ...g.Node) g.Node {
	return h.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2" class="duoicon-secondary-layer" opacity=".3"/><path fill="currentColor" d="M12 4a7.99 7.99 0 0 0-6.335 3.114l-.165.221V9.02c0 1.25.775 2.369 1.945 2.809l.178.06l1.29.395c1.373.42 2.71-.697 2.577-2.096l-.019-.145l-.175-1.049a1 1 0 0 1 .656-1.108l.108-.03l.612-.14a2.667 2.667 0 0 0 1.989-3.263A8 8 0 0 0 12 4m2 9.4l-1.564 1.251a.5.5 0 0 0-.041.744l1.239 1.239c.24.24.415.538.508.864l.175.613c.147.521.52.948 1.017 1.163a8 8 0 0 0 2.533-1.835l-.234-1.877a2 2 0 0 0-1.09-1.54l-1.47-.736A1 1 0 0 0 14 13.4" class="duoicon-primary-layer"/>`), g.Group(children),
	)
}
