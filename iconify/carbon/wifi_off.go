package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="16" cy="25" r="2" fill="currentColor"/><path fill="currentColor" d="M30 3.414L28.586 2L2 28.586L3.414 30l10.682-10.682a5.936 5.936 0 0 1 6.01 1.32l1.414-1.414a7.967 7.967 0 0 0-5.125-2.204l3.388-3.388a11.99 11.99 0 0 1 4.564 2.765l1.413-1.414a13.975 13.975 0 0 0-4.426-2.903l2.997-2.997a17.936 17.936 0 0 1 4.254 3.075L30 10.743v-.002a20.02 20.02 0 0 0-4.19-3.138zm-15.32 9.664l2.042-2.042C16.48 11.023 16.243 11 16 11a13.945 13.945 0 0 0-9.771 3.993l1.414 1.413a11.97 11.97 0 0 1 7.037-3.328zM16 7a17.87 17.87 0 0 1 4.232.525l1.643-1.642A19.954 19.954 0 0 0 2 10.74v.023l1.404 1.404A17.92 17.92 0 0 1 16 7z"/>`),
		g.Group(children),
	)
}