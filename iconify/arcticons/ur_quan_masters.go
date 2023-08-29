package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UrQuanMasters(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.151 8.848L36.88 6.577l-8.634 1.022l-4.901 3.87l.67 4.393l-7.529 7.529l-1.49-1.491l5.645-5.646l-4.04-4.04l-4.196 4.196l-.406-.405l-5.094 5.094v1.243s-3.17 4.313 0 7.484"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m6.906 29.826l5.828-5.828l1.52 1.52l-2.942 2.944l.687.687l-.95.95H7.606a4.302 4.302 0 0 1-.903 3.37c-1.249 1.249 1.852-.246 4.963 2.865"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M7.98 34.338s-1.277 2.004 1.2 4.482m27.676-27.676l-1.629-1.63l-5.999.425m4.759 4.074l-.682-.682h-4.759l1.294 4.829l-9.338 9.338"/><circle cx="18.914" cy="29.086" r="2.061" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.521 27.567l-1.601-1.601M7.713 24.77l9.273-9.272m6.359-4.029l3.109 1.214m-2.439 3.179l3.456 1.185m9.849-10.03l1.513-1.513m-21.753 7.19a2.06 2.06 0 1 1 2.914 2.913m19.157-6.759l2.272 2.272l-1.022 8.634l-3.87 4.901l-4.394-.67l-7.529 7.53l1.49 1.49l5.646-5.646l4.04 4.04l-4.196 4.196l.406.406l-5.093 5.093h-1.244s-4.312 3.171-7.483 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m18.174 41.094l5.828-5.828l-1.52-1.521l-2.944 2.943l-.687-.687l-.95.95v3.443a4.302 4.302 0 0 0-3.37.903c-1.25 1.25.246-1.851-2.866-4.963"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.662 40.02s-2.004 1.277-4.482-1.2m27.676-27.676l1.629 1.628l-.424 6m-4.074-4.759l.682.682v4.759L29.84 18.16l-9.338 9.338"/><circle cx="18.914" cy="29.086" r="2.061" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m20.433 30.479l1.601 1.601m1.196 8.207l9.273-9.272m4.027-6.36l-1.213-3.109m-3.179 2.439l-1.185-3.456m10.03-9.849l1.513-1.513m-7.19 21.753a2.06 2.06 0 0 0-2.913-2.914"/>`),
		g.Group(children),
	)
}