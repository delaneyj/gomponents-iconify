package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LargeBlueDiamond(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f816id1)"><path fill="url(#f816id0)" d="M3.608 18.353a2 2 0 0 1 0-2.829L14.922 4.211a2 2 0 0 1 2.829 0l11.313 11.314a2 2 0 0 1 0 2.828L17.751 29.667a2 2 0 0 1-2.829 0L3.608 18.353Z"/></g><defs><linearGradient id="f816id0" x1="10.875" x2="23.125" y1="11.75" y2="25.125" gradientUnits="userSpaceOnUse"><stop stop-color="#3D8BDD"/><stop offset="1" stop-color="#437FDB"/></linearGradient><filter id="f816id1" width="28.627" height="28.127" x="2.023" y="2.875" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.25098 0 0 0 0 0.462745 0 0 0 0 0.831373 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18590_2891"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="1"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.176471 0 0 0 0 0.521569 0 0 0 0 0.803922 0 0 0 1 0"/><feBlend in2="effect1_innerShadow_18590_2891" result="effect2_innerShadow_18590_2891"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx=".75" dy="-.75"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.231373 0 0 0 0 0.352941 0 0 0 0 0.784314 0 0 0 1 0"/><feBlend in2="effect2_innerShadow_18590_2891" result="effect3_innerShadow_18590_2891"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-.75" dy=".75"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.329412 0 0 0 0 0.698039 0 0 0 0 0.952941 0 0 0 1 0"/><feBlend in2="effect3_innerShadow_18590_2891" result="effect4_innerShadow_18590_2891"/></filter></defs></g>`),
		g.Group(children),
	)
}