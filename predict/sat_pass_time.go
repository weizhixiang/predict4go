/**
  predict4java: An SDP4 / SGP4 library for satellite orbit predictions

  Copyright (C)  2004-2010  David A. B. Johnson, G4DPZ.

  This class is a Java port of one of the core elements of
  the Predict program, Copyright John A. Magliacane,
  KD2BD 1991-2003: http://www.qsl.net/kd2bd/predict.html

  Dr. T.S. Kelso is the author of the SGP4/SDP4 orbital models,
  originally written in Fortran and Pascal, and released into the
  public domain through his website (http://www.celestrak.com/).
  Neoklis Kyriazis, 5B4AZ, later re-wrote Dr. Kelso's code in C,
  and released it under the GNU GPL in 2002.
  PREDICT's core is based on 5B4AZ's code translation efforts.

  Author: David A. B. Johnson, G4DPZ <dave@g4dpz.me.uk>

  Comments, questions and bugreports should be submitted via
  http://sourceforge.net/projects/websat/
  More details can be found at the project home page:

  http://websat.sourceforge.net

  This program is free software; you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation; either version 2 of the License, or
  (at your option) any later version.

  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.

  You should have received a copy of the GNU General Public License
  along with this program; if not, visit http://www.fsf.org/
*/
package predict

import (
	"fmt"
	"strconv"
	"time"
)

type SatPassTime struct {
	StartTime  time.Time
	EndTime    time.Time
	Tca        time.Time
	PolePassed string
	Aos        int
	Los        int
	MaxEl      float64
}

func NewSatPassTime(startTime time.Time, endTime time.Time, tca time.Time, polePassed string, aos int, los int, maxEl float64) *SatPassTime {
	return &SatPassTime{
		StartTime:  startTime,
		EndTime:    endTime,
		Tca:        tca, //time.Unix((startTime.Unix() + endTime.Unix()) / 2, 0),
		PolePassed: polePassed,
		Aos:        aos,
		Los:        los,
		MaxEl:      maxEl,
	}
}

/**
 * Returns a string representing the contents of the object.
 */
func (s *SatPassTime) String() string {
	// Get duration in decimal minutes.
	duration := float64(s.EndTime.Sub(s.StartTime) / time.Second) / 60.0

	return "Date: " + s.StartTime.String() + "\n" +
		"Start Time: " + s.StartTime.String() + "\n" +
		fmt.Sprintf("Duration: %4.1f min.\n", duration) +
		"AOS Azimuth: " + strconv.Itoa(s.Aos) + DEG_CR +
		fmt.Sprintf("Max Elevation: %4.1f deg.\n", s.MaxEl) +
		"LOS Azimuth: " + strconv.Itoa(s.Los) + " deg."
}
