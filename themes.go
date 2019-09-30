package main

import "github.com/inkyblackness/imgui-go"

type theme interface {
	SetStyleColors(s imgui.Style)
	LatencyColors() [3]imgui.Vec4

	PushMainWindowStyle()
	PopMainWindowStyle()
}

type lightTheme struct {
}

func (l *lightTheme) SetStyleColors(s imgui.Style) {
	s.SetColor(imgui.StyleColorText, imgui.Vec4{0.00, 0.00, 0.00, 1.00})
	s.SetColor(imgui.StyleColorTextDisabled, imgui.Vec4{0.60, 0.60, 0.60, 1.00})
	s.SetColor(imgui.StyleColorWindowBg, imgui.Vec4{0.94, 0.94, 0.94, 0.94})
	s.SetColor(imgui.StyleColorChildBg, imgui.Vec4{0.00, 0.00, 0.00, 0.00})
	s.SetColor(imgui.StyleColorPopupBg, imgui.Vec4{1.00, 1.00, 1.00, 0.94})
	s.SetColor(imgui.StyleColorBorder, imgui.Vec4{0.00, 0.00, 0.00, 0.39})
	s.SetColor(imgui.StyleColorBorderShadow, imgui.Vec4{1.00, 1.00, 1.00, 0.10})
	s.SetColor(imgui.StyleColorFrameBg, imgui.Vec4{1.00, 1.00, 1.00, 0.94})
	s.SetColor(imgui.StyleColorFrameBgHovered, imgui.Vec4{0.26, 0.59, 0.98, 0.40})
	s.SetColor(imgui.StyleColorFrameBgActive, imgui.Vec4{0.26, 0.59, 0.98, 0.67})
	s.SetColor(imgui.StyleColorTitleBg, imgui.Vec4{0.96, 0.96, 0.96, 1.00})
	s.SetColor(imgui.StyleColorTitleBgCollapsed, imgui.Vec4{1.00, 1.00, 1.00, 0.51})
	s.SetColor(imgui.StyleColorTitleBgActive, imgui.Vec4{0.82, 0.82, 0.82, 1.00})
	s.SetColor(imgui.StyleColorMenuBarBg, imgui.Vec4{0.92, 0.92, 0.92, 1.00})
	s.SetColor(imgui.StyleColorScrollbarBg, imgui.Vec4{0.98, 0.98, 0.98, 0.53})
	s.SetColor(imgui.StyleColorScrollbarGrab, imgui.Vec4{0.69, 0.69, 0.69, 1.00})
	s.SetColor(imgui.StyleColorScrollbarGrabHovered, imgui.Vec4{0.59, 0.59, 0.59, 1.00})
	s.SetColor(imgui.StyleColorScrollbarGrabActive, imgui.Vec4{0.49, 0.49, 0.49, 1.00})
	// ComboBg?
	s.SetColor(imgui.StyleColorCheckMark, imgui.Vec4{0.26, 0.59, 0.98, 1.00})
	s.SetColor(imgui.StyleColorSliderGrab, imgui.Vec4{0.24, 0.52, 0.88, 1.00})
	s.SetColor(imgui.StyleColorSliderGrabActive, imgui.Vec4{0.26, 0.59, 0.98, 1.00})
	s.SetColor(imgui.StyleColorButton, imgui.Vec4{0.26, 0.59, 0.98, 0.40})
	s.SetColor(imgui.StyleColorButtonHovered, imgui.Vec4{0.26, 0.59, 0.98, 1.00})
	s.SetColor(imgui.StyleColorButtonActive, imgui.Vec4{0.06, 0.53, 0.98, 1.00})
	s.SetColor(imgui.StyleColorHeader, imgui.Vec4{0.26, 0.59, 0.98, 0.31})
	s.SetColor(imgui.StyleColorHeaderHovered, imgui.Vec4{0.26, 0.59, 0.98, 0.40})
	s.SetColor(imgui.StyleColorHeaderActive, imgui.Vec4{0.26, 0.59, 0.98, 0.80})
	// Column?
	s.SetColor(imgui.StyleColorResizeGrip, imgui.Vec4{1.00, 1.00, 1.00, 0.50})
	s.SetColor(imgui.StyleColorResizeGripHovered, imgui.Vec4{0.26, 0.59, 0.98, 0.67})
	s.SetColor(imgui.StyleColorResizeGripActive, imgui.Vec4{0.26, 0.59, 0.98, 0.95})
	// CloseButton?
	s.SetColor(imgui.StyleColorPlotLines, imgui.Vec4{0.39, 0.39, 0.39, 1.00})
	s.SetColor(imgui.StyleColorPlotLinesHovered, imgui.Vec4{1.00, 0.43, 0.35, 1.00})
	s.SetColor(imgui.StyleColorPlotHistogram, imgui.Vec4{0.90, 0.70, 0.00, 1.00})
	s.SetColor(imgui.StyleColorPlotHistogramHovered, imgui.Vec4{1.00, 0.60, 0.00, 1.00})
	s.SetColor(imgui.StyleColorTextSelectedBg, imgui.Vec4{0.26, 0.59, 0.98, 0.35})
	s.SetColor(imgui.StyleColorModalWindowDarkening, imgui.Vec4{0.20, 0.20, 0.20, 0.35})
}

func (l *lightTheme) LatencyColors() [3]imgui.Vec4 {
	return [3]imgui.Vec4{
		{0, 0.60, 0.10, 1},
		{0.60, 0.60, 0, 1},
		{0.8, 0, 0, 1},
	}
}

func (l *lightTheme) PushMainWindowStyle() {
	imgui.PushStyleVarFloat(imgui.StyleVarWindowBorderSize, 0)
	imgui.PushStyleVarFloat(imgui.StyleVarWindowRounding, 0)
	imgui.PushStyleColor(imgui.StyleColorWindowBg, imgui.Vec4{0.96, 0.96, 0.96, 1})
}

func (l *lightTheme) PopMainWindowStyle() {
	imgui.PopStyleColor()
	imgui.PopStyleVarV(2)
}
