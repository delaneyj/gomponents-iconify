package entypo

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

const IconifyVersion = ""

var (
	hAttr   = g.Attr("height", "1em")
	viewbox = g.Attr("viewbox", "0 0 20 20")
)

func AddToList(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.4 9H16V5.6c0-.6-.4-.6-1-.6s-1 0-1 .6V9h-3.4c-.6 0-.6.4-.6 1s0 1 .6 1H14v3.4c0 .6.4.6 1 .6s1 0 1-.6V11h3.4c.6 0 .6-.4.6-1s0-1-.6-1zm-12 0H.6C0 9 0 9.4 0 10s0 1 .6 1h6.8c.6 0 .6-.4.6-1s0-1-.6-1zm0 5H.6c-.6 0-.6.4-.6 1s0 1 .6 1h6.8c.6 0 .6-.4.6-1s0-1-.6-1zm0-10H.6C0 4 0 4.4 0 5s0 1 .6 1h6.8C8 6 8 5.6 8 5s0-1-.6-1z"/>`), g.Group(children),
	)
}

func AddUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.989 19.129C16 17 13.803 15.74 11.672 14.822c-2.123-.914-2.801-1.684-2.801-3.334c0-.989.648-.667.932-2.481c.12-.752.692-.012.802-1.729c0-.684-.313-.854-.313-.854s.159-1.013.221-1.793c.064-.817-.398-2.56-2.301-3.095c-.332-.341-.557-.882.467-1.424c-2.24-.104-2.761 1.068-3.954 1.93c-1.015.756-1.289 1.953-1.24 2.59c.065.78.223 1.793.223 1.793s-.314.17-.314.854c.11 1.718.684.977.803 1.729c.284 1.814.933 1.492.933 2.481c0 1.65-.212 2.21-2.336 3.124C.663 15.53 0 17 .011 19.129C.014 19.766 0 20 0 20h16s-.014-.234-.011-.871zM17 10V7h-2v3h-3v2h3v3h2v-3h3v-2h-3z"/>`), g.Group(children),
	)
}

func Address(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19.799 5.165l-2.375-1.83a1.997 1.997 0 0 0-.521-.237A2.035 2.035 0 0 0 16.336 3H9.5l.801 5h6.035c.164 0 .369-.037.566-.098s.387-.145.521-.236l2.375-1.832c.135-.091.202-.212.202-.334s-.067-.243-.201-.335zM8.5 1h-1a.5.5 0 0 0-.5.5V5H3.664c-.166 0-.37.037-.567.099c-.198.06-.387.143-.521.236L.201 7.165C.066 7.256 0 7.378 0 7.5c0 .121.066.242.201.335l2.375 1.832c.134.091.323.175.521.235c.197.061.401.098.567.098H7v8.5a.5.5 0 0 0 .5.5h1a.5.5 0 0 0 .5-.5v-17a.5.5 0 0 0-.5-.5z"/>`), g.Group(children),
	)
}

func Adjust(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 9.199h-.98c-.553 0-1 .359-1 .801c0 .441.447.799 1 .799H19c.552 0 1-.357 1-.799c0-.441-.449-.801-1-.801zM10 4.5A5.483 5.483 0 0 0 4.5 10c0 3.051 2.449 5.5 5.5 5.5c3.05 0 5.5-2.449 5.5-5.5S13.049 4.5 10 4.5zm0 9.5c-2.211 0-4-1.791-4-4c0-2.211 1.789-4 4-4v8zm-7-4c0-.441-.449-.801-1-.801H1c-.553 0-1 .359-1 .801c0 .441.447.799 1 .799h1c.551 0 1-.358 1-.799zm7-7c.441 0 .799-.447.799-1V1c0-.553-.358-1-.799-1c-.442 0-.801.447-.801 1v1c0 .553.359 1 .801 1zm0 14c-.442 0-.801.447-.801 1v1c0 .553.359 1 .801 1c.441 0 .799-.447.799-1v-1c0-.553-.358-1-.799-1zm7.365-13.234c.391-.391.454-.961.142-1.273s-.883-.248-1.272.143l-.7.699c-.391.391-.454.961-.142 1.273s.883.248 1.273-.143l.699-.699zM3.334 15.533l-.7.701c-.391.391-.454.959-.142 1.271s.883.25 1.272-.141l.7-.699c.391-.391.454-.961.142-1.274s-.883-.247-1.272.142zm.431-12.898c-.39-.391-.961-.455-1.273-.143s-.248.883.141 1.274l.7.699c.391.391.96.455 1.272.143s.249-.883-.141-1.273l-.699-.7zm11.769 14.031l.7.699c.391.391.96.453 1.272.143c.312-.312.249-.883-.142-1.273l-.699-.699c-.391-.391-.961-.455-1.274-.143s-.248.882.143 1.273z"/>`), g.Group(children),
	)
}

func Air(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.643 6.357c1.747-1.5 3.127-2.686 6.872-.57c1.799 1.016 3.25 1.4 4.457 1.398c2.115 0 3.486-1.176 4.671-2.193a1.037 1.037 0 0 0 .122-1.439a.987.987 0 0 0-1.41-.125c-1.746 1.502-3.127 2.688-6.872.57c-4.948-2.793-7.266-.803-9.128.797a1.037 1.037 0 0 0-.121 1.439a.986.986 0 0 0 1.409.123zm14.712 2.178c-1.746 1.5-3.127 2.688-6.872.57c-4.948-2.795-7.266-.804-9.128.795a1.037 1.037 0 0 0-.121 1.439a.986.986 0 0 0 1.409.125c1.747-1.501 3.127-2.687 6.872-.572c1.799 1.018 3.25 1.4 4.457 1.4c2.115 0 3.486-1.176 4.671-2.195a1.035 1.035 0 0 0 .122-1.438a.986.986 0 0 0-1.41-.124zm0 5.106c-1.746 1.502-3.127 2.688-6.872.572c-4.948-2.795-7.266-.805-9.128.795a1.037 1.037 0 0 0-.121 1.439a.985.985 0 0 0 1.409.123c1.747-1.5 3.127-2.685 6.872-.57c1.799 1.016 3.25 1.4 4.457 1.4c2.115 0 3.486-1.178 4.671-2.195a1.037 1.037 0 0 0 .122-1.439a.988.988 0 0 0-1.41-.125z"/>`), g.Group(children),
	)
}

func Aircraft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.496 17.414c-.394-1.096-1.805-4.775-2.39-6.297c-1.103.737-2.334 1.435-3.512 1.928c-.366 1.28-1.094 3.709-1.446 4.033c-.604.557-.832.485-.925-.279c-.093-.764-.485-3.236-.485-3.236s-2.162-1.219-2.84-1.568s-.667-.591.057-.974c.422-.223 2.927-.085 4.242.005c.861-.951 1.931-1.882 2.993-2.679c-1.215-1.076-4.15-3.675-5.034-4.424c-.776-.658.079-.797.079-.797c.39-.07 1.222-.132 1.628-.009c2.524.763 6.442 2.068 7.363 2.376l1.162-.821c4.702-3.33 5.887-2.593 6.111-2.27s.503 1.701-4.199 5.032l-1.16.823c-.029.98-.157 5.151-.311 7.811c-.025.428-.367 1.198-.565 1.544c-.001 0-.423.765-.768-.198z"/>`), g.Group(children),
	)
}

func AircraftLanding(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.752 16.038c-.097.266-.822 1.002-6.029-.878l-5.105-1.843C5.841 12.676 3.34 11.668 2.36 11.1c-.686-.397-.836-1.282-.836-1.282s-.163-2.956-.263-3.684c-.1-.728.095-.853.796-.492c.436.225 1.865 2.562 2.464 3.567c1.512.381 2.862.761 3.493.949c-.257-1.717-.74-4.928-.913-5.933c-.166-.963.55-.535.55-.535c.331.19.983.661 1.206 1.002c1.522 2.326 3.672 6.6 3.836 6.928c.896.28 2.277.733 3.102 1.03c2.156.779 3.087 3.034 2.957 3.388z"/>`), g.Group(children),
	)
}

func AircraftTakeOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.87 6.453c.119.257.127 1.29-4.884 3.642l-4.913 2.306c-1.71.803-4.191 1.859-5.285 2.151c-.766.204-1.497-.316-1.497-.316S1.085 12.261.499 11.817c-.585-.444-.535-.67.215-.91c.467-.149 3.13.493 4.265.78A91.697 91.697 0 0 1 8.12 9.889c-1.396-1.033-4.008-2.962-4.841-3.55c-.799-.565.01-.768.01-.768c.368-.099 1.162-.228 1.562-.144c2.721.569 7.263 2.071 7.611 2.186a90.641 90.641 0 0 1 2.922-1.465c2.075-.974 4.327-.037 4.486.305z"/>`), g.Group(children),
	)
}

func AlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 11h-2V3H9v8H7l3 3l3-3zm4.4 4H2.6c-.552 0-.6.447-.6 1c0 .553.048 1 .6 1h14.8c.552 0 .6-.447.6-1c0-.553-.048-1-.6-1z"/>`), g.Group(children),
	)
}

func AlignHorizontalMiddle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 10L5 7v2H1v2h4v2l3-3zm7 3v-2h4V9h-4V7l-3 3l3 3zm-5 5c.553 0 1-.049 1-.6V2.6c0-.553-.447-.6-1-.6c-.552 0-1 .047-1 .6v14.8c0 .551.448.6 1 .6z"/>`), g.Group(children),
	)
}

func AlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6 10l3 3v-2h8V9H9V7l-3 3zM4 2c-.553 0-1 .047-1 .6v14.8c0 .551.447.6 1 .6c.552 0 1-.049 1-.6V2.6c0-.553-.448-.6-1-.6z"/>`), g.Group(children),
	)
}

func AlignRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 7v2H3v2h8v2l3-3l-3-3zm4-4.4v14.8c0 .551.448.6 1 .6c.553 0 1-.049 1-.6V2.6c0-.553-.447-.6-1-.6c-.552 0-1 .047-1 .6z"/>`), g.Group(children),
	)
}

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 6L7 9h2v8h2V9h2l-3-3zm8-2c0-.553-.048-1-.6-1H2.6c-.552 0-.6.447-.6 1c0 .553.048 1 .6 1h14.8c.552 0 .6-.447.6-1z"/>`), g.Group(children),
	)
}

func AlignVerticalMiddle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10 12l-3 3h2v4h2v-4h2l-3-3zm3-7h-2V1H9v4H7l3 3l3-3zm5 5c0-.553-.048-1-.6-1H2.6c-.552 0-.6.447-.6 1c0 .551.048 1 .6 1h14.8c.552 0 .6-.449.6-1z"/>`), g.Group(children),
	)
}

func Archive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.981 2H6.018s-.996 0-.996 1h9.955c0-1-.996-1-.996-1zm2.987 3c0-1-.995-1-.995-1H4.027s-.995 0-.995 1v1h13.936V5zm1.99 1l-.588-.592V7H1.63V5.408L1.041 6C.452 6.592.03 6.75.267 8c.236 1.246 1.379 8.076 1.549 9c.186 1.014 1.217 1 1.217 1h13.936s1.03.014 1.217-1c.17-.924 1.312-7.754 1.549-9c.235-1.25-.187-1.408-.777-2zM14 11.997c0 .554-.449 1.003-1.003 1.003H7.003A1.003 1.003 0 0 1 6 11.997V10h1v2h6v-2h1v1.997z"/>`), g.Group(children),
	)
}

func AreaGraph(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 2v16H.32c-.318 0-.416-.209-.216-.465l4.469-5.748a.526.526 0 0 1 .789-.062l1.419 1.334a.473.473 0 0 0 .747-.096l3.047-4.74a.466.466 0 0 1 .741-.09l2.171 2.096c.232.225.559.18.724-.1l5.133-7.785C19.51 2.062 19.75 2 20 2z"/>`), g.Group(children),
	)
}

func ArrowBoldDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.5 10H6V3h8v7h3.5L10 17.5L2.5 10z"/>`), g.Group(children),
	)
}

func ArrowBoldLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 2.5V6h7v8h-7v3.5L2.5 10L10 2.5z"/>`), g.Group(children),
	)
}

func ArrowBoldRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.5 10L10 17.5V14H3V6h7V2.5l7.5 7.5z"/>`), g.Group(children),
	)
}

func ArrowBoldUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10 2.5l7.5 7.5H14v7H6v-7H2.5L10 2.5z"/>`), g.Group(children),
	)
}

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 17.5L3.5 11H7V3h6v8h3.5L10 17.5z"/>`), g.Group(children),
	)
}

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.5 10L9 3.5V7h8v6H9v3.5L2.5 10z"/>`), g.Group(children),
	)
}

func ArrowLongDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 19.25L4.5 14H8V1h4v13h3.5L10 19.25z"/>`), g.Group(children),
	)
}

func ArrowLongLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M.75 10L6 4.5V8h13v4H6v3.5L.75 10z"/>`), g.Group(children),
	)
}

func ArrowLongRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 15.5V12H1V8h13V4.5l5.25 5.5L14 15.5z"/>`), g.Group(children),
	)
}

func ArrowLongUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .75L15.5 6H12v13H8V6H4.5L10 .75z"/>`), g.Group(children),
	)
}

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 16.5V13H3V7h8V3.5l6.5 6.5l-6.5 6.5z"/>`), g.Group(children),
	)
}

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 2.5L16.5 9H13v8H7V9H3.5L10 2.5z"/>`), g.Group(children),
	)
}

func ArrowWithCircleDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4C4.697.4.399 4.698.399 10A9.6 9.6 0 0 0 10 19.601c5.301 0 9.6-4.298 9.6-9.601c0-5.302-4.299-9.6-9.6-9.6zm-.001 17.2a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 1 1 0 15.2zM12 6H8v4H5.5l4.5 4.5l4.5-4.5H12V6z"/>`), g.Group(children),
	)
}

func ArrowWithCircleLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4C4.697.4.399 4.698.399 10A9.6 9.6 0 0 0 10 19.601c5.301 0 9.6-4.298 9.6-9.601c0-5.302-4.299-9.6-9.6-9.6zm-.001 17.2a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 1 1 0 15.2zM10 5.5L5.5 10l4.5 4.5V12h4V8h-4V5.5z"/>`), g.Group(children),
	)
}

func ArrowWithCircleRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4C4.697.4.399 4.698.399 10A9.6 9.6 0 0 0 10 19.601c5.301 0 9.6-4.298 9.6-9.601c0-5.302-4.299-9.6-9.6-9.6zm-.001 17.2a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 1 1 0 15.2zM10 8H6v4h4v2.5l4.5-4.5L10 5.5V8z"/>`), g.Group(children),
	)
}

func ArrowWithCircleUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4C4.697.4.399 4.698.399 10A9.6 9.6 0 0 0 10 19.601c5.301 0 9.6-4.298 9.6-9.601c0-5.302-4.299-9.6-9.6-9.6zm-.001 17.2a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 1 1 0 15.2zM10 5.5l4.5 4.5H12v4H8v-4H5.5L10 5.5z"/>`), g.Group(children),
	)
}

func Attachment(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.602 19.8c-1.293 0-2.504-.555-3.378-1.44c-1.695-1.716-2.167-4.711.209-7.116l9.748-9.87c.988-1 2.245-1.387 3.448-1.06c1.183.32 2.151 1.301 2.468 2.498c.322 1.22-.059 2.493-1.046 3.493l-9.323 9.44c-.532.539-1.134.858-1.738.922c-.599.064-1.17-.13-1.57-.535c-.724-.736-.828-2.117.378-3.337l6.548-6.63c.269-.272.705-.272.974 0s.269.714 0 .986l-6.549 6.631c-.566.572-.618 1.119-.377 1.364c.106.106.266.155.451.134c.283-.029.606-.216.909-.521l9.323-9.439c.64-.648.885-1.41.69-2.145a2.162 2.162 0 0 0-1.493-1.513c-.726-.197-1.48.052-2.12.7l-9.748 9.87c-1.816 1.839-1.381 3.956-.209 5.143c1.173 1.187 3.262 1.629 5.079-.212l9.748-9.87a.683.683 0 0 1 .974 0a.704.704 0 0 1 0 .987L9.25 18.15c-1.149 1.162-2.436 1.65-3.648 1.65z"/>`), g.Group(children),
	)
}

func AwarenessRibbon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.574 16.338c-.757-1.051-2.851-3.824-4.57-6.106c.696-.999 1.251-1.815 1.505-2.242c1.545-2.594.874-4.26.022-5.67C12.677.909 12.542.094 10 .094c-2.543 0-2.678.815-3.531 2.227c-.854 1.41-1.524 3.076.021 5.67c.254.426.809 1.243 1.506 2.242c-1.72 2.281-3.814 5.055-4.571 6.106c-.176.244-.16.664.009 1.082c.13.322.63 1.762.752 2.064c.156.389.664.67 1.082.092c.241-.334 2.582-3.525 4.732-6.522c2.149 2.996 4.491 6.188 4.732 6.522c.417.578.926.297 1.082-.092c.122-.303.622-1.742.752-2.064c.167-.419.184-.839.008-1.083zm-6.94-9.275C8.566 5.579 7.802 3.852 7.802 3.852s.42-.758 2.198-.758s2.198.758 2.198.758s-.766 1.727-1.833 3.211L10 7.56a40.64 40.64 0 0 1-.366-.497z"/>`), g.Group(children),
	)
}

func Back(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 7v6c0 1.103-.896 2-2 2H3v-3h13V8H5v2L1 6.5L5 3v2h12a2 2 0 0 1 2 2z"/>`), g.Group(children),
	)
}

func BackInTime(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 1.799c-4.445 0-8.061 3.562-8.169 7.996V10H.459l3.594 3.894L7.547 10H4.875v-.205C4.982 6.492 7.683 3.85 11 3.85c3.386 0 6.131 2.754 6.131 6.15c0 3.396-2.745 6.15-6.131 6.15a6.099 6.099 0 0 1-3.627-1.193l-1.406 1.504A8.13 8.13 0 0 0 11 18.199c4.515 0 8.174-3.67 8.174-8.199S15.515 1.799 11 1.799zM10 5v5a1.01 1.01 0 0 0 .293.707l3.2 3.2c.283-.183.55-.389.787-.628L12 11V5h-2z"/>`), g.Group(children),
	)
}

func BarGraph(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 1h-2a1 1 0 0 0-1 1v16.992h4V2a1 1 0 0 0-1-1zm-6 6H9a1 1 0 0 0-1 1v10.992h4V8a1 1 0 0 0-1-1zm-6 6H3a1 1 0 0 0-1 1v4.992h4V14a1 1 0 0 0-1-1z"/>`), g.Group(children),
	)
}

func Battery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.408 10c0-2.766 1.277-4.32 2.277-4.32H19C18.332 4.621 17.779 4 15.342 4H5.334C1.6 4 0 7.441 0 10s1.6 6 5.334 6h10.008c2.438 0 2.99-.621 3.658-1.68h-1.315c-1 0-2.277-1.554-2.277-4.32zm-2.72 1.795c-.164.25-.676.016-.676.016l-2.957-1.338s-.264.67-.467 1.141c-.205.471-.361 1.004-1.209.408c-.849-.598-3.581-3.25-3.581-3.25s-.345-.284-.173-.551c.163-.252.676-.016.676-.016l2.956 1.336s.265-.668.468-1.139c.205-.47.361-1.006 1.209-.408c.849.596 3.58 3.25 3.58 3.25s.345.283.174.551zm6.186-3.867h-.749c-.559 0-1.105.754-1.105 1.979c0 1.227.547 1.98 1.105 1.98h.749c.56 0 1.126-.754 1.126-1.98c0-1.225-.566-1.979-1.126-1.979z"/>`), g.Group(children),
	)
}

func BeamedNote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17 1l-.002 13c0 1.243-1.301 3-3.748 3c-1.243 0-2.25-.653-2.25-1.875c0-1.589 1.445-2.55 3-2.55c.432 0 .754.059 1 .123V5.364L8 6.637V16h-.002c0 1.243-1.301 3-3.748 3C3.007 19 2 18.347 2 17.125c0-1.589 1.445-2.55 3-2.55c.432 0 .754.059 1 .123V3l11-2z"/>`), g.Group(children),
	)
}

func Bell(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.65 8.512c-2.28-4.907-3.466-6.771-7.191-6.693c-1.327.027-1.009-.962-2.021-.587c-1.01.375-.143.924-1.177 1.773c-2.902 2.383-2.635 4.587-1.289 9.84c.567 2.213-1.367 2.321-.602 4.465c.559 1.564 4.679 2.219 9.025.607c4.347-1.613 7.086-4.814 6.527-6.378c-.765-2.145-2.311-.961-3.272-3.027zm-3.726 8.083c-3.882 1.44-7.072.594-7.207.217c-.232-.65 1.253-2.816 5.691-4.463c4.438-1.647 6.915-1.036 7.174-.311c.153.429-1.775 3.116-5.658 4.557zm-1.248-3.494c-2.029.753-3.439 1.614-4.353 2.389c.643.584 1.847.726 3.046.281c1.527-.565 2.466-1.866 2.095-2.904l-.016-.036c-.251.082-.508.171-.772.27z"/>`), g.Group(children),
	)
}

func Blackboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M2.539 20H6l1.406-3.698l-2.966-1.004L2.539 20zm10.055-3.698L14 20h3.461l-1.901-4.702l-2.966 1.004zM18 2h-6.5L11 0H9l-.5 2H2a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func Block(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4C4.697.4.399 4.698.399 10A9.6 9.6 0 0 0 10 19.601c5.301 0 9.6-4.298 9.6-9.601c0-5.302-4.299-9.6-9.6-9.6zM2.399 10a7.6 7.6 0 0 1 12.417-5.877L4.122 14.817A7.568 7.568 0 0 1 2.399 10zm7.6 7.599a7.56 7.56 0 0 1-4.815-1.722L15.878 5.184a7.6 7.6 0 0 1-5.879 12.415z"/>`), g.Group(children),
	)
}

func Book(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 5.95v10.351c0 .522-.452.771-1 1.16c-.44.313-1-.075-1-.587V6.76c0-.211-.074-.412-.314-.535c-.24-.123-7.738-4.065-7.738-4.065c-.121-.045-.649-.378-1.353-.016c-.669.344-1.033.718-1.126.894l8.18 4.482c.217.114.351.29.351.516v10.802a.67.67 0 0 1-.369.585a.746.746 0 0 1-.333.077a.736.736 0 0 1-.386-.104c-.215-.131-7.774-4.766-8.273-5.067c-.24-.144-.521-.439-.527-.658L3 3.385c0-.198-.023-.547.289-1.032C3.986 1.269 6.418.036 7.649.675l8.999 4.555c.217.112.352.336.352.72z"/>`), g.Group(children),
	)
}

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 2v17l-4-4l-4 4V2c0-.553.585-1.02 1-1h6c.689-.02 1 .447 1 1z"/>`), g.Group(children),
	)
}

func Bookmarks(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 0h-4a1 1 0 0 0-1 1l.023.222c1.102 0 2 .897 2 2v11.359L13 13.4l3 3.6V1a1 1 0 0 0-1-1zM9.023 3H5a1 1 0 0 0-1 1v16l3-3.6l3 3.6V4c0-.553-.424-1-.977-1z"/>`), g.Group(children),
	)
}

func Bowl(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.949 7.472c-2.176 2.902-4.095 3.002-7.046 3.152h-.101c-3.591-.002-6.138-1.336-6.138-1.832c-.002-.471 2.298-1.697 5.605-1.819l.59-1.473l-.057-.002c-4.908 0-7.791 1.562-7.791 3.051v2c0 .918.582 8.949 7.582 8.949s8-8.031 8-8.949v-2c0-.391-.201-.787-.584-1.158l-.06.081zm.64-4.77a1 1 0 0 0-1.399.201l-3.608 4.809l2.336-5.838a1 1 0 1 0-1.857-.742L9.802 9.274c2.882-.147 4.277-.227 6.067-2.611l1.919-2.561a1 1 0 0 0-.199-1.4z"/>`), g.Group(children),
	)
}

func Box(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.399 2H1.6c-.332 0-.6.267-.6.6V5h18V2.6a.6.6 0 0 0-.601-.6zM2 16.6c0 .77.629 1.4 1.399 1.4h13.2c.77 0 1.4-.631 1.4-1.4V6H2v10.6zM7 8h6v2H7V8z"/>`), g.Group(children),
	)
}

func Briefcase(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 10h2v2h9s-.149-4.459-.2-5.854C19.75 4.82 19.275 4 17.8 4h-3.208l-1.197-2.256C13.064 1.121 12.951 1 12.216 1H7.783c-.735 0-.847.121-1.179.744c-.165.311-.7 1.318-1.196 2.256H2.199c-1.476 0-1.945.82-2 2.146C.145 7.473 0 12 0 12h9v-2zM7.649 2.916c.23-.432.308-.516.817-.516h3.067c.509 0 .588.084.816.516L12.924 4h-5.85l.575-1.084zM11 15H9v-2H.5s.124 1.797.199 3.322C.73 16.955.917 18 2.499 18H17.5c1.582 0 1.765-1.047 1.8-1.678c.087-1.568.2-3.322.2-3.322H11v2z"/>`), g.Group(children),
	)
}

func Browser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 2H2C.9 2 0 2.9 0 4v12c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zM4.5 3.75a.75.75 0 1 1 0 1.5a.75.75 0 0 1 0-1.5zm-2.75.75a.75.75 0 1 1 1.5 0a.75.75 0 0 1-1.5 0zM18 16H2V7h16v9zm0-11H6V4h12.019L18 5z"/>`), g.Group(children),
	)
}

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.763 13.563c-1.515 1.488-.235 3.016-2.247 5.279c-.908 1.023 3.738.711 6.039-1.551c.977-.961.701-2.359-.346-3.389c-1.047-1.028-2.47-1.3-3.446-.339zM19.539.659C18.763-.105 10.16 6.788 7.6 9.305c-1.271 1.25-1.695 1.92-2.084 2.42c-.17.219.055.285.154.336c.504.258.856.496 1.311.943c.456.447.699.793.959 1.289c.053.098.121.318.342.152c.51-.383 1.191-.801 2.462-2.049C13.305 9.88 20.317 1.422 19.539.659z"/>`), g.Group(children),
	)
}

func Bucket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 1C6.092 1 3.002 2.592 3.21 3.95c.06.389.225 1.945.434 3.273C1.239 8.157.442 9.672.549 10.907c.127 1.461 1.441 3.025 4.328 3.295c1.648.154 3.631-.75 4.916-2.295a1.4 1.4 0 1 1 1.238.691c-1.529 1.973-3.858 3.164-6.064 3.025c.051.324.07.947.096 1.113c.09.579 2.347 2.26 5.937 2.264c3.59-.004 5.847-1.685 5.938-2.263c.088-.577 1.641-11.409 1.852-12.787C18.998 2.592 15.907 1 11 1zm-9.057 9.785c-.055-.643.455-1.498 1.924-2.139l.643 4.074c-1.604-.313-2.498-1.149-2.567-1.935zM11 6.024C7.41 6.022 4.863 4.69 4.863 4.192C4.861 3.698 7.41 2.402 11 2.404c3.59-.002 6.139 1.294 6.137 1.788c0 .498-2.547 1.83-6.137 1.832z"/>`), g.Group(children),
	)
}

func Cake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.584 6.036c1.952 0 2.591-1.381 1.839-2.843c-.871-1.693 1.895-3.155.521-3.155c-1.301 0-3.736 1.418-4.19 3.183c-.339 1.324.296 2.815 1.83 2.815zm5.212 8.951l-.444-.383a1.355 1.355 0 0 0-1.735 0l-.442.382a3.326 3.326 0 0 1-2.174.801a3.325 3.325 0 0 1-2.173-.8l-.444-.384a1.353 1.353 0 0 0-1.734.001l-.444.383c-1.193 1.028-2.967 1.056-4.204.1V19a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-3.912c-1.237.954-3.011.929-4.206-.101zM10 7c-7.574 0-9 3.361-9 5v.469l1.164 1.003a1.355 1.355 0 0 0 1.735 0l.444-.383a3.353 3.353 0 0 1 4.345 0l.444.384c.484.417 1.245.42 1.735-.001l.442-.382a3.352 3.352 0 0 1 4.346-.001l.444.383c.487.421 1.25.417 1.735 0L19 12.469V12c0-1.639-1.426-5-9-5z"/>`), g.Group(children),
	)
}

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.6 1H5.398C4.629 1 4 1.629 4 2.4v15.2c0 .77.629 1.4 1.398 1.4H14.6c.769 0 1.4-.631 1.4-1.4V2.4c0-.771-.631-1.4-1.4-1.4zM7 12c.689 0 1.25.447 1.25 1S7.689 14 7 14c-.69 0-1.25-.447-1.25-1s.56-1 1.25-1zm-1.25-2c0-.553.56-1 1.25-1c.689 0 1.25.447 1.25 1c0 .553-.561 1-1.25 1c-.69 0-1.25-.447-1.25-1zM7 15c.689 0 1.25.447 1.25 1S7.689 17 7 17c-.69 0-1.25-.447-1.25-1s.56-1 1.25-1zm3-3c.689 0 1.25.447 1.25 1s-.561 1-1.25 1c-.69 0-1.25-.447-1.25-1s.56-1 1.25-1zm-1.25-2c0-.553.56-1 1.25-1c.689 0 1.25.447 1.25 1c0 .553-.561 1-1.25 1c-.69 0-1.25-.447-1.25-1zM10 15c.689 0 1.25.447 1.25 1s-.561 1-1.25 1c-.69 0-1.25-.447-1.25-1s.56-1 1.25-1zm3-3c.689 0 1.25.447 1.25 1s-.561 1-1.25 1c-.69 0-1.25-.447-1.25-1s.56-1 1.25-1zm-1.25-2c0-.553.56-1 1.25-1c.689 0 1.25.447 1.25 1c0 .553-.561 1-1.25 1c-.69 0-1.25-.447-1.25-1zM13 15c.689 0 1.25.447 1.25 1s-.561 1-1.25 1c-.69 0-1.25-.447-1.25-1s.56-1 1.25-1zM5 7V4h10v3H5z"/>`), g.Group(children),
	)
}

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 3h-1v2h-3V3H7v2H4V3H3c-1.101 0-2 .9-2 2v12c0 1.1.899 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 14H3V9h14v8zM6.5 1h-2v3.5h2V1zm9 0h-2v3.5h2V1z"/>`), g.Group(children),
	)
}

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 8a3 3 0 1 0 0 6a3 3 0 0 0 0-6zm8-3h-2.4a.888.888 0 0 1-.789-.57l-.621-1.861A.89.89 0 0 0 13.4 2H6.6c-.33 0-.686.256-.789.568L5.189 4.43A.889.889 0 0 1 4.4 5H2C.9 5 0 5.9 0 7v9c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V7c0-1.1-.9-2-2-2zm-8 11a5 5 0 0 1-5-5a5 5 0 1 1 10 0a5 5 0 0 1-5 5zm7.5-7.8a.7.7 0 1 1 0-1.4a.7.7 0 0 1 0 1.4z"/>`), g.Group(children),
	)
}

func Ccw(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M.685 10h2.372v-.205c.108-4.434 3.724-7.996 8.169-7.996c4.515 0 8.174 3.672 8.174 8.201s-3.659 8.199-8.174 8.199a8.13 8.13 0 0 1-5.033-1.738l1.406-1.504a6.099 6.099 0 0 0 3.627 1.193c3.386 0 6.131-2.754 6.131-6.15c0-3.396-2.745-6.15-6.131-6.15c-3.317 0-6.018 2.643-6.125 5.945V10h2.672l-3.494 3.894L.685 10z"/>`), g.Group(children),
	)
}

func Chat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.8 12.2V6H2C.9 6 0 6.9 0 8v6c0 1.1.9 2 2 2h1v3l3-3h5c1.1 0 2-.9 2-2v-1.82a.943.943 0 0 1-.2.021h-7V12.2zM18 1H9c-1.1 0-2 .9-2 2v8h7l3 3v-3h1c1.1 0 2-.899 2-2V3c0-1.1-.9-2-2-2z"/>`), g.Group(children),
	)
}

func Check(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.294 16.998c-.435 0-.847-.203-1.111-.553L3.61 11.724a1.392 1.392 0 0 1 .27-1.951a1.392 1.392 0 0 1 1.953.27l2.351 3.104l5.911-9.492a1.396 1.396 0 0 1 1.921-.445c.653.406.854 1.266.446 1.92L9.478 16.34a1.39 1.39 0 0 1-1.12.656c-.022.002-.042.002-.064.002z"/>`), g.Group(children),
	)
}

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.516 7.548c.436-.446 1.043-.481 1.576 0L10 11.295l3.908-3.747c.533-.481 1.141-.446 1.574 0c.436.445.408 1.197 0 1.615c-.406.418-4.695 4.502-4.695 4.502a1.095 1.095 0 0 1-1.576 0S4.924 9.581 4.516 9.163c-.409-.418-.436-1.17 0-1.615z"/>`), g.Group(children),
	)
}

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.452 4.516c.446.436.481 1.043 0 1.576L8.705 10l3.747 3.908c.481.533.446 1.141 0 1.574c-.445.436-1.197.408-1.615 0c-.418-.406-4.502-4.695-4.502-4.695a1.095 1.095 0 0 1 0-1.576s4.084-4.287 4.502-4.695c.418-.409 1.17-.436 1.615 0z"/>`), g.Group(children),
	)
}

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.163 4.516c.418.408 4.502 4.695 4.502 4.695a1.095 1.095 0 0 1 0 1.576s-4.084 4.289-4.502 4.695c-.418.408-1.17.436-1.615 0c-.446-.434-.481-1.041 0-1.574L11.295 10L7.548 6.092c-.481-.533-.446-1.141 0-1.576c.445-.436 1.197-.409 1.615 0z"/>`), g.Group(children),
	)
}

func ChevronSmallDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.418 7.859a.695.695 0 0 1 .978 0a.68.68 0 0 1 0 .969l-3.908 3.83a.697.697 0 0 1-.979 0l-3.908-3.83a.68.68 0 0 1 0-.969a.695.695 0 0 1 .978 0L10 11l3.418-3.141z"/>`), g.Group(children),
	)
}

func ChevronSmallLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.141 13.418a.695.695 0 0 1 0 .978a.68.68 0 0 1-.969 0l-3.83-3.908a.697.697 0 0 1 0-.979l3.83-3.908a.68.68 0 0 1 .969 0a.695.695 0 0 1 0 .978L9 10l3.141 3.418z"/>`), g.Group(children),
	)
}

func ChevronSmallRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 10L7.859 6.58a.695.695 0 0 1 0-.978a.68.68 0 0 1 .969 0l3.83 3.908a.697.697 0 0 1 0 .979l-3.83 3.908a.68.68 0 0 1-.969 0a.695.695 0 0 1 0-.978L11 10z"/>`), g.Group(children),
	)
}

func ChevronSmallUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.582 12.141a.695.695 0 0 1-.978 0a.68.68 0 0 1 0-.969l3.908-3.83a.697.697 0 0 1 .979 0l3.908 3.83a.68.68 0 0 1 0 .969a.697.697 0 0 1-.979 0L10 9l-3.418 3.141z"/>`), g.Group(children),
	)
}

func ChevronThinDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.418 6.109a.697.697 0 0 1 .979 0a.68.68 0 0 1 0 .969l-7.908 7.83a.697.697 0 0 1-.979 0l-7.908-7.83a.68.68 0 0 1 0-.969a.697.697 0 0 1 .979 0L10 13.25l7.418-7.141z"/>`), g.Group(children),
	)
}

func ChevronThinLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.891 17.418a.697.697 0 0 1 0 .979a.68.68 0 0 1-.969 0l-7.83-7.908a.697.697 0 0 1 0-.979l7.83-7.908a.68.68 0 0 1 .969 0a.697.697 0 0 1 0 .979L6.75 10l7.141 7.418z"/>`), g.Group(children),
	)
}

func ChevronThinRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.25 10L6.109 2.58a.697.697 0 0 1 0-.979a.68.68 0 0 1 .969 0l7.83 7.908a.697.697 0 0 1 0 .979l-7.83 7.908a.68.68 0 0 1-.969 0a.697.697 0 0 1 0-.979L13.25 10z"/>`), g.Group(children),
	)
}

func ChevronThinUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.582 13.891c-.272.268-.709.268-.979 0s-.271-.701 0-.969l7.908-7.83a.697.697 0 0 1 .979 0l7.908 7.83a.68.68 0 0 1 0 .969a.695.695 0 0 1-.978 0L10 6.75l-7.418 7.141z"/>`), g.Group(children),
	)
}

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.484 12.452c-.436.446-1.043.481-1.576 0L10 8.705l-3.908 3.747c-.533.481-1.141.446-1.574 0c-.436-.445-.408-1.197 0-1.615c.406-.418 4.695-4.502 4.695-4.502a1.095 1.095 0 0 1 1.576 0s4.287 4.084 4.695 4.502c.409.418.436 1.17 0 1.615z"/>`), g.Group(children),
	)
}

func ChevronWithCircleDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.505 8.698L10 11L7.494 8.698a.512.512 0 0 0-.718 0a.5.5 0 0 0 0 .71l2.864 2.807a.51.51 0 0 0 .717 0l2.864-2.807a.498.498 0 0 0 0-.71a.51.51 0 0 0-.716 0zM10 .4A9.6 9.6 0 0 0 .4 10c0 5.303 4.298 9.6 9.6 9.6s9.6-4.297 9.6-9.6A9.6 9.6 0 0 0 10 .4zm0 17.954A8.353 8.353 0 0 1 1.646 10A8.354 8.354 0 1 1 10 18.354z"/>`), g.Group(children),
	)
}

func ChevronWithCircleLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.302 6.776a.5.5 0 0 0-.71 0L7.785 9.641a.51.51 0 0 0 0 .717l2.807 2.864a.498.498 0 0 0 .71 0a.51.51 0 0 0 0-.717L9 10l2.302-2.506a.512.512 0 0 0 0-.718zM10 .4A9.6 9.6 0 0 0 .4 10c0 5.303 4.298 9.6 9.6 9.6s9.6-4.297 9.6-9.6A9.6 9.6 0 0 0 10 .4zm0 17.954A8.353 8.353 0 0 1 1.646 10A8.354 8.354 0 1 1 10 18.354z"/>`), g.Group(children),
	)
}

func ChevronWithCircleRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 10L8.698 7.494a.512.512 0 0 1 0-.718a.5.5 0 0 1 .71 0l2.807 2.864a.51.51 0 0 1 0 .717l-2.807 2.864a.498.498 0 0 1-.71 0a.51.51 0 0 1 0-.717L11 10zM10 .4a9.6 9.6 0 0 1 9.6 9.6c0 5.303-4.298 9.6-9.6 9.6S.4 15.303.4 10A9.6 9.6 0 0 1 10 .4zm0 17.954a8.354 8.354 0 1 0 0-16.709a8.354 8.354 0 0 0 0 16.709z"/>`), g.Group(children),
	)
}

func ChevronWithCircleUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.359 7.785a.51.51 0 0 0-.717 0l-2.864 2.807a.498.498 0 0 0 0 .71a.51.51 0 0 0 .717 0L10 9l2.506 2.302a.512.512 0 0 0 .718 0a.5.5 0 0 0 0-.71l-2.865-2.807zM10 .4A9.6 9.6 0 0 0 .4 10c0 5.303 4.298 9.6 9.6 9.6s9.6-4.297 9.6-9.6A9.6 9.6 0 0 0 10 .4zm0 17.954A8.353 8.353 0 0 1 1.646 10A8.354 8.354 0 1 1 10 18.354z"/>`), g.Group(children),
	)
}

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4A9.6 9.6 0 0 0 .4 10a9.6 9.6 0 1 0 19.2-.001C19.6 4.698 15.301.4 10 .4zm0 17.199A7.6 7.6 0 1 1 10 2.4a7.6 7.6 0 1 1 0 15.199z"/>`), g.Group(children),
	)
}

func CircleWithCross(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 1.6a8.4 8.4 0 1 0 0 16.8a8.4 8.4 0 0 0 0-16.8zm4.789 11.461L13.06 14.79L10 11.729l-3.061 3.06L5.21 13.06L8.272 10L5.211 6.939L6.94 5.211L10 8.271l3.061-3.061l1.729 1.729L11.728 10l3.061 3.061z"/>`), g.Group(children),
	)
}

func CircleWithMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 1.6a8.4 8.4 0 1 0 0 16.8a8.4 8.4 0 0 0 0-16.8zm5 9.4H5V9h10v2z"/>`), g.Group(children),
	)
}

func CircleWithPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 1.6a8.4 8.4 0 1 0 0 16.8a8.4 8.4 0 0 0 0-16.8zm5 9.4h-4v4H9v-4H5V9h4V5h2v4h4v2z"/>`), g.Group(children),
	)
}

func CircularGraph(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.584 9.372h2a9.554 9.554 0 0 0-.668-2.984L17.16 7.402c.224.623.371 1.283.424 1.97zm-3.483-8.077a9.492 9.492 0 0 0-3.086-.87v2.021a7.548 7.548 0 0 1 2.084.585l1.002-1.736zm2.141 4.327l1.741-1.005a9.643 9.643 0 0 0-2.172-2.285l-1.006 1.742a7.625 7.625 0 0 1 1.437 1.548zm-6.228 11.949a7.6 7.6 0 0 1-7.6-7.6c0-3.858 2.877-7.036 6.601-7.526V.424C4.182.924.414 5.007.414 9.971a9.6 9.6 0 0 0 9.601 9.601c4.824 0 8.807-3.563 9.486-8.2H17.48c-.658 3.527-3.748 6.199-7.466 6.199z"/>`), g.Group(children),
	)
}

func Clapperboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M20 3v14a1 1 0 0 1-1 1H1a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h1l3 3h2.5l-3-3h3l3 3H13l-3-3h3l3 3h2.5l-3-3H19a1 1 0 0 1 1 1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ClassicComputer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 0H4C2.9 0 2 .899 2 2v15a1 1 0 0 0 1 1v2h14v-2a1 1 0 0 0 1-1V2c0-1.101-.899-2-2-2zm-2 15h-4v-1h4v1zm1-4H5V3h10v8z"/>`), g.Group(children),
	)
}

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15.6 2l-1.2 3H5.6L4.4 2C3.629 2 3 2.629 3 3.4v15.2c0 .77.629 1.4 1.399 1.4h11.2c.77 0 1.4-.631 1.4-1.4V3.4C17 2.629 16.369 2 15.6 2zm-2 2l.9-2h-2.181L11.6 0H8.4l-.72 2H5.5l.899 2H13.6z"/>`), g.Group(children),
	)
}

func Clock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4C4.697.4.399 4.698.399 10A9.6 9.6 0 0 0 10 19.601c5.301 0 9.6-4.298 9.6-9.601c0-5.302-4.299-9.6-9.6-9.6zm-.001 17.2a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 1 1 0 15.2zM11 9.33V4H9v6.245l-3.546 2.048l1 1.732l4.115-2.377A.955.955 0 0 0 11 10.9v-.168l4.24-4.166a6.584 6.584 0 0 0-.647-.766L11 9.33z"/>`), g.Group(children),
	)
}

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 11.32c0 2.584-2.144 4.68-4.787 4.68H3.617C1.619 16 0 14.416 0 12.463c0-1.951 1.619-3.535 3.617-3.535c.146 0 .288.012.429.027a5.076 5.076 0 0 1-.057-.756C3.989 5.328 6.37 3 9.309 3c2.407 0 4.439 1.562 5.096 3.707a5 5 0 0 1 .809-.066C17.856 6.641 20 8.734 20 11.32z"/>`), g.Group(children),
	)
}

func Code(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.719 14.75a.997.997 0 0 1-.664-.252L-.005 10l5.341-4.748a1 1 0 0 1 1.328 1.495L3.005 10l3.378 3.002a1 1 0 0 1-.664 1.748zm8.945-.002L20.005 10l-5.06-4.498a.999.999 0 1 0-1.328 1.495L16.995 10l-3.659 3.252a1 1 0 0 0 1.328 1.496zm-4.678 1.417l2-12a1 1 0 1 0-1.972-.329l-2 12a1 1 0 1 0 1.972.329z"/>`), g.Group(children),
	)
}

func Cog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.783 10c0-1.049.646-1.875 1.617-2.443a8.932 8.932 0 0 0-.692-1.672c-1.089.285-1.97-.141-2.711-.883c-.741-.74-.968-1.621-.683-2.711a8.732 8.732 0 0 0-1.672-.691c-.568.97-1.595 1.615-2.642 1.615c-1.048 0-2.074-.645-2.643-1.615a8.697 8.697 0 0 0-1.671.691c.285 1.09.059 1.971-.684 2.711c-.74.742-1.621 1.168-2.711.883A8.797 8.797 0 0 0 1.6 7.557c.97.568 1.615 1.394 1.615 2.443c0 1.047-.645 2.074-1.615 2.643a8.89 8.89 0 0 0 .691 1.672c1.09-.285 1.971-.059 2.711.682c.741.742.969 1.623.684 2.711a8.841 8.841 0 0 0 1.672.693c.568-.973 1.595-1.617 2.643-1.617c1.047 0 2.074.645 2.643 1.617a8.963 8.963 0 0 0 1.672-.693c-.285-1.088-.059-1.969.683-2.711c.741-.74 1.622-1.166 2.711-.883a8.811 8.811 0 0 0 .692-1.672c-.973-.569-1.619-1.395-1.619-2.442zM10 13.652a3.652 3.652 0 1 1 0-7.306a3.653 3.653 0 0 1 0 7.306z"/>`), g.Group(children),
	)
}

func Colours(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.179 5.998a1.005 1.005 0 0 0-1.408.132L.494 7.669a1.004 1.004 0 0 0 .131 1.407l7.888 6.542l-3.807-8.354l-1.527-1.266zm3.834-3.315l-1.82.829a1.005 1.005 0 0 0-.495 1.324l4.25 9.325l.213-9.179l-.822-1.804c-.23-.5-.826-.723-1.326-.495zm7.198.204a1.003 1.003 0 0 0-.976-1.023l-2-.046a1.003 1.003 0 0 0-1.022.976l-.239 10.243l4.19-8.167l.047-1.983zm4.98.95l-1.779-.913a1.005 1.005 0 0 0-1.347.434L9.674 15.814a1.004 1.004 0 0 0 .434 1.347l1.779.913a1.003 1.003 0 0 0 1.346-.433l6.391-12.456a1.005 1.005 0 0 0-.433-1.348zm-6.392 12.456a1 1 0 1 1-1.78-.911a1 1 0 0 1 1.78.911z"/>`), g.Group(children),
	)
}

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.454 14.548s4.568-.627 6.518-2.576s2.576-6.518 2.576-6.518s-4.569.627-6.518 2.576s-2.576 6.518-2.576 6.518zm3.563-5.533c.818-.818 2.385-1.4 3.729-1.762c-.361 1.342-.945 2.92-1.76 3.732a1.39 1.39 0 0 1-1.969 0a1.391 1.391 0 0 1 0-1.97zM10.001.4C4.698.4.4 4.698.4 10a9.6 9.6 0 0 0 9.601 9.601c5.301 0 9.6-4.298 9.6-9.601c0-5.302-4.299-9.6-9.6-9.6zM10 17.6a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 1 1 0 15.2z"/>`), g.Group(children),
	)
}

func ControllerFastBackward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.959 4.571L10.756 9.52s-.279.201-.279.481s.279.479.279.479l7.203 4.951c.572.38 1.041.099 1.041-.626V5.196c0-.727-.469-1.008-1.041-.625zm-9.076 0L1.68 9.52s-.279.201-.279.481s.279.479.279.479l7.203 4.951c.572.381 1.041.1 1.041-.625v-9.61c0-.727-.469-1.008-1.041-.625z"/>`), g.Group(children),
	)
}

func ControllerFastForward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.244 9.52L2.041 4.571C1.469 4.188 1 4.469 1 5.196v9.609c0 .725.469 1.006 1.041.625l7.203-4.951s.279-.199.279-.478c0-.28-.279-.481-.279-.481zm9.356.481c0 .279-.279.478-.279.478l-7.203 4.951c-.572.381-1.041.1-1.041-.625V5.196c0-.727.469-1.008 1.041-.625L18.32 9.52s.28.201.28.481z"/>`), g.Group(children),
	)
}

func ControllerJumpToStart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.959 4.571L7.756 9.52s-.279.201-.279.481s.279.479.279.479l7.203 4.951c.572.38 1.041.099 1.041-.626V5.196c0-.727-.469-1.008-1.041-.625zM6 4H5c-.553 0-1 .048-1 .6v10.8c0 .552.447.6 1 .6h1c.553 0 1-.048 1-.6V4.6c0-.552-.447-.6-1-.6z"/>`), g.Group(children),
	)
}

func ControllerNext(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.244 9.52L5.041 4.571C4.469 4.188 4 4.469 4 5.196v9.609c0 .725.469 1.006 1.041.625l7.203-4.951s.279-.199.279-.478c0-.28-.279-.481-.279-.481zM14 4h1c.553 0 1 .048 1 .6v10.8c0 .552-.447.6-1 .6h-1c-.553 0-1-.048-1-.6V4.6c0-.552.447-.6 1-.6z"/>`), g.Group(children),
	)
}

func ControllerPaus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 3h-2c-.553 0-1 .048-1 .6v12.8c0 .552.447.6 1 .6h2c.553 0 1-.048 1-.6V3.6c0-.552-.447-.6-1-.6zM7 3H5c-.553 0-1 .048-1 .6v12.8c0 .552.447.6 1 .6h2c.553 0 1-.048 1-.6V3.6c0-.552-.447-.6-1-.6z"/>`), g.Group(children),
	)
}

func ControllerPlay(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 10.001c0 .299-.305.514-.305.514l-8.561 5.303C5.51 16.227 5 15.924 5 15.149V4.852c0-.777.51-1.078 1.135-.67l8.561 5.305c-.001 0 .304.215.304.514z"/>`), g.Group(children),
	)
}

func ControllerRecord(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 3a7 7 0 1 0 .001 13.999A7 7 0 0 0 10 3z"/>`), g.Group(children),
	)
}

func ControllerStop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 4.995v9.808c0 .661-.536 1.197-1.196 1.197H4.997A.997.997 0 0 1 4 15.003V5.196C4 4.536 4.536 4 5.196 4h9.808c.55 0 .996.446.996.995z"/>`), g.Group(children),
	)
}

func ControllerVolume(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 13.805c0 .657-.538 1.195-1.195 1.195H1.533c-.88 0-.982-.371-.229-.822l16.323-9.055C18.382 4.67 19 5.019 19 5.9v7.905z"/>`), g.Group(children),
	)
}

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 0H3a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h5v2h2v-2H8.001v-2H10v-2H8v2H4V2h6v4h2V1a1 1 0 0 0-1-1zM8 7v1h2V6H9a1 1 0 0 0-1 1zm4 13h2v-2h-2v2zm0-12h2V6h-2v2zM8 19a1 1 0 0 0 1 1h1v-2H8v1zm9-13h-1v2h2V7a1 1 0 0 0-1-1zm-1 14h1a1 1 0 0 0 1-1v-1h-2v2zm0-8h2v-2h-2v2zm0 4h2v-2h-2v2z"/>`), g.Group(children),
	)
}

func CreativeCommons(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.651 11.628c-.406 0-.713-.149-.92-.445c-.209-.295-.312-.69-.312-1.182c0-1.084.41-1.627 1.232-1.627c.164 0 .342.055.534.164s.353.3.484.574l1.232-.64c-.492-.887-1.308-1.33-2.447-1.33c-.778 0-1.422.257-1.93.771c-.51.516-.765 1.211-.765 2.088c0 .898.253 1.6.756 2.103c.504.504 1.168.754 1.988.754a2.697 2.697 0 0 0 2.416-1.445l-1.135-.574c-.219.525-.597.789-1.133.789zm5.307 0c-.406 0-.713-.149-.92-.445c-.209-.295-.312-.69-.312-1.182c0-1.084.41-1.627 1.232-1.627c.174 0 .357.055.549.164c.192.11.353.3.486.574l1.215-.64c-.482-.887-1.293-1.33-2.432-1.33c-.777 0-1.42.257-1.93.771c-.509.516-.763 1.211-.763 2.088c0 .898.248 1.6.747 2.103c.498.504 1.163.754 1.996.754c.503 0 .97-.129 1.396-.384a2.831 2.831 0 0 0 1.02-1.06l-1.151-.575c-.219.525-.598.789-1.133.789zm3.855-8.444C14.954 1.328 12.679.4 9.987.4c-2.659 0-4.91.928-6.752 2.784C1.345 5.104.4 7.375.4 10c0 2.624.945 4.88 2.835 6.768c1.89 1.888 4.142 2.832 6.752 2.832c2.644 0 4.935-.952 6.874-2.856C18.687 14.936 19.6 12.688 19.6 10c0-2.688-.93-4.96-2.787-6.816zM15.61 15.496c-1.586 1.568-3.452 2.352-5.6 2.352c-2.146 0-3.996-.776-5.55-2.329c-1.554-1.551-2.33-3.391-2.33-5.52c0-2.127.784-3.983 2.354-5.567C5.99 2.896 7.832 2.127 10.01 2.127c2.18 0 4.03.769 5.552 2.305C17.101 5.952 17.87 7.807 17.87 10c0 2.208-.753 4.04-2.259 5.496z"/>`), g.Group(children),
	)
}

func CreativeCommonsAttribution(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.583 7.623a.62.62 0 0 0-.62-.62H8.037a.62.62 0 0 0-.62.62v3.927h1.095v4.65h2.976v-4.65h1.095V7.623z"/><circle cx="10" cy="5.143" r="1.343" fill="currentColor"/><path fill="currentColor" fill-rule="evenodd" d="M9.988.4c-2.66 0-4.91.928-6.753 2.784C1.345 5.104.4 7.376.4 10c0 2.624.945 4.88 2.835 6.768c1.89 1.888 4.142 2.832 6.753 2.832c2.643 0 4.934-.952 6.872-2.856c1.827-1.808 2.74-4.056 2.74-6.744s-.93-4.96-2.788-6.816C14.954 1.328 12.68.4 9.988.4zm.024 1.728c2.179 0 4.029.768 5.55 2.304c1.54 1.52 2.308 3.376 2.308 5.568c0 2.208-.752 4.04-2.258 5.496c-1.586 1.568-3.453 2.352-5.6 2.352c-2.146 0-3.996-.776-5.55-2.328C2.907 13.968 2.13 12.128 2.13 10c0-2.128.785-3.984 2.355-5.568c1.506-1.536 3.348-2.304 5.527-2.304z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CreativeCommonsNoderivs(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.988.4c-2.66 0-4.91.928-6.753 2.784C1.345 5.104.4 7.376.4 10c0 2.624.945 4.88 2.835 6.768c1.89 1.888 4.142 2.832 6.753 2.832c2.643 0 4.934-.952 6.872-2.856c1.827-1.808 2.74-4.056 2.74-6.744c0-2.688-.93-4.96-2.788-6.816C14.954 1.328 12.68.4 9.988.4zm.024 1.727c2.179 0 4.029.769 5.55 2.305c1.54 1.52 2.308 3.375 2.308 5.568c0 2.208-.752 4.04-2.258 5.496c-1.586 1.568-3.453 2.351-5.6 2.351c-2.146 0-3.996-.775-5.55-2.328C2.907 13.967 2.13 12.128 2.13 10c0-2.129.785-3.984 2.355-5.568c1.506-1.536 3.348-2.305 5.527-2.305z" clip-rule="evenodd"/><path fill="currentColor" fill-rule="evenodd" d="M13.627 7.725H6.649v1.653h6.978V7.725zm0 3.085H6.649v1.653h6.978V10.81z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func CreativeCommonsNoncommercialEu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.812 3.184C14.954 1.328 12.68.4 9.988.4c-2.66 0-4.91.928-6.753 2.784C1.345 5.104.4 7.376.4 10c0 2.624.945 4.88 2.835 6.768c1.89 1.888 4.141 2.832 6.753 2.832c2.643 0 4.934-.952 6.872-2.856c1.827-1.808 2.74-4.056 2.74-6.744c0-2.688-.93-4.96-2.788-6.816zM15.61 15.496c-1.586 1.568-3.452 2.351-5.598 2.351c-2.147 0-3.998-.775-5.551-2.327C2.907 13.967 2.13 12.128 2.13 10c0-.9.142-1.75.423-2.553l2.544 1.126h-.184v1.14h.9c0 .162-.016.322-.016.483v.274h-.884v1.14h1.045c.145.852.466 1.543.9 2.09c.9 1.19 2.346 1.832 3.921 1.832c1.03 0 1.96-.305 2.507-.611l-.386-1.784c-.337.177-1.092.418-1.831.418c-.804 0-1.56-.24-2.074-.82c-.24-.273-.417-.642-.53-1.125h3.494l4.965 2.198a7.407 7.407 0 0 1-1.314 1.688zM9.367 10.47l-.022-.016l.037.016h-.015zm2.988-.756h.144V8.573H9.778l-1.105-.49a2.19 2.19 0 0 1 .339-.554c.498-.612 1.205-.869 1.976-.869c.708 0 1.366.21 1.784.386l.45-1.832c-.579-.257-1.43-.482-2.41-.482c-1.511 0-2.797.61-3.713 1.639a4.96 4.96 0 0 0-.54.778L3.402 5.75a8.83 8.83 0 0 1 1.083-1.32c1.506-1.535 3.348-2.304 5.527-2.304c2.178 0 4.029.769 5.55 2.305C17.1 5.952 17.87 7.807 17.87 10c0 .722-.082 1.404-.244 2.046l-5.27-2.332z"/>`), g.Group(children),
	)
}

func CreativeCommonsNoncommercialUs(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.988.4c2.69 0 4.966.928 6.825 2.784C18.67 5.04 19.6 7.312 19.6 10s-.913 4.936-2.74 6.744c-1.938 1.904-4.23 2.856-6.872 2.856c-2.611 0-4.862-.944-6.753-2.832C1.345 14.88.4 12.624.4 10s.945-4.896 2.835-6.816C5.078 1.328 7.33.4 9.988.4zM2.56 7.421c-.287.81-.43 1.67-.43 2.579c0 2.128.777 3.968 2.33 5.52c1.555 1.552 3.405 2.328 5.552 2.328s4.013-.784 5.599-2.352a7.41 7.41 0 0 0 1.311-1.68l-3.618-1.611c-.246 1.217-1.331 2.04-2.643 2.137v1.48H9.559v-1.48c-1.077-.013-2.118-.453-2.914-1.15l1.322-1.333c.637.6 1.274.869 2.143.869c.563 0 1.188-.22 1.188-.955c0-.26-.1-.44-.26-.577l-.915-.407l-1.14-.507c-.563-.252-1.04-.464-1.52-.677L2.56 7.42zm7.452-5.293c-2.179 0-4.021.768-5.527 2.304c-.41.414-.766.846-1.07 1.297l3.67 1.633c.332-1.017 1.3-1.635 2.474-1.704v-1.48h1.102v1.48a4.14 4.14 0 0 1 2.412.88l-1.26 1.297c-.466-.33-1.054-.563-1.642-.563c-.477 0-1.15.148-1.15.747c0 .091.03.172.085.243l1.228.547l.83.37l1.542.686l4.92 2.19c.162-.644.244-1.33.244-2.055c0-2.192-.769-4.048-2.307-5.568c-1.522-1.536-3.372-2.304-5.551-2.304z"/>`), g.Group(children),
	)
}

func CreativeCommonsPublicDomain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M9.987.4c-2.659 0-4.91.928-6.752 2.784C1.345 5.104.4 7.376.4 10c0 2.624.945 4.88 2.835 6.768c1.89 1.888 4.141 2.832 6.752 2.832c2.644 0 4.935-.952 6.873-2.856c1.826-1.808 2.74-4.056 2.74-6.744c0-2.688-.93-4.96-2.788-6.816C14.954 1.328 12.68.4 9.987.4zm.025 1.728c2.178 0 4.028.768 5.55 2.304c1.54 1.52 2.308 3.376 2.308 5.568c0 2.208-.752 4.04-2.259 5.496c-1.586 1.568-3.452 2.352-5.598 2.352c-2.147 0-3.997-.776-5.551-2.328C2.907 13.968 2.13 12.128 2.13 10c0-2.128.785-3.984 2.355-5.568c1.506-1.536 3.348-2.304 5.527-2.304z"/><path d="M6.82 8.279H7.9c.604 0 .857.29.857.82c0 .41-.216.804-.82.804H6.82V8.279zM5.493 12.6h1.342v-1.714H7.87c1.819 0 2.214-1.088 2.214-1.796c0-1.05-.53-1.796-1.96-1.796H5.493V12.6zm4.952 0h2.057c1.617 0 2.705-.767 2.705-2.652c0-1.886-1.088-2.653-2.705-2.653h-2.057V12.6zm1.326-4.187h.753c.999 0 1.357.678 1.357 1.535c0 .857-.358 1.535-1.35 1.535h-.76v-3.07z"/></g>`), g.Group(children),
	)
}

func CreativeCommonsRemix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.987.4c-2.658 0-4.91.928-6.752 2.784C1.345 5.104.4 7.376.4 10c0 2.624.945 4.88 2.835 6.767c1.89 1.889 4.142 2.833 6.752 2.833c2.644 0 4.935-.952 6.874-2.856C18.687 14.936 19.6 12.688 19.6 10c0-2.688-.93-4.96-2.787-6.816C14.954 1.328 12.679.4 9.987.4zm.024 1.728c2.18 0 4.03.768 5.552 2.304C17.101 5.952 17.87 7.807 17.87 10c0 2.208-.753 4.04-2.259 5.496c-1.586 1.568-3.452 2.352-5.6 2.352c-2.146 0-3.996-.777-5.55-2.329c-1.554-1.551-2.33-3.392-2.33-5.519c0-2.128.784-3.984 2.354-5.568c1.506-1.536 3.348-2.304 5.526-2.304z"/><path fill="currentColor" d="m16.278 10.17l-1.91-.795v-2.38L8.495 4.55L5.878 5.7v2.723l.028.011l-2.309.961v2.468l2.468 1.057l2.48-1.026l.168.067l5.025 2.071l.112.047l.113-.047l2.325-.993l.178-.075v-2.72l-.188-.074zm-2.728 3.142l-.023-.01v.015l-4.401-1.823V9.758l4.4 1.799v.056l.024.01v1.69zm.297-2.217l-1.396-.598l1.503-.643l1.433.595l-1.54.646zm2.019 1.465l-1.716.74v-1.681l1.716-.722v1.663z"/>`), g.Group(children),
	)
}

func CreativeCommonsShare(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.987.4c-2.658 0-4.91.928-6.752 2.784C1.345 5.104.4 7.376.4 10c0 2.624.945 4.88 2.835 6.768c1.89 1.888 4.142 2.832 6.752 2.832c2.644 0 4.935-.952 6.874-2.856C18.687 14.936 19.6 12.688 19.6 10c0-2.688-.93-4.96-2.787-6.816C14.954 1.328 12.679.4 9.987.4zm.024 1.727c2.18 0 4.03.769 5.552 2.305C17.101 5.952 17.87 7.807 17.87 10c0 2.208-.753 4.04-2.259 5.496c-1.586 1.568-3.452 2.352-5.6 2.352c-2.146 0-3.996-.776-5.55-2.328c-1.554-1.552-2.33-3.392-2.33-5.52c0-2.128.784-3.984 2.354-5.568c1.506-1.536 3.348-2.305 5.526-2.305z"/><path fill="currentColor" d="M13.932 7.25h-2.1V5.159a.525.525 0 0 0-.525-.525H6.02a.525.525 0 0 0-.479.525v7.069c0 .29.235.525.525.525h2.1v2.09c0 .29.235.525.525.525h5.24c.29 0 .526-.235.526-.525V7.775a.525.525 0 0 0-.525-.525zm-7.34 4.453V5.684h4.19V7.25H8.635a.525.525 0 0 0-.469.525v3.928H6.591zM13.406 8.3v6.018h-4.19V8.3h4.19z"/>`), g.Group(children),
	)
}

func CreativeCommonsSharealike(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9.988.4c-2.66 0-4.91.928-6.753 2.784C1.345 5.104.4 7.376.4 10s.945 4.88 2.835 6.768c1.89 1.888 4.142 2.832 6.753 2.832c2.643 0 4.934-.952 6.872-2.856c1.827-1.808 2.74-4.056 2.74-6.744s-.93-4.96-2.788-6.816C14.954 1.328 12.68.4 9.988.4zm.024 1.728c2.179 0 4.029.768 5.55 2.304c1.54 1.52 2.308 3.376 2.308 5.568c0 2.208-.752 4.04-2.258 5.496c-1.586 1.568-3.453 2.352-5.6 2.352c-2.146 0-3.996-.776-5.55-2.328C2.907 13.968 2.13 12.128 2.13 10c0-2.128.785-3.984 2.355-5.568c1.506-1.536 3.348-2.304 5.527-2.304z" clip-rule="evenodd"/><path fill="currentColor" d="M5.733 8.645c.383-2.415 2.083-3.707 4.214-3.707c3.064 0 4.932 2.225 4.932 5.19c0 2.893-1.987 5.141-4.98 5.141c-2.06 0-3.903-1.267-4.238-3.754h2.418c.072 1.291.91 1.746 2.107 1.746c1.365 0 2.251-1.268 2.251-3.205c0-2.033-.766-3.109-2.203-3.109c-1.054 0-1.963.383-2.155 1.698l.703-.004l-1.903 1.904L4.975 8.64l.758.004z"/>`), g.Group(children),
	)
}

func Credit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 16.755V19H9v-2.143c-1.712-.1-3.066-.589-4.241-1.797l1.718-1.74c.859.87 2.023 1.16 3.282 1.16c1.565 0 2.405-.599 2.405-1.702c0-.483-.133-.889-.42-1.16c-.267-.251-.572-.387-1.202-.483L8.9 10.903c-1.164-.174-2.022-.541-2.634-1.141c-.648-.657-.973-1.546-.973-2.707c0-2.155 1.382-3.743 3.707-4.1V1h2v1.932c1.382.145 2.465.62 3.415 1.551l-1.679 1.682c-.859-.832-1.889-.947-2.787-.947c-1.412 0-2.099.792-2.099 1.74c0 .348.115.716.401.986c.267.252.706.464 1.26.541l1.602.232c1.241.174 2.023.522 2.596 1.063c.726.696 1.05 1.702 1.05 2.92c0 2.25-1.567 3.662-3.759 4.055z"/>`), g.Group(children),
	)
}

func CreditCard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 3H2C.899 3 0 3.9 0 5v10c0 1.1.899 2 2 2h16c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 12H2V9h16v6zm0-9H2V5h16v1zM4 11.1v.6h.6v-.6H4zm3.6 1.199v.601h1.2v-.601h.6v-.6h.6v-.6H8.8v.6h-.601v.6H7.6zm2.4.601v-.601h-.601v.601H10zm-3 0v-.601H5.8v.601H7zm.6-1.201h.6v-.6H7v1.199h.6v-.599zm-2.401.6H5.8v-.6h.6v-.6H5.2v.6h-.6v.6H4v.601h1.199v-.601z"/>`), g.Group(children),
	)
}

func Cross(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.348 14.849a1.2 1.2 0 0 1-1.697 0L10 11.819l-2.651 3.029a1.2 1.2 0 1 1-1.697-1.697l2.758-3.15l-2.759-3.152a1.2 1.2 0 1 1 1.697-1.697L10 8.183l2.651-3.031a1.2 1.2 0 1 1 1.697 1.697l-2.758 3.152l2.758 3.15a1.2 1.2 0 0 1 0 1.698z"/>`), g.Group(children),
	)
}

func Cup(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 1C5.721 1 3.06 2.41 3.205 3.555l1.442 13.467c.058.46 2.221 1.976 5.353 1.978c3.131-.002 5.295-1.518 5.351-1.979l1.442-13.467C16.938 2.41 14.279 1 10 1zm0 4.291c-3.132-.002-5.353-1.117-5.353-1.535C4.646 3.342 6.869 2.225 10 2.227c3.131-.002 5.354 1.115 5.351 1.529c0 .418-2.22 1.533-5.351 1.535z"/>`), g.Group(children),
	)
}

func Cw(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.315 10h-2.372v-.205c-.108-4.434-3.724-7.996-8.169-7.996C4.259 1.799.6 5.471.6 10s3.659 8.199 8.174 8.199a8.13 8.13 0 0 0 5.033-1.738l-1.406-1.504a6.099 6.099 0 0 1-3.627 1.193c-3.386 0-6.131-2.754-6.131-6.15s2.745-6.15 6.131-6.15c3.317 0 6.018 2.643 6.125 5.945V10h-2.672l3.494 3.894L19.315 10z"/>`), g.Group(children),
	)
}

func Cycle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.516 14.224c-2.262-2.432-2.222-6.244.128-8.611a6.074 6.074 0 0 1 3.414-1.736L8.989 1.8a8.112 8.112 0 0 0-4.797 2.351c-3.149 3.17-3.187 8.289-.123 11.531l-1.741 1.752l5.51.301l-.015-5.834l-2.307 2.323zm6.647-11.959l.015 5.834l2.307-2.322c2.262 2.434 2.222 6.246-.128 8.611a6.07 6.07 0 0 1-3.414 1.736l.069 2.076a8.122 8.122 0 0 0 4.798-2.35c3.148-3.172 3.186-8.291.122-11.531l1.741-1.754l-5.51-.3z"/>`), g.Group(children),
	)
}

func Database(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.726 12.641c-.843 1.363-3.535 2.361-6.726 2.361s-5.883-.998-6.727-2.361c-.178-.29-.273-.135-.273.007v2.002c0 1.94 3.134 3.95 7 3.95s7-2.01 7-3.949v-2.002c0-.143-.096-.298-.274-.008zm.011-5.116c-.83 1.205-3.532 2.09-6.737 2.09s-5.908-.885-6.738-2.09C3.091 7.277 3 7.412 3 7.523V9.88c0 1.762 3.134 3.189 7 3.189s7-1.428 7-3.189V7.523c0-.111-.092-.246-.263.002zM10 1C6.134 1 3 2.18 3 3.633v1.26c0 1.541 3.134 2.791 7 2.791s7-1.25 7-2.791v-1.26C17 2.18 13.866 1 10 1z"/>`), g.Group(children),
	)
}

func DialPad(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6 0H4a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V1a1 1 0 0 0-1-1zm5 0H9a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V1a1 1 0 0 0-1-1zm5 0h-2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V1a1 1 0 0 0-1-1zM6 5H4a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1zm5 0H9a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1zm5 0h-2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1V6a1 1 0 0 0-1-1zM6 10H4a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1zm5 0H9a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1zm0 6H9a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1zm5-6h-2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func Direction(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.06 1.941c-.586-.586-1.144-.033-3.041.879C9.944 5.259 1.1 10.216 1.1 10.216L8.699 11.3l1.085 7.599s4.958-8.843 7.396-13.916c.912-1.898 1.465-2.456.88-3.042zm-1.824 1.955l-5.519 10.247l-.561-4.655l6.08-5.592z"/>`), g.Group(children),
	)
}

func Document(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 1H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zm-1 16H5V3h10v14z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func DocumentLandscape(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 3H1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zm-1 12H2V5h16v10z"/>`), g.Group(children),
	)
}

func Documents(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19.398 7.415l-7.444-1.996L10.651.558c-.109-.406-.545-.642-.973-.529L.602 2.461c-.428.114-.686.538-.577.944l3.23 12.051c.109.406.544.643.971.527l3.613-.967l-.492 1.838c-.109.406.149.83.577.943l8.11 2.174c.428.115.862-.121.972-.529l2.97-11.084c.108-.406-.15-.83-.578-.943zM1.633 3.631l7.83-2.096l2.898 10.818l-7.83 2.096L1.633 3.631zm14.045 14.832L8.864 16.6l.536-2.002l3.901-1.047c.428-.113.688-.537.578-.943l-1.508-5.627l5.947 1.631l-2.64 9.851z"/>`), g.Group(children),
	)
}

func DotSingle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.8 10a2.2 2.2 0 0 0 4.4 0a2.2 2.2 0 0 0-4.4 0z"/>`), g.Group(children),
	)
}

func DotsThreeHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.001 7.8a2.2 2.2 0 1 0 0 4.402A2.2 2.2 0 0 0 10 7.8zm-7 0a2.2 2.2 0 1 0 0 4.402A2.2 2.2 0 0 0 3 7.8zm14 0a2.2 2.2 0 1 0 0 4.402A2.2 2.2 0 0 0 17 7.8z"/>`), g.Group(children),
	)
}

func DotsThreeVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.001 7.8a2.2 2.2 0 1 0 0 4.402A2.2 2.2 0 0 0 10 7.8zm0-2.6A2.2 2.2 0 1 0 9.999.8a2.2 2.2 0 0 0 .002 4.4zm0 9.6a2.2 2.2 0 1 0 0 4.402a2.2 2.2 0 0 0 0-4.402z"/>`), g.Group(children),
	)
}

func DotsTwoHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.001 7.8a2.2 2.2 0 1 0 0 4.402A2.2 2.2 0 0 0 14 7.8zm-8 0a2.2 2.2 0 1 0 0 4.402A2.2 2.2 0 0 0 6 7.8z"/>`), g.Group(children),
	)
}

func DotsTwoVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.001 8.2a2.2 2.2 0 1 0-.002-4.4a2.2 2.2 0 0 0 .002 4.4zm0 3.6a2.2 2.2 0 1 0 0 4.402a2.2 2.2 0 0 0 0-4.402z"/>`), g.Group(children),
	)
}

func Download(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 7h-3V1H8v6H5l5 5l5-5zm4.338 6.532c-.21-.224-1.611-1.723-2.011-2.114A1.503 1.503 0 0 0 16.285 11h-1.757l3.064 2.994h-3.544a.274.274 0 0 0-.24.133L12.992 16H7.008l-.816-1.873a.276.276 0 0 0-.24-.133H2.408L5.471 11H3.715c-.397 0-.776.159-1.042.418c-.4.392-1.801 1.891-2.011 2.114c-.489.521-.758.936-.63 1.449l.561 3.074c.128.514.691.936 1.252.936h16.312c.561 0 1.124-.422 1.252-.936l.561-3.074c.126-.513-.142-.928-.632-1.449z"/>`), g.Group(children),
	)
}

func Drink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.538 2.639C17.932 2.094 18 1 18 1H2s.068 1.094.462 1.639L9 11v6H7c-2 0-2 2-2 2h10s0-2-2-2h-2v-6l6.538-8.361zM9.4 6a1.599 1.599 0 1 1 3.2 0a1.6 1.6 0 0 1-3.2 0z"/>`), g.Group(children),
	)
}

func Drive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19.059 10.898l-3.171-7.927A1.543 1.543 0 0 0 14.454 2H5.546c-.632 0-1.2.384-1.434.971L.941 10.898a4.25 4.25 0 0 0-.246 2.272l.59 3.539A1.544 1.544 0 0 0 2.808 18h14.383c.755 0 1.399-.546 1.523-1.291l.59-3.539a4.22 4.22 0 0 0-.245-2.272zm-2.1 4.347a.902.902 0 0 1-.891.755H3.932a.902.902 0 0 1-.891-.755l-.365-2.193A.902.902 0 0 1 3.567 12h12.867c.558 0 .983.501.891 1.052l-.366 2.193z"/>`), g.Group(children),
	)
}

func Drop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.203.561c-.027-.215-.38-.215-.406 0c-.883 7.107-5.398 8.572-5.398 13.512c0 3.053 2.564 5.527 5.601 5.527c3.036 0 5.6-2.475 5.6-5.527c0-4.94-4.514-6.405-5.397-13.512zM9.35 8.418c-.059.219-.123.444-.189.678c-.401 1.424-.856 3.039-.856 4.906c0 1.012-.598 1.371-1.156 1.371a1.161 1.161 0 0 1-1.156-1.166c0-2.207 1.062-3.649 2-4.92c.295-.398.572-.775.797-1.15c.103-.172.38-.164.506.006c.059.08.079.182.054.275z"/>`), g.Group(children),
	)
}

func Edit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.561 2.439c-1.442-1.443-2.525-1.227-2.525-1.227L8.984 7.264L2.21 14.037L1.2 18.799l4.763-1.01l6.774-6.771l6.052-6.052c-.001 0 .216-1.083-1.228-2.527zM5.68 17.217l-1.624.35a3.71 3.71 0 0 0-.69-.932a3.742 3.742 0 0 0-.932-.691l.35-1.623l.47-.469s.883.018 1.881 1.016c.997.996 1.016 1.881 1.016 1.881l-.471.468z"/>`), g.Group(children),
	)
}

func Email(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.608 12.172c0 .84.239 1.175.864 1.175c1.393 0 2.28-1.775 2.28-4.727c0-4.512-3.288-6.672-7.393-6.672c-4.223 0-8.064 2.832-8.064 8.184c0 5.112 3.36 7.896 8.52 7.896c1.752 0 2.928-.192 4.727-.792l.386 1.607c-1.776.577-3.674.744-5.137.744c-6.768 0-10.393-3.72-10.393-9.456c0-5.784 4.201-9.72 9.985-9.72c6.024 0 9.215 3.6 9.215 8.016c0 3.744-1.175 6.6-4.871 6.6c-1.681 0-2.784-.672-2.928-2.161c-.432 1.656-1.584 2.161-3.145 2.161c-2.088 0-3.84-1.609-3.84-4.848c0-3.264 1.537-5.28 4.297-5.28c1.464 0 2.376.576 2.782 1.488l.697-1.272h2.016v7.057h.002zm-2.951-3.168c0-1.319-.985-1.872-1.801-1.872c-.888 0-1.871.719-1.871 2.832c0 1.68.744 2.616 1.871 2.616c.792 0 1.801-.504 1.801-1.896v-1.68z"/>`), g.Group(children),
	)
}

func EmojiFlirt(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 9.75C8.329 9.75 9 8.967 9 8s-.671-1.75-1.5-1.75S6 7.034 6 8s.672 1.75 1.5 1.75zM10 .4A9.6 9.6 0 0 0 .4 10a9.6 9.6 0 1 0 19.2-.001C19.6 4.698 15.301.4 10 .4zm0 17.199a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 0 1 0 15.2zm4.341-6.263a.756.756 0 0 0-1.008.32c-.034.065-.869 1.593-3.332 1.593c-2.451 0-3.291-1.513-3.333-1.592a.75.75 0 0 0-1.339.678c.05.099 1.248 2.414 4.672 2.414c3.425 0 4.621-2.316 4.67-2.415a.745.745 0 0 0-.33-.998zM11.25 8.75h2.5a.75.75 0 0 0 0-1.5h-2.5a.75.75 0 0 0 0 1.5z"/>`), g.Group(children),
	)
}

func EmojiHappy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4A9.6 9.6 0 0 0 .4 10a9.6 9.6 0 1 0 19.2-.001C19.6 4.698 15.301.4 10 .4zm0 17.199A7.6 7.6 0 1 1 10 2.4a7.6 7.6 0 1 1 0 15.199zM7.501 9.75C8.329 9.75 9 8.967 9 8s-.672-1.75-1.5-1.75S6 7.033 6 8s.672 1.75 1.501 1.75zm4.999 0c.829 0 1.5-.783 1.5-1.75s-.672-1.75-1.5-1.75S11 7.034 11 8s.672 1.75 1.5 1.75zm1.841 1.586a.756.756 0 0 0-1.008.32c-.034.066-.869 1.593-3.332 1.593c-2.451 0-3.291-1.513-3.333-1.592a.75.75 0 0 0-1.339.678c.05.099 1.248 2.414 4.672 2.414c3.425 0 4.621-2.316 4.67-2.415a.745.745 0 0 0-.33-.998z"/>`), g.Group(children),
	)
}

func EmojiNeutral(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4A9.6 9.6 0 0 0 .4 10a9.6 9.6 0 1 0 19.2-.001C19.6 4.698 15.301.4 10 .4zm0 17.199a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 0 1 0 15.2zm2.5-11.348c-.828 0-1.5.783-1.5 1.749s.672 1.75 1.5 1.75c.829 0 1.5-.783 1.5-1.75s-.671-1.749-1.5-1.749zM7.501 9.75C8.329 9.75 9 8.967 9 8s-.672-1.75-1.5-1.75S6 7.033 6 8s.672 1.75 1.501 1.75zM13 12.25H7a.75.75 0 0 0 0 1.5h6a.75.75 0 0 0 0-1.5z"/>`), g.Group(children),
	)
}

func EmojiSad(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.001.4C4.698.4.4 4.698.4 10a9.6 9.6 0 0 0 9.601 9.601c5.301 0 9.6-4.298 9.6-9.601c-.001-5.302-4.3-9.6-9.6-9.6zM10 17.599a7.6 7.6 0 1 1 0-15.2a7.6 7.6 0 0 1 0 15.2zm2.501-7.849c.828 0 1.5-.783 1.5-1.75s-.672-1.75-1.5-1.75s-1.5.783-1.5 1.75s.671 1.75 1.5 1.75zm-5 0c.828 0 1.5-.783 1.5-1.75s-.672-1.75-1.5-1.75s-1.5.783-1.5 1.75s.671 1.75 1.5 1.75zm2.501 1.5c-3.424 0-4.622 2.315-4.672 2.414a.75.75 0 0 0 1.342.672c.008-.017.822-1.586 3.33-1.586c2.463 0 3.298 1.527 3.328 1.585a.75.75 0 1 0 1.342-.67c-.049-.099-1.246-2.415-4.67-2.415z"/>`), g.Group(children),
	)
}

func Erase(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 3H8.446c-.44 0-1.071.236-1.402.525L.248 9.473a.682.682 0 0 0 0 1.053l6.796 5.947c.331.289.962.527 1.402.527H18c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm-2.809 11l-2.557-2.557L10.078 14l-1.443-1.443L11.191 10L8.635 7.443L10.078 6l2.557 2.555L15.19 6l1.444 1.443L14.078 10l2.557 2.555L15.191 14z"/>`), g.Group(children),
	)
}

func Eraser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.998 4.18l-3.154-2.425a2.01 2.01 0 0 0-2.807.365l-8.4 10.897a2.003 2.003 0 0 0 .365 2.803l3.153 2.425a2.01 2.01 0 0 0 2.807-.365l8.401-10.897a2.003 2.003 0 0 0-.365-2.803zm-8.45 12.287l-.537.681a.8.8 0 0 1-.639.31a.793.793 0 0 1-.485-.164l-3.153-2.425a.792.792 0 0 1-.303-.53a.788.788 0 0 1 .157-.589l.537-.681a.801.801 0 0 1 .64-.311c.124 0 .309.029.485.164l3.154 2.425a.802.802 0 0 1 .144 1.12z"/>`), g.Group(children),
	)
}

func Export(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 15H2V6h2.595s.689-.896 2.17-2H1a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h15a1 1 0 0 0 1-1v-3.746l-2 1.645V15zm-1.639-6.95v3.551L20 6.4l-6.639-4.999v3.131C5.3 4.532 5.3 12.5 5.3 12.5c2.282-3.748 3.686-4.45 8.061-4.45z"/>`), g.Group(children),
	)
}

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 4.4C3.439 4.4 0 9.232 0 10c0 .766 3.439 5.6 10 5.6c6.56 0 10-4.834 10-5.6c0-.768-3.44-5.6-10-5.6zm0 9.907c-2.455 0-4.445-1.928-4.445-4.307c0-2.379 1.99-4.309 4.445-4.309s4.444 1.93 4.444 4.309c0 2.379-1.989 4.307-4.444 4.307zM10 10c-.407-.447.663-2.154 0-2.154c-1.228 0-2.223.965-2.223 2.154s.995 2.154 2.223 2.154c1.227 0 2.223-.965 2.223-2.154c0-.547-1.877.379-2.223 0z"/>`), g.Group(children),
	)
}

func EyeWithLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.521 1.478a1 1 0 0 0-1.414 0L1.48 17.107a1 1 0 1 0 1.414 1.414L18.52 2.892a1 1 0 0 0 0-1.414zM3.108 13.498l2.56-2.56A4.18 4.18 0 0 1 5.555 10c0-2.379 1.99-4.309 4.445-4.309c.286 0 .564.032.835.082l1.203-1.202A12.645 12.645 0 0 0 10 4.401C3.44 4.4 0 9.231 0 10c0 .423 1.057 2.09 3.108 3.497zm13.787-6.993l-2.562 2.56c.069.302.111.613.111.935c0 2.379-1.989 4.307-4.444 4.307c-.284 0-.56-.032-.829-.081l-1.204 1.203c.642.104 1.316.17 2.033.17c6.56 0 10-4.833 10-5.599c0-.424-1.056-2.09-3.105-3.495z"/>`), g.Group(children),
	)
}

func Feather(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.254 19.567c.307-.982.77-2.364 1.391-4.362c2.707-.429 3.827.341 5.546-2.729c-1.395.427-3.077-.792-2.987-1.321c.091-.528 3.913.381 6.416-3.173c-3.155.696-4.164-.836-3.757-1.067c.939-.534 3.726-.222 5.212-1.669c.766-.745 1.125-2.556.813-3.202c-.374-.781-2.656-1.946-3.914-1.836c-1.258.109-3.231 4.79-3.817 4.754c-.584-.037-.703-2.098.319-4.013c-1.077.477-3.051 1.959-3.67 3.226c-1.153 2.357.108 7.766-.296 7.958c-.405.193-1.766-2.481-2.172-3.694c-.555 1.859-.568 3.721 1.053 6.194c-.611 1.623-.945 3.491-.996 4.441c-.024.759.724.922.859.493z"/>`), g.Group(children),
	)
}

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.926 5.584c-9.339 13.568-6.142-.26-14.037 6.357L6.684 19H4.665L1 4.59l1.85-.664c8.849-6.471 4.228 5.82 15.637 1.254c.364-.147.655.09.439.404z"/>`), g.Group(children),
	)
}

func Flash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.803 18.998c-.194-.127 3.153-7.16 3.038-7.469c-.116-.309-3.665-1.436-3.838-1.979c-.174-.543 7.007-8.707 7.196-8.549c.188.158-3.129 7.238-3.039 7.469c.091.23 3.728 1.404 3.838 1.979c.111.575-7.002 8.676-7.195 8.549z"/>`), g.Group(children),
	)
}

func Flashlight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.115 2.884c-1.676-1.676-3.779-2.288-4.449-1.618L9.97 3.962c-.409.41-.766 1.779-.602 3.164l-8.161 8.161c-.484.484-.092 1.66.876 2.629c.968.969 2.146 1.359 2.629.877l8.161-8.162c1.386.164 2.755-.193 3.164-.601l2.696-2.697c.67-.67.058-2.774-1.618-4.449zm-8.974 8.155c-.373-.372-.251-1.096.269-1.617c.521-.521 1.246-.643 1.618-.27c.372.371.251 1.097-.27 1.617c-.521.522-1.245.643-1.617.27zm6.75-5.931c-1.298-1.297-1.623-3.01-1.508-3.125c.115-.116 1.76.277 3.059 1.575c1.298 1.298 1.688 2.946 1.575 3.059c-.112.112-1.829-.21-3.126-1.509z"/>`), g.Group(children),
	)
}

func FlatBrush(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.844 14.889c.498.376 1.594-1.178 1.942-.915c.348.263-.82 1.762-.528 1.982c.292.22 1.513-1.239 1.852-.984c.338.255-.803 1.774-.437 2.051c.367.277 1.562-1.202 1.852-.983c.29.219-.773 1.797-.437 2.05c.336.254 1.481-1.263 1.76-1.052c.28.211-.844 1.743-.346 2.119c.498.375 5.37-8.964 5.37-8.964L8.996 7.254s-7.65 7.26-7.152 7.635zM13.023.831L9.661 5.987l4.121 3.109l4.396-4.246c-.527-1.503-3.44-3.843-5.155-4.019z"/>`), g.Group(children),
	)
}

func FlowBranch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.4 4a2.4 2.4 0 1 0-4.8 0c0 .961.568 1.784 1.384 2.167c-.082 1.584-1.27 2.122-3.335 2.896c-.87.327-1.829.689-2.649 1.234V6.176A2.396 2.396 0 0 0 6 1.6a2.397 2.397 0 0 0-1 4.576v7.649A2.393 2.393 0 0 0 3.6 16a2.4 2.4 0 1 0 4.8 0c0-.961-.568-1.784-1.384-2.167c.082-1.583 1.271-2.122 3.335-2.896c2.03-.762 4.541-1.711 4.64-4.756A2.398 2.398 0 0 0 16.4 4zM6 2.615a1.384 1.384 0 1 1 0 2.768a1.384 1.384 0 0 1 0-2.768zm0 14.77a1.385 1.385 0 1 1 0-2.77a1.385 1.385 0 0 1 0 2.77zm8-12a1.385 1.385 0 1 1 0-2.77a1.385 1.385 0 0 1 0 2.77z"/>`), g.Group(children),
	)
}

func FlowCascade(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 14.6c-.967 0-1.796.576-2.176 1.4H8.5A1.5 1.5 0 0 1 7 14.5v-3.85c.456.218.961.35 1.5.35h3.324a2.396 2.396 0 0 0 4.576-1a2.397 2.397 0 0 0-4.576-1H8.5A1.5 1.5 0 0 1 7 7.5V5.176A2.396 2.396 0 0 0 6 .6a2.397 2.397 0 0 0-1 4.576V14.5A3.5 3.5 0 0 0 8.5 18h3.324a2.396 2.396 0 0 0 4.576-1a2.4 2.4 0 0 0-2.4-2.4zm0-5.985a1.384 1.384 0 1 1 0 2.768a1.384 1.384 0 0 1 0-2.768zm-8-7a1.384 1.384 0 1 1 0 2.768a1.384 1.384 0 0 1 0-2.768zm8 16.77a1.385 1.385 0 1 1 0-2.77a1.385 1.385 0 0 1 0 2.77z"/>`), g.Group(children),
	)
}

func FlowLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 13.824V6.176A2.395 2.395 0 0 0 12.4 4a2.4 2.4 0 1 0-4.8 0c0 .967.576 1.796 1.4 2.176v7.649A2.393 2.393 0 0 0 7.6 16a2.4 2.4 0 1 0 4.8 0c0-.967-.575-1.796-1.4-2.176zM10 2.615a1.384 1.384 0 1 1 0 2.768a1.384 1.384 0 0 1 0-2.768zm0 14.77a1.385 1.385 0 1 1 0-2.77a1.385 1.385 0 0 1 0 2.77z"/>`), g.Group(children),
	)
}

func FlowParallel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.4 4a2.4 2.4 0 1 0-4.8 0c0 .967.576 1.796 1.4 2.176v7.649A2.393 2.393 0 0 0 3.6 16a2.4 2.4 0 1 0 4.8 0c0-.967-.576-1.796-1.4-2.176V6.176A2.396 2.396 0 0 0 8.4 4zM7.384 16a1.385 1.385 0 1 1-2.77 0a1.385 1.385 0 0 1 2.77 0zM6 5.385a1.385 1.385 0 1 1 .001-2.77A1.385 1.385 0 0 1 6 5.386zm9 8.439V6.176A2.395 2.395 0 0 0 16.4 4a2.4 2.4 0 1 0-4.8 0c0 .967.576 1.796 1.4 2.176v7.649a2.395 2.395 0 0 0-1.4 2.176a2.4 2.4 0 1 0 4.8 0c0-.968-.575-1.797-1.4-2.177zM12.616 4a1.384 1.384 0 1 1 2.769 0a1.385 1.385 0 0 1-2.769 0zM14 17.385a1.385 1.385 0 1 1 0-2.77a1.385 1.385 0 0 1 0 2.77z"/>`), g.Group(children),
	)
}

func FlowTree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 14.824V12.5A3.5 3.5 0 0 0 14.5 9h-2A1.5 1.5 0 0 1 11 7.5V5.176A2.395 2.395 0 0 0 12.4 3a2.4 2.4 0 1 0-4.8 0c0 .967.576 1.796 1.4 2.176V7.5A1.5 1.5 0 0 1 7.5 9h-2A3.5 3.5 0 0 0 2 12.5v2.324A2.396 2.396 0 0 0 3 19.4a2.397 2.397 0 0 0 1-4.576V12.5A1.5 1.5 0 0 1 5.5 11h2c.539 0 1.044-.132 1.5-.35v4.174a2.396 2.396 0 0 0 1 4.576a2.397 2.397 0 0 0 1-4.576V10.65c.456.218.961.35 1.5.35h2a1.5 1.5 0 0 1 1.5 1.5v2.324A2.395 2.395 0 0 0 14.6 17a2.4 2.4 0 1 0 4.8 0c0-.967-.575-1.796-1.4-2.176zM10 1.615a1.384 1.384 0 1 1 0 2.768a1.384 1.384 0 0 1 0-2.768zm-7 16.77a1.385 1.385 0 1 1 0-2.77a1.385 1.385 0 0 1 0 2.77zm7 0a1.385 1.385 0 1 1 0-2.77a1.385 1.385 0 0 1 0 2.77zm7 0a1.385 1.385 0 1 1 0-2.77a1.385 1.385 0 0 1 0 2.77z"/>`), g.Group(children),
	)
}

func Flower(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.66 8.76c.439-1.087-.364-2.509-1.21-3.406c.328-1.19.371-2.822-.526-3.575c-.897-.754-2.497-.43-3.613.098c-1.03-.68-2.568-1.226-3.563-.606c-.994.62-1.18 2.243-1.023 3.467c-.964.769-1.959 2.064-1.676 3.201s1.768 1.816 2.98 2.044c.434 1.155 1.358 2.501 2.527 2.584c.828.058 1.62-.531 2.224-1.263c1.22 1.52 2.082 3.127 2.555 4.706C10.612 14.93 9.115 14 6 14c0 4.155 3.042 5.003 5 5.003l3.2.002c0-2.723-.986-5.91-3.29-8.901c.783-.247 1.479-.677 1.75-1.344zm-4.377-.22a2.447 2.447 0 1 1-2.591-4.153a2.447 2.447 0 0 1 2.59 4.153zm5.737 3.708c1.083 2.17 1.669 4.453 1.678 6.705C18.996 16.582 19 12.206 19 8.003c0 0-3.67 1.034-4.98 4.245z"/>`), g.Group(children),
	)
}

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.405 4.799c-.111-.44-.655-.799-1.21-.799h-6.814c-.554 0-1.33-.318-1.722-.707l-.596-.588C7.671 2.316 6.896 2 6.342 2H3.087c-.555 0-1.059.447-1.12.994L1.675 6h16.931l-.201-1.201zM19.412 7H.588a.58.58 0 0 0-.577.635l.923 9.669A.77.77 0 0 0 1.7 18h16.6a.77.77 0 0 0 .766-.696l.923-9.669A.58.58 0 0 0 19.412 7z"/>`), g.Group(children),
	)
}

func FolderImages(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.405 2.799c-.112-.44-.656-.799-1.21-.799H2.805c-.555 0-1.099.359-1.21.799L1.394 4h17.211l-.2-1.201zM19.412 5H.587a.58.58 0 0 0-.577.635l.923 11.669a.77.77 0 0 0 .766.696H18.3a.77.77 0 0 0 .766-.696l.923-11.669A.58.58 0 0 0 19.412 5zm-6.974 3.375a.938.938 0 1 1 0 1.876a.938.938 0 0 1 0-1.876zM5.5 14l2.486-5.714l2.827 4.576l2.424-1.204L14.5 14h-9z"/>`), g.Group(children),
	)
}

func FolderMusic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.405 2.799c-.112-.44-.656-.799-1.21-.799H2.805c-.555 0-1.099.359-1.21.799L1.394 4h17.211l-.2-1.201zM19.412 5H.587a.58.58 0 0 0-.577.635l.923 11.669a.77.77 0 0 0 .766.696H18.3a.77.77 0 0 0 .766-.696l.923-11.669A.58.58 0 0 0 19.412 5zm-7.47 7.521c-.128.265-.258.279-.202 0c.146-.721.047-2.269-1.043-2.441v3.294c0 .674-.311 1.262-1.136 1.528c-.802.256-1.699-.011-1.908-.586c-.21-.576.261-1.276 1.052-1.564c.442-.161.954-.203 1.299-.07V8h.694c-.001 1.633 2.818 1.275 1.244 4.521z"/>`), g.Group(children),
	)
}

func FolderVideo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.405 2.799c-.112-.44-.656-.799-1.21-.799H2.805c-.555 0-1.099.359-1.21.799L1.394 4h17.211l-.2-1.201zM19.412 5H.587a.58.58 0 0 0-.577.635l.923 11.669a.77.77 0 0 0 .766.696H18.3a.77.77 0 0 0 .766-.696l.923-11.669A.58.58 0 0 0 19.412 5zM8 14V9l4.383 2.5L8 14z"/>`), g.Group(children),
	)
}

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 11.874v4.357l7-6.69l-7-6.572v3.983c-8.775 0-11 9.732-11 9.732c2.484-4.388 6.237-4.81 11-4.81z"/>`), g.Group(children),
	)
}

func Funnel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 1C5.092 1 2 2.512 2 4.001v2c0 .918 6 6 6 6v6c-.001.684 1 1 2 1s2.001-.316 2-1v-6s6-5.082 6-6v-2C18 2.512 14.908 1 10 1zm0 5.123C6.409 6.122 3.862 4.79 3.862 4.292C3.86 3.797 6.41 2.461 10 2.463c3.59-.002 6.14 1.334 6.138 1.828c0 .499-2.547 1.831-6.138 1.832z"/>`), g.Group(children),
	)
}

func GameController(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19.444 9.361c-.882-4.926-2.854-6.379-3.903-6.379c-1.637 0-2.057 1.217-5.541 1.258c-3.484-.041-3.904-1.258-5.541-1.258c-1.049 0-3.022 1.453-3.904 6.379c-.503 2.812-1.049 7.01.252 7.514c1.619.627 2.168-.941 3.946-2.266C6.558 13.266 7.424 12.95 10 12.95s3.442.316 5.247 1.659c1.778 1.324 2.327 2.893 3.946 2.266c1.301-.504.755-4.701.251-7.514zM6 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4zm7 0a1 1 0 1 1 0-2a1 1 0 1 1 0 2zm2-2a1 1 0 1 1 0-2a1 1 0 1 1 0 2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func Gauge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.127 13.6c-.689 1.197-.225 2.18.732 2.732c.956.553 2.041.465 2.732-.732c.689-1.195 5.047-11.865 4.668-12.084c-.379-.219-7.442 8.888-8.132 10.084zM10 6c.438 0 .864.037 1.281.109c.438-.549.928-1.154 1.405-1.728A9.664 9.664 0 0 0 10 4C4.393 4 0 8.729 0 14.766c0 .371.016.742.049 1.103c.049.551.54.955 1.084.908c.551-.051.957-.535.908-1.086A10.462 10.462 0 0 1 2 14.766C2 9.85 5.514 6 10 6zm7.219 1.25c-.279.75-.574 1.514-.834 2.174C17.4 10.894 18 12.738 18 14.766c0 .316-.015.635-.043.943a1.001 1.001 0 0 0 1.992.182c.033-.37.051-.748.051-1.125c0-2.954-1.053-5.59-2.781-7.516z"/>`), g.Group(children),
	)
}

func Globe(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4C4.705.4.399 4.707.399 10c0 5.293 4.306 9.6 9.601 9.6c5.293 0 9.6-4.307 9.6-9.6c0-5.293-4.307-9.6-9.6-9.6zm8.188 9.6c0 1.873-.636 3.6-1.696 4.98c-.3-.234-.619-.867-.319-1.523c.303-.66.382-2.188.312-2.783c-.066-.594-.375-2.025-1.214-2.039c-.838-.012-1.413-.289-1.911-1.283c-1.033-2.068 1.939-2.465.906-3.609c-.289-.322-1.783 1.322-2.002-.869c-.014-.157.135-.392.336-.636c3.244 1.09 5.588 4.157 5.588 7.762zM8.875 1.893c-.196.382-.713.537-1.027.824c-.684.619-.978.533-1.346 1.127c-.371.594-1.567 1.449-1.567 1.879s.604.936.906.838c.302-.1 1.099-.094 1.567.07c.469.166 3.914.332 2.816 3.244c-.348.926-1.873.77-2.279 2.303c-.061.225-.272 1.186-.285 1.5c-.025.486.344 2.318-.125 2.318c-.471 0-1.738-1.639-1.738-1.936c0-.297-.328-1.338-.328-2.23c0-.891-1.518-.877-1.518-2.062c0-1.068.823-1.6.638-2.113c-.181-.51-1.627-.527-2.23-.59a8.213 8.213 0 0 1 6.516-5.172zM7.424 17.77c.492-.26.542-.596.988-.613c.51-.023.925-.199 1.5-.326c.51-.111 1.423-.629 2.226-.695c.678-.055 2.015.035 2.375.689a8.159 8.159 0 0 1-7.089.945z"/>`), g.Group(children),
	)
}

func GraduationCap(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.302 12.238c.464 1.879 1.054 2.701 3.022 3.562c1.969.86 2.904 1.8 3.676 1.8c.771 0 1.648-.822 3.616-1.684c1.969-.861 1.443-1.123 1.907-3.002L10 15.6l-6.698-3.362zm16.209-4.902l-8.325-4.662c-.652-.365-1.72-.365-2.372 0L.488 7.336c-.652.365-.652.963 0 1.328l8.325 4.662c.652.365 1.72.365 2.372 0l5.382-3.014l-5.836-1.367a3.09 3.09 0 0 1-.731.086c-1.052 0-1.904-.506-1.904-1.131c0-.627.853-1.133 1.904-1.133c.816 0 1.51.307 1.78.734l6.182 2.029l1.549-.867c.651-.364.651-.962 0-1.327zm-2.544 8.834c-.065.385 1.283 1.018 1.411-.107c.579-5.072-.416-6.531-.416-6.531l-1.395.781c0-.001 1.183 1.125.4 5.857z"/>`), g.Group(children),
	)
}

func Grid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M8 4H5a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1zm7 0h-3a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1V5a1 1 0 0 0-1-1zm-7 7H5a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1zm7 0h-3a1 1 0 0 0-1 1v3a1 1 0 0 0 1 1h3a1 1 0 0 0 1-1v-3a1 1 0 0 0-1-1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func HairCross(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4C4.697.4.399 4.698.399 10A9.6 9.6 0 0 0 10 19.601c5.301 0 9.6-4.298 9.6-9.601c0-5.302-4.299-9.6-9.6-9.6zm1 17.125V13H9v4.525A7.589 7.589 0 0 1 2.473 11H7V9H2.473A7.589 7.589 0 0 1 9 2.475V7h2V2.475A7.59 7.59 0 0 1 17.525 9H13v2h4.525A7.592 7.592 0 0 1 11 17.525z"/>`), g.Group(children),
	)
}

func Hand(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.924 17.315c-.057.174-.193.367-.416.432c-.161.047-5.488 1.59-5.652 1.633c-.469.125-.795.033-1.009-.156c-.326-.287-4.093-2.85-8.845-3.092c-.508-.025-.259-1.951 1.193-1.951c.995 0 3.904.723 4.255.371c.271-.272.394-1.879-.737-4.683L4.438 4.232a1.045 1.045 0 0 1 1.937-.781L8.361 8.37c.193.48.431.662.69.562c.231-.088.279-.242.139-.709L7.144 2.195A1.043 1.043 0 0 1 7.796.871a1.042 1.042 0 0 1 1.325.652l1.946 5.732c.172.504.354.768.642.646c.173-.073.161-.338.115-.569l-1.366-5.471a1.045 1.045 0 1 1 2.027-.506l1.26 5.042c.184.741.353 1.008.646.935c.299-.073.285-.319.244-.522l-.872-4.328a.95.95 0 0 1 1.86-.375l.948 4.711l.001.001v.001l.568 2.825c.124.533.266 1.035.45 1.527c1.085 2.889.519 5.564.334 6.143z"/>`), g.Group(children),
	)
}

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.19 4.155c-1.672-1.534-4.383-1.534-6.055 0L10 5.197L8.864 4.155c-1.672-1.534-4.382-1.534-6.054 0c-1.881 1.727-1.881 4.52 0 6.246L10 17l7.19-6.599c1.88-1.726 1.88-4.52 0-6.246z"/>`), g.Group(children),
	)
}

func HeartOutlined(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.19 4.156c-1.672-1.535-4.383-1.535-6.055 0L10 5.197L8.864 4.156c-1.672-1.535-4.382-1.535-6.054 0c-1.881 1.726-1.881 4.519 0 6.245L10 17l7.19-6.599c1.88-1.726 1.88-4.52 0-6.245zm-1.066 5.219L10 15.09L3.875 9.375c-.617-.567-.856-1.307-.856-2.094s.138-1.433.756-1.999c.545-.501 1.278-.777 2.063-.777c.784 0 1.517.476 2.062.978L10 7.308l2.099-1.826c.546-.502 1.278-.978 2.063-.978s1.518.276 2.063.777c.618.566.755 1.212.755 1.999s-.238 1.528-.856 2.095z"/>`), g.Group(children),
	)
}

func Help(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.09 2.233C12.95 1.411 11.518 1 9.794 1c-1.311 0-2.418.289-3.317.868C5.05 2.774 4.292 4.313 4.2 6.483h3.307c0-.633.185-1.24.553-1.828c.369-.586.995-.879 1.878-.879c.898 0 1.517.238 1.854.713c.339.477.508 1.004.508 1.582c0 .504-.252.965-.557 1.383a2.88 2.88 0 0 1-.661.674s-1.793 1.15-2.58 2.074c-.456.535-.497 1.338-.538 2.488c-.002.082.029.252.315.252h2.571c.256 0 .309-.189.312-.274c.018-.418.064-.633.141-.875c.144-.457.538-.855.979-1.199l.91-.627c.822-.641 1.477-1.166 1.767-1.578c.494-.676.842-1.51.842-2.5c-.001-1.615-.571-2.832-1.711-3.656zM9.741 14.924c-1.139-.035-2.079.754-2.115 1.99c-.035 1.234.858 2.051 1.998 2.084c1.189.035 2.104-.727 2.141-1.963c.034-1.236-.834-2.076-2.024-2.111z"/>`), g.Group(children),
	)
}

func HelpWithCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4A9.6 9.6 0 0 0 .4 10a9.6 9.6 0 1 0 19.2-.001C19.6 4.698 15.301.4 10 .4zm-.151 15.199h-.051c-.782-.023-1.334-.6-1.311-1.371c.022-.758.587-1.309 1.343-1.309l.046.002c.804.023 1.35.594 1.327 1.387c-.023.76-.578 1.291-1.354 1.291zm3.291-6.531c-.184.26-.588.586-1.098.983l-.562.387c-.308.24-.494.467-.563.688c-.056.174-.082.221-.087.576v.09H8.685l.006-.182c.027-.744.045-1.184.354-1.547c.485-.568 1.555-1.258 1.6-1.287a1.65 1.65 0 0 0 .379-.387c.225-.311.324-.555.324-.793c0-.334-.098-.643-.293-.916c-.188-.266-.545-.398-1.061-.398c-.512 0-.863.162-1.072.496c-.216.341-.325.7-.325 1.067v.092H6.386l.004-.096c.057-1.353.541-2.328 1.435-2.897c.563-.361 1.264-.544 2.081-.544c1.068 0 1.972.26 2.682.772c.721.519 1.086 1.297 1.086 2.311c-.001.567-.18 1.1-.534 1.585z"/>`), g.Group(children),
	)
}

func Home(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.672 11H17v6c0 .445-.194 1-1 1h-4v-6H8v6H4c-.806 0-1-.555-1-1v-6H1.328c-.598 0-.47-.324-.06-.748L9.292 2.22c.195-.202.451-.302.708-.312c.257.01.513.109.708.312l8.023 8.031c.411.425.539.749-.059.749z"/>`), g.Group(children),
	)
}

func HourGlass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.6 4.576V2.228C15.6 1.439 13.092 0 10 0C6.907 0 4.4 1.439 4.4 2.228v2.348C4.4 6.717 8.277 8.484 8.277 10c0 1.514-3.877 3.281-3.877 5.422v2.35C4.4 18.56 6.907 20 10 20c3.092 0 5.6-1.44 5.6-2.229v-2.35c0-2.141-3.877-3.908-3.877-5.422c0-1.515 3.877-3.282 3.877-5.423zM5.941 2.328c.696-.439 2-1.082 4.114-1.082c2.113 0 4.006 1.082 4.006 1.082c.142.086.698.383.317.609c-.838.497-2.478 1.02-4.378 1.02s-3.484-.576-4.324-1.074c-.381-.225.265-.555.265-.555zM10.501 10c0 1.193.996 1.961 2.051 2.986c.771.748 1.826 1.773 1.826 2.435v1.328c-.97-.483-3.872-.955-3.872-2.504c0-.783-1.013-.783-1.013 0c0 1.549-2.902 2.021-3.872 2.504v-1.328c0-.662 1.056-1.688 1.826-2.435C8.502 11.961 9.498 11.193 9.498 10c0-1.193-.996-1.961-2.051-2.986c-.771-.75-1.826-1.775-1.826-2.438l-.046-.998C6.601 4.131 8.227 4.656 10 4.656c1.772 0 3.406-.525 4.433-1.078l-.055.998c0 .662-1.056 1.688-1.826 2.438c-1.054 1.025-2.051 1.793-2.051 2.986z"/>`), g.Group(children),
	)
}

func Image(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M19 2H1a1 1 0 0 0-1 1v14a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V3a1 1 0 0 0-1-1zm-1 14H2V4h16v12zm-3.685-5.123l-3.231 1.605l-3.77-6.101L4 14h12l-1.685-3.123zM13.25 9a1.25 1.25 0 1 0 0-2.5a1.25 1.25 0 0 0 0 2.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func ImageInverted(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 3H2a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zm-4.75 3.5a1.25 1.25 0 1 1 0 2.5a1.25 1.25 0 0 1 0-2.5zM4 14l3.314-7.619l3.769 6.102l3.231-1.605L16 14H4z"/>`), g.Group(children),
	)
}

func Images(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.125 6.17L15.079.535c-.151-.416-.595-.637-.989-.492L.492 5.006c-.394.144-.593.597-.441 1.013l2.156 5.941V8.777c0-1.438 1.148-2.607 2.56-2.607H8.36l4.285-3.008l2.479 3.008h2.001zM19.238 8H4.767a.761.761 0 0 0-.762.777v9.42c.001.444.343.803.762.803h14.471c.42 0 .762-.359.762-.803v-9.42A.761.761 0 0 0 19.238 8zM18 17H6v-2l1.984-4.018l2.768 3.436l2.598-2.662l3.338-1.205L18 14v3z"/>`), g.Group(children),
	)
}

func Inbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.253 9.542c-.388-.416-2.265-2.271-3.122-3.118A1.49 1.49 0 0 0 15.098 6H4.902c-.394 0-.77.165-1.033.424c-.858.847-2.734 2.701-3.122 3.118c-.485.521-.723.902-.624 1.449s.466 2.654.556 3.074c.088.419.684.935 1.24.935h16.162c.556 0 1.152-.516 1.241-.935c.089-.42.457-2.527.556-3.074s-.139-.929-.625-1.449zm-5.239.461a.271.271 0 0 0-.238.133L12.966 12H7.034l-.809-1.864a.271.271 0 0 0-.238-.133H2.473L4.495 8h11.01l2.023 2.003h-3.514z"/>`), g.Group(children),
	)
}

func Infinity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.902 5.6c-2.079 0-4.358 1.158-5.902 2.916C8.455 6.758 6.175 5.6 4.096 5.6C2.116 5.6 0 6.756 0 10c0 3.244 2.116 4.398 4.096 4.4c2.079 0 4.358-1.158 5.903-2.916c1.544 1.758 3.823 2.916 5.902 2.916C17.882 14.4 20 13.244 20 10c0-3.244-2.118-4.4-4.098-4.4zM4.096 12.641C2.584 12.641 1.8 11.752 1.8 10s.784-2.641 2.296-2.641c1.673 0 3.614 1.086 4.807 2.641c-1.193 1.555-3.134 2.641-4.807 2.641zm11.806 0c-1.673 0-3.614-1.086-4.807-2.641c1.192-1.555 3.135-2.641 4.807-2.641c1.512 0 2.298.889 2.298 2.641s-.786 2.641-2.298 2.641z"/>`), g.Group(children),
	)
}

func Info(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.432 0c1.34 0 2.01.912 2.01 1.957c0 1.305-1.164 2.512-2.679 2.512c-1.269 0-2.009-.75-1.974-1.99C9.789 1.436 10.67 0 12.432 0zM8.309 20c-1.058 0-1.833-.652-1.093-3.524l1.214-5.092c.211-.814.246-1.141 0-1.141c-.317 0-1.689.562-2.502 1.117l-.528-.88c2.572-2.186 5.531-3.467 6.801-3.467c1.057 0 1.233 1.273.705 3.23l-1.391 5.352c-.246.945-.141 1.271.106 1.271c.317 0 1.357-.392 2.379-1.207l.6.814C12.098 19.02 9.365 20 8.309 20z"/>`), g.Group(children),
	)
}

func InfoWithCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4C4.697.4.399 4.698.399 10A9.6 9.6 0 0 0 10 19.601c5.301 0 9.6-4.298 9.6-9.601c0-5.302-4.299-9.6-9.6-9.6zm.896 3.466c.936 0 1.211.543 1.211 1.164c0 .775-.62 1.492-1.679 1.492c-.886 0-1.308-.445-1.282-1.182c0-.621.519-1.474 1.75-1.474zM8.498 15.75c-.64 0-1.107-.389-.66-2.094l.733-3.025c.127-.484.148-.678 0-.678c-.191 0-1.022.334-1.512.664l-.319-.523c1.555-1.299 3.343-2.061 4.108-2.061c.64 0 .746.756.427 1.92l-.84 3.18c-.149.562-.085.756.064.756c.192 0 .82-.232 1.438-.719l.362.486c-1.513 1.512-3.162 2.094-3.801 2.094z"/>`), g.Group(children),
	)
}

func Install(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19.059 10.898l-3.171-7.927A1.543 1.543 0 0 0 14.454 2H12.02l.38 4.065h2.7L10 10.293L4.9 6.065h2.7L7.98 2H5.546c-.632 0-1.2.384-1.434.971L.941 10.898a4.25 4.25 0 0 0-.246 2.272l.59 3.539A1.544 1.544 0 0 0 2.808 18h14.383c.755 0 1.399-.546 1.523-1.291l.59-3.539a4.22 4.22 0 0 0-.245-2.272zm-2.1 4.347a.902.902 0 0 1-.891.755H3.932a.902.902 0 0 1-.891-.755l-.365-2.193A.902.902 0 0 1 3.567 12h12.867c.558 0 .983.501.891 1.052l-.366 2.193z"/>`), g.Group(children),
	)
}

func Key(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.691 4.725c-.503-2.977-3.22-4.967-6.069-4.441C8.772.809 6.366 3.1 6.869 6.079c.107.641.408 1.644.763 2.365l-5.175 7.723c-.191.285-.299.799-.242 1.141l.333 1.971a.612.612 0 0 0 .7.514l1.516-.281c.328-.059.744-.348.924-.639l2.047-3.311l.018-.022l1.386-.256l2.39-3.879c.785.139 1.912.092 2.578-.031c2.848-.526 4.087-3.67 3.584-6.649zm-2.525 1.527c-.784 1.17-1.584.346-2.703-.475c-1.119-.818-2.135-1.322-1.352-2.492c.784-1.17 2.326-1.455 3.447-.635c1.12.819 1.391 2.432.608 3.602z"/>`), g.Group(children),
	)
}

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.6 4H1.4C.629 4 0 4.629 0 5.4v9.2c0 .769.629 1.4 1.399 1.4h17.2c.77 0 1.4-.631 1.4-1.4V5.4C20 4.629 19.369 4 18.6 4zM11 6h2v2h-2V6zm3 3v2h-2V9h2zM8 6h2v2H8V6zm3 3v2H9V9h2zM5 6h2v2H5V6zm3 3v2H6V9h2zM2 6h2v2H2V6zm3 3v2H3V9h2zm-1 5H2v-2h2v2zm11 0H5v-2h10v2zm3 0h-2v-2h2v2zm-3-3V9h2v2h-2zm3-3h-4V6h4v2z"/>`), g.Group(children),
	)
}

func LabFlask(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.432 15C14.387 9.893 12 8.547 12 6V3h.5a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5h-5a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5H8v3c0 2.547-2.387 3.893-4.432 9c-.651 1.625-2.323 4 6.432 4s7.083-2.375 6.432-4zm-1.617 1.751c-.702.21-2.099.449-4.815.449s-4.113-.239-4.815-.449c-.249-.074-.346-.363-.258-.628c.22-.67.635-1.828 1.411-3.121c1.896-3.159 3.863.497 5.5.497s1.188-1.561 1.824-.497a15.353 15.353 0 0 1 1.411 3.121c.088.265-.009.553-.258.628z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func Landline(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.902.25C3.498-.027 2.115.875 1.833 2.273c-1.105 5.455-1.105 9.997 0 15.454C2.08 18.952 3.17 19.8 4.388 19.8c.17 0 .342-.016.515-.05c1.412-.279 2.329-1.638 2.046-3.036c-.978-4.832-.978-8.598 0-13.43C7.231 1.888 6.314.529 4.902.25zM17 2H8.436a4.08 4.08 0 0 1-.017 1.44c-.936 4.72-.936 8.398 0 13.12c.098.49.09.973.017 1.44H17c1.1 0 2-.9 2-2V4c0-1.1-.899-2-2-2zm-5 12.5a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 12 14.5zM17 9h-7V5h7v4z"/>`), g.Group(children),
	)
}

func Language(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.753 10.909c-.624-1.707-2.366-2.726-4.661-2.726c-.09 0-.176.002-.262.006l-.016-2.063l3.525-.607c.115-.019.133-.119.109-.231c-.023-.111-.167-.883-.188-.976c-.027-.131-.102-.127-.207-.109c-.104.018-3.25.461-3.25.461l-.013-2.078c-.001-.125-.069-.158-.194-.156l-1.025.016c-.105.002-.164.049-.162.148l.033 2.307s-3.061.527-3.144.543c-.084.014-.17.053-.151.143c.019.09.19 1.094.208 1.172c.018.08.072.129.188.107l2.924-.504l.035 2.018c-1.077.281-1.801.824-2.256 1.303c-.768.807-1.207 1.887-1.207 2.963c0 1.586.971 2.529 2.328 2.695c3.162.387 5.119-3.06 5.769-4.715c1.097 1.506.256 4.354-2.094 5.98c-.043.029-.098.129-.033.207l.619.756c.08.096.206.059.256.023c2.51-1.73 3.661-4.515 2.869-6.683zm-7.386 3.188c-.966-.121-.944-.914-.944-1.453c0-.773.327-1.58.876-2.156a3.21 3.21 0 0 1 1.229-.799l.082 4.277a2.773 2.773 0 0 1-1.243.131zm2.427-.553l.046-4.109c.084-.004.166-.01.252-.01c.773 0 1.494.145 1.885.361c.391.217-1.023 2.713-2.183 3.758zm-8.95-7.668a.196.196 0 0 0-.196-.145h-1.95a.194.194 0 0 0-.194.144L.008 16.916c-.017.051-.011.076.062.076h1.733c.075 0 .099-.023.114-.072l1.008-3.318h3.496l1.008 3.318c.016.049.039.072.113.072h1.734c.072 0 .078-.025.062-.076c-.014-.05-3.083-9.741-3.494-11.04zm-2.618 6.318l1.447-5.25l1.447 5.25H3.226z"/>`), g.Group(children),
	)
}

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.754 15.631L18 13V4c0-1.102-.9-2-2-2H4c-1.101 0-2 .898-2 2v9L.246 15.631C0 16 0 16.213 0 16.5v.5c0 .5.5 1 .999 1h18.002c.499 0 .999-.5.999-1v-.5c0-.287 0-.5-.246-.869zM7 16l.6-1h4.8l.6 1H7zm9-4H4V4h12v8z"/>`), g.Group(children),
	)
}

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.248 11.601c.45.313 1.05.313 1.5 0l9.088-5.281a.375.375 0 0 0-.048-.646l-9.205-3.537a1.315 1.315 0 0 0-1.17 0L.208 5.674a.375.375 0 0 0-.048.646l9.088 5.281zm10.54-.79l-2.486-1.233l-5.725 3.327c-.469.309-1.014.471-1.579.471s-1.11-.163-1.579-.471L2.698 9.576L.208 10.81a.375.375 0 0 0-.048.646l9.088 6.309c.45.313 1.05.313 1.5 0l9.088-6.309a.374.374 0 0 0-.048-.645z"/>`), g.Group(children),
	)
}

func Leaf(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.025 3.587c-4.356 2.556-4.044 7.806-7.096 10.175c-2.297 1.783-5.538.88-7.412.113c0 0-1.27 1.603-2.181 3.74c-.305.717-1.644-.073-1.409-.68C3.905 9.25 14.037 5.416 14.037 5.416s-7.149-.303-11.927 5.94c-.128-1.426-.34-5.284 3.36-7.65c5.016-3.211 14.572-.715 13.555-.119z"/>`), g.Group(children),
	)
}

func LevelDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 12V5h3v6h10V8l5 4.5l-5 4.5v-3H3a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func LevelUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 9v7h-3v-6H6v3L1 8.5L6 4v3h11c1.104 0 2 .897 2 2z"/>`), g.Group(children),
	)
}

func Lifebuoy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.788 3.212c-3.749-3.749-9.827-3.749-13.575 0c-3.75 3.75-3.75 9.828-.002 13.576A9.6 9.6 0 1 0 16.788 3.212zm-10.04 10.04a4.598 4.598 0 0 1 0-6.505a4.6 4.6 0 1 1 0 6.505zm8.599-.373a6.07 6.07 0 0 0 0-5.759l1.783-.96a8.111 8.111 0 0 1 .002 7.678l-1.785-.959zm-1.508-10.01l-.961 1.784a6.073 6.073 0 0 0-5.756 0L6.161 2.87a8.114 8.114 0 0 1 7.678-.001zM2.87 6.16l1.784.961a6.07 6.07 0 0 0-.001 5.756l-1.784.961A8.111 8.111 0 0 1 2.87 6.16zm3.289 10.969l.961-1.783a6.068 6.068 0 0 0 5.759 0l.961 1.785a8.117 8.117 0 0 1-7.681-.002z"/>`), g.Group(children),
	)
}

func LightBulb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7.186 19.172c.789.51 1.701.855 2.814.828c1.111.027 2.025-.318 2.814-.828L12.797 17H7.203l-.017 2.172zM12.697 16c0-4.357 4.63-5.848 4.283-10.188c-.218-2.738-2.073-5.81-6.98-5.81S3.238 3.074 3.019 5.813C2.672 10.152 7.303 11.643 7.303 16h5.394zM5 6c.207-2.598 2.113-4 5-4c2.886 0 4.654 1.371 4.861 3.969c.113 1.424-.705 2.373-1.809 3.926C12.238 11.041 11.449 12.238 11 14H9c-.449-1.762-1.238-2.959-2.053-4.106C5.844 8.342 4.886 7.424 5 6z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LightDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 6.797a3.191 3.191 0 0 0-3.2 3.201a3.19 3.19 0 0 0 3.2 3.199a3.19 3.19 0 0 0 3.199-3.199A3.19 3.19 0 0 0 10 6.797zm0 5.25a2.049 2.049 0 1 1 0-4.1a2.05 2.05 0 0 1 0 4.1zM15 5c-.312-.312-.883-.248-1.273.142c-.39.391-.453.959-.141 1.272s.882.25 1.273-.141c.39-.39.453-.961.141-1.273zm-9.858 8.729c-.391.39-.454.959-.142 1.271s.882.25 1.273-.141c.391-.391.454-.961.142-1.273s-.883-.248-1.273.143zM5 5c-.312.312-.249.883.141 1.273c.391.391.961.453 1.273.141s.249-.883-.142-1.273C5.883 4.752 5.312 4.688 5 5zm8.727 9.857c.39.391.96.455 1.273.143s.249-.883-.142-1.274s-.96-.453-1.273-.141s-.248.882.142 1.272zM10 4.998c.441 0 .8-.447.8-1C10.799 3.445 10.441 3 10 3c-.442 0-.801.445-.801.998c0 .553.358 1 .801 1zM10 17c.441 0 .8-.447.8-1c0-.553-.358-.998-.799-.998c-.442 0-.801.445-.801.998c-.001.553.357 1 .8 1zm-5-7c0-.441-.45-.8-1.003-.8c-.553 0-.997.359-.997.8c0 .442.444.8.997.8C4.55 10.8 5 10.442 5 10zm12 0c0-.441-.448-.8-1.001-.8c-.553 0-.999.359-.999.8c0 .442.446.8.999.8c.553 0 1.001-.358 1.001-.8z"/>`), g.Group(children),
	)
}

func LightUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 9.199h-.98c-.553 0-1 .359-1 .801c0 .441.447.799 1 .799H19c.552 0 1-.357 1-.799c0-.441-.449-.801-1-.801zM10 4.5A5.483 5.483 0 0 0 4.5 10c0 3.051 2.449 5.5 5.5 5.5c3.05 0 5.5-2.449 5.5-5.5S13.049 4.5 10 4.5zm0 9.5c-2.211 0-4-1.791-4-4c0-2.211 1.789-4 4-4a4 4 0 0 1 0 8zm-7-4c0-.441-.449-.801-1-.801H1c-.553 0-1 .359-1 .801c0 .441.447.799 1 .799h1c.551 0 1-.358 1-.799zm7-7c.441 0 .799-.447.799-1V1c0-.553-.358-1-.799-1c-.442 0-.801.447-.801 1v1c0 .553.359 1 .801 1zm0 14c-.442 0-.801.447-.801 1v1c0 .553.359 1 .801 1c.441 0 .799-.447.799-1v-1c0-.553-.358-1-.799-1zm7.365-13.234c.391-.391.454-.961.142-1.273s-.883-.248-1.272.143l-.7.699c-.391.391-.454.961-.142 1.273s.883.248 1.273-.143l.699-.699zM3.334 15.533l-.7.701c-.391.391-.454.959-.142 1.271s.883.25 1.272-.141l.7-.699c.391-.391.454-.961.142-1.274s-.883-.247-1.272.142zm.431-12.898c-.39-.391-.961-.455-1.273-.143s-.248.883.141 1.274l.7.699c.391.391.96.455 1.272.143s.249-.883-.141-1.273l-.699-.7zm11.769 14.031l.7.699c.391.391.96.453 1.272.143c.312-.312.249-.883-.142-1.273l-.699-.699c-.391-.391-.961-.455-1.274-.143s-.248.882.143 1.273z"/>`), g.Group(children),
	)
}

func LineGraph(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m.69 11.331l1.363.338l1.026-1.611l-1.95-.482a.904.904 0 1 0-.439 1.755zm17.791.261l-4.463 4.016l-5.247-4.061a.905.905 0 0 0-.338-.162l-.698-.174l-1.027 1.611l1.1.273l5.697 4.408a.915.915 0 0 0 1.168-.043l5.028-4.527a.9.9 0 0 0 .064-1.277a.912.912 0 0 0-1.284-.064zM8.684 7.18l4.887 3.129a.913.913 0 0 0 1.24-.246l5.027-7.242a.902.902 0 0 0-.231-1.26a.91.91 0 0 0-1.265.23l-4.528 6.521l-4.916-3.147a.915.915 0 0 0-.688-.123a.914.914 0 0 0-.571.4L.142 17.209A.903.903 0 0 0 .908 18.6a.908.908 0 0 0 .768-.42l7.008-11z"/>`), g.Group(children),
	)
}

func Link(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7.859 14.691l-.81.805a1.814 1.814 0 0 1-2.545 0a1.762 1.762 0 0 1 0-2.504l2.98-2.955c.617-.613 1.779-1.515 2.626-.675a.992.992 0 1 0 1.397-1.407c-1.438-1.428-3.566-1.164-5.419.675l-2.98 2.956A3.719 3.719 0 0 0 2 14.244a3.72 3.72 0 0 0 1.108 2.658a3.779 3.779 0 0 0 2.669 1.096c.967 0 1.934-.365 2.669-1.096l.811-.805a.988.988 0 0 0 .005-1.4a.995.995 0 0 0-1.403-.006zm9.032-11.484c-1.547-1.534-3.709-1.617-5.139-.197l-1.009 1.002a.99.99 0 1 0 1.396 1.406l1.01-1.001c.74-.736 1.711-.431 2.346.197c.336.335.522.779.522 1.252s-.186.917-.522 1.251l-3.18 3.154c-1.454 1.441-2.136.766-2.427.477a.99.99 0 1 0-1.396 1.406c.668.662 1.43.99 2.228.99c.977 0 2.01-.492 2.993-1.467l3.18-3.153A3.732 3.732 0 0 0 18 5.866a3.726 3.726 0 0 0-1.109-2.659z"/>`), g.Group(children),
	)
}

func List(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.4 9H8.6c-.552 0-.6.447-.6 1s.048 1 .6 1h5.8c.552 0 .6-.447.6-1s-.048-1-.6-1zm2 5H8.6c-.552 0-.6.447-.6 1s.048 1 .6 1h7.8c.552 0 .6-.447.6-1s-.048-1-.6-1zM8.6 6h7.8c.552 0 .6-.447.6-1s-.048-1-.6-1H8.6c-.552 0-.6.447-.6 1s.048 1 .6 1zM5.4 9H3.6c-.552 0-.6.447-.6 1s.048 1 .6 1h1.8c.552 0 .6-.447.6-1s-.048-1-.6-1zm0 5H3.6c-.552 0-.6.447-.6 1s.048 1 .6 1h1.8c.552 0 .6-.447.6-1s-.048-1-.6-1zm0-10H3.6c-.552 0-.6.447-.6 1s.048 1 .6 1h1.8c.552 0 .6-.447.6-1s-.048-1-.6-1z"/>`), g.Group(children),
	)
}

func Location(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.367 18.102L18 14h-1.5l.833 4H2.667l.833-4H2L.632 18.102C.285 19.146.9 20 2 20h16c1.1 0 1.715-.854 1.367-1.898zM15 5A5 5 0 1 0 5 5c0 4.775 5 10 5 10s5-5.225 5-10zm-7.7.06A2.699 2.699 0 0 1 10 2.361a2.699 2.699 0 1 1 0 5.399a2.7 2.7 0 0 1-2.7-2.7z"/>`), g.Group(children),
	)
}

func LocationPin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 2.009c-2.762 0-5 2.229-5 4.99c0 4.774 5 11 5 11s5-6.227 5-11c0-2.76-2.238-4.99-5-4.99zm0 7.751a2.7 2.7 0 1 1 0-5.4a2.7 2.7 0 0 1 0 5.4z"/>`), g.Group(children),
	)
}

func Lock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1C7.334 1 6 2.703 6 5.6V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387A7.731 7.731 0 0 0 7.1 19h5.8a7.68 7.68 0 0 0 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8zM12 8H8V5.199C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8z"/>`), g.Group(children),
	)
}

func LockOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.8 8H14V5.6C14 2.703 12.665 1 10 1C7.334 1 6 2.703 6 5.6V6h2v-.801C8 3.754 8.797 3 10 3c1.203 0 2 .754 2 2.199V8H4c-.553 0-1 .646-1 1.199V17c0 .549.428 1.139.951 1.307l1.197.387A7.731 7.731 0 0 0 7.1 19h5.8a7.68 7.68 0 0 0 1.951-.307l1.196-.387c.524-.167.953-.757.953-1.306V9.199C17 8.646 16.352 8 15.8 8z"/>`), g.Group(children),
	)
}

func LogOut(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19 10l-6-5v3H6v4h7v3l6-5zM3 3h8V1H3c-1.1 0-2 .9-2 2v14c0 1.1.9 2 2 2h8v-2H3V3z"/>`), g.Group(children),
	)
}

func Login(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 10L8 5v3H1v4h7v3l6-5zm3 7H9v2h8c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2H9v2h8v14z"/>`), g.Group(children),
	)
}

func Loop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 7v7c0 1.103-.896 2-2 2H2c-1.104 0-2-.897-2-2V7a2 2 0 0 1 2-2h7V3l4 3.5L9 10V8H3v5h14V8h-3V5h4a2 2 0 0 1 2 2z"/>`), g.Group(children),
	)
}

func Magnet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.165 17.86c-.028.309.217.584.545.611l3.985.326c.329.027.621-.203.65-.512l.311-3.287l-5.18-.426l-.311 3.288zm-9.821-2.861l.312 3.287c.028.309.321.539.65.512l3.985-.326c.328-.027.573-.303.546-.611l-.312-3.287l-5.181.425zm-.513-5.416l.321 3.391l5.181-.426l-.322-3.387A2.949 2.949 0 0 1 7 8.911c0-1.555 1.346-2.82 3-2.82s3 1.266 3 2.82c0 .084-.004.168-.012.25l-.321 3.387l5.181.426l.321-3.391c.021-.225.03-.449.03-.672C18.2 4.659 14.522 1.2 10 1.2S1.8 4.659 1.8 8.911c0 .223.011.447.031.672z"/>`), g.Group(children),
	)
}

func MagnifyingGlass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.545 15.467l-3.779-3.779a6.15 6.15 0 0 0 .898-3.21c0-3.417-2.961-6.377-6.378-6.377A6.185 6.185 0 0 0 2.1 8.287c0 3.416 2.961 6.377 6.377 6.377a6.15 6.15 0 0 0 3.115-.844l3.799 3.801a.953.953 0 0 0 1.346 0l.943-.943c.371-.371.236-.84-.135-1.211zM4.004 8.287a4.282 4.282 0 0 1 4.282-4.283c2.366 0 4.474 2.107 4.474 4.474a4.284 4.284 0 0 1-4.283 4.283c-2.366-.001-4.473-2.109-4.473-4.474z"/>`), g.Group(children),
	)
}

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m1.574 5.286l7.5 4.029c.252.135.578.199.906.199c.328 0 .654-.064.906-.199l7.5-4.029c.489-.263.951-1.286.054-1.286H1.521c-.897 0-.435 1.023.053 1.286zm17.039 2.203l-7.727 4.027c-.34.178-.578.199-.906.199s-.566-.021-.906-.199s-7.133-3.739-7.688-4.028C.996 7.284 1 7.523 1 7.707V15c0 .42.566 1 1 1h16c.434 0 1-.58 1-1V7.708c0-.184.004-.423-.387-.219z"/>`), g.Group(children),
	)
}

func Man(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 4a2 2 0 1 0-.001-4.001A2 2 0 0 0 10 4zm5.978 7.583c-.385-1.775-1.058-4.688-2.042-5.894c-.957-1.173-2.885-1.222-3.936-1.222c-1.051 0-2.979.049-3.936 1.222c-.984 1.206-1.657 4.119-2.042 5.894c-.213.983 1.154 1.344 1.511.355c.531-1.473.941-2.71 1.839-3.736C7.844 11.109 6.102 16.168 6 19a1 1 0 0 0 1.934.358C8.391 17.771 10 13.355 10 13.355s1.609 4.416 2.066 6.003A1 1 0 0 0 14 19c-.102-2.832-1.844-7.891-1.372-10.797c.898 1.026 1.308 2.263 1.839 3.736c.356.988 1.724.627 1.511-.356z"/>`), g.Group(children),
	)
}

func Map(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19.447 3.718l-6-3a1 1 0 0 0-.895 0l-5.63 2.815l-5.606-1.869A1 1 0 0 0 0 2.613v13.001c0 .379.214.725.553.894l6 3a1.006 1.006 0 0 0 .894 0l5.63-2.814l5.606 1.869a.999.999 0 0 0 1.316-.949V4.612a.996.996 0 0 0-.552-.894zM8 5.231l4-2v11.763l-4 2V5.231zM2 4l4 1.333v11.661l-4-2V4zm16 12.227l-4-1.334V3.231l4 2v10.996z"/>`), g.Group(children),
	)
}

func Mask(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.581 5.84a.802.802 0 0 0-.92-.73C16.919 5.388 12.835 7 10 7S3.081 5.388 1.339 5.11a.801.801 0 0 0-.92.729C.277 7.371 0 11.45 0 12.068c0 .83 3.472 2.732 6 2.732c2.452 0 2.95-2.732 4-2.732s1.548 2.732 4 2.732c2.528 0 6-1.902 6-2.732c0-.618-.277-4.697-.419-6.228zM7.66 10.72c-.353.318-1.335 1.07-2.531.835c-1.196-.235-1.919-1.323-2.166-1.758a.259.259 0 0 1 .044-.317c.353-.318 1.335-1.07 2.532-.835c1.196.235 1.919 1.323 2.166 1.758a.26.26 0 0 1-.045.317zm9.377-.923c-.246.436-.969 1.523-2.166 1.758c-1.196.235-2.179-.517-2.531-.835a.26.26 0 0 1-.045-.317c.246-.436.969-1.523 2.166-1.758c1.196-.235 2.179.517 2.531.835a.258.258 0 0 1 .045.317z"/>`), g.Group(children),
	)
}

func Medal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 10c.528 0 1.026.104 1.504.256L5.427 1.141A.316.316 0 0 0 5.164 1H1.608a.147.147 0 0 0-.122.229l6.231 9.347A4.94 4.94 0 0 1 10 10zm8.392-9h-3.556a.316.316 0 0 0-.263.141L10.75 6.875l2 3l5.764-8.646A.147.147 0 0 0 18.392 1zM10 11a4 4 0 1 0 0 8a4 4 0 0 0 0-8zm2.112 4.117a.132.132 0 0 1-.022.208a.172.172 0 0 0-.049.229c.047.076.018.165-.065.199s-.125.13-.095.214s-.017.165-.104.181s-.149.101-.137.189s-.051.158-.14.155c-.089-.003-.167.068-.174.156s-.083.144-.169.123s-.178.031-.203.117s-.111.124-.191.085c-.08-.039-.18-.006-.222.072s-.134.098-.205.043s-.175-.044-.232.024s-.151.068-.209 0s-.162-.079-.232-.024s-.162.035-.205-.043s-.142-.111-.222-.072c-.08.039-.166 0-.191-.085s-.116-.138-.203-.117s-.163-.034-.169-.123s-.084-.159-.173-.157c-.089.003-.152-.067-.14-.155s-.05-.173-.137-.189s-.135-.097-.104-.181s-.013-.18-.095-.214s-.111-.123-.065-.199a.17.17 0 0 0-.049-.229a.133.133 0 0 1-.022-.208c.062-.064.062-.169 0-.234s-.052-.158.022-.208s.095-.153.049-.229c-.047-.076-.018-.165.065-.199s.125-.13.095-.214s.017-.165.104-.181s.149-.101.137-.189s.051-.158.14-.155c.089.003.167-.068.174-.156s.083-.144.169-.123s.178-.031.203-.117s.111-.124.191-.085c.08.039.18.006.222-.072s.134-.098.205-.043s.175.044.232-.024s.151-.068.209 0s.162.079.232.024s.162-.035.205.043s.142.111.222.072c.08-.039.166 0 .191.085s.116.138.203.117s.163.034.169.123s.085.159.174.156c.089-.003.152.067.14.155s.05.173.137.189s.135.097.104.181s.013.18.095.214s.111.123.065.199c-.047.076-.025.179.049.229s.083.144.022.208s-.063.171-.001.235z"/>`), g.Group(children),
	)
}

func Megaphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.223 7.03c-1.584-3.686-4.132-6.49-5.421-5.967c-2.189.891 1.304 5.164-9.447 9.533c-.929.379-1.164 1.888-.775 2.792c.388.902 1.658 1.801 2.587 1.424c.161-.066.751-.256.751-.256c.663.891 1.357.363 1.604.928l1.158 2.66c.219.502.715.967 1.075.83l2.05-.779c.468-.178.579-.596.436-.924c-.154-.355-.786-.459-.967-.873c-.18-.412-.769-1.738-.938-2.156c-.23-.568.259-1.031.97-1.104c4.894-.512 5.809 2.512 7.475 1.834c1.287-.525 1.025-4.259-.558-7.942zm-.551 5.976c-.287.115-2.213-1.402-3.443-4.267c-1.231-2.863-1.076-5.48-.79-5.597c.286-.115 2.165 1.717 3.395 4.58c1.231 2.863 1.124 5.167.838 5.284z"/>`), g.Group(children),
	)
}

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.4 9H3.6c-.552 0-.6.447-.6 1c0 .553.048 1 .6 1h12.8c.552 0 .6-.447.6-1c0-.553-.048-1-.6-1zm0 4H3.6c-.552 0-.6.447-.6 1c0 .553.048 1 .6 1h12.8c.552 0 .6-.447.6-1c0-.553-.048-1-.6-1zM3.6 7h12.8c.552 0 .6-.447.6-1c0-.553-.048-1-.6-1H3.6c-.552 0-.6.447-.6 1c0 .553.048 1 .6 1z"/>`), g.Group(children),
	)
}

func Merge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.89 17.707L16.892 20c-3.137-1.366-5.496-3.152-6.892-5.275c-1.396 2.123-3.755 3.91-6.892 5.275l-.998-2.293C5.14 16.389 8.55 14.102 8.55 10V7H5.5L10 0l4.5 7h-3.05v3c0 4.102 3.41 6.389 6.44 7.707z"/>`), g.Group(children),
	)
}

func Message(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 6v7c0 1.1-.9 2-2 2h-4v3l-4-3H4c-1.101 0-2-.9-2-2V6c0-1.1.899-2 2-2h12c1.1 0 2 .9 2 2z"/>`), g.Group(children),
	)
}

func Mic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.399 7.643V10.4c0 2.236-1.643 4.629-5.399 4.959V18h2.6c.22 0 .4.18.4.4v1.2c0 .221-.181.4-.4.4H6.4c-.22 0-.4-.18-.4-.4v-1.2c0-.22.18-.4.399-.4H9v-2.641c-3.758-.33-5.4-2.723-5.4-4.959V7.643a.4.4 0 0 1 .4-.4h.6c.22 0 .4.18.4.4V10.4c0 1.336 1.053 3.6 5 3.6c3.946 0 5-2.264 5-3.6V7.643a.4.4 0 0 1 .399-.4H16a.399.399 0 0 1 .399.4zM10 12c2.346 0 3-.965 3-1.6V7.242H7V10.4c0 .635.652 1.6 3 1.6zm3-10.4c0-.637-.654-1.6-3-1.6c-2.348 0-3 .963-3 1.6v4.242h6V1.6z"/>`), g.Group(children),
	)
}

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 10c0 .553-.048 1-.601 1H4.601C4.049 11 4 10.553 4 10c0-.553.049-1 .601-1H15.4c.552 0 .6.447.6 1z"/>`), g.Group(children),
	)
}

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.004 0H5.996A1.996 1.996 0 0 0 4 1.996v16.007C4 19.106 4.894 20 5.996 20h8.007A1.997 1.997 0 0 0 16 18.004V1.996A1.996 1.996 0 0 0 14.004 0zM10 19c-.69 0-1.25-.447-1.25-1s.56-1 1.25-1s1.25.447 1.25 1s-.56 1-1.25 1zm4-3H6V2h8v14z"/>`), g.Group(children),
	)
}

func ModernMic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.228 10.891a.528.528 0 0 0-.159.69l1.296 2.244c.133.23.438.325.677.208L7 12.116V19h2v-7.854l4.071-1.973l-2.62-4.54l-9.223 6.258zm17.229-7.854a4.061 4.061 0 0 0-5.546-1.484c-.91.525-1.508 1.359-1.801 2.289l2.976 5.156c.951.212 1.973.11 2.885-.415a4.06 4.06 0 0 0 1.486-5.546z"/>`), g.Group(children),
	)
}

func Moon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.719 1.8A8.759 8.759 0 1 1 1.8 13.719c3.335 1.867 7.633 1.387 10.469-1.449c2.837-2.837 3.318-7.134 1.45-10.47z"/>`), g.Group(children),
	)
}

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15.402 14.402l-2.627-7.535c-.722-2.073-2.966-3.22-5.092-2.653L5.707.379a.687.687 0 0 0-.938-.296a.719.719 0 0 0-.289.961l1.929 3.742C4.872 5.806 4.073 7.74 4.58 9.56l2.139 7.696c.602 2.162 3.08 3.264 5.571 2.502c2.459-.863 3.85-3.237 3.112-5.356zM8.899 8.923a1.38 1.38 0 0 1-1.745-.921c-.235-.748.168-1.548.897-1.788c.73-.24 1.512.172 1.746.92s-.168 1.549-.898 1.789z"/>`), g.Group(children),
	)
}

func Music(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 1H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zm-3.205 10.519c-.185.382-.373.402-.291 0C12.715 10.48 12.572 8.248 11 8v4.75c0 .973-.448 1.82-1.639 2.203c-1.156.369-2.449-.016-2.752-.846c-.303-.83.377-1.84 1.518-2.256c.637-.232 1.375-.292 1.873-.101V5h1c0 2.355 4.065 1.839 1.795 6.519z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func Network(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.274 6.915c.2 0 .394.029.576.086a15.774 15.774 0 0 1 2.283-2.1a1.954 1.954 0 0 1 .048-1.076A14.407 14.407 0 0 0 5.17 2.171a9.25 9.25 0 0 0-2.582 2.381c.519.92 1.136 1.777 1.838 2.557c.256-.124.543-.194.848-.194zM3.316 8.872c0-.275.058-.537.159-.773A15.91 15.91 0 0 1 1.78 5.87a9.165 9.165 0 0 0-.98 4.131a9.16 9.16 0 0 0 1.295 4.705a15.614 15.614 0 0 1 1.62-4.652a1.947 1.947 0 0 1-.399-1.182zm6.72-6.383c.517 0 .985.201 1.336.529a15.578 15.578 0 0 1 3.215-.992A9.154 9.154 0 0 0 10 .8a9.167 9.167 0 0 0-3.236.588a15.76 15.76 0 0 1 2.277 1.375c.292-.174.631-.274.995-.274zm2.926 9.219a1.94 1.94 0 0 1 .509-.656a14.336 14.336 0 0 0-2.672-4.803a1.956 1.956 0 0 1-1.901-.211a14.343 14.343 0 0 0-1.964 1.803a1.93 1.93 0 0 1 .207 1.617a14.252 14.252 0 0 0 5.821 2.25zm2.539 2.643a15.872 15.872 0 0 1-.081 3.082a9.216 9.216 0 0 0 3.347-4.639a15.39 15.39 0 0 1-2.181.365a1.958 1.958 0 0 1-1.085 1.192zm-2.997-1.327a15.643 15.643 0 0 1-6.21-2.484a1.953 1.953 0 0 1-1.423.248a14.219 14.219 0 0 0-1.599 5.484a9.203 9.203 0 0 0 3.145 2.205a15.662 15.662 0 0 1 6.087-5.453zm3.672-9.843a14.296 14.296 0 0 0-4.193 1.068a1.946 1.946 0 0 1-.191 1.056a15.68 15.68 0 0 1 2.969 5.291a1.961 1.961 0 0 1 1.77 1.195c.886-.09 1.748-.26 2.578-.504a9.178 9.178 0 0 0-2.933-8.106zm-2.687 10.888a14.291 14.291 0 0 0-5.723 4.856A9.187 9.187 0 0 0 10 19.2a9.165 9.165 0 0 0 3.882-.859c.19-.928.29-1.887.29-2.869c0-.355-.016-.707-.043-1.055a1.923 1.923 0 0 1-.64-.348z"/>`), g.Group(children),
	)
}

func New(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m18.69 12.344l-1.727-1.538c-.5-.445-.5-1.174 0-1.619l1.727-1.539c.5-.445.367-.859-.296-.924l-2.29-.217c-.662-.062-1.02-.633-.791-1.266l1.215-3.383c.228-.635-.051-.865-.619-.514l-2.701 1.67a1.158 1.158 0 0 1-1.631-.426L10.599.842C10.27.254 9.727.252 9.392.834l-.909 1.58c-.337.585-1.108.833-1.713.556l-1.6-.734c-.608-.28-1.073.042-1.037.716l.086 1.615c.037.674-.461 1.367-1.104 1.541l-1.545.414c-.642.174-.76.68-.26 1.125l1.727 1.539c.5.445.5 1.174 0 1.619L1.31 12.344c-.5.445-.368.877.293.957l2.095.254c.661.08 1.029.67.818 1.311l-1.074 3.258c-.211.641.09.889.668.555l2.463-1.426a1.321 1.321 0 0 1 1.729.408L9.324 19.2c.372.559.931.529 1.24-.068l.899-1.733a1.243 1.243 0 0 1 1.648-.543l1.734.867c.598.297 1.057-.01 1.021-.682l-.087-1.617c-.035-.674.461-1.365 1.106-1.539l1.543-.416c.644-.174.762-.68.262-1.125zM11 14H9v-2h2v2zm0-3H9V6h2v5z"/>`), g.Group(children),
	)
}

func NewMessage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.174 1.826c-1.102-1.102-2.082-.777-2.082-.777L7.453 9.681L6 14l4.317-1.454l8.634-8.638s.324-.98-.777-2.082zm-7.569 9.779l-.471.47l-1.473.5a2.216 2.216 0 0 0-.498-.74a2.226 2.226 0 0 0-.74-.498l.5-1.473l.471-.47s.776-.089 1.537.673c.762.761.674 1.538.674 1.538zM16 17H3V4h5l2-2H3c-1.1 0-2 .9-2 2v13c0 1.1.9 2 2 2h13c1.1 0 2-.9 2-2v-7l-2 2v5z"/>`), g.Group(children),
	)
}

func News(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M14 5h-4v2h4V5zm0 3h-4v1h4V8zM9 5H6v4h3V5zm0 6h5v-1H9v1zm3 2h2v-1h-2v1zm2 1H6v1h8v-1zm-3-2H6v1h5v-1zm-3-2H6v1h2v-1zm9-9H3a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zm-1 16H4V3h12v14z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func Note(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.971 9.438c-.422.656-.646.375-.52 0c.336-.993.348-4.528-2.451-4.969L11.998 16c0 1.657-1.735 4-4.998 4c-1.657 0-3-.871-3-2.5c0-2.119 1.927-3.4 4-3.4c1.328 0 2 .4 2 .4V0h2c0 2.676 5.986 4.744 2.971 9.438z"/>`), g.Group(children),
	)
}

func Notification(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 8.38V17H3V5h8.62c-.073-.322-.12-.655-.12-1s.047-.678.12-1H3c-1.102 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V8.38c-.322.073-.655.12-1 .12s-.678-.047-1-.12zM16 1a3 3 0 1 0 0 6a3 3 0 0 0 0-6z"/>`), g.Group(children),
	)
}

func OldMobile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.6 3H7V0H5v18.6c0 .77.629 1.4 1.398 1.4H13.6c.769 0 1.4-.631 1.4-1.4V4.401C15 3.629 14.369 3 13.6 3zM8 15c-.691 0-1.25-.447-1.25-1s.559-1 1.25-1s1.25.447 1.25 1s-.559 1-1.25 1zm1.25 2c0 .553-.559 1-1.25 1s-1.25-.447-1.25-1s.559-1 1.25-1s1.25.447 1.25 1zM7 11V5h6v6H7zm5 4c-.691 0-1.25-.447-1.25-1s.559-1 1.25-1s1.25.447 1.25 1s-.559 1-1.25 1zm1.25 2c0 .553-.559 1-1.25 1s-1.25-.447-1.25-1s.559-1 1.25-1s1.25.447 1.25 1z"/>`), g.Group(children),
	)
}

func OldPhone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.256 12.253c-.096-.667-.611-1.187-1.274-1.342c-2.577-.604-3.223-2.088-3.332-3.734C12.193 7.092 11.38 7 10 7s-2.193.092-2.65.177c-.109 1.646-.755 3.13-3.332 3.734c-.663.156-1.178.675-1.274 1.342l-.497 3.442C2.072 16.907 2.962 18 4.2 18h11.6c1.237 0 2.128-1.093 1.953-2.305l-.497-3.442zM10 15.492c-1.395 0-2.526-1.12-2.526-2.5s1.131-2.5 2.526-2.5s2.526 1.12 2.526 2.5s-1.132 2.5-2.526 2.5zM19.95 6c-.024-1.5-3.842-3.999-9.95-4C3.891 2.001.073 4.5.05 6s.021 3.452 2.535 3.127c2.941-.381 2.76-1.408 2.76-2.876C5.345 5.227 7.737 4.98 10 4.98s4.654.247 4.655 1.271c0 1.468-.181 2.495 2.76 2.876C19.928 9.452 19.973 7.5 19.95 6z"/>`), g.Group(children),
	)
}

func OpenBook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10.595 5.196l.446 1.371a4.135 4.135 0 0 1 1.441-.795c.59-.192 1.111-.3 1.582-.362l-.43-1.323a9.465 9.465 0 0 0-1.58.368a5.25 5.25 0 0 0-1.459.741zm.927 2.855l.446 1.371a4.135 4.135 0 0 1 1.441-.795c.59-.192 1.111-.3 1.582-.362l-.43-1.323a9.465 9.465 0 0 0-1.58.368a5.21 5.21 0 0 0-1.459.741zm.928 2.854l.446 1.371a4.135 4.135 0 0 1 1.441-.795c.59-.192 1.111-.3 1.582-.362l-.43-1.323a9.465 9.465 0 0 0-1.58.368a5.21 5.21 0 0 0-1.459.741zm-7.062 2.172l.43 1.323a8.745 8.745 0 0 1 1.492-.636a4.141 4.141 0 0 1 1.633-.203l-.446-1.371a5.25 5.25 0 0 0-1.615.257a9.406 9.406 0 0 0-1.494.63zM3.533 7.368l.43 1.323a8.825 8.825 0 0 1 1.492-.636a4.141 4.141 0 0 1 1.633-.203L6.643 6.48a5.263 5.263 0 0 0-1.616.258a9.406 9.406 0 0 0-1.494.63zm.927 2.855l.43 1.323a8.745 8.745 0 0 1 1.492-.636a4.141 4.141 0 0 1 1.633-.203L7.57 9.335a5.25 5.25 0 0 0-1.615.257a9.417 9.417 0 0 0-1.495.631zm6.604-8.813a5.26 5.26 0 0 0-3.053 2.559a5.257 5.257 0 0 0-3.973-.275C1.515 4.514.069 6.321.069 6.321l4.095 12.587c.126.387.646.477.878.143c.499-.719 1.46-1.658 3.257-2.242c1.718-.558 2.969.054 3.655.578c.272.208.662.06.762-.268c.252-.827.907-2.04 2.61-2.593c1.799-.585 3.129-.389 3.956-.1c.385.134.75-.242.625-.629L15.819 1.203s-2.232-.612-4.755.207zm-.113 13.846a5.208 5.208 0 0 0-3.141.044c-1.251.406-2.127.949-2.699 1.404L1.866 6.722c.358-.358 1.187-1.042 2.662-1.521c1.389-.451 2.528-.065 3.279.378l3.144 9.677zm6.894-2.689c-.731-.032-1.759.044-3.01.451a5.205 5.205 0 0 0-2.567 1.81L9.124 5.151c.346-.8 1.04-1.782 2.43-2.233c1.474-.479 2.547-.413 3.047-.334l3.244 9.983z"/>`), g.Group(children),
	)
}

func Palette(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.74 2.608c-3.528-1.186-7.066-.961-10.72 1.274C2.167 5.625.302 9.958.917 13.064c.728 3.671 4.351 5.995 9.243 4.651c5.275-1.449 6.549-4.546 6.379-5.334c-.17-.788-2.665-1.652-1.718-3.498c1.188-2.313 3.129-1.149 3.982-1.622c.855-.472.539-3.442-3.063-4.653zm-3.646 10.706a1.504 1.504 0 0 1-1.843-1.059a1.5 1.5 0 0 1 1.046-1.849a1.503 1.503 0 0 1 1.843 1.059a1.501 1.501 0 0 1-1.046 1.849z"/>`), g.Group(children),
	)
}

func PaperPlane(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.64 2.634L.984 8.856c-.284.1-.347.345-.01.479l3.796 1.521l2.25.901l10.984-8.066c.148-.108.318.095.211.211l-7.871 8.513v.002l-.452.503l.599.322l4.982 2.682c.291.156.668.027.752-.334l2.906-12.525c.079-.343-.148-.552-.491-.431zM7 17.162c0 .246.139.315.331.141c.251-.229 2.85-2.561 2.85-2.561L7 13.098v4.064z"/>`), g.Group(children),
	)
}

func Pencil(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.69 2.661c-1.894-1.379-3.242-1.349-3.754-1.266a.538.538 0 0 0-.35.223l-4.62 6.374l-2.263 3.123a2.447 2.447 0 0 0-.462 1.307l-.296 5.624a.56.56 0 0 0 .76.553l5.256-2.01c.443-.17.828-.465 1.106-.849l1.844-2.545l5.036-6.949a.56.56 0 0 0 .1-.423c-.084-.526-.487-1.802-2.357-3.162zM8.977 15.465l-2.043.789a.19.19 0 0 1-.221-.062a5.145 5.145 0 0 0-1.075-1.03a5.217 5.217 0 0 0-1.31-.706a.192.192 0 0 1-.126-.192l.122-2.186l.549-.755s1.229-.169 2.833.998c1.602 1.166 1.821 2.388 1.821 2.388l-.55.756z"/>`), g.Group(children),
	)
}

func Phone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.229 11.229c-1.583 1.582-3.417 3.096-4.142 2.371c-1.037-1.037-1.677-1.941-3.965-.102c-2.287 1.838-.53 3.064.475 4.068c1.16 1.16 5.484.062 9.758-4.211c4.273-4.274 5.368-8.598 4.207-9.758c-1.005-1.006-2.225-2.762-4.063-.475c-1.839 2.287-.936 2.927.103 3.965c.722.725-.791 2.559-2.373 4.142z"/>`), g.Group(children),
	)
}

func PieChart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 .958v9.039C11 10.551 10.551 11 9.997 11H.958A9.1 9.1 0 1 0 11 .958zm-2 0A9.098 9.098 0 0 0 .958 9H9V.958z"/>`), g.Group(children),
	)
}

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m4.774 15.287l-2.105 3.25l.224 1.063l1.06-.227l2.104-3.248a8.352 8.352 0 0 1-1.283-.838zm8.912-1.135c.014-.029.023-.061.036-.092c.053-.117.1-.234.138-.357c.006-.022.009-.044.016-.064a4.48 4.48 0 0 0 .098-.408v-.021c.195-1.169-.145-2.473-.923-3.651l1.11-1.714c1.279.163 2.385-.159 2.917-.982c.923-1.423-.2-3.792-2.505-5.293C12.266.068 9.65.005 8.729 1.426c-.534.824-.378 1.967.293 3.073L7.91 6.213c-1.389-.233-2.716-.016-3.703.64c-.006.002-.013.004-.017.008a3.735 3.735 0 0 0-.332.254c-.017.014-.037.027-.051.041a3.024 3.024 0 0 0-.271.272c-.02.024-.048.045-.067.07a3.102 3.102 0 0 0-.29.385c-1.384 2.133-.203 5.361 2.633 7.209c2.838 1.848 6.26 1.614 7.641-.519c.087-.135.167-.276.233-.421zm-.815-9.958c-.887-.577-1.32-1.487-.965-2.036c.354-.547 1.361-.522 2.246.055c.889.577 1.318 1.489.965 2.036c-.353.547-1.358.522-2.246-.055z"/>`), g.Group(children),
	)
}

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 10c0 .553-.048 1-.601 1H11v4.399c0 .552-.447.601-1 .601c-.553 0-1-.049-1-.601V11H4.601C4.049 11 4 10.553 4 10c0-.553.049-1 .601-1H9V4.601C9 4.048 9.447 4 10 4c.553 0 1 .048 1 .601V9h4.399c.553 0 .601.447.601 1z"/>`), g.Group(children),
	)
}

func Popup(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 2H7.979C6.88 2 6 2.88 6 3.98V12c0 1.1.9 2 2 2h8c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm0 10H8V4h8v8zM4 10H2v6c0 1.1.9 2 2 2h6v-2H4v-6z"/>`), g.Group(children),
	)
}

func PowerPlug(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 14v1.498c0 .277.225.502.502.502h.997A.502.502 0 0 0 2 15.498V14c0-.959.801-2.273 2-2.779V9.116C1.684 9.652 0 11.97 0 14zm12.065-9.299l-2.53 1.898c-.347.26-.769.401-1.203.401H6.005C5.45 7 5 7.45 5 8.005v3.991C5 12.55 5.45 13 6.005 13h2.327c.434 0 .856.141 1.203.401l2.531 1.898a3.502 3.502 0 0 0 2.102.701H16V4h-1.832c-.758 0-1.496.246-2.103.701zM17 6v2h3V6h-3zm0 8h3v-2h-3v2z"/>`), g.Group(children),
	)
}

func PriceRibbon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.825 10.653c.118-.258.445-.497.727-.529s.539-.29.571-.572c.034-.28.272-.608.529-.727a.69.69 0 0 0 .369-.72c-.058-.278.068-.663.276-.854a.689.689 0 0 0 .127-.801a1.017 1.017 0 0 1 0-.897a.688.688 0 0 0-.127-.801c-.208-.193-.333-.577-.276-.854a.691.691 0 0 0-.369-.722a1.03 1.03 0 0 1-.529-.727a.689.689 0 0 0-.571-.572a1.024 1.024 0 0 1-.727-.528a.686.686 0 0 0-.722-.366a1.024 1.024 0 0 1-.854-.278c-.193-.21-.553-.266-.8-.127s-.652.139-.898 0a.684.684 0 0 0-.801.125a1.022 1.022 0 0 1-.854.278a.685.685 0 0 0-.72.367c-.119.256-.446.495-.728.527a.69.69 0 0 0-.572.573a1.023 1.023 0 0 1-.529.726a.69.69 0 0 0-.366.722c.055.277-.07.662-.278.854s-.266.552-.127.801c.139.246.139.651 0 .897a.69.69 0 0 0 .127.802c.209.19.333.575.278.854a.687.687 0 0 0 .366.72c.258.119.495.447.528.727c.034.282.29.54.572.572s.609.272.728.529a.688.688 0 0 0 .72.366c.278-.055.663.069.854.278a.69.69 0 0 0 .801.127c.246-.139.651-.139.898 0s.607.081.8-.127c.193-.21.576-.333.854-.278a.69.69 0 0 0 .723-.365zM10 9.399a3.4 3.4 0 1 1 0-6.8a3.4 3.4 0 0 1 0 6.8zm-4.025 2.01l-1.243 7.049l3.128-.464l2.781 1.506l1.238-7.021a6.707 6.707 0 0 1-5.904-1.07zm7.986.048a6.741 6.741 0 0 1-.99.597l-.748 4.236l3.369-1.828l-1.631-3.005z"/>`), g.Group(children),
	)
}

func PriceTag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.388.405a.605.605 0 0 0-1.141.399c.929 2.67-.915 4.664-2.321 5.732l-.568-.814c-.191-.273-.618-.5-.95-.504l-3.188.014a2.162 2.162 0 0 0-1.097.338L.729 12.157a1.01 1.01 0 0 0-.247 1.404l4.269 6.108c.32.455.831.4 1.287.082l9.394-6.588c.27-.191.582-.603.692-.918l.998-3.145c.11-.314.043-.793-.148-1.066l-.346-.496c1.888-1.447 3.848-4.004 2.76-7.133zm-4.371 9.358a1.608 1.608 0 0 1-2.24-.396a1.614 1.614 0 0 1 .395-2.246a1.607 1.607 0 0 1 1.868.017c-.272.164-.459.26-.494.275a.606.606 0 0 0 .259 1.153c.086 0 .174-.02.257-.059c.194-.092.402-.201.619-.33a1.615 1.615 0 0 1-.664 1.586z"/>`), g.Group(children),
	)
}

func Print(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.501 6h17c.57 0 .477-.608.193-.707C18.409 5.194 15.251 4 14.7 4H14V1H6v3h-.699c-.55 0-3.709 1.194-3.993 1.293c-.284.099-.377.707.193.707zM19 7H1c-.55 0-1 .45-1 1v5c0 .551.45 1 1 1h2.283l-.882 5H17.6l-.883-5H19c.551 0 1-.449 1-1V8c0-.55-.449-1-1-1zM4.603 17l1.198-7.003H14.2L15.399 17H4.603z"/>`), g.Group(children),
	)
}

func ProgressEmpty(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 5H2C.9 5 0 5.9 0 7v6c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V7c0-1.1-.9-2-2-2zm0 8H2V7h16v6z"/>`), g.Group(children),
	)
}

func ProgressFull(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 5H2C.9 5 0 5.9 0 7v6c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V7c0-1.1-.9-2-2-2zm0 8H2V7h16v6zM7 8H3v4h4V8zm5 0H8v4h4V8zm5 0h-4v4h4V8z"/>`), g.Group(children),
	)
}

func ProgressOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 5H2C.9 5 0 5.9 0 7v6c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V7c0-1.1-.9-2-2-2zm0 8H2V7h16v6zM7 8H3v4h4V8z"/>`), g.Group(children),
	)
}

func ProgressTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 5H2C.9 5 0 5.9 0 7v6c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V7c0-1.1-.9-2-2-2zm0 8H2V7h16v6zM7 8H3v4h4V8zm5 0H8v4h4V8z"/>`), g.Group(children),
	)
}

func Publish(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.967 8.193L5 13h3v6h4v-6h3L9.967 8.193zM18 1H2C.9 1 0 1.9 0 3v12c0 1.1.9 2 2 2h4v-2H2V6h16v9h-4v2h4c1.1 0 2-.9 2-2V3c0-1.1-.9-2-2-2zM2.5 4.25a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5zm2 0a.75.75 0 1 1 0-1.5a.75.75 0 0 1 0 1.5zM18 4H6V3h12.019L18 4z"/>`), g.Group(children),
	)
}

func Quote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.315 3.401c-1.61 0-2.916 1.343-2.916 3c0 1.656 1.306 3 2.916 3c2.915 0 .972 5.799-2.916 5.799v1.4c6.939.001 9.658-13.199 2.916-13.199zm8.4 0c-1.609 0-2.915 1.343-2.915 3c0 1.656 1.306 3 2.915 3c2.916 0 .973 5.799-2.915 5.799v1.4c6.938.001 9.657-13.199 2.915-13.199z"/>`), g.Group(children),
	)
}

func Radio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 8H5.021l8.974-5.265L13 1L1.736 7.571A1.482 1.482 0 0 0 1 8.852V17a2 2 0 0 0 2 2h14a2 2 0 0 0 2-2v-7a2 2 0 0 0-2-2zm-1.5 9a1.5 1.5 0 1 1 .001-3.001A1.5 1.5 0 0 1 15.5 17zm1.5-5H3v-2h14v2z"/>`), g.Group(children),
	)
}

func RemoveUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.989 19.129c0-2.246-2.187-3.389-4.317-4.307c-2.123-.914-2.801-1.684-2.801-3.334c0-.989.648-.667.932-2.481c.12-.752.692-.012.802-1.729c0-.684-.313-.854-.313-.854s.159-1.013.221-1.793c.064-.817-.398-2.56-2.301-3.095c-.332-.341-.557-.882.467-1.424c-2.24-.104-2.761 1.068-3.954 1.93c-1.015.756-1.289 1.953-1.24 2.59c.065.78.223 1.793.223 1.793s-.314.17-.314.854c.11 1.718.684.977.803 1.729c.284 1.814.933 1.492.933 2.481c0 1.65-.212 2.21-2.336 3.124C.663 15.53 0 17 .011 19.129C.014 19.766 0 20 0 20h16s-.011-.234-.011-.871zm.011-9.09l-2.299-2.398l-1.061 1.061L15.039 11l-2.398 2.298l1.061 1.061L16 11.961l2.298 2.398l1.061-1.061L16.961 11l2.397-2.298l-1.061-1.061L16 10.039z"/>`), g.Group(children),
	)
}

func Reply(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 16.685S16.775 6.953 8 6.953V2.969L1 9.542l7 6.69v-4.357c4.763-.001 8.516.421 11 4.81z"/>`), g.Group(children),
	)
}

func ReplyAll(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.225 5.767V3.086L0 9.542l7.225 6.691v-2.777L3 9.542l4.225-3.775zm5 1.186V3.086L5 9.542l7.225 6.691v-4.357c3.292 0 5.291.422 7.775 4.81c0-.001-.368-9.733-7.775-9.733z"/>`), g.Group(children),
	)
}

func ResizeFullScreen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6.987 10.987l-2.931 3.031L2 11.589V18h6.387l-2.43-2.081l3.03-2.932l-2-2zM11.613 2l2.43 2.081l-3.03 2.932l2 2l2.931-3.031L18 8.411V2h-6.387z"/>`), g.Group(children),
	)
}

func ResizeOneHundred(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.1 14.1L1 17l2 2l2.9-3.1L8 18v-6H2l2.1 2.1zM19 3l-2-2l-2.9 3.1L12 2v6h6l-2.1-2.1L19 3z"/>`), g.Group(children),
	)
}

func Retweet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 13V8h2L3.5 4L0 8h2v6a2 2 0 0 0 2 2h9.482l-2.638-3H5zm4.156-6L6.518 4H16c1.104 0 2 .897 2 2v6h2l-3.5 4l-3.5-4h2V7H9.156z"/>`), g.Group(children),
	)
}

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.933 13.069s7.059-5.094 6.276-10.924a.465.465 0 0 0-.112-.268a.436.436 0 0 0-.263-.115C12.137.961 7.16 8.184 7.16 8.184c-4.318-.517-4.004.344-5.974 5.076c-.377.902.234 1.213.904.959l2.148-.811l2.59 2.648l-.793 2.199c-.248.686.055 1.311.938.926c4.624-2.016 5.466-1.694 4.96-6.112zm1.009-5.916a1.594 1.594 0 0 1 0-2.217a1.509 1.509 0 0 1 2.166 0a1.594 1.594 0 0 1 0 2.217a1.509 1.509 0 0 1-2.166 0z"/>`), g.Group(children),
	)
}

func RoundBrush(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.135 9.304c-2.558-1.879-6.7 1.17-7.632 5.284c-.718 3.17-4.043 3.04-3.996 3.464c.046.424 7.473 1.103 10.156-1.123c2.506-2.08 4.277-5.564 1.472-7.625zm2.203-7.796l-3.363 5.156c1.102.179 3.635 1.885 4.121 3.109l4.396-4.246c-.526-1.503-3.438-3.844-5.154-4.019z"/>`), g.Group(children),
	)
}

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.4 2.4v2.367c7.086 0 12.83 5.746 12.83 12.832h2.369C17.599 9.205 10.794 2.4 2.4 2.4zm0 4.737v2.369a8.093 8.093 0 0 1 8.093 8.094h2.368c0-5.778-4.684-10.463-10.461-10.463zm2.269 5.922a2.271 2.271 0 0 0 0 4.541c1.254 0 2.269-1.016 2.269-2.27s-1.015-2.271-2.269-2.271z"/>`), g.Group(children),
	)
}

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.249.438L.438 14.251a1.505 1.505 0 0 0 .002 2.124l3.185 3.187a1.506 1.506 0 0 0 2.124.002L19.562 5.751a1.508 1.508 0 0 0 0-2.125L16.376.438a1.51 1.51 0 0 0-2.127 0zM3.929 15.312l-.759.759l-1.896-1.897l.759-.759l1.896 1.897zm3.036 0l-.759.759l-3.415-3.415l.759-.76l3.415 3.416zm0-3.036l-.759.759l-1.898-1.896l.76-.76l1.897 1.897zm1.518-1.518l-.759.759l-1.896-1.896l.759-.76l1.896 1.897zm3.035 0l-.759.759l-3.414-3.414l.759-.759l3.414 3.414zm0-3.035l-.759.759l-1.896-1.896l.759-.759l1.896 1.896zm1.518-1.517l-.759.759l-1.897-1.897l.759-.759l1.897 1.897zm3.036 0l-.76.759l-3.414-3.415l.759-.76l3.415 3.416zm-.001-3.035l-.759.759l-1.896-1.898l.759-.758l1.896 1.897z"/>`), g.Group(children),
	)
}

func Save(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.173 2H4c-1.101 0-2 .9-2 2v12c0 1.1.899 2 2 2h12c1.101 0 2-.9 2-2V5.127L15.173 2zM14 8c0 .549-.45 1-1 1H7c-.55 0-1-.451-1-1V3h8v5zm-1-4h-2v4h2V4z"/>`), g.Group(children),
	)
}

func Scissors(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.38 5.59a3.69 3.69 0 1 0-3.69 3.69a3.67 3.67 0 0 0 2.483-.976L9 9.991l.012.009l-.004.003l-1.836 1.693a3.665 3.665 0 0 0-2.482-.976a3.69 3.69 0 1 0 3.69 3.69c0-.297-.044-.582-.111-.858l2.844-1.991l4.127 3.065c2.212 1.549 3.76-.663 3.76-.663L8.269 6.448c.066-.276.111-.561.111-.858zm-3.69 1.8a1.8 1.8 0 1 1 0-3.6a1.8 1.8 0 0 1 0 3.6zm0 8.82a1.8 1.8 0 1 1 0-3.6a1.8 1.8 0 0 1 0 3.6zM19 6.038s-1.548-2.212-3.76-.663L12.035 7.61l2.354 1.648L19 6.038z"/>`), g.Group(children),
	)
}

func SelectArrows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 1L5 8h10l-5-7zm0 18l5-7H5l5 7z"/>`), g.Group(children),
	)
}

func Share(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 13.442c-.633 0-1.204.246-1.637.642l-5.938-3.463c.046-.188.075-.384.075-.584s-.029-.395-.075-.583L13.3 6.025A2.48 2.48 0 0 0 15 6.7c1.379 0 2.5-1.121 2.5-2.5S16.379 1.7 15 1.7s-2.5 1.121-2.5 2.5c0 .2.029.396.075.583L6.7 8.212A2.485 2.485 0 0 0 5 7.537c-1.379 0-2.5 1.121-2.5 2.5s1.121 2.5 2.5 2.5a2.48 2.48 0 0 0 1.7-.675l5.938 3.463a2.339 2.339 0 0 0-.067.546A2.428 2.428 0 1 0 15 13.442z"/>`), g.Group(children),
	)
}

func ShareAlternitive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 13h2V4h2l-3-4l-3 4h2v9zm8-6h-3v2h2v9H4V9h2V7H3a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h14a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1z"/>`), g.Group(children),
	)
}

func Shareable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.8 10a3.2 3.2 0 1 0 6.401 0A3.2 3.2 0 0 0 6.8 10zM4.529 8.8a5.6 5.6 0 0 1 9.43-2.76a1.2 1.2 0 1 0 1.697-1.697A8.002 8.002 0 0 0 2.367 7.601H0V10h3.199c.999 0 1.245-.813 1.33-1.2zM16.8 10c-.999 0-1.245.814-1.329 1.199a5.6 5.6 0 0 1-9.43 2.759a1.2 1.2 0 0 0-1.698 1.697A7.972 7.972 0 0 0 10 18a8.005 8.005 0 0 0 7.633-5.6H20V10h-3.2z"/>`), g.Group(children),
	)
}

func Shield(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.604 3.332C12.99 4 12.075 2.833 10 1C7.925 2.833 7.01 4 2.396 3.332C-.063 15.58 10 19 10 19s10.063-3.42 7.604-15.668zm-5.131 9.977L10 12.009l-2.472 1.3L8 10.556l-2-1.95l2.764-.401L10 5.7l1.236 2.505L14 8.606l-2 1.949l.473 2.754z"/>`), g.Group(children),
	)
}

func Shop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.123 7.25L6.914 2H2.8L1.081 6.5C1.028 6.66 1 6.826 1 7c0 1.104 1.15 2 2.571 2c1.31 0 2.393-.764 2.552-1.75zM10 9c1.42 0 2.571-.896 2.571-2c0-.041-.003-.082-.005-.121L12.057 2H7.943l-.51 4.875A2.527 2.527 0 0 0 7.429 7c0 1.104 1.151 2 2.571 2zm5 1.046V14H5v-3.948c-.438.158-.92.248-1.429.248c-.195 0-.384-.023-.571-.049V16.6c0 .77.629 1.4 1.398 1.4H15.6c.77 0 1.4-.631 1.4-1.4v-6.348a4.297 4.297 0 0 1-.571.049A4.155 4.155 0 0 1 15 10.046zM18.92 6.5L17.199 2h-4.113l.79 5.242C14.03 8.232 15.113 9 16.429 9C17.849 9 19 8.104 19 7c0-.174-.028-.34-.08-.5z"/>`), g.Group(children),
	)
}

func ShoppingBag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.121 3.271c-.295-.256-1.906-1.731-2.207-1.991c-.299-.259-.756-.28-1.102-.28H5.188c-.345 0-.802.021-1.102.28c-.301.26-1.912 1.736-2.207 1.991c-.297.256-.543.643-.464 1.192c.079.551 1.89 13.661 1.937 13.973A.67.67 0 0 0 4 19h12a.67.67 0 0 0 .648-.565c.047-.311 1.858-13.422 1.938-13.973c.078-.548-.168-.935-.465-1.191zM10 11.973c-3.248 0-3.943-4.596-4.087-5.543H7.75c.276 1.381.904 3.744 2.25 3.744s1.975-2.363 2.25-3.744h1.838c-.145.947-.84 5.543-4.088 5.543zM3.17 4.006L5 2h10l1.83 2.006H3.17z"/>`), g.Group(children),
	)
}

func ShoppingBasket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.399 7h-5.007L11.58 8.812a2.384 2.384 0 0 1-1.696.702c-.642 0-1.244-.25-1.698-.703a2.387 2.387 0 0 1-.703-1.695c0-.039.01-.077.011-.116H1.6a.6.6 0 0 0-.6.6V10h18V7.6a.6.6 0 0 0-.601-.6zm-7.631.999l5.055-5.055a.6.6 0 0 0 .002-.849l-.92-.92a.603.603 0 0 0-.85 0L9 6.231a1.25 1.25 0 0 0 1.768 1.768zm-6.945 9.272c.097.401.515.729.927.729h10.5c.412 0 .83-.328.927-.729L17.7 11H2.3l1.523 6.271z"/>`), g.Group(children),
	)
}

func ShoppingCart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 17a2 2 0 1 0 3.999.001A2 2 0 0 0 13 17zM3 17a2 2 0 1 0 4 0a2 2 0 0 0-4 0zm3.547-4.828L17.615 9.01A.564.564 0 0 0 18 8.5V3H4V1.4c0-.22-.181-.4-.399-.4H.399A.401.401 0 0 0 0 1.4V3h2l1.91 8.957l.09.943v1.649c0 .219.18.4.4.4h13.2c.22 0 .4-.182.4-.4V13H6.752c-1.15 0-1.174-.551-.205-.828z"/>`), g.Group(children),
	)
}

func Shuffle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.093 6.694h.92v2.862L20 5.532l-3.988-4.025v2.387h-.92c-3.694 0-5.776 2.738-7.614 5.152c-1.652 2.172-3.08 4.049-5.386 4.049H0v2.799h2.093c3.694 0 5.776-2.736 7.614-5.152c1.652-2.173 3.08-4.048 5.386-4.048zM5.41 8.458c.158-.203.316-.412.477-.623a47.33 47.33 0 0 1 1.252-1.596C5.817 5.005 4.224 4.095 2.093 4.095H0v2.799h2.093c1.327 0 2.362.623 3.317 1.564zm10.602 4.836h-.92c-1.407 0-2.487-.701-3.491-1.738l-.303.397c-.441.58-.915 1.201-1.439 1.818c1.356 1.324 3 2.324 5.232 2.324h.92v2.398L20 14.468l-3.988-4.025v2.851z"/>`), g.Group(children),
	)
}

func Signal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 14a1.99 1.99 0 0 0-1.981 2c0 1.104.887 2 1.981 2a1.99 1.99 0 0 0 1.98-2c0-1.105-.886-2-1.98-2zm-4.2-2.242l1.4 1.414a3.933 3.933 0 0 1 5.601 0l1.399-1.414a5.898 5.898 0 0 0-8.4 0zM3 8.928l1.4 1.414a7.864 7.864 0 0 1 11.199 0L17 8.928a9.831 9.831 0 0 0-14 0zM.199 6.1l1.4 1.414a11.797 11.797 0 0 1 16.801 0L19.8 6.1a13.763 13.763 0 0 0-19.601 0z"/>`), g.Group(children),
	)
}

func Sound(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.312 4.566C4.19 5.685-.715 12.681 3.523 16.918c4.236 4.238 11.23-.668 12.354-1.789c1.121-1.119-.335-4.395-3.252-7.312c-2.919-2.919-6.191-4.376-7.313-3.251zm9.264 9.59c-.332.328-2.895-.457-5.364-2.928c-2.467-2.469-3.256-5.033-2.924-5.363c.328-.332 2.894.457 5.36 2.926c2.471 2.467 3.258 5.033 2.928 5.365zm.858-8.174l1.904-1.906a.999.999 0 1 0-1.414-1.414L14.02 4.568a.999.999 0 1 0 1.414 1.414zM11.124 3.8a1 1 0 0 0 1.36-.388l1.087-1.926a1 1 0 0 0-1.748-.972L10.736 2.44a1 1 0 0 0 .388 1.36zm8.748 3.016a.999.999 0 0 0-1.36-.388l-1.94 1.061a1 1 0 1 0 .972 1.748l1.94-1.061a1 1 0 0 0 .388-1.36z"/>`), g.Group(children),
	)
}

func SoundMix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 1.6c0-.553-.448-.6-1-.6c-.553 0-1 .047-1 .6V10h2V1.6zM3 18.4c0 .551.447.6 1 .6c.552 0 1-.049 1-.6V15H3v3.4zM6.399 11h-4.8C1.046 11 1 11.448 1 12v1c0 .553.046 1 .599 1H6.4c.55 0 .6-.447.6-1v-1c0-.552-.05-1-.601-1zm12 1h-4.801c-.552 0-.598.448-.598 1v1c0 .553.046 1 .599 1H18.4c.55 0 .6-.447.6-1v-1c0-.552-.05-1-.601-1zM13 7c0-.552-.05-1-.601-1h-4.8C7.046 6 7 6.448 7 7v1c0 .553.046 1 .599 1H12.4c.55 0 .6-.447.6-1V7zm-2-5.4c0-.553-.448-.6-1-.6c-.553 0-1 .047-1 .6V5h2V1.6zM9 18.4c0 .551.447.6 1 .6c.552 0 1-.049 1-.6V10H9v8.4zm8-16.8c0-.553-.448-.6-1-.6c-.553 0-1 .047-1 .6V11h2V1.6zm-2 16.8c0 .551.447.6 1 .6c.552 0 1-.049 1-.6V16h-2v2.4z"/>`), g.Group(children),
	)
}

func SoundMute(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.201 9.194c1.389 1.883 1.818 3.517 1.559 3.777c-.26.258-1.893-.17-3.778-1.559l-5.526 5.527c4.186 1.838 9.627-2.018 10.605-2.996c.925-.922.097-3.309-1.856-5.754l-1.004 1.005zM8.667 7.941c-1.099-1.658-1.431-3.023-1.194-3.26c.233-.234 1.6.096 3.257 1.197l1.023-1.025C9.489 3.179 7.358 2.519 6.496 3.384c-.928.926-4.448 5.877-3.231 9.957l5.402-5.4zm9.854-6.463a.999.999 0 0 0-1.414 0L1.478 17.108a.999.999 0 1 0 1.414 1.414l15.629-15.63a.999.999 0 0 0 0-1.414z"/>`), g.Group(children),
	)
}

func SportsClub(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6 13.5l4 2.5l4-2.5V5H6v8.5zM4.5 10a2 2 0 1 0-4 0a2 2 0 0 0 4 0zm13-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4zM4.485 6.199A6.71 6.71 0 0 1 10 3.3a6.715 6.715 0 0 1 5.456 2.823a1.4 1.4 0 0 0 2.281-1.624A9.518 9.518 0 0 0 10 .5a9.506 9.506 0 0 0-7.817 4.107a1.402 1.402 0 0 0 .355 1.948a1.401 1.401 0 0 0 1.947-.356zm10.971 7.678A6.713 6.713 0 0 1 10 16.7a6.71 6.71 0 0 1-5.515-2.899a1.4 1.4 0 0 0-2.302 1.592A9.506 9.506 0 0 0 10 19.5a9.518 9.518 0 0 0 7.737-3.999a1.401 1.401 0 0 0-2.281-1.624z"/>`), g.Group(children),
	)
}

func Spreadsheet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 1H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zm-1 7H9v9H8V8H5V7h3V3h1v4h6v1z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func SquaredCross(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 2H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-2.939 12.789L10 11.729l-3.061 3.06l-1.729-1.728L8.271 10l-3.06-3.061L6.94 5.21L10 8.271l3.059-3.061l1.729 1.729L11.729 10l3.06 3.061l-1.728 1.728z"/>`), g.Group(children),
	)
}

func SquaredMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 2H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-1 9H5V9h10v2z"/>`), g.Group(children),
	)
}

func SquaredPlus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 2H4c-1.1 0-2 .9-2 2v12c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-1 9h-4v4H9v-4H5V9h4V5h2v4h4v2z"/>`), g.Group(children),
	)
}

func Star(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10 1.3l2.388 6.722H18.8l-5.232 3.948l1.871 6.928L10 14.744l-5.438 4.154l1.87-6.928l-5.233-3.948h6.412L10 1.3z"/>`), g.Group(children),
	)
}

func StarOutlined(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.8 8.022h-6.413L10 1.3L7.611 8.022H1.199l5.232 3.947l-1.871 6.929L10 14.744l5.438 4.154l-1.869-6.929L18.8 8.022zM10 12.783l-3.014 2.5l1.243-3.562l-2.851-2.3l3.522.101l1.1-4.04l1.099 4.04l3.521-.101l-2.851 2.3l1.243 3.562l-3.012-2.5z"/>`), g.Group(children),
	)
}

func Stopwatch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.376 6.745c-.447.275 1.197 4.242 1.598 4.888a1.206 1.206 0 0 0 2.053-1.266c-.397-.648-3.205-3.898-3.651-3.622zm-.335-4.343a8.98 8.98 0 0 1 5.918 0c.329.114.765-.115.572-.611c-.141-.36-.277-.712-.332-.855c-.131-.339-.6-.619-.804-.665C11.623.097 10.823 0 10 0S8.377.097 7.604.271c-.204.046-.672.326-.803.665l-.332.855c-.193.496.243.726.572.611zm12.057.784a10.132 10.132 0 0 0-1.283-1.285c-.153-.129-.603-.234-.888.051l-1.648 1.647a9.27 9.27 0 0 1 1.155.966c.362.361.677.752.966 1.155l1.647-1.647c.286-.286.181-.735.051-.887zM10 2.9A8.1 8.1 0 0 0 1.899 11A8.1 8.1 0 0 0 10 19.101A8.1 8.1 0 0 0 10 2.9zm0 14.201A6.1 6.1 0 1 1 10.001 4.9A6.1 6.1 0 0 1 10 17.1z"/>`), g.Group(children),
	)
}

func Suitcase(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 4h-1v15h1c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zM0 6v11c0 1.1.899 2 2 2h1V4H2C.899 4 0 4.9 0 6zm13.5-4.094C12.819 1.59 11.611 1 9.981 1c-1.633 0-2.8.59-3.481.906V4H4v15h12V4h-2.5V1.906zM12 4H8V2.664c.534-.23 1.078-.465 1.981-.465c.902 0 1.486.234 2.019.465V4z"/>`), g.Group(children),
	)
}

func Swap(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 5H4V3L0 6.5L4 10V8h10V5zm6 8.5L16 10v2H6v3h10v2l4-3.5z"/>`), g.Group(children),
	)
}

func Sweden(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 4H9v5h10V5a1 1 0 0 0-1-1zM1 15c0 .553.248 1 .8 1H7v-5H1v4zm8 1h9a1 1 0 0 0 1-1v-4H9v5zM1 5v4h6V4H1.8c-.552 0-.8.447-.8 1z"/>`), g.Group(children),
	)
}

func Switch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 3H7a7 7 0 1 0 0 14h6a7 7 0 1 0 0-14zm0 12a5 5 0 1 1 .001-10.001A5 5 0 0 1 13 15z"/>`), g.Group(children),
	)
}

func Tablet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 0H4C2.9 0 2 .899 2 2v16c0 1.1.9 2 2 2h12c1.101 0 2-.9 2-2V2c0-1.101-.899-2-2-2zm-6 19c-.69 0-1.25-.447-1.25-1s.56-1 1.25-1c.689 0 1.25.447 1.25 1s-.561 1-1.25 1zm6-3H4V2h12v14z"/>`), g.Group(children),
	)
}

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.662 5.521L5.237 19l.707-4.967l-4.945.709L14.424 1.263c.391-.392 1.133-.308 1.412 0l2.826 2.839c.5.473.391 1.026 0 1.419z"/>`), g.Group(children),
	)
}

func Text(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M15.5 11h-11c-.275 0-.5.225-.5.5v1a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5zm0-4h-11c-.275 0-.5.225-.5.5v1a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5zm-5 8h-6c-.275 0-.5.225-.5.5v1a.5.5 0 0 0 .5.5h6a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5zm5-12h-11c-.275 0-.5.225-.5.5v1a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func TextDocument(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 1H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zm-1 16H5V3h10v14zM13 5H7v2h6V5zm0 8H7v2h6v-2zm0-4H7v2h6V9z"/>`), g.Group(children),
	)
}

func TextDocumentInverted(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16 1H4a1 1 0 0 0-1 1v16a1 1 0 0 0 1 1h12a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1zm-3 14H7v-2h6v2zm0-4H7V9h6v2zm0-4H7V5h6v2z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func Thermometer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 10.123V1a1 1 0 0 0-1-1H7.799C7.247 0 7 .447 7 1v9.123A5.383 5.383 0 0 0 4.6 14.6a5.4 5.4 0 0 0 10.8 0a5.383 5.383 0 0 0-2.4-4.477zM10 17.9a3.3 3.3 0 0 1-3.3-3.3A3.29 3.29 0 0 1 9 11.471V4h2v7.471a3.29 3.29 0 0 1 2.3 3.129a3.3 3.3 0 0 1-3.3 3.3z"/>`), g.Group(children),
	)
}

func ThumbsDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.352 12.638c.133.356-3.539 3.634-1.397 6.291c.501.621 2.201-2.975 4.615-4.602c1.331-.899 4.43-2.811 4.43-3.868V3.617C14 2.346 9.086 1 5.352 1C3.983 1 2 9.576 2 10.939c0 1.367 4.221 1.343 4.352 1.699zM15 12.543c.658 0 3-.4 3-3.123V4.572c0-2.721-2.342-3.021-3-3.021c-.657 0 1 .572 1 2.26v6.373c0 1.768-1.657 2.359-1 2.359z"/>`), g.Group(children),
	)
}

func ThumbsUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.648 7.362c-.133-.355 3.539-3.634 1.398-6.291c-.501-.621-2.201 2.975-4.615 4.603C9.099 6.572 6 8.484 6 9.541v6.842C6 17.654 10.914 19 14.648 19C16.017 19 18 10.424 18 9.062c0-1.368-4.221-1.344-4.352-1.7zM5 7.457c-.658 0-3 .4-3 3.123v4.848c0 2.721 2.342 3.021 3 3.021c.657 0-1-.572-1-2.26V9.816c0-1.768 1.657-2.359 1-2.359z"/>`), g.Group(children),
	)
}

func ThunderCloud(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.213 6.641c-.276 0-.546.023-.809.066C13.748 4.562 11.715 3 9.309 3c-2.939 0-5.32 2.328-5.32 5.199c0 .258.02.51.057.756a3.815 3.815 0 0 0-.429-.027C1.619 8.928 0 10.512 0 12.463C0 14.416 1.619 16 3.617 16h11.596C17.856 16 20 13.904 20 11.32c0-2.586-2.144-4.679-4.787-4.679zm-3.842 4.31c-.494.703-2.614 2.889-2.704 2.98c-.104.129-.391.344-.663.166c-.079-.051-.172-.152-.172-.354c0-.193.088-.391.098-.412l1.033-2.287c-.193-.078-.527-.211-.785-.324l-.068-.029c-.262-.111-.588-.25-.588-.607c0-.172.081-.373.249-.609c.495-.705 2.614-2.889 2.705-2.982c.103-.127.39-.342.663-.166c.078.051.171.154.171.354c0 .193-.088.391-.098.414L10.178 9.38c.195.078.528.213.787.324l.068.029c.262.111.588.25.588.609c0 .172-.082.371-.25.609z"/>`), g.Group(children),
	)
}

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m4.906 11.541l3.551 3.553l6.518-6.518l-3.553-3.551l-6.516 6.516zm14.198-4.877l-1.511-1.512a2.024 2.024 0 0 1-2.747-2.746L13.335.894a1.017 1.017 0 0 0-1.432 0L.893 11.904a1.017 1.017 0 0 0 0 1.432l1.512 1.51a2.024 2.024 0 0 1 2.747 2.748l1.512 1.51a1.015 1.015 0 0 0 1.432 0L19.104 8.096a1.015 1.015 0 0 0 0-1.432zM8.457 16.719l-5.176-5.178L11.423 3.4l5.176 5.176l-8.142 8.143z"/>`), g.Group(children),
	)
}

func TimeSlot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 .4A9.6 9.6 0 0 0 .4 10a9.6 9.6 0 1 0 19.2-.001C19.6 4.698 15.301.4 10 .4zm0 17.199a7.6 7.6 0 1 1 0-15.2V10l6.792-3.396A7.548 7.548 0 0 1 17.6 10a7.6 7.6 0 0 1-7.6 7.599z"/>`), g.Group(children),
	)
}

func Tools(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.135 6.89c.933-.725 1.707-.225 2.74.971c.116.135.272-.023.361-.1c.088-.078 1.451-1.305 1.518-1.361c.066-.059.146-.169.041-.292a36.238 36.238 0 0 1-.743-.951c-1.808-2.365 4.946-3.969 3.909-3.994c-.528-.014-2.646-.039-2.963-.004c-1.283.135-2.894 1.334-3.705 1.893c-1.061.726-1.457 1.152-1.522 1.211c-.3.262-.048.867-.592 1.344c-.575.503-.934.122-1.267.414c-.165.146-.627.492-.759.607c-.133.117-.157.314-.021.471c0 0 1.264 1.396 1.37 1.52c.105.122.391.228.567.071c.177-.156.632-.553.708-.623c.078-.066-.05-.861.358-1.177zm5.708.517c-.12-.139-.269-.143-.397-.029L7.012 8.63a.289.289 0 0 0-.027.4l8.294 9.439c.194.223.53.246.751.053l.97-.813a.54.54 0 0 0 .052-.758L8.843 7.407zM19.902 3.39c-.074-.494-.33-.391-.463-.182c-.133.211-.721 1.102-.963 1.506c-.24.4-.832 1.191-1.934.41c-1.148-.811-.749-1.377-.549-1.758c.201-.383.818-1.457.907-1.59c.089-.135-.015-.527-.371-.363c-.357.164-2.523 1.025-2.823 2.26c-.307 1.256.257 2.379-.85 3.494l-1.343 1.4l1.349 1.566l1.654-1.57c.394-.396 1.236-.781 1.998-.607c1.633.369 2.524-.244 3.061-1.258c.482-.906.402-2.814.327-3.308zM2.739 17.053a.538.538 0 0 0 0 .758l.951.93c.208.209.538.121.746-.088l4.907-4.824l-1.503-1.714l-5.101 4.938z"/>`), g.Group(children),
	)
}

func TrafficCone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 12.078c2.39 0 4.392-.812 4.513-1.873l-1.125-3.152c-.264.761-1.725 1.301-3.388 1.301s-3.124-.54-3.388-1.301l-1.124 3.152c.121 1.061 2.122 1.873 4.512 1.873zm0-6.705c1.124 0 2.167-.348 2.473-.889c-.421-1.182-.782-2.197-1.011-2.836C11.31 1.221 10.621 1 10 1s-1.31.221-1.462.648L7.527 4.484c.306.541 1.35.889 2.473.889zm8.78 7.693l-3.755-1.514l.433 1.207c-.022 1.279-2.504 2.299-5.458 2.299c-2.953 0-5.437-1.019-5.458-2.299l.433-1.207l-3.755 1.514c-1.053.424-1.098 1.209-.098 1.744l7.062 3.787c.998.535 2.633.535 3.632 0l7.063-3.787c.999-.535.954-1.32-.099-1.744z"/>`), g.Group(children),
	)
}

func Trash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.389 7.113L4.49 18.021c.061.461 2.287 1.977 5.51 1.979c3.225-.002 5.451-1.518 5.511-1.979l1.102-10.908C14.929 8.055 12.412 8.5 10 8.5c-2.41 0-4.928-.445-6.611-1.387zm9.779-5.603l-.859-.951C11.977.086 11.617 0 10.916 0H9.085c-.7 0-1.061.086-1.392.559l-.859.951C4.264 1.959 2.4 3.15 2.4 4.029v.17C2.4 5.746 5.803 7 10 7c4.198 0 7.601-1.254 7.601-2.801v-.17c0-.879-1.863-2.07-4.433-2.519zM12.07 4.34L11 3H9L7.932 4.34h-1.7s1.862-2.221 2.111-2.522c.19-.23.384-.318.636-.318h2.043c.253 0 .447.088.637.318c.248.301 2.111 2.522 2.111 2.522h-1.7z"/>`), g.Group(children),
	)
}

func Tree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 10c0-1.361-.758-2.616-2.031-3.622c-.002-.001-.004-.001-.005-.003C17.602 2.803 14.177 0 10 0S2.398 2.803 2.036 6.375c-.001.002-.003.002-.005.003C.758 7.384 0 8.639 0 10c0 3.112 3.947 5.669 9 5.97V17c0 1-1.821 1.911-1.821 1.911a.227.227 0 0 0-.109.277S7.375 20 8 20s1.124-.5 2.374-.5s2.439.432 2.439.432a.342.342 0 0 0 .329-.073l.717-.717c.078-.078.058-.173-.046-.212c0 0-1.812-.68-1.812-1.93v-1.121C16.565 15.324 20 12.903 20 10zM2 10c0-1.019.768-1.945 2.022-2.651C4.012 7.233 4 7.117 4 7c0-2.762 2.687-5 6-5s6 2.238 6 5c0 .117-.012.233-.021.349C17.232 8.055 18 8.981 18 10c0 1.864-2.551 3.424-5.999 3.869v-.668a.53.53 0 0 1 .145-.337l1.833-1.726a.534.534 0 0 0 .146-.337V9.95c0-.11-.078-.155-.172-.099l-1.779 1.047c-.096.056-.173.012-.173-.099V7.2c0-.11-.085-.172-.19-.137l-2.621.874a.297.297 0 0 0-.189.263v2.6c0 .11-.079.158-.177.107L6.802 9.843a.289.289 0 0 0-.318.048l-.342.342a.185.185 0 0 0 .009.273l2.7 2.361c.083.073.15.222.15.332v.765C5.056 13.719 2 12.04 2 10z"/>`), g.Group(children),
	)
}

func TriangleDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 6h10l-5 9l-5-9z"/>`), g.Group(children),
	)
}

func TriangleLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 5v10l-9-5l9-5z"/>`), g.Group(children),
	)
}

func TriangleRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15 10l-9 5V5l9 5z"/>`), g.Group(children),
	)
}

func TriangleUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 14H5l5-9l5 9z"/>`), g.Group(children),
	)
}

func Trophy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.18 14.356c0-1.451 1.1-2.254 2.894-3.442C16.268 9.458 19 7.649 19 3.354a.703.703 0 0 0-.709-.699h-3.43C14.377 1.759 12.932.8 10 .8c-2.934 0-4.377.959-4.862 1.855H1.707A.703.703 0 0 0 1 3.354c0 4.295 2.73 6.104 4.926 7.559c1.794 1.188 2.894 1.991 2.894 3.442v1.311c-1.884.209-3.269.906-3.269 1.736c0 .994 1.992 1.799 4.449 1.799s4.449-.805 4.449-1.799c0-.83-1.385-1.527-3.269-1.736v-1.31zM13.957 9.3c.566-1.199 1.016-2.826 1.088-5.246h2.51c-.24 2.701-1.862 4.064-3.598 5.246zM10 2.026c2.732-.002 3.799 1.115 3.798 1.529c0 .418-1.066 1.533-3.798 1.535c-2.732-.001-3.799-1.116-3.799-1.534C6.2 3.142 7.268 2.024 10 2.026zM2.445 4.054h2.509c.073 2.42.521 4.047 1.089 5.246c-1.736-1.182-3.359-2.545-3.598-5.246z"/>`), g.Group(children),
	)
}

func Tv(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 1H2C.899 1 0 1.9 0 3v11c0 1.1.882 2.178 1.961 2.393l4.372.875S2.57 19 5 19h10c2.43 0-1.334-1.732-1.334-1.732l4.373-.875C19.117 16.178 20 15.1 20 14V3c0-1.1-.9-2-2-2zm0 13H2V3h16v11z"/>`), g.Group(children),
	)
}

func Typing(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 4H4c-1.101 0-2 .9-2 2v7c0 1.1.899 2 2 2h4l4 3v-3h4c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zM6 10.6a1.1 1.1 0 1 1 0-2.2a1.1 1.1 0 0 1 0 2.2zm4 0a1.1 1.1 0 1 1 0-2.2a1.1 1.1 0 0 1 0 2.2zm4 0a1.1 1.1 0 1 1 0-2.2a1.1 1.1 0 0 1 0 2.2z"/>`), g.Group(children),
	)
}

func Uninstall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19.059 10.898l-3.171-7.927A1.543 1.543 0 0 0 14.454 2H5.546c-.632 0-1.2.384-1.434.971L.941 10.898a4.25 4.25 0 0 0-.246 2.272l.59 3.539A1.544 1.544 0 0 0 2.808 18h14.383c.755 0 1.399-.546 1.523-1.291l.59-3.539a4.22 4.22 0 0 0-.245-2.272zM5.52 4.786l1.639-1.132l2.868 2.011l2.868-2.011l1.639 1.132l-2.869 2.033l2.928 2.06l-1.639 1.171l-2.927-2.076L7.1 10.05L5.461 8.879l2.928-2.06L5.52 4.786zm11.439 10.459a.902.902 0 0 1-.891.755H3.932a.902.902 0 0 1-.891-.755l-.365-2.193A.902.902 0 0 1 3.567 12h12.867c.558 0 .983.501.891 1.052l-.366 2.193z"/>`), g.Group(children),
	)
}

func Unread(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M17 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4zm-4.5 1h-11a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5zm0 5h-11a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5zm0 5h-11a.5.5 0 0 0-.5.5v1a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-1a.5.5 0 0 0-.5-.5z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func Untag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m1 14.742l4.945-.709L5.239 19l5.962-5.985l-4.069-4.429L1 14.742zm17.664-9.221c.391-.393.5-.945 0-1.419l-2.826-2.839c-.279-.308-1.021-.392-1.412 0l-3.766 3.78l4.068 4.429l3.936-3.951zm.042 9.772l-14.001-14a.999.999 0 1 0-1.414 1.414l14.001 14a.996.996 0 0 0 1.414 0a.999.999 0 0 0 0-1.414z"/>`), g.Group(children),
	)
}

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 12h4V6h3l-5-5l-5 5h3v6zm11.338 1.532c-.21-.224-1.611-1.723-2.011-2.114A1.503 1.503 0 0 0 16.285 11h-1.757l3.064 2.994h-3.544a.274.274 0 0 0-.24.133L12.992 16H7.008l-.816-1.873a.276.276 0 0 0-.24-.133H2.408L5.471 11H3.715c-.397 0-.776.159-1.042.418c-.4.392-1.801 1.891-2.011 2.114c-.489.521-.758.936-.63 1.449l.561 3.074c.128.514.691.936 1.252.936h16.312c.561 0 1.124-.422 1.252-.936l.561-3.074c.126-.513-.142-.928-.632-1.449z"/>`), g.Group(children),
	)
}

func UploadToCloud(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.213 6.639c-.276 0-.546.025-.809.068C13.748 4.562 11.716 3 9.309 3c-2.939 0-5.32 2.328-5.32 5.199c0 .256.02.508.057.756a3.567 3.567 0 0 0-.429-.027C1.619 8.928 0 10.51 0 12.463S1.619 16 3.617 16H8v-4H5.5L10 7l4.5 5H12v4h3.213C17.856 16 20 13.904 20 11.32c0-2.586-2.144-4.681-4.787-4.681z"/>`), g.Group(children),
	)
}

func User(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.725 2.146c-1.016.756-1.289 1.953-1.239 2.59c.064.779.222 1.793.222 1.793s-.313.17-.313.854c.109 1.717.683.976.801 1.729c.284 1.814.933 1.491.933 2.481c0 1.649-.68 2.42-2.803 3.334C3.196 15.845 1 17 1 19v1h18v-1c0-2-2.197-3.155-4.328-4.072c-2.123-.914-2.801-1.684-2.801-3.334c0-.99.647-.667.932-2.481c.119-.753.692-.012.803-1.729c0-.684-.314-.854-.314-.854s.158-1.014.221-1.793c.065-.817-.398-2.561-2.3-3.096c-.333-.34-.558-.881.466-1.424c-2.24-.105-2.761 1.067-3.954 1.929z"/>`), g.Group(children),
	)
}

func Users(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.989 19.129c0-2.246-2.187-3.389-4.317-4.307c-2.123-.914-2.801-1.684-2.801-3.334c0-.989.648-.667.932-2.481c.12-.752.692-.012.802-1.729c0-.684-.313-.854-.313-.854s.159-1.013.221-1.793c.064-.817-.398-2.56-2.301-3.095c-.332-.341-.557-.882.467-1.424c-2.24-.104-2.761 1.068-3.954 1.93c-1.015.756-1.289 1.953-1.24 2.59c.065.78.223 1.793.223 1.793s-.314.17-.314.854c.11 1.718.684.977.803 1.729c.284 1.814.933 1.492.933 2.481c0 1.65-.212 2.21-2.336 3.124C.663 15.53 0 17 .011 19.129C.014 19.766 0 20 0 20h16s-.011-.234-.011-.871zm2.539-5.764c-1.135-.457-1.605-1.002-1.605-2.066c0-.641.418-.432.602-1.603c.077-.484.447-.008.518-1.115c0-.441-.202-.551-.202-.551s.103-.656.143-1.159c.05-.627-.364-2.247-2.268-2.247c-1.903 0-2.318 1.62-2.269 2.247c.042.502.144 1.159.144 1.159s-.202.109-.202.551c.071 1.107.441.631.518 1.115c.184 1.172.602.963.602 1.603c0 1.064-.438 1.562-1.809 2.152c-.069.029-.12.068-.183.102c1.64.712 4.226 1.941 4.838 4.447H20v-2.318c0-1-.273-1.834-1.472-2.317z"/>`), g.Group(children),
	)
}

func VCard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 3H1a1 1 0 0 0-1 1v12a1 1 0 0 0 1 1h18a1 1 0 0 0 1-1V4a1 1 0 0 0-1-1zm-6 4h4v1h-4V7zm-2 7.803a2.31 2.31 0 0 0-.529-.303c-1.18-.508-2.961-1.26-2.961-2.176c0-.551.359-.371.518-1.379c.066-.418.385-.007.445-.961c0-.38-.174-.475-.174-.475s.088-.562.123-.996c.036-.453-.221-1.8-1.277-2.097c-.186-.188-.311-.111.258-.412c-1.244-.059-1.534.592-2.196 1.071c-.564.42-.717 1.085-.689 1.439c.037.433.125.996.125.996s-.175.094-.175.474c.061.954.38.543.445.961c.158 1.008.519.828.519 1.379c0 .916-1.781 1.668-2.961 2.176a2.503 2.503 0 0 0-.471.26V5h9v9.803zM18 11h-5v-1h5v1z"/>`), g.Group(children),
	)
}

func Video(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 5V3.799A.798.798 0 0 0 19.201 3H.801A.8.8 0 0 0 0 3.799V5h2v2H0v2h2v2H0v2h2v2H0v1.199A.8.8 0 0 0 .801 17h18.4a.8.8 0 0 0 .799-.801V15h-2v-2h2v-2h-2V9h2V7h-2V5h2zM8 13V7l5 3l-5 3z"/>`), g.Group(children),
	)
}

func Vinyl(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.999.8A9.2 9.2 0 0 0 .8 10.001a9.2 9.2 0 0 0 18.399 0A9.2 9.2 0 0 0 9.999.8zM10 13.001a3 3 0 1 1 0-6a3 3 0 0 1 0 6z"/>`), g.Group(children),
	)
}

func Voicemail(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.4 5.801a4.6 4.6 0 0 0-3.795 7.2H8.394A4.6 4.6 0 1 0 4.6 15h10.8a4.6 4.6 0 0 0 0-9.199zM2 10.4a2.6 2.6 0 1 1 5.2 0a2.6 2.6 0 0 1-5.2 0zM15.4 13a2.6 2.6 0 1 1-.002-5.2A2.6 2.6 0 0 1 15.4 13z"/>`), g.Group(children),
	)
}

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 6H3.5v-.5l11-.88v.88H16V4c0-1.1-.891-1.872-1.979-1.717L3.98 3.717C2.891 3.873 2 4.9 2 6v10a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2zm-1.5 7.006a1.5 1.5 0 1 1 .001-3.001a1.5 1.5 0 0 1-.001 3.001z"/>`), g.Group(children),
	)
}

func Warning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.511 17.98L10.604 1.348a.697.697 0 0 0-1.208 0L.49 17.98a.675.675 0 0 0 .005.68c.125.211.352.34.598.34h17.814a.694.694 0 0 0 .598-.34a.677.677 0 0 0 .006-.68zM11 17H9v-2h2v2zm0-3.5H9V7h2v6.5z"/>`), g.Group(children),
	)
}

func Water(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.882 9.093c-.511 4.115-3.121 4.847-3.121 7.708C6.761 18.567 8.244 20 10 20c1.756 0 3.238-1.434 3.238-3.199c0-2.861-2.61-3.593-3.121-7.708c-.016-.123-.219-.123-.235 0zm-5.999-9C3.372 4.208.762 4.939.762 7.801C.762 9.566 2.244 11 4 11s3.238-1.434 3.238-3.199c0-2.861-2.61-3.593-3.121-7.708c-.015-.123-.219-.123-.234 0zm12 0c-.511 4.115-3.121 4.847-3.121 7.708C12.762 9.566 14.244 11 16 11s3.238-1.434 3.238-3.199c0-2.861-2.61-3.593-3.121-7.708c-.016-.123-.219-.123-.234 0z"/>`), g.Group(children),
	)
}
