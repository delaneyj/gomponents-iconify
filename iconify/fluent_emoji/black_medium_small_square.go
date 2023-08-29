package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BlackMediumSmallSquare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f138id1)"><path fill="url(#f138id0)" d="M8.407 9.078c0-.626.507-1.133 1.132-1.133h13.594c.626 0 1.133.507 1.133 1.133v13.594c0 .626-.507 1.133-1.133 1.133H9.54a1.133 1.133 0 0 1-1.132-1.133V9.078Z"/></g><defs><linearGradient id="f138id0" x1="16.336" x2="16.336" y1="7.945" y2="23.805" gradientUnits="userSpaceOnUse"><stop stop-color="#4C405A"/><stop offset="1" stop-color="#473B54"/></linearGradient><filter id="f138id1" width="17.86" height="17.86" x="7.407" y="6.945" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.407843 0 0 0 0 0.384314 0 0 0 0 0.443137 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18590_2841"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dy="-1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.258824 0 0 0 0 0.141176 0 0 0 0 0.356863 0 0 0 1 0"/><feBlend in2="effect1_innerShadow_18590_2841" result="effect2_innerShadow_18590_2841"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.196078 0 0 0 0 0.176471 0 0 0 0 0.223529 0 0 0 1 0"/><feBlend in2="effect2_innerShadow_18590_2841" result="effect3_innerShadow_18590_2841"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dy="1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.278431 0 0 0 0 0.247059 0 0 0 0 0.317647 0 0 0 1 0"/><feBlend in2="effect3_innerShadow_18590_2841" result="effect4_innerShadow_18590_2841"/></filter></defs></g>`),
		g.Group(children),
	)
}