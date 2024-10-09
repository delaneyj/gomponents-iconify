package tdesign

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
		g.Raw(`<path fill="currentColor" d="M12 2a1.25 1.25 0 1 1 0 2.5A1.25 1.25 0 0 1 12 2M8.75 3.25a3.25 3.25 0 1 0 6.5 0a3.25 3.25 0 0 0-6.5 0m-6.5 6.375A2.625 2.625 0 0 1 4.875 7h14.25a2.625 2.625 0 0 1 0 5.25h-2.628l1.823 8.208a2.5 2.5 0 0 1-2.44 3.042h-2.142L12 17.198L10.261 23.5H8.12a2.5 2.5 0 0 1-2.44-3.042l1.823-8.208H4.875A2.625 2.625 0 0 1 2.25 9.625M4.875 9a.625.625 0 1 0 0 1.25h5.122L7.632 20.892a.5.5 0 0 0 .488.608h.619l2-7.25h2.522l2 7.25h.62a.5.5 0 0 0 .487-.608L14.003 10.25h5.122a.625.625 0 1 0 0-1.25z"/>`), g.Group(children),
	)
}

func AccessibilityFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4.5A2.25 2.25 0 1 0 12 0a2.25 2.25 0 0 0 0 4.5M21.75 7V5.5H2.25V7l6.5 1.75L9 12L4.75 22.75l1.5.75l5.745-9h.01l5.745 9l1.5-.75L15 12l.25-3.25z"/>`), g.Group(children),
	)
}

func Activity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 2H2v20h20zm-2 2v7h-2.554l-2.021 3.233l-5.865-7.82L5.546 11H4V4zM4 13h2.454L9.44 9.587l6.135 8.18L18.555 13H20v7H4z"/>`), g.Group(children),
	)
}

func Add(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 4v7h7v2h-7v7h-2v-7H4v-2h7V4z"/>`), g.Group(children),
	)
}

func AddAndSubtract(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v8h8v2h-8v8h-2v-8H3V9h8V1zM3 20h18v2H3z"/>`), g.Group(children),
	)
}

func AddCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m12-5.5V11h4.5v2H13v4.5h-2V13H6.5v-2H11V6.5z"/>`), g.Group(children),
	)
}

func AddRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm9 2.5V11h4.5v2H13v4.5h-2V13H6.5v-2H11V6.5z"/>`), g.Group(children),
	)
}

func AddressBook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 6a4 4 0 0 1 4-4h14v20H7a4 4 0 0 1-4-4zm2 8.535A4 4 0 0 1 7 14h12V4h-2v6.766l-3.5-2.1l-3.5 2.1V4H7a2 2 0 0 0-2 2zM19 16H7a2 2 0 1 0 0 4h12zM15 4h-3v3.234l1.5-.9l1.5.9z"/>`), g.Group(children),
	)
}

func Adjustment(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2M5.17 4a3.001 3.001 0 0 1 5.66 0H22v2H10.83a3.001 3.001 0 0 1-5.66 0H2V4zm8 7a3.001 3.001 0 0 1 5.66 0H22v2h-3.17a3.001 3.001 0 0 1-5.66 0H2v-2zM16 11a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-8 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-2.83 0a3.001 3.001 0 0 1 5.66 0H22v2H10.83a3.001 3.001 0 0 1-5.66 0H2v-2z"/>`), g.Group(children),
	)
}

func AirplayWave(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v17h-5.5v-2H21V5H3l.001 13h3.5v2H1zm15.95 10.383a7 7 0 0 0-9.9 0l-.706.707l-1.414-1.414l.707-.707a9 9 0 0 1 12.728 0l.707.707l-1.415 1.414zm-2.828 2.828a3 3 0 0 0-4.243 0l-.707.707l-1.414-1.414l.707-.707a5 5 0 0 1 7.07 0l.708.707l-1.414 1.414zM12 18.086L15.914 22H8.086z"/>`), g.Group(children),
	)
}

func Alarm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.914 2.5L1.5 7.914L.086 6.5L5.5 1.086zM18.5 1.086L23.914 6.5L22.5 7.914L17.086 2.5zM12 5a8 8 0 1 0 0 16a8 8 0 0 0 0-16M2 13C2 7.477 6.477 3 12 3s10 4.477 10 10s-4.477 10-10 10S2 18.523 2 13m11-5.5v5.086L16.414 16L15 17.414l-4-4V7.5z"/>`), g.Group(children),
	)
}

func AlarmAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.914 2.5L1.5 7.914L.086 6.5L5.5 1.086zM18.5 1.086L23.914 6.5L22.5 7.914L17.086 2.5zM12 5a8 8 0 1 0 0 16a8 8 0 0 0 0-16M2 13C2 7.477 6.477 3 12 3s10 4.477 10 10s-4.477 10-10 10S2 18.523 2 13m11-5v4h4v2h-4v4h-2v-4H7v-2h4V8z"/>`), g.Group(children),
	)
}

func AlarmOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m.086 6.5l2.5-2.5l-2-2L2 .586l2 2l1.5-1.5L6.914 2.5L5.414 4l18 18L22 23.414l-3.139-3.139A9.97 9.97 0 0 1 12.001 23C6.476 23 2 18.523 2 13a9.97 9.97 0 0 1 2.724-6.861L4 5.414l-2.5 2.5zM6.14 7.554A8 8 0 0 0 17.446 18.86zm2.421-4l.97-.246A10 10 0 0 1 12 3c5.523 0 10 4.477 10 10c0 .851-.107 1.679-.308 2.47l-.246.969l-1.938-.493l.246-.969a8 8 0 0 0-9.731-9.731l-.97.246zm9.94-2.468L23.913 6.5L22.5 7.914L17.086 2.5z"/>`), g.Group(children),
	)
}

func AlignTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 3h18v2H3zm9 2.586l-.707.707l-4 4l-.707.707L8 12.414l.707-.707L11 9.414V21h2V9.414l2.293 2.293l.707.707L17.414 11l-.707-.707l-4-4z"/>`), g.Group(children),
	)
}

func AlignVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 6.586V0h-2v6.586L9.707 5.293L9 4.586L7.586 6l.707.707l3 3l.707.707l.707-.707l3-3L16.414 6L15 4.586l-.707.707zM21 13H3v-2h18zm-9 .586l.707.707l3 3l.707.707L15 19.414l-.707-.707L13 17.414V24h-2v-6.586l-1.293 1.293l-.707.707L7.586 18l.707-.707l3-3z"/>`), g.Group(children),
	)
}

func Alpha(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 2v3.866l5.336-3.24l1.038 1.71L14 8.206v8.588l6.374 3.87l-1.038 1.71L14 19.134V22h-2v-4.08L3.073 12.5L12 7.08V2zm-2 7.42L6.927 12.5L12 15.58z"/>`), g.Group(children),
	)
}

func Analytics(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm14 4v3h-2V8zm-5 2v3h-2v-3zm-5 2v6H6v-6zm10 1v5h-2v-5zm-5 2v3h-2v-3z"/>`), g.Group(children),
	)
}

func Anchor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4M8 5a4 4 0 1 1 5 3.874V10.5h4.5v2H13v8.458c3.133-.267 5.643-1.796 6.802-4.228l-1.23-1.23l4.048-4.048l-.135 2.6C22.19 19.74 17.455 23 12 23S1.81 19.74 1.515 14.052l-.135-2.6L5.427 15.5l-1.23 1.23c1.159 2.432 3.67 3.96 6.802 4.228V12.5H6.5v-2H11V8.874A4 4 0 0 1 8 5"/>`), g.Group(children),
	)
}

func Angry(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5.769-3.866l3.464 2l-1 1.732l-3.464-2zm11.464 1.732l-3.464 2l-1-1.732l3.464-2zM7.67 15.499A5 5 0 0 1 12 13a5 5 0 0 1 4.33 2.5l.501.865l-1.731 1.001l-.5-.865A3 3 0 0 0 12 15a3 3 0 0 0-2.6 1.5l-.5.866l-1.731-1.001z"/>`), g.Group(children),
	)
}

func Animation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.09 7.054l.11-.46q.027-.091.057-.181l.371-.841a5 5 0 1 1 6.491 6.958l-.532.213l-.117.037l-.511.128Q16.493 13 16 13a5 5 0 0 1-4.91-5.946M9.47 5.471a7.02 7.02 0 0 0-4 4A7.002 7.002 0 0 0 8 23a7 7 0 0 0 6.529-4.47q.198-.078.39-.166l.338-.135l-.01-.026a7.03 7.03 0 0 0 3.282-3.674q.198-.076.392-.165l.336-.135l-.01-.026A7 7 0 1 0 9.47 5.471M16 15a5 5 0 0 1-1.882 1.53l-.53.212l-.12.04l-.507.126q-.467.091-.961.092a5 5 0 0 1-4.91-5.946l.11-.46q.027-.091.057-.181l.372-.842A5 5 0 0 1 9 7.999V8a7 7 0 0 0 7 7M5 12a7 7 0 0 0 7 7a5 5 0 1 1-7-7"/>`), g.Group(children),
	)
}

func AnimationOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.41.06l3.716 6.174l7.02 1.626l-4.724 5.44l.623 7.18l-6.635-2.812l-6.634 2.811l.623-7.178L.676 7.86l7.02-1.626zm0 3.88L8.972 7.99L4.365 9.058l3.1 3.572l-.41 4.711l4.355-1.845l4.355 1.845l-.409-4.711l3.1-3.572l-4.607-1.067zm9.453 10.071l2.475 2.475l-1.414 1.414l-2.475-2.474zm-8.296 6.116l2.474 2.475l-1.414 1.414l-2.475-2.475zm6.578 0l2.474 2.475l-1.414 1.414l-2.475-2.475z"/>`), g.Group(children),
	)
}

func Anticlockwise(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.333 5a3 3 0 0 1 3-3H13v2H7.333a1 1 0 0 0-1 1v6.5H4.34L.586 7.446l1.467-1.36l2.28 2.462zM8 6h15v15H8zm2 2v11h11V8z"/>`), g.Group(children),
	)
}

func Api(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M.586 12L5 7.586L6.414 9l-3 3l3 3L5 16.414zm7-7L12 .586L16.414 5L15 6.414l-3-3l-3 3zM9 17.586l3 3l3-3L16.414 19L12 23.414L7.586 19zm1.998-4.584v-2.004h2.004v2.004zM17.586 15l3-3l-3-3L19 7.586L23.414 12L19 16.414z"/>`), g.Group(children),
	)
}

func App(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h9v9H2zm2 2v5h5V4zm13.5 0a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5M13 6.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0M2 13h9v9H2zm2 2v5h5v-5zm9-2h9v9h-9zm2 2v5h5v-5z"/>`), g.Group(children),
	)
}

func Apple(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 2v2.404l.152-.025l.275-.05c.329-.059.698-.126 1.065-.165c1.025-.111 2.153-.03 3.492.711C20.622 6.335 22 9.1 22 12.5c0 4.28-2.71 8.3-6.189 9.439c-1.04.34-1.785.156-2.404-.072l-.272-.102c-.416-.159-.695-.265-1.135-.265s-.718.106-1.134.265l-.272.103c-.62.228-1.366.413-2.407.07C4.72 20.797 2 16.782 2 12.5c0-3.4 1.378-6.16 4.015-7.624c1.34-.744 2.468-.825 3.494-.714c.366.04.736.107 1.064.167l.275.05l.152.025V2zm-2 4.427q-.258-.036-.494-.078l-.34-.061a12 12 0 0 0-.873-.137c-.7-.076-1.397-.032-2.308.473C5.12 7.66 4 9.657 4 12.5c0 3.557 2.288 6.708 4.813 7.538c.458.151.711.092 1.09-.047l.19-.073c.429-.168 1.068-.418 1.907-.418s1.478.25 1.906.417l.19.073c.38.14.634.198 1.093.048C17.718 19.21 20 16.058 20 12.5c0-2.844-1.12-4.843-2.985-5.875c-.91-.505-1.608-.549-2.308-.473c-.29.031-.562.08-.874.137q-.161.03-.34.06q-.235.041-.493.078v1.402q.261-.093.5-.23l.865-.5l1.001 1.73l-.866.501A5 5 0 0 1 12 10c-.91 0-1.764-.244-2.5-.67l-.866-.5l1.001-1.731l.866.5q.237.137.499.23z"/>`), g.Group(children),
	)
}

func Application(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .845l9.66 5.578v11.154L12 23.155l-9.66-5.578V6.423zm0 2.31L4.34 7.577v8.846L12 20.845l7.66-4.422V7.577zM8.723 8.613L12 10.798l3.277-2.185l1.11 1.664L13 12.535V16h-2v-3.465l-3.387-2.258z"/>`), g.Group(children),
	)
}

func ArchitectureHuiStyle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.8 1.4L12 .333L11.2 1.4l-.004.005l-.012.016l-.049.065l-.19.249a77 77 0 0 1-2.948 3.607c-.865.988-1.77 1.952-2.568 2.66c-.4.356-.75.627-1.035.803c-.246.151-.37.185-.396.193v-.002l-.004-1l-2 .008l.004 1c.001.405.106.908.459 1.327c.378.45.928.67 1.543.67v9H2v2h20v-2h-2v-9c.614 0 1.165-.22 1.543-.67c.353-.42.457-.924.457-1.33V8h-2v.996a1.6 1.6 0 0 1-.394-.192a8 8 0 0 1-1.036-.802c-.797-.71-1.703-1.673-2.568-2.661a73 73 0 0 1-3.138-3.856l-.049-.065l-.012-.016zM16.704 9H7.295a38 38 0 0 0 2.207-2.34A75 75 0 0 0 12 3.637l.184.234a75 75 0 0 0 2.313 2.788A38 38 0 0 0 16.704 9M6 11h12v9h-3v-3a3 3 0 0 0-6 0v3H6zm7 9h-2v-3a1 1 0 1 1 2 0z"/>`), g.Group(children),
	)
}

func Archway(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 2v1h12V2h2v1h1v2h-.78l.6 3H22v2h-1v10h1v2H2v-2h1V10H2V8h1.18l.6-3H3V3h1V2zm-.18 3l-.6 3h13.56l-.6-3zM19 10h-3v10h3zm-5 10V10h-4v10zm-6 0V10H5v10z"/>`), g.Group(children),
	)
}

func ArchwayOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v5h-2.915l.385 3H22v2h-2.274l1.393 10.865l-1.984.254L17.71 12H6.258l-1.39 11.116l-1.984-.248L4.242 12H2v-2h2.492l.375-3H2zm4.883 5l-.375 3H11V7zM13 7v3h4.454l-.385-3zM4 4v1h16V4z"/>`), g.Group(children),
	)
}

func ArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 4.5v11.586l-4.5-4.5L5.086 13L12 19.914L18.914 13L17.5 11.586l-4.5 4.5V4.5z"/>`), g.Group(children),
	)
}

func ArrowDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m12-5.5v7.586l3-3l1.414 1.414L12 17.914L6.586 12.5L8 11.086l3 3V6.5z"/>`), g.Group(children),
	)
}

func ArrowDownRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm9 2.5v7.586l3-3l1.414 1.414L12 17.914L6.586 12.5L8 11.086l3 3V6.5z"/>`), g.Group(children),
	)
}

func ArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.5 13H7.914l4.5 4.5L11 18.914L4.086 12L11 5.086L12.414 6.5l-4.5 4.5H19.5z"/>`), g.Group(children),
	)
}

func ArrowLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21 12a9 9 0 1 0-18 0a9 9 0 0 0 18 0M12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1m5.5 12H9.914l3 3l-1.414 1.414L6.086 12L11.5 6.586L12.914 8l-3 3H17.5z"/>`), g.Group(children),
	)
}

func ArrowLeftDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m18.01 7.404l-8.19 8.192h6.364v2h-9.78V7.818h2v6.364l8.193-8.192z"/>`), g.Group(children),
	)
}

func ArrowLeftDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0m9 11C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11m-4.182-6.819V8.524h2v4.243l5.364-5.364l1.414 1.414l-5.364 5.364h4.243v2z"/>`), g.Group(children),
	)
}

func ArrowLeftRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 12a9 9 0 1 0 18 0a9 9 0 0 0-18 0m9 11C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11m-6.914-8L9 11.086l1.414 1.414l-1.5 1.5H14v2H8.914l1.5 1.5L9 18.914zM10 8h5.086l-1.5-1.5L15 5.086L18.914 9L15 12.914L13.586 11.5l1.5-1.5H10z"/>`), g.Group(children),
	)
}

func ArrowLeftRightOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.586 16.5L8.5 11.586L9.914 13l-2.5 2.5H19.5v2H7.414l2.5 2.5L8.5 21.414zm.914-10h12.086l-2.5-2.5L15.5 2.586L20.414 7.5L15.5 12.414L14.086 11l2.5-2.5H4.5z"/>`), g.Group(children),
	)
}

func ArrowLeftRightThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.914 7.5L15.5 12.914L14.086 11.5l3-3H8.5v-2h8.586l-3-3L15.5 2.086zm-5.414 10H6.914l3 3L8.5 21.914L3.086 16.5L8.5 11.086L9.914 12.5l-3 3H15.5z"/>`), g.Group(children),
	)
}

func ArrowLeftRightTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21.5 17.5h-8.586l3 3l-1.414 1.414L9.086 16.5l5.414-5.414l1.414 1.414l-3 3H21.5zm-6.586-10L9.5 12.914L8.086 11.5l3-3H2.5v-2h8.586l-3-3L9.5 2.086z"/>`), g.Group(children),
	)
}

func ArrowLeftUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.596 18.01L8.403 9.818v6.364h-2V6.404h9.779v2H9.817l8.192 8.192z"/>`), g.Group(children),
	)
}

func ArrowLeftUpCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m6.818-4.182h7.657v2h-4.243l5.364 5.364l-1.414 1.414l-5.364-5.364v4.243h-2z"/>`), g.Group(children),
	)
}

func ArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.5 11h11.586l-4.5-4.5L13 5.086L19.914 12L13 18.914L11.586 17.5l4.5-4.5H4.5z"/>`), g.Group(children),
	)
}

func ArrowRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 12a9 9 0 1 1 18 0a9 9 0 0 1-18 0m9-11C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1M6.5 13h7.586l-3 3l1.414 1.414L17.914 12L12.5 6.586L11.086 8l3 3H6.5z"/>`), g.Group(children),
	)
}

func ArrowRightDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7.403 5.99l8.193 8.192V7.818h2v9.778H7.818v-2h6.364L5.989 7.404z"/>`), g.Group(children),
	)
}

func ArrowRightDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-8.818.768V8.525h2v7.657H8.525v-2h4.243L7.404 8.818l1.414-1.414z"/>`), g.Group(children),
	)
}

func ArrowRightUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m5.99 16.596l8.192-8.192H7.818v-2h9.778v9.778h-2V9.818L7.403 18.01z"/>`), g.Group(children),
	)
}

func ArrowRightUpCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21 12a9 9 0 1 0-18 0a9 9 0 0 0 18 0M12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1m.768 8.819H8.525v-2h7.657v7.657h-2v-4.243l-5.364 5.364l-1.414-1.414z"/>`), g.Group(children),
	)
}

func ArrowTriangleDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 12V2h6v10h4.5L12 22L4.5 12zm-.5 2l3.5 4.667L15.5 14H13V4h-2v10z"/>`), g.Group(children),
	)
}

func ArrowTriangleDownFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 12h4.5L12 22L4.5 12H9V2h6z"/>`), g.Group(children),
	)
}

func ArrowTriangleUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 12v10h6V12h4.5L12 2L4.5 12zm-.5-2L12 5.333L15.5 10H13v10h-2V10z"/>`), g.Group(children),
	)
}

func ArrowTriangleUpFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 12h4.5L12 2L4.5 12H9v10h6z"/>`), g.Group(children),
	)
}

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 19.5V7.914l-4.5 4.5L5.086 11L12 4.086L18.914 11L17.5 12.414l-4.5-4.5V19.5z"/>`), g.Group(children),
	)
}

func ArrowUpCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m11-5.914l5.414 5.414L16 12.914l-3-3V17.5h-2V9.914l-3 3L6.586 11.5z"/>`), g.Group(children),
	)
}

func ArrowUpDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-8 6.914L11.086 15l1.414-1.414l1.5 1.5V10h2v5.086l1.5-1.5L18.914 15zM8 14V8.914l-1.5 1.5L5.086 9L9 5.086L12.914 9L11.5 10.414l-1.5-1.5V14z"/>`), g.Group(children),
	)
}

func ArrowUpDownOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.5 20.414L11.586 15.5L13 14.086l2.5 2.5V4.5h2v12.086l2.5-2.5l1.414 1.414zm-10-.914V7.414L4 9.914L2.586 8.5L7.5 3.586L12.414 8.5L11 9.914l-2.5-2.5V19.5z"/>`), g.Group(children),
	)
}

func ArrowUpDownThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 3.086L12.914 8.5L11.5 9.914l-3-3V15.5h-2V6.914l-3 3L2.086 8.5zm10 5.414v8.586l3-3l1.414 1.414l-5.414 5.414l-5.414-5.414l1.414-1.414l3 3V8.5z"/>`), g.Group(children),
	)
}

func ArrowUpDownTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.5 2.5v8.586l3-3L21.914 9.5L16.5 14.914L11.086 9.5L12.5 8.086l3 3V2.5zm-10 6.586l5.414 5.414l-1.414 1.414l-3-3V21.5h-2v-8.586l-3 3L2.086 14.5z"/>`), g.Group(children),
	)
}

func Artboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 2v4h8V2h2v4h4v2h-4v8h4v2h-4v4h-2v-4H8v4H6v-4H2v-2h4V8H2V6h4V2zm0 6v8h8V8z"/>`), g.Group(children),
	)
}

func Article(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.5 2h19v20h-19zm2 2v16h15V4zM7 7h10v2H7zm0 4h10v2H7zm0 4h7v2H7z"/>`), g.Group(children),
	)
}

func Assignment(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2.5a1.5 1.5 0 0 0-1.376.9l-.262.6H4.5v16h15V4h-5.862l-.262-.6A1.5 1.5 0 0 0 12 2.5M9.128 2A3.5 3.5 0 0 1 12 .5c1.19 0 2.24.594 2.872 1.5H21.5v20h-19V2zM7 7h10v2H7zm0 4h10v2H7zm0 4h7v2H7z"/>`), g.Group(children),
	)
}

func AssignmentChecked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2.5a1.5 1.5 0 0 0-1.376.9l-.262.6H4.5v16h15V4h-5.862l-.262-.6A1.5 1.5 0 0 0 12 2.5M9.128 2A3.5 3.5 0 0 1 12 .5c1.19 0 2.24.594 2.872 1.5H21.5v20h-19V2zm8.53 7.586l-7.072 7.07l-4.242-4.242L7.758 11l2.828 2.829l5.657-5.657z"/>`), g.Group(children),
	)
}

func AssignmentCode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.5 2h6.628A3.5 3.5 0 0 1 12 .5c1.19 0 2.24.594 2.872 1.5H21.5v20h-19zm2 18h15V4h-5.862l-.262-.6a1.5 1.5 0 0 0-2.752 0l-.262.6H4.5zm1.086-8L9.5 8.086L10.914 9.5l-2.5 2.5l2.5 2.5L9.5 15.914zM14.5 8.086L18.414 12L14.5 15.914L13.086 14.5l2.5-2.5l-2.5-2.5z"/>`), g.Group(children),
	)
}

func AssignmentError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2.5a1.5 1.5 0 0 0-1.376.9l-.262.6H4.5v16h15V4h-5.862l-.262-.6A1.5 1.5 0 0 0 12 2.5M9.128 2A3.5 3.5 0 0 1 12 .5c1.19 0 2.24.594 2.872 1.5H21.5v20h-19V2zM13 7.5v6h-2v-6zm-2 7.496h2.004V17H11z"/>`), g.Group(children),
	)
}

func AssignmentUser(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2.5a1.5 1.5 0 0 0-1.376.9l-.262.6H4.5v16h15V4h-5.862l-.262-.6A1.5 1.5 0 0 0 12 2.5M9.128 2A3.5 3.5 0 0 1 12 .5c1.19 0 2.24.594 2.872 1.5H21.5v20h-19V2zM12 8a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M8.5 9.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0M6 18a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v1h-2v-1a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v1H6z"/>`), g.Group(children),
	)
}

func Attach(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.314 3.121a5 5 0 1 1 7.07 7.071l-7.777 7.778a3 3 0 1 1-4.243-4.242l7.778-7.778l1.414 1.414l-7.778 7.778a1 1 0 1 0 1.414 1.414l7.779-7.778a3 3 0 1 0-4.243-4.243L4.95 12.314a5 5 0 0 0 7.07 7.07l8.486-8.485l1.414 1.415l-8.485 8.485a7 7 0 0 1-9.9-9.9z"/>`), g.Group(children),
	)
}

func Attic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.81 1.414L12 .294l-.81 1.12l-.002.003l-.009.011l-.033.046l-.133.18a55 55 0 0 1-2.214 2.774c-.687.803-1.45 1.64-2.206 2.364c-.767.735-1.48 1.306-2.067 1.622a7.5 7.5 0 0 1-1.002.459C3.231 8.977 3.062 9 3 9H2v2h1c.328 0 .671-.074 1-.178V20H2v2h20v-2h-2v-9.178c.328.104.672.178 1 .178h1V9h-1c-.062 0-.23-.023-.524-.127a7.5 7.5 0 0 1-1.002-.46c-.587-.315-1.3-.886-2.067-1.62a33 33 0 0 1-2.206-2.365a55 55 0 0 1-2.347-2.954l-.033-.046l-.008-.011zM16.867 9H7.133c.285-.242.568-.5.843-.763a35 35 0 0 0 2.343-2.509A56 56 0 0 0 12 3.66a57 57 0 0 0 1.681 2.068a35 35 0 0 0 2.343 2.509q.414.397.843.763M6 11h12v9h-3v-3a3 3 0 0 0-6 0v3H6zm5 9v-3a1 1 0 1 1 2 0v3z"/>`), g.Group(children),
	)
}

func AtticOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.169 1.444L11.999.2l.832 1.243l.005.008l.02.03l.087.125a28 28 0 0 0 1.6 2.04c1.071 1.239 2.484 2.638 3.954 3.486c.466.269.956.367 1.501.367h1v2h-1c-.326 0-.66-.025-1-.087V17h2v2h-1v3h-2v-3h-5v3h-2v-3h-5v3h-2v-3h-1v-2H5l-.002-7.586a5.6 5.6 0 0 1-1 .086h-1v-2h1c.545 0 1.035-.098 1.501-.367c1.471-.847 2.885-2.247 3.956-3.486a28 28 0 0 0 1.601-2.04l.087-.125l.02-.03zM6.999 9L7 17h2v-2.2a3 3 0 0 1 6 0V17h1.998V9zm8.007-2a24 24 0 0 1-1.976-2.045c-.4-.462-.747-.895-1.03-1.261a31 31 0 0 1-1.032 1.262A24 24 0 0 1 8.99 7zM13 17v-2.2a1 1 0 1 0-2 0V17z"/>`), g.Group(children),
	)
}

func Audio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h18v22H3zm2 2v18h14V3zm5.996 1.996H13V7h-2.004zM12 11a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0"/>`), g.Group(children),
	)
}

func Awkward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h4v2H6zm8 0h4v2h-4zm2 3v2.667h-2V12zm3 0v3h-2v-3z"/>`), g.Group(children),
	)
}

func Backtop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 4h16v2H4zm8 3.586l6.914 6.914l-1.414 1.414l-4.5-4.5V21h-2v-9.586l-4.5 4.5L5.086 14.5z"/>`), g.Group(children),
	)
}

func BacktopRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm13.5 4h-11V6h11zm-5.5.808l4.596 4.596l-1.414 1.414L13 12.636V18.5h-2v-5.864l-2.182 2.182l-1.414-1.414z"/>`), g.Group(children),
	)
}

func Backup(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4a5.5 5.5 0 0 0-5.49 5.15l-.048.783l-.77.14A4.502 4.502 0 0 0 6.5 19h11a4.5 4.5 0 0 0 .809-8.928l-.771-.14l-.049-.781A5.5 5.5 0 0 0 12 4M4.598 8.283a7.502 7.502 0 0 1 14.804 0A6.502 6.502 0 0 1 17.5 21h-11A6.5 6.5 0 0 1 4.598 8.283M12 7.086l4.414 4.414L15 12.914l-2-2V17h-2v-6.086l-2 2L7.586 11.5z"/>`), g.Group(children),
	)
}

func Backward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.25 4.336v7l7-7v15.328l-7-7v7L3.586 12zM6.414 12l2.836 2.836V9.164zm7 0l2.836 2.836V9.164z"/>`), g.Group(children),
	)
}

func BadLaugh(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5.769-3.866l3.464 2l-1 1.732l-3.464-2zm11.464 1.732l-3.464 2l-1-1.732l3.464-2zM7 13h10v1a5 5 0 0 1-10 0zm2.17 2a3.001 3.001 0 0 0 5.66 0z"/>`), g.Group(children),
	)
}

func BambooShoot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19.843 4.228l-.211.037a21 21 0 0 0-3.178.866l.53 2.621c.09.45.173.854.238 1.173l1.167.117c.74-1.766 1.136-3.216 1.347-4.244q.064-.312.107-.57M17.5 10.963l-1.24-.124l-7.036-1.13l1.796 10.362c3.124-3.37 5.157-6.49 6.48-9.108M9.029 20.29L7.291 10.265a47 47 0 0 0-4.19 3.64a18 18 0 0 0 3.028 4.037a18.4 18.4 0 0 0 2.9 2.348m1.666-12.37l4.427.71l-.097-.48l-.46-2.268a32.5 32.5 0 0 0-3.87 2.037m11.3-5.844l-.018 1.087l-.005.116a9 9 0 0 1-.036.414c-.037.355-.107.864-.24 1.508c-.265 1.287-.777 3.111-1.767 5.318c-1.532 3.412-4.202 7.726-8.85 12.373l-.478.479l-.623-.267c-.933-.4-3.187-1.672-5.263-3.748C2.64 17.28 1.368 15.026.968 14.092l-.268-.62l.479-.479C3.352 10.82 5.452 9.08 7.422 7.685c2.839-2.01 5.403-3.298 7.52-4.124c1.81-.706 3.29-1.072 4.33-1.263c.52-.095.93-.145 1.216-.173a8 8 0 0 1 .43-.03z"/>`), g.Group(children),
	)
}

func Banana(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.487.613v7.15c-.304.593-.731 1.264-1.504 1.847c-.862.648-2.243 1.253-4.556 1.392c-2.553.154-4.66-1.15-5.241-1.717c-.591-.577-1.42-.503-1.96-.064c-.51.415-.739 1.082-.739 1.78c0 .727.071 1.403.264 2.023a1.6 1.6 0 0 0-.818-.102a1.52 1.52 0 0 0-.916.492c-.41.452-.53 1.072-.53 1.586c0 4.238 5.048 9.767 13.323 6.947l.024-.008l.023-.01c4.425-1.759 6.656-4.465 7.589-7.23c.887-2.633.562-5.2.041-6.856V2.279zM8.755 14.625c-1.475-.467-2.242-.918-2.666-1.409c-.359-.417-.564-.975-.597-1.931c1.236.861 3.458 1.87 6.055 1.713c1.607-.096 2.898-.397 3.94-.83c-.619.884-1.27 1.498-1.928 1.913c-1.41.89-3.023.97-4.804.544M20.487 3.721v4.442l.052.153c.45 1.353.754 3.542.011 5.743c-.725 2.15-2.489 4.438-6.409 6.003c-6.968 2.361-10.556-2.177-10.652-4.948c1.404.644 2.984 1.021 4.31 1.339l.453.108c2.11.512 4.343.494 6.375-.79c2.009-1.267 3.636-3.657 4.816-7.476l.044-.144V3.388z"/>`), g.Group(children),
	)
}

func Barbecue(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.316 2.796a2.717 2.717 0 0 1 3.842 0l1.316 1.316l1.619-1.619l1.414 1.414l-1.619 1.619l1.316 1.316a2.717 2.717 0 0 1 0 3.842l-1.214 1.214c-.46.46-1.045.72-1.645.782a2.7 2.7 0 0 1-.782 1.646l-1.214 1.213c-.46.46-1.045.722-1.646.782a2.7 2.7 0 0 1-.782 1.646l-1.214 1.214a2.717 2.717 0 0 1-3.842 0L7.55 17.865l-1.618 1.619l-.708-.708a.717.717 0 1 0-.267 1.183l.943-.332l.665 1.887l-.944.332a2.717 2.717 0 1 1 .176-5.056l.339-.34l-1.316-1.315a2.717 2.717 0 0 1 0-3.842l1.214-1.214a2.7 2.7 0 0 1 1.646-.782c.06-.6.321-1.185.781-1.646l1.214-1.214a2.7 2.7 0 0 1 1.646-.781a2.7 2.7 0 0 1 .782-1.646zm.2 3.641l4.047 4.047c.28.28.733.28 1.013 0L19.79 9.27a.717.717 0 0 0 0-1.014L18.474 6.94l-.974.974L16.086 6.5l.974-.974l-1.316-1.316a.717.717 0 0 0-1.014 0l-1.214 1.214a.717.717 0 0 0 0 1.013m2.632 5.46l-4.046-4.045a.717.717 0 0 0-1.013 0L9.875 9.065a.717.717 0 0 0 0 1.014l4.046 4.046c.28.28.734.28 1.014 0l1.213-1.214a.717.717 0 0 0 0-1.013m-3.641 3.643L8.46 11.493a.717.717 0 0 0-1.014 0l-1.214 1.214a.717.717 0 0 0 0 1.014l4.046 4.046c.28.28.734.28 1.014 0l1.214-1.214a.717.717 0 0 0 0-1.014"/>`), g.Group(children),
	)
}

func Barcode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h3v19H2zm6.75 0v19h-2V2zm1.75 0h3v19h-3zm6.75 0v19h-2V2zM19 2h3v19h-3z"/>`), g.Group(children),
	)
}

func BarcodeOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h7v2H4v5H2zm13 0h7v7h-2V4h-5zM9 8v8H7V8zm4 0v8h-2V8zm4 0v8h-2V8zM4 15v5h5v2H2v-7zm18 0v7h-7v-2h5v-5z"/>`), g.Group(children),
	)
}

func BaseStation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6.344 3.93l-.707.707a9 9 0 0 0 0 12.728l.707.707l-1.414 1.414l-.707-.707c-4.296-4.296-4.296-11.26 0-15.557l.707-.707zM19.07 2.515l.708.707c4.295 4.296 4.295 11.261 0 15.557l-.707.707l-1.415-1.414l.707-.707a9 9 0 0 0 0-12.728l-.707-.707zM9.526 7.111l-.707.707a4.5 4.5 0 0 0 0 6.364l.707.707l-1.414 1.415l-.708-.707a6.5 6.5 0 0 1 0-9.193l.708-.707zm6.363-1.414l.707.707a6.5 6.5 0 0 1 0 9.193l-.707.707l-1.414-1.415l.707-.707a4.5 4.5 0 0 0 0-6.364l-.707-.707zM10 11a2 2 0 1 1 3 1.732V23h-2V12.732A2 2 0 0 1 10 11"/>`), g.Group(children),
	)
}

func Battery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 5h21v14H0zm2 2v10h17V7zm22 2v6h-2V9z"/>`), g.Group(children),
	)
}

func BatteryAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 .5h8V3h5v20.5H3V3h5zm2 2V5H5v16.5h14V5h-5V2.5zm3 7v3h3v2h-3v3h-2v-3H8v-2h3v-3z"/>`), g.Group(children),
	)
}

func BatteryCharging(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.376 5.172L8.79 11h6l-5.462 8.876l-1.704-1.048L11.21 13h-6l5.462-8.876zM0 5h7.5v2H2v10h4v2H0zm14 0h7v14h-8v-2h6V7h-5zm10 4v6h-2V9z"/>`), g.Group(children),
	)
}

func BatteryLow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 5h21v14H0zm2 2v10h17V7zm4 2v6H4V9zm18 0v6h-2V9z"/>`), g.Group(children),
	)
}

func Bean(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.164 2.636C13.166 2.118 14.28 2 15.56 2c1.908 0 3.603.782 4.777 2.173c1.169 1.385 1.77 3.306 1.646 5.514c-.218 3.917-1.81 7.018-4.259 9.132C15.288 20.923 12.063 22 8.645 22a8.4 8.4 0 0 1-5.165-1.82C2.042 19.029 1 17.34 1 15.317c0-1.39.34-2.598 1.07-3.615c.722-1.005 1.767-1.742 3.041-2.304q.516-.227 1.087-.42c1.353-.46 1.88-.817 2.137-1.34c.457-.935.878-1.727 1.289-2.392c.755-1.22 1.526-2.085 2.54-2.61m-.433 3.05c1.702 2.194 1.298 4.748-.243 6.224a4.1 4.1 0 0 1-3.228 1.142c-1.045-.098-2.074-.596-2.986-1.508c-.714.392-1.227.832-1.58 1.324C3.25 13.486 3 14.265 3 15.316c0 1.29.658 2.444 1.731 3.304A6.4 6.4 0 0 0 8.645 20c3.006 0 5.748-.945 7.773-2.694c2.014-1.74 3.38-4.323 3.57-7.73c.099-1.785-.39-3.18-1.178-4.113C18.028 4.537 16.899 4 15.56 4c-1.193 0-1.912.12-2.479.412c-.438.227-.864.599-1.35 1.274M7.373 10.68c.386.238.752.35 1.075.38a2.1 2.1 0 0 0 1.657-.595c.646-.619 1.006-1.695.479-2.85q-.217.418-.453.9c-.563 1.15-1.603 1.73-2.758 2.165"/>`), g.Group(children),
	)
}

func Beer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.152 4c.52-1.695 2.316-3 4.348-3c.988 0 1.915.316 2.664.77A5 5 0 0 1 13.55 3h.95A3.5 3.5 0 0 1 18 6.5v.014c0 .237 0 .863-.117 1.486H19c.497 0 1.054.126 1.474.56c.415.428.526.98.526 1.457V18c0 .476-.111 1.031-.535 1.457c-.424.428-.981.543-1.465.543h-2v1c0 .476-.111 1.031-.535 1.457c-.424.428-.981.543-1.465.543H6c-.484 0-1.04-.115-1.466-.543C4.111 22.032 4 21.476 4 21v-5.17A3 3 0 0 1 2 13V7a3 3 0 0 1 3-3zM6 15.83V21h9V10H8v3a3 3 0 0 1-2 2.83M17 10v8h2v-8zm-1.185-2q.04-.091.084-.28c.098-.441.101-.963.101-1.22A1.5 1.5 0 0 0 14.5 5h-2.118l-.276-.553c-.147-.294-.483-.666-.98-.967C10.64 3.184 10.067 3 9.5 3C7.973 3 7 4.154 7 5v1H5a1 1 0 0 0-1 1v6a1 1 0 1 0 2 0V8z"/>`), g.Group(children),
	)
}

func Beta(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 2.491v18.51H4v-2h14v-3.058L3.49 9.476zm-2 11.263V5.509L8.51 9.524z"/>`), g.Group(children),
	)
}

func Bifurcate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2M3 5a3 3 0 1 1 4.086 2.797c.128.667.412 1.506.934 2.256C8.752 11.103 9.958 12 12 12s3.248-.897 3.98-1.947a6 6 0 0 0 .934-2.256a3.001 3.001 0 1 1 2.019.055a8 8 0 0 1-1.312 3.345c-.934 1.34-2.421 2.478-4.621 2.744v2.23a3.001 3.001 0 1 1-2 0v-2.23c-2.2-.266-3.687-1.405-4.62-2.744a8 8 0 0 1-1.313-3.345A3 3 0 0 1 3 5m15-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-6 14a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`), g.Group(children),
	)
}

func Bill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v2h-2v18.08l-1.625-1.3l-1.841-1.472l-1.62 1.44l-.664.59l-.664-.59L12 19.337l-1.586 1.41l-.664.59l-.664-.59l-1.62-1.44l-1.841 1.473L4 22.08V4H2zm4 2v13.92l.875-.7l.659-.528l.63.56l1.586 1.41l1.586-1.41l.664-.59l.664.59l1.586 1.41l1.586-1.41l.63-.56l.659.527l.875.7V4zm2 3h8v2H8zm2 4H9v2h6v-2z"/>`), g.Group(children),
	)
}

func Blockchain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h7v2.5h6V2h7v7h-2.5v6H22v7h-7v-2.5H9V22H2v-7h2.5V9H2zm5 5V4H4v3zm-.5 2v6H9v2.5h6V15h2.5V9H15V6.5H9V9zM17 17v3h3v-3zM7 17H4v3h3zM17 4v3h3V4z"/>`), g.Group(children),
	)
}

func Bluetooth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11 .255l7.453 6.707L13.414 12l5.039 5.038L11 23.745v-9.33l-4 4L5.586 17l5-5l-5-5L7 5.586l4 4zm2 14.16v4.84l2.548-2.293zm0-4.83l2.548-2.547L13 4.745z"/>`), g.Group(children),
	)
}

func Bone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.5 3a1.5 1.5 0 0 0-1.318 2.217l.358.657L5.874 16.54l-.657-.358a1.5 1.5 0 1 0-.993 2.793l.676.125l.125.676a1.5 1.5 0 1 0 2.793-.993l-.358-.657L18.126 7.46l.657.358a1.5 1.5 0 1 0 .993-2.793L19.1 4.9l-.125-.676A1.5 1.5 0 0 0 17.5 3M14 4.5a3.5 3.5 0 0 1 6.764-1.264a3.5 3.5 0 0 1-2.218 6.632l-8.678 8.678a3.5 3.5 0 0 1-6.633 2.218a3.5 3.5 0 0 1 2.219-6.632l8.678-8.678A3.5 3.5 0 0 1 14 4.5"/>`), g.Group(children),
	)
}

func Book(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 6a4 4 0 0 1 4-4h14v20H7a4 4 0 0 1-4-4zm2 8.535A4 4 0 0 1 7 14h12V4H7a2 2 0 0 0-2 2zM19 16H7a2 2 0 1 0 0 4h12zM10 6h7v2h-7z"/>`), g.Group(children),
	)
}

func BookOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2h8a4 4 0 0 1 3 1.354A4 4 0 0 1 15 2h8v19h-1c-2.944 0-5.14.245-6.586.486c-.723.12-1.26.24-1.609.328a11 11 0 0 0-.472.13l-.017.005h-.002l-.152.051h-2.324l-.152-.05l-.001-.001h-.001l-.017-.006l-.088-.026a11 11 0 0 0-.384-.103a21 21 0 0 0-1.61-.328C7.14 21.246 4.946 21 2 21H1zm2 2v15.01c2.563.047 4.535.274 5.914.504c.777.13 1.366.26 1.766.36q.188.047.32.084V6a2 2 0 0 0-2-2zm10 2v13.958q.132-.037.32-.084c.4-.1.989-.23 1.766-.36c1.379-.23 3.35-.457 5.914-.505V4h-6a2 2 0 0 0-2 2m2 2h4v2h-4zm0 3h4v2h-4z"/>`), g.Group(children),
	)
}

func Bookmark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 3h16v19.943l-8-5.714l-8 5.714zm2 2v14.057l6-4.286l6 4.286V5z"/>`), g.Group(children),
	)
}

func BookmarkAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 0v3h3v2h-3v3h-2V5h-3V3h3V0zM4 3h9v2H6v14.057l6-4.286l6 4.286V10h2v12.943l-8-5.714l-8 5.714z"/>`), g.Group(children),
	)
}

func BookmarkChecked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22.596 2.94L16.94 8.595L13.405 5.06l1.414-1.415l2.121 2.122l4.243-4.243zM4 3h8v2H6v14.057l6-4.286l6 4.286V10h2v12.943l-8-5.714l-8 5.714z"/>`), g.Group(children),
	)
}

func BookmarkDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m23 0l.003 18.419L21 16.415V2.001l-10.999.001v-2zM3 4h16v19.943l-8-5.714l-8 5.714zm2 2v14.057l6-4.286l6 4.286V6z"/>`), g.Group(children),
	)
}

func BookmarkMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 3h9v2H6v14.057l6-4.286l6 4.286V7h2v15.943l-8-5.714l-8 5.714zm11 0h8v2h-8z"/>`), g.Group(children),
	)
}

func Braces(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.5 17.6A3.4 3.4 0 0 0 5.9 21H7v-2H5.9a1.4 1.4 0 0 1-1.4-1.4V14c0-.768-.288-1.47-.764-2c.476-.53.764-1.232.764-2V6.4A1.4 1.4 0 0 1 5.9 5H7V3H5.9a3.4 3.4 0 0 0-3.4 3.4V10a1 1 0 0 1-1 1H.4v2h1.1a1 1 0 0 1 1 1zM17 21h1.1a3.4 3.4 0 0 0 3.4-3.4V14a1 1 0 0 1 1-1h1.1v-2h-1.1a1 1 0 0 1-1-1V6.4A3.4 3.4 0 0 0 18.1 3H17v2h1.1a1.4 1.4 0 0 1 1.4 1.4V10c0 .768.288 1.47.764 2a3 3 0 0 0-.764 2v3.6a1.4 1.4 0 0 1-1.4 1.4H17z"/>`), g.Group(children),
	)
}

func Brackets(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 3.5h5v2H4v13h3v2H2zm15 0h5v17h-5v-2h3v-13h-3z"/>`), g.Group(children),
	)
}

func Bread(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.023 11.278c-.051.463.064.86.387 1.183l3.776 3.776l-1.414 1.414l-3.627-3.627l-.03.082c-.558 1.52-.036 3.24 1.175 4.527c1.213 1.289 2.943 1.955 4.538 1.55c2.124-.538 4.425-2.026 6.429-3.946c1.92-2.004 3.407-4.305 3.946-6.429c.405-1.595-.261-3.325-1.55-4.538c-1.287-1.211-3.006-1.733-4.527-1.176l-.082.03l3.627 3.628l-1.414 1.414L12.48 5.39c-.323-.323-.72-.438-1.183-.387c-.489.054-1.03.297-1.49.677c-.463.383-.757.832-.853 1.209c-.085.328-.031.6.23.863l4.243 4.242l-1.414 1.414l-4.242-4.242c-.262-.262-.535-.316-.863-.231c-.377.096-.826.39-1.21.854c-.379.46-.622 1-.676 1.489m1.906-4.369a3 3 0 0 1 .088-.518c.229-.888.82-1.676 1.518-2.253c.702-.58 1.599-1.02 2.544-1.123c.31-.035.627-.032.941.014c.146-.167.312-.285.445-.368c.267-.167.604-.31.973-.445c2.413-.883 4.904.015 6.586 1.598c1.681 1.582 2.743 4.022 2.118 6.486c-.656 2.587-2.395 5.19-4.457 7.337l-.014.014l-.014.014c-2.147 2.062-4.75 3.8-7.337 4.457c-2.464.625-4.904-.437-6.486-2.118c-1.583-1.682-2.481-4.173-1.598-6.586c.136-.37.278-.706.445-.973c.083-.133.201-.299.368-.445a3.7 3.7 0 0 1-.014-.941c.104-.945.543-1.842 1.123-2.544c.577-.698 1.365-1.29 2.253-1.518a3 3 0 0 1 .518-.088"/>`), g.Group(children),
	)
}

func Bridge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v3.413q.172.086.406.192c.61.27 1.547.62 2.81.9C8.482 6.786 10.076 7 12 7s3.52-.214 4.783-.495c1.264-.28 2.2-.63 2.81-.9q.236-.106.407-.192V2h2v13h1v2h-1v5h-2v-5H4v5H2v-5H1v-2h1V2zm0 13h2V8.267a17 17 0 0 1-2-.662zm4-6.31V15h3V8.982a24 24 0 0 1-3-.291m5 .292V15h3V8.69c-.893.145-1.893.251-3 .292m5-.715V15h2V7.605c-.531.215-1.198.449-2 .662m2.446-3.1l.003-.002l-.002.002z"/>`), g.Group(children),
	)
}

func BridgeFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 3h24v18.001h-6v-7.603l-.14-.132a8.4 8.4 0 0 0-1.206-.92A8.64 8.64 0 0 0 12 11c-2.02 0-3.586.671-4.654 1.345A8.4 8.4 0 0 0 6 13.398V21H0zm6 7.836q.134-.09.279-.181A10.64 10.64 0 0 1 12 9a10.64 10.64 0 0 1 6 1.836V5.001h-.625L12 5H6zm14-5.835v14h2v-14zM4 5H2v14h2z"/>`), g.Group(children),
	)
}

func BridgeFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23 2.001v2h-1.18L21 8.098V22h-2V7.901l.78-3.9h-3.56L17 7.9V22h-2V8.099L14.18 4H9.82L9 8.1V22H7V7.902l.78-3.9H4.22L5 7.9V22H3V8.099L2.18 4L1 4.001v-2z"/>`), g.Group(children),
	)
}

func BridgeOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1.999h6V5.53c.208.154.524.363.947.575c.878.438 2.226.894 4.053.894s3.175-.456 4.053-.894c.423-.212.739-.42.947-.575V2h6v20h-6v-5H7v5H1zm6 13h4v-6.04c-1.686-.134-3.004-.594-3.947-1.066L7 7.867zm6-6.04V15h4V7.867l-.053.026c-.943.472-2.26.932-3.947 1.067M3 4v16h2V4zm16 0v16h2V4z"/>`), g.Group(children),
	)
}

func BridgeSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m1 1.754l11 9.9l11-9.9V21h-2v-7h-2v7h-2v-7H7v7H5v-7H3v7H1zM3 12h6.394L3 6.245zm11.606 0h6.393L21 6.245z"/>`), g.Group(children),
	)
}

func BridgeThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v3.465l7 4.667V2h2v8.132l7-4.667V2h2v13h1v2h-1v5h-2v-5H4v5H2v-5H1v-2h1V2zm0 13h7v-2.465L4 7.87zm9 0h7V7.87l-7 4.666z"/>`), g.Group(children),
	)
}

func BridgeTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.298 6.49c-1.208 1.098-2.383 2.927-3.35 5.826c-.082.248-.168.524-.26.817c-.244.783-.525 1.684-.87 2.47c-.39.882-.944 1.836-1.818 2.433V19h6v-5a4 4 0 0 1 8 0v5h6v-.964c-.874-.597-1.429-1.55-1.817-2.433c-.346-.786-.627-1.687-.872-2.47c-.09-.293-.177-.57-.26-.817c-.966-2.9-2.141-4.728-3.349-5.826C14.515 5.41 13.256 5 12 5s-2.515.41-3.702 1.49M6.952 5.01C8.515 3.59 10.256 3 12 3s3.485.59 5.048 2.01c1.542 1.402 2.867 3.573 3.9 6.674q.162.487.31.969c.232.743.458 1.468.756 2.144c.428.974.877 1.523 1.368 1.726l.618.255V21H14v-7a2 2 0 1 0-4 0v7H0v-4.222l.618-.255c.491-.203.94-.752 1.368-1.726c.298-.676.524-1.401.756-2.144q.148-.481.31-.97c1.033-3.1 2.358-5.27 3.9-6.673"/>`), g.Group(children),
	)
}

func Brightness(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .086l3.49 3.49h4.935V8.51l3.49 3.49l-3.49 3.49v4.934H15.49L12 23.914l-3.49-3.49H3.577V15.49L.086 12l3.49-3.49V3.575H8.51zm0 2.828L9.34 5.575H5.576V9.34L2.914 12l2.662 2.662v3.763h3.763L12 21.086l2.661-2.661h3.764V14.66l2.66-2.66l-2.66-2.661V5.575H14.66zM12 8a4 4 0 1 0 0 8a4 4 0 0 0 0-8m-6 4a6 6 0 1 1 12 0a6 6 0 0 1-12 0"/>`), g.Group(children),
	)
}

func BrightnessOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .086l3.49 3.49h4.935V8.51l3.49 3.49l-3.49 3.49v4.934H15.49L12 23.914l-3.49-3.49H3.577V15.49L.086 12l3.49-3.49V3.575H8.51zm0 2.828L9.34 5.575H5.576V9.34L2.914 12l2.662 2.662v3.763h3.763L12 21.086l2.661-2.661h3.764V14.66l2.66-2.66l-2.66-2.661V5.575H14.66zM12.065 8A9 9 0 0 1 13 12a9 9 0 0 1-.936 4a4 4 0 0 0 0-8m-1.939-1.7a6 6 0 1 1 0 11.402l-1.314-.432l.823-1.113A6.96 6.96 0 0 0 11 12a6.96 6.96 0 0 0-1.367-4.157L8.81 6.731z"/>`), g.Group(children),
	)
}

func Broccoli(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.616 4.17c1.303-1.124 3.213-1.07 4.373.091a3.11 3.11 0 0 1 .893 2.46a5 5 0 0 0-.374.047l-.029.005l-.01.002h-.003l-.003.002l-.981.194l.388 1.962l.977-.194h.005l.04-.007a2.6 2.6 0 0 1 .662-.01a3.07 3.07 0 0 1 1.746.87c1.218 1.219 1.218 3.261-.084 4.563a3.4 3.4 0 0 1-.971.68l-.805.373l.274.844c.283.868.071 1.89-.653 2.615c-1.023 1.023-2.62 1.017-3.566.07c-.377-.376-.614-.896-.683-1.471c-.08-.668.074-1.351.43-1.876l.464-.684l-1.832-1.832l-.437-1.31l-1.311-.438l-1.016-1.016l-.62.261c-.873.369-1.803.278-2.628-.222a3.8 3.8 0 0 1-.704-.556C4.94 8.375 4.94 6.333 6.242 5.03c1.21-1.21 3.11-1.195 4.387-.06q.067.06.132.125c.723.722.856 1.32.954 2.083l.128.991l1.983-.254l-.127-.992c-.108-.843-.296-1.771-1.083-2.753M19.865 7a5.11 5.11 0 0 0-1.462-4.153c-2.008-2.009-5.256-1.963-7.307.002c-1.956-1.142-4.52-.982-6.268.767c-2.045 2.045-2.12 5.354-.084 7.39a6 6 0 0 0 .578.509a6 6 0 0 1-.141.624a15 15 0 0 1-.649 1.813c-.569 1.339-1.257 2.53-1.724 2.998l-.707.707L6.343 21.9l.707-.707a18.5 18.5 0 0 1 2.458-2.001c.567-.38 1.129-.692 1.626-.874c.378-.139.643-.174.816-.164a4.4 4.4 0 0 0 1.13 1.997c1.765 1.766 4.629 1.697 6.395-.07a4.6 4.6 0 0 0 1.324-3.83a5.4 5.4 0 0 0 .831-.682c2.046-2.045 2.12-5.354.084-7.39a5.05 5.05 0 0 0-1.849-1.18m-8.014 9.152c-.495.008-.979.132-1.405.288c-.706.259-1.416.666-2.05 1.09a20 20 0 0 0-2.026 1.568l-1.514-1.514c.555-.799 1.097-1.86 1.517-2.85c.3-.705.56-1.425.736-2.065l.062-.234c.788.21 1.62.22 2.44.004l.268.268l-1.414 1.414l1.414 1.415l1.414-1.415l.887.888a4.8 4.8 0 0 0-.33 1.143"/>`), g.Group(children),
	)
}

func Browse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.097 12c1.441 4.08 5.332 7 9.903 7c4.57 0 8.46-2.92 9.902-7C20.461 7.92 16.57 5 12 5s-8.462 2.92-9.903 7m-2.008-.304C1.7 6.654 6.421 3 12 3c5.58 0 10.302 3.654 11.91 8.696l.098.304l-.097.304C22.3 17.346 17.578 21 12 21S1.698 17.346.09 12.304L-.009 12zM12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0"/>`), g.Group(children),
	)
}

func BrowseGallery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19.41 1.586l.692.722A13.96 13.96 0 0 1 24 12c0 3.761-1.485 7.178-3.898 9.692l-.692.722l-1.443-1.385l.692-.721A11.96 11.96 0 0 0 22 12c0-3.225-1.27-6.15-3.34-8.308l-.693-.721zM10 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16M0 12C0 6.477 4.477 2 10 2s10 4.477 10 10s-4.477 10-10 10S0 17.523 0 12m11-5.5v5.086L14.414 15L13 16.414l-4-4V6.5z"/>`), g.Group(children),
	)
}

func BrowseOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m4 1.586l6.171 6.172l7.072 7.07L23.413 21L22 22.414l-2.965-2.965A12.45 12.45 0 0 1 13 21c-5.58 0-10.302-3.654-11.91-8.696L.991 12l.097-.304A12.5 12.5 0 0 1 5.265 5.68L2.585 3zm2.691 5.52A10.53 10.53 0 0 0 3.097 12c1.441 4.08 5.332 7 9.903 7c1.631 0 3.174-.372 4.55-1.035l-1.793-1.793a5 5 0 0 1-6.929-6.929zm3.601 3.6a3 3 0 0 0 4.001 4.001zM13 5q-.866 0-1.692.135l-.987.16l-.32-1.974l.988-.16A12.6 12.6 0 0 1 13 3c5.579 0 10.301 3.654 11.91 8.696l.097.304l-.097.304a12.5 12.5 0 0 1-1.817 3.572l-.591.807l-1.614-1.181l.59-.807A10.5 10.5 0 0 0 22.904 12C21.46 7.92 17.57 5 13 5m.513 1.926l.956.294a5.01 5.01 0 0 1 3.312 3.311l.293.956l-1.912.587l-.293-.956a3.01 3.01 0 0 0-1.987-1.986l-.956-.294z"/>`), g.Group(children),
	)
}

func Brush(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.92 23.494L.794 12.368L10.3 2.86l2.833 2.833l3.274-3.275L21.87 7.88l-3.274 3.275l2.832 2.832zm3.074-5.903L6.697 9.293l-3.075 3.075l8.298 8.297zM8.111 7.879l8.297 8.298l2.191-2.19l-2.833-2.833l3.275-3.275l-2.633-2.632l-3.274 3.274l-2.833-2.832z"/>`), g.Group(children),
	)
}

func Bug(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.17 5h5.66a3.001 3.001 0 0 0-5.66 0M7 6a5 5 0 0 1 10 0v1H7zM5 4v2.586L6.414 8h11.172L19 6.586V4h2v3.414l-3 3V12h3v2h-3v1.586l3 3V22h-2v-2.586l-1.36-1.36A6.01 6.01 0 0 1 13 21.917V22h-1a6 6 0 0 1-5.64-3.946L5 19.414V22H3v-3.414l3-3V14H3v-2h3v-1.586l-3-3V4zm3 6.414V16a4 4 0 0 0 3 3.874V10H8.414zM13 10v9.874c1.725-.444 3-2.01 3-3.874v-5.586L15.586 10z"/>`), g.Group(children),
	)
}

func BugReport(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.128 1.592L9.493 4.43A7.5 7.5 0 0 1 12 4c.878 0 1.722.151 2.507.43l2.365-2.838l1.536 1.28l-2.083 2.5A7.5 7.5 0 0 1 19.073 9H22v2h-2.516q.015.248.016.5v5q0 .252-.016.5H22v2h-2.927a7.503 7.503 0 0 1-14.146 0H2v-2h2.516a8 8 0 0 1-.016-.5v-5q0-.252.016-.5H2V9h2.927a7.5 7.5 0 0 1 2.748-3.628l-2.083-2.5zM12 6a5.5 5.5 0 0 0-5.5 5.5v5a5.5 5.5 0 0 0 11 0v-5A5.5 5.5 0 0 0 12 6m-3 4h6v2H9zm0 6h6v2H9z"/>`), g.Group(children),
	)
}

func Building(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 2h10v3h3v4h2v2h-1v9h1v2H2v-2h1v-9H2V9h2V5h3zm2 3h6V4H9zm-4 6v9h3v-6h8v6h3v-9zm13-2V7H6v2zm-4 11v-4h-4v4z"/>`), g.Group(children),
	)
}

func BuildingFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.998.811l8.384 5.388L19.3 7.882l-.3-.193v3.417l3.375 2.062l-1.043 1.707l-.332-.203V22H3v-7.328l-.332.203l-1.043-1.707L5 11.106V7.689l-.3.193L3.617 6.2zM7 6.403v3.48l5-3.055l5 3.055v-3.48L11.998 3.19zM5 13.45V20h6v-4h2v4h6v-6.55l-7-4.278z"/>`), g.Group(children),
	)
}

func BuildingFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 2h12v2h-1v4h2v14H2V8h8V4H9zm3 2v4h6V4zm8 6H4v10h4v-6h8v6h4zm-6 10v-4h-4v4zM13 4.998h2.004v2.004H13z"/>`), g.Group(children),
	)
}

func BuildingOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 1.5v1.625q.595.16 1.065.51c.549.407.897.951 1.113 1.494c.247.615.344 1.277.364 1.871H11.5v2H10v1h12v12H3V9H1.5V7h1.458c.02-.594.117-1.256.363-1.871c.217-.543.565-1.087 1.114-1.494q.47-.35 1.065-.51V1.5zM4.96 7h3.08c-.018-.389-.083-.789-.219-1.129c-.116-.29-.268-.497-.449-.63C7.201 5.112 6.937 5 6.5 5s-.7.113-.873.24c-.18.134-.332.34-.449.631c-.135.34-.2.74-.219 1.129M8 9H5v11h3zm2 11h2v-5h6v5h2v-8H10zm6 0v-3h-2v3z"/>`), g.Group(children),
	)
}

func BuildingThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 2h8v3h3v4h3v13H2V9h3V5h3zm2 3h4V4h-4zM7 9h10V7H7zm-3 2v9h4v-6h8v6h4v-9zm10 9v-4h-4v4z"/>`), g.Group(children),
	)
}

func BuildingTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .834l6.372 3.823l-1.029 1.715L17 6.166V8h5v14H2V8h5V6.166l-.343.206l-1.029-1.715zM9 4.966V8h6V4.966l-3-1.8zm2 .032h2.004v2.004H11zM4 10v10h4v-7h8v7h4V10zm10 10v-5h-4v5z"/>`), g.Group(children),
	)
}

func Bulletpoint(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 6.004H2V4h2.004v2.004zm0 7H2V11h2.004v2.004zm-1 7h2.004V18H2zM8 4H7v2h15V4zm-1 7h15v2H7zm1 7H7v2h15v-2z"/>`), g.Group(children),
	)
}

func Button(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 5h20v14H2zm2 2v10h16V7zm2 2h12v6H6zm2 2v2h8v-2z"/>`), g.Group(children),
	)
}

func Cabbage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.116 3.468a4 4 0 0 0-1.363-1.28c-1.508-.865-3.451-.713-5.244.291C3.18 4.344 1 7.85 1 12c0 5.54 4.01 10.013 9.247 10.86q.858.139 1.753.14a10.96 10.96 0 0 0 11-11.085c-.023-1.932-.798-4.303-2.023-6.198c-1.22-1.888-3.096-3.608-5.427-3.586c-1.138.002-1.915.628-2.434 1.338m6.655 4.145a4.8 4.8 0 0 0-1.263-.033l-1.16-1.95L15.63 6.65l.892 1.5a7 7 0 0 0-1.065.625l-.018-.058c-.267-.81-.783-1.475-1.471-2.136l.004-.02c.069-.323.179-.74.344-1.145c.373-.916.806-1.287 1.24-1.287h.011c1.291-.014 2.636.977 3.731 2.672q.255.394.474.81M13.61 10.61c-.33.008-.669.045-1 .117l-.032-.052a3.8 3.8 0 0 0-1.03-1.093c-.79-.555-1.872-.83-3.042-.393a14 14 0 0 0-.324-.181l3.233-3.169l-1.4-1.428L6.18 8.17c-.81-.235-1.567-.319-2.247-.225a8.7 8.7 0 0 1 3.555-3.722c1.364-.764 2.537-.722 3.271-.301c.7.401 1.281 1.311 1.135 2.937l-.045.503l.378.335c.822.727 1.164 1.197 1.312 1.646c.104.315.138.696.071 1.267m-2.81 1.005c-.854.715-1.427 1.718-1.76 2.802l-2.974-3.265L4.587 12.5l4.088 4.488c-.005 1.192.148 2.412.37 3.515A8.95 8.95 0 0 1 3 12c0-1.04.289-1.607.629-1.855c.321-.236 1.026-.429 2.47.104c.522.193 1.126.489 1.817.919l.502.312l.516-.29c.623-.35 1.099-.226 1.464.03c.16.112.297.252.402.395m1.765 9.367l7.292-8.037l-1.481-1.344l-2.932 3.231l.007-1.014l-2-.014l-.023 3.25l-2.483 2.737c-.187-1.021-.3-2.118-.264-3.147c.064-1.755.765-3.197 1.802-3.768c.43-.237 1.189-.33 1.789-.224l.89.155l.245-.869c.128-.453.467-.938.997-1.37a4.8 4.8 0 0 1 1.834-.916c.801-.196 1.484-.094 1.934.216c.393.27.788.822.825 1.954v.018L21 12a8.96 8.96 0 0 1-8.435 8.982"/>`), g.Group(children),
	)
}

func Cake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.998 1.998h2.004v2.004H6.998zm4 0h2.004v2.004h-2.004zm4 0h2.004v2.004h-2.004zM9 5v5h2V5h2v5h2V5h2v5h4v10h2v2H1v-2h2V10h4V5zm-4 7v3h1.162L9 15.946l3-1l3 1L17.838 15H19v-3zm14 5h-.838L15 18.054l-3-1l-3 1L5.838 17H5v3h14z"/>`), g.Group(children),
	)
}

func Calculation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 2v3.5H11v2H7.5V11h-2V7.5H2v-2h3.5V2zM13 5.5h9v2h-9zm-2.224 9.511L7.948 17.84l2.989 2.988l-1.415 1.415l-2.989-2.99l-2.828 2.83l-1.414-1.415l2.828-2.828l-2.668-2.668l1.415-1.415l2.667 2.668l2.829-2.828zm5.724-1.013h2.004v2.004H16.5zM13 17h9v2h-9zm3.5 2.998h2.004v2.004H16.5z"/>`), g.Group(children),
	)
}

func CalculationOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h22v22H1zm2 2v18h18V3zm6 2v2h2v2H9v2H7V9H5V7h2V5zm4 2h6v2h-6zm-6.414 6.172L8 14.586l1.414-1.414l1.414 1.414L9.414 16l1.415 1.414l-1.415 1.415L8 17.414L6.586 18.83L5.17 17.414L6.586 16l-1.414-1.414zM13 13.5h6v2h-6zm0 3h6v2h-6z"/>`), g.Group(children),
	)
}

func Calculator(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h18v22H3zm2 2v4h14V3zm14 6H5v12h14zM7 12h4v2H7zm6 0h4v2h-4zm-6 4h4v2H7zm6 0h4v2h-4z"/>`), g.Group(children),
	)
}

func CalculatorOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2h16v20H4zm2 2v4h12V4zm12 6h-3v2h3zm0 4h-3v2h3zm0 4h-3v2h3zm-5 2v-2h-2v2zm-4 0v-2H6v2zm-3-4h3v-2H6zm0-4h3v-2H6zm5-2v2h2v-2zm2 4h-2v2h2z"/>`), g.Group(children),
	)
}

func Calendar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 4V1.5h2V4h8V1.5h2V4h4v18H2V4zM4 6v3h16V6zm16 5H4v9h16z"/>`), g.Group(children),
	)
}

func CalendarEdit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 1v3h8V1h2v3h4v2.001h.001l-.001 6L20 12v-1H4v9h8v2H2V4h4V1zM4 9h16V6H4zm15.287 4.086l4.127 4.127l-5.286 5.287H14v-4.128zm-.672 3.5l1.299 1.3l.672-.673l-1.3-1.299zm-.115 2.713L17.2 18L16 19.2v1.3h1.3z"/>`), g.Group(children),
	)
}

func CalendarEvent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 1v3h8V1h2v3h4v18H2V4h4V1zM4 6v3h16V6zm16 5H4v9h16zm-7.5 1.5h6v6h-6zm2 2v2h2v-2z"/>`), g.Group(children),
	)
}

func CalendarOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 1v3h8V1h2v3h4v18H2V4h4V1zM4 6v3h16V6zm16 5H4v9h16zM7 13h2.004v2.004H7zm4 0h2.004v2.004H11zm4 0h2.004v2.004H15zm-8 3h2.004v2.004H7zm4 0h2.004v2.004H11zm4 0h2.004v2.004H15z"/>`), g.Group(children),
	)
}

func CalendarTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 1v3h8V1h2v3h4v18H2V4h4V1zM4 6v3h16V6zm16 5H4v9h16z"/><path fill="currentColor" d="m16.914 13.25l-5.657 5.657l-3.535-3.536l1.414-1.414l2.121 2.121l4.243-4.242z"/>`), g.Group(children),
	)
}

func Call(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2h8.58l1.487 6.69l-1.86 1.86a14.1 14.1 0 0 0 4.243 4.242l1.86-1.859L22 14.42V23h-1a19.9 19.9 0 0 1-10.85-3.196a20.1 20.1 0 0 1-5.954-5.954A19.9 19.9 0 0 1 1 3zm2.027 2a17.9 17.9 0 0 0 2.849 8.764a18.1 18.1 0 0 0 5.36 5.36A17.9 17.9 0 0 0 20 20.973v-4.949l-4.053-.9l-2.174 2.175l-.663-.377a16.07 16.07 0 0 1-6.032-6.032l-.377-.663l2.175-2.174L7.976 4z"/>`), g.Group(children),
	)
}

func CallCancel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.38 1.966l2.12 2.121l2.121-2.121l1.415 1.414l-2.122 2.121l2.122 2.122l-1.415 1.414l-2.12-2.121l-2.122 2.12l-1.414-1.413L17.087 5.5l-2.121-2.12zM1 2h8.58l1.487 6.69l-1.86 1.86a14.1 14.1 0 0 0 4.243 4.242l1.86-1.859L22 14.42V23h-1a19.9 19.9 0 0 1-10.85-3.196a20.1 20.1 0 0 1-5.954-5.954A19.9 19.9 0 0 1 1 3zm2.027 2a17.9 17.9 0 0 0 2.849 8.764a18.1 18.1 0 0 0 5.36 5.36A17.9 17.9 0 0 0 20 20.973v-4.949l-4.053-.9l-2.174 2.175l-.663-.377a16.07 16.07 0 0 1-6.032-6.032l-.377-.663l2.175-2.174L7.976 4z"/>`), g.Group(children),
	)
}

func CallForwarded(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2h8.58l1.487 6.69l-1.86 1.86a14.1 14.1 0 0 0 4.243 4.242l1.86-1.859L22 14.42V23h-1a19.9 19.9 0 0 1-10.85-3.196a20.1 20.1 0 0 1-5.954-5.954A19.9 19.9 0 0 1 1 3zm2.027 2a17.9 17.9 0 0 0 2.849 8.764a18.1 18.1 0 0 0 5.36 5.36A17.9 17.9 0 0 0 20 20.973v-4.949l-4.053-.9l-2.174 2.175l-.663-.377a16.07 16.07 0 0 1-6.032-6.032l-.377-.663l2.175-2.174L7.976 4zm13.159-1.716l5.418.116l.117 5.419l-2 .043l-.045-2.119l-3.293 3.293l-1.415-1.414l3.294-3.293l-2.12-.046z"/>`), g.Group(children),
	)
}

func CallIncoming(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22.032 3.378L18.738 6.67l2.12.046l-.044 2l-5.418-.117l-.117-5.419l2-.043l.045 2.119l3.293-3.293zM1 2h8.58l1.487 6.69l-1.86 1.86a14.1 14.1 0 0 0 4.243 4.242l1.86-1.859L22 14.42V23h-1a19.9 19.9 0 0 1-10.85-3.196a20.1 20.1 0 0 1-5.954-5.954A19.9 19.9 0 0 1 1 3zm2.027 2a17.9 17.9 0 0 0 2.849 8.764a18.1 18.1 0 0 0 5.36 5.36A17.9 17.9 0 0 0 20 20.974v-4.948l-4.053-.901l-2.174 2.175l-.663-.377a16.07 16.07 0 0 1-6.032-6.032l-.377-.663l2.175-2.174L7.976 4z"/>`), g.Group(children),
	)
}

func CallOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22.414 3L11.833 13.58q.764.662 1.617 1.211l1.86-1.859L22 14.42V23h-1a19.9 19.9 0 0 1-10.85-3.197a20 20 0 0 1-2.567-1.971L3 22.414L1.586 21L21 1.586zM9 16.415a18 18 0 0 0 2.237 1.71A17.9 17.9 0 0 0 20 20.972v-4.949l-4.053-.9l-2.174 2.175l-.663-.377A16 16 0 0 1 10.415 15zM1 2h8.58l1.487 6.69l-1.865 1.866l.297.504l-1.723 1.015l-1.084-1.839l2.184-2.183L7.976 4H3.027a17.9 17.9 0 0 0 3.097 9.138l.564.825l-1.652 1.128l-.564-.825A19.9 19.9 0 0 1 1 3z"/>`), g.Group(children),
	)
}

func CallOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.052 1.787l.966.261a7 7 0 0 1 4.93 4.934l.26.965l-1.93.521l-.261-.965a5 5 0 0 0-3.522-3.524l-.965-.262zM1 2h8.58l1.487 6.69l-1.86 1.86a14.1 14.1 0 0 0 4.243 4.242l1.86-1.859L22 14.42V23h-1a19.9 19.9 0 0 1-10.85-3.196a20.1 20.1 0 0 1-5.954-5.954A19.9 19.9 0 0 1 1 3zm2.027 2a17.9 17.9 0 0 0 2.849 8.764a18.1 18.1 0 0 0 5.36 5.36A17.9 17.9 0 0 0 20 20.973v-4.949l-4.053-.9l-2.174 2.175l-.663-.377a16.07 16.07 0 0 1-6.032-6.032l-.377-.663l2.175-2.174L7.976 4zm12.111 1.165l.966.261a3.5 3.5 0 0 1 2.465 2.467l.26.965l-1.93.521l-.261-.965a1.5 1.5 0 0 0-1.057-1.057l-.965-.261z"/>`), g.Group(children),
	)
}

func Calm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m8-4v4H7V8zm8 0v4h-2V8zm-6.5 6H16v2h-5.5z"/>`), g.Group(children),
	)
}

func CalmOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h4v2H9v2H7v-2H6zm8 0h4v2h-1v2h-2v-2h-1zm-3 5h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func Camera(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.882 2h8.236l1.5 3H23v16H1V5h5.382zm1.236 2l-1.5 3H3v12h18V7h-4.618l-1.5-3zM12 9.5a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0"/>`), g.Group(children),
	)
}

func CameraOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 .586L22.414 21L21 22.414L19.586 21H1V5h2.586l-3-3zM5.586 7H3v12h14.586l-2.537-2.537a5 5 0 0 1-7.012-7.012zm3.885 3.885a3 3 0 0 0 4.144 4.144zM7.882 2h8.236l1.5 3H23v14h-2V7h-4.618l-1.5-3H9.118l-.64 1.279l-1.788-.894z"/>`), g.Group(children),
	)
}

func CameraOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a7 7 0 1 0 0 14a7 7 0 0 0 0-14m-9 7a9 9 0 1 1 13.148 7.989L16.805 21H19v2H5v-2h2.195l.657-3.011A9 9 0 0 1 3 10m6.74 8.714L9.243 21h5.516l-.499-2.286A9 9 0 0 1 12 19c-.78 0-1.537-.1-2.26-.286M12 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0"/>`), g.Group(children),
	)
}

func CameraTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3.5a7 7 0 1 0 0 14a7 7 0 0 0 0-14m-9 7a9 9 0 1 1 10 8.945v1.01a9.94 9.94 0 0 0 4.626-1.682l.826-.563l1.127 1.652l-.826.564A11.94 11.94 0 0 1 13 22.464v.036h-.652a12 12 0 0 1-.696 0H11v-.036a11.94 11.94 0 0 1-5.753-2.038l-.826-.564l1.127-1.652l.826.563A9.94 9.94 0 0 0 11 20.456v-1.01c-4.5-.498-8-4.313-8-8.946m9-2a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0"/>`), g.Group(children),
	)
}

func Candy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.31 3.536c.154-1.69 2.208-2.433 3.407-1.234l3.98 3.981c1.2 1.199.456 3.253-1.233 3.406l-1.314.12a3 3 0 0 1-.357 3.812l-5.172 5.172a3 3 0 0 1-3.812.357l-.12 1.314c-.153 1.689-2.207 2.432-3.406 1.233l-3.98-3.98c-1.2-1.2-.456-3.253 1.233-3.406l1.314-.12a3 3 0 0 1 .357-3.812l5.172-5.172a3 3 0 0 1 3.812-.357zm2.058 3.418l.001.001l.676.675l.001.002a1 1 0 0 0 .796.287l2.441-.222l-3.98-3.98l-.223 2.44a1 1 0 0 0 .288.797m-1.415 1.413l-1.746-1.746a1 1 0 0 0-1.414 0L9.914 8.5l5.586 5.586l1.878-1.88a1 1 0 0 0 0-1.413l-1.745-1.746l-.003-.002l-.675-.676zm-.867 7.133L8.5 9.914l-1.88 1.879a1 1 0 0 0 0 1.414l4.172 4.171a1 1 0 0 0 1.414 0zm-7.462.538l-2.907.264l3.98 3.981l.265-2.907z"/>`), g.Group(children),
	)
}

func Card(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 3h12v2h4v14h-4v2H6v-2H2V5h4zm0 4H4v10h2zm2 12h8V5H8zM18 7v10h2V7z"/>`), g.Group(children),
	)
}

func Cardmembership(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23 3v18H1V3zm-2 2H3v4h18zm0 6h-2v6.766l-3.5-2.1l-3.5 2.1V11H3v8h18zm-4 0h-3v3.234l1.5-.9l1.5.9z"/>`), g.Group(children),
	)
}

func CaretDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 17.914L21.414 8.5H2.586zM7.414 10.5h9.172L12 15.086z"/>`), g.Group(children),
	)
}

func CaretDownSmall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 16.5l7-7H5z"/>`), g.Group(children),
	)
}

func CaretLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.5 2.586v18.828L6.086 12zM8.914 12l4.586 4.586V7.414z"/>`), g.Group(children),
	)
}

func CaretLeftSmall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.5 5v14l-7-7z"/>`), g.Group(children),
	)
}

func CaretRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.5 21.414L17.914 12L8.5 2.586zm2-4.828V7.414L15.086 12z"/>`), g.Group(children),
	)
}

func CaretRightSmall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.5 12l-7-7v14z"/>`), g.Group(children),
	)
}

func CaretUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 6.086l9.414 9.414H2.586zM7.414 13.5h9.172L12 8.914z"/>`), g.Group(children),
	)
}

func CaretUpSmall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 7.5l7 7H5z"/>`), g.Group(children),
	)
}

func Cart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 1h4.764l.545 2h18.078l-3.666 11H7.78l-.5 2H22v2H4.72l1.246-4.989L3.236 3H0zm7.764 11h10.515l2.334-7H5.855zM4 21a2 2 0 1 1 4 0a2 2 0 0 1-4 0m14 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func CartAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 1h4.764l3 11h10.515l3.089-9.265l1.897.633L19.72 14H7.78l-.5 2H22v2H4.72l1.246-4.989L3.236 3H0zm14 1v3h3v2h-3v3h-2V7H9V5h3V2zM4 21a2 2 0 1 1 4 0a2 2 0 0 1-4 0m14 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func Cast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H13v-2h8V5H3v4H1zm0 8h1a9 9 0 0 1 9 9v1H9v-1a7 7 0 0 0-7-7H1zm0 4h1a5 5 0 0 1 5 5v1H5v-1a3 3 0 0 0-3-3H1zm0 4h2.004v2.004H1z"/>`), g.Group(children),
	)
}

func Castle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 2v2h2V2h2v2h2V2h2v2h2V2h2v12h2v-1h2v9H1v-9h2v1h2V2zM5 16H3v4h2zm2 4h2v-7h6v7h2V6H7zm12 0h2v-4h-2zm-6 0v-5h-2v5z"/>`), g.Group(children),
	)
}

func CastleFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 2v1h-4V2h-2v7.131l-2-1.333l-2 1.333V2H8v1H4V2H2v20h20V2zM4 5h4v5H6v10H4zm4 15v-8h1.303L12 10.202L14.697 12H16v8h-3v-5h-2v5zm10 0V10h-2V5h4v15z"/>`), g.Group(children),
	)
}

func CastleFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 2h-8v6.342a6 6 0 0 0-1-.259V7h-2v1.083a6 6 0 0 0-1 .259V2H2v20h20zM8 9.528A6 6 0 0 0 6 14v6H4V4h4zM8 20v-5h8v5h-3v-3h-2v3zm10 0v-6a6 6 0 0 0-2-4.472V4h4v16zm-2.126-7H8.126a4.003 4.003 0 0 1 7.748 0"/>`), g.Group(children),
	)
}

func CastleOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 2v2h2V2h2v2h2V2h2v2h2V2h2v4h-1v8h3v-1h2v9H1v-9h2v1h3V6H5V2zm1 4v14h3v-7h2v7h3V6zm10 14h3v-4h-3zM6 20v-4H3v4zM9 7.998h2.004v2.004H9zm4 0h2.004v2.004H13z"/>`), g.Group(children),
	)
}

func CastleSeven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v1.083c2.838.476 5 2.944 5 5.917h4v14H2V8h4a6 6 0 0 1 5-5.917V1zM8 8h8a4 4 0 0 0-8 0m-4 4.535A4 4 0 0 1 6 12a4 4 0 0 1 3 1.354A4 4 0 0 1 12 12a4 4 0 0 1 3 1.354A4 4 0 0 1 18 12c.729 0 1.412.195 2 .535V10H4zM8 16a2 2 0 1 0-4 0v4h4zm2 4h4v-4a2 2 0 1 0-4 0zm6 0h4v-4a2 2 0 1 0-4 0z"/>`), g.Group(children),
	)
}

func CastleSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 2h12v2h-1v4h5v14H2V8h5V4H6zm3 2v4h2V6h2v2h2V4zm-5 8.535A4 4 0 0 1 6 12a4 4 0 0 1 3 1.354A4 4 0 0 1 12 12a4 4 0 0 1 3 1.354A4 4 0 0 1 18 12c.729 0 1.412.195 2 .535V10H4zM8 16a2 2 0 1 0-4 0v4h4zm2 4h4v-4a2 2 0 1 0-4 0zm6 0h4v-4a2 2 0 1 0-4 0z"/>`), g.Group(children),
	)
}

func CastleThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 2v1h-4V2h-2v8h-4V2H8v1H4V2H2v20h20V2zM4 5h4v5H6v10H4zm4 15v-8h8v8h-3v-5h-2v5zm10 0V10h-2V5h4v15z"/>`), g.Group(children),
	)
}

func CastleTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 2v1.083A6 6 0 0 0 6 9a5 5 0 0 0-5 5v8h22v-8a5 5 0 0 0-5-5a6 6 0 0 0-5-5.917V2zm7 9a3 3 0 0 1 3 3v6h-3zm-2 9h-3v-5h-2v5H8V10h8zM6 20H3v-6a3 3 0 0 1 3-3zM8.126 8a4.002 4.002 0 0 1 7.748 0z"/>`), g.Group(children),
	)
}

func Cat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 .661L7.48 2.58c1.38-.605 2.911-.941 4.52-.941s3.14.336 4.52.94L23 .662v12.053h-.007C22.778 18.458 17.91 23 12 23S1.222 18.458 1.007 12.714H1zm2 5.516a11 11 0 0 1 2.137-2.205L3 3.339zm0 6.143C3 17.083 6.999 21 12 21s9-3.917 9-8.68c0-3.39-2.017-6.344-4.99-7.774A9.2 9.2 0 0 0 12 3.64a9.2 9.2 0 0 0-4.01.907C5.016 5.976 3 8.93 3 12.32m18-6.143V3.34l-2.137.633A11 11 0 0 1 21 6.177M9 9v3H7V9zm8 0v3h-2V9zm-5 3.264l.894 1.789c.379.757.857.985 1.189 1.013c.34.028.759-.131 1.085-.62l.555-.833l1.664 1.11l-.555.832c-.674 1.01-1.755 1.6-2.915 1.504c-.713-.06-1.371-.37-1.917-.885c-.546.515-1.204.826-1.917.885c-1.16.097-2.241-.493-2.915-1.504l-.555-.832l1.664-1.11l.555.832c.326.49.745.65 1.085.621c.332-.028.81-.256 1.189-1.013z"/>`), g.Group(children),
	)
}

func Catalog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h18v22H3zm2 2v18h14V3zm3 4h8v2H8zm0 4h8v2H8zm0 4h8v2H8z"/>`), g.Group(children),
	)
}

func Cd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m10-8h1a8 8 0 0 1 8 8v1h-2v-1a6 6 0 0 0-6-6h-1zm1 6a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0"/>`), g.Group(children),
	)
}

func Celsius(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.5 6a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1M4 6.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0m7-.5a2 2 0 0 1 2-2h7v2h-7v12h7v2h-7a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func CenterFocusStrong(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h7v2H4v5H2zm13 0h7v7h-2V4h-5zm-3 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0m-4 3v5h5v2H2v-7zm18 0v7h-7v-2h5v-5z"/>`), g.Group(children),
	)
}

func Centimeter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 4h10a2 2 0 0 1 2 2v14h-2V6h-3v14h-2V6h-3v14h-2zM1 6a2 2 0 0 1 2-2h7v2H3v12h7v2H3a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func Certificate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v14h18V5h-2v6.5l-3-2.25l-3 2.25V5zm12 0v2.5l1-.75l1 .75V5zM5 11h6v2H5zm0 4h14v2H5z"/>`), g.Group(children),
	)
}

func CertificateOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21 1v22H3V1zm-8 2v6.5l-3-2.25L7 9.5V3H5v18h14V3zM9 3v2.5l1-.75l1 .75V3zm-2 9h10v2H7zm0 4h8v2H7z"/>`), g.Group(children),
	)
}

func Chart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm9 3v11h-2V7zm-5 4v7H6v-7zm10 2v5h-2v-5z"/>`), g.Group(children),
	)
}

func ChartAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 0v3h3v2h-3v3h-2V5h-3V3h3V0zM2 2h11v2H4v16h16V10h2v12H2zm11 6v10h-2V8zm-5 3v7H6v-7zm10 2v5h-2v-5z"/>`), g.Group(children),
	)
}

func ChartAnalytics(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm12 4h6v6h-2V9.423l-1.579 1.575l-2.714 2.708l-.707.707l-4-4l-5 5L4.586 14L11 7.585l4 4l2.009-2.003L18.594 8H16z"/>`), g.Group(children),
	)
}

func ChartArea(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm17 2.086V18.21H6v-6.152l6.59-5.99l2.967 3.461zm-2 4.828l-3.556 3.557l-3.033-3.538L8 12.943v3.267h11z"/>`), g.Group(children),
	)
}

func ChartAreaMulti(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm17 2.086V18.21H6v-6.152l6.59-5.99l2.967 3.461zM8.258 16.21H19v-3.796l-3.73 3.73l-3.034-3.539zM19 9.586v-.672l-3.556 3.557l-3.033-3.538L8 12.943v.801l4.415-4.001l2.967 3.46z"/>`), g.Group(children),
	)
}

func ChartBar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h2V9h6v11h2V5h6v15h2v2H2V2zm14 18V7h-2v13zm-8 0v-9H8v9z"/>`), g.Group(children),
	)
}

func ChartBubble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm12 3a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0m-3 6a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0m11 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func ChartColumn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v2h15v6H4v2h11v6H4v2h18v2H2V2zm0 14h9v-2H4zm0-8h13V6H4z"/>`), g.Group(children),
	)
}

func ChartCombo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h2v-6h6v6h2v-9h6v9h2v2H2V2zm14 18v-7h-2v7zm-8 0v-4H8v4zm6.673-16.273L21.246 8.3l-1.414 1.414l-3.163-3.163l-6.782 6.74l-4.594-4.594l1.414-1.415l3.184 3.185z"/>`), g.Group(children),
	)
}

func ChartLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm17.914 6L15.5 14.414l-4-4l-5 5L5.086 14L11.5 7.586l4 4l5-5z"/>`), g.Group(children),
	)
}

func ChartLineData(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm15.5 5a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1m-2.5.5a2.5 2.5 0 1 1 1.026 2.02l-8.041 4.696a2.5 2.5 0 1 1-1.003-1.73l8.035-4.693A3 3 0 0 1 17 7.5m-9.067 6.75a.5.5 0 1 0-.866.5a.5.5 0 0 0 .866-.5"/>`), g.Group(children),
	)
}

func ChartLineDataOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm16.307 5.27l1.782.91l-.454.89l-.891-.454l.89.454v.002l-.003.005l-.007.015l-.027.05l-.099.184a19 19 0 0 1-.364.635c-.31.516-.756 1.205-1.291 1.878c-.527.663-1.18 1.36-1.914 1.844c-.73.482-1.68.842-2.703.54c-1.704-.504-2.248-1.847-2.615-2.751l-.032-.078c-.39-.962-.626-1.434-1.375-1.646c-.236-.067-.572-.024-1.06.299c-.484.32-.985.833-1.454 1.423a16 16 0 0 0-1.138 1.662a17 17 0 0 0-.407.72l-.02.038l-.005.009l-.453.892l-1.783-.904l.452-.892l.892.452l-.892-.452l.001-.003l.003-.004l.007-.015l.027-.051l.098-.184c.085-.155.208-.374.364-.633a17.5 17.5 0 0 1 1.288-1.879c.528-.664 1.181-1.36 1.918-1.848c.732-.484 1.683-.845 2.708-.554c1.738.492 2.288 1.846 2.66 2.76l.022.056c.396.972.626 1.447 1.362 1.665c.222.066.547.03 1.033-.29c.48-.319.98-.83 1.45-1.421c.463-.582.86-1.193 1.142-1.663a17 17 0 0 0 .409-.721l.02-.039l.005-.008v-.001z"/>`), g.Group(children),
	)
}

func ChartLineMulti(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm17.914 3.5L15.5 11.914l-4-4l-5 5L5.086 11.5L11.5 5.086l4 4l5-5zm0 5L15.5 16.914l-4-4l-5 5L5.086 16.5l6.414-6.414l4 4l5-5z"/>`), g.Group(children),
	)
}

func ChartMaximum(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm1.998.998h2.004v2.004H5.998zm3 0h2.004v2.004H8.998zm3 0h2.004v2.004h-2.004zm3 0h2.004v2.004h-2.004zm3 0h2.004v2.004h-2.004zm-2.567 5.789C14.816 8.304 14.035 7.999 13 8c-1.034 0-1.816.306-2.432.79c-.63.495-1.136 1.218-1.53 2.116C8.247 12.721 8 15.059 8 16.998v1H6v-1c0-2.047.252-4.707 1.207-6.893c.482-1.102 1.163-2.131 2.127-2.888C10.312 6.448 11.532 6 12.999 6s2.688.445 3.667 1.214c.964.757 1.645 1.787 2.127 2.89c.955 2.189 1.207 4.852 1.207 6.9v1h-2v-1c0-1.94-.246-4.282-1.04-6.1c-.393-.899-.899-1.622-1.53-2.117"/>`), g.Group(children),
	)
}

func ChartMedian(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm12 2.404c-1.075 0-2 .909-2 2.096v2.998h.004v2.004H14v3.07C14 16.887 12.33 19 10 19s-4-2.112-4-4.429v-1h2v1C8 16.042 9.017 17 10 17s2-.958 2-2.429V6.5c0-2.233 1.762-4.096 4-4.096s4 1.863 4 4.096v1h-2v-1c0-1.187-.925-2.096-2-2.096M6 9.496h2.004V11.5H6zm2.996.002H11v2.004H8.996zm6 0H17v2.004h-2.004zm3.004 0h2.004v2.004H18z"/>`), g.Group(children),
	)
}

func ChartMinimum(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm16 1v1c0 2.048-.252 4.71-1.207 6.899c-.482 1.103-1.163 2.133-2.127 2.89c-.979.77-2.199 1.216-3.667 1.215c-1.468-.002-2.687-.449-3.665-1.217C8.37 13.03 7.689 12 7.207 10.899C6.252 8.712 6 6.053 6 4.006v-1h2v1c0 1.939.246 4.276 1.04 6.092c.393.898.899 1.62 1.53 2.116c.615.484 1.396.789 2.43.79c1.034 0 1.816-.304 2.43-.787c.631-.495 1.137-1.219 1.53-2.118C17.754 8.28 18 5.94 18 3.999V3zM5.998 15.998h2.004v2.004H5.998zm3 0h2.004v2.004H8.998zm3 0h2.004v2.004h-2.004zm3 0h2.004v2.004h-2.004zm3 0h2.004v2.004h-2.004z"/>`), g.Group(children),
	)
}

func ChartPie(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 3.055A9.001 9.001 0 1 0 20.945 13H11zm2 0V11h7.945A9.004 9.004 0 0 0 13 3.055M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12"/>`), g.Group(children),
	)
}

func ChartRadar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 .764L23.186 8.89l-4.273 13.15H5.087L.814 8.89zM3.548 10.83l2.61 8.035l4.224-5.814zm4.228 9.21h8.448L12 14.226zm5.842-6.99l4.224 5.815l2.61-8.035zm6.216-4.122L13 3.963v7.186zM11 3.963L4.166 8.928L11 11.148z"/>`), g.Group(children),
	)
}

func ChartRadial(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 9 9v-1h2v1c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1h1v2zm0 4a5 5 0 1 0 3.515 1.444l-.711-.703l1.406-1.423l.71.703A7 7 0 1 1 12 5h1v2zm0 4q-.133 0-.25.031A1 1 0 1 0 13 12v-1h2v1a3 3 0 1 1-3-3h1v2z"/>`), g.Group(children),
	)
}

func ChartRing(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 3.055a9.001 9.001 0 1 0 6.631 15.966l-2.123-2.153A6 6 0 1 1 11 6.083zm2 0v3.028A6.005 6.005 0 0 1 17.917 11h3.028A9.004 9.004 0 0 0 13 3.055M20.945 13h-3.028a6 6 0 0 1-1.004 2.445l2.13 2.16A8.96 8.96 0 0 0 20.945 13M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m11-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8"/>`), g.Group(children),
	)
}

func ChartRingOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.174 2.983l-.98.2a9 9 0 1 0 3.612 17.635a9 9 0 0 0 7.057-7.237l.176-.984l1.969.351l-.176.985c-.769 4.302-4.079 7.913-8.625 8.844c-5.951 1.22-11.764-2.617-12.983-8.569C.004 8.257 3.841 2.444 9.793 1.225l.98-.201zM13.31 1.03l.978.209a11 11 0 0 1 8.488 8.555l.201.98l-1.96.4l-.2-.98a9 9 0 0 0-6.946-7l-.978-.208zm2.61 10.168a4 4 0 1 0-7.838 1.605a4 4 0 0 0 7.838-1.605m-5.123-5.076a6 6 0 1 1 2.408 11.756a6 6 0 0 1-2.408-11.756"/>`), g.Group(children),
	)
}

func ChartScatter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h18v2H2V2zm13 4a2 2 0 1 1 4 0a2 2 0 0 1-4 0m-6.002-1h2.004v2.004h-2.004zm-5 3h2.004v2.004H5.998zm9 2h2.004v2.004h-2.004zM9 13a2 2 0 1 1 4 0a2 2 0 0 1-4 0m7.998 2h2.004v2.004h-2.004zm-11 1h2.004v2.004H5.998z"/>`), g.Group(children),
	)
}

func ChartStacked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v18h2V9h6v11h2V5h6v15h2v2H2V2zm14 18v-3.5h-2V20zm-2-5.5h2V7h-2zM10 20v-3.5H8V20zm-2-5.5h2V11H8z"/>`), g.Group(children),
	)
}

func ChartThreeD(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .845l9.66 5.578v11.154L12 23.155l-9.66-5.578V6.423zm0 2.31L5.34 7L12 10.845L18.66 7zm7.66 5.577L13 12.577v7.69l6.66-3.844zM11 20.268v-7.69L4.34 8.731v7.69zM13 5v4h-2V5zm-3.098 9.366l-3.464 2l-1-1.732l3.464-2zm5.196-1.732l3.464 2l-1 1.732l-3.464-2z"/>`), g.Group(children),
	)
}

func Chat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704zm2 2v14.296L6.124 16H20.5V4z"/>`), g.Group(children),
	)
}

func ChatAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21 0v3h3v2h-3v3h-2V5h-3V3h3V0zM1.5 2H14v2H3.5v14.296L6.124 16H20.5v-6h2v8H6.876L1.5 22.704z"/>`), g.Group(children),
	)
}

func ChatBubble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9c0 2.397.935 4.573 2.463 6.186l.504.532L4.7 21H12a9 9 0 1 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11H1.3l2.22-3.994A10.96 10.96 0 0 1 1 12"/>`), g.Group(children),
	)
}

func ChatBubbleAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9c0 2.397.935 4.573 2.463 6.186l.504.532L4.7 21H12a9 9 0 1 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11H1.3l2.22-3.994A10.96 10.96 0 0 1 1 12m12-4v3h3v2h-3v3h-2v-3H8v-2h3V8z"/>`), g.Group(children),
	)
}

func ChatBubbleError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9c0 2.397.935 4.573 2.463 6.186l.504.532L4.7 21H12a9 9 0 1 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11H1.3l2.22-3.994A10.96 10.96 0 0 1 1 12m12-5.5V14h-2V6.5zm-2 9h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func ChatBubbleHelp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9c0 2.397.935 4.573 2.463 6.186l.504.532L4.7 21H12a9 9 0 1 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11H1.3l2.22-3.994A10.96 10.96 0 0 1 1 12m11-4.5a2 2 0 0 0-1.886 1.333l-.334.943l-1.885-.666l.333-.943a4.001 4.001 0 1 1 6.246 4.476c-.431.34-.817.666-1.096 1.009c-.274.338-.378.61-.378.848v1.25h-2V14.5c0-.867.39-1.573.826-2.11c.432-.53.974-.974 1.41-1.318A2 2 0 0 0 12 7.5m-1 9.25h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func ChatBubbleHistory(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9c0 2.397.935 4.573 2.463 6.186l.504.532L4.7 21H12a9 9 0 1 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11H1.3l2.22-3.994A10.96 10.96 0 0 1 1 12m12-6.5v6.086L16.414 15L15 16.414l-4-4V5.5z"/>`), g.Group(children),
	)
}

func ChatBubbleLocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9c0 2.397.935 4.573 2.463 6.186l.504.532L4.7 21H12a9 9 0 1 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11H1.3l2.22-3.994A10.96 10.96 0 0 1 1 12m11-3.5c.69 0 1.25.56 1.25 1.25v.75h-2.5v-.75c0-.69.56-1.25 1.25-1.25m3.25 2v-.75a3.25 3.25 0 0 0-6.5 0v.75H7.499V17h9v-6.5zm-.751 2V15h-5v-2.5z"/>`), g.Group(children),
	)
}

func ChatBubbleOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9c0 2.397.935 4.573 2.463 6.186l.504.532L4.7 21H12a9 9 0 1 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11H1.3l2.22-3.994A10.96 10.96 0 0 1 1 12m7-3h8v2H8zm0 4h6v2H8z"/>`), g.Group(children),
	)
}

func ChatBubbleSmile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m8.333 13.6l-.4-.917l-1.833.801l.4.916a6.001 6.001 0 0 0 11 0l.401-.916l-1.833-.8l-.4.916a4.001 4.001 0 0 1-7.335 0"/><path fill="currentColor" d="M12 1C5.925 1 1 5.925 1 12c0 2.662.946 5.104 2.52 7.006L1.3 23H12c6.075 0 11-4.925 11-11S18.075 1 12 1M3 12a9 9 0 1 1 9 9H4.7l1.267-2.282l-.504-.532A8.97 8.97 0 0 1 3 12"/>`), g.Group(children),
	)
}

func ChatChecked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704zm2 2v14.296L6.124 16H20.5V4zm14.157 3.586l-7.071 7.07l-4.243-4.242L7.757 9l2.829 2.828l5.657-5.657z"/>`), g.Group(children),
	)
}

func ChatClear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704zm2 2v14.296L6.124 16H20.5V4zm5.671 1.757L12 8.586l2.828-2.829l1.414 1.415L13.414 10l2.828 2.828l-1.414 1.415L12 11.414l-2.83 2.829l-1.414-1.415L10.585 10L7.757 7.172z"/>`), g.Group(children),
	)
}

func ChatDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2h19v14H5.82L1 19.443zm2 2v11.557L5.18 14H18V4zm20.5.5v18.443L18.68 19.5H7.5v-2h11.82l2.18 1.557V4.5z"/>`), g.Group(children),
	)
}

func ChatError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704zm2 2v14.296L6.124 16H20.5V4zM13 6v5h-2V6zm-2 6h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func ChatHeart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704zm2 2v14.296L6.124 16H20.5V4zM12 5.772a3.251 3.251 0 0 0-4.065 5.026L12 14.864l4.065-4.066A3.25 3.25 0 0 0 12 5.772m-.883 1.844L12 8.5l.884-.884a1.25 1.25 0 0 1 1.768 1.768l-2.651 2.652l-2.652-2.652a1.25 1.25 0 0 1 1.768-1.768"/>`), g.Group(children),
	)
}

func ChatMessage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 22.704V2h21v16H6.876zm2-4.408L6.124 16H20.5V4h-17zM13.004 11H11V8.996h2.004zm-5 0H6V8.996h2.004zm10 0H16V8.996h2.004z"/>`), g.Group(children),
	)
}

func ChatOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23.414 22.001L2 .587L.586 2l.914.914v19.79L6.876 18h9.71L22 23.415zm-8.828-6H6.124L3.5 18.297V4.915zM22.503 18L22.5 2H9.655l-3.418-.003L8.24 4H20.5v11.995z"/>`), g.Group(children),
	)
}

func ChatPoll(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704zm2 2v14.296L6.124 16H20.5V4zM13 6v8h-2V6zm5 3v5h-2V9zM8 11v3H6v-3z"/>`), g.Group(children),
	)
}

func ChatSetting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v9h-2V4h-17v14.296L6.124 16H13v2H6.876L1.5 22.704zM20 12.5v1.14a3.5 3.5 0 0 1 1.405.814l.99-.571l1 1.732l-.99.571a3.5 3.5 0 0 1 0 1.623l.99.572l-1 1.732l-.993-.573a3.5 3.5 0 0 1-1.403.81v1.145h-2V20.35a3.5 3.5 0 0 1-1.403-.81l-.992.573l-1-1.732l.99-.572a3.5 3.5 0 0 1 0-1.623l-.99-.571l1-1.732l.989.57a3.5 3.5 0 0 1 1.406-.813V12.5zm-1 2.995a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/>`), g.Group(children),
	)
}

func Check(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.985 7.378L10.38 17.985l-6.364-6.364l1.414-1.414l4.95 4.95l9.192-9.193z"/>`), g.Group(children),
	)
}

func CheckCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-12.5 4.414L6.086 12L7.5 10.586l3 3l6-6L17.914 9z"/>`), g.Group(children),
	)
}

func CheckCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11M7.5 10.586l3 3l6-6L17.914 9L10.5 16.414L6.086 12z"/>`), g.Group(children),
	)
}

func CheckDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.657 7.171l-5.303 5.304l-1.414-1.415l5.303-5.303zm5.657 0L12.708 17.778l-6.364-6.364L7.758 10l4.95 4.95L21.9 5.757zM2.101 10l4.95 4.95l.353-.354l1.414 1.414l-1.768 1.768l-6.363-6.364z"/>`), g.Group(children),
	)
}

func CheckRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm13.914 5L10.5 16.414L6.086 12L7.5 10.586l3 3l6-6z"/>`), g.Group(children),
	)
}

func CheckRectangleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 2H2v20h20zM6.086 12L7.5 10.586l3 3l6-6L17.915 9L10.5 16.414z"/>`), g.Group(children),
	)
}

func Cheese(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m22.533 1.467l-3.998 12.658l-.833-.112a1.5 1.5 0 0 0-.89 2.82l.746.387l-1.885 5.968l-1.662-.395c-6.318-1.505-11.29-6.446-12.795-12.769L.812 8.327l10.97-3.465l.266 1.018a1.501 1.501 0 1 0 2.858-.903l-.366-.986zm-5.538 3.846A3.5 3.5 0 0 1 13.5 9a3.5 3.5 0 0 1-2.969-1.646L3.19 9.674c1.349 5.48 5.657 9.774 11.138 11.137l.85-2.693a3.5 3.5 0 0 1 1.925-6.096l2.365-7.49zM9.5 11a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M6 12.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0"/>`), g.Group(children),
	)
}

func Cherry(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 .586L23.414 7L22 8.414L20.056 6.47c-.645.8-1.5 1.978-2.26 3.343c-1.03 1.845-1.819 3.908-1.809 5.788c1.472 1.773 1.378 4.448-.28 6.106c-1.758 1.758-4.657 1.758-6.414 0c-1.758-1.757-1.758-4.657 0-6.414c1.283-1.283 3.176-1.63 4.788-1.038c.27-1.939 1.085-3.835 1.967-5.417a26 26 0 0 1 1.71-2.658a27 27 0 0 0-2.1.341c-2.172.445-4.572 1.237-6.296 2.604c1.065 1.745.847 4.08-.655 5.582c-1.758 1.758-4.657 1.758-6.414 0c-1.758-1.757-1.758-4.657 0-6.414c1.525-1.526 3.912-1.727 5.664-.603c2.124-1.756 4.988-2.655 7.3-3.128a29 29 0 0 1 2.495-.396L15.586 2zM7.293 9.707c-.977-.976-2.61-.976-3.586 0c-.976.977-.976 2.61 0 3.586s2.61.976 3.586 0s.976-2.61 0-3.586m7 7c-.977-.976-2.61-.976-3.586 0c-.976.977-.976 2.61 0 3.586s2.61.976 3.586 0s.976-2.61 0-3.586"/>`), g.Group(children),
	)
}

func ChevronDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.5 8.086l-5.5 5.5l-5.5-5.5L5.086 9.5L12 16.414L18.914 9.5z"/>`), g.Group(children),
	)
}

func ChevronDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m7-2.914l4 4l4-4l1.414 1.414L12 15.914L6.586 10.5z"/>`), g.Group(children),
	)
}

func ChevronDownDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.5 5.586l-4.5 4.5l-4.5-4.5L6.086 7L12 12.914L17.914 7zm0 6.5l-4.5 4.5l-4.5-4.5L6.086 13.5L12 19.414l5.914-5.914z"/>`), g.Group(children),
	)
}

func ChevronDownDoubleS(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 17.909l-4.95-4.95l1.415-1.414l3.536 3.535l3.535-3.535l1.414 1.414zm0-5.461l-4.95-4.95l1.415-1.414l3.536 3.535l3.535-3.535l1.414 1.414z"/>`), g.Group(children),
	)
}

func ChevronDownRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm4 5.086l4 4l4-4l1.414 1.414L12 15.914L6.586 10.5z"/>`), g.Group(children),
	)
}

func ChevronDownS(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 15.684l-4.95-4.95L8.464 9.32L12 12.856l3.535-3.536l1.414 1.414z"/>`), g.Group(children),
	)
}

func ChevronLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15.914 17.5l-5.5-5.5l5.5-5.5L14.5 5.086L7.586 12l6.914 6.914z"/>`), g.Group(children),
	)
}

func ChevronLeftCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m13.914-4l-4 4l4 4l-1.414 1.414L8.086 12L13.5 6.586z"/>`), g.Group(children),
	)
}

func ChevronLeftDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m18.414 7.5l-4.5 4.5l4.5 4.5L17 17.914L11.086 12L17 6.086zm-6.5 0l-4.5 4.5l4.5 4.5l-1.414 1.414L4.586 12L10.5 6.086z"/>`), g.Group(children),
	)
}

func ChevronLeftDoubleS(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6.088 11.498l4.95-4.95l1.414 1.415l-3.536 3.535l3.536 3.536l-1.414 1.414zm5.46 0l4.95-4.95l1.415 1.415l-3.536 3.535l3.536 3.536l-1.414 1.414z"/>`), g.Group(children),
	)
}

func ChevronLeftRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm10.914 4l-4 4l4 4l-1.414 1.414L8.086 12L13.5 6.586z"/>`), g.Group(children),
	)
}

func ChevronLeftS(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m8.315 12l4.95-4.95l1.414 1.415L11.144 12l3.535 3.536l-1.414 1.414z"/>`), g.Group(children),
	)
}

func ChevronRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m8.086 17.5l5.5-5.5l-5.5-5.5L9.5 5.086L16.414 12L9.5 18.914z"/>`), g.Group(children),
	)
}

func ChevronRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m9.5-5.414L15.914 12L10.5 17.414L9.086 16l4-4l-4-4z"/>`), g.Group(children),
	)
}

func ChevronRightDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m5.586 16.5l4.5-4.5l-4.5-4.5L7 6.086L12.914 12L7 17.914zm6.5 0l4.5-4.5l-4.5-4.5L13.5 6.086L19.414 12L13.5 17.914z"/>`), g.Group(children),
	)
}

func ChevronRightDoubleS(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.912 11.498l-4.95 4.95l-1.414-1.414l3.535-3.536l-3.535-3.535l1.414-1.415zm-5.461 0l-4.95 4.95l-1.414-1.414l3.535-3.536l-3.535-3.535L7.5 6.548z"/>`), g.Group(children),
	)
}

func ChevronRightRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm6.5 2.586L15.914 12L10.5 17.414L9.086 16l4-4l-4-4z"/>`), g.Group(children),
	)
}

func ChevronRightS(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15.684 12l-4.95 4.95l-1.414-1.414L12.856 12L9.32 8.465l1.415-1.415z"/>`), g.Group(children),
	)
}

func ChevronUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.5 15.914l-5.5-5.5l-5.5 5.5L5.086 14.5L12 7.586l6.914 6.914z"/>`), g.Group(children),
	)
}

func ChevronUpCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m11-3.914l5.414 5.414L16 14.914l-4-4l-4 4L6.586 13.5z"/>`), g.Group(children),
	)
}

func ChevronUpDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.5 18.414l-4.5-4.5l-4.5 4.5L6.086 17L12 11.086L17.914 17zm0-6.5l-4.5-4.5l-4.5 4.5L6.086 10.5L12 4.586l5.914 5.914z"/>`), g.Group(children),
	)
}

func ChevronUpDoubleS(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 6.084l4.95 4.95l-1.414 1.414l-3.535-3.536l-3.536 3.536l-1.414-1.414zm0 5.46l4.95 4.95l-1.414 1.415l-3.535-3.536l-3.536 3.536l-1.414-1.415z"/>`), g.Group(children),
	)
}

func ChevronUpRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm8 4.086l5.414 5.414L16 14.914l-4-4l-4 4L6.586 13.5z"/>`), g.Group(children),
	)
}

func ChevronUpS(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 8.316l4.95 4.95l-1.415 1.414L12 11.144L8.464 14.68L7.05 13.265z"/>`), g.Group(children),
	)
}

func Chicken(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.032 3.38a4 4 0 0 1 .473.01a2.65 2.65 0 0 1 1.25.416c.446.299.782.757.931 1.367c.144.59.083 1.15-.178 1.642c-.247.467-.63.789-.99 1.012c-.529.327-1.167.531-1.654.667l-1.072 1.855A5.99 5.99 0 0 1 21 14.999a5.98 5.98 0 0 1-1.528 4H22v2H2v-2h2.708a8 8 0 0 1 5.293-14H13v2h-2.999a6 6 0 0 0-.001 12h5a4 4 0 1 0-.43-7.977A3.973 3.973 0 0 0 11 14.999v1H9v-1a5.976 5.976 0 0 1 4.945-5.91l1.738-3.009c-.127-.49-.27-1.144-.251-1.765c.012-.424.1-.916.381-1.363c.296-.47.751-.803 1.331-.975c.604-.178 1.17-.12 1.654.118c.458.224.776.573.988.877a4 4 0 0 1 .245.408m-3.798 5.746q.42.088.82.234l1.493-2.583l.4-.107c.57-.153 1.129-.303 1.519-.544c.181-.112.25-.2.275-.247c.011-.021.04-.076.002-.23c-.033-.135-.078-.166-.101-.181a.7.7 0 0 0-.311-.085a2 2 0 0 0-.668.06l-.99.283l-.248-1.002v-.002l-.01-.035a2.2 2.2 0 0 0-.27-.571a.7.7 0 0 0-.226-.225c-.024-.011-.073-.036-.207.004c-.158.047-.192.1-.206.122c-.027.044-.069.147-.075.358c-.014.452.134 1.003.287 1.576l.11.412z"/>`), g.Group(children),
	)
}

func Chili(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.701 2.129L19.226 4.1a5.5 5.5 0 0 1 2.76 5.714q-.088.522-.18 1.032c-.38 2.076-1.605 6.033-4.1 8.527c-2.72 2.72-6.304 3.235-9.076 3.12a19.4 19.4 0 0 1-4.558-.755a11 11 0 0 1-.372-.117l-.023-.008l-.007-.003l-.004-.001l-3.19-1.128l3.285-.786h.001l.005-.002l.04-.011q.061-.016.191-.06c.173-.059.43-.155.746-.304a9.7 9.7 0 0 0 2.379-1.6c1.748-1.586 3.66-4.426 3.99-9.447l.01-.155v-.01c.215-3.16 3.337-5.04 6.199-4.64l.435-1.807zM7.008 20.335c.521.076 1.096.135 1.704.16c2.51.104 5.425-.38 7.58-2.535c1.897-1.897 3-4.978 3.442-6.961l-.862-.85l-1.655 1.104l-1-1.5l-1.5 1l-1.164-1.745l-.504.126c-.495 5.098-2.542 8.214-4.581 10.067c-.5.453-.995.827-1.46 1.134m6.549-13.39l.879-.22l.836 1.255l1.5-1l1 1.5l1.345-.897l.918.904a3.505 3.505 0 0 0-2.901-3.026c-1.473-.245-2.884.413-3.577 1.484"/>`), g.Group(children),
	)
}

func Chimney(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.719 2H13.28l-1.27 5.08l.358 4.118L15 9.882l7 3.5V22H1.927L2.99 7.088zM4.93 8l-.857 12H8v-6.618l2.444-1.223L10.083 8zm5.288-2l.5-2H4.28l.5 2zm-.22 14h4v-3h2v3h4v-5.382l-5-2.5l-5 2.5z"/>`), g.Group(children),
	)
}

func ChimneyOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.28 2H1.72l1.27 5.088L1.926 22H22V10h-9.661l-.326-2.932zm-3.175 6l.222 2H8v10H4.074L4.93 8zM4.78 6l-.5-2h6.439l-.5 2zM10 20v-8h10v8h-4v-4h-2v4z"/>`), g.Group(children),
	)
}

func ChimneyTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 2H4v5.924L1.834 22H22V10H11.32L11 7.924zM9 4v4.076L9.296 10H8v10H4.165L6 8.076V4zm1 16v-8h10v8h-4v-4h-2v4z"/>`), g.Group(children),
	)
}

func ChineseCabbage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.412 3.71C10.11 4.127 10 4.66 10 5v.52l-.427.3c-.366.256-.77.71-1.007 1.333l-.033.086c-.229.6-.407 1.067-.402 1.593c.004.484.174 1.125.952 1.964c.45.428.766.644 1.146.896l.06.04c.165.108.342.225.54.365L11 11.76v-1.346L8.586 8L10 6.586l1 1V5h2v2.586l1-1L15.414 8L13 10.414v1.345l.172.338c.197-.14.374-.257.538-.366l.06-.04a7.4 7.4 0 0 0 1.147-.895c.778-.839.948-1.48.952-1.964c.005-.526-.173-.994-.402-1.593l-.033-.086a2.9 2.9 0 0 0-1.007-1.334L14 5.521V5c0-.34-.11-.872-.412-1.29C13.324 3.341 12.874 3 12 3s-1.324.342-1.588.71m5.86 8.557a9.3 9.3 0 0 1-1.455 1.13q-.183.12-.376.253c.392-.11.802-.24 1.218-.391l.371-.135c1.344-.485 2.557-.924 3.486-1.552C20.46 10.932 21 10.172 21 9c0-.905-.195-1.39-.386-1.645c-.177-.234-.425-.375-.794-.43c-.614-.09-1.419.087-2.198.396c.117.377.217.809.242 1.277l1.344-.997l1.19 1.607zm.547-6.78c.992-.4 2.192-.702 3.292-.54c.779.115 1.547.47 2.101 1.207c.54.717.788 1.68.788 2.846c0 2.006-1.026 3.324-2.364 4.229c-.916.62-2.03 1.079-3.096 1.475c.276.425.546.864.785 1.302c.439.803.831 1.7.918 2.569c.094.933-.17 1.892-1.084 2.558c-.742.541-1.801.801-3.137.856C14.325 22.629 13.319 23 12 23c-1.32 0-2.326-.37-3.022-1.011c-1.336-.055-2.395-.315-3.137-.856c-.914-.666-1.178-1.625-1.084-2.558c.087-.869.48-1.766.918-2.569c.24-.438.51-.877.785-1.303c-1.067-.395-2.18-.854-3.096-1.474C2.026 12.324 1 11.006 1 9c0-1.167.249-2.13.788-2.846c.554-.738 1.322-1.092 2.1-1.207c1.101-.162 2.3.14 3.293.54a5 5 0 0 1 .853-.996a4.3 4.3 0 0 1 .754-1.95C9.424 1.658 10.474 1 12 1s2.576.658 3.211 1.54a4.3 4.3 0 0 1 .755 1.951c.308.273.6.607.853.997M6.38 7.322c-.78-.309-1.585-.486-2.199-.395c-.369.054-.618.195-.794.43C3.195 7.61 3 8.094 3 9c0 1.172.54 1.933 1.484 2.572c.929.628 2.142 1.067 3.486 1.552l.371.135c.416.15.826.28 1.218.39a23 23 0 0 0-.435-.29a9.3 9.3 0 0 1-1.397-1.092l-4.126-3.06l1.191-1.606l1.344.997c.025-.468.125-.9.242-1.277m2.913 8.33a20 20 0 0 1-.89-.26c-.353.528-.69 1.058-.972 1.574c-.403.738-.637 1.35-.683 1.81c-.04.394.056.585.272.742c.167.122.446.25.89.343a3.9 3.9 0 0 1 .142-1.676l.022-.065zM12 14.32l-.319 1.01l-1.749 3.54a2 2 0 0 0-.045.688c.033.307.133.592.313.816c.226.282.705.625 1.8.625s1.574-.343 1.8-.625c.18-.224.28-.509.313-.816c.03-.279-.002-.53-.045-.689l-1.75-3.54zm4.091 5.539c.444-.093.723-.22.89-.343c.216-.157.312-.348.272-.743c-.046-.46-.28-1.07-.683-1.809a21 21 0 0 0-.971-1.574a20 20 0 0 1-.892.26l1.22 2.468l.022.065c.14.419.225 1.027.142 1.676"/>`), g.Group(children),
	)
}

func Church(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 1v2H9v2h2v.98l-5 3.5V12H2v10h20V12h-4V9.48l-5-3.5V5h2V3h-2V1zm1 6.72l4 2.8V20h-1v-4a3 3 0 1 0-6 0v4H8v-9.48zM6 20H4v-6h2zm5 0v-4a1 1 0 1 1 2 0v4zm7 0v-6h2v6z"/>`), g.Group(children),
	)
}

func Circle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12"/>`), g.Group(children),
	)
}

func City(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m2 1.5l8 3.333V10h12v12H2zM10 12v8h2v-5h6v5h2v-8zm6 8v-3h-2v3zm-8 0V6.167L4 4.5V20z"/>`), g.Group(children),
	)
}

func CityAncient(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.146 1.483l.856-1.417l.857 1.419l.005.01l.024.039l.097.154a36 36 0 0 0 1.59 2.333c.49.66 1.024 1.32 1.54 1.852c.544.559.97.87 1.233.967c.337.125.554.16.652.16h1v2h-1v2h4v2h-1v9H3v-9H2v-2h4.002V9H5V7h1c.098 0 .316-.035.654-.16c.264-.099.69-.41 1.233-.968A19 19 0 0 0 9.43 4.02a36 36 0 0 0 1.686-2.487l.024-.04l.005-.01zM8.002 9v2H16V9zm6.427-2a22 22 0 0 1-1.46-1.785c-.361-.486-.691-.96-.968-1.37c-.276.41-.606.883-.967 1.369A22 22 0 0 1 9.573 7zM5 13v7h3v-5h8v5h3v-7zm9 7v-3h-4v3z"/>`), g.Group(children),
	)
}

func CityAncientOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.52 4.915c.922-.936 1.482-1.96 1.482-2.915h2c0 .956.56 1.98 1.482 2.916c.91.923 2.037 1.617 2.864 1.924c.337.125.554.16.652.16h1v2h-1v2h4v2h-1v7h1v2H2v-2h1v-7H2v-2h4.002V9H5V7h1c.098 0 .316-.035.654-.16c.827-.308 1.955-1.002 2.865-1.925M8.001 9v2H16V9zM13.8 7q-.385-.32-.74-.68a8.7 8.7 0 0 1-1.058-1.291a8.7 8.7 0 0 1-1.058 1.29q-.355.36-.74.681zM5 13v7h6.002v-4h2v4H19v-7z"/>`), g.Group(children),
	)
}

func CityAncientTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.999.802l6.053 3.993L17 6.39V9.5l3.334 2.5h1.667v2h-1v8h-18v-8h-1v-2h1.667L7 9.5V6.39L5.948 4.794zm-3 4.375V9h6V5.177l-3-1.98zM15.667 11H8.333L7 12h10zm3.335 3h-14v6H11v-3h2v3h6zm-8-9.002h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func CityEight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h12v8h8v12H2zm10 8V8H9v2zm-3 2v2h3v-2zm-2-2V8H4v2zm-3 2v2h3v-2zm0 4v4h8v-4zm10 4h2v-4h2v4h2v-8h-6zM4 6h3V4H4zm5-2v2h3V4z"/>`), g.Group(children),
	)
}

func CityEleven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 2H2v20h20v-9.72l-7-2.334l-1 .333zm-2 8.946l-4 1.333V20H4V4h8zM10 20v-6.28l5-1.666l5 1.667V20h-4v-4h-2v4z"/>`), g.Group(children),
	)
}

func CityFifteen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 2H2v20h20V10.687L14 9.02zm-2 6.604L8 7.77V20H4V4h8zM10 20v-9.77l10 2.083V20h-4v-4h-2v4z"/>`), g.Group(children),
	)
}

func CityFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.5 1.72L15 4.52v5.108a4.124 4.124 0 0 1-6 0V4.519l-3.5-2.8L2 4.52V22h20V4.52zM9 12.14a6.15 6.15 0 0 0 6 0V20h-2v-5h-2v5H9zM7 20H4V5.48l1.5-1.2L7 5.48zm10 0V5.48l1.5-1.2l1.5 1.2V20z"/>`), g.Group(children),
	)
}

func CityFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.5 1.72L15 4.52v3.862l-3-1.5l-3 1.5V4.519l-3.5-2.8L2 4.52V22h20V4.52zM9 10.617l3-1.5l3 1.5V20h-2v-4h-2v4H9zM7 20H4V5.48l1.5-1.2L7 5.48zm10 0V5.48l1.5-1.2l1.5 1.2V20zm-3-8h-4v2h4z"/>`), g.Group(children),
	)
}

func CityFourteen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 1.923l-6.3 2.52l.743 1.857L7 6.077V8H1v14h22V8h-6V6.077l.557.223l.743-1.857zm3 3.354V20h-2v-5h-2v5H9V5.277l3-1.2zM7 20H3V10h4zm10 0V10h4v10zM13.004 6.998H11v2.004h2.004z"/>`), g.Group(children),
	)
}

func CityNine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m8.01 1.828l3.988 2.492l3.99-2.492l6.384 3.829l-1.029 1.715L21 7.166V22H2.997V7.166l-.343.206l-1.029-1.715zM4.996 5.967V20H7v-5h2v5h1.997V6.053l-3.01-1.88zM13 20h2v-5h2v5h2V5.967l-2.991-1.794l-3.01 1.88zM7 8.998h2.004v2.004H7zm8 0h2.004v2.004H15z"/>`), g.Group(children),
	)
}

func CityOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h10v8h10v12H2zm10 10v8h4v-5h4v-3zm8 5h-2v3h2zm-10 3V4H4v16zM8.004 6v2.004H6v-2h.004V6zm0 5v2.004H6v-2h.004V11zm0 5v2.004H6v-2h.004V16z"/>`), g.Group(children),
	)
}

func CitySeven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 2H2v20h20V10h-8zm-2 8H8v10H4V4h8zm-2 10v-8h10v8h-4v-4h-2v4z"/>`), g.Group(children),
	)
}

func CitySix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 2H6v6H1v14h22V6h-5zm0 6h3v12h-3zm-2 12h-3v-5h-2v5H8V4h8zM6 20H3V10h3z"/>`), g.Group(children),
	)
}

func CityTen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 .72V10h8v12H2V3.72zM14 20h2v-5h2v5h2v-8h-6zM4 5.28V20h8V3.28zm2 2.718h2.004v2.004H6z"/>`), g.Group(children),
	)
}

func CityThirteen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 2H6v2h1v4H1v14h22V8h-6V4h1zm-3 2v16h-2v-5h-2v5H9V4zM7 20H3V10h4zm10 0V10h4v10zM13.004 5.998H11v2.004h2.004z"/>`), g.Group(children),
	)
}

func CityThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 2h-7v5H9V2H2v20h20zM9 9h6v11h-2v-4h-2v4H9zM7 20H4V4h3zm10 0V4h3v16zm-3-9h-4v2h4z"/>`), g.Group(children),
	)
}

func CityTwelve(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 2H2v20h20V5h-8v5h-4zM8 12h8v8h-3v-5h-2v5H8zm-2 8H4V4h4v6H6zm12 0V10h-2V7h4v13z"/>`), g.Group(children),
	)
}

func CityTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6 .698l4 3.334v6.833l5-4.167l7 5.834V22H2V4.032zM10 20h2v-6h6v6h2v-6.532l-5-4.166l-5 4.166zm6 0v-4h-2v4zm-8 0V4.968L6 3.302L4 4.968V20z"/>`), g.Group(children),
	)
}

func Clear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 1h6v8.5h6V23H3V9.5h6zm2 2v8.5H5V14h14v-2.5h-6V3zm8 13H5v5h9v-3h2v3h3z"/>`), g.Group(children),
	)
}

func ClearFormatting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m14.031 1.888l9.657 9.657l-8.345 8.345l-.27.272H20v2H6.748L.253 15.667zm.322 16.164l6.507-6.507l-6.829-6.828l-6.828 6.828l6.828 6.828l.32-.323zM5.788 12.96l-2.707 2.707l4.495 4.495h4.68l.365-.37z"/>`), g.Group(children),
	)
}

func ClearFormattingOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m14.03 1.889l9.657 9.657l-8.345 8.345l-.27.27H20v2H6.747l-3.666-3.666a4 4 0 0 1 0-5.657zm.322 16.163l6.507-6.506l-6.829-6.829l-6.828 6.829l6.828 6.828l.32-.323zM5.788 12.96l-1.293 1.293a2 2 0 0 0 0 2.828l3.08 3.08h4.68l.366-.368z"/>`), g.Group(children),
	)
}

func Close(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7.05 5.636l4.95 4.95l4.95-4.95l1.414 1.414l-4.95 4.95l4.95 4.95l-1.415 1.414l-4.95-4.95l-4.949 4.95l-1.414-1.414l4.95-4.95l-4.95-4.95z"/>`), g.Group(children),
	)
}

func CloseCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11M7.403 15.182L10.586 12L7.403 8.818l1.415-1.415L12 10.586l3.182-3.182l1.414 1.414L13.414 12l3.182 3.182l-1.414 1.414L12 13.414l-3.182 3.182z"/>`), g.Group(children),
	)
}

func CloseCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11M8.818 7.403L12 10.586l3.182-3.182l1.414 1.414L13.414 12l3.182 3.182l-1.415 1.414L12 13.414l-3.182 3.182l-1.415-1.414L10.586 12L7.403 8.818z"/>`), g.Group(children),
	)
}

func CloseOctagon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.35 1.5l6.15 6.15v8.699l-6.15 6.15h-8.7L1.5 16.35v-8.7L7.65 1.5zm-.83 2H8.48L3.5 8.479v7.041l4.98 4.98h7.04l4.98-4.98V8.48zm1.076 5.318L13.414 12l3.182 3.181l-1.414 1.415L12 13.414l-3.182 3.182l-1.415-1.414L10.586 12L7.403 8.817l1.415-1.414L12 10.585l3.182-3.181z"/>`), g.Group(children),
	)
}

func CloseRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm5.172 3.757L12 10.586l2.828-2.829l1.415 1.415L13.414 12l2.829 2.828l-1.415 1.415L12 13.414l-2.828 2.829l-1.415-1.415L10.586 12L7.757 9.172z"/>`), g.Group(children),
	)
}

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 5a5 5 0 0 0-4.916 5.919l.175.942l-.934.215A3.001 3.001 0 0 0 6 18h11a4 4 0 1 0-.067-8l-.86.014l-.142-.848A5 5 0 0 0 11 5m-7 5a7 7 0 0 1 13.723-1.957A6.001 6.001 0 0 1 17 20H6a5 5 0 0 1-1.988-9.589A7 7 0 0 1 4 10"/>`), g.Group(children),
	)
}

func CloudDownload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.598 6.283a7.502 7.502 0 0 1 14.804 0a6.502 6.502 0 0 1 1.053 12.008l-.89.455l-.91-1.78l.89-.456a4.502 4.502 0 0 0-1.236-8.438l-.771-.14l-.049-.781a5.5 5.5 0 0 0-10.978 0l-.049.782l-.77.14a4.502 4.502 0 0 0-1.237 8.437l.89.455l-.91 1.78l-.89-.454A6.502 6.502 0 0 1 4.599 6.283M13 10v9.586l3-3L17.414 18L12 23.414L6.586 18L8 16.586l3 3V10z"/>`), g.Group(children),
	)
}

func CloudUpload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.598 6.283a7.502 7.502 0 0 1 14.804 0a6.502 6.502 0 0 1 1.053 12.008l-.89.455l-.91-1.78l.89-.456a4.502 4.502 0 0 0-1.236-8.438l-.771-.14l-.049-.781a5.5 5.5 0 0 0-10.978 0l-.049.782l-.77.14a4.502 4.502 0 0 0-1.237 8.437l.89.455l-.91 1.78l-.89-.454A6.502 6.502 0 0 1 4.599 6.283M12 9.586L17.414 15L16 16.414l-3-3V23h-2v-9.586l-3 3L6.586 15z"/>`), g.Group(children),
	)
}

func CloudyDay(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.414 6.005A5 5 0 0 1 18.584 6A4.5 4.5 0 0 1 23 10.5c0 1.216-.617 2.33-1.387 3.113a5 5 0 0 1-1.69 1.139q.076.446.077.915C20 18.594 17.668 21 14.75 21H5.4C2.952 21 1 18.982 1 16.533a4.47 4.47 0 0 1 2.556-4.056a6 6 0 0 1-.006-.277c0-3.326 2.588-6.07 5.864-6.195m2.114.294c1.895.622 3.376 2.158 3.956 4.086a5.25 5.25 0 0 1 3.74 2.492a2.9 2.9 0 0 0 .963-.667c.518-.527.813-1.163.813-1.71a2.5 2.5 0 0 0-2.939-2.462l-.924.163l-.221-.912a3.002 3.002 0 0 0-5.388-.99M9.65 8c-2.246 0-4.1 1.862-4.1 4.2q0 .4.07.776l.17.935l-.926.217C3.808 14.375 3 15.348 3 16.533C3 17.914 4.093 19 5.4 19h9.35c1.777 0 3.25-1.474 3.25-3.333c0-.46-.09-.894-.251-1.288c-.496-1.212-1.66-2.046-2.999-2.046h-.054l-.863.015l-.14-.852c-.308-1.872-1.82-3.306-3.665-3.478A4 4 0 0 0 9.65 8"/>`), g.Group(children),
	)
}

func CloudyNight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.133.92l-.102 1.186q-.017.194-.017.393a4.506 4.506 0 0 0 4.9 4.49l1.185-.103l-.103 1.185a6.5 6.5 0 0 1-2.162 4.3a5.2 5.2 0 0 1 1.186 3.31c0 2.983-2.501 5.338-5.505 5.338H6.604C4.097 21.019 2 19.052 2 16.549c0-1.833 1.126-3.38 2.709-4.076a6 6 0 0 1-.006-.263c0-3.147 2.398-5.7 5.464-6.139a6.51 6.51 0 0 1 5.782-5.048zM12.232 6.1q.235.04.464.097c2.15.53 3.883 2.121 4.538 4.19c.723.092 1.404.32 2.013.658l.032.018a4.5 4.5 0 0 0 1.5-2.1a6.51 6.51 0 0 1-5.723-5.723a4.52 4.52 0 0 0-2.824 2.86m-1.123 1.904h-.064c-2.44.034-4.342 1.938-4.342 4.206q0 .395.073.766l.184.956l-.95.21C4.833 14.402 4 15.403 4 16.548c0 1.33 1.13 2.47 2.604 2.47h9.91c1.972 0 3.506-1.53 3.506-3.338c0-1.052-.512-2-1.333-2.62a3.6 3.6 0 0 0-2.231-.718l-.855.014l-.146-.841c-.283-1.624-1.553-2.96-3.239-3.377a4.6 4.6 0 0 0-1.107-.134"/>`), g.Group(children),
	)
}

func CloudyNightRain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.131.901l-.102 1.185q-.017.194-.017.394a4.505 4.505 0 0 0 4.899 4.488l1.185-.102l-.103 1.184a6.5 6.5 0 0 1-2.162 4.3a5.2 5.2 0 0 1 1.186 3.309c0 2.212-1.383 4.085-3.316 4.898l-.921.388l-.776-1.844l.922-.387c1.25-.526 2.09-1.71 2.09-3.055c0-1.052-.51-2-1.332-2.62a3.6 3.6 0 0 0-2.172-.718h-.059l-.854.014l-.146-.842C15.17 9.87 13.9 8.535 12.215 8.118a4.6 4.6 0 0 0-1.172-.134c-2.438.034-4.34 1.938-4.34 4.205q0 .393.072.766l.184.956l-.95.21C4.833 14.38 4 15.38 4 16.526c0 .948.567 1.793 1.434 2.207l.902.432l-.864 1.804l-.902-.431C3.064 19.817 2 18.303 2 16.526c0-1.831 1.126-3.378 2.708-4.074a6 6 0 0 1-.005-.263c0-3.146 2.398-5.699 5.463-6.138a6.51 6.51 0 0 1 5.78-5.047zm-4.9 5.178q.234.041.463.097c2.15.531 3.883 2.122 4.538 4.19a5.6 5.6 0 0 1 2.045.676a4.5 4.5 0 0 0 1.5-2.099a6.51 6.51 0 0 1-5.723-5.722a4.52 4.52 0 0 0-2.823 2.858m.765 7.897H15v2.003h-2.004zm-4.996 2h2.004v2.003H8zm5 3h2.004v2.003H13zm-5 2h2.004v2.003H8z"/>`), g.Group(children),
	)
}

func CloudyRain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.004 1v2.004H15V1zm-5.952 1.635l1.418 1.417l-1.417 1.417l-1.417-1.417zm9.9 0l1.415 1.416l-1.415 1.417l-1.418-1.416zm-8.993 3.42A7 7 0 0 0 11.1 6c-3.499 0-6.4 2.74-6.4 6.2q0 .132.006.263C3.125 13.158 2 14.703 2 16.533c0 1.776 1.063 3.288 2.568 4.009l.902.431l.864-1.804l-.902-.431C4.566 18.323 4 17.48 4 16.534c0-1.143.832-2.143 2.007-2.403l.95-.21l-.184-.955A4 4 0 0 1 6.7 12.2C6.7 9.916 8.634 8 11.1 8q.576.002 1.106.134c1.683.416 2.952 1.75 3.234 3.372l.147.841l.854-.013h.059a3.6 3.6 0 0 1 1.759.45C19.314 13.37 20 14.451 20 15.667c0 1.344-.84 2.526-2.088 3.051l-.922.388l.775 1.843l.922-.387C20.618 19.749 22 17.877 22 15.667c0-1.575-.702-2.979-1.803-3.949a5 5 0 0 0-8.238-5.663m2.108.65a3 3 0 0 1 4.413 3.984a5.6 5.6 0 0 0-1.26-.31c-.507-1.601-1.66-2.916-3.153-3.673M22 8h2.004v2.004H22zm-9.002 5.998h2.004v2.004h-2.004zm-5 2h2.004v2.004H7.998zm5 3h2.004v2.004h-2.004zm-5 2h2.004v2.004H7.998z"/>`), g.Group(children),
	)
}

func CloudySunny(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.004 1v2.004H15V1zm-5.952 1.635l1.418 1.417l-1.417 1.417l-1.417-1.417zm9.9 0l1.415 1.416l-1.415 1.417l-1.418-1.416zm-8.993 3.42a5 5 0 0 1 8.238 5.663c1.101.97 1.803 2.374 1.803 3.949C22 18.647 19.502 21 16.5 21H6.6C4.095 21 2 19.036 2 16.534c0-1.83 1.125-3.376 2.706-4.071A6 6 0 0 1 4.7 12.2C4.7 8.74 7.601 6 11.1 6q.436 0 .859.056m2.108.65c1.492.758 2.646 2.073 3.152 3.674q.659.084 1.26.31a3 3 0 0 0-4.413-3.984M11.1 8c-2.466 0-4.4 1.916-4.4 4.2q0 .394.073.766l.184.955l-.95.21C4.832 14.391 4 15.391 4 16.534C4 17.86 5.128 19 6.6 19h9.9c1.969 0 3.5-1.527 3.5-3.333c0-1.216-.686-2.297-1.741-2.883a3.6 3.6 0 0 0-1.818-.45l-.854.013l-.147-.841c-.282-1.622-1.55-2.956-3.234-3.372A4.6 4.6 0 0 0 11.1 8M22 8h2.004v2.004H22z"/>`), g.Group(children),
	)
}

func Code(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m.586 12l4.95-4.95L6.95 8.464L3.414 12l3.536 3.536l-1.414 1.414zm8.201 8.728l4.486-17.94l1.94.485l-4.485 17.94zm8.263-5.192L20.586 12L17.05 8.464l1.415-1.414l4.95 4.95l-4.95 4.95z"/>`), g.Group(children),
	)
}

func CodeOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m23.414 12l-4.95-4.95l-1.414 1.414L20.586 12l-2.5 2.501l1.415 1.414zm-2 8L4 2.587L2.586 4.001l3 3l-5 5l4.95 4.949l1.414-1.414L3.414 12L7 8.415l13 13z"/>`), g.Group(children),
	)
}

func CodeOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m.586 12l4.95-4.95L6.95 8.464L3.414 12l3.536 3.535l-1.414 1.414zm16.464 3.535L20.586 12L17.05 8.464l1.415-1.414l4.95 4.95l-4.95 4.95z"/>`), g.Group(children),
	)
}

func Cola(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.307.893l5.8 5.8l-1.204 4.818l-3.173 3.174q-.2.225-.47.47l-6.749 6.748l-4.818 1.205l-5.8-5.801l1.204-4.818L12.49 2.097zm-.426 12.812L19.586 11l-1.53-1.53c-2.105 2.3-5.903 5.192-11.633 5.538l1.041 1.042c4.912.56 8.137-1.213 9.417-2.345m-7.313 4.45L11 19.584l1.805-1.804c-.978.218-2.059.355-3.237.373m-.524 2.303l-5.502-5.502l-.434 1.737l4.199 4.2zm-4.607-7.435c6.198.288 10.145-2.691 12.204-4.968L13 4.415L4.414 13zM14.956 3.54l5.502 5.502l.435-1.737l-4.2-4.2z"/>`), g.Group(children),
	)
}

func Collage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h4.198l3.555-16zm9.802 0l-1.641 7.387L20 14.523V4zM20 16.677l-8.279-3.312L10.247 20H20z"/>`), g.Group(children),
	)
}

func Collection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 4h1V2H5v2zm3 3.5H3v-2h18zM23 9v13H1V9zm-2 2H3v9h18z"/>`), g.Group(children),
	)
}

func ColorInvert(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 1.586l6.01 6.01a8.5 8.5 0 0 1 1.906 9.114a7.5 7.5 0 0 1-1.918 2.92a8.5 8.5 0 0 1-2.63 1.784A8.503 8.503 0 0 1 5.99 7.595zm3.157 17.704c.453-.258.927-.612 1.451-1.1a6.5 6.5 0 0 0 1.445-2.21a6.7 6.7 0 0 0 .381-2.387c-.032-1.742-.695-3.44-1.838-4.583L13 5.414v14.542a5.5 5.5 0 0 0 1.548-.368q.31-.132.61-.298M11 5.414L7.404 9.01A6.5 6.5 0 0 0 11 20.03z"/>`), g.Group(children),
	)
}

func Combination(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 1.148l6.16 9.602H5.84zm0 3.704L9.5 8.75h5zM2 13h9v9H2zm2 2v5h5v-5zm13.5 0a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5M13 17.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0"/>`), g.Group(children),
	)
}

func Command(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.5 4.5a2 2 0 0 1 2 2v2h-2a2 2 0 1 1 0-4m4 4v-2a4 4 0 1 0-4 4h2v3h-2a4 4 0 1 0 4 4v-2h3v2a4 4 0 1 0 4-4h-2v-3h2a4 4 0 1 0-4-4v2zm0 2h3v3h-3zm5-2v-2a2 2 0 1 1 2 2zm0 7h2a2 2 0 1 1-2 2zm-7 0v2a2 2 0 1 1-2-2z"/>`), g.Group(children),
	)
}

func Compass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21 12a9 9 0 1 0-18 0a9 9 0 0 0 18 0M12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1m5.435 5.565l-2.64 8.23l-8.23 2.64l2.64-8.23zm-6.64 4.23l-1.138 3.548l3.547-1.138l1.138-3.548z"/>`), g.Group(children),
	)
}

func CompassOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m5.891 12l4.341 1.155A2.1 2.1 0 0 1 9.89 12c0-.426.126-.823.343-1.155zm4.954-1.768A2.1 2.1 0 0 1 12 9.89c.426 0 .823.126 1.155.343L12 5.891zm2.481-7.135l1.592 5.985l5.985 1.592a9.01 9.01 0 0 0-7.576-7.577M18.11 12l-4.341-1.155c.217.332.343.729.343 1.155s-.126.823-.343 1.155zm-4.955 1.768A2.1 2.1 0 0 1 12 14.11c-.426 0-.823-.126-1.155-.343L12 18.109zm-2.48 7.135l-1.592-5.985l-5.985-1.591a9.01 9.01 0 0 0 7.577 7.576m-7.577-10.23l5.985-1.591l1.592-5.985a9.01 9.01 0 0 0-7.577 7.577m10.23 10.23a9.01 9.01 0 0 0 7.576-7.576l-5.985 1.591zM1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m11-.111a.111.111 0 1 0 0 .222a.111.111 0 0 0 0-.222"/>`), g.Group(children),
	)
}

func ComponentBreadcrumb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 6.5v11h7.914l5.5-5.5l-5.5-5.5zm2 2h5.086l3.5 3.5l-3.5 3.5H4zm10.586-1l4.5 4.5l-4.5 4.5L16 17.914L21.914 12L16 6.086z"/>`), g.Group(children),
	)
}

func ComponentCheckbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h13v5.5h-2V4H4v9h3.5v2H2zm7 7h13v13H9zm2 2v9h9v-9zm8.414 3L15 18.414L12.086 15.5l1.414-1.414l1.5 1.5l3-3z"/>`), g.Group(children),
	)
}

func ComponentDividerHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 2v20h-2V2zM2 5h7v14H2zm2 2v10h3V7zm11-2h7v14h-7zm2 2v10h3V7z"/>`), g.Group(children),
	)
}

func ComponentDividerVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 2h14v7H5zm2 2v3h10V4zm-5 7h20v2H2zm3 4h14v7H5zm2 2v3h10v-3z"/>`), g.Group(children),
	)
}

func ComponentDropdown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v7H2zm2 2v3h16V4zm0 7v9h16v-9h2v11H2V11zm2 1h12v2H6zm0 4h12v2H6z"/>`), g.Group(children),
	)
}

func ComponentGrid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v4h4V4zm6 0v4h4V4zm6 0v4h4V4zm4 6h-4v4h4zm0 6h-4v4h4zm-6 4v-4h-4v4zm-6 0v-4H4v4zm-4-6h4v-4H4zm6-4v4h4v-4z"/>`), g.Group(children),
	)
}

func ComponentInput(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 7h20v9H2zm2 2v5h16V9zm4 1v3H6v-3z"/>`), g.Group(children),
	)
}

func ComponentLayout(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v7h7V4zm9 0v16h7V4zm-2 16v-7H4v7z"/>`), g.Group(children),
	)
}

func ComponentRadio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4a8 8 0 1 0 0 16a8 8 0 0 0 0-16M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12m10-4a4 4 0 1 0 0 8a4 4 0 0 0 0-8m-6 4a6 6 0 1 1 12 0a6 6 0 0 1-12 0"/>`), g.Group(children),
	)
}

func ComponentSpace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 3h20v8H2zm2 2v4h16V5zm-2 8h20v8H2zm2 2v4h16v-4z"/>`), g.Group(children),
	)
}

func ComponentSteps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 12a3.5 3.5 0 0 1 6.855-1h.79a3.502 3.502 0 0 1 6.71 0h.79A3.502 3.502 0 0 1 23 12a3.5 3.5 0 0 1-6.855 1h-.79a3.502 3.502 0 0 1-6.71 0h-.79A3.502 3.502 0 0 1 1 12m3.5-1.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m7.5 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m7.5 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/>`), g.Group(children),
	)
}

func ComponentSwitch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 8.5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7m4.243 0A5.48 5.48 0 0 1 13 12c0 1.33-.472 2.55-1.257 3.5H16.5a3.5 3.5 0 1 0 0-7zM2 12a5.5 5.5 0 0 1 5.5-5.5h9a5.5 5.5 0 1 1 0 11h-9A5.5 5.5 0 0 1 2 12"/>`), g.Group(children),
	)
}

func Constraint(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 8H6.2C3.83 8.026 2 9.851 2 12s1.83 3.975 4.2 4H10v2H6.186C2.809 17.967 0 15.338 0 12s2.81-5.967 6.186-6H10zm4-2h3.75C21.155 6 24 8.641 24 12s-2.845 6-6.25 6H14v-2h3.75c2.394 0 4.25-1.836 4.25-4s-1.856-4-4.25-4H14zm-7 5h10v2H7z"/>`), g.Group(children),
	)
}

func Contrast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m10-7.5h1a7.5 7.5 0 0 1 0 15h-1zm2 2.09v10.82a5.502 5.502 0 0 0 0-10.82"/>`), g.Group(children),
	)
}

func ContrastOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 3.055a9.001 9.001 0 0 0 0 17.89zm2 0V4.84l2.184-1.26A9 9 0 0 0 13 3.054m4.255 1.638L13 7.149v4.042l6.693-3.864a9.05 9.05 0 0 0-2.438-2.634m3.284 4.455L13 13.5v4.042l7.95-4.59a9.1 9.1 0 0 0-.411-3.804m-.326 6.539L13 19.85v1.094a9.01 9.01 0 0 0 7.213-5.258M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12"/>`), g.Group(children),
	)
}

func ControlPlatform(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .856l10 5.556v11.176l-10 5.556l-10-5.556V6.412zm-8 8.04v7.515l7 3.89v-7.699zM13 20.3l7-3.889V8.897l-7 3.706zm-1-9.432L19.12 7.1L12 3.144L4.88 7.099z"/>`), g.Group(children),
	)
}

func Cooperate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.586l3 3l3-3l8.914 8.914L12 22.414L.086 10.5zM10.586 6L9 4.414L2.914 10.5L12 19.586l.954-.954l-2.368-2.369L12 14.85l2.368 2.369l.955-.955l-2.369-2.368l1.414-1.414l2.369 2.368l1.349-1.349L15 10.414l-3 3L7.586 9zm8.914 6.086l1.586-1.586L15 4.414L10.414 9L12 10.586l3-3z"/>`), g.Group(children),
	)
}

func CoordinateSystem(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 1l11.258 19.5H.74zM4.204 18.5h15.589L11.999 5zM13 9v3.532l3.408 2.84l-1.28 1.536l-3.129-2.606l-3.128 2.606l-1.28-1.536L11 12.532V9z"/>`), g.Group(children),
	)
}

func Copy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h13v5.5h-2V4H4v9h3.5v2H2zm7 7h13v13H9zm2 2v9h9v-9z"/>`), g.Group(children),
	)
}

func Copyright(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11M9.525 9.526a3.5 3.5 0 0 0 4.95 4.95l.707-.708l1.414 1.415l-.707.707a5.5 5.5 0 1 1 0-7.778l.707.707l-1.414 1.414l-.707-.707a3.5 3.5 0 0 0-4.95 0"/>`), g.Group(children),
	)
}

func Corn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.525 4.11l.979.98l2.086-2.085c-.37-.062-.785-.006-1.257.152c-.61.205-1.232.559-1.808.953m4.473.316l-2.08 2.078l.972.972c.394-.575.747-1.194.951-1.802c.158-.469.215-.88.157-1.248m-2.336 4.65l-1.159-1.159L15.93 9.49l1.338 1.338l.989-1.243zm-2.65 3.326l-1.497-1.499l-2.06 2.058c.996-.053 1.994.04 2.902.263zm1.31 1.566a9 9 0 0 1 1.01.614a8 8 0 0 1 .8.64l.05.046l.016.016l.005.005l.002.002h.001l.001.002l1.374 1.374l-1.904.317h-.003l-.011.003a2 2 0 0 0-.146.042a6 6 0 0 0-.725.286c-.678.31-1.742.897-3.192 1.985c-.15.112-.325.25-.523.404c-1.004.786-2.576 2.017-4.257 2.687c-1.029.41-2.196.656-3.381.423c-1.173-.23-2.247-.905-3.154-2.1c-1.194-.906-1.869-1.98-2.1-3.153c-.232-1.186.013-2.352.424-3.381c.67-1.68 1.9-3.254 2.686-4.257c.155-.198.293-.373.405-.523c1.088-1.45 1.674-2.515 1.985-3.192c.155-.338.24-.58.286-.725a2 2 0 0 0 .045-.16l.317-1.904l1.374 1.374l.002.001l.002.002l.005.006l.015.015a4 4 0 0 1 .2.221a9 9 0 0 1 1.091 1.622l4.356-3.443c.825-.66 2.026-1.523 3.32-1.956c1.328-.445 2.933-.48 4.23.825c1.292 1.3 1.255 2.9.81 4.225c-.435 1.291-1.297 2.489-1.956 3.313h-.002l-.002.005l-.957 1.202l-1.688 2.12zm-2.4-8.632l-1.757 1.387l1.35 1.352l1.574-1.572zM13.1 9.489l-1.515-1.517l-.816.646c.227.917.323 1.925.268 2.932zm-3.474 6.298L5.36 20.053c.495.476.99.706 1.463.798c.677.133 1.435.01 2.256-.318c1.382-.55 2.667-1.552 3.685-2.344q.34-.265.635-.49c1.242-.931 2.251-1.542 3.026-1.942a7 7 0 0 0-1.02-.44c-1.822-.61-4.228-.477-5.779.47m-5.68 2.852l4.266-4.266c.95-1.557 1.081-3.975.462-5.8a7 7 0 0 0-.432-1c-.4.776-1.011 1.784-1.943 3.027q-.223.297-.489.635c-.793 1.018-1.793 2.302-2.345 3.685c-.327.82-.45 1.58-.318 2.256c.093.472.323.968.799 1.463"/>`), g.Group(children),
	)
}

func Coupon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 21H3V1.132l1.555 1.036l2.391 1.594l1.93-1.543l.624-.5l.625.5L12 3.72l1.875-1.5l.625-.5l.625.5l1.929 1.543l2.391-1.594L21 1.132V21zM5 4.87V19h14V4.87l-1.445.963l-.609.406l-.57-.457l-1.876-1.5l-1.875 1.5l-.625.5l-.625-.5l-1.875-1.5l-1.875 1.5l-.571.457l-.609-.406zM8 16H7v-2h10v2zm3-4h6v-2h-6zm-3-2H7v2.004h2.004V10z"/>`), g.Group(children),
	)
}

func Course(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 3h-2v6.5l-3-2.25l-3 2.25V3H5v18h14zm-6 0v2.5l1-.75l1 .75V3zm8 20H3V1h18z"/>`), g.Group(children),
	)
}

func Cpu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 .5V3h6V.5h2V3h4v4h2.5v2H21v6h2.5v2H21v4h-4v2.5h-2V21H9v2.5H7V21H3v-4H.5v-2H3V9H.5V7H3V3h4V.5zM5 5v14h14V5zm3 3h8v8H8zm2 2v4h4v-4z"/>`), g.Group(children),
	)
}

func Crack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.388 3.02a9 9 0 0 0-.958 17.844l1.45-2.865A5 5 0 0 1 7.67 15.5l-.501-.866l1.731-1l.5.865c.31.536.78.967 1.345 1.227L9.882 14l2-4l-2.001-4.003zm2.182.116l-1.45 2.867L14.117 10l-2 4l.91 1.82a3 3 0 0 0 1.571-1.32l.501-.866l1.731 1.001l-.5.866a5 5 0 0 1-2.405 2.115l.193.387l-1.506 2.977a9 9 0 0 0 .957-17.844M1 12C1 5.925 5.925 1 12 1q.555 0 1.099.054C18.659 1.606 23 6.296 23 12c0 6.075-4.925 11-11 11q-.555 0-1.099-.054C5.341 22.394 1 17.704 1 12m8-4v4H7V8zm8 0v4h-2V8z"/>`), g.Group(children),
	)
}

func Creditcard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23 3v18H1V3zm-2 2H3v4h18zm0 6H3v8h18zm-11 5H5v-2h5z"/>`), g.Group(children),
	)
}

func CreditcardAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v9h-2v-1H3v8h11v2H1zm2 6h18V5H3zm2 5h5v2H5zm16 0v3h3v2h-3v3h-2v-3h-3v-2h3v-3z"/>`), g.Group(children),
	)
}

func CreditcardOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 .586L23.414 22L22 23.414L19.586 21H1V3h.586l-1-1zM3 5v4h4.586l-4-4zm0 6v8h14.586l-8-8zm4.581-8.001L23.001 3l.002 9.418l-.003-.004l.003 6.004L21 16.415v-5.414l-5.413.001l-2.005-2.003L21 9V5l-11.413.001zM5.001 14h5v2H5z"/>`), g.Group(children),
	)
}

func CrookedSmile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m8-4v4H7V8zm8 0v4h-2V8zm-.282 5.78l-.25.97c-.269 1.045-.793 1.895-1.613 2.467c-.806.563-1.792.783-2.855.783h-1v-2h1c.8 0 1.343-.167 1.71-.423c.353-.246.647-.646.822-1.326l.249-.969z"/>`), g.Group(children),
	)
}

func CryAndLaugh(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h4v2H9v2h6v-2h-1V9h4v2h-1v3h-.98l-.04.195C15.582 16.199 14.09 18 12 18s-3.581-1.802-3.98-3.805L7.98 14H7v-3H6zm4.445 6c.422.662 1.02 1 1.555 1s1.133-.338 1.555-1z"/>`), g.Group(children),
	)
}

func CryLoudly(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h4v2H9v4H7v-4H6zm8 0h4v2h-1v4h-2v-4h-1zm-1 4v3h-2v-3z"/>`), g.Group(children),
	)
}

func CssThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.941 1H22.06l-1.098 19.208l-8.96 3.36l-8.962-3.36zM4.06 3l.902 15.792l7.04 2.64l7.038-2.64L19.941 3zm1.423 2H18.33l-.108 1.887l-7.29 3.49h7.09l-.404 7.084L12 19.568l-5.618-2.107l-.114-1.994v-.004l-.075-1.39h2.01l.062 1.284l.04.69L12 17.431l3.696-1.386l.21-3.67H6.09l-.106-1.848L13.355 7H12l-6.255.012z"/>`), g.Group(children),
	)
}

func Cucumber(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.726 2.18c2.16-.683 3.992-.404 5.36.583a5.1 5.1 0 0 1 1.244 1.284c.93 1.38 1.158 3.201.456 5.333c-.693 2.104-2.296 4.545-5.078 7.327c-2.828 2.828-5.304 4.438-7.433 5.112c-2.16.684-3.992.405-5.36-.582a5.1 5.1 0 0 1-1.244-1.284c-.93-1.38-1.158-3.201-.456-5.333c.693-2.104 2.296-4.545 5.078-7.327c2.828-2.828 5.304-4.438 7.433-5.112m.603 1.908c-1.705.54-3.913 1.91-6.621 4.62c-2.664 2.663-4.035 4.843-4.593 6.538a6 6 0 0 0-.305 1.53c3.124-5.112 8.103-9.938 13.188-12.972c-.487 0-1.041.085-1.669.284m3.872.766C13.545 7.884 7.797 13.469 4.746 19.16c2.679-1.493 5.528-3.782 8.11-6.372c2.558-2.566 4.81-5.383 6.345-7.935M6.922 20.196c.507.01 1.087-.074 1.75-.284c1.705-.54 3.913-1.911 6.621-4.62c2.664-2.663 4.035-4.844 4.593-6.538c.215-.653.309-1.228.31-1.732c-1.576 2.391-3.644 4.893-5.923 7.179c-2.297 2.303-4.84 4.42-7.35 5.995"/>`), g.Group(children),
	)
}

func CurrencyExchange(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.5 5.835A10.49 10.49 0 0 0 12 1.5c-5.427 0-9.89 4.115-10.443 9.396l-.104.994l1.99.209l.103-.995A8.501 8.501 0 0 1 19.213 7.5H17.5v2h5v-7h-2zM11 6v1a3 3 0 0 0 0 6h2a1 1 0 1 1 0 2H8.5v2H11v1h2v-1a3 3 0 1 0 0-6h-2a1 1 0 0 1 0-2h4.5V7H13V6zm9.557 5.901l-.104.995A8.501 8.501 0 0 1 4.787 16.5H6.5v-2h-5v7h2v-3.335A10.49 10.49 0 0 0 12 22.5c5.426 0 9.89-4.115 10.442-9.396l.104-.994z"/>`), g.Group(children),
	)
}

func Cursor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m3.108 3.114l16.864 5.925l-4.319 3.084l5.707 5.706l-3.536 3.536l-5.706-5.706l-3.085 4.318zm3.269 3.268l3.267 9.3l2.219-3.107l5.96 5.961l.708-.707l-5.961-5.961l3.106-2.219z"/>`), g.Group(children),
	)
}

func Curtain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 7.465V20h2.914L6.093 9.999L6 10a4 4 0 0 1-2-.535m4.053-.032L8.921 20h6.158l.868-10.567A4 4 0 0 1 15 8.646A4 4 0 0 1 12 10a4 4 0 0 1-3-1.354a4 4 0 0 1-.947.787M8 6h2a2 2 0 1 0 4 0h2a2 2 0 0 0 2 2a2 2 0 0 0 2-2V4H4v2a2 2 0 1 0 4 0m12 3.465a4 4 0 0 1-2.093.534L17.085 20H20z"/>`), g.Group(children),
	)
}

func Curve(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.328 7.438C3.504 8.994 3 10.768 3 11.758v1H1v-1c0-1.436.652-3.54 1.56-5.256c.462-.871 1.023-1.707 1.66-2.34C4.84 3.547 5.658 3 6.622 3c.938 0 1.735.49 2.354 1.079c.627.596 1.174 1.39 1.635 2.227c.922 1.676 1.599 3.72 1.874 5.278l.002.013l.002.013c.173 1.164.804 3.118 1.674 4.773c.434.826.9 1.524 1.357 2c.48.501.811.617.98.617c.186 0 .513-.115.96-.576c.433-.445.869-1.103 1.26-1.896c.788-1.6 1.28-3.508 1.28-4.77v-1h2v1c0 1.647-.602 3.86-1.486 5.654c-.446.904-.991 1.76-1.619 2.406C18.283 20.448 17.47 21 16.5 21c-.986 0-1.81-.591-2.424-1.232c-.637-.665-1.204-1.542-1.683-2.454c-.953-1.813-1.666-3.973-1.88-5.395c-.24-1.35-.845-3.178-1.654-4.65c-.406-.737-.838-1.338-1.262-1.741C7.165 5.116 6.84 5 6.622 5c-.192 0-.529.12-.993.58c-.446.444-.896 1.092-1.301 1.858M4 11h5v2H4zm10 0h5v2h-5z"/>`), g.Group(children),
	)
}

func Cut(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.414 4.586a2 2 0 1 0-2.828 2.828a2 2 0 0 0 2.828-2.828M3.172 3.172a4 4 0 0 1 6.274 4.86L12 10.586l8-8L21.414 4L9.446 15.968a4.002 4.002 0 0 1-6.274 4.86a4 4 0 0 1 4.86-6.274L10.586 12L8.032 9.446a4.002 4.002 0 0 1-4.86-6.274M15 13.586L21.414 20L20 21.414L13.586 15zm-7.586 3a2 2 0 1 0-2.828 2.828a2 2 0 0 0 2.828-2.828"/>`), g.Group(children),
	)
}

func CutOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v2.5h-2V1zm8.414 2.5L15 9.914L13.586 8.5L20 2.086zM4 2.086l11.968 11.968a4.001 4.001 0 0 1 4.86 6.274a4 4 0 0 1-6.274-4.86L12 12.914l-2.554 2.554a4.002 4.002 0 0 1-6.274 4.86a4 4 0 0 1 4.86-6.274l2.554-2.554l-8-8zM13 5v2.5h-2V5zM7.414 16.086a2 2 0 1 0-2.828 2.828a2 2 0 0 0 2.828-2.828m12 0a2 2 0 1 0-2.828 2.828a2 2 0 0 0 2.828-2.828"/>`), g.Group(children),
	)
}

func Dam(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .671l9 7.875V20h1v2h-9v-2h1V10h-4v10h1v2H2v-2h1V8.546zm-7 19.33h3V10H5zM6.662 8H17.34L12 3.328zM19 10h-3v10h3z"/>`), g.Group(children),
	)
}

func DamFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v2h-1v16h1v2H2v-2h1V4H2zm3 2v16h2V4zm4 0v7c.836-.628 1.874-1 3-1s2.164.372 3 1V4zm8 0v16h2V4zm-2 16v-5a3 3 0 1 0-6 0v5z"/>`), g.Group(children),
	)
}

func DamFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h4V10h3V4zm9 0v6h3v10h4V4zm1 16v-8h-4v8z"/>`), g.Group(children),
	)
}

func DamOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3c3.905 0 7 3.162 7 7.1V20h-2v-9.9a5 5 0 0 0-10 0V20H5v-9.9C5 6.162 8.095 3 12 3m9 17v-9.9C21 5.09 17.042 1 12 1s-9 4.09-9 9.1V20H2v2h8v-2H9v-9.9a3 3 0 1 1 6 0V20h-1v2h8v-2z"/>`), g.Group(children),
	)
}

func DamSeven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.22 1h11.56l1 4H21v15h1v2H2v-2h1V5h2.22zm1.06 4h9.44l-.5-2H7.78zM5 20h3v-8h8v8h3V7H5zm9 0v-6h-4v6z"/>`), g.Group(children),
	)
}

func DamSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 2h18v5h1v2h-1v11h1v2h-9v-2h1v-5a2 2 0 1 0-4 0v5h1v2H2v-2h1V9H2V7h1zm2 7v11h3v-5a4 4 0 0 1 8 0v5h3V9zm14-2V4H5v3z"/>`), g.Group(children),
	)
}

func DamThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.064 9h13.872C18.45 5.255 15.25 3 12 3S5.55 5.255 5.064 9M19 11h-2v9h2zm2 9h1v2h-8v-2h1v-9H9v9h1v2H2v-2h1V10c0-5.6 4.529-9 9-9s9 3.4 9 9zM7 11H5v9h2zm3-6h4v2h-4z"/>`), g.Group(children),
	)
}

func DamTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 2h18v18h1v2h-9v-2h1V10h-4v10h1v2H2v-2h1zm2 18h3V10H5zM5 8h14V4H5zm14 2h-3v10h3z"/>`), g.Group(children),
	)
}

func DartBoard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m11-5a5 5 0 1 0 0 10a5 5 0 0 0 0-10m-7 5a7 7 0 1 1 14 0a7 7 0 0 1-14 0m5 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func Dashboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-4.5 16.796l.865.5l-1.001 1.732l-.866-.5A11 11 0 0 1 1 12C1 5.925 5.925 1 12 1c1.203 0 2.362.193 3.448.552l.95.313l-.627 1.9l-.95-.314A9 9 0 0 0 12 3m8.981 1.414l-5.542 5.542a4 4 0 1 1-1.42-1.41L19.567 3zm1.154 3.188l.313.95C22.807 9.638 23 10.797 23 12c0 4.071-2.212 7.625-5.496 9.526l-.865.5l-1.002-1.73l.865-.501A9 9 0 0 0 21 12a9 9 0 0 0-.45-2.822l-.314-.95zM12 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4"/>`), g.Group(children),
	)
}

func DashboardOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h9v7H2zm2 2v3h5V4zm9-2h9v11h-9zm2 2v7h5V4zM2 11h9v11H2zm2 2v7h5v-7zm9 2h9v7h-9zm2 2v3h5v-3z"/>`), g.Group(children),
	)
}

func Data(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v7h1.998v-.002h2.004V11H20V4zm16 9H8.002v.002H5.998V13H4v7h16zM5.998 6.5h2.004v2.004H5.998zm0 9h2.004v2.004H5.998z"/>`), g.Group(children),
	)
}

func DataBase(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 5c0-.002-.003-.019.035-.081c.043-.07.13-.174.291-.301c.33-.26.87-.539 1.626-.79C7.454 3.326 9.592 3 12 3s4.546.327 6.048.827c.756.252 1.296.53 1.626.79c.162.128.248.232.29.302c.04.063.036.08.036.081c0 .003.003.02-.035.081c-.043.07-.13.174-.291.301c-.33.26-.87.539-1.626.79C16.546 6.674 14.408 7 12 7s-4.546-.327-6.048-.827c-.756-.252-1.296-.53-1.626-.79a1.2 1.2 0 0 1-.29-.302C3.997 5.021 4 5.003 4 5m16 2.527V12l-.003.015a.3.3 0 0 1-.032.066c-.043.07-.13.174-.291.301c-.33.26-.87.539-1.626.79c-1.502.501-3.64.828-6.048.828s-4.546-.327-6.048-.827c-.756-.252-1.296-.53-1.626-.79a1.2 1.2 0 0 1-.29-.302a.3.3 0 0 1-.033-.066L4 12V7.527a9.6 9.6 0 0 0 1.32.543C7.075 8.655 9.437 9 12 9s4.925-.345 6.68-.93A9.5 9.5 0 0 0 20 7.527m0 7V19l-.003.015a.3.3 0 0 1-.032.066c-.043.07-.13.174-.291.301c-.33.26-.87.539-1.626.79c-1.502.501-3.64.828-6.048.828s-4.546-.327-6.048-.827c-.756-.252-1.296-.53-1.626-.79a1.2 1.2 0 0 1-.29-.302a.3.3 0 0 1-.033-.066L4 19v-4.473a9.5 9.5 0 0 0 1.32.543c1.755.585 4.117.93 6.68.93s4.925-.345 6.68-.93a9.5 9.5 0 0 0 1.32-.543M2 19c0 .852.519 1.504 1.088 1.953c.581.458 1.36.826 2.232 1.117c1.755.585 4.117.93 6.68.93s4.925-.345 6.68-.93c.873-.29 1.651-.66 2.232-1.117C21.482 20.504 22 19.852 22 19V5c0-.852-.519-1.504-1.088-1.953c-.581-.458-1.36-.826-2.232-1.117C16.925 1.345 14.563 1 12 1s-4.925.345-6.68.93c-.873.29-1.651.66-2.232 1.117C2.518 3.496 2 4.148 2 5z"/>`), g.Group(children),
	)
}

func DataChecked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v12h-2v-1H4v7h8.5v2H2zm2 9h16V4H4zm1.998-4.5h2.004v2.004H5.998zm0 9h2.004v2.004H5.998zm17.598 1.44l-5.657 5.656l-3.535-3.535l1.414-1.415l2.121 2.122l4.243-4.243z"/>`), g.Group(children),
	)
}

func DataDisplay(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 18H3v2h19zm-2-2a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2h-4v2h4v3h-3.5v2H20v3h-4v2zm-5-2h-4v-3h2a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H9v2h4v3h-2a2 2 0 0 0-2 2v5h6zm-7 0H6.5V6a2 2 0 0 0-2-2H3v2h1.5v8H3v2h5z"/>`), g.Group(children),
	)
}

func DataError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v12h-2v-1H4v7h10v2H2zm2 9h16V4H4zm1.998-4.5h2.004v2.004H5.998zm10.88 8.964L19 17.586l2.122-2.122l1.414 1.415L20.414 19l2.122 2.121l-1.415 1.415l-2.12-2.122l-2.122 2.121l-1.414-1.414L17.586 19l-2.121-2.121zm-10.88.036h2.004v2.004H5.998z"/>`), g.Group(children),
	)
}

func DataSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v11H4v7h9v2H2zm2 9h16V4H4zm1.998-4.5h2.004v2.004H5.998zM17.75 16a1.75 1.75 0 1 0 0 3.5a1.75 1.75 0 0 0 0-3.5M14 17.75a3.75 3.75 0 1 1 7.013 1.849l1.651 1.651l-1.414 1.414l-1.651-1.65A3.75 3.75 0 0 1 14 17.75M5.998 15.5h2.004v2.004H5.998z"/>`), g.Group(children),
	)
}

func Delete(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 1h9v3H22v2h-2.029l-.5 17H4.529l-.5-17H2V4h5.5zm2 3h5V3h-5zM6.03 6l.441 15h11.058l.441-15zM13 8v11h-2V8z"/>`), g.Group(children),
	)
}

func DeleteOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 1h9v3H22v2h-2.029l-.5 17H4.529l-.5-17H2V4h5.5zm2 3h5V3h-5zM6.03 6l.441 15h11.058l.441-15zm3.142 3.257L12 12.086l2.828-2.829l1.415 1.415l-2.829 2.828l2.829 2.828l-1.415 1.415L12 14.914l-2.828 2.829l-1.415-1.415l2.829-2.828l-2.829-2.828z"/>`), g.Group(children),
	)
}

func DeleteTime(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 1h9v3H22v2h-2.029l-.163 5.529l-1.999-.059L17.97 6H6.03l.441 15H11.5v2H4.529l-.5-17H2V4h5.5zm2 3h5V3h-5zM13 8v7h-2V8zm5.5 7a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M13 18.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m6.5-2.248v1.834l1.414 1.414l-1.414 1.414l-2-2v-2.662z"/>`), g.Group(children),
	)
}

func Delta(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 2.091V22.79L3.172 13.05zM6.828 12.951L17 19.21V5.907z"/>`), g.Group(children),
	)
}

func Depressed(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m9.233-2.134l-3.464 2l-1-1.732l3.464-2zm4.536-1.732l3.464 2l-1 1.732l-3.464-2zm-7.1 7.365A5 5 0 0 1 12 13a5 5 0 0 1 4.33 2.5l.501.865l-1.731 1.001l-.5-.865A3 3 0 0 0 12 15a3 3 0 0 0-2.6 1.5l-.5.866l-1.731-1.001z"/>`), g.Group(children),
	)
}

func Desktop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v15H13v2h4v2H7v-2h4v-2H1zm2 2v11h18V5z"/>`), g.Group(children),
	)
}

func DesktopOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v16H13v1h4v2H7v-2h4v-1H1zm2 2v9h18V5zm18 11H3v1h18z"/>`), g.Group(children),
	)
}

func Despise(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h4v2H9v2H7v-2H6zm8 0h4v2h-1v2h-2v-2h-1zm-6.33 6.5A5 5 0 0 1 12 13a5 5 0 0 1 4.33 2.5l.501.865l-1.731 1.001l-.5-.865A3 3 0 0 0 12 15a3 3 0 0 0-2.6 1.5l-.5.866l-1.731-1.001z"/>`), g.Group(children),
	)
}

func Device(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2h22v4h-2V4H3v13h9v2H1zm13 6h10v14H14zm2 2v10h6V10zm1.998 6.998h2.004v2.004h-2.004zM5 20h7v2H5z"/>`), g.Group(children),
	)
}

func Discount(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.496 1.19l10.506.808l.808 10.505l-10.932 10.932L.564 12.121zm.764 2.064l-8.867 8.867l8.485 8.486l8.867-8.868l-.606-7.879zm3.86 4.624a1 1 0 1 0-1.414 1.414a1 1 0 0 0 1.414-1.414m-2.828-1.414a3 3 0 1 1 4.243 4.242a3 3 0 0 1-4.243-4.242"/>`), g.Group(children),
	)
}

func DiscountFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.474 7.524a1.5 1.5 0 1 0-2.122 2.122a1.5 1.5 0 0 0 2.122-2.122"/><path fill="currentColor" d="m22.003 1.998l.808 10.505l-10.932 10.932L.565 12.121L11.497 1.19zm-8.711 4.466a3 3 0 1 0 4.243 4.242a3 3 0 0 0-4.243-4.242"/>`), g.Group(children),
	)
}

func Dissatisfaction(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m7-3.5v1a.5.5 0 0 0 1 0v-1h2v1a2.5 2.5 0 0 1-5 0v-1zm7 0v1a.5.5 0 0 0 1 0v-1h2v1a2.5 2.5 0 0 1-5 0v-1zm-7.33 7A5 5 0 0 1 12 13a5 5 0 0 1 4.33 2.5l.501.865l-1.731 1.001l-.5-.865A3 3 0 0 0 12 15a3 3 0 0 0-2.6 1.5l-.5.866l-1.731-1.001z"/>`), g.Group(children),
	)
}

func Divide(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 7h2.004v2.004H11zm-4.5 4h11v2h-11zm4.5 4h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func Dividers(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v1.232a4.36 4.36 0 0 1 3.375 4.239a4.3 4.3 0 0 1-.678 2.33L18.967 15H22v2h-1.979l2.33 4.418l-1.768.933L17.76 17H13v2h-2v-2H6.24l-2.822 5.351l-1.769-.933L3.98 17H2v-2h3.035l3.27-6.2a4.3 4.3 0 0 1-.68-2.33A4.36 4.36 0 0 1 11 2.233V1zm-3.193 9.238L7.296 15H11v-2h2v2h3.705l-2.511-4.762a4.4 4.4 0 0 1-2.194.585a4.4 4.4 0 0 1-2.193-.585M12 4.118A2.364 2.364 0 0 0 9.625 6.47c0 .6.227 1.148.603 1.566A2.38 2.38 0 0 0 12 8.824c.706 0 1.338-.304 1.773-.787a2.33 2.33 0 0 0 .602-1.566A2.364 2.364 0 0 0 12 4.118"/>`), g.Group(children),
	)
}

func DividersOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v3.532c1.945.44 3.427 2.133 3.427 4.205c0 .929-.298 1.782-.803 2.48l2.487 4.878a19 19 0 0 0 1.776-.703l.91-.415l.83 1.82l-.91.415q-.833.38-1.688.682L21.632 23h-2.249l-2.303-4.52a21.1 21.1 0 0 1-10.16 0L4.615 23H2.367l2.604-5.106q-.856-.303-1.688-.682l-.91-.415l.83-1.82l.91.415q.873.398 1.775.703l2.487-4.878a4.2 4.2 0 0 1-.802-2.48c0-2.072 1.481-3.765 3.427-4.205V1zM7.854 16.645c2.728.607 5.564.607 8.291 0l-2.083-4.09a4.5 4.5 0 0 1-2.062.498c-.74 0-1.444-.18-2.062-.497zM12 6.421c-1.372 0-2.427 1.068-2.427 2.316s1.055 2.316 2.427 2.316c1.371 0 2.427-1.068 2.427-2.316S13.37 6.42 12 6.42"/>`), g.Group(children),
	)
}

func Doge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h4.739l2.034 1.973A11.5 11.5 0 0 1 12 2.176c1.456 0 2.85.269 4.125.757L17.62 1H23v12h-.008c-.23 5.611-5.118 10-10.992 10S1.237 18.611 1.008 13H1zm2 5.6a10.8 10.8 0 0 1 2.888-2.67L4.928 3H3zm0 5.988C3 17.178 6.971 21 12 21s9-3.822 9-8.412c0-3.268-2.001-6.135-4.98-7.527A9.5 9.5 0 0 0 12 4.176c-1.449 0-2.813.32-4.02.885C5 6.453 3 9.32 3 12.588M21 6.6V3h-2.398l-.645.834A10.8 10.8 0 0 1 21 6.6M9 9v3H7V9zm8 0v3h-2V9zm-5 2.264l.895 1.789c.379.757.857.985 1.189 1.013c.216.018.447-.038.668-.189l.827-.562l1.124 1.654l-.826.562a3.1 3.1 0 0 1-1.167.487a11 11 0 0 1-.671 1.386c-.2.34-.443.697-.728.982c-.26.26-.708.614-1.31.614s-1.05-.355-1.31-.614a5 5 0 0 1-.728-.982a11 11 0 0 1-.672-1.386a3.1 3.1 0 0 1-1.166-.487l-.827-.562l1.125-1.654l.827.562c.221.151.452.207.668.189c.332-.028.81-.256 1.189-1.013zm-.678 4.413c.111.244.236.492.367.716q.173.292.312.462q.14-.17.312-.462a8 8 0 0 0 .367-.716a3.5 3.5 0 0 1-.68-.503a3.5 3.5 0 0 1-.678.503"/>`), g.Group(children),
	)
}

func DoubleStorey(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 2h14v2h-1v4h4v2h-1v10h1v2H2v-2h1V10H2V8h4V4H5zm3 2v4h8V4zm-3 6v10h3v-7h8v7h3V10zm9 10v-5h-4v5zM10.998 4.998h2.004v2.004h-2.004z"/>`), g.Group(children),
	)
}

func Download(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 3v9.586l3.5-3.5l1.414 1.414L12 16.414L6.086 10.5L7.5 9.086l3.5 3.5V3zM4.5 14v5h15v-5h2v7h-19v-7z"/>`), g.Group(children),
	)
}

func DownloadOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 3v9.586l4-4L18.414 10L12 16.414L5.586 10L7 8.586l4 4V3zM3 18h18v2H3z"/>`), g.Group(children),
	)
}

func Downscale(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h9v2H4v6H2zm19.414 2l-5.027 5.027l3.527.02l-.011 2l-6.91-.04l-.04-6.91l2-.011l.02 3.526L20 2.585zM2 12h4v2H4v2H2zm6 0h4v4h-2v-2H8zm14 1v9h-8v-2h6v-7zM4 18v2h2v2H2v-4zm8 0v4H8v-2h2v-2z"/>`), g.Group(children),
	)
}

func DragDrop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.999 15V2h13v5h-2V4h-9v9h3v2zm6 5V8h12v5.5h-2V10h-8v8h3.5v2zm8.778 3.684L13.41 13.378l10.258 3.407l-4.656 2.227z"/>`), g.Group(children),
	)
}

func DragMove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 23.414L7.586 19L9 17.586l2 2V13H4.414l2 2L5 16.414L.586 12L5 7.586L6.414 9l-2 2H11V4.414l-2 2L7.586 5L12 .586L16.414 5L15 6.414l-2-2V11h6.586l-2-2L19 7.586L23.414 12L19 16.414L17.586 15l2-2H13v6.586l2-2L16.414 19z"/>`), g.Group(children),
	)
}

func Drink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.48 1H20v2h-4.02l-.65 2.222h4.74L18.858 23H6.217l-.87-12.722C3.475 9.863 2 8.239 2 6.222c0-2.32 1.914-4.166 4.231-4.166c1.973 0 3.653 1.337 4.11 3.166h2.904zm-1.82 6.222H7.145l.236 3.482l4.067.661zm-.664 6.258l-4.475-.727L8.085 21h8.904l.451-6.616l-5.44-.903zm5.581-1.1l.352-5.158h-3.185l-1.308 4.47zM8.211 5.221a2.23 2.23 0 0 0-1.98-1.166C4.98 4.056 4 5.045 4 6.222c0 .797.48 1.523 1.201 1.896l-.197-2.896z"/>`), g.Group(children),
	)
}

func Drumstick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 3a6 6 0 0 0-4.199 10.286l.715.7l-1.4 1.428l-.714-.7a8 8 0 0 1-1.718-2.475l-.478.47l-.008.013a.3.3 0 0 0-.021.059a1 1 0 0 0 .001.353c.054.352.24.832.6 1.374a6.8 6.8 0 0 0 1.716 1.766c.571.405 1.092.625 1.478.696c.19.035.32.028.395.012a.3.3 0 0 0 .068-.022l.015-.01l.006-.006l2.128-2.09l.495.076q.45.07.921.07a6 6 0 0 0 0-12m-7.932 7.04a8 8 0 1 1 7.217 6.928l-1.424 1.4c-.636.632-1.512.704-2.248.57a5.2 5.2 0 0 1-1.554-.58L8 19.416c-.005.506-.036 1.175-.214 1.77c-.122.406-.334.858-.722 1.218c-.407.377-.933.581-1.538.597c-.629.016-1.16-.188-1.567-.542c-.384-.335-.6-.755-.726-1.103a4 4 0 0 1-.132-.458a4 4 0 0 1-.455-.132a2.65 2.65 0 0 1-1.1-.725c-.353-.404-.559-.934-.545-1.562c.013-.606.217-1.133.596-1.54c.36-.387.813-.599 1.219-.72c.595-.18 1.265-.212 1.77-.217l1.107-1.104a5 5 0 0 1-.492-1.46c-.108-.71-.014-1.538.592-2.144l.007-.006zm-.22 6.529L5.415 18H5c-.591 0-1.17 0-1.608.133c-.204.06-.294.129-.33.167c-.016.017-.058.063-.062.222c-.003.139.033.18.051.2a.7.7 0 0 0 .279.163a2 2 0 0 0 .656.115h.004l1.03-.017l-.02 1.03v.004l.001.036a2.3 2.3 0 0 0 .112.622c.057.157.119.24.16.276c.02.017.062.053.202.05c.164-.005.211-.048.23-.065c.037-.035.104-.125.165-.326c.13-.434.13-1.004.13-1.598v-.427l1.427-1.423a9 9 0 0 1-.578-.593"/>`), g.Group(children),
	)
}

func Dv(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h10v6.682c.456.687.775 1.473.917 2.318H14V7h8v10h-8v-4h-1.083A6 6 0 0 1 12 15.318V22H2v-6.682c-.632-.95-1-2.093-1-3.318s.368-2.367 1-3.318zm2 4.803A6 6 0 0 1 7 6c1.093 0 2.118.293 3 .803V4H4zm0 10.394V20h6v-2.803A6 6 0 0 1 7 18a6 6 0 0 1-3-.803M7 8a4 4 0 0 0-3.2 1.6A3.98 3.98 0 0 0 3 12c0 .902.297 1.731.8 2.4A4 4 0 0 0 7 16a4 4 0 0 0 3.2-1.6c.503-.669.8-1.498.8-2.4s-.297-1.731-.8-2.4A4 4 0 0 0 7 8m9 1v6h4V9z"/>`), g.Group(children),
	)
}

func Dvd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm8 4a4 4 0 1 0 0 8a4 4 0 0 0 0-8m-6 4a6 6 0 1 1 12 0a6 6 0 0 1-12 0m4 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func Earphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4c-4.411 0-8 3.605-8 8.067V13h2a3 3 0 0 1 3 3v2a3 3 0 0 1-3 3H2v-8.933C2 6.514 6.47 2 12 2s10 4.514 10 10.067V21h-4a3 3 0 0 1-3-3v-2a3 3 0 0 1 3-3h2v-.933C20 7.605 16.411 4 12 4m8 11h-2a1 1 0 0 0-1 1v2a1 1 0 0 0 1 1h2zM4 15v4h2a1 1 0 0 0 1-1v-2a1 1 0 0 0-1-1z"/>`), g.Group(children),
	)
}

func Earth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m1.045 11.003l-.076.777l.033.012L1 12c0 6.075 4.925 11 11 11s11-4.925 11-11S18.075 1 12 1C8.007 1 4.51 3.128 2.584 6.311l-.038-.014l-.381.772a10.9 10.9 0 0 0-1.12 3.934m2.1-.596c.1-.546.245-1.076.434-1.582L7.22 10.2l2.02-2.052l1.735.586v4.067l3.294 1.114l1.987 4.033l4.386-5.084l.333-.193A9 9 0 0 1 6.18 18.866l3.195-6.342zm17.718.016l-1.515.88l-2.632 3.051l-1.02-2.068l-2.72-.92V7.3l-4.28-1.447l-1.988 2.016l-2.214-.835A8.99 8.99 0 0 1 12 3a9 9 0 0 1 8.863 7.423m-16.135 6.88a8.95 8.95 0 0 1-1.715-4.828l3.542 1.203z"/>`), g.Group(children),
	)
}

func Edit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.748 2.947a2 2 0 0 1 2.828 0l2.475 2.475a2 2 0 0 1 0 2.829L9.158 20.144l-6.38 1.076l1.077-6.38zm-.229 3.057l2.475 2.475l1.643-1.643l-2.475-2.474zm1.06 3.89l-2.474-2.475l-8.384 8.384l-.503 2.977l2.977-.502z"/>`), g.Group(children),
	)
}

func EditOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.748 2.947a2 2 0 0 1 2.829 0l2.474 2.475a2 2 0 0 1 0 2.829l-2.35 2.35l-2.202 2.202l-1.415-1.414l1.496-1.496l-2.475-2.474l-1.495 1.495L11.195 7.5l2.203-2.203zm-.228 3.057l2.474 2.475l1.643-1.643l-2.475-2.474zM4 2.586L21.414 20L20 21.414l-6.056-6.056l-4.786 4.786l-6.38 1.076l1.077-6.38l4.785-4.785L2.586 4zm6.055 8.883L5.72 15.803l-.503 2.977l2.977-.502l4.335-4.334z"/>`), g.Group(children),
	)
}

func EditOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.828 1.416l5.755 5.755L7.755 22H2v-5.756zm0 8.681l2.927-2.926l-2.927-2.927l-2.926 2.927zm-4.34-1.512L4 17.074V20h2.926l8.488-8.488z"/>`), g.Group(children),
	)
}

func EditTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.428 1.963l5.61 5.61L7.61 22.001H2v-5.61zm0 2.828l-2.782 2.782l2.781 2.782l2.782-2.782zm-1.415 6.978l-2.782-2.782L4 17.22V20h2.782zm7.212 10.232h-9.543v-2h9.543z"/>`), g.Group(children),
	)
}

func Education(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 .807L23.835 8.5L22 9.693V16h-2v-5.007l-1 .65V17.5c0 1.47-1.014 2.615-2.253 3.339c-1.265.737-2.945 1.16-4.747 1.16s-3.483-.423-4.747-1.16C6.013 20.115 5 18.969 5 17.499v-5.856L.165 8.5zM7 12.943V17.5c0 .463.33 1.067 1.261 1.61c.908.53 2.227.89 3.739.89s2.831-.36 3.739-.89c.932-.543 1.26-1.147 1.26-1.61v-4.557l-5 3.25zM20.165 8.5L12 3.193L3.835 8.5L12 13.807z"/>`), g.Group(children),
	)
}

func Eggplant(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.087 5.899l-1.485-.693l1.01-1.01a6.5 6.5 0 0 1 8.506-.598l1.488-.744l.894 1.79l-.934.466a6.5 6.5 0 0 1-.761 8.28l-1.01 1.01l-.694-1.488a73 73 0 0 1-1.48 1.536c-.457.458-.661 1.073-.994 2.075l-.091.274c-.367 1.096-.868 2.429-2.094 3.66c-1.209 1.212-2.845 2.044-4.623 2.149c-1.801.106-3.668-.542-5.275-2.149s-2.256-3.47-2.151-5.272c.103-1.777.935-3.414 2.148-4.627c1.232-1.232 2.566-1.734 3.663-2.1l.282-.093c.997-.33 1.608-.532 2.065-.988A72 72 0 0 1 11.087 5.9m2.204-1.18l1.585.74l.5 3.166l3.166.5l.74 1.585a4.503 4.503 0 0 0-5.991-5.99m3.888 6.216l-3.552-.561l-.56-3.552l-.026-.012l-.419.39c-.52.488-1.16 1.094-1.657 1.591c-.816.816-1.907 1.169-2.81 1.461l-.319.104c-1.024.341-1.989.724-2.881 1.617c-.912.912-1.494 2.102-1.566 3.328c-.07 1.202.345 2.518 1.569 3.742s2.54 1.637 3.744 1.566c1.228-.072 2.417-.655 3.323-1.564c.888-.891 1.271-1.857 1.614-2.882l.103-.316c.294-.905.65-1.997 1.466-2.813a75 75 0 0 0 1.982-2.077z"/>`), g.Group(children),
	)
}

func Ellipsis(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 10.5h3v3H3zm7.5 0h3v3h-3zm7.5 0h3v3h-3z"/>`), g.Group(children),
	)
}

func EmoEmotional(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m8-4v4H7V8zm8 0v4h-2V8zm-9 6h8v2H8z"/>`), g.Group(children),
	)
}

func Enter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 4v9a4 4 0 0 1-4 4H6.914l2.5 2.5L8 20.914L3.086 16L8 11.086L9.414 12.5l-2.5 2.5H16a2 2 0 0 0 2-2V4z"/>`), g.Group(children),
	)
}

func Equal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 7h16v2H4zm0 8h16v2H4z"/>`), g.Group(children),
	)
}

func Error(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 2v14.5h-2V2zm-2 17h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func ErrorCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m12-5.5V14h-2V6.5zm-2 9h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func ErrorCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1m-1 13h2V6.5h-2zm2.004 1.5H11v2.004h2.004z"/>`), g.Group(children),
	)
}

func ErrorTriangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 1l11.951 20.7H.05zM3.513 19.7h16.974L12 5zM13 9.5V15h-2V9.5zm-2 7h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func Excited(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m6.619-4.862L11 9.04v1.46l-3.2 2.4l-1.2-1.6l1.86-1.394L6.637 8.88zm9.743 1.743L15.54 9.906L17.4 11.3l-1.2 1.6l-3.2-2.4V9.04l3.381-1.902zM12 11.264l.894 1.789c.379.757.857.985 1.189 1.013c.216.018.447-.038.668-.189l.827-.562l1.124 1.654l-.826.562a3.1 3.1 0 0 1-1.167.487a11 11 0 0 1-.671 1.386c-.2.34-.443.697-.728.982c-.26.26-.708.614-1.31.614s-1.05-.355-1.31-.614a5 5 0 0 1-.728-.982a11 11 0 0 1-.672-1.386a3.1 3.1 0 0 1-1.166-.487l-.827-.562l1.125-1.654l.827.562c.221.151.452.207.668.189c.332-.028.81-.256 1.189-1.013zm-.68 4.413c.112.244.236.492.368.716q.173.292.312.462q.14-.17.312-.462a8 8 0 0 0 .367-.716a3.5 3.5 0 0 1-.68-.503a3.5 3.5 0 0 1-.678.503"/>`), g.Group(children),
	)
}

func ExcitedOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m6.619-4.862L11 9.04v1.46l-3.2 2.4l-1.2-1.6l1.86-1.394L6.637 8.88zm9.743 1.743L15.54 9.906L17.4 11.3l-1.2 1.6l-3.2-2.4V9.04l3.381-1.902zM8.9 13.634l.5.865C9.922 15.4 10.89 16 12 16s2.08-.601 2.6-1.5l.5-.866l1.731 1.001l-.5.866A5 5 0 0 1 12 18a5 5 0 0 1-4.33-2.5l-.501-.865z"/>`), g.Group(children),
	)
}

func ExpandHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.5 4v16h-2V4zm.586 8L9 7.086L10.414 8.5l-2.5 2.5h8.172l-2.5-2.5L15 7.086L19.914 12L15 16.914L13.586 15.5l2.5-2.5H7.914l2.5 2.5L9 16.914zM22.5 4v16h-2V4z"/>`), g.Group(children),
	)
}

func ExpandVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 3.5H4v-2h16zm-8 .586L16.914 9L15.5 10.414l-2.5-2.5v8.172l2.5-2.5L16.914 15L12 19.914L7.086 15L8.5 13.586l2.5 2.5V7.914l-2.5 2.5L7.086 9zM20 22.5H4v-2h16z"/>`), g.Group(children),
	)
}

func Explore(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.364 5.636A9 9 0 1 0 5.636 18.364A9 9 0 0 0 18.364 5.636M4.222 4.222c4.296-4.296 11.26-4.296 15.556 0s4.296 11.26 0 15.556s-11.26 4.296-15.556 0s-4.296-11.26 0-15.556m13.22 2.337l-4.965 12.91l-2.1-5.844l-5.845-2.1zm-7.174 4.902l1.672.6l.6 1.672l1.42-3.692z"/>`), g.Group(children),
	)
}

func ExploreOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.472 1.973A11 11 0 0 1 12 1c6.075 0 11 4.925 11 11c0 1.612-.348 3.146-.973 4.528l-.412.912l-1.822-.825l.412-.91A9 9 0 0 0 21 12a9 9 0 0 0-9-9a9 9 0 0 0-3.704.795l-.911.412l-.824-1.822zM3 1.586l7.604 7.603h.001l4.208 4.208v.001l7.6 7.602L21 22.414l-1.961-1.96A10.96 10.96 0 0 1 12 23C5.925 23 1 18.075 1 12c0-2.677.957-5.132 2.547-7.039L1.586 3zm1.968 4.796A8.96 8.96 0 0 0 3 12a9 9 0 0 0 9 9a8.96 8.96 0 0 0 5.618-1.968l-3.591-3.591l-1.55 4.028l-2.1-5.844l-5.845-2.1l4.029-1.55zm12.475.177l-1.937 5.036l-1.867-.718l.322-.836l-.836.321l-.718-1.866z"/>`), g.Group(children),
	)
}

func Exposure(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m12-6v3h3v2h-3v3h-2v-3H8V9h3V6zM8 16h8v2H8z"/>`), g.Group(children),
	)
}

func Extension(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 4a3 3 0 0 1 6 0h6v6a3 3 0 1 1 0 6v6h-6.403l-.122-.858a2.5 2.5 0 0 0-4.95 0L8.403 22H2v-6.403l.858-.122a2.5 2.5 0 0 0 0-4.95L2 10.403V4zm3-1a1 1 0 0 0-1 1v2H4v2.756a4.501 4.501 0 0 1 0 8.488V20h2.756a4.501 4.501 0 0 1 8.488 0H18v-6h2a1 1 0 1 0 0-2h-2V6h-6V4a1 1 0 0 0-1-1"/>`), g.Group(children),
	)
}

func ExtensionOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 3.99A3 3 0 0 1 14 4h6v6a3 3 0 0 1 .01 6l.05 2.516L18 16.408V14h2a1 1 0 1 0 0-2h-2V6h-6V4a1 1 0 1 0-2 0v2H7.586L5.528 3.942zM1 1.586L22.414 23L21 24.414L18.586 22h-4.988l-.123-.858a2.5 2.5 0 0 0-4.95 0L8.404 22H2v-6.403l.859-.122a2.5 2.5 0 0 0 0-4.95L2 10.403V5.414L-.414 3zm3 5.828v1.342a4.501 4.501 0 0 1 0 8.488V20h2.756a4.501 4.501 0 0 1 8.488 0h1.342z"/>`), g.Group(children),
	)
}

func FaceRetouching(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.996.695L5.094 2.9l2.204 1.097l-2.204 1.097l-1.098 2.204L2.9 5.093L.695 3.996L2.9 2.9zM12 3c-.622 0-1.23.063-1.815.183l-.98.2l-.4-1.96l.98-.2A11 11 0 0 1 12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12c0-.885.105-1.746.303-2.573l.233-.972l1.945.466l-.233.973A9 9 0 1 0 12 3m2.512 5.934h2.004v2.004h-2.003zm-7.049 0h2.004v2.004H7.463zm2.447 4.252l.488.872a1.834 1.834 0 0 0 3.206 0l.488-.872l1.745.977l-.488.872a3.834 3.834 0 0 1-6.696 0l-.488-.872zm10.92 7.633l-1.568.78l1.569.782l.78 1.569l.782-1.57l1.569-.78l-1.57-.781l-.78-1.569z"/>`), g.Group(children),
	)
}

func FactCheck(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23 3v18H1V3zm-2 2H3v14h18zm-1.086 5.5L15 15.414L12.086 12.5l1.414-1.414l1.5 1.5l3.5-3.5zM11 17H5v-2h6zm0-8H5V7h6zm0 4H5v-2h6z"/>`), g.Group(children),
	)
}

func FahrenheitScale(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.5 6a.5.5 0 1 0 0 1a.5.5 0 0 0 0-1M4 6.5a2.5 2.5 0 1 1 5 0a2.5 2.5 0 0 1-5 0m7-.5a2 2 0 0 1 2-2h7v2h-7v5h7v2h-7v7h-2z"/>`), g.Group(children),
	)
}

func FeelAtEase(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m7-3.5v1a.5.5 0 0 0 1 0v-1h2v1a2.5 2.5 0 0 1-5 0v-1zm7 0v1a.5.5 0 0 0 1 0v-1h2v1a2.5 2.5 0 0 1-5 0v-1zm-6.1 5.134l.5.865C9.922 15.4 10.89 16 12 16s2.08-.601 2.6-1.5l.5-.866l1.731 1.001l-.5.866A5 5 0 0 1 12 18a5 5 0 0 1-4.33-2.5l-.501-.865z"/>`), g.Group(children),
	)
}

func Ferocious(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5.769-3.866l3.464 2l-1 1.732l-3.464-2zm11.464 1.732l-3.464 2l-1-1.732l3.464-2zM16 12.791V17.1l-8.095-.771l-.092-1.978z"/>`), g.Group(children),
	)
}

func FerrisWheel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v2.062a8 8 0 0 1 2.102.562l1.03-1.785l1.733 1l-1.031 1.786c.58.44 1.098.959 1.54 1.539l1.784-1.03l1 1.731l-1.783 1.03c.279.66.471 1.367.563 2.105H22v2h-2.062a8 8 0 0 1-.563 2.104l1.783 1.029l-1 1.732l-1.784-1.03A8 8 0 0 1 16 17.929V21h3v2H5v-2h3v-3.07a8 8 0 0 1-2.375-2.097l-1.787 1.032l-1-1.732l1.786-1.031A8 8 0 0 1 4.062 12H2v-2h2.062a8 8 0 0 1 .562-2.103L2.837 6.865l1-1.732l1.789 1.032c.44-.58.958-1.098 1.539-1.539L6.133 2.84l1.732-1l1.03 1.786A8 8 0 0 1 11 3.062V1zm-3 17.748V21h4v-2.252a8 8 0 0 1-4 0m5.439-2.83a6 6 0 0 0 1.655-1.746l.205-.356c.448-.84.7-1.799.7-2.816a6 6 0 0 0-.72-2.856l-.161-.278a6 6 0 0 0-2.054-2.026l-.138-.08a6 6 0 0 0-6.365 10.157L12 8.668zm-5.071.858A6 6 0 0 0 12 17a6 6 0 0 0 1.632-.224L12 13.334z"/>`), g.Group(children),
	)
}

func File(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm2 2v18h14V9h-6V3zm10 .414V7h3.586z"/>`), g.Group(children),
	)
}

func FileAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V12h-2V9h-6V3H5v18h7v2H3zm12 2.414V7h3.586zM20 14v4h4v2h-4v4h-2v-4h-4v-2h4v-4z"/>`), g.Group(children),
	)
}

func FileAddOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm15.586 6L15 3.414V7zM13 3H5v18h14V9h-6zm0 8v3h3v2h-3v3h-2v-3H8v-2h3v-3z"/>`), g.Group(children),
	)
}

func FileAttachment(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V12h-2V9h-6V3H5v18h7v2H3zm12 2.414V7h3.586zM14 15.5a2.5 2.5 0 0 1 5 0V20h-2v-4.5a.5.5 0 0 0-1 0V20a2 2 0 1 0 4 0v-4h2v4a4 4 0 0 1-8 0z"/>`), g.Group(children),
	)
}

func FileBlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11h-2V9h-6V3H5v18h6v2H3zm12 2.414V7h3.586zM18 14.5a3.5 3.5 0 0 0-3.08 5.165l4.745-4.744A3.5 3.5 0 0 0 18 14.5m3.08 1.835l-4.745 4.744a3.5 3.5 0 0 0 4.745-4.745M12.5 18a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0"/>`), g.Group(children),
	)
}

func FileCode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.001 1h12.414l5.586 5.587V13h-2V9h-6V3h-8v18h6v2h-8zm12 2.415V7h3.586zm8.663 11.841l-2.776 2.749l2.776 2.748l-1.407 1.421l-4.212-4.17l4.212-4.17zM13 21h4.5v2H13z"/>`), g.Group(children),
	)
}

func FileCodeOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm2 2v18h14V9h-6V3zm10 .414V7h3.586zm-3.793 8.793L9.414 14l1.793 1.793l-1.414 1.414L6.586 14l3.207-3.207zm3-1.414L17.414 14l-3.207 3.207l-1.414-1.414L14.586 14l-1.793-1.793z"/>`), g.Group(children),
	)
}

func FileCopy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 1h9.414L21 6.586V19H6zm2 2v14h11V9h-6V3zm7 .414V7h3.586zM4 5v16h11v2H2V5z"/>`), g.Group(children),
	)
}

func FileDownload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11.5h-2V9h-6V3H5v18h8v2H3zm12 2.414V7h3.586zM20 13v7.11l2.508-2.48l1.406 1.422L19 23.91l-4.914-4.858l1.406-1.422L18 20.11V13z"/>`), g.Group(children),
	)
}

func FileExcel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm2 2v18h14V9h-6V3zm10 .414V7h3.586zM15 10v1.646a2 2 0 0 1-.612 1.44l-.947.914l.947.914a2 2 0 0 1 .612 1.44V18h-2v-1.646l-1-.965l-1 .965V18H9v-1.646a2 2 0 0 1 .612-1.44l.947-.914l-.947-.914A2 2 0 0 1 9 11.646V10h2v1.646l1 .965l1-.965V10z"/>`), g.Group(children),
	)
}

func FileExport(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V12h-2V9h-6V3H5v18h6v2H3zm12 2.414V7h3.586zm4.05 10.674l4.858 4.914l-4.858 4.914l-1.422-1.406l2.48-2.508h-7.11v-2h7.11l-2.48-2.508z"/>`), g.Group(children),
	)
}

func FileIcon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11h-2V9h-6V3H5v18h16v2H3zm12 2.414V7h3.586zM8 12v8H6v-8zm.5 2a2 2 0 0 1 2-2H12v2h-1.5v4H12v2h-1.5a2 2 0 0 1-2-2zm4 0a2 2 0 0 1 2-2h1a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-1a2 2 0 0 1-2-2zm5.5-2h3a2 2 0 0 1 2 2v6h-2v-6h-1v6h-2zm-3.5 2v4h1v-4z"/>`), g.Group(children),
	)
}

func FileImage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm2 2v18h14V9h-6V3zm10 .414V7h3.586zM8 11h2.004v2.004H8zm4.728 1.938L16.79 17l-1.414 1.414l-2.648-2.648l-1.612 1.613l-.652-.784l-.7.677l-1.142 1.142L7.208 17l1.154-1.154l2.262-2.188l.628.756z"/>`), g.Group(children),
	)
}

func FileImport(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23h-9v-2h7V9h-6V3H5v10H3zm12 2.414V7h3.586zM7.05 14.088l4.858 4.914l-4.858 4.914l-1.422-1.406l2.48-2.508H.997v-2h7.11l-2.48-2.508z"/>`), g.Group(children),
	)
}

func FileLocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11h-2V9h-6V3H5v18h6.5v2H3zm12 2.414V7h3.586zM18 14.5c.69 0 1.25.56 1.25 1.25v.75h-2.5v-.75c0-.69.56-1.25 1.25-1.25m3.25 2v-.75a3.25 3.25 0 0 0-6.5 0v.75h-1.251V23h9v-6.5zm-.751 2V21h-5v-2.5z"/>`), g.Group(children),
	)
}

func FileMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm15.586 6L15 3.414V7zM13 3H5v18h14V9h-6zM8 14h8v2H8z"/>`), g.Group(children),
	)
}

func FileMusic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm15.586 6L15 3.414V7zM13 3H5v18h14V9h-6zm-1 7h4v2h-2v4a2.5 2.5 0 1 1-2-2.45zm0 6a.5.5 0 1 0-1 0a.5.5 0 0 0 1 0"/>`), g.Group(children),
	)
}

func FileOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm15.586 6L15 3.414V7zM13 3H5v18h14V9h-6zm-6 9h10v2H7zm0 4h10v2H7z"/>`), g.Group(children),
	)
}

func FileOnenote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm2 2v18h14V9h-6V3zm10 .414V7h3.586zM9 10h4a2 2 0 0 1 2 2v6h-2v-6h-2v6H9z"/>`), g.Group(children),
	)
}

func FileOutlook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm2 2v18h14V9h-6V3zm10 .414V7h3.586zM8.5 12a2 2 0 0 1 2-2h3a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-3a2 2 0 0 1-2-2zm5 0h-3v4h3z"/>`), g.Group(children),
	)
}

func FilePaste(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm16 8h-6V3H5v18h14zm-4-5.583V7h3.585zM7 10h2.6v2H9v.6H7zm3.4 0H13v2.6h-2V12h-.6zM9 14v-.6H7V16h2.6v-2zm2 0h6v6h-6zm2 2v2h2v-2z"/>`), g.Group(children),
	)
}

func FilePdf(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11h-2V9h-6V3H5v18h16v2H3zm12 2.414V7h3.586zM6 12h3.714c.71 0 1.286.576 1.286 1.286v2.428c0 .71-.576 1.286-1.286 1.286H8v3H6zm2 3h1v-1H8zm3.5-3h3.714c.71 0 1.286.576 1.286 1.286v5.428c0 .71-.576 1.286-1.286 1.286H11.5zm2 2v4h1v-4zm3.5-.714c0-.71.576-1.286 1.286-1.286h3.38v2H19v1h2.667v2H19v3h-2z"/>`), g.Group(children),
	)
}

func FilePowerpoint(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm2 2v18h14V9h-6V3zm10 .414V7h3.586zM9 10h4a2 2 0 0 1 2 2v1.6a2 2 0 0 1-2 2h-2V18H9zm2 3.6h2V12h-2z"/>`), g.Group(children),
	)
}

func FileRestore(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V12h-2V9h-6V3H5v18h7v2H3zm12 2.414V7h3.586zM18.414 13l-1 1H18a5 5 0 1 1-5 5v-1h2v1a3 3 0 1 0 3-3h-.586l1 1L17 18.414L13.586 15L17 11.586z"/>`), g.Group(children),
	)
}

func FileSafety(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11.5h-2V9h-6V3H5v18h6.5v2H3zm12 2.414V7h3.586zM13.498 13.5h9v5.634a3 3 0 0 1-1.36 2.511l-3.14 2.052l-3.14-2.052a3 3 0 0 1-1.36-2.511zm2 2v3.634a1 1 0 0 0 .453.837l2.047 1.337l2.047-1.337a1 1 0 0 0 .453-.837V15.5z"/>`), g.Group(children),
	)
}

func FileSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11.5h-2V9h-6V3H5v18h7.5v2H3zm12 2.414V7h3.586zM17.25 14.5a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5m-4.75 2.75a4.75 4.75 0 1 1 8.74 2.578l1.674 1.671l-1.413 1.415l-1.675-1.673A4.75 4.75 0 0 1 12.5 17.25"/>`), g.Group(children),
	)
}

func FileSetting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11h-2V9h-6V3H5v18h6v2H3zm12 2.414V7h3.586zm4 9.336v1.376c.715.184 1.352.56 1.854 1.072l1.193-.689l1 1.732l-1.192.688a4 4 0 0 1 0 2.142l1.192.688l-1 1.732l-1.193-.689A4 4 0 0 1 19 21.874v1.376h-2v-1.376a4 4 0 0 1-1.854-1.072l-1.193.689l-1-1.732l1.192-.688a4 4 0 0 1 0-2.142l-1.192-.688l1-1.732l1.193.688A4 4 0 0 1 17 14.127V12.75zm-2.751 4.283a2 2 0 0 0-.25.967c0 .35.091.68.25.967l.036.063a2 2 0 0 0 3.43 0l.036-.063A2 2 0 0 0 20 18c0-.35-.09-.68-.249-.967l-.036-.063a2 2 0 0 0-3.43 0z"/>`), g.Group(children),
	)
}

func FileTeams(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm2 2v18h14V9h-6V3zm10 .414V7h3.586zM9 10h6v2h-2v6h-2v-6H9z"/>`), g.Group(children),
	)
}

func FileUnknown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11.5h-2V9h-6V3H5v18h9.5v2H3zm12 2.414V7h3.586zM18 15c-.93 0-1.5.656-1.5 1.249v1h-2v-1C14.5 14.358 16.17 13 18 13s3.5 1.358 3.5 3.249a3.13 3.13 0 0 1-1.027 2.3L19 19.939v.683h-2v-1.546l2.112-1.993c.256-.235.388-.53.388-.834c0-.593-.57-1.249-1.5-1.249m-1 6.996h2.003V24h-2.004z"/>`), g.Group(children),
	)
}

func FileUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11h-2V9h-6V3H5v18h6.5v2H3zm12 2.414V7h3.586zM18 14.5c-.69 0-1.25.56-1.25 1.25v.75h5.749V23h-9v-6.5h1.251v-.75a3.25 3.25 0 0 1 5.541-2.305l.71.705l-1.41 1.418l-.71-.705A1.24 1.24 0 0 0 18 14.5m-2.501 4V21h5v-2.5z"/>`), g.Group(children),
	)
}

func FileWord(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm2 2v18h14V9h-6V3zm10 .414V7h3.586zM10 11v5h1v-5h2v5h1v-5h2v5a2 2 0 0 1-2 2h-4a2 2 0 0 1-2-2v-5z"/>`), g.Group(children),
	)
}

func FileZip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm2 2v18h14V9h-6V3h-1.996v2.004h-2V7h2v2.004h-2V11h2v2.004h-2v2H7V13h2v-1.996H7V9h2V7.004H7V5h2V3zm10 .414V7h3.586z"/>`), g.Group(children),
	)
}

func FillColor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9.525.626l.707.707l2.475 2.475l8.485 8.485l.708.707l-.707.707l-8.486 8.486L12 22.9l-.707-.707l-8.485-8.486L2.1 13l.707-.707l7.778-7.778l-1.768-1.768l-.707-.707zM5.93 12h12.142l-6.07-6.07zm12.142 2H5.93L12 20.071zm2.679 3.39l.784.99l.53.67c.581.733.581 1.847 0 2.58a1.68 1.68 0 0 1-1.314.657c-.53 0-1-.26-1.314-.657c-.581-.733-.581-1.847 0-2.58l.53-.67z"/>`), g.Group(children),
	)
}

func FillColorOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 .052L23.948 12L22 13.948V14h-.052l-7.12 7.12a4 4 0 0 1-5.656 0L2.88 14.829a4 4 0 0 1 0-5.657l4.438-4.439L4.586 2L6 .586l2.733 2.733zM8.733 6.148l-4.438 4.438c-.39.39-.586.902-.586 1.414h17.41L12 2.88l-1.852 1.853L13.414 8L12 9.414zM4.88 14l5.706 5.706a2 2 0 0 0 2.828 0L19.12 14zm16.87 4.389l1.314 1.66c.581.734.581 1.848 0 2.581a1.68 1.68 0 0 1-1.314.657c-.53 0-1-.26-1.314-.657c-.581-.733-.581-1.847 0-2.581z"/>`), g.Group(children),
	)
}

func Film(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 1 9 9c0 2.453-.98 4.676-2.573 6.3A8.97 8.97 0 0 1 12 21a9 9 0 1 1 0-18m6.326 18q.826-.582 1.53-1.3A10.97 10.97 0 0 0 23 12c0-6.075-4.925-11-11-11S1 5.925 1 12s4.925 11 11 11h12v-2zM12 6a1 1 0 1 0 0 2a1 1 0 0 0 0-2M9 7a3 3 0 1 1 6 0a3 3 0 0 1-6 0m-2 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0m13-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0m-2 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0"/>`), g.Group(children),
	)
}

func FilmOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v2h2V4zm4 0v7h8V4zm10 0v2h2V4zm2 4h-2v3h2zm0 5h-2v3h2zm0 5h-2v2h2zm-4 2v-7H8v7zM6 20v-2H4v2zm-2-4h2v-3H4zm0-5h2V8H4z"/>`), g.Group(children),
	)
}

func Filter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.57 3h18.86l-6.93 9.817V21h-5v-8.183zm3.86 2l5.07 7.183V19h1v-6.817L17.57 5z"/>`), g.Group(children),
	)
}

func FilterClear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.57 3h18.86l-6.93 9.817V21h-5v-8.183zm3.86 2l5.07 7.183V19h1v-6.817L17.57 5zm11.45 8.465L20 15.586l2.122-2.121l1.414 1.414L21.415 17l2.121 2.122l-1.414 1.414L20 18.414l-2.12 2.122l-1.415-1.415l2.121-2.12l-2.121-2.122z"/>`), g.Group(children),
	)
}

func FilterOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1.586L21.414 20L20 21.414l-5.5-5.5V21h-5v-8.182L4.933 6.347L1.586 3zm8.5 11.328V19h1v-5.086zM6.545 3H21.43l-6.034 8.549l-1.634-1.153L17.57 5H9.51l-.899.018z"/>`), g.Group(children),
	)
}

func FilterOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.55 5.5c.039.038.098.089.189.152c.306.213.815.445 1.533.656C7.695 6.727 9.72 7 12 7s4.305-.273 5.728-.692c.718-.21 1.227-.443 1.533-.656c.091-.063.15-.114.188-.152a1.4 1.4 0 0 0-.188-.152c-.306-.213-.814-.445-1.533-.656C16.305 4.273 14.28 4 12 4s-4.305.273-5.728.692c-.718.21-1.227.443-1.533.656c-.091.063-.15.114-.188.152m12.624 3.007C15.694 8.821 13.906 9 12 9s-3.693-.179-5.174-.493L11 13.645v6.184A3 3 0 0 0 13 17v-3.355zM21.5 5.5v.855l-6.5 8V17a5 5 0 0 1-5 5H9v-7.645l-6.5-8V5.5c0-.84.572-1.43 1.097-1.794c.554-.385 1.29-.692 2.11-.933C7.36 2.287 9.585 2 12 2s4.64.287 6.293.773c.82.241 1.556.548 2.11.933c.524.365 1.097.955 1.097 1.794"/>`), g.Group(children),
	)
}

func FilterThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12.02.002l3.49 7.988l7.987 3.49l-7.988 3.49l-3.49 7.989l-3.49-7.989l-7.988-3.49l7.988-3.49zm9.304 3.323l-1.568.78l1.568.781l.781 1.57l.781-1.57l1.57-.78l-1.57-.781l-.78-1.57zM12.02 4.998l-1.97 4.511l-4.512 1.971l4.511 1.971l1.971 4.512l1.971-4.512l4.512-1.97l-4.512-1.972zm7.316 9.758l1.3 2.61l2.61 1.3l-2.61 1.299l-1.3 2.61l-1.3-2.61l-2.61-1.3l2.61-1.299z"/>`), g.Group(children),
	)
}

func FilterTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m20.888 2.27l-1.57.78l1.57.782l.78 1.569l.781-1.57l1.57-.78l-1.57-.781l-.78-1.569zm-12.93-.755l1.098 2.204l2.204 1.097l-2.204 1.097l-1.097 2.204l-1.097-2.204l-2.204-1.097l2.204-1.097zm9.28 1.887l5.15 5.149L7.297 23.64L2.15 18.491zm-2.004 4.833l2.32 2.32l2.005-2.004l-2.32-2.32zm.906 3.735l-2.32-2.32l-8.842 8.841l2.32 2.32z"/>`), g.Group(children),
	)
}

func Fingerprint(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 11C2 5.477 6.477 1 12 1s10 4.477 10 10v3c0 2.64-.57 5.15-1.592 7.412l-.413.911l-1.822-.824l.412-.911A15.9 15.9 0 0 0 20 14v-3a8 8 0 1 0-16 0v3H2zm4.5 0a5.5 5.5 0 1 1 11 0v3c0 2.8-.853 5.403-2.314 7.56l-.561.829l-1.656-1.122l.56-.828A11.44 11.44 0 0 0 15.5 14v-3a3.5 3.5 0 1 0-7 0v3A4.5 4.5 0 0 1 4 18.5H2v-2h2A2.5 2.5 0 0 0 6.5 14zm6.5-1v4c0 3.39-1.875 6.34-4.639 7.874l-.874.486l-.97-1.75l.874-.484A7 7 0 0 0 11 14v-4z"/>`), g.Group(children),
	)
}

func FingerprintOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 12C2 6.477 6.477 2 12 2s10 4.477 10 10v9h-2v-9a8 8 0 1 0-16 0v9H2zm4.5 0a5.5 5.5 0 1 1 11 0v9h-2v-9a3.5 3.5 0 1 0-7 0v9h-2zm6.5-1v10h-2V11z"/>`), g.Group(children),
	)
}

func FingerprintThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 11C2 5.477 6.477 1 12 1c1.484 0 2.894.324 4.163.905l.909.417l-.833 1.818l-.91-.417A8 8 0 0 0 4 11v3H2.001zm16.957-7.253l.649.76A9.96 9.96 0 0 1 22 11v3c0 2.64-.57 5.15-1.592 7.412l-.413.911l-1.822-.824l.412-.911A15.9 15.9 0 0 0 20 14v-3a7.96 7.96 0 0 0-1.915-5.193l-.65-.76zM11 5.5h1a5.5 5.5 0 0 1 5.5 5.5v1h-2v-1A3.5 3.5 0 0 0 12 7.5h-1zM9.83 8.079l-.61.793A3.48 3.48 0 0 0 8.5 11v2h-2v-2a5.48 5.48 0 0 1 1.134-3.345l.608-.793zM13 10v4c0 3.39-1.875 6.34-4.639 7.874l-.874.486l-.97-1.75l.874-.484A7 7 0 0 0 11 14v-4zm4.536 4.082l-.079.997a13.43 13.43 0 0 1-2.271 6.482l-.561.828l-1.656-1.122l.56-.828a11.4 11.4 0 0 0 1.935-5.518l.079-.997zm-9.163 1.381l-.516.857A4.5 4.5 0 0 1 4 18.5H2v-2h2c.908 0 1.704-.484 2.143-1.212l.517-.857z"/>`), g.Group(children),
	)
}

func FingerprintTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 11c0-5.523-4.477-10-10-10S2 5.477 2 11v3c0 2.64.57 5.15 1.592 7.412l.413.911l1.822-.824l-.412-.911A15.9 15.9 0 0 1 4 14v-3a8 8 0 1 1 16 0v3h2zm-4.5 0a5.5 5.5 0 1 0-11 0v3c0 2.8.853 5.403 2.314 7.56l.561.829l1.656-1.122l-.56-.828A11.44 11.44 0 0 1 8.5 14v-3a3.5 3.5 0 1 1 7 0v3a4.5 4.5 0 0 0 4.5 4.5h2v-2h-2a2.5 2.5 0 0 1-2.5-2.5zM11 10v4c0 3.39 1.875 6.34 4.639 7.874l.874.486l.97-1.75l-.874-.484A7 7 0 0 1 13 14v-4z"/>`), g.Group(children),
	)
}

func Fish(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.497 4.564c1.649-.906 3.859-1.137 6.669-.694c.119.685.221 1.711.147 2.86c-.102 1.572-.53 3.278-1.602 4.656l-.165.212l-.036.263v.002l-.002.014l-.013.074a9.3 9.3 0 0 1-1.294 3.217l-.84-1.688l-.96.72C13.65 15.513 10.903 16 9 16H8v1c0 .77-.004 1.293-.106 1.804a4 4 0 0 1-.147.53l-4.011-4.012a8 8 0 0 1 .978-.209A10 10 0 0 1 5.985 15H7v-1c0-2.697.864-3.993 1.83-5.442L9.202 8L7.428 5.339l.112-.04a9.7 9.7 0 0 1 1.653-.372c.609-.087 1.228-.13 1.773-.122q.303.004.548.026c.09.768.373 1.643.861 2.475c.725 1.236 1.938 2.442 3.774 3.13l.936.351l.703-1.872l-.937-.351c-1.364-.512-2.234-1.39-2.75-2.27c-.385-.655-.557-1.28-.604-1.73m6.947 7.845c1.285-1.759 1.752-3.81 1.865-5.55c.117-1.806-.14-3.371-.344-4.121l-.164-.605l-.616-.116c-3.425-.643-6.471-.492-8.855.91a7.7 7.7 0 0 0-1.338-.122c-.66-.009-1.383.042-2.083.143c-.698.1-1.397.252-2.004.455c-.575.193-1.193.471-1.612.89l-.58.58L6.8 8.003c-.813 1.256-1.6 2.711-1.767 5.054q-.287.029-.622.08c-.857.131-2.032.409-2.965 1.03l-1.015.678l7.725 7.725l.676-1.015c.563-.845.87-1.59 1.024-2.359c.083-.413.118-.823.133-1.231c1.704-.117 3.837-.545 5.612-1.523l1.12 2.25l.983-.983c1.188-1.189 1.88-2.582 2.273-3.653a11 11 0 0 0 .467-1.646M17.5 4.58l1.417 1.417L17.5 7.414l-1.417-1.417z"/>`), g.Group(children),
	)
}

func Flag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 2h8.414l2 2h8.014l-2.364 6.5l2.364 6.5h-8.842l-2-2H5v7.5H3zm2 11h6.414l2 2h5.158l-1.636-4.5L18.572 6h-5.986l-2-2H5z"/>`), g.Group(children),
	)
}

func FlagFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h14v1h6v14h-8v-1H4v6H2zm2 12h10V4H4zm12-9v10h4V5z"/>`), g.Group(children),
	)
}

func FlagOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 2h8.414l2 2h8.014l-2.364 6.5l2.364 6.5h-8.842l-2-2H5v7.5H3zm2 11h6.414l2 2h5.158L17.3 11.5h-4.714l-2-2H5zm0-5.5h6.414l2 2H17.3L18.572 6h-5.986l-2-2H5z"/>`), g.Group(children),
	)
}

func FlagThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23 10L8 18.66V1.34zM10 4.804v10.392L19 10zM4 1h2v22H4z"/>`), g.Group(children),
	)
}

func FlagTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 2h8.414l2 2h8.014l-2.364 6.5l2.364 6.5h-8.842l-2-2H5v7.5H3zm2 11h2V4H5zm4-9v9h2.414l2 2h5.158l-1.636-4.5L18.572 6h-5.986l-2-2z"/>`), g.Group(children),
	)
}

func Flashlight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 0h16v7.414l-3 3V24H7V10.414l-3-3zm2 2v1.5h12V2zm12 3.5H6v1.086l3 3V22h6V9.586l3-3zM11 10h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func FlightLanding(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9.302 1.203l3.716.995l2.566 8.97l4.958 1.329a2.5 2.5 0 1 1-1.294 4.83L1.553 12.584l.46-7.124l3.173.85l.642 2.243l3.335.894zm1.956 2.594l-.138 8.246l-6.904-1.85l-.412-1.441l-.15 2.325l16.112 4.318a.5.5 0 1 0 .259-.966l-6.053-1.622l-2.566-8.97zM2 19h20v2H2z"/>`), g.Group(children),
	)
}

func FlightTakeoff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m8.324 2.387l6.707 6.486l4.96-1.33a2.5 2.5 0 0 1 1.293 4.83L3.59 17.115l-3.164-6.4l3.174-.85l1.677 1.622l3.335-.894l-4.002-7.21zM7.6 4.652l4.003 7.21l-6.904 1.85l-1.077-1.042l1.033 2.089l16.112-4.317a.5.5 0 0 0-.26-.966l-6.052 1.621l-6.707-6.485zM2 19h20v2H2z"/>`), g.Group(children),
	)
}

func FlipSmilingFace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m6.67-3.5A5 5 0 0 1 12 6a5 5 0 0 1 4.33 2.5l.501.865l-1.731 1.001l-.5-.865A3 3 0 0 0 12 8a3 3 0 0 0-2.6 1.5l-.5.866l-1.731-1.001zM9 12v4H7v-4zm8 0v4h-2v-4z"/>`), g.Group(children),
	)
}

func FlipToBack(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 3h3.5v2H9v1.5H7zm5.5 0h3v2h-3zm5 0H21v3.5h-2V5h-1.5zM5 7.5V19h11.5v2H3V7.5zm4 1v3H7v-3zm12 0v3h-2v-3zm-12 5V15h1.5v2H7v-3.5zm12 0V17h-3.5v-2H19v-1.5zM12.5 15h3v2h-3z"/>`), g.Group(children),
	)
}

func FlipToFront(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 3h14v14H7zm2 2v10h10V5zM5 7.5v3H3v-3zm0 5v3H3v-3zm0 5V19h1.5v2H3v-3.5zm6.5 1.5v2h-3v-2zm2 0h3v2h-3z"/>`), g.Group(children),
	)
}

func Focus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v3.062A8.004 8.004 0 0 1 19.938 11H23v2h-3.062A8.004 8.004 0 0 1 13 19.938V23h-2v-3.062A8.004 8.004 0 0 1 4.062 13H1v-2h3.062A8.004 8.004 0 0 1 11 4.062V1zm-1 5a6 6 0 1 0 0 12a6 6 0 0 0 0-12m-2 6a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func Fog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.1 4C8.634 4 6.7 5.916 6.7 8.2q0 .394.073.765l.229 1.19H5.79C4.839 10.155 4 11.018 4 12v1H2v-1c0-1.657 1.124-3.185 2.701-3.68L4.7 8.2C4.7 4.74 7.601 2 11.1 2c.546 0 1.078.066 1.586.192c2.147.53 3.88 2.12 4.533 4.187a5.6 5.6 0 0 1 2.01.657C20.942 7.986 22 10.086 22 12v1h-2v-1c0-1.298-.756-2.669-1.741-3.216a3.6 3.6 0 0 0-1.818-.45l-.854.013l-.147-.842c-.282-1.62-1.55-2.956-3.234-3.371A4.6 4.6 0 0 0 11.1 4M2 14h6v2H2zm8 0h12v2H10zm7 4h5v2h-5zM2 20v-2h13v2z"/>`), g.Group(children),
	)
}

func FogNight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10.916.585l-.45 1.675c-.3 1.116-.27 2.337.07 3.604a7 7 0 0 0 9.84 4.476l1.541-.752l-.103 1.71c-.063 1.034-.375 2.06-.92 3.15l-.446.894l-1.79-.895l.448-.894q.199-.398.338-.76A9 9 0 0 1 8.604 6.38a9.6 9.6 0 0 1-.338-2.852a8 8 0 0 0-4.164 9.235c.017.061.085.251.174.48l.105.267l.033.082l.01.022l.002.007l.378.926l-1.852.756l-.378-.926l-.003-.01l-.01-.024l-.037-.09l-.113-.288a8 8 0 0 1-.24-.684C.74 7.947 3.906 2.464 9.24 1.034zM2.001 17h6v2h-6zm8 0h12v2h-12zm7 4h5v2h-5zm-15 0h13v2h-13z"/>`), g.Group(children),
	)
}

func FogSunny(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v3h-2V1zm7.485 3.928L18.364 7.05L16.95 5.636l2.121-2.122zM4.93 3.514l2.12 2.122L5.636 7.05L3.515 4.929zM12 8a4 4 0 0 0-3.668 5.6l.4.916l-1.832.8l-.4-.916a6 6 0 1 1 11 0l-.4.917l-1.833-.801l.4-.916A4 4 0 0 0 12 8M1 11h3v2H1zm19 0h3v2h-3zM2 17h6v2H2zm8 0h12v2H10zm7 4h5v2h-5zM2 21h13v2H2z"/>`), g.Group(children),
	)
}

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1zm2 2V19h18V7H10.52l-2-2.5z"/>`), g.Group(children),
	)
}

func FolderAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v7h-2V7H10.52l-2-2.5H3V19h11v2H1zM21 14v3h3v2h-3v3h-2v-3h-3v-2h3v-3z"/>`), g.Group(children),
	)
}

func FolderAddOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1zm2 2V19h18V7H10.52l-2-2.5zM13 9v3h3v2h-3v3h-2v-3H8v-2h3V9z"/>`), g.Group(children),
	)
}

func FolderBlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v5h-2V7H10.52l-2-2.5H3V19h8.5v2H1zm17.5 11a3.5 3.5 0 0 0-3.08 5.165l4.745-4.744A3.5 3.5 0 0 0 18.5 13.5m3.08 1.835l-4.745 4.744a3.5 3.5 0 0 0 4.745-4.745M13 17a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0"/>`), g.Group(children),
	)
}

func FolderDetails(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1zm2 2V19h18V7H10.52l-2-2.5zm3.998 7.498h2.004v2.004H6.998zm4 0h2.004v2.004h-2.004zm4 0h2.004v2.004h-2.004z"/>`), g.Group(children),
	)
}

func FolderExport(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v8h-2V7H10.52l-2-2.5H3V19h8.5v2H1zm18.05 10.588l4.858 4.914l-4.858 4.914l-1.422-1.406l2.48-2.508h-7.11v-2h7.11l-2.48-2.508z"/>`), g.Group(children),
	)
}

func FolderImport(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2.5h8.48l2 2.5H24v16H13v-2h9V7H11.52l-2-2.5H4v8.25H2zm5.05 10.588l4.858 4.914l-4.858 4.914l-1.422-1.406l2.48-2.508H.997v-2h7.11l-2.48-2.508z"/>`), g.Group(children),
	)
}

func FolderLocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v5h-2V7H10.52l-2-2.5H3V19h10v2H1zm18.5 11c.69 0 1.25.56 1.25 1.25v.75h-2.5v-.75c0-.69.56-1.25 1.25-1.25m3.25 2v-.75a3.25 3.25 0 0 0-6.5 0v.75h-1.251V22h9v-6.5zm-.751 2V20h-5v-2.5z"/>`), g.Group(children),
	)
}

func FolderMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1zm2 2V19h18V7H10.52l-2-2.5zM8 12h8v2H8z"/>`), g.Group(children),
	)
}

func FolderMove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1zm2 2V19h18V7H10.52l-2-2.5z"/><path fill="currentColor" d="M8 11.999h4.657l-1.466-1.53l1.445-1.384l3.749 3.914l-3.75 3.914l-1.444-1.384L12.657 14H8z"/>`), g.Group(children),
	)
}

func FolderOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9.482 2.5l2 2.5H14v2h-3.48l-2-2.5H3v11.501l.003 1.412l-2 2L1 16.002V2.5zM21.414 4l-1 1H23v15H5.414L3 22.414L1.586 21L20 2.586zm-3 3l-11 11H21V7z"/>`), g.Group(children),
	)
}

func FolderOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1zm2 2v4.25h18V7H10.52l-2-2.5zm18 6.25H3V19h18z"/>`), g.Group(children),
	)
}

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h10.48l2 2.5H23v2H12.52l-2-2.5H1zm0 4h8.48l2 2.5H23v12H1zm2 2V19h18v-8H10.52l-2-2.5z"/>`), g.Group(children),
	)
}

func FolderOpenOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m23.844 9l-3.6 12H0V2.5h7.362l3 2.5H20.5v4zM18.5 9V7H9.638l-3-2.5H2v8.687L3.256 9zM2.344 19h16.412l2.4-8H4.744z"/>`), g.Group(children),
	)
}

func FolderSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v7.5h-2V7H10.52l-2-2.5H3V19h8.5v2H1zm16.25 12a2.75 2.75 0 0 1 1.946 4.693l-.008.008A2.75 2.75 0 1 1 17.25 14.5m3.992 5.325a4.75 4.75 0 1 0-1.414 1.415l1.67 1.674l1.416-1.412z"/>`), g.Group(children),
	)
}

func FolderSetting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v6h-2V7H10.52l-2-2.5H3V19h8v2H1zm18 9.25v1.376c.715.184 1.352.56 1.854 1.072l1.193-.689l1 1.732l-1.192.688a4 4 0 0 1 0 2.142l1.192.688l-1 1.732l-1.193-.689A4 4 0 0 1 19 20.874v1.376h-2v-1.376a4 4 0 0 1-1.854-1.072l-1.193.689l-1-1.732l1.192-.688a4 4 0 0 1 0-2.142l-1.192-.688l1-1.732l1.193.688A4 4 0 0 1 17 13.127V11.75zm-2.751 4.283a2 2 0 0 0-.25.967c0 .35.091.68.25.967l.036.063a2 2 0 0 0 3.43 0l.036-.063A2 2 0 0 0 20 17c0-.35-.09-.68-.249-.967l-.036-.063a2 2 0 0 0-3.43 0z"/>`), g.Group(children),
	)
}

func FolderShared(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1zm2 2V19h18V7H10.52l-2-2.5zm9 6a1 1 0 1 1 0 2a1 1 0 0 1 0-2m2.4 2.8a3 3 0 1 0-4.8 0A4 4 0 0 0 8 16.5v1h2v-1a2 2 0 1 1 4 0v1h2v-1c0-1.309-.628-2.47-1.6-3.2"/>`), g.Group(children),
	)
}

func FolderUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v5h-2V7H10.52l-2-2.5H3V19h10v2H1zm18.5 11c-.69 0-1.25.56-1.25 1.25v.75h5.749V22h-9v-6.5h1.251v-.75a3.25 3.25 0 0 1 5.541-2.305l.71.705l-1.41 1.418l-.71-.705a1.24 1.24 0 0 0-.881-.363m-2.501 4V20h5v-2.5z"/>`), g.Group(children),
	)
}

func FolderZip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2.5h8.48l2 2.5H23v16H1zm2 2V19h11v-2h2v-1.996h-2V13h2v-1.996h-2V9h2V7h-5.48l-2-2.5zm13.004 4.504V11h2v2.004h-2V15h2v2.004h-2V19H21V7h-2.996v2.004z"/>`), g.Group(children),
	)
}

func Forest(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7 .214l5 13l5-13L23.456 17H18v5h-2v-5H8v5H6v-5H.544zM8 15h2.544L8 8.385zM6 8.385L3.456 15H6zM13.456 15H16V8.385zM18 8.385V15h2.544z"/>`), g.Group(children),
	)
}

func Fork(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a2 2 0 1 0 0-4a2 2 0 0 0 0 4m1-5.874A4.002 4.002 0 0 1 12 23a4 4 0 0 1-1-7.874V13H7a3 3 0 0 1-3-3V8.874A4.002 4.002 0 0 1 5 1a4 4 0 0 1 1 7.874V10a1 1 0 0 0 1 1h10a1 1 0 0 0 1-1V8.874A4.002 4.002 0 0 1 19 1a4 4 0 0 1 1 7.874V10a3 3 0 0 1-3 3h-4zM4.997 7h.006a2 2 0 1 0-.006 0M19 7a2 2 0 1 0 0-4a2 2 0 0 0 0 4"/>`), g.Group(children),
	)
}

func Form(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v3h16V4zm16 5H4v11h16zM6 11.5h12v2H6zm0 4h8v2H6z"/>`), g.Group(children),
	)
}

func FormatHorizontalAlignBottom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 2v20H4V2zm14 0v20h-2V2zm-7 6v14h-2V8z"/>`), g.Group(children),
	)
}

func FormatHorizontalAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 2v20H4V2zm14 0v20h-2V2zm-7 3v14h-2V5z"/>`), g.Group(children),
	)
}

func FormatHorizontalAlignTop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 2v20H4V2zm7 0v14h-2V2zm7 0v20h-2V2z"/>`), g.Group(children),
	)
}

func FormatVerticalAlignCenter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 4h20v2H2zm3 7h14v2H5zm-2 7H2v2h20v-2z"/>`), g.Group(children),
	)
}

func FormatVerticalAlignLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 4h20v2H2zm0 7h14v2H2zm1 7H2v2h20v-2z"/>`), g.Group(children),
	)
}

func FormatVerticalAlignRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 4h20v2H2zm6 7h14v2H8zm-5 7H2v2h20v-2z"/>`), g.Group(children),
	)
}

func Forward(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.75 4.336v7l-7-7v15.328l7-7v7L20.414 12zM17.586 12l-2.836 2.836V9.164zm-7 0L7.75 14.836V9.164z"/>`), g.Group(children),
	)
}

func Frame(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v7h7V4zm9 0v7h7V4zm7 9h-7v7h7zm-9 7v-7H4v7z"/>`), g.Group(children),
	)
}

func FrameOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h4V4zm6 0v16h4V4zm6 0v16h4V4z"/>`), g.Group(children),
	)
}

func Fries(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.465.994L21 4.53l-.707.707l2.122 2.121l-2.122 2.121l2.122 2.122l-3.2 3.2l1.397 1.397L8.845 23.26L.741 15.156l7.06-11.768l1.514 1.514L13.222.994l2.122 2.121zM12.85 8.437l.707.707l4.614-4.614l-.707-.707zm6.028-1.786l-3.907 3.907l.707.707l3.907-3.907zm-1.786 6.028l.707.707l1.786-1.785l-.707-.707zm.294 3.122L8.198 6.613l-1.94 3.232l7.897 7.896zm-5 3l-7.189-7.188l-1.939 3.232l5.896 5.896zM10.729 6.315l.708.707l2.492-2.492l-.707-.707z"/>`), g.Group(children),
	)
}

func Fullscreen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4.5h7.5V12h-2V7.914L7.914 17.5H12v2H4.5V12h2v4.086L16.086 6.5H12z"/>`), g.Group(children),
	)
}

func FullscreenExit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m20.414 5l-4.5 4.5H20v2h-7.5V4h2v4.086l4.5-4.5zM4 12.5h7.5V20h-2v-4.086l-4.5 4.5L3.586 19l4.5-4.5H4z"/>`), g.Group(children),
	)
}

func FullscreenExitOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14 20v-6h6v2h-4v4zm-6 0v-4H4v-2h6v6zm12-10h-6V4h2v4h4zm-10 0H4V8h4V4h2z"/>`), g.Group(children),
	)
}

func FullscreenOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 4h6v2H6v4H4zm10 0h6v6h-2V6h-4zM6 14v4h4v2H4v-6zm14 0v6h-6v-2h4v-4z"/>`), g.Group(children),
	)
}

func FullscreenTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 4h6v2H7.414l3 3L9 10.414l-3-3V10H4zm10 0h6v6h-2V7.414l-3 3L13.586 9l3-3H14zm-3.586 11l-3 3H10v2H4v-6h2v2.586l3-3zM15 13.586l3 3V14h2v6h-6v-2h2.586l-3-3z"/>`), g.Group(children),
	)
}

func FunctionCurve(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 2v13h8.79q.177-.387.34-.8c1.193-3.015 1.741-6.74 1.871-9.252l.052-.998l1.997.103l-.051.999c-.137 2.64-.709 6.598-2.01 9.884l-.025.064H22v2h-3.985c-.476.843-1.029 1.61-1.67 2.243c-1.16 1.143-2.63 1.86-4.41 1.745c-1.57-.102-2.867-.773-3.935-1.766V22H6v-5H2v-2h3.203q-.1-.228-.19-.457c-1.3-3.227-1.874-6.997-2.011-9.488l-.055-.998l1.996-.11l.055.998A32 32 0 0 0 6 11.165V2zm.66 15c.958 1.188 2.09 1.907 3.405 1.992c1.1.072 2.036-.346 2.875-1.174q.367-.362.702-.818z"/>`), g.Group(children),
	)
}

func Functions(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 3h16v2H7.14l8.275 7l-8.275 7H20v2H4v-1.964L12.318 12L4 4.964z"/>`), g.Group(children),
	)
}

func FunctionsOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 5a2 2 0 0 1 2-2h8v2H4v6h8v2H4v8H2zm20 5v2.136a2 2 0 0 1-.726 1.542L19.07 15.5l2.204 1.822A2 2 0 0 1 22 18.864V21h-2v-2.136l-2.5-2.067l-2.5 2.067V21h-2v-2.136a2 2 0 0 1 .726-1.542L15.93 15.5l-2.204-1.822A2 2 0 0 1 13 12.136V10h2v2.136l2.5 2.067l2.5-2.067V10z"/>`), g.Group(children),
	)
}

func Gamepad(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.06 3h19.88l1.124 18H16.28l-1-3H8.721l-1 3H.936zm1.88 2l-.876 14H6.28l1-3h9.442l1 3h3.215L20.06 5zM9 7.5v2h2v2H9v2H7v-2H5v-2h2v-2zm7 0h2.004v2.004H16zm0 3.996h2.004V13.5H16z"/>`), g.Group(children),
	)
}

func GamepadOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 4a2 2 0 0 1 2-2h4v2h-4v2h9v14H2V6h9zM4 8v10h16V8zm5 2v2h2v2H9v2H7v-2H5v-2h2v-2zm6 0h2.003v2h2v2.004h-2v2H15v-2h-2V12h2zm.004 2.004V14H17v-1.996z"/>`), g.Group(children),
	)
}

func Gamma(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6.262 1.457l6.569 12.747L18.5 1.677l1.822.824l-6.434 14.215V22h-2v-5.257L6.182 5.668L4.38 8.862L2.638 7.88z"/>`), g.Group(children),
	)
}

func Garlic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.5.28V2l-.003.342c-.01.734-.02 1.632.307 2.707c.371 1.22 1.205 2.72 3.141 4.294c2.158 1.754 4.011 4.802 4.055 7.642c.024 1.588-.517 3.138-1.817 4.276C19.9 22.385 17.998 23 15.5 23c-.499 0-.956-.207-1.204-.334c-.718.246-1.504.334-2.296.334c-.918 0-1.65-.193-2.227-.478A8.8 8.8 0 0 1 7 23c-2.948 0-4.737-2.526-5.125-5.16c-.396-2.69.55-5.89 3.384-7.865a29 29 0 0 0 3.057-2.467c.407-.376.745-.712.995-.978q.159-.17.262-.29c-.008-.674-.108-1.16-.197-1.591l-.017-.08l-.145-.711zM9.025 9.56a31 31 0 0 1-2.622 2.056c-2.127 1.481-2.851 3.883-2.55 5.932C4.164 19.652 5.44 21 7 21a6 6 0 0 0 1.001-.086c-1.36-2.077-1.444-5.483.377-9.494c.27-.686.482-1.303.647-1.86m7.44 11.402c1.64-.134 2.722-.611 3.4-1.206c.788-.689 1.152-1.638 1.135-2.74c-.032-2.102-1.49-4.637-3.316-6.121c-2.225-1.809-3.298-3.637-3.794-5.263a9 9 0 0 1-.36-1.929l-2.074 1.18c.074.455.134 1.024.112 1.73c-.04 1.336-.37 3.104-1.339 5.564l-.01.024l-.01.024c-2.154 4.727-1.054 7.504.213 8.366c.287.195.77.409 1.578.409c.83 0 1.472-.12 1.943-.348c.497-.241.868-.632 1.07-1.335c.215-.753.242-1.893-.105-3.585c-.168-.82-.434-1.512-.737-2.304l-.111-.29c-.345-.905-.706-1.918-.896-3.213l-.145-.99L15 8.647l.144.99c.157 1.073.455 1.92.786 2.79l.11.287c.3.78.625 1.627.828 2.617c.378 1.844.409 3.347.068 4.537a4.5 4.5 0 0 1-.47 1.095"/>`), g.Group(children),
	)
}

func GenderFemale(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4.006A5 5 0 0 1 12.252 14h-.504A5 5 0 0 1 12 4.006m1 11.93a7.002 7.002 0 0 0-1-13.93a7 7 0 0 0-1 13.93V17H8v2h3v3h2v-3h3v-2h-3z"/>`), g.Group(children),
	)
}

func GenderMale(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.5 3H21v7.5h-2V6.414l-3.394 3.394a7 7 0 1 1-1.414-1.414L17.586 5H13.5zM10 9a5 5 0 1 0 0 10a5 5 0 0 0 0-10"/>`), g.Group(children),
	)
}

func GestureApplause(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 .586L5.414 3L4 4.414L1.586 2zM22.414 2L20 4.414L18.586 3L21 .586zM10.357 3.303a.65.65 0 0 0-.651.516l-2.28 10.674l-1.218 1.217l2.702 2.701l1.527-1.527a1.9 1.9 0 0 0 .56-1.35V3.954a.653.653 0 0 0-.64-.652M7.496 19.826l-2.702-2.702l-.643.644l2.701 2.701zM7.75 3.4A2.653 2.653 0 0 1 12 1.882a2.652 2.652 0 0 1 4.25 1.519l2.156 10.096l4.27 4.27l-5.53 5.53l-4.999-4.999q-.075-.076-.148-.156a4 4 0 0 1-.148.156l-4.999 5l-5.53-5.53l4.27-4.27zm6.543.418a.652.652 0 0 0-1.291.136v11.579c0 .506.201.992.56 1.35l1.527 1.527l2.7-2.701l-1.216-1.217zm4.911 13.305l-2.701 2.702l.643.643l2.702-2.701zM5.414 7L3 9.414L1.586 8L4 5.586zM20 5.586L22.414 8L21 9.414L18.586 7z"/>`), g.Group(children),
	)
}

func GestureClick(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 .5v3.414h-2V.5zM5 3.085l2.414 2.414L6 6.913L3.586 4.5zm13.414 1.414L16 6.913L14.586 5.5L17 3.085zM11 6.81a.756.756 0 0 0-.756.756v10.031l-4.498-1l-.1.151l3.424 4.456a.76.76 0 0 0 .6.295h7.086a.76.76 0 0 0 .718-.517l1.782-5.35a.76.76 0 0 0-.336-.894l-3.625-2.115a.76.76 0 0 0-.381-.103h-3.159V7.566A.756.756 0 0 0 11 6.81m-2.756.756a2.756 2.756 0 0 1 5.511 0v2.954h1.159c.488 0 .968.13 1.39.375l3.624 2.116a2.76 2.76 0 0 1 1.226 3.252l-1.783 5.35a2.76 2.76 0 0 1-2.615 1.886H9.67a2.76 2.76 0 0 1-2.185-1.076L3.187 16.83l.94-1.412a1.88 1.88 0 0 1 1.972-.792l2.145.477z"/>`), g.Group(children),
	)
}

func GestureDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.57 19.5a3 3 0 1 0 6 0V16h1.46a3 3 0 0 0 1.511-.409l4.13-2.409a3 3 0 0 0 1.334-3.54l-2.03-6.09A3 3 0 0 0 17.129 1.5H9.055a3 3 0 0 0-2.378 1.17L1.841 8.959l1.017 1.527a2 2 0 0 0 2.098.843l2.614-.581zm3 1a1 1 0 0 1-1-1V8.253L4.522 9.375L4.3 9.042L8.263 3.89a1 1 0 0 1 .792-.39h8.074a1 1 0 0 1 .949.684l2.03 6.09a1 1 0 0 1-.445 1.18l-4.13 2.41a1 1 0 0 1-.503.136h-3.46v5.5a1 1 0 0 1-1 1"/>`), g.Group(children),
	)
}

func GestureExpansion(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 0h5v5H9V3.414L5.414 7H7v2H2V4h2v1.586L7.586 2H6zm4.244 7.566a2.756 2.756 0 0 1 5.512 0v2.954h1.158c.488 0 .968.13 1.39.375l3.624 2.116a2.76 2.76 0 0 1 1.226 3.252l-1.782 5.35a2.76 2.76 0 0 1-2.616 1.886h-7.085a2.76 2.76 0 0 1-2.186-1.076L5.187 16.83l.94-1.412a1.88 1.88 0 0 1 1.972-.792l2.145.477zM13 6.81a.756.756 0 0 0-.756.756v10.031l-4.498-1l-.1.151l3.425 4.456a.76.76 0 0 0 .6.295h7.085a.76.76 0 0 0 .718-.517l1.782-5.35a.76.76 0 0 0-.336-.894l-3.625-2.115a.76.76 0 0 0-.381-.103h-3.159V7.566A.756.756 0 0 0 13 6.81"/>`), g.Group(children),
	)
}

func GestureLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.5 7.57a3 3 0 1 0 0 6H8v1.46a3 3 0 0 0 .409 1.511l2.409 4.13a3 3 0 0 0 3.54 1.334l6.09-2.03a3 3 0 0 0 2.052-2.846V9.055a3 3 0 0 0-1.17-2.378l-6.288-4.836l-1.527 1.017a2 2 0 0 0-.843 2.098l.581 2.614zm-1 3a1 1 0 0 1 1-1h11.247l-1.122-5.048l.333-.222l5.152 3.963a1 1 0 0 1 .39.792v8.074a1 1 0 0 1-.684.949l-6.09 2.03a1 1 0 0 1-1.18-.445l-2.41-4.13A1 1 0 0 1 10 15.03v-3.46H4.5a1 1 0 0 1-1-1"/>`), g.Group(children),
	)
}

func GestureLeftSlip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.4 5.2c4.308-3.239 10.455-2.898 14.377 1.024l.707.707l1.414-1.414l-.707-.707C16.651.27 9.587-.21 4.513 3.371v-1.87h-2v5.7H7.5v-2zm3.844 2.366a2.756 2.756 0 1 1 5.512 0v2.954h1.158c.488 0 .968.13 1.39.375l3.624 2.116a2.76 2.76 0 0 1 1.226 3.252l-1.782 5.35a2.76 2.76 0 0 1-2.616 1.886h-7.085a2.76 2.76 0 0 1-2.186-1.076L4.187 16.83l.94-1.412a1.88 1.88 0 0 1 1.972-.792l2.145.477zM12 6.81a.756.756 0 0 0-.756.756v10.031l-4.498-1l-.1.151l3.425 4.456a.76.76 0 0 0 .6.295h7.085a.76.76 0 0 0 .718-.517l1.783-5.35a.76.76 0 0 0-.337-.894l-3.625-2.115a.76.76 0 0 0-.381-.103h-3.158V7.566A.756.756 0 0 0 12 6.81"/>`), g.Group(children),
	)
}

func GesturePray(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.357 3.302a.65.65 0 0 0-.651.516l-2.28 10.675l-1.218 1.217l2.702 2.7l1.527-1.527a1.9 1.9 0 0 0 .56-1.35V3.953a.653.653 0 0 0-.64-.652M7.496 19.825l-2.702-2.701l-.643.643l2.701 2.702zM7.75 3.401A2.653 2.653 0 0 1 12 1.882a2.652 2.652 0 0 1 4.25 1.519l2.156 10.096l4.27 4.27l-5.53 5.53l-4.999-4.999q-.075-.076-.148-.156a4 4 0 0 1-.148.156l-4.999 5l-5.53-5.53l4.27-4.271zm6.543.417a.652.652 0 0 0-1.291.137v11.578c0 .507.201.993.56 1.35l1.527 1.528l2.7-2.701l-1.216-1.217zm4.911 13.306l-2.701 2.701l.643.644l2.702-2.702z"/>`), g.Group(children),
	)
}

func GesturePrayOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.036 4.986a.454.454 0 0 0-.74.496a2 2 0 0 0 .09.185c.033.063.079.152.121.244l1.96 4.004l-2.158 4.653l3.643 1.116l1.002-5.974l.01-.039c.027-.11.035-.179.036-.22v-.013q0-.014-.002-.023a1.56 1.56 0 0 0-.445-.91zm2.582 12.687l-4.85-1.484H4v2.957h5.37zM4.28 14.189l1.972-4.25l-1.555-3.177l-.007-.015a4 4 0 0 0-.084-.168l-.009-.018a4 4 0 0 1-.147-.309a2.454 2.454 0 0 1 4-2.68l3.518 3.518l.032.033l.032-.033l3.517-3.518a2.454 2.454 0 0 1 4 2.68c-.05.122-.106.232-.146.309l-.01.018c-.039.077-.062.122-.083.168l-.007.015l-1.555 3.176l1.972 4.25l2.28.001v6.957h-9.063L12 15.562l-.937 5.584H2v-6.957zM13 9.45c0 .04.009.11.037.22l.01.039l1.001 5.974l3.644-1.116l-2.159-4.653l1.96-4.004c.042-.093.088-.18.12-.244l.012-.022a2 2 0 0 0 .08-.163a.454.454 0 0 0-.741-.496l-3.517 3.518a1.56 1.56 0 0 0-.447.934m6.231 6.75l-4.849 1.485l.247 1.473H20v-2.957z"/>`), g.Group(children),
	)
}

func GesturePress(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 2.5a5 5 0 0 1 3.985 8.02h-1.23V7.566a2.756 2.756 0 1 0-5.511 0v4.107A5 5 0 0 1 11 2.5m5.911 8.75a7 7 0 1 0-8.667 2.686v1.168l-2.145-.477a1.88 1.88 0 0 0-1.971.792l-.941 1.412l4.298 5.592a2.76 2.76 0 0 0 2.185 1.076h7.086c1.187 0 2.24-.76 2.615-1.885l1.783-5.35a2.76 2.76 0 0 0-1.226-3.253zM11 6.81c.417 0 .755.338.755.756v4.954h3.159a.76.76 0 0 1 .381.103l3.625 2.115c.309.18.45.553.336.893l-1.782 5.35a.76.76 0 0 1-.718.518H9.67a.76.76 0 0 1-.6-.295l-3.425-4.456l.1-.15l4.499 1V7.565c0-.418.338-.756.756-.756"/>`), g.Group(children),
	)
}

func GestureRanslation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.865 2.186a2.57 2.57 0 0 1 5.092.068q.284-.067.588-.068a2.57 2.57 0 0 1 2.57 2.57v.83q.274-.063.567-.063a2.567 2.567 0 0 1 2.567 2.567V15a8 8 0 0 1-8 8h-1.35a8 8 0 0 1-7.738-5.968l-1.269-4.829A2.296 2.296 0 0 1 6.698 9.96V4.755a2.57 2.57 0 0 1 2.57-2.569zm-.028 2.57a.57.57 0 1 0-1.139 0v8.749c0 1.071-1.432 1.373-1.89.462l-1.44-2.495a.296.296 0 0 0-.541.223l1.268 4.83A6 6 0 0 0 11.898 21h1.35a6 6 0 0 0 6-6V8.09a.567.567 0 1 0-1.134 0V12h-2V4.755a.57.57 0 1 0-1.138 0V12h-2V2.57a.57.57 0 0 0-1.139 0V12h-2z"/>`), g.Group(children),
	)
}

func GestureRanslationOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.98 1.545l.823.568a12.1 12.1 0 0 1 3.084 3.084l.568.824l-1.647 1.135l-.568-.824a10.1 10.1 0 0 0-2.572-2.572l-.823-.568zm-6.524 5.332a.388.388 0 0 0-.548.548l5.237 5.237l-1.415 1.415l-6.603-6.604a.388.388 0 0 0-.548.548l6.603 6.604l-1.414 1.414l-5.237-5.237a.388.388 0 0 0-.548.549l5.237 5.236l-1.415 1.415l-3.151-3.152a.386.386 0 0 0-.546.545l4.32 4.32a5.19 5.19 0 0 0 7.34 0l.844-.844a5.19 5.19 0 0 0 .812-6.286l-2.227-3.813a.145.145 0 0 0-.266.111l.669 2.496a1.01 1.01 0 0 1-.277.986a1 1 0 0 1-1.397-.017zM5.99 13.185l-.42-.42a2.388 2.388 0 0 1 .326-3.65A2.388 2.388 0 0 1 9.22 5.79q.12-.172.274-.327a2.39 2.39 0 0 1 3.376 0l3.078 3.078a2.145 2.145 0 0 1 3.976-.777l2.227 3.812a7.19 7.19 0 0 1-1.125 8.71l-.844.844a7.19 7.19 0 0 1-10.168 0l-4.32-4.32a2.386 2.386 0 0 1 .295-3.625m-2.796 3.659l.567.823a10.1 10.1 0 0 0 2.573 2.573l.823.567l-1.135 1.647l-.823-.567a12.1 12.1 0 0 1-3.085-3.085l-.567-.823z"/>`), g.Group(children),
	)
}

func GestureRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.5 7.57a3 3 0 1 1 0 6H16v1.46a3 3 0 0 1-.409 1.511l-2.409 4.13a3 3 0 0 1-3.54 1.334l-6.09-2.03A3 3 0 0 1 1.5 17.129V9.055a3 3 0 0 1 1.17-2.378l6.288-4.836l1.527 1.017a2 2 0 0 1 .843 2.098l-.581 2.614zm1 3a1 1 0 0 0-1-1H8.253l1.122-5.048l-.333-.222L3.89 8.263a1 1 0 0 0-.39.792v8.074a1 1 0 0 0 .684.949l6.09 2.03a1 1 0 0 0 1.18-.445l2.41-4.13A1 1 0 0 0 14 15.03v-3.46h5.5a1 1 0 0 0 1-1"/>`), g.Group(children),
	)
}

func GestureRightSlip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.4 5.2c3.898-2.93 9.3-2.93 13.198 0h-2.1v2h4.986l.003-5.699l-2-.001l-.002 1.872a13.01 13.01 0 0 0-14.972 0V1.5h-2v5.7H7.5v-2zm3.844 2.366a2.756 2.756 0 1 1 5.512 0v2.954h1.158c.488 0 .968.13 1.39.375l3.624 2.116a2.76 2.76 0 0 1 1.226 3.252l-1.782 5.35a2.76 2.76 0 0 1-2.616 1.886h-7.085a2.76 2.76 0 0 1-2.186-1.076L4.187 16.83l.94-1.412a1.88 1.88 0 0 1 1.972-.792l2.145.477zM12 6.81a.756.756 0 0 0-.756.756v10.031l-4.498-1l-.1.151l3.425 4.456a.76.76 0 0 0 .6.295h7.085a.76.76 0 0 0 .718-.517l1.783-5.35a.76.76 0 0 0-.337-.894l-3.625-2.115a.76.76 0 0 0-.381-.103h-3.158V7.566A.756.756 0 0 0 12 6.81"/>`), g.Group(children),
	)
}

func GestureSlideUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m.499 1.513l5.699.002v4.986h-2V4.4C.958 8.711 1.3 14.857 5.222 18.78l.707.707L4.515 20.9l-.707-.707C-.733 15.653-1.212 8.588 2.37 3.514L.498 3.513zm9.508 7.133a.756.756 0 0 0-1.31.756l5.017 8.688l-4.396 1.383l-.012.18l5.195 2.147c.217.09.463.073.667-.044l6.136-3.543a.76.76 0 0 0 .363-.807l-1.132-5.525a.76.76 0 0 0-.737-.605l-4.197-.02a.76.76 0 0 0-.382.101l-2.735 1.58zM7.974 6.638a2.756 2.756 0 0 1 3.765 1.008l1.477 2.558l1.003-.579a2.76 2.76 0 0 1 1.391-.369l4.197.02a2.76 2.76 0 0 1 2.688 2.204l1.131 5.524a2.76 2.76 0 0 1-1.322 2.941l-6.136 3.543a2.76 2.76 0 0 1-2.431.16l-6.518-2.693l.109-1.694c.05-.774.57-1.438 1.31-1.671l2.097-.66l-3.77-6.528a2.756 2.756 0 0 1 1.01-3.764"/>`), g.Group(children),
	)
}

func GestureUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.25 4.5a3 3 0 0 1 6 0V8h1.46a3 3 0 0 1 1.511.409l4.13 2.409a3 3 0 0 1 1.334 3.54l-2.03 6.09a3 3 0 0 1-2.846 2.052H8.735a3 3 0 0 1-2.378-1.17l-4.836-6.288l1.017-1.527a2 2 0 0 1 2.098-.843l2.614.581zm3-1a1 1 0 0 0-1 1v11.247l-5.048-1.122l-.222.333l3.962 5.152a1 1 0 0 0 .793.39h8.074a1 1 0 0 0 .948-.684l2.03-6.09a1 1 0 0 0-.444-1.18l-4.13-2.41A1 1 0 0 0 14.71 10h-3.46V4.5a1 1 0 0 0-1-1"/>`), g.Group(children),
	)
}

func GestureUpAndDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m.499.513l5.699.002v4.986h-2V3.4a11.01 11.01 0 0 0 0 13.2v-2.1h2v4.985h-5.7v-2h1.87A13.01 13.01 0 0 1 2.37 2.514L.498 2.513zm9.508 7.133a.756.756 0 0 0-1.31.756l5.017 8.688l-4.396 1.383l-.012.18l5.195 2.147c.217.09.463.073.667-.044l6.136-3.543a.76.76 0 0 0 .363-.807l-1.132-5.525a.76.76 0 0 0-.737-.605l-4.197-.02a.76.76 0 0 0-.382.101l-2.735 1.58zM7.974 5.638a2.756 2.756 0 0 1 3.765 1.008l1.477 2.558l1.003-.579a2.76 2.76 0 0 1 1.391-.37l4.197.02a2.76 2.76 0 0 1 2.688 2.204l1.131 5.525a2.76 2.76 0 0 1-1.322 2.94l-6.136 3.544a2.76 2.76 0 0 1-2.431.16l-6.518-2.693l.109-1.694c.05-.774.57-1.438 1.31-1.671l2.097-.66l-3.77-6.528a2.756 2.756 0 0 1 1.01-3.764"/>`), g.Group(children),
	)
}

func GestureUpOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2a2 2 0 1 1 4 0a2 2 0 0 1-4 0m7 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0m7 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0M8.244 8.066a2.756 2.756 0 0 1 5.512 0v2.954h1.158c.488 0 .968.13 1.39.375l3.624 2.116a2.76 2.76 0 0 1 1.226 3.252l-1.782 5.35A2.76 2.76 0 0 1 16.756 24H9.671a2.76 2.76 0 0 1-2.186-1.076L3.187 17.33l.94-1.412a1.88 1.88 0 0 1 1.972-.792l2.145.477zM11 7.31a.756.756 0 0 0-.756.756v10.031l-4.498-1l-.1.151l3.425 4.456a.76.76 0 0 0 .6.295h7.085a.76.76 0 0 0 .718-.517l1.782-5.35a.76.76 0 0 0-.336-.894l-3.625-2.115a.76.76 0 0 0-.381-.103h-3.158V8.066A.756.756 0 0 0 11 7.31M2 8a2 2 0 1 1 4 0a2 2 0 0 1-4 0m14 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func GestureUpTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.598 5.2C14.288 1.961 8.142 2.302 4.22 6.224l-.707.707l-1.414-1.414l.707-.707C7.346.27 14.41-.21 19.485 3.372l.001-1.872l2 .001l-.003 5.7h-4.985v-2zM9.244 7.566a2.756 2.756 0 1 1 5.511 0v2.954h1.159c.488 0 .967.13 1.39.375l3.624 2.116a2.76 2.76 0 0 1 1.226 3.252l-1.783 5.35a2.76 2.76 0 0 1-2.615 1.886H10.67a2.76 2.76 0 0 1-2.185-1.076L4.186 16.83l.941-1.412a1.88 1.88 0 0 1 1.971-.792l2.146.477zm2.755-.756a.756.756 0 0 0-.755.756v10.031l-4.498-1l-.101.151l3.425 4.456a.76.76 0 0 0 .6.295h7.086a.76.76 0 0 0 .718-.517l1.782-5.35a.76.76 0 0 0-.336-.894l-3.625-2.115a.76.76 0 0 0-.382-.103h-3.158V7.566A.756.756 0 0 0 12 6.81"/>`), g.Group(children),
	)
}

func GestureWipeDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m5.929 2.515l-.707.707C1.3 7.144.959 13.291 4.198 17.6v-2.1h2v4.986h-5.7v-2h1.87C-1.211 13.413-.732 6.35 3.809 1.808l.707-.707zm4.078 4.131a.756.756 0 0 0-1.31.756l5.017 8.688l-4.396 1.383l-.012.18l5.195 2.147c.217.09.463.073.667-.044l6.136-3.543a.76.76 0 0 0 .363-.807L20.535 9.88a.76.76 0 0 0-.737-.605l-4.197-.02a.76.76 0 0 0-.382.101l-2.735 1.58zM7.974 4.638a2.756 2.756 0 0 1 3.765 1.008l1.477 2.559l1.003-.58a2.76 2.76 0 0 1 1.391-.369l4.197.02a2.76 2.76 0 0 1 2.688 2.204l1.131 5.524a2.76 2.76 0 0 1-1.322 2.941l-6.136 3.543a2.76 2.76 0 0 1-2.431.16l-6.518-2.693l.109-1.694c.05-.774.57-1.438 1.31-1.671l2.097-.66l-3.77-6.528a2.756 2.756 0 0 1 1.01-3.764"/>`), g.Group(children),
	)
}

func Gift(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23 22V6h-4.535A4 4 0 0 0 12 1.354A4 4 0 0 0 5.535 6H1v16zM13 4a2 2 0 1 1 4 0a2 2 0 0 1-4 0m1.414 4H21v6H3V8h6.586l-3 3L8 12.414l4-4l4 4L17.414 11zM3 16h18v4H3zm8-12a2 2 0 1 1-4 0a2 2 0 0 1 4 0"/>`), g.Group(children),
	)
}

func Giggle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m16.362-3.119L15.54 9.906L17.4 11.3l-1.2 1.6l-3.2-2.4V9.04l3.381-1.902zM9 8v4H7V8zm-2 5h10v1a5 5 0 0 1-10 0zm2.17 2a3.001 3.001 0 0 0 5.66 0z"/>`), g.Group(children),
	)
}

func GitBranch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.5 19.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3M2 18a3.5 3.5 0 1 0 4.667-3.3A2 2 0 0 1 8.5 13.5h7a4 4 0 0 0 4-4v-.145a3.502 3.502 0 0 0-1-6.855a3.5 3.5 0 0 0-1 6.855V9.5a2 2 0 0 1-2 2h-7c-.729 0-1.412.195-2 .535v-2.68a3.502 3.502 0 0 0-1-6.855a3.5 3.5 0 0 0-1 6.855v5.29A3.5 3.5 0 0 0 2 18M18.5 7.5a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3m-13 0a1.5 1.5 0 1 1 0-3a1.5 1.5 0 0 1 0 3"/>`), g.Group(children),
	)
}

func GitCommit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-4.9 2a5.002 5.002 0 0 1 9.8 0H23v2h-6.1a5.002 5.002 0 0 1-9.8 0H1v-2z"/>`), g.Group(children),
	)
}

func GitMerge(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.5 4.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M2 6a3.5 3.5 0 1 1 4.667 3.3A2 2 0 0 0 8.5 10.5h7a4 4 0 0 1 4 4v.145a3.502 3.502 0 0 1-1 6.855a3.5 3.5 0 0 1-1-6.855V14.5a2 2 0 0 0-2-2h-7a4 4 0 0 1-2-.535v2.68a3.502 3.502 0 0 1-1 6.855a3.5 3.5 0 0 1-1-6.855v-5.29A3.5 3.5 0 0 1 2 6m16.5 10.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m-13 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/>`), g.Group(children),
	)
}

func GitPullRequest(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.414 3l-2 2H19.5v9.645a3.501 3.501 0 0 1-1 6.855a3.5 3.5 0 0 1-1-6.855V7h-3.086l2 2L15 10.414L10.586 6L15 1.586zM5.5 4.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M2 6a3.5 3.5 0 1 1 4.5 3.355v5.29a3.501 3.501 0 0 1-1 6.855a3.5 3.5 0 0 1-1-6.855v-5.29A3.5 3.5 0 0 1 2 6m3.5 10.5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m13 0a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/>`), g.Group(children),
	)
}

func GitRepository(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 5.5A4.5 4.5 0 0 1 7.5 1H21v22H7.5A4.5 4.5 0 0 1 3 18.5zm2 9.258A4.5 4.5 0 0 1 7.5 14H19V3H7.5A2.5 2.5 0 0 0 5 5.5zM8 16h-.5a2.5 2.5 0 0 0 0 5H19v-5h-5v4.618l-3-1.5l-3 1.5zm4 0h-2v1.382l1-.5l1 .5zM8 5h2.004v2.004H8zm0 3h2.004v2.004H8z"/>`), g.Group(children),
	)
}

func GitRepositoryCommits(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 5.5A4.5 4.5 0 0 1 7.5 1H21v22h-5v-2h3v-5h-2.5v-2H19V3H7.5A2.5 2.5 0 0 0 5 5.5v9.258A4.5 4.5 0 0 1 7.5 14h2v2h-2a2.5 2.5 0 0 0 0 5H10v2H7.5A4.5 4.5 0 0 1 3 18.5zM8 5h2.004v2.004H8zm0 3h2.004v2.004H8zm5 6.615l3.914 3.75l-1.384 1.444L14 18.343V23h-2v-4.657l-1.53 1.466l-1.384-1.445z"/>`), g.Group(children),
	)
}

func GitRepositoryPrivate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.5 7.5a6.5 6.5 0 0 1 13 0V9H21v14H3V9h2.5zm2 1.5h9V7.5a4.5 4.5 0 1 0-9 0zM5 11v10h14V11zm2 1.504h2.004V15H7zM7 16.5h2.004v2.504H7z"/>`), g.Group(children),
	)
}

func Gps(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3.001a8.98 8.98 0 0 0-8.981 8.98c0 2.35.9 4.488 2.378 6.09l.678.734l-1.47 1.356l-.678-.735a10.95 10.95 0 0 1-2.908-7.444C1.019 5.917 5.935 1 11.999 1C18.064 1 22.98 5.916 22.98 11.98c0 2.873-1.103 5.49-2.908 7.445l-.678.735l-1.47-1.356l.678-.735a8.95 8.95 0 0 0 2.378-6.088A8.98 8.98 0 0 0 12 3m0 4.5a4.5 4.5 0 0 0-3.308 7.55l.678.735l-1.47 1.356l-.678-.735a6.5 6.5 0 1 1 9.557 0l-.679.735l-1.47-1.356l.678-.735A4.5 4.5 0 0 0 12 7.5M10 12a2 2 0 1 1 4 0a2 2 0 0 1-4 0m1.947 3.966L16.936 23H6.959zm0 3.457l-1.117 1.575h2.235z"/>`), g.Group(children),
	)
}

func Grape(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.588 4c1.202 0 2.177.974 2.177 2.176q-.002.33-.133.68a4.2 4.2 0 0 0-1.377 1.384q-.325.111-.667.113c-.572 0-1.09-.22-1.48-.58a2.17 2.17 0 0 1-.696-1.597A2.174 2.174 0 0 1 13.589 4m3.97.874a4.178 4.178 0 0 0-6.948-1.625a4.2 4.2 0 0 0-1.257-.19a4.176 4.176 0 0 0-3.891 5.695a4.16 4.16 0 0 0-2.403 3.775c0 .737.191 1.43.526 2.031A4.14 4.14 0 0 0 2 17.823A4.176 4.176 0 0 0 6.176 22c1.323 0 2.485-.65 3.245-1.596a4.16 4.16 0 0 0 2.05.537c1.667 0 3.125-.97 3.796-2.394a4.176 4.176 0 0 0 5.674-3.9c0-.43-.06-.856-.184-1.263a4.178 4.178 0 0 0-1.872-7.013c.89-.838 1.836-1.508 2.968-1.935l.936-.353l-.706-1.872l-.936.353c-1.419.536-2.576 1.363-3.59 2.31m1.375 9.565q.008.102.008.208a2.176 2.176 0 1 1-4.35-.122c.027-.4.159-.764.371-1.07a4.18 4.18 0 0 0 3.97.985m-6.218-.8a2.17 2.17 0 0 1-1.663-.57a2.17 2.17 0 0 1-.699-1.598a2.14 2.14 0 0 1 .957-1.793a4.16 4.16 0 0 0 2.337.675v.059c0 .446.07.877.2 1.281a4.1 4.1 0 0 0-1.132 1.946M9.937 8.205c-.408.32-.754.714-1.017 1.164a2.17 2.17 0 0 1-1.18-.672a2.176 2.176 0 0 1 1.82-3.629a4.18 4.18 0 0 0 .377 3.137m-1.58 3.087q-.004.09-.004.179c0 .738.192 1.432.528 2.034a4.1 4.1 0 0 0-.984 1.099q-.313.1-.662.102c-.572 0-1.09-.22-1.48-.581a2.17 2.17 0 0 1-.696-1.596c0-1.003.68-1.846 1.61-2.094a4.2 4.2 0 0 0 1.688.857m1.901 3.684a4.16 4.16 0 0 0 2.45.667c.157.641.462 1.224.875 1.71a2.18 2.18 0 0 1-2.112 1.588c-.63 0-1.198-.267-1.596-.697a2.17 2.17 0 0 1-.58-1.48c0-.331.073-.644.204-.923l.002-.005l.004-.008a2.13 2.13 0 0 1 .753-.852m-2.963 1.73v.059c0 .832.243 1.608.663 2.26c-.409.585-1.075.975-1.782.975A2.176 2.176 0 0 1 4 17.823c0-.746.377-1.402.96-1.791a4.16 4.16 0 0 0 2.335.674m10.529-8.47a2.177 2.177 0 1 1 0 4.352a2.177 2.177 0 0 1 0-4.353"/>`), g.Group(children),
	)
}

func GreaterThan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.636 4.372L16.985 12l-13.35 7.628l.993 1.736L21.016 12L4.628 2.636z"/>`), g.Group(children),
	)
}

func GreaterThanOrEqual(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m4.435 1.704l17.3 6.796l-17.3 6.796l-.731-1.861L16.265 8.5L3.704 3.565zM3 19h18v2H3z"/>`), g.Group(children),
	)
}

func GreenOnion(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m14.545.914l.028 1l.005.175c.081 2.902.131 4.674-.485 6.214a6 6 0 0 1-.148.338l6.343-6.344l1.415 1.415l-6.344 6.343q.168-.08.338-.148c1.54-.616 3.312-.566 6.214-.485l.176.005l1 .028l-.056 1.999l-1-.028c-3.185-.089-4.496-.1-5.59.338c-.529.212-1.021.533-1.641 1.063a8.3 8.3 0 0 1 2.17-.147c1.698.111 3.312.761 4.942 1.873l.827.563l-1.127 1.652l-.826-.563c-1.421-.968-2.696-1.447-3.948-1.53s-2.578.227-4.113 1.002c-.956.482-1.685 1.346-2.403 2.407c-.212.313-.386.585-.564.863c-.15.234-.303.472-.483.743c-.35.527-.736 1.064-1.179 1.517c-.79.808-1.751 1.487-2.81 1.713c-1.129.242-2.262-.056-3.206-1c-.944-.945-1.242-2.078-1-3.206c.227-1.06.905-2.02 1.713-2.81c.453-.443.99-.828 1.518-1.18c.27-.18.508-.332.742-.482c.278-.178.55-.352.864-.564c1.06-.718 1.924-1.446 2.407-2.402c.774-1.535 1.083-2.862 1-4.114c-.081-1.251-.56-2.527-1.53-3.948l-.562-.826L8.884 1.26l.563.827c1.112 1.63 1.762 3.244 1.873 4.943a8.3 8.3 0 0 1-.147 2.169c.53-.62.852-1.112 1.063-1.64c.438-1.095.427-2.406.338-5.591l-.028-1zm-5.501 12.77c-.62.664-1.334 1.195-2.006 1.65c-.292.198-.636.42-.96.627c-.242.156-.473.304-.66.429c-.5.333-.912.636-1.227.944c-.66.646-1.043 1.274-1.155 1.799c-.098.455-.013.9.459 1.372s.917.557 1.372.46c.525-.113 1.152-.495 1.8-1.156c.307-.315.61-.726.943-1.228c.125-.186.273-.417.43-.66c.207-.322.428-.667.626-.959c.455-.672.987-1.386 1.651-2.005l-.282-.28l.003-.003l-.712-.712l-.003.003z"/>`), g.Group(children),
	)
}

func GridAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h9v9H2zm2 2v5h5V4zm9-2h9v9h-9zm2 2v5h5V4zM2 13h9v9H2zm2 2v5h5v-5zm12.5-1.5h2v3h3v2h-3v3h-2v-3h-3v-2h3z"/>`), g.Group(children),
	)
}

func GridView(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h9v9H2zm2 2v5h5V4zm9-2h9v9h-9zm2 2v5h5V4zM2 13h9v9H2zm2 2v5h5v-5zm9-2h9v9h-9zm2 2v5h5v-5z"/>`), g.Group(children),
	)
}

func Guitar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 1.586L22.414 5L21 6.414l-1-1l-.586.586l.5.5L18.5 7.914l-.5-.5l-.828.829a5 5 0 0 1-2.183 7.346q.01.204.011.411a7 7 0 1 1-6.589-6.988a5 5 0 0 1 7.346-2.183L16.586 6l-.5-.5L17.5 4.086l.5.5l.586-.586l-1-1zM13 8a3 3 0 0 0-2.924 2.325l-.215.934l-.942-.175a5 5 0 1 0 3.997 3.997l-.175-.942l.934-.215A3.002 3.002 0 0 0 13 8m-5 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0"/>`), g.Group(children),
	)
}

func Hamburger(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.706 4.073C7.366 2.7 9.605 2 12 2c2.396 0 4.635.7 6.294 2.073C19.974 5.462 21 7.495 21 10v1H3v-1c0-2.505 1.027-4.538 2.706-5.927M5.074 9h13.852c-.219-1.432-.911-2.563-1.907-3.386C15.784 4.59 14.023 4 12 4s-3.784.591-5.02 1.614C5.986 6.437 5.294 7.568 5.075 9M6 11.798l3 2l3-2l3 2l3-2l4.387 2.925l-1.11 1.664L18 14.202l-3 2l-3-2l-3 2l-3-2l-3.277 2.185l-1.11-1.664zM3 17h18v1a5 5 0 0 1-5 5H8a5 5 0 0 1-5-5zm2.17 2A3 3 0 0 0 8 21h8a3 3 0 0 0 2.83-2z"/>`), g.Group(children),
	)
}

func Happy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m8-4v4H7V8zm8 0v4h-2V8zM7 13h10v1a5 5 0 0 1-10 0zm2.17 2a3.001 3.001 0 0 0 5.66 0z"/>`), g.Group(children),
	)
}

func HardDiskStorage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4h-2.5v7.875h-11V4zm4.5 0v5.875h7V4zM14 6v3h-2V6z"/>`), g.Group(children),
	)
}

func HardDrive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v10h16V4zm16 12H4v4h16zm-14.002.998h2.004v2.004H5.998zm3 0h2.004v2.004H8.998z"/>`), g.Group(children),
	)
}

func Hashtag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.604 2.117L10.006 7.5h4.988l.623-5.604l1.987.22l-.598 5.384H22v2h-5.216l-.556 5H22v2h-5.994l-.623 5.604l-1.987-.22l.598-5.384H9.006l-.623 5.604l-1.987-.22l.598-5.384H2v-2h5.216l.556-5H2v-2h5.994l.623-5.604zM9.784 9.5l-.556 5h4.988l.556-5z"/>`), g.Group(children),
	)
}

func Hd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v14h18V5zm4 3v3h2V8h2v8H9v-3H7v3H5V8zm6 0h4a2 2 0 0 1 2 2v4a2 2 0 0 1-2 2h-4zm2 2v4h2v-4z"/>`), g.Group(children),
	)
}

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.001 3.818a6.228 6.228 0 0 1 8.51 9.087l-5.224 5.225h-.001L12 21.415l-7.28-7.279l-1.23-1.232A6.228 6.228 0 0 1 12 3.818m3.285 11.485l3.811-3.812a4.228 4.228 0 1 0-5.98-5.98L12 6.627L10.883 5.51a4.228 4.228 0 1 0-5.98 5.98l1.232 1.232L12 18.587l3.285-3.285"/>`), g.Group(children),
	)
}

func HeartFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.002 4.818a6.228 6.228 0 0 1 8.51 9.087l-5.225 5.225L12 22.415l-7.28-7.279l-1.23-1.232a6.228 6.228 0 0 1 8.511-9.086"/>`), g.Group(children),
	)
}

func Help(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12.995 21.006l-2-.006v-2.006h2.012zm-2-3.556v-2.38c0-1.15.518-2.11 1.137-2.871c.614-.756 1.395-1.397 2.055-1.917a3.544 3.544 0 1 0-5.534-3.968l-.334.943l-1.885-.666l.333-.943A5.546 5.546 0 0 1 17.54 7.497a5.54 5.54 0 0 1-2.116 4.357c-.655.516-1.279 1.04-1.74 1.608c-.458.562-.689 1.087-.689 1.609v2.379z"/>`), g.Group(children),
	)
}

func HelpCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-9.996 6.254H11V16.25h2.004zM11 15.25V14c0-.867.39-1.573.826-2.11c.432-.53.974-.974 1.41-1.318a2 2 0 1 0-3.123-2.24l-.333.944l-1.885-.666l.333-.943a4.001 4.001 0 1 1 6.246 4.476c-.431.34-.817.666-1.096 1.009c-.274.338-.378.61-.378.848v1.25z"/>`), g.Group(children),
	)
}

func HelpCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11m-.174-11.11c.432-.53.974-.974 1.41-1.318a2 2 0 1 0-3.123-2.24l-.332.944l-1.886-.666l.333-.943a4.001 4.001 0 1 1 6.246 4.476c-.431.34-.817.666-1.096 1.009c-.274.338-.378.61-.378.848v1.25h-2V14c0-.867.39-1.573.826-2.11M11 18.254V16.25h2.004v2.004z"/>`), g.Group(children),
	)
}

func HelpRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm8 3.25a2 2 0 0 0-1.886 1.333l-.334.943l-1.885-.666l.333-.943a4.001 4.001 0 1 1 6.246 4.476c-.431.34-.817.666-1.096 1.009c-.274.338-.378.61-.378.848v1.25h-2v-1.25c0-.867.39-1.573.826-2.11c.432-.53.974-.974 1.41-1.318A2 2 0 0 0 12 7.25m-1 9.25h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func Highlight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h3.5v2H4v1.5H2zm5.5 0h3v2h-3zm5 0h3v2h-3zm5 0H21v3.5h-2V4h-1.5zM4 7.5v3H2v-3zm17 0v3h-2v-3zm-17 5v3H2v-3zM20 14v2h-2.586l4 4L20 21.414l-4-4V20h-2v-6zM4 17.5V19h1.5v2H2v-3.5zM7.5 19h3v2h-3z"/>`), g.Group(children),
	)
}

func HighlightOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m11-5a5 5 0 0 0-4.354 2.54l-.493.871l-1.74-.985l.492-.87A7 7 0 0 1 12 5h1v2zm-4.996 4.133v2.076h-2v-2.076z"/>`), g.Group(children),
	)
}

func History(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22.5 12c0-5.799-4.701-10.5-10.5-10.5c-1.798 0-3.493.453-4.975 1.251A10.55 10.55 0 0 0 3.5 5.834V2.5h-2v7h7v-2H4.787a8.55 8.55 0 0 1 3.187-2.988A8.46 8.46 0 0 1 12 3.5a8.5 8.5 0 1 1-8.454 9.396l-.104-.995l-1.989.209l.104.994C2.11 18.384 6.573 22.5 12 22.5c5.799 0 10.5-4.701 10.5-10.5M11 6v6.414l3.5 3.5l1.414-1.414L13 11.586V6z"/>`), g.Group(children),
	)
}

func HistorySetting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3.5a8.46 8.46 0 0 0-4.026 1.012A8.55 8.55 0 0 0 4.787 7.5H8.5v2h-7v-7h2v3.334a10.55 10.55 0 0 1 3.525-3.083A10.46 10.46 0 0 1 12 1.5c5.799 0 10.5 4.701 10.5 10.5v1h-2v-1A8.5 8.5 0 0 0 12 3.5M13 6v5.586l1.664 1.664l-1.414 1.414l-2.25-2.25V6zm-9.558 5.901l.104.995A8.5 8.5 0 0 0 12 20.5h1v2h-1c-5.427 0-9.89-4.115-10.443-9.396l-.104-.994zM19.5 14.005v1.14c.533.159 1.013.44 1.406.813l.989-.57l1 1.731l-.99.572a3.5 3.5 0 0 1 0 1.623l.99.571l-1 1.732l-.992-.572a3.5 3.5 0 0 1-1.403.81V23h-2v-1.145a3.5 3.5 0 0 1-1.404-.81l-.992.572l-1-1.732l.99-.572a3.5 3.5 0 0 1 0-1.622l-.99-.571l1-1.732l.99.57a3.5 3.5 0 0 1 1.406-.813v-1.14zM18.5 17a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3"/>`), g.Group(children),
	)
}

func Home(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 1.198l10 8.334V22H2V9.532zM10 20h4v-5h-4zm6 0h4v-9.532l-8-6.666l-8 6.666V20h4v-7h8z"/>`), g.Group(children),
	)
}

func Hospital(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2h16v2h-1v4h3v14H2V8h3V4H4zm3 2v6H4v10h5v-6h6v6h5V10h-3V4zm6 16v-4h-2v4zm0-15v2h2v2h-2v2h-2V9H9V7h2V5z"/>`), g.Group(children),
	)
}

func HospitalOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.999 1.66l10.413 9.257l-1.329 1.495L20 11.449v10.55H4V11.455l-1.094.957l-1.317-1.505L4.338 8.5zm-6 8.038V20H9v-5h6v5h3V9.67l-6-5.33zM13 20v-3h-2v3zm0-13v2h2v2h-2v2h-2v-2H9V9h2V7z"/>`), g.Group(children),
	)
}

func HotspotWave(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 0h1a9 9 0 0 1 9 9v1h-2V9a7 7 0 0 0-7-7h-1zM4 2h6v2H6v18h12V12h2v12H4zm8 2h1a5 5 0 0 1 5 5v1h-2V9a3 3 0 0 0-3-3h-1zm0 4h2.004v2.004H12z"/>`), g.Group(children),
	)
}

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2h16v3a8 8 0 0 1-4.124 7A8 8 0 0 1 20 19v3H4v-3a8 8 0 0 1 4.124-7A8 8 0 0 1 4 5zm8 11a6 6 0 0 0-6 6v1h12v-1a6 6 0 0 0-6-6M6 4v1a6 6 0 1 0 12 0V4z"/>`), g.Group(children),
	)
}

func Houses(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.996 1.658l10.416 9.259l-1.329 1.495L20 11.449v10.55H4V11.455l-1.094.957l-1.317-1.505l4.41-3.86V3H8v2.254zM6 9.704V20h3v-6h6v6h3V9.67l-5.996-5.33L7.66 8.252zM13 20v-4h-2v4z"/>`), g.Group(children),
	)
}

func HousesOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h4v-7h8v7h4V4zm10 16v-5h-4v5zM8 6v3H6V6zm5 0v3h-2V6zm5 0v3h-2V6z"/>`), g.Group(children),
	)
}

func HousesTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 2v2h9.566l4.2 7H21v11H3V11H1.233l4.2-7H7V2zm-4 9v9h4v-5h6v5h4v-9zm14.233-2l-1.8-3H6.566l-1.8 3zM13 20v-3h-2v3z"/>`), g.Group(children),
	)
}

func HtmlFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.941 1H22.06l-1.098 19.208l-8.96 3.36l-8.962-3.36zM4.06 3l.902 15.792l7.04 2.64l7.038-2.64L19.941 3zm1.61 2h12.66l-.115 2.017L16.495 7H7.787l.193 3.377h10.043l-.405 7.084L12 19.568l-5.618-2.107l-.194-3.388h2.024l.044 1.12l.048.853L12 17.432l3.696-1.386l.21-3.67H6.09z"/>`), g.Group(children),
	)
}

func Https(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.5 7.5a6.5 6.5 0 0 1 13 0V9h2v13h-17V9h2zm2 1.5h9V7.5a4.5 4.5 0 1 0-9 0zm-2 2v9h13v-9zm4.5 4.5a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func IceCream(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 2.717c-1.757-1.153-4.17-.961-5.708.576a4.55 4.55 0 0 0-1.316 3.531L2.073 19.291l-.008.017c-.711 1.66.967 3.337 2.626 2.626l.018-.007l12.467-5.904a4.55 4.55 0 0 0 3.53-1.316c1.538-1.537 1.73-3.95.576-5.707c1.154-1.756.962-4.17-.575-5.707S16.756 1.563 15 2.717m1.414 1.496c.94-.436 2.113-.272 2.878.494c.976.975.976 2.61 0 3.586c-.765.766-1.937.93-2.878.494a4.59 4.59 0 0 0 0-4.574m3.372 6.202c.437.94.272 2.112-.494 2.878c-.5.5-1.173.752-1.844.74l-7.482-7.48a2.55 2.55 0 0 1 .74-1.846c.976-.975 2.611-.975 3.587 0c.975.975.975 2.61 0 3.586L13.585 9l.707.707c1.476 1.476 3.76 1.712 5.494.708m-4.497 4.289l-11.384 5.39L9.296 8.712z"/>`), g.Group(children),
	)
}

func Icon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 2H8v5.15A7.5 7.5 0 1 0 16.85 16H22zm-5.016 12A7.5 7.5 0 0 0 10 7.016V4h10v10zM9.5 9a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11"/>`), g.Group(children),
	)
}

func Image(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 18h13.586L9 11.414l-5 5zm16-.414V4H4v9.586l5-5zM15.547 7a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0"/>`), g.Group(children),
	)
}

func ImageAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v10h-2V4H4v9.586l5-5L14.414 14L13 15.414l-4-4l-5 5V20h8v2H2zm13.547 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0M20 14v4h4v2h-4v4h-2v-4h-4v-2h4v-4z"/>`), g.Group(children),
	)
}

func ImageEdit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v8h-2V4H4v9.586l5-5l5.914 5.914l-1.414 1.414l-4.5-4.5l-5 5V20h6v2H2zm13.547 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0m7.24 4.086l4.127 4.127l-7.286 7.287H12.5l-.001-4.128zm-.922 3.75l1.299 1.3l.922-.923l-1.3-1.299zm-.115 2.713l-1.3-1.299l-2.95 2.95v1.3h1.3z"/>`), g.Group(children),
	)
}

func ImageError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v10h-2V4H4v9.586l5-5L14.414 14L13 15.414l-4-4l-5 5V20h8v2H2zm13.547 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0m3.625 6.757L19 17.586l2.828-2.829l1.415 1.415L20.414 19l2.829 2.828l-1.415 1.415L19 20.414l-2.828 2.829l-1.415-1.415L17.586 19l-2.829-2.828z"/>`), g.Group(children),
	)
}

func ImageOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 .586L.586 22L2 23.414L3.414 22H22V3.414L23.414 2zM11.5 13.914L17.586 20H5.414zm8.5 5.672L12.915 12.5l1.06-1.062a3.21 3.21 0 0 0 4.463-4.463L20 5.414zm-4.47-9.614l1.443-1.442q.027.125.028.26a1.21 1.21 0 0 1-1.47 1.182M2.004 2l-.005 15.658l2.005-1.696V4l11.362.002l2.37-2.004z"/>`), g.Group(children),
	)
}

func ImageOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 4a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0M9.928 8.124l4.419 7.015l2.581-3.975L23.966 22H1.188zm0 3.752L4.812 20h15.47l-3.354-5.164l-2.607 4.014z"/>`), g.Group(children),
	)
}

func ImageSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v9h-2V4H4v9.586l5-5l3.914 3.914l-1.414 1.414l-2.5-2.5l-5 5V20h7v2H2zm13.547 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0m5.107 6.5a3.154 3.154 0 1 0 0 6.308a3.154 3.154 0 0 0 0-6.308M12.5 17.654a5.154 5.154 0 1 1 9.437 2.867l1.977 1.98l-1.415 1.413l-1.976-1.978a5.154 5.154 0 0 1-8.023-4.282"/>`), g.Group(children),
	)
}

func IndentLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 4v2h20V4zm6 7v2h14v-2zm-6 7h20v2H2zm-.414-6l3.182 3.182l1.414-1.414L4.414 12l1.768-1.768l-1.414-1.414z"/>`), g.Group(children),
	)
}

func IndentRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 4H2v2h20V4zm6 7H8v2h14v-2zm-7 7h20v2H2zm3.805-5.293L6.512 12l-.707-.707l-1.768-1.768l-.707-.707l-1.414 1.414l.707.708L3.683 12l-1.06 1.06l-.707.708l1.414 1.414l.707-.707z"/>`), g.Group(children),
	)
}

func Indicator(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.353 1v3h3.294V1h2v3h1.89l4.097 3l-4.096 3h-1.89v2.111H22v6h-6.353V23h-2v-4.889h-3.294V23h-2v-4.889h-1.89l-4.097-3l4.096-3h1.89V10H2V4h6.353V1zM4 6v2h12.884l1.365-1l-1.365-1zm9.647 4h-3.294v2.111h3.294zm-6.53 4.111l-1.366 1l1.366 1H20v-2z"/>`), g.Group(children),
	)
}

func InfoCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-12 5.5V10h2v7.5zm2-9h-2.004V6.496H13z"/>`), g.Group(children),
	)
}

func InfoCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11M10.996 8.5V6.496H13V8.5zM13 10v7.5h-2V10z"/>`), g.Group(children),
	)
}

func Ink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 12.896V21H2v-8.104l2.196-4.182H6.2V3h11.6v5.714h2.004zm-6.2-4.182V5H8.2v3.714zM20 13.39l-1.404-2.675H5.404L4 13.39V19h16z"/>`), g.Group(children),
	)
}

func Install(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 2v7.11l2.508-2.48l1.406 1.422L12 12.91L7.086 8.052L8.492 6.63L11 9.11V2zM2 2h7v2H4v10h16V4h-5V2h7v20H2zm18 14H4v4h16zm-14.002.998h2.004v2.004H5.998zm3 0h2.004v2.004H8.998z"/>`), g.Group(children),
	)
}

func InstallDesktop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 2v4.657l1.53-1.466l1.384 1.445L19 10.385l-3.914-3.75l1.384-1.444L18 6.657V2zM1 3h13v2H3v11h18v-4h2v6H1zm4 17h14v2H5z"/>`), g.Group(children),
	)
}

func InstallMobile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 1h11v2H6v18h12V11h2v12H4zm16 0v4.657l1.53-1.466l1.384 1.445L19 9.385l-3.914-3.75l1.384-1.444L18 5.657V1zm-9 16h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func Institution(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .856l10 5.556V9H2V6.412zM5.06 7h13.88L12 3.144zM7 11v8H5v-8zm6 0v8h-2v-8zm6 0v8h-2v-8zM2 21h20v2H2z"/>`), g.Group(children),
	)
}

func InstitutionChecked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11 .856l10 5.556V9H1V6.412zM4.06 7h13.88L11 3.144zM6 11v8H4v-8zm6 0v8h-2v-8zm6 0v6h-2v-6zm4.596 6.94l-5.657 5.656l-3.535-3.535l1.414-1.414l2.121 2.12l4.243-4.242zM1 21h11v2H1z"/>`), g.Group(children),
	)
}

func Internet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.055 11a9.01 9.01 0 0 1 6.277-7.598A16.9 16.9 0 0 0 7.029 11zm7.937-9.954C5.39 1.554 1 6.265 1 12s4.39 10.446 9.992 10.955l.008.01l.425.02A13 13 0 0 0 12 23a11 11 0 0 0 .575-.015l.425-.02l.008-.01C18.61 22.444 23 17.735 23 12S18.61 1.554 13.008 1.046L13 1.036l-.426-.021a11 11 0 0 0-1.148 0l-.426.02zM12.002 3a14.9 14.9 0 0 1 2.965 8H9.033a14.9 14.9 0 0 1 2.966-8H12M7.028 13c.16 2.76.98 5.345 2.303 7.598A9.01 9.01 0 0 1 3.055 13zm4.97 8a14.9 14.9 0 0 1-2.966-8h5.934A14.9 14.9 0 0 1 12 21zm2.67-.402A16.9 16.9 0 0 0 16.97 13h3.974a9.01 9.01 0 0 1-6.276 7.598M16.97 11c-.16-2.76-.98-5.345-2.303-7.598A9.01 9.01 0 0 1 20.945 11z"/>`), g.Group(children),
	)
}

func Ipod(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h18v22H3zm2 2v10h2V9h2v4h2V5h2v8h2V7h2v6h2V3zm14 12H5v6h14zm-9 3a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func Joyful(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m6.619-4.862L11 9.04v1.46l-3.2 2.4l-1.2-1.6l1.86-1.394L6.637 8.88zm9.743 1.743L15.54 9.906L17.4 11.3l-1.2 1.6l-3.2-2.4V9.04l3.381-1.902zM7 13h10v1a5 5 0 0 1-10 0zm2.17 2a3.001 3.001 0 0 0 5.66 0z"/>`), g.Group(children),
	)
}

func Jump(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 21V3h7.09v2H5v14h14v-5.09h2V21zm7.586-9l7-7H13V3h8v8h-2V6.414l-7 7z"/>`), g.Group(children),
	)
}

func JumpOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22.414 21L3 1.586L1.586 3L3 4.414V21h16.586L21 22.414zm-4.828-2H5V6.414L10.586 12l-.502.502l1.417 1.417l.502-.502zm3.42-1.581L21 14.499v-1.01h-2v1.924zM21 2.999h-8v2h4.586l-4.5 4.5l1.414 1.415l4.5-4.5V11h2zm-10.49.003H9.5l-2.92-.006l2.007 2.006h1.923z"/>`), g.Group(children),
	)
}

func Keyboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v14h18V5zm1.996 2.5H7v2.004H4.996zm4 0H11v2.004H8.996zm4 0H15v2.004h-2.004zm4 0H19v2.004h-2.004zm-12 3H7v2.004H4.996zm4 0H11v2.004H8.996zm4 0H15v2.004h-2.004zm4 0H19v2.004h-2.004zM5 15h14v2H5z"/>`), g.Group(children),
	)
}

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v15H1zm2 2v11h18V5zM0 19h24v2H0z"/>`), g.Group(children),
	)
}

func Layers(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 1.917l10.1 4.208L12 10.333L1.9 6.125zM7.1 6.125L12 8.166l4.9-2.041L12 4.083zM2 9.98l10 4.308L22 9.98v2.181l-9.604 4.134l-.396.17l-.395-.17L2 12.162zm0 6l10 4.308l10-4.308v2.181l-9.603 4.134l-.397.17l-.395-.17L2 18.162z"/>`), g.Group(children),
	)
}

func Layout(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v5h16V4zm16 7h-7v9h7zm-9 9v-9H4v9z"/>`), g.Group(children),
	)
}

func Leaderboard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 3h8v8h6v10H2V9h6zm2 16h4V5h-4zm6 0h4v-6h-4zm-8 0v-8H4v8z"/>`), g.Group(children),
	)
}

func Lemon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.608 12.344c-.183 1.679.312 3.24 1.223 4.6l.25.373l-.567 2.179l2.165-.579l.376.251c1.36.907 2.923 1.4 4.603 1.218c1.68-.183 3.596-1.054 5.635-3.093c2.04-2.039 2.91-3.955 3.093-5.635s-.31-3.244-1.218-4.603l-.25-.376l.578-2.165l-2.18.567l-.372-.25c-1.36-.911-2.921-1.406-4.6-1.223s-3.594 1.056-5.637 3.1c-2.043 2.042-2.916 3.958-3.099 5.636m-1.988-.216c.243-2.231 1.387-4.548 3.673-6.835s4.604-3.43 6.834-3.673c2.062-.225 3.948.331 5.544 1.303l1.547-.403l.047-.007a1.95 1.95 0 0 1 2.23 2.256l-.009.044l-.406 1.521c.965 1.595 1.518 3.48 1.294 5.54c-.242 2.231-1.383 4.55-3.667 6.833s-4.602 3.425-6.833 3.667c-2.06.224-3.945-.329-5.54-1.294l-1.52.406l-.045.008a1.95 1.95 0 0 1-2.256-2.23l.007-.046l.403-1.547c-.972-1.596-1.528-3.482-1.303-5.543m6.257-2.833l1.417-1.417l1.417 1.417l-1.417 1.417zm-.707 3.536l1.417-1.417l1.417 1.417l-1.417 1.417zm4.242-4.243l1.417-1.417l1.417 1.417l-1.417 1.417z"/>`), g.Group(children),
	)
}

func LemonSlice(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.995.59L1.59 16.996l.697.707c1.04 1.056 2.523 2.117 4.125 2.918C8.007 21.416 9.81 22 11.488 22a11.48 11.48 0 0 0 8.14-3.372A11.48 11.48 0 0 0 23 10.488c0-1.678-.584-3.481-1.38-5.075c-.8-1.603-1.862-3.085-2.918-4.126zm2.904 8.8l-8.172-.703l5.243-5.243a15 15 0 0 1 1.861 2.863c.519 1.038.897 2.101 1.068 3.083m-7.754 1.34l7.812.672a9.46 9.46 0 0 1-2.073 5.068zm4.325 7.154a9.46 9.46 0 0 1-5.068 2.072l-.671-7.811zm-7.782-6.158l.702 8.173c-.982-.171-2.045-.55-3.083-1.068a15 15 0 0 1-2.862-1.861z"/>`), g.Group(children),
	)
}

func LessThan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.364 4.372L7.015 12l13.35 7.628l-.993 1.736L2.984 12l16.388-9.364z"/>`), g.Group(children),
	)
}

func LessThanOrEqual(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.297 3.565L7.735 8.5l12.562 4.935l-.731 1.861L2.266 8.5l17.3-6.796zM3 19h18v2H3z"/>`), g.Group(children),
	)
}

func LettersA(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 20V6a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v14h2v-7h6v7zm-2-9H9V6h6z"/>`), g.Group(children),
	)
}

func LettersB(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 4h8a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H7zm2 2v5h6V6zm6 7H9v5h6z"/>`), g.Group(children),
	)
}

func LettersC(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 4H9a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h8v-2H9V6h8z"/>`), g.Group(children),
	)
}

func LettersD(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 4h8a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H7zm2 2v12h6V6z"/>`), g.Group(children),
	)
}

func LettersE(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 6a2 2 0 0 1 2-2h8v2H9v5h8v2H9v5h8v2H9a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func LettersF(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 6a2 2 0 0 1 2-2h7v2h-7v5h7v2h-7v7H8z"/>`), g.Group(children),
	)
}

func LettersG(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 4H8.5A1.5 1.5 0 0 0 7 5.5v13A1.5 1.5 0 0 0 8.5 20H17v-9h-5.2v2H15v5H9V6h7z"/>`), g.Group(children),
	)
}

func LettersH(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 4v7h6V4h2v16h-2v-7H9v7H7V4z"/>`), g.Group(children),
	)
}

func LettersI(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.5 4h7v2H13v12h2.5v2h-7v-2H11V6H8.5z"/>`), g.Group(children),
	)
}

func LettersJ(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.667 4H15v14.5a1.5 1.5 0 0 1-1.5 1.5H8v-2h5V6H9.667z"/>`), g.Group(children),
	)
}

func LettersK(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 20v-6.057l5 3.572V20h2v-2.743a1.5 1.5 0 0 0-.628-1.22L10.72 12l5.65-4.037a1.5 1.5 0 0 0 .63-1.22V4h-2v2.485l-5 3.572V4H8v16z"/>`), g.Group(children),
	)
}

func LettersL(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 4v14a2 2 0 0 0 2 2h8v-2H9V4z"/>`), g.Group(children),
	)
}

func LettersM(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 20V6h-5v14h-2V6H6v14H4V4h14a2 2 0 0 1 2 2v14z"/>`), g.Group(children),
	)
}

func LettersN(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.5 20V6h7v14h2V6a2 2 0 0 0-2-2h-9v16z"/>`), g.Group(children),
	)
}

func LettersO(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 18V6h6v12zM7 6v12a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2V6a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2"/>`), g.Group(children),
	)
}

func LettersP(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 4h8a2 2 0 0 1 2 2v5.5a2 2 0 0 1-2 2H9V20H7zm2 7.5h6V6H9z"/>`), g.Group(children),
	)
}

func LettersQ(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 6v12h6V6zm8 12h1.5v2h-10A1.5 1.5 0 0 1 7 18.5v-13A1.5 1.5 0 0 1 8.5 4h7A1.5 1.5 0 0 1 17 5.5z"/>`), g.Group(children),
	)
}

func LettersR(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 4h8.5A1.5 1.5 0 0 1 17 5.5v6a1.5 1.5 0 0 1-1.5 1.5h-2.77l4.239 3.587a1.5 1.5 0 0 1 .531 1.145V20h-2v-2.036L9.634 13H9v7H7zm2 7h6V6H9z"/>`), g.Group(children),
	)
}

func LettersS(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 6a2 2 0 0 1 2-2H16v2H9.5v5h5a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2H8v-2h6.5v-5h-5a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func LettersT(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.5 4h11v2H13v14h-2V6H6.5z"/>`), g.Group(children),
	)
}

func LettersU(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 4v14H9V4H7v14a2 2 0 0 0 2 2h8V4z"/>`), g.Group(children),
	)
}

func LettersV(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 4v10.273l-3 3.652l-3-3.652V4H7v10.273a2 2 0 0 0 .455 1.27l3.772 4.592h1.546l3.772-4.592a2 2 0 0 0 .455-1.27V4z"/>`), g.Group(children),
	)
}

func LettersW(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 4v14h5V4h2v14h5V4h2v14a2 2 0 0 1-2 2H6a2 2 0 0 1-2-2V4z"/>`), g.Group(children),
	)
}

func LettersX(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.495 8.449A2 2 0 0 0 17 7.119V4h-2v3.12l-3 3.375L9 7.12V4H7v3.12a2 2 0 0 0 .505 1.329L10.662 12l-3.157 3.552A2 2 0 0 0 7 16.88V20h2v-3.12l3-3.375l3 3.375V20h2v-3.12a2 2 0 0 0-.505-1.328L13.338 12z"/>`), g.Group(children),
	)
}

func LettersY(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.713 8.22A1.5 1.5 0 0 0 17 7.337V4h-2v3.175L12 11.3L9 7.175V4H7v3.337c0 .317.1.626.287.883L11 13.325V20h2v-6.675z"/>`), g.Group(children),
	)
}

func LettersZ(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.571 5.1a1.1 1.1 0 0 0-1.1-1.1H7v2h7.571v1.135l-7.571 9V18.9A1.1 1.1 0 0 0 8.1 20H17v-2H9v-1.135l7.571-9z"/>`), g.Group(children),
	)
}

func Lightbulb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2c-3.418-.002-6.2 2.782-6.2 6.2c0 .917.169 1.61.632 2.723c.16.385.405.91.696 1.533q.213.456.452.975A48 48 0 0 1 8.69 16h6.629c.341-.853.741-1.749 1.114-2.565l.39-.85c.312-.677.576-1.25.745-1.66c.46-1.114.632-1.81.632-2.725c0-3.417-2.782-6.198-6.2-6.2M3.8 8.2C3.8 3.679 7.477-.002 12 0s8.2 3.678 8.2 8.2c0 1.264-.262 2.227-.783 3.489c-.186.45-.477 1.08-.797 1.777q-.181.391-.368.8c-.467 1.023-.953 2.121-1.31 3.083l-.243.651h-9.41l-.235-.666c-.339-.955-.82-2.046-1.29-3.067l-.423-.91a50 50 0 0 1-.756-1.665C4.058 10.425 3.8 9.462 3.8 8.2M7.5 19h9v2h-9zm2 3h5v2h-5z"/>`), g.Group(children),
	)
}

func LightbulbCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-8.5 7h-5v-2h5zm.735-3h-6.47l-.235-.668c-.435-1.236-.978-2.398-1.488-3.456A5.4 5.4 0 0 1 6.5 9.5C6.5 6.467 8.967 4 12 4s5.5 2.467 5.5 5.5c0 .83-.181 1.631-.542 2.377c-.51 1.057-1.052 2.219-1.488 3.455zm-1.405-2a40 40 0 0 1 1.328-2.993q.341-.706.342-1.507C15.5 7.571 13.93 6 12 6a3.506 3.506 0 0 0-3.5 3.5q0 .799.343 1.507c.429.888.906 1.9 1.327 2.993z"/>`), g.Group(children),
	)
}

func Lighthouse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 1v1h6V1h2v9c0 2.393.482 5.311.978 7.669a71 71 0 0 0 .962 3.982l.016.057l.004.014v.003L19.326 23H13.22l-1-4h-.439l-1 4H4.674l.364-1.275l.001-.003l.004-.014l.016-.057l.062-.224a71 71 0 0 0 .9-3.758C6.517 15.311 7 12.393 7 10V1zm0 6.002V10c0 2.607-.518 5.689-1.022 8.081A73 73 0 0 1 7.3 21h1.92l1-4h3.561l1 4H16.7a74 74 0 0 1-.678-2.919C15.517 15.689 15 12.607 15 10V7.002zM15 5V4H9v1zm-4 3.998h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func LighthouseOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .923l6.3 2.52l-.743 1.857L17 5.077V10h2v2h-1.907c.164 1.905.52 3.938.885 5.669a71 71 0 0 0 .962 3.982l.016.057l.004.014v.003L19.326 23H13.22l-1-4h-.439l-1 4H4.674l.364-1.275l.001-.004l.004-.013l.016-.057l.062-.224a71 71 0 0 0 .9-3.758c.364-1.731.721-3.764.885-5.67H5v-2h2V5.078l-.558.223L5.7 3.443zM9 4.277V10h6V4.277l-3-1.2zM15.086 12H8.913c-.168 2.086-.556 4.28-.935 6.08A73 73 0 0 1 7.3 21h1.92l1-4h3.561l1 4H16.7a74 74 0 0 1-.678-2.92c-.38-1.8-.767-3.994-.935-6.08M11 5.998h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func LighthouseTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 1v1.002h6V1h2v4h1v2h-1v3c0 .94.074 1.963.195 3H19v2h-1.523c.151.928.325 1.834.501 2.669a71 71 0 0 0 .9 3.758c.144.526.298 1.049.447 1.573H13.22l-1-4h-.439l-1 4H4.674c.15-.524.304-1.047.447-1.573a71 71 0 0 0 .9-3.758c.176-.835.35-1.741.501-2.669H5v-2h1.804c.121-1.037.196-2.06.196-3V7H6V5h1V1zm0 3.002V5h6v-.998zM15 7H9v3c0 .955-.07 1.974-.183 3h6.365A28 28 0 0 1 15 10zm.451 8H8.548a61 61 0 0 1-.57 3.081A73 73 0 0 1 7.3 21h1.92l1-4h3.561l1 4H16.7a74 74 0 0 1-.678-2.919a61 61 0 0 1-.57-3.081M11 8.998h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func LightingCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-10.915 7.5H10v-5H6.29l5.625-10H14v5h3.71zM12 15.571l2.29-4.071H12V8.429L9.71 12.5H12z"/>`), g.Group(children),
	)
}

func LineHeight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.707 3.293L5 2.586l-.707.707L2 5.586l-.707.707l1.414 1.414L3.414 7L4 6.414v11.171L3.414 17l-.707-.707l-1.414 1.414l.707.707l2.293 2.293l.707.707l.707-.707L8 18.414l.707-.707l-1.414-1.414l-.707.707l-.586.585V6.414L6.586 7l.707.707l1.414-1.414L8 5.586zM13 4h-1v2h10V4zm-2 7h-1v2h12v-2zm1 7h10v2H12z"/>`), g.Group(children),
	)
}

func Link(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.293 3.96a4.771 4.771 0 1 1 6.747 6.747l-3.03 3.03l-1.415-1.413l3.03-3.031a2.771 2.771 0 1 0-3.918-3.92l-3.031 3.031l-1.414-1.414zm2.12 6.04l-5.415 5.414L8.584 14l5.414-5.414zm-7.01 1.676l-3.03 3.031a2.771 2.771 0 1 0 3.92 3.92l3.03-3.031l1.414 1.414l-3.03 3.03a4.771 4.771 0 1 1-6.748-6.747l3.03-3.03z"/>`), g.Group(children),
	)
}

func LinkOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.778 4.929a3.5 3.5 0 0 0-4.95 0L12.707 7.05l-1.415-1.414l2.122-2.121a5.5 5.5 0 0 1 7.778 7.778l-3.182 3.182a5.5 5.5 0 0 1-7.778 0l-1.533-1.533l1.415-1.414l1.532 1.533a3.5 3.5 0 0 0 4.95 0l3.182-3.182a3.5 3.5 0 0 0 0-4.95m-7.425 6.01a3.5 3.5 0 0 0-4.95 0l-3.182 3.182a3.5 3.5 0 0 0 4.95 4.95l2.122-2.121l1.414 1.414l-2.122 2.121a5.5 5.5 0 0 1-7.778-7.778L5.99 9.525a5.5 5.5 0 0 1 7.778 0l1.296 1.296l-1.414 1.414z"/>`), g.Group(children),
	)
}

func LinkUnlink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m4.304 2.89l3.535 3.536L6.425 7.84L2.89 4.304zm5.949-1.709v4.5h-2v-4.5zM1.347 8.09h4.5v2h-4.5zm13.482-2.454L11.785 8.68l-1.414-1.415l3.044-3.044a4.5 4.5 0 0 1 6.364 6.364l-3.044 3.044l-1.415-1.414l3.044-3.044a2.5 2.5 0 1 0-3.535-3.535m-9.192 9.192a2.5 2.5 0 1 0 3.535 3.536l3.044-3.044l1.414 1.414l-3.044 3.044a4.5 4.5 0 0 1-6.364-6.364l3.044-3.044l1.415 1.414zm12.678-1.081h4.5v2h-4.5zm-2.406 4.406v4.5h-2v-4.5zm1.664-1.993l3.536 3.536l-1.415 1.414l-3.535-3.536z"/>`), g.Group(children),
	)
}

func Liquor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.5 1.086L22.914 4.5L18 9.414v4L9.414 22a2 2 0 0 1-2.828 0L2 17.414a2 2 0 0 1 0-2.828L10.586 6h4zM5 14.414L3.414 16L8 20.586L9.586 19zm10-.828l1-1v-4L20.086 4.5l-.586-.586L15.414 8h-4l-1 1zL13.586 15L9 10.414L6.414 13L11 17.586m-2.002-4.588h2.004v2.004H8.998z"/>`), g.Group(children),
	)
}

func List(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 19.004h2.004V17H2zM7 19h15v-2H7zm-5-5.996h2.004V11H2zM7 13h15v-2H7zM2 7.004h2.004V5H2zM7 7h15V5H7z"/>`), g.Group(children),
	)
}

func ListNumbered(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.5 6h1v2h-4V6h1V4h-1V2h1a2 2 0 0 1 2 2zM7 20v-2h15v2zm0-7v-2h15v2zm0-7V4h15v2zM1.5 9.25h2.47a1.53 1.53 0 0 1 1.082 2.612l-.888.888H5.5v2h-4v-1.336A2 2 0 0 1 2.086 12l.75-.75H1.5zm0 6.5v2h2V18H2v2h1.5v.25h-2v2h2a2 2 0 0 0 2-2v-2.5a2 2 0 0 0-2-2z"/>`), g.Group(children),
	)
}

func Load(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 1h1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12v-1h2v1a9 9 0 1 0 9-9h-1z"/>`), g.Group(children),
	)
}

func Loading(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-opacity=".9" d="M12 2.25c-5.384 0-9.75 4.366-9.75 9.75s4.366 9.75 9.75 9.75v-2.437A7.312 7.312 0 1 1 19.313 12h2.437c0-5.384-4.366-9.75-9.75-9.75"/>`), g.Group(children),
	)
}

func Location(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a7 7 0 0 0-7 7c0 2.862 1.782 5.623 3.738 7.762A26 26 0 0 0 12 20.758q.262-.201.615-.49a26 26 0 0 0 2.647-2.504C17.218 15.623 19 12.863 19 10a7 7 0 0 0-7-7m0 20.214l-.567-.39l-.003-.002l-.006-.005l-.02-.014l-.075-.053l-.27-.197a28 28 0 0 1-3.797-3.44C5.218 16.875 3 13.636 3 9.999a9 9 0 0 1 18 0c0 3.637-2.218 6.877-4.262 9.112a28 28 0 0 1-3.796 3.44a17 17 0 0 1-.345.251l-.021.014l-.006.005l-.002.001zM12 8a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0"/>`), g.Group(children),
	)
}

func LocationEnlargement(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a7 7 0 0 0-7 7c0 2.862 1.782 5.623 3.738 7.762A26 26 0 0 0 12 20.758q.262-.201.615-.49a26 26 0 0 0 2.647-2.504C17.218 15.623 19 12.863 19 10a7 7 0 0 0-7-7m0 20.214l-.567-.39l-.003-.002l-.006-.005l-.02-.014l-.075-.053l-.27-.197a28 28 0 0 1-3.797-3.44C5.218 16.875 3 13.636 3 9.999a9 9 0 0 1 18 0c0 3.637-2.218 6.877-4.262 9.112a28 28 0 0 1-3.796 3.44a17 17 0 0 1-.345.251l-.021.014l-.006.005l-.002.001zM13 6.5v3h3v2h-3v3h-2v-3H8v-2h3v-3z"/>`), g.Group(children),
	)
}

func LocationError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a7 7 0 0 0-7 7c0 2.862 1.782 5.623 3.738 7.762A26 26 0 0 0 12 20.758q.262-.201.615-.49a26 26 0 0 0 2.647-2.504C17.218 15.623 19 12.863 19 10a7 7 0 0 0-7-7m0 20.214l-.567-.39l-.003-.002l-.006-.005l-.02-.014l-.075-.053l-.27-.197a28 28 0 0 1-3.797-3.44C5.218 16.875 3 13.636 3 9.999a9 9 0 0 1 18 0c0 3.637-2.218 6.877-4.262 9.112a28 28 0 0 1-3.796 3.44a17 17 0 0 1-.345.251l-.021.014l-.006.005l-.002.001zM9.879 6.964L12 9.086l2.121-2.122l1.415 1.415l-2.122 2.121l2.122 2.121l-1.415 1.415L12 11.914l-2.121 2.122l-1.415-1.415l2.122-2.121l-2.122-2.121z"/>`), g.Group(children),
	)
}

func LocationOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a3 3 0 1 0 0 6a3 3 0 0 0 0-6M7 6a5 5 0 1 1 6 4.9V17h-2v-6.1A5 5 0 0 1 7 6m-3.895 5H8v2H4.895l-.778 7h15.766l-.778-7H16v-2h4.895l1.222 11H1.883z"/>`), g.Group(children),
	)
}

func LocationParkingPlace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a7 7 0 0 0-7 7c0 2.862 1.782 5.623 3.738 7.762A26 26 0 0 0 12 20.758q.262-.201.615-.49a26 26 0 0 0 2.647-2.504C17.218 15.623 19 12.863 19 10a7 7 0 0 0-7-7m0 20.214l-.567-.39l-.003-.002l-.006-.005l-.02-.014l-.075-.053l-.27-.197a28 28 0 0 1-3.797-3.44C5.218 16.875 3 13.636 3 9.999a9 9 0 0 1 18 0c0 3.637-2.218 6.877-4.262 9.112a28 28 0 0 1-3.796 3.44a17 17 0 0 1-.345.251l-.021.014l-.006.005l-.002.001zM9 7h4.636c.444 0 1.022.16 1.505.574c.52.447.859 1.131.859 2.026s-.34 1.579-.86 2.026a2.37 2.37 0 0 1-1.504.574H11V15H9zm2 3.2h2.636a.4.4 0 0 0 .2-.091a.4.4 0 0 0 .1-.138A.9.9 0 0 0 14 9.6a.9.9 0 0 0-.063-.37a.38.38 0 0 0-.3-.23H11z"/>`), g.Group(children),
	)
}

func LocationReduction(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a7 7 0 0 0-7 7c0 2.862 1.782 5.623 3.738 7.762A26 26 0 0 0 12 20.758q.262-.201.615-.49a26 26 0 0 0 2.647-2.504C17.218 15.623 19 12.863 19 10a7 7 0 0 0-7-7m0 20.214l-.567-.39l-.003-.002l-.006-.005l-.02-.014l-.075-.053l-.27-.197a28 28 0 0 1-3.797-3.44C5.218 16.875 3 13.636 3 9.999a9 9 0 0 1 18 0c0 3.637-2.218 6.877-4.262 9.112a28 28 0 0 1-3.796 3.44a17 17 0 0 1-.345.251l-.021.014l-.006.005l-.002.001zM8 9.5h8v2H8z"/>`), g.Group(children),
	)
}

func LocationSetting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 3a7 7 0 0 1 7 7v1h2v-1a9 9 0 1 0-18 0c0 3.637 2.218 6.876 4.262 9.112a28 28 0 0 0 3.797 3.44c.39.292.797.562 1.198.839l1.134-1.648c-.381-.262-.767-.517-1.137-.794a26 26 0 0 1-3.516-3.187C5.782 15.624 4 12.862 4 10a7 7 0 0 1 7-7m0 5a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0"/><path fill="currentColor" d="M19 12.75v1.376c.715.184 1.352.56 1.854 1.072l1.193-.689l1 1.732l-1.192.688a4 4 0 0 1 0 2.142l1.192.688l-1 1.732l-1.193-.689A4 4 0 0 1 19 21.874v1.376h-2v-1.376a4 4 0 0 1-1.854-1.072l-1.193.689l-1-1.732l1.192-.688a4 4 0 0 1 0-2.142l-1.192-.688l1-1.732l1.193.688A4 4 0 0 1 17 14.127V12.75zm-2.751 4.283a2 2 0 0 0-.25.967c0 .35.091.68.25.967l.036.063a2 2 0 0 0 3.43 0l.036-.063A2 2 0 0 0 20 18c0-.35-.09-.68-.249-.967l-.036-.063a2 2 0 0 0-3.43 0z"/>`), g.Group(children),
	)
}

func LockOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 3a4 4 0 0 0-4 4v3h12.5v12h-17V10H6V7a6 6 0 0 1 11.725-1.8l.3.954l-1.908.6l-.3-.954A4 4 0 0 0 12 3m-6.5 9v8h13v-8zM9 15h6v2H9z" clip-rule="evenodd"/>`), g.Group(children),
	)
}

func LockOffFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 7a4 4 0 0 1 7.817-1.2l.3.954l1.908-.6l-.3-.954A6.001 6.001 0 0 0 6 7v3H3.5v12h17V10H8zm7 8v2H9v-2z"/>`), g.Group(children),
	)
}

func LockOn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a4 4 0 0 1 4 4v3H8V7a4 4 0 0 1 4-4m6 7V7A6 6 0 0 0 6 7v3H3.5v12h17V10zM5.5 12h13v8h-13zM9 15h6v2H9z"/>`), g.Group(children),
	)
}

func LockOnFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 10H3.5v12h17V10H18V7A6 6 0 0 0 6 7zm2-3a4 4 0 1 1 8 0v3H8zm1 10v-2h6v2z"/>`), g.Group(children),
	)
}

func LockTime(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.5 7.5a6.5 6.5 0 0 1 13 0V9h2v2.5h-2V11h-13v9h6v2h-8V9h2zm2 1.5h9V7.5a4.5 4.5 0 1 0-9 0zm11 6a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M13 18.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m-4-4h2.5v2H9zm10.5 1.752v1.834l1.414 1.414l-1.414 1.414l-2-2v-2.662z"/>`), g.Group(children),
	)
}

func Login(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 11h9.586l-3.5-3.5L10.5 6.086L16.414 12L10.5 17.914L9.086 16.5l3.5-3.5H3zm11 8.5h5v-15h-5v-2h7v19h-7z"/>`), g.Group(children),
	)
}

func LogoAdobeIllustrate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm4.756 2h2.488l3.6 12h-2.091l-1.183-3.945H8.428L7.244 18H5.153zm.272 6.055h1.943L10 8.815zM15.996 9H18v2.004h-2.004zM18 12v6h-2v-6z"/>`), g.Group(children),
	)
}

func LogoAdobePhotoshop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 13v5H6V6h3.5a3.5 3.5 0 1 1 0 7zm0-5v3h1.5a1.5 1.5 0 0 0 0-3zm5.5 4.453A2.453 2.453 0 0 1 15.953 10H18v2h-2.047a.453.453 0 0 0-.143.883l1.013.337a2.453 2.453 0 0 1-.776 4.78H13.5v-2h2.547a.453.453 0 0 0 .143-.883l-1.013-.337a2.45 2.45 0 0 1-1.677-2.327"/><path fill="currentColor" d="M22 2H2v20h20zM4 20V4h16v16z"/>`), g.Group(children),
	)
}

func LogoAdobePhotoshopOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16v-6.5h-1c-1.015 0-1.99.403-2.707 1.121l-.293.293V18h-2v-7h2v1.331a5.83 5.83 0 0 1 3-.831h1V4zm4 2v10h5v2H6V6z"/>`), g.Group(children),
	)
}

func LogoAndroid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m4.372 3.636l2.116 3.702A11.95 11.95 0 0 1 12 6c1.986 0 3.86.483 5.512 1.338l2.116-3.702l1.736.992l-2.158 3.776A11.98 11.98 0 0 1 24 18v2H0v-2c0-3.924 1.884-7.407 4.793-9.596L2.636 4.628zM12 8a9.95 9.95 0 0 0-5.348 1.55A9.99 9.99 0 0 0 2 18h20a9.99 9.99 0 0 0-4.652-8.451A9.95 9.95 0 0 0 12 8m-6 5h2.004v2.004H6zm10 0h2.004v2.004H16z"/>`), g.Group(children),
	)
}

func LogoApple(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.862 3.94a3.5 3.5 0 0 1 3.5-3.5h1.05v1.05a3.5 3.5 0 0 1-3.5 3.5h-1.05zm4.452 1.074c.601.04 2.183.187 3.609 1.283c.418.322.82.723 1.176 1.22l.666.93l-1.01.536c-.014.008-.197.11-.443.307c-.603.484-1.43 1.403-1.417 2.805c.014 1.257.578 2.073 1.155 2.598c.508.462 1.01.68 1.106.722q.028.012.007.005l.937.298l-.282.942c-.01.032-.039.127-.09.273a12.5 12.5 0 0 1-1.585 3.024c-.451.631-.982 1.365-1.62 1.934c-.66.588-1.495 1.056-2.558 1.073h-.004h.008h-.004c-.936.021-1.583-.246-2.095-.457l-.024-.01c-.502-.207-.895-.366-1.547-.366c-.708 0-1.132.17-1.662.386l-.026.01c-.475.194-1.071.437-1.894.47l-.01.001c-1.09.034-1.97-.494-2.64-1.093c-.664-.594-1.23-1.357-1.69-1.987l-.002-.003c-1.037-1.431-1.933-3.418-2.316-5.49c-.382-2.065-.274-4.34.87-6.237c1.142-1.894 3.175-3.101 5.407-3.13H8.33h.011h-.005c1.004-.023 1.944.333 2.596.58l.058.023c.27.102.482.182.67.239c.193.057.295.07.34.07c.019 0 .099-.007.28-.062c.176-.053.379-.129.643-.228l.05-.019c.795-.298 1.989-.746 3.341-.647m-.143 1.995c-.884-.066-1.707.23-2.546.544l-.018.007c-.242.09-.503.188-.748.262A3 3 0 0 1 12 7.97c-.331 0-.65-.075-.916-.155c-.262-.08-.537-.183-.785-.277l-.018-.007c-.72-.273-1.323-.488-1.903-.474h-.012c-1.52.019-2.928.842-3.723 2.163v.001c-.805 1.333-.947 3.058-.618 4.84c.329 1.777 1.105 3.486 1.97 4.68c.472.647.922 1.241 1.405 1.674c.477.426.874.594 1.24.584c.454-.019.771-.142 1.243-.334l.012-.005c.574-.234 1.296-.529 2.404-.529c1.057 0 1.752.287 2.297.512l.013.005c.494.204.812.328 1.314.317h.008c.442-.007.836-.189 1.261-.567c.446-.398.855-.948 1.324-1.605l.001-.002a10.5 10.5 0 0 0 1.092-1.936a6 6 0 0 1-.904-.682a5.43 5.43 0 0 1-1.81-4.057v-.002c-.018-2.05 1.078-3.421 1.9-4.158l-.091-.073c-.943-.725-2.028-.841-2.524-.873z"/>`), g.Group(children),
	)
}

func LogoAppleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.718 5.492q-.862 0-1.603.264t-1.35.523t-1.057.258q-.47 0-1.045-.247a24 24 0 0 0-1.235-.488A4.1 4.1 0 0 0 8.02 5.56q-1.413 0-2.722.77q-1.31.768-2.149 2.274q-.838 1.505-.838 3.711q0 1.379.321 2.757q.322 1.38.862 2.608q.54 1.23 1.195 2.172q.85 1.182 1.66 2.033q.81.85 1.878.85q.7 0 1.218-.23q.517-.23 1.097-.465t1.454-.235q.666 0 1.114.137q.448.139.821.316q.375.179.787.316q.414.138 1 .138q.77 0 1.384-.408q.616-.407 1.138-1.05q.522-.645 1.017-1.345q.54-.793.878-1.522a13 13 0 0 0 .672-1.706q-.034-.012-.5-.259t-1.056-.787q-.593-.54-1.034-1.413q-.443-.873-.443-2.137q0-1.103.35-1.901q.351-.8.822-1.321q.471-.523.845-.799q.373-.276.43-.31q-.574-.827-1.252-1.287a5.2 5.2 0 0 0-1.332-.666a6 6 0 0 0-1.167-.259a8 8 0 0 0-.752-.051m-.804-1.862q.529-.631.861-1.464a4.6 4.6 0 0 0 .333-1.718q0-.253-.034-.448q-.861.034-1.798.517a5 5 0 0 0-1.545 1.206a5.5 5.5 0 0 0-.873 1.379a4 4 0 0 0-.38 1.7a1.7 1.7 0 0 0 .046.414q.15.034.31.034q.771 0 1.638-.46a4.7 4.7 0 0 0 1.442-1.16"/>`), g.Group(children),
	)
}

func LogoBehance(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.5 4c1.312 0 2.462.425 3.284 1.262C10.603 6.096 11 7.24 11 8.5c0 1.223-.375 2.339-1.147 3.166q.391.217.728.51c.937.816 1.419 1.98 1.419 3.324s-.482 2.508-1.419 3.324C9.661 19.626 8.408 20 7.014 20H2V4zM4 13v5h3.014c1.053 0 1.792-.282 2.254-.684c.444-.387.732-.972.732-1.816s-.288-1.43-.732-1.816C8.806 13.282 8.067 13 7.014 13zm0-7v5h2.5c.875 0 1.476-.275 1.857-.663C8.74 9.946 9 9.34 9 8.5s-.26-1.446-.643-1.837C7.976 6.275 7.375 6 6.5 6zm9.5.5h8v2h-8zm.661 5.027C14.954 10.576 16.111 10 17.5 10c1.643 0 2.821.718 3.552 1.758c.699.993.948 2.213.948 3.242v1h-6.883c.119.477.321.882.581 1.193c.41.494 1.004.807 1.802.807c.876 0 1.45-.263 1.79-.495a2.2 2.2 0 0 0 .44-.393l.005-.006l.54-.824l1.672 1.097l-.549.836v.002l-.002.002l-.003.005l-.007.01l-.074.102a3 3 0 0 1-.188.22a4 4 0 0 1-.712.605c-.642.435-1.6.839-2.912.839c-1.39 0-2.546-.576-3.339-1.527C13.384 17.54 13 16.3 13 15s.384-2.54 1.161-3.473M15.117 14h4.769a3.2 3.2 0 0 0-.47-1.091c-.362-.516-.934-.909-1.916-.909c-.798 0-1.391.313-1.802.807a2.96 2.96 0 0 0-.58 1.193"/>`), g.Group(children),
	)
}

func LogoChrome(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.001 1c6.075 0 11 4.925 11 11s-4.925 11-11 11s-11-4.924-11-11s4.925-11 11-11m-6.84 5.15L7.384 10l.212-.367A5.01 5.01 0 0 1 12.114 7h7.371a8.99 8.99 0 0 0-7.484-4a8.98 8.98 0 0 0-6.84 3.15M3.93 8.016A9 9 0 0 0 3.001 12c0 4.409 3.17 8.077 7.356 8.85L12.58 17h-.496a5 5 0 0 1-4.492-2.64zm8.662 12.965a9 9 0 0 0 7.898-11.98h-4.444l.24.415a5.01 5.01 0 0 1 0 5.169zm-3.186-7.48A3 3 0 0 0 12 15h.043a3 3 0 0 0 1.459-.403a3 3 0 0 0 1.069-1.047l.029-.05q.078-.136.14-.276c.172-.385.262-.803.263-1.226A3 3 0 0 0 14.6 10.5l-.03-.051a3 3 0 0 0-1.068-1.047c-.45-.26-.941-.39-1.428-.402H12q-.177 0-.35.02a3 3 0 0 0-2.246 1.48l-.054.094A3 3 0 0 0 9 12.001c0 .505.125.981.346 1.399z"/>`), g.Group(children),
	)
}

func LogoChromeFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 23c6.075 0 11-4.925 11-11c0-1.131-.17-2.223-.488-3.25h-7.048A4.73 4.73 0 0 1 16.75 12a4.73 4.73 0 0 1-.636 2.376l-4.96 8.592Q11.571 23 12 23"/><path fill="currentColor" d="m9.56 22.728l3.523-6.102q-.523.123-1.083.124a4.75 4.75 0 0 1-4.114-2.374L2.925 5.782A10.95 10.95 0 0 0 1 12c0 5.236 3.659 9.618 8.56 10.728"/><path fill="currentColor" d="M3.93 4.524A10.97 10.97 0 0 1 12 1a11 11 0 0 1 9.924 6.25H12a4.75 4.75 0 0 0-4.548 3.374z"/><path fill="currentColor" d="m9.185 13.625l-.004-.006a3.25 3.25 0 1 1 .003.005"/>`), g.Group(children),
	)
}

func LogoCinemaFourD(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m2.747 17.95l.02.041l.036.057q.263.4.567.774A10.98 10.98 0 0 0 12 23c6.075 0 11-4.925 11-11c0-2.479-.82-4.766-2.203-6.605a10.4 10.4 0 0 0-3.851-3.223A10.95 10.95 0 0 0 12 1C5.925 1 1 5.925 1 12c0 2.193.641 4.235 1.747 5.95m.674-8.678a9.004 9.004 0 0 1 11.609-5.75a.52.52 0 0 1-.396.361l-.007.002c-.988.269-1.959.52-2.93.705h-.006c-2.108.414-4.113 1.16-5.897 2.524A10 10 0 0 0 3.42 9.272m1.523 8.316a9 9 0 0 1-.447-.616c-.918-1.985-.878-3.615-.383-4.92c.514-1.357 1.564-2.473 2.806-3.285l.033-.02l.03-.024q.273-.21.555-.398c-.747 1.951-.572 4.203.709 6.132c1.527 2.299 4.242 3.396 6.836 3.051c-1.498 1.664-3.892 2.736-5.97 2.498c-1.55-.177-3.08-1.094-4.169-2.418M20.35 8.635a9.004 9.004 0 0 1-6.065 12.073c1.11-.661 2.087-1.528 2.814-2.51h.001c1.315-1.78 1.997-3.794 2.313-5.894c.142-.941.312-1.877.502-2.817q.106-.518.435-.852m-3.521-4.231a9 9 0 0 1 2.386 2.215q.084.117.165.236a2 2 0 0 0-.25.185l-.007.006l-.007.005a3.62 3.62 0 0 0-1.16 2.037v.001a59 59 0 0 0-.52 2.918c-.154 1.02-.4 1.985-.766 2.888c-2.267 1.32-5.297.653-6.757-1.545c-1.474-2.218-.898-5.05 1.297-6.603q.424-.108.863-.193c1.05-.2 2.08-.468 3.075-.738a2.5 2.5 0 0 0 1.68-1.412"/>`), g.Group(children),
	)
}

func LogoCodepen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .807l11 7.15v8.086l-11 7.15l-11-7.15V7.957zm-9 9.614v3.158L5.256 12zm.788 5.048L11 20.157V16.02l-4-2.8zM8.744 12L12 14.28L15.256 12L12 9.72zM13 7.98l4 2.8l3.212-2.248L13 3.842zm-2-4.137L3.788 8.532L7 10.779l4-2.8zm10 6.579L18.744 12L21 13.58zm-.788 5.048L17 13.221l-4 2.8v4.136z"/>`), g.Group(children),
	)
}

func LogoCodesandbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.34 6.423L12 .845l9.66 5.578v11.154L12 23.155l-9.66-5.578zM12 3.155L9.67 4.5L12 5.845L14.33 4.5zm4.33 2.5L12 8.155l-4.33-2.5L5.34 7L12 10.845L18.66 7zm3.33 3.077L13 12.577v7.69l2.34-1.35v-4.994l4.32-2.495zm0 5.006l-2.32 1.34v2.684l2.32-1.34zm-15.32-2.31l4.32 2.495v4.994l2.34 1.35v-7.69L4.34 8.732zm0 2.31v2.685l2.32 1.34v-2.686z"/>`), g.Group(children),
	)
}

func LogoDribbble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.908 3.982a9.02 9.02 0 0 0-4.679 5.99a26.9 26.9 0 0 0 7.882-1.548a34 34 0 0 0-3.203-4.442M3 11.983V12a8.96 8.96 0 0 0 2.005 5.664a16.04 16.04 0 0 1 7.909-5.838q-.381-.831-.806-1.637A28.9 28.9 0 0 1 3 11.983m3.45 7.103A8.96 8.96 0 0 0 12 21a9 9 0 0 0 3.522-.715a33.7 33.7 0 0 0-1.825-6.611a14.04 14.04 0 0 0-7.247 5.412m9.206-5.89a36 36 0 0 1 1.708 6.032a9 9 0 0 0 3.543-5.926a14.1 14.1 0 0 0-5.25-.107m5.315-1.92a8.95 8.95 0 0 0-1.893-4.836a29 29 0 0 1-5.113 2.984q.48.925.909 1.881a16.2 16.2 0 0 1 6.096-.028m-7.98-3.606A27 27 0 0 0 17.656 5A8.96 8.96 0 0 0 12 3c-.724 0-1.428.085-2.101.246A36 36 0 0 1 12.99 7.67M1.053 10.9a11.01 11.01 0 0 1 6.733-9.065A11 11 0 0 1 12 1c3.121 0 5.94 1.301 7.942 3.389a10.97 10.97 0 0 1 3.043 8.199a11 11 0 0 1-5.895 9.166A10.96 10.96 0 0 1 12 23c-2.921 0-5.578-1.14-7.547-2.997A10.97 10.97 0 0 1 1.054 10.9"/>`), g.Group(children),
	)
}

func LogoFacebook(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 6a6 6 0 0 1 6-6h5v6.5h-4v2h4.247L17.802 15H15v9H8v-9H4.25V8.5H8zm6-4a4 4 0 0 0-4 4v4.5H6.25V13H10v9h3v-9h3.198l.555-2.5H13v-4a2 2 0 0 1 2-2h2V2z"/>`), g.Group(children),
	)
}

func LogoFigma(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 5A4.5 4.5 0 0 1 8.5.5h7a4.5 4.5 0 0 1 2.829 8A4.5 4.5 0 0 1 13 15.742V19a4.5 4.5 0 1 1-7.329-3.5A4.5 4.5 0 0 1 4 12c0-1.414.652-2.675 1.671-3.5A4.5 4.5 0 0 1 4 5m4.5 2.5H11v-5H8.5a2.5 2.5 0 0 0 0 5m4.5-5v5h2.5a2.5 2.5 0 0 0 0-5zm-2 7H8.5a2.5 2.5 0 0 0 0 5H11zm0 7H8.5A2.5 2.5 0 1 0 11 19zm2-4.5a2.5 2.5 0 1 0 5 0a2.5 2.5 0 0 0-5 0"/>`), g.Group(children),
	)
}

func LogoFramer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.086 1H19.5v8.5h-5.086l6.5 6.5H13v7.914l-8.5-8.5V7.5h5.086zm9.328 6.5H17.5V3H7.914zm-.828 2H6.5v5.086l4.5 4.5V14h5.086z"/>`), g.Group(children),
	)
}

func LogoGithub(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.307 3.047c-.328 1.282-.005 2.175.069 2.358l.233.582l-.422.463C4.447 7.262 4 8.278 4 9.572c0 2.491.747 3.833 1.686 4.62c.987.829 2.355 1.19 3.869 1.36l2.276.255l-1.735 1.495c-.242.21-.606.751-.606 1.856v.342h.01v2h-.002l.005.732v.002l.001.187l.004.579v1h-2v-1l-.004-.561l-.001-.193v-.002l-.005-.744h-.556a4 4 0 0 1-3.124-1.501l-.55-.688a3 3 0 0 0-1.001-.81l-1.109-.554l.895-1.789l1.108.555A5 5 0 0 1 4.83 18.06l.55.688a2 2 0 0 0 1.562.751h.548v-.342c0-.708.11-1.334.303-1.869c-1.186-.262-2.387-.72-3.392-1.564C2.903 14.469 2 12.5 2 9.572c0-1.586.495-2.925 1.33-4.016c-.2-.801-.348-2.197.33-3.893l.173-.433l.444-.146h.002l.002-.001l.004-.001l.008-.003l.018-.005a1 1 0 0 1 .15-.036q.123-.024.304-.034c.243-.013.561.004.957.087c.718.15 1.684.515 2.928 1.31a14.2 14.2 0 0 1 6.7 0c1.244-.795 2.21-1.16 2.928-1.31c.395-.083.713-.1.957-.087a2.3 2.3 0 0 1 .411.058l.042.012l.019.005l.008.003l.004.001h.002l.002.001l.443.146l.174.433c.678 1.696.53 3.092.33 3.893c.834 1.09 1.33 2.43 1.33 4.016c0 2.929-.903 4.897-2.4 6.153c-1.006.843-2.208 1.302-3.393 1.564c.194.535.303 1.161.303 1.869c0 1.26-.008 2.357-.013 3.087l-.002.194l-.003.561v1h-2v-1l.003-.58l.002-.188c.005-.73.013-1.821.013-3.074c0-1.105-.364-1.646-.606-1.856l-1.735-1.495l2.276-.256c1.514-.17 2.881-.53 3.869-1.359c.939-.787 1.686-2.129 1.686-4.62c0-1.294-.448-2.31-1.187-3.122l-.422-.463l.233-.582c.073-.183.397-1.076.069-2.358l-.006.001c-.477.1-1.328.398-2.597 1.26l-.385.263l-.448-.126a12.13 12.13 0 0 0-6.514 0l-.448.126l-.385-.262c-1.27-.863-2.12-1.16-2.597-1.26z"/>`), g.Group(children),
	)
}

func LogoGithubFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 .999c-6.074 0-11 5.05-11 11.278c0 4.983 3.152 9.21 7.523 10.702c.55.104.727-.246.727-.543v-2.1c-3.06.683-3.697-1.33-3.697-1.33c-.5-1.304-1.222-1.65-1.222-1.65c-.998-.7.076-.686.076-.686c1.105.08 1.686 1.163 1.686 1.163c.98 1.724 2.573 1.226 3.201.937c.098-.728.383-1.226.698-1.508c-2.442-.286-5.01-1.253-5.01-5.574c0-1.232.429-2.237 1.132-3.027c-.114-.285-.49-1.432.107-2.985c0 0 .924-.303 3.026 1.156c.877-.25 1.818-.375 2.753-.38c.935.005 1.876.13 2.755.38c2.1-1.459 3.023-1.156 3.023-1.156c.598 1.554.222 2.701.108 2.985c.706.79 1.132 1.796 1.132 3.027c0 4.332-2.573 5.286-5.022 5.565c.394.35.754 1.036.754 2.088v3.095c0 .3.176.652.734.542C19.852 21.484 23 17.258 23 12.277C23 6.048 18.075.999 12 .999"/>`), g.Group(children),
	)
}

func LogoGitlab(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.78 1h1.936l2.5 7.333h5.568L17.284 1h1.937l4.454 13.362L12 23.257L.325 14.362zm.978 3.389l-3.083 9.249L12 20.743l9.325-7.105l-3.083-9.25l-2.026 5.945H7.784z"/>`), g.Group(children),
	)
}

func LogoIe(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.427 6.194c-.672.423-1.359.907-2.048 1.45a4.6 4.6 0 0 1 2.088.89h-3.164a29 29 0 0 0-2.076 2h8.505L17.35 9.25c-.387-1.303-1.492-2.408-2.924-3.056"/><path fill="currentColor" d="M8.325 22.28c1.143.419 2.377.647 3.664.647c3.48 0 6.582-1.67 8.524-4.26c.276-.36.512-.727.73-1.106c.112-.186.209-.369.3-.553l.724-1.447h-7.314l-.265.197a4.5 4.5 0 0 1-2.688.876c-1.077 0-1.909-.366-2.53-.951a4.2 4.2 0 0 1-.937-1.344l11.187-.105h2.927v-1.966q0-1.093-.213-2.141v-.002a11 11 0 0 0-.352-1.294c.516-1.248.823-2.46.846-3.555c.023-1.168-.28-2.286-1.1-3.105c-.778-.779-1.832-1.09-2.937-1.096c-1.1-.006-2.32.284-3.573.786q-.191.078-.385.16a10.7 10.7 0 0 0-2.944-.411C6.103 1.61 1.33 6.382 1.33 12.268c0 1.103.168 2.167.479 3.168c-.472 1.212-.743 2.39-.736 3.455c.008 1.105.319 2.158 1.096 2.937c.816.822 1.936 1.124 3.102 1.1c.948-.02 1.984-.252 3.054-.648M8.8 8.803c2.45-2.445 5.058-4.201 7.262-5.084c1.104-.442 2.06-.647 2.817-.643c.752.005 1.236.21 1.535.51c.319.318.53.837.514 1.65a6 6 0 0 1-.125 1.039a10.7 10.7 0 0 0-1.81-2.039a9.6 9.6 0 0 0-1.86.554q-.23.09-.463.195a8.68 8.68 0 0 1 3.804 5.542c.113.555.171 1.12.173 1.708h-.931l-12.055.114a24 24 0 0 0-1.314 1.828c.333 1.139.912 2.172 1.751 2.963c1.002.944 2.328 1.495 3.902 1.495a6.5 6.5 0 0 0 3.598-1.073h3.244a8.63 8.63 0 0 1-6.853 3.366a8.65 8.65 0 0 1-7.19-3.834l-.015.035c-.359.9-.56 1.7-.62 2.376c.473.512.996.978 1.561 1.389q-.258.03-.495.035c-.815.017-1.33-.196-1.642-.51l-.003-.003c-.3-.3-.507-.785-.512-1.537c-.005-.76.2-1.716.64-2.82c.882-2.205 2.638-4.812 5.087-7.256m-5.468 3.634l-.002-.168a8.66 8.66 0 0 1 8.68-8.658c-1.548 1.005-3.124 2.28-4.623 3.776c-1.638 1.635-3.012 3.365-4.055 5.05"/>`), g.Group(children),
	)
}

func LogoIeFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10.13 22.257l-.59.231c-1.323.57-2.614.911-3.772.935c-1.166.024-2.286-.278-3.102-1.1c-.777-.779-1.088-1.832-1.096-2.937c-.007-1.1.282-2.322.784-3.576c1.004-2.512 2.935-5.34 5.53-7.93c2.594-2.587 5.421-4.518 7.931-5.524c1.253-.502 2.474-.792 3.573-.786c1.105.007 2.16.317 2.938 1.096c.82.82 1.122 1.937 1.099 3.104c-.024 1.162-.369 2.456-.942 3.784l-.004.009l-.464 1.047q.133.813.133 1.658v1.466h-2.432l-11.901.112c.273.886.719 1.642 1.313 2.2c.716.676 1.671 1.088 2.872 1.088a5 5 0 0 0 2.986-.975l.133-.098h6.34l-.362.724a8 8 0 0 1-.284.523a10 10 0 0 1-.698 1.058a10.13 10.13 0 0 1-8.125 4.06q-.955 0-1.86-.169m2.81-16.032a27.6 27.6 0 0 0-3.643 3.071q-.368.367-.715.739h8.48l-.19-.643c-.44-1.486-2.003-2.749-3.932-3.167m-9.23 11.93a5.7 5.7 0 0 0-.14 1.217c.005.753.212 1.238.512 1.538l.003.003c.311.314.827.527 1.643.51c.423-.008.903-.08 1.43-.216a10.2 10.2 0 0 1-3.449-3.053M21.055 7.68c.238-.737.358-1.393.37-1.951c.016-.813-.196-1.332-.514-1.65c-.3-.3-.784-.506-1.535-.51c-.495-.003-1.073.083-1.722.266a10.2 10.2 0 0 1 3.401 3.845M6.82 6.817C4.704 8.93 3.01 11.21 1.89 13.37q-.059-.544-.059-1.102C1.831 6.658 6.38 2.11 11.99 2.11q.459 0 .907.04c-2.014 1.115-4.115 2.71-6.076 4.667"/>`), g.Group(children),
	)
}

func LogoInstagram(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.95 1h.1c1.827 0 3.266 0 4.42.105c1.178.106 2.156.328 3.03.833A7 7 0 0 1 22.062 4.5c.505.874.727 1.852.833 3.03C23 8.684 23 10.123 23 11.95v.1c0 1.827 0 3.266-.105 4.42c-.106 1.178-.328 2.156-.833 3.03a7 7 0 0 1-2.562 2.562c-.874.505-1.852.727-3.03.833c-1.154.105-2.593.105-4.42.105h-.1c-1.827 0-3.266 0-4.42-.105c-1.178-.106-2.156-.328-3.03-.833A7 7 0 0 1 1.938 19.5c-.505-.874-.727-1.852-.833-3.03C1 15.316 1 13.877 1 12.05v-.1c0-1.827 0-3.266.105-4.42c.106-1.178.328-2.156.833-3.03A7 7 0 0 1 4.5 1.938c.874-.505 1.852-.727 3.03-.833C8.684 1 10.123 1 11.95 1M7.71 3.096c-1.039.095-1.691.274-2.21.574A5 5 0 0 0 3.67 5.5c-.3.519-.48 1.171-.574 2.21C3.001 8.764 3 10.112 3 12s.001 3.236.096 4.29c.095 1.039.274 1.691.574 2.21a5 5 0 0 0 1.83 1.83c.519.3 1.171.48 2.21.574c1.054.095 2.402.096 4.29.096s3.236-.001 4.29-.096c1.039-.095 1.691-.274 2.21-.574a5 5 0 0 0 1.83-1.83c.3-.519.48-1.171.574-2.21c.095-1.053.096-2.402.096-4.29s-.001-3.236-.096-4.29c-.095-1.039-.274-1.691-.574-2.21a5 5 0 0 0-1.83-1.83c-.519-.3-1.171-.48-2.21-.574C15.237 3.001 13.888 3 12 3s-3.236.001-4.29.096m9.04 3.156c0-.552.45-1.002 1.002-1.002s1.002.45 1.002 1.002s-.45 1.002-1.002 1.002s-1.002-.45-1.002-1.002M12 8.5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6.5 12a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0"/>`), g.Group(children),
	)
}

func LogoQq(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.897 3.087C7.188 1.24 9.213.012 11.997.013c2.786 0 4.81 1.228 6.102 3.074c1.261 1.804 1.789 4.138 1.789 6.385c0 .137-.004.331-.007.487l-.002.096l.914 2.286l.003.007c.273.712.55 1.47.755 2.132c.486 1.563.669 2.723.684 3.535c.007.401-.025.756-.104 1.046a1.7 1.7 0 0 1-.209.478a1.13 1.13 0 0 1-.8.51c-.436.052-.772-.156-.9-.242a3 3 0 0 1-.446-.382a6 6 0 0 1-.217-.239a7.3 7.3 0 0 1-.936 1.596c.196.103.383.216.55.34c.348.26.601.6.692 1.008c.087.39.003.753-.153 1.017a1.26 1.26 0 0 1-.523.475a2 2 0 0 1-.408.15a5.5 5.5 0 0 1-.828.135c-.592.059-1.34.08-2.097.078c-1.445-.002-3.035-.084-3.855-.168c-.82.084-2.41.166-3.855.168a22 22 0 0 1-2.097-.078a5.5 5.5 0 0 1-.828-.135a2 2 0 0 1-.408-.15a1.25 1.25 0 0 1-.522-.474l.002.004a1.4 1.4 0 0 1-.157-1.023c.09-.41.346-.75.695-1.008l.001-.001c.166-.123.352-.235.546-.337a7.3 7.3 0 0 1-.936-1.597a6 6 0 0 1-.218.239a3 3 0 0 1-.445.382c-.128.085-.464.295-.901.242h-.003a1.13 1.13 0 0 1-.796-.506a1.7 1.7 0 0 1-.21-.48a3.8 3.8 0 0 1-.106-1.048c.014-.813.196-1.972.681-3.534c.205-.661.482-1.42.756-2.134l.005-.013l.911-2.279l-.009-.583c0-2.247.528-4.581 1.79-6.385m-1.825 13.18l1.851-3.004v3.53c0 1.042.557 2.521 1.837 3.605l1.394 1.18l-1.319.407h.308c1.467-.002 3.062-.09 3.746-.167l.112-.013l.112.013c.684.077 2.28.165 3.746.167h.307l-1.32-.408l1.395-1.18c1.28-1.083 1.836-2.562 1.836-3.605v-3.53l1.852 3.006l.003.004l.01.017l.043.069l.01.016a18 18 0 0 0-.354-1.3a35 35 0 0 0-.71-2.004l-1.057-2.64v-.193l.007-.32c.003-.157.007-.332.007-.445c0-1.973-.47-3.867-1.428-5.239c-.929-1.328-2.35-2.22-4.462-2.22c-2.113 0-3.534.892-4.463 2.22C6.576 5.604 6.107 7.5 6.107 9.472c0 .2.009.557.012.694l.001.051l.004.203l-1.06 2.65a34 34 0 0 0-.71 2.003m-.282 1.195l-.003.004l-.01.017l-.043.069l-.017.026q.129-.575.355-1.31"/>`), g.Group(children),
	)
}

func LogoTwitter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 4a4 4 0 0 0-4 4v2h-1a12.96 12.96 0 0 1-8.803-3.434a9.504 9.504 0 0 0 5.806 10.77l1.731.686l-1.528 1.064a13.5 13.5 0 0 1-2.554 1.396C6.707 20.819 7.832 21 9 21c6.075 0 11-4.925 11-11V8.473l.227-.277c.546-.666.93-1.166 1.279-1.657c-.278.111-.548.2-.784.254l-.774.176l-.348-.714A4 4 0 0 0 16 4m6.976 1.755l1.018.696l-.565.825c-.417.61-.832 1.171-1.429 1.91V10c0 7.18-5.82 13-13 13c-2.45 0-4.743-.678-6.7-1.858L.05 19.787l2.581-.484a11.4 11.4 0 0 0 2.912-.965A11.49 11.49 0 0 1 0 8.5c0-1.496.286-2.927.807-4.24l.624-1.573l1.077 1.306A10.98 10.98 0 0 0 10 7.955a6 6 0 0 1 10.947-3.351c.372-.166.764-.385 1.008-.582l.778-.627l1.255 1.557l-.778.627a4 4 0 0 1-.234.176"/>`), g.Group(children),
	)
}

func LogoWechat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.796 17.027H8.75c-1.153 0-2.254-.188-3.262-.53L2.65 17.92l.352-2.712C1.162 13.855 0 11.861 0 9.64c0-4.083 3.918-7.39 8.75-7.39c4.174 0 7.665 2.468 8.54 5.77a9 9 0 0 0-.6-.02c-4.364 0-8.19 3.037-8.19 7.11c0 .67.104 1.312.296 1.917M6 8a1 1 0 1 0 0-2a1 1 0 0 0 0 2m5.5.007a1 1 0 1 0 0-2a1 1 0 0 0 0 2"/><path fill="currentColor" d="M21.874 19.52C23.187 18.405 24 16.863 24 15.16C24 11.758 20.754 9 16.75 9S9.5 11.758 9.5 15.161s3.246 6.161 7.25 6.161c.95 0 1.856-.155 2.686-.437l2.438 1.407zm-7.564-5.362a1 1 0 1 1 0-2a1 1 0 0 1 0 2m4.88 0a1 1 0 1 1 0-2a1 1 0 0 1 0 2"/>`), g.Group(children),
	)
}

func LogoWechatStroke(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.75 4.25C4.857 4.25 2 6.884 2 9.794c0 1.6.835 3.086 2.252 4.13a1 1 0 0 1 .399.933l-.07.534l.83-.416a1 1 0 0 1 .77-.054a8 8 0 0 0 2.569.417q.381 0 .75-.034c.004-3.21 2.674-5.678 5.946-6.201c-.41-2.623-3.114-4.853-6.696-4.853m8.707 4.78C21.003 9.335 24 11.904 24 15.311c0 1.637-.711 3.102-1.833 4.199v1.95a1 1 0 0 1-1.5.865l-1.725-.996a8.3 8.3 0 0 1-2.192.293c-3.122 0-5.924-1.747-6.893-4.344a10.2 10.2 0 0 1-3.91-.334l-2.151 1.08A1 1 0 0 1 2.356 17l.239-1.841C1.018 13.815 0 11.926 0 9.794C0 5.476 4.083 2.25 8.75 2.25c4.388 0 8.259 2.85 8.707 6.78M5.247 7.496a1 1 0 0 1 1-1h.003a1 1 0 0 1 1 1V7.5a1 1 0 0 1-1 1h-.004a1 1 0 0 1-1-1zm5.003.011a1 1 0 0 1 1-1h.004a1 1 0 0 1 1 1v.004a1 1 0 0 1-1 1h-.004a1 1 0 0 1-1-1zM16.75 11q-.11 0-.22.004c-2.948.1-5.03 2.126-5.03 4.307q0 .337.061.66c.368 1.959 2.447 3.651 5.189 3.651a6.2 6.2 0 0 0 1.995-.323a1 1 0 0 1 .82.08l.602.348v-.658a1 1 0 0 1 .353-.763c.942-.799 1.48-1.862 1.48-2.995C22 13.077 19.808 11 16.75 11m-3.194 2.656a1 1 0 0 1 1-1h.004a1 1 0 0 1 1 1v.004a1 1 0 0 1-1 1h-.004a1 1 0 0 1-1-1zm4.38.004a1 1 0 0 1 1-1h.004a1 1 0 0 1 1 1v.004a1 1 0 0 1-1 1h-.004a1 1 0 0 1-1-1z"/>`), g.Group(children),
	)
}

func LogoWecom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19.443 17.523l-.075-.015l-.077-.01a3.9 3.9 0 0 1-2.165-1.196a.304.304 0 0 0-.428 0a.3.3 0 0 0-.021.4l.021.025q.021.019.042.035l.087.084a3.85 3.85 0 0 1 1.068 1.976q.005.061.014.122q.011.067.03.133c.055.203.164.395.325.555c.494.49 1.296.49 1.79 0a1.248 1.248 0 0 0-.611-2.109m4.187-3.203a1.273 1.273 0 0 0-1.79 0a1.25 1.25 0 0 0-.338.604q-.01.038-.016.075l-.01.076a3.86 3.86 0 0 1-1.078 2.027q-.062.061-.127.12a.3.3 0 0 0 0 .425c.11.11.284.116.404.021q.012-.009.023-.021l.036-.041q.04-.044.084-.087a3.9 3.9 0 0 1 1.993-1.06q.06-.003.123-.013q.068-.011.134-.03a1.25 1.25 0 0 0 .562-2.096m-5.358-3.548a1.25 1.25 0 0 0 0 1.774c.174.173.386.284.61.336l.075.016l.077.01a3.9 3.9 0 0 1 2.165 1.195a.3.3 0 0 0 .427 0a.3.3 0 0 0 .022-.401l-.022-.025l-.041-.035l-.088-.084a3.85 3.85 0 0 1-1.068-1.976a1.4 1.4 0 0 0-.043-.255a1.27 1.27 0 0 0-2.114-.556m-1.435 4.634l.01-.076a3.86 3.86 0 0 1 1.205-2.148a.3.3 0 0 0 0-.424a.306.306 0 0 0-.428 0q-.02.02-.035.041l-.085.087a3.9 3.9 0 0 1-1.992 1.06a1.4 1.4 0 0 0-.257.043a1.248 1.248 0 0 0-.56 2.097a1.275 1.275 0 0 0 1.79 0a1.26 1.26 0 0 0 .352-.68"/><path fill="currentColor" d="M11.465 3.615a7.865 7.865 0 0 1 7.572 5.738l1.926-.538a9.865 9.865 0 0 0-9.498-7.2c-5.446 0-9.862 4.415-9.862 9.862a9.8 9.8 0 0 0 1.54 5.293l-1.905 4.57h9.022l1.473-.005h.012a9.85 9.85 0 0 0 4.037-.99l-.876-1.797a7.8 7.8 0 0 1-3.205.787l-1.444.005H4.239l1.173-2.815l-.327-.454a7.862 7.862 0 0 1 6.38-12.456"/>`), g.Group(children),
	)
}

func LogoWindows(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 1.847V12H12V3.276zM14 5.01V10h6V4.153zm-3-1.188V12H2V5.303zM4 7.001v3h5V6.177zm-2 6h9V21.1l-9-.857zm2 2v3.423l5 .476V15zm8-2h10v9.156l-10-1.459zm2 2v3.968l6 .875V15z"/>`), g.Group(children),
	)
}

func LogoWindowsFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21.75 2.135v9.615h-9.5V3.492zm-11 1.982v7.633h-8.5V5.515zm-8.5 9.133h8.5v7.575l-8.5-.81zm10 0h9.5v8.617l-9.5-1.385z"/>`), g.Group(children),
	)
}

func LogoYoutube(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.927 2.5h6.146c1.824 0 3.293 0 4.45.155c1.2.162 2.21.507 3.012 1.31c.803.802 1.148 1.813 1.31 3.013C24 8.134 24 9.603 24 11.427v1.146c0 1.824 0 3.293-.155 4.45c-.162 1.2-.507 2.21-1.31 3.012c-.802.803-1.812 1.148-3.013 1.31c-1.156.155-2.625.155-4.449.155H8.927c-1.824 0-3.293 0-4.45-.155c-1.2-.162-2.21-.507-3.013-1.31c-.802-.802-1.147-1.812-1.309-3.013C0 15.866 0 14.397 0 12.573v-1.146c0-1.824 0-3.293.155-4.45c.162-1.2.507-2.21 1.31-3.013c.802-.802 1.813-1.147 3.013-1.309C5.634 2.5 7.103 2.5 8.927 2.5M4.744 4.638c-.978.131-1.496.372-1.865.74c-.37.37-.61.888-.741 1.866C2.002 8.251 2 9.586 2 11.5v1c0 1.914.002 3.249.138 4.256c.131.978.372 1.496.74 1.865c.37.37.888.61 1.866.742c1.007.135 2.342.137 4.256.137h6c1.914 0 3.249-.002 4.256-.137c.978-.132 1.496-.373 1.865-.742c.37-.369.61-.887.742-1.865c.135-1.007.137-2.342.137-4.256v-1c0-1.914-.002-3.249-.137-4.256c-.132-.978-.373-1.496-.742-1.865c-.369-.37-.887-.61-1.865-.741C18.249 4.502 16.914 4.5 15 4.5H9c-1.914 0-3.249.002-4.256.138m4.057 2.326l8.695 5.037l-8.695 5.036zm2 3.47v3.134L13.506 12z"/>`), g.Group(children),
	)
}

func Logout(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m21.207 11.793l-5.914 5.914l-1.414-1.414l3.5-3.5H7.793v-2h9.586l-3.5-3.5l1.414-1.414zm-11.414-7.5h-5v15h5v2h-7v-19h7z"/>`), g.Group(children),
	)
}

func LookAround(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m8-4v4H7V8zm4 0v4h-2V8zm-4 6h2.004v2.004H9z"/>`), g.Group(children),
	)
}

func Loudspeaker(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 1.382V15.5l-7.2-3h-.91L7.712 14H9.5v2H7.477l-.588 5H2V5h6.764zM5.875 12.5H4V19h1.11zM4 10.5h5.2l4.8 2V4.618L9.236 7H4zm16.333-5.914l.707.707c.937.937 1.328 2.342 1.328 3.624c0 1.281-.391 2.687-1.328 3.624l-.707.707l-1.414-1.414l.707-.708c.444-.443.742-1.264.742-2.21c0-.945-.298-1.765-.742-2.209L18.919 6zM18.25 6.255l.707.707a2.77 2.77 0 0 1 0 3.914l-.707.707l-1.415-1.414l.707-.707c.3-.3.3-.786 0-1.086l-.707-.707z"/>`), g.Group(children),
	)
}

func Mail(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v1.83l9 4.55l9-4.55V5zm18 4.07l-9 4.55l-9-4.55V19h18z"/>`), g.Group(children),
	)
}

func Map(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5v15.574l-7 4.084l-6.074-3.544L2 21.5V5.926zm1 15.084l4 2.333V7.074l-4-2.333zM8 4.74L4 7.074V18.5l4-1.667zm8 2.426v12.092l4-2.333V5.5z"/>`), g.Group(children),
	)
}

func MapAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7.886 1.859l8.086 3.537L22 2.382V11h-2V5.618l-3 1.5V11h-2V7.154L9 4.53v10.853l2.34 1.17l-.895 1.789l-2.401-1.201L2 20.766V5.98zM7 15.434V4.92l-3 2.1v10.213zM19 12v4h4v2h-4v4h-2v-4h-4v-2h4v-4z"/>`), g.Group(children),
	)
}

func MapAiming(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m12-7.667V5.52A6.56 6.56 0 0 1 18.48 11h1.187v2H18.48A6.56 6.56 0 0 1 13 18.48v1.187h-2V18.48A6.56 6.56 0 0 1 5.52 13H4.333v-2H5.52A6.56 6.56 0 0 1 11 5.52V4.333zm-1 3.111a4.556 4.556 0 1 0 0 9.112a4.556 4.556 0 0 0 0-9.112m0 4.445a.111.111 0 1 0 0 .222a.111.111 0 0 0 0-.222M9.889 12a2.111 2.111 0 1 1 4.222 0a2.111 2.111 0 0 1-4.222 0"/>`), g.Group(children),
	)
}

func MapBlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5v8h-2v-5l-4 1.667V10.5h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5zM18 13.5a3.5 3.5 0 0 0-3.08 5.165l4.745-4.744A3.5 3.5 0 0 0 18 13.5m3.08 1.835l-4.745 4.744a3.5 3.5 0 0 0 4.745-4.745M12.5 17a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0"/>`), g.Group(children),
	)
}

func MapBubble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5v15.574l-7 4.084l-6.074-3.544L2 21.5V5.926zM4 7.074V18.5l5.074-2.114L15 19.842l5-2.916V5.5l-5.074 2.114L9 4.158zM7.5 9.5a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5m-1.75-.25a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0M14 10.5a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5m-1.75-.25a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0M10 14a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5m-1.75-.25a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0M16.5 15a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5m-1.75-.25a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0"/>`), g.Group(children),
	)
}

func MapCancel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V11h-2V5.5l-4 1.667V11h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5zm7.172-4.076L18 15.586l2.829-2.829l1.414 1.415L19.415 17l2.828 2.828l-1.414 1.415L18 18.414l-2.828 2.829l-1.414-1.415L16.586 17l-2.828-2.828z"/>`), g.Group(children),
	)
}

func MapChat(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V11h-2V5.5l-4 1.667V10h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5zM13 13h10v7.497h-5.292L13 22.517zm2 2v4.483l2.296-.986H21V15z"/>`), g.Group(children),
	)
}

func MapChecked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V12h-2V5.5l-4 1.667V12h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5zm15.657-2.247l-7.07 7.071l-4.243-4.243L13.758 16l2.828 2.829l5.657-5.657z"/>`), g.Group(children),
	)
}

func MapCollection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5v8h-2v-5l-4 1.667V9.5h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5zm10-6.102l2.21 3.226l3.752 1.106l-2.385 3.1l.108 3.909L18 20.762l-3.685 1.31l.108-3.91l-2.385-3.1l3.751-1.105zm0 3.538l-.963 1.405l-1.634.482l1.039 1.35l-.047 1.703L18 18.64l1.605.57l-.047-1.703l1.04-1.35l-1.635-.482z"/>`), g.Group(children),
	)
}

func MapConnection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 4a2 2 0 1 0 0 4a2 2 0 0 0 0-4M2 6a4 4 0 0 1 7.874-1h8.209l.081.014a5.45 5.45 0 0 1 2.32 1.005C21.288 6.637 22 7.614 22 9s-.71 2.363-1.517 2.981a5.45 5.45 0 0 1-2.319 1.005l-.081.014H6c-.237 0-.772.148-1.25.53c-.447.358-.75.842-.75 1.47s.303 1.112.75 1.47c.478.382 1.013.53 1.25.53h11v-2.5l4.667 3.5L17 21.5V19H6c-.763 0-1.728-.352-2.5-.97C2.697 17.389 2 16.373 2 15s.697-2.388 1.5-3.03C4.272 11.351 5.237 11 6 11h11.91c.374-.075.917-.27 1.357-.606c.444-.34.733-.78.733-1.394s-.29-1.053-.733-1.394A3.5 3.5 0 0 0 17.91 7H9.874A4.002 4.002 0 0 1 2 6"/>`), g.Group(children),
	)
}

func MapDistance(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm12.5 3a.5.5 0 0 0-.5.5c0 .222.157.545.488.907l.012.013l.012-.013c.33-.362.488-.685.488-.907a.5.5 0 0 0-.5-.5m0 3.964c-.294-.202-.594-.396-.87-.622a6 6 0 0 1-.618-.585C14.592 9.298 14 8.497 14 7.5a2.5 2.5 0 0 1 5 0c0 .997-.593 1.798-1.012 2.257a6 6 0 0 1-.618.585c-.276.227-.576.42-.87.622M8.5 10a.5.5 0 0 0-.5.5c0 .222.157.545.488.907l.012.013l.012-.013c.33-.362.488-.685.488-.907a.5.5 0 0 0-.5-.5m0 3.964c-.221-.152-.447-.299-.662-.459a6 6 0 0 1-.826-.748C6.592 12.298 6 11.497 6 10.5a2.5 2.5 0 0 1 5 0c0 .997-.593 1.798-1.012 2.257a6 6 0 0 1-.826.748c-.215.16-.441.307-.662.459m11.288-.379l-5.873 2.202l-.702-1.872l5.872-2.202zm-8 3l-5.873 2.203l-.702-1.873l5.872-2.202z"/>`), g.Group(children),
	)
}

func MapDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7.5 1.842l4.574 2.669L17.5 2.25v6.375L22 6.75v12.2l-5.5 3.208l-4.574-2.669L6.5 21.75v-6.375L2 17.25V5.05zm-1 11.366V9.551l5.123-2.989L7.5 4.158L4 6.199v8.051zm5.866-6.652L15.5 8.384V5.25zm5.134 4.236v8.467l2.5-1.458V9.75zm-2 8.467v-8.56L13 9.241v8.56zm-4.5-1.55V9.24l-2.5 1.458v8.051z"/>`), g.Group(children),
	)
}

func MapEdit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V11h-2V5.5l-4 1.667V11h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5zm11.787-4.747l4.127 4.127l-7.286 7.287H12.5l-.001-4.128zm-.922 3.75l1.299 1.3l.922-.923l-1.3-1.299zm-.115 2.713l-1.3-1.299l-2.95 2.95v1.3h1.3z"/>`), g.Group(children),
	)
}

func MapGrid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm5 2v1h2V6h2v1h2V6h2v1h1v2h-1v2h1v2h-1v2h1v2h-1v1h-2v-1h-2v1h-2v-1H9v1H7v-1H6v-2h1v-2H6v-2h1V9H6V7h1V6zm0 3v2h2V9zm4 0v2h2V9zm2 4h-2v2h2zm-4 2v-2H9v2z"/>`), g.Group(children),
	)
}

func MapInformation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4a3 3 0 0 0-3 3c0 1.237.782 2.498 1.738 3.544c.456.498.914.908 1.262 1.195a13 13 0 0 0 1.262-1.195C14.218 9.498 15 8.238 15 7a3 3 0 0 0-3-3m0 10.214l-.567-.39l-.002-.002l-.004-.002l-.012-.009l-.041-.03l-.144-.104a14.6 14.6 0 0 1-1.968-1.784C8.218 10.751 7 9.013 7 7a5 5 0 0 1 10 0c0 2.012-1.218 3.752-2.262 4.893a14.6 14.6 0 0 1-2.112 1.889l-.04.029l-.013.009l-.004.002l-.001.001zm0-6.964a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5M10.25 7a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0M2 10h5v2H4v6.43l4.439-4.843l1.474 1.352L5.273 20h9.023l-2.707-3.094l1.505-1.317l.888 1.015L18.586 12H17v-2h5v12H2zm18 3.414l-4.698 4.698L16.954 20H20z"/>`), g.Group(children),
	)
}

func MapInformationOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4a3 3 0 0 0-3 3c0 1.237.782 2.498 1.738 3.544c.456.498.914.908 1.262 1.195a13 13 0 0 0 1.262-1.195C14.218 9.498 15 8.238 15 7a3 3 0 0 0-3-3m0 10.214l-.567-.39l-.002-.002l-.004-.002l-.012-.009l-.041-.03l-.144-.104a14.6 14.6 0 0 1-1.968-1.784C8.218 10.751 7 9.013 7 7a5 5 0 0 1 10 0c0 2.012-1.218 3.752-2.262 4.893a14.6 14.6 0 0 1-2.112 1.889l-.04.029l-.013.009l-.004.002l-.001.001zm0-6.964a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5M10.25 7a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0M2 10h5v2H4v6.497l5.075-2.116L15 19.842l5-2.916V12h-3v-2h5v8.074l-7 4.084l-6.074-3.548L2 21.498z"/>`), g.Group(children),
	)
}

func MapInformationTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4a3 3 0 0 0-3 3c0 1.237.782 2.498 1.738 3.544c.456.498.914.908 1.262 1.195a13 13 0 0 0 1.262-1.195C14.218 9.498 15 8.238 15 7a3 3 0 0 0-3-3m0 10.214l-.567-.39l-.002-.002l-.004-.002l-.012-.009l-.041-.03l-.144-.104a14.6 14.6 0 0 1-1.968-1.784C8.218 10.751 7 9.013 7 7a5 5 0 0 1 10 0c0 2.012-1.218 3.752-2.262 4.893a14.6 14.6 0 0 1-2.112 1.889l-.04.029l-.013.009l-.004.002l-.001.001zm0-6.964a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5M10.25 7a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0m-7.145 3H7v2H4.895l-.778 7h15.766l-.778-7H17v-2h3.895l1.222 11H1.883z"/>`), g.Group(children),
	)
}

func MapLocation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V11h-2V5.5l-4 1.667V11h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5zM18 14a2.75 2.75 0 0 0-2.75 2.75c0 1.252.735 2.454 1.615 3.422c.407.448.817.814 1.135 1.075c.318-.26.728-.627 1.135-1.075c.88-.968 1.615-2.17 1.615-3.422A2.75 2.75 0 0 0 18 14m0 9.701l-.555-.369l-.002-.001l-.004-.003l-.012-.008l-.04-.028l-.137-.097a13 13 0 0 1-1.865-1.677c-.995-1.094-2.135-2.767-2.135-4.768a4.75 4.75 0 1 1 9.5 0c0 2.001-1.14 3.674-2.135 4.768a13 13 0 0 1-2.002 1.774l-.04.028l-.012.008l-.004.003h-.002zM16.75 16h2.5v2h-2.5z"/>`), g.Group(children),
	)
}

func MapLocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V11h-2V5.5l-4 1.667V11h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5zM18 14c.69 0 1.25.56 1.25 1.25V16h-2.5v-.75c0-.69.56-1.25 1.25-1.25m3.25 2v-.75a3.25 3.25 0 0 0-6.5 0V16h-1.251v6.5h9V16zm-.751 2v2.5h-5V18z"/>`), g.Group(children),
	)
}

func MapMarked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V10h-2V5.5l-4 1.667V9h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5zm5.75-4.837h8.5v10.295l-4.247-2.617l-4.253 2.614zm2 2v4.715l2.254-1.385l2.246 1.383v-4.713z"/>`), g.Group(children),
	)
}

func MapNavigation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 2v3h10V2h2v3h2v2h-2v8h2v2h-2v4h-2v-4h-2v-2h2V7H8v4H6V7H2V5h4V2zm6.58 8.419l-4.375 13.126l-3.008-5.743l-5.743-3.008zm-8.032 4.785l2.13 1.117l1.117 2.13l1.624-4.87z"/>`), g.Group(children),
	)
}

func MapOutline(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5v15.574l-7 4.084l-6.074-3.544L2 21.5V5.926zM4 7.074V18.5l5.074-2.114L15 19.842l5-2.916V5.5l-5.074 2.114L9 4.158z"/>`), g.Group(children),
	)
}

func MapRoutePlanning(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4a3 3 0 0 0-3 3c0 1.237.782 2.498 1.738 3.544c.456.498.914.908 1.262 1.195a13 13 0 0 0 1.262-1.195C14.218 9.498 15 8.238 15 7a3 3 0 0 0-3-3m0 10.214c-.258-.178-.519-.35-.77-.537a14.6 14.6 0 0 1-1.968-1.784C8.218 10.751 7 9.013 7 7a5 5 0 0 1 10 0c0 2.012-1.218 3.752-2.262 4.893a14.6 14.6 0 0 1-1.968 1.784c-.251.187-.512.36-.77.537m0-6.964a.25.25 0 1 0 0-.5a.25.25 0 0 0 0 .5M10.25 7a1.75 1.75 0 1 1 3.5 0a1.75 1.75 0 0 1-3.5 0M2 10h3v2h-.634l.868 1.419L4 14.174V20h16v-6.248l-.783-1.035l.68-.514l.103.137V12h-2v-2h4v12H2zm2 2.194L4.317 12H4zm14.654 3.607l-.879.478q-.663.36-1.366.635l-.728-1.863q.585-.228 1.139-.53l.878-.477zm-11.318.974a12 12 0 0 1-1.346-.677l1.008-1.727q.543.317 1.123.564l.92.392l-.785 1.84zm6.303.897l-.999.05q-.753.037-1.507-.02l.152-1.994q.627.048 1.255.016l1-.05z"/>`), g.Group(children),
	)
}

func MapRuler(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15.857.216l7.812 7.811L7.86 23.835L.05 16.023zm0 2.828l-2.184 2.184l1.602 1.603l-1.414 1.414l-1.602-1.602l-2.184 2.184l2.402 2.402l-1.415 1.414l-2.401-2.402l-2.184 2.184l1.602 1.602l-1.414 1.415l-1.603-1.603l-2.184 2.184l4.983 4.983l12.98-12.979z"/>`), g.Group(children),
	)
}

func MapSafety(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V10h-2V5.5l-4 1.667V9h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5zm5.498-4.835h9v5.633a3 3 0 0 1-1.36 2.512l-3.14 2.052l-3.14-2.052a3 3 0 0 1-1.36-2.512zm2 2v3.633a1 1 0 0 0 .453.837l2.047 1.337l2.047-1.337a1 1 0 0 0 .453-.837v-3.633z"/>`), g.Group(children),
	)
}

func MapSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7.886 1.859l8.086 3.537L22 2.382V10h-2V5.618l-3 1.5V10h-2V7.154L9 4.53v10.853l2.345 1.172l-.895 1.79l-2.406-1.204L2 20.766V5.98zM7 15.434V4.92l-3 2.1v10.213zM17.25 13.5a2.75 2.75 0 0 1 1.946 4.693l-.008.008a2.75 2.75 0 1 1-1.938-4.7m3.992 5.325a4.75 4.75 0 1 0-1.414 1.415l1.67 1.674l1.416-1.412z"/>`), g.Group(children),
	)
}

func MapSearchOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 3c1.15 0 2 .887 2 2.09c0 .809-.523 1.662-1.227 2.405a9 9 0 0 1-.773.72a9 9 0 0 1-.773-.72C7.523 6.752 7 5.899 7 5.09C7 3.884 7.84 3 9 3m3.909 1.2C12.532 2.398 11.026 1 9 1C6.96 1 5.464 2.4 5.09 4.2H2V20h9v-2H4v-3.2l1.957-1.044L9.31 15.88l1.07-1.69l-4.337-2.746L4 12.534V6.2h1.155c.306 1.107 1.014 2.03 1.618 2.67a11 11 0 0 0 1.63 1.4l.033.022l.01.007l.003.002l.002.001l.549.36l.548-.359l.003-.002l.003-.002l.01-.007l.033-.022l.112-.078a11 11 0 0 0 1.517-1.323c.605-.639 1.313-1.562 1.62-2.669H20v4.3h2V4.2zm4.341 9.3a2.75 2.75 0 0 1 1.946 4.693l-.008.008a2.75 2.75 0 1 1-1.938-4.7m3.992 5.325a4.75 4.75 0 1 0-1.414 1.415l1.67 1.674l1.416-1.412z"/>`), g.Group(children),
	)
}

func MapSetting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V11h-2V5.5l-4 1.667V11h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5z"/><path fill="currentColor" d="M19 12.75v1.376c.715.184 1.352.56 1.854 1.072l1.193-.689l1 1.732l-1.192.688a4 4 0 0 1 0 2.142l1.192.688l-1 1.732l-1.193-.689A4 4 0 0 1 19 21.874v1.376h-2v-1.376a4 4 0 0 1-1.854-1.072l-1.193.689l-1-1.732l1.192-.688a4 4 0 0 1 0-2.142l-1.192-.688l1-1.732l1.193.688A4 4 0 0 1 17 14.127V12.75zm-2.751 4.283a2 2 0 0 0-.25.967c0 .35.091.68.25.967l.036.063a2 2 0 0 0 3.43 0l.036-.063A2 2 0 0 0 20 18c0-.35-.09-.68-.249-.967l-.036-.063a2 2 0 0 0-3.43 0z"/>`), g.Group(children),
	)
}

func MapThreeD(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 1.842l7 4.084v5.401c1.03.452 1.927 1.018 2.61 1.687c.828.81 1.39 1.825 1.39 2.986c0 1.963-1.566 3.472-3.482 4.43C17.542 21.418 14.883 22 12 22s-5.543-.582-7.518-1.57C2.566 19.472 1 17.964 1 16c0-1.161.562-2.175 1.39-2.986c.683-.67 1.58-1.235 2.61-1.687V5.926zM5 13.56a5.8 5.8 0 0 0-1.212.884C3.224 14.996 3 15.525 3 16c0 .798.673 1.79 2.376 2.641C7.02 19.463 9.36 20 12 20s4.98-.537 6.624-1.359C20.327 17.79 21 16.798 21 16c0-.475-.224-1.004-.788-1.557A5.8 5.8 0 0 0 19 13.56v.515l-7 4.084l-7-4.084zm12-.633V8.365l-4 2.223v4.671zm-6 2.333v-4.67L7 8.364v4.56zM7.876 6.563L12 8.856l4.124-2.293L12 4.158z"/>`), g.Group(children),
	)
}

func MapUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9 1.842l6.074 3.544L22 2.5V11h-2V5.5l-4 1.667V11h-2V7.074l-4-2.333v12.185l1.868 1.09l-1.008 1.727l-1.934-1.129L2 21.5V5.926zM8 16.833V4.741L4 7.074V18.5zM18 14c-.69 0-1.25.56-1.25 1.25V16h5.749v6.5h-9V16h1.251v-.75a3.25 3.25 0 0 1 5.541-2.305l.71.705l-1.41 1.418l-.71-.705A1.24 1.24 0 0 0 18 14m-2.501 4v2.5h5V18z"/>`), g.Group(children),
	)
}

func MarkAsUnread(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9.5 1.34l7.744 4.555L16.23 7.62L9.5 3.66L2 8.072V20H0V6.928zM4 9h19v14H4zm3.992 2l5.508 3.787L19.008 11zM21 12.057l-7.5 5.157L6 12.057V21h15z"/>`), g.Group(children),
	)
}

func Markup(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-5.61 16.038L8.228 12H10V8.586l2-2l2 2V12h1.773l1.837 7.038A9 9 0 0 0 12 3m3.832 17.146L14.227 14H9.772l-1.604 6.146A9 9 0 0 0 12 21c1.372 0 2.67-.306 3.832-.854M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11a10.99 10.99 0 0 1-5.5 9.528A10.95 10.95 0 0 1 12 23a10.95 10.95 0 0 1-6.013-1.787A10.99 10.99 0 0 1 1 12"/>`), g.Group(children),
	)
}

func Mathematics(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.125 1H22v22H2V1.83l10.125 10zm2 20H20v-1.333h-3.15v-2H20v-1.334h-3.15v-2H20V13h-3.15v-2H20V9.667h-3.15v-2H20V6.333h-3.15v-2H20V3h-5.875zm-2 0v-6.36L4 6.615V21z"/>`), g.Group(children),
	)
}

func Measurement(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 1h12v22H6zm2 2v3h2.5v2H8v3h4v2H8v3h2.5v2H8v3h8V3z"/>`), g.Group(children),
	)
}

func MeasurementOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h10v10h10v10H2zm2 2v2.5h2.004v2H4V11h2.004v2H4v2.5h2.004v2H4V20h2.5v-1.967h2V20H11v-1.967h2V20h2.5v-1.967h2V20H20v-6H10V4z"/>`), g.Group(children),
	)
}

func MeasurementTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.5 1.586L22.914 6L21.5 7.414l-2-2v13.172l2-2L22.914 18L18.5 22.414L14.086 18l1.414-1.414l2 2V5.414l-2 2L14.086 6zM2 2h10v20H2zm2 2v16h6V4z"/>`), g.Group(children),
	)
}

func MeatPepper(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m21.212 5.727l-.97.243c-.999.25-1.848.8-2.456 1.364c.47.167.938.392 1.384.668c1.036.64 2.022 1.591 2.667 2.838c.858 1.657.543 3.303-.339 4.705c-.857 1.363-2.276 2.564-3.82 3.547c-1.56.992-3.323 1.812-4.958 2.385c-1.602.562-3.162.917-4.324.923c-1.294.07-2.364-.255-3.149-.934c-.578-.5-.94-1.139-1.138-1.809c-.67-.198-1.308-.56-1.808-1.138c-.679-.785-1.005-1.854-.934-3.15c.006-1.164.36-2.677.927-4.217a22.6 22.6 0 0 1 2.413-4.749C5.71 4.92 6.93 3.561 8.314 2.708c1.404-.864 3.036-1.236 4.725-.6c2.116.725 3.133 2.17 3.599 3.57c.822-.71 1.89-1.34 3.119-1.648l.97-.243zM17 9.178v1.472h-2V8c0-1.393-.432-3.26-2.621-4.003l-.018-.006l-.018-.007c-.968-.37-1.952-.205-2.98.427C8.31 5.06 7.28 6.165 6.365 7.522a20.6 20.6 0 0 0-2.194 4.32c-.534 1.454-.804 2.726-.804 3.557v.03l-.002.03c-.053.906.176 1.437.449 1.752c.28.325.71.54 1.268.613l.76.1l.1.76c.074.558.289.988.614 1.27c.315.271.846.501 1.752.448l.03-.002h.029c.835 0 2.162-.274 3.69-.81c1.51-.53 3.133-1.286 4.547-2.186c1.429-.909 2.57-1.918 3.2-2.923c.608-.966.71-1.848.257-2.722c-.455-.88-1.167-1.577-1.942-2.056A5.2 5.2 0 0 0 17 9.178"/>`), g.Group(children),
	)
}

func MediaLibrary(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 2h14v2H5zM3 5.5h18v2H3zM1 9h22v13H1zm2 2v9h18v-9zm6.75 1.469L15 15.5l-5.25 3.031z"/>`), g.Group(children),
	)
}

func Member(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.5 7a3 3 0 1 1 6 0a3 3 0 0 1-6 0m3-5a5 5 0 1 0 0 10a5 5 0 0 0 0-10m7 0h-1v2h1a3 3 0 1 1 0 6h-1v2h1a5 5 0 0 0 0-10M0 19a5 5 0 0 1 5-5h7a5 5 0 0 1 5 5v2h-2v-2a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v2H0zm24 0a5 5 0 0 0-5-5h-1v2h1a3 3 0 0 1 3 3v2h2z"/>`), g.Group(children),
	)
}

func Menu(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v5.5h16V4zm16 7.5H4V20h16zM5.996 6H8v2h-.004v.004h-2zM10 6h8v2h-8z"/>`), g.Group(children),
	)
}

func MenuApplication(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 3h3v3H3zm7.5 0h3v3h-3zM18 3h3v3h-3zM3 10.5h3v3H3zm7.5 0h3v3h-3zm7.5 0h3v3h-3zM3 18h3v3H3zm7.5 0h3v3h-3zm7.5 0h3v3h-3z"/>`), g.Group(children),
	)
}

func MenuFold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 4h20v2H2zm0 5.57L5.887 12L2 14.43zM7 11h15v2H7zm-5 7h20v2H2z"/>`), g.Group(children),
	)
}

func MenuUnfold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 4h20v2H2zm20 5.57v4.86L18.113 12zM2 13v-2h15v2zm0 7v-2h20v2z"/>`), g.Group(children),
	)
}

func MergeCells(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h7v-4h2v4h7V4h-7v4h-2V4zm13.182 6.232L15.414 12l1.768 1.768l-1.414 1.414L12.586 12l3.182-3.182zM8.33 8.818L11.512 12L8.33 15.182l-1.414-1.414L8.683 12l-1.767-1.768z"/>`), g.Group(children),
	)
}

func Microphone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 7a6 6 0 0 1 6-6h2a6 6 0 0 1 6 6v6q0 .615-.12 1.199A6 6 0 0 1 13 19v1h5v2H6v-2h5v-1a6 6 0 0 1-6-6zm2 4.5V13h3v2H7.535A4 4 0 0 0 11 17h2a4 4 0 0 0 3.465-2H14v-2h3v-1.5h-3v-2h3V8h-3V6h2.874A4 4 0 0 0 13 3h-2a4 4 0 0 0-3.874 3H10v2H7v1.5h3v2z"/>`), g.Group(children),
	)
}

func MicrophoneOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 6a5 5 0 0 1 10 0v5a5 5 0 0 1-10 0zm5-3a3 3 0 0 0-3 3v5a3 3 0 1 0 6 0V6a3 3 0 0 0-3-3m-6 7v1a6 6 0 0 0 12 0v-1h2v1a8 8 0 0 1-7 7.938V20h5v2H6v-2h5v-1.062A8 8 0 0 1 4 11v-1z"/>`), g.Group(children),
	)
}

func MicrophoneTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16 3a5 5 0 1 0 0 10a5 5 0 0 0 0-10M9.009 7.641a7 7 0 1 1 7.35 7.35l-9.224 7.282a1 1 0 0 1-.115.087a2 2 0 0 1-.295.167a2.3 2.3 0 0 1-1.001.207c-.851-.01-1.836-.433-2.93-1.527s-1.52-2.079-1.528-2.93a2.3 2.3 0 0 1 .207-1.002a2 2 0 0 1 .254-.41zm.4 2.722l-6.122 7.754a.4.4 0 0 0-.021.139c.001.16.077.674.94 1.537c.865.863 1.38.94 1.54.941a.4.4 0 0 0 .137-.02l7.754-6.123a7.02 7.02 0 0 1-4.228-4.228M10.414 15L8 17.414L6.586 16L9 13.586z"/>`), g.Group(children),
	)
}

func Milk(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.85 1H6.15v3.59L4 9.49V23h16V9.49l-2.15-4.9zm-.918 6.476L18 9.91V21h-1.95V9.878zM14.05 10.7V21H6V10.7zm-7.52-2l1.273-2.9h7.614l-1.065 2.9zm1.62-4.9V3h7.7v.8zM13 13h-1.426l-1.576 1.584L8.415 13H7v6h2v-2.586l1.002 1.002l.998-1.003V19h2z"/>`), g.Group(children),
	)
}

func Minus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.502 11h11v2h-11z"/>`), g.Group(children),
	)
}

func MinusCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5.5-1h11v2h-11z"/>`), g.Group(children),
	)
}

func MinusCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 1c6.075 0 11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1m5.5 12v-2h-11v2z"/>`), g.Group(children),
	)
}

func MinusRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm2.5 7h11v2h-11z"/>`), g.Group(children),
	)
}

func MinusRectangleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm4.5 9v2h11v-2z"/>`), g.Group(children),
	)
}

func Mirror(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 2v20h-2V2zM9 4.64V18.5H1.3zm6 0l7.7 13.86H15zM4.7 16.5H7v-4.14zM17 12.36v4.14h2.3z"/>`), g.Group(children),
	)
}

func Mobile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 1h16v22H4zm2 2v18h12V3zm5 14h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func MobileBlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 1h16v4h-2V3H6v18h12v-2h2v4H4zm14 7.5a3.5 3.5 0 0 0-3.08 5.165l4.745-4.744A3.5 3.5 0 0 0 18 8.5m3.08 1.835l-4.745 4.744a3.5 3.5 0 0 0 4.745-4.745M12.5 12a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M11 17h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func MobileList(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 1h16v5h-2V3H6v18h12v-5h2v7H4zm9 7h11v2H13zm0 4h8v2h-8zm-2 5h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func MobileNavigation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 1h16v22H4zm2 2v18h12V3zm6 1.764l4.946 9.892L12 13.11l-4.946 1.546zm-1.054 6.58l1.054-.33l1.054.33L12 9.236zM11 17h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func MobileShortcut(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 23H4V1h16v5h-2V3H6v18h12v-3h2zm-6.996-3.996H11V17h2.004zm8.883-2.717l1.568-.781l-1.568-.78l-.781-1.57l-.781 1.57l-1.57.78l1.57.78l.78 1.57zm-5.83-.986l-1.097-2.204L12.756 12l2.204-1.098l1.097-2.204l1.097 2.204L19.358 12l-2.204 1.097zm5.83-6.024l1.568-.78l-1.568-.782l-.781-1.568l-.781 1.568l-1.57.781l1.57.781l.78 1.569z"/>`), g.Group(children),
	)
}

func MobileVibrate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 1h14v22H5zm2 2v18h10V3zM4 5v14H2V5zm18 0v14h-2V5zM11 17h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func ModeDark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15.844 3.344l-1.428.781l1.428.781l.781 1.428l.781-1.428l1.428-.781l-1.428-.781l-.781-1.428zm-5.432.814A8 8 0 1 0 18.93 16A9 9 0 0 1 10 7c0-.98.131-1.937.412-2.842M2 12C2 6.477 6.477 2 12 2h1.734l-.868 1.5C12.287 4.5 12 5.69 12 7a7 7 0 0 0 8.348 6.87l1.682-.327l-.543 1.626C20.162 19.137 16.417 22 12 22C6.477 22 2 17.523 2 12m18.5-5.584l.914 1.67l1.67.914l-1.67.914l-.914 1.67l-.914-1.67L17.916 9l1.67-.914z"/>`), g.Group(children),
	)
}

func ModeLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.999-.004h2.004V2h-2.004zM4.223 2.803L5.64 4.22L4.223 5.637L2.806 4.22zm15.556 0l1.417 1.417l-1.417 1.417l-1.417-1.417zM12 6a6 6 0 1 0 0 12a6 6 0 0 0 0-12m-8 6a8 8 0 1 1 16 0a8 8 0 0 1-16 0m-4.001-1.004h2.004V13H-.001zm22 0h2.004V13h-2.004zM4.223 18.36l1.417 1.417l-1.417 1.418l-1.417-1.418zm15.556 0l1.417 1.417l-1.417 1.417l-1.417-1.417zM11 21.997h2.004V24H11z"/>`), g.Group(children),
	)
}

func Module(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h7V4zm9 0v7h7V4zm7 9h-7v7h7z"/>`), g.Group(children),
	)
}

func Money(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 12.5a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M10.5 16a1.5 1.5 0 1 1 3 0a1.5 1.5 0 0 1-3 0"/><path fill="currentColor" d="M17.526 5.116L14.347.659L2.658 9.997L2.01 9.99V10H1.5v12h21V10h-.962l-1.914-5.599zM19.425 10H9.397l7.469-2.546l1.522-.487zM15.55 5.79L7.84 8.418l6.106-4.878zM3.5 18.169v-4.34A3 3 0 0 0 5.33 12h13.34a3 3 0 0 0 1.83 1.83v4.34A3 3 0 0 0 18.67 20H5.332A3.01 3.01 0 0 0 3.5 18.169"/>`), g.Group(children),
	)
}

func Monument(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 1v1h6V1h2v2.162l-1 3V15h1v3h2v5H5v-5h2v-3h1V6.162l-1-3V1zm1 6v8h4V7zm4.28-2l.333-1H9.387l.334 1zM9 17v1h6v-1zm8 3H7v1h10z"/>`), g.Group(children),
	)
}

func Moon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.412 4.158A8 8 0 1 0 18.93 16A9 9 0 0 1 10 7c0-.98.131-1.937.412-2.842M2 12C2 6.477 6.477 2 12 2h1.734l-.868 1.5C12.287 4.5 12 5.689 12 7a7 7 0 0 0 8.348 6.87l1.682-.327l-.543 1.626C20.162 19.137 16.417 22 12 22C6.477 22 2 17.523 2 12"/>`), g.Group(children),
	)
}

func MoonFall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10.916 1.585l-.45 1.675c-.3 1.116-.27 2.337.07 3.604a7 7 0 0 0 9.84 4.476l1.541-.752l-.104 1.71c-.062 1.034-.374 2.06-.919 3.15l-.447.894l-1.789-.895l.448-.894q.198-.398.338-.76A9 9 0 0 1 8.604 7.38a9.6 9.6 0 0 1-.338-2.852a8 8 0 0 0-4.164 9.235c.016.061.085.251.173.48l.106.267l.033.082l.009.022l.003.007l.378.926l-1.852.756l-.378-.926l-.004-.01l-.01-.024l-.036-.09l-.113-.288a8 8 0 0 1-.24-.684C.74 8.947 3.906 3.464 9.24 2.034zM1 18h8.303L12 19.798L14.697 18H23v2h-7.697L12 22.202L8.697 20H1z"/>`), g.Group(children),
	)
}

func MoonRising(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 20h8.303L12 18.202L14.697 20H23v-2h-7.697L12 15.798L8.697 18H1zm19.447-3.658l.447-.895c.545-1.088.857-2.115.92-3.148l.103-1.711l-1.54.752a7 7 0 0 1-9.841-4.476c-.34-1.267-.37-2.488-.07-3.604l.45-1.675l-1.675.45c-5.334 1.429-8.5 6.912-7.07 12.247c.05.187.163.486.24.684l.113.288l.036.09l.01.024l.382.936l1.852-.756l-.381-.933l-.009-.022l-.033-.082l-.106-.268a7 7 0 0 1-.173-.479a8 8 0 0 1 4.164-9.235a9.6 9.6 0 0 0 .337 2.852a9 9 0 0 0 10.84 6.411a9 9 0 0 1-.337.76l-.448.895z"/>`), g.Group(children),
	)
}

func More(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.5 3h3v3h-3zm0 7.5h3v3h-3zm0 7.5h3v3h-3z"/>`), g.Group(children),
	)
}

func Mosque(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m13.004 2.002l-.002 1.081A6.005 6.005 0 0 1 17.917 8H22v2h-1v12H3V10H2V8h4.083a6.005 6.005 0 0 1 4.919-4.917l.002-1.085zM8.126 8h7.748a4.002 4.002 0 0 0-7.748 0M5 10v10h3v-6h8v6h3V10zm9 10v-4h-4v4z"/>`), g.Group(children),
	)
}

func MosqueOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6 1.586l4 4v2.31q.185-.316.412-.62c.785-1.05 1.957-1.938 3.588-2.198V3.5h2v1.578c1.63.26 2.802 1.148 3.588 2.198c.658.88 1.05 1.873 1.252 2.724H22v12H2V5.586zM18.764 10a5.4 5.4 0 0 0-.777-1.526C17.367 7.644 16.437 7 15 7s-2.366.645-2.987 1.474A5.4 5.4 0 0 0 11.236 10zM10 12v8h2v-6h6v6h2v-8zm6 8v-4h-2v4zm-8 0V6.414l-2-2l-2 2V20zM7.004 8v2.004H5v-2h.004V8z"/>`), g.Group(children),
	)
}

func Mouse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 9a8 8 0 1 1 16 0v6a8 8 0 1 1-16 0zm7-5.917A6 6 0 0 0 6 9h5zm2 0V9h5a6 6 0 0 0-5-5.917M18 11H6v4a6 6 0 0 0 12 0z"/>`), g.Group(children),
	)
}

func Move(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.75 2.25h3v3h-3zm7.5 0h3v3h-3zm-7.5 5.5h3v3h-3zm7.5 0h3v3h-3zm-7.5 5.5h3v3h-3zm7.5 0h3v3h-3zm-7.5 5.5h3v3h-3zm7.5 0h3v3h-3z"/>`), g.Group(children),
	)
}

func MoveOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.086 12L5.5 7.586L6.914 9l-2 2H11V4.914l-2 2L7.586 5.5L12 1.086L16.414 5.5L15 6.914l-2-2V11h6.086l-2-2L18.5 7.586L22.914 12L18.5 16.414L17.086 15l2-2H13v6.086l2-2l1.414 1.414L12 22.914L7.586 18.5L9 17.086l2 2V13H4.914l2 2L5.5 16.414z"/>`), g.Group(children),
	)
}

func MovieClapper(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v4h4.865L5.532 4zm4.135 0l3.333 4h4.397l-3.333-4zm7 0l3.333 4H20V4zM20 10H4v10h16zm-5 4H9v-2h6z"/>`), g.Group(children),
	)
}

func Multiply(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.818 7.404L12 10.586l3.182-3.182l1.414 1.414L13.414 12l3.182 3.182l-1.415 1.414L12 13.414l-3.182 3.182l-1.415-1.414L10.585 12L7.403 8.818z"/>`), g.Group(children),
	)
}

func Museum(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 2h12v4.327l4-.444V22H2V8.105l4-.444zm2 5.438l8-.889V4H8zM18 20h2v-5h-2zm2-7V8.117L4 9.895V20h12v-7zM6 10.998h2.004v2.004H6zm4 0h2.004v2.004H10z"/>`), g.Group(children),
	)
}

func MuseumOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 4.765V20h2v-5h2v5h2V4.766A6.2 6.2 0 0 0 12 4c-1.089 0-2.11.277-3 .765m8 1.616v3.72A6.98 6.98 0 0 1 22 8h1v14H1V8h1c1.959 0 3.73.804 5 2.101v-3.72l-.033.03l-1.379-1.448l.724-.69q.526-.5 1.136-.904A8.2 8.2 0 0 1 12 2a8.22 8.22 0 0 1 5.69 2.276l.724.69l-1.38 1.448zM7 15a5 5 0 0 0-4-4.9V20h4zm10 5h4v-9.9a5 5 0 0 0-4 4.9zM11 6.998h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func MuseumTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6 1.798l4 2.667v6.355l12-2V22h-6v-4h-2v4H2V4.465zM4 5.535V20h8v-4h6v4h2v-8.82l-12 2V5.536L6 4.202z"/>`), g.Group(children),
	)
}

func Mushroom(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.308 3.102C12.987 4.779 14.63 6 16.499 6c1.157 0 2.198-.405 2.985-1.091a5 5 0 0 0-.191-.202c-1.626-1.626-4.38-2.158-6.985-1.605m8.228 3.53A6.52 6.52 0 0 1 16.499 8c-2.79 0-5.161-1.826-6.096-4.278c-.842.38-1.593.882-2.196 1.485c-.381.381-.55.72-.6 1.035c-.051.317 0 .736.287 1.31c.097.193.34.52.737.955c.379.415.84.869 1.31 1.306l.014.012l4.22 4.22l.012.013q.107.116.222.236C15.337 11.822 17.704 10 20.5 10q.303 0 .599.027c.037-1.197-.15-2.372-.562-3.394m.294 5.38A5 5 0 0 0 20.5 12c-2.211 0-4.089 1.693-4.448 3.85c.161.125.282.199.395.255l.021.011c.566.283.979.331 1.292.28c.314-.052.652-.223 1.033-.603c.979-.98 1.68-2.304 2.038-3.781m-8.761 2.756l-2.837-2.837a45 45 0 0 1-.708.624c-.83.715-1.937 1.614-2.984 2.287l-.013.008l-.013.007l-.104.063c-.554.33-1.34.8-1.73 1.472a1.4 1.4 0 0 0-.173 1.05c.094.433.4 1.05 1.2 1.85c.8.801 1.418 1.107 1.851 1.2c.407.088.741.008 1.05-.172c.672-.39 1.142-1.176 1.472-1.73l.062-.105l.008-.012l.008-.013c.672-1.047 1.572-2.154 2.287-2.984a45 45 0 0 1 .624-.708m-4.27-4.24a20 20 0 0 1-.647-.673c-.412-.452-.82-.951-1.047-1.408c-.42-.84-.607-1.683-.475-2.518c.133-.837.567-1.541 1.162-2.136c1.086-1.086 2.48-1.871 3.968-2.347c3.323-1.064 7.432-.668 9.946 1.847c.348.348.656.73.923 1.135c1.308 1.98 1.689 4.551 1.36 6.917c-.307 2.21-1.244 4.323-2.783 5.862c-.59.591-1.291 1.026-2.123 1.162c-.832.137-1.674-.046-2.51-.464l-.021-.01c-.42-.21-.781-.486-1.22-.88a23 23 0 0 1-.86-.818c-.154.174-.328.37-.513.585c-.692.803-1.515 1.82-2.11 2.745l-.061.104c-.298.51-1.022 1.75-2.177 2.42a3.4 3.4 0 0 1-2.474.397c-.933-.2-1.878-.774-2.845-1.741s-1.54-1.912-1.74-2.845a3.4 3.4 0 0 1 .397-2.474c.67-1.155 1.91-1.88 2.42-2.177l.103-.061c.924-.595 1.942-1.418 2.745-2.11c.213-.184.41-.357.582-.511"/>`), g.Group(children),
	)
}

func MushroomOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.543 5.455c-4.052-4.054-7.822-3.49-9.015-2.298c-.276.276-.395.647-.166 1.36c.244.758.841 1.713 1.761 2.838c.737.902 1.642 1.868 2.648 2.873c1.005 1.006 1.97 1.91 2.872 2.648c1.126.92 2.08 1.516 2.84 1.76c.711.23 1.083.11 1.358-.165c1.193-1.193 1.755-4.962-2.298-9.016M8.113 1.743C10.437-.58 15.526-.393 19.959 4.04c4.432 4.434 4.62 9.522 2.298 11.844c-.974.973-2.218 1.031-3.387.655c-.869-.28-1.778-.817-2.691-1.492q-.224.328-.468.742a49 49 0 0 0-.95 1.723l-.207.388c-.796 1.484-1.735 3.182-2.905 4.349c-.958.956-2.186 1.803-3.787 1.747c-1.57-.055-3.214-.969-5.051-2.81C.973 19.349.06 17.704.003 16.137c-.058-1.602.789-2.828 1.747-3.786c1.167-1.168 2.865-2.106 4.348-2.903l.384-.205a49 49 0 0 0 1.726-.953q.414-.244.743-.468c-.675-.913-1.213-1.823-1.493-2.692c-.376-1.168-.318-2.413.656-3.386m2.103 7.635c-.313.22-.649.431-.991.633c-.564.333-1.19.668-1.81 1l-.37.198c-1.53.822-2.957 1.63-3.881 2.555c-.8.8-1.19 1.515-1.163 2.3c.03.82.525 2.009 2.224 3.71c1.698 1.702 2.886 2.196 3.706 2.224c.786.028 1.503-.365 2.305-1.164c.925-.923 1.733-2.349 2.554-3.878l.2-.375c.332-.62.666-1.244.998-1.806c.202-.343.412-.68.633-.993a48 48 0 0 1-2.265-2.14a48 48 0 0 1-2.14-2.264"/>`), g.Group(children),
	)
}

func Music(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20 2.913V18a3 3 0 1 1-2-2.83v-5.083L8 10.92V19a3 3 0 1 1-2-2.83V4.08zM6 19a1 1 0 1 0-2 0a1 1 0 0 0 2 0M8 8.913l10-.833V5.087L8 5.92zM18 18a1 1 0 1 0-2 0a1 1 0 0 0 2 0"/>`), g.Group(children),
	)
}

func MusicOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 .847V17a4 4 0 1 1-2-3.465V3.153L8 4.867V19a4 4 0 1 1-2-3.465V3.133zM6 19a2 2 0 1 0-4 0a2 2 0 0 0 4 0m14-2a2 2 0 1 0-4 0a2 2 0 0 0 4 0"/>`), g.Group(children),
	)
}

func MusicRectangleAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v11h-2V4H4v16h9v2H2zm10 5h4v2h-2v5a3 3 0 1 1-2-2.83zm0 7a1 1 0 1 0-2 0a1 1 0 0 0 2 0m8 1v3h3v2h-3v3h-2v-3h-3v-2h3v-3z"/>`), g.Group(children),
	)
}

func MusicTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 3h7v2h-5v13a4 4 0 1 1-2-3.465zm0 15a2 2 0 1 0-4 0a2 2 0 0 0 4 0"/>`), g.Group(children),
	)
}

func NavigationArrow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 1.764l9.946 19.892L12 18.548l-9.946 3.108zm0 4.472L5.946 18.344L12 16.452l6.054 1.892z"/>`), g.Group(children),
	)
}

func Next(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.75 4.336L14.414 12L6.75 19.664zM17.5 5v14h-2V5zM8.75 9.164v5.672L11.586 12z"/>`), g.Group(children),
	)
}

func NoExpression(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m8-4v4H7V8zm8 0v4h-2V8z"/>`), g.Group(children),
	)
}

func Noodle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 3h7v2h-1v5h1v2c0 5.523-4.477 10-10 10S2 17.523 2 12v-2h1a6 6 0 1 1 12 0h1V5h-1zm3 2v5h1V5zM4 12a8 8 0 1 0 16 0zm1-2h1a3 3 0 0 1 6 0h1a4 4 0 0 0-8 0m5 0a1 1 0 0 0-2 0z"/>`), g.Group(children),
	)
}

func Notification(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 8a8 8 0 1 1 16 0v4.697l2 3V20h-5.611a4.502 4.502 0 0 1-8.777 0H2v-4.303l2-3zm5.708 12a2.5 2.5 0 0 0 4.584 0zM12 2a6 6 0 0 0-6 6v5.303l-2 3V18h16v-1.697l-2-3V8a6 6 0 0 0-6-6"/>`), g.Group(children),
	)
}

func NotificationAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 8a8 8 0 1 1 16 0v4.697l2 3V20h-5.611a4.502 4.502 0 0 1-8.777 0H2v-4.303l2-3zm5.708 12a2.5 2.5 0 0 0 4.584 0zM12 2a6 6 0 0 0-6 6v5.303l-2 3V18h16v-1.697l-2-3V8a6 6 0 0 0-6-6m1 4v3h3v2h-3v3h-2v-3H8V9h3V6z"/>`), g.Group(children),
	)
}

func NotificationCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-9.268 5a2 2 0 0 1-3.464 0H6v-2.736l1-2V10a5 5 0 0 1 10 0v2.264l1 2V17zM8 15h8v-.264l-1-2V10a3 3 0 1 0-6 0v2.736l-1 2z"/>`), g.Group(children),
	)
}

func NotificationError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 8a8 8 0 1 1 16 0v4.697l2 3V20h-5.611a4.502 4.502 0 0 1-8.777 0H2v-4.303l2-3zm16 10v-1.697l-2-3V8A6 6 0 0 0 6 8v5.303l-2 3V18zm-5.708 2H9.708a2.5 2.5 0 0 0 4.584 0M13 6v6h-2V6zm-2 7.5h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func NotificationFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.645 20.5a3.502 3.502 0 0 0 6.71 0zM3 19.5h18v-3l-2-3v-5a7 7 0 1 0-14 0v5l-2 3z"/>`), g.Group(children),
	)
}

func NumbersEight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 6a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2zM9 6h6v5H9zm0 7h6v5H9z"/>`), g.Group(children),
	)
}

func NumbersEightOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 5.5a2.25 2.25 0 1 0 0 4.5a2.25 2.25 0 0 0 0-4.5m2.857 5.396a5 5 0 1 1-5.714 0a4.25 4.25 0 1 1 5.714 0M12 12a3 3 0 1 0 0 6a3 3 0 0 0 0-6"/>`), g.Group(children),
	)
}

func NumbersFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 4H7v7a2 2 0 0 0 2 2h6v5H7v2h8a2 2 0 0 0 2-2v-5a2 2 0 0 0-2-2H9V6h8z"/>`), g.Group(children),
	)
}

func NumbersFiveOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 4H16v2H9.5v4H12a5 5 0 0 1 4.9 4h.1v1a5 5 0 0 1-5 5h-.5a5 5 0 0 1-5-5h2a3 3 0 0 0 3 3h.5a3 3 0 1 0 0-6H7.5z"/>`), g.Group(children),
	)
}

func NumbersFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.5 4v8h7V4h2v16h-2v-6h-7a2 2 0 0 1-2-2V4z"/>`), g.Group(children),
	)
}

func NumbersFourOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.989 4H15.5v10h2v2h-2v4h-2v-4H6v-2.323zm.511 10V6.708L8.234 14z"/>`), g.Group(children),
	)
}

func NumbersNine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 6a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v5a2 2 0 0 0 2 2h6v5H7.5v2H15a2 2 0 0 0 2-2zm-2 5H9V6h6z"/>`), g.Group(children),
	)
}

func NumbersNineOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.582 20.5H9.273l3.473-6.052a5.25 5.25 0 1 1 3.573-2.209zm3.402-9.953a3.25 3.25 0 1 0-.354.617z"/>`), g.Group(children),
	)
}

func NumbersOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 18h-2V6a2 2 0 0 0-2-2H9v2h2v12H9v2h6z"/>`), g.Group(children),
	)
}

func NumbersOneOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 20V3.5h-1.834L8.1 5.8l1.2 1.6L11 6.125V20z"/>`), g.Group(children),
	)
}

func NumbersSeven(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 4h10.5v3.113a2 2 0 0 1-.52 1.346l-4.48 4.928V20h-2v-6.613a2 2 0 0 1 .52-1.346l4.48-4.928V6H7z"/>`), g.Group(children),
	)
}

func NumbersSevenOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 4h10v2.203L11.663 20H9.488l5.926-14H7.5z"/>`), g.Group(children),
	)
}

func NumbersSix(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.5 4H9a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2v-5a2 2 0 0 0-2-2H9V6h7.5zM9 13h6v5H9z"/>`), g.Group(children),
	)
}

func NumbersSixOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.42 3.5h2.31l-3.473 6.052Q11.621 9.5 12 9.5a5.25 5.25 0 1 1-4.316 2.261zm-3.4 9.953a3.25 3.25 0 1 0 .354-.617z"/>`), g.Group(children),
	)
}

func NumbersThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 4h8a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H7v-2h8v-5H8v-2h7V6H7z"/>`), g.Group(children),
	)
}

func NumbersThreeOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 4h9v3.02l-4.264 2.986A5 5 0 0 1 12 20h-.5a5 5 0 0 1-5-5h2a3 3 0 0 0 3 3h.5a3 3 0 1 0 0-6H9.5V9.48L14.47 6H7.5z"/>`), g.Group(children),
	)
}

func NumbersTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 4h8a2 2 0 0 1 2 2v5a2 2 0 0 1-2 2H9v5h8v2H7v-7a2 2 0 0 1 2-2h6V6H7z"/>`), g.Group(children),
	)
}

func NumbersTwoOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 8a5 5 0 0 1 5-5h1v.1c2.282.463 4 2.481 4 4.9v.065c0 1.525-.687 2.97-1.871 3.931L9 16.976V18h8v2H7v-3.024a2 2 0 0 1 .739-1.552l6.129-4.98A3.07 3.07 0 0 0 15 8.065V8a3 3 0 1 0-6 0v1H7z"/>`), g.Group(children),
	)
}

func NumbersZero(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 6a2 2 0 0 0-2-2H9a2 2 0 0 0-2 2v12a2 2 0 0 0 2 2h6a2 2 0 0 0 2-2zM9 6h6v12H9z"/>`), g.Group(children),
	)
}

func NumbersZeroOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 9a5 5 0 0 0-5-5h-1v.1A5 5 0 0 0 7 9v6a5 5 0 0 0 10 0zm-5-3a3 3 0 0 1 3 3v6a3 3 0 1 1-6 0V9a3 3 0 0 1 3-3"/>`), g.Group(children),
	)
}

func Nut(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.474 4.805c.413-.623.83-1.38 1.077-2.121l1.898.632c-.36 1.078-.97 2.122-1.52 2.906c1.067 1.222 1.953 2.874 2.343 4.59c.468 2.067.243 4.4-1.494 6.138l-2.121 2.121l-.726-.726c-1.898 1.704-4.095 2.964-6.264 3.488c-2.468.596-4.998.25-6.874-1.626s-2.222-4.406-1.626-6.874c.524-2.169 1.784-4.366 3.488-6.264l-.726-.726L7.05 4.222c1.727-1.727 4.06-1.95 6.11-1.5a10.7 10.7 0 0 1 4.314 2.083M16.217 6.36a8.7 8.7 0 0 0-3.486-1.685c-1.646-.362-3.185-.121-4.266.96l-.708.707l5.364 5.364h3.415v3.414l1.12 1.121l.708-.707c1.092-1.091 1.331-2.629.957-4.28c-.29-1.281-.934-2.527-1.695-3.474c-.207.214-.424.431-.653.66l-.023.023l-1.414-1.414c.257-.257.48-.481.681-.69m-.703 10.567l-.978-.978v-2.243h-2.243L7.072 8.486c-1.49 1.681-2.536 3.56-2.96 5.317c-.484 2-.153 3.742 1.095 4.99s2.99 1.58 4.99 1.096c1.757-.425 3.636-1.47 5.317-2.96"/>`), g.Group(children),
	)
}

func ObjectStorage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4a5 5 0 0 0-5 5v2H6c-.78 0-1.805.302-2.62.964C2.598 12.6 2 13.572 2 15c0 1.414.615 2.388 1.42 3.032c.836.67 1.866.968 2.58.968c.973 0 2.239-.015 3.263-.521c.485-.24.9-.585 1.203-1.097c.306-.517.534-1.27.534-2.382v-1.086l-1.5 1.5L8.086 14L12 10.086L15.914 14L14.5 15.414l-1.5-1.5V15c0 1.387-.286 2.51-.812 3.4a4.8 4.8 0 0 1-2.04 1.872C8.674 21 6.979 21 6.056 21H6c-1.168 0-2.638-.452-3.83-1.407C.943 18.613 0 17.086 0 15c0-2.072.903-3.6 2.12-4.589A6.45 6.45 0 0 1 5 9.093V9a7 7 0 0 1 14 0v.093a6.45 6.45 0 0 1 2.88 1.318C23.099 11.401 24 12.928 24 15c0 2.076-.907 3.614-2.205 4.607C20.529 20.575 18.93 21 17.5 21h-4v-2h4c1.07 0 2.221-.325 3.08-.982C21.407 17.386 22 16.424 22 15c0-1.428-.598-2.4-1.38-3.036C19.804 11.302 18.78 11 18 11h-1V9a5 5 0 0 0-5-5"/>`), g.Group(children),
	)
}

func OpenMouth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m8-4v4H7V8zm8 0v4h-2V8zM7 13h10v5H7zm2 2v1h6v-1z"/>`), g.Group(children),
	)
}

func Opera(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.498 5h1.5c3.66 0 6.065 1.464 7.578 3.51c1.01 1.365 1.593 2.947 1.933 4.43a10.3 10.3 0 0 1 2.169-1.238a12.4 12.4 0 0 1 1.943-.642c.528-.123 1.067-.21 1.6-.312l-1.376 8.253L2.217 19l-1.5-6h1.28c1.022 0 1.899.105 2.596.236l-.921-3.225h1.326c1.296 0 2.554.231 3.503.467a38 38 0 0 0-1.426-4.094zm2.587 7.723c-.24-.078-.559-.174-.932-.27a15 15 0 0 0-1.804-.35L7.751 17h2.125q-.035-.26-.081-.573a47 47 0 0 0-.71-3.705M11.893 17h4.086a24 24 0 0 0-.124-1.884c-.21-1.79-.718-3.838-1.887-5.418c-.927-1.253-2.31-2.27-4.464-2.592a44 44 0 0 1 1.398 4.613a49 49 0 0 1 .99 5.281m6.087 0h2.17l.596-3.575q-.138.054-.284.116c-.894.381-1.872.953-2.578 1.763c.052.568.08 1.164.094 1.642zM5.67 17l-.445-1.555a9 9 0 0 0-.5-.136a11 11 0 0 0-1.43-.242L3.778 17z"/>`), g.Group(children),
	)
}

func OrderAdjustmentColumn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 3v8h5.586l-2-2L18 7.586L22.414 12L18 16.414L16.586 15l2-2H13v8h-2v-8H5.414l2 2L6 16.414L1.586 12L6 7.586L7.414 9l-2 2H11V3z"/>`), g.Group(children),
	)
}

func OrderAscending(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 2.586L23.414 7L22 8.414l-2-2V20h-2V6.414l-2 2L14.586 7zM2 4h11v2H2zm0 7h13v2H2zm0 7h13v2H2z"/>`), g.Group(children),
	)
}

func OrderDescending(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 4h13v2H2zm0 7h13v2H2zm1 7H2v2h11v-2zm16 3.414l.707-.707l3-3l.707-.707L22 15.586l-.707.707L20 17.586V4h-2v13.586l-1.293-1.293l-.707-.707L14.586 17l.707.707l3 3z"/>`), g.Group(children),
	)
}

func Outbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 2H2v20h20zm-2 14.5V20H4v-3.5h3.67A5 5 0 0 0 12 19a5 5 0 0 0 4.33-2.5zm-16-2V4h16v10.5h-4.965l-.253.625a3.001 3.001 0 0 1-5.564 0l-.253-.625zm8-9.414L7.586 9.5L9 10.914l2-2V14h2V8.914l2 2L16.414 9.5z"/>`), g.Group(children),
	)
}

func PageFirst(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.5 5v14h-2V5zm9.164 1.75L12.414 12l5.25 5.25l-1.414 1.414L9.586 12l6.664-6.664z"/>`), g.Group(children),
	)
}

func PageHead(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21 10v12h-2V12H5v10H3V10zm0-8v6H3V2zm-2 2H5v2h14z"/>`), g.Group(children),
	)
}

func PageLast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.5 5v14h-2V5zm-9.75.336L14.414 12L7.75 18.664L6.336 17.25l5.25-5.25l-5.25-5.25z"/>`), g.Group(children),
	)
}

func Palace(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 2v1.126c1.725.444 3 2.01 3 3.874h3v2h-1v1h4v2h-1v8h1v2H2v-2h1v-8H2v-2h4V9H5V7h3a4 4 0 0 1 3-3.874V2zM8 9v1h8V9zm6-2a2 2 0 1 0-4 0zm-9 5v8h3v-2a4 4 0 0 1 8 0v2h3v-8zm9 8v-2a2 2 0 1 0-4 0v2z"/>`), g.Group(children),
	)
}

func PalaceFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6 1.698l4 3.334V9h4V5.032l4-3.334l4 3.334V22H2V5.032zM16 9h4V5.968l-2-1.666l-2 1.666zm4 2H4v9h5v-3a3 3 0 1 1 6 0v3h5zm-7 9v-3a1 1 0 1 0-2 0v3zM4 9h4V5.968L6 4.302L4 5.968z"/>`), g.Group(children),
	)
}

func PalaceOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m13.004 2.002l-.002 1.081A6 6 0 0 1 18 9h4v13H2V9h4a6 6 0 0 1 5.002-5.917l.002-1.085zM8 9h8a4 4 0 0 0-8 0m2.998-3.002h2.004v2.004h-2.004zM4 11v9h1v-3a3 3 0 1 1 6 0v3h2v-3a3 3 0 1 1 6 0v3h1v-9zm13 9v-3a1 1 0 1 0-2 0v3zm-8 0v-3a1 1 0 1 0-2 0v3z"/>`), g.Group(children),
	)
}

func PalaceThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m13.004 2.002l-.002 1.081A6 6 0 0 1 18 9h4v13H2V9h4a6 6 0 0 1 5.002-5.917l.002-1.085zM12 5a4 4 0 0 0-4 4v2H4v9h4v-1.646a6.43 6.43 0 0 1 3.553-5.748l.447-.224l.447.224A6.43 6.43 0 0 1 16 18.354V20h4v-9h-4V9a4 4 0 0 0-4-4m2 15v-1.646a4.43 4.43 0 0 0-2-3.702a4.43 4.43 0 0 0-2 3.702V20zM10.998 7.998h2.004v2.004h-2.004z"/>`), g.Group(children),
	)
}

func PalaceTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m13.004 2.002l-.002 1.081A6 6 0 0 1 18 9h2V3h2v19H2V3h2v6h2a6 6 0 0 1 5.002-5.917l.002-1.085zM8 9h8a4 4 0 0 0-8 0m-4 2v9h4v-1.646a6.43 6.43 0 0 1 3.553-5.748l.447-.224l.447.224A6.43 6.43 0 0 1 16 18.354V20h4v-9zm10 9v-1.646a4.43 4.43 0 0 0-2-3.702a4.43 4.43 0 0 0-2 3.702V20zM10.998 5.998h2.004v2.004h-2.004z"/>`), g.Group(children),
	)
}

func Palette(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9c0 5.004 3.715 8.897 8.653 8.897h.003q-.071-.118-.156-.252l-.02-.032c-.4-.637-.71-1.466-.653-2.375c.059-.942.505-1.856 1.414-2.616c1.032-.863 2.263-.884 3.218-.847q.322.014.608.03c.758.041 1.358.074 1.947-.085c1.852-.5 2.881-1.73 2.988-3.323C21.323 6.62 17.115 3 12 3M1 12C1 5.925 5.925 1 12 1c5.931 0 11.418 4.286 10.998 10.53c-.173 2.566-1.908 4.431-4.463 5.12c-.932.252-1.932.193-2.708.148q-.24-.016-.447-.024c-.917-.036-1.45.043-1.855.381c-.527.441-.68.86-.702 1.208c-.024.38.106.796.35 1.185l.027.043c.096.153.236.374.352.605c.12.239.278.604.29 1.013a1.5 1.5 0 0 1-.7 1.338c-.425.275-.959.35-1.49.35C5.546 22.897 1 18.042 1 12m8.75-4.996a2 2 0 1 1 4 0a2 2 0 0 1-4 0m-4.72 3a2 2 0 1 1 4 0a2 2 0 0 1-4 0m9.488 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func PaletteOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4c-1.584 0-2.889 1.303-2.889 2.938c0 1.634 1.305 2.937 2.889 2.937s2.889-1.303 2.889-2.937C14.889 5.303 13.584 4 12 4M7.111 6.938C7.111 4.222 9.288 2 12 2s4.889 2.222 4.889 4.938s-2.177 4.937-4.889 4.937s-4.889-2.222-4.889-4.937M5.89 14.125c-1.585 0-2.89 1.303-2.89 2.938C3 18.697 4.305 20 5.889 20s2.889-1.303 2.889-2.937c0-1.635-1.305-2.938-2.89-2.938M1 17.063c0-2.716 2.177-4.938 4.889-4.938s4.889 2.222 4.889 4.938S8.6 22 5.888 22C3.179 22 1 19.778 1 17.063m17.111-2.938c-1.584 0-2.889 1.303-2.889 2.938c0 1.634 1.305 2.937 2.89 2.937C19.694 20 21 18.697 21 17.063c0-1.635-1.305-2.938-2.889-2.938m-4.889 2.938c0-2.716 2.177-4.938 4.89-4.938c2.71 0 4.888 2.222 4.888 4.938S20.823 22 18.111 22s-4.889-2.222-4.889-4.937"/>`), g.Group(children),
	)
}

func PanoramaHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23 3.156v17.688l-1.287-.386h-.002l-.009-.003l-.038-.011a23 23 0 0 0-.772-.21a42 42 0 0 0-2.213-.503C16.839 19.362 14.409 19 12 19s-4.838.363-6.679.73a42 42 0 0 0-2.827.668l-.158.046l-.038.01l-.01.004L1 20.844V3.156l1.287.386h.002l.009.003l.038.012l.158.045a42 42 0 0 0 2.827.668C7.161 4.637 9.591 5 12 5s4.838-.362 6.679-.73a42 42 0 0 0 2.827-.668l.158-.045l.038-.012l.01-.002zm-2 2.645c-.519.128-1.175.279-1.929.43C17.161 6.613 14.591 7 12 7s-5.162-.387-7.071-.77c-.754-.15-1.41-.3-1.929-.428V18.2a45 45 0 0 1 1.929-.43C6.839 17.389 9.409 17 12 17s5.162.388 7.071.77c.754.15 1.41.301 1.929.43z"/>`), g.Group(children),
	)
}

func PanoramaVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.156 1h17.688l-.386 1.287v.002l-.003.009l-.011.038l-.045.158c-.04.14-.096.348-.165.614a42 42 0 0 0-.503 2.213C19.363 7.161 19 9.591 19 12s.362 4.838.73 6.679a42 42 0 0 0 .668 2.827l.045.158l.011.038l.003.01l.386 1.288H3.156l.386-1.287v-.002l.004-.009l.01-.038l.046-.158a42 42 0 0 0 .668-2.827c.368-1.84.73-4.27.73-6.679s-.362-4.838-.73-6.679a42 42 0 0 0-.668-2.827l-.045-.158l-.011-.038l-.003-.01zm2.645 2c.128.519.279 1.175.43 1.929C6.613 6.839 7 9.409 7 12s-.388 5.162-.77 7.071A45 45 0 0 1 5.8 21h12.4a45 45 0 0 1-.43-1.929c-.381-1.91-.77-4.48-.77-7.071s.388-5.162.77-7.071c.15-.754.302-1.41.43-1.929z"/>`), g.Group(children),
	)
}

func Pantone(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m14.347.991l3.178 4.455l2.099-.715l1.915 5.602H23V22H1V10.333h1.652zM3.353 12.333l-.003.003H3V20h18v-7.667zm16.072-2L18.388 7.3l-1.604.513l-7.397 2.521zM15.549 6.12l-1.604-2.248l-6.102 4.874zM5.285 15.164H7.29v2.004H5.285z"/>`), g.Group(children),
	)
}

func Parabola(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.083 7.912c-.229.44-.47.904-.683 1.188c-.89 1.187-1.605 2.118-2.223 2.75C4.583 12.456 3.881 13 3 13H2v-2h1c.019 0 .217-.007.748-.55C4.255 9.933 4.89 9.114 5.8 7.9c.106-.141.252-.42.495-.885l.137-.261c.298-.568.677-1.27 1.145-1.947C8.48 3.5 9.903 2 12 2s3.52 1.499 4.423 2.807c.468.678.847 1.379 1.145 1.947l.137.26c.243.466.39.745.495.886c.91 1.213 1.544 2.032 2.052 2.55c.53.543.729.55.748.55h1v2h-1c-.881 0-1.583-.543-2.177-1.15c-.617-.632-1.333-1.563-2.223-2.75c-.213-.284-.454-.748-.683-1.188l-.12-.228c-.292-.557-.625-1.169-1.02-1.741C13.955 4.75 13.053 4 12 4s-1.955.751-2.777 1.943c-.395.572-.728 1.184-1.02 1.74zM4 16v2h7v-2h2v2h7v-2h2v5h-2v-1h-7v1h-2v-1H4v1H2v-5z"/>`), g.Group(children),
	)
}

func Parentheses(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m5.445 4.539l-.438.899A14.9 14.9 0 0 0 3.5 12c0 2.355.542 4.581 1.507 6.562l.438.899l-1.798.876l-.438-.899A16.9 16.9 0 0 1 1.5 12c0-2.665.614-5.19 1.71-7.438l.437-.899zm14.907-.876l.439.899A16.9 16.9 0 0 1 22.5 12c0 2.665-.614 5.19-1.71 7.438l-.437.899l-1.798-.876l.438-.9A14.9 14.9 0 0 0 20.5 12c0-2.355-.542-4.581-1.507-6.562l-.438-.899z"/>`), g.Group(children),
	)
}

func Paste(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h5v2H4v3H2zm9 0h5v5h-2V4h-3zM9 9h13v13H9zm2 2v9h9v-9zm-7-1v3h3v2H2v-5z"/>`), g.Group(children),
	)
}

func Patio(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.783 2.06C7.25 1.527 9.365 1 12 1s4.75.527 6.217 1.06a14 14 0 0 1 1.699.74a9 9 0 0 1 .62.356l.011.007l.005.003l.001.001l.002.001l.445.297V20h1v2H2v-2h1V3.465l.445-.297l.002-.001l.001-.001l.005-.003l.011-.007a3 3 0 0 1 .163-.101q.161-.098.457-.254c.395-.206.966-.474 1.7-.74M5 4.58V20h2V7.5h2v2h2v-2h2V20h6V4.579l-.01-.005c-.324-.17-.815-.4-1.457-.634C16.25 3.473 14.365 3 12 3s-4.25.473-5.533.94A12 12 0 0 0 5 4.579M9 11.5V13h2v-1.5zm2 3.5H9v1.5h2zm0 3.5H9V20h2z"/>`), g.Group(children),
	)
}

func Pause(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 17V7H8v10zm5 0V7h-3v10z"/>`), g.Group(children),
	)
}

func PauseCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12"/><path fill="currentColor" d="M13 7h3v10h-3zM8 7h3v10H8z"/>`), g.Group(children),
	)
}

func PauseCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11M8 7h3v10H8zm5 0h3v10h-3z"/>`), g.Group(children),
	)
}

func PauseCircleStroke(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m9.5-5v10h-2V7zm5 0v10h-2V7z"/>`), g.Group(children),
	)
}

func Pea(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m22.874 3.008l.662 1.836L21 6.452V8c0 3.146-.54 5.83-1.799 7.954q-1.91 3.227-5.878 4.466c-.913.285-2.423.562-3.432.578c-3.572.058-6.596-1.548-8.148-3.344l-.758-.878l.98-.62l4.44-2.816l-.035-.016c-.884-.407-1.758-1.002-2.528-1.686s-1.465-1.48-1.974-2.31C1.368 8.51 1 7.585 1 6.663v-.881zM8.681 11.896L13 9.156l.842-.537l.002.002l4.834-3.065L3.188 7.52q.138.355.385.762c.39.635.948 1.285 1.597 1.86c.648.576 1.359 1.054 2.036 1.365c.585.27 1.086.388 1.475.39m4.551-.518l-2.829 1.794l1.457 2.166l-1.66 1.116l-1.486-2.21l-4.594 2.913c1.317 1.026 3.358 1.88 5.739 1.842c.786-.013 2.105-.25 2.868-.488c2.229-.696 3.759-1.897 4.753-3.576C18.49 13.229 19 10.94 19 8v-.28l-4.079 2.587l1.45 2.272l-1.685 1.076z"/>`), g.Group(children),
	)
}

func Peach(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 1h1c1.196 0 2.05.632 2.616 1.364C13.08 1.644 13.829 1 15 1h1v2h-1c-.361 0-.676.225-.94.928a4 4 0 0 0-.127.402a6.4 6.4 0 0 1 .998.025c1.09.109 2.356.49 3.49 1.262c2.009 1.366 5.005 5.074 2.817 10.533c-1.31 3.269-4.344 5.063-7.815 6.005a11 11 0 0 1-.881.672l-.438.297l-.492-.194c-5.837-2.302-8.604-5.612-9.38-8.91c-.764-3.244.46-6.28 2.215-7.872c2.161-1.96 4.83-2.316 6.963-1.915q-.067-.15-.148-.298C10.908 3.301 10.485 3 10 3H9zm3.318 5.591c-1.685-.757-4.444-.851-6.527 1.038c-1.226 1.112-2.207 3.409-1.613 5.933c.564 2.393 2.6 5.182 7.662 7.302q.227-.177.444-.368a14 14 0 0 0 1.66-1.757a15 15 0 0 0 .596-.794l.029-.042l.006-.01l.014-.02l.015-.02c.797-1.085 1.484-3.33 1.29-5.647c-.19-2.262-1.197-4.422-3.576-5.615m4.069 12.199c1.394-.851 2.424-1.962 2.994-3.384c1.732-4.32-.586-7.116-2.084-8.135a5.6 5.6 0 0 0-1.782-.79c1.518 1.58 2.209 3.604 2.372 5.559c.215 2.571-.464 5.175-1.5 6.75"/>`), g.Group(children),
	)
}

func Pear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m21.323 2.534l-.447.895A14.3 14.3 0 0 1 19.65 5.45c.878 1.101 1.358 2.354 1.313 3.695c-.053 1.604-.852 3.156-2.257 4.561l-.055.055l-.059.043l-.005.004l-.04.032a4.7 4.7 0 0 0-.743.799c-.417.567-.807 1.364-.805 2.358c.004 1.686-.29 3.115-1.498 4.323c-2.398 2.399-6.307 2.368-8.674 0l-.068-.068c-.67-.669-1.34-1.338-1.637-2.376c-1.038-.298-1.707-.967-2.376-1.637l-.068-.068c-2.368-2.368-2.399-6.276 0-8.675C3.887 7.29 5.316 6.997 7.002 7c.994.002 1.79-.388 2.358-.804a4.7 4.7 0 0 0 .799-.744l.032-.04l.004-.005l.043-.06l.055-.054c1.385-1.386 2.89-2.229 4.458-2.329c1.241-.079 2.402.315 3.445 1.081c.338-.49.627-.984.891-1.51l.447-.895zm-4.38 3.073c-.705-.5-1.397-.69-2.065-.647c-.905.058-1.962.553-3.126 1.703l-.046.057a6.7 6.7 0 0 1-1.163 1.088C9.745 8.394 8.541 9.003 6.998 9c-1.44-.003-2.236.243-2.905.912c-1.611 1.611-1.593 4.253 0 5.846c.89.89 1.242 1.201 1.975 1.267a1 1 0 0 1 .907.907c.066.733.377 1.085 1.267 1.975c1.593 1.593 4.235 1.611 5.846 0c.67-.67.915-1.466.912-2.905c-.003-1.543.606-2.747 1.192-3.545a6.7 6.7 0 0 1 1.088-1.163l.057-.046c1.13-1.146 1.597-2.227 1.628-3.169c.023-.668-.171-1.354-.625-2.036q-.357.389-.763.794l-.707.707l-1.414-1.414l.707-.707q.422-.422.78-.816m1.651 8.197"/>`), g.Group(children),
	)
}

func PearlOfTheOrient(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.997.999L13 3.17a3.001 3.001 0 0 1 0 5.66v4.296a4.002 4.002 0 0 1 1.86 6.67L16.754 23h-2.325l-1.28-2.168a4 4 0 0 1-2.298 0L9.573 23H7.249l1.89-3.204A3.9 3.9 0 0 1 8 17a4 4 0 0 1 3-3.874V8.829a3.001 3.001 0 0 1 0-5.658L10.997 1zm-.997 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2m0 10a2 2 0 0 0-2 2c0 .518.165.945.502 1.327q.2.225.457.382c.303.185.658.292 1.04.292a1.99 1.99 0 0 0 1.526-.706A2 2 0 0 0 12 15"/>`), g.Group(children),
	)
}

func Pen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m16.362 2.595l-1.671 3.05a9 9 0 0 1-4.785 4.122l-3.829 1.41l-.737 6.066L11 11.582l1.414 1.414l-5.66 5.66l6.063-.737l1.418-3.884a9 9 0 0 1 4.027-4.75l3.052-1.724l.984 1.741l-3.052 1.725a7 7 0 0 0-3.132 3.694l-1.838 5.036l-11.426 1.39L4.239 9.72l4.976-1.83a7 7 0 0 0 3.722-3.206l1.671-3.05z"/>`), g.Group(children),
	)
}

func PenBall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.155 1.374l5.471 5.472l-1.414 1.414l-5.471-5.472zm-2.802 2.833l5.441 5.44L7.397 22.002H2v-5.397zm.002 2.83L4 17.43v2.571h2.57L16.964 9.645zm7.853 5.83l-6.244 6.243l-1.414-1.414l6.244-6.244z"/>`), g.Group(children),
	)
}

func PenBrush(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.999 22h3.623a3 3 0 0 0 2.12-.878l14.79-14.789l-4.866-4.865L2.878 16.256a3 3 0 0 0-.879 2.122zm2-2v-1.622a1 1 0 0 1 .293-.707l2.158-2.158l2.037 2.036l-2.158 2.158a1 1 0 0 1-.707.293zm5.902-3.865l-2.037-2.037l9.802-9.801l2.037 2.036z"/>`), g.Group(children),
	)
}

func PenMark(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23.555 6.88L17.328.651L3.979 14.002l-.565 2.826l3.965 3.966l2.827-.566zm-2.829 0L9.22 18.385l-1.184.236l-2.451-2.451l.236-1.184L17.328 3.481zM1.386 19.612l3.208 3.208l1.414-1.414L2.8 18.2z"/>`), g.Group(children),
	)
}

func PenQuill(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m23.302 8.118l-11.14 11.094l-5.416-.697l-3.673 3.673l-1.414-1.414l3.669-3.67l-.742-5.41L15.672.565l1.816 5.787zm-7.745-1.242l-.803-2.557l-8.052 8.085l.401 2.926zm-7.025 9.853l2.914.375l8.076-8.045l-2.546-.773z"/>`), g.Group(children),
	)
}

func Pending(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-9.996 1.004H11V11h2.004zm-5 0H6V11h2.004zm10 0H16V11h2.004z"/>`), g.Group(children),
	)
}

func Percent(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 5a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3M2.5 6.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0M20.414 5L5 20.414L3.586 19L19 3.586zM18 16a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m-3.5 1.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0"/>`), g.Group(children),
	)
}

func PersonalInformation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V23H3zm2 2v18h14V7.414L14.586 3zm7 6a1.5 1.5 0 1 0 0 3a1.5 1.5 0 0 0 0-3m-3.5 1.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0M6 19a4 4 0 0 1 4-4h4a4 4 0 0 1 4 4v1h-2v-1a2 2 0 0 0-2-2h-4a2 2 0 0 0-2 2v1H6z"/>`), g.Group(children),
	)
}

func PhoneLocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 1h16v10h-2V3H6v18h5v2H4zm13.5 13.5c.69 0 1.25.56 1.25 1.25v.75h-2.5v-.75c0-.69.56-1.25 1.25-1.25m3.25 2v-.75a3.25 3.25 0 0 0-6.5 0v.75h-1.251V23h9v-6.5zm-.751 2V21h-5v-2.5z"/>`), g.Group(children),
	)
}

func PhoneSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 1h16v10h-2V3H6v18h6.5v2H4zm13.25 13.5a2.75 2.75 0 1 1 0 5.5a2.75 2.75 0 0 1 0-5.5m3.992 5.325a4.75 4.75 0 1 0-1.414 1.415l1.67 1.674l1.416-1.412z"/>`), g.Group(children),
	)
}

func Pi(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.161 3h18.258v2h-5.61a111 111 0 0 0-.414 7.741c-.024 1.49-.006 2.85.07 3.9q.058.791.15 1.264c.055.284.102.383.1.385c.64.81 2.544 1.31 4.583-.699l.712-.702l1.404 1.425l-.713.702c-2.59 2.552-5.988 2.575-7.594.465c-.262-.345-.383-.821-.456-1.194c-.082-.424-.14-.938-.181-1.503c-.082-1.133-.1-2.56-.075-4.075A113 113 0 0 1 13.801 5h-3.466c-.163 2.463-.615 5.692-1.52 8.542c-.512 1.619-1.187 3.167-2.076 4.39c-.889 1.225-2.052 2.203-3.546 2.497l-.981.193l-.387-1.962l.981-.193c.829-.163 1.601-.726 2.315-1.71c.714-.983 1.31-2.31 1.788-3.82C7.737 10.33 8.168 7.33 8.331 5H3.16zm12.553 15.29l-.015-.02z"/>`), g.Group(children),
	)
}

func Piano(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v4h18V5zm18 6h-2v5h-2v-5h-2v5h-2v-5h-2v5H9v-5H7v5H5v-5H3v8h18z"/>`), g.Group(children),
	)
}

func Pin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.787.693l5.523 5.522l-6.365 7.774l2.188 2.188l-5.659 5.659l-4.95-4.95L2.16 23.25L.746 21.836l6.364-6.364l-4.95-4.95l5.66-5.659l2.188 2.189zm-4.313 18.314l2.83-2.83l-2.054-2.054l6.365-7.774l-2.963-2.962l-7.779 6.358L7.82 7.692l-2.83 2.83z"/>`), g.Group(children),
	)
}

func PinFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m18.076.981l4.949 4.95l-6.365 7.773l2.121 2.12l-5.305 5.306l-4.596-4.596l-6.718 6.718l-1.414-1.415l6.718-6.717l-4.597-4.596l5.306-5.306l2.121 2.122z"/>`), g.Group(children),
	)
}

func Play(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18.25 12L8.5 17.629V6.37z"/>`), g.Group(children),
	)
}

func PlayCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12"/><path fill="currentColor" d="M18.25 12L8.5 17.63V6.37z"/>`), g.Group(children),
	)
}

func PlayCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 23c6.075 0 11-4.925 11-11S18.075 1 12 1S1 5.925 1 12s4.925 11 11 11m-3.5-5.37V6.37L18.25 12z"/>`), g.Group(children),
	)
}

func PlayCircleStroke(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m7.5-5.63L18.25 12L8.5 17.63zm2 3.465v4.33L14.25 12z"/>`), g.Group(children),
	)
}

func PlayCircleStrokeAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0 0 18h1v2h-1C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11v1h-2v-1a9 9 0 0 0-9-9M9.5 7.131L16.803 12L9.5 16.869zm2 3.737v2.264L13.197 12zM20 15v3h3v2h-3v3h-2v-3h-3v-2h3v-3z"/>`), g.Group(children),
	)
}

func PlayDemo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v2h-1v14h-6.586l4 4L17 23.414l-5-5l-5 5L5.586 22l4-4H3V4H2zm3 2v12h14V4zm5 2.5l4.667 3.5L10 13.5z"/>`), g.Group(children),
	)
}

func PlayRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v14h18V5zm5 1.37L17.75 12L8 17.63zm2 3.465v4.33L13.75 12z"/>`), g.Group(children),
	)
}

func Plus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 6.5V11h4.5v2H13v4.5h-2V13H6.5v-2H11V6.5z"/>`), g.Group(children),
	)
}

func Popsicle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.151 2.95a7 7 0 0 1 9.9 9.9l-8.84 8.839l-2.828-2.829l-3.182 3.182a3 3 0 1 1-4.242-4.243l3.182-3.182l-2.829-2.828zm1.06 15.91l7.042-7.042a4.3 4.3 0 0 0-.431-.997c-.451-.757-1.287-1.619-2.915-1.94c-2.261-.447-3.547-1.706-4.245-2.879a6 6 0 0 1-.257-.477l-6.264 6.264zm.746-14.849c.079.265.21.61.423.968c.451.757 1.287 1.618 2.915 1.94c2.261.446 3.546 1.706 4.245 2.878q.056.095.107.188a5.002 5.002 0 0 0-7.69-5.974m-6.402 12.02l-3.182 3.183a1 1 0 1 0 1.414 1.414l3.182-3.182z"/>`), g.Group(children),
	)
}

func Portrait(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h7v2H4v5H2zm13 0h7v7h-2V4h-5zm-3 6.5a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3m2.665 3.769a3.5 3.5 0 1 0-5.33 0A5 5 0 0 0 7 16.5v1h2v-1a3 3 0 1 1 6 0v1h2v-1a5 5 0 0 0-2.335-4.231M4 15v5h5v2H2v-7zm18 0v7h-7v-2h5v-5z"/>`), g.Group(children),
	)
}

func Pout(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5.769-3.866l3.464 2l-1 1.732l-3.464-2zm11.464 1.732l-3.464 2l-1-1.732l3.464-2zM15 14v3.133l-6.116-.765l.248-1.984l3.868.483V14z"/>`), g.Group(children),
	)
}

func Poweroff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 2v10h-2V2zm3.78 1.728l.809.589a9.5 9.5 0 1 1-11.178 0l.808-.59l1.178 1.617l-.808.59a7.5 7.5 0 1 0 8.822 0l-.808-.59z"/>`), g.Group(children),
	)
}

func PreciseMonitor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 2v6h-2V2zm-9 .586L9.914 8.5L8.5 9.914L2.586 4zM21.414 4L15.5 9.914L14.086 8.5L20 2.586zm-11.146 7A2 2 0 0 1 14 12a2 2 0 0 1-3.732 1H2v-2zM16 11h6v2h-6zm-6.086 4.5L4 21.414L2.586 20L8.5 14.086zm5.586-1.414L21.414 20L20 21.414L14.086 15.5zM13 16v6h-2v-6z"/>`), g.Group(children),
	)
}

func Previous(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.25 4.336v15.328L9.586 12zM8.5 5v14h-2V5zm3.914 7l2.836 2.836V9.164z"/>`), g.Group(children),
	)
}

func Print(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2h16v5h3.5v11H19v4H5v-4H.5V7H4zm2 5h12V4H6zM2.5 9v7H5v-2h14v2h2.5V9zM17 16H7v4h10zm0-5.504h2.004V12.5H17z"/>`), g.Group(children),
	)
}

func Pumpkin(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.277 1.113l.555.832c.468.702.748 1.405.916 2.09a8.2 8.2 0 0 1 3.565 1.233c.708-.247 1.595-.297 2.434-.124c1.178.244 2.424.958 3.148 2.41c.464.93.6 2.24.563 3.55a18.7 18.7 0 0 1-.648 4.246c-.378 1.393-.924 2.752-1.633 3.787c-.686 1-1.685 1.922-3.008 1.922H15.28c-.23.218-.522.416-.902.572c-.595.244-1.364.369-2.375.369s-1.78-.125-2.375-.369a2.9 2.9 0 0 1-.902-.572H6.369c-1.268 0-2.179-1.008-2.765-1.964c-.63-1.028-1.127-2.377-1.475-3.764a20 20 0 0 1-.594-4.232c-.03-1.306.105-2.615.57-3.545c.724-1.453 1.972-2.166 3.152-2.41c.84-.173 1.729-.123 2.438.124a8.3 8.3 0 0 1 2.996-1.157a4.5 4.5 0 0 0-.523-1.056l-.555-.832zm-.286 4.975c-1.123.197-1.942.682-2.409.989c-.705 1.793-.69 4.433-.286 6.966c.207 1.3.508 2.516.828 3.493c.334 1.016.651 1.657.844 1.899l.027.034l.024.036c.053.08.135.18.369.276c.266.109.75.219 1.616.219s1.35-.11 1.616-.22c.234-.095.315-.195.369-.275l.023-.036l.028-.034c.193-.242.51-.883.843-1.9c.32-.976.621-2.193.829-3.492c.404-2.533.419-5.173-.287-6.966c-.47-.31-1.297-.799-2.432-.993q.009.436.007.845V8h-2V7c0-.317 0-.62-.01-.912m6.536.96c.649 2.216.535 4.956.16 7.31a25 25 0 0 1-.903 3.8q-.153.471-.323.9h.708c.32 0 .797-.232 1.358-1.051c.538-.785 1.008-1.912 1.353-3.18a16.7 16.7 0 0 0 .578-3.78c.035-1.216-.113-2.12-.353-2.6c-.403-.808-1.073-1.201-1.763-1.344a3 3 0 0 0-.815-.054M6.48 7.05a3 3 0 0 0-.82.054c-.693.143-1.365.537-1.767 1.343c-.24.482-.388 1.386-.36 2.606c.029 1.181.218 2.527.536 3.793c.319 1.275.753 2.412 1.239 3.204c.529.863.92 1.01 1.06 1.01h1.178a16 16 0 0 1-.323-.9a25 25 0 0 1-.903-3.801c-.376-2.354-.489-5.094.16-7.309"/>`), g.Group(children),
	)
}

func Pyramid(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11 2.96l4.44 7.893l1.06-1.928L23.69 22H.29zM13.227 11L11 7.04L8.772 11zm-5.58 2L3.71 20H10v-7zM12 13v4.107L14.259 13zm.691 7H15.5v-3h-1.16zm4.809 0h2.809l-1.65-3h-1.16zm.059-5l-1.06-1.925L15.442 15z"/>`), g.Group(children),
	)
}

func PyramidMaya(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 2h12v2h-1v4h2v3h1v3h1v3h1v6H2v-6h1v-3h1v-3h1V8h2V4H6zm3 2v4h5.999L15 4zm4 6h-2v11h2zm2 11h5v-2h-1v-3h-1v-3h-1v-3h-2zm-6 0V10H7v3H6v3H5v3H4v2zm2-16.002h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func Qrcode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h9v9H2zm2 2v5h5V4zm9-2h9v9h-9zm2 2v5h5V4zM5.5 5.5h2.004v2.004H5.5zm11 0h2.004v2.004H16.5zm-3.504 7.496H15V15h-2.004zm7 0H22V15h-2.004zM2 13h9v9H2zm2 2v5h5v-5zm11.996.996H18v2h2v2h2V22h-2.004v-2h-2v-2h-2zM5.5 16.5h2.004v2.004H5.5zm7.496 3.496H15V22h-2.004z"/>`), g.Group(children),
	)
}

func Quadratic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.253 3H23v2H12.747L7.882 21.418l-4.215-6.775H1v-2h3.778l2.451 3.94zM22 10v2.136a2 2 0 0 1-.726 1.542L19.07 15.5l2.204 1.822A2 2 0 0 1 22 18.864V21h-2v-2.136l-2.5-2.067l-2.5 2.067V21h-2v-2.136a2 2 0 0 1 .726-1.542L15.93 15.5l-2.204-1.822A2 2 0 0 1 13 12.136V10h2v2.136l2.5 2.067l2.5-2.067V10z"/>`), g.Group(children),
	)
}

func Questionnaire(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 2h21v16H6.876L1.5 22.704zm2 2v14.296L6.124 16H20.5V4zM12 7.5a1 1 0 0 0-1 1v1H9v-1a3 3 0 1 1 6 0c0 .676-.172 1.246-.474 1.71a2.96 2.96 0 0 1-1.029.95a4 4 0 0 1-.494.238v.352h-2v-1c0-.424.245-.687.361-.79c.12-.105.24-.165.296-.192c.107-.05.233-.094.309-.12l.018-.007c.19-.066.36-.127.52-.218a.96.96 0 0 0 .343-.305c.072-.11.15-.294.15-.618a1 1 0 0 0-1-1m-1 5h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func Queue(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m22 2.001l.003 18.418L20 18.415V4.001l-12.999.001v-2zm-20 4h16v16H2zm2 2v12h12V8zm7 1.5V13h3.5v2H11v3.5H9V15H5.5v-2H9V9.5z"/>`), g.Group(children),
	)
}

func Radar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 3.055A9.001 9.001 0 0 0 12 21a9 9 0 0 0 9-9c0-1.98-.638-3.808-1.72-5.293l-.59-.809l1.617-1.177l.59.808A10.96 10.96 0 0 1 23 12c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1h1v9.268A2.01 2.01 0 0 1 14 12a2 2 0 1 1-3-1.732V7.612a4.502 4.502 0 0 0 1 8.888a4.5 4.5 0 0 0 3.64-7.146l-.589-.809l1.617-1.177l.589.808A6.5 6.5 0 1 1 11 5.576z"/>`), g.Group(children),
	)
}

func RadioOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.342 2.447L6.236 6H23v16h-1v-1v1H1V6.382L12.447.658zM21 20V8H3v12zM9 12a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0m10-3h4v2h-4zm0 4h4v2h-4z"/>`), g.Group(children),
	)
}

func RadioTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v16h-4v3h-2v-3H7v3H5v-3H1zm2 2v12h18V5zm12 4a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 2a4 4 0 1 1 8 0a4 4 0 0 1-8 0M5 8h4v2H5zm0 4h4v2H5z"/>`), g.Group(children),
	)
}

func Radish(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.39 1.954v4.243L20 2.586L21.415 4l-3.611 3.611h4.242v2h-4.5c.611.815 1.333 1.931 1.551 2.982c.64 2.6-.633 5.871-3.934 7.913c-2.744 1.698-6.857 2.548-12.508 1.675l-.723-.112l-.112-.724c-.47-3.042-.446-5.631-.06-7.803c.666-3.755 2.416-6.26 4.514-7.612C7.92 4.869 9.81 4.51 11.407 4.903c1.051.218 2.168.94 2.982 1.552v-4.5zm-1.708 7.364l-.012-.012a11 11 0 0 0-.355-.326a18 18 0 0 0-.974-.81c-.828-.638-1.717-1.184-2.35-1.31l-.024-.005l-.023-.006c-.694-.174-1.562-.129-2.455.197l2.426 2.466l-1.426 1.402L6.71 8.09c-1.227 1.035-2.336 2.739-2.883 5.3l3.088 3.117l-1.421 1.407L3.5 15.901c-.073 1.304-.02 2.77.192 4.407c4.127.536 7.16.042 9.284-.905l-2.89-2.907l1.419-1.41l3.256 3.275c2.174-1.643 2.76-3.83 2.39-5.304l-.005-.024l-.005-.023c-.127-.634-.672-1.522-1.31-2.35a18 18 0 0 0-1.137-1.33z"/>`), g.Group(children),
	)
}

func RainHeavy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.1 4h-.064C8.6 4.033 6.7 5.935 6.7 8.2q0 .394.073.765l.184.956l-.95.21C4.832 10.391 4 11.391 4 12.533c0 .947.566 1.79 1.432 2.205l.902.431l-.864 1.804l-.902-.431C3.063 15.822 2 14.309 2 12.533c0-1.83 1.125-3.375 2.706-4.07A6 6 0 0 1 4.7 8.2c0-3.43 2.851-6.152 6.309-6.2h.091c.546 0 1.078.066 1.586.192c2.147.53 3.88 2.12 4.533 4.187a5.6 5.6 0 0 1 2.654 1.074C21.158 8.422 22 9.943 22 11.667c0 2.21-1.382 4.082-3.313 4.894l-.922.388l-.767-1.824V18h-2v-5.998h2v3.1l.914-.384C19.16 14.193 20 13.01 20 11.667c0-1.051-.511-1.999-1.331-2.616a3.6 3.6 0 0 0-2.228-.717l-.854.013l-.147-.842c-.282-1.62-1.55-2.956-3.234-3.371A4.6 4.6 0 0 0 11.1 4m1.898 6.002V20h-2v-9.998zm-4 2V18h-2v-5.998z"/>`), g.Group(children),
	)
}

func RainLight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.1 4h-.064C8.6 4.033 6.7 5.935 6.7 8.2q0 .394.073.765l.184.956l-.95.21C4.832 10.391 4 11.391 4 12.533c0 .947.566 1.79 1.432 2.205l.902.431l-.864 1.804l-.902-.431C3.063 15.822 2 14.309 2 12.533c0-1.83 1.125-3.375 2.706-4.07A6 6 0 0 1 4.7 8.2c0-3.43 2.851-6.152 6.309-6.2h.091c.546 0 1.078.066 1.586.192c2.147.53 3.88 2.12 4.533 4.187a5.6 5.6 0 0 1 2.654 1.074C21.158 8.422 22 9.943 22 11.667c0 2.21-1.382 4.082-3.313 4.894l-.922.388l-.775-1.843l.922-.388C19.16 14.193 20 13.01 20 11.667c0-1.051-.511-1.999-1.331-2.616a3.6 3.6 0 0 0-2.228-.717l-.854.013l-.147-.842c-.282-1.62-1.55-2.956-3.234-3.371A4.6 4.6 0 0 0 11.1 4m-.102 5.998h2.004v2.004h-2.004zm-4 2h2.004v2.004H6.998zm8 0h2.004v2.004h-2.004zm-4 2h2.004v2.004h-2.004zm-4 2h2.004v2.004H6.998zm8 0h2.004v2.004h-2.004zm-4 2h2.004v2.004h-2.004z"/>`), g.Group(children),
	)
}

func RainMedium(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.1 4h-.064C8.6 4.033 6.7 5.935 6.7 8.2q0 .394.073.765l.184.956l-.95.21C4.832 10.391 4 11.391 4 12.533c0 .947.566 1.79 1.432 2.205l.902.431l-.864 1.804l-.902-.431C3.063 15.822 2 14.309 2 12.533c0-1.83 1.125-3.375 2.706-4.07A6 6 0 0 1 4.7 8.2c0-3.43 2.851-6.152 6.309-6.2h.091c.546 0 1.078.066 1.586.192c2.147.53 3.88 2.12 4.533 4.187a5.6 5.6 0 0 1 2.654 1.074C21.158 8.422 22 9.943 22 11.667c0 2.21-1.382 4.082-3.313 4.894l-.922.388l-.775-1.843l.922-.388C19.16 14.193 20 13.01 20 11.667c0-1.051-.511-1.999-1.331-2.616a3.6 3.6 0 0 0-2.228-.717l-.854.013l-.147-.842c-.282-1.62-1.55-2.956-3.234-3.371A4.6 4.6 0 0 0 11.1 4m1.898 6.002v2.999h-2v-2.999zm-4 2V15h-2v-2.998zm8 0V15h-2v-2.998zm-4 1.999V20h-2v-6zm-6 1.997h2.004v2.004H6.998zm8 0h2.004v2.004h-2.004z"/>`), g.Group(children),
	)
}

func Rainbow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 8a9 9 0 0 0-9 9v1H1v-1C1 10.925 5.925 6 12 6s11 4.925 11 11v1h-2v-1a9 9 0 0 0-9-9m0 3a6 6 0 0 0-6 6v1H4v-1a8 8 0 1 1 16 0v1h-2v-1a6 6 0 0 0-6-6m0 3a3 3 0 0 0-3 3v1H7v-1a5 5 0 0 1 10 0v1h-2v-1a3 3 0 0 0-3-3"/>`), g.Group(children),
	)
}

func Rectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4z"/>`), g.Group(children),
	)
}

func Refresh(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.5 5.835A10.49 10.49 0 0 0 12 1.5c-5.427 0-9.89 4.115-10.443 9.396l-.104.994l1.99.209l.103-.995A8.501 8.501 0 0 1 19.213 7.5H15.5v2h7v-7h-2zm.057 6.066l-.104.995A8.501 8.501 0 0 1 4.787 16.5H8.5v-2h-7v7h2v-3.335A10.49 10.49 0 0 0 12 22.5c5.426 0 9.89-4.115 10.442-9.396l.104-.994z"/>`), g.Group(children),
	)
}

func Relation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10.239 1.633L12 4.684l1.761-3.05l1.732 1l-2.338 4.05L19.11 17H23v2h-2.735l1.128 1.954l-1.732 1L17.956 19H6.044L4.34 21.954l-1.732-1L3.735 19H1v-2h3.89l5.955-10.316l-2.338-4.05zM12 8.684L7.199 17H16.8z"/>`), g.Group(children),
	)
}

func Relativity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h13v7h7v13H9v-7H2zm9 13v5h9v-9h-5v4zm2-6V4H4v9h5V9zm0 2h-2v2h2z"/>`), g.Group(children),
	)
}

func RemoteWave(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.004 3a9.99 9.99 0 0 0-8.426 4.612l-.54.841l-1.684-1.079l.54-.842A11.99 11.99 0 0 1 12.004 1c4.247 0 7.978 2.207 10.11 5.532l.539.842l-1.684 1.08l-.54-.842A9.99 9.99 0 0 0 12.004 3m-.012 3.988A5 5 0 0 0 7.78 9.294l-.54.842l-1.683-1.079l.54-.842a7 7 0 0 1 5.896-3.227a7 7 0 0 1 5.897 3.227l.54.842l-1.684 1.08l-.54-.843a5 5 0 0 0-4.213-2.306M5 11h14v12h-2V13H7v10H5zm6 4h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func Remove(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 11h16v2H4z"/>`), g.Group(children),
	)
}

func Replay(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m23 14l-4.4-3.3l1.2-1.6l.928.696A9 9 0 1 0 12 21a9 9 0 0 0 8.252-5.4l.4-.917l1.833.801l-.4.916A11 11 0 0 1 12 23C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11zm-6.197-2L9.5 16.869V7.13zM11.5 10.868v2.264L13.197 12z"/>`), g.Group(children),
	)
}

func Rice(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4C8.855 4 6.17 5.211 5.493 7.714A15 15 0 0 0 5.057 10h13.886a15 15 0 0 0-.435-2.282C17.831 5.212 15.148 4 12 4m8 8H4a8 8 0 1 0 16 0M3.048 10c.08-.886.257-1.859.514-2.809C4.605 3.341 8.564 2 12 2c3.44 0 7.397 1.343 8.439 5.196c.256.948.433 1.92.513 2.804H22v2c0 5.523-4.477 10-10 10S2 17.523 2 12v-2zm7.95-5.002h2.004v2.004h-2.004zm-3 2h2.004v2.004H7.998zm6 0h2.004v2.004h-2.004z"/>`), g.Group(children),
	)
}

func RiceBall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.017 2.403C9.847 1.515 10.887 1 12 1s2.153.515 2.983 1.403c1.603 1.716 3.345 4.17 4.796 6.781c1.448 2.605 2.649 5.44 3.113 7.924c.15.805.169 1.628-.108 2.432c-.28.814-.828 1.504-1.626 2.103c-.367.275-.92.457-1.443.592c-.57.147-1.276.28-2.067.39c-1.585.223-3.578.37-5.645.375h-.005c-2.068-.005-4.06-.152-5.646-.375a19 19 0 0 1-2.067-.39c-.522-.135-1.076-.317-1.443-.592c-.798-.6-1.345-1.29-1.626-2.103c-.277-.804-.258-1.627-.108-2.432c.464-2.483 1.665-5.32 3.113-7.924c1.451-2.61 3.193-5.065 4.796-6.78M12 3c-.462 0-.998.209-1.521.769c-1.47 1.572-3.12 3.885-4.51 6.387c-1.395 2.508-2.486 5.131-2.895 7.32c-.112.6-.094 1.045.033 1.412c.123.357.383.74.936 1.156q-.004-.003 0 0c.011.006.064.036.182.082q.205.08.56.173c.474.122 1.1.241 1.845.346l.37.05V13h10v7.694q.19-.023.37-.05a17 17 0 0 0 1.845-.345a5 5 0 0 0 .56-.173c.118-.046.17-.076.181-.082q.005-.003 0 0c.553-.416.814-.799.937-1.156c.127-.367.145-.812.033-1.413c-.41-2.188-1.5-4.81-2.895-7.32c-1.39-2.5-3.04-4.814-4.51-6.386C12.998 3.209 12.462 3 12 3m3 17.89V15H9v5.89c.946.067 1.961.107 3 .11a44 44 0 0 0 3-.11M10.998 5.128h2.004v2.004h-2.004zM8.296 9.187h2.003v2.004H8.296zm5.405 0h2.003v2.004h-2.003z"/>`), g.Group(children),
	)
}

func Roast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h4v2H6zm8 0h4v2h-4zm-5 4h6v5H9zm2 2v1h2v-1z"/>`), g.Group(children),
	)
}

func Rocket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.557 3.1H20.9v1.343a7 7 0 0 1-2.05 4.95l-7.557 7.556l-4.243-4.243l7.557-7.556a7 7 0 0 1 4.95-2.05m-1.554 9.968l2.26-2.261A9 9 0 0 0 22.9 4.443V1.1h-3.343a9 9 0 0 0-6.364 2.636l-2.261 2.26l-5.657-.707L1.04 9.524L14.475 22.96l4.235-4.235zm-1.792 1.791l.393 3.143l-2.129 2.129l-1.768-1.768zm-7.07-7.071l-3.505 3.504L3.87 9.524l2.129-2.129zm-3.505 9.16l-3.535 3.536L.687 19.07l3.535-3.535zm2.829 2.83l-3.536 3.535l-1.414-1.414l3.535-3.536z"/>`), g.Group(children),
	)
}

func Rollback(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.93 14A7 7 0 0 1 14 20H5.5v-2H14a5 5 0 1 0 0-10H6.914l2.5 2.5L8 11.914L3.086 7L8 2.086L9.414 3.5L6.914 6H14a7 7 0 0 1 7 7v1z"/>`), g.Group(children),
	)
}

func Rollfront(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.07 14A7 7 0 0 0 10 20h8.5v-2H10a5 5 0 0 1 0-10h7.086l-2.5 2.5L16 11.914L20.914 7L16 2.086L14.586 3.5l2.5 2.5H10a7 7 0 0 0-7 7v1z"/>`), g.Group(children),
	)
}

func RootList(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm2 3.5h12v2H6zm0 4h8v2H6z"/>`), g.Group(children),
	)
}

func Rotate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0 0 18a9 9 0 0 0 8.252-5.4l.4-.917l1.833.801l-.4.916A11 11 0 0 1 12 23C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11v2l-4.4-3.3l1.2-1.6l.928.696A9.004 9.004 0 0 0 12 3"/>`), g.Group(children),
	)
}

func RotateLocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0 0 18a9 9 0 0 0 8.252-5.4l1.832.8A11 11 0 0 1 12 23C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11v2l-3.6-2.7l1.2-1.6l.128.096A9.004 9.004 0 0 0 12 3m0 5.5c.69 0 1.25.56 1.25 1.25v.75h-2.5v-.75c0-.69.56-1.25 1.25-1.25m3.25 2v-.75a3.25 3.25 0 0 0-6.5 0v.75H7.499V17h9v-6.5zm-.751 2V15h-5v-2.5z"/>`), g.Group(children),
	)
}

func Rotation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2v9a9 9 0 0 1 9 9h9v2H2V2zm0 18h7a7 7 0 0 0-7-7z"/>`), g.Group(children),
	)
}

func Round(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 6a6 6 0 1 0 0 12a6 6 0 0 0 0-12m-8 6a8 8 0 1 1 16 0a8 8 0 0 1-16 0"/>`), g.Group(children),
	)
}

func RouterWave(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.004 3a9.99 9.99 0 0 0-8.426 4.612l-.54.841l-1.684-1.079l.54-.842A11.99 11.99 0 0 1 12.004 1c4.247 0 7.978 2.207 10.11 5.532l.539.842l-1.684 1.08l-.54-.842A9.99 9.99 0 0 0 12.004 3m-.012 3.988A5 5 0 0 0 7.78 9.294l-.54.842l-1.683-1.079l.54-.842a7 7 0 0 1 5.896-3.227a7 7 0 0 1 5.897 3.227l.54.842l-1.684 1.08l-.54-.843a5 5 0 0 0-4.213-2.306M13 11v3h7v9H4v-9h7v-3zm-7 5v5h12v-5z"/>`), g.Group(children),
	)
}

func Rss(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 3h1c9.389 0 17 7.611 17 17v1h-2v-1c0-8.284-6.716-15-15-15H3zm0 7h1c5.523 0 10 4.477 10 10v1h-2v-1a8 8 0 0 0-8-8H3zm0 9a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func Ruler(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.719 2h11.035L13.64 20H22v2H1.719zm4.869 18l3.657-16H8.28l-.625 2.5h4.22v2h-4.72L6.53 11h4.22v2H6.03l-.625 2.5h4.22v2h-4.72L4.28 20z"/>`), g.Group(children),
	)
}

func SailingHotel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7 .811l1.171.204c3.987.693 6.584 2.801 8.154 5.48c.974 1.661 1.538 3.519 1.822 5.352c.28 1.812.29 3.624.135 5.248c-.139 1.463-.416 2.819-.774 3.905H21v2H4v-2h3zM9 21h6.376c.299-.708.588-1.747.78-3H9zm0-5h7.357c.049-.966.03-1.98-.077-3H9zm0-5h6.95a12.8 12.8 0 0 0-1.079-3H9zm0-5h4.513C12.426 4.77 10.961 3.793 9 3.248z"/>`), g.Group(children),
	)
}

func Sandwich(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.8 2.515a2 2 0 0 1 2.494.476L21 8.638V22H3V8.406zM5 10v4h14v-4zm12.865-2l-3.107-3.73L7.922 8zM19 16H5v4h14z"/>`), g.Group(children),
	)
}

func Saturation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .07l7.068 7a9.856 9.856 0 0 1 0 14.029c-3.905 3.867-10.231 3.867-14.136 0a9.856 9.856 0 0 1 0-14.029zm0 2.814L6.34 8.491a7.856 7.856 0 0 0 0 11.187c3.125 3.095 8.195 3.095 11.32 0a7.856 7.856 0 0 0 0-11.187zm-1 6.15h1a5.482 5.482 0 1 1 0 10.964h-1zm2 2.147v6.671a3.483 3.483 0 0 0 0-6.671"/>`), g.Group(children),
	)
}

func Sausage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19.998 2.003v1.996h2.004v2h-1.257a4.22 4.22 0 0 1-1.04 4.268l-9.44 9.44a4.22 4.22 0 0 1-4.267 1.04V22h-2v-2H2.002v-2h1.249a4.22 4.22 0 0 1 1.04-4.269l9.44-9.44a4.22 4.22 0 0 1 4.267-1.04v-1.25zm-1.707 3.704a2.225 2.225 0 0 0-3.146 0l-9.44 9.44a2.225 2.225 0 1 0 3.146 3.146l9.44-9.44a2.225 2.225 0 0 0 0-3.146"/>`), g.Group(children),
	)
}

func Save(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h14.414L22 7.586V22H2zm2 2v16h2v-6h12v6h2V8.414L15.586 4H13v4H6V4zm4 0v2h3V4zm8 16v-4H8v4z"/>`), g.Group(children),
	)
}

func SavingPot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.198 22.5h5.604l.445-2h2.506l.445 2h5.604l.632-2.843c3.65-3.061 4.69-8.39 2.226-12.658q-.34-.585-.74-1.108Q22 5.455 22 5a5 5 0 0 0-9.007-2.99c-1.344.084-2.7.402-4.005.954l-.92.39l.778 1.841l.921-.39a10.4 10.4 0 0 1 2.31-.678Q12 4.553 12 5a5 5 0 0 0 8.959 3.054a8 8 0 0 1-2.04 10.256l-.283.221l-.438 1.97h-2.396l-.444-2H9.642l-.444 2H6.802l-.669-3.01l-.501-.192c-1.064-.408-2.563-1.27-3.632-2.23V10.5h1.066l2.047-3.41l-.668-1.671c.806.104 1.735.501 2.601 1.454l.673.74l1.48-1.346l-.673-.74c-1.893-2.082-4.17-2.37-5.752-2.001l-1.135.264l1.248 3.12l-.953 1.59H0v7.416l.295.293c1.15 1.144 2.795 2.16 4.115 2.747zm.302-11h2.004V9.494H5.5zM17 8a3 3 0 1 1 0-6a3 3 0 0 1 0 6"/>`), g.Group(children),
	)
}

func Scan(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h7v2H4v5H2zm13 0h7v7h-2V4h-5zM2 11h20v2H2zm2 4v5h5v2H2v-7zm18 0v7h-7v-2h5v-5z"/>`), g.Group(children),
	)
}

func ScreenFourK(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v14h18V5zm12 3v2.048l1-.708V8h2v1.598a1.5 1.5 0 0 1-.633 1.225L15.703 12l1.664 1.177A1.5 1.5 0 0 1 18 14.402V16h-2v-1.34l-1-.708V16h-2V8zM7 8v3.429h2V8h2v8H9v-2.571H7a2 2 0 0 1-2-2V8z"/>`), g.Group(children),
	)
}

func Screencast(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2h22v17h-6v-2h4V4H3l.001 13h4v2H1zm4.586 20L12 15.585L18.414 22zm4.828-2h3.172L12 18.414z"/>`), g.Group(children),
	)
}

func Screenshot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v14h-2V4H4v12H2zm3.403 7h2.822L12 12.797L15.774 9h2.823l-5.186 5.216l2.688 2.704a3.3 3.3 0 0 1 1.615-.42c1.8 0 3.286 1.44 3.286 3.25S19.514 23 17.714 23s-3.285-1.44-3.285-3.25c0-.49.108-.953.303-1.367L12 15.635l-2.732 2.748c.194.414.303.877.303 1.367c0 1.81-1.486 3.25-3.285 3.25S3 21.56 3 19.75s1.486-3.25 3.286-3.25c.585 0 1.137.152 1.615.42l2.688-2.704zm12.311 9.5c-.725 0-1.285.574-1.285 1.25s.56 1.25 1.285 1.25S19 20.426 19 19.75s-.56-1.25-1.286-1.25m-11.428 0C5.56 18.5 5 19.074 5 19.75S5.56 21 6.286 21c.725 0 1.285-.574 1.285-1.25s-.56-1.25-1.285-1.25"/>`), g.Group(children),
	)
}

func ScrollBar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 2v20h-2V2zM2 2h16v20H2zm2 2v16h12V4z"/>`), g.Group(children),
	)
}

func SdCard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 1h16v22H3V11.586l2-2zm2 2v7.414l-2 2V21h14V3zm4 2v3H9V5zm3 0v3h-2V5zm3 0v3h-2V5z"/>`), g.Group(children),
	)
}

func SdCardOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 2h12.515L21 9.68V22H3zm9 2h-1v3.5H9V4H8v3.5H6V4H5v16h14v-9.68L14.485 4H14v3.5h-2z"/>`), g.Group(children),
	)
}

func Search(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.096 5.904a6.5 6.5 0 1 0-9.192 9.192a6.5 6.5 0 0 0 9.192-9.192M4.49 4.49a8.5 8.5 0 0 1 12.686 11.272l5.345 5.345l-1.414 1.414l-5.345-5.345A8.501 8.501 0 0 1 4.49 4.49"/>`), g.Group(children),
	)
}

func SearchError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.097 5.904A6.5 6.5 0 0 0 4 10.504l.001 1h-2v-1a8.5 8.5 0 1 1 15.176 5.258l5.344 5.345l-1.414 1.414l-5.344-5.345A8.48 8.48 0 0 1 10.5 19h-1v-2h1a6.5 6.5 0 0 0 4.596-11.096M1.672 13.257L4.5 16.086l2.829-2.829l1.414 1.415L5.915 17.5l2.828 2.828l-1.414 1.415L4.5 18.914l-2.828 2.829l-1.414-1.415L3.086 17.5L.258 14.672z"/>`), g.Group(children),
	)
}

func Secured(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .44l10 3.5V12c0 4.127-2.533 7.012-4.896 8.803a19.7 19.7 0 0 1-4.65 2.595l-.087.033l-.025.009l-.007.002l-.003.001c-.001 0-.002 0-.332-.943l-.331.944h-.001l-.003-.002l-.007-.002l-.025-.01l-.087-.032a18 18 0 0 1-1.39-.606a20 20 0 0 1-3.26-1.989C4.534 19.012 2 16.127 2 12V3.94zm0 22.06l-.331.944l.331.116l.331-.116zm0-1.072l.009-.004a17.8 17.8 0 0 0 3.887-2.215C18.034 17.59 20 15.223 20 12V5.36l-8-2.8l-8 2.8V12c0 3.223 1.966 5.588 4.104 7.21A17.8 17.8 0 0 0 12 21.428m6.072-13.085l-7.071 7.071l-4.243-4.242l1.415-1.415L11 12.586l5.657-5.657z"/>`), g.Group(children),
	)
}

func Send(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M.292 1.665L24.002 12L.293 22.336L3.94 12zM5.708 13l-2 5.665L18.999 12L3.708 5.336l2 5.664H11v2z"/>`), g.Group(children),
	)
}

func SendCancel(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M.292 1.665L24.002 12l-5.456 2.379l-.8-1.833L18.999 12L3.708 5.336l2 5.664H11v2H5.708l-2 5.665l8.56-3.731l.799 1.833L.292 22.336L3.94 12zm22.79 14.33l-2.832 2.851l2.821 2.822l-1.414 1.414l-2.819-2.82l-2.818 2.819l-1.414-1.415l2.818-2.817l-2.838-2.838L16 14.597l2.836 2.835l2.827-2.846z"/>`), g.Group(children),
	)
}

func Sensors(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 12c0-2.437-.87-4.667-2.317-6.402l-.64-.768l1.535-1.281l.64.767A11.96 11.96 0 0 1 24 12a11.95 11.95 0 0 1-2.782 7.683l-.64.768l-1.536-1.281l.64-.768A9.96 9.96 0 0 0 22 11.999m-4.623-4.483A6.98 6.98 0 0 1 19 12c0 1.704-.61 3.268-1.623 4.482l-.64.768l-1.536-1.281l.64-.768A4.98 4.98 0 0 0 17 12a4.98 4.98 0 0 0-1.158-3.201L15.2 8.03l1.536-1.281zM12 10a2 2 0 1 1 0 4a2 2 0 0 1 0-4M8.799 8.03l-.64.769A4.98 4.98 0 0 0 7 11.999c0 1.219.435 2.333 1.158 3.2l.641.769l-1.536 1.28l-.64-.767A6.98 6.98 0 0 1 5 12c0-1.704.61-3.268 1.623-4.482l.64-.767zm-3.841-3.2l-.64.768A9.96 9.96 0 0 0 2 11.999c0 2.436.87 4.667 2.317 6.402l.64.768l-1.535 1.28l-.64-.767A11.96 11.96 0 0 1 0 12c0-2.921 1.045-5.602 2.781-7.683l.641-.767z"/>`), g.Group(children),
	)
}

func SensorsOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21.219 4.317A11.96 11.96 0 0 1 24 12c0 2.953-1.068 5.659-2.837 7.749l1.251 1.25L21 22.415l-8.482-8.482a2 2 0 0 1-2.45-2.45l-2.24-2.24A4.97 4.97 0 0 0 7 12c0 1.219.435 2.333 1.158 3.2l.641.768l-1.536 1.282l-.64-.768A6.98 6.98 0 0 1 5 12a6.97 6.97 0 0 1 1.394-4.192L4.257 5.67A9.96 9.96 0 0 0 2 12a9.96 9.96 0 0 0 2.317 6.4l.64.768l-1.535 1.282l-.64-.768A11.96 11.96 0 0 1 0 12a11.96 11.96 0 0 1 2.836-7.75L1.586 3L3 1.586l8.482 8.482Q11.73 10 12 10a2 2 0 0 1 1.932 2.518l2.24 2.24A5 5 0 0 0 17 12a4.98 4.98 0 0 0-1.158-3.2l-.642-.77l1.536-1.28l.64.767A6.98 6.98 0 0 1 19 12a6.97 6.97 0 0 1-1.394 4.192l2.137 2.137A9.96 9.96 0 0 0 22 11.999c0-2.435-.87-4.666-2.317-6.4l-.64-.769l1.535-1.28z"/>`), g.Group(children),
	)
}

func SensorsOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m5.636 4.223l-.707.707A9.97 9.97 0 0 0 2 12a9.97 9.97 0 0 0 2.929 7.072l.707.707l-1.414 1.414l-.707-.707A11.97 11.97 0 0 1 0 12.001c0-3.314 1.344-6.315 3.515-8.486l.707-.707zm14.141-1.415l.707.707A11.97 11.97 0 0 1 24 12.001c0 3.313-1.344 6.315-3.514 8.485l-.708.707l-1.414-1.414l.707-.707A9.97 9.97 0 0 0 22 12a9.97 9.97 0 0 0-2.93-7.071l-.707-.707zM8.464 7.051l-.707.707A5.98 5.98 0 0 0 6 12c0 1.658.67 3.156 1.757 4.243l.707.707l-1.414 1.414l-.707-.707A7.98 7.98 0 0 1 4 12c0-2.208.897-4.21 2.343-5.656l.707-.708zm8.485-1.415l.707.708A7.98 7.98 0 0 1 20 12c0 2.21-.897 4.21-2.344 5.657l-.707.707l-1.414-1.414l.707-.707A5.98 5.98 0 0 0 18 12a5.98 5.98 0 0 0-1.758-4.242l-.707-.707zM12 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-4 2a4 4 0 1 1 5 3.874v4.905h-2v-4.905A4 4 0 0 1 8 12"/>`), g.Group(children),
	)
}

func SensorsTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6.343 4.929l-.707.707A8.97 8.97 0 0 0 3 12a8.97 8.97 0 0 0 2.636 6.364l.707.707l-1.414 1.414l-.707-.707A10.97 10.97 0 0 1 1 12c0-3.037 1.232-5.79 3.222-7.778l.707-.707zm12.728-1.414l.707.707A10.97 10.97 0 0 1 23 12c0 3.037-1.232 5.789-3.222 7.778l-.707.707l-1.414-1.414l.707-.707A8.97 8.97 0 0 0 21 12a8.97 8.97 0 0 0-2.636-6.364l-.707-.707zm-9.9 4.242l-.707.707A4.98 4.98 0 0 0 7 12c0 1.381.559 2.63 1.465 3.536l.707.707l-1.414 1.414l-.708-.707A6.98 6.98 0 0 1 5 12c0-1.933.785-3.684 2.05-4.95l.707-.707zm7.072-1.414l.707.707A6.98 6.98 0 0 1 19 12a6.98 6.98 0 0 1-2.05 4.95l-.707.707l-1.415-1.414l.707-.707A4.98 4.98 0 0 0 17 12a4.98 4.98 0 0 0-1.464-3.536l-.708-.707zM13 9v6h-2V9z"/>`), g.Group(children),
	)
}

func Serenity(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h4v2H6zm8 0h4v2h-4zm-5.1 4.634l.5.865C9.922 15.4 10.89 16 12 16s2.08-.601 2.6-1.5l.5-.866l1.731 1.001l-.5.866A5 5 0 0 1 12 18a5 5 0 0 1-4.33-2.5l-.501-.865z"/>`), g.Group(children),
	)
}

func Server(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 13v9H2v-9zm-2 2H4v5h16zm2-13v9H2V2zm-2 2H4v5h16zM7.504 16.5v2.004H5.5V16.5zm0-11v2.004H5.5V5.5z"/>`), g.Group(children),
	)
}

func Service(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 11C2 5.477 6.477 1 12 1s10 4.477 10 10v5.154C22 17.8 20.58 19 19 19h-3v-8h4a8 8 0 1 0-16 0h4v8H6.063A2 2 0 0 0 8 20.5h1.564c.316-.453.841-.75 1.436-.75h2a1.75 1.75 0 1 1 0 3.5h-2c-.595 0-1.12-.297-1.436-.75H8a4 4 0 0 1-3.986-3.66C2.874 18.463 2 17.446 2 16.155zm4 6v-4H4v3.154c0 .393.37.846 1 .846zm14-4h-2v4h1c.63 0 1-.453 1-.846z"/>`), g.Group(children),
	)
}

func Setting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .845l9.66 5.578v11.154L12 23.155l-9.66-5.578V6.423zm0 2.31L4.34 7.577v8.846L12 20.845l7.66-4.422V7.577zM12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0"/>`), g.Group(children),
	)
}

func SettingOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9.18 1h5.64l.647 3.237a8.5 8.5 0 0 1 1.52.88l3.129-1.059l2.82 4.884l-2.481 2.18a8.6 8.6 0 0 1 0 1.756l2.48 2.18l-2.819 4.884l-3.129-1.058a8.5 8.5 0 0 1-1.52.879L14.819 23H9.18l-.647-3.237a8.5 8.5 0 0 1-1.52-.88l-3.129 1.059l-2.82-4.884l2.481-2.18a8.6 8.6 0 0 1 0-1.756l-2.48-2.18l2.82-4.884l3.128 1.058a8.5 8.5 0 0 1 1.52-.879zm1.64 2l-.542 2.705l-.525.193a6.5 6.5 0 0 0-1.912 1.106l-.43.359l-2.615-.885l-1.18 2.044l2.072 1.821l-.095.551a6.6 6.6 0 0 0 0 2.212l.095.55l-2.073 1.822l1.18 2.044l2.616-.885l.43.359a6.5 6.5 0 0 0 1.912 1.106l.525.193L10.82 21h2.36l.542-2.705l.525-.193a6.5 6.5 0 0 0 1.912-1.106l.43-.359l2.616.885l1.18-2.044l-2.072-1.821l.094-.551a6.6 6.6 0 0 0 0-2.212l-.094-.55l2.072-1.822l-1.18-2.044l-2.616.885l-.43-.359a6.5 6.5 0 0 0-1.912-1.106l-.525-.193L13.18 3zM12 9a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0"/>`), g.Group(children),
	)
}

func Share(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.991 1.035a4 4 0 1 1-.855 6.267l-6.28 3.626q.147.533.145 1.072c0 .358-.047.719-.145 1.072l6.28 3.626a4.002 4.002 0 0 1 6.32 4.803a4 4 0 0 1-7.32-3.07l-6.28-3.627a4.002 4.002 0 1 1 0-5.608l6.28-3.626a4 4 0 0 1 1.855-4.535M19.723 3.5a2 2 0 1 0-3.464 2a2 2 0 0 0 3.464-2M3.071 12.527a2.002 2.002 0 0 0 2.93 1.204a2 2 0 1 0-2.93-1.204m15.92 5.242a2 2 0 1 0-2 3.464a2 2 0 0 0 2-3.464"/>`), g.Group(children),
	)
}

func ShareOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 1.85L23.553 12L11.5 22.15v-6.227c-4.194.2-7.073 1.7-9.186 4.658L.52 19.804c.523-2.617 1.58-5.295 3.478-7.462c1.761-2.014 4.209-3.543 7.502-4.187zm2 4.3v3.717l-.858.123c-3.27.467-5.551 1.853-7.14 3.669a12 12 0 0 0-1.744 2.658C6.096 14.666 8.978 13.9 12.5 13.9h1v3.95L20.448 12z"/>`), g.Group(children),
	)
}

func Sharpness(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.5.586V19.5H1.587zM6.416 17.5H18.5V5.414zM20.5 21v2h-19v-2z"/>`), g.Group(children),
	)
}

func ShieldError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.67 23.444L12 22.5l.33.944l-.33.116zm.33-2.016l.009-.004a17.8 17.8 0 0 0 3.887-2.215C18.034 17.59 20 15.223 20 12V5.36l-8-2.8l-8 2.8V12c0 3.223 1.966 5.588 4.104 7.21A17.8 17.8 0 0 0 12 21.428m-5.104-.625C4.534 19.012 2 16.127 2 12V3.94l10-3.5l10 3.5V12c0 4.127-2.533 7.012-4.896 8.803a19.8 19.8 0 0 1-4.65 2.595l-.087.033l-.025.009l-.007.002l-.005.002L12 22.5l-.33.944l-.005-.002l-.007-.002l-.025-.01l-.087-.032a18 18 0 0 1-1.39-.606a20 20 0 0 1-3.26-1.989M11 16.5V10h2v6.5zm2-8h-2.004V6.496H13z"/>`), g.Group(children),
	)
}

func Shimen(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11.999 1.996l11 1.1v4.81l-2.932.293L21.075 22h-7.162L15.01 8.713l-.621.06l-.121.011q-1.026.098-2.168.212l-.1.01l-2.914-.292l.991 13.287H2.923L3.931 8.199L.999 7.906v-4.81zM5.922 8.398L5.075 20h2.847L7.064 8.512zm11.11.113L16.085 20h2.838l-.846-11.6l-.435.046l-.014.001zM3 4.906v1.19l9 .9q1.095-.109 2.076-.203l.12-.011l1.81-.176c.492-.05.95-.098 1.415-.147l.012-.001c.469-.05.945-.1 1.468-.152L21 6.096v-1.19l-9-.9z"/>`), g.Group(children),
	)
}

func Shop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 5.5a4.5 4.5 0 0 1 9 0V7H21v16H3V7h4.5zm0 3.5H5v12h14V9h-2.5v3h-2V9h-5v3h-2zm7-2V5.5a2.5 2.5 0 0 0-5 0V7z"/>`), g.Group(children),
	)
}

func ShopFive(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.28 2h17.44l.972 2.914c.181.543.391 1.332.163 2.154A4 4 0 0 1 21 8.646V20h1v2H2v-2h1V8.646A4 4 0 0 1 2 6v-.162zM5 9.874V20h3v-7h8v7h3V9.874a4 4 0 0 1-4-1.228A4 4 0 0 1 12 10a4 4 0 0 1-3-1.354a3.99 3.99 0 0 1-4 1.228M10 6a2 2 0 1 0 4 0V4h-4zM8 4H4.72l-.715 2.146c.039.533.285 1.008.662 1.345A2 2 0 0 0 8 6zm8 0v2a2 2 0 0 0 3.928.535c.059-.213.026-.512-.133-.989L19.279 4zm-2 16v-5h-4v5z"/>`), g.Group(children),
	)
}

func ShopFour(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 4a2 2 0 0 0-2 2v1h3V4zm3 0v3h4V4zm6 0v3h3V6a2 2 0 0 0-2-2zm5 3h1v2h-1v11h1v2H2v-2h1V9H2V7h1V6a4 4 0 0 1 4-4h10a4 4 0 0 1 4 4zm-2 2H5v11h3v-5a4 4 0 0 1 8 0v5h3zm-5 11v-5a2 2 0 1 0-4 0v5z"/>`), g.Group(children),
	)
}

func ShopOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 7.465V20h4v-7h8v7h4V9.465A4 4 0 0 1 18 10a4 4 0 0 1-3-1.354A4 4 0 0 1 12 10a4 4 0 0 1-3-1.354A4 4 0 0 1 6 10a4 4 0 0 1-2-.535M10 6a2 2 0 1 0 4 0V4h-4zM8 4H4v2a2 2 0 1 0 4 0zm8 0v2a2 2 0 1 0 4 0V4zm-2 16v-5h-4v5z"/>`), g.Group(children),
	)
}

func ShopThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 7.465V20h4v-7h8v7h4V9.465A4 4 0 0 1 18 10a4 4 0 0 1-3-1.354A4 4 0 0 1 12 10a4 4 0 0 1-3-1.354A4 4 0 0 1 6 10a4 4 0 0 1-2-.535M8 6h2a2 2 0 1 0 4 0h2a2 2 0 1 0 4 0V4H4v2a2 2 0 1 0 4 0m6 14v-5h-4v5z"/>`), g.Group(children),
	)
}

func ShopTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4.982 5.11C6.409 3.397 8.684 2 12 2s5.591 1.397 7.018 3.11c1.06 1.27 1.655 2.72 1.879 3.89H22v2h-1v9h1v2H2v-2h1v-9H2V9h1.103c.224-1.17.82-2.62 1.879-3.89M5 11v9h3v-3a4 4 0 1 1 8 0v3h3v-9zm13.846-2a7 7 0 0 0-1.364-2.61A6.4 6.4 0 0 0 14.4 4.341q.272.481.512 1.011c.493 1.097.887 2.379 1.03 3.648zm-4.92 0a10.7 10.7 0 0 0-.838-2.827a9 9 0 0 0-.868-1.536q-.114-.159-.22-.286q-.106.127-.22.286c-.298.413-.6.939-.868 1.536A10.7 10.7 0 0 0 10.074 9zM8.058 9c.143-1.27.536-2.551 1.03-3.648q.239-.531.512-1.01A6.4 6.4 0 0 0 6.518 6.39A7 7 0 0 0 5.154 9zM14 20v-3a2 2 0 0 0-4 0v3z"/>`), g.Group(children),
	)
}

func Shrimp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.598 5.99L14 9.588V5c1.324 0 2.55.359 3.598.99M12 5v6h-2c-2.023 0-3.784-.591-5.02-1.614C3.765 8.38 3 6.913 3 5zm2 9.415l3.59 3.592A7 7 0 0 1 14 19zM12 17v2H8a2 2 0 0 1 2-2zm2 4c2.19 0 4.215-.798 5.773-2.095A8.94 8.94 0 0 0 23 12a9.03 9.03 0 0 0-3.218-6.898A8.9 8.9 0 0 0 14 3H1v2c0 2.505 1.027 4.538 2.706 5.927C5.366 12.301 7.604 13 10 13h2v2h-2a4 4 0 0 0-4 4v2zm5.153-4.258L15.413 13h5.517a6.9 6.9 0 0 1-1.777 3.742M20.928 11h-5.513l3.737-3.735A7.03 7.03 0 0 1 20.928 11M9.002 5.998H6.998v2.004h2.004z"/>`), g.Group(children),
	)
}

func ShrinkHorizontal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 11h5.086l-2.5-2.5L5.5 7.086L10.414 12L5.5 16.914L4.086 15.5l2.5-2.5H1.5zM13 3v18h-2V3zm.586 9L18.5 7.086L19.914 8.5l-2.5 2.5H22.5v2h-5.086l2.5 2.5l-1.414 1.414z"/>`), g.Group(children),
	)
}

func ShrinkVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1.5v5.086l2.5-2.5L16.914 5.5L12 10.414L7.086 5.5L8.5 4.086l2.5 2.5V1.5zM21 13H3v-2h18zm-9 .586l4.914 4.914l-1.414 1.414l-2.5-2.5V22.5h-2v-5.086l-2.5 2.5L7.086 18.5z"/>`), g.Group(children),
	)
}

func Shutter(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.288 6.004L7.504 9.96l3.924-6.943a8.98 8.98 0 0 0-6.14 2.986M13.65 3.15L11.475 7h8.01a9 9 0 0 0-5.835-3.85M20.488 9H15.96l4.025 7.156A8.96 8.96 0 0 0 21 12c0-1.053-.18-2.063-.512-3m-1.768 8.987L16.5 14.04l-3.904 6.94a8.98 8.98 0 0 0 6.124-2.993m-8.347 2.866L12.54 17H4.515a9 9 0 0 0 5.858 3.853M3.512 15H8.04l-1.413-2.511L4.02 7.833A8.96 8.96 0 0 0 3 12c0 1.053.18 2.063.512 3m6.823 0h3.33l1.688-3l-1.688-3h-3.32l-1.697 3.001zM3.289 5.282a10.98 10.98 0 0 1 9.939-4.214A11.01 11.01 0 0 1 22.084 7.6c.59 1.348.916 2.837.916 4.4c0 2.522-.85 4.85-2.28 6.705a10.98 10.98 0 0 1-9.92 4.23A11.01 11.01 0 0 1 1.916 16.4A11 11 0 0 1 1 12c0-2.528.854-4.86 2.289-6.718"/>`), g.Group(children),
	)
}

func Shutup(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h4v2H6zm8 0h4v2h-4zm-3 3.586l1 1l1-1L14.414 14l-1 1l1 1L13 17.414l-1-1l-1 1L9.586 16l1-1l-1-1z"/>`), g.Group(children),
	)
}

func SimCard(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.586 1H21v22H3V5.586zm.828 2L5 6.414V21h14V3zM8 8.998h2.004v2.004H8zm6 0h2.004v2.004H14zM13 9v5h-2V9zm-3 3v5H8v-5zm6 0v5h-2v-5zm-5.002 2.998h2.004v2.004h-2.004z"/>`), g.Group(children),
	)
}

func SimCardOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.586 1H21v22H3V5.586zm.828 2L5 6.414V21h14V3zM10.5 9.006h3v7.982h-2v-5.982h-1z"/>`), g.Group(children),
	)
}

func SimCardTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.586 1H21v22H3V5.586zm.828 2L5 6.414V21h14V3zM12 10a1 1 0 0 0-1 1v1H9v-1a3 3 0 1 1 5.149 2.094l-.028.028l-1.841 1.611H15v2H9v-1.787l3.739-3.272A1 1 0 0 0 12 10"/>`), g.Group(children),
	)
}

func SinisterSmile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h4v2H8v1H6zm8 0h4v2h-2v1h-2zm-5.1 4.634l.5.865C9.922 15.4 10.89 16 12 16s2.08-.601 2.6-1.5l.5-.866l1.731 1.001l-.5.866A5 5 0 0 1 12 18a5 5 0 0 1-4.33-2.5l-.501-.865z"/>`), g.Group(children),
	)
}

func Sip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M20.781 8.132A3.424 3.424 0 1 0 15.94 3.29l-1.748 1.748l-1.286-1.285l-1.414 1.414l1.286 1.285L2.072 17.157v4.842h4.843l10.704-10.704l1.285 1.286l1.415-1.415l-1.286-1.285zm-4.576 1.749L6.086 19.999H4.072v-2.014L14.191 7.867zm-.6-3.429l1.748-1.748a1.424 1.424 0 1 1 2.014 2.014L17.62 8.466z"/>`), g.Group(children),
	)
}

func Slash(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m18.366 2.974l-11 19.052l-1.732-1l11-19.052z"/>`), g.Group(children),
	)
}

func Sleep(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m7-3.5v1a.5.5 0 0 0 1 0v-1h2v1a2.5 2.5 0 0 1-5 0v-1zm7 0v1a.5.5 0 0 0 1 0v-1h2v1a2.5 2.5 0 0 1-5 0v-1zM11 14h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func Slice(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m.186 18.962l7.974 2.683l4.657-4.657l1.185 1.185l8.057-8.059a4.606 4.606 0 0 0-6.515-6.513zm10.13-7.303l6.643-6.644A2.606 2.606 0 1 1 20.645 8.7l-6.643 6.645zm1.087 3.915l-3.78 3.78l-3.742-1.26l5.021-5.02z"/>`), g.Group(children),
	)
}

func Slideshow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 3h12v18H6zm2 2v14h8V5zM1 5h3.5v14H1v-2h1.5V7H1zm18.5 0H23v2h-1.5v10H23v2h-3.5z"/>`), g.Group(children),
	)
}

func Smile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m8-4v4H7V8zm8 0v4h-2V8zm-8.1 5.634l.5.865C9.922 15.4 10.89 16 12 16s2.08-.601 2.6-1.5l.5-.866l1.731 1.001l-.5.866A5 5 0 0 1 12 18a5 5 0 0 1-4.33-2.5l-.501-.865z"/>`), g.Group(children),
	)
}

func Sneer(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h4v2H6zm8 0h4v2h-4zm2.718 4.78l-.25.97c-.269 1.045-.793 1.895-1.613 2.467c-.806.563-1.792.783-2.855.783h-1v-2h1c.8 0 1.343-.167 1.71-.423c.353-.246.647-.646.822-1.326l.249-.969z"/>`), g.Group(children),
	)
}

func Snowflake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v1.586l1-1L15.414 3L13 5.414v4.203l3.943-3.98l1.42 1.407l-3.949 3.987l4.169-.028L21 8.586L22.414 10l-1 1H23v2h-1.586l1 1L21 15.414l-2.411-2.411l-4.175.028l3.95 3.912l-1.407 1.421L13 14.446v4.14L15.414 21L14 22.414l-1-1V23h-2v-1.586l-1 1L8.586 21L11 18.586v-4.14l-3.957 3.918l-1.407-1.421l3.95-3.912l-4.175-.028L3 15.414L1.586 14l1-1H1v-2h1.586l-1-1L3 8.586l2.417 2.417l4.169.028l-3.95-3.988l1.421-1.407L11 9.617V5.414L8.586 3L10 1.586l1 1V1z"/>`), g.Group(children),
	)
}

func Sonic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 2v20h-2V2zM9 6v12H7V6zm8 0v12h-2V6zM5 9v6H3V9zm16 0v6h-2V9z"/>`), g.Group(children),
	)
}

func Sound(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 1.29v21.42L5.737 17.5H1v-11h4.737zM4.999 8.5H3v7h1.999zm2 7.415L13 19.29V4.71L6.999 8.084zm13.98-8.933l.603.798a7 7 0 0 1-.003 8.44l-.603.798l-1.595-1.206l.603-.798a5 5 0 0 0 .002-6.03l-.603-.797zM18.186 9.09l.603.798a3.5 3.5 0 0 1-.001 4.22l-.604.798L16.59 13.7l.603-.797a1.5 1.5 0 0 0 .001-1.809l-.603-.798z"/>`), g.Group(children),
	)
}

func SoundDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 1.29v21.42L5.737 17.5H1v-11h4.737zM4.999 8.5H3v7h1.999zm2 7.415L13 19.29V4.71L6.999 8.084zM16 11h8v2h-8z"/>`), g.Group(children),
	)
}

func SoundHigh(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 1.29v21.42L5.737 17.5H1v-11h4.737zM4.999 8.5H3v7h1.999zm2 7.415L13 19.29V4.71L6.999 8.084zm13.45-10.28l.708.708a8 8 0 0 1 0 11.314l-.707.707l-1.415-1.415l.708-.707a6 6 0 0 0 0-8.485l-.707-.707zm-3.181 3.183l1.414-1.414l.707.707a5.5 5.5 0 0 1 0 7.778l-.707.707l-1.414-1.414l.707-.707a3.5 3.5 0 0 0 0-4.95zm-.354.353l.707.707a3 3 0 0 1 0 4.243l-.707.707l-1.414-1.414l.707-.707a1 1 0 0 0 0-1.414l-.707-.707z"/>`), g.Group(children),
	)
}

func SoundLow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 1.29v21.42L5.737 17.5H1v-11h4.737zM4.999 8.5H3v7h1.999zm2 7.415L13 19.29V4.71L6.999 8.084zm11.683-8.512l.707.707a5.5 5.5 0 0 1 0 7.778l-.707.707l-1.414-1.414l.707-.707a3.5 3.5 0 0 0 0-4.95l-.707-.707z"/>`), g.Group(children),
	)
}

func SoundMute(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M19 22.71L9.737 17.5H5v-11h4.737L19 1.29zM8.999 8.5H7v7h1.999zm2 7.415L17 19.29V4.71l-6.001 3.375z"/>`), g.Group(children),
	)
}

func SoundMuteOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 1.29v21.42L5.737 17.5H1v-11h4.737zM4.999 8.5H3v7h1.999zm2 7.415L13 19.29V4.71L6.999 8.084zm10.88-7.45L20 10.584l2.121-2.12l1.415 1.413L21.414 12l2.121 2.121l-1.414 1.414L20 13.414l-2.121 2.121l-1.415-1.414L18.587 12l-2.121-2.122z"/>`), g.Group(children),
	)
}

func SoundUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 1.29v21.42L5.737 17.5H1v-11h4.737zM4.999 8.5H3v7h1.999zm2 7.415L13 19.29V4.71L6.999 8.084zM19 8h2v3h3v2h-3v3h-2v-3h-3v-2h3z"/>`), g.Group(children),
	)
}

func Space(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 9v5h18V9h2v7H1V9z"/>`), g.Group(children),
	)
}

func SpeechlessOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h5v6h3v2H9v-6H6zm8 0h4v2h-4z"/>`), g.Group(children),
	)
}

func Star(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .63l2.903 8.35l8.839.181l-7.045 5.341l2.56 8.462L12 17.914l-7.256 5.05l2.56-8.462L.259 9.161l8.839-.18zm0 6.092l-1.47 4.23l-4.478.091l3.569 2.706l-1.297 4.287L12 15.478l3.676 2.558l-1.296-4.287l3.568-2.706l-4.477-.09z"/>`), g.Group(children),
	)
}

func StarFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12.001.63l2.903 8.35l8.839.181l-7.045 5.341l2.56 8.462L12 17.914l-7.256 5.05l2.56-8.462L.26 9.161l8.839-.18z"/>`), g.Group(children),
	)
}

func StatueOfJesus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 1h4v4h-4zM2 6h20v2.66l-7 3V17h.78l1.5 6H6.72l1.5-6H9v-5.34l-7-3zm7 3.483V8H5.539zM11 8v2.586l2 2V8zm4 0v1.483L18.461 8zm-2 7.414l-2-2V17h2zM14.22 19H9.78l-.5 2h5.44z"/>`), g.Group(children),
	)
}

func StickyNote(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v12.414L14.414 22H2zm2 2v16h9v-7h7V4zm14.586 11H15v3.586zM6 8h12v2H6zm0 4h5v2H6z"/>`), g.Group(children),
	)
}

func Stop(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 8h8v8H8z"/>`), g.Group(children),
	)
}

func StopCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12"/><path fill="currentColor" d="M8 8h8v8H8z"/>`), g.Group(children),
	)
}

func StopCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 1C5.925 1 1 5.925 1 12s4.925 11 11 11s11-4.925 11-11S18.075 1 12 1m4 7v8H8V8z"/>`), g.Group(children),
	)
}

func StopCircleStroke(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m7-4h8v8H8zm2 2v4h4v-4z"/>`), g.Group(children),
	)
}

func Store(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.5 2.5h17v2h-17zm-.82 3h18.64l1.18 5.901V13.5H21v8h-2v-8h-5v8H3v-8H1.5v-2.099zm2.32 8v6h7v-6zm15.48-2l-.8-4H4.32l-.8 4z"/>`), g.Group(children),
	)
}

func StreetRoad(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm14.372 2.343L6.657 13.372l-1.029-1.715l2.111-1.267L5.64 6.611l1.749-.97l2.066 3.72l7.888-4.733zm0 6l-2.111 1.267l2.099 3.779l-1.748.97l-2.067-3.72l-7.888 4.733l-1.029-1.715l11.715-7.029z"/>`), g.Group(children),
	)
}

func StreetRoadOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm6.177 2.216l-2.393 11.96l-1.96-.391L8.215 5.823zm5.608-.393l2.392 11.962l-1.962.392l-2.392-11.961zM13 7v3h-2V7zm0 4v3h-2v-3zm0 4v3h-2v-3z"/>`), g.Group(children),
	)
}

func Subtitle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v14h18V5zm2 5a2 2 0 0 1 2-2h4v2H7v4h4v2H7a2 2 0 0 1-2-2zm8 0a2 2 0 0 1 2-2h4v2h-4v4h4v2h-4a2 2 0 0 1-2-2z"/>`), g.Group(children),
	)
}

func SubwayLine(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 2v3h9.17A3 3 0 0 1 18 3.17V2h2v1.17a3.001 3.001 0 0 1 0 5.66V19H8.83A3 3 0 0 1 7 20.83V22H5v-1.17A3 3 0 0 1 3.17 19H2v-2h1.17A3 3 0 0 1 5 15.17V7H2V5h3V2zm0 5v8.17A3 3 0 0 1 8.83 17H18V8.83A3 3 0 0 1 16.17 7zm12-2a1 1 0 1 0 0 2a1 1 0 0 0 0-2M6 17a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`), g.Group(children),
	)
}

func Sum(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.586 3H21v2H6.414l7 7l-7 7H21v2H1.586l9-9z"/>`), g.Group(children),
	)
}

func SunFall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 3v3h-2V3zm7.485 3.928L18.364 9.05L16.95 7.636l2.121-2.122zM4.93 5.514l2.12 2.122L5.636 9.05L3.515 6.929zM12 10a4 4 0 0 0-4 4v1H6v-1a6 6 0 0 1 12 0v1h-2v-1a4 4 0 0 0-4-4M1 13h3v2H1zm19 0h3v2h-3zM1 17h8.303L12 18.798L14.697 17H23v2h-7.697L12 21.202L8.697 19H1z"/>`), g.Group(children),
	)
}

func SunRising(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 3v3h-2V3zm7.485 3.928L18.364 9.05L16.95 7.636l2.121-2.122zM4.93 5.514l2.12 2.122L5.636 9.05L3.515 6.929zM12 10a4 4 0 0 0-4 4v1H6v-1a6 6 0 0 1 12 0v1h-2v-1a4 4 0 0 0-4-4M1 13h3v2H1zm19 0h3v2h-3zm-8 2.798L15.303 18H23v2h-8.303L12 18.202L9.303 20H1v-2h7.697z"/>`), g.Group(children),
	)
}

func Sunny(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v3h-2V1zm7.485 3.928L18.364 7.05L16.95 5.636l2.121-2.122zM4.93 3.514l2.12 2.122L5.636 7.05L3.515 4.929zM12 8a4 4 0 1 0 0 8a4 4 0 0 0 0-8m-6 4a6 6 0 1 1 12 0a6 6 0 0 1-12 0m-5-1h3v2H1zm19 0h3v2h-3zM7.05 18.363l-2.12 2.123l-1.415-1.416l2.121-2.122zm11.314-1.414l2.121 2.122l-1.414 1.414l-2.121-2.121zM13 20v3h-2v-3z"/>`), g.Group(children),
	)
}

func Support(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m6.382 4.968l2.86 2.86a5.01 5.01 0 0 1 5.516 0l2.86-2.86a9.004 9.004 0 0 0-11.236 0m12.65 1.414l-2.86 2.86a5.01 5.01 0 0 1 0 5.516l2.86 2.86a9.004 9.004 0 0 0 0-11.236m-1.414 12.65l-2.86-2.86a5.01 5.01 0 0 1-5.516 0l-2.86 2.86a9.004 9.004 0 0 0 11.236 0m-12.65-1.414l2.86-2.86a5.01 5.01 0 0 1 0-5.516l-2.86-2.86a9.004 9.004 0 0 0 0 11.236M4.222 4.222c4.296-4.296 11.26-4.296 15.556 0s4.296 11.26 0 15.556s-11.26 4.296-15.556 0s-4.296-11.26 0-15.556m9.9 5.657a3 3 0 1 0-4.243 4.242A3 3 0 0 0 14.12 9.88"/>`), g.Group(children),
	)
}

func Surprised(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-2a2 2 0 1 1 4 0a2 2 0 0 1-4 0m8 0a2 2 0 1 1 4 0a2 2 0 0 1-4 0m-2 4a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0"/>`), g.Group(children),
	)
}

func SurprisedOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m8-4v4H7V8zm8 0v4h-2V8zm-5 6c-.356 0-1 .452-1 1.5s.644 1.5 1 1.5s1-.452 1-1.5s-.644-1.5-1-1.5m-3 1.5c0-1.713 1.146-3.5 3-3.5s3 1.787 3 3.5s-1.146 3.5-3 3.5s-3-1.787-3-3.5"/>`), g.Group(children),
	)
}

func Swap(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15 3.086l7.414 7.414H2v-2h15.586l-4-4zM22 13.5v2H6.414l4 4L9 20.914L1.586 13.5z"/>`), g.Group(children),
	)
}

func SwapLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10.414 9l-4 4H22v2H1.586L9 7.586z"/>`), g.Group(children),
	)
}

func SwapRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15 7.586L22.414 15H2v-2h15.586l-4-4z"/>`), g.Group(children),
	)
}

func SwearOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5.769-3.866l3.464 2l-1 1.732l-3.464-2zm11.464 1.732l-3.464 2l-1-1.732l3.464-2zM8.019 16.805C8.42 14.802 9.91 13 12 13s3.582 1.802 3.98 3.805L16.22 18H7.781zM10.445 16h3.11c-.422-.662-1.02-1-1.555-1s-1.133.338-1.555 1"/>`), g.Group(children),
	)
}

func SwearTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5.769-3.866l3.464 2l-1 1.732l-3.464-2zm11.464 1.732l-3.464 2l-1-1.732l3.464-2zM12 14c-.356 0-1 .452-1 1.5s.644 1.5 1 1.5s1-.452 1-1.5s-.644-1.5-1-1.5m-3 1.5c0-1.713 1.146-3.5 3-3.5s3 1.787 3 3.5s-1.146 3.5-3 3.5s-3-1.787-3-3.5"/>`), g.Group(children),
	)
}

func SystemApplication(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 13v9H2v-9zm-2 2H4v5h16zm2-13v9H2V2zm-2 2H4v5h16zM7.504 16.5v2.004H5.5V16.5zm0-11v2.004H5.5V5.5z"/><path fill="currentColor" d="M11.504 7.504V5.5H9.5v2.004zm0 11V16.5H9.5v2.004z"/>`), g.Group(children),
	)
}

func SystemBlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h22v9.5h-2V3H3v13h8.5v2H1zm17.5 13a3.5 3.5 0 0 0-3.08 5.165l4.745-4.744A3.5 3.5 0 0 0 18.5 14m3.08 1.835l-4.745 4.744a3.5 3.5 0 0 0 4.745-4.745M13 17.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M2.25 21h9.25v2H2.25z"/>`), g.Group(children),
	)
}

func SystemCode(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h22v17H1zm2 2v13h18V3zm7.406 3.844L8.28 9.5l2.125 2.656l-1.562 1.25L5.719 9.5l3.125-3.906zm4.75-1.25L18.281 9.5l-3.125 3.906l-1.562-1.25L15.72 9.5l-2.125-2.656zM3.222 21h17.556v2H3.222z"/>`), g.Group(children),
	)
}

func SystemComponents(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4c-.663 0-1.222.548-1.222 1.25S11.337 6.5 12 6.5s1.222-.548 1.222-1.25S12.663 4 12 4M8.778 5.25C8.778 3.467 10.208 2 12 2c1.791 0 3.222 1.467 3.222 3.25A3.25 3.25 0 0 1 13 8.34v.836l3.849 2.25v4.457l.818.411a3.2 3.2 0 0 1 2.11-.794C21.57 15.5 23 16.967 23 18.75S21.569 22 19.778 22s-3.222-1.467-3.222-3.25q0-.379.082-.735l-.753-.378L12 19.908l-3.885-2.27l-.753.378q.082.355.082.734c0 1.783-1.43 3.25-3.222 3.25C2.431 22 1 20.533 1 18.75s1.431-3.25 3.222-3.25c.81 0 1.548.3 2.112.795l.817-.411v-4.458L11 9.176v-.835A3.25 3.25 0 0 1 8.778 5.25M12 10.908l-2.849 1.666v3.352L12 17.592l2.849-1.666v-3.352zm-6.775 7.126a1.22 1.22 0 0 0-1.003-.534C3.56 17.5 3 18.048 3 18.75S3.559 20 4.222 20c.664 0 1.222-.548 1.222-1.25q-.001-.22-.068-.416zm13.399.303a1.3 1.3 0 0 0-.068.413c0 .702.558 1.25 1.222 1.25S21 19.452 21 18.75s-.559-1.25-1.222-1.25c-.41 0-.779.209-1.002.533z"/>`), g.Group(children),
	)
}

func SystemCoordinate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.272 1v2.65l6.613 3.89v7.745l2.465 1.295l-.93 1.77l-2.507-1.317l-5.641 3.317V23h-2v-2.65l-5.69-3.346l-3.12 1.304l-.77-1.845l2.967-1.24V7.539l6.613-3.889V1zm-1 4.382l-4.64 2.73l4.64 2.728l4.64-2.729zm5.613 4.477l-4.613 2.713v5.458l4.613-2.713zm-6.613 8.17v-5.457L6.66 9.859v5.458z"/>`), g.Group(children),
	)
}

func SystemDevice(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.533 1H21v21H2V6.588h4.533zm2 5.588h4.534V20H19V3H8.533zM11.067 20V8.588H4V20zM6 16h3v2H6z"/>`), g.Group(children),
	)
}

func SystemInterface(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h22v22H1zm2 7.556V21h18V8.556zm18-2V3H3v3.556zM6 11h2.004v2.004H6zm4 0h2.004v2.004H10zm4 0h2.004v2.004H14z"/>`), g.Group(children),
	)
}

func SystemLocation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h22v10h-2V3H3v13h9v2H1zm17.5 13a2.75 2.75 0 0 0-2.75 2.75c0 1.252.735 2.454 1.615 3.422c.407.448.817.814 1.135 1.075c.318-.26.728-.627 1.135-1.075c.88-.968 1.615-2.17 1.615-3.422A2.75 2.75 0 0 0 18.5 14m0 9.701l-.555-.369l-.002-.001l-.004-.003l-.012-.008l-.04-.027l-.137-.098a13 13 0 0 1-1.865-1.677c-.995-1.094-2.135-2.767-2.135-4.768a4.75 4.75 0 1 1 9.5 0c0 2.001-1.14 3.674-2.135 4.768a13 13 0 0 1-2.002 1.774l-.04.028l-.012.008l-.004.003h-.002zM17.25 16h2.5v2h-2.5zm-15 5H12v2H2.25z"/>`), g.Group(children),
	)
}

func SystemLocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h22v10h-2V3H3v13h10v2H1zm18.502 13.5c.69 0 1.25.56 1.25 1.25v.75h-2.5v-.75c0-.69.56-1.25 1.25-1.25m3.25 2v-.75a3.25 3.25 0 0 0-6.5 0v.75H15V23h9v-6.5zm-.752 2V21h-5v-2.5zM2.25 21H13v2H2.25z"/>`), g.Group(children),
	)
}

func SystemLog(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h22v22H1zm2 8.667V21h18V9.667zm18-2V3H3v4.667zM5 4h2.004v2.004H5zm1 8h12v2H6zm0 4h6v2H6z"/>`), g.Group(children),
	)
}

func SystemMarked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h22v10h-2V3H3v13h9.5v2H1zm13.5 11.996H23v10.295l-4.247-2.617L14.5 23.29zm2 2v4.715l2.254-1.385L21 19.709v-4.713zM2.25 21H12.5v2H2.25z"/>`), g.Group(children),
	)
}

func SystemMessages(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m15-6.766v13.532L9.723 15H6V9h3.723zm-2 3.532L10.277 11H8v2h2.277L14 15.234z"/>`), g.Group(children),
	)
}

func SystemRegulation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h22v17H1zm2 2v13h18V3zm6 2v5H7V5zm4 0v3h-2V5zm4 0v5h-2V5zm-4 4v5h-2V9zm-4 2v3H7v-3zm8 0v3h-2v-3zM3.222 21h17.556v2H3.222z"/>`), g.Group(children),
	)
}

func SystemSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h22v11h-2V3H3v13h9v2H1zm17.25 14a2.75 2.75 0 0 1 1.947 4.693l-.01.008A2.75 2.75 0 1 1 18.25 15m3.992 5.325a4.75 4.75 0 1 0-1.414 1.415l1.67 1.674l1.416-1.412zM2.25 21H12v2H2.25z"/>`), g.Group(children),
	)
}

func SystemSetting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2h22v8.25h-2V4H3v12h8.5v2H1zm2 18h8.5v2H3z"/><path fill="currentColor" d="M19.5 12v1.376c.715.184 1.352.56 1.854 1.072l1.193-.689l1 1.732l-1.192.688a4 4 0 0 1 0 2.142l1.192.688l-1 1.732l-1.193-.689a4 4 0 0 1-1.854 1.072V22.5h-2v-1.376a4 4 0 0 1-1.854-1.072l-1.193.689l-1-1.732l1.192-.688a4 4 0 0 1 0-2.142l-1.192-.688l1-1.732l1.193.688a4 4 0 0 1 1.854-1.071V12zm-2.751 4.283a2 2 0 0 0-.25.967c0 .35.091.68.25.967l.036.063a2 2 0 0 0 3.43 0l.036-.063c.159-.287.249-.616.249-.967c0-.35-.09-.68-.249-.967l-.036-.063a2 2 0 0 0-3.43 0z"/>`), g.Group(children),
	)
}

func SystemStorage(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.5 4v3.625H11V4zM13 4v5.625H6.5V4H4v16h16V9.04L14.96 4zm2.79-2L22 8.212v13.79H2V2zM7 14v-2h10v2zm0 3v-2h6v2z"/>`), g.Group(children),
	)
}

func SystemSum(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.527 5.693a32 32 0 0 1 2.424 2.194a32 32 0 0 1 2.193 2.424a17.5 17.5 0 0 0 1.379-2.499c.454-1.033.681-1.914.71-2.598c.029-.682-.14-1.059-.346-1.264s-.582-.374-1.264-.345c-.683.028-1.564.255-2.598.71a17.6 17.6 0 0 0-2.498 1.378M11.839 4.42c1.163-.81 2.306-1.463 3.38-1.936c1.168-.513 2.302-.835 3.32-.878c1.02-.042 2.028.195 2.763.93s.972 1.743.93 2.762c-.044 1.018-.365 2.153-.879 3.32c-.472 1.074-1.126 2.217-1.935 3.38c.81 1.164 1.463 2.307 1.936 3.382c.513 1.167.835 2.302.878 3.32c.042 1.018-.195 2.027-.93 2.762s-1.743.972-2.762.93c-1.018-.043-2.153-.365-3.32-.879c-1.074-.472-2.218-1.126-3.381-1.936c-1.163.81-2.307 1.464-3.381 1.936c-1.167.514-2.302.835-3.32.878c-1.019.043-2.028-.194-2.762-.93c-.735-.734-.973-1.743-.93-2.762c.043-1.017.365-2.152.878-3.32c.473-1.074 1.127-2.217 1.936-3.38c-.809-1.164-1.463-2.307-1.935-3.381c-.514-1.167-.835-2.302-.878-3.32c-.043-1.019.194-2.027.93-2.762c.734-.735 1.743-.972 2.762-.93c1.017.043 2.152.365 3.32.878c1.073.473 2.217 1.127 3.38 1.936m-6.305 5.89a32 32 0 0 1 2.192-2.423a32 32 0 0 1 2.425-2.194a17.5 17.5 0 0 0-2.498-1.378c-1.034-.455-1.915-.682-2.598-.71c-.682-.029-1.06.14-1.265.345s-.374.582-.345 1.264c.029.684.256 1.565.71 2.598a17.5 17.5 0 0 0 1.379 2.499m1.212 1.688a30 30 0 0 0 2.395 2.699a30 30 0 0 0 2.698 2.395a30 30 0 0 0 2.698-2.395a30 30 0 0 0 2.395-2.699a30 30 0 0 0-2.395-2.697a30 30 0 0 0-2.698-2.396A30 30 0 0 0 9.14 9.301a30 30 0 0 0-2.395 2.697m-1.212 1.688a17.5 17.5 0 0 0-1.38 2.5c-.454 1.033-.68 1.914-.71 2.597c-.028.682.14 1.06.346 1.265s.582.374 1.264.345c.683-.029 1.565-.256 2.598-.71a17.5 17.5 0 0 0 2.5-1.38a32 32 0 0 1-2.425-2.192a32 32 0 0 1-2.193-2.425m7.993 4.618a17.5 17.5 0 0 0 2.499 1.379c1.033.454 1.915.681 2.598.71s1.059-.14 1.264-.345c.205-.206.374-.583.345-1.265c-.028-.683-.255-1.564-.71-2.598a17.5 17.5 0 0 0-1.379-2.499a32 32 0 0 1-2.193 2.425a32 32 0 0 1-2.424 2.193M11 11h2.004v2.004H11z"/>`), g.Group(children),
	)
}

func SystemThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h5.077A3.926 3.926 0 0 1 11 5.923V11H5.923A3.926 3.926 0 0 1 2 7.077zm2 2v3.077C4 8.137 4.864 9 5.923 9H9V5.923C9 4.863 8.136 4 7.077 4zm12.923 0C15.863 4 15 4.864 15 5.923V9h3.077C19.137 9 20 8.136 20 7.077V4zM13 5.923A3.926 3.926 0 0 1 16.923 2H22v5.077A3.926 3.926 0 0 1 18.077 11H13zM5.923 15C4.863 15 4 15.864 4 16.923V20h3.077C8.137 20 9 19.136 9 18.077V15zM2 16.923A3.926 3.926 0 0 1 5.923 13H11v5.077A3.926 3.926 0 0 1 7.077 22H2zM13 13h5.077A3.926 3.926 0 0 1 22 16.923V22h-5.077A3.926 3.926 0 0 1 13 18.077zm2 2v3.077c0 1.06.864 1.923 1.923 1.923H20v-3.077c0-1.06-.863-1.923-1.923-1.923z"/>`), g.Group(children),
	)
}

func SystemTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M10 4.37V13H8V4.052l.042-.14c.123-.413.45-1.115 1.063-1.728C9.74 1.546 10.695 1 12 1c1.306 0 2.259.546 2.895 1.184c.613.613.94 1.315 1.063 1.728l.042.14v2.385h-2V4.37a2.5 2.5 0 0 0-.52-.773C13.158 3.274 12.694 3 12 3s-1.158.274-1.48.597a2.5 2.5 0 0 0-.52.773M4.052 8h2.385v2H4.37a2.5 2.5 0 0 0-.773.52C3.274 10.842 3 11.306 3 12s.274 1.158.597 1.48c.276.276.586.442.773.52H13v2H4.052l-.14-.042c-.413-.123-1.115-.45-1.728-1.063C1.546 14.26 1 13.305 1 12c0-1.306.546-2.259 1.184-2.895c.613-.613 1.315-.94 1.728-1.063zM11 8h8.948l.14.042c.413.123 1.114.45 1.729 1.063C22.454 9.74 23 10.695 23 12c0 1.306-.546 2.259-1.183 2.895c-.614.613-1.316.94-1.729 1.063l-.14.042h-2.384v-2h2.066a2.5 2.5 0 0 0 .773-.52c.323-.322.597-.786.597-1.48s-.274-1.158-.597-1.48a2.5 2.5 0 0 0-.773-.52H11zm5 3v8.948l-.042.14c-.123.413-.45 1.114-1.063 1.729C14.26 22.454 13.305 23 12 23c-1.306 0-2.259-.546-2.895-1.183c-.613-.614-.94-1.316-1.063-1.729L8 19.948v-2.384h2v2.066c.078.187.244.497.52.773c.322.323.786.597 1.48.597s1.158-.274 1.48-.597a2.5 2.5 0 0 0 .52-.773V11z"/>`), g.Group(children),
	)
}

func SystemUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 1h22v10h-2V3H3v13h10v2H1zm18.502 13.5c-.69 0-1.25.56-1.25 1.25v.75H24V23h-9v-6.5h1.252v-.75a3.25 3.25 0 0 1 5.54-2.305l.71.705l-1.41 1.418l-.71-.705a1.24 1.24 0 0 0-.88-.363M17 18.5V21h5v-2.5zM2.25 21H13v2H2.25z"/>`), g.Group(children),
	)
}

func Tab(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22.5 20h-21v-2h21zm0-16v12h-6V4zm-2 2h-2v8h2zM15 4v12H9V4zm-2 2h-2v8h2zM7.5 4v12h-6V4zm-2 2h-2v8h2z"/>`), g.Group(children),
	)
}

func Table(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 22h20V2H2zm2-2v-5h4v5zm6 0v-5h4v5zm6 0v-5h4v5zm4-7h-4V8h4zm0-7H4V4h16zM4 8h4v5H4zm6 5V8h4v5z"/>`), g.Group(children),
	)
}

func TableAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v2h16V4zm16 4H4v12h16zm-7 2v3h3v2h-3v3h-2v-3H8v-2h3v-3z"/>`), g.Group(children),
	)
}

func TableOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v4h16V4zm16 6h-4v10h4zm-6 10V10h-4v10zm-6 0V10H4v10z"/>`), g.Group(children),
	)
}

func TableSplit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v2h7V4zm9 0v2h7V4zm7 4H4v12h16zm-9.002 1.998h2.004v2.004h-2.004zm-6 3h2.004v2.004H4.998zm3 0h2.004v2.004H7.998zm3 0h2.004v2.004h-2.004zm3 0h2.004v2.004h-2.004zm3 0h2.004v2.004h-2.004zm-6 3h2.004v2.004h-2.004z"/>`), g.Group(children),
	)
}

func TableTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 3.5h21v17h-21zm2 2v3H11v-3zm9.5 0v3h7.5v-3zm7.5 5H13v3h7.5zm0 5H13v3h7.5zm-9.5 3v-3H3.5v3zm-7.5-5H11v-3H3.5z"/>`), g.Group(children),
	)
}

func Tag(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10.878 23.02l-9.9-9.9L11.1 3.016l9.9-.017v9.917zm.001-2.827L19 12.085V5.002l-7.07.012l-8.122 8.108zm4.117-11.19V7H17v2.004z"/>`), g.Group(children),
	)
}

func Tangerinr(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9.504 1.586l2.521 2.535l2.712-2.534l1.365 1.46l-1.023.958c.786.037 1.572.297 2.206.587c.657.3 1.262.683 1.642 1.028c1.51 1.155 2.981 3.706 3.069 6.616c.091 3.042-1.322 6.402-5.492 9.04a4.9 4.9 0 0 1-1.596.625c-.542.112-1.149.145-1.695.015a5.6 5.6 0 0 1-.838-.26l-.07-.027a.2.2 0 0 0-.055-.016A2 2 0 0 0 12 21.6c-.155 0-.222.007-.25.012a.2.2 0 0 0-.055.016l-.07.027a5.6 5.6 0 0 1-.838.26c-.546.13-1.153.097-1.695-.015a4.9 4.9 0 0 1-1.596-.626c-4.17-2.637-5.583-5.996-5.492-9.038c.088-2.911 1.559-5.462 3.07-6.617c.379-.345.984-.728 1.641-1.028c.678-.31 1.529-.586 2.37-.592l-1-1.004zM7.546 6.41c-.555.253-.975.545-1.141.702l-.043.041l-.048.036c-1 .745-2.24 2.728-2.31 5.107c-.07 2.276.937 4.996 4.561 7.288c.244.154.578.285.93.357c.36.074.656.07.83.028c.303-.071.433-.122.569-.175l.103-.04c.298-.111.557-.154 1.003-.154s.706.043 1.003.154l.103.04c.136.053.266.104.568.175c.175.042.47.046.83-.028c.353-.072.687-.203.93-.357c3.625-2.292 4.631-5.012 4.563-7.289c-.072-2.378-1.312-4.361-2.31-5.106l-.049-.036l-.043-.04c-.166-.158-.586-.45-1.141-.703c-.553-.252-1.118-.41-1.567-.41c-.854 0-1.503.17-1.843.3A3 3 0 0 1 12 6.513c-.412 0-.775-.112-1.044-.214c-.34-.129-.99-.3-1.843-.3c-.449 0-1.014.159-1.567.411M8 8.001h2.004v2.003H8zm-2 3h2.004v2.003H6zm2 3h2.004v2.003H8z"/>`), g.Group(children),
	)
}

func Tape(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v14h18V5h-1.465l-2 3H6.465l-2-3zm3.869 0l.666 1h8.93l.666-1zM8 12a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0m11-1a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0"/>`), g.Group(children),
	)
}

func Task(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v20H3V3h4zm0 4H5v16h14V5h-2v2H7zm8-2H9v2h6zm-6 8h6v2H9zm0 4h6v2H9z"/>`), g.Group(children),
	)
}

func TaskAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v9h-2V5h-2v2H7V5H5v16h7v2H3V3h4zm2 4h6V3H9zm11 9v4h4v2h-4v4h-2v-4h-4v-2h4v-4z"/>`), g.Group(children),
	)
}

func TaskAddOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0 0 18h1v2h-1C5.925 23 1 18.075 1 12S5.925 1 12 1c1.498 0 2.928.3 4.232.844l.923.385l-.77 1.846l-.923-.385A9 9 0 0 0 12 3m11.414 1.5L12 15.914L6.086 10L7.5 8.586l4.5 4.5l10-10zM20 15v3h3v2h-3v3h-2v-3h-3v-2h3v-3z"/>`), g.Group(children),
	)
}

func TaskChecked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v11h-2V5h-2v2H7V5H5v16h6v2H3V3h4zm2 4h6V3H9zm14.657 11.586l-7.07 7.07l-4.243-4.242L13.758 18l2.828 2.828l5.657-5.656z"/>`), g.Group(children),
	)
}

func TaskError(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v9h-2V5h-2v2H7V5H5v16h7v2H3V3h4zm2 4h6V3H9zm7.172 9.757L19 17.586l2.828-2.829l1.415 1.415L20.414 19l2.829 2.828l-1.415 1.415L19 20.414l-2.828 2.829l-1.415-1.415L17.586 19l-2.829-2.828z"/>`), g.Group(children),
	)
}

func TaskLocation(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.5 0v2h7V0h2v2H21v8.5h-2V9H5v12h9v2H3V2h3.5V0zM5 7h14V4H5zm13.75 7A2.75 2.75 0 0 0 16 16.75c0 1.252.735 2.454 1.615 3.422c.407.448.817.814 1.135 1.075c.318-.26.728-.627 1.135-1.075c.88-.968 1.615-2.17 1.615-3.422A2.75 2.75 0 0 0 18.75 14m0 9.701c-.25-.167-.506-.329-.75-.506a13 13 0 0 1-1.865-1.677C15.14 20.424 14 18.75 14 16.75a4.75 4.75 0 1 1 9.5 0c0 2.001-1.14 3.674-2.135 4.768a13 13 0 0 1-1.865 1.677c-.244.177-.5.339-.75.506M17.5 16H20v2h-2.5z"/>`), g.Group(children),
	)
}

func TaskMarked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v8h-2V5h-2v2H7V5H5v16h7.5v2H3V3h4zm2 4h6V3H9zm5.75 7.996h8.5v10.295l-4.247-2.617l-4.253 2.615zm2 2v4.715l2.254-1.385l2.246 1.383v-4.713z"/>`), g.Group(children),
	)
}

func TaskOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v20H3V3h4zm0 4H5v16h14V5h-2v2H7zm8-2H9v2h6z"/>`), g.Group(children),
	)
}

func TaskSetting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 1h10v2h4v9h-2V5h-2v2H7V5H5v16h7v2H3V3h4zm2 4h6V3H9zm11 8.75v1.376c.715.184 1.352.56 1.854 1.072l1.193-.689l1 1.732l-1.192.688a4 4 0 0 1 0 2.142l1.192.688l-1 1.732l-1.193-.689A4 4 0 0 1 20 22.874v1.376h-2v-1.376a4 4 0 0 1-1.854-1.072l-1.193.689l-1-1.732l1.192-.688a4 4 0 0 1 0-2.142l-1.192-.688l1-1.732l1.193.689A4 4 0 0 1 18 15.126V13.75zm-2.751 4.283a2 2 0 0 0-.25.967c0 .35.091.68.25.967l.036.063a2 2 0 0 0 3.43 0l.036-.063A2 2 0 0 0 21 19c0-.35-.09-.68-.249-.967l-.036-.063a2 2 0 0 0-3.43 0z"/>`), g.Group(children),
	)
}

func TaskVisible(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 1H7v2H3v20h7v-2H5V5h2v2h10V5h2v8h2V3h-4zm-2 4H9V3h6zm.75 15v-2h2.5v2z"/><path fill="currentColor" d="M17.002 23.5c4.419 0 6.09-4.5 6.09-4.5s-1.673-4.5-6.09-4.5s-6.09 4.5-6.09 4.5s1.672 4.5 6.09 4.5m-.002-2c-2.615 0-3.87-2.5-3.87-2.5s1.25-2.5 3.87-2.5s3.87 2.5 3.87 2.5s-1.254 2.5-3.87 2.5"/>`), g.Group(children),
	)
}

func Tea(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 2v5H9V2zM7 3v4H5V3zm8 0v4h-2V3zM2.927 8H21v5a4 4 0 0 1-4 4h-.933a6.67 6.67 0 0 1-2.583 3H20v2H3v-2h3.506a6.6 6.6 0 0 1-3.084-5.039v-.007zm13.648 7H17a2 2 0 0 0 2-2v-3h-2.068l-.353 4.954v.007zm-1.648-5H5.074l.341 4.797A4.603 4.603 0 0 0 10 19c2.059 0 3.836-1.38 4.412-3.289c.09-.295.145-.587.173-.914z"/>`), g.Group(children),
	)
}

func Teahouse(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 2v1h8V2h2v1h1v2h-1v1.667L20.5 10H22v2h-1v8h1v2H2v-2h1v-8H2v-2h1.5L6 6.667V5H5V3h1V2zm0 3v1h8V5zm8.5 3h-9L6 10h12zm2.5 4h-3v8h3zm-5 8v-8h-4v8zm-6 0v-8H5v8z"/>`), g.Group(children),
	)
}

func Template(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v4h16V4zm16 6h-9v10h9zM9 20V10H4v10z"/>`), g.Group(children),
	)
}

func Temple(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 1.586L19.414 9H21v2h-1v9h2v2H2v-2h2v-9H3V9h1.586zM6 11v9h5v-9zm7 0v9h5v-9zm3.586-2L12 4.414L7.414 9z"/>`), g.Group(children),
	)
}

func Terminal(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.5 5.586L9.914 12L3.5 18.414L2.086 17l5-5l-5-5zM12 18h10v2H12z"/>`), g.Group(children),
	)
}

func TerminalRectangle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23 3v18H1V3zm-2 2H3v14h18zm-2 12h-7v-2h7zm-8.93-5l-3.739 3.744l-1.415-1.413L7.244 12L4.916 9.67L6.33 8.255z"/>`), g.Group(children),
	)
}

func TerminalRectangleOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v3h18V5zm18 5H3v9h18zM7 11.086l3.414 3.414L7 17.914L5.586 16.5l2-2l-2-2zM12 15h6v2h-6z"/>`), g.Group(children),
	)
}

func TerminalWindow(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23 3H1v18h22zM3 10h18v9H3zm0-2V5h18v3zm4 9H5v-5h2z"/>`), g.Group(children),
	)
}

func Textbox(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm7 4.5H6v-2h12v2h-5V18h-2z"/>`), g.Group(children),
	)
}

func TextformatBold(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 3H5v18h9a5 5 0 0 0 2.436-9.367A5 5 0 0 0 13 3zm7 8H7V5h6a3 3 0 1 1 0 6m-6 2h7a3 3 0 1 1 0 6H7z"/>`), g.Group(children),
	)
}

func TextformatColor(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19.813 17.525l-.394-.919l-6-14L13.16 2h-2.32l-.26.606l-6 14l-.393.92l1.838.787l.394-.92l1.824-4.254h7.515l1.823 4.255l.394.92zM9.791 11.14H9.1L12 4.372l2.9 6.767H9.791M19 22h1v-2H4v2z"/>`), g.Group(children),
	)
}

func TextformatItalic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 3H18v2h-3.67l-2.625 14H16.5v2H6v-2h3.67l2.625-14H7.5z"/>`), g.Group(children),
	)
}

func TextformatStrikethrough(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.977 3.083C9.3 2.276 10.87 2 12.003 2c2.454 0 4.286.991 6.053 2.161l.834.553l-1.104 1.667l-.834-.552C15.279 4.722 13.855 4 12.002 4c-.87 0-2.054.222-2.983.79c-.88.536-1.522 1.363-1.522 2.709c0 .456.078.827.2 1.132l.37.929l-1.858.74l-.37-.929A5 5 0 0 1 5.497 7.5c0-2.158 1.11-3.58 2.48-4.417M4 11h16v2h-2.963c.853.801 1.463 1.932 1.463 3.504c0 2.158-1.11 3.58-2.48 4.417c-1.323.807-2.892 1.082-4.025 1.082c-2.457 0-4.286-.989-6.051-2.17l-.832-.556l1.113-1.662l.83.556c1.666 1.114 3.086 1.832 4.94 1.832c.871 0 2.054-.223 2.983-.79c.88-.536 1.522-1.363 1.522-2.71c0-1.302-.638-2.05-1.549-2.564c-.948-.536-2.116-.767-3.047-.938h-3.9V13H4z"/>`), g.Group(children),
	)
}

func TextformatUnderline(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M18 3v9a6 6 0 0 1-12 0V3h2v9a4 4 0 0 0 8 0V3zM4 20h16v2H4z"/>`), g.Group(children),
	)
}

func TextformatWrap(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 4h20v2H2zm0 7h15c1.148 0 2.38.284 3.35 1.012C21.36 12.77 22 13.946 22 15.5s-.64 2.73-1.65 3.488C19.38 19.716 18.148 20 17 20h-5.414l1.707-1.707l2-2l.707-.707L17.414 17l-.707.707l-.293.293H17c.852 0 1.62-.216 2.15-.613c.49-.367.85-.941.85-1.887s-.36-1.52-.85-1.887C18.62 13.216 17.852 13 17 13H2zm1 7H2v2h8v-2z"/>`), g.Group(children),
	)
}

func Theaters(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M22 22H2V2h20zM20 4h-2.5v2.5H20zm0 4.5h-2.5V11H20zm0 4.5h-2.5v2.5H20zm-2.5 4.5V20H20v-2.5zm-2-2V13H15v-2h.5V8.5H15v-2h.5V4h-7v2.5H9v2h-.5V11H9v2h-.5v2.5H9v2h-.5V20h7v-2.5H15v-2zM6.5 4H4v2.5h2.5zm0 4.5H4V11h2.5zm0 4.5H4v2.5h2.5zm0 4.5H4V20h2.5z"/>`), g.Group(children),
	)
}

func ThumbDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10.879 22.617l1.279-.213A4 4 0 0 0 15.5 18.46V16h5.32a2 2 0 0 0 1.972-2.329l-1.666-10a2 2 0 0 0-1.973-1.67H7v11.197zm1.234-2.254L9 12.803V4h10.153l1.667 10H13.5v4.459a2 2 0 0 1-1.387 1.904M4 14V2H2v12z"/>`), g.Group(children),
	)
}

func ThumbDownOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m13.446 22.236l-1.716-.572a3.7 3.7 0 0 1-2.53-3.51V15.7H5.332a3 3 0 0 1-2.965-3.456l1.184-7.7A3 3 0 0 1 6.516 2H21v11.9h-3.85zm-1.09-2.472L15.85 11.9H19V4H6.515a1 1 0 0 0-.988.848l-1.185 7.7a1 1 0 0 0 .989 1.152H11.2v4.454a1.7 1.7 0 0 0 1.154 1.61"/>`), g.Group(children),
	)
}

func ThumbDownTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m13.446 22.236l-1.716-.572a3.7 3.7 0 0 1-2.53-3.51V15.7H5.332a3 3 0 0 1-2.965-3.456l1.184-7.7A3 3 0 0 1 6.516 2H22v11.9h-4.85zM17.5 11.9H20V4h-2.5zm-2-7.9H6.517a1 1 0 0 0-.988.848l-1.185 7.7a1 1 0 0 0 .989 1.152H11.2v4.454a1.7 1.7 0 0 0 1.154 1.61l3.146-7.076z"/>`), g.Group(children),
	)
}

func ThumbUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10.879 1.383l1.279.213A4 4 0 0 1 15.5 5.54V8h5.32a2 2 0 0 1 1.972 2.329l-1.666 10a2 2 0 0 1-1.973 1.67H7V10.803zm1.234 2.254L9 11.197V20h10.153l1.667-10H13.5V5.54a2 2 0 0 0-1.387-1.904M4 10v12H2V10z"/>`), g.Group(children),
	)
}

func ThumbUpOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10.555 1.764l1.715.572a3.7 3.7 0 0 1 2.53 3.51V8.3h3.869a3 3 0 0 1 2.965 3.456l-1.185 7.7A3 3 0 0 1 17.484 22H3V10.1h3.85zm1.09 2.472L8.15 12.1H5V20h12.484a1 1 0 0 0 .988-.848l1.185-7.7a1 1 0 0 0-.988-1.152H12.8V5.846a1.7 1.7 0 0 0-1.155-1.61"/>`), g.Group(children),
	)
}

func ThumbUpTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m10.555 1.764l1.715.572a3.7 3.7 0 0 1 2.53 3.51V8.3h3.869a3 3 0 0 1 2.965 3.456l-1.185 7.7A3 3 0 0 1 17.484 22H2V10.1h4.85zM6.5 12.1H4V20h2.5zm2 7.9h8.984a1 1 0 0 0 .988-.848l1.185-7.7a1 1 0 0 0-.988-1.152H12.8V5.846a1.7 1.7 0 0 0-1.155-1.61L8.5 11.312z"/>`), g.Group(children),
	)
}

func Thunder(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.1 5h-.064C8.6 5.033 6.7 6.935 6.7 9.2q0 .394.073.765l.184.956l-.95.21C4.832 11.391 4 12.391 4 13.533c0 .947.566 1.79 1.432 2.205l.902.431l-.864 1.804l-.902-.431C3.063 16.822 2 15.309 2 13.533c0-1.83 1.125-3.375 2.706-4.07A6 6 0 0 1 4.7 9.2c0-3.43 2.851-6.152 6.309-6.2h.091c.546 0 1.078.066 1.586.192c2.147.53 3.88 2.12 4.533 4.187a5.6 5.6 0 0 1 2.654 1.074c1.285.969 2.127 2.49 2.127 4.214c0 2.21-1.382 4.082-3.313 4.894l-.922.388l-.775-1.843l.922-.388C19.16 15.193 20 14.01 20 12.667c0-1.051-.511-1.998-1.331-2.616a3.6 3.6 0 0 0-2.228-.717l-.854.013l-.147-.842c-.282-1.62-1.55-2.956-3.234-3.371A4.6 4.6 0 0 0 11.1 5m2.777 5.677L11.817 14h4.018l-4.023 6.38l-1.691-1.067L12.21 16H8.226l3.95-6.377z"/>`), g.Group(children),
	)
}

func Thunderstorm(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.1 5h-.064C8.6 5.033 6.7 6.935 6.7 9.2q0 .394.073.765l.184.956l-.95.21C4.832 11.391 4 12.391 4 13.533c0 .947.566 1.79 1.432 2.205l.902.431l-.864 1.804l-.902-.431C3.063 16.822 2 15.309 2 13.533c0-1.83 1.125-3.375 2.706-4.07A6 6 0 0 1 4.7 9.2c0-3.43 2.851-6.152 6.309-6.2h.091c.546 0 1.078.066 1.586.192c2.147.53 3.88 2.12 4.533 4.187a5.6 5.6 0 0 1 2.654 1.074c1.285.969 2.127 2.49 2.127 4.214c0 2.21-1.382 4.082-3.313 4.894l-.922.388l-.775-1.843l.922-.388C19.16 15.193 20 14.01 20 12.667c0-1.051-.511-1.998-1.331-2.616a3.6 3.6 0 0 0-2.228-.717l-.854.013l-.147-.842c-.282-1.62-1.55-2.956-3.234-3.371A4.6 4.6 0 0 0 11.1 5m2.777 5.677L11.817 14h4.018l-4.023 6.38l-1.691-1.067L12.21 16H8.226l3.95-6.377zm-6.879 7.321h2.004v2.004H6.998zm8 0h2.004v2.004h-2.004z"/>`), g.Group(children),
	)
}

func ThunderstormNight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m17.131.901l-.102 1.185q-.017.194-.017.394a4.505 4.505 0 0 0 4.899 4.488l1.185-.102l-.103 1.184a6.5 6.5 0 0 1-2.162 4.3a5.2 5.2 0 0 1 1.186 3.309c0 2.212-1.383 4.085-3.316 4.898l-.921.388l-.776-1.844l.922-.387c1.25-.526 2.09-1.71 2.09-3.055c0-1.052-.51-2-1.332-2.62a3.6 3.6 0 0 0-2.172-.718h-.059l-.854.014l-.146-.842C15.17 9.87 13.9 8.535 12.215 8.118a4.6 4.6 0 0 0-1.172-.134c-2.438.034-4.34 1.938-4.34 4.205q0 .393.072.766l.184.956l-.95.21C4.833 14.38 4 15.38 4 16.526c0 .948.567 1.793 1.434 2.207l.902.432l-.864 1.804l-.902-.431C3.064 19.817 2 18.303 2 16.526c0-1.831 1.126-3.378 2.708-4.074a6 6 0 0 1-.005-.263c0-3.146 2.398-5.699 5.463-6.138a6.51 6.51 0 0 1 5.78-5.047zm-4.9 5.178q.234.041.463.097c2.15.531 3.883 2.122 4.538 4.19a5.6 5.6 0 0 1 2.045.676a4.5 4.5 0 0 0 1.5-2.099a6.51 6.51 0 0 1-5.723-5.722a4.52 4.52 0 0 0-2.823 2.858m1.646 7.577l-2.06 3.323h4.018l-4.023 6.38l-1.691-1.067l2.089-3.313H8.226l3.95-6.376z"/>`), g.Group(children),
	)
}

func ThunderstormSunny(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M16.003 1v2.004H14V1zm-5.95 1.635l1.417 1.417l-1.417 1.417l-1.417-1.417zm9.898 0l1.416 1.416l-1.415 1.417l-1.418-1.416zm-8.992 3.42A7 7 0 0 0 10.1 6c-3.499 0-6.4 2.74-6.4 6.2q0 .132.006.263C2.125 13.158 1 14.703 1 16.533c0 1.776 1.063 3.288 2.568 4.009l.902.431l.864-1.804l-.902-.431C3.566 18.323 3 17.48 3 16.534c0-1.143.832-2.143 2.007-2.403l.95-.21l-.184-.955A4 4 0 0 1 5.7 12.2C5.7 9.916 7.634 8 10.1 8q.576.002 1.106.134c1.683.416 2.952 1.75 3.234 3.372l.147.841l.854-.013h.059a3.6 3.6 0 0 1 1.759.45C18.314 13.37 19 14.451 19 15.667c0 1.344-.84 2.526-2.088 3.051l-.922.388l.775 1.843l.922-.387C19.618 19.749 21 17.877 21 15.667c0-1.575-.702-2.979-1.803-3.949a5 5 0 0 0-8.238-5.663m2.108.65a3 3 0 0 1 4.413 3.984a5.6 5.6 0 0 0-1.26-.31c-.507-1.601-1.66-2.916-3.153-3.673M21 8h2.004v2.004H21zm-8.123 5.677L10.817 17h4.018l-4.023 6.38l-1.691-1.067L11.21 19H7.226l3.95-6.377z"/>`), g.Group(children),
	)
}

func Ticket(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 4h22v5.469l-.57.271a2.5 2.5 0 0 0 0 4.52l.57.271V20H1v-5.469l.57-.271a2.5 2.5 0 0 0 0-4.52L1 9.47zm2 2v2.258c1.205.806 2 2.18 2 3.742a4.5 4.5 0 0 1-2 3.742V18h18v-2.258A4.5 4.5 0 0 1 19 12c0-1.561.795-2.936 2-3.742V6zm5 3h8v2H8zm0 4h8v2H8z"/>`), g.Group(children),
	)
}

func Time(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-8 4.414l-4-4V5.5h2v6.086L16.414 15z"/>`), g.Group(children),
	)
}

func TimeFilled(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 23C5.925 23 1 18.075 1 12S5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11m1-17.5h-2v6.914l4 4L16.414 15L13 11.586z"/>`), g.Group(children),
	)
}

func Tips(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M23 2v17h-7.586L12 22.414L8.586 19H1V2zm-2 2H3v13h6.414L12 19.586L14.586 17H21z"/>`), g.Group(children),
	)
}

func TipsDouble(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21 7H6v11h4.914l2.586 2.586L16.086 18H21zm2 13h-6.086L13.5 23.414L10.086 20H4V5h19zM19 3.5H2.5v12h-2v-14H19z"/>`), g.Group(children),
	)
}

func Tomato(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m9.504 1.586l2.521 2.535l2.712-2.534l1.365 1.46l-1.023.958c.415.019.816.097 1.203.211c1.06.312 2.108.916 2.644 1.404c1.511 1.154 2.982 3.705 3.07 6.616c.091 3.042-1.322 6.402-5.492 9.04a4.9 4.9 0 0 1-1.596.625c-.542.112-1.149.145-1.695.015a5.6 5.6 0 0 1-.838-.26l-.07-.027a.2.2 0 0 0-.055-.016A2 2 0 0 0 12 21.6c-.155 0-.222.007-.25.012a.2.2 0 0 0-.055.016l-.07.027a5.6 5.6 0 0 1-.838.26c-.546.13-1.153.097-1.695-.015a4.9 4.9 0 0 1-1.596-.626c-4.17-2.637-5.583-5.996-5.492-9.038c.088-2.911 1.559-5.462 3.07-6.617c.536-.489 1.566-1.096 2.651-1.406A5.2 5.2 0 0 1 9.085 4l-1-1.004zm.966 4.568l.839.802a1 1 0 0 0 1.382 0l.84-.802c-.2.047-.364.1-.487.146A3 3 0 0 1 12 6.514c-.412 0-.775-.112-1.044-.214a4 4 0 0 0-.486-.146m5.777.17L14.073 8.4a3 3 0 0 1-4.146 0L7.75 6.321c-.62.253-1.137.595-1.345.792l-.043.041l-.048.036c-1 .745-2.24 2.728-2.31 5.107c-.07 2.276.937 4.996 4.561 7.288c.244.154.578.285.93.357c.36.074.656.07.83.028c.303-.071.433-.122.569-.175l.103-.04c.298-.111.557-.154 1.003-.154s.706.043 1.003.154l.103.04c.136.053.266.104.568.175c.175.042.47.046.83-.028c.353-.072.687-.203.93-.357c3.625-2.292 4.631-5.012 4.563-7.289c-.072-2.378-1.312-4.361-2.31-5.106l-.049-.036l-.043-.04c-.2-.191-.726-.533-1.348-.79"/>`), g.Group(children),
	)
}

func Tools(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.534 1.866a8.502 8.502 0 0 1 11.487 9.793l5.154 5.153l-1.415 1.414l-5.987-5.987l.178-.576a6.5 6.5 0 0 0-1.615-6.518A6.5 6.5 0 0 0 8.34 3.392l4.228 4.228l-4.95 4.95L3.39 8.34a6.5 6.5 0 0 0 1.754 5.996a6.5 6.5 0 0 0 6.518 1.615l.575-.178l5.988 5.988l-1.414 1.414l-5.154-5.153A8.502 8.502 0 0 1 1.864 6.535l.254-.624l1.672.002l3.828 3.828L9.74 7.62L5.912 3.792L5.91 2.12zm9.216 12.471l4.95 4.95l-1.414 1.414l-4.95-4.95z"/>`), g.Group(children),
	)
}

func ToolsCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 21a9 9 0 1 0 0-18a9 9 0 0 0 0 18m11-9c0 6.075-4.925 11-11 11S1 18.075 1 12S5.925 1 12 1s11 4.925 11 11m-7.31 7.225l-4.207-4.207a5.001 5.001 0 0 1-5.885-6.814l.318-.738l1.355.262L9.114 9.57l.459-.46L7.73 7.27l-.263-1.354l.739-.319a5.001 5.001 0 0 1 6.813 5.886l4.207 4.208zm.708-3.535l-3.71-3.71l.263-.62a3 3 0 0 0-2.46-4.157l1.91 1.909l-3.287 3.287l-1.91-1.91a3 3 0 0 0 4.157 2.462l.62-.264l3.71 3.71z"/>`), g.Group(children),
	)
}

func Tornado(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.528 4A.514.514 0 0 0 5 4.5c0 .259.219.5.528.5H19v2H5.528C4.149 7 3 5.898 3 4.5S4.15 2 5.528 2H8v2zM21 8v2H6V8zM8 11h11v2H8zm-2 3h9v2H6zm-1 3h8v2H5zm3 3h7v2H8z"/>`), g.Group(children),
	)
}

func Tower(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v1h3v6h-.977c.079 1.872.37 4.353.903 6.731c.317 1.414.714 2.768 1.192 3.925c.355.86.74 1.574 1.145 2.118q.112.123.201.226H20v2h-6v-2h1.674q-.133-.123-.275-.246c-.54-.47-1.14-.921-1.749-1.25c-.615-.33-1.174-.504-1.65-.504s-1.035.173-1.65.505c-.608.328-1.209.778-1.75 1.249q-.14.123-.274.246H10v2H4v-2h1.536l.2-.226c.405-.544.79-1.259 1.146-2.118c.478-1.157.875-2.511 1.192-3.925c.533-2.378.824-4.859.903-6.731H8V2h3V1zm-2.021 7a42 42 0 0 1-.953 7.169c-.2.89-.433 1.774-.703 2.618l.076-.042C10.173 17.327 11.07 17 12 17c.932 0 1.827.327 2.6.745l.077.042a31 31 0 0 1-.703-2.618A42 42 0 0 1 13.021 8zM10 6h4V4h-4z"/>`), g.Group(children),
	)
}

func TowerClock(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .932l5.288 1.983l-.703 1.873l-.585-.22V5h3v12h-2v6H7v-6H5V5h3v-.432l-.585.22l-.702-1.873zm-2 2.886V5h4V3.818l-2-.75zM9 17v4h6v-4h-2v3h-2v-3zm8-2V7H7v8zm-5-5a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0"/>`), g.Group(children),
	)
}

func TowerOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 1v1h3v6h-.939c.2 1.785.874 4.199 1.819 6.581c1.055 2.662 2.37 5.084 3.552 6.419H22v2h-7.847l-.139-.831l-.002-.012l-.014-.066a4 4 0 0 0-.424-1.088C13.238 20.428 12.764 20 12 20c-.763 0-1.238.427-1.574 1.004a4 4 0 0 0-.438 1.153l-.002.011l-.139.832H2v-2h1.568c1.182-1.335 2.497-3.757 3.552-6.419C8.065 12.2 8.74 9.785 8.94 8H8V2h3V1zm-2.051 7c-.2 2.1-.972 4.803-1.97 7.319c-.816 2.059-1.83 4.107-2.898 5.681H8.24c.103-.294.251-.646.46-1.004C9.236 19.073 10.261 18 12 18s2.763 1.073 3.302 1.996c.208.358.356.71.459 1.004h2.158c-1.069-1.574-2.082-3.622-2.899-5.681c-.997-2.516-1.768-5.218-1.969-7.319zM10 6h4V4h-4z"/>`), g.Group(children),
	)
}

func TowerThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11 3.126V1h2v2.126a4.002 4.002 0 0 1 1.575 6.935L16.127 23H7.873l1.553-12.939A4.002 4.002 0 0 1 11 3.126m.334 7.819L10.127 21h3.746l-1.207-10.055a4 4 0 0 1-1.332 0M12 5a2 2 0 1 0 0 4a2 2 0 0 0 0-4"/>`), g.Group(children),
	)
}

func TowerTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .586l2.973 2.973L15.936 18H17v5H7v-5h1.064l.963-14.44zM10.069 18h3.862l-.904-13.558L12 3.414l-1.027 1.028zM9 20v1h6v-1z"/>`), g.Group(children),
	)
}

func Town(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h5.442L11 5.913V2h5.468L22 8.638V22H2zm11 2v16h7V9.362L15.532 4zm-2 16V8.887L6.558 4H4v16zM6 7.998h2.004v2.004H6zm9 0h2.004v2.004H15zm-9 4h2.004v2.004H6zm9 0h2.004v2.004H15zm-9 4h2.004v2.004H6zm9 0h2.004v2.004H15z"/>`), g.Group(children),
	)
}

func Traffic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 2h12v4h3v2h-3v3h3v2h-3v3h3v2h-3v4H6v-4H3v-2h3v-3H3v-2h3V8H3V6h3zm2 2v16h8V4zm2 3a2 2 0 1 1 4 0a2 2 0 0 1-4 0m0 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0m0 5a2 2 0 1 1 4 0a2 2 0 0 1-4 0"/>`), g.Group(children),
	)
}

func TrafficEvents(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 1.914l3.643 8.687l.406.899l-.025.01L19.165 19H21v2H3v-2h1.835l3.14-7.49l-.024-.01l.406-.9zM10.149 11.5l-1.145 2.73h5.992l-1.145-2.73zm2.863-2L12 7.086L10.988 9.5zm2.823 6.73h-7.67L7.004 19h9.992z"/>`), g.Group(children),
	)
}

func Transform(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 3a1 1 0 1 0 0 2a1 1 0 0 0 0-2M1 4a3 3 0 0 1 5.83-1h10.34A3.001 3.001 0 1 1 21 6.83v10.34A3.001 3.001 0 1 1 17.17 21H6.83A3.001 3.001 0 1 1 3 17.17V6.83A3 3 0 0 1 1 4m4 2.83v10.34A3 3 0 0 1 6.83 19h10.34A3 3 0 0 1 19 17.17V6.83A3 3 0 0 1 17.17 5H6.83A3 3 0 0 1 5 6.83M20 3a1 1 0 1 0 0 2a1 1 0 0 0 0-2M4 19a1 1 0 1 0 0 2a1 1 0 0 0 0-2m16 0a1 1 0 1 0 0 2a1 1 0 0 0 0-2"/>`), g.Group(children),
	)
}

func TransformOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h6v2h8V2h6v6h-2v8h2v6h-6v-2H8v2H2v-6h2V8H2zm4 6v8h2v2h8v-2h2V8h-2V6H8v2zm0-4H4v2h2zm14 2V4h-2v2zm-2 12v2h2v-2zM6 20v-2H4v2z"/>`), g.Group(children),
	)
}

func TransformThree(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 2.5V6h8.586L20 2.586L21.414 4L18 7.414V16h3.5v2H18v3.5h-2V18H6V8H2.5V6H6V2.5zM8 8v6.586L14.586 8zm8 1.414L9.414 16H16z"/>`), g.Group(children),
	)
}

func TransformTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8 1v15h6v2H6V8H1V6h5V1zm2 5h8v10h5v2h-5v5h-2V8h-6z"/>`), g.Group(children),
	)
}

func Translate(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 5a3 3 0 0 1 3-3h2a3 3 0 0 1 3 3v6H8V8.5H4V11H2zm2 1.5h4V5a1 1 0 0 0-1-1H5a1 1 0 0 0-1 1zm8-3h6a3 3 0 0 1 3 3V9h-2V6.5a1 1 0 0 0-1-1h-6zm6 8V13h4v2h-1.062a7.97 7.97 0 0 1-2.19 4.563A6 6 0 0 0 21 20h1v2h-1a7.96 7.96 0 0 1-4-1.07A7.96 7.96 0 0 1 13 22h-1v-2h1c.796 0 1.556-.155 2.251-.437a8 8 0 0 1-1.48-2.134l-.43-.903l1.807-.858l.429.903A6 6 0 0 0 17 18.472A6 6 0 0 0 18.917 15H12v-2h4v-1.5zM6 13v6a1 1 0 0 0 1 1h2.5v2H7a3 3 0 0 1-3-3v-6z"/>`), g.Group(children),
	)
}

func TranslateOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.5 2v2.5H2v2h7.46a15 15 0 0 1-2.085 3.665A15 15 0 0 1 5.91 7.868l-.459-.889l-1.777.917l.458.889a17 17 0 0 0 1.894 2.9a15 15 0 0 1-3.028 2.309l-.865.5l1.001 1.732l.866-.501a17 17 0 0 0 3.374-2.56a17 17 0 0 0 3.374 2.56l.866.501l1.001-1.731l-.865-.5a15 15 0 0 1-3.029-2.31A17 17 0 0 0 11.59 6.5H13v-2H8.5V2zm10 7.164l-5.832 12.312l1.808.856L13.58 20h5.84l1.104 2.332l1.808-.856zM18.472 18h-3.944l1.972-4.164z"/>`), g.Group(children),
	)
}

func TreeList(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 22a4 4 0 0 1-1-7.874V9.874A4.002 4.002 0 0 1 6 2a4 4 0 0 1 1 7.874v4.252A4.002 4.002 0 0 1 6 22m-2-4a2 2 0 1 0 4 0a2 2 0 0 0-4 0M4 6a2 2 0 1 0 4 0a2 2 0 0 0-4 0m8 12h7v2h-7zm0-7h10v2H12zm0-7h7v2h-7z"/>`), g.Group(children),
	)
}

func TreeRoundDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4m-1 5.874A4.002 4.002 0 0 1 12 1a4 4 0 0 1 1 7.874V11h4a3 3 0 0 1 3 3v1.126A4.002 4.002 0 0 1 19 23a4 4 0 0 1-1-7.874V14a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v1.126A4.002 4.002 0 0 1 5 23a4 4 0 0 1-1-7.874V14a3 3 0 0 1 3-3h4zM19.003 17h-.006a2 2 0 1 0 .006 0M5 17a2 2 0 1 0 0 4a2 2 0 0 0 0-4"/>`), g.Group(children),
	)
}

func TreeRoundDotVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.126 4A4.002 4.002 0 0 1 23 5a4 4 0 0 1-7.874 1H14a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h1.126A4.002 4.002 0 0 1 23 19a4 4 0 0 1-7.874 1H14a3 3 0 0 1-3-3v-4H8.874A4.002 4.002 0 0 1 1 12a4 4 0 0 1 7.874-1H11V7a3 3 0 0 1 3-3zM19 3a2 2 0 1 0 0 4a2 2 0 0 0 0-4M5 10a2 2 0 1 0 0 4a2 2 0 0 0 0-4m14 7a2 2 0 1 0 0 4a2 2 0 0 0 0-4"/>`), g.Group(children),
	)
}

func TreeSquareDot(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.5 2h7v7H13v2.333h6.5V15H22v7h-7v-7h2.5v-1.667h-11V15H9v7H2v-7h2.5v-3.667H11V9H8.5zm5 5V4h-3v3zM4 17v3h3v-3zm13 0v3h3v-3z"/>`), g.Group(children),
	)
}

func TreeSquareDotVertical(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 4h3v3h-3zm-2 2.5V9h7V2h-7v2.5h-3.667V11H9V8.5H2v7h7V13h2.333v6.5H15V22h7v-7h-7v2.5h-1.667v-11zM17 20v-3h3v3zM7 10.5v3H4v-3z"/>`), g.Group(children),
	)
}

func TrendingDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.086 7.5L8.5 13.914l4-4L17.586 15H13.5v2H21V9.5h-2v4.086l-6.5-6.5l-4 4l-5-5z"/>`), g.Group(children),
	)
}

func TrendingUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2.086 16.5L8.5 10.086l4 4L17.586 9H13.5V7H21v7.5h-2v-4.086l-6.5 6.5l-4-4l-5 5z"/>`), g.Group(children),
	)
}

func Tv(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 2h22v17h-6v-2h4V4H3l.001 13h4v2H1zm11 13.587l4.242 4.242l2.42 2.415h-2.833l-1-1L12 18.417l-2.827 2.828l-1 1H5.34l2.419-2.414z"/>`), g.Group(children),
	)
}

func TvOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 3h24v15h-4.046l2.698 2.063l-1.215 1.589L16.66 18H7.336l-4.771 3.652l-1.216-1.588L4.044 18H0zm2 2v11h20V5z"/>`), g.Group(children),
	)
}

func TvTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 3h24v15H0zm2 2v11h20V5zm2 15h16v2H4z"/>`), g.Group(children),
	)
}

func Typography(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 3h15v8H2zm2 2v4h11V5zm-2 9h20v2H2zm0 5h20v2H2z"/>`), g.Group(children),
	)
}

func Uncomfortable(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m9.233-2.134l-3.464 2l-1-1.732l3.464-2zm4.536-1.732l3.464 2l-1 1.732l-3.464-2zM9.5 15.5a1 1 0 0 0-1 1v1h-2v-1a3 3 0 0 1 5.065-2.177a1 1 0 0 0 .206.158a.2.2 0 0 0 .044.019h.37a.2.2 0 0 0 .044-.019a1 1 0 0 0 .207-.158A3 3 0 0 1 17.5 16.5v1h-2.001v-1a1 1 0 0 0-1.688-.726c-.322.306-.878.726-1.621.726h-.382c-.743 0-1.299-.42-1.621-.726A1 1 0 0 0 9.5 15.5m2.318-1"/>`), g.Group(children),
	)
}

func UncomfortableOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m6.619-4.862L11 9.04v1.46l-3.2 2.4l-1.2-1.6l1.86-1.394L6.637 8.88zm9.743 1.743L15.54 9.906L17.4 11.3l-1.2 1.6l-3.2-2.4V9.04l3.381-1.902zM7.67 15.5A5 5 0 0 1 12 13a5 5 0 0 1 4.33 2.5l.501.865l-1.731 1.001l-.5-.865A3 3 0 0 0 12 15a3 3 0 0 0-2.6 1.5l-.5.866l-1.731-1.001z"/>`), g.Group(children),
	)
}

func UncomfortableTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m6.619-4.862L11 9.04v1.46l-3.2 2.4l-1.2-1.6l1.86-1.394L6.637 8.88zm9.743 1.743L15.54 9.906L17.4 11.3l-1.2 1.6l-3.2-2.4V9.04l3.381-1.902zM9.5 15.5a1 1 0 0 0-1 1v1h-2v-1a3 3 0 0 1 5.065-2.177a1 1 0 0 0 .206.158a.2.2 0 0 0 .044.019h.37a.2.2 0 0 0 .044-.019a1 1 0 0 0 .207-.158A3 3 0 0 1 17.5 16.5v1h-2.001v-1a1 1 0 0 0-1.688-.726c-.322.306-.878.726-1.621.726h-.382c-.743 0-1.299-.42-1.621-.726A1 1 0 0 0 9.5 15.5m2.318-1"/>`), g.Group(children),
	)
}

func Undertake(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.861 8.777a1.38 1.38 0 0 0-.976.402l-2.184 2.185V15.5h6.014l6.147-1.537l3.741-1.596a.645.645 0 0 0-.48-1.188l-.02.006l-6.773 1.557h-3.757v-2h3.247a.983.983 0 0 0 0-1.965zm7.919 1.35l3.836-.883a2.647 2.647 0 0 1 1.86 4.924l-.027.013l-3.948 1.684l-6.54 1.635H0V9.95h4.286l2.187-2.187a3.38 3.38 0 0 1 2.392-.986h.001h-.002h4.956a2.983 2.983 0 0 1 2.96 3.35M3.7 11.949H2V15.5h1.7z"/>`), g.Group(children),
	)
}

func UndertakeDelivery(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 2h14v9.5h-2V4h-2v5.618l-3-1.5l-3 1.5V4h-2v5.5h-2zm6 2v2.382l1-.5l1 .5V4zm-5.065 9.25a1.25 1.25 0 0 0-.885.364l-2.05 2.05V19.5h5.627l5.803-1.45l3.532-1.508a.555.555 0 0 0-.416-1.022l-.02.005L13.614 17H10v-2h3.125a.875.875 0 1 0 0-1.75zm7.552 1.152l3.552-.817a2.56 2.56 0 0 1 3.211 2.47a2.56 2.56 0 0 1-1.414 2.287l-.027.014l-3.74 1.595l-6.196 1.549H0v-7.25h4.086l2.052-2.052a3.25 3.25 0 0 1 2.3-.948h.002h-.002h4.687a2.875 2.875 0 0 1 2.862 3.152M3.5 16.25H2v3.25h1.5z"/>`), g.Group(children),
	)
}

func UndertakeEnvironmentProtection(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17 3.513c.21.257.446.561.681.897C18.421 5.466 19 6.617 19 7.5a2 2 0 1 1-4 0c0-.883.579-2.034 1.319-3.09c.235-.336.471-.64.681-.897M17 .57l-.714.73h-.002l-.003.004l-.008.008l-.03.031l-.105.112a18 18 0 0 0-1.457 1.806C13.921 4.346 13 5.945 13 7.5a4 4 0 0 0 8 0c0-1.555-.921-3.154-1.681-4.238a18 18 0 0 0-1.562-1.917l-.03-.031l-.008-.01l-.003-.002l-.002-.002zM8.435 13.25a1.25 1.25 0 0 0-.885.364l-2.05 2.05V19.5h5.627l5.804-1.451l3.531-1.507a.555.555 0 0 0-.163-1.032a.6.6 0 0 0-.253.01l-.02.005L13.614 17H10v-2h3.125a.875.875 0 0 0 0-1.75zm7.552 1.152l3.552-.817a2.556 2.556 0 0 1 1.797 4.757l-.027.013l-3.74 1.596l-6.196 1.549H0v-7.25h4.086l2.052-2.052a3.25 3.25 0 0 1 2.3-.948h.002h-.002h4.687a2.875 2.875 0 0 1 2.862 3.152M3.5 16.25H2v3.25h1.5z"/>`), g.Group(children),
	)
}

func UndertakeHoldUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13.952 5.71a3.357 3.357 0 1 0-4.747 4.747l.793.793h3.127a2.875 2.875 0 0 1 2.862 3.152l1.124-.259l3.686-3.685A3.357 3.357 0 0 0 16.05 5.71L15 6.758zm6.605 7.816l.123.02a2.558 2.558 0 0 1 1.688 3.854c-.25.404-.607.73-1.032.942l-.027.014l-3.74 1.595l-6.196 1.549H0v-7.25h4.086l2.052-2.052a3.25 3.25 0 0 1 1.251-.776a5.358 5.358 0 0 1 7.613-7.46a5.357 5.357 0 0 1 7.21 7.91zM3.5 16.25H2v3.25h1.5zm2 3.25h5.627l5.804-1.45l3.531-1.508a.555.555 0 0 0-.416-1.022l-.02.005L13.614 17H10v-2h3.125a.874.874 0 1 0 0-1.75h-4.69a1.25 1.25 0 0 0-.885.364l-2.05 2.05z"/>`), g.Group(children),
	)
}

func UndertakeTransaction(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M14.5 3a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5M10 5.5a4.5 4.5 0 0 1 6.5-4.032a4.5 4.5 0 1 1 0 8.064A4.5 4.5 0 0 1 10 5.5m8.25 2.488q.123.012.25.012a2.5 2.5 0 1 0-.25-4.988A4.5 4.5 0 0 1 19 5.5a4.5 4.5 0 0 1-.75 2.488M8.435 13.25a1.25 1.25 0 0 0-.885.364l-2.05 2.05V19.5h5.627l5.803-1.45l3.532-1.508a.555.555 0 0 0-.416-1.022l-.02.005L13.614 17H10v-2h3.125a.875.875 0 1 0 0-1.75zm7.552 1.152l3.552-.817a2.56 2.56 0 0 1 3.211 2.47a2.56 2.56 0 0 1-1.414 2.287l-.027.014l-3.74 1.595l-6.196 1.549H0v-7.25h4.086l2.052-2.052a3.25 3.25 0 0 1 2.3-.948h.002h-.002h4.687a2.875 2.875 0 0 1 2.862 3.152M3.5 16.25H2v3.25h1.5z"/>`), g.Group(children),
	)
}

func UnfoldLess(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m8 3.586l4 4l4-4L17.414 5L12 10.414L6.586 5zm4 10L17.414 19L16 20.414l-4-4l-4 4L6.586 19z"/>`), g.Group(children),
	)
}

func UnfoldMore(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3.586L17.414 9L16 10.414l-4-4l-4 4L6.586 9zm-4 10l4 4l4-4L17.414 15L12 20.414L6.586 15z"/>`), g.Group(children),
	)
}

func Unhappy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m8-4v4H7V8zm8 0v4h-2V8zm-9.33 7.5A5 5 0 0 1 12 13a5 5 0 0 1 4.33 2.5l.501.865l-1.731 1.001l-.5-.865A3 3 0 0 0 12 15a3 3 0 0 0-2.6 1.5l-.5.866l-1.731-1.001z"/>`), g.Group(children),
	)
}

func UnhappyOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m5-3h4v2H6zm8 0h4v2h-4zm-6.33 6.5A5 5 0 0 1 12 13a5 5 0 0 1 4.33 2.5l.501.865l-1.731 1.001l-.5-.865A3 3 0 0 0 12 15a3 3 0 0 0-2.6 1.5l-.5.866l-1.731-1.001z"/>`), g.Group(children),
	)
}

func Uninstall(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 1.594l4.914 4.858l-1.406 1.422L13 5.394v7.11h-2v-7.11l-2.508 2.48l-1.406-1.422zM2 2h5.5v2H4v10h16V4h-3.5V2H22v20H2zm18 14H4v4h16zm-14.002.998h2.004v2.004H5.998zm3 0h2.004v2.004H8.998z"/>`), g.Group(children),
	)
}

func Upload(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2.586L17.914 8.5L16.5 9.914l-3.5-3.5V16h-2V6.414l-3.5 3.5L6.086 8.5zM4.5 14v5h15v-5h2v7h-19v-7z"/>`), g.Group(children),
	)
}

func UploadOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 2.586L18.414 9L17 10.414l-4-4V16h-2V6.414l-4 4L5.586 9zM3 18h18v2H3z"/>`), g.Group(children),
	)
}

func Upscale(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m15.056 1.994l6.91.04l.04 6.91l-2 .011l-.02-3.527l-5.027 5.028l-1.415-1.415l5.027-5.027l-3.526-.02zM2 2h10v2H4v6H2zm0 10h4v2H4v2H2zm6 0h4v4h-2v-2H8zm14 0v10h-8v-2h6v-8zM4 18v2h2v2H2v-4zm8 0v4H8v-2h2v-2z"/>`), g.Group(children),
	)
}

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M4 2h16v9h2v11H2V11h2zm0 11v7h16v-7zm14-2V4H6v7zM8 6.496h2.004V8.5H8zm6 0h2.004V8.5H14z"/>`), g.Group(children),
	)
}

func User(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6.5 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M3 19a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5v3H3zm5-3a3 3 0 0 0-3 3v1h14v-1a3 3 0 0 0-3-3z"/>`), g.Group(children),
	)
}

func UserAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M19 13v4h4v2h-4v4h-2v-4h-4v-2h4v-4zM8 16a4 4 0 0 0-4 4h8.05v2H2v-2a6 6 0 0 1 6-6h4v2z"/>`), g.Group(children),
	)
}

func UserArrowDown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M19 12v7.11l2.508-2.48l1.406 1.422L18 22.91l-4.914-4.858l1.406-1.422L17 19.11V12zM8 16a4 4 0 0 0-4 4h8.8v2H2v-2a6 6 0 0 1 6-6h4.75v2z"/>`), g.Group(children),
	)
}

func UserArrowLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m12.872 6.994l-2.48 2.508h7.11v2h-7.11l2.48 2.508l-1.422 1.406l-4.858-4.914l4.858-4.914zM8 16a4 4 0 0 0-4 4h8.05v2H2v-2a6 6 0 0 1 6-6h4v2z"/>`), g.Group(children),
	)
}

func UserArrowRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m13.05 5.588l4.858 4.914l-4.858 4.914l-1.422-1.406l2.48-2.508h-7.11v-2h7.11l-2.48-2.508zM8 16a4 4 0 0 0-4 4h8.05v2H2v-2a6 6 0 0 1 6-6h4v2z"/>`), g.Group(children),
	)
}

func UserArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m12 5.094l4.914 4.858l-1.406 1.422L19 16.394v7.11h-2v-7.11l-2.508 2.48l-1.406-1.422zM8 16a4 4 0 0 0-4 4h8.8v2H2v-2a6 6 0 0 1 6-6h4.75v2z"/>`), g.Group(children),
	)
}

func UserAvatar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm17.5 18h.5V4H4v16h.5a5 5 0 0 1 5-5h5a5 5 0 0 1 5 5M12 7a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5M7.5 9.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0m2 7.5a3 3 0 0 0-3 3h11a3 3 0 0 0-3-3z"/>`), g.Group(children),
	)
}

func UserBlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m12 7a3.5 3.5 0 0 0-3.08 5.165l4.745-4.744A3.5 3.5 0 0 0 18 14.5m3.08 1.835l-4.745 4.744a3.5 3.5 0 0 0 4.745-4.745M12.5 18a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M8 16a4 4 0 0 0-4 4h7.05v2H2v-2a6 6 0 0 1 6-6h3v2z"/>`), g.Group(children),
	)
}

func UserBusiness(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6.5 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M3 19a5 5 0 0 1 5-5h8a5 5 0 0 1 5 5v3H3zm5-3a3 3 0 0 0-3 3v1h5.613l-1.334-4zm3.387 0L12 17.838L12.613 16zm3.334 0l-1.334 4H19v-1a3 3 0 0 0-3-3z"/>`), g.Group(children),
	)
}

func UserChecked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m17.657 7.59l-7.071 7.071l-4.243-4.243l1.414-1.414l2.829 2.829l5.657-5.657zM8 16a4 4 0 0 0-4 4h8.05v2H2v-2a6 6 0 0 1 6-6h4v2z"/>`), g.Group(children),
	)
}

func UserCheckedOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.5 4a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0m20.096 1.44l-5.657 5.656l-3.535-3.535l1.414-1.415l2.121 2.122l4.243-4.243zM0 19a5 5 0 0 1 5-5h7a5 5 0 0 1 5 5v2h-2v-2a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v2H0z"/>`), g.Group(children),
	)
}

func UserCircle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 0 0-9 9a8.96 8.96 0 0 0 1.773 5.365A5 5 0 0 1 9.5 14h5a5 5 0 0 1 4.727 3.365A8.96 8.96 0 0 0 21 12a9 9 0 0 0-9-9m5.5 16.125V19a3 3 0 0 0-3-3h-5a3 3 0 0 0-3 3v.125A8.96 8.96 0 0 0 12 21c2.072 0 3.979-.7 5.5-1.875M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11a10.98 10.98 0 0 1-3.85 8.36A10.96 10.96 0 0 1 12 23a10.96 10.96 0 0 1-7.15-2.64A10.98 10.98 0 0 1 1 12m11-6a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5M7.5 8.5a4.5 4.5 0 1 1 9 0a4.5 4.5 0 0 1-9 0"/>`), g.Group(children),
	)
}

func UserClear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m9.172 6.257L18 16.586l2.828-2.829l1.415 1.415L19.414 18l2.829 2.828l-1.415 1.415L18 19.414l-2.828 2.829l-1.415-1.415L16.586 18l-2.829-2.828zM8 16a4 4 0 0 0-4 4h8.05v2H2v-2a6 6 0 0 1 6-6h4v2z"/>`), g.Group(children),
	)
}

func UserErrorOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.5 4a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0m13.879-.536L19.5 8.586l2.121-2.122l1.415 1.415L20.914 10l2.121 2.121l-1.414 1.415l-2.121-2.122l-2.121 2.122l-1.415-1.415L18.087 10l-2.121-2.121zM0 19a5 5 0 0 1 5-5h7a5 5 0 0 1 5 5v2h-2v-2a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v2H0z"/>`), g.Group(children),
	)
}

func UserInvisible(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 2a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11M8 7.5a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0M4 20a4 4 0 0 1 4-4h3v-2H8a6 6 0 0 0-6 6v2h9.05v-2zm19.59-2s-1.672 4.5-6.09 4.5a5.84 5.84 0 0 1-2.528-.557l-.972.971l-1.414-1.414l.717-.718C11.971 19.512 11.41 18 11.41 18s1.673-4.5 6.09-4.5c.972 0 1.81.218 2.528.558l.972-.972l1.414 1.414l-.718.718C23.028 16.49 23.59 18 23.59 18m-3.303-1.372l-3.753 3.753c.297.075.619.119.964.119c2.616 0 3.87-2.5 3.87-2.5s-.354-.71-1.081-1.372m-1.82-1.008a4 4 0 0 0-.969-.12c-2.62 0-3.87 2.5-3.87 2.5s.357.71 1.085 1.373z"/>`), g.Group(children),
	)
}

func UserList(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M8 16a4 4 0 0 0-4 4h7.8v2H2v-2a6 6 0 0 1 6-6h3.75v2zm5.75-2h8.5v2h-8.5zm0 3h8.5v2h-8.5zm0 3h8.5v2h-8.5z"/>`), g.Group(children),
	)
}

func UserLocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M18 14c.69 0 1.25.56 1.25 1.25V16h-2.5v-.75c0-.69.56-1.25 1.25-1.25m3.25 2v-.75a3.25 3.25 0 0 0-6.5 0V16h-1.251v6.5h9V16zm-.751 2v2.5h-5V18zM8 16a4 4 0 0 0-4 4h7.55v2H2v-2a6 6 0 0 1 6-6h3.5v2z"/>`), g.Group(children),
	)
}

func UserMarked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m8.75 6h8.5v10.295l-4.247-2.617l-4.253 2.614zm2 2v4.715l2.254-1.385l2.246 1.383V15.5zM8 16a4 4 0 0 0-4 4h8.8v2H2v-2a6 6 0 0 1 6-6h4.75v2z"/>`), g.Group(children),
	)
}

func UserOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 4a3 3 0 1 0 0 6a3 3 0 0 0 0-6M7 7a5 5 0 1 1 10 0A5 5 0 0 1 7 7M3.5 19a5 5 0 0 1 5-5h7a5 5 0 0 1 5 5v2h-2v-2a3 3 0 0 0-3-3h-7a3 3 0 0 0-3 3v2h-2z"/>`), g.Group(children),
	)
}

func UserPassword(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M17 12h4v2h-2v2.003h1.953L22.5 17.14v5.36h-9v-5.36l1.547-1.137H17zm-9 4a4 4 0 0 0-4 4h7.55v2H2v-2a6 6 0 0 1 6-6h3.5v2zm7.5 2.152V20.5h5v-2.348l-.203-.15h-4.594z"/>`), g.Group(children),
	)
}

func UserSafety(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m10.5 8v3.634a1 1 0 0 0 .453.837L19 21.308l2.047-1.337a1 1 0 0 0 .453-.837V15.5zm-2-2h9v5.634a3 3 0 0 1-1.36 2.511L19 23.697l-3.14-2.052a3 3 0 0 1-1.36-2.511zM8 16a4 4 0 0 0-4 4h8.55v2H2v-2a6 6 0 0 1 6-6h4.5v2z"/>`), g.Group(children),
	)
}

func UserSearch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M17.75 15a2.75 2.75 0 1 0 0 5.5a2.75 2.75 0 0 0 0-5.5M13 17.75a4.75 4.75 0 1 1 8.74 2.578l1.674 1.671l-1.413 1.415l-1.675-1.673A4.75 4.75 0 0 1 13 17.75M8 16a4 4 0 0 0-4 4h7.55v2H2v-2a6 6 0 0 1 6-6h3.5v2z"/>`), g.Group(children),
	)
}

func UserSetting(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M8 16a4 4 0 0 0-4 4h8.05v2H2v-2a6 6 0 0 1 6-6h4v2zm11.5-3.25v1.376c.715.184 1.352.56 1.854 1.072l1.193-.689l1 1.732l-1.192.688a4 4 0 0 1 0 2.142l1.192.688l-1 1.732l-1.193-.689a4 4 0 0 1-1.854 1.072v1.376h-2v-1.376a4 4 0 0 1-1.854-1.072l-1.193.689l-1-1.732l1.192-.688a4 4 0 0 1 0-2.142l-1.192-.688l1-1.732l1.193.688a4 4 0 0 1 1.854-1.071V12.75zm-2.751 4.283a2 2 0 0 0-.25.967c0 .35.091.68.25.967l.036.063a2 2 0 0 0 3.43 0l.036-.063c.159-.287.249-.616.249-.967c0-.35-.09-.68-.249-.967l-.036-.063a2 2 0 0 0-3.43 0z"/>`), g.Group(children),
	)
}

func UserTalk(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m12.983 5.482l.603.798a7 7 0 0 1-.003 8.44l-.603.798l-1.595-1.206l.603-.798a5 5 0 0 0 .002-6.029l-.603-.798zM8 16a4 4 0 0 0-4 4h9.05v2H2v-2a6 6 0 0 1 6-6h5v2zm8.19-.91l.603.799a3.5 3.5 0 0 1-.001 4.22l-.604.798l-1.595-1.206l.604-.798a1.5 1.5 0 0 0 0-1.809l-.602-.798z"/>`), g.Group(children),
	)
}

func UserTalkOffOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1.586L22.414 21L21 22.414l-4.039-4.038Q17 18.683 17 19v2h-2v-2a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v2H0v-2a5 5 0 0 1 5-5h7q.317 0 .624.039l-2.359-2.36a5 5 0 0 1-6.444-6.444L1.586 3zm2.501 5.33L5.5 7a3 3 0 0 0 3.085 3zM8.5 4q-.182 0-.356.02l-.993.118l-.234-1.987l.993-.116Q8.2 2 8.5 2a5 5 0 0 1 4.965 5.59l-.117.993l-1.986-.234l.117-.993q.02-.175.02-.356a3 3 0 0 0-3-3m12.075 1.648l.498.867a7 7 0 0 1-.002 6.975l-.499.867l-1.734-.997l.499-.867a5 5 0 0 0 .002-4.982l-.498-.867zM17.54 7.391l.498.867a3.5 3.5 0 0 1-.001 3.487l-.499.867l-1.733-.997l.498-.867a1.5 1.5 0 0 0 0-1.494l-.497-.868z"/>`), g.Group(children),
	)
}

func UserTalkOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M8.5 4a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0m17.073-1.352l.497.867a7 7 0 0 1-.002 6.975l-.499.867l-1.733-.997l.498-.867a5 5 0 0 0 .002-4.982l-.498-.867zM17.538 7.39l.497.868a3.5 3.5 0 0 1 0 3.487l-.5.867l-1.733-.997l.498-.867a1.5 1.5 0 0 0 0-1.495l-.497-.867zM0 19a5 5 0 0 1 5-5h7a5 5 0 0 1 5 5v2h-2v-2a3 3 0 0 0-3-3H5a3 3 0 0 0-3 3v2H0z"/>`), g.Group(children),
	)
}

func UserTime(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m12 7a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M12.5 18a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M8 16a4 4 0 0 0-4 4h7.05v2H2v-2a6 6 0 0 1 6-6h3v2zm11-.248v1.834L20.414 19L19 20.414l-2-2v-2.662z"/>`), g.Group(children),
	)
}

func UserTransmit(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m12.578 5.596l3.674 2.88v1.526H13.75v-2h4.655l-1.061-.832zM8 16a4 4 0 0 0-4 4h7.8v2H2v-2a6 6 0 0 1 6-6h3.75v2zm5.752 2.498h8.502v2H17.6l1.061.832l-1.234 1.574l-3.674-2.88z"/>`), g.Group(children),
	)
}

func UserUnknown(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M18 14c-.93 0-1.5.656-1.5 1.249v1h-2v-1C14.5 13.358 16.17 12 18 12s3.5 1.358 3.5 3.249a3.13 3.13 0 0 1-1.027 2.3L19 18.939v.683h-2v-1.546l2.112-1.993c.256-.235.388-.53.388-.834c0-.593-.57-1.249-1.5-1.249M8 16a4 4 0 0 0-4 4h8.8v2H2v-2a6 6 0 0 1 6-6h4.75v2zm9 4.996h2.003V23h-2.004z"/>`), g.Group(children),
	)
}

func UserUnlocked(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M18 14c-.69 0-1.25.56-1.25 1.25V16h5.749v6.5h-9V16h1.251v-.75a3.25 3.25 0 0 1 5.541-2.305l.71.705l-1.41 1.418l-.71-.705A1.24 1.24 0 0 0 18 14m-2.501 4v2.5h5V18zM8 16a4 4 0 0 0-4 4h7.55v2H2v-2a6 6 0 0 1 6-6h3.5v2z"/>`), g.Group(children),
	)
}

func UserVip(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m8.382 6h7.236l2.081 4.162L18 23.995l-5.7-6.333zm1.236 2l-.919 1.838L18 21.005l3.3-3.667l-.918-1.838zM8 16a4 4 0 0 0-4 4h7.05v2H2v-2a6 6 0 0 1 6-6h3v2z"/>`), g.Group(children),
	)
}

func UserVisible(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M11.5 4a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M6 7.5a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0M8 16a4 4 0 0 0-4 4h7.05v2H2v-2a6 6 0 0 1 6-6h3v2zm8.248 3v-2h2.5v2z"/><path fill="currentColor" d="M17.5 22.5c4.418 0 6.09-4.5 6.09-4.5s-1.674-4.5-6.09-4.5s-6.09 4.5-6.09 4.5s1.671 4.5 6.09 4.5m-.002-2c-2.616 0-3.87-2.5-3.87-2.5s1.25-2.5 3.87-2.5s3.87 2.5 3.87 2.5s-1.254 2.5-3.87 2.5"/>`), g.Group(children),
	)
}

func Usergroup(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7 5a3 3 0 0 0 0 6v2a5 5 0 0 0-5 5v4H0v-4a7 7 0 0 1 3.75-6.2A5 5 0 0 1 7 3h1v2a5 5 0 0 1 4-2a4.99 4.99 0 0 1 4 2V3h1a5 5 0 0 1 3.25 8.8A7 7 0 0 1 24 18v4h-2v-4a5 5 0 0 0-5-5v-2a3 3 0 1 0 0-6h-1a5 5 0 1 1-8 0zm5 0a3 3 0 1 0 0 6a3 3 0 0 0 0-6M4 19a5 5 0 0 1 5-5h6a5 5 0 0 1 5 5v3H4zm5-3a3 3 0 0 0-3 3v1h12v-1a3 3 0 0 0-3-3z"/>`), g.Group(children),
	)
}

func UsergroupAdd(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 5a3 3 0 0 0 0 6v2a5 5 0 0 0-5 5v4H2v-4a7 7 0 0 1 3.75-6.2A5 5 0 0 1 9 3h1v2zm6 0a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0M7 19a5 5 0 0 1 5-5h3v2h-3a3 3 0 0 0-3 3v1h6v2H7zm14-5v3h3v2h-3v3h-2v-3h-3v-2h3v-3z"/>`), g.Group(children),
	)
}

func UsergroupClear(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M9 5a3 3 0 0 0 0 6v2a5 5 0 0 0-5 5v4H2v-4a7 7 0 0 1 3.75-6.2A5 5 0 0 1 9 3h1v2zm6 0a3 3 0 1 0 0 6a3 3 0 0 0 0-6m-5 3a5 5 0 1 1 10 0a5 5 0 0 1-10 0M7 19a5 5 0 0 1 5-5h3v2h-3a3 3 0 0 0-3 3v1h6v2H7zm10.879-4.536L20 16.586l2.121-2.122l1.415 1.415L21.414 18l2.121 2.121l-1.414 1.415L20 19.414l-2.121 2.121l-1.415-1.414L18.587 18l-2.121-2.121z"/>`), g.Group(children),
	)
}

func Vehicle(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 5a2 2 0 0 1 2-2h11a2 2 0 0 1 2 2v2h-2V5H2v7h6v2H2v4h6v2H5.414L3.5 21.914L2.086 20.5l.5-.5H2a2 2 0 0 1-2-2zm11.323 3h10.354L24 13.807V21.5h-2V20H11v1.5H9v-7.693zM11 18h11v-3.807L21.923 14H11.077l-.077.193zm.877-6h9.246l-.8-2h-7.646zM3 15h2.004v2.004H3zm9 0h2.004v2.004H12zm7 0h2.004v2.004H19z"/>`), g.Group(children),
	)
}

func Verified(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m12 .186l3.617 3.082l4.737.378l.378 4.737L23.814 12l-3.082 3.617l-.378 4.737l-4.737.378L12 23.814l-3.616-3.082l-4.737-.378l-.378-4.737L.187 12l3.082-3.617l.378-4.737l4.737-.378zm0 2.628L9.188 5.21l-3.683.293l-.294 3.684L2.814 12l2.397 2.812l.294 3.684l3.683.294L12 21.186l2.813-2.396l3.683-.294l.294-3.684L21.186 12L18.79 9.188l-.294-3.684l-3.683-.293zM17.915 9.5L11 16.414L6.586 12L8 10.586l3 3l5.5-5.5z"/>`), g.Group(children),
	)
}

func Verify(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v14h18V5zm12.5 4.5a1 1 0 1 1 0 2a1 1 0 0 1 0-2m2.4 2.8a3 3 0 1 0-4.8 0a4 4 0 0 0-1.6 3.2v1h2v-1a2 2 0 1 1 4 0v1h2v-1c0-1.309-.628-2.47-1.6-3.2M5 9h4.5v2H5zm0 4h4.5v2H5z"/>`), g.Group(children),
	)
}

func Video(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h16V4zm4 2.37L17.75 12L8 17.63zm2 3.465v4.33L13.75 12z"/>`), g.Group(children),
	)
}

func VideoCamera(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 4h17v4.434l7-4.2v15.532l-7-4.2V20H0zm15 2H2v12h13zm2 7.234l5 3V7.766l-5 3z"/>`), g.Group(children),
	)
}

func VideoCameraDollar(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 3.998h17v4.434l7-4.2v15.49l-7-4v4.276H0zm17 9.42l5 2.857v-8.51l-5 3zM2 5.998v12h13v-12zM9.5 7v1h2v2h-4v1h2a2 2 0 0 1 2 2v1a2 2 0 0 1-2 2v1h-2v-1h-2v-2h4v-1h-2a2 2 0 0 1-2-2v-1a2 2 0 0 1 2-2V7z"/>`), g.Group(children),
	)
}

func VideoCameraMinus(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M24 4.234v15.49l-7-4V20H0V4h17v4.434zm-7 6.532v2.654l5 2.857v-8.51zM15 18V6H2v12zm-2.5-5h-8v-2h8z"/>`), g.Group(children),
	)
}

func VideoCameraMusic(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M24 4.234v15.49l-7-4V20H0V4h17v4.434zm-7 6.532v2.654l5 2.857v-8.51zM15 6H2v12h13zm-2.5 3h-2v5a3 3 0 1 1-2-2.83V7h4zm-4 5a1 1 0 1 0-2 0a1 1 0 0 0 2 0"/>`), g.Group(children),
	)
}

func VideoCameraOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1.586L22.414 21L21 22.414l-4-4V20H0V4h2.586l-1-1zM4.586 6H2v12h13v-1.586zm2.996-2H17v4.433l7-4.2V16l.003 3.416L22 17.415V7.766l-5 3l.004 2.655L15 11.415V6.001l-5.413.001z"/>`), g.Group(children),
	)
}

func VideoCameraOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M0 4h17v4.434l7-4.2v15.49l-7-4V20H0zm17 9.42l5 2.857v-8.51l-5 3zM2 6v12h13V6zm2 2h6v2H4z"/>`), g.Group(children),
	)
}

func VideoCameraTwo(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12.5 3a3.5 3.5 0 1 1 0 7a3.5 3.5 0 0 1 0-7m4.243 7a5.5 5.5 0 1 0-9.45-5.278A4 4 0 0 0 1.536 10H1v12h17v-2.523l5 2V10.523l-5 2V10zM18 14.677l3-1.2v5.046l-3-1.2zM16 12v8H3v-8zM5 10a2 2 0 1 1 0-4a2 2 0 0 1 0 4"/>`), g.Group(children),
	)
}

func VideoLibrary(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 2h12v2H6zM4 5h16v2H4zM1.839 8H22.16l-2.1 14H3.94zm2.322 2l1.5 10H18.34l1.5-10zm6.34 1.5l4.666 3.5l-4.667 3.5z"/>`), g.Group(children),
	)
}

func ViewAgenda(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 3h20v8H2zm2 2v4h16V5zm-2 8h20v8H2zm2 2v4h16v-4z"/>`), g.Group(children),
	)
}

func ViewColumn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6 2v20H4V2zm7 0v20h-2V2zm7 0v20h-2V2z"/>`), g.Group(children),
	)
}

func ViewInAr(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1.5 1.5h5v2h-3v3h-2zm16 0h5v5h-2v-3h-3zM12 3.845l7.061 4.077v8.155L12 20.155l-7.062-4.078V7.922zm-5.062 6.386v4.692L11 17.268v-4.691zM13 17.268l4.062-2.345v-4.691L13 12.577zm-1-6.423L16.061 8.5L12 6.155L7.938 8.499zM3.5 17.5v3h3v2h-5v-5zm19 0v5h-5v-2h3v-3z"/>`), g.Group(children),
	)
}

func ViewList(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 4h20v2H2zm0 7h20v2H2zm1 7H2v2h20v-2z"/>`), g.Group(children),
	)
}

func ViewModule(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h22v18H1zm2 2v6h4.667V5zm6.667 0v6h4.666V5zm6.666 0v6H21V5zM21 13h-4.667v6H21zm-6.667 6v-6H9.667v6zm-6.666 0v-6H3v6z"/>`), g.Group(children),
	)
}

func VisualRecognition(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v10h-2V4H4v9.586l5-5l4.914 4.914l-1.414 1.414l-3.5-3.5l-5 5V20h13v2H2zm13.547 5a1 1 0 1 0 0 2a1 1 0 0 0 0-2m-3 1a3 3 0 1 1 6 0a3 3 0 0 1-6 0M16 14h8v2h-3v7h-2v-7h-3z"/>`), g.Group(children),
	)
}

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M1 3h15a3 3 0 0 1 3 3v1h4v14H1zm2 6v10h18V9zm0-2h14V6a1 1 0 0 0-1-1H3zm13 6h3v2h-3z"/>`), g.Group(children),
	)
}

func Watch(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.117 0h9.766l.503 4.025A3 3 0 0 1 20 7v2h1v6h-1v2a3 3 0 0 1-2.614 2.975L16.883 24H7.117l-.503-4.025A3 3 0 0 1 4 17V7a3 3 0 0 1 2.614-2.975zm1.516 4h6.734l-.25-2H8.883zM7 18h10a1 1 0 0 0 1-1V7a1 1 0 0 0-1-1H7a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1m1.633 2l.25 2h6.234l.25-2z"/>`), g.Group(children),
	)
}

func Watermelon(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M17.995.59L1.59 16.996l.696.707c1.04 1.056 2.523 2.117 4.126 2.918C8.007 21.416 9.81 22 11.488 22C17.846 22 23 16.846 23 10.488c0-1.678-.584-3.481-1.38-5.075c-.8-1.603-1.862-3.085-2.918-4.126zm2.946 9.07q.059.428.059.828A9.51 9.51 0 0 1 11.488 20q-.4 0-.827-.06A11.516 11.516 0 0 0 20.94 9.66m-3.037-6.15c.463.788.745 1.765.906 2.707a13.6 13.6 0 0 1 .188 2.094v.133l.001.05v.021A9.51 9.51 0 0 1 9.489 18H9.43q-.033.002-.104.001c-.094 0-.236-.003-.417-.012a12 12 0 0 1-1.468-.173c-.89-.161-1.915-.455-2.867-.975zm-.902 4.488h-2.004v2.004h2.004zm-3 3h-2.004v2.004h2.004zm-3 3H8.998v2.004h2.004z"/>`), g.Group(children),
	)
}

func WaveLeft(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m7.156 3.192l-.823.568A10.1 10.1 0 0 0 3.76 6.332l-.567.824L1.546 6.02l.567-.823a12.1 12.1 0 0 1 3.085-3.084l.823-.568zm4.3 3.685a.388.388 0 0 0-.548.548l5.237 5.237l-1.415 1.414l-6.603-6.603a.388.388 0 0 0-.548.548l6.603 6.604l-1.414 1.414l-5.237-5.237a.388.388 0 0 0-.548.549l2.084 2.083v.001l3.153 3.152l-1.415 1.415l-3.151-3.152a.386.386 0 0 0-.546.545l4.32 4.32a5.19 5.19 0 0 0 7.34 0l.844-.844a5.19 5.19 0 0 0 .812-6.286l-2.227-3.813a.145.145 0 0 0-.266.111l.669 2.496a1.01 1.01 0 0 1-.276.986a1 1 0 0 1-1.398-.017zM5.99 13.185l-.42-.42a2.388 2.388 0 0 1 .326-3.65A2.388 2.388 0 0 1 9.22 5.79q.12-.173.274-.327a2.39 2.39 0 0 1 3.376 0l3.078 3.078a2.145 2.145 0 0 1 3.976-.777l2.227 3.812a7.19 7.19 0 0 1-1.125 8.71l-.844.844a7.19 7.19 0 0 1-10.168 0l-4.32-4.32a2.386 2.386 0 0 1 .295-3.625"/>`), g.Group(children),
	)
}

func WaveRight(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m19.047.374l.824.567a12.1 12.1 0 0 1 3.084 3.084l.568.824l-1.647 1.135l-.568-.824a10.1 10.1 0 0 0-2.572-2.572l-.824-.568zM12.198 4.29a2.388 2.388 0 0 1 3.65.327a2.388 2.388 0 0 1 3.325 3.325q.173.12.327.273a2.39 2.39 0 0 1 0 3.377l-.42.42q.155.111.295.25a2.386 2.386 0 0 1 0 3.374l-4.32 4.32a7.19 7.19 0 0 1-10.168 0l-.845-.844a7.19 7.19 0 0 1-1.124-8.709l2.226-3.812a2.146 2.146 0 0 1 3.977.777zm5.888 5.888a.388.388 0 0 0-.548-.549L12.3 14.867l-1.415-1.414l6.604-6.604a.388.388 0 0 0-.548-.548l-6.604 6.604l-1.414-1.414l5.237-5.237a.388.388 0 0 0-.549-.548l-5.468 5.468a1.001 1.001 0 0 1-1.675-.967l.668-2.495A.146.146 0 0 0 6.87 7.6l-2.226 3.813a5.19 5.19 0 0 0 .812 6.286l.844.845a5.19 5.19 0 0 0 7.339 0l4.32-4.32a.386.386 0 0 0-.545-.546l-3.152 3.152l-1.414-1.415z"/>`), g.Group(children),
	)
}

func Wealth(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M6.75 2.5h10.5v2.75c0 1.57-.69 2.98-1.783 3.942a9.03 9.03 0 0 1 4.41 3.951l.485.875l-1.75.97l-.484-.875A7 7 0 0 0 5 17.5v.5a2 2 0 0 0 2 2h7v2H7a4 4 0 0 1-4-4v-.5a9 9 0 0 1 5.533-8.308A5.24 5.24 0 0 1 6.75 5.25zm5.25 6a3.25 3.25 0 0 0 3.25-3.25V4.5h-6.5v.75A3.25 3.25 0 0 0 12 8.5m4 7.5h7v2h-7zm0 4h7v2h-7z"/>`), g.Group(children),
	)
}

func WealthOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5 7.967v3.53c.004.018.032.098.184.236c.205.187.558.401 1.079.601c1.035.399 2.533.666 4.237.666s3.202-.267 4.237-.666c.521-.2.874-.414 1.08-.6c.151-.139.179-.219.183-.237v-3.53a8 8 0 0 1-.545.234C14.138 8.708 12.385 9 10.5 9s-3.638-.292-4.955-.799A8 8 0 0 1 5 7.967M18 5.5v6.522a5.5 5.5 0 1 1-4.897 8.783c-.81.126-1.689.195-2.603.195c-1.886 0-3.638-.292-4.955-.799c-.656-.252-1.254-.577-1.707-.988S3 18.223 3 17.5v-12c0-.724.385-1.301.838-1.713c.453-.411 1.051-.736 1.707-.988C6.862 2.292 8.615 2 10.5 2s3.638.292 4.955.799c.655.252 1.254.577 1.707.988S18 4.777 18 5.5m-13 0c.007.023.04.102.184.233c.205.187.558.401 1.079.601C7.298 6.733 8.796 7 10.5 7s3.202-.267 4.237-.666c.521-.2.874-.414 1.08-.6c.144-.132.176-.21.182-.234c-.006-.023-.038-.102-.183-.233c-.205-.187-.558-.401-1.079-.601C13.702 4.267 12.204 4 10.5 4s-3.202.267-4.237.666c-.521.2-.874.414-1.08.6c-.144.132-.176.21-.182.234m0 8.467v3.53c.004.018.032.098.184.236c.205.187.558.401 1.079.601c1.035.399 2.533.666 4.237.666c.587 0 1.152-.032 1.683-.09A5.5 5.5 0 0 1 12 17.5c0-.954.243-1.852.67-2.634c-.69.088-1.419.134-2.17.134c-1.886 0-3.638-.292-4.955-.799A8 8 0 0 1 5 13.967M17.5 14q-.198 0-.39.021A3.5 3.5 0 1 0 17.5 14"/>`), g.Group(children),
	)
}

func Widget(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M5.35 3h13.3L23 12.788V22H1v-9.212zm1.3 2l-3.111 7H20.46l-3.11-7zM21 14H3v6h18zM5 16h2.004v2.004H5zm4 0h2.004v2.004H9z"/>`), g.Group(children),
	)
}

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M21.193 9.808c-5.077-5.077-13.308-5.077-18.385 0l-.707.707L.687 9.1l.707-.707c5.858-5.857 15.355-5.857 21.213 0l.707.707l-1.414 1.415zm-4.597 4.596a6.5 6.5 0 0 0-9.192 0l-.707.707l-1.414-1.414l.707-.707a8.5 8.5 0 0 1 12.02 0l.708.707l-1.415 1.414zm-6.01 3.182a2 2 0 0 1 2.829 0l.707.707L12 20.414l-2.12-2.121z"/>`), g.Group(children),
	)
}

func WifiOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1.586L22.414 21L21 22.414l-7.304-7.304l-2.554-2.554a6.47 6.47 0 0 0-3.739 1.847l-.707.707l-1.414-1.414l.707-.707a8.5 8.5 0 0 1 3.48-2.106l-.37-.37l-3.067-3.067c-1.16.6-2.25 1.386-3.225 2.36l-.707.708L.686 9.1l.707-.707A15 15 0 0 1 4.557 5.97L1.586 3zm18.192 8.222a12.98 12.98 0 0 0-10.934-3.692l-.99.134L9 4.267l.991-.133a14.98 14.98 0 0 1 12.614 4.26l.707.706l-1.414 1.415zm-10.606 7.778a2 2 0 0 1 2.828 0l.707.707l-2.12 2.121l-2.122-2.121z"/>`), g.Group(children),
	)
}

func WifiOffOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3 1.586L22.414 21l-1.415 1.414l-5.361-5.361l-3.639 4.548L.6 7.351l.779-.626c.747-.6 1.534-1.127 2.35-1.582L1.586 3zm2.207 5.036q-.917.465-1.778 1.064l8.57 10.713l2.216-2.77zm15.362 1.064A15 15 0 0 0 9.41 5.224l-.985.172l-.343-1.97l.985-.173a17 17 0 0 1 13.555 3.472l.779.625l-5.467 6.833l-1.562-1.25z"/>`), g.Group(children),
	)
}

func WifiOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M3.43 7.686L12 18.4l8.57-10.713a15.015 15.015 0 0 0-17.14 0m-2.051-.961c6.192-4.967 15.05-4.967 21.243 0l.779.625l-11.4 14.25L.6 7.35z"/>`), g.Group(children),
	)
}

func Window(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 2h20v20H2zm2 2v16h7V4zm9 1.414v6.172l7 7v-6.172zm7 4.172V4h-5.586zM18.586 20L13 14.414V20z"/>`), g.Group(children),
	)
}

func WindowOne(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="m11 .613l11 3.666V22H2V3h9zM11 5H4v15h10.838L11 18.721zm9 14.613V13.72l-7-2.333v5.891zm0-8V5.72l-7-2.333v5.891z"/>`), g.Group(children),
	)
}

func Windy(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 5a1 1 0 1 0 0 2h9v2h-9a3 3 0 1 1 0-6h4v2zM5.5 8a1.5 1.5 0 1 0 0 3H22v2H5.5a3.5 3.5 0 1 1 0-7H8v2zM5 18a3 3 0 0 1 3-3h10v2H8a1 1 0 1 0 0 2h4.5v2H8a3 3 0 0 1-3-3"/>`), g.Group(children),
	)
}

func WindyRain(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M13 5a1 1 0 1 0 0 2h5v2h-5a3 3 0 1 1 0-6h4v2zM5.5 8a1.5 1.5 0 1 0 0 3H12v2H5.5a3.5 3.5 0 1 1 0-7H8v2zM20 6.996h2.004V9H20zm0 4h2.004V13H20zM14 11h4v2h-4zm2 3.996h2.004V17H16zM5 18a3 3 0 0 1 3-3h1v2H8a1 1 0 1 0 0 2h4.5v2H8a3 3 0 0 1-3-3m6-3h3v2h-3z"/>`), g.Group(children),
	)
}

func Wink(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m16.362-3.119L15.54 9.906L17.4 11.3l-1.2 1.6l-3.2-2.4V9.04l3.381-1.902zM9 8v4H7V8zm-.1 5.634l.5.865C9.922 15.4 10.89 16 12 16s2.08-.601 2.6-1.5l.5-.866l1.731 1.001l-.5.866A5 5 0 0 1 12 18a5 5 0 0 1-4.33-2.5l-.501-.865z"/>`), g.Group(children),
	)
}

func Work(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 2.5h9v4H22v15H2v-15h5.5zm2 4h5v-2h-5zM4 8.5v11h16v-11z"/>`), g.Group(children),
	)
}

func WorkHistory(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M7.5 2.5h9v4H22V11h-2V8.5H4v11h7v2H2v-15h5.5zm2 4h5v-2h-5zm8.5 8a3.5 3.5 0 1 0 0 7a3.5 3.5 0 0 0 0-7M12.5 18a5.5 5.5 0 1 1 11 0a5.5 5.5 0 0 1-11 0m6.5-2.248v1.834L20.414 19L19 20.414l-2-2v-2.662z"/>`), g.Group(children),
	)
}

func WorkOff(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M2 .586L23.414 22L22 23.414L20.086 21.5H2v-15h3.086L.586 2zM7.086 8.5H4v11h14.086zm.42-6H16.5v4.002L22 6.5l.003 11.418L20 15.915V8.501h-7.413l-2.003-2.003l3.916.004V4.5h-5v1.622L7.497 4.078z"/>`), g.Group(children),
	)
}

func WrySmile(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M12 3a9 9 0 1 0 0 18a9 9 0 0 0 0-18M1 12C1 5.925 5.925 1 12 1s11 4.925 11 11s-4.925 11-11 11S1 18.075 1 12m9.233-2.134l-3.464 2l-1-1.732l3.464-2zm4.536-1.732l3.464 2l-1 1.732l-3.464-2zm-5.869 5.5l.5.865C9.922 15.4 10.89 16 12 16s2.08-.601 2.6-1.5l.5-.866l1.731 1.001l-.5.866A5 5 0 0 1 12 18a5 5 0 0 1-4.33-2.5l-.501-.865z"/>`), g.Group(children),
	)
}

func ZoomIn(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.096 5.904a6.5 6.5 0 1 0-9.192 9.192a6.5 6.5 0 0 0 9.192-9.192M4.49 4.49a8.5 8.5 0 0 1 12.686 11.272l5.345 5.345l-1.414 1.414l-5.345-5.345A8.501 8.501 0 0 1 4.49 4.49M11.5 6.5v3h3v2h-3v3h-2v-3h-3v-2h3v-3z"/>`), g.Group(children),
	)
}

func ZoomOut(children ...g.Node) g.Node {
	return s.SVG(
		viewbox, hAttr,
		g.Raw(`<path fill="currentColor" d="M15.096 5.904a6.5 6.5 0 1 0-9.192 9.192a6.5 6.5 0 0 0 9.192-9.192M4.49 4.49a8.5 8.5 0 0 1 12.686 11.272l5.345 5.345l-1.414 1.414l-5.345-5.345A8.501 8.501 0 0 1 4.49 4.49M6.5 11.5v-2h8v2z"/>`), g.Group(children),
	)
}
