package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmallOrangeDiamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f2236id1)"><path fill="url(#f2236id0)" d="M10.1 17.357a.976.976 0 0 1 0-1.38l5.524-5.524a.976.976 0 0 1 1.38 0l5.524 5.523c.382.381.382 1 0 1.38l-5.523 5.524a.976.976 0 0 1-1.381 0L10.1 17.358Z"/></g><defs><linearGradient id="f2236id0" x1="13.648" x2="19.628" y1="14.133" y2="20.663" gradientUnits="userSpaceOnUse"><stop stop-color="#F47947"/><stop offset="1" stop-color="#EF5D50"/></linearGradient><filter id="f2236id1" width="14" height="14" x="9.314" y="9.666" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-.5"/><feGaussianBlur stdDeviation=".25"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.917647 0 0 0 0 0.305882 0 0 0 0 0.305882 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18590_2901"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx=".5"/><feGaussianBlur stdDeviation=".25"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.862745 0 0 0 0 0.396078 0 0 0 0 0.243137 0 0 0 1 0"/><feBlend in2="effect1_innerShadow_18590_2901" result="effect2_innerShadow_18590_2901"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx=".5" dy="-.5"/><feGaussianBlur stdDeviation=".375"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.843137 0 0 0 0 0.282353 0 0 0 0 0.360784 0 0 0 1 0"/><feBlend in2="effect2_innerShadow_18590_2901" result="effect3_innerShadow_18590_2901"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-.5" dy=".5"/><feGaussianBlur stdDeviation=".375"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 1 0 0 0 0 0.592157 0 0 0 0 0.356863 0 0 0 1 0"/><feBlend in2="effect3_innerShadow_18590_2901" result="effect4_innerShadow_18590_2901"/></filter></defs></g>`),
		g.Group(children),
	)
}