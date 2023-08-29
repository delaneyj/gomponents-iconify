package fluent_emoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HollowRedCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none"><g filter="url(#f692id3)"><circle cx="15.969" cy="15.75" r="12.438" stroke="url(#f692id0)" stroke-width="3"/></g><g filter="url(#f692id4)"><circle cx="16.305" cy="15.414" r="12.438" stroke="url(#f692id1)"/><circle cx="16.305" cy="15.414" r="12.438" stroke="url(#f692id2)"/></g><defs><linearGradient id="f692id0" x1="23.438" x2="6.688" y1="3.875" y2="22.875" gradientUnits="userSpaceOnUse"><stop stop-color="#FF5C96"/><stop offset=".486" stop-color="#DB3051"/><stop offset="1" stop-color="#EF3741"/></linearGradient><linearGradient id="f692id1" x1="26.598" x2="16.305" y1="5.422" y2="19.179" gradientUnits="userSpaceOnUse"><stop stop-color="#FF75A8"/><stop offset="1" stop-color="#FF75A8" stop-opacity="0"/></linearGradient><linearGradient id="f692id2" x1="7.388" x2="12.589" y1="23.29" y2="17.25" gradientUnits="userSpaceOnUse"><stop stop-color="#FF6D6F"/><stop offset="1" stop-color="#FF6D6F" stop-opacity="0"/></linearGradient><filter id="f692id3" width="28.375" height="28.375" x="2.032" y="1.313" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feColorMatrix in="SourceAlpha" result="hardAlpha" values="0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 127 0"/><feOffset dx=".5" dy="-.5"/><feGaussianBlur stdDeviation=".5"/><feComposite in2="hardAlpha" k2="-1" k3="1" operator="arithmetic"/><feColorMatrix values="0 0 0 0 0.921569 0 0 0 0 0.164706 0 0 0 0 0.309804 0 0 0 1 0"/><feBlend in2="shape" result="effect1_innerShadow_18590_1981"/></filter><filter id="f692id4" width="27.375" height="27.375" x="2.617" y="1.727" color-interpolation-filters="sRGB" filterUnits="userSpaceOnUse"><feFlood flood-opacity="0" result="BackgroundImageFix"/><feBlend in="SourceGraphic" in2="BackgroundImageFix" result="shape"/><feGaussianBlur result="effect1_foregroundBlur_18590_1981" stdDeviation=".375"/></filter></defs></g>`),
		g.Group(children),
	)
}