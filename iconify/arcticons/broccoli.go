package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Broccoli(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.09 14.006a4.164 4.164 0 0 1 6.382-5.332"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.142 11.302a4.481 4.481 0 0 1 7.668-4.64h0a3.688 3.688 0 0 1 5.754 4.553M17.958 22.323c2.755 1.031 4.684 3.364 5.787 6.998m0-.001c.216-2.415.962-4.178 2.239-5.288"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.657 20.712a4.132 4.132 0 0 1-8.21.713"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.108 20.236c-.725 2.087-3.029 2.422-4.348.768c-.967 1.045-2.026 1.312-3.18.8c.962 1.006.054 3.082-1.86 2.78"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.09 14.006a5.775 5.775 0 1 0 4.64 10.567c1.73 1.384 3.45 2.758 4.283 5.894c.844 3.126.79 8.025.736 12.914l5.602.119c-.194-5.008-.4-10.004.433-13.281c.833-3.267 2.693-4.813 4.564-6.36c2.61 2.545 5.163 2.84 7.658.887c3.06-2.088 1.243-6.176-1.677-7.182a4.856 4.856 0 0 0-5.289-7.257"/>`),
		g.Group(children),
	)
}