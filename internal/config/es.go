/*
  Copyright (C) 2019 - 2022 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package config

// ES holds the configuration values for the Elasticsearch client.
type ES struct {
	Host    string `env:"ES_HOST" yaml:"host" default:"192.168.1.229"`
	Port    string `env:"ES_PORT" yaml:"port" default:"9200"`
	Cluster string `env:"ES_CLUSTER" yaml:"cluster" default:"superheromatch"`
	Index   string `env:"ES_INDEX" yaml:"index" default:"superhero"`
}
