package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RedTrianglePointedDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f2093id1)"><path fill="url(#f2093id0)" d="m15.174 21.273l-5.61-9.805a1.25 1.25 0 0 1 1.085-1.87h11.375a1.25 1.25 0 0 1 1.078 1.883l-5.766 9.805a1.25 1.25 0 0 1-2.162-.013Z"/></g><defs><linearGradient id="f2093id0" x1="14.379" x2="14.379" y1="9.597" y2="20.725" gradientUnits="userSpaceOnUse"><stop stop-color="#DD3859"/><stop offset="1" stop-color="#D63983"/></linearGradient><filter id="f2093id1" width="15.88" height="13.806" x="8.397" y="8.597" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-1" dy=".5"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.964706 0 0 0 0 0.384314 0 0 0 0 0.54902 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18590_3123"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dy="-1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.768627 0 0 0 0 0.129412 0 0 0 0 0.560784 0 0 0 1 0"/><feBlend in2="effect1_innerShadow_18590_3123" result="effect2_innerShadow_18590_3123"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="1" dy=".5"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.815686 0 0 0 0 0.247059 0 0 0 0 0.376471 0 0 0 1 0"/><feBlend in2="effect2_innerShadow_18590_3123" result="effect3_innerShadow_18590_3123"/></filter></defs></g>`),
		g.Group(children),
	)
}