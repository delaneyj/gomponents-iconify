package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ozobotevo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.203 23.42c0 4.366 1.149 2.616-1.073 14.928c-3.568 4.644-12.312 3.907-18.485 4.244c-7.635.416-9.739.292-18.613-4.474c-.49-7.021-.786-11.97-.234-13.686c.667-2.076 1.22-2.761 1.559-5.151C7.63 8.914 15.686 5.256 23.79 5.256c10.792 0 19.412 7.523 19.412 18.164Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.764 14.284c5.075-3.002 12.415-4.822 28.245-2.425m-24.721-1.836a38.904 38.904 0 0 1 21.999-.728"/><ellipse cx="9.823" cy="33.585" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.731" ry="6.52"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m10.575 17.19l4.47-1.22v12.197m-1.751 10.056a125.644 125.644 0 0 0 29.092-1.313"/><circle cx="19.294" cy="15.744" r="1.395" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.346" cy="15.841" r="1.395" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.531 29.398c.018 1.613-1.924 2.58-3.025 2.58c-.962 0-3.313-.48-3.324-1.806l-.07-9.043c-.012-1.707 1.771-2.402 2.917-2.402c1.159 0 3.39.98 3.41 2.713Z"/><circle cx="36.287" cy="15.755" r="1.395" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.361" cy="23.431" r="1.395" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.215" cy="28.556" r="1.395" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.902 40.557a128.442 128.442 0 0 0 29.414-1.32m1.662-18.747c-1.555.71-.938 2.566-3.28 2.422l-.259 9.018c2.622.283.69 3.155 3.122 3.957m-32.625-2.41l1.412-3.364m-1.659 4.83l1.155 3.494m-2.192-4.648l-2.56-.886"/>`),
		g.Group(children),
	)
}