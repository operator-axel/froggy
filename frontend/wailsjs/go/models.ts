export namespace main {
	
	export class Task {
	    id: number;
	    title: string;
	    description: string;
	    tag: string;
	    completed: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Task(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.title = source["title"];
	        this.description = source["description"];
	        this.tag = source["tag"];
	        this.completed = source["completed"];
	    }
	}

}

