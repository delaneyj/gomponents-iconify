package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NameBadge(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f1423id2)"><path fill="url(#f1423id0)" d="M15.97 3.888c-.384 0-.609.242-.766.418l-3.49 3.978a1 1 0 0 1-1.489 0L7.252 4.973c-.37-.412-1.007-.447-1.39-.048a13.953 13.953 0 0 0-3.893 9.687c0 7.732 6.268 14 14 14s14-6.268 14-14c0-3.759-1.481-7.172-3.892-9.687c-.383-.4-1.02-.364-1.39.048l-2.974 3.31a1 1 0 0 1-1.488 0L16.74 4.307c-.153-.17-.36-.418-.77-.418Z"/></g><path fill="url(#f1423id1)" d="M5.97 15.612a1 1 0 0 1 1-1h18a1 1 0 0 1 1 1v2a1 1 0 0 1-1 1h-18a1 1 0 0 1-1-1v-2Z"/><defs><linearGradient id="f1423id0" x1="15.969" x2="15.969" y1="3.888" y2="28.612" gradientUnits="userSpaceOnUse"><stop stop-color="#FF4E92"/><stop offset="1" stop-color="#F54C5B"/></linearGradient><linearGradient id="f1423id1" x1="15.969" x2="15.969" y1="15.312" y2="18.612" gradientUnits="userSpaceOnUse"><stop stop-color="#E7D8FC"/><stop offset="1" stop-color="#F5EFFF"/></linearGradient><filter id="f1423id2" width="29" height="25.224" x="1.469" y="3.388" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx="-.75"/><feGaussianBlur stdDeviation=".25"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 1 0 0 0 0 0.466667 0 0 0 0 0.670588 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18590_1959"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dy="-.5"/><feGaussianBlur stdDeviation=".375"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.87451 0 0 0 0 0.14902 0 0 0 0 0.356863 0 0 0 1 0"/><feBlend in2="effect1_innerShadow_18590_1959" result="effect2_innerShadow_18590_1959"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx=".5"/><feGaussianBlur stdDeviation=".375"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.901961 0 0 0 0 0.164706 0 0 0 0 0.313726 0 0 0 1 0"/><feBlend in2="effect2_innerShadow_18590_1959" result="effect3_innerShadow_18590_1959"/></filter></defs></g>`),
		g.Group(children),
	)
}