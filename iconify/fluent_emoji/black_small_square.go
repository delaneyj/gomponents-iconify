package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackSmallSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f140id1)"><path fill="url(#f140id0)" d="M11.296 11.674c0-.395.32-.715.714-.715h8.571c.395 0 .715.32.715.715v8.57c0 .395-.32.715-.715.715H12.01a.714.714 0 0 1-.714-.714v-8.572Z"/></g><defs><linearGradient id="f140id0" x1="16.296" x2="16.296" y1="10.959" y2="20.959" gradientUnits="userSpaceOnUse"><stop stop-color="#4C405A"/><stop offset="1" stop-color="#473B54"/></linearGradient><filter id="f140id1" width="12" height="12" x="10.296" y="9.959" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.407843 0 0 0 0 0.384314 0 0 0 0 0.443137 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18590_2861"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dy="-1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.258824 0 0 0 0 0.141176 0 0 0 0 0.356863 0 0 0 1 0"/><feBlend in2="effect1_innerShadow_18590_2861" result="effect2_innerShadow_18590_2861"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.196078 0 0 0 0 0.176471 0 0 0 0 0.223529 0 0 0 1 0"/><feBlend in2="effect2_innerShadow_18590_2861" result="effect3_innerShadow_18590_2861"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dy="1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.278431 0 0 0 0 0.247059 0 0 0 0 0.317647 0 0 0 1 0"/><feBlend in2="effect3_innerShadow_18590_2861" result="effect4_innerShadow_18590_2861"/></filter></defs></g>`),
		g.Group(children),
	)
}