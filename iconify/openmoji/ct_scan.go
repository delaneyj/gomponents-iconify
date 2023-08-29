package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CtScan(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="M11.797 32.146h12.968a1.694 1.694 0 0 1 1.207.502a14.389 14.389 0 0 1 2.554 3.097l.015.053a1.691 1.691 0 0 1-1.655 2.106l-15.067-.008a1.701 1.701 0 0 1-1.7-1.686l-.023-2.347a1.701 1.701 0 0 1 1.701-1.717Z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M16.985 42.761a20.43 20.43 0 0 0 39.601-7.483a20.43 20.43 0 0 0-39.684-7.257s8.798.275 8.887.167c3.72-4.512 6.023-5.59 10.318-5.59a12.682 12.682 0 0 1 0 25.36a12.364 12.364 0 0 1-10.198-5.412m-8.577-.134l7.241.044"/><path stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10.186 40.904l26.07.078a3.232 3.232 0 0 1-2.036 1.41c-1.252.156-24.068.035-24.068.035Z"/><path fill="none" stroke="#000" stroke-miterlimit="10" stroke-width="2" d="M11.797 32.146h12.968a1.694 1.694 0 0 1 1.207.502a14.389 14.389 0 0 1 2.554 3.097l.015.053a1.691 1.691 0 0 1-1.655 2.106l-15.067-.008a1.701 1.701 0 0 1-1.7-1.686l-.023-2.347a1.701 1.701 0 0 1 1.701-1.717Z"/><ellipse cx="32.8" cy="34.819" rx="3.217" ry="3.199"/>`),
		g.Group(children),
	)
}