package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StraightRuler(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f2304id1)"><rect width="31.576" height="10.482" x="8.542" y="1.201" fill="url(#f2304id0)" rx="2.25" transform="rotate(45 8.542 1.2)"/></g><path fill="#9D73E8" d="M11.117 3.776L9.172 5.72a.544.544 0 0 1-.77-.77l1.945-1.945l.77.77Zm1.491 1.491l.77.77l-1.11 1.11a.544.544 0 1 1-.77-.769l1.11-1.111Zm2.217 2.217l.77.77l-1.875 1.875a.545.545 0 0 1-.77-.77l1.875-1.875Zm2.281 2.28l.77.77l-1.15 1.15a.544.544 0 0 1-.77-.77l1.15-1.15Zm2.248 2.248l.77.77l-1.87 1.87a.544.544 0 1 1-.769-.77l1.87-1.87Zm2.268 2.268l.77.77l-1.143 1.143a.544.544 0 1 1-.77-.77l1.143-1.143Zm2.212 2.212l.77.77l-1.9 1.9a.544.544 0 0 1-.77-.77l1.9-1.9Zm2.253 2.253l.77.77l-1.162 1.162a.544.544 0 1 1-.77-.77l1.162-1.162Zm2.273 2.274l.77.77l-1.881 1.88a.544.544 0 0 1-.77-.77l1.881-1.88Z"/><defs><linearGradient id="f2304id0" x1="24.33" x2="24.33" y1="1.201" y2="11.682" gradientUnits="userSpaceOnUse"><stop stop-color="#CAC8D1"/><stop offset="1" stop-color="#BAB8C0"/></linearGradient><filter id="f2304id1" width="28.475" height="28.475" x="1.762" y="1.833" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx=".3" dy="-.3"/><feGaussianBlur stdDeviation=".3"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.611765 0 0 0 0 0.596078 0 0 0 0 0.658824 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18_23242"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-.3" dy=".3"/><feGaussianBlur stdDeviation=".3"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.913725 0 0 0 0 0.909804 0 0 0 0 0.92549 0 0 0 1 0"/><feBlend in2="effect1_innerShadow_18_23242" result="effect2_innerShadow_18_23242"/></filter></defs></g>`),
		g.Group(children),
	)
}