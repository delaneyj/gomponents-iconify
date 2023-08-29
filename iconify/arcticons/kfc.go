package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kfc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.746 23.277c.326-2.448 2.204-2.57 2.693-1.387s-.204 8.897-3.877 8.285c.082 3.55-1.02 7.795-3.469 8.938s-4.612 3.02-4.979 5.387c-1.837-1.388-9.142-4.572-10.57-13.162c-.776-.102-4.204-4.143-3.796-7.49s1.995-2.839 2.974-1.737"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.443 21.608c.795-1.235 1.297-3.232 1.221-4.738c-.408-8.122-6.53-12.897-12.517-13.346c-5.988-.449-8.91 5.51-11.807 4.571c2.245 3.102 3.755.612 7.02.653s4.787 2.327 4.787 2.327"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.685 9.653s-2.284 1.36-3.97.993c.42 1.408 1.625 2.55 1.625 2.55c-1.674 1.286-1.918 2.164-1.918 2.164c.102 1.918 1.918 6.01 1.918 6.01m11.047-11.115c-4.089-1.609-8.702 2.493-9.966 9.798m0 0s6.51.041 6.959.245s1.755.653 1.796 2.653s-1.102 2.938-2.327 3.265m13.958-6.009H27.89c-1.469 0-3.795.5-2.652 3.56"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.95 25.91c1.674-.245 3.53-2.365 1.98-5.702m-15.754 5.396c-2.391-1.08-3.584-3.263-1.247-5.324"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.544 31.338c-.98-3.204-1.067-5.104-.35-7.276m13.016 3.474c.598-.572 0-1.184 0-1.184"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.746 27.536c-.06 1.646 2.539 2.013 2.539 2.013m12.657 2.816c.39 1.02.158 2.789.158 2.789m-6.176.245c-.245-.721-1.468-1.23-2.258-1.261"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.312 34.623c.122-.857.204-1.537 1.156-2.163c1.592 1.415 1.864 3.7 1.592 5.034s-.62 2.522.03 4.308m5.94-19.568s1.41-.247 2.34.295m-11.571.259c1.03-.96 2.663-.381 2.663-.381m3.712 7.441c-1.116-1.605-2.293-1.459-3.307-.896c-.723.402-.965 1.685-2.407 1.55c.952 1.605 3.514 2.23 5.714.435"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.174 29.848c1.115-1.605 2.292-1.459 3.306-.896c.724.402.965 1.685 2.407 1.55c-.952 1.605-3.514 2.23-5.714.435m2.074 11.066c.704-1.074.255-2.727.255-2.727m-4.467-17.289c.794-.587 2.145-.882 2.918 0m7.97-8.028c.136 3.319 3.191 6.249 3.191 6.249M21.64 25.006c-.47.857-.807 2.719.395 3.556"/>`),
		g.Group(children),
	)
}