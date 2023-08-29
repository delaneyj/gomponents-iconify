package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Melon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b1cc33" d="M52.6 12.6a28.14 28.14 0 0 1-39.8 39.8z"/><path fill="#f4aa41" d="M49.62 15.76c8.22 8.247 8.118 25.21-.115 33.45c-8.233 8.233-24.75 8.335-32.99.115z"/><path fill="#e27022" d="M41.67 23.53a12.9 12.9 0 0 1-17.94 17.94z"/><path fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2.251" d="M52.01 12.6c.168.168.333.337.495.508c10.49 11.02 10.32 28.47-.498 39.29c-10.99 10.99-28.81 10.99-39.8.002z"/><path fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2.251" d="M49.11 15.44c8.167 8.299 9.065 22.67-.053 32.99c-8.942 10.13-25.08 8.628-33.27.357z"/><path fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2.345" d="M41.09 23.53a12.9 12.9 0 0 1-17.94 17.94z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2.251" d="M31.79 51.06c3.369.537 9.099-.913 13.52-5.585c3.186-3.371 5.996-9.423 5.673-13.39"/><circle cx="50.54" cy="28.03" r="1.173"/>`),
		g.Group(children),
	)
}