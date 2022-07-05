export namespace main {
	
	export class Config {
	    port: number;
	    lock2lock: number;
	    window_x: number;
	    window_y: number;
	    window_w: number;
	    window_h: number;
	
	    static createFrom(source: any = {}) {
	        return new Config(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.port = source["port"];
	        this.lock2lock = source["lock2lock"];
	        this.window_x = source["window_x"];
	        this.window_y = source["window_y"];
	        this.window_w = source["window_w"];
	        this.window_h = source["window_h"];
	    }
	}

}

