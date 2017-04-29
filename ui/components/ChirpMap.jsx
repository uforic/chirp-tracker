"use strict";

var React = require('react');
import { CircleMarker, Map, Marker, Popup, TileLayer } from 'react-leaflet';

const COORDS_MUHANGA = [-1.9209437,29.5748399];
class ChirpMap extends React.Component {
	constructor() {
		super();
		this.state = {
			zips: [],
		}
		this.update = this.update.bind(this);
		setInterval(this.update, 1000);
	}

	update(){
		var promise = fetch("influx", {credentials: "include"}).then(
			function(data) {
				return data.json();
			}
		).then(
			(function(data) {
				this.setState({
					zips: [
						{
							zipId: data.ZipID,
							lat: data.Lat,
							lon: data.Lon,
						}],
					});
			}).bind(this)
		);
	}

	render(){
		let position = null;
		if (this.state.zips.length > 0) {
			position = [this.state.zips[0].lat, this.state.zips[0].lon];
		}
		return <Map center={COORDS_MUHANGA} zoom={11}>
			<TileLayer
			url='http://{s}.tile.osm.org/{z}/{x}/{y}.png'
			/>
			{ position ?
				<CircleMarker 
					center={position}
					radius={7}
				/>
				: console.log("NO")
			}
		</Map>;
	}
}

module.exports = ChirpMap;
