export namespace main {
	
	export class DependencyStatus {
	    ddInstalled: boolean;
	    iso2godInstalled: boolean;
	    platform: string;
	
	    static createFrom(source: any = {}) {
	        return new DependencyStatus(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.ddInstalled = source["ddInstalled"];
	        this.iso2godInstalled = source["iso2godInstalled"];
	        this.platform = source["platform"];
	    }
	}
	export class DriveInfo {
	    name: string;
	    path: string;
	    size: string;
	
	    static createFrom(source: any = {}) {
	        return new DriveInfo(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.name = source["name"];
	        this.path = source["path"];
	        this.size = source["size"];
	    }
	}

}

