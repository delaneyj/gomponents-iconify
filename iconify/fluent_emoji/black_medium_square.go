package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackMediumSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f137id1)"><path fill="url(#f137id0)" d="M5 6.571C5 5.704 5.704 5 6.571 5H25.43c.866 0 1.57.704 1.57 1.571V25.43c0 .867-.704 1.571-1.571 1.571H6.57A1.571 1.571 0 0 1 5 25.429V6.57Z"/></g><defs><linearGradient id="f137id0" x1="16" x2="16" y1="5" y2="27" gradientUnits="userSpaceOnUse"><stop stop-color="#4C405A"/><stop offset="1" stop-color="#473B54"/></linearGradient><filter id="f137id1" width="24" height="24" x="4" y="4" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.407843 0 0 0 0 0.384314 0 0 0 0 0.443137 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18590_2821"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dy="-1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.258824 0 0 0 0 0.141176 0 0 0 0 0.356863 0 0 0 1 0"/><feBlend in2="effect1_innerShadow_18590_2821" result="effect2_innerShadow_18590_2821"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.196078 0 0 0 0 0.176471 0 0 0 0 0.223529 0 0 0 1 0"/><feBlend in2="effect2_innerShadow_18590_2821" result="effect3_innerShadow_18590_2821"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dy="1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.278431 0 0 0 0 0.247059 0 0 0 0 0.317647 0 0 0 1 0"/><feBlend in2="effect3_innerShadow_18590_2821" result="effect4_innerShadow_18590_2821"/></filter></defs></g>`),
		g.Group(children),
	)
}