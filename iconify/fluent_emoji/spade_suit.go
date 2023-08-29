package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpadeSuit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f2267id1)"><path fill="url(#f2267id0)" d="M28.43 14.965L18.712 3.35a3.791 3.791 0 0 0-5.8 0L3.193 14.965c-3.212 3.839-.463 9.657 4.562 9.657a5.961 5.961 0 0 0 5.408-3.459l.645-1.41v3.17c0 2.099-1.692 3.798-3.797 3.838h-.08c-.897 0-1.632.73-1.632 1.62c0 .89.735 1.619 1.632 1.619h11.772c.896 0 1.632-.73 1.632-1.62c0-.89-.736-1.619-1.632-1.619h-.08c-2.105-.04-3.797-1.74-3.797-3.838v-3.17l.645 1.41c.956 2.11 3.071 3.459 5.407 3.459c5.016 0 7.765-5.818 4.552-9.657Z"/></g><defs><linearGradient id="f2267id0" x1="24.129" x2="11.266" y1="11.078" y2="29.196" gradientUnits="userSpaceOnUse"><stop stop-color="#534165"/><stop offset="1" stop-color="#3F3946"/></linearGradient><filter id="f2267id1" width="29.75" height="29.5" x=".812" y="1.25" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.427451 0 0 0 0 0.372549 0 0 0 0 0.482353 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18_4023"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx=".75" dy="-.75"/><feGaussianBlur stdDeviation=".625"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.172549 0 0 0 0 0.109804 0 0 0 0 0.227451 0 0 0 1 0"/><feBlend in2="effect1_innerShadow_18_4023" result="effect2_innerShadow_18_4023"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx=".75" dy=".75"/><feGaussianBlur stdDeviation=".625"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.196078 0 0 0 0 0.192157 0 0 0 0 0.2 0 0 0 1 0"/><feBlend in2="effect2_innerShadow_18_4023" result="effect3_innerShadow_18_4023"/></filter></defs></g>`),
		g.Group(children),
	)
}