package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Husky(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 14.145a12.619 12.619 0 0 1 4.785.997l-.004-.004s.767-6.941 5.159-9.638c4.315 3.544 4.448 11.804 4.448 17.891l-.001-.001c1.684 2.489 2.68 4.944 2.68 6.589c0 6.973-3.197 9.516-3.197 12.521a17.068 17.068 0 0 1-.964-4.777c0-2.12 1.156-3.063 1.156-6.126c0-4.315-4.816-8.206-8.514-8.206c-4.045 0-5.548 4.777-5.548 9.477c0-4.7-1.502-9.477-5.548-9.477c-3.698 0-8.514 3.891-8.514 8.206c0 3.063 1.156 4.007 1.156 6.126a17.068 17.068 0 0 1-.963 4.777c0-3.005-3.198-5.548-3.198-12.521c0-1.645.996-4.1 2.68-6.59l-.001.002c0-6.087.133-14.347 4.448-17.891c4.392 2.697 5.16 9.638 5.16 9.638l-.005.004A12.619 12.619 0 0 1 24 14.145Z"/><circle cx="30.781" cy="31.713" r="1.695" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.687 20.042c0-2.15-.453-8.835-1.977-10.087c-1.552.953-2.412 6.475-2.412 6.475h0a12.986 12.986 0 0 1 4.389 3.612Z"/><circle cx="16.853" cy="32.13" r="1.787" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 37.267l-2.802 3.068h5.604L24 37.267zM11.682 19.828c0-2.266.476-9.312 2.083-10.632c1.636 1.004 2.543 6.825 2.543 6.825h0a13.688 13.688 0 0 0-4.626 3.807Z"/>`),
		g.Group(children),
	)
}