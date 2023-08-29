package noto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GemStone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="#E1F5FE" d="m4.01 47.94l17.48-26.51L35.03 36.9z"/><path fill="#81D4FA" d="M44.11 68.26L4.01 47.94L35.03 36.9z"/><path fill="#64B5F6" d="M63.94 43.06L35.03 36.9l9.08 31.36z"/><path fill="#0288D1" d="m123.87 47.94l-17.48-26.51L92.85 36.9z"/><path fill="#81D4FA" d="m83.77 68.26l40.1-20.32L92.85 36.9z"/><path fill="#E1F5FE" d="m63.94 43.06l28.91-6.16l-9.08 31.36z"/><path fill="#B2EBF2" d="m83.77 68.26l-19.83-25.2l-19.83 25.2z"/><path fill="#B3E5FC" d="M43 10.06h41.88l21.51 11.37L92.85 36.9l-28.91 6.16l-28.91-6.16l-13.54-15.47z"/><path fill="#1E88E5" d="M63.94 117.27L4.01 47.94l40.1 20.32z"/><path fill="#B3E5FC" d="m63.94 117.27l59.93-69.33l-40.1 20.32z"/><path fill="#E1F5FE" d="m83.77 68.26l-19.83 49.01l-19.83-49.01z"/>`),
		g.Group(children),
	)
}